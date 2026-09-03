// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllClassifyUsers(v bool) *CreateDesktopGroupRequest
	GetAllClassifyUsers() *bool
	SetAllowAutoSetup(v int32) *CreateDesktopGroupRequest
	GetAllowAutoSetup() *int32
	SetAllowBufferCount(v int32) *CreateDesktopGroupRequest
	GetAllowBufferCount() *int32
	SetAutoPay(v bool) *CreateDesktopGroupRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateDesktopGroupRequest
	GetAutoRenew() *bool
	SetBindAmount(v int64) *CreateDesktopGroupRequest
	GetBindAmount() *int64
	SetBundleId(v string) *CreateDesktopGroupRequest
	GetBundleId() *string
	SetBuyDesktopsCount(v int32) *CreateDesktopGroupRequest
	GetBuyDesktopsCount() *int32
	SetChargeType(v string) *CreateDesktopGroupRequest
	GetChargeType() *string
	SetClassify(v string) *CreateDesktopGroupRequest
	GetClassify() *string
	SetClientToken(v string) *CreateDesktopGroupRequest
	GetClientToken() *string
	SetComments(v string) *CreateDesktopGroupRequest
	GetComments() *string
	SetConnectDuration(v int64) *CreateDesktopGroupRequest
	GetConnectDuration() *int64
	SetDataDiskCategory(v string) *CreateDesktopGroupRequest
	GetDataDiskCategory() *string
	SetDataDiskPerLevel(v string) *CreateDesktopGroupRequest
	GetDataDiskPerLevel() *string
	SetDataDiskSize(v int32) *CreateDesktopGroupRequest
	GetDataDiskSize() *int32
	SetDefaultInitDesktopCount(v int32) *CreateDesktopGroupRequest
	GetDefaultInitDesktopCount() *int32
	SetDefaultLanguage(v string) *CreateDesktopGroupRequest
	GetDefaultLanguage() *string
	SetDeleteDuration(v int64) *CreateDesktopGroupRequest
	GetDeleteDuration() *int64
	SetDesktopGroupName(v string) *CreateDesktopGroupRequest
	GetDesktopGroupName() *string
	SetDesktopType(v string) *CreateDesktopGroupRequest
	GetDesktopType() *string
	SetDirectoryId(v string) *CreateDesktopGroupRequest
	GetDirectoryId() *string
	SetEndUserIds(v []*string) *CreateDesktopGroupRequest
	GetEndUserIds() []*string
	SetExclusiveType(v string) *CreateDesktopGroupRequest
	GetExclusiveType() *string
	SetFileSystemId(v string) *CreateDesktopGroupRequest
	GetFileSystemId() *string
	SetGroupAmount(v int32) *CreateDesktopGroupRequest
	GetGroupAmount() *int32
	SetGroupVersion(v int32) *CreateDesktopGroupRequest
	GetGroupVersion() *int32
	SetHostname(v string) *CreateDesktopGroupRequest
	GetHostname() *string
	SetIdleDisconnectDuration(v int64) *CreateDesktopGroupRequest
	GetIdleDisconnectDuration() *int64
	SetImageId(v string) *CreateDesktopGroupRequest
	GetImageId() *string
	SetKeepDuration(v int64) *CreateDesktopGroupRequest
	GetKeepDuration() *int64
	SetLoadPolicy(v int64) *CreateDesktopGroupRequest
	GetLoadPolicy() *int64
	SetMaxDesktopsCount(v int32) *CreateDesktopGroupRequest
	GetMaxDesktopsCount() *int32
	SetMinDesktopsCount(v int32) *CreateDesktopGroupRequest
	GetMinDesktopsCount() *int32
	SetMultiResource(v bool) *CreateDesktopGroupRequest
	GetMultiResource() *bool
	SetOfficeSiteId(v string) *CreateDesktopGroupRequest
	GetOfficeSiteId() *string
	SetOwnType(v int32) *CreateDesktopGroupRequest
	GetOwnType() *int32
	SetPeriod(v int32) *CreateDesktopGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateDesktopGroupRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateDesktopGroupRequest
	GetPolicyGroupId() *string
	SetProfileFollowSwitch(v bool) *CreateDesktopGroupRequest
	GetProfileFollowSwitch() *bool
	SetPromotionId(v string) *CreateDesktopGroupRequest
	GetPromotionId() *string
	SetRatioThreshold(v float32) *CreateDesktopGroupRequest
	GetRatioThreshold() *float32
	SetRegionId(v string) *CreateDesktopGroupRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *CreateDesktopGroupRequest
	GetResellerOwnerUid() *int64
	SetResetType(v int64) *CreateDesktopGroupRequest
	GetResetType() *int64
	SetScaleStrategyId(v string) *CreateDesktopGroupRequest
	GetScaleStrategyId() *string
	SetSessionType(v string) *CreateDesktopGroupRequest
	GetSessionType() *string
	SetSimpleUserGroupId(v string) *CreateDesktopGroupRequest
	GetSimpleUserGroupId() *string
	SetSnapshotPolicyId(v string) *CreateDesktopGroupRequest
	GetSnapshotPolicyId() *string
	SetStopDuration(v int64) *CreateDesktopGroupRequest
	GetStopDuration() *int64
	SetSystemDiskCategory(v string) *CreateDesktopGroupRequest
	GetSystemDiskCategory() *string
	SetSystemDiskPerLevel(v string) *CreateDesktopGroupRequest
	GetSystemDiskPerLevel() *string
	SetSystemDiskSize(v int32) *CreateDesktopGroupRequest
	GetSystemDiskSize() *int32
	SetTag(v []*CreateDesktopGroupRequestTag) *CreateDesktopGroupRequest
	GetTag() []*CreateDesktopGroupRequestTag
	SetTimerGroupId(v string) *CreateDesktopGroupRequest
	GetTimerGroupId() *string
	SetUserGroupName(v string) *CreateDesktopGroupRequest
	GetUserGroupName() *string
	SetUserOuPath(v string) *CreateDesktopGroupRequest
	GetUserOuPath() *string
	SetVolumeEncryptionEnabled(v bool) *CreateDesktopGroupRequest
	GetVolumeEncryptionEnabled() *bool
	SetVolumeEncryptionKey(v string) *CreateDesktopGroupRequest
	GetVolumeEncryptionKey() *string
	SetVpcId(v string) *CreateDesktopGroupRequest
	GetVpcId() *string
}

