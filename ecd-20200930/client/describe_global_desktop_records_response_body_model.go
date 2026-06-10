// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDesktopRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeGlobalDesktopRecordsResponseBody
	GetRequestId() *string
	SetSessions(v []*DescribeGlobalDesktopRecordsResponseBodySessions) *DescribeGlobalDesktopRecordsResponseBody
	GetSessions() []*DescribeGlobalDesktopRecordsResponseBodySessions
	SetTotalCount(v int32) *DescribeGlobalDesktopRecordsResponseBody
	GetTotalCount() *int32
}

type DescribeGlobalDesktopRecordsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of sessions.
	Sessions []*DescribeGlobalDesktopRecordsResponseBodySessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGlobalDesktopRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalDesktopRecordsResponseBody) GetSessions() []*DescribeGlobalDesktopRecordsResponseBodySessions {
	return s.Sessions
}

func (s *DescribeGlobalDesktopRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGlobalDesktopRecordsResponseBody) SetRequestId(v string) *DescribeGlobalDesktopRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBody) SetSessions(v []*DescribeGlobalDesktopRecordsResponseBodySessions) *DescribeGlobalDesktopRecordsResponseBody {
	s.Sessions = v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBody) SetTotalCount(v int32) *DescribeGlobalDesktopRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBody) Validate() error {
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

type DescribeGlobalDesktopRecordsResponseBodySessions struct {
	// The connection status.
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The ID of the desktop group.
	//
	// example:
	//
	// dg-iaqu3bi2xtie****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The name of the desktop group.
	//
	// example:
	//
	// DemoCCGroup
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// The ID of the cloud desktop.
	//
	// example:
	//
	// ecd-g6t1ukbaea****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud desktop.
	//
	// example:
	//
	// DemoComputer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The desktop status.
	//
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// TestUser
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// A list of assigned end user IDs.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The GPU memory size.
	//
	// example:
	//
	// 8GiB
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// The duration of the last connection, in seconds.
	//
	// example:
	//
	// 120
	LatestConnectionTime *int64 `json:"LatestConnectionTime,omitempty" xml:"LatestConnectionTime,omitempty"`
	// The memory size of the cloud desktop, in MiB.
	//
	// example:
	//
	// 4096
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ID of the office site.
	//
	// example:
	//
	// cn-hangzhou+dir-8904****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The name of the office site.
	//
	// example:
	//
	// TestOfficeSite
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The office site type.
	//
	// example:
	//
	// Simple
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The operating system type. Valid values:
	//
	// - `Windows`
	//
	// - `Linux`
	//
	// example:
	//
	// Linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The operating system version.
	//
	// example:
	//
	// Windows 10
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The protocol type. Valid values:
	//
	// - `HDX`
	//
	// - `ASP`
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// A list of resource groups.
	ResourceGroups []*DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// The idle duration of the session, in minutes.
	//
	// example:
	//
	// 120
	SessionIdleTime *int64 `json:"SessionIdleTime,omitempty" xml:"SessionIdleTime,omitempty"`
	// A list of sessions.
	Sessions []*DescribeGlobalDesktopRecordsResponseBodySessionsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The time when the cloud desktop status changed.
	//
	// example:
	//
	// 1760583xxxx
	StatusChangeTime *int64 `json:"StatusChangeTime,omitempty" xml:"StatusChangeTime,omitempty"`
	// The billing method for the cloud desktop. Valid values:
	//
	// - `prePaid`: Subscription.
	//
	// - `postPaid`: Pay-as-you-go.
	//
	// - `monthPackage`: Monthly usage package.
	//
	// example:
	//
	// monthPackage
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// The total connection duration, in seconds.
	//
	// example:
	//
	// 240
	TotalConnectionTime *int64 `json:"TotalConnectionTime,omitempty" xml:"TotalConnectionTime,omitempty"`
	// The cloud desktop uptime, in seconds.
	//
	// example:
	//
	// 86400
	UpTime *int64 `json:"UpTime,omitempty" xml:"UpTime,omitempty"`
}

func (s DescribeGlobalDesktopRecordsResponseBodySessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopRecordsResponseBodySessions) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetLatestConnectionTime() *int64 {
	return s.LatestConnectionTime
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetOsType() *string {
	return s.OsType
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetResourceGroups() []*DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups {
	return s.ResourceGroups
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetSessionIdleTime() *int64 {
	return s.SessionIdleTime
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetSessions() []*DescribeGlobalDesktopRecordsResponseBodySessionsSessions {
	return s.Sessions
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetStatusChangeTime() *int64 {
	return s.StatusChangeTime
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetSubPayType() *string {
	return s.SubPayType
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetTotalConnectionTime() *int64 {
	return s.TotalConnectionTime
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) GetUpTime() *int64 {
	return s.UpTime
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetConnectionStatus(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetCpu(v int32) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.Cpu = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetDesktopGroupId(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetDesktopGroupName(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetDesktopId(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.DesktopId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetDesktopName(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetDesktopStatus(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetEndUserId(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetEndUserIds(v []*string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.EndUserIds = v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetGpuSpec(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.GpuSpec = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetLatestConnectionTime(v int64) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.LatestConnectionTime = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetMemory(v int64) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.Memory = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetOfficeSiteId(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetOfficeSiteName(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetOfficeSiteType(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetOsType(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.OsType = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetPlatform(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.Platform = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetProtocolType(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.ProtocolType = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetRegionId(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetResourceGroups(v []*DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.ResourceGroups = v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetSessionIdleTime(v int64) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.SessionIdleTime = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetSessions(v []*DescribeGlobalDesktopRecordsResponseBodySessionsSessions) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.Sessions = v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetStatusChangeTime(v int64) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.StatusChangeTime = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetSubPayType(v string) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.SubPayType = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetTotalConnectionTime(v int64) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.TotalConnectionTime = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) SetUpTime(v int64) *DescribeGlobalDesktopRecordsResponseBodySessions {
	s.UpTime = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessions) Validate() error {
	if s.ResourceGroups != nil {
		for _, item := range s.ResourceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups struct {
	// The ID of the resource group.
	//
	// example:
	//
	// rg-f3s3dgt8dtb0vlqc8
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// dms_test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups) SetResourceGroupId(v string) *DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups) SetResourceGroupName(v string) *DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsResourceGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDesktopRecordsResponseBodySessionsSessions struct {
	// The ID of the end user.
	//
	// example:
	//
	// TestUser
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time the session was created.
	//
	// example:
	//
	// 2022-08-31T06:56:45Z
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty"`
}

func (s DescribeGlobalDesktopRecordsResponseBodySessionsSessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopRecordsResponseBodySessionsSessions) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsSessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsSessions) GetEstablishmentTime() *string {
	return s.EstablishmentTime
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsSessions) SetEndUserId(v string) *DescribeGlobalDesktopRecordsResponseBodySessionsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsSessions) SetEstablishmentTime(v string) *DescribeGlobalDesktopRecordsResponseBodySessionsSessions {
	s.EstablishmentTime = &v
	return s
}

func (s *DescribeGlobalDesktopRecordsResponseBodySessionsSessions) Validate() error {
	return dara.Validate(s)
}
