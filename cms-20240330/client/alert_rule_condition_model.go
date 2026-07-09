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
	SetCountOperator(v string) *AlertRuleCondition
	GetCountOperator() *string
	SetCountThreshold(v int64) *AlertRuleCondition
	GetCountThreshold() *int64
	SetEnableSeveritySuppression(v bool) *AlertRuleCondition
	GetEnableSeveritySuppression() *bool
	SetEscalationType(v string) *AlertRuleCondition
	GetEscalationType() *string
	SetExpressEscalation(v *AlertRuleConditionExpressEscalation) *AlertRuleCondition
	GetExpressEscalation() *AlertRuleConditionExpressEscalation
	SetMatchField(v string) *AlertRuleCondition
	GetMatchField() *string
	SetMatchOperator(v string) *AlertRuleCondition
	GetMatchOperator() *string
	SetMatchValue(v string) *AlertRuleCondition
	GetMatchValue() *string
	SetMax(v float64) *AlertRuleCondition
	GetMax() *float64
	SetMin(v float64) *AlertRuleCondition
	GetMin() *float64
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
	SetThresholdList(v []*AlertRuleConditionThresholdList) *AlertRuleCondition
	GetThresholdList() []*AlertRuleConditionThresholdList
	SetTriggers(v []*AlertRuleConditionTriggers) *AlertRuleCondition
	GetTriggers() []*AlertRuleConditionTriggers
	SetType(v string) *AlertRuleCondition
	GetType() *string
	SetValue(v float64) *AlertRuleCondition
	GetValue() *float64
}

