// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue interface {
	dara.Model
	String() string
	GoString() string
	SetBaggageAmount(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetBaggageAmount() *int32
	SetBaggageWeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetBaggageWeight() *int32
	SetBaggageWeightUnit(v string) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetBaggageWeightUnit() *string
	SetIsAllWeight(v bool) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetIsAllWeight() *bool
	SetCarryOnAmount(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryOnAmount() *int32
	SetCarryOnWeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryOnWeight() *int32
	SetCarryOnWeightUnit(v string) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryOnWeightUnit() *string
	SetIsAllCarryOnWeight(v bool) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetIsAllCarryOnWeight() *bool
}

type DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue struct {
	// checked baggage quantity
	//
	// example:
	//
	// 22
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// checked baggage weight
	//
	// example:
	//
	// 2
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// checked baggage weight unit
	//
	// example:
	//
	// kg
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
	// 2
	CarryOnWeight *int32 `json:"carry_on_weight,omitempty" xml:"carry_on_weight,omitempty"`
	// carry-on baggage weight unit
	//
	// example:
	//
	// kg
	CarryOnWeightUnit *string `json:"carry_on_weight_unit,omitempty" xml:"carry_on_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllCarryOnWeight *bool `json:"is_all_carry_on_weight,omitempty" xml:"is_all_carry_on_weight,omitempty"`
}

func (s DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GoString() string {
	return s.String()
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetBaggageAmount() *int32 {
	return s.BaggageAmount
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetBaggageWeight() *int32 {
	return s.BaggageWeight
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetBaggageWeightUnit() *string {
	return s.BaggageWeightUnit
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetIsAllWeight() *bool {
	return s.IsAllWeight
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryOnAmount() *int32 {
	return s.CarryOnAmount
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryOnWeight() *int32 {
	return s.CarryOnWeight
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryOnWeightUnit() *string {
	return s.CarryOnWeightUnit
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetIsAllCarryOnWeight() *bool {
	return s.IsAllCarryOnWeight
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageAmount(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageAmount = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageWeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageWeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageWeightUnit(v string) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageWeightUnit = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetIsAllWeight(v bool) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.IsAllWeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnAmount(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnAmount = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnWeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnWeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnWeightUnit(v string) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnWeightUnit = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetIsAllCarryOnWeight(v bool) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.IsAllCarryOnWeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) Validate() error {
	return dara.Validate(s)
}
