// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iControlRoomDeviceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCmd(v string) *ControlRoomDeviceShrinkRequest
	GetCmd() *string
	SetDeviceIndex(v int32) *ControlRoomDeviceShrinkRequest
	GetDeviceIndex() *int32
	SetDeviceNumber(v string) *ControlRoomDeviceShrinkRequest
	GetDeviceNumber() *string
	SetHotelId(v string) *ControlRoomDeviceShrinkRequest
	GetHotelId() *string
	SetPropertiesShrink(v string) *ControlRoomDeviceShrinkRequest
	GetPropertiesShrink() *string
	SetRoomNo(v string) *ControlRoomDeviceShrinkRequest
	GetRoomNo() *string
}

type ControlRoomDeviceShrinkRequest struct {
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
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ControlRoomDeviceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ControlRoomDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceShrinkRequest) GetCmd() *string {
	return s.Cmd
}

func (s *ControlRoomDeviceShrinkRequest) GetDeviceIndex() *int32 {
	return s.DeviceIndex
}

func (s *ControlRoomDeviceShrinkRequest) GetDeviceNumber() *string {
	return s.DeviceNumber
}

func (s *ControlRoomDeviceShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ControlRoomDeviceShrinkRequest) GetPropertiesShrink() *string {
	return s.PropertiesShrink
}

func (s *ControlRoomDeviceShrinkRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ControlRoomDeviceShrinkRequest) SetCmd(v string) *ControlRoomDeviceShrinkRequest {
	s.Cmd = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetDeviceIndex(v int32) *ControlRoomDeviceShrinkRequest {
	s.DeviceIndex = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetDeviceNumber(v string) *ControlRoomDeviceShrinkRequest {
	s.DeviceNumber = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetHotelId(v string) *ControlRoomDeviceShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetPropertiesShrink(v string) *ControlRoomDeviceShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) SetRoomNo(v string) *ControlRoomDeviceShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *ControlRoomDeviceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