type AlertRuleCondition struct {
	// Applicable condition type: SLS_CONDITION.
	//
	// The number of times the condition must be met before an alert is triggered. Default value: 1.
	//
	// example:
	//
	// 1
	AlertCount *int32 `json:"alertCount,omitempty" xml:"alertCount,omitempty"`
	// Applicable condition type: SLS_CONDITION.
	//
	// The list of Simple Log Service alert conditions.
	CaseList []*AlertRuleConditionCaseList `json:"caseList,omitempty" xml:"caseList,omitempty" type:"Repeated"`
	// Applicable condition type: APM_CONDITION.
	//
	// The list of Application Performance Monitoring (APM) alert comparison conditions.
	CompareList []*AlertRuleConditionCompareList `json:"compareList,omitempty" xml:"compareList,omitempty" type:"Repeated"`
	// Applicable condition type: CMS_BASIC_CONDITION.
	//
	// This parameter takes effect only when escalationType is set to composite. The composite metric alert condition.
	CompositeEscalation       *AlertRuleConditionCompositeEscalation `json:"compositeEscalation,omitempty" xml:"compositeEscalation,omitempty" type:"Struct"`
	CountOperator             *string                                `json:"countOperator,omitempty" xml:"countOperator,omitempty"`
	CountThreshold            *int64                                 `json:"countThreshold,omitempty" xml:"countThreshold,omitempty"`
	EnableSeveritySuppression *bool                                  `json:"enableSeveritySuppression,omitempty" xml:"enableSeveritySuppression,omitempty"`
	// Applicable condition type: CMS_BASIC_CONDITION.
	//
	// Valid values:
	//
	// - simple: simple metric condition.
	//
	// - composite: composite metric condition.
	//
	// - express: expression condition.
	//
	// example:
	//
	// simple
	EscalationType *string `json:"escalationType,omitempty" xml:"escalationType,omitempty"`
	// Applicable condition type: CMS_BASIC_CONDITION.
	//
	// This parameter takes effect only when escalationType is set to composite. The multi-metric composite alert condition.
	ExpressEscalation *AlertRuleConditionExpressEscalation `json:"expressEscalation,omitempty" xml:"expressEscalation,omitempty" type:"Struct"`
	MatchField        *string                              `json:"matchField,omitempty" xml:"matchField,omitempty"`
	MatchOperator     *string                              `json:"matchOperator,omitempty" xml:"matchOperator,omitempty"`
	MatchValue        *string                              `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
	Max               *float64                             `json:"max,omitempty" xml:"max,omitempty"`
	Min               *float64                             `json:"min,omitempty" xml:"min,omitempty"`
	// Applicable condition type: APM_CONDITION.
	//
	// The alert level when no data is available. If this parameter is not specified, no alert is triggered when no data is available.
	//
	// example:
	//
	// INFO
	NoDataAlertLevel *string `json:"noDataAlertLevel,omitempty" xml:"noDataAlertLevel,omitempty"`
	// Applicable condition type: APM_CONDITION.
	//
	// The compensation value when no data is available.
	//
	// example:
	//
	// 1
	NoDataAppendValue *string `json:"noDataAppendValue,omitempty" xml:"noDataAppendValue,omitempty"`
	// Applicable condition type: CMS_BASIC_CONDITION.
	//
	// The method used to handle alerts when no monitoring data is available. Valid values:
	//
	// - KEEP_LAST_STATE (default): No action is taken.
	//
	// - INSUFFICIENT_DATA: The alert content indicates that no data is available.
	//
	// - OK: Normal.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"noDataPolicy,omitempty" xml:"noDataPolicy,omitempty"`
	// The comparison operator. Specifies whether to use year-over-year or period-over-period comparison. Valid values:
	//
	// - GT: greater than.
	//
	// - GTE: greater than or equal to.
	//
	// - LT: less than.
	//
	// - LTE: less than or equal to.
	//
	// - EQ: equal to.
	//
	// - NE: not equal to.
	//
	// - YOY_UP: year-over-year increase.
	//
	// - YOY_DOWN: year-over-year decrease.
	//
	// example:
	//
	// LT
	Oper *string `json:"oper,omitempty" xml:"oper,omitempty"`
	// Applicable condition type: APM_CONDITION.
	//
	// The logical relationship between multiple conditions. Valid values:
	//
	// - and
	//
	// - or
	//
	// example:
	//
	// and
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// Applicable condition type: CMS_BASIC_CONDITION.
	//
	// This parameter takes effect only when escalationType is set to simple. The alert condition configured for a single metric.
	SimpleEscalation *AlertRuleConditionSimpleEscalation `json:"simpleEscalation,omitempty" xml:"simpleEscalation,omitempty" type:"Struct"`
	ThresholdList    []*AlertRuleConditionThresholdList  `json:"thresholdList,omitempty" xml:"thresholdList,omitempty" type:"Repeated"`
	Triggers         []*AlertRuleConditionTriggers       `json:"triggers,omitempty" xml:"triggers,omitempty" type:"Repeated"`
	// The rule condition type. Valid values:
	//
	// - SLS_CONDITION: Simple Log Service alert condition.
	//
	// - APM_CONDITION: APM alert condition.
	//
	// - CMS_BASIC_CONDITION: CloudMonitor Basic monitoring alert condition.
	//
	// This parameter is required.
	//
	// example:
	//
	// SLS_CONDITION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The threshold that triggers the alert.
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

func (s *AlertRuleCondition) GetCountOperator() *string {
	return s.CountOperator
}

func (s *AlertRuleCondition) GetCountThreshold() *int64 {
	return s.CountThreshold
}

func (s *AlertRuleCondition) GetEnableSeveritySuppression() *bool {
	return s.EnableSeveritySuppression
}

func (s *AlertRuleCondition) GetEscalationType() *string {
	return s.EscalationType
}

func (s *AlertRuleCondition) GetExpressEscalation() *AlertRuleConditionExpressEscalation {
	return s.ExpressEscalation
}

func (s *AlertRuleCondition) GetMatchField() *string {
	return s.MatchField
}

func (s *AlertRuleCondition) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *AlertRuleCondition) GetMatchValue() *string {
	return s.MatchValue
}

func (s *AlertRuleCondition) GetMax() *float64 {
	return s.Max
}

