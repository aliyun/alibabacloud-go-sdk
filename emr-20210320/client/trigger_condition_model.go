// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerCondition interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *TriggerCondition
	GetComparisonOperator() *string
	SetMetricName(v string) *TriggerCondition
	GetMetricName() *string
	SetStatistics(v string) *TriggerCondition
	GetStatistics() *string
	SetTags(v []*Tag) *TriggerCondition
	GetTags() []*Tag
	SetThreshold(v float64) *TriggerCondition
	GetThreshold() *float64
}

type TriggerCondition struct {
	// 比较符。取值范围：
	//
	// - EQ:等于。
	//
	// - NE:不等于。
	//
	// - GT:大于。
	//
	// - LT:小于。
	//
	// - GE:大于等于。
	//
	// - LE:小于等于。
	//
	// This parameter is required.
	//
	// example:
	//
	// LT
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// 指标名称。指标名称需要在 ListAutoScalingMetrics 接口返回的指标名称列表中。
	//
	// This parameter is required.
	//
	// example:
	//
	// yarn_resourcemanager_root_availablememoryusage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// 统计量名称。取值范围：
	//
	// - MAX：最大值。
	//
	// - MIN：最小值。
	//
	// - AVG：平均值。
	//
	// This parameter is required.
	//
	// example:
	//
	// AVG
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// 指标Tag。
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 阈值。
	//
	// This parameter is required.
	//
	// example:
	//
	// 12.5
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s TriggerCondition) String() string {
	return dara.Prettify(s)
}

func (s TriggerCondition) GoString() string {
	return s.String()
}

func (s *TriggerCondition) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *TriggerCondition) GetMetricName() *string {
	return s.MetricName
}

func (s *TriggerCondition) GetStatistics() *string {
	return s.Statistics
}

func (s *TriggerCondition) GetTags() []*Tag {
	return s.Tags
}

func (s *TriggerCondition) GetThreshold() *float64 {
	return s.Threshold
}

func (s *TriggerCondition) SetComparisonOperator(v string) *TriggerCondition {
	s.ComparisonOperator = &v
	return s
}

func (s *TriggerCondition) SetMetricName(v string) *TriggerCondition {
	s.MetricName = &v
	return s
}

func (s *TriggerCondition) SetStatistics(v string) *TriggerCondition {
	s.Statistics = &v
	return s
}

func (s *TriggerCondition) SetTags(v []*Tag) *TriggerCondition {
	s.Tags = v
	return s
}

func (s *TriggerCondition) SetThreshold(v float64) *TriggerCondition {
	s.Threshold = &v
	return s
}

func (s *TriggerCondition) Validate() error {
	return dara.Validate(s)
}
