// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelRoomDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *GetHotelRoomDeviceRequest
	GetHotelId() *string
	SetRoomNo(v string) *GetHotelRoomDeviceRequest
	GetRoomNo() *string
}

type GetHotelRoomDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s GetHotelRoomDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelRoomDeviceRequest) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelRoomDeviceRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *GetHotelRoomDeviceRequest) SetHotelId(v string) *GetHotelRoomDeviceRequest {
	s.HotelId = &v
	return s
}

func (s *GetHotelRoomDeviceRequest) SetRoomNo(v string) *GetHotelRoomDeviceRequest {
	s.RoomNo = &v
	return s
}

func (s *GetHotelRoomDeviceRequest) Validate() error {
	return dara.Validate(s)
}
