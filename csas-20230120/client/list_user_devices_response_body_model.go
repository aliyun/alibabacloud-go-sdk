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
	// The list of endpoint devices.
	Devices []*ListUserDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of endpoint devices.
	//
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

type ListUserDevicesResponseBodyDevices struct {
	// The client status. Valid values:
	//
	// - **Online**: online.
	//
	// - **Offline**: offline.
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
	// The auto-logon status of the client.
	//
	// example:
	//
	// Enabled
	AutoLoginStatus *string `json:"AutoLoginStatus,omitempty" xml:"AutoLoginStatus,omitempty"`
	// The CPU model of the endpoint device.
	//
	// example:
	//
	// Apple M1
	CPU  *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
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
	CityZh    *string `json:"CityZh,omitempty" xml:"CityZh,omitempty"`
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
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
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
	// The registration time of the endpoint device.
	//
	// example:
	//
	// 2023-07-17 18:46:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The department to which the user belongs.
	//
	// example:
	//
	// 测试部
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The ownership of the endpoint device. Valid values:
	//
	// - **Personal**: personal device.
	//
	// - **Company**: company device.
	//
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// The model of the endpoint device.
	//
	// example:
	//
	// MacBookPro17,1
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// The status of the endpoint device. Valid values:
	//
	// - **Online**: online.
	//
	// - **Offline**: offline.
	//
	// - **LongTermOffline**: long-term offline.
	//
	// - **Locked**: locked.
	//
	// - **Lost**: reported as lost.
	//
	// - **Unbound**: unbound.
	//
	// example:
	//
	// Online
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// The endpoint device ID.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// The operating system type of the endpoint device. Valid values:
	//
	// - **Windows**: Windows.
	//
	// - **macOS**: macOS.
	//
	// - **Linux**: Linux.
	//
	// - **Android**: Android.
	//
	// - **iOS**: iOS.
	//
	// - **Windows_Wuying**: WUYING Workspace.
	//
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The operating system version of the endpoint device.
	//
	// example:
	//
	// 3.5.1
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	// The disk model of the endpoint device.
	//
	// example:
	//
	// APPLE SSD AP0512Q Media
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// The office data protection status. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// - **Unprovisioned**: not configured.
	//
	// - **Unauthorized**: unauthorized.
	//
	// example:
	//
	// Enabled
	DlpStatus *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	// The anti-intrusion status. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// example:
	//
	// Enabled
	EdrStatus *string `json:"EdrStatus,omitempty" xml:"EdrStatus,omitempty"`
	// The list of full department paths.
	FullDepartment []*string `json:"FullDepartment,omitempty" xml:"FullDepartment,omitempty" type:"Repeated"`
	// The name of the endpoint device.
	//
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The Internet access status. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// - **Unprovisioned**: not configured.
	//
	// example:
	//
	// Enabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// The internal IP address of the endpoint device.
	//
	// example:
	//
	// 192.168.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// Indicates whether the device is joined to an Active Directory (AD) domain.
	//
	// example:
	//
	// true
	JoinAdDomain *bool `json:"JoinAdDomain,omitempty" xml:"JoinAdDomain,omitempty"`
	// The MAC address of the endpoint device.
	//
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The collection of matched device group IDs.
	MatchDeviceGroupIds []*string `json:"MatchDeviceGroupIds,omitempty" xml:"MatchDeviceGroupIds,omitempty" type:"Repeated"`
	// The memory capacity of the endpoint device. Unit: GB.
	//
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network access control status. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// - **Unprovisioned**: not configured.
	//
	// example:
	//
	// Enabled
	NacStatus *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	// The list of network interfaces of the endpoint device.
	NetInterfaceInfo []*ListUserDevicesResponseBodyDevicesNetInterfaceInfo `json:"NetInterfaceInfo,omitempty" xml:"NetInterfaceInfo,omitempty" type:"Repeated"`
	// The private access status. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// - **Unprovisioned**: not configured.
	//
	// example:
	//
	// Enabled
	PaStatus *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
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
	// Specifies whether sharing is enabled for the device. Valid values:
	//
	// - **true**: Sharing is enabled.
	//
	// - **false**: Sharing is disabled.
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
	// The BIOS system serial number.
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
	// The processor serial number.
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
	// The logon IP address of the endpoint device.
	//
	// example:
	//
	// 11.49.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The last online time of the endpoint device.
	//
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The username.
	//
	// example:
	//
	// 王先生
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the office area.
	//
	// example:
	//
	// 测试办公区
	Workshop *string `json:"Workshop,omitempty" xml:"Workshop,omitempty"`
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

func (s *ListUserDevicesResponseBodyDevices) GetCity() *string {
	return s.City
}

func (s *ListUserDevicesResponseBodyDevices) GetCityEn() *string {
	return s.CityEn
}

func (s *ListUserDevicesResponseBodyDevices) GetCityZh() *string {
	return s.CityZh
}

func (s *ListUserDevicesResponseBodyDevices) GetContinent() *string {
	return s.Continent
}

func (s *ListUserDevicesResponseBodyDevices) GetContinentEn() *string {
	return s.ContinentEn
}

func (s *ListUserDevicesResponseBodyDevices) GetContinentZh() *string {
	return s.ContinentZh
}

func (s *ListUserDevicesResponseBodyDevices) GetCountry() *string {
	return s.Country
}

func (s *ListUserDevicesResponseBodyDevices) GetCountryEn() *string {
	return s.CountryEn
}

func (s *ListUserDevicesResponseBodyDevices) GetCountryZh() *string {
	return s.CountryZh
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

func (s *ListUserDevicesResponseBodyDevices) GetFullDepartment() []*string {
	return s.FullDepartment
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

func (s *ListUserDevicesResponseBodyDevices) GetJoinAdDomain() *bool {
	return s.JoinAdDomain
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

func (s *ListUserDevicesResponseBodyDevices) GetProvince() *string {
	return s.Province
}

func (s *ListUserDevicesResponseBodyDevices) GetProvinceEn() *string {
	return s.ProvinceEn
}

func (s *ListUserDevicesResponseBodyDevices) GetProvinceZh() *string {
	return s.ProvinceZh
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

func (s *ListUserDevicesResponseBodyDevices) SetCity(v string) *ListUserDevicesResponseBodyDevices {
	s.City = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCityEn(v string) *ListUserDevicesResponseBodyDevices {
	s.CityEn = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCityZh(v string) *ListUserDevicesResponseBodyDevices {
	s.CityZh = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetContinent(v string) *ListUserDevicesResponseBodyDevices {
	s.Continent = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetContinentEn(v string) *ListUserDevicesResponseBodyDevices {
	s.ContinentEn = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetContinentZh(v string) *ListUserDevicesResponseBodyDevices {
	s.ContinentZh = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCountry(v string) *ListUserDevicesResponseBodyDevices {
	s.Country = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCountryEn(v string) *ListUserDevicesResponseBodyDevices {
	s.CountryEn = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCountryZh(v string) *ListUserDevicesResponseBodyDevices {
	s.CountryZh = &v
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

func (s *ListUserDevicesResponseBodyDevices) SetFullDepartment(v []*string) *ListUserDevicesResponseBodyDevices {
	s.FullDepartment = v
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

func (s *ListUserDevicesResponseBodyDevices) SetJoinAdDomain(v bool) *ListUserDevicesResponseBodyDevices {
	s.JoinAdDomain = &v
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

func (s *ListUserDevicesResponseBodyDevices) SetProvince(v string) *ListUserDevicesResponseBodyDevices {
	s.Province = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetProvinceEn(v string) *ListUserDevicesResponseBodyDevices {
	s.ProvinceEn = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetProvinceZh(v string) *ListUserDevicesResponseBodyDevices {
	s.ProvinceZh = &v
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
	if s.NetInterfaceInfo != nil {
		for _, item := range s.NetInterfaceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserDevicesResponseBodyDevicesNetInterfaceInfo struct {
	// The MAC address of the network interface.
	//
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The name of the network interface.
	//
	// example:
	//
	// eth0
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
