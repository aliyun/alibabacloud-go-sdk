// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateDesktopsRequest
	GetAmount() *int32
	SetAppRuleId(v string) *CreateDesktopsRequest
	GetAppRuleId() *string
	SetAutoPay(v bool) *CreateDesktopsRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateDesktopsRequest
	GetAutoRenew() *bool
	SetBundleId(v string) *CreateDesktopsRequest
	GetBundleId() *string
	SetBundleModels(v []*CreateDesktopsRequestBundleModels) *CreateDesktopsRequest
	GetBundleModels() []*CreateDesktopsRequestBundleModels
	SetChannelCookie(v string) *CreateDesktopsRequest
	GetChannelCookie() *string
	SetChargeType(v string) *CreateDesktopsRequest
	GetChargeType() *string
	SetDesktopAttachment(v *CreateDesktopsRequestDesktopAttachment) *CreateDesktopsRequest
	GetDesktopAttachment() *CreateDesktopsRequestDesktopAttachment
	SetDesktopMemberIp(v string) *CreateDesktopsRequest
	GetDesktopMemberIp() *string
	SetDesktopName(v string) *CreateDesktopsRequest
	GetDesktopName() *string
	SetDesktopNameSuffix(v bool) *CreateDesktopsRequest
	GetDesktopNameSuffix() *bool
	SetDesktopTimers(v []*CreateDesktopsRequestDesktopTimers) *CreateDesktopsRequest
	GetDesktopTimers() []*CreateDesktopsRequestDesktopTimers
	SetDirectoryId(v string) *CreateDesktopsRequest
	GetDirectoryId() *string
	SetEndUserId(v []*string) *CreateDesktopsRequest
	GetEndUserId() []*string
	SetExtendInfo(v string) *CreateDesktopsRequest
	GetExtendInfo() *string
	SetGroupId(v string) *CreateDesktopsRequest
	GetGroupId() *string
	SetHostname(v string) *CreateDesktopsRequest
	GetHostname() *string
	SetMonthDesktopSetting(v *CreateDesktopsRequestMonthDesktopSetting) *CreateDesktopsRequest
	GetMonthDesktopSetting() *CreateDesktopsRequestMonthDesktopSetting
	SetOfficeSiteId(v string) *CreateDesktopsRequest
	GetOfficeSiteId() *string
	SetOuPath(v string) *CreateDesktopsRequest
	GetOuPath() *string
	SetPeriod(v int32) *CreateDesktopsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateDesktopsRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateDesktopsRequest
	GetPolicyGroupId() *string
	SetPromotionId(v string) *CreateDesktopsRequest
	GetPromotionId() *string
	SetPurchaseOptions(v *CreateDesktopsRequestPurchaseOptions) *CreateDesktopsRequest
	GetPurchaseOptions() *CreateDesktopsRequestPurchaseOptions
	SetQosRuleId(v string) *CreateDesktopsRequest
	GetQosRuleId() *string
	SetRegionId(v string) *CreateDesktopsRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *CreateDesktopsRequest
	GetResellerOwnerUid() *int64
	SetResourceGroupId(v string) *CreateDesktopsRequest
	GetResourceGroupId() *string
	SetSavingPlanId(v string) *CreateDesktopsRequest
	GetSavingPlanId() *string
	SetSnapshotPolicyId(v string) *CreateDesktopsRequest
	GetSnapshotPolicyId() *string
	SetSubPayType(v string) *CreateDesktopsRequest
	GetSubPayType() *string
	SetSubnetId(v string) *CreateDesktopsRequest
	GetSubnetId() *string
	SetTag(v []*CreateDesktopsRequestTag) *CreateDesktopsRequest
	GetTag() []*CreateDesktopsRequestTag
	SetTimerGroupId(v string) *CreateDesktopsRequest
	GetTimerGroupId() *string
	SetUserAssignMode(v string) *CreateDesktopsRequest
	GetUserAssignMode() *string
	SetUserCommands(v []*CreateDesktopsRequestUserCommands) *CreateDesktopsRequest
	GetUserCommands() []*CreateDesktopsRequestUserCommands
	SetUserName(v string) *CreateDesktopsRequest
	GetUserName() *string
	SetVolumeEncryptionEnabled(v bool) *CreateDesktopsRequest
	GetVolumeEncryptionEnabled() *bool
	SetVolumeEncryptionKey(v string) *CreateDesktopsRequest
	GetVolumeEncryptionKey() *string
	SetVpcId(v string) *CreateDesktopsRequest
	GetVpcId() *string
}

