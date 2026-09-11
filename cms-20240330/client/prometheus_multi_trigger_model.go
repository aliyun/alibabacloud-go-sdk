// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrometheusMultiTrigger interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*PrometheusSimpleExpression) *PrometheusMultiTrigger
	GetConditions() []*PrometheusSimpleExpression
	SetDurationSecs(v int32) *PrometheusMultiTrigger
	GetDurationSecs() *int32
	SetExpressionType(v string) *PrometheusMultiTrigger
	GetExpressionType() *string
	SetLogicOperator(v string) *PrometheusMultiTrigger
	GetLogicOperator() *string
	SetOperator(v string) *PrometheusMultiTrigger
	GetOperator() *string
	SetQueryName(v string) *PrometheusMultiTrigger
	GetQueryName() *string
	SetSeverity(v string) *PrometheusMultiTrigger
	GetSeverity() *string
	SetThreshold(v float64) *PrometheusMultiTrigger
	GetThreshold() *float64
}

type PrometheusMultiTrigger struct {
	// The list of sub-conditions. This parameter is used when expressionType is set to COMPOSITE. Each item contains queryName, operator, and threshold.
	Conditions []*PrometheusSimpleExpression `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The duration, in seconds, for which the data must continuously meet the condition before the alert is triggered. If this parameter is not specified, the value of conditionConfig.durationSecs is inherited.
	//
	// example:
	//
	// 1
	DurationSecs *int32 `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	// The expression type. Valid values:
	//
	// - SIMPLE: single-query threshold.
	//
	// - COMPOSITE: multi-query AND/OR/UNLESS combination.
	//
	// example:
	//
	// SIMPLE
	ExpressionType *string `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	// The logical operator. This parameter is used when expressionType is set to COMPOSITE. Valid values: AND, OR, and UNLESS.
	//
	// example:
	//
	// AND
	LogicOperator *string `json:"logicOperator,omitempty" xml:"logicOperator,omitempty"`
	// The comparison operator. This parameter is used when expressionType is set to SIMPLE. Valid values: GT, GE, LT, LE, EQ, and NE.
	//
	// example:
	//
	// GT
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The referenced query name. This parameter is used when expressionType is set to SIMPLE. The value corresponds to QueryConfigUnified.queries[].name.
	//
	// example:
	//
	// SampleName
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// The alert severity level. Severity levels in descending priority: CRITICAL > ERROR > WARN / WARNING > INFO. When multiple triggers are configured, they are sorted by this priority, and the first match fires the alert.
	//
	// example:
	//
	// INFO
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The comparison threshold. This parameter is used when expressionType is set to SIMPLE.
	//
	// example:
	//
	// 1.0
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s PrometheusMultiTrigger) String() string {
	return dara.Prettify(s)
}

func (s PrometheusMultiTrigger) GoString() string {
	return s.String()
}

func (s *PrometheusMultiTrigger) GetConditions() []*PrometheusSimpleExpression {
	return s.Conditions
}

func (s *PrometheusMultiTrigger) GetDurationSecs() *int32 {
	return s.DurationSecs
}

func (s *PrometheusMultiTrigger) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *PrometheusMultiTrigger) GetLogicOperator() *string {
	return s.LogicOperator
}

func (s *PrometheusMultiTrigger) GetOperator() *string {
	return s.Operator
}

func (s *PrometheusMultiTrigger) GetQueryName() *string {
	return s.QueryName
}

func (s *PrometheusMultiTrigger) GetSeverity() *string {
	return s.Severity
}

func (s *PrometheusMultiTrigger) GetThreshold() *float64 {
	return s.Threshold
}

func (s *PrometheusMultiTrigger) SetConditions(v []*PrometheusSimpleExpression) *PrometheusMultiTrigger {
	s.Conditions = v
	return s
}

func (s *PrometheusMultiTrigger) SetDurationSecs(v int32) *PrometheusMultiTrigger {
	s.DurationSecs = &v
	return s
}

func (s *PrometheusMultiTrigger) SetExpressionType(v string) *PrometheusMultiTrigger {
	s.ExpressionType = &v
	return s
}

func (s *PrometheusMultiTrigger) SetLogicOperator(v string) *PrometheusMultiTrigger {
	s.LogicOperator = &v
	return s
}

func (s *PrometheusMultiTrigger) SetOperator(v string) *PrometheusMultiTrigger {
	s.Operator = &v
	return s
}

func (s *PrometheusMultiTrigger) SetQueryName(v string) *PrometheusMultiTrigger {
	s.QueryName = &v
	return s
}

func (s *PrometheusMultiTrigger) SetSeverity(v string) *PrometheusMultiTrigger {
	s.Severity = &v
	return s
}

func (s *PrometheusMultiTrigger) SetThreshold(v float64) *PrometheusMultiTrigger {
	s.Threshold = &v
	return s
}

func (s *PrometheusMultiTrigger) Validate() error {
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
