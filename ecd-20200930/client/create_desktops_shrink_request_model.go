// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateDesktopsShrinkRequest
	GetAmount() *int32
	SetAppRuleId(v string) *CreateDesktopsShrinkRequest
	GetAppRuleId() *string
	SetAutoPay(v bool) *CreateDesktopsShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateDesktopsShrinkRequest
	GetAutoRenew() *bool
	SetBundleId(v string) *CreateDesktopsShrinkRequest
	GetBundleId() *string
	SetBundleModels(v []*CreateDesktopsShrinkRequestBundleModels) *CreateDesktopsShrinkRequest
	GetBundleModels() []*CreateDesktopsShrinkRequestBundleModels
	SetChannelCookie(v string) *CreateDesktopsShrinkRequest
	GetChannelCookie() *string
	SetChargeType(v string) *CreateDesktopsShrinkRequest
	GetChargeType() *string
	SetDesktopAttachmentShrink(v string) *CreateDesktopsShrinkRequest
	GetDesktopAttachmentShrink() *string
	SetDesktopMemberIp(v string) *CreateDesktopsShrinkRequest
	GetDesktopMemberIp() *string
	SetDesktopName(v string) *CreateDesktopsShrinkRequest
	GetDesktopName() *string
	SetDesktopNameSuffix(v bool) *CreateDesktopsShrinkRequest
	GetDesktopNameSuffix() *bool
	SetDesktopTimers(v []*CreateDesktopsShrinkRequestDesktopTimers) *CreateDesktopsShrinkRequest
	GetDesktopTimers() []*CreateDesktopsShrinkRequestDesktopTimers
	SetDirectoryId(v string) *CreateDesktopsShrinkRequest
	GetDirectoryId() *string
	SetEndUserId(v []*string) *CreateDesktopsShrinkRequest
	GetEndUserId() []*string
	SetExtendInfo(v string) *CreateDesktopsShrinkRequest
	GetExtendInfo() *string
	SetGroupId(v string) *CreateDesktopsShrinkRequest
	GetGroupId() *string
	SetHostname(v string) *CreateDesktopsShrinkRequest
	GetHostname() *string
	SetMonthDesktopSetting(v *CreateDesktopsShrinkRequestMonthDesktopSetting) *CreateDesktopsShrinkRequest
	GetMonthDesktopSetting() *CreateDesktopsShrinkRequestMonthDesktopSetting
	SetOfficeSiteId(v string) *CreateDesktopsShrinkRequest
	GetOfficeSiteId() *string
	SetPeriod(v int32) *CreateDesktopsShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateDesktopsShrinkRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateDesktopsShrinkRequest
	GetPolicyGroupId() *string
	SetPromotionId(v string) *CreateDesktopsShrinkRequest
	GetPromotionId() *string
	SetPurchaseOptionsShrink(v string) *CreateDesktopsShrinkRequest
	GetPurchaseOptionsShrink() *string
	SetQosRuleId(v string) *CreateDesktopsShrinkRequest
	GetQosRuleId() *string
	SetRegionId(v string) *CreateDesktopsShrinkRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *CreateDesktopsShrinkRequest
	GetResellerOwnerUid() *int64
	SetResourceGroupId(v string) *CreateDesktopsShrinkRequest
	GetResourceGroupId() *string
	SetSavingPlanId(v string) *CreateDesktopsShrinkRequest
	GetSavingPlanId() *string
	SetSnapshotPolicyId(v string) *CreateDesktopsShrinkRequest
	GetSnapshotPolicyId() *string
	SetSubnetId(v string) *CreateDesktopsShrinkRequest
	GetSubnetId() *string
	SetTag(v []*CreateDesktopsShrinkRequestTag) *CreateDesktopsShrinkRequest
	GetTag() []*CreateDesktopsShrinkRequestTag
	SetTimerGroupId(v string) *CreateDesktopsShrinkRequest
	GetTimerGroupId() *string
	SetUserAssignMode(v string) *CreateDesktopsShrinkRequest
	GetUserAssignMode() *string
	SetUserCommands(v []*CreateDesktopsShrinkRequestUserCommands) *CreateDesktopsShrinkRequest
	GetUserCommands() []*CreateDesktopsShrinkRequestUserCommands
	SetUserName(v string) *CreateDesktopsShrinkRequest
	GetUserName() *string
	SetVolumeEncryptionEnabled(v bool) *CreateDesktopsShrinkRequest
	GetVolumeEncryptionEnabled() *bool
	SetVolumeEncryptionKey(v string) *CreateDesktopsShrinkRequest
	GetVolumeEncryptionKey() *string
	SetVpcId(v string) *CreateDesktopsShrinkRequest
	GetVpcId() *string
}

