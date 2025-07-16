// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMeetingRoomsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *RemoveMeetingRoomsRequest
	GetCalendarId() *string
	SetEventId(v string) *RemoveMeetingRoomsRequest
	GetEventId() *string
	SetMeetingRoomsToRemove(v []*RemoveMeetingRoomsRequestMeetingRoomsToRemove) *RemoveMeetingRoomsRequest
	GetMeetingRoomsToRemove() []*RemoveMeetingRoomsRequestMeetingRoomsToRemove
}

type RemoveMeetingRoomsRequest struct {
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
	EventId              *string                                          `json:"EventId,omitempty" xml:"EventId,omitempty"`
	MeetingRoomsToRemove []*RemoveMeetingRoomsRequestMeetingRoomsToRemove `json:"MeetingRoomsToRemove,omitempty" xml:"MeetingRoomsToRemove,omitempty" type:"Repeated"`
}

func (s RemoveMeetingRoomsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveMeetingRoomsRequest) GoString() string {
	return s.String()
}

func (s *RemoveMeetingRoomsRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *RemoveMeetingRoomsRequest) GetEventId() *string {
	return s.EventId
}

func (s *RemoveMeetingRoomsRequest) GetMeetingRoomsToRemove() []*RemoveMeetingRoomsRequestMeetingRoomsToRemove {
	return s.MeetingRoomsToRemove
}

func (s *RemoveMeetingRoomsRequest) SetCalendarId(v string) *RemoveMeetingRoomsRequest {
	s.CalendarId = &v
	return s
}

func (s *RemoveMeetingRoomsRequest) SetEventId(v string) *RemoveMeetingRoomsRequest {
	s.EventId = &v
	return s
}

func (s *RemoveMeetingRoomsRequest) SetMeetingRoomsToRemove(v []*RemoveMeetingRoomsRequestMeetingRoomsToRemove) *RemoveMeetingRoomsRequest {
	s.MeetingRoomsToRemove = v
	return s
}

func (s *RemoveMeetingRoomsRequest) Validate() error {
	return dara.Validate(s)
}

type RemoveMeetingRoomsRequestMeetingRoomsToRemove struct {
	// This parameter is required.
	//
	// example:
	//
	// U5Kxxxxx
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s RemoveMeetingRoomsRequestMeetingRoomsToRemove) String() string {
	return dara.Prettify(s)
}

func (s RemoveMeetingRoomsRequestMeetingRoomsToRemove) GoString() string {
	return s.String()
}

func (s *RemoveMeetingRoomsRequestMeetingRoomsToRemove) GetRoomId() *string {
	return s.RoomId
}

func (s *RemoveMeetingRoomsRequestMeetingRoomsToRemove) SetRoomId(v string) *RemoveMeetingRoomsRequestMeetingRoomsToRemove {
	s.RoomId = &v
	return s
}

func (s *RemoveMeetingRoomsRequestMeetingRoomsToRemove) Validate() error {
	return dara.Validate(s)
}
