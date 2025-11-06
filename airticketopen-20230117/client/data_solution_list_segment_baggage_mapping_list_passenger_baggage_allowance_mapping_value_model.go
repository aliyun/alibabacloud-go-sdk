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
	SetCarryLength(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryLength() *int32
	SetCarryWidth(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryWidth() *int32
	SetCarryHeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarryHeight() *int32
	SetCarrySumOfLengthWidthHeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetCarrySumOfLengthWidthHeight() *int32
	SetLength(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetLength() *int32
	SetWidth(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetWidth() *int32
	SetHeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetHeight() *int32
	SetSumOfLengthWidthHeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue
	GetSumOfLengthWidthHeight() *int32
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

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryLength() *int32 {
	return s.CarryLength
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryWidth() *int32 {
	return s.CarryWidth
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarryHeight() *int32 {
	return s.CarryHeight
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetCarrySumOfLengthWidthHeight() *int32 {
	return s.CarrySumOfLengthWidthHeight
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetLength() *int32 {
	return s.Length
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetWidth() *int32 {
	return s.Width
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetHeight() *int32 {
	return s.Height
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GetSumOfLengthWidthHeight() *int32 {
	return s.SumOfLengthWidthHeight
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

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryLength(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryLength = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryWidth(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryWidth = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryHeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryHeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarrySumOfLengthWidthHeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarrySumOfLengthWidthHeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetLength(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.Length = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetWidth(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.Width = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetHeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.Height = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetSumOfLengthWidthHeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.SumOfLengthWidthHeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) Validate() error {
	return dara.Validate(s)
}
