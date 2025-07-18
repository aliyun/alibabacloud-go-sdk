// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevice(v *GetUserDeviceResponseBodyDevice) *GetUserDeviceResponseBody
	GetDevice() *GetUserDeviceResponseBodyDevice
	SetRequestId(v string) *GetUserDeviceResponseBody
	GetRequestId() *string
}

type GetUserDeviceResponseBody struct {
	Device *GetUserDeviceResponseBodyDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponseBody) GetDevice() *GetUserDeviceResponseBodyDevice {
	return s.Device
}

func (s *GetUserDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserDeviceResponseBody) SetDevice(v *GetUserDeviceResponseBodyDevice) *GetUserDeviceResponseBody {
	s.Device = v
	return s
}

func (s *GetUserDeviceResponseBody) SetRequestId(v string) *GetUserDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserDeviceResponseBodyDevice struct {
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
	// Unauthorized
	DlpStatus    *string                                        `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	EdrStatus    *string                                        `json:"EdrStatus,omitempty" xml:"EdrStatus,omitempty"`
	HistoryUsers []*GetUserDeviceResponseBodyDeviceHistoryUsers `json:"HistoryUsers,omitempty" xml:"HistoryUsers,omitempty" type:"Repeated"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// Disabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// example:
	//
	// 172.16.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// example:
	//
	// 48:9e:XX:XX:02:80
	Mac                 *string   `json:"Mac,omitempty" xml:"Mac,omitempty"`
	MatchDeviceGroupIds []*string `json:"MatchDeviceGroupIds,omitempty" xml:"MatchDeviceGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Unprovisioned
	NacStatus        *string                                            `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	NetInterfaceInfo []*GetUserDeviceResponseBodyDeviceNetInterfaceInfo `json:"NetInterfaceInfo,omitempty" xml:"NetInterfaceInfo,omitempty" type:"Repeated"`
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
	// 106.14.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Workshop   *string `json:"Workshop,omitempty" xml:"Workshop,omitempty"`
}

func (s GetUserDeviceResponseBodyDevice) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceResponseBodyDevice) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponseBodyDevice) GetAppStatus() *string {
	return s.AppStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetUserDeviceResponseBodyDevice) GetCPU() *string {
	return s.CPU
}

func (s *GetUserDeviceResponseBodyDevice) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserDeviceResponseBodyDevice) GetDepartment() *string {
	return s.Department
}

func (s *GetUserDeviceResponseBodyDevice) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *GetUserDeviceResponseBodyDevice) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *GetUserDeviceResponseBodyDevice) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *GetUserDeviceResponseBodyDevice) GetDeviceType() *string {
	return s.DeviceType
}

func (s *GetUserDeviceResponseBodyDevice) GetDeviceVersion() *string {
	return s.DeviceVersion
}

func (s *GetUserDeviceResponseBodyDevice) GetDisk() *string {
	return s.Disk
}

