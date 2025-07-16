// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMeetingRoomsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *AddMeetingRoomsShrinkRequest
	GetCalendarId() *string
	SetEventId(v string) *AddMeetingRoomsShrinkRequest
	GetEventId() *string
	SetMeetingRoomsToAddShrink(v string) *AddMeetingRoomsShrinkRequest
	GetMeetingRoomsToAddShrink() *string
}

type AddMeetingRoomsShrinkRequest struct {
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
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// This parameter is required.
	MeetingRoomsToAddShrink *string `json:"MeetingRoomsToAdd,omitempty" xml:"MeetingRoomsToAdd,omitempty"`
}

func (s AddMeetingRoomsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMeetingRoomsShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddMeetingRoomsShrinkRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *AddMeetingRoomsShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *AddMeetingRoomsShrinkRequest) GetMeetingRoomsToAddShrink() *string {
	return s.MeetingRoomsToAddShrink
}

func (s *AddMeetingRoomsShrinkRequest) SetCalendarId(v string) *AddMeetingRoomsShrinkRequest {
	s.CalendarId = &v
	return s
}

func (s *AddMeetingRoomsShrinkRequest) SetEventId(v string) *AddMeetingRoomsShrinkRequest {
	s.EventId = &v
	return s
}

func (s *AddMeetingRoomsShrinkRequest) SetMeetingRoomsToAddShrink(v string) *AddMeetingRoomsShrinkRequest {
	s.MeetingRoomsToAddShrink = &v
	return s
}

func (s *AddMeetingRoomsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
