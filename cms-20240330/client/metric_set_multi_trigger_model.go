// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricSetMultiTrigger interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*MetricSetTriggerSimpleExpression) *MetricSetMultiTrigger
	GetConditions() []*MetricSetTriggerSimpleExpression
	SetDurationSecs(v int32) *MetricSetMultiTrigger
	GetDurationSecs() *int32
	SetExpressionType(v string) *MetricSetMultiTrigger
	GetExpressionType() *string
	SetLogicOperator(v string) *MetricSetMultiTrigger
	GetLogicOperator() *string
	SetMax(v float64) *MetricSetMultiTrigger
	GetMax() *float64
	SetMin(v float64) *MetricSetMultiTrigger
	GetMin() *float64
	SetOperator(v string) *MetricSetMultiTrigger
	GetOperator() *string
	SetQueryName(v string) *MetricSetMultiTrigger
	GetQueryName() *string
	SetSeverity(v string) *MetricSetMultiTrigger
	GetSeverity() *string
	SetThreshold(v float64) *MetricSetMultiTrigger
	GetThreshold() *float64
}

type MetricSetMultiTrigger struct {
	// The list of sub-conditions (used when expressionType=COMPOSITE). Each item contains queryName, operator, and threshold.
	Conditions []*MetricSetTriggerSimpleExpression `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The duration in seconds that data must continuously meet the condition to trigger an alert. If not specified, the value is inherited from conditionConfig.durationSecs.
	DurationSecs *int32 `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	// The expression type. Valid values: SIMPLE (single-metric threshold) or COMPOSITE (multi-metric AND/OR/UNLESS combination).
	ExpressionType *string `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	// The logic operator (used when expressionType=COMPOSITE). Valid values: AND (all conditions met), OR (any condition met), UNLESS (first condition met and all others not met).
	LogicOperator *string `json:"logicOperator,omitempty" xml:"logicOperator,omitempty"`
	// The upper bound of the range. Required when expressionType=SIMPLE and operator is IN_RANGE or OUT_OF_RANGE. The value must be greater than or equal to min.
	Max *float64 `json:"max,omitempty" xml:"max,omitempty"`
	// The lower bound of the range. Required when expressionType=SIMPLE and operator is IN_RANGE or OUT_OF_RANGE.
	Min *float64 `json:"min,omitempty" xml:"min,omitempty"`
	// The comparison operator (used when expressionType=SIMPLE). Valid values: GT (greater than), GE (greater than or equal to), LT (less than), LE (less than or equal to), EQ (equal to), NE (not equal to), IN_RANGE (within range, requires min/max), OUT_OF_RANGE (outside range, requires min/max), PRESENT (field exists, no threshold/min/max needed), NOT_PRESENT (field does not exist, no threshold/min/max needed).
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The referenced query name (used when expressionType=SIMPLE), corresponding to QueryConfigUnified.queries[].name.
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// The alert severity level: CRITICAL > ERROR > WARN / WARNING > INFO. Multiple triggers are sorted by this priority, and the first match fires.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The comparison threshold. Used when expressionType=SIMPLE and operator is GT/GE/LT/LE/EQ/NE. For IN_RANGE/OUT_OF_RANGE, use min/max instead. For PRESENT/NOT_PRESENT, leave this field empty.
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s MetricSetMultiTrigger) String() string {
	return dara.Prettify(s)
}

func (s MetricSetMultiTrigger) GoString() string {
	return s.String()
}

func (s *MetricSetMultiTrigger) GetConditions() []*MetricSetTriggerSimpleExpression {
	return s.Conditions
}

func (s *MetricSetMultiTrigger) GetDurationSecs() *int32 {
	return s.DurationSecs
}

func (s *MetricSetMultiTrigger) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *MetricSetMultiTrigger) GetLogicOperator() *string {
	return s.LogicOperator
}

func (s *MetricSetMultiTrigger) GetMax() *float64 {
	return s.Max
}

func (s *MetricSetMultiTrigger) GetMin() *float64 {
	return s.Min
}

func (s *MetricSetMultiTrigger) GetOperator() *string {
	return s.Operator
}

func (s *MetricSetMultiTrigger) GetQueryName() *string {
	return s.QueryName
}

func (s *MetricSetMultiTrigger) GetSeverity() *string {
	return s.Severity
}

func (s *MetricSetMultiTrigger) GetThreshold() *float64 {
	return s.Threshold
}

func (s *MetricSetMultiTrigger) SetConditions(v []*MetricSetTriggerSimpleExpression) *MetricSetMultiTrigger {
	s.Conditions = v
	return s
}

func (s *MetricSetMultiTrigger) SetDurationSecs(v int32) *MetricSetMultiTrigger {
	s.DurationSecs = &v
	return s
}

func (s *MetricSetMultiTrigger) SetExpressionType(v string) *MetricSetMultiTrigger {
	s.ExpressionType = &v
	return s
}

func (s *MetricSetMultiTrigger) SetLogicOperator(v string) *MetricSetMultiTrigger {
	s.LogicOperator = &v
	return s
}

func (s *MetricSetMultiTrigger) SetMax(v float64) *MetricSetMultiTrigger {
	s.Max = &v
	return s
}

func (s *MetricSetMultiTrigger) SetMin(v float64) *MetricSetMultiTrigger {
	s.Min = &v
	return s
}

func (s *MetricSetMultiTrigger) SetOperator(v string) *MetricSetMultiTrigger {
	s.Operator = &v
	return s
}

func (s *MetricSetMultiTrigger) SetQueryName(v string) *MetricSetMultiTrigger {
	s.QueryName = &v
	return s
}

func (s *MetricSetMultiTrigger) SetSeverity(v string) *MetricSetMultiTrigger {
	s.Severity = &v
	return s
}

func (s *MetricSetMultiTrigger) SetThreshold(v float64) *MetricSetMultiTrigger {
	s.Threshold = &v
	return s
}

func (s *MetricSetMultiTrigger) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
