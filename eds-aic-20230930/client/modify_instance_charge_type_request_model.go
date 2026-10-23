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
	SetPromotionId(v string) *ModifyInstanceChargeTypeRequest
	GetPromotionId() *string
}

type ModifyInstanceChargeTypeRequest struct {
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
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The billing method of the instance. Currently, only switching from **pay-as-you-go to subscription*	- is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The list of instance group IDs.
	//
	// This parameter is required.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The purchase duration of the resource. The unit is specified by PeriodUnit. Valid values: 1 month, 2 months, 3 months, 6 months, and 1 year.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the purchase duration. Valid values:
	//
	// - **Month**: month.
	//
	// - **Year**: year.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
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

func (s *ModifyInstanceChargeTypeRequest) GetPromotionId() *string {
	return s.PromotionId
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

func (s *ModifyInstanceChargeTypeRequest) SetPromotionId(v string) *ModifyInstanceChargeTypeRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
