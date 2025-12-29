// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelAlarmShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *CreateHotelAlarmShrinkRequest
	GetHotelId() *string
	SetMusicType(v string) *CreateHotelAlarmShrinkRequest
	GetMusicType() *string
	SetRoomsShrink(v string) *CreateHotelAlarmShrinkRequest
	GetRoomsShrink() *string
	SetScheduleInfoShrink(v string) *CreateHotelAlarmShrinkRequest
	GetScheduleInfoShrink() *string
}

type CreateHotelAlarmShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cf2446fc9d144c85aaee4f9ae20a96e7
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// DOU_YIN
	MusicType *string `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// This parameter is required.
	RoomsShrink *string `json:"Rooms,omitempty" xml:"Rooms,omitempty"`
	// This parameter is required.
	ScheduleInfoShrink *string `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty"`
}

func (s CreateHotelAlarmShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *CreateHotelAlarmShrinkRequest) GetMusicType() *string {
	return s.MusicType
}

func (s *CreateHotelAlarmShrinkRequest) GetRoomsShrink() *string {
	return s.RoomsShrink
}

func (s *CreateHotelAlarmShrinkRequest) GetScheduleInfoShrink() *string {
	return s.ScheduleInfoShrink
}

func (s *CreateHotelAlarmShrinkRequest) SetHotelId(v string) *CreateHotelAlarmShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *CreateHotelAlarmShrinkRequest) SetMusicType(v string) *CreateHotelAlarmShrinkRequest {
	s.MusicType = &v
	return s
}

func (s *CreateHotelAlarmShrinkRequest) SetRoomsShrink(v string) *CreateHotelAlarmShrinkRequest {
	s.RoomsShrink = &v
	return s
}

func (s *CreateHotelAlarmShrinkRequest) SetScheduleInfoShrink(v string) *CreateHotelAlarmShrinkRequest {
	s.ScheduleInfoShrink = &v
	return s
}

func (s *CreateHotelAlarmShrinkRequest) Validate() error {
	return dara.Validate(s)
}
