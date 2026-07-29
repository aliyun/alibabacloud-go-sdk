// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktops(v []*DescribeGlobalDesktopsResponseBodyDesktops) *DescribeGlobalDesktopsResponseBody
	GetDesktops() []*DescribeGlobalDesktopsResponseBodyDesktops
	SetNextToken(v string) *DescribeGlobalDesktopsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeGlobalDesktopsResponseBody
	GetRequestId() *string
}

type DescribeGlobalDesktopsResponseBody struct {
	// Detailed cloud desktop information.
	Desktops []*DescribeGlobalDesktopsResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
	// Token that marks the start of the next query. If NextToken is empty, no more results are available.
	//
	// example:
	//
	// eyJkZWZhdWx0IjpbIjIwMjItMDgtMTdUM****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4686A731-D601-548C-83E2-4CB6371E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBody) GetDesktops() []*DescribeGlobalDesktopsResponseBodyDesktops {
	return s.Desktops
}

func (s *DescribeGlobalDesktopsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGlobalDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalDesktopsResponseBody) SetDesktops(v []*DescribeGlobalDesktopsResponseBodyDesktops) *DescribeGlobalDesktopsResponseBody {
	s.Desktops = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBody) SetNextToken(v string) *DescribeGlobalDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBody) SetRequestId(v string) *DescribeGlobalDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBody) Validate() error {
	if s.Desktops != nil {
		for _, item := range s.Desktops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalDesktopsResponseBodyDesktops struct {
	// Billing method.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Supported client information.
	Clients []*DescribeGlobalDesktopsResponseBodyDesktopsClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// Endpoint connection status.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// Number of CPUs.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Time when the cloud desktop was created.
	//
	// example:
	//
	// 2020-11-06T08:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Shared cloud desktop ID.
	//
	// example:
	//
	// dg-3uiojcc0j4kh7****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// Cloud desktop ID.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// Cloud desktop name.
	//
	// example:
	//
	// testDesktopName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// Cloud desktop status.
	//
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// Cloud desktop timer object.
	DesktopTimers []*DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	// Cloud desktop specification.
	//
	// example:
	//
	// ecd.basic.large
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// Office site ID. Same as `OfficeSiteId`.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Disk information.
	Disks []*DescribeGlobalDesktopsResponseBodyDesktopsDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// End user name.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// List of end users.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// Time when the cloud desktop expires.
	//
	// - Valid for subscription cloud desktops.
	//
	// - For pay-as-you-go cloud desktops, always returns `2099-12-31T15:59Z`.
	//
	// example:
	//
	// 2021-12-31T15:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Image update information.
	FotaUpdate *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate `json:"FotaUpdate,omitempty" xml:"FotaUpdate,omitempty" type:"Struct"`
	// GPU memory size. Valid for GPU cloud desktops. Unit: MB.
	//
	// example:
	//
	// 2048
	GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// Whether this is a hibernation beta version.
	//
	// example:
	//
	// true
	HibernationBeta *bool `json:"HibernationBeta,omitempty" xml:"HibernationBeta,omitempty"`
	// Host name.
	//
	// example:
	//
	// testName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// Image ID.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Last startup time — the most recent time the cloud desktop started.
	//
	// example:
	//
	// 2021-07-13T15:59Z
	LastStartTime *string `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	// Region name.
	//
	// example:
	//
	// China (Shanghai)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// List of cloud desktop management statuses.
	ManagementFlags []*string `json:"ManagementFlags,omitempty" xml:"ManagementFlags,omitempty" type:"Repeated"`
	// Cloud desktop memory size, in MiB.
	//
	// example:
	//
	// 4096
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Network interface IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// Office site ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// Operating system type
	//
	// example:
	//
	// Windows
	Os            *string `json:"Os,omitempty" xml:"Os,omitempty"`
	OsDescription *string `json:"OsDescription,omitempty" xml:"OsDescription,omitempty"`
	// Operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// Operating system platform.
	//
	// example:
	//
	// Ubuntu
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// Cloud desktop policy ID.
	//
	// example:
	//
	// pg-9cktlowtxfl6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// Protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// If this is a shared cloud desktop and a real cloud desktop has been assigned, this field shows the cloud desktop ID.
	//
	// example:
	//
	// ecd-gx2x1dhsm****
	RealDesktopId *string `json:"RealDesktopId,omitempty" xml:"RealDesktopId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionLocation *string `json:"RegionLocation,omitempty" xml:"RegionLocation,omitempty"`
	// Session type.
	//
	// example:
	//
	// SINGLE_SESSION
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// List of session information.
	Sessions []*DescribeGlobalDesktopsResponseBodyDesktopsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// Whether hibernation is supported.
	//
	// example:
	//
	// true
	SupportHibernation *bool `json:"SupportHibernation,omitempty" xml:"SupportHibernation,omitempty"`
	// User-defined cloud desktop name.
	//
	// example:
	//
	// testDesktop
	UserCustomName *string `json:"UserCustomName,omitempty" xml:"UserCustomName,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetClients() []*DescribeGlobalDesktopsResponseBodyDesktopsClients {
	return s.Clients
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetDesktopTimers() []*DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	return s.DesktopTimers
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetDisks() []*DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	return s.Disks
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetFotaUpdate() *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	return s.FotaUpdate
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetGpuMemory() *int32 {
	return s.GpuMemory
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetHibernationBeta() *bool {
	return s.HibernationBeta
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetHostName() *string {
	return s.HostName
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetLastStartTime() *string {
	return s.LastStartTime
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetManagementFlags() []*string {
	return s.ManagementFlags
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetNetworkInterfaceIp() *string {
	return s.NetworkInterfaceIp
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetOs() *string {
	return s.Os
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetOsDescription() *string {
	return s.OsDescription
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetOsType() *string {
	return s.OsType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetRealDesktopId() *string {
	return s.RealDesktopId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetRegionLocation() *string {
	return s.RegionLocation
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetSessionType() *string {
	return s.SessionType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetSessions() []*DescribeGlobalDesktopsResponseBodyDesktopsSessions {
	return s.Sessions
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetSupportHibernation() *bool {
	return s.SupportHibernation
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) GetUserCustomName() *string {
	return s.UserCustomName
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetChargeType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetClients(v []*DescribeGlobalDesktopsResponseBodyDesktopsClients) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Clients = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetConnectionStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetCpu(v int32) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetCreationTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopGroupId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopTimers(v []*DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopTimers = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDirectoryId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDisks(v []*DescribeGlobalDesktopsResponseBodyDesktopsDisks) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Disks = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetEndUserId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetEndUserIds(v []*string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetExpiredTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetFotaUpdate(v *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.FotaUpdate = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetGpuMemory(v int32) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.GpuMemory = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetHibernationBeta(v bool) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.HibernationBeta = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetHostName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.HostName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetImageId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetLastStartTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.LastStartTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetLocalName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.LocalName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetManagementFlags(v []*string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ManagementFlags = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetMemory(v int64) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetNetworkInterfaceIp(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOfficeSiteId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOs(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Os = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOsDescription(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.OsDescription = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOsType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetPlatform(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Platform = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetPolicyGroupId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetProtocolType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ProtocolType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetRealDesktopId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.RealDesktopId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetRegionId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetRegionLocation(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.RegionLocation = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSessionType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.SessionType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSessions(v []*DescribeGlobalDesktopsResponseBodyDesktopsSessions) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSupportHibernation(v bool) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.SupportHibernation = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetUserCustomName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.UserCustomName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) Validate() error {
	if s.Clients != nil {
		for _, item := range s.Clients {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DesktopTimers != nil {
		for _, item := range s.DesktopTimers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Disks != nil {
		for _, item := range s.Disks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FotaUpdate != nil {
		if err := s.FotaUpdate.Validate(); err != nil {
			return err
		}
	}
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalDesktopsResponseBodyDesktopsClients struct {
	// Client type.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Client status.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsClients) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsClients) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) GetStatus() *string {
	return s.Status
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) SetClientType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsClients {
	s.ClientType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) SetStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktopsClients {
	s.Status = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers struct {
	// Whether clients can set policies.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// Cron expression for the scheduled task.
	//
	// Example: `0 0 4 1/1 	- ?` means run daily starting at 4:00 AM on the first day of each month.
	//
	// example:
	//
	// 0 0 0 ? 	- 1
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Whether to enforce execution. If `true`, the scheduled task runs regardless of cloud desktop or connection status.
	//
	// example:
	//
	// false
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// Task execution time.
	//
	// example:
	//
	// 2021-12-31T15:59Z
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// Time interval, in seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Scheduled task type.
	//
	// example:
	//
	// SHUTDOWN
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// Reset type for reset tasks.
	//
	// example:
	//
	// RESET_TYPE_BOTH
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// Scheduled task type.
	//
	// example:
	//
	// NoConnectShutdown
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GetAllowClientSetting() *bool {
	return s.AllowClientSetting
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GetResetType() *string {
	return s.ResetType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetAllowClientSetting(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetCronExpression(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.CronExpression = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetEnforce(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.Enforce = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetExecutionTime(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetInterval(v int32) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.Interval = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetOperationType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.OperationType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetResetType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.ResetType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetTimerType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.TimerType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDesktopsResponseBodyDesktopsDisks struct {
	// Disk ID.
	//
	// example:
	//
	// d-jedbpr4sl9l37****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Disk size, in GiB.
	//
	// example:
	//
	// 80
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// Disk type.
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskId(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskSize(v int32) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate struct {
	// Subscription channel.
	//
	// example:
	//
	// Enterprise
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Current version number of the cloud desktop.
	//
	// example:
	//
	// 0.0.0-D-20220102.xxxx
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	// Whether the upgrade is forced.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// New application version number after the update.
	//
	// example:
	//
	// 0.0.0-R-20220307.xxxx
	NewAppVersion *string `json:"NewAppVersion,omitempty" xml:"NewAppVersion,omitempty"`
	NewDcdVersion *string `json:"NewDcdVersion,omitempty" xml:"NewDcdVersion,omitempty"`
	// Project name.
	//
	// example:
	//
	// testProject
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// Description of the upgradable version.
	//
	// example:
	//
	// Test upgrade package 03-07
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// English description of the upgradable version.
	//
	// example:
	//
	// Release note
	ReleaseNoteEn *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	// Japanese description of the upgradable version.
	//
	// example:
	//
	// リリースノート
	ReleaseNoteJp *string `json:"ReleaseNoteJp,omitempty" xml:"ReleaseNoteJp,omitempty"`
	// Size of the upgradable version package, in MiB.
	//
	// example:
	//
	// 100
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetChannel() *string {
	return s.Channel
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetCurrentAppVersion() *string {
	return s.CurrentAppVersion
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetForce() *bool {
	return s.Force
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetNewAppVersion() *string {
	return s.NewAppVersion
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetNewDcdVersion() *string {
	return s.NewDcdVersion
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetProject() *string {
	return s.Project
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetReleaseNoteJp() *string {
	return s.ReleaseNoteJp
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GetSize() *string {
	return s.Size
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetChannel(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Channel = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetCurrentAppVersion(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetForce(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Force = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetNewAppVersion(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.NewAppVersion = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetNewDcdVersion(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.NewDcdVersion = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetProject(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Project = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNote(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNoteEn(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNoteJp(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNoteJp = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetSize(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Size = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDesktopsResponseBodyDesktopsSessions struct {
	// End user information.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Time when the session was created.
	//
	// example:
	//
	// 2021-03-07T08:23Z
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsSessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) GetEstablishmentTime() *string {
	return s.EstablishmentTime
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) SetEndUserId(v string) *DescribeGlobalDesktopsResponseBodyDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) SetEstablishmentTime(v string) *DescribeGlobalDesktopsResponseBodyDesktopsSessions {
	s.EstablishmentTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) Validate() error {
	return dara.Validate(s)
}