type CreateDesktopGroupRequest struct {
	// The users of all shared cloud computer categories.
	//
	// example:
	//
	// true
	AllClassifyUsers *bool `json:"AllClassifyUsers,omitempty" xml:"AllClassifyUsers,omitempty"`
	// Specifies whether to allow automatic creation of cloud computers within subscription shared cloud computers. This parameter takes effect and is required only when ChargeType is set to PrePaid.
	//
	// example:
	//
	// 1
	AllowAutoSetup *int32 `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	// The number of reserved cloud computers allowed in pay-as-you-go shared cloud computers. This parameter takes effect and is required only when ChargeType is set to PostPaid. Valid values:
	//
	// example:
	//
	// 1
	AllowBufferCount *int32 `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	// Specifies whether automatic payment is enabled for the subscription order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the subscription shared cloud computer.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The number of concurrent sessions allowed per cloud computer in multi-session shared cloud computers.
	//
	// example:
	//
	// 2
	BindAmount *int64 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	// The cloud computer template ID.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// - For subscription shared cloud computers: the initial number of cloud computers to create. Valid values: 0 to 200.
	//
	// example:
	//
	// 3
	BuyDesktopsCount *int32 `json:"BuyDesktopsCount,omitempty" xml:"BuyDesktopsCount,omitempty"`
	// The billing method of the cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The type of the shared cloud computer.
	//
	// example:
	//
	// teacher
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The remarks.
	//
	// example:
	//
	// comment
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The maximum duration that a session can remain in the connected state. The session is automatically disconnected when this duration is reached. Unit: milliseconds. Valid values: 900000 (15 minutes) to 345600000 (4 days).
	//
	// example:
	//
	// 900000
	ConnectDuration *int64 `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	// The data cloud disk type.
	//
	// example:
	//
	// cloud_auto
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The performance level of the ESSD. Default value: PL0.
	//
	// example:
	//
	// PL0
	DataDiskPerLevel *string `json:"DataDiskPerLevel,omitempty" xml:"DataDiskPerLevel,omitempty"`
	// The size of the attached data cloud disk. Unit: GB. Valid values: 0 to 16380. The value must be a multiple of 20.
	//
	// example:
	//
	// 80
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The default number of cloud computers to create when you create multiple shared cloud computers. Default value: 1.
	//
	// example:
	//
	// 1
	DefaultInitDesktopCount *int32 `json:"DefaultInitDesktopCount,omitempty" xml:"DefaultInitDesktopCount,omitempty"`
	// The system language.
	//
	// example:
	//
	// zh-CN
	DefaultLanguage *string `json:"DefaultLanguage,omitempty" xml:"DefaultLanguage,omitempty"`
	// The retention period before cloud computers in the cloud computer pool are automatically deleted.
	//
	// example:
	//
	// 30
	DeleteDuration *int64 `json:"DeleteDuration,omitempty" xml:"DeleteDuration,omitempty"`
	// The name of the shared cloud computer. The name can be up to 30 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. The name can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// SharedComputers01
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// The cloud computer specification. You can call [DescribeDesktopTypes](~~DescribeDesktopTypes~~) to query the specification IDs supported by cloud computers.
	//
	// example:
	//
	// eds.enterprise_office.16c64g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// dri-uf62w3qzt4aigvlcb****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The list of user IDs for the shared cloud computer.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// Creates a static pool. This parameter is required when the `SessionType` parameter is set to `MultipleSession`. Set the value to `Exclusive`.
	//
	// example:
	//
	// Exclusive
	ExclusiveType *string `json:"ExclusiveType,omitempty" xml:"ExclusiveType,omitempty"`
	// The ID of the NAS file system used for user data roaming.
	//
	// example:
	//
	// kegd-nas-****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of single shared cloud computers to create. This parameter is required only when the `MultiResource` parameter is set to `false`. Valid values: 1 to 5. Default value: 1.
	//
	// example:
	//
	// 1
	GroupAmount *int32 `json:"GroupAmount,omitempty" xml:"GroupAmount,omitempty"`
	// The version of the shared cloud computer.
	//
	// example:
	//
	// 2
	GroupVersion *int32 `json:"GroupVersion,omitempty" xml:"GroupVersion,omitempty"`
	// The custom hostname of the cloud computer. Only Settings for cloud computers that run the Windows operating system in AD office networks are supported.
	//
	// example:
	//
	// testhost
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The maximum idle duration after a user session is established. If no keyboard or mouse activity occurs within this duration, the session is disconnected. Unit: milliseconds. Valid values: 360000 (6 minutes) to 3600000 (60 minutes).
	//
	// 30 seconds before this duration is reached, the end user in the session receives a prompt to save document data. The end user must save document data promptly to avoid data loss.
	//
	// > Applicable only to cloud computers with an image version of 1.0.2 or later.
	//
	// example:
	//
	// 360000
	IdleDisconnectDuration *int64 `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The retention period after a session is disconnected. Unit: milliseconds. Valid values: 180000 (3 minutes) to 345600000 (4 days). A value of 0 indicates that the session is always retained.
	//
	// example:
	//
	// 180000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// The load balancing policy for multi-session shared cloud computers.
	//
	// example:
	//
	// 0
	LoadPolicy *int64 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// The maximum number of pay-as-you-go shared cloud computers. Valid values: 0 to 500.
	//
	// example:
	//
	// 50
	MaxDesktopsCount *int32 `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	// The maximum number of cloud computers that can be used for automatic creation for subscription shared cloud computers. This parameter takes effect and is required only when ChargeType is set to PrePaid. Default value: 1. Valid values: 0 to the value of MaxDesktopsCount.
	//
	// example:
	//
	// 1
	MinDesktopsCount *int32 `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	// Specifies whether the cloud computers are multi-resource shared cloud computers.
	//
	// example:
	//
	// false
	MultiResource *bool `json:"MultiResource,omitempty" xml:"MultiResource,omitempty"`
	// The ID of the office network to which the shared cloud computer belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+os-c5cy7q578s8jc****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The type of the shared cloud computer.
	//
	// example:
	//
	// 0
	OwnType *int32 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The subscription duration of the shared cloud computer. This parameter takes effect and is required only when ChargeType is set to PrePaid. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription billable methods duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the policy associated with the shared cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-9c2d6t2dwflqr****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// Specifies whether to enable user data roaming.
	//
	// example:
	//
	// false
	ProfileFollowSwitch *bool `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_*****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The session occupancy threshold used as the automatic scaling trigger condition for multi-session shared cloud computers. The session occupancy is calculated by using the following formula:
	//
	// ```Session occupancy = Number of bound sessions / (Total number of cloud computer resources × Maximum number of sessions supported per cloud computer) × 100%```
	//
	// When the session occupancy reaches this threshold, new cloud computers are created. When the session occupancy is below this threshold, excess cloud computers are deleted.
	//
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// 0.5
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the resource ownership in reseller pattern. You do not need to specify this parameter in non-reseller pattern.
	//
	// example:
	//
	// 1422724566551XXX
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The reset type of the cloud computer.
	//
	// example:
	//
	// 0
	ResetType *int64 `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The scaling policy ID.
	//
	// example:
	//
	// ss-f9dkjz6vw3aaw****
	ScaleStrategyId *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	// The session type.
	//
	// example:
	//
	// SingleSession
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The ID of the convenience user group.
	//
	// example:
	//
	// ug-3f6c8a2b****
	SimpleUserGroupId *string `json:"SimpleUserGroupId,omitempty" xml:"SimpleUserGroupId,omitempty"`
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-28mp6my0l6zow****
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	// The idle shutdown duration. When the cloud computer has been idle for this duration, it is automatically shut down. If a user connects after shutdown, the cloud computer automatically starts. Unit: milliseconds.
	//
	// example:
	//
	// 300000
	StopDuration *int64 `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
	// The system cloud disk type.
	//
	// example:
	//
	// cloud_auto
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The performance level of the ESSD. Default value: PL0.
	//
	// example:
	//
	// PL0
	SystemDiskPerLevel *string `json:"SystemDiskPerLevel,omitempty" xml:"SystemDiskPerLevel,omitempty"`
	// The system cloud disk size. Unit: GiB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The list of tags. A maximum of 20 tags can be specified.
	Tag []*CreateDesktopGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the scheduled task group.
	//
	// example:
	//
	// ccg-0caoeogrk9m5****
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// R&D Group
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	// The organizational unit (OU) path of the user.
	//
	// example:
	//
	// example.com
	UserOuPath *string `json:"UserOuPath,omitempty" xml:"UserOuPath,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the KMS key used for disk encryption. You can call [ListKeys](https://help.aliyun.com/document_detail/28951.html) to obtain the key ID.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
	// The VPC ID of the office network to which the shared cloud computer belongs.
	//
	// example:
	//
	// vpc-uf6w8u60n8xbkg5el****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateDesktopGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupRequest) GetAllClassifyUsers() *bool {
	return s.AllClassifyUsers
}

