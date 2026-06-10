// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDesktopGroupDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktops(v *GetDesktopGroupDetailResponseBodyDesktops) *GetDesktopGroupDetailResponseBody
	GetDesktops() *GetDesktopGroupDetailResponseBodyDesktops
	SetRequestId(v string) *GetDesktopGroupDetailResponseBody
	GetRequestId() *string
}

type GetDesktopGroupDetailResponseBody struct {
	// Information about the cloud computer share.
	Desktops *GetDesktopGroupDetailResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 1B5268CE-5EB3-545F-9F38-A8BCF710****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDesktopGroupDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBody) GetDesktops() *GetDesktopGroupDetailResponseBodyDesktops {
	return s.Desktops
}

func (s *GetDesktopGroupDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDesktopGroupDetailResponseBody) SetDesktops(v *GetDesktopGroupDetailResponseBodyDesktops) *GetDesktopGroupDetailResponseBody {
	s.Desktops = v
	return s
}

func (s *GetDesktopGroupDetailResponseBody) SetRequestId(v string) *GetDesktopGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBody) Validate() error {
	if s.Desktops != nil {
		if err := s.Desktops.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDesktopGroupDetailResponseBodyDesktops struct {
	// Whether to allow automatic creation of subscription cloud computers.
	//
	// example:
	//
	// 1
	AllowAutoSetup *int32 `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	// Number of pre-started, idle cloud computers reserved for immediate connection. Applies only to pay-as-you-go cloud computers. Valid values:
	//
	// - 0: No reservation
	//
	// - N: Reserve N cloud computer(s) (1 ≤ N ≤ 100)
	//
	// example:
	//
	// 1
	AllowBufferCount *int32 `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	// Maximum concurrent sessions per cloud computer in multi-session, multi-cloud computer deployments.
	//
	// example:
	//
	// 2
	BindAmount *int32 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	// Initial purchase count for subscription cloud computers. Valid values: 0 to 200.
	//
	// example:
	//
	// 5
	BuyDesktopsCount *int32 `json:"BuyDesktopsCount,omitempty" xml:"BuyDesktopsCount,omitempty"`
	// Remarks.
	//
	// example:
	//
	// comment
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// Maximum time a session remains connected. The session disconnects automatically when this duration is reached. Unit: milliseconds.
	//
	// example:
	//
	// 3600000
	ConnectDuration *int64 `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	// vCPU count.
	//
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2020-11-06T08:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Alibaba Cloud account ID of the creator.
	//
	// example:
	//
	// 155177335370****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// User disk type.
	//
	// example:
	//
	// cloud_essd
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// User disk capacity in GiB.
	//
	// example:
	//
	// 80
	DataDiskSize *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// Cloud computer share ID.
	//
	// example:
	//
	// dg-3uiojcc0j4kh7****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The name of the cloud computer share that you want to query.
	//
	// example:
	//
	// CloudComputerPool01
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// Directory ID (office network ID).
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Directory type.
	//
	// example:
	//
	// SIMPLE
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	EnvId         *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvType       *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Expiration time for subscription cloud computers.
	//
	// example:
	//
	// 2021-12-31T15:59Z
	ExpiredTime  *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ExpiredTimes []*string `json:"ExpiredTimes,omitempty" xml:"ExpiredTimes,omitempty" type:"Repeated"`
	// Number of GPU cores.
	//
	// example:
	//
	// 4
	GpuCount *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// GPU specification.
	//
	// example:
	//
	// NVIDIA T4
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// Time after which an idle session disconnects. If no keyboard or mouse activity occurs during this period, the session disconnects. Unit: milliseconds.
	//
	// example:
	//
	// 900000
	IdleDisconnectDuration *int64 `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	// Image ID.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Time to keep a disconnected session active. Unit: milliseconds. Valid range: 180000 (3 minutes) to 345600000 (4 days). A value of 0 means keep indefinitely.
	//
	// If a session disconnects due to user action or other reasons, the timer starts at disconnection. If no reconnection occurs within this duration, the session logs off and unsaved data is destroyed. If the user reconnects successfully within this duration, they resume the original session and access all previously saved data.
	//
	// example:
	//
	// 180000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// Traffic steering policy for multi-session, multi-cloud computer deployments.
	//
	// example:
	//
	// 0
	LoadPolicy *int32 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// - For pay-as-you-go cloud computers, this is the maximum number of cloud computers that can be created.
	//
	// - For subscription cloud computers, this is the sum of the initial purchase count (`BuyDesktopsCount`) and the number of cloud computers allowed for automatic creation.
	//
	// example:
	//
	// 3
	MaxDesktopsCount *int32 `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	// Memory size in MiB.
	//
	// example:
	//
	// 4096
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// - For pay-as-you-go cloud computers, this is the minimum number of cloud computers to create.
	//
	// - For subscription cloud computers, this equals `BuyDesktopsCount`, the initial purchase count.
	//
	// example:
	//
	// 1
	MinDesktopsCount *int32 `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	// NAS file system ID used for user profile roaming.
	//
	// example:
	//
	// 0783b4****
	NasFileSystemID *string `json:"NasFileSystemID,omitempty" xml:"NasFileSystemID,omitempty"`
	// NAS file system name used for user profile roaming.
	//
	// example:
	//
	// abcd
	NasFileSystemName *string `json:"NasFileSystemName,omitempty" xml:"NasFileSystemName,omitempty"`
	// Office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-990541****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// Name of the office network where the cloud computer share resides.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// Account system type of the office network.
	//
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	OsType         *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// Cloud computer template ID.
	//
	// example:
	//
	// b-1se9fb37r5tfq****
	OwnBundleId *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	// Cloud computer template name.
	//
	// example:
	//
	// BundleDemo
	OwnBundleName *string `json:"OwnBundleName,omitempty" xml:"OwnBundleName,omitempty"`
	// Cloud computer share type.
	//
	// example:
	//
	// 0
	OwnType *int32 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// Billing method.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// ID of the policy associated with the cloud computer share.
	//
	// example:
	//
	// pg-9cktlowtxfl6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The IDs of the policies that are associated with the cloud computer share.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// The name of the policy that is associated with the cloud computer share.
	//
	// example:
	//
	// All enabled policy
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// A list of policy names associated with cloud computer share.
	PolicyGroupNames []*string `json:"PolicyGroupNames,omitempty" xml:"PolicyGroupNames,omitempty" type:"Repeated"`
	// Whether to enable user profile roaming.
	//
	// example:
	//
	// false
	ProfileFollowSwitch *bool   `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
	ProtocolType        *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// Session occupancy threshold used to trigger auto scaling for multi-session, multi-cloud computer deployments. Session occupancy is calculated as:
	//
	// `Session occupancy = (Bound sessions / (Total cloud computers × Max sessions per cloud computer)) × 100%`
	//
	// When occupancy reaches this threshold, new cloud computers are created. When occupancy falls below this threshold, excess cloud computers are deleted.
	//
	// example:
	//
	// 0.85
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// Resource type. Only Elastic Compute Service (ECS) is supported.
	//
	// example:
	//
	// 0
	ResType *int32 `json:"ResType,omitempty" xml:"ResType,omitempty"`
	// Cloud computer reset type.
	//
	// example:
	//
	// 0
	ResetType *int32 `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// Scheduled scaling task information.
	ScaleTimerInfos []*GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos `json:"ScaleTimerInfos,omitempty" xml:"ScaleTimerInfos,omitempty" type:"Repeated"`
	// Cloud computer share status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Idle shutdown time. The cloud computer shuts down automatically after being idle for this duration. If a user connects after shutdown, the cloud computer starts automatically. Unit: milliseconds.
	//
	// example:
	//
	// 180000
	StopDuration *int64 `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
	// System disk type.
	//
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// System disk capacity in GiB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// List of scheduled tasks.
	TimerInfos []*GetDesktopGroupDetailResponseBodyDesktopsTimerInfos `json:"TimerInfos,omitempty" xml:"TimerInfos,omitempty" type:"Repeated"`
	// Scheduled application information.
	//
	// example:
	//
	// abcd
	TimingStrategyInfo *string `json:"TimingStrategyInfo,omitempty" xml:"TimingStrategyInfo,omitempty"`
	// The version number of the cloud computer share.
	//
	// example:
	//
	// 2
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetDesktopGroupDetailResponseBodyDesktops) String() string {
	return dara.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetAllowAutoSetup() *int32 {
	return s.AllowAutoSetup
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetAllowBufferCount() *int32 {
	return s.AllowBufferCount
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetBindAmount() *int32 {
	return s.BindAmount
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetBuyDesktopsCount() *int32 {
	return s.BuyDesktopsCount
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetComments() *string {
	return s.Comments
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetConnectDuration() *int64 {
	return s.ConnectDuration
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetCreator() *string {
	return s.Creator
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetDataDiskSize() *string {
	return s.DataDiskSize
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetEnvId() *string {
	return s.EnvId
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetExpiredTimes() []*string {
	return s.ExpiredTimes
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetGpuCount() *float32 {
	return s.GpuCount
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetIdleDisconnectDuration() *int64 {
	return s.IdleDisconnectDuration
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetImageId() *string {
	return s.ImageId
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetKeepDuration() *int64 {
	return s.KeepDuration
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetLoadPolicy() *int32 {
	return s.LoadPolicy
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetMaxDesktopsCount() *int32 {
	return s.MaxDesktopsCount
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetMemory() *int64 {
	return s.Memory
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetMinDesktopsCount() *int32 {
	return s.MinDesktopsCount
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetNasFileSystemID() *string {
	return s.NasFileSystemID
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetNasFileSystemName() *string {
	return s.NasFileSystemName
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetOsType() *string {
	return s.OsType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetOwnBundleId() *string {
	return s.OwnBundleId
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetOwnBundleName() *string {
	return s.OwnBundleName
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetOwnType() *int32 {
	return s.OwnType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetPayType() *string {
	return s.PayType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetPolicyGroupIds() []*string {
	return s.PolicyGroupIds
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetPolicyGroupNames() []*string {
	return s.PolicyGroupNames
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetProfileFollowSwitch() *bool {
	return s.ProfileFollowSwitch
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetRatioThreshold() *float32 {
	return s.RatioThreshold
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetResType() *int32 {
	return s.ResType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetResetType() *int32 {
	return s.ResetType
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetScaleTimerInfos() []*GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	return s.ScaleTimerInfos
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetStatus() *int32 {
	return s.Status
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetStopDuration() *int64 {
	return s.StopDuration
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetTimerInfos() []*GetDesktopGroupDetailResponseBodyDesktopsTimerInfos {
	return s.TimerInfos
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetTimingStrategyInfo() *string {
	return s.TimingStrategyInfo
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) GetVersion() *int64 {
	return s.Version
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetAllowAutoSetup(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.AllowAutoSetup = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetAllowBufferCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.AllowBufferCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetBindAmount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.BindAmount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetBuyDesktopsCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.BuyDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetComments(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Comments = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetConnectDuration(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ConnectDuration = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCpu(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCreationTime(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCreator(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Creator = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDataDiskCategory(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDataDiskSize(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDesktopGroupId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDesktopGroupName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DesktopGroupName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDirectoryId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDirectoryType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DirectoryType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetEnvId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.EnvId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetEnvType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.EnvType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetExpiredTime(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetExpiredTimes(v []*string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ExpiredTimes = v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetGpuCount(v float32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.GpuCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetGpuSpec(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.GpuSpec = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetIdleDisconnectDuration(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.IdleDisconnectDuration = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetImageId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetKeepDuration(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.KeepDuration = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetLoadPolicy(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.LoadPolicy = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMaxDesktopsCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.MaxDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMemory(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMinDesktopsCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.MinDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetNasFileSystemID(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.NasFileSystemID = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetNasFileSystemName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.NasFileSystemName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOsType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OsType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOwnBundleId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OwnBundleId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOwnBundleName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OwnBundleName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOwnType(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OwnType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPayType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PayType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPolicyGroupId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPolicyGroupIds(v []*string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PolicyGroupIds = v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPolicyGroupName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PolicyGroupName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPolicyGroupNames(v []*string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PolicyGroupNames = v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetProfileFollowSwitch(v bool) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ProfileFollowSwitch = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetProtocolType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ProtocolType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetRatioThreshold(v float32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.RatioThreshold = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetResType(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ResType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetResetType(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ResetType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetScaleTimerInfos(v []*GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ScaleTimerInfos = v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetStatus(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Status = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetStopDuration(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.StopDuration = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetSystemDiskCategory(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetSystemDiskSize(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetTimerInfos(v []*GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) *GetDesktopGroupDetailResponseBodyDesktops {
	s.TimerInfos = v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetTimingStrategyInfo(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.TimingStrategyInfo = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetVersion(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Version = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) Validate() error {
	if s.ScaleTimerInfos != nil {
		for _, item := range s.ScaleTimerInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TimerInfos != nil {
		for _, item := range s.TimerInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos struct {
	// Scaling parameter: Number of cloud computers to buy. Valid values: 0 to 200.
	//
	// example:
	//
	// 5
	BuyResAmount *int32 `json:"BuyResAmount,omitempty" xml:"BuyResAmount,omitempty"`
	// Cron expression for the scheduled scaling task.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// Time to keep a disconnected session active. Unit: milliseconds. Valid range: 180000 (3 minutes) to 345600000 (4 days). A value of 0 means keep indefinitely.
	//
	// If a session disconnects due to user action or other reasons, the timer starts at disconnection. If no reconnection occurs within this duration, the session logs off and unsaved data is destroyed. If the user reconnects successfully within this duration, they resume the original session and access all previously saved data.
	//
	// example:
	//
	// 180000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// Traffic steering policy for multi-session, multi-cloud computer deployments.
	//
	// example:
	//
	// 1
	LoadPolicy *int32 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// Scaling parameter: Maximum number of cloud computers. Valid values: 0 to 200.
	//
	// example:
	//
	// 100
	MaxResAmount *int32 `json:"MaxResAmount,omitempty" xml:"MaxResAmount,omitempty"`
	// Scaling parameter: Minimum number of cloud computers. Valid values: 0 to 200.
	//
	// example:
	//
	// 1
	MinResAmount *int32 `json:"MinResAmount,omitempty" xml:"MinResAmount,omitempty"`
	// Session occupancy threshold used to trigger auto scaling for multi-session, multi-cloud computer deployments. Session occupancy is calculated as:
	//
	// `Session occupancy = (Bound sessions / (Total cloud computers × Max sessions per cloud computer)) × 100%`
	//
	// When occupancy reaches this threshold, new cloud computers are created. When occupancy falls below this threshold, excess cloud computers are deleted.
	//
	// example:
	//
	// 0.85
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// Scheduled scaling task type.
	//
	// example:
	//
	// rise
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) String() string {
	return dara.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GetBuyResAmount() *int32 {
	return s.BuyResAmount
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GetCron() *string {
	return s.Cron
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GetKeepDuration() *int64 {
	return s.KeepDuration
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GetLoadPolicy() *int32 {
	return s.LoadPolicy
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GetMaxResAmount() *int32 {
	return s.MaxResAmount
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GetMinResAmount() *int32 {
	return s.MinResAmount
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GetRatioThreshold() *float32 {
	return s.RatioThreshold
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) GetType() *string {
	return s.Type
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) SetBuyResAmount(v int32) *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	s.BuyResAmount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) SetCron(v string) *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	s.Cron = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) SetKeepDuration(v int64) *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	s.KeepDuration = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) SetLoadPolicy(v int32) *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	s.LoadPolicy = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) SetMaxResAmount(v int32) *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	s.MaxResAmount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) SetMinResAmount(v int32) *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	s.MinResAmount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) SetRatioThreshold(v float32) *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	s.RatioThreshold = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) SetType(v string) *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos {
	s.Type = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsScaleTimerInfos) Validate() error {
	return dara.Validate(s)
}

type GetDesktopGroupDetailResponseBodyDesktopsTimerInfos struct {
	// Cron expression.
	//
	// example:
	//
	// 0 58 11 ? 	- 2
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Whether to force execute this scheduled task.
	//
	// example:
	//
	// false
	Forced *bool `json:"Forced,omitempty" xml:"Forced,omitempty"`
	// Status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Scheduled task type.
	//
	// example:
	//
	// 1
	TimerType *int32 `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) String() string {
	return dara.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) GetForced() *bool {
	return s.Forced
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) GetStatus() *int32 {
	return s.Status
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) GetTimerType() *int32 {
	return s.TimerType
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) SetCronExpression(v string) *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos {
	s.CronExpression = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) SetForced(v bool) *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos {
	s.Forced = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) SetStatus(v int32) *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos {
	s.Status = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) SetTimerType(v int32) *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos {
	s.TimerType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktopsTimerInfos) Validate() error {
	return dara.Validate(s)
}
