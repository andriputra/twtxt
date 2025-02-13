package types

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// AuthRequest ...
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewAuthRequest ...
func NewAuthRequest(r io.Reader) (req AuthRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// AuthResponse ...
type AuthResponse struct {
	Token string `json:"token"`
}

// Bytes ...
func (res AuthResponse) Bytes() ([]byte, error) {
	body, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// RegisterRequest ...
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// NewRegisterRequest ...
func NewRegisterRequest(r io.Reader) (req RegisterRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// PostRequest ...
type PostRequest struct {
	PostAs string `json:"post_as"`
	Text   string `json:"text"`
}

// NewPostRequest ...
func NewPostRequest(r io.Reader) (req PostRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// PagedRequest ...
type PagedRequest struct {
	Page int `json:"page"`
}

// NewPagedRequest ...
func NewPagedRequest(r io.Reader) (req PagedRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// PagerResponse ...
type PagerResponse struct {
	Current   int `json:"current_page"`
	MaxPages  int `json:"max_pages"`
	TotalTwts int `json:"total_twts"`
}

// PagedResponse ...
type PagedResponse struct {
	Twts  []Twt `json:"twts"`
	Pager PagerResponse
}

// Bytes ...
func (res PagedResponse) Bytes() ([]byte, error) {
	body, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// FollowRequest ...
type FollowRequest struct {
	Nick string `json:"nick"`
	URL  string `json:"url"`
}

// NewFollowRequest ...
func NewFollowRequest(r io.Reader) (req FollowRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// UnfollowRequest ...
type UnfollowRequest struct {
	Nick string `json:"nick"`
}

// NewUnfollowRequest ...
func NewUnfollowRequest(r io.Reader) (req UnfollowRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// ProfileResponse ...
type ProfileResponse struct {
	Profile      Profile      `json:"profile"`
	Links        Links        `json:"links"`
	Alternatives Alternatives `json:"alternatives"`
	Twter        Twter        `json:"twter"`
}

// ConversationRequest ...
type ConversationRequest struct {
	Hash string `json:"hash"`
	Page int    `json:"page"`
}

// NewConversationRequest ...
func NewConversationRequest(r io.Reader) (req ConversationRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// FetchTwtsRequest ...
type FetchTwtsRequest struct {
	URL  string `json:"url"`
	Nick string `json:"nick"`
	Page int    `json:"page"`
}

// NewFetchTwtsRequest ...
func NewFetchTwtsRequest(r io.Reader) (req FetchTwtsRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// MuteRequest ...
type MuteRequest struct {
	Nick string `json:"nick"`
	URL  string `json:"url"`
}

// NewMuteRequest ...
func NewMuteRequest(r io.Reader) (req MuteRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// UnmuteRequest ...
type UnmuteRequest struct {
	Nick string `json:"nick"`
}

// NewUnmuteRequest ...
func NewUnmuteRequest(r io.Reader) (req UnmuteRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// ExternalProfileRequest ...
type ExternalProfileRequest struct {
	URL  string `json:"url"`
	Nick string `json:"nick"`
}

// NewExternalProfileRequest ...
func NewExternalProfileRequest(r io.Reader) (req ExternalProfileRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// SupportRequest ...
type SupportRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// NewSupportRequest ...
func NewSupportRequest(r io.Reader) (req SupportRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}

// ReportRequest ...
type ReportRequest struct {
	Nick string `json:"nick"`
	URL  string `json:"url"`

	Name     string `json:"name"`
	Email    string `json:"email"`
	Category string `json:"subject"`
	Message  string `json:"message"`
}

// NewReportRequest ...
func NewReportRequest(r io.Reader) (req ReportRequest, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	return
}
