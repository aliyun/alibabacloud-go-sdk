// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelAlarmShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListHotelAlarmShrinkRequest
	GetHotelId() *string
	SetRoomsShrink(v string) *ListHotelAlarmShrinkRequest
	GetRoomsShrink() *string
}

type ListHotelAlarmShrinkRequest struct {
	// example:
	//
	// a7a3***013
	HotelId     *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	RoomsShrink *string `json:"Rooms,omitempty" xml:"Rooms,omitempty"`
}

func (s ListHotelAlarmShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelAlarmShrinkRequest) GetRoomsShrink() *string {
	return s.RoomsShrink
}

func (s *ListHotelAlarmShrinkRequest) SetHotelId(v string) *ListHotelAlarmShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelAlarmShrinkRequest) SetRoomsShrink(v string) *ListHotelAlarmShrinkRequest {
	s.RoomsShrink = &v
	return s
}

func (s *ListHotelAlarmShrinkRequest) Validate() error {
	return dara.Validate(s)
}
