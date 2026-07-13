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
	SetBandwidthPackageId(v string) *CreateAndroidInstanceGroupRequest
	GetBandwidthPackageId() *string
	SetBandwidthPackageType(v string) *CreateAndroidInstanceGroupRequest
	GetBandwidthPackageType() *string
	SetBizRegionId(v string) *CreateAndroidInstanceGroupRequest
	GetBizRegionId() *string
	SetChannelCookie(v string) *CreateAndroidInstanceGroupRequest
	GetChannelCookie() *string
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
	SetInstanceVersion(v string) *CreateAndroidInstanceGroupRequest
	GetInstanceVersion() *string
	SetIpv6Bandwidth(v int32) *CreateAndroidInstanceGroupRequest
	GetIpv6Bandwidth() *int32
	SetKeyPairId(v string) *CreateAndroidInstanceGroupRequest
	GetKeyPairId() *string
	SetNetworkInfo(v *CreateAndroidInstanceGroupRequestNetworkInfo) *CreateAndroidInstanceGroupRequest
	GetNetworkInfo() *CreateAndroidInstanceGroupRequestNetworkInfo
	SetNetworkType(v string) *CreateAndroidInstanceGroupRequest
	GetNetworkType() *string
	SetNumberOfInstances(v int32) *CreateAndroidInstanceGroupRequest
	GetNumberOfInstances() *int32
	SetOfficeSiteId(v string) *CreateAndroidInstanceGroupRequest
	GetOfficeSiteId() *string
	SetPaidCallBackUrl(v string) *CreateAndroidInstanceGroupRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *CreateAndroidInstanceGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateAndroidInstanceGroupRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateAndroidInstanceGroupRequest
	GetPolicyGroupId() *string
	SetPromotionId(v string) *CreateAndroidInstanceGroupRequest
	GetPromotionId() *string
	SetSaleMode(v string) *CreateAndroidInstanceGroupRequest
	GetSaleMode() *string
	SetStreamMode(v int32) *CreateAndroidInstanceGroupRequest
	GetStreamMode() *int32
	SetTag(v []*CreateAndroidInstanceGroupRequestTag) *CreateAndroidInstanceGroupRequest
	GetTag() []*CreateAndroidInstanceGroupRequestTag
	SetVSwitchId(v string) *CreateAndroidInstanceGroupRequest
	GetVSwitchId() *string
}

type CreateAndroidInstanceGroupRequest struct {
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
	KeyPairId   *string                                       `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	NetworkInfo *CreateAndroidInstanceGroupRequestNetworkInfo `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty" type:"Struct"`
	NetworkType *string                                       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
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
	Tag []*CreateAndroidInstanceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateAndroidInstanceGroupRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateAndroidInstanceGroupRequest) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *CreateAndroidInstanceGroupRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateAndroidInstanceGroupRequest) GetChannelCookie() *string {
	return s.ChannelCookie
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

func (s *CreateAndroidInstanceGroupRequest) GetInstanceVersion() *string {
	return s.InstanceVersion
}

func (s *CreateAndroidInstanceGroupRequest) GetIpv6Bandwidth() *int32 {
	return s.Ipv6Bandwidth
}

func (s *CreateAndroidInstanceGroupRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *CreateAndroidInstanceGroupRequest) GetNetworkInfo() *CreateAndroidInstanceGroupRequestNetworkInfo {
	return s.NetworkInfo
}

func (s *CreateAndroidInstanceGroupRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateAndroidInstanceGroupRequest) GetNumberOfInstances() *int32 {
	return s.NumberOfInstances
}

func (s *CreateAndroidInstanceGroupRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateAndroidInstanceGroupRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
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

func (s *CreateAndroidInstanceGroupRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateAndroidInstanceGroupRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *CreateAndroidInstanceGroupRequest) GetStreamMode() *int32 {
	return s.StreamMode
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

func (s *CreateAndroidInstanceGroupRequest) SetBandwidthPackageId(v string) *CreateAndroidInstanceGroupRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetBandwidthPackageType(v string) *CreateAndroidInstanceGroupRequest {
	s.BandwidthPackageType = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetBizRegionId(v string) *CreateAndroidInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetChannelCookie(v string) *CreateAndroidInstanceGroupRequest {
	s.ChannelCookie = &v
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

func (s *CreateAndroidInstanceGroupRequest) SetInstanceVersion(v string) *CreateAndroidInstanceGroupRequest {
	s.InstanceVersion = &v
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

func (s *CreateAndroidInstanceGroupRequest) SetNetworkInfo(v *CreateAndroidInstanceGroupRequestNetworkInfo) *CreateAndroidInstanceGroupRequest {
	s.NetworkInfo = v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetNetworkType(v string) *CreateAndroidInstanceGroupRequest {
	s.NetworkType = &v
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

func (s *CreateAndroidInstanceGroupRequest) SetPaidCallBackUrl(v string) *CreateAndroidInstanceGroupRequest {
	s.PaidCallBackUrl = &v
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

func (s *CreateAndroidInstanceGroupRequest) SetPromotionId(v string) *CreateAndroidInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetSaleMode(v string) *CreateAndroidInstanceGroupRequest {
	s.SaleMode = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetStreamMode(v int32) *CreateAndroidInstanceGroupRequest {
	s.StreamMode = &v
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
	if s.NetworkInfo != nil {
		if err := s.NetworkInfo.Validate(); err != nil {
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
	return nil
}

type CreateAndroidInstanceGroupRequestNetworkInfo struct {
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" xml:"BandwidthPackageName,omitempty"`
	CidrBlock            *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	InternetChargeType   *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpRatio              *int32  `json:"IpRatio,omitempty" xml:"IpRatio,omitempty"`
	Isp                  *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	LimitedBandwidth     *int32  `json:"LimitedBandwidth,omitempty" xml:"LimitedBandwidth,omitempty"`
	PaidCallbackUrl      *string `json:"PaidCallbackUrl,omitempty" xml:"PaidCallbackUrl,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit           *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId          *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	VisibleType          *string `json:"VisibleType,omitempty" xml:"VisibleType,omitempty"`
}

func (s CreateAndroidInstanceGroupRequestNetworkInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupRequestNetworkInfo) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetBandwidthPackageName() *string {
	return s.BandwidthPackageName
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetIpRatio() *int32 {
	return s.IpRatio
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetIsp() *string {
	return s.Isp
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetLimitedBandwidth() *int32 {
	return s.LimitedBandwidth
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetPaidCallbackUrl() *string {
	return s.PaidCallbackUrl
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetPayType() *string {
	return s.PayType
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) GetVisibleType() *string {
	return s.VisibleType
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetAutoPay(v bool) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.AutoPay = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetAutoRenew(v bool) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.AutoRenew = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetBandwidthPackageName(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.BandwidthPackageName = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetCidrBlock(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.CidrBlock = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetInternetChargeType(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.InternetChargeType = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetIpRatio(v int32) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.IpRatio = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetIsp(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.Isp = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetLimitedBandwidth(v int32) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.LimitedBandwidth = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetPaidCallbackUrl(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.PaidCallbackUrl = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetPayType(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.PayType = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetPeriod(v int32) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.Period = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetPeriodUnit(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetPromotionId(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.PromotionId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) SetVisibleType(v string) *CreateAndroidInstanceGroupRequestNetworkInfo {
	s.VisibleType = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestNetworkInfo) Validate() error {
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