func (s *CreateDesktopGroupRequest) GetAllowAutoSetup() *int32 {
	return s.AllowAutoSetup
}

func (s *CreateDesktopGroupRequest) GetAllowBufferCount() *int32 {
	return s.AllowBufferCount
}

func (s *CreateDesktopGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateDesktopGroupRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDesktopGroupRequest) GetBindAmount() *int64 {
	return s.BindAmount
}

func (s *CreateDesktopGroupRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateDesktopGroupRequest) GetBuyDesktopsCount() *int32 {
	return s.BuyDesktopsCount
}

func (s *CreateDesktopGroupRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateDesktopGroupRequest) GetClassify() *string {
	return s.Classify
}

func (s *CreateDesktopGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDesktopGroupRequest) GetComments() *string {
	return s.Comments
}

func (s *CreateDesktopGroupRequest) GetConnectDuration() *int64 {
	return s.ConnectDuration
}

func (s *CreateDesktopGroupRequest) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *CreateDesktopGroupRequest) GetDataDiskPerLevel() *string {
	return s.DataDiskPerLevel
}

func (s *CreateDesktopGroupRequest) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *CreateDesktopGroupRequest) GetDefaultInitDesktopCount() *int32 {
	return s.DefaultInitDesktopCount
}

func (s *CreateDesktopGroupRequest) GetDefaultLanguage() *string {
	return s.DefaultLanguage
}