type CreateDesktopsRequest struct {
	// The number of cloud desktops to create. Valid values: 1 to 300. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The application control policy ID.
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
	// Specifies whether to enable auto-renewal. This parameter takes effect only when `ChargeType` is set to `PrePaid`.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The cloud desktop template ID. If no template ID is specified, you can create a cloud desktop by specifying the required fields.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The list of cloud desktop templates.
	//
	// if can be null:
	// true
	BundleModels []*CreateDesktopsRequestBundleModels `json:"BundleModels,omitempty" xml:"BundleModels,omitempty" type:"Repeated"`
	// > This field is not available for use.
	//
	// example:
	//
	// PBKB1QbqEl2tslEuU6gRrLxvCFBU2M%2FVD0Eru6Oo%2FI9LTU3XQhvq3PGMWarE%2BPJdkNvCqT3blqlRSthNy4A%2BJQ%3D%3D
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The billing method of the cloud desktop.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The parameters for creating a cloud desktop without a template. This parameter is invalid when the BundleId parameter is specified.
	DesktopAttachment *CreateDesktopsRequestDesktopAttachment `json:"DesktopAttachment,omitempty" xml:"DesktopAttachment,omitempty" type:"Struct"`
	// The private IP address of the cloud desktop.
	//
	// example:
	//
	// 10.0.0.1
	DesktopMemberIp *string `json:"DesktopMemberIp,omitempty" xml:"DesktopMemberIp,omitempty"`
	// The cloud desktop name. The naming rules are as follows:
	//
	// - The name can be up to 64 characters in length.
	//
	// - The name must start with a letter or a Chinese character and cannot start with `http://` or `https://`.
	//
	// - The name can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// DemoComputer01
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// Specifies whether to automatically append a suffix to the cloud desktop name when you create multiple cloud desktops in a batch.
	//
	// example:
	//
	// false
	DesktopNameSuffix *bool `json:"DesktopNameSuffix,omitempty" xml:"DesktopNameSuffix,omitempty"`
	// The scheduled task details of the cloud desktop. This parameter is being deprecated. Use the TimerGroupId parameter instead.
	//
	// if can be null:
	// true
	DesktopTimers []*CreateDesktopsRequestDesktopTimers `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	// > This parameter is not available for use.
	//
	// example:
	//
	// cn-hangzhou+dir-300943****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The list of authorized user IDs to add to the cloud desktops. You can specify 1 to 100 user IDs.
	//
	// example:
	//
	// 123456789
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The extended information in JSON string format. This parameter is available only for internal customers.
	//
	// example:
	//
	// {}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The cloud desktop pool ID.
	//
	// example:
	//
	// dg-boyczi8enfyc5****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The custom hostname of the cloud desktop. Settings for this parameter are supported only for cloud desktops that run the Windows operating system in an AD office network.
	//
	// The naming rules for the hostname are as follows:
	//
	// - The hostname must be 2 to 15 characters in length.
	//
	// - The hostname can contain uppercase letters, lowercase letters, digits, or hyphens (-). It cannot start or end with a hyphen, contain consecutive hyphens, or consist of only digits.
	//
	// When you create multiple cloud desktops, you can use the `name_prefix[begin_number,bits]name_suffix` format to uniformly name the cloud desktops. For example, if you set Hostname to ecd-[1,4]-test, the hostname of the first cloud desktop is ecd-0001-test, the hostname of the second cloud desktop is ecd-0002-test, and so on.
	//
	// - `name_prefix`: the prefix of the hostname.
	//
	// - `[begin_number,bits]`: the sequential number in the hostname. `begin_number` is the starting number. Valid values: 0 to 999999. Default value: 0. `bits` is the number of digits. Valid values: 1 to 6. Default value: 6.
	//
	// - `name_suffix`: the suffix of the hostname.
	//
	// example:
	//
	// testhost
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The purchase parameters for the monthly hourly package.
	MonthDesktopSetting *CreateDesktopsRequestMonthDesktopSetting `json:"MonthDesktopSetting,omitempty" xml:"MonthDesktopSetting,omitempty" type:"Struct"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-387822****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The OU path. If specified, the cloud desktop is added to the corresponding organizational unit (OU) in Active Directory (AD).
	//
	// example:
	//
	// test.com/wuyingtest/computers
	OuPath *string `json:"OuPath,omitempty" xml:"OuPath,omitempty"`
	// The subscription duration of the resource. The unit is specified by `PeriodUnit`. This parameter takes effect and is required only when `ChargeType` is set to `PrePaid`.
	//
	// - If `PeriodUnit` is set to `Month`, valid values of this parameter:
	//
	//      - 1
	//
	//     -  2
	//
	//     - 3
	//
	//     - 6
	//
	// - If `PeriodUnit` is set to `Year`, valid values of this parameter:
	//
	//     - 1
	//
	//     - 2
	//
	//     - 3
	//
	//     - 4
	//
	//     - 5
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
	// The policy ID.
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
	PurchaseOptions *CreateDesktopsRequestPurchaseOptions `json:"PurchaseOptions,omitempty" xml:"PurchaseOptions,omitempty" type:"Struct"`
	// The public network rate limiting rule ID.
	//
	// example:
	//
	// qos-52fqmg6kvyro7zu4l
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID for resource ownership in reseller pattern. This parameter is not required in non-reseller pattern.
	//
	// example:
	//
	// 1828644634819902
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The WUYING resource group ID.
	//
	// example:
	//
	// rg-3mtuc28rx95lx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// > This field is not available for use.
	//
	// example:
	//
	// spn-26c1b7bcrjcI****
	SavingPlanId *string `json:"SavingPlanId,omitempty" xml:"SavingPlanId,omitempty"`
	// The WUYING automatic snapshot policy ID.
	//
	// example:
	//
	// sp-28mp6my0l6zow****
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	SubPayType       *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// The subnet ID.
	//
	// example:
	//
	// vsw-bp1m*****
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The tags.
	Tag []*CreateDesktopsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The scheduled task group ID.
	//
	// example:
	//
	// ccg-0caoeogrk9m5****
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	// The cloud desktop assignment mode.
	//
	// > If `EndUserId` is not specified, the created cloud desktops are not assigned to any user.
	//
	// example:
	//
	// ALL
	UserAssignMode *string `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
	// The custom command script data.
	UserCommands []*CreateDesktopsRequestUserCommands `json:"UserCommands,omitempty" xml:"UserCommands,omitempty" type:"Repeated"`
	// > This parameter is not available for use.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Specifies whether to enable cloud disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the Key Management Service (KMS) key used for cloud disk encryption. You can call [ListKeys](https://help.aliyun.com/document_detail/28951.html) to obtain the key ID.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
	// > This parameter is not available for use.
	//
	// example:
	//
	// vpc-uf6w8u60n8xbkg5el****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateDesktopsRequest) GetAppRuleId() *string {
	return s.AppRuleId
}

func (s *CreateDesktopsRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateDesktopsRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDesktopsRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateDesktopsRequest) GetBundleModels() []*CreateDesktopsRequestBundleModels {
	return s.BundleModels
}

func (s *CreateDesktopsRequest) GetChannelCookie() *string {
	return s.ChannelCookie
}

func (s *CreateDesktopsRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateDesktopsRequest) GetDesktopAttachment() *CreateDesktopsRequestDesktopAttachment {
	return s.DesktopAttachment
}

func (s *CreateDesktopsRequest) GetDesktopMemberIp() *string {
	return s.DesktopMemberIp
}

func (s *CreateDesktopsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *CreateDesktopsRequest) GetDesktopNameSuffix() *bool {
	return s.DesktopNameSuffix
}

func (s *CreateDesktopsRequest) GetDesktopTimers() []*CreateDesktopsRequestDesktopTimers {
	return s.DesktopTimers
}

func (s *CreateDesktopsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateDesktopsRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *CreateDesktopsRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *CreateDesktopsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateDesktopsRequest) GetHostname() *string {
	return s.Hostname
}

func (s *CreateDesktopsRequest) GetMonthDesktopSetting() *CreateDesktopsRequestMonthDesktopSetting {
	return s.MonthDesktopSetting
}

func (s *CreateDesktopsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateDesktopsRequest) GetOuPath() *string {
	return s.OuPath
}

func (s *CreateDesktopsRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateDesktopsRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateDesktopsRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateDesktopsRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateDesktopsRequest) GetPurchaseOptions() *CreateDesktopsRequestPurchaseOptions {
	return s.PurchaseOptions
}

func (s *CreateDesktopsRequest) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *CreateDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDesktopsRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *CreateDesktopsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDesktopsRequest) GetSavingPlanId() *string {
	return s.SavingPlanId
}

func (s *CreateDesktopsRequest) GetSnapshotPolicyId() *string {
	return s.SnapshotPolicyId
}

func (s *CreateDesktopsRequest) GetSubPayType() *string {
	return s.SubPayType
}

func (s *CreateDesktopsRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *CreateDesktopsRequest) GetTag() []*CreateDesktopsRequestTag {
	return s.Tag
}

func (s *CreateDesktopsRequest) GetTimerGroupId() *string {
	return s.TimerGroupId
}

func (s *CreateDesktopsRequest) GetUserAssignMode() *string {
	return s.UserAssignMode
}

func (s *CreateDesktopsRequest) GetUserCommands() []*CreateDesktopsRequestUserCommands {
	return s.UserCommands
}

func (s *CreateDesktopsRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateDesktopsRequest) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *CreateDesktopsRequest) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *CreateDesktopsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDesktopsRequest) SetAmount(v int32) *CreateDesktopsRequest {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsRequest) SetAppRuleId(v string) *CreateDesktopsRequest {
	s.AppRuleId = &v
	return s
}

func (s *CreateDesktopsRequest) SetAutoPay(v bool) *CreateDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopsRequest) SetAutoRenew(v bool) *CreateDesktopsRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDesktopsRequest) SetBundleId(v string) *CreateDesktopsRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsRequest) SetBundleModels(v []*CreateDesktopsRequestBundleModels) *CreateDesktopsRequest {
	s.BundleModels = v
	return s
}

func (s *CreateDesktopsRequest) SetChannelCookie(v string) *CreateDesktopsRequest {
	s.ChannelCookie = &v
	return s
}

func (s *CreateDesktopsRequest) SetChargeType(v string) *CreateDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopAttachment(v *CreateDesktopsRequestDesktopAttachment) *CreateDesktopsRequest {
	s.DesktopAttachment = v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopMemberIp(v string) *CreateDesktopsRequest {
	s.DesktopMemberIp = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopName(v string) *CreateDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopNameSuffix(v bool) *CreateDesktopsRequest {
	s.DesktopNameSuffix = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopTimers(v []*CreateDesktopsRequestDesktopTimers) *CreateDesktopsRequest {
	s.DesktopTimers = v
	return s
}

func (s *CreateDesktopsRequest) SetDirectoryId(v string) *CreateDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopsRequest) SetEndUserId(v []*string) *CreateDesktopsRequest {
	s.EndUserId = v
	return s
}

func (s *CreateDesktopsRequest) SetExtendInfo(v string) *CreateDesktopsRequest {
	s.ExtendInfo = &v
	return s
}

func (s *CreateDesktopsRequest) SetGroupId(v string) *CreateDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetHostname(v string) *CreateDesktopsRequest {
	s.Hostname = &v
	return s
}

func (s *CreateDesktopsRequest) SetMonthDesktopSetting(v *CreateDesktopsRequestMonthDesktopSetting) *CreateDesktopsRequest {
	s.MonthDesktopSetting = v
	return s
}

func (s *CreateDesktopsRequest) SetOfficeSiteId(v string) *CreateDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopsRequest) SetOuPath(v string) *CreateDesktopsRequest {
	s.OuPath = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriod(v int32) *CreateDesktopsRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriodUnit(v string) *CreateDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopsRequest) SetPolicyGroupId(v string) *CreateDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetPromotionId(v string) *CreateDesktopsRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetPurchaseOptions(v *CreateDesktopsRequestPurchaseOptions) *CreateDesktopsRequest {
	s.PurchaseOptions = v
	return s
}

func (s *CreateDesktopsRequest) SetQosRuleId(v string) *CreateDesktopsRequest {
	s.QosRuleId = &v
	return s
}

func (s *CreateDesktopsRequest) SetRegionId(v string) *CreateDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetResellerOwnerUid(v int64) *CreateDesktopsRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *CreateDesktopsRequest) SetResourceGroupId(v string) *CreateDesktopsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetSavingPlanId(v string) *CreateDesktopsRequest {
	s.SavingPlanId = &v
	return s
}

func (s *CreateDesktopsRequest) SetSnapshotPolicyId(v string) *CreateDesktopsRequest {
	s.SnapshotPolicyId = &v
	return s
}

func (s *CreateDesktopsRequest) SetSubPayType(v string) *CreateDesktopsRequest {
	s.SubPayType = &v
	return s
}

func (s *CreateDesktopsRequest) SetSubnetId(v string) *CreateDesktopsRequest {
	s.SubnetId = &v
	return s
}

func (s *CreateDesktopsRequest) SetTag(v []*CreateDesktopsRequestTag) *CreateDesktopsRequest {
	s.Tag = v
	return s
}

func (s *CreateDesktopsRequest) SetTimerGroupId(v string) *CreateDesktopsRequest {
	s.TimerGroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetUserAssignMode(v string) *CreateDesktopsRequest {
	s.UserAssignMode = &v
	return s
}

func (s *CreateDesktopsRequest) SetUserCommands(v []*CreateDesktopsRequestUserCommands) *CreateDesktopsRequest {
	s.UserCommands = v
	return s
}

func (s *CreateDesktopsRequest) SetUserName(v string) *CreateDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *CreateDesktopsRequest) SetVolumeEncryptionEnabled(v bool) *CreateDesktopsRequest {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *CreateDesktopsRequest) SetVolumeEncryptionKey(v string) *CreateDesktopsRequest {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *CreateDesktopsRequest) SetVpcId(v string) *CreateDesktopsRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDesktopsRequest) Validate() error {
	if s.BundleModels != nil {
		for _, item := range s.BundleModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DesktopAttachment != nil {
		if err := s.DesktopAttachment.Validate(); err != nil {
			return err
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
	if s.PurchaseOptions != nil {
		if err := s.PurchaseOptions.Validate(); err != nil {
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

type CreateDesktopsRequestBundleModels struct {
	// The number of cloud desktops to create. Valid values: 1 to 300. Default value: 0.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The cloud desktop template ID.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The cloud desktop name. The naming rules are as follows:
	//
	// - The name can be up to 64 characters in length.
	//
	// - The name must start with a letter or a Chinese character and cannot start with `http://` or `https://`.
	//
	// - The name can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// DemoComputer02
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The list of users to whom the cloud desktops are assigned.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The custom hostname of the cloud desktop. Settings for this parameter are supported only for cloud desktops that run the Windows operating system in an AD office network.
	//
	// The naming rules for the hostname are as follows:
	//
	// - The hostname must be 2 to 15 characters in length.
	//
	// - The hostname can contain uppercase letters, lowercase letters, digits, or hyphens (-). It cannot start or end with a hyphen, contain consecutive hyphens, or consist of only digits.
	//
	// When you create multiple cloud desktops, you can use the `name_prefix[begin_number,bits]name_suffix` format to uniformly name the cloud desktops. For example, if you set Hostname to ecd-[1,4]-test, the hostname of the first cloud desktop is ecd-0001-test, the hostname of the second cloud desktop is ecd-0002-test, and so on.
	//
	// - `name_prefix`: the prefix of the hostname.
	//
	// - `[begin_number,bits]`: the sequential number in the hostname. `begin_number` is the starting number. Valid values: 0 to 999999. Default value: 0. `bits` is the number of digits. Valid values: 1 to 6. Default value: 6.
	//
	// - `name_suffix`: the suffix of the hostname.
	//
	// example:
	//
	// testhost
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Specifies whether to enable cloud disk encryption.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the Key Management Service (KMS) key used for cloud disk encryption. You can call [ListKeys](https://help.aliyun.com/document_detail/28951.html) to obtain the key ID.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s CreateDesktopsRequestBundleModels) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsRequestBundleModels) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestBundleModels) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateDesktopsRequestBundleModels) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateDesktopsRequestBundleModels) GetDesktopName() *string {
	return s.DesktopName
}

func (s *CreateDesktopsRequestBundleModels) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *CreateDesktopsRequestBundleModels) GetHostname() *string {
	return s.Hostname
}

