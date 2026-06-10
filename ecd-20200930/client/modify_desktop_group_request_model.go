// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowAutoSetup(v int32) *ModifyDesktopGroupRequest
	GetAllowAutoSetup() *int32
	SetAllowBufferCount(v int32) *ModifyDesktopGroupRequest
	GetAllowBufferCount() *int32
	SetBindAmount(v int64) *ModifyDesktopGroupRequest
	GetBindAmount() *int64
	SetBuyDesktopsCount(v int32) *ModifyDesktopGroupRequest
	GetBuyDesktopsCount() *int32
	SetClassify(v string) *ModifyDesktopGroupRequest
	GetClassify() *string
	SetComments(v string) *ModifyDesktopGroupRequest
	GetComments() *string
	SetConnectDuration(v int64) *ModifyDesktopGroupRequest
	GetConnectDuration() *int64
	SetDeleteDuration(v int64) *ModifyDesktopGroupRequest
	GetDeleteDuration() *int64
	SetDesktopGroupId(v string) *ModifyDesktopGroupRequest
	GetDesktopGroupId() *string
	SetDesktopGroupName(v string) *ModifyDesktopGroupRequest
	GetDesktopGroupName() *string
	SetDisableSessionConfig(v bool) *ModifyDesktopGroupRequest
	GetDisableSessionConfig() *bool
	SetFileSystemId(v string) *ModifyDesktopGroupRequest
	GetFileSystemId() *string
	SetIdleDisconnectDuration(v int64) *ModifyDesktopGroupRequest
	GetIdleDisconnectDuration() *int64
	SetImageId(v string) *ModifyDesktopGroupRequest
	GetImageId() *string
	SetKeepDuration(v int64) *ModifyDesktopGroupRequest
	GetKeepDuration() *int64
	SetLoadPolicy(v int64) *ModifyDesktopGroupRequest
	GetLoadPolicy() *int64
	SetMaxDesktopsCount(v int32) *ModifyDesktopGroupRequest
	GetMaxDesktopsCount() *int32
	SetMinDesktopsCount(v int32) *ModifyDesktopGroupRequest
	GetMinDesktopsCount() *int32
	SetOwnBundleId(v string) *ModifyDesktopGroupRequest
	GetOwnBundleId() *string
	SetPolicyGroupId(v string) *ModifyDesktopGroupRequest
	GetPolicyGroupId() *string
	SetPolicyGroupIds(v []*string) *ModifyDesktopGroupRequest
	GetPolicyGroupIds() []*string
	SetProfileFollowSwitch(v bool) *ModifyDesktopGroupRequest
	GetProfileFollowSwitch() *bool
	SetRatioThreshold(v float32) *ModifyDesktopGroupRequest
	GetRatioThreshold() *float32
	SetRegionId(v string) *ModifyDesktopGroupRequest
	GetRegionId() *string
	SetResetType(v int64) *ModifyDesktopGroupRequest
	GetResetType() *int64
	SetScaleStrategyId(v string) *ModifyDesktopGroupRequest
	GetScaleStrategyId() *string
	SetStopDuration(v int64) *ModifyDesktopGroupRequest
	GetStopDuration() *int64
}

