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
	// The number of instance groups to create. Valid values: 1 to 100. Default value: 1.
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
	// Specifies whether to enable auto-renewal for subscription resources. Default value: false.
	//
	// example:
	//
	// false
	AutoRenew            *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BandwidthPackageId   *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions where Cloud Phone instances are available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// A client-generated token to ensure request idempotence. This parameter prevents duplicate requests. The token can be up to 100 characters in length.
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
	// The image ID. You can call the [DescribeImageList](~~DescribeImageList~~) operation to query available images for Cloud Phone instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-06zyt9m93zwax****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the instance group.
	//
	// > The name can be up to 30 characters in length. It must start with an uppercase or lowercase letter or a Chinese character, and cannot start with `http://` or `https://`. The name can contain only Chinese characters, letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// Cloud phoneA
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The instance group specification. You can call the [DescribeSpec](~~DescribeSpec~~) operation to query the specifications that are available for Cloud Phone instances.
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
	// The key pair ID. If you specify a valid key pair ID when you create the instance group, the system attaches the key pair to all successfully created instances. No separate API call is required to attach the key pair.
	//
	// > Attaching a key pair during a scale-out operation is not supported.
	//
	// example:
	//
	// kp-7o9xywwfutc1l****
	KeyPairId         *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	NetworkInfoShrink *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The number of instances in the instance group. The maximum value is 100.
	//
	// example:
	//
	// 1
	NumberOfInstances *int32 `json:"NumberOfInstances,omitempty" xml:"NumberOfInstances,omitempty"`
	// The network ID.
	//
	// - To create instances in a Shared Network: This parameter is optional. Specify the ID of a **Shared Network**. You can find the ID on the [Cloud Phone console > Network](https://wya.wuying.aliyun.com/network) page. If no Shared Network is available in the console, you can omit this parameter. The system automatically creates a Shared Network when you create the instance group.
	//
	// - To create instances in a VPC: This parameter is required. Specify the ID of a **VPC**. You can find the ID on the [Cloud Phone console > Network](https://wya.wuying.aliyun.com/network) page. If no VPC is available in the console, you must create one first.
	//
	// example:
	//
	// cn-hangzhou+dir-745976****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// wya.wuying.aliyun.com/instanceGroup
	PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// The subscription duration. The PeriodUnit parameter specifies the unit.
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
	// The policy ID. You can call the [ListPolicyGroups](~~ListPolicyGroups~~) operation to query available policies.
	//
	// example:
	//
	// pg-b7bxrrwxkijjh****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	SaleMode      *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	StreamMode    *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
	// The resource tags.
	Tag []*CreateAndroidInstanceGroupShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/448774.html) operation to query available vSwitches.
	//
	// - If you create instances in a Shared Network, omit this parameter.
	//
	// - If you create instances in a VPC, this parameter is required. The system creates the instances in the specified vSwitch.
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
