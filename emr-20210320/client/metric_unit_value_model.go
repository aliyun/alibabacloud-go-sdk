// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricUnitValue interface {
	dara.Model
	String() string
	GoString() string
	SetMetricName(v string) *MetricUnitValue
	GetMetricName() *string
	SetMetricUnit(v string) *MetricUnitValue
	GetMetricUnit() *string
}

type MetricUnitValue struct {
	// 指标名称。
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// 指标单位。
	MetricUnit *string `json:"MetricUnit,omitempty" xml:"MetricUnit,omitempty"`
}

func (s MetricUnitValue) String() string {
	return dara.Prettify(s)
}

func (s MetricUnitValue) GoString() string {
	return s.String()
}

func (s *MetricUnitValue) GetMetricName() *string {
	return s.MetricName
}

func (s *MetricUnitValue) GetMetricUnit() *string {
	return s.MetricUnit
}

func (s *MetricUnitValue) SetMetricName(v string) *MetricUnitValue {
	s.MetricName = &v
	return s
}

func (s *MetricUnitValue) SetMetricUnit(v string) *MetricUnitValue {
	s.MetricUnit = &v
	return s
}

func (s *MetricUnitValue) Validate() error {
	return dara.Validate(s)
}
