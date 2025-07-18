// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDevicesSharingStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevices(v []*UpdateUserDevicesSharingStatusResponseBodyDevices) *UpdateUserDevicesSharingStatusResponseBody
	GetDevices() []*UpdateUserDevicesSharingStatusResponseBodyDevices
	SetRequestId(v string) *UpdateUserDevicesSharingStatusResponseBody
	GetRequestId() *string
}

type UpdateUserDevicesSharingStatusResponseBody struct {
	Devices []*UpdateUserDevicesSharingStatusResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserDevicesSharingStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusResponseBody) GetDevices() []*UpdateUserDevicesSharingStatusResponseBodyDevices {
	return s.Devices
}

func (s *UpdateUserDevicesSharingStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserDevicesSharingStatusResponseBody) SetDevices(v []*UpdateUserDevicesSharingStatusResponseBodyDevices) *UpdateUserDevicesSharingStatusResponseBody {
	s.Devices = v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBody) SetRequestId(v string) *UpdateUserDevicesSharingStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateUserDevicesSharingStatusResponseBodyDevices struct {
	// example:
	//
	// Online
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// example:
	//
	// 2.2.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// Apple M1
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// example:
	//
	// MacBookPro17,1
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// example:
	//
	// Online
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 3.5.1
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	// example:
	//
	// APPLE SSD AP0512Q Media
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// example:
	//
	// Enabled
	DlpStatus *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// Enabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Enabled
	NacStatus        *string                                                              `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	NetInterfaceInfo []*UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo `json:"NetInterfaceInfo,omitempty" xml:"NetInterfaceInfo,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	PaStatus *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	// example:
	//
	// 11.49.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateUserDevicesSharingStatusResponseBodyDevices) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetAppStatus() *string {
	return s.AppStatus
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetAppVersion() *string {
	return s.AppVersion
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetCPU() *string {
	return s.CPU
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDepartment() *string {
	return s.Department
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDeviceType() *string {
	return s.DeviceType
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDeviceVersion() *string {
	return s.DeviceVersion
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDisk() *string {
	return s.Disk
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetDlpStatus() *string {
	return s.DlpStatus
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetHostname() *string {
	return s.Hostname
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetIaStatus() *string {
	return s.IaStatus
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetInnerIP() *string {
	return s.InnerIP
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetMac() *string {
	return s.Mac
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetMemory() *string {
	return s.Memory
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetNacStatus() *string {
	return s.NacStatus
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetNetInterfaceInfo() []*UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo {
	return s.NetInterfaceInfo
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetPaStatus() *string {
	return s.PaStatus
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetSharingStatus() *bool {
	return s.SharingStatus
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetSrcIP() *string {
	return s.SrcIP
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) GetUsername() *string {
	return s.Username
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetAppStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.AppStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetAppVersion(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.AppVersion = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetCPU(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.CPU = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetCreateTime(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDepartment(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Department = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceBelong(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceBelong = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceModel(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceModel = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceTag(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceTag = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceType(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceType = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceVersion(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceVersion = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDisk(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Disk = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDlpStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DlpStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetHostname(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Hostname = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetIaStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.IaStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetInnerIP(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.InnerIP = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetMac(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Mac = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetMemory(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Memory = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetNacStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.NacStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetNetInterfaceInfo(v []*UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.NetInterfaceInfo = v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetPaStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.PaStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetSaseUserId(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.SaseUserId = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetSharingStatus(v bool) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.SharingStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetSrcIP(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.SrcIP = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetUpdateTime(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.UpdateTime = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetUsername(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Username = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) Validate() error {
	return dara.Validate(s)
}

type UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo struct {
	Mac  *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo) GetMac() *string {
	return s.Mac
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo) GetName() *string {
	return s.Name
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo) SetMac(v string) *UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo {
	s.Mac = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo) SetName(v string) *UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo {
	s.Name = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevicesNetInterfaceInfo) Validate() error {
	return dara.Validate(s)
}
