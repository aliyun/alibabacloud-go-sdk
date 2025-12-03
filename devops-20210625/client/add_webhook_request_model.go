// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *AddWebhookRequest
	GetAccessToken() *string
	SetDescription(v string) *AddWebhookRequest
	GetDescription() *string
	SetEnableSslVerification(v bool) *AddWebhookRequest
	GetEnableSslVerification() *bool
	SetMergeRequestsEvents(v bool) *AddWebhookRequest
	GetMergeRequestsEvents() *bool
	SetNoteEvents(v bool) *AddWebhookRequest
	GetNoteEvents() *bool
	SetPushEvents(v bool) *AddWebhookRequest
	GetPushEvents() *bool
	SetSecretToken(v string) *AddWebhookRequest
	GetSecretToken() *string
	SetTagPushEvents(v bool) *AddWebhookRequest
	GetTagPushEvents() *bool
	SetUrl(v string) *AddWebhookRequest
	GetUrl() *string
	SetOrganizationId(v string) *AddWebhookRequest
	GetOrganizationId() *string
}

type AddWebhookRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	EnableSslVerification *bool `json:"enableSslVerification,omitempty" xml:"enableSslVerification,omitempty"`
	// example:
	//
	// true
	MergeRequestsEvents *bool `json:"mergeRequestsEvents,omitempty" xml:"mergeRequestsEvents,omitempty"`
	// example:
	//
	// false
	NoteEvents *bool `json:"noteEvents,omitempty" xml:"noteEvents,omitempty"`
	// example:
	//
	// true
	PushEvents *bool `json:"pushEvents,omitempty" xml:"pushEvents,omitempty"`
	// example:
	//
	// xxxx
	SecretToken *string `json:"secretToken,omitempty" xml:"secretToken,omitempty"`
	// example:
	//
	// false
	TagPushEvents *bool `json:"tagPushEvents,omitempty" xml:"tagPushEvents,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s AddWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWebhookRequest) GoString() string {
	return s.String()
}

func (s *AddWebhookRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *AddWebhookRequest) GetDescription() *string {
	return s.Description
}

func (s *AddWebhookRequest) GetEnableSslVerification() *bool {
	return s.EnableSslVerification
}

func (s *AddWebhookRequest) GetMergeRequestsEvents() *bool {
	return s.MergeRequestsEvents
}

func (s *AddWebhookRequest) GetNoteEvents() *bool {
	return s.NoteEvents
}

func (s *AddWebhookRequest) GetPushEvents() *bool {
	return s.PushEvents
}

func (s *AddWebhookRequest) GetSecretToken() *string {
	return s.SecretToken
}

func (s *AddWebhookRequest) GetTagPushEvents() *bool {
	return s.TagPushEvents
}

func (s *AddWebhookRequest) GetUrl() *string {
	return s.Url
}

func (s *AddWebhookRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AddWebhookRequest) SetAccessToken(v string) *AddWebhookRequest {
	s.AccessToken = &v
	return s
}

func (s *AddWebhookRequest) SetDescription(v string) *AddWebhookRequest {
	s.Description = &v
	return s
}

func (s *AddWebhookRequest) SetEnableSslVerification(v bool) *AddWebhookRequest {
	s.EnableSslVerification = &v
	return s
}

func (s *AddWebhookRequest) SetMergeRequestsEvents(v bool) *AddWebhookRequest {
	s.MergeRequestsEvents = &v
	return s
}

func (s *AddWebhookRequest) SetNoteEvents(v bool) *AddWebhookRequest {
	s.NoteEvents = &v
	return s
}

func (s *AddWebhookRequest) SetPushEvents(v bool) *AddWebhookRequest {
	s.PushEvents = &v
	return s
}

func (s *AddWebhookRequest) SetSecretToken(v string) *AddWebhookRequest {
	s.SecretToken = &v
	return s
}

func (s *AddWebhookRequest) SetTagPushEvents(v bool) *AddWebhookRequest {
	s.TagPushEvents = &v
	return s
}

func (s *AddWebhookRequest) SetUrl(v string) *AddWebhookRequest {
	s.Url = &v
	return s
}

func (s *AddWebhookRequest) SetOrganizationId(v string) *AddWebhookRequest {
	s.OrganizationId = &v
	return s
}

func (s *AddWebhookRequest) Validate() error {
	return dara.Validate(s)
}