func (s *GetUserDeviceResponseBodyDevice) GetDlpStatus() *string {
	return s.DlpStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetEdrStatus() *string {
	return s.EdrStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetHistoryUsers() []*GetUserDeviceResponseBodyDeviceHistoryUsers {
	return s.HistoryUsers
}

func (s *GetUserDeviceResponseBodyDevice) GetHostname() *string {
	return s.Hostname
}

func (s *GetUserDeviceResponseBodyDevice) GetIaStatus() *string {
	return s.IaStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetInnerIP() *string {
	return s.InnerIP
}

func (s *GetUserDeviceResponseBodyDevice) GetMac() *string {
	return s.Mac
}

func (s *GetUserDeviceResponseBodyDevice) GetMatchDeviceGroupIds() []*string {
	return s.MatchDeviceGroupIds
}

func (s *GetUserDeviceResponseBodyDevice) GetMemory() *string {
	return s.Memory
}

func (s *GetUserDeviceResponseBodyDevice) GetNacStatus() *string {
	return s.NacStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetNetInterfaceInfo() []*GetUserDeviceResponseBodyDeviceNetInterfaceInfo {
	return s.NetInterfaceInfo
}

func (s *GetUserDeviceResponseBodyDevice) GetPaStatus() *string {
	return s.PaStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *GetUserDeviceResponseBodyDevice) GetSharingStatus() *bool {
	return s.SharingStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetSnBaseBoard() *string {
	return s.SnBaseBoard
}

func (s *GetUserDeviceResponseBodyDevice) GetSnBios() *string {
	return s.SnBios
}

func (s *GetUserDeviceResponseBodyDevice) GetSnDiskDrive() *string {
	return s.SnDiskDrive
}

func (s *GetUserDeviceResponseBodyDevice) GetSnProcessor() *string {
	return s.SnProcessor
}

func (s *GetUserDeviceResponseBodyDevice) GetSnSystem() *string {
	return s.SnSystem
}

func (s *GetUserDeviceResponseBodyDevice) GetSrcIP() *string {
	return s.SrcIP
}

func (s *GetUserDeviceResponseBodyDevice) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserDeviceResponseBodyDevice) GetUsername() *string {
	return s.Username
}

func (s *GetUserDeviceResponseBodyDevice) GetWorkshop() *string {
	return s.Workshop
}

func (s *GetUserDeviceResponseBodyDevice) SetAppStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.AppStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetAppVersion(v string) *GetUserDeviceResponseBodyDevice {
	s.AppVersion = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCPU(v string) *GetUserDeviceResponseBodyDevice {
	s.CPU = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCreateTime(v string) *GetUserDeviceResponseBodyDevice {
	s.CreateTime = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDepartment(v string) *GetUserDeviceResponseBodyDevice {
	s.Department = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceBelong(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceBelong = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceModel(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceModel = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceTag(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceTag = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceType(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceType = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceVersion(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceVersion = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDisk(v string) *GetUserDeviceResponseBodyDevice {
	s.Disk = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDlpStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.DlpStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetEdrStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.EdrStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetHistoryUsers(v []*GetUserDeviceResponseBodyDeviceHistoryUsers) *GetUserDeviceResponseBodyDevice {
	s.HistoryUsers = v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetHostname(v string) *GetUserDeviceResponseBodyDevice {
	s.Hostname = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetIaStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.IaStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetInnerIP(v string) *GetUserDeviceResponseBodyDevice {
	s.InnerIP = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetMac(v string) *GetUserDeviceResponseBodyDevice {
	s.Mac = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetMatchDeviceGroupIds(v []*string) *GetUserDeviceResponseBodyDevice {
	s.MatchDeviceGroupIds = v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetMemory(v string) *GetUserDeviceResponseBodyDevice {
	s.Memory = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetNacStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.NacStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetNetInterfaceInfo(v []*GetUserDeviceResponseBodyDeviceNetInterfaceInfo) *GetUserDeviceResponseBodyDevice {
	s.NetInterfaceInfo = v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetPaStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.PaStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSaseUserId(v string) *GetUserDeviceResponseBodyDevice {
	s.SaseUserId = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSharingStatus(v bool) *GetUserDeviceResponseBodyDevice {
	s.SharingStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSnBaseBoard(v string) *GetUserDeviceResponseBodyDevice {
	s.SnBaseBoard = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSnBios(v string) *GetUserDeviceResponseBodyDevice {
	s.SnBios = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSnDiskDrive(v string) *GetUserDeviceResponseBodyDevice {
	s.SnDiskDrive = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSnProcessor(v string) *GetUserDeviceResponseBodyDevice {
	s.SnProcessor = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSnSystem(v string) *GetUserDeviceResponseBodyDevice {
	s.SnSystem = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSrcIP(v string) *GetUserDeviceResponseBodyDevice {
	s.SrcIP = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetUpdateTime(v string) *GetUserDeviceResponseBodyDevice {
	s.UpdateTime = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetUsername(v string) *GetUserDeviceResponseBodyDevice {
	s.Username = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetWorkshop(v string) *GetUserDeviceResponseBodyDevice {
	s.Workshop = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) Validate() error {
	return dara.Validate(s)
}

type GetUserDeviceResponseBodyDeviceHistoryUsers struct {
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserDeviceResponseBodyDeviceHistoryUsers) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceResponseBodyDeviceHistoryUsers) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponseBodyDeviceHistoryUsers) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *GetUserDeviceResponseBodyDeviceHistoryUsers) GetUsername() *string {
	return s.Username
}

func (s *GetUserDeviceResponseBodyDeviceHistoryUsers) SetSaseUserId(v string) *GetUserDeviceResponseBodyDeviceHistoryUsers {
	s.SaseUserId = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceHistoryUsers) SetUsername(v string) *GetUserDeviceResponseBodyDeviceHistoryUsers {
	s.Username = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceHistoryUsers) Validate() error {
	return dara.Validate(s)
}

type GetUserDeviceResponseBodyDeviceNetInterfaceInfo struct {
	Mac  *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserDeviceResponseBodyDeviceNetInterfaceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceResponseBodyDeviceNetInterfaceInfo) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponseBodyDeviceNetInterfaceInfo) GetMac() *string {
	return s.Mac
}

func (s *GetUserDeviceResponseBodyDeviceNetInterfaceInfo) GetName() *string {
	return s.Name
}

func (s *GetUserDeviceResponseBodyDeviceNetInterfaceInfo) SetMac(v string) *GetUserDeviceResponseBodyDeviceNetInterfaceInfo {
	s.Mac = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceNetInterfaceInfo) SetName(v string) *GetUserDeviceResponseBodyDeviceNetInterfaceInfo {
	s.Name = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceNetInterfaceInfo) Validate() error {
	return dara.Validate(s)
}
