// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAttendeeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttendeesToRemove(v []*string) *RemoveAttendeeRequest
	GetAttendeesToRemove() []*string
	SetCalendarId(v string) *RemoveAttendeeRequest
	GetCalendarId() *string
	SetEventId(v string) *RemoveAttendeeRequest
	GetEventId() *string
}

type RemoveAttendeeRequest struct {
	AttendeesToRemove []*string `json:"AttendeesToRemove,omitempty" xml:"AttendeesToRemove,omitempty" type:"Repeated"`
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
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s RemoveAttendeeRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAttendeeRequest) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeRequest) GetAttendeesToRemove() []*string {
	return s.AttendeesToRemove
}

func (s *RemoveAttendeeRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *RemoveAttendeeRequest) GetEventId() *string {
	return s.EventId
}

func (s *RemoveAttendeeRequest) SetAttendeesToRemove(v []*string) *RemoveAttendeeRequest {
	s.AttendeesToRemove = v
	return s
}

func (s *RemoveAttendeeRequest) SetCalendarId(v string) *RemoveAttendeeRequest {
	s.CalendarId = &v
	return s
}

func (s *RemoveAttendeeRequest) SetEventId(v string) *RemoveAttendeeRequest {
	s.EventId = &v
	return s
}

func (s *RemoveAttendeeRequest) Validate() error {
	return dara.Validate(s)
}
