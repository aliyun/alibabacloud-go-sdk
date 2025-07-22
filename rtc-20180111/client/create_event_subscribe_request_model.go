// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSubscribeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateEventSubscribeRequest
	GetAppId() *string
	SetCallbackUrl(v string) *CreateEventSubscribeRequest
	GetCallbackUrl() *string
	SetChannelId(v string) *CreateEventSubscribeRequest
	GetChannelId() *string
	SetClientToken(v string) *CreateEventSubscribeRequest
	GetClientToken() *string
	SetEvents(v []*string) *CreateEventSubscribeRequest
	GetEvents() []*string
	SetNeedCallbackAuth(v bool) *CreateEventSubscribeRequest
	GetNeedCallbackAuth() *bool
	SetOwnerId(v int64) *CreateEventSubscribeRequest
	GetOwnerId() *int64
	SetRole(v int64) *CreateEventSubscribeRequest
	GetRole() *int64
	SetUsers(v []*string) *CreateEventSubscribeRequest
	GetUsers() []*string
}

type CreateEventSubscribeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// 123333
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ChannelEvent
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// false
	NeedCallbackAuth *bool  `json:"NeedCallbackAuth,omitempty" xml:"NeedCallbackAuth,omitempty"`
	OwnerId          *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// user1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateEventSubscribeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSubscribeRequest) GoString() string {
	return s.String()
}

func (s *CreateEventSubscribeRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateEventSubscribeRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateEventSubscribeRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateEventSubscribeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEventSubscribeRequest) GetEvents() []*string {
	return s.Events
}

func (s *CreateEventSubscribeRequest) GetNeedCallbackAuth() *bool {
	return s.NeedCallbackAuth
}

func (s *CreateEventSubscribeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateEventSubscribeRequest) GetRole() *int64 {
	return s.Role
}

func (s *CreateEventSubscribeRequest) GetUsers() []*string {
	return s.Users
}

func (s *CreateEventSubscribeRequest) SetAppId(v string) *CreateEventSubscribeRequest {
	s.AppId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetCallbackUrl(v string) *CreateEventSubscribeRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetChannelId(v string) *CreateEventSubscribeRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetClientToken(v string) *CreateEventSubscribeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetEvents(v []*string) *CreateEventSubscribeRequest {
	s.Events = v
	return s
}

func (s *CreateEventSubscribeRequest) SetNeedCallbackAuth(v bool) *CreateEventSubscribeRequest {
	s.NeedCallbackAuth = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetOwnerId(v int64) *CreateEventSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetRole(v int64) *CreateEventSubscribeRequest {
	s.Role = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetUsers(v []*string) *CreateEventSubscribeRequest {
	s.Users = v
	return s
}

func (s *CreateEventSubscribeRequest) Validate() error {
	return dara.Validate(s)
}
