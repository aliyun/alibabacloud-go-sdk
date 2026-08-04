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
	// A list of endpoint devices.
	Devices []*UpdateUserDevicesStatusResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// The ID of this request.
	//
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

type UpdateUserDevicesStatusResponseBodyDevices struct {
	// The client status. Valid values:
	//
	// - **Online**: Online.
	//
	// - **Offline**: Offline.
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
	// The CPU model of the endpoint device.
	//
	// example:
	//
	// Apple M1
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The time when the endpoint device was registered.
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
	// - **Personal**: Personal device.
	//
	// - **Company**: Company device.
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
	// - **Online**: Online.
	//
	// - **Offline**: Offline.
	//
	// - **LongTermOffline**: Long-term offline.
	//
	// - **Locked**: Locked.
	//
	// - **Lost**: Reported as lost.
	//
	// - **Unbound**: Detached.
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
	// - **Windows_Wuying**: Alibaba Cloud Cloud Desktop.
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
	// The data protection status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// - **Unprovisioned**: Not configured.
	//
	// - **Unauthorized**: Unauthorized.
	//
	// example:
	//
	// Enabled
	DlpStatus *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	// The device name.
	//
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Internet access status. Valid values:
	//
	// - **Enabled**: Internet access is enabled.
	//
	// - **Disabled**: Internet access is disabled.
	//
	// - **Unprovisioned**: The device is unconfigured.
	//
	// example:
	//
	// Enabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// The private network IP address of the endpoint device.
	//
	// example:
	//
	// 192.168.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// The MAC address of the endpoint device.
	//
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The memory capacity of the endpoint device, in GB.
	//
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network admission control status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// - **Unprovisioned**: Not configured.
	//
	// example:
	//
	// Enabled
	NacStatus *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	// A list of network interface controllers (NICs) on the endpoint device.
	NetInterfaceInfo []*UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo `json:"NetInterfaceInfo,omitempty" xml:"NetInterfaceInfo,omitempty" type:"Repeated"`
	// The private network access status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// - **Unprovisioned**: Not configured.
	//
	// example:
	//
	// Enabled
	PaStatus *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	// The user ID.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// Whether sharing is enabled for the device. Valid values:
	//
	// - **true**: Sharing is enabled.
	//
	// - **false**: Sharing is disabled.
	//
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	// The IP address used to log on to the endpoint device.
	//
	// example:
	//
	// 11.49.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The last time the endpoint device was online.
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

type UpdateUserDevicesStatusResponseBodyDevicesNetInterfaceInfo struct {
	// The MAC address of the NIC.
	//
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The NIC name.
	//
	// example:
	//
	// eth0
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
