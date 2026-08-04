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
	// The terminal device.
	Device *GetUserDeviceResponseBodyDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	// The ID of the request.
	//
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
	if s.Device != nil {
		if err := s.Device.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserDeviceResponseBodyDevice struct {
	// The client status. Valid values:
	//
	// example:
	//
	// Online
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// The client version.
	//
	// example:
	//
	// 2.2.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The auto-logon status of the device. Valid values:
	//
	// example:
	//
	// Enabled
	AutoLoginStatus *string `json:"AutoLoginStatus,omitempty" xml:"AutoLoginStatus,omitempty"`
	// The battery health percentage.
	//
	// example:
	//
	// 100
	BatteryHealthPercentage *int32 `json:"BatteryHealthPercentage,omitempty" xml:"BatteryHealthPercentage,omitempty"`
	// The battery remaining charge percentage.
	//
	// example:
	//
	// 90
	BatteryRemainingPercentage *int32 `json:"BatteryRemainingPercentage,omitempty" xml:"BatteryRemainingPercentage,omitempty"`
	// The CPU model of the terminal device.
	//
	// example:
	//
	// Apple M1
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The city to which the public IP address belongs.
	//
	// example:
	//
	// Hangzhou City
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The city name in English.
	//
	// example:
	//
	// Beijing City
	CityEn *string `json:"CityEn,omitempty" xml:"CityEn,omitempty"`
	// The city name in Chinese.
	//
	// example:
	//
	// 北京市
	CityZh *string `json:"CityZh,omitempty" xml:"CityZh,omitempty"`
	// The continent to which the public IP address belongs.
	//
	// example:
	//
	// Asia
	Continent *string `json:"Continent,omitempty" xml:"Continent,omitempty"`
	// The continent name in English.
	//
	// example:
	//
	// Asia
	ContinentEn *string `json:"ContinentEn,omitempty" xml:"ContinentEn,omitempty"`
	// The continent name in Chinese.
	//
	// example:
	//
	// 亚洲
	ContinentZh *string `json:"ContinentZh,omitempty" xml:"ContinentZh,omitempty"`
	// The country to which the public IP address belongs.
	//
	// example:
	//
	// China
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The country name in English.
	//
	// example:
	//
	// China
	CountryEn *string `json:"CountryEn,omitempty" xml:"CountryEn,omitempty"`
	// The country name in Chinese.
	//
	// example:
	//
	// 中国
	CountryZh *string `json:"CountryZh,omitempty" xml:"CountryZh,omitempty"`
	// The registration time of the terminal device.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The department to which the user belongs.
	//
	// example:
	//
	// QA Department
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The ownership of the terminal device. Valid values:
	//
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// The model of the terminal device.
	//
	// example:
	//
	// MacBookPro17,1
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// The status of the terminal device. Valid values:
	//
	// example:
	//
	// Online
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// The ID of the terminal device.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// The operating system type of the terminal device. Valid values:
	//
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The operating system version of the terminal device.
	//
	// example:
	//
	// 3.5.1
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	// The disk model of the terminal device.
	//
	// example:
	//
	// APPLE SSD AP0512Q Media
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// The available disk space, in GB.
	//
	// example:
	//
	// 200
	DiskAvailable *int32 `json:"DiskAvailable,omitempty" xml:"DiskAvailable,omitempty"`
	// The used disk space, in GB.
	//
	// example:
	//
	// 103
	DiskUsed *int32 `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	// The office data protection status. Valid values:
	//
	// example:
	//
	// Unauthorized
	DlpStatus *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	// The anti-intrusion status. Valid values:
	//
	// example:
	//
	// Disabled
	EdrStatus *string `json:"EdrStatus,omitempty" xml:"EdrStatus,omitempty"`
	// The historical users of the terminal device.
	HistoryUsers []*GetUserDeviceResponseBodyDeviceHistoryUsers `json:"HistoryUsers,omitempty" xml:"HistoryUsers,omitempty" type:"Repeated"`
	// The name of the terminal device.
	//
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The Internet access status. Valid values:
	//
	// example:
	//
	// Disabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// The internal IP address of the terminal device.
	//
	// example:
	//
	// 172.16.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// Indicates whether the terminal has joined an AD domain.
	//
	// example:
	//
	// true
	JoinAdDomain *bool `json:"JoinAdDomain,omitempty" xml:"JoinAdDomain,omitempty"`
	// The MAC address of the terminal device.
	//
	// example:
	//
	// 48:9e:XX:XX:02:80
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The IDs of matched device groups.
	MatchDeviceGroupIds []*string `json:"MatchDeviceGroupIds,omitempty" xml:"MatchDeviceGroupIds,omitempty" type:"Repeated"`
	// The memory capacity of the terminal device. Unit: GB.
	//
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network access control status. Valid values:
	//
	// example:
	//
	// Unprovisioned
	NacStatus *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	// The list of network interface controllers (NICs) of the terminal device.
	NetInterfaceInfo []*GetUserDeviceResponseBodyDeviceNetInterfaceInfo `json:"NetInterfaceInfo,omitempty" xml:"NetInterfaceInfo,omitempty" type:"Repeated"`
	// The private access status. Valid values:
	//
	// example:
	//
	// Enabled
	PaStatus *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	// The list of processes running on the terminal.
	Processes []*GetUserDeviceResponseBodyDeviceProcesses `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	// The province to which the public IP address belongs.
	//
	// example:
	//
	// Zhejiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The province name in English.
	//
	// example:
	//
	// Beijing
	ProvinceEn *string `json:"ProvinceEn,omitempty" xml:"ProvinceEn,omitempty"`
	// The province name in Chinese.
	//
	// example:
	//
	// 北京市
	ProvinceZh *string `json:"ProvinceZh,omitempty" xml:"ProvinceZh,omitempty"`
	// The user ID.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// Indicates whether sharing is enabled for the device. Valid values:
	//
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	// The motherboard serial number.
	//
	// example:
	//
	// PJGGU00WBD****
	SnBaseBoard *string `json:"SnBaseBoard,omitempty" xml:"SnBaseBoard,omitempty"`
	// The serial number (SN) of the BIOS system.
	//
	// example:
	//
	// 5CG003****
	SnBios *string `json:"SnBios,omitempty" xml:"SnBios,omitempty"`
	// The hard disk serial number.
	//
	// example:
	//
	// WD-WXR1A99A****
	SnDiskDrive *string `json:"SnDiskDrive,omitempty" xml:"SnDiskDrive,omitempty"`
	// The serial number (SN) of the processor.
	//
	// example:
	//
	// BFEBFBFF0008****
	SnProcessor *string `json:"SnProcessor,omitempty" xml:"SnProcessor,omitempty"`
	// The system serial number.
	//
	// example:
	//
	// KVN9C9****
	SnSystem *string `json:"SnSystem,omitempty" xml:"SnSystem,omitempty"`
	// The logon IP address of the terminal device.
	//
	// example:
	//
	// 106.14.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The timestamp when the terminal process information was collected.
	//
	// example:
	//
	// 1775096714
	TerminalInfoCollectTime *int64 `json:"TerminalInfoCollectTime,omitempty" xml:"TerminalInfoCollectTime,omitempty"`
	// The last online time of the terminal device.
	//
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The username.
	//
	// example:
	//
	// Mr. Wang
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the office area.
	//
	// example:
	//
	// Test Office Area
	Workshop *string `json:"Workshop,omitempty" xml:"Workshop,omitempty"`
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

func (s *GetUserDeviceResponseBodyDevice) GetAutoLoginStatus() *string {
	return s.AutoLoginStatus
}

func (s *GetUserDeviceResponseBodyDevice) GetBatteryHealthPercentage() *int32 {
	return s.BatteryHealthPercentage
}

func (s *GetUserDeviceResponseBodyDevice) GetBatteryRemainingPercentage() *int32 {
	return s.BatteryRemainingPercentage
}

func (s *GetUserDeviceResponseBodyDevice) GetCPU() *string {
	return s.CPU
}

func (s *GetUserDeviceResponseBodyDevice) GetCity() *string {
	return s.City
}

func (s *GetUserDeviceResponseBodyDevice) GetCityEn() *string {
	return s.CityEn
}

func (s *GetUserDeviceResponseBodyDevice) GetCityZh() *string {
	return s.CityZh
}

func (s *GetUserDeviceResponseBodyDevice) GetContinent() *string {
	return s.Continent
}

func (s *GetUserDeviceResponseBodyDevice) GetContinentEn() *string {
	return s.ContinentEn
}

func (s *GetUserDeviceResponseBodyDevice) GetContinentZh() *string {
	return s.ContinentZh
}

func (s *GetUserDeviceResponseBodyDevice) GetCountry() *string {
	return s.Country
}

func (s *GetUserDeviceResponseBodyDevice) GetCountryEn() *string {
	return s.CountryEn
}

func (s *GetUserDeviceResponseBodyDevice) GetCountryZh() *string {
	return s.CountryZh
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

func (s *GetUserDeviceResponseBodyDevice) GetDiskAvailable() *int32 {
	return s.DiskAvailable
}

func (s *GetUserDeviceResponseBodyDevice) GetDiskUsed() *int32 {
	return s.DiskUsed
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

func (s *GetUserDeviceResponseBodyDevice) GetJoinAdDomain() *bool {
	return s.JoinAdDomain
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

func (s *GetUserDeviceResponseBodyDevice) GetProcesses() []*GetUserDeviceResponseBodyDeviceProcesses {
	return s.Processes
}

func (s *GetUserDeviceResponseBodyDevice) GetProvince() *string {
	return s.Province
}

func (s *GetUserDeviceResponseBodyDevice) GetProvinceEn() *string {
	return s.ProvinceEn
}

func (s *GetUserDeviceResponseBodyDevice) GetProvinceZh() *string {
	return s.ProvinceZh
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

func (s *GetUserDeviceResponseBodyDevice) GetTerminalInfoCollectTime() *int64 {
	return s.TerminalInfoCollectTime
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

func (s *GetUserDeviceResponseBodyDevice) SetAutoLoginStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.AutoLoginStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetBatteryHealthPercentage(v int32) *GetUserDeviceResponseBodyDevice {
	s.BatteryHealthPercentage = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetBatteryRemainingPercentage(v int32) *GetUserDeviceResponseBodyDevice {
	s.BatteryRemainingPercentage = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCPU(v string) *GetUserDeviceResponseBodyDevice {
	s.CPU = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCity(v string) *GetUserDeviceResponseBodyDevice {
	s.City = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCityEn(v string) *GetUserDeviceResponseBodyDevice {
	s.CityEn = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCityZh(v string) *GetUserDeviceResponseBodyDevice {
	s.CityZh = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetContinent(v string) *GetUserDeviceResponseBodyDevice {
	s.Continent = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetContinentEn(v string) *GetUserDeviceResponseBodyDevice {
	s.ContinentEn = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetContinentZh(v string) *GetUserDeviceResponseBodyDevice {
	s.ContinentZh = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCountry(v string) *GetUserDeviceResponseBodyDevice {
	s.Country = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCountryEn(v string) *GetUserDeviceResponseBodyDevice {
	s.CountryEn = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCountryZh(v string) *GetUserDeviceResponseBodyDevice {
	s.CountryZh = &v
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

func (s *GetUserDeviceResponseBodyDevice) SetDiskAvailable(v int32) *GetUserDeviceResponseBodyDevice {
	s.DiskAvailable = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDiskUsed(v int32) *GetUserDeviceResponseBodyDevice {
	s.DiskUsed = &v
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

func (s *GetUserDeviceResponseBodyDevice) SetJoinAdDomain(v bool) *GetUserDeviceResponseBodyDevice {
	s.JoinAdDomain = &v
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

func (s *GetUserDeviceResponseBodyDevice) SetProcesses(v []*GetUserDeviceResponseBodyDeviceProcesses) *GetUserDeviceResponseBodyDevice {
	s.Processes = v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetProvince(v string) *GetUserDeviceResponseBodyDevice {
	s.Province = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetProvinceEn(v string) *GetUserDeviceResponseBodyDevice {
	s.ProvinceEn = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetProvinceZh(v string) *GetUserDeviceResponseBodyDevice {
	s.ProvinceZh = &v
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

func (s *GetUserDeviceResponseBodyDevice) SetTerminalInfoCollectTime(v int64) *GetUserDeviceResponseBodyDevice {
	s.TerminalInfoCollectTime = &v
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
	if s.HistoryUsers != nil {
		for _, item := range s.HistoryUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetInterfaceInfo != nil {
		for _, item := range s.NetInterfaceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Processes != nil {
		for _, item := range s.Processes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserDeviceResponseBodyDeviceHistoryUsers struct {
	// The user ID.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// The username.
	//
	// example:
	//
	// Ms. Zhang
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	// The MAC address of the NIC.
	//
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The name of the NIC.
	//
	// example:
	//
	// eth0
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

type GetUserDeviceResponseBodyDeviceProcesses struct {
	// The CPU usage percentage of the process.
	//
	// example:
	//
	// 0.05
	Cpu *float64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The process running description.
	//
	// example:
	//
	// C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The memory usage of the process, in MB.
	//
	// example:
	//
	// 233
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The process name.
	//
	// example:
	//
	// chrome.exe
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserDeviceResponseBodyDeviceProcesses) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceResponseBodyDeviceProcesses) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) GetCpu() *float64 {
	return s.Cpu
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) GetDescription() *string {
	return s.Description
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) GetMemory() *int32 {
	return s.Memory
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) GetName() *string {
	return s.Name
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) SetCpu(v float64) *GetUserDeviceResponseBodyDeviceProcesses {
	s.Cpu = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) SetDescription(v string) *GetUserDeviceResponseBodyDeviceProcesses {
	s.Description = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) SetMemory(v int32) *GetUserDeviceResponseBodyDeviceProcesses {
	s.Memory = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) SetName(v string) *GetUserDeviceResponseBodyDeviceProcesses {
	s.Name = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceProcesses) Validate() error {
	return dara.Validate(s)
}
