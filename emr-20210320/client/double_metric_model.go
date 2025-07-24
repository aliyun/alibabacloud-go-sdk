// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoubleMetric interface {
	dara.Model
	String() string
	GoString() string
	SetUnit(v string) *DoubleMetric
	GetUnit() *string
	SetValue(v float64) *DoubleMetric
	GetValue() *float64
}

type DoubleMetric struct {
	Unit  *string  `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DoubleMetric) String() string {
	return dara.Prettify(s)
}

func (s DoubleMetric) GoString() string {
	return s.String()
}

func (s *DoubleMetric) GetUnit() *string {
	return s.Unit
}

func (s *DoubleMetric) GetValue() *float64 {
	return s.Value
}

func (s *DoubleMetric) SetUnit(v string) *DoubleMetric {
	s.Unit = &v
	return s
}

func (s *DoubleMetric) SetValue(v float64) *DoubleMetric {
	s.Value = &v
	return s
}

func (s *DoubleMetric) Validate() error {
	return dara.Validate(s)
}
