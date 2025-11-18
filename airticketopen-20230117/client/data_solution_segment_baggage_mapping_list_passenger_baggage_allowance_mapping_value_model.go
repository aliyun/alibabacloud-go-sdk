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
	SetCarryLength(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryLength() *int32
	SetCarryWidth(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryWidth() *int32
	SetCarryHeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryHeight() *int32
	SetCarrySumOfLengthWidthHeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarrySumOfLengthWidthHeight() *int32
	SetLength(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetLength() *int32
	SetWidth(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetWidth() *int32
	SetHeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetHeight() *int32
	SetSumOfLengthWidthHeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetSumOfLengthWidthHeight() *int32
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
	// example:
	//
	// 55
	CarryLength *int32 `json:"carry_length,omitempty" xml:"carry_length,omitempty"`
	// example:
	//
	// 40
	CarryWidth *int32 `json:"carry_width,omitempty" xml:"carry_width,omitempty"`
	// example:
	//
	// 20
	CarryHeight *int32 `json:"carry_height,omitempty" xml:"carry_height,omitempty"`
	// example:
	//
	// 115
	CarrySumOfLengthWidthHeight *int32 `json:"carry_sum_of_length_width_height,omitempty" xml:"carry_sum_of_length_width_height,omitempty"`
	// example:
	//
	// 60
	Length *int32 `json:"length,omitempty" xml:"length,omitempty"`
	// example:
	//
	// 40
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
	// example:
	//
	// 60
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// 200
	SumOfLengthWidthHeight *int32 `json:"sum_of_length_width_height,omitempty" xml:"sum_of_length_width_height,omitempty"`
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

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryLength() *int32 {
	return s.CarryLength
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryWidth() *int32 {
	return s.CarryWidth
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryHeight() *int32 {
	return s.CarryHeight
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarrySumOfLengthWidthHeight() *int32 {
	return s.CarrySumOfLengthWidthHeight
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetLength() *int32 {
	return s.Length
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetWidth() *int32 {
	return s.Width
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetHeight() *int32 {
	return s.Height
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetSumOfLengthWidthHeight() *int32 {
	return s.SumOfLengthWidthHeight
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

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryLength(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryLength = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryWidth(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryWidth = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryHeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryHeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarrySumOfLengthWidthHeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarrySumOfLengthWidthHeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetLength(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.Length = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetWidth(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.Width = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetHeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.Height = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetSumOfLengthWidthHeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.SumOfLengthWidthHeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) Validate() error {
	return dara.Validate(s)
}
