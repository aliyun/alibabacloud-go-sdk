// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue interface {
	dara.Model
	String() string
	GoString() string
	SetBaggageAmount(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetBaggageAmount() *int32
	SetBaggageWeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetBaggageWeight() *int32
	SetBaggageWeightUnit(v string) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetBaggageWeightUnit() *string
	SetIsAllWeight(v bool) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetIsAllWeight() *bool
	SetCarryOnAmount(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryOnAmount() *int32
	SetCarryOnWeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryOnWeight() *int32
	SetCarryOnWeightUnit(v string) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryOnWeightUnit() *string
	SetIsAllCarryOnWeight(v bool) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetIsAllCarryOnWeight() *bool
}

type DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue struct {
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

func (s DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GoString() string {
	return s.String()
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetBaggageAmount() *int32 {
	return s.BaggageAmount
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetBaggageWeight() *int32 {
	return s.BaggageWeight
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetBaggageWeightUnit() *string {
	return s.BaggageWeightUnit
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetIsAllWeight() *bool {
	return s.IsAllWeight
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryOnAmount() *int32 {
	return s.CarryOnAmount
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryOnWeight() *int32 {
	return s.CarryOnWeight
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryOnWeightUnit() *string {
	return s.CarryOnWeightUnit
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetIsAllCarryOnWeight() *bool {
	return s.IsAllCarryOnWeight
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageAmount(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageAmount = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageWeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageWeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageWeightUnit(v string) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageWeightUnit = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetIsAllWeight(v bool) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.IsAllWeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnAmount(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnAmount = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnWeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnWeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnWeightUnit(v string) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnWeightUnit = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetIsAllCarryOnWeight(v bool) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.IsAllCarryOnWeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) Validate() error {
	return dara.Validate(s)
}
