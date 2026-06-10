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
	// Specifies whether to authorize all users in the desktop group\\"s categories.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// Alice
	AllClassifyUsers *bool `json:"AllClassifyUsers,omitempty" xml:"AllClassifyUsers,omitempty"`
	// Specifies whether to allow automatic creation of desktops in the subscription desktop group. This parameter is required and applies only when `ChargeType` is set to `PrePaid`.
	//
	// example:
	//
	// 1
	AllowAutoSetup *int32 `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	// The number of desktops to reserve in the pay-as-you-go desktop group. This parameter is required and applies only when `ChargeType` is set to `PostPaid`. Valid values:
	//
	// - 0: Does not reserve desktops.
	//
	// - N: Reserves N desktops, where N is an integer from 1 to 100.
	//
	// > If no desktops are reserved, a user must wait for a new desktop to be created and started, which can cause connection delays. We recommend reserving an appropriate number of desktops to improve connection times.
	//
	// example:
	//
	// 1
	AllowBufferCount *int32 `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	// Specifies whether to automatically pay for subscription orders.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the subscription desktop group.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The number of concurrent sessions allowed per desktop in a multi-session desktop group.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 2
	BindAmount *int64 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	// The bundle ID.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// - For `subscription` desktop groups: The number of desktops to purchase. Valid values: 0 to 200.
	//
	// - For `pay-as-you-go` desktop groups: The minimum number of desktops in the group. Valid values: 0 to `MaxDesktopsCount`. The default value is 1.
	//
	// example:
	//
	// 3
	BuyDesktopsCount *int32 `json:"BuyDesktopsCount,omitempty" xml:"BuyDesktopsCount,omitempty"`
	// The billing method of the desktops.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The type of the desktop group.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// teacher
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// A client token to ensure the idempotence of the request. You can use your client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// A description or comments for the desktop group.
	//
	// example:
	//
	// comment
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The maximum duration of a connected session. When the session duration reaches this value, the session is automatically disconnected. Unit: milliseconds. Valid values: 900000 (15 minutes) to 345600000 (4 days).
	//
	// example:
	//
	// 900000
	ConnectDuration *int64 `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	// The type of the data disk.
	//
	// example:
	//
	// cloud_auto
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The performance level (PL) of the ESSD. Default value: PL0.
	//
	// example:
	//
	// PL0
	DataDiskPerLevel *string `json:"DataDiskPerLevel,omitempty" xml:"DataDiskPerLevel,omitempty"`
	// The size of the data disk. Unit: GiB. The value must be a multiple of 20 and in the range of 0 to 16,380.
	//
	// <props="china">
	//
	// - A value of 0 indicates that no data disk is attached.
	//
	// - If the selected bundle uses an Enhanced SSD (ESSD) at PL0, the minimum data disk size is 40 GiB.
	//
	// - If the selected bundle uses an SSD, the minimum data disk size is 20 GiB.
	//
	//
	//
	// <props="intl">
	//
	// - A value of 0 indicates that no data disk is attached.
	//
	// - If the selected bundle uses an SSD, the minimum data disk size is 20 GiB.
	//
	//
	//
	// Default value: 0
	//
	// example:
	//
	// 80
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The default number of desktops to create in the desktop group. The default value is 1.
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
	DeleteDuration  *int64  `json:"DeleteDuration,omitempty" xml:"DeleteDuration,omitempty"`
	// The name of the desktop group. The name must be 1 to 30 characters long, start with a letter or a Chinese character, and must not begin with `http://` or `https://`. The name can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// SharedComputers01
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// The desktop type. You can call the [DescribeDesktopTypes](~~DescribeDesktopTypes~~) operation to query supported desktop types.
	//
	// example:
	//
	// eds.enterprise_office.16c64g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The directory ID.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// dri-uf62w3qzt4aigvlcb****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// An array of user IDs to authorize for the desktop group.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// Specifies the pool type. To create a static pool, set this parameter to `Exclusive`. This is required if `SessionType` is `MultipleSession`.
	//
	// example:
	//
	// Exclusive
	ExclusiveType *string `json:"ExclusiveType,omitempty" xml:"ExclusiveType,omitempty"`
	// The ID of the Apsara File Storage NAS file system used for user data roaming.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// kegd-nas-****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of individual desktops to create. This parameter is required only if `MultiResource` is set to `false`. Valid values: 1 to 5. Default value: 1.
	//
	// example:
	//
	// 1
	GroupAmount *int32 `json:"GroupAmount,omitempty" xml:"GroupAmount,omitempty"`
	// The version of the desktop group.
	//
	// example:
	//
	// 2
	GroupVersion *int32 `json:"GroupVersion,omitempty" xml:"GroupVersion,omitempty"`
	// The custom hostname for the desktops. This parameter is applicable only to Windows desktops in an AD office network.
	//
	// The hostname must meet the following naming conventions:
	//
	// - Must be 2 to 15 characters in length.
	//
	// - Can contain letters, digits, and hyphens (-). It cannot start or end with a hyphen, contain consecutive hyphens, or consist only of digits.
	//
	// To generate sequential hostnames when creating multiple desktops, use the format `name_prefix[begin_number,bits]name_suffix`. For example, if you set the Hostname parameter to `ecd-[1,4]-test`, the first desktop is named ecd-0001-test, the second is named ecd-0002-test, and so on.
	//
	// - `name_prefix`: The prefix of the hostname.
	//
	// - `[begin_number,bits]`: The sequential number in the hostname. `begin_number` is the starting number, which can be an integer from 0 to 999999. The default value is 0. `bits` is the number of digits, which can be an integer from 1 to 6. The default value is 6.
	//
	// - `name_suffix`: The suffix of the hostname.
	//
	// example:
	//
	// testhost
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The maximum duration that a session can be idle before it is automatically disconnected. A session is considered idle if there is no keyboard or mouse input. Unit: milliseconds. Valid values: 360000 (6 minutes) to 3600000 (60 minutes).
	//
	// Thirty seconds before disconnection, the user is prompted to save their work to prevent data loss.
	//
	// > This parameter applies only to desktops created from image version 1.0.2 or later.
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
	// The duration for which a session is kept active after a user disconnects. Unit: milliseconds. Valid values: 180000 (3 minutes) to 345600000 (4 days). A value of 0 indicates that the session is retained indefinitely.
	//
	// If a user reconnects within this period, they can resume their session. If they fail to reconnect, the session is terminated, and any unsaved data is lost.
	//
	// example:
	//
	// 180000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// The load balancing policy for the multi-session desktop group.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 0
	LoadPolicy *int64 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// The maximum number of desktops in the pay-as-you-go desktop group. Valid values: 0 to 500.
	//
	// example:
	//
	// 50
	MaxDesktopsCount *int32 `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	// The minimum number of desktops in the subscription desktop group. This parameter is required only if `ChargeType` is `PrePaid`. Valid values: 0 to `MaxDesktopsCount`. Default value: 1.
	//
	// example:
	//
	// 1
	MinDesktopsCount *int32 `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	// Specifies whether to create a desktop group.
	//
	// example:
	//
	// false
	MultiResource *bool `json:"MultiResource,omitempty" xml:"MultiResource,omitempty"`
	// The ID of the office network for the desktops.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+os-c5cy7q578s8jc****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The type of the desktop.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 0
	OwnType *int32 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The subscription duration for the desktops. This parameter is required only if `ChargeType` is set to `PrePaid`. The `PeriodUnit` parameter specifies the time unit for this duration.
	//
	// - If `PeriodUnit` is `Month`, the valid values are:
	//
	//   - 1
	//
	//   - 2
	//
	//   - 3
	//
	//   - 6
	//
	// - If `PeriodUnit` is `Year`, the valid values are:
	//
	//   - 1
	//
	//   - 2
	//
	//   - 3
	//
	//   - 4
	//
	//   - 5
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit of the subscription period.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the policy to apply to the desktops.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-9c2d6t2dwflqr****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// Specifies whether to enable user data roaming.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// false
	ProfileFollowSwitch *bool `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_*****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The session usage threshold that triggers auto scaling for multi-session desktop groups. Session usage is calculated by using the following formula:
	//
	// `Session usage = (Number of connected sessions / (Total number of desktops × Maximum number of sessions per desktop)) × 100%`
	//
	// When session usage reaches this threshold, new desktops are created. When session usage falls below this threshold, the group scales in by deleting surplus desktops.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 0.5
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// The ID of the region. To find the regions supported by Elastic Desktop Service (EDS), call the [DescribeRegions](~~DescribeRegions~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The desktop reset type.
	//
	// example:
	//
	// 0
	ResetType *int64 `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The ID of the scaling policy.
	//
	// > This parameter is not yet available.
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
	SessionType       *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	SimpleUserGroupId *string `json:"SimpleUserGroupId,omitempty" xml:"SimpleUserGroupId,omitempty"`
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-28mp6my0l6zow****
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	// The amount of time a desktop can be idle before it is automatically stopped. Connecting to a stopped desktop automatically starts it. Unit: milliseconds.
	//
	// example:
	//
	// 300000
	StopDuration *int64 `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
	// The type of the system disk.
	//
	// example:
	//
	// cloud_auto
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The performance level (PL) of the ESSD. Default value: PL0.
	//
	// example:
	//
	// PL0
	SystemDiskPerLevel *string `json:"SystemDiskPerLevel,omitempty" xml:"SystemDiskPerLevel,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// > The system disk size must be at least the size of the image.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The list of tags. You can specify up to 20 tags.
	Tag []*CreateDesktopGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the scheduled task group.
	//
	// example:
	//
	// ccg-0caoeogrk9m5****
	TimerGroupId  *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	UserOuPath    *string `json:"UserOuPath,omitempty" xml:"UserOuPath,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the key from Key Management Service (KMS) used for disk encryption. You can call the [ListKeys](https://help.aliyun.com/document_detail/28951.html) operation to obtain the key ID.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
	// The ID of the Virtual Private Cloud (VPC) that contains the office network for the desktops.
	//
	// > This parameter is not yet available.
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
	// The tag key. The key cannot be an empty string, can be up to 128 characters long, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The value can be an empty string. The value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
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
