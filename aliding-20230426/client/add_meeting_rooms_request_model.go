// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMeetingRoomsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *AddMeetingRoomsRequest
	GetCalendarId() *string
	SetEventId(v string) *AddMeetingRoomsRequest
	GetEventId() *string
	SetMeetingRoomsToAdd(v []*AddMeetingRoomsRequestMeetingRoomsToAdd) *AddMeetingRoomsRequest
	GetMeetingRoomsToAdd() []*AddMeetingRoomsRequestMeetingRoomsToAdd
}

type AddMeetingRoomsRequest struct {
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
	MeetingRoomsToAdd []*AddMeetingRoomsRequestMeetingRoomsToAdd `json:"MeetingRoomsToAdd,omitempty" xml:"MeetingRoomsToAdd,omitempty" type:"Repeated"`
}

func (s AddMeetingRoomsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMeetingRoomsRequest) GoString() string {
	return s.String()
}

func (s *AddMeetingRoomsRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *AddMeetingRoomsRequest) GetEventId() *string {
	return s.EventId
}

func (s *AddMeetingRoomsRequest) GetMeetingRoomsToAdd() []*AddMeetingRoomsRequestMeetingRoomsToAdd {
	return s.MeetingRoomsToAdd
}

func (s *AddMeetingRoomsRequest) SetCalendarId(v string) *AddMeetingRoomsRequest {
	s.CalendarId = &v
	return s
}

func (s *AddMeetingRoomsRequest) SetEventId(v string) *AddMeetingRoomsRequest {
	s.EventId = &v
	return s
}

func (s *AddMeetingRoomsRequest) SetMeetingRoomsToAdd(v []*AddMeetingRoomsRequestMeetingRoomsToAdd) *AddMeetingRoomsRequest {
	s.MeetingRoomsToAdd = v
	return s
}

func (s *AddMeetingRoomsRequest) Validate() error {
	if s.MeetingRoomsToAdd != nil {
		for _, item := range s.MeetingRoomsToAdd {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddMeetingRoomsRequestMeetingRoomsToAdd struct {
	// example:
	//
	// 4002fxxxxx
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s AddMeetingRoomsRequestMeetingRoomsToAdd) String() string {
	return dara.Prettify(s)
}

func (s AddMeetingRoomsRequestMeetingRoomsToAdd) GoString() string {
	return s.String()
}

func (s *AddMeetingRoomsRequestMeetingRoomsToAdd) GetRoomId() *string {
	return s.RoomId
}

func (s *AddMeetingRoomsRequestMeetingRoomsToAdd) SetRoomId(v string) *AddMeetingRoomsRequestMeetingRoomsToAdd {
	s.RoomId = &v
	return s
}

func (s *AddMeetingRoomsRequestMeetingRoomsToAdd) Validate() error {
	return dara.Validate(s)
}
