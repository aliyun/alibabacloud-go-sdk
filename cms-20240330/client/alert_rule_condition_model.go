// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleCondition interface {
	dara.Model
	String() string
	GoString() string
	SetAlertCount(v int32) *AlertRuleCondition
	GetAlertCount() *int32
	SetCaseList(v []*AlertRuleConditionCaseList) *AlertRuleCondition
	GetCaseList() []*AlertRuleConditionCaseList
	SetCompareList(v []*AlertRuleConditionCompareList) *AlertRuleCondition
	GetCompareList() []*AlertRuleConditionCompareList
	SetCompositeEscalation(v *AlertRuleConditionCompositeEscalation) *AlertRuleCondition
	GetCompositeEscalation() *AlertRuleConditionCompositeEscalation
	SetEscalationType(v string) *AlertRuleCondition
	GetEscalationType() *string
	SetExpressEscalation(v *AlertRuleConditionExpressEscalation) *AlertRuleCondition
	GetExpressEscalation() *AlertRuleConditionExpressEscalation
	SetNoDataAlertLevel(v string) *AlertRuleCondition
	GetNoDataAlertLevel() *string
	SetNoDataAppendValue(v string) *AlertRuleCondition
	GetNoDataAppendValue() *string
	SetNoDataPolicy(v string) *AlertRuleCondition
	GetNoDataPolicy() *string
	SetOper(v string) *AlertRuleCondition
	GetOper() *string
	SetRelation(v string) *AlertRuleCondition
	GetRelation() *string
	SetSimpleEscalation(v *AlertRuleConditionSimpleEscalation) *AlertRuleCondition
	GetSimpleEscalation() *AlertRuleConditionSimpleEscalation
	SetType(v string) *AlertRuleCondition
	GetType() *string
	SetValue(v float64) *AlertRuleCondition
	GetValue() *float64
}