func (s *CreateDesktopsRequestBundleModels) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *CreateDesktopsRequestBundleModels) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *CreateDesktopsRequestBundleModels) SetAmount(v int32) *CreateDesktopsRequestBundleModels {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsRequestBundleModels) SetBundleId(v string) *CreateDesktopsRequestBundleModels {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsRequestBundleModels) SetDesktopName(v string) *CreateDesktopsRequestBundleModels {
	s.DesktopName = &v
	return s
}

func (s *CreateDesktopsRequestBundleModels) SetEndUserIds(v []*string) *CreateDesktopsRequestBundleModels {
	s.EndUserIds = v
	return s
}

func (s *CreateDesktopsRequestBundleModels) SetHostname(v string) *CreateDesktopsRequestBundleModels {
	s.Hostname = &v
	return s
}

func (s *CreateDesktopsRequestBundleModels) SetVolumeEncryptionEnabled(v bool) *CreateDesktopsRequestBundleModels {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *CreateDesktopsRequestBundleModels) SetVolumeEncryptionKey(v string) *CreateDesktopsRequestBundleModels {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *CreateDesktopsRequestBundleModels) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsRequestDesktopAttachment struct {
	// The data cloud disk type. The system cloud disk type must be the same as the data cloud disk type. Valid values:
	//
	// - cloud_auto: standard SSD ultra cloud disk
	//
	// - cloud_essd: ESSD cloud disk
	//
	// example:
	//
	// cloud_auto
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. This parameter is required when an ESSD cloud disk is selected. Valid values:
	//
	// - PL0
	//
	// - PL1
	//
	// example:
	//
	// PL0
	DataDiskPerLevel *string `json:"DataDiskPerLevel,omitempty" xml:"DataDiskPerLevel,omitempty"`
	// The user cloud disk capacity. Valid values: 40 to 2040 GiB, in increments of 10 GiB.
	//
	// example:
	//
	// 40
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The language. Valid values:
	//
	// - zh-CN
	//
	// - zh-HK
	//
	// - en-US
	//
	// - ja-JP
	//
	// example:
	//
	// zh-CN
	DefaultLanguage *string `json:"DefaultLanguage,omitempty" xml:"DefaultLanguage,omitempty"`
	// The cloud desktop specification. You can call [DescribeDesktopTypes](https://help.aliyun.com/document_detail/188882.html) to query the supported specification IDs.
	//
	// example:
	//
	// eds.enterprise_office.8c16g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-39ddhdb0ggzjx*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The system cloud disk type. The system cloud disk type must be the same as the data cloud disk type. Valid values:
	//
	// - cloud_auto: standard SSD ultra cloud disk
	//
	// - cloud_essd: ESSD cloud disk
	//
	// example:
	//
	// cloud_auto
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. This parameter is required when an ESSD cloud disk is selected. Valid values:
	//
	// - PL0
	//
	// - PL1
	//
	// example:
	//
	// PL0
	SystemDiskPerLevel *string `json:"SystemDiskPerLevel,omitempty" xml:"SystemDiskPerLevel,omitempty"`
	// The system cloud disk capacity. Valid values: 60 to 500 GiB, in increments of 10 GiB.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s CreateDesktopsRequestDesktopAttachment) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsRequestDesktopAttachment) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestDesktopAttachment) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *CreateDesktopsRequestDesktopAttachment) GetDataDiskPerLevel() *string {
	return s.DataDiskPerLevel
}

