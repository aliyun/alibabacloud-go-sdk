// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDevicesStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevices(v []*UpdateUserDevicesStatusResponseBodyDevices) *UpdateUserDevicesStatusResponseBody
	GetDevices() []*UpdateUserDevicesStatusResponseBodyDevices
	SetRequestId(v string) *UpdateUserDevicesStatusResponseBody
	GetRequestId() *string
}

type UpdateUserDevicesStatusResponseBody struct {
	Devices []*UpdateUserDevicesStatusResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserDevicesStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusResponseBody) GetDevices() []*UpdateUserDevicesStatusResponseBodyDevices {
	return s.Devices
}

func (s *UpdateUserDevicesStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserDevicesStatusResponseBody) SetDevices(v []*UpdateUserDevicesStatusResponseBodyDevices) *UpdateUserDevicesStatusResponseBody {
	s.Devices = v
	return s
}

func (s *UpdateUserDevicesStatusResponseBody) SetRequestId(v string) *UpdateUserDevicesStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateUserDevicesStatusResponseBodyDevices struct {
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
	// 2023-07-17 18:46:55
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
	NacStatus        *string                                                       `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	NetInterfaceInfo []*UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo `json:"NetInterfaceInfo,omitempty" xml:"NetInterfaceInfo,omitempty" type:"Repeated"`
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

func (s UpdateUserDevicesStatusResponseBodyDevices) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesStatusResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetAppStatus() *string {
	return s.AppStatus
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetAppVersion() *string {
	return s.AppVersion
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetCPU() *string {
	return s.CPU
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDepartment() *string {
	return s.Department
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDeviceType() *string {
	return s.DeviceType
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDeviceVersion() *string {
	return s.DeviceVersion
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDisk() *string {
	return s.Disk
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetDlpStatus() *string {
	return s.DlpStatus
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetHostname() *string {
	return s.Hostname
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetIaStatus() *string {
	return s.IaStatus
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetInnerIP() *string {
	return s.InnerIP
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetMac() *string {
	return s.Mac
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetMemory() *string {
	return s.Memory
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetNacStatus() *string {
	return s.NacStatus
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetNetInterfaceInfo() []*UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo {
	return s.NetInterfaceInfo
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetPaStatus() *string {
	return s.PaStatus
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetSharingStatus() *bool {
	return s.SharingStatus
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetSrcIP() *string {
	return s.SrcIP
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) GetUsername() *string {
	return s.Username
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetAppStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.AppStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetAppVersion(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.AppVersion = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetCPU(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.CPU = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetCreateTime(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDepartment(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Department = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceBelong(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceBelong = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceModel(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceModel = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceTag(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceTag = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceType(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceType = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceVersion(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceVersion = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDisk(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Disk = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDlpStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DlpStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetHostname(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Hostname = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetIaStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.IaStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetInnerIP(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.InnerIP = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetMac(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Mac = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetMemory(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Memory = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetNacStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.NacStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetNetInterfaceInfo(v []*UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo) *UpdateUserDevicesStatusResponseBodyDevices {
	s.NetInterfaceInfo = v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetPaStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.PaStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetSaseUserId(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.SaseUserId = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetSharingStatus(v bool) *UpdateUserDevicesStatusResponseBodyDevices {
	s.SharingStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetSrcIP(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.SrcIP = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetUpdateTime(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.UpdateTime = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetUsername(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Username = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) Validate() error {
	return dara.Validate(s)
}

type UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo struct {
	Mac  *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo) GetMac() *string {
	return s.Mac
}

func (s *UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo) GetName() *string {
	return s.Name
}

func (s *UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo) SetMac(v string) *UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo {
	s.Mac = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo) SetName(v string) *UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo {
	s.Name = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo) Validate() error {
	return dara.Validate(s)
}
