// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *GetEventRequest
	GetCalendarId() *string
	SetEventId(v string) *GetEventRequest
	GetEventId() *string
	SetMaxAttendees(v int64) *GetEventRequest
	GetMaxAttendees() *int64
}

type GetEventRequest struct {
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
	// 311525211
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// 100
	MaxAttendees *int64 `json:"MaxAttendees,omitempty" xml:"MaxAttendees,omitempty"`
}

func (s GetEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEventRequest) GoString() string {
	return s.String()
}

func (s *GetEventRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *GetEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *GetEventRequest) GetMaxAttendees() *int64 {
	return s.MaxAttendees
}

func (s *GetEventRequest) SetCalendarId(v string) *GetEventRequest {
	s.CalendarId = &v
	return s
}

func (s *GetEventRequest) SetEventId(v string) *GetEventRequest {
	s.EventId = &v
	return s
}

func (s *GetEventRequest) SetMaxAttendees(v int64) *GetEventRequest {
	s.MaxAttendees = &v
	return s
}

func (s *GetEventRequest) Validate() error {
	return dara.Validate(s)
}
