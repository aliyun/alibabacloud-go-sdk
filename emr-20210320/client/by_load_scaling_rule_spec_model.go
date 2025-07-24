// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iByLoadScalingRuleSpec interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *ByLoadScalingRuleSpec
	GetComparisonOperator() *string
	SetEvaluationCount(v int32) *ByLoadScalingRuleSpec
	GetEvaluationCount() *int32
	SetMetricName(v string) *ByLoadScalingRuleSpec
	GetMetricName() *string
	SetStatistics(v string) *ByLoadScalingRuleSpec
	GetStatistics() *string
	SetThreshold(v float64) *ByLoadScalingRuleSpec
	GetThreshold() *float64
	SetTimeWindow(v int32) *ByLoadScalingRuleSpec
	GetTimeWindow() *int32
}

type ByLoadScalingRuleSpec struct {
	// 比较符。
	//
	// This parameter is required.
	//
	// example:
	//
	// LT
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// 统计次数。
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// 指标名称。指标名称需要在 ListAutoScalingMetrics 接口返回的指标名称列表中。
	//
	// This parameter is required.
	//
	// example:
	//
	// yarn_resourcemanager_root_availablememoryusage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// 统计量名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// AVG
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// 阈值。
	//
	// This parameter is required.
	//
	// example:
	//
	// 12.5
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// 统计窗口。单位为秒。
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	TimeWindow *int32 `json:"TimeWindow,omitempty" xml:"TimeWindow,omitempty"`
}

func (s ByLoadScalingRuleSpec) String() string {
	return dara.Prettify(s)
}

func (s ByLoadScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ByLoadScalingRuleSpec) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ByLoadScalingRuleSpec) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *ByLoadScalingRuleSpec) GetMetricName() *string {
	return s.MetricName
}

func (s *ByLoadScalingRuleSpec) GetStatistics() *string {
	return s.Statistics
}

func (s *ByLoadScalingRuleSpec) GetThreshold() *float64 {
	return s.Threshold
}

func (s *ByLoadScalingRuleSpec) GetTimeWindow() *int32 {
	return s.TimeWindow
}

func (s *ByLoadScalingRuleSpec) SetComparisonOperator(v string) *ByLoadScalingRuleSpec {
	s.ComparisonOperator = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetEvaluationCount(v int32) *ByLoadScalingRuleSpec {
	s.EvaluationCount = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetMetricName(v string) *ByLoadScalingRuleSpec {
	s.MetricName = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetStatistics(v string) *ByLoadScalingRuleSpec {
	s.Statistics = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetThreshold(v float64) *ByLoadScalingRuleSpec {
	s.Threshold = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetTimeWindow(v int32) *ByLoadScalingRuleSpec {
	s.TimeWindow = &v
	return s
}

func (s *ByLoadScalingRuleSpec) Validate() error {
	return dara.Validate(s)
}
