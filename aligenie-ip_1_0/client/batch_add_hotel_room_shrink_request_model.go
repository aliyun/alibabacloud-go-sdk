// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddHotelRoomShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *BatchAddHotelRoomShrinkRequest
	GetHotelId() *string
	SetRoomNoListShrink(v string) *BatchAddHotelRoomShrinkRequest
	GetRoomNoListShrink() *string
}

type BatchAddHotelRoomShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNoListShrink *string `json:"RoomNoList,omitempty" xml:"RoomNoList,omitempty"`
}

func (s BatchAddHotelRoomShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddHotelRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *BatchAddHotelRoomShrinkRequest) GetRoomNoListShrink() *string {
	return s.RoomNoListShrink
}

func (s *BatchAddHotelRoomShrinkRequest) SetHotelId(v string) *BatchAddHotelRoomShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *BatchAddHotelRoomShrinkRequest) SetRoomNoListShrink(v string) *BatchAddHotelRoomShrinkRequest {
	s.RoomNoListShrink = &v
	return s
}

func (s *BatchAddHotelRoomShrinkRequest) Validate() error {
	return dara.Validate(s)
}
