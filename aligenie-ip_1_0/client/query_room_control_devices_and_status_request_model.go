// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomControlDevicesAndStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *QueryRoomControlDevicesAndStatusRequest
	GetHotelId() *string
	SetRoomNo(v string) *QueryRoomControlDevicesAndStatusRequest
	GetRoomNo() *string
}

type QueryRoomControlDevicesAndStatusRequest struct {
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *QueryRoomControlDevicesAndStatusRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *QueryRoomControlDevicesAndStatusRequest) SetHotelId(v string) *QueryRoomControlDevicesAndStatusRequest {
	s.HotelId = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusRequest) SetRoomNo(v string) *QueryRoomControlDevicesAndStatusRequest {
	s.RoomNo = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusRequest) Validate() error {
	return dara.Validate(s)
}
