// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricSetTriggerSimpleExpression interface {
	dara.Model
	String() string
	GoString() string
	SetAbsDeviation(v float64) *MetricSetTriggerSimpleExpression
	GetAbsDeviation() *float64
	SetBaselinePeriod(v string) *MetricSetTriggerSimpleExpression
	GetBaselinePeriod() *string
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
	SetSensitivity(v string) *MetricSetTriggerSimpleExpression
	GetSensitivity() *string
	SetThreshold(v float64) *MetricSetTriggerSimpleExpression
	GetThreshold() *float64
}

type MetricSetTriggerSimpleExpression struct {
	// The minimum deviation or absolute deviation dead zone for the dynamic baseline. Takes effect only for baseline operators. The unit is the same as the metric. The value must be greater than or equal to 0. A value of 0 indicates no restriction.
	//
	// example:
	//
	// 0.0
	AbsDeviation *float64 `json:"absDeviation,omitempty" xml:"absDeviation,omitempty"`
	// The baseline period. Takes effect only for baseline operators. Valid values:
	//
	// - AUTO: Automatically identifies the period.
	//
	// - DAILY: Daily period.
	//
	// - WEEKLY: Weekly period. The backend automatically expands the historical training window to at least 14 days.
	//
	// - NONE: No period.
	//
	// example:
	//
	// AUTO
	BaselinePeriod *string `json:"baselinePeriod,omitempty" xml:"baselinePeriod,omitempty"`
	// The expression type. Fixed value: SIMPLE.
	//
	// example:
	//
	// SIMPLE
	ExpressionType *string `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	// The upper bound of the range. Required when operator is set to IN_RANGE or OUT_OF_RANGE. The value must be greater than or equal to min.
	//
	// example:
	//
	// 100
	Max *float64 `json:"max,omitempty" xml:"max,omitempty"`
	// The lower bound of the range. Required when operator is set to IN_RANGE or OUT_OF_RANGE.
	//
	// example:
	//
	// 0
	Min *float64 `json:"min,omitempty" xml:"min,omitempty"`
	// The comparison operator. Valid values:
	//
	// - GT: Greater than.
	//
	// - GE: Greater than or equal to.
	//
	// - LT: Less than.
	//
	// - LE: Less than or equal to.
	//
	// - EQ: Equal to.
	//
	// - NE: Not equal to.
	//
	// - IN_RANGE: Within the range. Both min and max must be specified.
	//
	// - OUT_OF_RANGE: Outside the range. Both min and max must be specified.
	//
	// - PRESENT: The field exists. The threshold, min, and max parameters are not required.
	//
	// - NOT_PRESENT: The field does not exist. The threshold, min, and max parameters are not required.
	//
	// - ABOVE_UPPER: Dynamic baseline spike. The sensitivity parameter is required. The threshold, min, and max parameters are not required.
	//
	// - BELOW_LOWER: Dynamic baseline drop. The sensitivity parameter is required. The threshold, min, and max parameters are not required.
	//
	// - OUT_OF_BAND: Dynamic baseline bidirectional deviation. The sensitivity parameter is required. The threshold, min, and max parameters are not required.
	//
	// example:
	//
	// OUT_OF_BAND
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The referenced query name, corresponding to QueryConfigUnified.queries[].name.
	//
	// example:
	//
	// cpuQuery
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// The dynamic baseline sensitivity. Takes effect only for baseline operators. Valid values:
	//
	// - HIGH: The narrowest band and highest sensitivity.
	//
	// - MEDIUM: Medium sensitivity.
	//
	// - LOW: The widest band and lowest sensitivity.
	//
	// example:
	//
	// MEDIUM
	Sensitivity *string `json:"sensitivity,omitempty" xml:"sensitivity,omitempty"`
	// The comparison threshold. Used when operator is set to GT, GE, LT, LE, EQ, or NE. For IN_RANGE or OUT_OF_RANGE, use min and max instead. Not required for PRESENT or NOT_PRESENT.
	//
	// example:
	//
	// 80
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s MetricSetTriggerSimpleExpression) String() string {
	return dara.Prettify(s)
}

func (s MetricSetTriggerSimpleExpression) GoString() string {
	return s.String()
}

func (s *MetricSetTriggerSimpleExpression) GetAbsDeviation() *float64 {
	return s.AbsDeviation
}

func (s *MetricSetTriggerSimpleExpression) GetBaselinePeriod() *string {
	return s.BaselinePeriod
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

func (s *MetricSetTriggerSimpleExpression) GetSensitivity() *string {
	return s.Sensitivity
}

func (s *MetricSetTriggerSimpleExpression) GetThreshold() *float64 {
	return s.Threshold
}

func (s *MetricSetTriggerSimpleExpression) SetAbsDeviation(v float64) *MetricSetTriggerSimpleExpression {
	s.AbsDeviation = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) SetBaselinePeriod(v string) *MetricSetTriggerSimpleExpression {
	s.BaselinePeriod = &v
	return s
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

func (s *MetricSetTriggerSimpleExpression) SetSensitivity(v string) *MetricSetTriggerSimpleExpression {
	s.Sensitivity = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) SetThreshold(v float64) *MetricSetTriggerSimpleExpression {
	s.Threshold = &v
	return s
}

func (s *MetricSetTriggerSimpleExpression) Validate() error {
	return dara.Validate(s)
}
