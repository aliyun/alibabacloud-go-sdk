// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomControlDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryRoomControlDevicesResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryRoomControlDevicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRoomControlDevicesResponseBody
	GetRequestId() *string
	SetResult(v []*QueryRoomControlDevicesResponseBodyResult) *QueryRoomControlDevicesResponseBody
	GetResult() []*QueryRoomControlDevicesResponseBodyResult
}

type QueryRoomControlDevicesResponseBody struct {
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
	// fdsgfdscvre
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryRoomControlDevicesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryRoomControlDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryRoomControlDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRoomControlDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRoomControlDevicesResponseBody) GetResult() []*QueryRoomControlDevicesResponseBodyResult {
	return s.Result
}

func (s *QueryRoomControlDevicesResponseBody) SetCode(v int32) *QueryRoomControlDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBody) SetMessage(v string) *QueryRoomControlDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBody) SetRequestId(v string) *QueryRoomControlDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBody) SetResult(v []*QueryRoomControlDevicesResponseBodyResult) *QueryRoomControlDevicesResponseBody {
	s.Result = v
	return s
}

func (s *QueryRoomControlDevicesResponseBody) Validate() error {
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

type QueryRoomControlDevicesResponseBodyResult struct {
	Devices []*QueryRoomControlDevicesResponseBodyResultDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// room
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
}

func (s QueryRoomControlDevicesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBodyResult) GetDevices() []*QueryRoomControlDevicesResponseBodyResultDevices {
	return s.Devices
}

func (s *QueryRoomControlDevicesResponseBodyResult) GetLocation() *string {
	return s.Location
}

func (s *QueryRoomControlDevicesResponseBodyResult) GetLocationName() *string {
	return s.LocationName
}

func (s *QueryRoomControlDevicesResponseBodyResult) SetDevices(v []*QueryRoomControlDevicesResponseBodyResultDevices) *QueryRoomControlDevicesResponseBodyResult {
	s.Devices = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResult) SetLocation(v string) *QueryRoomControlDevicesResponseBodyResult {
	s.Location = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResult) SetLocationName(v string) *QueryRoomControlDevicesResponseBodyResult {
	s.LocationName = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResult) Validate() error {
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

type QueryRoomControlDevicesResponseBodyResultDevices struct {
	AliasList         []*string                                                          `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	ConnectType       *string                                                            `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	DN                *string                                                            `json:"DN,omitempty" xml:"DN,omitempty"`
	DeviceName        *string                                                            `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceStatus      *string                                                            `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	MultiKeySwitchExt *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt `json:"MultiKeySwitchExt,omitempty" xml:"MultiKeySwitchExt,omitempty" type:"Struct"`
	// example:
	//
	// light
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// night_light
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	PK     *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryRoomControlDevicesResponseBodyResultDevices) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBodyResultDevices) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetAliasList() []*string {
	return s.AliasList
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetConnectType() *string {
	return s.ConnectType
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetDN() *string {
	return s.DN
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetMultiKeySwitchExt() *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt {
	return s.MultiKeySwitchExt
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetName() *string {
	return s.Name
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetNumber() *string {
	return s.Number
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) GetPK() *string {
	return s.PK
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetAliasList(v []*string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.AliasList = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetConnectType(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.ConnectType = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetDN(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.DN = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetDeviceName(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.DeviceName = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetDeviceStatus(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.DeviceStatus = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetMultiKeySwitchExt(v *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.MultiKeySwitchExt = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetName(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.Name = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetNumber(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.Number = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) SetPK(v string) *QueryRoomControlDevicesResponseBodyResultDevices {
	s.PK = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevices) Validate() error {
	if s.MultiKeySwitchExt != nil {
		if err := s.MultiKeySwitchExt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt struct {
	SwitchList []*QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList `json:"SwitchList,omitempty" xml:"SwitchList,omitempty" type:"Repeated"`
}

func (s QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) GetSwitchList() []*QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	return s.SwitchList
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) SetSwitchList(v []*QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt {
	s.SwitchList = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExt) Validate() error {
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

type QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList struct {
	AliasList    []*string `json:"AliasList,omitempty" xml:"AliasList,omitempty" type:"Repeated"`
	Category     *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeviceIndex  *int32    `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	DeviceName   *string   `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceStatus *string   `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	ElementCode  *string   `json:"ElementCode,omitempty" xml:"ElementCode,omitempty"`
	Location     *string   `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetAliasList() []*string {
	return s.AliasList
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetCategory() *string {
	return s.Category
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetDeviceIndex() *int32 {
	return s.DeviceIndex
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetElementCode() *string {
	return s.ElementCode
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) GetLocation() *string {
	return s.Location
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetAliasList(v []*string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.AliasList = v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetCategory(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Category = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceIndex(v int32) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceIndex = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceName(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceName = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetDeviceStatus(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.DeviceStatus = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetElementCode(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.ElementCode = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) SetLocation(v string) *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList {
	s.Location = &v
	return s
}

func (s *QueryRoomControlDevicesResponseBodyResultDevicesMultiKeySwitchExtSwitchList) Validate() error {
	return dara.Validate(s)
}
