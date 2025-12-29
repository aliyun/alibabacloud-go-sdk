// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomControlDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableInfraredDeviceImport(v string) *ImportRoomControlDevicesRequest
	GetEnableInfraredDeviceImport() *string
	SetEnableMeshDeviceImport(v string) *ImportRoomControlDevicesRequest
	GetEnableMeshDeviceImport() *string
	SetHotelId(v string) *ImportRoomControlDevicesRequest
	GetHotelId() *string
	SetLocationDevices(v []*ImportRoomControlDevicesRequestLocationDevices) *ImportRoomControlDevicesRequest
	GetLocationDevices() []*ImportRoomControlDevicesRequestLocationDevices
	SetRoomNo(v string) *ImportRoomControlDevicesRequest
	GetRoomNo() *string
}

type ImportRoomControlDevicesRequest struct {
	EnableInfraredDeviceImport *string `json:"EnableInfraredDeviceImport,omitempty" xml:"EnableInfraredDeviceImport,omitempty"`
	EnableMeshDeviceImport     *string `json:"EnableMeshDeviceImport,omitempty" xml:"EnableMeshDeviceImport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vdgrefds
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	LocationDevices []*ImportRoomControlDevicesRequestLocationDevices `json:"LocationDevices,omitempty" xml:"LocationDevices,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ImportRoomControlDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesRequest) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequest) GetEnableInfraredDeviceImport() *string {
	return s.EnableInfraredDeviceImport
}

func (s *ImportRoomControlDevicesRequest) GetEnableMeshDeviceImport() *string {
	return s.EnableMeshDeviceImport
}

func (s *ImportRoomControlDevicesRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ImportRoomControlDevicesRequest) GetLocationDevices() []*ImportRoomControlDevicesRequestLocationDevices {
	return s.LocationDevices
}

func (s *ImportRoomControlDevicesRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ImportRoomControlDevicesRequest) SetEnableInfraredDeviceImport(v string) *ImportRoomControlDevicesRequest {
	s.EnableInfraredDeviceImport = &v
	return s
}

func (s *ImportRoomControlDevicesRequest) SetEnableMeshDeviceImport(v string) *ImportRoomControlDevicesRequest {
	s.EnableMeshDeviceImport = &v
	return s
}

func (s *ImportRoomControlDevicesRequest) SetHotelId(v string) *ImportRoomControlDevicesRequest {
	s.HotelId = &v
	return s
}

func (s *ImportRoomControlDevicesRequest) SetLocationDevices(v []*ImportRoomControlDevicesRequestLocationDevices) *ImportRoomControlDevicesRequest {
	s.LocationDevices = v
	return s
}

func (s *ImportRoomControlDevicesRequest) SetRoomNo(v string) *ImportRoomControlDevicesRequest {
	s.RoomNo = &v
	return s
}

func (s *ImportRoomControlDevicesRequest) Validate() error {
	if s.LocationDevices != nil {
		for _, item := range s.LocationDevices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportRoomControlDevicesRequestLocationDevices struct {
	Devices []*ImportRoomControlDevicesRequestLocationDevicesDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// room
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
}

func (s ImportRoomControlDevicesRequestLocationDevices) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesRequestLocationDevices) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequestLocationDevices) GetDevices() []*ImportRoomControlDevicesRequestLocationDevicesDevices {
	return s.Devices
}

func (s *ImportRoomControlDevicesRequestLocationDevices) GetLocation() *string {
	return s.Location
}

func (s *ImportRoomControlDevicesRequestLocationDevices) GetLocationName() *string {
	return s.LocationName
}