type AlertRuleCondition struct {
	// 适用条件类型：SLS_CONDITION
	//
	// 满足条件几次后告警，默认为1
	//
	// example:
	//
	// 1
	AlertCount *int32 `json:"alertCount,omitempty" xml:"alertCount,omitempty"`
	// 适用条件类型： SLS_CONDITION。
	//
	// SLS告警条件列表
	CaseList []*AlertRuleConditionCaseList `json:"caseList,omitempty" xml:"caseList,omitempty" type:"Repeated"`
	// 适用条件类型：APM_CONDITION。
	//
	// APM告警比较条件列表
	CompareList []*AlertRuleConditionCompareList `json:"compareList,omitempty" xml:"compareList,omitempty" type:"Repeated"`
	// 适用条件类型：CMS_BASIC_CONDITION
	//
	// escalationType=composite时有效，组合指标告警条件
	CompositeEscalation *AlertRuleConditionCompositeEscalation `json:"compositeEscalation,omitempty" xml:"compositeEscalation,omitempty" type:"Struct"`
	// 适用条件类型：CMS_BASIC_CONDITION
	//
	//
	// 取值范围:
	//
	// - simple: 简单指标条件
	//
	// - composite: 组合指标条件
	//
	// - express: 表达式条件
	//
	// example:
	//
	// simple
	EscalationType *string `json:"escalationType,omitempty" xml:"escalationType,omitempty"`
	// 适用条件类型：CMS_BASIC_CONDITION。
	//
	// escalationType=composite时有效，多指标组合告警条件。
	ExpressEscalation *AlertRuleConditionExpressEscalation `json:"expressEscalation,omitempty" xml:"expressEscalation,omitempty" type:"Struct"`
	// 适用条件类型：APM_CONDITION。
	//
	// 无数据时的告警级别，不指定则不对无数据报警
	//
	// example:
	//
	// INFO
	NoDataAlertLevel *string `json:"noDataAlertLevel,omitempty" xml:"noDataAlertLevel,omitempty"`
	// 适用条件类型：APM_CONDITION。
	//
	// 无数据时补偿的值。
	//
	// example:
	//
	// 1
	NoDataAppendValue *string `json:"noDataAppendValue,omitempty" xml:"noDataAppendValue,omitempty"`
	// 适用条件类型：CMS_BASIC_CONDITION。
	//
	//
	// 无监控数据时报警的处理方式。取值：
	//
	// - KEEP_LAST_STATE（默认值）：不做任何处理。
	//
	// - INSUFFICIENT_DATA：报警内容为无数据。
	//
	// - OK：正常。
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"noDataPolicy,omitempty" xml:"noDataPolicy,omitempty"`
	// 比较操作，判断是否是同比、环比
	//
	// - 大于 GT
	//
	// - 大于等于 GTE
	//
	// - 小于 LT
	//
	// - 小于等于 LTE
	//
	// - 等于 EQ
	//
	// - 不等于 NE
	//
	// - 同比增加 YOY_UP
	//
	// - 同比减少 YOY_DOWN
	//
	// example:
	//
	// LT
	Oper *string `json:"oper,omitempty" xml:"oper,omitempty"`
	// 适用条件类型：APM_CONDITION。
	//
	// 多个条件之间的逻辑关系。 取值：
	//
	// - and
	//
	// - or
	//
	// example:
	//
	// and
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// 适用条件类型：CMS_BASIC_CONDITION。
	//
	// 仅当escalationType=simple时有效，针对单一指标设置的告警条件
	SimpleEscalation *AlertRuleConditionSimpleEscalation `json:"simpleEscalation,omitempty" xml:"simpleEscalation,omitempty" type:"Struct"`
	// 规则条件类型，取值范围：
	//
	// - SLS_CONDITION(SLS告警条件)
	//
	// - APM_CONDITION(APM告警条件)
	//
	// - CMS_BASIC_CONDITION(基础云监控告警条件)
	//
	// This parameter is required.
	//
	// example:
	//
	// SLS_CONDITION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 告警触发的阈值。
	//
	// example:
	//
	// 60
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleCondition) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleCondition) GoString() string {
	return s.String()
}

func (s *AlertRuleCondition) GetAlertCount() *int32 {
	return s.AlertCount
}

func (s *AlertRuleCondition) GetCaseList() []*AlertRuleConditionCaseList {
	return s.CaseList
}

func (s *AlertRuleCondition) GetCompareList() []*AlertRuleConditionCompareList {
	return s.CompareList
}

func (s *AlertRuleCondition) GetCompositeEscalation() *AlertRuleConditionCompositeEscalation {
	return s.CompositeEscalation
}

func (s *AlertRuleCondition) GetEscalationType() *string {
	return s.EscalationType
}

func (s *AlertRuleCondition) GetExpressEscalation() *AlertRuleConditionExpressEscalation {
	return s.ExpressEscalation
}

func (s *AlertRuleCondition) GetNoDataAlertLevel() *string {
	return s.NoDataAlertLevel
}

func (s *AlertRuleCondition) GetNoDataAppendValue() *string {
	return s.NoDataAppendValue
}

func (s *AlertRuleCondition) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *AlertRuleCondition) GetOper() *string {
	return s.Oper
}

func (s *AlertRuleCondition) GetRelation() *string {
	return s.Relation
}

func (s *AlertRuleCondition) GetSimpleEscalation() *AlertRuleConditionSimpleEscalation {
	return s.SimpleEscalation
}

func (s *AlertRuleCondition) GetType() *string {
	return s.Type
}

func (s *AlertRuleCondition) GetValue() *float64 {
	return s.Value
}

func (s *AlertRuleCondition) SetAlertCount(v int32) *AlertRuleCondition {
	s.AlertCount = &v
	return s
}

func (s *AlertRuleCondition) SetCaseList(v []*AlertRuleConditionCaseList) *AlertRuleCondition {
	s.CaseList = v
	return s
}

func (s *AlertRuleCondition) SetCompareList(v []*AlertRuleConditionCompareList) *AlertRuleCondition {
	s.CompareList = v
	return s
}

func (s *AlertRuleCondition) SetCompositeEscalation(v *AlertRuleConditionCompositeEscalation) *AlertRuleCondition {
	s.CompositeEscalation = v
	return s
}

func (s *AlertRuleCondition) SetEscalationType(v string) *AlertRuleCondition {
	s.EscalationType = &v
	return s
}

func (s *AlertRuleCondition) SetExpressEscalation(v *AlertRuleConditionExpressEscalation) *AlertRuleCondition {
	s.ExpressEscalation = v
	return s
}

func (s *AlertRuleCondition) SetNoDataAlertLevel(v string) *AlertRuleCondition {
	s.NoDataAlertLevel = &v
	return s
}

func (s *AlertRuleCondition) SetNoDataAppendValue(v string) *AlertRuleCondition {
	s.NoDataAppendValue = &v
	return s
}

func (s *AlertRuleCondition) SetNoDataPolicy(v string) *AlertRuleCondition {
	s.NoDataPolicy = &v
	return s
}

func (s *AlertRuleCondition) SetOper(v string) *AlertRuleCondition {
	s.Oper = &v
	return s
}

func (s *AlertRuleCondition) SetRelation(v string) *AlertRuleCondition {
	s.Relation = &v
	return s
}

func (s *AlertRuleCondition) SetSimpleEscalation(v *AlertRuleConditionSimpleEscalation) *AlertRuleCondition {
	s.SimpleEscalation = v
	return s
}

func (s *AlertRuleCondition) SetType(v string) *AlertRuleCondition {
	s.Type = &v
	return s
}

func (s *AlertRuleCondition) SetValue(v float64) *AlertRuleCondition {
	s.Value = &v
	return s
}

func (s *AlertRuleCondition) Validate() error {
	if s.CaseList != nil {
		for _, item := range s.CaseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CompareList != nil {
		for _, item := range s.CompareList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CompositeEscalation != nil {
		if err := s.CompositeEscalation.Validate(); err != nil {
			return err
		}
	}
	if s.ExpressEscalation != nil {
		if err := s.ExpressEscalation.Validate(); err != nil {
			return err
		}
	}
	if s.SimpleEscalation != nil {
		if err := s.SimpleEscalation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AlertRuleConditionCaseList struct {
	// 匹配表达式，示例：  logLevel: error
	//
	// example:
	//
	// logLevel: error
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// 数量匹配表达式，示例：
	//
	// 区间组合： __count__ >= 3 && __count__ <= 10
	//
	// 单区间： __count__ >= 3
	//
	// example:
	//
	// count >= 3
	CountCondition *string `json:"countCondition,omitempty" xml:"countCondition,omitempty"`
	// 满足条件后的告警级别
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// 匹配类型： 有数据/有特定条数据/有数据匹配/有特定条数匹配。
	//
	// 取值范围：
	//
	// - HasData: 有数据
	//
	// - HasDataCount:  有特定条数据
	//
	// - HasDataMatch：有数据匹配
	//
	// - HasDataMatchCount：有特定条数匹配
	//
	// example:
	//
	// HasData
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleConditionCaseList) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionCaseList) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionCaseList) GetCondition() *string {
	return s.Condition
}

func (s *AlertRuleConditionCaseList) GetCountCondition() *string {
	return s.CountCondition
}

func (s *AlertRuleConditionCaseList) GetLevel() *string {
	return s.Level
}

func (s *AlertRuleConditionCaseList) GetType() *string {
	return s.Type
}

func (s *AlertRuleConditionCaseList) SetCondition(v string) *AlertRuleConditionCaseList {
	s.Condition = &v
	return s
}

func (s *AlertRuleConditionCaseList) SetCountCondition(v string) *AlertRuleConditionCaseList {
	s.CountCondition = &v
	return s
}

func (s *AlertRuleConditionCaseList) SetLevel(v string) *AlertRuleConditionCaseList {
	s.Level = &v
	return s
}

func (s *AlertRuleConditionCaseList) SetType(v string) *AlertRuleConditionCaseList {
	s.Type = &v
	return s
}

func (s *AlertRuleConditionCaseList) Validate() error {
	return dara.Validate(s)
}

type AlertRuleConditionCompareList struct {
	// 时间序列后聚合函数
	//
	// - count
	//
	// - sum
	//
	// - avg
	//
	// - min
	//
	// - max
	//
	// - p90
	//
	// - p95
	//
	// - p99
	//
	// example:
	//
	// count
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// 数据单位
	//
	// example:
	//
	// %
	BaseUnit *string `json:"baseUnit,omitempty" xml:"baseUnit,omitempty"`
	// 展示单位
	//
	// example:
	//
	// %
	DisplayUnit *string `json:"displayUnit,omitempty" xml:"displayUnit,omitempty"`
	// 比较操作，判断是否是同比、环比
	//
	// - 大于 GT
	//
	// - 大于等于 GTE
	//
	// - 小于 LT
	//
	// - 小于等于 LTE
	//
	// - 等于 EQ
	//
	// - 不等于 NE
	//
	// - 同比增加 YOY_UP
	//
	// - 同比减少 YOY_DOWN
	//
	// example:
	//
	// GT
	Oper *string `json:"oper,omitempty" xml:"oper,omitempty"`
	// 对比的阈值
	//
	// example:
	//
	// 50
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
	// 不同值的报警级别的列表。
	ValueLevelList []*AlertRuleConditionCompareListValueLevelList `json:"valueLevelList,omitempty" xml:"valueLevelList,omitempty" type:"Repeated"`
	// 同比时间单位（仅对oper=YOY_UP/YOY_DOWN生效）
	//
	//  minute、hour、day、week、month
	//
	// example:
	//
	// month
	YoyTimeUnit *string `json:"yoyTimeUnit,omitempty" xml:"yoyTimeUnit,omitempty"`
	// 同比时间的值，与yoyTimeUnit配合使用
	//
	// example:
	//
	// 1
	YoyTimeValue *int32 `json:"yoyTimeValue,omitempty" xml:"yoyTimeValue,omitempty"`
}

func (s AlertRuleConditionCompareList) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionCompareList) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionCompareList) GetAggregate() *string {
	return s.Aggregate
}

func (s *AlertRuleConditionCompareList) GetBaseUnit() *string {
	return s.BaseUnit
}

func (s *AlertRuleConditionCompareList) GetDisplayUnit() *string {
	return s.DisplayUnit
}

func (s *AlertRuleConditionCompareList) GetOper() *string {
	return s.Oper
}

func (s *AlertRuleConditionCompareList) GetValue() *float64 {
	return s.Value
}

func (s *AlertRuleConditionCompareList) GetValueLevelList() []*AlertRuleConditionCompareListValueLevelList {
	return s.ValueLevelList
}

func (s *AlertRuleConditionCompareList) GetYoyTimeUnit() *string {
	return s.YoyTimeUnit
}

func (s *AlertRuleConditionCompareList) GetYoyTimeValue() *int32 {
	return s.YoyTimeValue
}

func (s *AlertRuleConditionCompareList) SetAggregate(v string) *AlertRuleConditionCompareList {
	s.Aggregate = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetBaseUnit(v string) *AlertRuleConditionCompareList {
	s.BaseUnit = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetDisplayUnit(v string) *AlertRuleConditionCompareList {
	s.DisplayUnit = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetOper(v string) *AlertRuleConditionCompareList {
	s.Oper = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetValue(v float64) *AlertRuleConditionCompareList {
	s.Value = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetValueLevelList(v []*AlertRuleConditionCompareListValueLevelList) *AlertRuleConditionCompareList {
	s.ValueLevelList = v
	return s
}

func (s *AlertRuleConditionCompareList) SetYoyTimeUnit(v string) *AlertRuleConditionCompareList {
	s.YoyTimeUnit = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetYoyTimeValue(v int32) *AlertRuleConditionCompareList {
	s.YoyTimeValue = &v
	return s
}

func (s *AlertRuleConditionCompareList) Validate() error {
	if s.ValueLevelList != nil {
		for _, item := range s.ValueLevelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleConditionCompareListValueLevelList struct {
	// 阈值对应的级别
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// 对比的阈值
	//
	// example:
	//
	// 120
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleConditionCompareListValueLevelList) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionCompareListValueLevelList) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionCompareListValueLevelList) GetLevel() *string {
	return s.Level
}

func (s *AlertRuleConditionCompareListValueLevelList) GetValue() *float64 {
	return s.Value
}

func (s *AlertRuleConditionCompareListValueLevelList) SetLevel(v string) *AlertRuleConditionCompareListValueLevelList {
	s.Level = &v
	return s
}

func (s *AlertRuleConditionCompareListValueLevelList) SetValue(v float64) *AlertRuleConditionCompareListValueLevelList {
	s.Value = &v
	return s
}

func (s *AlertRuleConditionCompareListValueLevelList) Validate() error {
	return dara.Validate(s)
}

type AlertRuleConditionCompositeEscalation struct {
	// 多指标的组合条件列表
	Escalations []*AlertRuleConditionCompositeEscalationEscalations `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// 满足条件时触发的告警级别(多指标组合告警仅支持一个级别)
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// 多个指标条件之间的关系，取值为and或or
	//
	// example:
	//
	// and
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// 触发告警需满足条件的次数
	//
	// example:
	//
	// 3
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s AlertRuleConditionCompositeEscalation) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionCompositeEscalation) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionCompositeEscalation) GetEscalations() []*AlertRuleConditionCompositeEscalationEscalations {
	return s.Escalations
}

