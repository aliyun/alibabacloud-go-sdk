// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeetingRoomsScheduleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetMeetingRoomsScheduleShrinkRequest
	GetEndTime() *string
	SetRoomIdsShrink(v string) *GetMeetingRoomsScheduleShrinkRequest
	GetRoomIdsShrink() *string
	SetStartTime(v string) *GetMeetingRoomsScheduleShrinkRequest
	GetStartTime() *string
}

type GetMeetingRoomsScheduleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["4002xxxxx"]
	RoomIdsShrink *string `json:"RoomIds,omitempty" xml:"RoomIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetMeetingRoomsScheduleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMeetingRoomsScheduleShrinkRequest) GetRoomIdsShrink() *string {
	return s.RoomIdsShrink
}

func (s *GetMeetingRoomsScheduleShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMeetingRoomsScheduleShrinkRequest) SetEndTime(v string) *GetMeetingRoomsScheduleShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetMeetingRoomsScheduleShrinkRequest) SetRoomIdsShrink(v string) *GetMeetingRoomsScheduleShrinkRequest {
	s.RoomIdsShrink = &v
	return s
}

func (s *GetMeetingRoomsScheduleShrinkRequest) SetStartTime(v string) *GetMeetingRoomsScheduleShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetMeetingRoomsScheduleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
