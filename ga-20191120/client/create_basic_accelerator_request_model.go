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
	// 	- **false:*	- disables automatic payment. If you select this option, you must go to the Order Center to complete the payment after an order is generated. This is the default value.
	//
	// 	- **true:*	- enables automatic payment. Payments are automatically completed.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the basic GA instance. Valid values:
	//
	// 	- **true:*	- enables auto-renewal for the basic GA instance.
	//
	// 	- **false:*	- disables auto-renewal for the basic GA instance. This is the default value.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: months.
	//
	// Valid values: **1*	- to **12**. Default value: **1**.
	//
	// >  This parameter takes effect only when the **AutoPay*	- parameter is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to automatically apply coupons to your bills. Valid values:
	//
	// 	- **true:*	- automatically applies coupons to your bills.
	//
	// 	- **false:*	- does not automatically apply coupons to your bills. This is the default value.
	//
	// >  This parameter takes effect only when the **AutoPay*	- parameter is set to **true**.
	//
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The bandwidth billing method. Valid values:
	//
	// 	- **BandwidthPackage:*	- billed based on bandwidth plans.
	//
	// 	- **CDT:*	- billed based on data transfer. The bills are managed by using Cloud Data Transfer (CDT).
	//
	// 	- **CDT95:*	- billed based on the 95th percentile bandwidth. The bills are managed by using Cloud Data Transfer (CDT). This bandwidth billing method is not available by default. Contact your Alibaba Cloud account manager for more information.
	//
	// example:
	//
	// CDT
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PREPAY*	- (default)
	//
	// 	- **POSTPAY**
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
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true:*	- performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The subscription duration of the GA instance.
	//
	// 	- If you set **PricingCycle*	- to **Month**, the valid values for **Duration*	- are **1*	- to **9**.
	//
	// 	- If you set **PricingCycle*	- to **Year**, the valid values for **Duration*	- are **1*	- to **3**.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The billing cycle. Valid values:
	//
	// 	- **Month**
	//
	// 	- **Year**
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The code of the coupon.
	//
	// >  This parameter takes effect only for accounts registered on the international site (alibabacloud.com).
	//
	// example:
	//
	// 50003298014****
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The ID of the region where the basic GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the basic GA instance belongs.
	//
	// example:
	//
	// rg-acfmxshhcsn****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the basic GA instance.
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
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// You can specify up to 20 tag values.
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
