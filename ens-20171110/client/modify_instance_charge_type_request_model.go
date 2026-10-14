// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyInstanceChargeTypeRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *ModifyInstanceChargeTypeRequest
	GetAutoRenew() *bool
	SetBillingCycle(v string) *ModifyInstanceChargeTypeRequest
	GetBillingCycle() *string
	SetIncludeDataDisks(v bool) *ModifyInstanceChargeTypeRequest
	GetIncludeDataDisks() *bool
	SetInstanceChargeType(v string) *ModifyInstanceChargeTypeRequest
	GetInstanceChargeType() *string
	SetInstanceIds(v []*string) *ModifyInstanceChargeTypeRequest
	GetInstanceIds() []*string
	SetPeriod(v string) *ModifyInstanceChargeTypeRequest
	GetPeriod() *string
	SetPeriodUnit(v string) *ModifyInstanceChargeTypeRequest
	GetPeriodUnit() *string
}

type ModifyInstanceChargeTypeRequest struct {
	// Specifies whether to automatically complete the payment when you switch the billing method of the instance from pay-as-you-go to subscription. Valid values:
	//
	// true: The payment is automatically completed. Make sure that your account balance is sufficient.
	//
	// false: Only an order is generated but no payment is made. This is the default value.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal when converting the instance from pay-as-you-go to subscription. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is not enabled (default).
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
	// Specifies whether to convert all pay-as-you-go data cloud disks created with the instance to subscription data cloud disks when converting an instance-level pay-as-you-go instance to a subscription instance. Valid values:
	//
	// example:
	//
	// true
	IncludeDataDisks *bool `json:"IncludeDataDisks,omitempty" xml:"IncludeDataDisks,omitempty"`
	// The target billing method of the instance. Valid values:
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
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The subscription duration when converting a pay-as-you-go instance to subscription. This parameter is required in this scenario. Valid values:
	//
	// If PeriodUnit is set to Day, Period can only be set to 3.
	//
	// If PeriodUnit is set to Month, Period can be set to 1-9 or 12.
	//
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration when converting an instance-level pay-as-you-go instance to a subscription instance. This parameter is required in this scenario. Valid values:
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s ModifyInstanceChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceChargeTypeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyInstanceChargeTypeRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *ModifyInstanceChargeTypeRequest) GetIncludeDataDisks() *bool {
	return s.IncludeDataDisks
}

func (s *ModifyInstanceChargeTypeRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyInstanceChargeTypeRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyInstanceChargeTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyInstanceChargeTypeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyInstanceChargeTypeRequest) SetAutoPay(v bool) *ModifyInstanceChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetAutoRenew(v bool) *ModifyInstanceChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetBillingCycle(v string) *ModifyInstanceChargeTypeRequest {
	s.BillingCycle = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetIncludeDataDisks(v bool) *ModifyInstanceChargeTypeRequest {
	s.IncludeDataDisks = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetInstanceChargeType(v string) *ModifyInstanceChargeTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetInstanceIds(v []*string) *ModifyInstanceChargeTypeRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriod(v string) *ModifyInstanceChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriodUnit(v string) *ModifyInstanceChargeTypeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
