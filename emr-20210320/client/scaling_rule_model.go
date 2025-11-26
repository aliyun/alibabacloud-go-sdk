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
	// 伸缩活动类型。取值范围：
	//
	// - SCALE_OUT：扩容。
	//
	// - SCALE_IN：缩容。
	//
	// This parameter is required.
	//
	// example:
	//
	// SCALE_IN
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// 调整值。需要为正数，代表需要扩容或者缩容的实例数量。
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// 按照负载伸缩描述。
	//
	// <p>
	MetricsTrigger *MetricsTrigger `json:"MetricsTrigger,omitempty" xml:"MetricsTrigger,omitempty"`
	// example:
	//
	// 1
	MinAdjustmentValue *int32 `json:"MinAdjustmentValue,omitempty" xml:"MinAdjustmentValue,omitempty"`
	// 规则名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// scale-out-memory
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 按照时间伸缩描述。
	//
	// <p>
	TimeTrigger *TimeTrigger `json:"TimeTrigger,omitempty" xml:"TimeTrigger,omitempty"`
	// 伸缩规则类型。 取值范围：
	//
	// - TIME_TRIGGER: 按时间伸缩。
	//
	// - METRICS_TRIGGER: 按负载伸缩。
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
