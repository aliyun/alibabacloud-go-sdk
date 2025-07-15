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
	SetChargeType(v string) *ModifyInstanceChargeTypeRequest
	GetChargeType() *string
	SetInstanceGroupIds(v []*string) *ModifyInstanceChargeTypeRequest
	GetInstanceGroupIds() []*string
	SetPeriod(v int32) *ModifyInstanceChargeTypeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyInstanceChargeTypeRequest
	GetPeriodUnit() *string
}

type ModifyInstanceChargeTypeRequest struct {
	// Specifies whether to enable the auto-payment feature. Default value: false.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature. Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The billing method. Valid values:
	//
	// >  Currently, this operation only allows you to change the billing method from **pay-as-you-go to subscription**.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The IDs of the instance groups.
	//
	// This parameter is required.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The subscription duration. The unit is specified by PeriodUnit. Valid values: 1 Month, 2 Months, 3 Months, 6 Months, and 1 Year.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Month**
	//
	// 	- **Year**
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

func (s *ModifyInstanceChargeTypeRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyInstanceChargeTypeRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *ModifyInstanceChargeTypeRequest) GetPeriod() *int32 {
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

func (s *ModifyInstanceChargeTypeRequest) SetChargeType(v string) *ModifyInstanceChargeTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetInstanceGroupIds(v []*string) *ModifyInstanceChargeTypeRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriod(v int32) *ModifyInstanceChargeTypeRequest {
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
