// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppStatuses(v []*string) *ListUserDevicesRequest
	GetAppStatuses() []*string
	SetAppVersions(v []*string) *ListUserDevicesRequest
	GetAppVersions() []*string
	SetAutoLoginStatuses(v []*string) *ListUserDevicesRequest
	GetAutoLoginStatuses() []*string
	SetCurrentPage(v int64) *ListUserDevicesRequest
	GetCurrentPage() *int64
	SetDepartment(v string) *ListUserDevicesRequest
	GetDepartment() *string
	SetDeviceBelong(v string) *ListUserDevicesRequest
	GetDeviceBelong() *string
	SetDeviceGroupId(v string) *ListUserDevicesRequest
	GetDeviceGroupId() *string
	SetDeviceStatuses(v []*string) *ListUserDevicesRequest
	GetDeviceStatuses() []*string
	SetDeviceTags(v []*string) *ListUserDevicesRequest
	GetDeviceTags() []*string
	SetDeviceTypes(v []*string) *ListUserDevicesRequest
	GetDeviceTypes() []*string
	SetDiskEncryptStatuses(v []*string) *ListUserDevicesRequest
	GetDiskEncryptStatuses() []*string
	SetDlpStatuses(v []*string) *ListUserDevicesRequest
	GetDlpStatuses() []*string
	SetHostname(v string) *ListUserDevicesRequest
	GetHostname() *string
	SetIaStatuses(v []*string) *ListUserDevicesRequest
	GetIaStatuses() []*string
	SetInnerIp(v string) *ListUserDevicesRequest
	GetInnerIp() *string
	SetMac(v string) *ListUserDevicesRequest
	GetMac() *string
	SetNacStatuses(v []*string) *ListUserDevicesRequest
	GetNacStatuses() []*string
	SetPaStatuses(v []*string) *ListUserDevicesRequest
	GetPaStatuses() []*string
	SetPageSize(v int64) *ListUserDevicesRequest
	GetPageSize() *int64
	SetSaseUserId(v string) *ListUserDevicesRequest
	GetSaseUserId() *string
	SetSharingStatus(v bool) *ListUserDevicesRequest
	GetSharingStatus() *bool
	SetSnBios(v string) *ListUserDevicesRequest
	GetSnBios() *string
	SetSnSystem(v string) *ListUserDevicesRequest
	GetSnSystem() *string
	SetSortBy(v string) *ListUserDevicesRequest
	GetSortBy() *string
	SetUsername(v string) *ListUserDevicesRequest
	GetUsername() *string
	SetWorkshop(v string) *ListUserDevicesRequest
	GetWorkshop() *string
}