func (s *CreateDesktopGroupRequest) GetDeleteDuration() *int64 {
	return s.DeleteDuration
}

func (s *CreateDesktopGroupRequest) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *CreateDesktopGroupRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *CreateDesktopGroupRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateDesktopGroupRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *CreateDesktopGroupRequest) GetExclusiveType() *string {
	return s.ExclusiveType
}

func (s *CreateDesktopGroupRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateDesktopGroupRequest) GetGroupAmount() *int32 {
	return s.GroupAmount
}

func (s *CreateDesktopGroupRequest) GetGroupVersion() *int32 {
	return s.GroupVersion
}

func (s *CreateDesktopGroupRequest) GetHostname() *string {
	return s.Hostname
}

func (s *CreateDesktopGroupRequest) GetIdleDisconnectDuration() *int64 {
	return s.IdleDisconnectDuration
}

func (s *CreateDesktopGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateDesktopGroupRequest) GetKeepDuration() *int64 {
	return s.KeepDuration
}

func (s *CreateDesktopGroupRequest) GetLoadPolicy() *int64 {
	return s.LoadPolicy
}

func (s *CreateDesktopGroupRequest) GetMaxDesktopsCount() *int32 {
	return s.MaxDesktopsCount
}

func (s *CreateDesktopGroupRequest) GetMinDesktopsCount() *int32 {
	return s.MinDesktopsCount
}

