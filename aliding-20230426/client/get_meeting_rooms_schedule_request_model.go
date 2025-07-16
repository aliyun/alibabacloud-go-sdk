// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeetingRoomsScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetMeetingRoomsScheduleRequest
	GetEndTime() *string
	SetRoomIds(v []*string) *GetMeetingRoomsScheduleRequest
	GetRoomIds() []*string
	SetStartTime(v string) *GetMeetingRoomsScheduleRequest
	GetStartTime() *string
}

type GetMeetingRoomsScheduleRequest struct {
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
	RoomIds []*string `json:"RoomIds,omitempty" xml:"RoomIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetMeetingRoomsScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMeetingRoomsScheduleRequest) GetRoomIds() []*string {
	return s.RoomIds
}

func (s *GetMeetingRoomsScheduleRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMeetingRoomsScheduleRequest) SetEndTime(v string) *GetMeetingRoomsScheduleRequest {
	s.EndTime = &v
	return s
}

func (s *GetMeetingRoomsScheduleRequest) SetRoomIds(v []*string) *GetMeetingRoomsScheduleRequest {
	s.RoomIds = v
	return s
}

func (s *GetMeetingRoomsScheduleRequest) SetStartTime(v string) *GetMeetingRoomsScheduleRequest {
	s.StartTime = &v
	return s
}

func (s *GetMeetingRoomsScheduleRequest) Validate() error {
	return dara.Validate(s)
}