type CreateDesktopsShrinkRequest struct {
	// The number of cloud desktops to create. Valid values: 1 to 300. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of the application control policy.
	//
	// example:
	//
	// bwr-245d4e0e6b7d42f5afa97eb3fbc7e488
	AppRuleId *string `json:"AppRuleId,omitempty" xml:"AppRuleId,omitempty"`
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the cloud desktops. This parameter is valid only when `ChargeType` is set to `PrePaid`.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The bundle ID. If you do not specify this parameter, you must use the `DesktopAttachment` parameter to configure the cloud desktop.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// An array of bundle objects. Use this parameter to create cloud desktops from one or more bundles in a single call.
	//
	// if can be null:
	// true
	BundleModels []*CreateDesktopsShrinkRequestBundleModels `json:"BundleModels,omitempty" xml:"BundleModels,omitempty" type:"Repeated"`
	// > This parameter is for internal use only.
	//
	// example:
	//
	// PBKB1QbqEl2tslEuU6gRrLxvCFBU2M%2FVD0Eru6Oo%2FI9LTU3XQhvq3PGMWarE%2BPJdkNvCqT3blqlRSthNy4A%2BJQ%3D%3D
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The billing method of the cloud desktops.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The parameters for creating a cloud desktop without a bundle. This parameter is used only if `BundleId` is not specified.
	DesktopAttachmentShrink *string `json:"DesktopAttachment,omitempty" xml:"DesktopAttachment,omitempty"`
	// The private IP address of the cloud desktop.
	//
	// example:
	//
	// 10.0.0.1
	DesktopMemberIp *string `json:"DesktopMemberIp,omitempty" xml:"DesktopMemberIp,omitempty"`
	// The name of the cloud desktop. The name must meet the following requirements:
	//
	// - The name must be 1 to 64 characters in length.
	//
	// - The name must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// - The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// DemoComputer01
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// Specifies whether to automatically append a suffix to the value of `DesktopName` when you create multiple cloud desktops.
	//
	// example:
	//
	// false
	DesktopNameSuffix *bool `json:"DesktopNameSuffix,omitempty" xml:"DesktopNameSuffix,omitempty"`
	// The details of the scheduled tasks for the cloud desktops. This parameter is being deprecated. We recommend that you use `TimerGroupId` instead.
	//
	// if can be null:
	// true
	DesktopTimers []*CreateDesktopsShrinkRequestDesktopTimers `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	// > This parameter is for internal use only.
	//
	// example:
	//
	// cn-hangzhou+dir-300943****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The IDs of the end users to assign to the cloud desktops. You can specify up to 100 IDs.
	//
	// example:
	//
	// 123456789
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The extended information in a JSON string. This parameter is for internal use only.
	//
	// example:
	//
	// {}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The ID of the desktop pool.
	//
	// example:
	//
	// dg-boyczi8enfyc5****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The custom hostname of the cloud desktop. This parameter is supported only for Windows cloud desktops that are in an AD office network.
	//
	// The hostname must meet the following requirements:
	//
	// - The hostname must be 2 to 15 characters in length.
	//
	// - The hostname can contain letters, digits, and hyphens (-). The hostname cannot start or end with a hyphen, contain consecutive hyphens, or consist of only digits.
	//
	// If you create multiple cloud desktops, you can use the `name_prefix[begin_number,bits]name_suffix` format to specify sequential hostnames for the cloud desktops. For example, if you set the `Hostname` parameter to `ecd-[1,4]-test`, the hostname of the first cloud desktop is `ecd-0001-test`, the second is `ecd-0002-test`, and so on.
	//
	// - `name_prefix`: the prefix of the hostname.
	//
	// - `[begin_number,bits]`: The sequential part of the hostname.
	//
	// - `name_suffix`: the suffix of the hostname.
	//
	// example:
	//
	// testhost
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The parameters for purchasing a monthly usage package.
	MonthDesktopSetting *CreateDesktopsShrinkRequestMonthDesktopSetting `json:"MonthDesktopSetting,omitempty" xml:"MonthDesktopSetting,omitempty" type:"Struct"`
	// The ID of the office network.
	//
	// example:
	//
	// cn-hangzhou+dir-387822****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The subscription duration. The unit is specified by the `PeriodUnit` parameter. This parameter is required only when `ChargeType` is set to `PrePaid`.
	//
	// - If `PeriodUnit` is set to `Month`, valid values are:
	//
	//   - 1
	//
	//   - 2
	//
	//   - 3
	//
	//   - 6
	//
	// - If `PeriodUnit` is set to `Year`, valid values are:
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
	// The unit of the subscription duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// system-all-enabled-policy
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 23141
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The additional parameters for a specific purchase type.
	PurchaseOptionsShrink *string `json:"PurchaseOptions,omitempty" xml:"PurchaseOptions,omitempty"`
	// The ID of the bandwidth QoS policy.
	//
	// example:
	//
	// qos-52fqmg6kvyro7zu4l
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions that support Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UID of the resource owner in reseller mode. This parameter is required only in reseller mode.
	//
	// example:
	//
	// 1828644634819902
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-3mtuc28rx95lx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// > This parameter is for internal use only.
	//
	// example:
	//
	// spn-26c1b7bcrjcI****
	SavingPlanId *string `json:"SavingPlanId,omitempty" xml:"SavingPlanId,omitempty"`
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-28mp6my0l6zow****
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1m*****
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The tags to add to the cloud desktops. A resource can have up to 20 tags.
	Tag []*CreateDesktopsShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the scheduled task group.
	//
	// example:
	//
	// ccg-0caoeogrk9m5****
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	// The user assignment mode for the cloud desktops.
	//
	// > If you do not specify the `EndUserId` parameter, the created cloud desktops are unassigned.
	//
	// example:
	//
	// ALL
	UserAssignMode *string `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
	// The custom scripts to run on the cloud desktops after they start.
	UserCommands []*CreateDesktopsShrinkRequestUserCommands `json:"UserCommands,omitempty" xml:"UserCommands,omitempty" type:"Repeated"`
	// > This parameter is for internal use only.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the KMS key to use for disk encryption. You can call the [ListKeys](https://help.aliyun.com/document_detail/28951.html) operation to obtain a list of key IDs.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
	// > This parameter is for internal use only.
	//
	// example:
	//
	// vpc-uf6w8u60n8xbkg5el****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateDesktopsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopsShrinkRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateDesktopsShrinkRequest) GetAppRuleId() *string {
	return s.AppRuleId
}

