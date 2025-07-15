// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndroidInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateAndroidInstanceGroupRequest
	GetAmount() *int32
	SetAutoPay(v bool) *CreateAndroidInstanceGroupRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateAndroidInstanceGroupRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *CreateAndroidInstanceGroupRequest
	GetBizRegionId() *string
	SetChargeType(v string) *CreateAndroidInstanceGroupRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateAndroidInstanceGroupRequest
	GetClientToken() *string
	SetEnableIpv6(v bool) *CreateAndroidInstanceGroupRequest
	GetEnableIpv6() *bool
	SetGpuAcceleration(v bool) *CreateAndroidInstanceGroupRequest
	GetGpuAcceleration() *bool
	SetImageId(v string) *CreateAndroidInstanceGroupRequest
	GetImageId() *string
	SetInstanceGroupName(v string) *CreateAndroidInstanceGroupRequest
	GetInstanceGroupName() *string
	SetInstanceGroupSpec(v string) *CreateAndroidInstanceGroupRequest
	GetInstanceGroupSpec() *string
	SetIpv6Bandwidth(v int32) *CreateAndroidInstanceGroupRequest
	GetIpv6Bandwidth() *int32
	SetKeyPairId(v string) *CreateAndroidInstanceGroupRequest
	GetKeyPairId() *string
	SetNumberOfInstances(v int32) *CreateAndroidInstanceGroupRequest
	GetNumberOfInstances() *int32
	SetOfficeSiteId(v string) *CreateAndroidInstanceGroupRequest
	GetOfficeSiteId() *string
	SetPeriod(v int32) *CreateAndroidInstanceGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateAndroidInstanceGroupRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateAndroidInstanceGroupRequest
	GetPolicyGroupId() *string
	SetTag(v []*CreateAndroidInstanceGroupRequestTag) *CreateAndroidInstanceGroupRequest
	GetTag() []*CreateAndroidInstanceGroupRequestTag
	SetVSwitchId(v string) *CreateAndroidInstanceGroupRequest
	GetVSwitchId() *string
}

type CreateAndroidInstanceGroupRequest struct {
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
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
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
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
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
	// The tags
	Tag []*CreateAndroidInstanceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s CreateAndroidInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateAndroidInstanceGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateAndroidInstanceGroupRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAndroidInstanceGroupRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateAndroidInstanceGroupRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateAndroidInstanceGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAndroidInstanceGroupRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *CreateAndroidInstanceGroupRequest) GetGpuAcceleration() *bool {
	return s.GpuAcceleration
}

func (s *CreateAndroidInstanceGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAndroidInstanceGroupRequest) GetInstanceGroupName() *string {
	return s.InstanceGroupName
}

func (s *CreateAndroidInstanceGroupRequest) GetInstanceGroupSpec() *string {
	return s.InstanceGroupSpec
}

func (s *CreateAndroidInstanceGroupRequest) GetIpv6Bandwidth() *int32 {
	return s.Ipv6Bandwidth
}

func (s *CreateAndroidInstanceGroupRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *CreateAndroidInstanceGroupRequest) GetNumberOfInstances() *int32 {
	return s.NumberOfInstances
}

func (s *CreateAndroidInstanceGroupRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateAndroidInstanceGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAndroidInstanceGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateAndroidInstanceGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateAndroidInstanceGroupRequest) GetTag() []*CreateAndroidInstanceGroupRequestTag {
	return s.Tag
}

func (s *CreateAndroidInstanceGroupRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAndroidInstanceGroupRequest) SetAmount(v int32) *CreateAndroidInstanceGroupRequest {
	s.Amount = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetAutoPay(v bool) *CreateAndroidInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetAutoRenew(v bool) *CreateAndroidInstanceGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetBizRegionId(v string) *CreateAndroidInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetChargeType(v string) *CreateAndroidInstanceGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetClientToken(v string) *CreateAndroidInstanceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetEnableIpv6(v bool) *CreateAndroidInstanceGroupRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetGpuAcceleration(v bool) *CreateAndroidInstanceGroupRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetImageId(v string) *CreateAndroidInstanceGroupRequest {
	s.ImageId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetInstanceGroupName(v string) *CreateAndroidInstanceGroupRequest {
	s.InstanceGroupName = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetInstanceGroupSpec(v string) *CreateAndroidInstanceGroupRequest {
	s.InstanceGroupSpec = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetIpv6Bandwidth(v int32) *CreateAndroidInstanceGroupRequest {
	s.Ipv6Bandwidth = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetKeyPairId(v string) *CreateAndroidInstanceGroupRequest {
	s.KeyPairId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetNumberOfInstances(v int32) *CreateAndroidInstanceGroupRequest {
	s.NumberOfInstances = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetOfficeSiteId(v string) *CreateAndroidInstanceGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetPeriod(v int32) *CreateAndroidInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetPeriodUnit(v string) *CreateAndroidInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetPolicyGroupId(v string) *CreateAndroidInstanceGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetTag(v []*CreateAndroidInstanceGroupRequestTag) *CreateAndroidInstanceGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetVSwitchId(v string) *CreateAndroidInstanceGroupRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAndroidInstanceGroupRequestTag struct {
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

func (s CreateAndroidInstanceGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAndroidInstanceGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAndroidInstanceGroupRequestTag) SetKey(v string) *CreateAndroidInstanceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestTag) SetValue(v string) *CreateAndroidInstanceGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
