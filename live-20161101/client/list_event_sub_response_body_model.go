// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListEventSubResponseBody
	GetRequestId() *string
	SetSubscribers(v []*ListEventSubResponseBodySubscribers) *ListEventSubResponseBody
	GetSubscribers() []*ListEventSubResponseBodySubscribers
}

type ListEventSubResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AE050E24-BE9B-1E79-BB30-7EA0BBAE7F08
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The events.
	Subscribers []*ListEventSubResponseBodySubscribers `json:"Subscribers,omitempty" xml:"Subscribers,omitempty" type:"Repeated"`
}

func (s ListEventSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventSubResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventSubResponseBody) GetSubscribers() []*ListEventSubResponseBodySubscribers {
	return s.Subscribers
}

func (s *ListEventSubResponseBody) SetRequestId(v string) *ListEventSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventSubResponseBody) SetSubscribers(v []*ListEventSubResponseBodySubscribers) *ListEventSubResponseBody {
	s.Subscribers = v
	return s
}

func (s *ListEventSubResponseBody) Validate() error {
	if s.Subscribers != nil {
		for _, item := range s.Subscribers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEventSubResponseBodySubscribers struct {
	// The callback URL.
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The ID of the channel to which you subscribe.
	//
	// example:
	//
	// testmodify
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The time when the subscription was created. The time is displayed in UTC+8. Format: yyyy-MM-dd hh:mm:ss.
	//
	// example:
	//
	// 2023-08-18 10:14:49
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- ChannelEvent: channel event. For example, the channel is opened or closed.
	//
	// 	- UserEvent: user event. For example, a user joins or leaves the channel.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The time when the subscription was modified. The time is displayed in UTC+8. Format: yyyy-MM-dd hh:mm:ss.
	//
	// example:
	//
	// 2023-08-18 10:14:49
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// Deprecated
	//
	// The role of the user whose events are returned. Valid values:
	//
	// 	- 1: streamer
	//
	// 	- 2: viewer
	//
	// An empty value or a value other than 1 and 2 indicates that events of all users in the channel are returned.
	//
	// >  This parameter is deprecated. When you create a subscription, this parameter is no longer supported.
	//
	// example:
	//
	// 1
	Roles *int32 `json:"Roles,omitempty" xml:"Roles,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 09be0d2254cb5a89f4cbd86403ec****
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The user whose events are returned. We recommend that you use either this parameter or Roles.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListEventSubResponseBodySubscribers) String() string {
	return dara.Prettify(s)
}

func (s ListEventSubResponseBodySubscribers) GoString() string {
	return s.String()
}

func (s *ListEventSubResponseBodySubscribers) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *ListEventSubResponseBodySubscribers) GetChannelId() *string {
	return s.ChannelId
}

func (s *ListEventSubResponseBodySubscribers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEventSubResponseBodySubscribers) GetEvents() []*string {
	return s.Events
}

func (s *ListEventSubResponseBodySubscribers) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListEventSubResponseBodySubscribers) GetRoles() *int32 {
	return s.Roles
}

func (s *ListEventSubResponseBodySubscribers) GetSubId() *string {
	return s.SubId
}

func (s *ListEventSubResponseBodySubscribers) GetUsers() []*string {
	return s.Users
}

func (s *ListEventSubResponseBodySubscribers) SetCallbackUrl(v string) *ListEventSubResponseBodySubscribers {
	s.CallbackUrl = &v
	return s
}

func (s *ListEventSubResponseBodySubscribers) SetChannelId(v string) *ListEventSubResponseBodySubscribers {
	s.ChannelId = &v
	return s
}

func (s *ListEventSubResponseBodySubscribers) SetCreateTime(v string) *ListEventSubResponseBodySubscribers {
	s.CreateTime = &v
	return s
}

func (s *ListEventSubResponseBodySubscribers) SetEvents(v []*string) *ListEventSubResponseBodySubscribers {
	s.Events = v
	return s
}

func (s *ListEventSubResponseBodySubscribers) SetModifyTime(v string) *ListEventSubResponseBodySubscribers {
	s.ModifyTime = &v
	return s
}

func (s *ListEventSubResponseBodySubscribers) SetRoles(v int32) *ListEventSubResponseBodySubscribers {
	s.Roles = &v
	return s
}

func (s *ListEventSubResponseBodySubscribers) SetSubId(v string) *ListEventSubResponseBodySubscribers {
	s.SubId = &v
	return s
}

func (s *ListEventSubResponseBodySubscribers) SetUsers(v []*string) *ListEventSubResponseBodySubscribers {
	s.Users = v
	return s
}

func (s *ListEventSubResponseBodySubscribers) Validate() error {
	return dara.Validate(s)
}
