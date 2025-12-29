// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *QueryRoomStatusRequest
	GetHotelId() *string
	SetRoomNo(v string) *QueryRoomStatusRequest
	GetRoomNo() *string
}

type QueryRoomStatusRequest struct {
	// example:
	//
	// cf2446fc9d144c85aaee4f9ae20a96e7
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s QueryRoomStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryRoomStatusRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *QueryRoomStatusRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *QueryRoomStatusRequest) SetHotelId(v string) *QueryRoomStatusRequest {
	s.HotelId = &v
	return s
}

func (s *QueryRoomStatusRequest) SetRoomNo(v string) *QueryRoomStatusRequest {
	s.RoomNo = &v
	return s
}

func (s *QueryRoomStatusRequest) Validate() error {
	return dara.Validate(s)
}
