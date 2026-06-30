// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateBandwidthPackageRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateBandwidthPackageRequest
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *CreateBandwidthPackageRequest
	GetAutoRenewDuration() *int32
	SetAutoUseCoupon(v string) *CreateBandwidthPackageRequest
	GetAutoUseCoupon() *string
	SetBandwidth(v int32) *CreateBandwidthPackageRequest
	GetBandwidth() *int32
	SetBandwidthType(v string) *CreateBandwidthPackageRequest
	GetBandwidthType() *string
	SetBillingType(v string) *CreateBandwidthPackageRequest
	GetBillingType() *string
	SetCbnGeographicRegionIdA(v string) *CreateBandwidthPackageRequest
	GetCbnGeographicRegionIdA() *string
	SetCbnGeographicRegionIdB(v string) *CreateBandwidthPackageRequest
	GetCbnGeographicRegionIdB() *string
	SetChargeType(v string) *CreateBandwidthPackageRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateBandwidthPackageRequest
	GetClientToken() *string
	SetDuration(v string) *CreateBandwidthPackageRequest
	GetDuration() *string
	SetPricingCycle(v string) *CreateBandwidthPackageRequest
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *CreateBandwidthPackageRequest
	GetPromotionOptionNo() *string
	SetRatio(v int32) *CreateBandwidthPackageRequest
	GetRatio() *int32
	SetRegionId(v string) *CreateBandwidthPackageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateBandwidthPackageRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateBandwidthPackageRequestTag) *CreateBandwidthPackageRequest
	GetTag() []*CreateBandwidthPackageRequestTag
	SetType(v string) *CreateBandwidthPackageRequest
	GetType() *string
}

type CreateBandwidthPackageRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **false*	- (default): Disables automatic payment. After an order is generated, go to the Order Hub to complete the payment.
	//
	// - **true**: Enables automatic payment. Payments are automatically completed.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: months.
	//
	// Valid values: **1*	- to **12**. Default value: **1**.
	//
	// > This parameter takes effect only if **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to use a coupon to automatically pay for the bill. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// > This parameter takes effect only if **AutoPay*	- is set to **true**.
	//
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The bandwidth of the bandwidth plan. Unit: Mbps.
	//
	// Valid values: **2*	- to **2000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The type of bandwidth. Valid values:
	//
	// - **Basic**: Basic bandwidth.
	//
	// - **Enhanced**: Enhanced bandwidth.
	//
	// - **Advanced**: Advanced bandwidth.
	//
	// This parameter is required if you set **Type*	- to **Basic**.
	//
	// example:
	//
	// Basic
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The billing method for a pay-as-you-go bandwidth plan. Valid values:
	//
	// - **PayByTraffic*	- (default): pay-by-traffic.
	//
	// - **PayBY95**: pay-by-95th-percentile. This billing method is not available by default. Contact your account manager to use this billing method.
	//
	// > This parameter takes effect only if **ChargeType*	- is set to **POSTPAY**.
	//
	// example:
	//
	// PayByTraffic
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	// Connected area A of the cross-region acceleration bandwidth plan. Set the value to **China-mainland**.
	//
	// This parameter is available only on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// example:
	//
	// China-mainland
	CbnGeographicRegionIdA *string `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	// Connected area B of the cross-region acceleration bandwidth plan. Set the value to **Global**.
	//
	// This parameter is available only on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// example:
	//
	// Global
	CbnGeographicRegionIdB *string `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	// The billing method. Valid values:
	//
	// - **PREPAY*	- (default): subscription.
	//
	// - **POSTPAY**: pay-as-you-go. The pay-as-you-go billing method is not available by default. Contact your account manager to use this billing method.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a client token from your client to make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 123e4567****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subscription duration.
	//
	// - If you set **PricingCycle*	- to **Month**, valid values for **Duration*	- are **1*	- to **9**.
	//
	// - If you set **PricingCycle*	- to **Year**, valid values for **Duration*	- are **1*	- to **3**.
	//
	// This parameter is required if you set **ChargeType*	- to **PREPAY**.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The billing cycle. Valid values:
	//
	// - **Month**: monthly billing.
	//
	// - **Year**: yearly billing.
	//
	// This parameter is required if you set **ChargeType*	- to **PREPAY**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The coupon code.
	//
	// > This parameter is available only on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// example:
	//
	// 50003298014****
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The percentage of the guaranteed minimum bandwidth if you use the pay-by-95th-percentile metering method. Valid values: **30*	- to **100**.
	//
	// > This parameter takes effect only if **BillingType*	- is set to **PayBY95**.
	//
	// example:
	//
	// 30
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzrnd67gq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the bandwidth plan.
	Tag []*CreateBandwidthPackageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the bandwidth plan. Valid values:
	//
	// - **Basic**: a basic bandwidth plan.
	//
	// - **CrossDomain**: a cross-region acceleration bandwidth plan.
	//
	// Only **Basic*	- is supported on the Alibaba Cloud China Website (www\\.aliyun.com).
	//
	// This parameter is required.
	//
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateBandwidthPackageRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateBandwidthPackageRequest) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *CreateBandwidthPackageRequest) GetAutoUseCoupon() *string {
	return s.AutoUseCoupon
}

