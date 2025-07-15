// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateEventSubRequest
	GetAppId() *string
	SetCallbackUrl(v string) *UpdateEventSubRequest
	GetCallbackUrl() *string
	SetChannelId(v string) *UpdateEventSubRequest
	GetChannelId() *string
	SetEvents(v []*string) *UpdateEventSubRequest
	GetEvents() []*string
	SetSubscribeId(v string) *UpdateEventSubRequest
	GetSubscribeId() *string
	SetUsers(v []*string) *UpdateEventSubRequest
	GetUsers() []*string
}

type UpdateEventSubRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL. For more information about the callback content, see CreateEventSub.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The channel ID. You can call the [ListEventSub](https://help.aliyun.com/document_detail/2848210.html) operation to query the channel ID.
	//
	// >
	//
	// 	- This parameter is required if you specify the Users.N parameter.
	//
	// 	- If you set this parameter to \\	- or do not specify this parameter, all channels are subscribed to.
	//
	// 	- You can create up to 20 subscriptions for each application ID.
	//
	// example:
	//
	// 123333
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The type of the events to which you want to subscribe.
	//
	// This parameter is required.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The subscription ID. You can obtain the ID from the response to the [CreateEventSub](https://help.aliyun.com/document_detail/2848209.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ad53276431c****
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
	// The user whose events you want to subscribe to.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s UpdateEventSubRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSubRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventSubRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateEventSubRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *UpdateEventSubRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateEventSubRequest) GetEvents() []*string {
	return s.Events
}

func (s *UpdateEventSubRequest) GetSubscribeId() *string {
	return s.SubscribeId
}

func (s *UpdateEventSubRequest) GetUsers() []*string {
	return s.Users
}

func (s *UpdateEventSubRequest) SetAppId(v string) *UpdateEventSubRequest {
	s.AppId = &v
	return s
}

func (s *UpdateEventSubRequest) SetCallbackUrl(v string) *UpdateEventSubRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UpdateEventSubRequest) SetChannelId(v string) *UpdateEventSubRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateEventSubRequest) SetEvents(v []*string) *UpdateEventSubRequest {
	s.Events = v
	return s
}

func (s *UpdateEventSubRequest) SetSubscribeId(v string) *UpdateEventSubRequest {
	s.SubscribeId = &v
	return s
}

func (s *UpdateEventSubRequest) SetUsers(v []*string) *UpdateEventSubRequest {
	s.Users = v
	return s
}

func (s *UpdateEventSubRequest) Validate() error {
	return dara.Validate(s)
}
