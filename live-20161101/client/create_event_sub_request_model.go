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
	// The ID of the application to subscribe to. You can view your application IDs by navigating to **ApsaraVideo Live > Live+ > ApsaraVideo Real-time Communication > Application Management**. If no application exists, create one by clicking [Create Application].
	//
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL. For the callback content, see the callback content examples below.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The ID of the channel to subscribe to. You can call the [ListEventSub](https://help.aliyun.com/document_detail/2848210.html) operation to query the subscribed channel IDs.
	//
	// >- If the Users.N parameter is not empty, this parameter is required.
	//
	// >- If ChannelId is set to \\	- or left empty, all channels are subscribed. Each AppId allows only one all-channel subscription.
	//
	// >- Each AppId allows a maximum of 20 subscriptions at the same time.
	//
	// example:
	//
	// 123333
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The subscription events.
	//
	// This parameter is required.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The users whose messages you want to subscribe to. If this parameter is empty, all users in the channel (including streamers and viewers) are subscribed. Format:
	//
	// ```
	//
	// Users.1=****
	//
	// Users.2=****
	//
	// ......
	//
	// ```
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
