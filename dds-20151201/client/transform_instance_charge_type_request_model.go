// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformInstanceChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *TransformInstanceChargeTypeRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *TransformInstanceChargeTypeRequest
	GetAutoRenew() *string
	SetBusinessInfo(v string) *TransformInstanceChargeTypeRequest
	GetBusinessInfo() *string
	SetChargeType(v string) *TransformInstanceChargeTypeRequest
	GetChargeType() *string
	SetCouponNo(v string) *TransformInstanceChargeTypeRequest
	GetCouponNo() *string
	SetInstanceId(v string) *TransformInstanceChargeTypeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *TransformInstanceChargeTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransformInstanceChargeTypeRequest
	GetOwnerId() *int64
	SetPeriod(v int64) *TransformInstanceChargeTypeRequest
	GetPeriod() *int64
	SetPricingCycle(v string) *TransformInstanceChargeTypeRequest
	GetPricingCycle() *string
	SetResourceOwnerAccount(v string) *TransformInstanceChargeTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransformInstanceChargeTypeRequest
	GetResourceOwnerId() *int64
}

type TransformInstanceChargeTypeRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > Default value: **true**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > Default value: **false**.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The business information. This is an additional parameter.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid:*	- subscription.
	//
	// 	- **PostPaid:*	- pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-2ze55b3ec56c****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the instance. Unit: months. Valid values: **1, 2, 3, 4, 5, 6, 7, 8, 9******, **12**, **24**, and **36**.
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// 实例付费时长单位
	//
	// 取值说明：
	//
	// - **Month：*	- 月
	//
	// -  **Year：*	- 年
	//
	// 默认值：Month
	//
	// example:
	//
	// Month
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TransformInstanceChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *TransformInstanceChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *TransformInstanceChargeTypeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *TransformInstanceChargeTypeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *TransformInstanceChargeTypeRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *TransformInstanceChargeTypeRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *TransformInstanceChargeTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransformInstanceChargeTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransformInstanceChargeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransformInstanceChargeTypeRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *TransformInstanceChargeTypeRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *TransformInstanceChargeTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransformInstanceChargeTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransformInstanceChargeTypeRequest) SetAutoPay(v bool) *TransformInstanceChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetAutoRenew(v string) *TransformInstanceChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetBusinessInfo(v string) *TransformInstanceChargeTypeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetChargeType(v string) *TransformInstanceChargeTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetCouponNo(v string) *TransformInstanceChargeTypeRequest {
	s.CouponNo = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetInstanceId(v string) *TransformInstanceChargeTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetOwnerAccount(v string) *TransformInstanceChargeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetOwnerId(v int64) *TransformInstanceChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetPeriod(v int64) *TransformInstanceChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetPricingCycle(v string) *TransformInstanceChargeTypeRequest {
	s.PricingCycle = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetResourceOwnerAccount(v string) *TransformInstanceChargeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetResourceOwnerId(v int64) *TransformInstanceChargeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