func (s *CreateDesktopsShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateDesktopsShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDesktopsShrinkRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateDesktopsShrinkRequest) GetBundleModels() []*CreateDesktopsShrinkRequestBundleModels {
	return s.BundleModels
}

func (s *CreateDesktopsShrinkRequest) GetChannelCookie() *string {
	return s.ChannelCookie
}

func (s *CreateDesktopsShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateDesktopsShrinkRequest) GetDesktopAttachmentShrink() *string {
	return s.DesktopAttachmentShrink
}

func (s *CreateDesktopsShrinkRequest) GetDesktopMemberIp() *string {
	return s.DesktopMemberIp
}

func (s *CreateDesktopsShrinkRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *CreateDesktopsShrinkRequest) GetDesktopNameSuffix() *bool {
	return s.DesktopNameSuffix
}

func (s *CreateDesktopsShrinkRequest) GetDesktopTimers() []*CreateDesktopsShrinkRequestDesktopTimers {
	return s.DesktopTimers
}

func (s *CreateDesktopsShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateDesktopsShrinkRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *CreateDesktopsShrinkRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *CreateDesktopsShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateDesktopsShrinkRequest) GetHostname() *string {
	return s.Hostname
}

func (s *CreateDesktopsShrinkRequest) GetMonthDesktopSetting() *CreateDesktopsShrinkRequestMonthDesktopSetting {
	return s.MonthDesktopSetting
}

func (s *CreateDesktopsShrinkRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateDesktopsShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateDesktopsShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateDesktopsShrinkRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateDesktopsShrinkRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateDesktopsShrinkRequest) GetPurchaseOptionsShrink() *string {
	return s.PurchaseOptionsShrink
}

func (s *CreateDesktopsShrinkRequest) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *CreateDesktopsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDesktopsShrinkRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *CreateDesktopsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDesktopsShrinkRequest) GetSavingPlanId() *string {
	return s.SavingPlanId
}

func (s *CreateDesktopsShrinkRequest) GetSnapshotPolicyId() *string {
	return s.SnapshotPolicyId
}

func (s *CreateDesktopsShrinkRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *CreateDesktopsShrinkRequest) GetTag() []*CreateDesktopsShrinkRequestTag {
	return s.Tag
}

