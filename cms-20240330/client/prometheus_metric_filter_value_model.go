// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrometheusMetricFilterValue interface {
	dara.Model
	String() string
	GoString() string
	SetDim(v string) *PrometheusMetricFilterValue
	GetDim() *string
	SetOpt(v string) *PrometheusMetricFilterValue
	GetOpt() *string
	SetValue(v string) *PrometheusMetricFilterValue
	GetValue() *string
}

type PrometheusMetricFilterValue struct {
	Dim   *string `json:"dim,omitempty" xml:"dim,omitempty"`
	Opt   *string `json:"opt,omitempty" xml:"opt,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PrometheusMetricFilterValue) String() string {
	return dara.Prettify(s)
}

func (s PrometheusMetricFilterValue) GoString() string {
	return s.String()
}

func (s *PrometheusMetricFilterValue) GetDim() *string {
	return s.Dim
}

func (s *PrometheusMetricFilterValue) GetOpt() *string {
	return s.Opt
}

func (s *PrometheusMetricFilterValue) GetValue() *string {
	return s.Value
}

func (s *PrometheusMetricFilterValue) SetDim(v string) *PrometheusMetricFilterValue {
	s.Dim = &v
	return s
}

func (s *PrometheusMetricFilterValue) SetOpt(v string) *PrometheusMetricFilterValue {
	s.Opt = &v
	return s
}

func (s *PrometheusMetricFilterValue) SetValue(v string) *PrometheusMetricFilterValue {
	s.Value = &v
	return s
}

func (s *PrometheusMetricFilterValue) Validate() error {
	return dara.Validate(s)
}