type ModifyDesktopGroupRequest struct {
	// Specifies whether to enable automatic creation of cloud computers in a subscription shared cloud computer group. This parameter is required and takes effect only when `ChargeType` is set to `PrePaid`.
	//
	// example:
	//
	// 1
	AllowAutoSetup *int32 `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	// The number of cloud computers to reserve in a pay-as-you-go shared cloud computer group. This parameter is required and takes effect only when `ChargeType` is set to `PostPaid`. Valid values:
	//
	// - 0: No cloud computers are reserved.
	//
	// - N: N cloud computers are reserved (1 ≤ N ≤ 100).
	//
	// > If you do not reserve any cloud computers, the system must create and start one when an end user requests a connection. This process takes longer. Reserve a specific number of cloud computers to ensure a good user experience.
	//
	// example:
	//
	// 1
	AllowBufferCount *int32 `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	// The number of concurrent sessions that each cloud computer in a multi-session shared cloud computer group can support.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 2
	BindAmount *int64 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	// - For a subscription shared cloud computer group: the number of cloud computers to purchase. Valid values: 0 to 200.
	//
	// - For a pay-as-you-go shared cloud computer group: the minimum number of cloud computers to create in the pool. Default value: 1. Valid values: 0 to the value of `MaxDesktopsCount`.
	//
	// example:
	//
	// 5
	BuyDesktopsCount *int32 `json:"BuyDesktopsCount,omitempty" xml:"BuyDesktopsCount,omitempty"`
	// The type of the shared cloud computer group.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// teacher
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The remarks.
	//
	// example:
	//
	// comment
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The maximum duration of a session. When the session duration reaches this value, the session is automatically disconnected. Unit: milliseconds. Valid values: 900000 (15 minutes) to 345600000 (4 days).
	//
	// example:
	//
	// 900000
	ConnectDuration *int64 `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	DeleteDuration  *int64 `json:"DeleteDuration,omitempty" xml:"DeleteDuration,omitempty"`
	// The ID of the shared cloud computer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// Shared cloud desktop name.
	//
	// example:
	//
	// desktopGroupName1
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// Specifies whether to disable session management.
	//
	// example:
	//
	// true
	DisableSessionConfig *bool `json:"DisableSessionConfig,omitempty" xml:"DisableSessionConfig,omitempty"`
	// The ID of the NAS file system used for user data roaming.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 04f314****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum idle time for a session. If there is no keyboard or mouse input within this time, the session disconnects. Unit: milliseconds. Valid values: 360000 (6 minutes) to 3600000 (60 minutes).
	//
	// Thirty seconds before the session disconnects, the end user receives a message to save their data. The end user must save their data to prevent data loss.
	//
	// > This parameter is applicable only to cloud computers with an image version of 1.0.2 or later.
	//
	// example:
	//
	// 360000
	IdleDisconnectDuration *int64 `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	// The image ID.
	//
	// example:
	//
	// desktopimage-windows-server-2016-64-ch
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The duration to keep a session active after it disconnects. Unit: milliseconds. Valid values range from 180000 (3 minutes) to 345600000 (4 days). A value of 0 means the session is always kept active.
	//
	// When a session disconnects, either intentionally or unexpectedly, a timer begins. If the user fails to reconnect within this duration, the session is logged off, and any unsaved data is destroyed. If the user reconnects within this duration, they can resume the original session and access the data from before the disconnection.
	//
	// example:
	//
	// 180000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// The load balancing policy for a multi-session shared cloud computer group that contains multiple cloud computers.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 0
	LoadPolicy *int64 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// The maximum number of cloud computers that a pay-as-you-go shared cloud computer group can contain. Valid values: 0 to 500.
	//
	// example:
	//
	// 10
	MaxDesktopsCount *int32 `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	// The maximum number of cloud computers that are automatically created in a subscription shared cloud computer group. This parameter is required and takes effect only when `ChargeType` is set to `PrePaid`. Default value: 1. Valid values: 0 to the value of `MaxDesktopsCount`.
	//
	// example:
	//
	// 1
	MinDesktopsCount *int32 `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	// The cloud computer template ID.
	//
	// example:
	//
	// b-7t275tpgjueeu****
	OwnBundleId *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-53iyi2aar0nd6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The list of policy group IDs.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to enable user data roaming.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// false
	ProfileFollowSwitch *bool `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
	// The session usage threshold. This threshold is a condition for triggering auto scaling in a multi-session shared cloud computer group. The session usage is calculated using the following formula:
	//
	// `Session usage = Number of active sessions / (Total number of cloud computers × Maximum number of sessions per cloud computer) × 100%`
	//
	// When the session usage reaches this threshold, new cloud computers are created. If the session usage is below this threshold, idle cloud computers are deleted.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 0.85
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to get a list of regions that WUYING Workspace supports.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reset type for the cloud computers.
	//
	// example:
	//
	// 0
	ResetType *int64 `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The ID of the auto scaling policy group.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// s-kakowkdl****
	ScaleStrategyId *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	// The idle shutdown time. The cloud computer automatically shuts down when it is idle for this amount of time. If a user connects to a shutdown cloud computer, it automatically starts. Unit: milliseconds.
	//
	// example:
	//
	// 300000
	StopDuration *int64 `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
}

func (s ModifyDesktopGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupRequest) GetAllowAutoSetup() *int32 {
	return s.AllowAutoSetup
}

func (s *ModifyDesktopGroupRequest) GetAllowBufferCount() *int32 {
	return s.AllowBufferCount
}

func (s *ModifyDesktopGroupRequest) GetBindAmount() *int64 {
	return s.BindAmount
}

func (s *ModifyDesktopGroupRequest) GetBuyDesktopsCount() *int32 {
	return s.BuyDesktopsCount
}

func (s *ModifyDesktopGroupRequest) GetClassify() *string {
	return s.Classify
}

func (s *ModifyDesktopGroupRequest) GetComments() *string {
	return s.Comments
}

func (s *ModifyDesktopGroupRequest) GetConnectDuration() *int64 {
	return s.ConnectDuration
}

func (s *ModifyDesktopGroupRequest) GetDeleteDuration() *int64 {
	return s.DeleteDuration
}

func (s *ModifyDesktopGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *ModifyDesktopGroupRequest) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *ModifyDesktopGroupRequest) GetDisableSessionConfig() *bool {
	return s.DisableSessionConfig
}

func (s *ModifyDesktopGroupRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyDesktopGroupRequest) GetIdleDisconnectDuration() *int64 {
	return s.IdleDisconnectDuration
}

func (s *ModifyDesktopGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyDesktopGroupRequest) GetKeepDuration() *int64 {
	return s.KeepDuration
}

func (s *ModifyDesktopGroupRequest) GetLoadPolicy() *int64 {
	return s.LoadPolicy
}

func (s *ModifyDesktopGroupRequest) GetMaxDesktopsCount() *int32 {
	return s.MaxDesktopsCount
}

func (s *ModifyDesktopGroupRequest) GetMinDesktopsCount() *int32 {
	return s.MinDesktopsCount
}

func (s *ModifyDesktopGroupRequest) GetOwnBundleId() *string {
	return s.OwnBundleId
}

func (s *ModifyDesktopGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyDesktopGroupRequest) GetPolicyGroupIds() []*string {
	return s.PolicyGroupIds
}

func (s *ModifyDesktopGroupRequest) GetProfileFollowSwitch() *bool {
	return s.ProfileFollowSwitch
}

func (s *ModifyDesktopGroupRequest) GetRatioThreshold() *float32 {
	return s.RatioThreshold
}

func (s *ModifyDesktopGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDesktopGroupRequest) GetResetType() *int64 {
	return s.ResetType
}

func (s *ModifyDesktopGroupRequest) GetScaleStrategyId() *string {
	return s.ScaleStrategyId
}

func (s *ModifyDesktopGroupRequest) GetStopDuration() *int64 {
	return s.StopDuration
}

func (s *ModifyDesktopGroupRequest) SetAllowAutoSetup(v int32) *ModifyDesktopGroupRequest {
	s.AllowAutoSetup = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetAllowBufferCount(v int32) *ModifyDesktopGroupRequest {
	s.AllowBufferCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetBindAmount(v int64) *ModifyDesktopGroupRequest {
	s.BindAmount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetBuyDesktopsCount(v int32) *ModifyDesktopGroupRequest {
	s.BuyDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetClassify(v string) *ModifyDesktopGroupRequest {
	s.Classify = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetComments(v string) *ModifyDesktopGroupRequest {
	s.Comments = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetConnectDuration(v int64) *ModifyDesktopGroupRequest {
	s.ConnectDuration = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDeleteDuration(v int64) *ModifyDesktopGroupRequest {
	s.DeleteDuration = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDesktopGroupId(v string) *ModifyDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDesktopGroupName(v string) *ModifyDesktopGroupRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDisableSessionConfig(v bool) *ModifyDesktopGroupRequest {
	s.DisableSessionConfig = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetFileSystemId(v string) *ModifyDesktopGroupRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetIdleDisconnectDuration(v int64) *ModifyDesktopGroupRequest {
	s.IdleDisconnectDuration = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetImageId(v string) *ModifyDesktopGroupRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetKeepDuration(v int64) *ModifyDesktopGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetLoadPolicy(v int64) *ModifyDesktopGroupRequest {
	s.LoadPolicy = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetMaxDesktopsCount(v int32) *ModifyDesktopGroupRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetMinDesktopsCount(v int32) *ModifyDesktopGroupRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetOwnBundleId(v string) *ModifyDesktopGroupRequest {
	s.OwnBundleId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetPolicyGroupIds(v []*string) *ModifyDesktopGroupRequest {
	s.PolicyGroupIds = v
	return s
}

func (s *ModifyDesktopGroupRequest) SetProfileFollowSwitch(v bool) *ModifyDesktopGroupRequest {
	s.ProfileFollowSwitch = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetRatioThreshold(v float32) *ModifyDesktopGroupRequest {
	s.RatioThreshold = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetRegionId(v string) *ModifyDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetResetType(v int64) *ModifyDesktopGroupRequest {
	s.ResetType = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetScaleStrategyId(v string) *ModifyDesktopGroupRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetStopDuration(v int64) *ModifyDesktopGroupRequest {
	s.StopDuration = &v
	return s
}

func (s *ModifyDesktopGroupRequest) Validate() error {
	return dara.Validate(s)
}
