// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingRuleV1 interface {
	dara.Model
	String() string
	GoString() string
	SetAdjustmentType(v string) *ScalingRuleV1
	GetAdjustmentType() *string
	SetAdjustmentValue(v int32) *ScalingRuleV1
	GetAdjustmentValue() *int32
	SetCoolDownTime(v int32) *ScalingRuleV1
	GetCoolDownTime() *int32
	SetRuleName(v string) *ScalingRuleV1
	GetRuleName() *string
	SetRuleParam(v *ScalingRuleV1RuleParam) *ScalingRuleV1
	GetRuleParam() *ScalingRuleV1RuleParam
	SetRuleType(v string) *ScalingRuleV1
	GetRuleType() *string
	SetScalingConfigBizId(v string) *ScalingRuleV1
	GetScalingConfigBizId() *string
}

type ScalingRuleV1 struct {
	// 调整类型。
	//
	// example:
	//
	// QUANTITY_CHANGE_IN_CAPACITY
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// 调整值,正数为扩容,负数为缩容。
	//
	// example:
	//
	// 1
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// 冷却时间,单位秒。
	//
	// example:
	//
	// 4
	CoolDownTime *int32 `json:"CoolDownTime,omitempty" xml:"CoolDownTime,omitempty"`
	// 规则名称。
	//
	// example:
	//
	// tule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 规则参数。
	RuleParam *ScalingRuleV1RuleParam `json:"RuleParam,omitempty" xml:"RuleParam,omitempty" type:"Struct"`
	// 规则类型。
	//
	// example:
	//
	// BY_LOAD
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// 弹性规则配置ID。
	//
	// example:
	//
	// SCB-DCD96BCCFED1****
	ScalingConfigBizId *string `json:"ScalingConfigBizId,omitempty" xml:"ScalingConfigBizId,omitempty"`
}

func (s ScalingRuleV1) String() string {
	return dara.Prettify(s)
}

func (s ScalingRuleV1) GoString() string {
	return s.String()
}

func (s *ScalingRuleV1) GetAdjustmentType() *string {
	return s.AdjustmentType
}

func (s *ScalingRuleV1) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *ScalingRuleV1) GetCoolDownTime() *int32 {
	return s.CoolDownTime
}

func (s *ScalingRuleV1) GetRuleName() *string {
	return s.RuleName
}

func (s *ScalingRuleV1) GetRuleParam() *ScalingRuleV1RuleParam {
	return s.RuleParam
}

func (s *ScalingRuleV1) GetRuleType() *string {
	return s.RuleType
}

func (s *ScalingRuleV1) GetScalingConfigBizId() *string {
	return s.ScalingConfigBizId
}

func (s *ScalingRuleV1) SetAdjustmentType(v string) *ScalingRuleV1 {
	s.AdjustmentType = &v
	return s
}

func (s *ScalingRuleV1) SetAdjustmentValue(v int32) *ScalingRuleV1 {
	s.AdjustmentValue = &v
	return s
}

func (s *ScalingRuleV1) SetCoolDownTime(v int32) *ScalingRuleV1 {
	s.CoolDownTime = &v
	return s
}

func (s *ScalingRuleV1) SetRuleName(v string) *ScalingRuleV1 {
	s.RuleName = &v
	return s
}

func (s *ScalingRuleV1) SetRuleParam(v *ScalingRuleV1RuleParam) *ScalingRuleV1 {
	s.RuleParam = v
	return s
}

func (s *ScalingRuleV1) SetRuleType(v string) *ScalingRuleV1 {
	s.RuleType = &v
	return s
}

func (s *ScalingRuleV1) SetScalingConfigBizId(v string) *ScalingRuleV1 {
	s.ScalingConfigBizId = &v
	return s
}

func (s *ScalingRuleV1) Validate() error {
	return dara.Validate(s)
}

type ScalingRuleV1RuleParam struct {
	// [负载触发参数] 比较符。
	//
	// example:
	//
	// >
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// [负载触发参数] 统计次数。
	//
	// example:
	//
	// 1
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// [时间调度参数] 周期类型周期过期时间。
	//
	// example:
	//
	// 0
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	// [时间调度参数] 周期类型周期开始时间。
	//
	// example:
	//
	// 2021-09-15T04:02Z
	LaunchTime *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// [负载触发参数] 度量名称。
	//
	// example:
	//
	// YarnRootAvailableVCores
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// [负载触发参数] 统计时长,单位分钟。
	//
	// example:
	//
	// 5
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// [时间调度参数] 周期类型周期结束时间。
	//
	// example:
	//
	// 2021-09-16T05:02Z
	RecurrenceEndTime *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	// [时间调度参数] 周期类型。
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// [时间调度参数] 周期类型周期值。
	//
	// example:
	//
	// 1
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// [负载触发参数] 统计方式。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// [负载触发参数] 阈值。
	//
	// example:
	//
	// 1
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ScalingRuleV1RuleParam) String() string {
	return dara.Prettify(s)
}

func (s ScalingRuleV1RuleParam) GoString() string {
	return s.String()
}

func (s *ScalingRuleV1RuleParam) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ScalingRuleV1RuleParam) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *ScalingRuleV1RuleParam) GetLaunchExpirationTime() *int32 {
	return s.LaunchExpirationTime
}

func (s *ScalingRuleV1RuleParam) GetLaunchTime() *string {
	return s.LaunchTime
}

func (s *ScalingRuleV1RuleParam) GetMetricName() *string {
	return s.MetricName
}

func (s *ScalingRuleV1RuleParam) GetPeriod() *int32 {
	return s.Period
}

func (s *ScalingRuleV1RuleParam) GetRecurrenceEndTime() *string {
	return s.RecurrenceEndTime
}

func (s *ScalingRuleV1RuleParam) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *ScalingRuleV1RuleParam) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *ScalingRuleV1RuleParam) GetStatistics() *string {
	return s.Statistics
}

func (s *ScalingRuleV1RuleParam) GetThreshold() *int32 {
	return s.Threshold
}

func (s *ScalingRuleV1RuleParam) SetComparisonOperator(v string) *ScalingRuleV1RuleParam {
	s.ComparisonOperator = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetEvaluationCount(v int32) *ScalingRuleV1RuleParam {
	s.EvaluationCount = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetLaunchExpirationTime(v int32) *ScalingRuleV1RuleParam {
	s.LaunchExpirationTime = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetLaunchTime(v string) *ScalingRuleV1RuleParam {
	s.LaunchTime = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetMetricName(v string) *ScalingRuleV1RuleParam {
	s.MetricName = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetPeriod(v int32) *ScalingRuleV1RuleParam {
	s.Period = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetRecurrenceEndTime(v string) *ScalingRuleV1RuleParam {
	s.RecurrenceEndTime = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetRecurrenceType(v string) *ScalingRuleV1RuleParam {
	s.RecurrenceType = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetRecurrenceValue(v string) *ScalingRuleV1RuleParam {
	s.RecurrenceValue = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetStatistics(v string) *ScalingRuleV1RuleParam {
	s.Statistics = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetThreshold(v int32) *ScalingRuleV1RuleParam {
	s.Threshold = &v
	return s
}

func (s *ScalingRuleV1RuleParam) Validate() error {
	return dara.Validate(s)
}
