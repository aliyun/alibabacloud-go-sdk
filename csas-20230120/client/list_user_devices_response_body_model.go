// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevices(v []*ListUserDevicesResponseBodyDevices) *ListUserDevicesResponseBody
	GetDevices() []*ListUserDevicesResponseBodyDevices
	SetRequestId(v string) *ListUserDevicesResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListUserDevicesResponseBody
	GetTotalNum() *int64
}

type ListUserDevicesResponseBody struct {
	Devices []*ListUserDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListUserDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDevicesResponseBody) GetDevices() []*ListUserDevicesResponseBodyDevices {
	return s.Devices
}

func (s *ListUserDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserDevicesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListUserDevicesResponseBody) SetDevices(v []*ListUserDevicesResponseBodyDevices) *ListUserDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *ListUserDevicesResponseBody) SetRequestId(v string) *ListUserDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDevicesResponseBody) SetTotalNum(v int64) *ListUserDevicesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListUserDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserDevicesResponseBodyDevices struct {
	// example:
	//
	// Online
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// example:
	//
	// 2.2.0
	AppVersion      *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AutoLoginStatus *string `json:"AutoLoginStatus,omitempty" xml:"AutoLoginStatus,omitempty"`
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
	EdrStatus *string `json:"EdrStatus,omitempty" xml:"EdrStatus,omitempty"`
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
	Mac                 *string   `json:"Mac,omitempty" xml:"Mac,omitempty"`
	MatchDeviceGroupIds []*string `json:"MatchDeviceGroupIds,omitempty" xml:"MatchDeviceGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Enabled
	NacStatus        *string                                               `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	NetInterfaceInfo []*ListUserDevicesResponseBodyDevicesNetInterfaceInfo `json:"NetInterfaceInfo,omitempty" xml:"NetInterfaceInfo,omitempty" type:"Repeated"`
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
	SharingStatus *bool   `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	SnBaseBoard   *string `json:"SnBaseBoard,omitempty" xml:"SnBaseBoard,omitempty"`
	SnBios        *string `json:"SnBios,omitempty" xml:"SnBios,omitempty"`
	SnDiskDrive   *string `json:"SnDiskDrive,omitempty" xml:"SnDiskDrive,omitempty"`
	SnProcessor   *string `json:"SnProcessor,omitempty" xml:"SnProcessor,omitempty"`
	SnSystem      *string `json:"SnSystem,omitempty" xml:"SnSystem,omitempty"`
	// example:
	//
	// 11.49.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Workshop   *string `json:"Workshop,omitempty" xml:"Workshop,omitempty"`
}

func (s ListUserDevicesResponseBodyDevices) String() string {
	return dara.Prettify(s)
}

func (s ListUserDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *ListUserDevicesResponseBodyDevices) GetAppStatus() *string {
	return s.AppStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListUserDevicesResponseBodyDevices) GetAutoLoginStatus() *string {
	return s.AutoLoginStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetCPU() *string {
	return s.CPU
}

func (s *ListUserDevicesResponseBodyDevices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserDevicesResponseBodyDevices) GetDepartment() *string {
	return s.Department
}

func (s *ListUserDevicesResponseBodyDevices) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *ListUserDevicesResponseBodyDevices) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *ListUserDevicesResponseBodyDevices) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *ListUserDevicesResponseBodyDevices) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListUserDevicesResponseBodyDevices) GetDeviceVersion() *string {
	return s.DeviceVersion
}

func (s *ListUserDevicesResponseBodyDevices) GetDisk() *string {
	return s.Disk
}

func (s *ListUserDevicesResponseBodyDevices) GetDlpStatus() *string {
	return s.DlpStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetEdrStatus() *string {
	return s.EdrStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetHostname() *string {
	return s.Hostname
}

func (s *ListUserDevicesResponseBodyDevices) GetIaStatus() *string {
	return s.IaStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetInnerIP() *string {
	return s.InnerIP
}

func (s *ListUserDevicesResponseBodyDevices) GetMac() *string {
	return s.Mac
}

func (s *ListUserDevicesResponseBodyDevices) GetMatchDeviceGroupIds() []*string {
	return s.MatchDeviceGroupIds
}

func (s *ListUserDevicesResponseBodyDevices) GetMemory() *string {
	return s.Memory
}

func (s *ListUserDevicesResponseBodyDevices) GetNacStatus() *string {
	return s.NacStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetNetInterfaceInfo() []*ListUserDevicesResponseBodyDevicesNetInterfaceInfo {
	return s.NetInterfaceInfo
}

func (s *ListUserDevicesResponseBodyDevices) GetPaStatus() *string {
	return s.PaStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListUserDevicesResponseBodyDevices) GetSharingStatus() *bool {
	return s.SharingStatus
}

func (s *ListUserDevicesResponseBodyDevices) GetSnBaseBoard() *string {
	return s.SnBaseBoard
}

func (s *ListUserDevicesResponseBodyDevices) GetSnBios() *string {
	return s.SnBios
}

func (s *ListUserDevicesResponseBodyDevices) GetSnDiskDrive() *string {
	return s.SnDiskDrive
}

func (s *ListUserDevicesResponseBodyDevices) GetSnProcessor() *string {
	return s.SnProcessor
}

func (s *ListUserDevicesResponseBodyDevices) GetSnSystem() *string {
	return s.SnSystem
}

func (s *ListUserDevicesResponseBodyDevices) GetSrcIP() *string {
	return s.SrcIP
}

func (s *ListUserDevicesResponseBodyDevices) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListUserDevicesResponseBodyDevices) GetUsername() *string {
	return s.Username
}

func (s *ListUserDevicesResponseBodyDevices) GetWorkshop() *string {
	return s.Workshop
}

func (s *ListUserDevicesResponseBodyDevices) SetAppStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.AppStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetAppVersion(v string) *ListUserDevicesResponseBodyDevices {
	s.AppVersion = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetAutoLoginStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.AutoLoginStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCPU(v string) *ListUserDevicesResponseBodyDevices {
	s.CPU = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCreateTime(v string) *ListUserDevicesResponseBodyDevices {
	s.CreateTime = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDepartment(v string) *ListUserDevicesResponseBodyDevices {
	s.Department = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceBelong(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceBelong = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceModel(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceModel = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceTag(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceTag = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceType(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceType = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceVersion(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceVersion = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDisk(v string) *ListUserDevicesResponseBodyDevices {
	s.Disk = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDlpStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.DlpStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetEdrStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.EdrStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetHostname(v string) *ListUserDevicesResponseBodyDevices {
	s.Hostname = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetIaStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.IaStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetInnerIP(v string) *ListUserDevicesResponseBodyDevices {
	s.InnerIP = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetMac(v string) *ListUserDevicesResponseBodyDevices {
	s.Mac = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetMatchDeviceGroupIds(v []*string) *ListUserDevicesResponseBodyDevices {
	s.MatchDeviceGroupIds = v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetMemory(v string) *ListUserDevicesResponseBodyDevices {
	s.Memory = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetNacStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.NacStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetNetInterfaceInfo(v []*ListUserDevicesResponseBodyDevicesNetInterfaceInfo) *ListUserDevicesResponseBodyDevices {
	s.NetInterfaceInfo = v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetPaStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.PaStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSaseUserId(v string) *ListUserDevicesResponseBodyDevices {
	s.SaseUserId = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSharingStatus(v bool) *ListUserDevicesResponseBodyDevices {
	s.SharingStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSnBaseBoard(v string) *ListUserDevicesResponseBodyDevices {
	s.SnBaseBoard = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSnBios(v string) *ListUserDevicesResponseBodyDevices {
	s.SnBios = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSnDiskDrive(v string) *ListUserDevicesResponseBodyDevices {
	s.SnDiskDrive = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSnProcessor(v string) *ListUserDevicesResponseBodyDevices {
	s.SnProcessor = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSnSystem(v string) *ListUserDevicesResponseBodyDevices {
	s.SnSystem = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSrcIP(v string) *ListUserDevicesResponseBodyDevices {
	s.SrcIP = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetUpdateTime(v string) *ListUserDevicesResponseBodyDevices {
	s.UpdateTime = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetUsername(v string) *ListUserDevicesResponseBodyDevices {
	s.Username = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetWorkshop(v string) *ListUserDevicesResponseBodyDevices {
	s.Workshop = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) Validate() error {
	return dara.Validate(s)
}

type ListUserDevicesResponseBodyDevicesNetInterfaceInfo struct {
	Mac  *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUserDevicesResponseBodyDevicesNetInterfaceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUserDevicesResponseBodyDevicesNetInterfaceInfo) GoString() string {
	return s.String()
}

func (s *ListUserDevicesResponseBodyDevicesNetInterfaceInfo) GetMac() *string {
	return s.Mac
}

func (s *ListUserDevicesResponseBodyDevicesNetInterfaceInfo) GetName() *string {
	return s.Name
}

func (s *ListUserDevicesResponseBodyDevicesNetInterfaceInfo) SetMac(v string) *ListUserDevicesResponseBodyDevicesNetInterfaceInfo {
	s.Mac = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevicesNetInterfaceInfo) SetName(v string) *ListUserDevicesResponseBodyDevicesNetInterfaceInfo {
	s.Name = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevicesNetInterfaceInfo) Validate() error {
	return dara.Validate(s)
}
