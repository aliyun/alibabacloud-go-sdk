// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateEventSubRequest
	GetAppId() *string
	SetCallbackUrl(v string) *CreateEventSubRequest
	GetCallbackUrl() *string
	SetChannelId(v string) *CreateEventSubRequest
	GetChannelId() *string
	SetEvents(v []*string) *CreateEventSubRequest
	GetEvents() []*string
	SetUsers(v []*string) *CreateEventSubRequest
	GetUsers() []*string
}

type CreateEventSubRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL. For more information about the content of the messages that are sent to the callback URL, see the Callback section in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The channel ID. You can call the [ListEventSub](https://help.aliyun.com/document_detail/2628135.html) operation to query the channel ID.
	//
	// >
	//
	// 	- This parameter is required if you specify the Users.N parameter.
	//
	// 	- If you set this parameter to \\	- or do not specify this parameter, all channels are subscribed to.
	//
	// 	- Each application ID allows only one all-channel subscription.
	//
	// example:
	//
	// 123333
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Subscribe to events.
	//
	// This parameter is required.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The user whose events you want to subscribe to. If you leave this parameter empty, the events of all users in the channel are subscribed to, including the events of the streamer and viewers. Specify this parameter in the following format:
	//
	//     Users.1=****
	//
	//     Users.2=****
	//
	//     ......
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateEventSubRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSubRequest) GoString() string {
	return s.String()
}

func (s *CreateEventSubRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateEventSubRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateEventSubRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateEventSubRequest) GetEvents() []*string {
	return s.Events
}

func (s *CreateEventSubRequest) GetUsers() []*string {
	return s.Users
}

func (s *CreateEventSubRequest) SetAppId(v string) *CreateEventSubRequest {
	s.AppId = &v
	return s
}

func (s *CreateEventSubRequest) SetCallbackUrl(v string) *CreateEventSubRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateEventSubRequest) SetChannelId(v string) *CreateEventSubRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateEventSubRequest) SetEvents(v []*string) *CreateEventSubRequest {
	s.Events = v
	return s
}

func (s *CreateEventSubRequest) SetUsers(v []*string) *CreateEventSubRequest {
	s.Users = v
	return s
}

func (s *CreateEventSubRequest) Validate() error {
	return dara.Validate(s)
}
