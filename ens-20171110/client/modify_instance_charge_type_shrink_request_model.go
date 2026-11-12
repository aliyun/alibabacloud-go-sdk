// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceChargeTypeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyInstanceChargeTypeShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *ModifyInstanceChargeTypeShrinkRequest
	GetAutoRenew() *bool
	SetBillingCycle(v string) *ModifyInstanceChargeTypeShrinkRequest
	GetBillingCycle() *string
	SetIncludeDataDisks(v bool) *ModifyInstanceChargeTypeShrinkRequest
	GetIncludeDataDisks() *bool
	SetInstanceChargeType(v string) *ModifyInstanceChargeTypeShrinkRequest
	GetInstanceChargeType() *string
	SetInstanceIdsShrink(v string) *ModifyInstanceChargeTypeShrinkRequest
	GetInstanceIdsShrink() *string
	SetPeriod(v string) *ModifyInstanceChargeTypeShrinkRequest
	GetPeriod() *string
	SetPeriodUnit(v string) *ModifyInstanceChargeTypeShrinkRequest
	GetPeriodUnit() *string
}

type ModifyInstanceChargeTypeShrinkRequest struct {
	// Specifies whether to enable automatic payment when the billing method of the instance is changed from pay-as-you-go to subscription. Valid values:
	//
	// - true: enables automatic payment. Make sure that your account balance is sufficient.
	//
	// - false: only generates the order without deducting the fee. This is the default value.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal when the billing method of the instance is changed from pay-as-you-go to subscription. Valid values:
	//
	// - true: enables auto-renewal.
	//
	// - false: disables auto-renewal. This is the default value.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The billing cycle. Valid values: Month, Day, and Hour.
	//
	// example:
	//
	// Month
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// Specifies whether to convert all pay-as-you-go data cloud disks created along with the instance into subscription data cloud disks when the billing method of the instance is changed from pay-as-you-go to subscription. Valid values:
	//
	// - true: yes.
	//
	// - false: no. This is the default value.
	//
	// example:
	//
	// true
	IncludeDataDisks *bool `json:"IncludeDataDisks,omitempty" xml:"IncludeDataDisks,omitempty"`
	// The target billing method of the instance. Valid values:
	//
	// - PrePaid: converts the instance-level pay-as-you-go instance to a subscription instance.
	//
	// - PostPaid: converts the subscription instance to an instance-level pay-as-you-go instance. This is the default value.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The list of instance IDs.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The subscription duration when the billing method of the instance is changed from pay-as-you-go to subscription. This parameter is required in this scenario. Valid values:
	//
	// - If PeriodUnit is set to Day, Period can only be set to 3.
	//
	// - If PeriodUnit is set to Month, Period can be set to 1, 2, 3, 4, 5, 6, 7, 8, 9, or 12.
	//
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration when the billing method of the instance is changed from pay-as-you-go to subscription. This parameter is required in this scenario. Valid values:
	//
	// - Month: month.
	//
	// - Day: day.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s ModifyInstanceChargeTypeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceChargeTypeShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyInstanceChargeTypeShrinkRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *ModifyInstanceChargeTypeShrinkRequest) GetIncludeDataDisks() *bool {
	return s.IncludeDataDisks
}

func (s *ModifyInstanceChargeTypeShrinkRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyInstanceChargeTypeShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ModifyInstanceChargeTypeShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyInstanceChargeTypeShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyInstanceChargeTypeShrinkRequest) SetAutoPay(v bool) *ModifyInstanceChargeTypeShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceChargeTypeShrinkRequest) SetAutoRenew(v bool) *ModifyInstanceChargeTypeShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyInstanceChargeTypeShrinkRequest) SetBillingCycle(v string) *ModifyInstanceChargeTypeShrinkRequest {
	s.BillingCycle = &v
	return s
}

func (s *ModifyInstanceChargeTypeShrinkRequest) SetIncludeDataDisks(v bool) *ModifyInstanceChargeTypeShrinkRequest {
	s.IncludeDataDisks = &v
	return s
}

func (s *ModifyInstanceChargeTypeShrinkRequest) SetInstanceChargeType(v string) *ModifyInstanceChargeTypeShrinkRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyInstanceChargeTypeShrinkRequest) SetInstanceIdsShrink(v string) *ModifyInstanceChargeTypeShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ModifyInstanceChargeTypeShrinkRequest) SetPeriod(v string) *ModifyInstanceChargeTypeShrinkRequest {
	s.Period = &v
	return s
}

func (s *ModifyInstanceChargeTypeShrinkRequest) SetPeriodUnit(v string) *ModifyInstanceChargeTypeShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyInstanceChargeTypeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