func (s *CreateDesktopGroupRequest) GetMultiResource() *bool {
	return s.MultiResource
}

func (s *CreateDesktopGroupRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateDesktopGroupRequest) GetOwnType() *int32 {
	return s.OwnType
}

func (s *CreateDesktopGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateDesktopGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateDesktopGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateDesktopGroupRequest) GetProfileFollowSwitch() *bool {
	return s.ProfileFollowSwitch
}

func (s *CreateDesktopGroupRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateDesktopGroupRequest) GetRatioThreshold() *float32 {
	return s.RatioThreshold
}

func (s *CreateDesktopGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDesktopGroupRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *CreateDesktopGroupRequest) GetResetType() *int64 {
	return s.ResetType
}

func (s *CreateDesktopGroupRequest) GetScaleStrategyId() *string {
	return s.ScaleStrategyId
}

func (s *CreateDesktopGroupRequest) GetSessionType() *string {
	return s.SessionType
}

func (s *CreateDesktopGroupRequest) GetSimpleUserGroupId() *string {
	return s.SimpleUserGroupId
}

func (s *CreateDesktopGroupRequest) GetSnapshotPolicyId() *string {
	return s.SnapshotPolicyId
}

func (s *CreateDesktopGroupRequest) GetStopDuration() *int64 {
	return s.StopDuration
}

func (s *CreateDesktopGroupRequest) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateDesktopGroupRequest) GetSystemDiskPerLevel() *string {
	return s.SystemDiskPerLevel
}

func (s *CreateDesktopGroupRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateDesktopGroupRequest) GetTag() []*CreateDesktopGroupRequestTag {
	return s.Tag
}

func (s *CreateDesktopGroupRequest) GetTimerGroupId() *string {
	return s.TimerGroupId
}

func (s *CreateDesktopGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *CreateDesktopGroupRequest) GetUserOuPath() *string {
	return s.UserOuPath
}

