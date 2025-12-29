// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDeviceStatusWithAkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryCnName(v string) *SyncDeviceStatusWithAkRequest
	GetCategoryCnName() *string
	SetCategoryEnName(v string) *SyncDeviceStatusWithAkRequest
	GetCategoryEnName() *string
	SetDeviceName(v string) *SyncDeviceStatusWithAkRequest
	GetDeviceName() *string
	SetHotelId(v string) *SyncDeviceStatusWithAkRequest
	GetHotelId() *string
	SetLocation(v string) *SyncDeviceStatusWithAkRequest
	GetLocation() *string
	SetLocationCnName(v string) *SyncDeviceStatusWithAkRequest
	GetLocationCnName() *string
	SetNumber(v string) *SyncDeviceStatusWithAkRequest
	GetNumber() *string
	SetRoomNo(v string) *SyncDeviceStatusWithAkRequest
	GetRoomNo() *string
	SetSwitch(v int32) *SyncDeviceStatusWithAkRequest
	GetSwitch() *int32
	SetFanSpeed(v string) *SyncDeviceStatusWithAkRequest
	GetFanSpeed() *string
	SetMode(v string) *SyncDeviceStatusWithAkRequest
	GetMode() *string
	SetRoomTemperature(v string) *SyncDeviceStatusWithAkRequest
	GetRoomTemperature() *string
	SetTemperature(v string) *SyncDeviceStatusWithAkRequest
	GetTemperature() *string
	SetValue(v int32) *SyncDeviceStatusWithAkRequest
	GetValue() *int32
}

type SyncDeviceStatusWithAkRequest struct {
	CategoryCnName *string `json:"CategoryCnName,omitempty" xml:"CategoryCnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// light
	CategoryEnName *string `json:"CategoryEnName,omitempty" xml:"CategoryEnName,omitempty"`
	DeviceName     *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
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
	// room
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationCnName *string `json:"LocationCnName,omitempty" xml:"LocationCnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bedLight
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Switch          *int32  `json:"Switch,omitempty" xml:"Switch,omitempty"`
	FanSpeed        *string `json:"fanSpeed,omitempty" xml:"fanSpeed,omitempty"`
	Mode            *string `json:"mode,omitempty" xml:"mode,omitempty"`
	RoomTemperature *string `json:"roomTemperature,omitempty" xml:"roomTemperature,omitempty"`
	Temperature     *string `json:"temperature,omitempty" xml:"temperature,omitempty"`
	Value           *int32  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SyncDeviceStatusWithAkRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDeviceStatusWithAkRequest) GoString() string {
	return s.String()
}

func (s *SyncDeviceStatusWithAkRequest) GetCategoryCnName() *string {
	return s.CategoryCnName
}

func (s *SyncDeviceStatusWithAkRequest) GetCategoryEnName() *string {
	return s.CategoryEnName
}

func (s *SyncDeviceStatusWithAkRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *SyncDeviceStatusWithAkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *SyncDeviceStatusWithAkRequest) GetLocation() *string {
	return s.Location
}

func (s *SyncDeviceStatusWithAkRequest) GetLocationCnName() *string {
	return s.LocationCnName
}

func (s *SyncDeviceStatusWithAkRequest) GetNumber() *string {
	return s.Number
}

func (s *SyncDeviceStatusWithAkRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *SyncDeviceStatusWithAkRequest) GetSwitch() *int32 {
	return s.Switch
}

func (s *SyncDeviceStatusWithAkRequest) GetFanSpeed() *string {
	return s.FanSpeed
}

func (s *SyncDeviceStatusWithAkRequest) GetMode() *string {
	return s.Mode
}

func (s *SyncDeviceStatusWithAkRequest) GetRoomTemperature() *string {
	return s.RoomTemperature
}

func (s *SyncDeviceStatusWithAkRequest) GetTemperature() *string {
	return s.Temperature
}

func (s *SyncDeviceStatusWithAkRequest) GetValue() *int32 {
	return s.Value
}

func (s *SyncDeviceStatusWithAkRequest) SetCategoryCnName(v string) *SyncDeviceStatusWithAkRequest {
	s.CategoryCnName = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetCategoryEnName(v string) *SyncDeviceStatusWithAkRequest {
	s.CategoryEnName = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetDeviceName(v string) *SyncDeviceStatusWithAkRequest {
	s.DeviceName = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetHotelId(v string) *SyncDeviceStatusWithAkRequest {
	s.HotelId = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetLocation(v string) *SyncDeviceStatusWithAkRequest {
	s.Location = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetLocationCnName(v string) *SyncDeviceStatusWithAkRequest {
	s.LocationCnName = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetNumber(v string) *SyncDeviceStatusWithAkRequest {
	s.Number = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetRoomNo(v string) *SyncDeviceStatusWithAkRequest {
	s.RoomNo = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetSwitch(v int32) *SyncDeviceStatusWithAkRequest {
	s.Switch = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetFanSpeed(v string) *SyncDeviceStatusWithAkRequest {
	s.FanSpeed = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetMode(v string) *SyncDeviceStatusWithAkRequest {
	s.Mode = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetRoomTemperature(v string) *SyncDeviceStatusWithAkRequest {
	s.RoomTemperature = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetTemperature(v string) *SyncDeviceStatusWithAkRequest {
	s.Temperature = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) SetValue(v int32) *SyncDeviceStatusWithAkRequest {
	s.Value = &v
	return s
}

func (s *SyncDeviceStatusWithAkRequest) Validate() error {
	return dara.Validate(s)
}