func (s *AlertRuleConditionCompositeEscalation) GetLevel() *string {
	return s.Level
}

func (s *AlertRuleConditionCompositeEscalation) GetRelation() *string {
	return s.Relation
}

func (s *AlertRuleConditionCompositeEscalation) GetTimes() *int32 {
	return s.Times
}

func (s *AlertRuleConditionCompositeEscalation) SetEscalations(v []*AlertRuleConditionCompositeEscalationEscalations) *AlertRuleConditionCompositeEscalation {
	s.Escalations = v
	return s
}

func (s *AlertRuleConditionCompositeEscalation) SetLevel(v string) *AlertRuleConditionCompositeEscalation {
	s.Level = &v
	return s
}

func (s *AlertRuleConditionCompositeEscalation) SetRelation(v string) *AlertRuleConditionCompositeEscalation {
	s.Relation = &v
	return s
}

func (s *AlertRuleConditionCompositeEscalation) SetTimes(v int32) *AlertRuleConditionCompositeEscalation {
	s.Times = &v
	return s
}

func (s *AlertRuleConditionCompositeEscalation) Validate() error {
	if s.Escalations != nil {
		for _, item := range s.Escalations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleConditionCompositeEscalationEscalations struct {
	// 阈值比较符，取值范围：
	//
	// - GreaterThanOrEqualToThreshold：大于等于。
	//
	// - GreaterThanThreshold：大于。
	//
	// - LessThanOrEqualToThreshold：小于等于。
	//
	// - LessThanThreshold：小于。
	//
	// - NotEqualToThreshold：不等于。
	//
	// - EqualToThreshold：等于。
	//
	// - GreaterThanYesterday：同比昨天时间上涨。
	//
	// - LessThanYesterday：同比昨天时间下降。
	//
	// - GreaterThanLastWeek：同比上周同一时间上涨。
	//
	// - LessThanLastWeek：同比上周同一时间下降。
	//
	// - GreaterThanLastPeriod：环比上周期上涨。
	//
	// - LessThanLastPeriod：环比上周期下降。
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	// 指标名称
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// 指标的时间窗口
	//
	// example:
	//
	// 60
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// 统计方法，该参数的取值由指定云产品的MetricName对应的Statistics列决定。  监控项的统计方法。取值示例：
	//
	// - $Maximum：最大值。
	//
	// - $Minimum：最小值。
	//
	// - $Average：平均值。
	//
	// - $Availability：可用率（通常用于站点监控）
	//
	// 说明 $为监控项的统一前缀符号。
	//
	// example:
	//
	// $Maximum
	Statistics *string `json:"statistics,omitempty" xml:"statistics,omitempty"`
	// 告警阈值
	//
	// example:
	//
	// 50
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s AlertRuleConditionCompositeEscalationEscalations) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionCompositeEscalationEscalations) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionCompositeEscalationEscalations) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *AlertRuleConditionCompositeEscalationEscalations) GetMetricName() *string {
	return s.MetricName
}

