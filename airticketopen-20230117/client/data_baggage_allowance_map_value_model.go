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
	SetCarryLength(v int32) *DataBaggageAllowanceMapValue
	GetCarryLength() *int32
	SetCarryWidth(v int32) *DataBaggageAllowanceMapValue
	GetCarryWidth() *int32
	SetCarryHeight(v int32) *DataBaggageAllowanceMapValue
	GetCarryHeight() *int32
	SetCarrySumOfLengthWidthHeight(v int32) *DataBaggageAllowanceMapValue
	GetCarrySumOfLengthWidthHeight() *int32
	SetLength(v int32) *DataBaggageAllowanceMapValue
	GetLength() *int32
	SetWidth(v int32) *DataBaggageAllowanceMapValue
	GetWidth() *int32
	SetHeight(v int32) *DataBaggageAllowanceMapValue
	GetHeight() *int32
	SetSumOfLengthWidthHeight(v int32) *DataBaggageAllowanceMapValue
	GetSumOfLengthWidthHeight() *int32
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

func (s *DataBaggageAllowanceMapValue) GetCarryLength() *int32 {
	return s.CarryLength
}

func (s *DataBaggageAllowanceMapValue) GetCarryWidth() *int32 {
	return s.CarryWidth
}

func (s *DataBaggageAllowanceMapValue) GetCarryHeight() *int32 {
	return s.CarryHeight
}

func (s *DataBaggageAllowanceMapValue) GetCarrySumOfLengthWidthHeight() *int32 {
	return s.CarrySumOfLengthWidthHeight
}

func (s *DataBaggageAllowanceMapValue) GetLength() *int32 {
	return s.Length
}

func (s *DataBaggageAllowanceMapValue) GetWidth() *int32 {
	return s.Width
}

func (s *DataBaggageAllowanceMapValue) GetHeight() *int32 {
	return s.Height
}

func (s *DataBaggageAllowanceMapValue) GetSumOfLengthWidthHeight() *int32 {
	return s.SumOfLengthWidthHeight
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

func (s *DataBaggageAllowanceMapValue) SetCarryLength(v int32) *DataBaggageAllowanceMapValue {
	s.CarryLength = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarryWidth(v int32) *DataBaggageAllowanceMapValue {
	s.CarryWidth = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarryHeight(v int32) *DataBaggageAllowanceMapValue {
	s.CarryHeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarrySumOfLengthWidthHeight(v int32) *DataBaggageAllowanceMapValue {
	s.CarrySumOfLengthWidthHeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetLength(v int32) *DataBaggageAllowanceMapValue {
	s.Length = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetWidth(v int32) *DataBaggageAllowanceMapValue {
	s.Width = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetHeight(v int32) *DataBaggageAllowanceMapValue {
	s.Height = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetSumOfLengthWidthHeight(v int32) *DataBaggageAllowanceMapValue {
	s.SumOfLengthWidthHeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) Validate() error {
	return dara.Validate(s)
}
