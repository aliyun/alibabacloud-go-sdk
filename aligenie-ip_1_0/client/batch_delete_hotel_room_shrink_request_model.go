// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteHotelRoomShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *BatchDeleteHotelRoomShrinkRequest
	GetHotelId() *string
	SetRoomNoListShrink(v string) *BatchDeleteHotelRoomShrinkRequest
	GetRoomNoListShrink() *string
}

type BatchDeleteHotelRoomShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNoListShrink *string `json:"RoomNoList,omitempty" xml:"RoomNoList,omitempty"`
}

func (s BatchDeleteHotelRoomShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteHotelRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *BatchDeleteHotelRoomShrinkRequest) GetRoomNoListShrink() *string {
	return s.RoomNoListShrink
}

func (s *BatchDeleteHotelRoomShrinkRequest) SetHotelId(v string) *BatchDeleteHotelRoomShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *BatchDeleteHotelRoomShrinkRequest) SetRoomNoListShrink(v string) *BatchDeleteHotelRoomShrinkRequest {
	s.RoomNoListShrink = &v
	return s
}

func (s *BatchDeleteHotelRoomShrinkRequest) Validate() error {
	return dara.Validate(s)
}