func (s *CreateDesktopGroupRequest) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *CreateDesktopGroupRequest) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *CreateDesktopGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDesktopGroupRequest) SetAllClassifyUsers(v bool) *CreateDesktopGroupRequest {
	s.AllClassifyUsers = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAllowAutoSetup(v int32) *CreateDesktopGroupRequest {
	s.AllowAutoSetup = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAllowBufferCount(v int32) *CreateDesktopGroupRequest {
	s.AllowBufferCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAutoPay(v bool) *CreateDesktopGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAutoRenew(v bool) *CreateDesktopGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetBindAmount(v int64) *CreateDesktopGroupRequest {
	s.BindAmount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetBundleId(v string) *CreateDesktopGroupRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetBuyDesktopsCount(v int32) *CreateDesktopGroupRequest {
	s.BuyDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetChargeType(v string) *CreateDesktopGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetClassify(v string) *CreateDesktopGroupRequest {
	s.Classify = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetClientToken(v string) *CreateDesktopGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetComments(v string) *CreateDesktopGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetConnectDuration(v int64) *CreateDesktopGroupRequest {
	s.ConnectDuration = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDataDiskCategory(v string) *CreateDesktopGroupRequest {
	s.DataDiskCategory = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDataDiskPerLevel(v string) *CreateDesktopGroupRequest {
	s.DataDiskPerLevel = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDataDiskSize(v int32) *CreateDesktopGroupRequest {
	s.DataDiskSize = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDefaultInitDesktopCount(v int32) *CreateDesktopGroupRequest {
	s.DefaultInitDesktopCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDefaultLanguage(v string) *CreateDesktopGroupRequest {
	s.DefaultLanguage = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDeleteDuration(v int64) *CreateDesktopGroupRequest {
	s.DeleteDuration = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDesktopGroupName(v string) *CreateDesktopGroupRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDesktopType(v string) *CreateDesktopGroupRequest {
	s.DesktopType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDirectoryId(v string) *CreateDesktopGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetEndUserIds(v []*string) *CreateDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *CreateDesktopGroupRequest) SetExclusiveType(v string) *CreateDesktopGroupRequest {
	s.ExclusiveType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetFileSystemId(v string) *CreateDesktopGroupRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetGroupAmount(v int32) *CreateDesktopGroupRequest {
	s.GroupAmount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetGroupVersion(v int32) *CreateDesktopGroupRequest {
	s.GroupVersion = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetHostname(v string) *CreateDesktopGroupRequest {
	s.Hostname = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetIdleDisconnectDuration(v int64) *CreateDesktopGroupRequest {
	s.IdleDisconnectDuration = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetImageId(v string) *CreateDesktopGroupRequest {
	s.ImageId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetKeepDuration(v int64) *CreateDesktopGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetLoadPolicy(v int64) *CreateDesktopGroupRequest {
	s.LoadPolicy = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMaxDesktopsCount(v int32) *CreateDesktopGroupRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMinDesktopsCount(v int32) *CreateDesktopGroupRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMultiResource(v bool) *CreateDesktopGroupRequest {
	s.MultiResource = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetOfficeSiteId(v string) *CreateDesktopGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetOwnType(v int32) *CreateDesktopGroupRequest {
	s.OwnType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPeriod(v int32) *CreateDesktopGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPeriodUnit(v string) *CreateDesktopGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPolicyGroupId(v string) *CreateDesktopGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetProfileFollowSwitch(v bool) *CreateDesktopGroupRequest {
	s.ProfileFollowSwitch = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPromotionId(v string) *CreateDesktopGroupRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetRatioThreshold(v float32) *CreateDesktopGroupRequest {
	s.RatioThreshold = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetRegionId(v string) *CreateDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetResellerOwnerUid(v int64) *CreateDesktopGroupRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetResetType(v int64) *CreateDesktopGroupRequest {
	s.ResetType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetScaleStrategyId(v string) *CreateDesktopGroupRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetSessionType(v string) *CreateDesktopGroupRequest {
	s.SessionType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetSimpleUserGroupId(v string) *CreateDesktopGroupRequest {
	s.SimpleUserGroupId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetSnapshotPolicyId(v string) *CreateDesktopGroupRequest {
	s.SnapshotPolicyId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetStopDuration(v int64) *CreateDesktopGroupRequest {
	s.StopDuration = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetSystemDiskCategory(v string) *CreateDesktopGroupRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetSystemDiskPerLevel(v string) *CreateDesktopGroupRequest {
	s.SystemDiskPerLevel = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetSystemDiskSize(v int32) *CreateDesktopGroupRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetTag(v []*CreateDesktopGroupRequestTag) *CreateDesktopGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateDesktopGroupRequest) SetTimerGroupId(v string) *CreateDesktopGroupRequest {
	s.TimerGroupId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetUserGroupName(v string) *CreateDesktopGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetUserOuPath(v string) *CreateDesktopGroupRequest {
	s.UserOuPath = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetVolumeEncryptionEnabled(v bool) *CreateDesktopGroupRequest {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetVolumeEncryptionKey(v string) *CreateDesktopGroupRequest {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetVpcId(v string) *CreateDesktopGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDesktopGroupRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDesktopGroupRequestTag struct {
	// The tag key. If you specify this parameter, the value cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDesktopGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDesktopGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDesktopGroupRequestTag) SetKey(v string) *CreateDesktopGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDesktopGroupRequestTag) SetValue(v string) *CreateDesktopGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDesktopGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
