// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecommendScalingRule interface {
	dara.Model
	String() string
	GoString() string
	SetActivityType(v string) *RecommendScalingRule
	GetActivityType() *string
	SetAdjustmentValue(v int32) *RecommendScalingRule
	GetAdjustmentValue() *int32
	SetInstanceType(v string) *RecommendScalingRule
	GetInstanceType() *string
	SetMaxSave(v float32) *RecommendScalingRule
	GetMaxSave() *float32
	SetMetricsTrigger(v *MetricsTrigger) *RecommendScalingRule
	GetMetricsTrigger() *MetricsTrigger
	SetRuleName(v string) *RecommendScalingRule
	GetRuleName() *string
	SetTimeTrigger(v *TimeTrigger) *RecommendScalingRule
	GetTimeTrigger() *TimeTrigger
	SetTriggerType(v string) *RecommendScalingRule
	GetTriggerType() *string
}

type RecommendScalingRule struct {
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
	// 推荐的规格类型。
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 最大节约成本。
	//
	// example:
	//
	// 0.12
	MaxSave *float32 `json:"MaxSave,omitempty" xml:"MaxSave,omitempty"`
	// 按照负载伸缩描述。
	//
	// <p>
	MetricsTrigger *MetricsTrigger `json:"MetricsTrigger,omitempty" xml:"MetricsTrigger,omitempty"`
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

func (s RecommendScalingRule) String() string {
	return dara.Prettify(s)
}

func (s RecommendScalingRule) GoString() string {
	return s.String()
}

func (s *RecommendScalingRule) GetActivityType() *string {
	return s.ActivityType
}

func (s *RecommendScalingRule) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *RecommendScalingRule) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RecommendScalingRule) GetMaxSave() *float32 {
	return s.MaxSave
}

func (s *RecommendScalingRule) GetMetricsTrigger() *MetricsTrigger {
	return s.MetricsTrigger
}

func (s *RecommendScalingRule) GetRuleName() *string {
	return s.RuleName
}

func (s *RecommendScalingRule) GetTimeTrigger() *TimeTrigger {
	return s.TimeTrigger
}

func (s *RecommendScalingRule) GetTriggerType() *string {
	return s.TriggerType
}

func (s *RecommendScalingRule) SetActivityType(v string) *RecommendScalingRule {
	s.ActivityType = &v
	return s
}

func (s *RecommendScalingRule) SetAdjustmentValue(v int32) *RecommendScalingRule {
	s.AdjustmentValue = &v
	return s
}

func (s *RecommendScalingRule) SetInstanceType(v string) *RecommendScalingRule {
	s.InstanceType = &v
	return s
}

func (s *RecommendScalingRule) SetMaxSave(v float32) *RecommendScalingRule {
	s.MaxSave = &v
	return s
}

func (s *RecommendScalingRule) SetMetricsTrigger(v *MetricsTrigger) *RecommendScalingRule {
	s.MetricsTrigger = v
	return s
}

func (s *RecommendScalingRule) SetRuleName(v string) *RecommendScalingRule {
	s.RuleName = &v
	return s
}

func (s *RecommendScalingRule) SetTimeTrigger(v *TimeTrigger) *RecommendScalingRule {
	s.TimeTrigger = v
	return s
}

func (s *RecommendScalingRule) SetTriggerType(v string) *RecommendScalingRule {
	s.TriggerType = &v
	return s
}

func (s *RecommendScalingRule) Validate() error {
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