func (s *AlertRuleCondition) GetMin() *float64 {
	return s.Min
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

func (s *AlertRuleCondition) GetThresholdList() []*AlertRuleConditionThresholdList {
	return s.ThresholdList
}

func (s *AlertRuleCondition) GetTriggers() []*AlertRuleConditionTriggers {
	return s.Triggers
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

func (s *AlertRuleCondition) SetCountOperator(v string) *AlertRuleCondition {
	s.CountOperator = &v
	return s
}

func (s *AlertRuleCondition) SetCountThreshold(v int64) *AlertRuleCondition {
	s.CountThreshold = &v
	return s
}

func (s *AlertRuleCondition) SetEnableSeveritySuppression(v bool) *AlertRuleCondition {
	s.EnableSeveritySuppression = &v
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

func (s *AlertRuleCondition) SetMatchField(v string) *AlertRuleCondition {
	s.MatchField = &v
	return s
}

func (s *AlertRuleCondition) SetMatchOperator(v string) *AlertRuleCondition {
	s.MatchOperator = &v
	return s
}

func (s *AlertRuleCondition) SetMatchValue(v string) *AlertRuleCondition {
	s.MatchValue = &v
	return s
}

func (s *AlertRuleCondition) SetMax(v float64) *AlertRuleCondition {
	s.Max = &v
	return s
}

func (s *AlertRuleCondition) SetMin(v float64) *AlertRuleCondition {
	s.Min = &v
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

func (s *AlertRuleCondition) SetThresholdList(v []*AlertRuleConditionThresholdList) *AlertRuleCondition {
	s.ThresholdList = v
	return s
}

func (s *AlertRuleCondition) SetTriggers(v []*AlertRuleConditionTriggers) *AlertRuleCondition {
	s.Triggers = v
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
	if s.ThresholdList != nil {
		for _, item := range s.ThresholdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Triggers != nil {
		for _, item := range s.Triggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleConditionCaseList struct {
	// The match expression. Example: logLevel: error.
	//
	// example:
	//
	// logLevel: error
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The count match expression. Examples:
	//
	// Range combination: __count__ >= 3 && __count__ <= 10
	//
	// Single range: __count__ >= 3
	//
	// example:
	//
	// count >= 3
	CountCondition *string `json:"countCondition,omitempty" xml:"countCondition,omitempty"`
	// The alert level when the condition is met.
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The match type: has data, has a specific number of data entries, has data match, or has a specific number of data matches.
	//
	// Valid values:
	//
	// - HasData: has data
	//
	// - HasDataCount: has a specific number of data entries
	//
	// - HasDataMatch: has data match
	//
	// - HasDataMatchCount: has a specific number of data matches
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
	// The aggregate functions applied after time series aggregation.
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
	// The data unit.
	//
	// example:
	//
	// %
	BaseUnit *string `json:"baseUnit,omitempty" xml:"baseUnit,omitempty"`
	// The display unit.
	//
	// example:
	//
	// %
	DisplayUnit *string `json:"displayUnit,omitempty" xml:"displayUnit,omitempty"`
	// The comparison operator. Determines whether year-over-year or period-over-period comparison is used.
	//
	// - Greater than: GT
	//
	// - Greater than or equal to: GTE
	//
	// - Less than: LT
	//
	// - Less than or equal to: LTE
	//
	// - Equal to: EQ
	//
	// - Not equal to: NE
	//
	// - Year-over-year increase: YOY_UP
	//
	// - Year-over-year decrease: YOY_DOWN
	//
	// example:
	//
	// GT
	Oper *string `json:"oper,omitempty" xml:"oper,omitempty"`
	// The comparison threshold.
	//
	// example:
	//
	// 50
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
	// The list of alert levels for different values.
	ValueLevelList []*AlertRuleConditionCompareListValueLevelList `json:"valueLevelList,omitempty" xml:"valueLevelList,omitempty" type:"Repeated"`
	// The time unit for year-over-year comparison. Valid only when oper is set to YOY_UP or YOY_DOWN.
	//
	// Valid values: minute, hour, day, week, month.
	//
	// example:
	//
	// month
	YoyTimeUnit *string `json:"yoyTimeUnit,omitempty" xml:"yoyTimeUnit,omitempty"`
	// The value of the year-over-year time period. Used together with yoyTimeUnit.
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
	// The level corresponding to the threshold.
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The comparison threshold.
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
	// The list of composite conditions for multiple metrics.
	Escalations []*AlertRuleConditionCompositeEscalationEscalations `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// The alert level triggered when the condition is met. Composite metric alerts support only one level.
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The relationship between multiple metric conditions. Valid values: and, or.
	//
	// example:
	//
	// and
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// The number of times the condition must be met to trigger an alert.
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
	// The threshold comparison operator. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: Greater than or equal to.
	//
	// - GreaterThanThreshold: Greater than.
	//
	// - LessThanOrEqualToThreshold: Less than or equal to.
	//
	// - LessThanThreshold: Less than.
	//
	// - NotEqualToThreshold: Not equal to.
	//
	// - EqualToThreshold: Equal to.
	//
	// - GreaterThanYesterday: Year-over-year increase compared with the same time yesterday.
	//
	// - LessThanYesterday: Year-over-year decrease compared with the same time yesterday.
	//
	// - GreaterThanLastWeek: Year-over-year increase compared with the same time last week.
	//
	// - LessThanLastWeek: Year-over-year decrease compared with the same time last week.
	//
	// - GreaterThanLastPeriod: Period-over-period increase compared with the previous period.
	//
	// - LessThanLastPeriod: Period-over-period decrease compared with the previous period.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	// The metric name.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The time window of the metric.
	//
	// example:
	//
	// 60
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The statistical method. The value of this parameter is determined by the Statistics column corresponding to the MetricName of the specified cloud service. Example values for the statistical method of a metric:
	//
	// - $Maximum: Maximum value.
	//
	// - $Minimum: Minimum value.
	//
	// - $Average: Average value.
	//
	// - $Availability: Availability rate (typically used for site monitoring).
	//
	// Note: $ is the unified prefix symbol for metrics.
	//
	// example:
	//
	// $Maximum
	Statistics *string `json:"statistics,omitempty" xml:"statistics,omitempty"`
	// The alert threshold.
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
	// The alert level triggered when the condition is met. Expression-based alerts support only one level.
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
	// The alert condition expression.
	//
	// example:
	//
	// @cpu_total[60].$Average > 60
	RawExpression *string `json:"rawExpression,omitempty" xml:"rawExpression,omitempty"`
	// The number of times the condition must be met to trigger an alert.
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
	// The list of conditions. When an alert rule corresponds to multiple levels, each level has a condition object.
	Escalations []*AlertRuleConditionSimpleEscalationEscalations `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// Applicable condition type: CMS_BASIC_CONDITION.
	//
	// The metric associated with the alert condition.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The time window of the metric. Unit: seconds.
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
	// The threshold comparison operator. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: Greater than or equal to.
	//
	// - GreaterThanThreshold: Greater than.
	//
	// - LessThanOrEqualToThreshold: Less than or equal to.
	//
	// - LessThanThreshold: Less than.
	//
	// - NotEqualToThreshold: Not equal to.
	//
	// - EqualToThreshold: Equal to.
	//
	// - GreaterThanYesterday: Year-over-year increase compared with the same time yesterday.
	//
	// - LessThanYesterday: Year-over-year decrease compared with the same time yesterday.
	//
	// - GreaterThanLastWeek: Year-over-year increase compared with the same time last week.
	//
	// - LessThanLastWeek: Year-over-year decrease compared with the same time last week.
	//
	// - GreaterThanLastPeriod: Period-over-period increase compared with the previous period.
	//
	// - LessThanLastPeriod: Period-over-period decrease compared with the previous period.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	// The alert level triggered when the condition is met. Expression-based alerts support only one level.
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
	// The statistical method. The value of this parameter is determined by the Statistics column corresponding to the MetricName of the specified cloud service. Examples: Maximum, Minimum, and Average.
	//
	// example:
	//
	// Average
	Statistics *string `json:"statistics,omitempty" xml:"statistics,omitempty"`
	// The alert threshold.
	//
	// example:
	//
	// 100
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// The number of times the condition must be met to trigger an alert.
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

type AlertRuleConditionThresholdList struct {
	Max       *float64 `json:"max,omitempty" xml:"max,omitempty"`
	Min       *float64 `json:"min,omitempty" xml:"min,omitempty"`
	Severity  *string  `json:"severity,omitempty" xml:"severity,omitempty"`
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s AlertRuleConditionThresholdList) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionThresholdList) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionThresholdList) GetMax() *float64 {
	return s.Max
}

func (s *AlertRuleConditionThresholdList) GetMin() *float64 {
	return s.Min
}

func (s *AlertRuleConditionThresholdList) GetSeverity() *string {
	return s.Severity
}

func (s *AlertRuleConditionThresholdList) GetThreshold() *float64 {
	return s.Threshold
}

func (s *AlertRuleConditionThresholdList) SetMax(v float64) *AlertRuleConditionThresholdList {
	s.Max = &v
	return s
}

func (s *AlertRuleConditionThresholdList) SetMin(v float64) *AlertRuleConditionThresholdList {
	s.Min = &v
	return s
}

func (s *AlertRuleConditionThresholdList) SetSeverity(v string) *AlertRuleConditionThresholdList {
	s.Severity = &v
	return s
}

func (s *AlertRuleConditionThresholdList) SetThreshold(v float64) *AlertRuleConditionThresholdList {
	s.Threshold = &v
	return s
}

func (s *AlertRuleConditionThresholdList) Validate() error {
	return dara.Validate(s)
}

type AlertRuleConditionTriggers struct {
	DurationSecs *int32                                `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	Expression   *AlertRuleConditionTriggersExpression `json:"expression,omitempty" xml:"expression,omitempty" type:"Struct"`
	Severity     *string                               `json:"severity,omitempty" xml:"severity,omitempty"`
}

func (s AlertRuleConditionTriggers) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionTriggers) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionTriggers) GetDurationSecs() *int32 {
	return s.DurationSecs
}

func (s *AlertRuleConditionTriggers) GetExpression() *AlertRuleConditionTriggersExpression {
	return s.Expression
}

func (s *AlertRuleConditionTriggers) GetSeverity() *string {
	return s.Severity
}

func (s *AlertRuleConditionTriggers) SetDurationSecs(v int32) *AlertRuleConditionTriggers {
	s.DurationSecs = &v
	return s
}

func (s *AlertRuleConditionTriggers) SetExpression(v *AlertRuleConditionTriggersExpression) *AlertRuleConditionTriggers {
	s.Expression = v
	return s
}

func (s *AlertRuleConditionTriggers) SetSeverity(v string) *AlertRuleConditionTriggers {
	s.Severity = &v
	return s
}

func (s *AlertRuleConditionTriggers) Validate() error {
	if s.Expression != nil {
		if err := s.Expression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AlertRuleConditionTriggersExpression struct {
	Conditions     []*AlertRuleConditionTriggersExpressionConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	ExpressionType *string                                           `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	LogicOperator  *string                                           `json:"logicOperator,omitempty" xml:"logicOperator,omitempty"`
}

func (s AlertRuleConditionTriggersExpression) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionTriggersExpression) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionTriggersExpression) GetConditions() []*AlertRuleConditionTriggersExpressionConditions {
	return s.Conditions
}

func (s *AlertRuleConditionTriggersExpression) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *AlertRuleConditionTriggersExpression) GetLogicOperator() *string {
	return s.LogicOperator
}

func (s *AlertRuleConditionTriggersExpression) SetConditions(v []*AlertRuleConditionTriggersExpressionConditions) *AlertRuleConditionTriggersExpression {
	s.Conditions = v
	return s
}

func (s *AlertRuleConditionTriggersExpression) SetExpressionType(v string) *AlertRuleConditionTriggersExpression {
	s.ExpressionType = &v
	return s
}

func (s *AlertRuleConditionTriggersExpression) SetLogicOperator(v string) *AlertRuleConditionTriggersExpression {
	s.LogicOperator = &v
	return s
}

func (s *AlertRuleConditionTriggersExpression) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleConditionTriggersExpressionConditions struct {
	ExpressionType *string  `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	Operator       *string  `json:"operator,omitempty" xml:"operator,omitempty"`
	QueryName      *string  `json:"queryName,omitempty" xml:"queryName,omitempty"`
	Threshold      *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s AlertRuleConditionTriggersExpressionConditions) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleConditionTriggersExpressionConditions) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionTriggersExpressionConditions) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *AlertRuleConditionTriggersExpressionConditions) GetOperator() *string {
	return s.Operator
}

func (s *AlertRuleConditionTriggersExpressionConditions) GetQueryName() *string {
	return s.QueryName
}

func (s *AlertRuleConditionTriggersExpressionConditions) GetThreshold() *float64 {
	return s.Threshold
}

func (s *AlertRuleConditionTriggersExpressionConditions) SetExpressionType(v string) *AlertRuleConditionTriggersExpressionConditions {
	s.ExpressionType = &v
	return s
}

func (s *AlertRuleConditionTriggersExpressionConditions) SetOperator(v string) *AlertRuleConditionTriggersExpressionConditions {
	s.Operator = &v
	return s
}

func (s *AlertRuleConditionTriggersExpressionConditions) SetQueryName(v string) *AlertRuleConditionTriggersExpressionConditions {
	s.QueryName = &v
	return s
}

func (s *AlertRuleConditionTriggersExpressionConditions) SetThreshold(v float64) *AlertRuleConditionTriggersExpressionConditions {
	s.Threshold = &v
	return s
}

func (s *AlertRuleConditionTriggersExpressionConditions) Validate() error {
	return dara.Validate(s)
}
