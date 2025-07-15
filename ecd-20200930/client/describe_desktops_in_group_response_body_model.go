// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopsInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeDesktopsInGroupResponseBody
	GetNextToken() *string
	SetOnlinePrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody
	GetOnlinePrePaidDesktopsCount() *int32
	SetPaidDesktops(v []*DescribeDesktopsInGroupResponseBodyPaidDesktops) *DescribeDesktopsInGroupResponseBody
	GetPaidDesktops() []*DescribeDesktopsInGroupResponseBodyPaidDesktops
	SetPaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody
	GetPaidDesktopsCount() *int32
	SetPostPaidDesktops(v []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops) *DescribeDesktopsInGroupResponseBody
	GetPostPaidDesktops() []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops
	SetPostPaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody
	GetPostPaidDesktopsCount() *int32
	SetPostPaidDesktopsTotalAmount(v int32) *DescribeDesktopsInGroupResponseBody
	GetPostPaidDesktopsTotalAmount() *int32
	SetRequestId(v string) *DescribeDesktopsInGroupResponseBody
	GetRequestId() *string
	SetRunningPrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody
	GetRunningPrePaidDesktopsCount() *int32
	SetStopedPrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody
	GetStopedPrePaidDesktopsCount() *int32
	SetStoppedPrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody
	GetStoppedPrePaidDesktopsCount() *int32
}

type DescribeDesktopsInGroupResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of subscription cloud computers that are in the Connected state.
	//
	// example:
	//
	// 100
	OnlinePrePaidDesktopsCount *int32 `json:"OnlinePrePaidDesktopsCount,omitempty" xml:"OnlinePrePaidDesktopsCount,omitempty"`
	// The subscription cloud computers.
	PaidDesktops []*DescribeDesktopsInGroupResponseBodyPaidDesktops `json:"PaidDesktops,omitempty" xml:"PaidDesktops,omitempty" type:"Repeated"`
	// The total number of subscription cloud computers.
	//
	// example:
	//
	// 10
	PaidDesktopsCount *int32 `json:"PaidDesktopsCount,omitempty" xml:"PaidDesktopsCount,omitempty"`
	// The pay-as-you-go cloud computers.
	PostPaidDesktops []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops `json:"PostPaidDesktops,omitempty" xml:"PostPaidDesktops,omitempty" type:"Repeated"`
	// The total number of pay-as-you-go cloud computers.
	//
	// example:
	//
	// 10
	PostPaidDesktopsCount *int32 `json:"PostPaidDesktopsCount,omitempty" xml:"PostPaidDesktopsCount,omitempty"`
	// The total amount of bills for pay-as-you-go cloud computers.
	//
	// example:
	//
	// 10000
	PostPaidDesktopsTotalAmount *int32 `json:"PostPaidDesktopsTotalAmount,omitempty" xml:"PostPaidDesktopsTotalAmount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of subscription cloud computers that are in the Running state.
	//
	// example:
	//
	// 100
	RunningPrePaidDesktopsCount *int32 `json:"RunningPrePaidDesktopsCount,omitempty" xml:"RunningPrePaidDesktopsCount,omitempty"`
	// The number of subscription cloud computers that are in the Stopped state.
	//
	// example:
	//
	// 100
	StopedPrePaidDesktopsCount *int32 `json:"StopedPrePaidDesktopsCount,omitempty" xml:"StopedPrePaidDesktopsCount,omitempty"`
	// The number of subscription cloud computers that are in the Stopped state.
	//
	// example:
	//
	// 100
	StoppedPrePaidDesktopsCount *int32 `json:"StoppedPrePaidDesktopsCount,omitempty" xml:"StoppedPrePaidDesktopsCount,omitempty"`
}

func (s DescribeDesktopsInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopsInGroupResponseBody) GetOnlinePrePaidDesktopsCount() *int32 {
	return s.OnlinePrePaidDesktopsCount
}

func (s *DescribeDesktopsInGroupResponseBody) GetPaidDesktops() []*DescribeDesktopsInGroupResponseBodyPaidDesktops {
	return s.PaidDesktops
}

func (s *DescribeDesktopsInGroupResponseBody) GetPaidDesktopsCount() *int32 {
	return s.PaidDesktopsCount
}

func (s *DescribeDesktopsInGroupResponseBody) GetPostPaidDesktops() []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	return s.PostPaidDesktops
}

func (s *DescribeDesktopsInGroupResponseBody) GetPostPaidDesktopsCount() *int32 {
	return s.PostPaidDesktopsCount
}

func (s *DescribeDesktopsInGroupResponseBody) GetPostPaidDesktopsTotalAmount() *int32 {
	return s.PostPaidDesktopsTotalAmount
}

func (s *DescribeDesktopsInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopsInGroupResponseBody) GetRunningPrePaidDesktopsCount() *int32 {
	return s.RunningPrePaidDesktopsCount
}

func (s *DescribeDesktopsInGroupResponseBody) GetStopedPrePaidDesktopsCount() *int32 {
	return s.StopedPrePaidDesktopsCount
}

func (s *DescribeDesktopsInGroupResponseBody) GetStoppedPrePaidDesktopsCount() *int32 {
	return s.StoppedPrePaidDesktopsCount
}

func (s *DescribeDesktopsInGroupResponseBody) SetNextToken(v string) *DescribeDesktopsInGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetOnlinePrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.OnlinePrePaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPaidDesktops(v []*DescribeDesktopsInGroupResponseBodyPaidDesktops) *DescribeDesktopsInGroupResponseBody {
	s.PaidDesktops = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktops(v []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktops = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktopsTotalAmount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktopsTotalAmount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetRequestId(v string) *DescribeDesktopsInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetRunningPrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.RunningPrePaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetStopedPrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.StopedPrePaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetStoppedPrePaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.StoppedPrePaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsInGroupResponseBodyPaidDesktops struct {
	// The connection status of the cloud computer.
	//
	// Valid values:
	//
	// 	- Unknown
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Connected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Disconnected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The ID of the cloud computer.
	//
	// example:
	//
	// ud-7ftf5b6yu77b0****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// testName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The status of the cloud computer.
	//
	// Valid values:
	//
	// 	- Stopped
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Rebuilding
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Expired
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleted
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Pending
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// The type of the disk.
	//
	// Valid values:
	//
	// 	- SYSTEM: system disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DATA: data disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the authorized user.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The IDs of the end users who are connected to the cloud computers in the cloud computer share. If no end users are connected, no values are returned for this parameter.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The username of the authorized user.
	//
	// example:
	//
	// alice
	EndUserName *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
	// The usernames of the end users who are connected to the cloud computers in the cloud computer share. If no end users are connected, no values are returned for this parameter.
	EndUserNames []*string `json:"EndUserNames,omitempty" xml:"EndUserNames,omitempty" type:"Repeated"`
	ExpiredTime  *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The image version.
	//
	// example:
	//
	// 0.1.0-R-20220914.17****
	FotaVersion *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	// The version of the GPU driver.
	//
	// example:
	//
	// 1.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-hn5v2mmk0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Win10_ZC
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The flag that is used to manage the cloud computer.
	//
	// Valid values:
	//
	// 	- Updating: The configurations of the cloud computer are being updated.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NoFlag: No flags are attached to the cloud computer.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// NoFlag
	ManagementFlag *string `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	// The flags that are used to manage the cloud computers.
	ManagementFlags []*string `json:"ManagementFlags,omitempty" xml:"ManagementFlags,omitempty" type:"Repeated"`
	// The IP address of the member network interface controller (NIC) of the instance.
	//
	// example:
	//
	// 192.168.XX.XX
	MemberEniIp *string `json:"MemberEniIp,omitempty" xml:"MemberEniIp,omitempty"`
	// The OS.
	//
	// Valid values:
	//
	// 	- Linux
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The IP address of the primary NIC of the instance.
	//
	// example:
	//
	// 192.168.XX.XX
	PrimaryEniIp *string `json:"PrimaryEniIp,omitempty" xml:"PrimaryEniIp,omitempty"`
	// The protocol.
	//
	// Valid values:
	//
	// 	- HDX: High-definition Experience (HDX) protocol
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ASP: Adaptive Streaming Protocol (ASP) protocol provided by Alibaba Cloud
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The time when the cloud computer was reset.
	//
	// example:
	//
	// 2021-03-03 08:48:08
	ResetTime *string `json:"ResetTime,omitempty" xml:"ResetTime,omitempty"`
	// The system disk size. Unit: GiB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopsInGroupResponseBodyPaidDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBodyPaidDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetEndUserName() *string {
	return s.EndUserName
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetEndUserNames() []*string {
	return s.EndUserNames
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetFotaVersion() *string {
	return s.FotaVersion
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetManagementFlag() *string {
	return s.ManagementFlag
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetManagementFlags() []*string {
	return s.ManagementFlags
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetMemberEniIp() *string {
	return s.MemberEniIp
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetOsType() *string {
	return s.OsType
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetPrimaryEniIp() *string {
	return s.PrimaryEniIp
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetResetTime() *string {
	return s.ResetTime
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetConnectionStatus(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopId(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopName(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopStatus(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDiskType(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DiskType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserId(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserIds(v []*string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserName(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserNames(v []*string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserNames = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetExpiredTime(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetFotaVersion(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.FotaVersion = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetGpuDriverVersion(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetImageId(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetImageName(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ImageName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetManagementFlag(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetManagementFlags(v []*string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ManagementFlags = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetMemberEniIp(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.MemberEniIp = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetOsType(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetPrimaryEniIp(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.PrimaryEniIp = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetProtocolType(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetResetTime(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ResetTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetSystemDiskSize(v int32) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsInGroupResponseBodyPostPaidDesktops struct {
	// The connection status of the cloud computer.
	//
	// Valid values:
	//
	// 	- Unknown
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Connected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Disconnected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The retention period. Unit: milliseconds.
	//
	// example:
	//
	// 4153958447
	CreateDuration *string `json:"CreateDuration,omitempty" xml:"CreateDuration,omitempty"`
	// The time when the cloud computer was created.
	//
	// example:
	//
	// 2022-01-21T06:34:57Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the cloud computer.
	//
	// example:
	//
	// ud-2i8qxpv6t1a07****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// testName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The status of the cloud computer.
	//
	// Valid values:
	//
	// 	- Stopped
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Rebuilding
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Expired
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleted
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Pending
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Stopped
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// The type of the disk.
	//
	// Valid values:
	//
	// 	- SYSTEM: system disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DATA: data disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the authorized user.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The IDs of the end users who are connected to the cloud computers in the cloud computer pool. If no end users are connected, no values are returned for this parameter.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The username of the authorized user.
	//
	// example:
	//
	// alice
	EndUserName *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
	// The usernames of the end users who are connected to the cloud computers in the cloud computer pool. If no end users are connected, no values are returned for this parameter.
	EndUserNames []*string `json:"EndUserNames,omitempty" xml:"EndUserNames,omitempty" type:"Repeated"`
	// The image version.
	//
	// example:
	//
	// 0.1.0-R-20220914.17****
	FotaVersion *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	// The version of the GPU driver.
	//
	// example:
	//
	// 1.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-hn5v2mmk0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Win10_ZC
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The flag that is used to manage the cloud computer.
	//
	// Valid values:
	//
	// 	- Updating: The configurations of the cloud computer are being updated.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NoFlag: No flags are attached to the cloud computer.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// NoFlag
	ManagementFlag *string `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	// The flags that are used to manage the cloud computers.
	ManagementFlags []*string `json:"ManagementFlags,omitempty" xml:"ManagementFlags,omitempty" type:"Repeated"`
	// The IP address of the member NIC of the instance.
	//
	// example:
	//
	// 192.168.XX.XX
	MemberEniIp *string `json:"MemberEniIp,omitempty" xml:"MemberEniIp,omitempty"`
	// The OS.
	//
	// Valid values:
	//
	// 	- Linux
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The IP address of the primary NIC of the instance.
	//
	// example:
	//
	// 192.168.XX.XX
	PrimaryEniIp *string `json:"PrimaryEniIp,omitempty" xml:"PrimaryEniIp,omitempty"`
	// The protocol.
	//
	// Valid values:
	//
	// 	- HDX: HDX protocol
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ASP: ASP protocol provided by Alibaba Cloud
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// HDX
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The time when the cloud computer was released.
	//
	// example:
	//
	// 2022-01-21T16:34:57Z
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The time when the cloud computer was reset.
	//
	// example:
	//
	// 2021-03-03 08:48:08
	ResetTime *string `json:"ResetTime,omitempty" xml:"ResetTime,omitempty"`
	// The system disk size. Unit: GiB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopsInGroupResponseBodyPostPaidDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetCreateDuration() *string {
	return s.CreateDuration
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetEndUserName() *string {
	return s.EndUserName
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetEndUserNames() []*string {
	return s.EndUserNames
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetFotaVersion() *string {
	return s.FotaVersion
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetManagementFlag() *string {
	return s.ManagementFlag
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetManagementFlags() []*string {
	return s.ManagementFlags
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetMemberEniIp() *string {
	return s.MemberEniIp
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetOsType() *string {
	return s.OsType
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetPrimaryEniIp() *string {
	return s.PrimaryEniIp
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetResetTime() *string {
	return s.ResetTime
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetConnectionStatus(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetCreateDuration(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.CreateDuration = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetCreateTime(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.CreateTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopId(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopName(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopStatus(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDiskType(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DiskType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserId(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserIds(v []*string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserName(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserNames(v []*string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserNames = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetFotaVersion(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.FotaVersion = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetGpuDriverVersion(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetImageId(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetImageName(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ImageName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetManagementFlag(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetManagementFlags(v []*string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ManagementFlags = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetMemberEniIp(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.MemberEniIp = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetOsType(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetPrimaryEniIp(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.PrimaryEniIp = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetProtocolType(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetReleaseTime(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetResetTime(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ResetTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetSystemDiskSize(v int32) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) Validate() error {
	return dara.Validate(s)
}
