// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrometheusMetricParamValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *PrometheusMetricParamValue
	GetName() *string
	SetValue(v string) *PrometheusMetricParamValue
	GetValue() *string
}

type PrometheusMetricParamValue struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PrometheusMetricParamValue) String() string {
	return dara.Prettify(s)
}

func (s PrometheusMetricParamValue) GoString() string {
	return s.String()
}

func (s *PrometheusMetricParamValue) GetName() *string {
	return s.Name
}

func (s *PrometheusMetricParamValue) GetValue() *string {
	return s.Value
}

func (s *PrometheusMetricParamValue) SetName(v string) *PrometheusMetricParamValue {
	s.Name = &v
	return s
}

func (s *PrometheusMetricParamValue) SetValue(v string) *PrometheusMetricParamValue {
	s.Value = &v
	return s
}

func (s *PrometheusMetricParamValue) Validate() error {
	return dara.Validate(s)
}