func (s *ImportRoomControlDevicesRequestLocationDevices) SetDevices(v []*ImportRoomControlDevicesRequestLocationDevicesDevices) *ImportRoomControlDevicesRequestLocationDevices {
	s.Devices = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevices) SetLocation(v string) *ImportRoomControlDevicesRequestLocationDevices {
	s.Location = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevices) SetLocationName(v string) *ImportRoomControlDevicesRequestLocationDevices {
	s.LocationName = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevices) Validate() error {
	if s.Devices != nil {
		for _, item := range s.Devices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportRoomControlDevicesRequestLocationDevicesDevices struct {
	AliasList   []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	Brand       *string   `json:"Brand,omitempty" xml:"Brand,omitempty"`
	City        *string   `json:"City,omitempty" xml:"City,omitempty"`
	ConnectType *string   `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// This parameter is required.
	DeviceName        *string                                                                 `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	Dn                *string                                                                 `json:"Dn,omitempty" xml:"Dn,omitempty"`
	InfraredId        *string                                                                 `json:"InfraredId,omitempty" xml:"InfraredId,omitempty"`
	InfraredIndex     *string                                                                 `json:"InfraredIndex,omitempty" xml:"InfraredIndex,omitempty"`
	InfraredVersion   *string                                                                 `json:"InfraredVersion,omitempty" xml:"InfraredVersion,omitempty"`
	MultiKeySwitchExt *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt `json:"MultiKeySwitchExt,omitempty" xml:"MultiKeySwitchExt,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// light
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// night_light
	Number          *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Pk              *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Province        *string `json:"Province,omitempty" xml:"Province,omitempty"`
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevices) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevices) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetAliasList() []*string {
	return s.AliasList
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetBrand() *string {
	return s.Brand
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetCity() *string {
	return s.City
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetConnectType() *string {
	return s.ConnectType
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetDn() *string {
	return s.Dn
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetInfraredId() *string {
	return s.InfraredId
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetInfraredIndex() *string {
	return s.InfraredIndex
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetInfraredVersion() *string {
	return s.InfraredVersion
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetMultiKeySwitchExt() *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt {
	return s.MultiKeySwitchExt
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetName() *string {
	return s.Name
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetNumber() *string {
	return s.Number
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetPk() *string {
	return s.Pk
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetProvince() *string {
	return s.Province
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetAliasList(v []*string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.AliasList = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetBrand(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Brand = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetCity(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.City = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetConnectType(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.ConnectType = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetDeviceName(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.DeviceName = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetDn(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Dn = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetInfraredId(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.InfraredId = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetInfraredIndex(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.InfraredIndex = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetInfraredVersion(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.InfraredVersion = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetMultiKeySwitchExt(v *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.MultiKeySwitchExt = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetName(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Name = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetNumber(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Number = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetPk(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Pk = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetProvince(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.Province = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) SetServiceProvider(v string) *ImportRoomControlDevicesRequestLocationDevicesDevices {
	s.ServiceProvider = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevices) Validate() error {
	if s.MultiKeySwitchExt != nil {
		if err := s.MultiKeySwitchExt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt struct {
	SwitchList []*ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList `json:"SwitchList,omitempty" xml:"SwitchList,omitempty" type:"Repeated"`
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) GetSwitchList() []*ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	return s.SwitchList
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) SetSwitchList(v []*ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt {
	s.SwitchList = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExt) Validate() error {
	if s.SwitchList != nil {
		for _, item := range s.SwitchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList struct {
	AliasList   []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	Category    *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeviceIndex *int32    `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	DeviceName  *string   `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	Location    *string   `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) GetAliasList() []*string {
	return s.AliasList
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) GetCategory() *string {
	return s.Category
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) GetDeviceIndex() *int32 {
	return s.DeviceIndex
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) GetLocation() *string {
	return s.Location
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetAliasList(v []*string) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.AliasList = v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetCategory(v string) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.Category = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetDeviceIndex(v int32) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.DeviceIndex = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetDeviceName(v string) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.DeviceName = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) SetLocation(v string) *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList {
	s.Location = &v
	return s
}

func (s *ImportRoomControlDevicesRequestLocationDevicesDevicesMultiKeySwitchExtSwitchList) Validate() error {
	return dara.Validate(s)
}
