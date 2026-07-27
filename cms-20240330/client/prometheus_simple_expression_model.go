// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrometheusSimpleExpression interface {
	dara.Model
	String() string
	GoString() string
	SetOperator(v string) *PrometheusSimpleExpression
	GetOperator() *string
	SetQueryName(v string) *PrometheusSimpleExpression
	GetQueryName() *string
	SetThreshold(v float64) *PrometheusSimpleExpression
	GetThreshold() *float64
}

type PrometheusSimpleExpression struct {
	Operator  *string  `json:"operator,omitempty" xml:"operator,omitempty"`
	QueryName *string  `json:"queryName,omitempty" xml:"queryName,omitempty"`
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s PrometheusSimpleExpression) String() string {
	return dara.Prettify(s)
}

func (s PrometheusSimpleExpression) GoString() string {
	return s.String()
}

func (s *PrometheusSimpleExpression) GetOperator() *string {
	return s.Operator
}

func (s *PrometheusSimpleExpression) GetQueryName() *string {
	return s.QueryName
}

func (s *PrometheusSimpleExpression) GetThreshold() *float64 {
	return s.Threshold
}

func (s *PrometheusSimpleExpression) SetOperator(v string) *PrometheusSimpleExpression {
	s.Operator = &v
	return s
}

func (s *PrometheusSimpleExpression) SetQueryName(v string) *PrometheusSimpleExpression {
	s.QueryName = &v
	return s
}

func (s *PrometheusSimpleExpression) SetThreshold(v float64) *PrometheusSimpleExpression {
	s.Threshold = &v
	return s
}

func (s *PrometheusSimpleExpression) Validate() error {
	return dara.Validate(s)
}