func (s *AlertRuleConditionCompositeEscalationEscalations) GetPeriod() *int64 {
	return s.Period
}

func (s *AlertRuleConditionCompositeEscalationEscalations) GetStatistics() *string {
	return s.Statistics
}

func (s *AlertRuleConditionCompositeEscalationEscalations) GetThreshold() *float64 {
	return s.Threshold
}

func (s *AlertRuleConditionCompositeEscalationEscalations) SetComparisonOperator(v string) *AlertRuleConditionCompositeEscalationEscalations {
	s.ComparisonOperator = &v
	return s
}

func (s *AlertRuleConditionCompositeEscalationEscalations) SetMetricName(v string) *AlertRuleConditionCompositeEscalationEscalations {
	s.MetricName = &v
	return s
}

func (s *AlertRuleConditionCompositeEscalationEscalations) SetPeriod(v int64) *AlertRuleConditionCompositeEscalationEscalations {
	s.Period = &v
	return s
}

func (s *AlertRuleConditionCompositeEscalationEscalations) SetStatistics(v string) *AlertRuleConditionCompositeEscalationEscalations {
	s.Statistics = &v
	return s
}

func (s *AlertRuleConditionCompositeEscalationEscalations) SetThreshold(v float64) *AlertRuleConditionCompositeEscalationEscalations {
	s.Threshold = &v
	return s
}

