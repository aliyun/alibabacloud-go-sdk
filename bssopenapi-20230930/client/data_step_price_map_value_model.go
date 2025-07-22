// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataStepPriceMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetRightClose(v bool) *DataStepPriceMapValue
	GetRightClose() *bool
	SetMin(v string) *DataStepPriceMapValue
	GetMin() *string
	SetMax(v string) *DataStepPriceMapValue
	GetMax() *string
	SetCurrency(v string) *DataStepPriceMapValue
	GetCurrency() *string
	SetLeftClose(v bool) *DataStepPriceMapValue
	GetLeftClose() *bool
	SetStepPriceValue(v string) *DataStepPriceMapValue
	GetStepPriceValue() *string
	SetPriceValueType(v string) *DataStepPriceMapValue
	GetPriceValueType() *string
	SetPriceValue(v string) *DataStepPriceMapValue
	GetPriceValue() *string
	SetDeductCycleType(v string) *DataStepPriceMapValue
	GetDeductCycleType() *string
}

type DataStepPriceMapValue struct {
	RightClose      *bool   `json:"RightClose,omitempty" xml:"RightClose,omitempty"`
	Min             *string `json:"Min,omitempty" xml:"Min,omitempty"`
	Max             *string `json:"Max,omitempty" xml:"Max,omitempty"`
	Currency        *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	LeftClose       *bool   `json:"LeftClose,omitempty" xml:"LeftClose,omitempty"`
	StepPriceValue  *string `json:"StepPriceValue,omitempty" xml:"StepPriceValue,omitempty"`
	PriceValueType  *string `json:"PriceValueType,omitempty" xml:"PriceValueType,omitempty"`
	PriceValue      *string `json:"PriceValue,omitempty" xml:"PriceValue,omitempty"`
	DeductCycleType *string `json:"DeductCycleType,omitempty" xml:"DeductCycleType,omitempty"`
}

func (s DataStepPriceMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataStepPriceMapValue) GoString() string {
	return s.String()
}

func (s *DataStepPriceMapValue) GetRightClose() *bool {
	return s.RightClose
}

func (s *DataStepPriceMapValue) GetMin() *string {
	return s.Min
}

func (s *DataStepPriceMapValue) GetMax() *string {
	return s.Max
}

func (s *DataStepPriceMapValue) GetCurrency() *string {
	return s.Currency
}

func (s *DataStepPriceMapValue) GetLeftClose() *bool {
	return s.LeftClose
}

func (s *DataStepPriceMapValue) GetStepPriceValue() *string {
	return s.StepPriceValue
}

func (s *DataStepPriceMapValue) GetPriceValueType() *string {
	return s.PriceValueType
}

func (s *DataStepPriceMapValue) GetPriceValue() *string {
	return s.PriceValue
}

func (s *DataStepPriceMapValue) GetDeductCycleType() *string {
	return s.DeductCycleType
}

func (s *DataStepPriceMapValue) SetRightClose(v bool) *DataStepPriceMapValue {
	s.RightClose = &v
	return s
}

func (s *DataStepPriceMapValue) SetMin(v string) *DataStepPriceMapValue {
	s.Min = &v
	return s
}

func (s *DataStepPriceMapValue) SetMax(v string) *DataStepPriceMapValue {
	s.Max = &v
	return s
}

func (s *DataStepPriceMapValue) SetCurrency(v string) *DataStepPriceMapValue {
	s.Currency = &v
	return s
}

func (s *DataStepPriceMapValue) SetLeftClose(v bool) *DataStepPriceMapValue {
	s.LeftClose = &v
	return s
}

func (s *DataStepPriceMapValue) SetStepPriceValue(v string) *DataStepPriceMapValue {
	s.StepPriceValue = &v
	return s
}

func (s *DataStepPriceMapValue) SetPriceValueType(v string) *DataStepPriceMapValue {
	s.PriceValueType = &v
	return s
}

func (s *DataStepPriceMapValue) SetPriceValue(v string) *DataStepPriceMapValue {
	s.PriceValue = &v
	return s
}

func (s *DataStepPriceMapValue) SetDeductCycleType(v string) *DataStepPriceMapValue {
	s.DeductCycleType = &v
	return s
}

func (s *DataStepPriceMapValue) Validate() error {
	return dara.Validate(s)
}
