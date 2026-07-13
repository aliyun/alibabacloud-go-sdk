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
	SetChannelCookie(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetChannelCookie() *string
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
	SetInstanceVersion(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetInstanceVersion() *string
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
	SetPaidCallBackUrl(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *CreateAndroidInstanceGroupShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetPolicyGroupId() *string
	SetPromotionId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetPromotionId() *string
	SetSaleMode(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetSaleMode() *string
	SetStreamMode(v int32) *CreateAndroidInstanceGroupShrinkRequest
	GetStreamMode() *int32
	SetTag(v []*CreateAndroidInstanceGroupShrinkRequestTag) *CreateAndroidInstanceGroupShrinkRequest
	GetTag() []*CreateAndroidInstanceGroupShrinkRequestTag
	SetVSwitchId(v string) *CreateAndroidInstanceGroupShrinkRequest
	GetVSwitchId() *string
}

type CreateAndroidInstanceGroupShrinkRequest struct {
	// The number of instance groups. Default value: 1. Maximum value: 100.
	//
	// example:
	//
	// 8
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable automatic payment. Default value: false.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Default value: false.
	//
	// example:
	//
	// false
	AutoRenew            *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BandwidthPackageId   *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the list of regions where cloud phone instances can be purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId   *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The billing type.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent repeated submissions. The value cannot exceed 100 characters in length.
	//
	// example:
	//
	// asadbuvwiabdbvchj****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// Specifies whether to enable GPU acceleration.
	//
	// example:
	//
	// false
	GpuAcceleration *bool `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// The image ID. You can call [DescribeImageList](~~DescribeImageList~~) to query the list of cloud phone images.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-06zyt9m93zwax****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance group name.
	//
	// > The instance group name cannot exceed 30 characters in length. It must start with an uppercase letter, lowercase letter, or Chinese character. It cannot start with `http://` or `https://`. It can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// Cloud phoneA
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The instance group specification. You can call [DescribeSpec](~~DescribeSpec~~) to query the specifications available for cloud phone instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// acp.basic.small
	InstanceGroupSpec *string `json:"InstanceGroupSpec,omitempty" xml:"InstanceGroupSpec,omitempty"`
	InstanceVersion   *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	Ipv6Bandwidth *int32 `json:"Ipv6Bandwidth,omitempty" xml:"Ipv6Bandwidth,omitempty"`
	// The key pair ID. If you specify a valid key pair ID when creating an instance group, the key pair is bound to all instances that are successfully created, without the need to call the bindng operation again.
	//
	// > Binding a key pair during scale-out is not supported.
	//
	// example:
	//
	// kp-7o9xywwfutc1l****
	KeyPairId         *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	NetworkInfoShrink *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The number of instances in the instance group. Maximum value: 100.
	//
	// example:
	//
	// 1
	NumberOfInstances *int32 `json:"NumberOfInstances,omitempty" xml:"NumberOfInstances,omitempty"`
	// The network ID.
	//
	// - To create a shared network instance: the network ID is optional. Specify the network ID of the **Shared Network*	- type on the [Cloud Phone console > Network](https://wya.wuying.aliyun.com/network) page. If no shared network exists in the console, you can leave this parameter empty. A shared network is automatically created when the instance group is created.
	//
	// - To create a VPC network instance: the network ID is required. Specify the network ID of the **VPC Network*	- type on the [Cloud Phone console > Network](https://wya.wuying.aliyun.com/network) page. If no VPC network exists in the console, create a network first.
	//
	// example:
	//
	// cn-hangzhou+dir-745976****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// wya.wuying.aliyun.com/instanceGroup
	PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// The subscription duration of the resource. The unit is specified by PeriodUnit.
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
	// The policy ID. You can call [ListPolicyGroups](~~ListPolicyGroups~~) to query the list of policies.
	//
	// example:
	//
	// pg-b7bxrrwxkijjh****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	SaleMode      *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	StreamMode    *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
	// The tags of the resource.
	Tag []*CreateAndroidInstanceGroupShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID. You can call [DescribeVSwitches](https://help.aliyun.com/document_detail/448774.html) to query the list of vSwitches.
	//
	// - To create a shared network instance: leave this parameter empty.
	//
	// - To create a VPC network instance: the vSwitch ID is required. The specified vSwitch is used to create the instance.
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

func (s *CreateAndroidInstanceGroupShrinkRequest) GetChannelCookie() *string {
	return s.ChannelCookie
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

func (s *CreateAndroidInstanceGroupShrinkRequest) GetInstanceVersion() *string {
	return s.InstanceVersion
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

func (s *CreateAndroidInstanceGroupShrinkRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
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

func (s *CreateAndroidInstanceGroupShrinkRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *CreateAndroidInstanceGroupShrinkRequest) GetStreamMode() *int32 {
	return s.StreamMode
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

func (s *CreateAndroidInstanceGroupShrinkRequest) SetChannelCookie(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.ChannelCookie = &v
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

func (s *CreateAndroidInstanceGroupShrinkRequest) SetInstanceVersion(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.InstanceVersion = &v
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

func (s *CreateAndroidInstanceGroupShrinkRequest) SetPaidCallBackUrl(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.PaidCallBackUrl = &v
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

func (s *CreateAndroidInstanceGroupShrinkRequest) SetSaleMode(v string) *CreateAndroidInstanceGroupShrinkRequest {
	s.SaleMode = &v
	return s
}

func (s *CreateAndroidInstanceGroupShrinkRequest) SetStreamMode(v int32) *CreateAndroidInstanceGroupShrinkRequest {
	s.StreamMode = &v
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
