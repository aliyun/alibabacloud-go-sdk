// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricsTrigger interface {
	dara.Model
	String() string
	GoString() string
	SetConditionLogicOperator(v string) *MetricsTrigger
	GetConditionLogicOperator() *string
	SetConditions(v []*TriggerCondition) *MetricsTrigger
	GetConditions() []*TriggerCondition
	SetCoolDownInterval(v int32) *MetricsTrigger
	GetCoolDownInterval() *int32
	SetEvaluationCount(v int32) *MetricsTrigger
	GetEvaluationCount() *int32
	SetTimeConstraints(v []*TimeConstraint) *MetricsTrigger
	GetTimeConstraints() []*TimeConstraint
	SetTimeWindow(v int32) *MetricsTrigger
	GetTimeWindow() *int32
}

type MetricsTrigger struct {
	// The logical relationship between multiple metrics. Valid values:
	//
	// 	- And
	//
	// 	- Or (default)
	//
	// example:
	//
	// Or
	ConditionLogicOperator *string `json:"ConditionLogicOperator,omitempty" xml:"ConditionLogicOperator,omitempty"`
	// The trigger conditions for the metric.
	Conditions []*TriggerCondition `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The cooldown interval. Unit: seconds. Valid values: 0 to 10800.
	//
	// example:
	//
	// 300
	CoolDownInterval *int32 `json:"CoolDownInterval,omitempty" xml:"CoolDownInterval,omitempty"`
	// The number of times that the statistics are collected. This parameter is required. Valid values: 1 to 5.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The limits on time.
	TimeConstraints []*TimeConstraint `json:"TimeConstraints,omitempty" xml:"TimeConstraints,omitempty" type:"Repeated"`
	// The time window for statistics. This parameter is required. Unit: seconds. Valid values: 30 to 1800.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	TimeWindow *int32 `json:"TimeWindow,omitempty" xml:"TimeWindow,omitempty"`
}

func (s MetricsTrigger) String() string {
	return dara.Prettify(s)
}

func (s MetricsTrigger) GoString() string {
	return s.String()
}

func (s *MetricsTrigger) GetConditionLogicOperator() *string {
	return s.ConditionLogicOperator
}

func (s *MetricsTrigger) GetConditions() []*TriggerCondition {
	return s.Conditions
}

func (s *MetricsTrigger) GetCoolDownInterval() *int32 {
	return s.CoolDownInterval
}

func (s *MetricsTrigger) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *MetricsTrigger) GetTimeConstraints() []*TimeConstraint {
	return s.TimeConstraints
}

func (s *MetricsTrigger) GetTimeWindow() *int32 {
	return s.TimeWindow
}

func (s *MetricsTrigger) SetConditionLogicOperator(v string) *MetricsTrigger {
	s.ConditionLogicOperator = &v
	return s
}

func (s *MetricsTrigger) SetConditions(v []*TriggerCondition) *MetricsTrigger {
	s.Conditions = v
	return s
}

func (s *MetricsTrigger) SetCoolDownInterval(v int32) *MetricsTrigger {
	s.CoolDownInterval = &v
	return s
}

func (s *MetricsTrigger) SetEvaluationCount(v int32) *MetricsTrigger {
	s.EvaluationCount = &v
	return s
}

func (s *MetricsTrigger) SetTimeConstraints(v []*TimeConstraint) *MetricsTrigger {
	s.TimeConstraints = v
	return s
}

func (s *MetricsTrigger) SetTimeWindow(v int32) *MetricsTrigger {
	s.TimeWindow = &v
	return s
}

func (s *MetricsTrigger) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TimeConstraints != nil {
		for _, item := range s.TimeConstraints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
