// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iControlRoomDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCmd(v string) *ControlRoomDeviceRequest
	GetCmd() *string
	SetDeviceIndex(v int32) *ControlRoomDeviceRequest
	GetDeviceIndex() *int32
	SetDeviceNumber(v string) *ControlRoomDeviceRequest
	GetDeviceNumber() *string
	SetHotelId(v string) *ControlRoomDeviceRequest
	GetHotelId() *string
	SetProperties(v map[string]*string) *ControlRoomDeviceRequest
	GetProperties() map[string]*string
	SetRoomNo(v string) *ControlRoomDeviceRequest
	GetRoomNo() *string
}

type ControlRoomDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// thing.attribute.set
	//
	// thing.attribute.adjust
	Cmd         *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	DeviceIndex *int32  `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INFRARED49122575595
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ControlRoomDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s ControlRoomDeviceRequest) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceRequest) GetCmd() *string {
	return s.Cmd
}

func (s *ControlRoomDeviceRequest) GetDeviceIndex() *int32 {
	return s.DeviceIndex
}

func (s *ControlRoomDeviceRequest) GetDeviceNumber() *string {
	return s.DeviceNumber
}

func (s *ControlRoomDeviceRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ControlRoomDeviceRequest) GetProperties() map[string]*string {
	return s.Properties
}

func (s *ControlRoomDeviceRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ControlRoomDeviceRequest) SetCmd(v string) *ControlRoomDeviceRequest {
	s.Cmd = &v
	return s
}

func (s *ControlRoomDeviceRequest) SetDeviceIndex(v int32) *ControlRoomDeviceRequest {
	s.DeviceIndex = &v
	return s
}

func (s *ControlRoomDeviceRequest) SetDeviceNumber(v string) *ControlRoomDeviceRequest {
	s.DeviceNumber = &v
	return s
}

func (s *ControlRoomDeviceRequest) SetHotelId(v string) *ControlRoomDeviceRequest {
	s.HotelId = &v
	return s
}

func (s *ControlRoomDeviceRequest) SetProperties(v map[string]*string) *ControlRoomDeviceRequest {
	s.Properties = v
	return s
}

func (s *ControlRoomDeviceRequest) SetRoomNo(v string) *ControlRoomDeviceRequest {
	s.RoomNo = &v
	return s
}

func (s *ControlRoomDeviceRequest) Validate() error {
	return dara.Validate(s)
}