func (s *CreateDesktopsRequestDesktopAttachment) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *CreateDesktopsRequestDesktopAttachment) GetDefaultLanguage() *string {
	return s.DefaultLanguage
}

func (s *CreateDesktopsRequestDesktopAttachment) GetDesktopType() *string {
	return s.DesktopType
}

func (s *CreateDesktopsRequestDesktopAttachment) GetImageId() *string {
	return s.ImageId
}

func (s *CreateDesktopsRequestDesktopAttachment) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateDesktopsRequestDesktopAttachment) GetSystemDiskPerLevel() *string {
	return s.SystemDiskPerLevel
}

func (s *CreateDesktopsRequestDesktopAttachment) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateDesktopsRequestDesktopAttachment) SetDataDiskCategory(v string) *CreateDesktopsRequestDesktopAttachment {
	s.DataDiskCategory = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) SetDataDiskPerLevel(v string) *CreateDesktopsRequestDesktopAttachment {
	s.DataDiskPerLevel = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) SetDataDiskSize(v int32) *CreateDesktopsRequestDesktopAttachment {
	s.DataDiskSize = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) SetDefaultLanguage(v string) *CreateDesktopsRequestDesktopAttachment {
	s.DefaultLanguage = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) SetDesktopType(v string) *CreateDesktopsRequestDesktopAttachment {
	s.DesktopType = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) SetImageId(v string) *CreateDesktopsRequestDesktopAttachment {
	s.ImageId = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) SetSystemDiskCategory(v string) *CreateDesktopsRequestDesktopAttachment {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) SetSystemDiskPerLevel(v string) *CreateDesktopsRequestDesktopAttachment {
	s.SystemDiskPerLevel = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) SetSystemDiskSize(v int32) *CreateDesktopsRequestDesktopAttachment {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateDesktopsRequestDesktopAttachment) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsRequestDesktopTimers struct {
	// Specifies whether to allow end users to configure scheduled tasks.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The cron expression of the scheduled task.
	//
	// 	Notice: Specify the time in UTC. For example, to schedule a task at 00:00 (UTC+8) every day, set the value to 0 0 16 ? 	- 1,2,3,4,5,6,7.</notice>
	//
	// example:
	//
	// 0 40 7 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to forcefully execute the task.
	//
	// example:
	//
	// true
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The time interval, in minutes.
	//
	// example:
	//
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The operation type of the scheduled task. Currently, only the disconnection scheduled task is supported.
	//
	// example:
	//
	// Shutdown
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The reset type of the cloud desktop.
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

func (s CreateDesktopsRequestDesktopTimers) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsRequestDesktopTimers) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestDesktopTimers) GetAllowClientSetting() *bool {
	return s.AllowClientSetting
}

