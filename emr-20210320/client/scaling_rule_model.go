// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingRule interface {
	dara.Model
	String() string
	GoString() string
	SetActivityType(v string) *ScalingRule
	GetActivityType() *string
	SetAdjustmentValue(v int32) *ScalingRule
	GetAdjustmentValue() *int32
	SetMetricsTrigger(v *MetricsTrigger) *ScalingRule
	GetMetricsTrigger() *MetricsTrigger
	SetMinAdjustmentValue(v int32) *ScalingRule
	GetMinAdjustmentValue() *int32
	SetRuleName(v string) *ScalingRule
	GetRuleName() *string
	SetTimeTrigger(v *TimeTrigger) *ScalingRule
	GetTimeTrigger() *TimeTrigger
	SetTriggerType(v string) *ScalingRule
	GetTriggerType() *string
}

type ScalingRule struct {
	// The type of the scaling activity. This parameter is required. Valid values:
	//
	// - SCALE_OUT: scale-out.
	//
	// - SCALE_IN: scale-in.
	//
	// This parameter is required.
	//
	// example:
	//
	// SCALE_IN
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// The adjustment value. This parameter is required and must be a positive integer. It specifies the number of instances to add for a scale-out activity or remove for a scale-in activity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The configurations for load-based scaling.
	MetricsTrigger *MetricsTrigger `json:"MetricsTrigger,omitempty" xml:"MetricsTrigger,omitempty"`
	// The minimum number of instances to add during a scale-out activity.
	//
	// example:
	//
	// 1
	MinAdjustmentValue *int32 `json:"MinAdjustmentValue,omitempty" xml:"MinAdjustmentValue,omitempty"`
	// The name of the rule. This parameter is required and cannot be an empty string.
	//
	// This parameter is required.
	//
	// example:
	//
	// scalingByYarnMemory
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The configurations for time-based scaling.
	TimeTrigger *TimeTrigger `json:"TimeTrigger,omitempty" xml:"TimeTrigger,omitempty"`
	// The type of the scaling rule. This parameter is required. Valid values:
	//
	// - TIME_TRIGGER: time-based scaling.
	//
	// - METRICS_TRIGGER: load-based scaling.
	//
	// This parameter is required.
	//
	// example:
	//
	// TIME_TRIGGER
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ScalingRule) String() string {
	return dara.Prettify(s)
}

func (s ScalingRule) GoString() string {
	return s.String()
}

func (s *ScalingRule) GetActivityType() *string {
	return s.ActivityType
}

func (s *ScalingRule) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *ScalingRule) GetMetricsTrigger() *MetricsTrigger {
	return s.MetricsTrigger
}

func (s *ScalingRule) GetMinAdjustmentValue() *int32 {
	return s.MinAdjustmentValue
}

func (s *ScalingRule) GetRuleName() *string {
	return s.RuleName
}

func (s *ScalingRule) GetTimeTrigger() *TimeTrigger {
	return s.TimeTrigger
}

func (s *ScalingRule) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ScalingRule) SetActivityType(v string) *ScalingRule {
	s.ActivityType = &v
	return s
}

func (s *ScalingRule) SetAdjustmentValue(v int32) *ScalingRule {
	s.AdjustmentValue = &v
	return s
}

func (s *ScalingRule) SetMetricsTrigger(v *MetricsTrigger) *ScalingRule {
	s.MetricsTrigger = v
	return s
}

func (s *ScalingRule) SetMinAdjustmentValue(v int32) *ScalingRule {
	s.MinAdjustmentValue = &v
	return s
}

func (s *ScalingRule) SetRuleName(v string) *ScalingRule {
	s.RuleName = &v
	return s
}

func (s *ScalingRule) SetTimeTrigger(v *TimeTrigger) *ScalingRule {
	s.TimeTrigger = v
	return s
}

func (s *ScalingRule) SetTriggerType(v string) *ScalingRule {
	s.TriggerType = &v
	return s
}

func (s *ScalingRule) Validate() error {
	if s.MetricsTrigger != nil {
		if err := s.MetricsTrigger.Validate(); err != nil {
			return err
		}
	}
	if s.TimeTrigger != nil {
		if err := s.TimeTrigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}
