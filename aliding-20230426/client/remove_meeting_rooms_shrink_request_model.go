// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMeetingRoomsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *RemoveMeetingRoomsShrinkRequest
	GetCalendarId() *string
	SetEventId(v string) *RemoveMeetingRoomsShrinkRequest
	GetEventId() *string
	SetMeetingRoomsToRemoveShrink(v string) *RemoveMeetingRoomsShrinkRequest
	GetMeetingRoomsToRemoveShrink() *string
}

type RemoveMeetingRoomsShrinkRequest struct {
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
	// U5Kxxxxx
	EventId                    *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	MeetingRoomsToRemoveShrink *string `json:"MeetingRoomsToRemove,omitempty" xml:"MeetingRoomsToRemove,omitempty"`
}

func (s RemoveMeetingRoomsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveMeetingRoomsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveMeetingRoomsShrinkRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *RemoveMeetingRoomsShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *RemoveMeetingRoomsShrinkRequest) GetMeetingRoomsToRemoveShrink() *string {
	return s.MeetingRoomsToRemoveShrink
}

func (s *RemoveMeetingRoomsShrinkRequest) SetCalendarId(v string) *RemoveMeetingRoomsShrinkRequest {
	s.CalendarId = &v
	return s
}

func (s *RemoveMeetingRoomsShrinkRequest) SetEventId(v string) *RemoveMeetingRoomsShrinkRequest {
	s.EventId = &v
	return s
}

func (s *RemoveMeetingRoomsShrinkRequest) SetMeetingRoomsToRemoveShrink(v string) *RemoveMeetingRoomsShrinkRequest {
	s.MeetingRoomsToRemoveShrink = &v
	return s
}

func (s *RemoveMeetingRoomsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
