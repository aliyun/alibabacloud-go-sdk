// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAttendeeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttendeesToAdd(v []*AddAttendeeRequestAttendeesToAdd) *AddAttendeeRequest
	GetAttendeesToAdd() []*AddAttendeeRequestAttendeesToAdd
	SetCalendarId(v string) *AddAttendeeRequest
	GetCalendarId() *string
	SetEventId(v string) *AddAttendeeRequest
	GetEventId() *string
	SetChatNotification(v bool) *AddAttendeeRequest
	GetChatNotification() *bool
	SetPushNotification(v bool) *AddAttendeeRequest
	GetPushNotification() *bool
}

type AddAttendeeRequest struct {
	// This parameter is required.
	AttendeesToAdd []*AddAttendeeRequestAttendeesToAdd `json:"AttendeesToAdd,omitempty" xml:"AttendeesToAdd,omitempty" type:"Repeated"`
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

func (s AddAttendeeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeRequest) GoString() string {
	return s.String()
}

func (s *AddAttendeeRequest) GetAttendeesToAdd() []*AddAttendeeRequestAttendeesToAdd {
	return s.AttendeesToAdd
}

func (s *AddAttendeeRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *AddAttendeeRequest) GetEventId() *string {
	return s.EventId
}

func (s *AddAttendeeRequest) GetChatNotification() *bool {
	return s.ChatNotification
}

func (s *AddAttendeeRequest) GetPushNotification() *bool {
	return s.PushNotification
}

func (s *AddAttendeeRequest) SetAttendeesToAdd(v []*AddAttendeeRequestAttendeesToAdd) *AddAttendeeRequest {
	s.AttendeesToAdd = v
	return s
}

func (s *AddAttendeeRequest) SetCalendarId(v string) *AddAttendeeRequest {
	s.CalendarId = &v
	return s
}

func (s *AddAttendeeRequest) SetEventId(v string) *AddAttendeeRequest {
	s.EventId = &v
	return s
}

func (s *AddAttendeeRequest) SetChatNotification(v bool) *AddAttendeeRequest {
	s.ChatNotification = &v
	return s
}

func (s *AddAttendeeRequest) SetPushNotification(v bool) *AddAttendeeRequest {
	s.PushNotification = &v
	return s
}

func (s *AddAttendeeRequest) Validate() error {
	return dara.Validate(s)
}

type AddAttendeeRequestAttendeesToAdd struct {
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsOptional *bool `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
}

func (s AddAttendeeRequestAttendeesToAdd) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeRequestAttendeesToAdd) GoString() string {
	return s.String()
}

func (s *AddAttendeeRequestAttendeesToAdd) GetId() *string {
	return s.Id
}

func (s *AddAttendeeRequestAttendeesToAdd) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *AddAttendeeRequestAttendeesToAdd) SetId(v string) *AddAttendeeRequestAttendeesToAdd {
	s.Id = &v
	return s
}

func (s *AddAttendeeRequestAttendeesToAdd) SetIsOptional(v bool) *AddAttendeeRequestAttendeesToAdd {
	s.IsOptional = &v
	return s
}

func (s *AddAttendeeRequestAttendeesToAdd) Validate() error {
	return dara.Validate(s)
}