func (s *AlertRuleConditionCompositeEscalationEscalations) Validate() error {
	return dara.Validate(s)
}

type AlertRuleConditionExpressEscalation struct {
	// 满足条件时触发的告警级别(表达式告警仅支持一个级别)
	//
	// - CRITICAL
	//
	// - WARNING
	//
	// - INFO
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// 告警条件表达式
	//
	// example:
	//
	// @cpu_total[60].$Average > 60
	RawExpression *string `json:"rawExpression,omitempty" xml:"rawExpression,omitempty"`
	// 触发告警需满足条件的次数
	//
	// example:
	//
	// 3
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s AlertRuleConditionExpressEscalation) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionExpressEscalation) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionExpressEscalation) GetLevel() *string {
	return s.Level
}

func (s *AlertRuleConditionExpressEscalation) GetRawExpression() *string {
	return s.RawExpression
}

func (s *AlertRuleConditionExpressEscalation) GetTimes() *int32 {
	return s.Times
}

func (s *AlertRuleConditionExpressEscalation) SetLevel(v string) *AlertRuleConditionExpressEscalation {
	s.Level = &v
	return s
}

func (s *AlertRuleConditionExpressEscalation) SetRawExpression(v string) *AlertRuleConditionExpressEscalation {
	s.RawExpression = &v
	return s
}