func (s *CreateDesktopsRequestDesktopTimers) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateDesktopsRequestDesktopTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *CreateDesktopsRequestDesktopTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateDesktopsRequestDesktopTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateDesktopsRequestDesktopTimers) GetResetType() *string {
	return s.ResetType
}

func (s *CreateDesktopsRequestDesktopTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *CreateDesktopsRequestDesktopTimers) SetAllowClientSetting(v bool) *CreateDesktopsRequestDesktopTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *CreateDesktopsRequestDesktopTimers) SetCronExpression(v string) *CreateDesktopsRequestDesktopTimers {
	s.CronExpression = &v
	return s
}

func (s *CreateDesktopsRequestDesktopTimers) SetEnforce(v bool) *CreateDesktopsRequestDesktopTimers {
	s.Enforce = &v
	return s
}

func (s *CreateDesktopsRequestDesktopTimers) SetInterval(v int32) *CreateDesktopsRequestDesktopTimers {
	s.Interval = &v
	return s
}

func (s *CreateDesktopsRequestDesktopTimers) SetOperationType(v string) *CreateDesktopsRequestDesktopTimers {
	s.OperationType = &v
	return s
}

func (s *CreateDesktopsRequestDesktopTimers) SetResetType(v string) *CreateDesktopsRequestDesktopTimers {
	s.ResetType = &v
	return s
}

