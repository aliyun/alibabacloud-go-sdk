// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndroidInstanceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateAndroidInstanceGroupShrinkRequest
	GetAmount() *int32
	SetAutoPay(v bool) *CreateAndroidInstanceGroupShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateAndroidInstanceGroupShrinkRequest
	GetAutoRenew() *bool
	SetBandwidthPackageId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetBandwidthPackageId() *string
	SetBandwidthPackageType(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetBandwidthPackageType() *string
	SetBizRegionId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetBizRegionId() *string
	SetChargeType(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetClientToken() *string
	SetEnableIpv6(v bool) *CreateAndroidInstanceGroupShrinkRequest
	GetEnableIpv6() *bool
	SetGpuAcceleration(v bool) *CreateAndroidInstanceGroupShrinkRequest
	GetGpuAcceleration() *bool
	SetImageId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetImageId() *string
	SetInstanceGroupName(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetInstanceGroupName() *string
	SetInstanceGroupSpec(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetInstanceGroupSpec() *string
	SetIpv6Bandwidth(v int32) *CreateAndroidInstanceGroupShrinkRequest
	GetIpv6Bandwidth() *int32
	SetKeyPairId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetKeyPairId() *string
	SetNetworkInfoShrink(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetNetworkInfoShrink() *string
	SetNetworkType(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetNetworkType() *string
	SetNumberOfInstances(v int32) *CreateAndroidInstanceGroupShrinkRequest
	GetNumberOfInstances() *int32
	SetOfficeSiteId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetOfficeSiteId() *string
	SetPeriod(v int32) *CreateAndroidInstanceGroupShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetPolicyGroupId() *string
	SetPromotionId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetPromotionId() *string
	SetTag(v []*CreateAndroidInstanceGroupShrinkRequestTag) *CreateAndroidInstanceGroupShrinkRequest
	GetTag() []*CreateAndroidInstanceGroupShrinkRequestTag
	SetVSwitchId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetVSwitchId() *string
}

type CreateAndroidInstanceGroupShrinkRequest struct {
	// The number of instance groups. Default value: 1. Maximum value: 1.
	//
	// example:
	//
	// 8
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable automatic payment. Default value: false.
	//
	// Valid values:
	//
	// 	- true: enables automatic payment. Make sure that your Alibaba Cloud account has sufficient balance.
	//
	// 	- false: disables automatic payment. You must manually complete the payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Default value: false.
	//
	// Valid values:
	//
	// 	- true: automatically renew resource upon expiration.
	//
	// 	- false: manually renew resources upon expiration.
	//
	// example:
	//
	// false
	AutoRenew            *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BandwidthPackageId   *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	// The ID of the region. You can call the DescribeRegions operation to query the regions where Cloud Phone is supported.
	//
	// Valid values:
	//
	// 	- cn-shenzhen: China (Shenzhen).
	//
	// 	- cn-beijing: China (Beijing).
	//
	// 	- cn-shanghai: China (Shanghai).
	//
	// 	- cn-hongkong: China (Hong Kong).
	//
	// 	- ap-southeast-1: Singapore.
	//
	// 	- cn-hangzhou: China (Hangzhou).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The value cannot exceed 100 characters in length.
	//
	// example:
	//
	// asadbuvwiabdbvchjsbj
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// Specifies whether to enable GPU acceleration.
	//
	// Valid values:
	//
	// 	- true: enables GPU acceleration.
	//
	// 	- false (default): disables GPU acceleration.
	//
	// example:
	//
	// false
	GpuAcceleration *bool `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// The ID of the image. You can call the [DescribeImageList](https://help.aliyun.com/document_detail/2807324.html) operation to query images.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-06zyt9m93zwax****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the instance group.
	//
	// >  The name can be up to 30 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), or hyphens (-). It must start with letters but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// defaultInstanceGroup
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The specifications of the instance group. You can call the [DescribeSpec](https://help.aliyun.com/document_detail/2807299.html) operation to query the available specifications.
	//
	// Valid values:
	//
	// 	- acp.perf.large: Performance (8 vCPUs, 16 GiB of memory, and 32 GiB of storage.
	//
	// 	- acp.basic.small: Lightweight (2 vCPUs, 4 GiB of memory, and 32 GiB of storage).
	//
	// 	- acp.std.large: Standard (4 vCPUs, 8 GiB of memory, and 32 GiB of storage).
	//
	// This parameter is required.
	//
	// example:
	//
	// acp.basic.small
	InstanceGroupSpec *string `json:"InstanceGroupSpec,omitempty" xml:"InstanceGroupSpec,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	Ipv6Bandwidth *int32 `json:"Ipv6Bandwidth,omitempty" xml:"Ipv6Bandwidth,omitempty"`
	// The ID of the key pair. When you create an instance group and specify a valid key pair ID, all cloud phone instances within the group will automatically be bound to that key pair upon creation. This eliminates the need to manually bind key pairs to individual cloud phone instances.
	//
	// >  Binding key pairs to cloud phone instances is currently not supported during instance group resizing.
	//
	// example:
	//
	// kp-7o9xywwfutc1l****
	KeyPairId         *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	NetworkInfoShrink *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The number of cloud phones in the instance group. Maximum value: 100.
	//
	// example:
	//
	// 1
	NumberOfInstances *int32 `json:"NumberOfInstances,omitempty" xml:"NumberOfInstances,omitempty"`
	// The ID of the network.
	//
	// 	- This parameter is required if you assign a shared network to cloud phones. You can go to the [Network](https://wya.wuying.aliyun.com/network) page of the Cloud Phone console to retrieve the ID of a **shared network**. If no shared network is available in the Cloud Phone console, you can leave this parameter empty. The system automatically creates one when you create an instance group.
	//
	// 	- This parameter is required if you assign a virtual private cloud (VPC) to cloud phones. You can go to the [Network](https://wya.wuying.aliyun.com/network) page of the Cloud Phone console to retrieve the ID of a **VPC**. If no VPC is available in the Cloud Phone console, you must first create one.
	//
	// example:
	//
	// cn-hangzhou+dir-745976****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The subscription duration. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// 	- Hour (Note that this unit is supported only by pay-as-you-go.)
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the policy. You can call the [ListPolicyGroups](https://help.aliyun.com/document_detail/2807352.html) operation to query policies.
	//
	// example:
	//
	// pg-b7bxrrwxkijjh****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The tags
	Tag []*CreateAndroidInstanceGroupShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/448774.html) operation to query vSwitches.
	//
	// 	- This parameter is not required if you assign a shared network to cloud phones.
	//
	// 	- This parameter is required if you assign a VPC to cloud phones. The vSwitch specified by this parameter is used to create cloud phones.
	//
	// example:
	//
	// vsw-uf61uvzhz8ejaw776****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateAndroidInstanceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetGpuAcceleration() *bool {
	return s.GpuAcceleration
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetInstanceGroupName() *string {
	return s.InstanceGroupName
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetInstanceGroupSpec() *string {
	return s.InstanceGroupSpec
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetIpv6Bandwidth() *int32 {
	return s.Ipv6Bandwidth
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetNetworkInfoShrink() *string {
	return s.NetworkInfoShrink
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetNumberOfInstances() *int32 {
	return s.NumberOfInstances
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetTag() []*CreateAndroidInstanceGroupShrinkRequestTag {
	return s.Tag
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetAmount(v int32) *CreateAndroidInstanceGroupShrinkRequest {
	s.Amount = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetAutoPay(v bool) *CreateAndroidInstanceGroupShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetAutoRenew(v bool) *CreateAndroidInstanceGroupShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetBandwidthPackageId(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetBandwidthPackageType(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.BandwidthPackageType = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetBizRegionId(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetChargeType(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetClientToken(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetEnableIpv6(v bool) *CreateAndroidInstanceGroupShrinkRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetGpuAcceleration(v bool) *CreateAndroidInstanceGroupShrinkRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetImageId(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetInstanceGroupName(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.InstanceGroupName = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetInstanceGroupSpec(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.InstanceGroupSpec = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetIpv6Bandwidth(v int32) *CreateAndroidInstanceGroupShrinkRequest {
	s.Ipv6Bandwidth = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetKeyPairId(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.KeyPairId = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetNetworkInfoShrink(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.NetworkInfoShrink = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetNetworkType(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetNumberOfInstances(v int32) *CreateAndroidInstanceGroupShrinkRequest {
	s.NumberOfInstances = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetOfficeSiteId(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetPeriod(v int32) *CreateAndroidInstanceGroupShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetPeriodUnit(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetPolicyGroupId(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetPromotionId(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetTag(v []*CreateAndroidInstanceGroupShrinkRequestTag) *CreateAndroidInstanceGroupShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetVSwitchId(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAndroidInstanceGroupShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAndroidInstanceGroupShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAndroidInstanceGroupShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAndroidInstanceGroupShrinkRequestTag) SetKey(v string) *CreateAndroidInstanceGroupShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequestTag) SetValue(v string) *CreateAndroidInstanceGroupShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
