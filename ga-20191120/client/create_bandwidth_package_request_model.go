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
	// 	- **false*	- (default): disables automatic payment. If you select this option, you must go to the Order Center to complete the payment after an order is generated.
	//
	// 	- **true**: enables automatic payment. Payments are automatically completed.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the bandwidth plan. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false*	- (default): does not enable auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: months.
	//
	// Valid values: **1*	- to **12**. Default value: **1**.
	//
	// >  This parameter is required only if **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to automatically pay bills by using coupons. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// >  This parameter is required only if **AutoPay*	- is set to **true**.
	//
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The bandwidth of the bandwidth plan. Unit: Mbit/s.
	//
	// Valid values: **2*	- to **2000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The type of the bandwidth. Valid values:
	//
	// 	- **Basic**: standard bandwidth
	//
	// 	- **Enhanced**: enhanced bandwidth
	//
	// 	- **Advanced**: premium bandwidth
	//
	// If **Type*	- is set to **Basic**, this parameter is required.
	//
	// example:
	//
	// Basic
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The metering method that is used when you use the pay-as-you-go billing method. Valid values:
	//
	// 	- **PayByTraffic*	- (default)
	//
	// 	- **PayBY95*	- By default, the pay-by-95th-percentile metering method is unavailable. If you want to use the metering method, contact your account manager.
	//
	// >  This parameter takes effect only if you set **ChargeType*	- to **POSTPAY**.
	//
	// example:
	//
	// PayByTraffic
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	// Area A to be connected. Set the value to **China-mainland**.
	//
	// You can set this parameter only if you call this operation on the international site (alibabacloud.com).
	//
	// example:
	//
	// China-mainland
	CbnGeographicRegionIdA *string `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	// Area B to be connected. Set the value to **Global**.
	//
	// You can set this parameter only if you call this operation on the international site (alibabacloud.com).
	//
	// example:
	//
	// Global
	CbnGeographicRegionIdB *string `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	// The billing method of the bandwidth plan. Valid values:
	//
	// 	- **PREPAY*	- (default): subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go. By default, the pay-as-you-go billing method is unavailable. If you want to use the billing method, contact your account manager.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subscription duration.
	//
	// 	- If the **PricingCycle*	- parameter is set to **Month**, the valid values for the **Duration*	- parameter are **1*	- to **9**.
	//
	// 	- If the **PricingCycle*	- parameter is set to **Year**, the valid values for the **Duration*	- parameter are **1*	- to **3**.
	//
	// If **ChargeType*	- is set to **PREPAY**, this parameter is required.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The billing cycle. Valid values:
	//
	// 	- **Month**: billed on a monthly basis.
	//
	// 	- **Year**: billed on an annual basis.
	//
	// If **ChargeType*	- is set to **PREPAY**, this parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The coupon code.
	//
	// >  This parameter is only available on the international site (alibabacloud.com).
	//
	// example:
	//
	// 50003298014****
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The percentage of the minimum bandwidth guaranteed if the pay-by-95th-percentile-bandwidth metering method is used. Valid values: **30*	- to **100**.
	//
	// >  This parameter is required only if **BillingType*	- is set to **PayBY95**.
	//
	// example:
	//
	// 30
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The ID of the region where the GA instance is deployed. **cn-hangzhou*	- is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzrnd67gq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags to add to the bandwidth plan.
	Tag []*CreateBandwidthPackageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the bandwidth plan. Valid values:
	//
	// 	- **Basic**: a basic bandwidth plan
	//
	// 	- **CrossDomain**: a cross-region acceleration bandwidth plan
	//
	// If you call this operation on the Alibaba Cloud China site, only **Basic*	- is supported.
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
	// The tag key.
	//
	// The tag keys cannot be an empty string. The tag key can be up to 64 characters in length, and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Each tag key corresponds to a tag value. Valid values of **N**: **1*	- to **20**.
	//
	// The value cannot exceed 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
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
