// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *DeleteEventRequest
	GetCalendarId() *string
	SetEventId(v string) *DeleteEventRequest
	GetEventId() *string
	SetPushNotification(v bool) *DeleteEventRequest
	GetPushNotification() *bool
}

type DeleteEventRequest struct {
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
	// iiiP35sJadba8aBSgjrwPRKgiEiF
	EventId          *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	PushNotification *bool   `json:"pushNotification,omitempty" xml:"pushNotification,omitempty"`
}

func (s DeleteEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *DeleteEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *DeleteEventRequest) GetPushNotification() *bool {
	return s.PushNotification
}

func (s *DeleteEventRequest) SetCalendarId(v string) *DeleteEventRequest {
	s.CalendarId = &v
	return s
}

func (s *DeleteEventRequest) SetEventId(v string) *DeleteEventRequest {
	s.EventId = &v
	return s
}

func (s *DeleteEventRequest) SetPushNotification(v bool) *DeleteEventRequest {
	s.PushNotification = &v
	return s
}

func (s *DeleteEventRequest) Validate() error {
	return dara.Validate(s)
}