func (s *AlertRuleConditionExpressEscalation) SetTimes(v int32) *AlertRuleConditionExpressEscalation {
	s.Times = &v
	return s
}

func (s *AlertRuleConditionExpressEscalation) Validate() error {
	return dara.Validate(s)
}

type AlertRuleConditionSimpleEscalation struct {
	// 条件列表，同一个告警规则对应多个级别的，每一个级别会有一个条件对象。
	Escalations []*AlertRuleConditionSimpleEscalationEscalations `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// 适用条件类型：CMS_BASIC_CONDITION。
	//
	// 告警条件关联的指标
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// 指标的时间窗口，单位秒
	//
	// example:
	//
	// 60
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
}

func (s AlertRuleConditionSimpleEscalation) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionSimpleEscalation) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionSimpleEscalation) GetEscalations() []*AlertRuleConditionSimpleEscalationEscalations {
	return s.Escalations
}

func (s *AlertRuleConditionSimpleEscalation) GetMetricName() *string {
	return s.MetricName
}

func (s *AlertRuleConditionSimpleEscalation) GetPeriod() *int64 {
	return s.Period
}

func (s *AlertRuleConditionSimpleEscalation) SetEscalations(v []*AlertRuleConditionSimpleEscalationEscalations) *AlertRuleConditionSimpleEscalation {
	s.Escalations = v
	return s
}

func (s *AlertRuleConditionSimpleEscalation) SetMetricName(v string) *AlertRuleConditionSimpleEscalation {
	s.MetricName = &v
	return s
}

func (s *AlertRuleConditionSimpleEscalation) SetPeriod(v int64) *AlertRuleConditionSimpleEscalation {
	s.Period = &v
	return s
}

func (s *AlertRuleConditionSimpleEscalation) Validate() error {
	if s.Escalations != nil {
		for _, item := range s.Escalations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleConditionSimpleEscalationEscalations struct {
	// 阈值比较符，取值范围：
	//
	// - GreaterThanOrEqualToThreshold：大于等于。
	//
	// - GreaterThanThreshold：大于。
	//
	// - LessThanOrEqualToThreshold：小于等于。
	//
	// - LessThanThreshold：小于。
	//
	// - NotEqualToThreshold：不等于。
	//
	// - EqualToThreshold：等于。
	//
	// - GreaterThanYesterday：同比昨天时间上涨。
	//
	// - LessThanYesterday：同比昨天时间下降。
	//
	// - GreaterThanLastWeek：同比上周同一时间上涨。
	//
	// - LessThanLastWeek：同比上周同一时间下降。
	//
	// - GreaterThanLastPeriod：环比上周期上涨。
	//
	// - LessThanLastPeriod：环比上周期下降。
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	// 满足条件时触发的告警级别(表达式告警仅支持一个级别)
	//
	// - CRITICAL
	//
	// - WARNING
	//
	// - INFO
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// 统计方法，该参数的取值由指定云产品的MetricName对应的Statistics列决定，例如：Maximum、Minimum 和 Average
	//
	// example:
	//
	// Average
	Statistics *string `json:"statistics,omitempty" xml:"statistics,omitempty"`
	// 告警阈值
	//
	// example:
	//
	// 100
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// 触发告警需满足条件的次数
	//
	// example:
	//
	// 3
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s AlertRuleConditionSimpleEscalationEscalations) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionSimpleEscalationEscalations) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionSimpleEscalationEscalations) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *AlertRuleConditionSimpleEscalationEscalations) GetLevel() *string {
	return s.Level
}

func (s *AlertRuleConditionSimpleEscalationEscalations) GetStatistics() *string {
	return s.Statistics
}

func (s *AlertRuleConditionSimpleEscalationEscalations) GetThreshold() *float64 {
	return s.Threshold
}

func (s *AlertRuleConditionSimpleEscalationEscalations) GetTimes() *int32 {
	return s.Times
}

func (s *AlertRuleConditionSimpleEscalationEscalations) SetComparisonOperator(v string) *AlertRuleConditionSimpleEscalationEscalations {
	s.ComparisonOperator = &v
	return s
}

func (s *AlertRuleConditionSimpleEscalationEscalations) SetLevel(v string) *AlertRuleConditionSimpleEscalationEscalations {
	s.Level = &v
	return s
}

func (s *AlertRuleConditionSimpleEscalationEscalations) SetStatistics(v string) *AlertRuleConditionSimpleEscalationEscalations {
	s.Statistics = &v
	return s
}

func (s *AlertRuleConditionSimpleEscalationEscalations) SetThreshold(v float64) *AlertRuleConditionSimpleEscalationEscalations {
	s.Threshold = &v
	return s
}

func (s *AlertRuleConditionSimpleEscalationEscalations) SetTimes(v int32) *AlertRuleConditionSimpleEscalationEscalations {
	s.Times = &v
	return s
}

func (s *AlertRuleConditionSimpleEscalationEscalations) Validate() error {
	return dara.Validate(s)
}
