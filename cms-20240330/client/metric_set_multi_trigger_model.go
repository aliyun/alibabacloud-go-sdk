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
	Conditions     []*MetricSetTriggerSimpleExpression `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	DurationSecs   *int32                              `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	ExpressionType *string                             `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	LogicOperator  *string                             `json:"logicOperator,omitempty" xml:"logicOperator,omitempty"`
	Max            *float64                            `json:"max,omitempty" xml:"max,omitempty"`
	Min            *float64                            `json:"min,omitempty" xml:"min,omitempty"`
	Operator       *string                             `json:"operator,omitempty" xml:"operator,omitempty"`
	QueryName      *string                             `json:"queryName,omitempty" xml:"queryName,omitempty"`
	Severity       *string                             `json:"severity,omitempty" xml:"severity,omitempty"`
	Threshold      *float64                            `json:"threshold,omitempty" xml:"threshold,omitempty"`
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
