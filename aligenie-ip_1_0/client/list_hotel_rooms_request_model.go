// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelRoomsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelAdminRoom(v *ListHotelRoomsRequestHotelAdminRoom) *ListHotelRoomsRequest
	GetHotelAdminRoom() *ListHotelRoomsRequestHotelAdminRoom
	SetHotelId(v string) *ListHotelRoomsRequest
	GetHotelId() *string
}

type ListHotelRoomsRequest struct {
	// if can be null:
	// true
	HotelAdminRoom *ListHotelRoomsRequestHotelAdminRoom `json:"HotelAdminRoom,omitempty" xml:"HotelAdminRoom,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ListHotelRoomsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsRequest) GetHotelAdminRoom() *ListHotelRoomsRequestHotelAdminRoom {
	return s.HotelAdminRoom
}

func (s *ListHotelRoomsRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelRoomsRequest) SetHotelAdminRoom(v *ListHotelRoomsRequestHotelAdminRoom) *ListHotelRoomsRequest {
	s.HotelAdminRoom = v
	return s
}

func (s *ListHotelRoomsRequest) SetHotelId(v string) *ListHotelRoomsRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelRoomsRequest) Validate() error {
	if s.HotelAdminRoom != nil {
		if err := s.HotelAdminRoom.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelRoomsRequestHotelAdminRoom struct {
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ListHotelRoomsRequestHotelAdminRoom) String() string {
	return dara.Prettify(s)
}

func (s ListHotelRoomsRequestHotelAdminRoom) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsRequestHotelAdminRoom) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ListHotelRoomsRequestHotelAdminRoom) SetRoomNo(v string) *ListHotelRoomsRequestHotelAdminRoom {
	s.RoomNo = &v
	return s
}

func (s *ListHotelRoomsRequestHotelAdminRoom) Validate() error {
	return dara.Validate(s)
}
