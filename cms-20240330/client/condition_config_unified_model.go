// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConditionConfigUnified interface {
	dara.Model
	String() string
	GoString() string
	SetAggregate(v string) *ConditionConfigUnified
	GetAggregate() *string
	SetCompareList(v []*ApmCompositeCompareConfig) *ConditionConfigUnified
	GetCompareList() []*ApmCompositeCompareConfig
	SetDurationSecs(v int32) *ConditionConfigUnified
	GetDurationSecs() *int32
	SetOperator(v string) *ConditionConfigUnified
	GetOperator() *string
	SetRelation(v string) *ConditionConfigUnified
	GetRelation() *string
	SetSeverity(v string) *ConditionConfigUnified
	GetSeverity() *string
	SetThreshold(v float64) *ConditionConfigUnified
	GetThreshold() *float64
	SetThresholdList(v []*ApmThresholdConfig) *ConditionConfigUnified
	GetThresholdList() []*ApmThresholdConfig
	SetType(v string) *ConditionConfigUnified
	GetType() *string
}

type ConditionConfigUnified struct {
	// 聚合函数（APM_SIMPLE_CONDITION）
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// 多条比较（APM_COMPOSITE_CONDITION）
	CompareList []*ApmCompositeCompareConfig `json:"compareList,omitempty" xml:"compareList,omitempty" type:"Repeated"`
	// 持续时间（秒），PROMETHEUS_SIMPLE / UMODEL_METRICSET 使用
	DurationSecs *int32 `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	// 比较操作符（UMODEL_METRICSET_CONDITION 或 APM_SIMPLE_CONDITION）
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 条件间逻辑关系（APM_COMPOSITE_CONDITION）
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// 严重等级（UMODEL / PROMETHEUS_SIMPLE / APM_COMPOSITE）
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// 阈值（UMODEL_METRICSET_CONDITION）
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// 多阈值列表（APM_SIMPLE_CONDITION）
	ThresholdList []*ApmThresholdConfig `json:"thresholdList,omitempty" xml:"thresholdList,omitempty" type:"Repeated"`
	// 检测条件类型
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConditionConfigUnified) String() string {
	return dara.Prettify(s)
}

func (s ConditionConfigUnified) GoString() string {
	return s.String()
}

func (s *ConditionConfigUnified) GetAggregate() *string {
	return s.Aggregate
}

func (s *ConditionConfigUnified) GetCompareList() []*ApmCompositeCompareConfig {
	return s.CompareList
}

func (s *ConditionConfigUnified) GetDurationSecs() *int32 {
	return s.DurationSecs
}

func (s *ConditionConfigUnified) GetOperator() *string {
	return s.Operator
}

func (s *ConditionConfigUnified) GetRelation() *string {
	return s.Relation
}

func (s *ConditionConfigUnified) GetSeverity() *string {
	return s.Severity
}

func (s *ConditionConfigUnified) GetThreshold() *float64 {
	return s.Threshold
}

func (s *ConditionConfigUnified) GetThresholdList() []*ApmThresholdConfig {
	return s.ThresholdList
}

func (s *ConditionConfigUnified) GetType() *string {
	return s.Type
}

func (s *ConditionConfigUnified) SetAggregate(v string) *ConditionConfigUnified {
	s.Aggregate = &v
	return s
}

func (s *ConditionConfigUnified) SetCompareList(v []*ApmCompositeCompareConfig) *ConditionConfigUnified {
	s.CompareList = v
	return s
}

func (s *ConditionConfigUnified) SetDurationSecs(v int32) *ConditionConfigUnified {
	s.DurationSecs = &v
	return s
}

func (s *ConditionConfigUnified) SetOperator(v string) *ConditionConfigUnified {
	s.Operator = &v
	return s
}

func (s *ConditionConfigUnified) SetRelation(v string) *ConditionConfigUnified {
	s.Relation = &v
	return s
}

func (s *ConditionConfigUnified) SetSeverity(v string) *ConditionConfigUnified {
	s.Severity = &v
	return s
}

func (s *ConditionConfigUnified) SetThreshold(v float64) *ConditionConfigUnified {
	s.Threshold = &v
	return s
}

func (s *ConditionConfigUnified) SetThresholdList(v []*ApmThresholdConfig) *ConditionConfigUnified {
	s.ThresholdList = v
	return s
}

func (s *ConditionConfigUnified) SetType(v string) *ConditionConfigUnified {
	s.Type = &v
	return s
}

func (s *ConditionConfigUnified) Validate() error {
	if s.CompareList != nil {
		for _, item := range s.CompareList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ThresholdList != nil {
		for _, item := range s.ThresholdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
