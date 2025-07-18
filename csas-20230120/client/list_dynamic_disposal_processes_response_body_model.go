// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicDisposalProcessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisposalProcesses(v []*ListDynamicDisposalProcessesResponseBodyDisposalProcesses) *ListDynamicDisposalProcessesResponseBody
	GetDisposalProcesses() []*ListDynamicDisposalProcessesResponseBodyDisposalProcesses
	SetRequestId(v string) *ListDynamicDisposalProcessesResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListDynamicDisposalProcessesResponseBody
	GetTotalNum() *int32
}

type ListDynamicDisposalProcessesResponseBody struct {
	// List of disposal processes.
	DisposalProcesses []*ListDynamicDisposalProcessesResponseBodyDisposalProcesses `json:"DisposalProcesses,omitempty" xml:"DisposalProcesses,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of dynamic disposal processes.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDynamicDisposalProcessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicDisposalProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicDisposalProcessesResponseBody) GetDisposalProcesses() []*ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	return s.DisposalProcesses
}

func (s *ListDynamicDisposalProcessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDynamicDisposalProcessesResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListDynamicDisposalProcessesResponseBody) SetDisposalProcesses(v []*ListDynamicDisposalProcessesResponseBodyDisposalProcesses) *ListDynamicDisposalProcessesResponseBody {
	s.DisposalProcesses = v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBody) SetRequestId(v string) *ListDynamicDisposalProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBody) SetTotalNum(v int32) *ListDynamicDisposalProcessesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDynamicDisposalProcessesResponseBodyDisposalProcesses struct {
	// User\\"s department.
	//
	// example:
	//
	// IT
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// Device ID.
	//
	// example:
	//
	// FD7554AD-4CDE-6359-6B49-4FE950606C2C
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// Basic device information.
	DeviceBasicInfo *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo `json:"DeviceBasicInfo,omitempty" xml:"DeviceBasicInfo,omitempty" type:"Struct"`
	// 设备状态信息。
	DeviceStatusInfo *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo `json:"DeviceStatusInfo,omitempty" xml:"DeviceStatusInfo,omitempty" type:"Struct"`
	// List of disposal actions.
	DisposalActions []*string `json:"DisposalActions,omitempty" xml:"DisposalActions,omitempty" type:"Repeated"`
	// Disposal process ID.
	//
	// example:
	//
	// dp-xxxxxxxx
	DisposalProcessId *string `json:"DisposalProcessId,omitempty" xml:"DisposalProcessId,omitempty"`
	// Disposal time, in seconds since the epoch.
	//
	// example:
	//
	// 1743059249
	DisposalTime *string `json:"DisposalTime,omitempty" xml:"DisposalTime,omitempty"`
	// Dynamic policy ID.
	//
	// example:
	//
	// dynamic-policy-xxxxxxxx
	DynamicPolicyId *string `json:"DynamicPolicyId,omitempty" xml:"DynamicPolicyId,omitempty"`
	// Dynamic policy name.
	//
	// example:
	//
	// 动态策略1
	DynamicPolicyName *string `json:"DynamicPolicyName,omitempty" xml:"DynamicPolicyName,omitempty"`
	// Terminal device name. Length: 1~128 characters, supporting Chinese and uppercase/lowercase English letters, and can include numbers, half-width periods (.), commas (,), semicolons (;), hyphens (-), underscores (_), slashes (/), at (@) symbols, and spaces. Entering an underscore (_) alone will additionally query all terminal devices with 4-byte UTF-8 characters in their names.
	//
	// example:
	//
	// WANGCHENCHENNBB
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Recovery type.
	//
	// - **auto**：Automatic recovery.
	//
	// - **console**：Console recovery.
	//
	// - **auth**：Certification and reporting recovery.
	//
	// example:
	//
	// auto
	RecoveryType *string `json:"RecoveryType,omitempty" xml:"RecoveryType,omitempty"`
	// Rule content.
	//
	// example:
	//
	// {
	//
	//   "Combinator": "OR",
	//
	//   "Rules": [
	//
	//     {
	//
	//       "Operator": "version_gt",
	//
	//       "Values": [
	//
	//         "1"
	//
	//       ],
	//
	//       "RuleType": "device_info",
	//
	//       "Id": "1",
	//
	//       "RuleSubType": "basic_info",
	//
	//       "Name": "app_version"
	//
	//     }
	//
	//   ]
	//
	// }
	RuleContent interface{} `json:"RuleContent,omitempty" xml:"RuleContent,omitempty"`
	// SASE用户ID。
	//
	// example:
	//
	// asdqwedg-xzczvzdaf-asfafs
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// Disposal status. Values:
	//
	// - **disposal**: In the disposal state.
	//
	// - **finished**: Already automatically recovered.
	//
	// - **recovery**: Recovered by authentication and reporting or console recovery.
	//
	// example:
	//
	// disposal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Username.
	//
	// example:
	//
	// xiaoming
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListDynamicDisposalProcessesResponseBodyDisposalProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GoString() string {
	return s.String()
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDepartment() *string {
	return s.Department
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDevTag() *string {
	return s.DevTag
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDeviceBasicInfo() *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	return s.DeviceBasicInfo
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDeviceStatusInfo() *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	return s.DeviceStatusInfo
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDisposalActions() []*string {
	return s.DisposalActions
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDisposalProcessId() *string {
	return s.DisposalProcessId
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDisposalTime() *string {
	return s.DisposalTime
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDynamicPolicyId() *string {
	return s.DynamicPolicyId
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetDynamicPolicyName() *string {
	return s.DynamicPolicyName
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetHostname() *string {
	return s.Hostname
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetRecoveryType() *string {
	return s.RecoveryType
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetRuleContent() interface{} {
	return s.RuleContent
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetStatus() *string {
	return s.Status
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) GetUserName() *string {
	return s.UserName
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDepartment(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.Department = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDevTag(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.DevTag = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDeviceBasicInfo(v *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.DeviceBasicInfo = v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDeviceStatusInfo(v *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.DeviceStatusInfo = v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDisposalActions(v []*string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.DisposalActions = v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDisposalProcessId(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.DisposalProcessId = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDisposalTime(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.DisposalTime = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDynamicPolicyId(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.DynamicPolicyId = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetDynamicPolicyName(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.DynamicPolicyName = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetHostname(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.Hostname = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetRecoveryType(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.RecoveryType = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetRuleContent(v interface{}) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.RuleContent = v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetSaseUserId(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.SaseUserId = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetStatus(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.Status = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) SetUserName(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcesses {
	s.UserName = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcesses) Validate() error {
	return dara.Validate(s)
}

type ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo struct {
	// CPU model.
	//
	// example:
	//
	// Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Device ID.
	//
	// example:
	//
	// A84D0AF0-1ACC-02B8-6A07-FC898F71BE09
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// Device operating system type. Values:
	//
	// - **Windows**：Windows system.
	//
	// - **macOS**：macOS system.
	//
	// - **Linux**：Linux system.
	//
	// - **Android**：Android system.
	//
	// - **iOS**：iOS system.
	//
	// - **Windows_Wuying**：Wuying cloud desktop system.
	//
	// example:
	//
	// windows
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
	// Device disk model.
	//
	// example:
	//
	// KXG6AZNV512G TOSHIBA
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// Device name.
	//
	// example:
	//
	// DESKTOP-ERLV3AK
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Device MAC address.
	//
	// example:
	//
	// CE:3B:**:**:FD:FB
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// Device memory capacity. Unit: GB.
	//
	// example:
	//
	// 2
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Operating system version
	//
	// example:
	//
	// 1
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
}

func (s ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GoString() string {
	return s.String()
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GetCpu() *string {
	return s.Cpu
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GetDevTag() *string {
	return s.DevTag
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GetDevType() *string {
	return s.DevType
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GetDisk() *string {
	return s.Disk
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GetHostname() *string {
	return s.Hostname
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GetMac() *string {
	return s.Mac
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GetMemory() *string {
	return s.Memory
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) GetOsVersion() *string {
	return s.OsVersion
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) SetCpu(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	s.Cpu = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) SetDevTag(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	s.DevTag = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) SetDevType(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	s.DevType = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) SetDisk(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	s.Disk = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) SetHostname(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	s.Hostname = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) SetMac(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	s.Mac = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) SetMemory(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	s.Memory = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) SetOsVersion(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo {
	s.OsVersion = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceBasicInfo) Validate() error {
	return dara.Validate(s)
}

type ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo struct {
	// Client version.
	//
	// example:
	//
	// 4.5.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// Department to which the user belongs.
	//
	// example:
	//
	// IT运维部
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// Office data protection status. Values:
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
	// enabled
	DlpStatus *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	// Public IP address.
	//
	// example:
	//
	// 120.26.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Internet behavior management enablement status.
	//
	// example:
	//
	// enabled
	LaStatus *string `json:"LaStatus,omitempty" xml:"LaStatus,omitempty"`
	// Login status.
	//
	// example:
	//
	// online
	LoginStatus *string `json:"LoginStatus,omitempty" xml:"LoginStatus,omitempty"`
	// Network access control status. Values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// - **Unprovisioned**: Not configured.
	//
	// example:
	//
	// enabled
	NacStatus *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	// Private IP address.
	//
	// example:
	//
	// 172.20.XX.XX
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// Unique ID of the SASE user.
	//
	// example:
	//
	// su_dfsdfsdgasgsgag
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// Username.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// Identified office area name.
	//
	// example:
	//
	// office
	Workshop *string `json:"Workshop,omitempty" xml:"Workshop,omitempty"`
	// ZTNA enablement status.
	//
	// example:
	//
	// enabled
	ZtnaStatus *string `json:"ZtnaStatus,omitempty" xml:"ZtnaStatus,omitempty"`
}

func (s ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GoString() string {
	return s.String()
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetDepartment() *string {
	return s.Department
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetDlpStatus() *string {
	return s.DlpStatus
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetLaStatus() *string {
	return s.LaStatus
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetLoginStatus() *string {
	return s.LoginStatus
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetNacStatus() *string {
	return s.NacStatus
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetUsername() *string {
	return s.Username
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetWorkshop() *string {
	return s.Workshop
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) GetZtnaStatus() *string {
	return s.ZtnaStatus
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetAppVersion(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.AppVersion = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetDepartment(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.Department = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetDlpStatus(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.DlpStatus = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetInternetIp(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.InternetIp = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetLaStatus(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.LaStatus = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetLoginStatus(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.LoginStatus = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetNacStatus(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.NacStatus = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetPrivateIp(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.PrivateIp = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetSaseUserId(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.SaseUserId = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetUsername(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.Username = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetWorkshop(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.Workshop = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) SetZtnaStatus(v string) *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo {
	s.ZtnaStatus = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponseBodyDisposalProcessesDeviceStatusInfo) Validate() error {
	return dara.Validate(s)
}
