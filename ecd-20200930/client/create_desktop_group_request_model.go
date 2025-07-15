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
	// The types of the users.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// Alice
	AllClassifyUsers *bool `json:"AllClassifyUsers,omitempty" xml:"AllClassifyUsers,omitempty"`
	// Specifies whether to enable batch-based automatic creation of subscription cloud computers for the shared group. This parameter is required if you set `ChargeType` to `PrePaid`.
	//
	// Valid values:
	//
	// 	- 0: enables batch-based automatic creation of subscription cloud computers.
	//
	// 	- 1: disables batch-based automatic creation of subscription cloud computers.
	//
	// example:
	//
	// 1
	AllowAutoSetup *int32 `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	// The maximum number of pay-as-you-go cloud computers that can be reserved in the shared group. This parameter is required if you set `ChargeType` to `PostPaid`. Valid values:
	//
	// 	- 0: does not reserve any cloud computers.
	//
	// 	- N: reserves N cloud computers (1≤ N ≤ 100).
	//
	// >  Setting this parameter to 0 means no cloud computers will be reserved in the shared group. In this case, the system must create, start, and assign cloud computers to end users upon request, which can be time-consuming. To improve user experience, we recommend that you reserve a specific number of cloud computers.
	//
	// example:
	//
	// 1
	AllowBufferCount *int32 `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	// Specifies whether to automatically complete the payment for subscription orders.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the shared subscription group.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The number of concurrent sessions of the multi-session shared group.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 1
	BindAmount *int64 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	// The ID of the cloud computer template.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// 	- For shared subscription groups, this parameter defines the initial number of cloud computers to be created. Valid values: 0 to 200.
	//
	// 	- For shared pay-as-you-go groups, this parameter defines the minimum initial number of cloud computers to be created. Valid values: 0 to `MaxDesktopsCount`. Default value: 1.
	//
	// example:
	//
	// 3
	BuyDesktopsCount *int32 `json:"BuyDesktopsCount,omitempty" xml:"BuyDesktopsCount,omitempty"`
	// The billing method of the shared group.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The type of the cloud computers in the shared group.
	//
	// >  This parameter is not publicly available.
	//
	// Valid values:
	//
	// 	- teacher: cloud computers designed for teachers.
	//
	// 	- student: cloud computers designed for students.
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
	// The remarks of the shared group.
	//
	// example:
	//
	// test
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The maximum duration for which each session remains connected. The session is automatically disconnected once the specified maximum time limit is reached. Unit: milliseconds. Valid values: 900000 to 345600000. That is, the session can be connected for 15 to 5,760 minutes (4 days).
	//
	// example:
	//
	// 300000
	ConnectDuration *int64 `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	// The category of the data disk.
	//
	// Valid values:
	//
	// 	- cloud_auto: the standard SSD.
	//
	// 	- cloud_essd: the ESSD.
	//
	// example:
	//
	// cloud_auto
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The PL of the data disk of the ESSD category. Default value: PL0.
	//
	// Valid values:
	//
	// 	- PL1
	//
	// 	- PL0
	//
	// example:
	//
	// PL0
	DataDiskPerLevel *string `json:"DataDiskPerLevel,omitempty" xml:"DataDiskPerLevel,omitempty"`
	// The size of the data disk. Unit: GB. Valid values: 0 to 16380. The value must be an integral multiple of 20.
	//
	// 	- A value of 0 means no data disk is attached.
	//
	// 	- If the selected plan includes a standard SSD, the data disk size must be at least 20 GB.
	//
	// Default value: 0.
	//
	// example:
	//
	// 80
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The default number of cloud computers that you want to create at the same time in the shared group. Default value: 1.
	//
	// example:
	//
	// 1
	DefaultInitDesktopCount *int32 `json:"DefaultInitDesktopCount,omitempty" xml:"DefaultInitDesktopCount,omitempty"`
	// The language of the OS.
	//
	// Valid values:
	//
	// 	- en-US: English.
	//
	// 	- zh-HK: Traditional Chinese.
	//
	// 	- zh-CN: Simplified Chinese
	//
	// 	- ja-JP: Japanese.
	//
	// example:
	//
	// zh-CN
	DefaultLanguage *string `json:"DefaultLanguage,omitempty" xml:"DefaultLanguage,omitempty"`
	DeleteDuration  *int64  `json:"DeleteDuration,omitempty" xml:"DeleteDuration,omitempty"`
	// The name of the shared group. The name can be up to 30 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desktopGroupName1
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// The specifications of the cloud computer. You can call the [DescribeDesktopTypes](~~DescribeDesktopTypes~~) operation to query all the supported specifications.
	//
	// example:
	//
	// eds.enterprise_office.16c64g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The ID of the directory.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// hide
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The IDs of the end users.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// Specifies whether the shared group is exclusive. You must set this parameter to `Exclusive` when `SessionType` is set to `MultipleSession`.
	//
	// example:
	//
	// Exclusive
	ExclusiveType *string `json:"ExclusiveType,omitempty" xml:"ExclusiveType,omitempty"`
	// The ID of the File Storage NAS (NAS) file system for the user data roaming feature.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 04f314****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of shared groups for the single-cloud computer type. You must specify this parameter if you set `MultiResource` to `false`. Valid values: 1 to 5. Default value: 1.
	//
	// example:
	//
	// 1
	GroupAmount *int32 `json:"GroupAmount,omitempty" xml:"GroupAmount,omitempty"`
	// The version of the shared group.
	//
	// example:
	//
	// 2
	GroupVersion *int32 `json:"GroupVersion,omitempty" xml:"GroupVersion,omitempty"`
	// The hostname series of the cloud computer. This parameter is supported exclusively when the office network operates on Active Directory (AD) and the cloud computer runs on a Windows operating system.
	//
	// Naming conventions:
	//
	// 	- A hostname must be 2 to 15 characters in length
	//
	// 	- and can contain only letters, digits, and hyphens (-). It cannot start or end with a hyphen (-), contain consecutive hyphens (-), or contain only digits.
	//
	// If you want to create multiple cloud computers, specify their hostnames in the `name_prefix[begin_number,bits]name_suffix` format. For example, if you set Hostname to ecd-[1,4]-test, the hostnames of the cloud computers will be assigned sequentially as ecd-0001-test, ecd-0002-test, and so on.
	//
	// 	- `name_prefix`: the prefix of the hostname.
	//
	// 	- `[begin_number,bits]`: the sequential number in the hostname. The `begin_number` value is the starting number. Valid values of begin_number: 0 to 999999. Default value: 0. The `bits` value is the number of digits. Valid values: 1 to 6. Default value: 6.
	//
	// 	- `name_suffix`: the suffix of the hostname.
	//
	// example:
	//
	// testhost
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The duration after which a session is terminated if no keyboard or mouse activity is detected. When an end user connects to a cloud computer, a session is initiated. If no input from the keyboard or mouse is detected within this specified timeframe, the session is automatically closed. Unit: milliseconds. Valid values: 360000 to 3600000 (6 minutes to 60 minutes)
	//
	// The system prompts end users to save their data 30 seconds before a session is disconnected. To avoid data loss, end users must save their session data upon receiving the prompt.
	//
	// >  This parameter is suitable only for cloud computers whose image version is v1.0.2 or later.
	//
	// example:
	//
	// 300000
	IdleDisconnectDuration *int64 `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The duration for which each session remains active after disconnection. Valid values: 180000 (3 minutes) to 345600000 (4 days). Unit: milliseconds. If you set this parameter to 0, the session is permanently retained after disconnection.
	//
	// When a session is disconnected, take note of the following items: 1. If the end user does not resume the session within the specified duration, the session will close, and all unsaved data will be cleared. 2. If the end user resumes the session within the specified duration, the session data will remain accessible for continued use.
	//
	// example:
	//
	// 6000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// The load balancing policy of the multi-session shared group.
	//
	// >  This parameter is not publicly available.
	//
	// Valid values:
	//
	// 	- 0: depth-first
	//
	// 	- 1: breadth first
	//
	// example:
	//
	// 0
	LoadPolicy *int64 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// The maximum number of pay-as-you-go cloud computers that can be automatically provisioned at the same time in the shared group. Valid values: 0 to 500.
	//
	// example:
	//
	// 50
	MaxDesktopsCount *int32 `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	// The minimum number of subscription cloud computers that can be automatically provisioned at the same time in the shared group. This parameter is required if you set `ChargeType` to `PrePaid`. Default value: 1. Valid values: 0 to `MaxDesktopsCount`.
	//
	// example:
	//
	// 1
	MinDesktopsCount *int32 `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	// Specifies whether the shared group is a multi-cloud computer type.
	//
	// Valid values:
	//
	// 	- true: a multi-cloud computer type.
	//
	// 	- false: a single-cloud computer type.
	//
	// example:
	//
	// false
	MultiResource *bool `json:"MultiResource,omitempty" xml:"MultiResource,omitempty"`
	// The ID of the office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+os-c5cy7q578s8jc****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The session type of the shared group.
	//
	// >  This parameter is not publicly available.
	//
	// Valid values:
	//
	// 	- 0: single-session.
	//
	// 	- 1: multi-session.
	//
	// example:
	//
	// 0
	OwnType *int32 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The subscription duration of the shared group. This parameter is required if you set `ChargeType` to `PrePaid`. You must specify the subscription duration unit by using `PeriodUnit`.
	//
	// 	- If you set `PeriodUnit` to `Month`, valid values of this parameter:
	//
	//     	- 1
	//
	//     	- 2
	//
	//     	- 3
	//
	//     	- 6
	//
	// 	- If you set `PeriodUnit` to `Year`, valid values of this parameter:
	//
	//     	- 1
	//
	//     	- 2
	//
	//     	- 3
	//
	//     	- 4
	//
	//     	- 5
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-9c2d6t2dwflqr****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// Specifies whether to enable user data roaming.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// false
	ProfileFollowSwitch *bool `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
	// The ID of the coupon.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_*****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The threshold for the ratio of connected sessions. This parameter defines the condition that activates automatic scaling of cloud computers in a multi-session shared group. The ratio of connected sessions is calculated by using the following formula:
	//
	// `Ratio of connected sessions = Number of connected sessions/(Total number of cloud computers × Maximum number of sessions allowed for each cloud computer) × 100%`.
	//
	// If the connected session ratio exceeds the specified threshold, new cloud computers are provisioned. If the ratio falls below the threshold, idle cloud computers are deleted.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 0.5
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The reset option of the shared group.
	//
	// Valid values:
	//
	// 	- 0: Reset is not required.
	//
	// 	- 1: Only the system disk is reset.
	//
	// 	- 2: Only the data disk is reset.
	//
	// 	- 3: Both the system disk and the data disk are reset.
	//
	// example:
	//
	// 0
	ResetType *int64 `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The ID of the scaling policy.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// hide
	ScaleStrategyId *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	// The type of the session.
	//
	// Valid values:
	//
	// 	- SingleSession
	//
	// 	- MultipleSession
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
	// The maximum period of inactivity allowed before a cloud computer is automatically stopped. If the idle duration reaches the specified limit, the system stops the cloud computer. When an end user reconnects to the stopped cloud computer, it automatically restarts. Unit: milliseconds.
	//
	// example:
	//
	// 300000
	StopDuration *int64 `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
	// The category of the system disk.
	//
	// Valid values:
	//
	// 	- cloud_auto: the standard SSD.
	//
	// 	- cloud_essd: the Enterprise SSD (ESSD).
	//
	// example:
	//
	// cloud_auto
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The performance level (PL) of the system disk of the ESSD category. Default value: PL0.
	//
	// Valid values:
	//
	// 	- PL1
	//
	// 	- PL0
	//
	// example:
	//
	// PL0
	SystemDiskPerLevel *string `json:"SystemDiskPerLevel,omitempty" xml:"SystemDiskPerLevel,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// >  The system disk must be at least as large as the image.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*CreateDesktopGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the timer group.
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
	// The ID of the Key Management Service (KMS) key that you want to use when disk encryption is enabled. You can call the [ListKeys](https://help.aliyun.com/document_detail/28951.html) operation to obtain a list of KMS keys.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// hide
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
	return dara.Validate(s)
}

type CreateDesktopGroupRequestTag struct {
	// The tag key. You cannot specify an empty string as a tag key. A tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify an empty string as a tag key. A tag value can be up to 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
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
