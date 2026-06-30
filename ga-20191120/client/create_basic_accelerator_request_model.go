// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateBasicAcceleratorRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateBasicAcceleratorRequest
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *CreateBasicAcceleratorRequest
	GetAutoRenewDuration() *int32
	SetAutoUseCoupon(v string) *CreateBasicAcceleratorRequest
	GetAutoUseCoupon() *string
	SetBandwidthBillingType(v string) *CreateBasicAcceleratorRequest
	GetBandwidthBillingType() *string
	SetChargeType(v string) *CreateBasicAcceleratorRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateBasicAcceleratorRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateBasicAcceleratorRequest
	GetDryRun() *bool
	SetDuration(v int32) *CreateBasicAcceleratorRequest
	GetDuration() *int32
	SetPricingCycle(v string) *CreateBasicAcceleratorRequest
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *CreateBasicAcceleratorRequest
	GetPromotionOptionNo() *string
	SetRegionId(v string) *CreateBasicAcceleratorRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateBasicAcceleratorRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateBasicAcceleratorRequestTag) *CreateBasicAcceleratorRequest
	GetTag() []*CreateBasicAcceleratorRequestTag
}

type CreateBasicAcceleratorRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **false*	- (default): disables automatic payment. After an order is generated, go to the Order Center to complete the payment.
	//
	// - **true**: enables automatic payment. The order is automatically paid.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// - **true**: enables auto-renewal.
	//
	// - **false*	- (default): disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: months.
	//
	// Valid values: **1*	- to **12**. Default value: **1**.
	//
	// > This parameter takes effect only when **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to use coupons for automatic payment of the bill. Valid values:
	//
	// - **true**: uses coupons.
	//
	// - **false*	- (default): does not use coupons.
	//
	// > This parameter takes effect only when **AutoPay*	- is set to **true**.
	//
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The bandwidth billing method. Valid values:
	//
	// - **BandwidthPackage**: billed by bandwidth plan.
	//
	// - **CDT**: billed by traffic and settled through unified settlement by Cloud Data Transfer (CDT).
	//
	// - **CDT95**: billed by the 95th percentile bandwidth and settled through unified settlement by CDT. This bandwidth billing method is available only to users in the whitelist.
	//
	// example:
	//
	// CDT
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	// The billing method. Valid values:
	//
	// - **PREPAY (default)**: subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the resource. The system checks the required parameters, request syntax, and business limitations. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check succeeds, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The subscription duration.
	//
	// - If **PricingCycle*	- is set to **Month**, valid values of **Duration*	- are **1*	- to **9**.
	//
	// - If **PricingCycle*	- is set to **Year**, valid values of **Duration*	- are **1*	- to **3**.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The billing cycle. Valid values:
	//
	// - **Month**: billed on a monthly basis.
	//
	// - **Year**: billed on a yearly basis.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The coupon number.
	//
	// > This parameter is applicable only to the China site (aliyun.com).
	//
	// example:
	//
	// 50003298014****
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The region ID of the basic Alibaba Cloud Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the basic Alibaba Cloud Global Accelerator (GA) instance belongs.
	//
	// example:
	//
	// rg-acfmxshhcsn****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The labels of the basic Alibaba Cloud Global Accelerator (GA) instance.
	Tag []*CreateBasicAcceleratorRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateBasicAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicAcceleratorRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateBasicAcceleratorRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateBasicAcceleratorRequest) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *CreateBasicAcceleratorRequest) GetAutoUseCoupon() *string {
	return s.AutoUseCoupon
}

func (s *CreateBasicAcceleratorRequest) GetBandwidthBillingType() *string {
	return s.BandwidthBillingType
}

func (s *CreateBasicAcceleratorRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateBasicAcceleratorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBasicAcceleratorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateBasicAcceleratorRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateBasicAcceleratorRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateBasicAcceleratorRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *CreateBasicAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBasicAcceleratorRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateBasicAcceleratorRequest) GetTag() []*CreateBasicAcceleratorRequestTag {
	return s.Tag
}

func (s *CreateBasicAcceleratorRequest) SetAutoPay(v bool) *CreateBasicAcceleratorRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetAutoRenew(v bool) *CreateBasicAcceleratorRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetAutoRenewDuration(v int32) *CreateBasicAcceleratorRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetAutoUseCoupon(v string) *CreateBasicAcceleratorRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetBandwidthBillingType(v string) *CreateBasicAcceleratorRequest {
	s.BandwidthBillingType = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetChargeType(v string) *CreateBasicAcceleratorRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetClientToken(v string) *CreateBasicAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetDryRun(v bool) *CreateBasicAcceleratorRequest {
	s.DryRun = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetDuration(v int32) *CreateBasicAcceleratorRequest {
	s.Duration = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetPricingCycle(v string) *CreateBasicAcceleratorRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetPromotionOptionNo(v string) *CreateBasicAcceleratorRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetRegionId(v string) *CreateBasicAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetResourceGroupId(v string) *CreateBasicAcceleratorRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetTag(v []*CreateBasicAcceleratorRequestTag) *CreateBasicAcceleratorRequest {
	s.Tag = v
	return s
}

func (s *CreateBasicAcceleratorRequest) Validate() error {
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

type CreateBasicAcceleratorRequestTag struct {
	// The label key of the basic Alibaba Cloud Global Accelerator (GA) instance. If you specify this parameter, the value cannot be an empty string.
	//
	// The label key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 label keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The label value of the basic Alibaba Cloud Global Accelerator (GA) instance. If you specify this parameter, the value cannot be an empty string.
	//
	// The label value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 label values.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateBasicAcceleratorRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAcceleratorRequestTag) GoString() string {
	return s.String()
}

func (s *CreateBasicAcceleratorRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateBasicAcceleratorRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateBasicAcceleratorRequestTag) SetKey(v string) *CreateBasicAcceleratorRequestTag {
	s.Key = &v
	return s
}

func (s *CreateBasicAcceleratorRequestTag) SetValue(v string) *CreateBasicAcceleratorRequestTag {
	s.Value = &v
	return s
}

func (s *CreateBasicAcceleratorRequestTag) Validate() error {
	return dara.Validate(s)
}
