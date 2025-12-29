// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomControlDevicesAndStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryRoomControlDevicesAndStatusResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryRoomControlDevicesAndStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRoomControlDevicesAndStatusResponseBody
	GetRequestId() *string
	SetResult(v []*QueryRoomControlDevicesAndStatusResponseBodyResult) *QueryRoomControlDevicesAndStatusResponseBody
	GetResult() []*QueryRoomControlDevicesAndStatusResponseBodyResult
	SetStatusCode(v int32) *QueryRoomControlDevicesAndStatusResponseBody
	GetStatusCode() *int32
}

type QueryRoomControlDevicesAndStatusResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67***6FA
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryRoomControlDevicesAndStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) GetResult() []*QueryRoomControlDevicesAndStatusResponseBodyResult {
	return s.Result
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetCode(v int32) *QueryRoomControlDevicesAndStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetMessage(v string) *QueryRoomControlDevicesAndStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetRequestId(v string) *QueryRoomControlDevicesAndStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetResult(v []*QueryRoomControlDevicesAndStatusResponseBodyResult) *QueryRoomControlDevicesAndStatusResponseBody {
	s.Result = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) SetStatusCode(v int32) *QueryRoomControlDevicesAndStatusResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRoomControlDevicesAndStatusResponseBodyResult struct {
	Devices []*QueryRoomControlDevicesAndStatusResponseBodyResultDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// room
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) GetDevices() []*QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	return s.Devices
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) GetLocation() *string {
	return s.Location
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) GetLocationName() *string {
	return s.LocationName
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) SetDevices(v []*QueryRoomControlDevicesAndStatusResponseBodyResultDevices) *QueryRoomControlDevicesAndStatusResponseBodyResult {
	s.Devices = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) SetLocation(v string) *QueryRoomControlDevicesAndStatusResponseBodyResult {
	s.Location = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) SetLocationName(v string) *QueryRoomControlDevicesAndStatusResponseBodyResult {
	s.LocationName = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) SetRoomNo(v string) *QueryRoomControlDevicesAndStatusResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResult) Validate() error {
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

type QueryRoomControlDevicesAndStatusResponseBodyResultDevices struct {
	AliasList []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	Brand     *string   `json:"Brand,omitempty" xml:"Brand,omitempty"`
	City      *string   `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// rcu
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	DeviceName  *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// {"powerstate": "1"}
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 3c5d***9ec
	Dn *string `json:"Dn,omitempty" xml:"Dn,omitempty"`
	// example:
	//
	// 9**7
	InfraredId *string `json:"InfraredId,omitempty" xml:"InfraredId,omitempty"`
	// example:
	//
	// 2
	InfraredIndex *string `json:"InfraredIndex,omitempty" xml:"InfraredIndex,omitempty"`
	// example:
	//
	// 3.0
	InfraredVersion   *string                                                                     `json:"InfraredVersion,omitempty" xml:"InfraredVersion,omitempty"`
	MultiKeySwitchExt *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt `json:"MultiKeySwitchExt,omitempty" xml:"MultiKeySwitchExt,omitempty" type:"Struct"`
	// example:
	//
	// light
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// night_light
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 50255129
	Pk              *string            `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Province        *string            `json:"Province,omitempty" xml:"Province,omitempty"`
	ServiceProvider *string            `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
	Status          map[string]*string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevices) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetAliasList() []*string {
	return s.AliasList
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetBrand() *string {
	return s.Brand
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetCity() *string {
	return s.City
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetConnectType() *string {
	return s.ConnectType
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetDn() *string {
	return s.Dn
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetInfraredId() *string {
	return s.InfraredId
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetInfraredIndex() *string {
	return s.InfraredIndex
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetInfraredVersion() *string {
	return s.InfraredVersion
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetMultiKeySwitchExt() *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt {
	return s.MultiKeySwitchExt
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetName() *string {
	return s.Name
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetNumber() *string {
	return s.Number
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetPk() *string {
	return s.Pk
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetProvince() *string {
	return s.Province
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) GetStatus() map[string]*string {
	return s.Status
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetAliasList(v []*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.AliasList = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetBrand(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Brand = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetCity(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.City = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetConnectType(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.ConnectType = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetDeviceName(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.DeviceName = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetDeviceStatus(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.DeviceStatus = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetDn(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Dn = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetInfraredId(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.InfraredId = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetInfraredIndex(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.InfraredIndex = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetInfraredVersion(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.InfraredVersion = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetMultiKeySwitchExt(v *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.MultiKeySwitchExt = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetName(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Name = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetNumber(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Number = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetPk(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Pk = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetProvince(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Province = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetServiceProvider(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.ServiceProvider = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) SetStatus(v map[string]*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevices {
	s.Status = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevices) Validate() error {
	if s.MultiKeySwitchExt != nil {
		if err := s.MultiKeySwitchExt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt struct {
	SwitchList []*QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList `json:"SwitchList,omitempty" xml:"SwitchList,omitempty" type:"Repeated"`
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) GetSwitchList() []*QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	return s.SwitchList
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) SetSwitchList(v []*QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt {
	s.SwitchList = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExt) Validate() error {
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

type QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList struct {
	AliasList []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	// example:
	//
	// light
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 1
	DeviceIndex *int32  `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	DeviceName  *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// {
	//
	//       "powerstate": "0"
	//
	// }
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// e2
	ElementCode *string `json:"ElementCode,omitempty" xml:"ElementCode,omitempty"`
	// example:
	//
	// room
	Location *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	Status   map[string]*string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags     []*string          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetAliasList() []*string {
	return s.AliasList
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetCategory() *string {
	return s.Category
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetDeviceIndex() *int32 {
	return s.DeviceIndex
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetElementCode() *string {
	return s.ElementCode
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetLocation() *string {
	return s.Location
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetStatus() map[string]*string {
	return s.Status
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetTags() []*string {
	return s.Tags
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetAliasList(v []*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.AliasList = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetCategory(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Category = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceIndex(v int32) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceIndex = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceName(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceName = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceStatus(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceStatus = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetElementCode(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.ElementCode = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetLocation(v string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Location = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetStatus(v map[string]*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Status = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetTags(v []*string) *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Tags = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponseBodyResultDevicesMultiKeySwitchExtSwitchList) Validate() error {
	return dara.Validate(s)
}
