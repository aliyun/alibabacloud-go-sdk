// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareList interface {
	dara.Model
	String() string
	GoString() string
	SetAggregate(v string) *CompareList
	GetAggregate() *string
	SetOperator(v string) *CompareList
	GetOperator() *string
	SetThreshold(v float32) *CompareList
	GetThreshold() *float32
	SetYoyTimeUnit(v string) *CompareList
	GetYoyTimeUnit() *string
	SetYoyTimeValue(v int32) *CompareList
	GetYoyTimeValue() *int32
}

type CompareList struct {
	// This parameter is required.
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// This parameter is required.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// This parameter is required.
	Threshold    *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	YoyTimeUnit  *string  `json:"yoyTimeUnit,omitempty" xml:"yoyTimeUnit,omitempty"`
	YoyTimeValue *int32   `json:"yoyTimeValue,omitempty" xml:"yoyTimeValue,omitempty"`
}

func (s CompareList) String() string {
	return dara.Prettify(s)
}

func (s CompareList) GoString() string {
	return s.String()
}

func (s *CompareList) GetAggregate() *string {
	return s.Aggregate
}

func (s *CompareList) GetOperator() *string {
	return s.Operator
}

func (s *CompareList) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CompareList) GetYoyTimeUnit() *string {
	return s.YoyTimeUnit
}

func (s *CompareList) GetYoyTimeValue() *int32 {
	return s.YoyTimeValue
}

func (s *CompareList) SetAggregate(v string) *CompareList {
	s.Aggregate = &v
	return s
}

func (s *CompareList) SetOperator(v string) *CompareList {
	s.Operator = &v
	return s
}

func (s *CompareList) SetThreshold(v float32) *CompareList {
	s.Threshold = &v
	return s
}

func (s *CompareList) SetYoyTimeUnit(v string) *CompareList {
	s.YoyTimeUnit = &v
	return s
}

func (s *CompareList) SetYoyTimeValue(v int32) *CompareList {
	s.YoyTimeValue = &v
	return s
}

func (s *CompareList) Validate() error {
	return dara.Validate(s)
}
