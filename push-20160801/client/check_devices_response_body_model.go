// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceCheckInfos(v *CheckDevicesResponseBodyDeviceCheckInfos) *CheckDevicesResponseBody
	GetDeviceCheckInfos() *CheckDevicesResponseBodyDeviceCheckInfos
	SetRequestId(v string) *CheckDevicesResponseBody
	GetRequestId() *string
}

type CheckDevicesResponseBody struct {
	DeviceCheckInfos *CheckDevicesResponseBodyDeviceCheckInfos `json:"DeviceCheckInfos,omitempty" xml:"DeviceCheckInfos,omitempty" type:"Struct"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDevicesResponseBody) GetDeviceCheckInfos() *CheckDevicesResponseBodyDeviceCheckInfos {
	return s.DeviceCheckInfos
}

func (s *CheckDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDevicesResponseBody) SetDeviceCheckInfos(v *CheckDevicesResponseBodyDeviceCheckInfos) *CheckDevicesResponseBody {
	s.DeviceCheckInfos = v
	return s
}

func (s *CheckDevicesResponseBody) SetRequestId(v string) *CheckDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckDevicesResponseBodyDeviceCheckInfos struct {
	DeviceCheckInfo []*CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo `json:"DeviceCheckInfo,omitempty" xml:"DeviceCheckInfo,omitempty" type:"Repeated"`
}

func (s CheckDevicesResponseBodyDeviceCheckInfos) String() string {
	return dara.Prettify(s)
}

func (s CheckDevicesResponseBodyDeviceCheckInfos) GoString() string {
	return s.String()
}

func (s *CheckDevicesResponseBodyDeviceCheckInfos) GetDeviceCheckInfo() []*CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo {
	return s.DeviceCheckInfo
}

func (s *CheckDevicesResponseBodyDeviceCheckInfos) SetDeviceCheckInfo(v []*CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) *CheckDevicesResponseBodyDeviceCheckInfos {
	s.DeviceCheckInfo = v
	return s
}

func (s *CheckDevicesResponseBodyDeviceCheckInfos) Validate() error {
	return dara.Validate(s)
}

type CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo struct {
	// example:
	//
	// true
	Available *bool `json:"Available,omitempty" xml:"Available,omitempty"`
	// example:
	//
	// ae296f3b04a58a05b30c95f****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) GoString() string {
	return s.String()
}

func (s *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) GetAvailable() *bool {
	return s.Available
}

func (s *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) SetAvailable(v bool) *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo {
	s.Available = &v
	return s
}

func (s *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) SetDeviceId(v string) *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo {
	s.DeviceId = &v
	return s
}

func (s *CheckDevicesResponseBodyDeviceCheckInfosDeviceCheckInfo) Validate() error {
	return dara.Validate(s)
}