func (s *CreateBandwidthPackageRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateBandwidthPackageRequest) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *CreateBandwidthPackageRequest) GetBillingType() *string {
	return s.BillingType
}

func (s *CreateBandwidthPackageRequest) GetCbnGeographicRegionIdA() *string {
	return s.CbnGeographicRegionIdA
}

func (s *CreateBandwidthPackageRequest) GetCbnGeographicRegionIdB() *string {
	return s.CbnGeographicRegionIdB
}

func (s *CreateBandwidthPackageRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateBandwidthPackageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBandwidthPackageRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateBandwidthPackageRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateBandwidthPackageRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *CreateBandwidthPackageRequest) GetRatio() *int32 {
	return s.Ratio
}

func (s *CreateBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBandwidthPackageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateBandwidthPackageRequest) GetTag() []*CreateBandwidthPackageRequestTag {
	return s.Tag
}

func (s *CreateBandwidthPackageRequest) GetType() *string {
	return s.Type
}

func (s *CreateBandwidthPackageRequest) SetAutoPay(v bool) *CreateBandwidthPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetAutoRenew(v bool) *CreateBandwidthPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetAutoRenewDuration(v int32) *CreateBandwidthPackageRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetAutoUseCoupon(v string) *CreateBandwidthPackageRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBandwidth(v int32) *CreateBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBandwidthType(v string) *CreateBandwidthPackageRequest {
	s.BandwidthType = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBillingType(v string) *CreateBandwidthPackageRequest {
	s.BillingType = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetCbnGeographicRegionIdA(v string) *CreateBandwidthPackageRequest {
	s.CbnGeographicRegionIdA = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetCbnGeographicRegionIdB(v string) *CreateBandwidthPackageRequest {
	s.CbnGeographicRegionIdB = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetChargeType(v string) *CreateBandwidthPackageRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetClientToken(v string) *CreateBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetDuration(v string) *CreateBandwidthPackageRequest {
	s.Duration = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetPricingCycle(v string) *CreateBandwidthPackageRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetPromotionOptionNo(v string) *CreateBandwidthPackageRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetRatio(v int32) *CreateBandwidthPackageRequest {
	s.Ratio = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetRegionId(v string) *CreateBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetResourceGroupId(v string) *CreateBandwidthPackageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetTag(v []*CreateBandwidthPackageRequestTag) *CreateBandwidthPackageRequest {
	s.Tag = v
	return s
}

func (s *CreateBandwidthPackageRequest) SetType(v string) *CreateBandwidthPackageRequest {
	s.Type = &v
	return s
}

func (s *CreateBandwidthPackageRequest) Validate() error {
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

type CreateBandwidthPackageRequestTag struct {
	// The tag key of the bandwidth plan. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the bandwidth plan. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag values.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateBandwidthPackageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateBandwidthPackageRequestTag) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateBandwidthPackageRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateBandwidthPackageRequestTag) SetKey(v string) *CreateBandwidthPackageRequestTag {
	s.Key = &v
	return s
}

func (s *CreateBandwidthPackageRequestTag) SetValue(v string) *CreateBandwidthPackageRequestTag {
	s.Value = &v
	return s
}

func (s *CreateBandwidthPackageRequestTag) Validate() error {
	return dara.Validate(s)
}
