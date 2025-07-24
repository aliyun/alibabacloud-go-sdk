// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntegerMetric interface {
	dara.Model
	String() string
	GoString() string
	SetUnit(v string) *IntegerMetric
	GetUnit() *string
	SetValue(v int32) *IntegerMetric
	GetValue() *int32
}

type IntegerMetric struct {
	Unit  *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Value *int32  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s IntegerMetric) String() string {
	return dara.Prettify(s)
}

func (s IntegerMetric) GoString() string {
	return s.String()
}

func (s *IntegerMetric) GetUnit() *string {
	return s.Unit
}

func (s *IntegerMetric) GetValue() *int32 {
	return s.Value
}

func (s *IntegerMetric) SetUnit(v string) *IntegerMetric {
	s.Unit = &v
	return s
}

func (s *IntegerMetric) SetValue(v int32) *IntegerMetric {
	s.Value = &v
	return s
}

func (s *IntegerMetric) Validate() error {
	return dara.Validate(s)
}