type ListUserDevicesRequest struct {
	// The client statuses.
	AppStatuses []*string `json:"AppStatuses,omitempty" xml:"AppStatuses,omitempty" type:"Repeated"`
	// The client versions.
	AppVersions []*string `json:"AppVersions,omitempty" xml:"AppVersions,omitempty" type:"Repeated"`
	// Filters the results by the automatic logon status of the user terminal devices.
	AutoLoginStatuses []*string `json:"AutoLoginStatuses,omitempty" xml:"AutoLoginStatuses,omitempty" type:"Repeated"`
	// The page number of the current page for paging in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The department to which the user belongs. The value must be 1 to 128 characters in length, and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), commas (,), semicolons (;), hyphens (-), underscores (_), forward slashes (/), at signs (@), and spaces.
	//
	// example:
	//
	// Test Department
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The ownership of the terminal device. Valid values:
	//
	// - **Personal**: Personal device.
	//
	// - **Company**: Corporate device.
	//
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// The device group ID.
	//
	// example:
	//
	// device-group-5191cf830a5e****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// The terminal device statuses.
	DeviceStatuses []*string `json:"DeviceStatuses,omitempty" xml:"DeviceStatuses,omitempty" type:"Repeated"`
	// The terminal device IDs.
	DeviceTags []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
	// The terminal device operating system types.
	DeviceTypes []*string `json:"DeviceTypes,omitempty" xml:"DeviceTypes,omitempty" type:"Repeated"`
	// Filters the results by the disk encryption status of the user terminal devices. You can select multiple statuses, which are evaluated with an OR logical operator. If you select Unencrypted, the result includes user terminal devices that have not reported their disk encryption status.
	DiskEncryptStatuses []*string `json:"DiskEncryptStatuses,omitempty" xml:"DiskEncryptStatuses,omitempty" type:"Repeated"`
	// The office data protection statuses.
	DlpStatuses []*string `json:"DlpStatuses,omitempty" xml:"DlpStatuses,omitempty" type:"Repeated"`
	// The name of the terminal device. The value must be 1 to 128 characters in length, and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), commas (,), semicolons (;), hyphens (-), underscores (_), forward slashes (/), at signs (@), and spaces. If you enter only an underscore (_), the system also queries all terminal devices whose names contain 4-byte UTF-8 characters.
	//
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The Internet access statuses.
	IaStatuses []*string `json:"IaStatuses,omitempty" xml:"IaStatuses,omitempty" type:"Repeated"`
	// The internal IP address of the terminal device.
	//
	// example:
	//
	// 192.168.XX.XX
	InnerIp *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// The MAC address of the terminal device.
	//
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The network access control statuses.
	NacStatuses []*string `json:"NacStatuses,omitempty" xml:"NacStatuses,omitempty" type:"Repeated"`
	// The internal network access statuses.
	PaStatuses []*string `json:"PaStatuses,omitempty" xml:"PaStatuses,omitempty" type:"Repeated"`
	// The number of entries per page in a paged query. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user ID. You can obtain the value from the following operations:
	//
	// - [GetUserDevice](~~GetUserDevice~~): Queries the details of a user terminal device.
	//
	// - [ListUserDevices](~~ListUserDevices~~): Lists user terminal devices.
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
	// The BIOS system serial number.
	//
	// example:
	//
	// 5CG003****
	SnBios *string `json:"SnBios,omitempty" xml:"SnBios,omitempty"`
	// The system serial number.
	//
	// example:
	//
	// KVN9C9****
	SnSystem *string `json:"SnSystem,omitempty" xml:"SnSystem,omitempty"`
	// The sorting parameter. Valid values:
	//
	// - **Username**: Sorts by Username in ascending order.
	//
	// - **AppVersion**: Sorts by AppVersion in descending order.
	//
	// - **UpdateTime**: Sorts by UpdateTime in descending order.
	//
	// - **CreateTime**: Sorts by CreateTime in descending order.
	//
	// example:
	//
	// UpdateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The username. The value must be 1 to 128 characters in length, and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), hyphens (-), asterisks (*), at signs (@), and spaces.
	//
	// example:
	//
	// Mr. Wang
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the office area.
	//
	// example:
	//
	// Test office area
	Workshop *string `json:"Workshop,omitempty" xml:"Workshop,omitempty"`
}

func (s ListUserDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListUserDevicesRequest) GetAppStatuses() []*string {
	return s.AppStatuses
}

func (s *ListUserDevicesRequest) GetAppVersions() []*string {
	return s.AppVersions
}

func (s *ListUserDevicesRequest) GetAutoLoginStatuses() []*string {
	return s.AutoLoginStatuses
}

func (s *ListUserDevicesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListUserDevicesRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListUserDevicesRequest) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *ListUserDevicesRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListUserDevicesRequest) GetDeviceStatuses() []*string {
	return s.DeviceStatuses
}

func (s *ListUserDevicesRequest) GetDeviceTags() []*string {
	return s.DeviceTags
}

func (s *ListUserDevicesRequest) GetDeviceTypes() []*string {
	return s.DeviceTypes
}

func (s *ListUserDevicesRequest) GetDiskEncryptStatuses() []*string {
	return s.DiskEncryptStatuses
}

func (s *ListUserDevicesRequest) GetDlpStatuses() []*string {
	return s.DlpStatuses
}

func (s *ListUserDevicesRequest) GetHostname() *string {
	return s.Hostname
}

