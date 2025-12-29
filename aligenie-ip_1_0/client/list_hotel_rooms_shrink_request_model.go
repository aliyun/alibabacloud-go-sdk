// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelRoomsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelAdminRoomShrink(v string) *ListHotelRoomsShrinkRequest
	GetHotelAdminRoomShrink() *string
	SetHotelId(v string) *ListHotelRoomsShrinkRequest
	GetHotelId() *string
}

type ListHotelRoomsShrinkRequest struct {
	// if can be null:
	// true
	HotelAdminRoomShrink *string `json:"HotelAdminRoom,omitempty" xml:"HotelAdminRoom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ListHotelRoomsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelRoomsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsShrinkRequest) GetHotelAdminRoomShrink() *string {
	return s.HotelAdminRoomShrink
}

func (s *ListHotelRoomsShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelRoomsShrinkRequest) SetHotelAdminRoomShrink(v string) *ListHotelRoomsShrinkRequest {
	s.HotelAdminRoomShrink = &v
	return s
}

func (s *ListHotelRoomsShrinkRequest) SetHotelId(v string) *ListHotelRoomsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelRoomsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
