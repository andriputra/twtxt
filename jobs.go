package twtxt

import (
	"fmt"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

var Jobs map[string]JobFactory

func init() {
	Jobs = map[string]JobFactory{
		"@every 5m":  NewUpdateFeedsJob,
		"@every 15m": NewUpdateFeedSourcesJob,
		"@hourly":    NewFixUserAccountsJob,
		"@daily":     NewStatsJob,
	}
}

type JobFactory func(conf *Config, store Store) cron.Job

type StatsJob struct {
	conf *Config
	db   Store
}

func NewStatsJob(conf *Config, db Store) cron.Job {
	return &StatsJob{conf: conf, db: db}
}

func (job *StatsJob) Run() {
	users, err := job.db.GetAllUsers()
	if err != nil {
		log.WithError(err).Warn("unable to get all users from database")
		return
	}

	log.Infof("updating stats")

	var feeds int
	for _, user := range users {
		feeds += len(user.Feeds)
	}

	tweets, err := GetAllTweets(job.conf)
	if err != nil {
		log.WithError(err).Warnf("error calculating number of tweets")
		return
	}

	text := fmt.Sprintf(
		"🧮  USERS:%d FEEDS:%d POSTS:%d",
		len(users), feeds, len(tweets),
	)

	if err := AppendSpecial(job.conf.Data, "stats", text); err != nil {
		log.WithError(err).Warn("error updating stats feed")
	}
}

type UpdateFeedsJob struct {
	conf *Config
	db   Store
}

func NewUpdateFeedsJob(conf *Config, db Store) cron.Job {
	return &UpdateFeedsJob{conf: conf, db: db}
}

func (job *UpdateFeedsJob) Run() {
	users, err := job.db.GetAllUsers()
	if err != nil {
		log.WithError(err).Warn("unable to get all users from database")
		return
	}

	log.Infof("updating feeds for %d users", len(users))

	sources := make(map[string]string)

	for _, user := range users {
		for u, n := range user.sources {
			sources[n] = u
		}
	}

	log.Infof("updating %d sources", len(sources))

	cache, err := LoadCache(job.conf.Data)
	if err != nil {
		log.WithError(err).Warn("error loading feed cache")
		return
	}

	cache.FetchTweets(job.conf, sources)

	if err := cache.Store(job.conf.Data); err != nil {
		log.WithError(err).Warn("error saving feed cache")
	} else {
		log.Info("updated feed cache")
	}
}

type UpdateFeedSourcesJob struct {
	conf *Config
	db   Store
}

func NewUpdateFeedSourcesJob(conf *Config, db Store) cron.Job {
	return &UpdateFeedSourcesJob{conf: conf, db: db}
}

func (job *UpdateFeedSourcesJob) Run() {
	log.Infof("updating %d feed sources", len(job.conf.FeedSources))

	feeds := FetchFeeds(job.conf.FeedSources)

	log.Infof("fetched %d feeds", len(feeds))

	if err := SaveFeeds(feeds, job.conf.Data); err != nil {
		log.WithError(err).Warn("error saving feeds")
	} else {
		log.Info("updated feeds")
	}
}

type FixUserAccountsJob struct {
	conf *Config
	db   Store
}

func NewFixUserAccountsJob(conf *Config, db Store) cron.Job {
	return &FixUserAccountsJob{conf: conf, db: db}
}

func (job *FixUserAccountsJob) Run() {
	fixMissingUserFeeds := func(username string, feeds []string) error {
		user, err := job.db.GetUser(username)
		if err != nil {
			log.WithError(err).Errorf("error loading user object for %s", username)
			return err
		}

		user.Feeds = feeds

		if err := job.db.SetUser(username, user); err != nil {
			log.WithError(err).Errorf("error updating user object %s", username)
			return err
		}

		log.Infof("fixed missing feeds for %s", username)

		return nil
	}

	// Fix missing Feeds for @rob @kt84
	if err := fixMissingUserFeeds("kt84", []string{"recipes", "local_wonders"}); err != nil {
		log.WithError(err).Errorf("error fixing missing user feeds")
	}
	if err := fixMissingUserFeeds("rob", []string{"off_grid_living"}); err != nil {
		log.WithError(err).Errorf("error fixing missing user feeds")
	}
}
