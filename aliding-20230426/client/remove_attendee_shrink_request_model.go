// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAttendeeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttendeesToRemoveShrink(v string) *RemoveAttendeeShrinkRequest
	GetAttendeesToRemoveShrink() *string
	SetCalendarId(v string) *RemoveAttendeeShrinkRequest
	GetCalendarId() *string
	SetEventId(v string) *RemoveAttendeeShrinkRequest
	GetEventId() *string
}

type RemoveAttendeeShrinkRequest struct {
	AttendeesToRemoveShrink *string `json:"AttendeesToRemove,omitempty" xml:"AttendeesToRemove,omitempty"`
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

func (s RemoveAttendeeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAttendeeShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeShrinkRequest) GetAttendeesToRemoveShrink() *string {
	return s.AttendeesToRemoveShrink
}

func (s *RemoveAttendeeShrinkRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *RemoveAttendeeShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *RemoveAttendeeShrinkRequest) SetAttendeesToRemoveShrink(v string) *RemoveAttendeeShrinkRequest {
	s.AttendeesToRemoveShrink = &v
	return s
}

func (s *RemoveAttendeeShrinkRequest) SetCalendarId(v string) *RemoveAttendeeShrinkRequest {
	s.CalendarId = &v
	return s
}

func (s *RemoveAttendeeShrinkRequest) SetEventId(v string) *RemoveAttendeeShrinkRequest {
	s.EventId = &v
	return s
}

func (s *RemoveAttendeeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
