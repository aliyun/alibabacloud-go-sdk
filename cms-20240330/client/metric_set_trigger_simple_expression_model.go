// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricSetTriggerSimpleExpression interface {
	dara.Model
	String() string
	GoString() string
	SetExpressionType(v string) *MetricSetTriggerSimpleExpression
	GetExpressionType() *string
	SetMax(v float64) *MetricSetTriggerSimpleExpression
	GetMax() *float64
	SetMin(v float64) *MetricSetTriggerSimpleExpression
	GetMin() *float64
	SetOperator(v string) *MetricSetTriggerSimpleExpression
	GetOperator() *string
	SetQueryName(v string) *MetricSetTriggerSimpleExpression
	GetQueryName() *string
	SetThreshold(v float64) *MetricSetTriggerSimpleExpression
	GetThreshold() *float64
}

type MetricSetTriggerSimpleExpression struct {
	ExpressionType *string  `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	Max            *float64 `json:"max,omitempty" xml:"max,omitempty"`
	Min            *float64 `json:"min,omitempty" xml:"min,omitempty"`
	Operator       *string  `json:"operator,omitempty" xml:"operator,omitempty"`
	QueryName      *string  `json:"queryName,omitempty" xml:"queryName,omitempty"`
	Threshold      *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s MetricSetTriggerSimpleExpression) String() string {
	return dara.Prettify(s)
}

func (s MetricSetTriggerSimpleExpression) GoString() string {
	return s.String()
}

func (s *MetricSetTriggerSimpleExpression) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *MetricSetTriggerSimpleExpression) GetMax() *float64 {
	return s.Max
}

func (s *MetricSetTriggerSimpleExpression) GetMin() *float64 {
	return s.Min
}

func (s *MetricSetTriggerSimpleExpression) GetOperator() *string {
	return s.Operator
}

func (s *MetricSetTriggerSimpleExpression) GetQueryName() *string {
	return s.QueryName
}

func (s *MetricSetTriggerSimpleExpression) GetThreshold() *float64 {
	return s.Threshold
}

func (s *MetricSetTriggerSimpleExpression) SetExpressionType(v string) *MetricSetTriggerSimpleExpression {
	s.ExpressionType = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) SetMax(v float64) *MetricSetTriggerSimpleExpression {
	s.Max = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) SetMin(v float64) *MetricSetTriggerSimpleExpression {
	s.Min = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) SetOperator(v string) *MetricSetTriggerSimpleExpression {
	s.Operator = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) SetQueryName(v string) *MetricSetTriggerSimpleExpression {
	s.QueryName = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) SetThreshold(v float64) *MetricSetTriggerSimpleExpression {
	s.Threshold = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) Validate() error {
	return dara.Validate(s)
}