func (s *CreateDesktopsShrinkRequest) GetTimerGroupId() *string {
	return s.TimerGroupId
}

func (s *CreateDesktopsShrinkRequest) GetUserAssignMode() *string {
	return s.UserAssignMode
}

func (s *CreateDesktopsShrinkRequest) GetUserCommands() []*CreateDesktopsShrinkRequestUserCommands {
	return s.UserCommands
}

func (s *CreateDesktopsShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateDesktopsShrinkRequest) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *CreateDesktopsShrinkRequest) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *CreateDesktopsShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDesktopsShrinkRequest) SetAmount(v int32) *CreateDesktopsShrinkRequest {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetAppRuleId(v string) *CreateDesktopsShrinkRequest {
	s.AppRuleId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetAutoPay(v bool) *CreateDesktopsShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetAutoRenew(v bool) *CreateDesktopsShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetBundleId(v string) *CreateDesktopsShrinkRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetBundleModels(v []*CreateDesktopsShrinkRequestBundleModels) *CreateDesktopsShrinkRequest {
	s.BundleModels = v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetChannelCookie(v string) *CreateDesktopsShrinkRequest {
	s.ChannelCookie = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetChargeType(v string) *CreateDesktopsShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetDesktopAttachmentShrink(v string) *CreateDesktopsShrinkRequest {
	s.DesktopAttachmentShrink = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetDesktopMemberIp(v string) *CreateDesktopsShrinkRequest {
	s.DesktopMemberIp = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetDesktopName(v string) *CreateDesktopsShrinkRequest {
	s.DesktopName = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetDesktopNameSuffix(v bool) *CreateDesktopsShrinkRequest {
	s.DesktopNameSuffix = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetDesktopTimers(v []*CreateDesktopsShrinkRequestDesktopTimers) *CreateDesktopsShrinkRequest {
	s.DesktopTimers = v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetDirectoryId(v string) *CreateDesktopsShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetEndUserId(v []*string) *CreateDesktopsShrinkRequest {
	s.EndUserId = v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetExtendInfo(v string) *CreateDesktopsShrinkRequest {
	s.ExtendInfo = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetGroupId(v string) *CreateDesktopsShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetHostname(v string) *CreateDesktopsShrinkRequest {
	s.Hostname = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetMonthDesktopSetting(v *CreateDesktopsShrinkRequestMonthDesktopSetting) *CreateDesktopsShrinkRequest {
	s.MonthDesktopSetting = v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetOfficeSiteId(v string) *CreateDesktopsShrinkRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetPeriod(v int32) *CreateDesktopsShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetPeriodUnit(v string) *CreateDesktopsShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetPolicyGroupId(v string) *CreateDesktopsShrinkRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetPromotionId(v string) *CreateDesktopsShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetPurchaseOptionsShrink(v string) *CreateDesktopsShrinkRequest {
	s.PurchaseOptionsShrink = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetQosRuleId(v string) *CreateDesktopsShrinkRequest {
	s.QosRuleId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetRegionId(v string) *CreateDesktopsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetResellerOwnerUid(v int64) *CreateDesktopsShrinkRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetResourceGroupId(v string) *CreateDesktopsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetSavingPlanId(v string) *CreateDesktopsShrinkRequest {
	s.SavingPlanId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetSnapshotPolicyId(v string) *CreateDesktopsShrinkRequest {
	s.SnapshotPolicyId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetSubnetId(v string) *CreateDesktopsShrinkRequest {
	s.SubnetId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetTag(v []*CreateDesktopsShrinkRequestTag) *CreateDesktopsShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetTimerGroupId(v string) *CreateDesktopsShrinkRequest {
	s.TimerGroupId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetUserAssignMode(v string) *CreateDesktopsShrinkRequest {
	s.UserAssignMode = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetUserCommands(v []*CreateDesktopsShrinkRequestUserCommands) *CreateDesktopsShrinkRequest {
	s.UserCommands = v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetUserName(v string) *CreateDesktopsShrinkRequest {
	s.UserName = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetVolumeEncryptionEnabled(v bool) *CreateDesktopsShrinkRequest {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetVolumeEncryptionKey(v string) *CreateDesktopsShrinkRequest {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) SetVpcId(v string) *CreateDesktopsShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDesktopsShrinkRequest) Validate() error {
	if s.BundleModels != nil {
		for _, item := range s.BundleModels {
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
	if s.MonthDesktopSetting != nil {
		if err := s.MonthDesktopSetting.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserCommands != nil {
		for _, item := range s.UserCommands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDesktopsShrinkRequestBundleModels struct {
	// The number of cloud desktops to create. Valid values: 1 to 300. Default value: 0.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The bundle ID.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The name of the cloud desktop. The name must meet the following requirements:
	//
	// - The name must be 1 to 64 characters in length.
	//
	// - The name must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// - The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// DemoComputer02
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The list of end user IDs to whom the cloud desktops are assigned.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The custom hostname of the cloud desktop. This parameter is supported only for Windows cloud desktops that are in an AD office network.
	//
	// The hostname must meet the following requirements:
	//
	// - The hostname must be 2 to 15 characters in length.
	//
	// - The hostname can contain letters, digits, and hyphens (-). The hostname cannot start or end with a hyphen, contain consecutive hyphens, or consist of only digits.
	//
	// If you create multiple cloud desktops, you can use the `name_prefix[begin_number,bits]name_suffix` format to specify sequential hostnames for the cloud desktops. For example, if you set the `Hostname` parameter to `ecd-[1,4]-test`, the hostname of the first cloud desktop is `ecd-0001-test`, the second is `ecd-0002-test`, and so on.
	//
	// - `name_prefix`: the prefix of the hostname.
	//
	// - `[begin_number,bits]`: The sequential part of the hostname.
	//
	// - `name_suffix`: the suffix of the hostname.
	//
	// example:
	//
	// testhost
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the KMS key to use for disk encryption. You can call the [ListKeys](https://help.aliyun.com/document_detail/28951.html) operation to obtain a list of key IDs.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s CreateDesktopsShrinkRequestBundleModels) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsShrinkRequestBundleModels) GoString() string {
	return s.String()
}

func (s *CreateDesktopsShrinkRequestBundleModels) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateDesktopsShrinkRequestBundleModels) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateDesktopsShrinkRequestBundleModels) GetDesktopName() *string {
	return s.DesktopName
}

func (s *CreateDesktopsShrinkRequestBundleModels) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *CreateDesktopsShrinkRequestBundleModels) GetHostname() *string {
	return s.Hostname
}

func (s *CreateDesktopsShrinkRequestBundleModels) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *CreateDesktopsShrinkRequestBundleModels) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *CreateDesktopsShrinkRequestBundleModels) SetAmount(v int32) *CreateDesktopsShrinkRequestBundleModels {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsShrinkRequestBundleModels) SetBundleId(v string) *CreateDesktopsShrinkRequestBundleModels {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsShrinkRequestBundleModels) SetDesktopName(v string) *CreateDesktopsShrinkRequestBundleModels {
	s.DesktopName = &v
	return s
}

func (s *CreateDesktopsShrinkRequestBundleModels) SetEndUserIds(v []*string) *CreateDesktopsShrinkRequestBundleModels {
	s.EndUserIds = v
	return s
}

func (s *CreateDesktopsShrinkRequestBundleModels) SetHostname(v string) *CreateDesktopsShrinkRequestBundleModels {
	s.Hostname = &v
	return s
}

func (s *CreateDesktopsShrinkRequestBundleModels) SetVolumeEncryptionEnabled(v bool) *CreateDesktopsShrinkRequestBundleModels {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *CreateDesktopsShrinkRequestBundleModels) SetVolumeEncryptionKey(v string) *CreateDesktopsShrinkRequestBundleModels {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *CreateDesktopsShrinkRequestBundleModels) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsShrinkRequestDesktopTimers struct {
	// Specifies whether to allow end users to configure the scheduled task.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The cron expression for the scheduled task.
	//
	// 	Notice:
	//
	// The cron expression is evaluated in UTC. For example, to specify 00:00 (UTC+8) every day, set the value to `0 0 16 ? 	- 1,2,3,4,5,6,7`.
	//
	// example:
	//
	// 0 40 7 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to forcibly execute the scheduled task.
	//
	// example:
	//
	// true
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The interval in minutes.
	//
	// example:
	//
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The operation to perform for the scheduled task. This parameter is valid only for scheduled tasks that are triggered by user disconnection.
	//
	// example:
	//
	// Shutdown
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The disk reset type.
	//
	// example:
	//
	// RESET_TYPE_SYSTEM
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The type of the scheduled task.
	//
	// example:
	//
	// NoOperationReboot
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s CreateDesktopsShrinkRequestDesktopTimers) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsShrinkRequestDesktopTimers) GoString() string {
	return s.String()
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) GetAllowClientSetting() *bool {
	return s.AllowClientSetting
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) GetResetType() *string {
	return s.ResetType
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) SetAllowClientSetting(v bool) *CreateDesktopsShrinkRequestDesktopTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) SetCronExpression(v string) *CreateDesktopsShrinkRequestDesktopTimers {
	s.CronExpression = &v
	return s
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) SetEnforce(v bool) *CreateDesktopsShrinkRequestDesktopTimers {
	s.Enforce = &v
	return s
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) SetInterval(v int32) *CreateDesktopsShrinkRequestDesktopTimers {
	s.Interval = &v
	return s
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) SetOperationType(v string) *CreateDesktopsShrinkRequestDesktopTimers {
	s.OperationType = &v
	return s
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) SetResetType(v string) *CreateDesktopsShrinkRequestDesktopTimers {
	s.ResetType = &v
	return s
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) SetTimerType(v string) *CreateDesktopsShrinkRequestDesktopTimers {
	s.TimerType = &v
	return s
}

func (s *CreateDesktopsShrinkRequestDesktopTimers) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsShrinkRequestMonthDesktopSetting struct {
	// > This parameter is for internal use only.
	//
	// example:
	//
	// null
	BuyerId *int64 `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	// > This parameter is for internal use only.
	//
	// example:
	//
	// null
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The duration in hours for the monthly usage package. Valid values: 120, 250, and 360.
	//
	// example:
	//
	// null
	UseDuration *int32 `json:"UseDuration,omitempty" xml:"UseDuration,omitempty"`
}

func (s CreateDesktopsShrinkRequestMonthDesktopSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsShrinkRequestMonthDesktopSetting) GoString() string {
	return s.String()
}

func (s *CreateDesktopsShrinkRequestMonthDesktopSetting) GetBuyerId() *int64 {
	return s.BuyerId
}

func (s *CreateDesktopsShrinkRequestMonthDesktopSetting) GetDesktopId() *string {
	return s.DesktopId
}

func (s *CreateDesktopsShrinkRequestMonthDesktopSetting) GetUseDuration() *int32 {
	return s.UseDuration
}

func (s *CreateDesktopsShrinkRequestMonthDesktopSetting) SetBuyerId(v int64) *CreateDesktopsShrinkRequestMonthDesktopSetting {
	s.BuyerId = &v
	return s
}

func (s *CreateDesktopsShrinkRequestMonthDesktopSetting) SetDesktopId(v string) *CreateDesktopsShrinkRequestMonthDesktopSetting {
	s.DesktopId = &v
	return s
}

func (s *CreateDesktopsShrinkRequestMonthDesktopSetting) SetUseDuration(v int32) *CreateDesktopsShrinkRequestMonthDesktopSetting {
	s.UseDuration = &v
	return s
}

func (s *CreateDesktopsShrinkRequestMonthDesktopSetting) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsShrinkRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDesktopsShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDesktopsShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDesktopsShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDesktopsShrinkRequestTag) SetKey(v string) *CreateDesktopsShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDesktopsShrinkRequestTag) SetValue(v string) *CreateDesktopsShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDesktopsShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsShrinkRequestUserCommands struct {
	// The content of the script.
	//
	// example:
	//
	// bmV3LWl0ZW0gZDpcdGVzdF91c2VyX2NvbW1hbmRzLnR4dCAtdHlwZSBm****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The encoding format of the script content specified in the `Content` parameter.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The script type.
	//
	// example:
	//
	// RunPowerShellScript
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s CreateDesktopsShrinkRequestUserCommands) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsShrinkRequestUserCommands) GoString() string {
	return s.String()
}

func (s *CreateDesktopsShrinkRequestUserCommands) GetContent() *string {
	return s.Content
}

func (s *CreateDesktopsShrinkRequestUserCommands) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *CreateDesktopsShrinkRequestUserCommands) GetContentType() *string {
	return s.ContentType
}

func (s *CreateDesktopsShrinkRequestUserCommands) SetContent(v string) *CreateDesktopsShrinkRequestUserCommands {
	s.Content = &v
	return s
}

func (s *CreateDesktopsShrinkRequestUserCommands) SetContentEncoding(v string) *CreateDesktopsShrinkRequestUserCommands {
	s.ContentEncoding = &v
	return s
}

func (s *CreateDesktopsShrinkRequestUserCommands) SetContentType(v string) *CreateDesktopsShrinkRequestUserCommands {
	s.ContentType = &v
	return s
}

func (s *CreateDesktopsShrinkRequestUserCommands) Validate() error {
	return dara.Validate(s)
}
