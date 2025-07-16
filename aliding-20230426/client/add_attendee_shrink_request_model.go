// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAttendeeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttendeesToAddShrink(v string) *AddAttendeeShrinkRequest
	GetAttendeesToAddShrink() *string
	SetCalendarId(v string) *AddAttendeeShrinkRequest
	GetCalendarId() *string
	SetEventId(v string) *AddAttendeeShrinkRequest
	GetEventId() *string
	SetChatNotification(v bool) *AddAttendeeShrinkRequest
	GetChatNotification() *bool
	SetPushNotification(v bool) *AddAttendeeShrinkRequest
	GetPushNotification() *bool
}

type AddAttendeeShrinkRequest struct {
	// This parameter is required.
	AttendeesToAddShrink *string `json:"AttendeesToAdd,omitempty" xml:"AttendeesToAdd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// primary
	CalendarId *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cnNTbW1YbU9sL2p6aFJZdEgvdlQrQT08
	EventId          *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	ChatNotification *bool   `json:"chatNotification,omitempty" xml:"chatNotification,omitempty"`
	PushNotification *bool   `json:"pushNotification,omitempty" xml:"pushNotification,omitempty"`
}

func (s AddAttendeeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAttendeeShrinkRequest) GetAttendeesToAddShrink() *string {
	return s.AttendeesToAddShrink
}

func (s *AddAttendeeShrinkRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *AddAttendeeShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *AddAttendeeShrinkRequest) GetChatNotification() *bool {
	return s.ChatNotification
}

func (s *AddAttendeeShrinkRequest) GetPushNotification() *bool {
	return s.PushNotification
}

func (s *AddAttendeeShrinkRequest) SetAttendeesToAddShrink(v string) *AddAttendeeShrinkRequest {
	s.AttendeesToAddShrink = &v
	return s
}

func (s *AddAttendeeShrinkRequest) SetCalendarId(v string) *AddAttendeeShrinkRequest {
	s.CalendarId = &v
	return s
}

func (s *AddAttendeeShrinkRequest) SetEventId(v string) *AddAttendeeShrinkRequest {
	s.EventId = &v
	return s
}

func (s *AddAttendeeShrinkRequest) SetChatNotification(v bool) *AddAttendeeShrinkRequest {
	s.ChatNotification = &v
	return s
}

func (s *AddAttendeeShrinkRequest) SetPushNotification(v bool) *AddAttendeeShrinkRequest {
	s.PushNotification = &v
	return s
}

func (s *AddAttendeeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
