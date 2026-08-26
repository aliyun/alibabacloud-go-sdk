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
	// The expression type. Fixed as SIMPLE.
	ExpressionType *string `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	// The upper bound of the range. Required when operator is IN_RANGE or OUT_OF_RANGE. Must be greater than or equal to min.
	Max *float64 `json:"max,omitempty" xml:"max,omitempty"`
	// The lower bound of the range. Required when operator is IN_RANGE or OUT_OF_RANGE.
	Min *float64 `json:"min,omitempty" xml:"min,omitempty"`
	// The comparison operator. Valid values: GT (greater than), GE (greater than or equal to), LT (less than), LE (less than or equal to), EQ (equal to), NE (not equal to), IN_RANGE (within range, requires both min and max), OUT_OF_RANGE (outside range, requires both min and max), PRESENT (field exists, does not require threshold/min/max), NOT_PRESENT (field does not exist, does not require threshold/min/max).
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The referenced query name, corresponding to QueryConfigUnified.queries[].name.
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// The comparison threshold. Used when operator is GT, GE, LT, LE, EQ, or NE. IN_RANGE and OUT_OF_RANGE use min/max instead. PRESENT and NOT_PRESENT do not require this field.
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
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
