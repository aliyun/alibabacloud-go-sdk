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
	SetPeriod(v int32) *CreateDesktopsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateDesktopsRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateDesktopsRequest
	GetPolicyGroupId() *string
	SetPromotionId(v string) *CreateDesktopsRequest
	GetPromotionId() *string
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
	// The number of cloud computers that you want to create. Valid values: 1 to 300. Default value: 1.
	//
	// example:
	//
	// 1
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppRuleId *string `json:"AppRuleId,omitempty" xml:"AppRuleId,omitempty"`
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when the ChargeType parameter is set to PrePaid.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the cloud computer template.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The cloud computer templates.
	//
	// if can be null:
	// true
	BundleModels  []*CreateDesktopsRequestBundleModels `json:"BundleModels,omitempty" xml:"BundleModels,omitempty" type:"Repeated"`
	ChannelCookie *string                              `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The billing method of the cloud computers.
	//
	// Default value: PostPaid. Valid values:
	//
	// 	- Postpaid: pay-as-you-go
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PrePaid: subscription
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The input parameters used when no templates are used.
	DesktopAttachment *CreateDesktopsRequestDesktopAttachment `json:"DesktopAttachment,omitempty" xml:"DesktopAttachment,omitempty" type:"Struct"`
	// The private IP address of the cloud computer.
	//
	// example:
	//
	// 10.0.0.1
	DesktopMemberIp *string `json:"DesktopMemberIp,omitempty" xml:"DesktopMemberIp,omitempty"`
	// The name of the cloud computer. The name must meet the following requirements:
	//
	// 	- The name must be 1 to 64 characters in length.
	//
	// 	- The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The name can only contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// testDesktopName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// Specifies whether to automatically add suffixes to the names of cloud computers when you create multiple cloud computers at the same time.
	//
	// Default value: true. Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- False
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	DesktopNameSuffix *bool `json:"DesktopNameSuffix,omitempty" xml:"DesktopNameSuffix,omitempty"`
	// The details of the scheduled task on cloud computers.
	//
	// if can be null:
	// true
	DesktopTimers []*CreateDesktopsRequestDesktopTimers `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// To be hidden.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The IDs of the end users to which you want to assign the cloud computers. You can specify 1 to 100 IDs.
	//
	// example:
	//
	// 123456789
	EndUserId  []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	ExtendInfo *string   `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The ID of the cloud computer pool.
	//
	// example:
	//
	// dg-boyczi8enfyc5****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The custom hostnames of the cloud computers. This parameter is valid only if the office network is an AD office network and the operating system type of the cloud computers is Windows.
	//
	// The hostnames must meet the following requirements:
	//
	// 	- The hostnames must be 2 to 15 characters in length.
	//
	// 	- The hostnames can contain only letters, digits, and hyphens (-). The hostnames cannot start or end with a hyphen (-), contain consecutive hyphens (-), or contain only digits.
	//
	// When you create multiple cloud computers, you can use the `name_prefix[begin_number,bits]name_suffix` naming format to name the cloud computers. For example, if you set the value of the Hostname parameter to ecd-[1,4]-test, the hostname of the first cloud computer is ecd-0001-test, the hostname of the second cloud computer is ecd-0002-test, and so on.
	//
	// 	- `name_prefix`: the prefix of the hostname.
	//
	// 	- `[begin_number,bits]`: the sequential number in the hostname. The `begin_number` value is the starting digit. Valid values of begin_number: 0 to 999999. Default value: 0. The `bits` value is the number of digits. Valid values: 1 to 6. Default value: 6.
	//
	// 	- `name_suffix`: the suffix of the hostname.
	//
	// example:
	//
	// testhost
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// > This parameter is not publicly available.
	MonthDesktopSetting *CreateDesktopsRequestMonthDesktopSetting `json:"MonthDesktopSetting,omitempty" xml:"MonthDesktopSetting,omitempty" type:"Struct"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+os-c5cy7q578s8jc****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The subscription duration of the cloud desktop that you want to create. The unit is specified by the `PeriodUnit` parameter. This parameter takes effect and is required only when the `ChargeType` parameter is set to `PrePaid`.
	//
	// 	- Valid values if the `PeriodUnit` parameter is set to `Month`:
	//
	//     	- 1
	//
	//     	- 2
	//
	//     	- 3
	//
	//     	- 6
	//
	// 	- Valid values if the `PeriodUnit` parameter is set to `Year`:
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
	// example:
	//
	// system-all-enabled-policy
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The ID of the sales promotion.
	//
	// example:
	//
	// 23141
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	QosRuleId   *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-3mtuc28rx95lx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the saving plan.
	//
	// example:
	//
	// spn-4b945dc4Wktd****
	SavingPlanId *string `json:"SavingPlanId,omitempty" xml:"SavingPlanId,omitempty"`
	// The ID of the auto-snapshot policy.
	//
	// example:
	//
	// sp-28mp6my0l6zow****
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	SubnetId         *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The tags that you want to add to the cloud desktop.
	Tag []*CreateDesktopsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the timer group.
	//
	// example:
	//
	// ccg-0caoeogrk9m5****
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	// How the cloud computers are assigned.
	//
	// >  If you do not specify the `EndUserId` parameter, the cloud computers are not assigned to end users after the cloud computers are created.
	//
	// Default value: ALL. Valid values:
	//
	// 	- ALL: If you specify the EndUserId parameter, the cloud computers are assigned to all specified end users after the cloud computers are created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PER_USER: If you specify the EndUserId parameter, the cloud computers are evenly assigned to the specified end users after the cloud computers are created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     In this case, you must make sure that the value of the Amount parameter can be divided by the N value of the EndUserId.N parameter that you specify.
	//
	//     <!-- -->
	//
	// example:
	//
	// ALL
	UserAssignMode *string `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
	// Details about the custom command scripts.
	UserCommands []*CreateDesktopsRequestUserCommands `json:"UserCommands,omitempty" xml:"UserCommands,omitempty" type:"Repeated"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// To be hidden.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// To be hidden.
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
	// The number of cloud computers that you want to create. Valid values: 1 to 300. Default value: null.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of a cloud computer template.
	//
	// example:
	//
	// b-je9hani001wfn****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The name of the cloud computer. The name must meet the following requirements:
	//
	// 	- The name must be 1 to 64 characters in length.
	//
	// 	- The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The name can only contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// testDesktopName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The IDs of the end users to whom the cloud computer are assigned.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The custom hostnames of the cloud computers. This parameter is valid only if the office network is an AD office network and the operating system type of the cloud computers is Windows.
	//
	// The hostnames must meet the following requirements:
	//
	// 	- The hostnames must be 2 to 15 characters in length.
	//
	// 	- The hostnames can contain only letters, digits, and hyphens (-). The hostnames cannot start or end with a hyphen (-), contain consecutive hyphens (-), or contain only digits.
	//
	// When you create multiple cloud computers, you can use the `name_prefix[begin_number,bits]name_suffix` naming format to name the cloud computers. For example, if you set the value of the Hostname parameter to ecd-[1,4]-test, the hostname of the first cloud computer is ecd-0001-test, the hostname of the second cloud computer is ecd-0002-test, and so on.
	//
	// 	- `name_prefix`: the prefix of the hostname.
	//
	// 	- `[begin_number,bits]`: the sequential number in the hostname. The `begin_number` value is the starting digit. Valid values of begin_number: 0 to 999999. Default value: 0. The `bits` value is the number of digits. Valid values: 1 to 6. Default value: 6.
	//
	// 	- `name_suffix`: the suffix of the hostname.
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
	// The ID of the Key Management Service (KMS) key that is used when disk encryption is enabled. You can call the [ListKeys](https://help.aliyun.com/document_detail/28951.html) operation to query the list of KMS keys.
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
	// The category of the data disk. Valid values:
	//
	// 	- cloud_auto: SSD
	//
	// 	- cloud_essd: ESSD
	//
	// example:
	//
	// cloud_auto
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The performance level of the data disk. Valid values:
	//
	// - PL0 (default)
	//
	// - PL1
	//
	// example:
	//
	// PL0
	DataDiskPerLevel *string `json:"DataDiskPerLevel,omitempty" xml:"DataDiskPerLevel,omitempty"`
	// The size of the data disk. Unit: GiB.
	//
	// example:
	//
	// 40
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The default display language:
	//
	// - zh-CN: Simplified Chinese
	//
	// - zh-HK: Traditional Chinese
	//
	// - en-US: English
	//
	// - ja-JP: Japanese
	//
	// example:
	//
	// zh-CN
	DefaultLanguage *string `json:"DefaultLanguage,omitempty" xml:"DefaultLanguage,omitempty"`
	// The desktop type. You can call the [DescribeDesktopTypes](~~DescribeDesktopTypes~~) operation to query the IDs of supported desktop types.
	//
	// example:
	//
	// eds.enterprise_office.8c16g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-39ddhdb0ggzjx*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The category of the system disk. Valid values:
	//
	// 	- cloud_auto: SSD
	//
	// 	- cloud_essd: ESSD
	//
	// example:
	//
	// cloud_auto
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The performance level of the system disk. Valid values:
	//
	// - PL0 (default)
	//
	// - PL1
	//
	// example:
	//
	// PL0
	SystemDiskPerLevel *string `json:"SystemDiskPerLevel,omitempty" xml:"SystemDiskPerLevel,omitempty"`
	// The size of the system disk. Unit: GiB.
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
	// Specifies whether to allow the end user to configure the scheduled task.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The cron expression for the scheduled task.
	//
	// >  The time must be in UTC. For example, for 24:00 (UTC+8), you must set the value to 0 0 16 ? \\	- 1,2,3,4,5,6,7
	//
	// example:
	//
	// 0 40 7 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to forcibly execute the scheduled task.
	//
	// Valid values:
	//
	// 	- true: forcibly executes the scheduled task regardless of the status and connection of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false: does not forcibly execute the scheduled task.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// True
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The interval at which cloud computers are created. Unit: minutes.
	//
	// example:
	//
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The operations that scheduled tasks support. This parameter is valid only when TimerType is set to NoConnect.
	//
	// Valid values:
	//
	// 	- Hibernate: hibernates the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Shutdown: stops the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Shutdown
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The reset type of the cloud computers.
	//
	// Valid values:
	//
	// 	- RESET_TYPE_SYSTEM: resets the system disks.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- RESET_TYPE_BOTH: resets the system disks and data disks.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
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
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	BuyerId *int64 `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// > This parameter is not publicly available.
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

type CreateDesktopsRequestTag struct {
	// The key of the tag. You can specify 1 to 20 keys for a tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify 1 to 20 values for a tag.
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
	// The encoding mode of the command content.
	//
	// Valid values:
	//
	// 	- Base64: encodes the command content in Base64.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PlainText: does not encode the command content.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The language type of the command.
	//
	// Valid values:
	//
	// 	- RunPowerShellScript: PowerShell commands (applicable to Windows cloud computers).
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- RunShellScript: shell commands (applicable to Linux cloud computers).
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- RunBatScript: batch commands (applicable to Windows cloud computers).
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
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
