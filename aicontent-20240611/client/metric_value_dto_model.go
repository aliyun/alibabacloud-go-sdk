// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricValueDTO interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *MetricValueDTO
	GetKey() *string
	SetLabel(v string) *MetricValueDTO
	GetLabel() *string
	SetUnit(v string) *MetricValueDTO
	GetUnit() *string
	SetValue(v float32) *MetricValueDTO
	GetValue() *float32
}

type MetricValueDTO struct {
	// example:
	//
	// total_calls
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 调用次数
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// 次
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 100
	Value *float32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MetricValueDTO) String() string {
	return dara.Prettify(s)
}

func (s MetricValueDTO) GoString() string {
	return s.String()
}

func (s *MetricValueDTO) GetKey() *string {
	return s.Key
}

func (s *MetricValueDTO) GetLabel() *string {
	return s.Label
}

func (s *MetricValueDTO) GetUnit() *string {
	return s.Unit
}

func (s *MetricValueDTO) GetValue() *float32 {
	return s.Value
}

func (s *MetricValueDTO) SetKey(v string) *MetricValueDTO {
	s.Key = &v
	return s
}

func (s *MetricValueDTO) SetLabel(v string) *MetricValueDTO {
	s.Label = &v
	return s
}

func (s *MetricValueDTO) SetUnit(v string) *MetricValueDTO {
	s.Unit = &v
	return s
}

func (s *MetricValueDTO) SetValue(v float32) *MetricValueDTO {
	s.Value = &v
	return s
}

func (s *MetricValueDTO) Validate() error {
	return dara.Validate(s)
}