func (s *ListUserDevicesRequest) GetIaStatuses() []*string {
	return s.IaStatuses
}

func (s *ListUserDevicesRequest) GetInnerIp() *string {
	return s.InnerIp
}

func (s *ListUserDevicesRequest) GetMac() *string {
	return s.Mac
}

func (s *ListUserDevicesRequest) GetNacStatuses() []*string {
	return s.NacStatuses
}

func (s *ListUserDevicesRequest) GetPaStatuses() []*string {
	return s.PaStatuses
}

func (s *ListUserDevicesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserDevicesRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListUserDevicesRequest) GetSharingStatus() *bool {
	return s.SharingStatus
}

func (s *ListUserDevicesRequest) GetSnBios() *string {
	return s.SnBios
}

func (s *ListUserDevicesRequest) GetSnSystem() *string {
	return s.SnSystem
}

func (s *ListUserDevicesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListUserDevicesRequest) GetUsername() *string {
	return s.Username
}

func (s *ListUserDevicesRequest) GetWorkshop() *string {
	return s.Workshop
}

func (s *ListUserDevicesRequest) SetAppStatuses(v []*string) *ListUserDevicesRequest {
	s.AppStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetAppVersions(v []*string) *ListUserDevicesRequest {
	s.AppVersions = v
	return s
}

func (s *ListUserDevicesRequest) SetAutoLoginStatuses(v []*string) *ListUserDevicesRequest {
	s.AutoLoginStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetCurrentPage(v int64) *ListUserDevicesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserDevicesRequest) SetDepartment(v string) *ListUserDevicesRequest {
	s.Department = &v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceBelong(v string) *ListUserDevicesRequest {
	s.DeviceBelong = &v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceGroupId(v string) *ListUserDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceStatuses(v []*string) *ListUserDevicesRequest {
	s.DeviceStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceTags(v []*string) *ListUserDevicesRequest {
	s.DeviceTags = v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceTypes(v []*string) *ListUserDevicesRequest {
	s.DeviceTypes = v
	return s
}

func (s *ListUserDevicesRequest) SetDiskEncryptStatuses(v []*string) *ListUserDevicesRequest {
	s.DiskEncryptStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetDlpStatuses(v []*string) *ListUserDevicesRequest {
	s.DlpStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetHostname(v string) *ListUserDevicesRequest {
	s.Hostname = &v
	return s
}

func (s *ListUserDevicesRequest) SetIaStatuses(v []*string) *ListUserDevicesRequest {
	s.IaStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetInnerIp(v string) *ListUserDevicesRequest {
	s.InnerIp = &v
	return s
}

func (s *ListUserDevicesRequest) SetMac(v string) *ListUserDevicesRequest {
	s.Mac = &v
	return s
}

func (s *ListUserDevicesRequest) SetNacStatuses(v []*string) *ListUserDevicesRequest {
	s.NacStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetPaStatuses(v []*string) *ListUserDevicesRequest {
	s.PaStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetPageSize(v int64) *ListUserDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserDevicesRequest) SetSaseUserId(v string) *ListUserDevicesRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListUserDevicesRequest) SetSharingStatus(v bool) *ListUserDevicesRequest {
	s.SharingStatus = &v
	return s
}

func (s *ListUserDevicesRequest) SetSnBios(v string) *ListUserDevicesRequest {
	s.SnBios = &v
	return s
}

func (s *ListUserDevicesRequest) SetSnSystem(v string) *ListUserDevicesRequest {
	s.SnSystem = &v
	return s
}

func (s *ListUserDevicesRequest) SetSortBy(v string) *ListUserDevicesRequest {
	s.SortBy = &v
	return s
}

func (s *ListUserDevicesRequest) SetUsername(v string) *ListUserDevicesRequest {
	s.Username = &v
	return s
}

func (s *ListUserDevicesRequest) SetWorkshop(v string) *ListUserDevicesRequest {
	s.Workshop = &v
	return s
}

func (s *ListUserDevicesRequest) Validate() error {
	return dara.Validate(s)
}
