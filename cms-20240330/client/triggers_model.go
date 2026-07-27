// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggers interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *Triggers
	GetComparisonOperator() *string
	SetConditions(v []*TriggerConditions) *Triggers
	GetConditions() []*TriggerConditions
	SetCountOperator(v string) *Triggers
	GetCountOperator() *string
	SetCountThreshold(v int64) *Triggers
	GetCountThreshold() *int64
	SetDurationSecs(v int32) *Triggers
	GetDurationSecs() *int32
	SetExpressionType(v string) *Triggers
	GetExpressionType() *string
	SetLogicOperator(v string) *Triggers
	GetLogicOperator() *string
	SetMatchField(v string) *Triggers
	GetMatchField() *string
	SetMatchOperator(v string) *Triggers
	GetMatchOperator() *string
	SetMatchValue(v string) *Triggers
	GetMatchValue() *string
	SetMax(v float64) *Triggers
	GetMax() *float64
	SetMetricName(v string) *Triggers
	GetMetricName() *string
	SetMin(v float64) *Triggers
	GetMin() *float64
	SetOperator(v string) *Triggers
	GetOperator() *string
	SetPeriod(v int32) *Triggers
	GetPeriod() *int32
	SetPreCondition(v string) *Triggers
	GetPreCondition() *string
	SetQueryName(v string) *Triggers
	GetQueryName() *string
	SetSeverity(v string) *Triggers
	GetSeverity() *string
	SetStatistics(v string) *Triggers
	GetStatistics() *string
	SetThreshold(v interface{}) *Triggers
	GetThreshold() interface{}
	SetTimes(v int32) *Triggers
	GetTimes() *int32
}

type Triggers struct {
	ComparisonOperator *string              `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	Conditions         []*TriggerConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	CountOperator      *string              `json:"countOperator,omitempty" xml:"countOperator,omitempty"`
	CountThreshold     *int64               `json:"countThreshold,omitempty" xml:"countThreshold,omitempty"`
	DurationSecs       *int32               `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	ExpressionType     *string              `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	LogicOperator      *string              `json:"logicOperator,omitempty" xml:"logicOperator,omitempty"`
	MatchField         *string              `json:"matchField,omitempty" xml:"matchField,omitempty"`
	MatchOperator      *string              `json:"matchOperator,omitempty" xml:"matchOperator,omitempty"`
	MatchValue         *string              `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
	Max                *float64             `json:"max,omitempty" xml:"max,omitempty"`
	MetricName         *string              `json:"metricName,omitempty" xml:"metricName,omitempty"`
	Min                *float64             `json:"min,omitempty" xml:"min,omitempty"`
	Operator           *string              `json:"operator,omitempty" xml:"operator,omitempty"`
	Period             *int32               `json:"period,omitempty" xml:"period,omitempty"`
	PreCondition       *string              `json:"preCondition,omitempty" xml:"preCondition,omitempty"`
	QueryName          *string              `json:"queryName,omitempty" xml:"queryName,omitempty"`
	Severity           *string              `json:"severity,omitempty" xml:"severity,omitempty"`
	Statistics         *string              `json:"statistics,omitempty" xml:"statistics,omitempty"`
	Threshold          interface{}          `json:"threshold,omitempty" xml:"threshold,omitempty"`
	Times              *int32               `json:"times,omitempty" xml:"times,omitempty"`
}

func (s Triggers) String() string {
	return dara.Prettify(s)
}

func (s Triggers) GoString() string {
	return s.String()
}

func (s *Triggers) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *Triggers) GetConditions() []*TriggerConditions {
	return s.Conditions
}

func (s *Triggers) GetCountOperator() *string {
	return s.CountOperator
}

func (s *Triggers) GetCountThreshold() *int64 {
	return s.CountThreshold
}

func (s *Triggers) GetDurationSecs() *int32 {
	return s.DurationSecs
}

func (s *Triggers) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *Triggers) GetLogicOperator() *string {
	return s.LogicOperator
}

func (s *Triggers) GetMatchField() *string {
	return s.MatchField
}

func (s *Triggers) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *Triggers) GetMatchValue() *string {
	return s.MatchValue
}

func (s *Triggers) GetMax() *float64 {
	return s.Max
}

func (s *Triggers) GetMetricName() *string {
	return s.MetricName
}

func (s *Triggers) GetMin() *float64 {
	return s.Min
}

func (s *Triggers) GetOperator() *string {
	return s.Operator
}

func (s *Triggers) GetPeriod() *int32 {
	return s.Period
}

func (s *Triggers) GetPreCondition() *string {
	return s.PreCondition
}

func (s *Triggers) GetQueryName() *string {
	return s.QueryName
}

func (s *Triggers) GetSeverity() *string {
	return s.Severity
}

func (s *Triggers) GetStatistics() *string {
	return s.Statistics
}

func (s *Triggers) GetThreshold() interface{} {
	return s.Threshold
}

func (s *Triggers) GetTimes() *int32 {
	return s.Times
}

func (s *Triggers) SetComparisonOperator(v string) *Triggers {
	s.ComparisonOperator = &v
	return s
}

func (s *Triggers) SetConditions(v []*TriggerConditions) *Triggers {
	s.Conditions = v
	return s
}

func (s *Triggers) SetCountOperator(v string) *Triggers {
	s.CountOperator = &v
	return s
}

func (s *Triggers) SetCountThreshold(v int64) *Triggers {
	s.CountThreshold = &v
	return s
}

func (s *Triggers) SetDurationSecs(v int32) *Triggers {
	s.DurationSecs = &v
	return s
}

func (s *Triggers) SetExpressionType(v string) *Triggers {
	s.ExpressionType = &v
	return s
}

func (s *Triggers) SetLogicOperator(v string) *Triggers {
	s.LogicOperator = &v
	return s
}

func (s *Triggers) SetMatchField(v string) *Triggers {
	s.MatchField = &v
	return s
}

func (s *Triggers) SetMatchOperator(v string) *Triggers {
	s.MatchOperator = &v
	return s
}

func (s *Triggers) SetMatchValue(v string) *Triggers {
	s.MatchValue = &v
	return s
}

func (s *Triggers) SetMax(v float64) *Triggers {
	s.Max = &v
	return s
}

func (s *Triggers) SetMetricName(v string) *Triggers {
	s.MetricName = &v
	return s
}

func (s *Triggers) SetMin(v float64) *Triggers {
	s.Min = &v
	return s
}

func (s *Triggers) SetOperator(v string) *Triggers {
	s.Operator = &v
	return s
}

func (s *Triggers) SetPeriod(v int32) *Triggers {
	s.Period = &v
	return s
}

func (s *Triggers) SetPreCondition(v string) *Triggers {
	s.PreCondition = &v
	return s
}

func (s *Triggers) SetQueryName(v string) *Triggers {
	s.QueryName = &v
	return s
}

func (s *Triggers) SetSeverity(v string) *Triggers {
	s.Severity = &v
	return s
}

func (s *Triggers) SetStatistics(v string) *Triggers {
	s.Statistics = &v
	return s
}

func (s *Triggers) SetThreshold(v interface{}) *Triggers {
	s.Threshold = v
	return s
}

func (s *Triggers) SetTimes(v int32) *Triggers {
	s.Times = &v
	return s
}

func (s *Triggers) Validate() error {
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
