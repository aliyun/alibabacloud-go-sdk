// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingRuleSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAdjustmentValue(v int32) *ScalingRuleSpec
	GetAdjustmentValue() *int32
	SetByLoadScalingRuleSpec(v *ScalingRuleSpecByLoadScalingRuleSpec) *ScalingRuleSpec
	GetByLoadScalingRuleSpec() *ScalingRuleSpecByLoadScalingRuleSpec
	SetByTimeScalingRuleSpec(v *ScalingRuleSpecByTimeScalingRuleSpec) *ScalingRuleSpec
	GetByTimeScalingRuleSpec() *ScalingRuleSpecByTimeScalingRuleSpec
	SetCoolDownInterval(v int32) *ScalingRuleSpec
	GetCoolDownInterval() *int32
	SetScalingActivityType(v string) *ScalingRuleSpec
	GetScalingActivityType() *string
	SetScalingRuleName(v string) *ScalingRuleSpec
	GetScalingRuleName() *string
	SetScalingRuleType(v string) *ScalingRuleSpec
	GetScalingRuleType() *string
}

type ScalingRuleSpec struct {
	// 调整值。需要为正数，代表需要扩容或者缩容的实例数量。
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// 按照负载伸缩描述。
	ByLoadScalingRuleSpec *ScalingRuleSpecByLoadScalingRuleSpec `json:"ByLoadScalingRuleSpec,omitempty" xml:"ByLoadScalingRuleSpec,omitempty" type:"Struct"`
	// 按照时间伸缩描述。
	ByTimeScalingRuleSpec *ScalingRuleSpecByTimeScalingRuleSpec `json:"ByTimeScalingRuleSpec,omitempty" xml:"ByTimeScalingRuleSpec,omitempty" type:"Struct"`
	// 冷却时间。单位为秒，取值范围在30~10800秒之间。
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	CoolDownInterval *int32 `json:"CoolDownInterval,omitempty" xml:"CoolDownInterval,omitempty"`
	// 伸缩活动类型。
	//
	// This parameter is required.
	//
	// example:
	//
	// SCALE_IN
	ScalingActivityType *string `json:"ScalingActivityType,omitempty" xml:"ScalingActivityType,omitempty"`
	// 规则名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// scale-out-memory
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// 伸缩规则类型。
	//
	// This parameter is required.
	//
	// example:
	//
	// BY_TIME
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
}

func (s ScalingRuleSpec) String() string {
	return dara.Prettify(s)
}

func (s ScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ScalingRuleSpec) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *ScalingRuleSpec) GetByLoadScalingRuleSpec() *ScalingRuleSpecByLoadScalingRuleSpec {
	return s.ByLoadScalingRuleSpec
}

func (s *ScalingRuleSpec) GetByTimeScalingRuleSpec() *ScalingRuleSpecByTimeScalingRuleSpec {
	return s.ByTimeScalingRuleSpec
}

func (s *ScalingRuleSpec) GetCoolDownInterval() *int32 {
	return s.CoolDownInterval
}

func (s *ScalingRuleSpec) GetScalingActivityType() *string {
	return s.ScalingActivityType
}

func (s *ScalingRuleSpec) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *ScalingRuleSpec) GetScalingRuleType() *string {
	return s.ScalingRuleType
}

func (s *ScalingRuleSpec) SetAdjustmentValue(v int32) *ScalingRuleSpec {
	s.AdjustmentValue = &v
	return s
}

func (s *ScalingRuleSpec) SetByLoadScalingRuleSpec(v *ScalingRuleSpecByLoadScalingRuleSpec) *ScalingRuleSpec {
	s.ByLoadScalingRuleSpec = v
	return s
}

func (s *ScalingRuleSpec) SetByTimeScalingRuleSpec(v *ScalingRuleSpecByTimeScalingRuleSpec) *ScalingRuleSpec {
	s.ByTimeScalingRuleSpec = v
	return s
}

func (s *ScalingRuleSpec) SetCoolDownInterval(v int32) *ScalingRuleSpec {
	s.CoolDownInterval = &v
	return s
}

func (s *ScalingRuleSpec) SetScalingActivityType(v string) *ScalingRuleSpec {
	s.ScalingActivityType = &v
	return s
}

func (s *ScalingRuleSpec) SetScalingRuleName(v string) *ScalingRuleSpec {
	s.ScalingRuleName = &v
	return s
}

func (s *ScalingRuleSpec) SetScalingRuleType(v string) *ScalingRuleSpec {
	s.ScalingRuleType = &v
	return s
}

func (s *ScalingRuleSpec) Validate() error {
	if s.ByLoadScalingRuleSpec != nil {
		if err := s.ByLoadScalingRuleSpec.Validate(); err != nil {
			return err
		}
	}
	if s.ByTimeScalingRuleSpec != nil {
		if err := s.ByTimeScalingRuleSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScalingRuleSpecByLoadScalingRuleSpec struct {
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

func (s ScalingRuleSpecByLoadScalingRuleSpec) String() string {
	return dara.Prettify(s)
}

func (s ScalingRuleSpecByLoadScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) GetMetricName() *string {
	return s.MetricName
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) GetStatistics() *string {
	return s.Statistics
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) GetThreshold() *float64 {
	return s.Threshold
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) GetTimeWindow() *int32 {
	return s.TimeWindow
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetComparisonOperator(v string) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.ComparisonOperator = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetEvaluationCount(v int32) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.EvaluationCount = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetMetricName(v string) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.MetricName = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetStatistics(v string) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.Statistics = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetThreshold(v float64) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.Threshold = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetTimeWindow(v int32) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.TimeWindow = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) Validate() error {
	return dara.Validate(s)
}

type ScalingRuleSpecByTimeScalingRuleSpec struct {
	// 重复执行定时任务的结束时间戳。单位为毫秒。
	//
	// example:
	//
	// 1639714800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 启动时间戳。单位为毫秒。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1639714634819
	LaunchTime *int64 `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// 指定时间规则的执行类型。
	//
	// example:
	//
	// WEEKLY
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// 重复执行定时任务的数值。具体取值取决于 recurrenceType 设置。
	//
	// example:
	//
	// MON,FRI,SUN
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
}

func (s ScalingRuleSpecByTimeScalingRuleSpec) String() string {
	return dara.Prettify(s)
}

func (s ScalingRuleSpecByTimeScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) GetLaunchTime() *int64 {
	return s.LaunchTime
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) SetEndTime(v int64) *ScalingRuleSpecByTimeScalingRuleSpec {
	s.EndTime = &v
	return s
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) SetLaunchTime(v int64) *ScalingRuleSpecByTimeScalingRuleSpec {
	s.LaunchTime = &v
	return s
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) SetRecurrenceType(v string) *ScalingRuleSpecByTimeScalingRuleSpec {
	s.RecurrenceType = &v
	return s
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) SetRecurrenceValue(v string) *ScalingRuleSpecByTimeScalingRuleSpec {
	s.RecurrenceValue = &v
	return s
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) Validate() error {
	return dara.Validate(s)
}
