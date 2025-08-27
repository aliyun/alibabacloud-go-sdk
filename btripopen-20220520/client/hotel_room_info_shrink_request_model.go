// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelRoomInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoomIdsShrink(v string) *HotelRoomInfoShrinkRequest
	GetRoomIdsShrink() *string
}

type HotelRoomInfoShrinkRequest struct {
	// This parameter is required.
	RoomIdsShrink *string `json:"room_ids,omitempty" xml:"room_ids,omitempty"`
}

func (s HotelRoomInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoShrinkRequest) GetRoomIdsShrink() *string {
	return s.RoomIdsShrink
}

func (s *HotelRoomInfoShrinkRequest) SetRoomIdsShrink(v string) *HotelRoomInfoShrinkRequest {
	s.RoomIdsShrink = &v
	return s
}

func (s *HotelRoomInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