func (s *CreateDesktopsRequestDesktopTimers) SetTimerType(v string) *CreateDesktopsRequestDesktopTimers {
	s.TimerType = &v
	return s
}

func (s *CreateDesktopsRequestDesktopTimers) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsRequestMonthDesktopSetting struct {
	// > This field is not available for use.
	//
	// example:
	//
	// null
	BuyerId *int64 `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	// > This field is not available for use.
	//
	// example:
	//
	// null
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The package option when purchasing a monthly hourly package. Valid values: 120, 250, and 360.
	//
	// example:
	//
	// null
	UseDuration *int32 `json:"UseDuration,omitempty" xml:"UseDuration,omitempty"`
}

func (s CreateDesktopsRequestMonthDesktopSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsRequestMonthDesktopSetting) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestMonthDesktopSetting) GetBuyerId() *int64 {
	return s.BuyerId
}

func (s *CreateDesktopsRequestMonthDesktopSetting) GetDesktopId() *string {
	return s.DesktopId
}

func (s *CreateDesktopsRequestMonthDesktopSetting) GetUseDuration() *int32 {
	return s.UseDuration
}

func (s *CreateDesktopsRequestMonthDesktopSetting) SetBuyerId(v int64) *CreateDesktopsRequestMonthDesktopSetting {
	s.BuyerId = &v
	return s
}

func (s *CreateDesktopsRequestMonthDesktopSetting) SetDesktopId(v string) *CreateDesktopsRequestMonthDesktopSetting {
	s.DesktopId = &v
	return s
}

func (s *CreateDesktopsRequestMonthDesktopSetting) SetUseDuration(v int32) *CreateDesktopsRequestMonthDesktopSetting {
	s.UseDuration = &v
	return s
}

func (s *CreateDesktopsRequestMonthDesktopSetting) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsRequestPurchaseOptions struct {
	// The monthly credit package for purchasing Agent resources. Valid values: 200, 1600, and 4000.
	//
	// example:
	//
	// 200
	MonthlyCredits *int32 `json:"MonthlyCredits,omitempty" xml:"MonthlyCredits,omitempty"`
}

func (s CreateDesktopsRequestPurchaseOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsRequestPurchaseOptions) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestPurchaseOptions) GetMonthlyCredits() *int32 {
	return s.MonthlyCredits
}

func (s *CreateDesktopsRequestPurchaseOptions) SetMonthlyCredits(v int32) *CreateDesktopsRequestPurchaseOptions {
	s.MonthlyCredits = &v
	return s
}

func (s *CreateDesktopsRequestPurchaseOptions) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsRequestTag struct {
	// The tag key. You can specify 1 to 20 tag keys.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify 1 to 20 tag values.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDesktopsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDesktopsRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDesktopsRequestTag) SetKey(v string) *CreateDesktopsRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDesktopsRequestTag) SetValue(v string) *CreateDesktopsRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDesktopsRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopsRequestUserCommands struct {
	// The command content.
	//
	// example:
	//
	// bmV3LWl0ZW0gZDpcdGVzdF91c2VyX2NvbW1hbmRzLnR4dCAtdHlwZSBm****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The encoding method of the command content (CommandContent).
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The language type of the command.
	//
	// example:
	//
	// RunPowerShellScript
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s CreateDesktopsRequestUserCommands) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsRequestUserCommands) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestUserCommands) GetContent() *string {
	return s.Content
}

func (s *CreateDesktopsRequestUserCommands) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *CreateDesktopsRequestUserCommands) GetContentType() *string {
	return s.ContentType
}

func (s *CreateDesktopsRequestUserCommands) SetContent(v string) *CreateDesktopsRequestUserCommands {
	s.Content = &v
	return s
}

func (s *CreateDesktopsRequestUserCommands) SetContentEncoding(v string) *CreateDesktopsRequestUserCommands {
	s.ContentEncoding = &v
	return s
}

func (s *CreateDesktopsRequestUserCommands) SetContentType(v string) *CreateDesktopsRequestUserCommands {
	s.ContentType = &v
	return s
}

func (s *CreateDesktopsRequestUserCommands) Validate() error {
	return dara.Validate(s)
}
