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
	// Specifies whether to enable auto-creation of cloud computers for the subscription cloud computer share. You must specify this parameter when `ChargeType` is set to `PrePaid`.
	//
	// Valid values:
	//
	// 	- 0: disable auto-creation of cloud computers.
	//
	// 	- 1: enables auto-creation of cloud computers.
	//
	// example:
	//
	// 1
	AllowAutoSetup *int32 `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	// The maximum number of standby cloud computers that can be reserved within the pay-as-you-go cloud computer share. You must specify this property only when `ChargeType` is set to `PostPaid`. Valid values:
	//
	// 	- 0: does not reserve any cloud computer.
	//
	// 	- N: reserves N cloud computers (1≤ N ≤ 100).
	//
	// >  Setting this parameter to 0 means no cloud computers will be reserved within the cloud computer share. In this case, the system must create, start, and assign cloud computers to end users upon request, which can be time-consuming. To improve user experience, we recommend that you reserve a specific number of cloud computers.
	//
	// example:
	//
	// 1
	AllowBufferCount *int32 `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	// The number of concurrent sessions allowed for each cloud computer within the multi-session many-to-many share.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 1
	BindAmount *int64 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	// 	- For subscription cloud computer shares, this parameter specifies the number of purchased cloud computers. Valid values: 0 to 200.
	//
	// 	- For pay-as-you-go cloud computer shares, this parameter specifies the minimum number of cloud computers created in the initial batch. Default value: 1. Valid values: 0 to `MaxDesktopsCount`.
	//
	// example:
	//
	// 5
	BuyDesktopsCount *int32 `json:"BuyDesktopsCount,omitempty" xml:"BuyDesktopsCount,omitempty"`
	// The type of the cloud computer share.
	//
	// >  This parameter is not publicly available.
	//
	// Valid values:
	//
	// 	- teacher: teacher-oriented.
	//
	// 	- student: student-oriented.
	//
	// example:
	//
	// teacher
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The maximum period of time during which the session is connected. When the specified maximum period of time is reached, the session is automatically disconnected. Unit: milliseconds. Valid values: 900000 to 345600000. That is, the session can be connected for 15 to 5,760 minutes (4 days).
	//
	// example:
	//
	// 600000
	ConnectDuration *int64 `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	DeleteDuration  *int64 `json:"DeleteDuration,omitempty" xml:"DeleteDuration,omitempty"`
	// The ID of the cloud computer share.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The name of the cloud computer share.
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
	// The ID of the File Storage NAS (NAS) file system for the user data roaming feature.
	//
	// >  This parameter is unavailable.
	//
	// example:
	//
	// 04f314****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// After an end user connects to a cloud computer, the session is established. If the system does not detect inputs from the keyboard or mouse within the specified period of time, the session is closed. Unit: milliseconds. Valid values: 360000 to 3600000 (6 minutes to 60 minutes)
	//
	// End users can receive a prompt to save data before sessions are disconnected. The system sends the prompt 30 seconds before the specified period of time is reached. To prevent data loss, end users must save the data of the sessions.
	//
	// >  This parameter is suitable only for cloud computers whose image version is v1.0.2 or later.
	//
	// example:
	//
	// 120000
	IdleDisconnectDuration *int64 `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	// The IDs of the images.
	//
	// example:
	//
	// desktopimage-windows-server-2016-64-ch
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The retention period of a session after it is disconnected. Unit: milliseconds. Valid values: 180000 to 345600000. That is, the session can be retained for 3 to 5,760 minutes (4 days) after it is disconnected. If you set this parameter to 0, the session is permanently retained after it is disconnected.
	//
	// When a session is disconnected, take note of the following situations: If an end user does not resume the session within the specified duration, the session is closed and all unsaved data is cleared. If the end user resumes the session within the specified duration, the end user can continue to access data of the session.
	//
	// example:
	//
	// 1000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// The load balancing policy for the multi-session many-to-many share.
	//
	// >  This parameter is not publicly available.
	//
	// Valid values:
	//
	// 	- 0: depth first.
	//
	// 	- 1: breadth first.
	//
	// example:
	//
	// 0
	LoadPolicy *int64 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// The maximum number of cloud computers allowed in the pay-as-you-go cloud computer share. Valid values: 0 to 500.
	//
	// example:
	//
	// 10
	MaxDesktopsCount *int32 `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	// The maximum number of auto-created cloud computers allowed in the subscription cloud computer share. You must specify this parameter when `ChargeType` is set to `PrePaid`. Default value: 1. Valid values: 0 to `MaxDesktopsCount`.
	//
	// example:
	//
	// 1
	MinDesktopsCount *int32 `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	// The ID of the cloud computer template.
	//
	// example:
	//
	// b-7t275tpgjueeu****
	OwnBundleId *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	// The ID of the security policy.
	//
	// example:
	//
	// pg-53iyi2aar0nd6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The IDs of policy groups.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to enable user data roaming.
	//
	// >  This parameter is unavailable.
	//
	// example:
	//
	// false
	ProfileFollowSwitch *bool `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
	// The threshold for the ratio of connected sessions, which triggers automatic scaling of cloud computers within the multi-session many-to-many share. To calculate the ratio of connected sessions, use the following formula:
	//
	// `Ratio of connected sessions = Number of connected sessions/(Total number of cloud computers × Maximum number of sessions allowed for each cloud computer) × 100%`
	//
	// If the session ratio exceeds the threshold, new cloud computers are provisioned. If it falls below the threshold, additional cloud computers are removed.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 0.5
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The disk reset type of cloud computers.
	//
	// Valid values:
	//
	// 	- 0: does not reset disks.
	//
	// 	- 1: resets only the system disks.
	//
	// 	- 2: resets only the user disks.
	//
	// 	- 3: resets the system disks and user disks.
	//
	// example:
	//
	// 0
	ResetType *int64 `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The ID of the scaling policy group.
	//
	// >  This parameter is unavailable.
	//
	// example:
	//
	// s-kakowkdl****
	ScaleStrategyId *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	// The period of time before the idle cloud computer enters the Stopped state. When the specified period of time is reached, the cloud computer is automatically stopped. If an end user connects to the stopped cloud computer, the cloud computer automatically starts. Unit: milliseconds.
	//
	// example:
	//
	// 180000
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
