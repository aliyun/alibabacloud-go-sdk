// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataBaggageAllowanceMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetBaggageAmount(v int32) *DataBaggageAllowanceMapValue
	GetBaggageAmount() *int32
	SetBaggageWeight(v int32) *DataBaggageAllowanceMapValue
	GetBaggageWeight() *int32
	SetBaggageWeightUnit(v string) *DataBaggageAllowanceMapValue
	GetBaggageWeightUnit() *string
	SetIsAllWeight(v bool) *DataBaggageAllowanceMapValue
	GetIsAllWeight() *bool
	SetCarryOnAmount(v int32) *DataBaggageAllowanceMapValue
	GetCarryOnAmount() *int32
	SetCarryOnWeight(v int32) *DataBaggageAllowanceMapValue
	GetCarryOnWeight() *int32
	SetCarryOnWeightUnit(v string) *DataBaggageAllowanceMapValue
	GetCarryOnWeightUnit() *string
	SetIsAllCarryOnWeight(v bool) *DataBaggageAllowanceMapValue
	GetIsAllCarryOnWeight() *bool
}

type DataBaggageAllowanceMapValue struct {
	// checked baggage quantity
	//
	// example:
	//
	// 1
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// checked baggage weight
	//
	// example:
	//
	// 10
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// checked baggage weight unit
	//
	// example:
	//
	// KG
	BaggageWeightUnit *string `json:"baggage_weight_unit,omitempty" xml:"baggage_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllWeight *bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
	// carry-on baggage quantity
	//
	// example:
	//
	// 1
	CarryOnAmount *int32 `json:"carry_on_amount,omitempty" xml:"carry_on_amount,omitempty"`
	// carry-on baggage weight
	//
	// example:
	//
	// 5
	CarryOnWeight *int32 `json:"carry_on_weight,omitempty" xml:"carry_on_weight,omitempty"`
	// carry-on baggage weight unit
	//
	// example:
	//
	// KG
	CarryOnWeightUnit *string `json:"carry_on_weight_unit,omitempty" xml:"carry_on_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllCarryOnWeight *bool `json:"is_all_carry_on_weight,omitempty" xml:"is_all_carry_on_weight,omitempty"`
}

func (s DataBaggageAllowanceMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataBaggageAllowanceMapValue) GoString() string {
	return s.String()
}

func (s *DataBaggageAllowanceMapValue) GetBaggageAmount() *int32 {
	return s.BaggageAmount
}

func (s *DataBaggageAllowanceMapValue) GetBaggageWeight() *int32 {
	return s.BaggageWeight
}

func (s *DataBaggageAllowanceMapValue) GetBaggageWeightUnit() *string {
	return s.BaggageWeightUnit
}

func (s *DataBaggageAllowanceMapValue) GetIsAllWeight() *bool {
	return s.IsAllWeight
}

func (s *DataBaggageAllowanceMapValue) GetCarryOnAmount() *int32 {
	return s.CarryOnAmount
}

func (s *DataBaggageAllowanceMapValue) GetCarryOnWeight() *int32 {
	return s.CarryOnWeight
}

func (s *DataBaggageAllowanceMapValue) GetCarryOnWeightUnit() *string {
	return s.CarryOnWeightUnit
}

func (s *DataBaggageAllowanceMapValue) GetIsAllCarryOnWeight() *bool {
	return s.IsAllCarryOnWeight
}

func (s *DataBaggageAllowanceMapValue) SetBaggageAmount(v int32) *DataBaggageAllowanceMapValue {
	s.BaggageAmount = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetBaggageWeight(v int32) *DataBaggageAllowanceMapValue {
	s.BaggageWeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetBaggageWeightUnit(v string) *DataBaggageAllowanceMapValue {
	s.BaggageWeightUnit = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetIsAllWeight(v bool) *DataBaggageAllowanceMapValue {
	s.IsAllWeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarryOnAmount(v int32) *DataBaggageAllowanceMapValue {
	s.CarryOnAmount = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarryOnWeight(v int32) *DataBaggageAllowanceMapValue {
	s.CarryOnWeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarryOnWeightUnit(v string) *DataBaggageAllowanceMapValue {
	s.CarryOnWeightUnit = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetIsAllCarryOnWeight(v bool) *DataBaggageAllowanceMapValue {
	s.IsAllCarryOnWeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) Validate() error {
	return dara.Validate(s)
}
