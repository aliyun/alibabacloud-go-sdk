// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iByLoadScalingRule interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *ByLoadScalingRule
	GetComparisonOperator() *string
	SetCoolDownInterval(v int32) *ByLoadScalingRule
	GetCoolDownInterval() *int32
	SetEvaluationCount(v int32) *ByLoadScalingRule
	GetEvaluationCount() *int32
	SetMetricName(v string) *ByLoadScalingRule
	GetMetricName() *string
	SetStatistics(v string) *ByLoadScalingRule
	GetStatistics() *string
	SetThreshold(v float64) *ByLoadScalingRule
	GetThreshold() *float64
	SetTimeWindow(v int32) *ByLoadScalingRule
	GetTimeWindow() *int32
}

type ByLoadScalingRule struct {
	// 比较符。
	//
	// This parameter is required.
	//
	// example:
	//
	// LT
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	CoolDownInterval   *int32  `json:"CoolDownInterval,omitempty" xml:"CoolDownInterval,omitempty"`
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

func (s ByLoadScalingRule) String() string {
	return dara.Prettify(s)
}

func (s ByLoadScalingRule) GoString() string {
	return s.String()
}

func (s *ByLoadScalingRule) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ByLoadScalingRule) GetCoolDownInterval() *int32 {
	return s.CoolDownInterval
}

func (s *ByLoadScalingRule) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *ByLoadScalingRule) GetMetricName() *string {
	return s.MetricName
}

func (s *ByLoadScalingRule) GetStatistics() *string {
	return s.Statistics
}

func (s *ByLoadScalingRule) GetThreshold() *float64 {
	return s.Threshold
}

func (s *ByLoadScalingRule) GetTimeWindow() *int32 {
	return s.TimeWindow
}

func (s *ByLoadScalingRule) SetComparisonOperator(v string) *ByLoadScalingRule {
	s.ComparisonOperator = &v
	return s
}

func (s *ByLoadScalingRule) SetCoolDownInterval(v int32) *ByLoadScalingRule {
	s.CoolDownInterval = &v
	return s
}

func (s *ByLoadScalingRule) SetEvaluationCount(v int32) *ByLoadScalingRule {
	s.EvaluationCount = &v
	return s
}

func (s *ByLoadScalingRule) SetMetricName(v string) *ByLoadScalingRule {
	s.MetricName = &v
	return s
}

func (s *ByLoadScalingRule) SetStatistics(v string) *ByLoadScalingRule {
	s.Statistics = &v
	return s
}

func (s *ByLoadScalingRule) SetThreshold(v float64) *ByLoadScalingRule {
	s.Threshold = &v
	return s
}

func (s *ByLoadScalingRule) SetTimeWindow(v int32) *ByLoadScalingRule {
	s.TimeWindow = &v
	return s
}

func (s *ByLoadScalingRule) Validate() error {
	return dara.Validate(s)
}
