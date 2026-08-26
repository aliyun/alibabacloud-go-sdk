// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggers interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *Triggers
	GetComparisonOperator() *string
	SetCondition(v string) *Triggers
	GetCondition() *string
	SetConditions(v []*TriggerConditions) *Triggers
	GetConditions() []*TriggerConditions
	SetCountCondition(v string) *Triggers
	GetCountCondition() *string
	SetCountOperator(v string) *Triggers
	GetCountOperator() *string
	SetCountThreshold(v int64) *Triggers
	GetCountThreshold() *int64
	SetDurationSecs(v int32) *Triggers
	GetDurationSecs() *int32
	SetExpressionType(v string) *Triggers
	GetExpressionType() *string
	SetLogicOperator(v string) *Triggers
	GetLogicOperator() *string
	SetMatchField(v string) *Triggers
	GetMatchField() *string
	SetMatchOperator(v string) *Triggers
	GetMatchOperator() *string
	SetMatchValue(v string) *Triggers
	GetMatchValue() *string
	SetMax(v float64) *Triggers
	GetMax() *float64
	SetMetricName(v string) *Triggers
	GetMetricName() *string
	SetMin(v float64) *Triggers
	GetMin() *float64
	SetOperator(v string) *Triggers
	GetOperator() *string
	SetPeriod(v int32) *Triggers
	GetPeriod() *int32
	SetPreCondition(v string) *Triggers
	GetPreCondition() *string
	SetQueryName(v string) *Triggers
	GetQueryName() *string
	SetSeverity(v string) *Triggers
	GetSeverity() *string
	SetStatistics(v string) *Triggers
	GetStatistics() *string
	SetThreshold(v interface{}) *Triggers
	GetThreshold() interface{}
	SetTimes(v int32) *Triggers
	GetTimes() *int32
}

type Triggers struct {
	// The comparison operator. This parameter applies to CLOUD_MONITORING_CONDITION.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	// The match expression for SLS_MULTI_CONDITION. This corresponds to the V1 condition parameter and is preserved as-is without parsing.
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The list of sub-conditions. This parameter applies to UMODEL_METRICSET_MULTI and PROMETHEUS_MULTI with expressionType=COMPOSITE. Each item contains queryName, operator, and threshold.
	Conditions []*TriggerConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The count match expression for SLS_MULTI_CONDITION. This corresponds to the V1 countCondition parameter and is preserved as-is without parsing.
	CountCondition *string `json:"countCondition,omitempty" xml:"countCondition,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The SLS_MULTI_CONDITION write path is disabled. Use the countCondition parameter instead.
	CountOperator *string `json:"countOperator,omitempty" xml:"countOperator,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The SLS_MULTI_CONDITION write path is disabled. Use the countCondition parameter instead.
	CountThreshold *int64 `json:"countThreshold,omitempty" xml:"countThreshold,omitempty"`
	// The duration in seconds for which data must continuously meet the condition to trigger an alert. If not specified, the value is inherited from conditionConfig.durationSecs. This parameter is used by UMODEL_METRICSET_MULTI_CONDITION and PROMETHEUS_MULTI_CONDITION.
	DurationSecs *int32 `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	// The expression type. Valid values: SIMPLE and COMPOSITE. This parameter applies to UMODEL_METRICSET_MULTI_CONDITION and PROMETHEUS_MULTI_CONDITION.
	ExpressionType *string `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	// The logical operator. This parameter applies to UMODEL_METRICSET_MULTI and PROMETHEUS_MULTI with expressionType=COMPOSITE. Valid values: AND, OR, and UNLESS.
	LogicOperator *string `json:"logicOperator,omitempty" xml:"logicOperator,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The SLS_MULTI_CONDITION write path is disabled. Use the condition parameter instead.
	MatchField *string `json:"matchField,omitempty" xml:"matchField,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The SLS_MULTI_CONDITION write path is disabled. Use the condition parameter instead.
	MatchOperator *string `json:"matchOperator,omitempty" xml:"matchOperator,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The SLS_MULTI_CONDITION write path is disabled. Use the condition parameter instead.
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
	// The upper bound of the range. This parameter applies to UMODEL_METRICSET_MULTI with expressionType=SIMPLE. This parameter is required when operator is set to IN_RANGE or OUT_OF_RANGE. The value must be greater than or equal to min.
	Max *float64 `json:"max,omitempty" xml:"max,omitempty"`
	// The metric name. This parameter applies to CLOUD_MONITORING_CONDITION with expressionType=COMPOSITE. For SIMPLE, the metric name is specified at the conditionConfig level by the metricName parameter.
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The lower bound of the range. This parameter applies to UMODEL_METRICSET_MULTI with expressionType=SIMPLE. This parameter is required when operator is set to IN_RANGE or OUT_OF_RANGE.
	Min *float64 `json:"min,omitempty" xml:"min,omitempty"`
	// The operator. For UMODEL_METRICSET_MULTI and PROMETHEUS_MULTI with expressionType=SIMPLE, this is a comparison operator. Valid values: GT, GE, LT, LE, EQ, NE, IN_RANGE, OUT_OF_RANGE, PRESENT, and NOT_PRESENT. For SLS_MULTI_CONDITION, this is aligned with V1 caseList.type. Valid values: HAS_DATA, HAS_DATA_COUNT, HAS_DATA_MATCH, and HAS_DATA_MATCH_COUNT.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The aggregation period in seconds. This parameter applies to CLOUD_MONITORING_CONDITION with expressionType=COMPOSITE. For SIMPLE, the period is specified at the conditionConfig level by the period parameter.
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The precondition. This parameter applies to CLOUD_MONITORING_CONDITION.
	PreCondition *string `json:"preCondition,omitempty" xml:"preCondition,omitempty"`
	// The referenced query name. This parameter applies to UMODEL_METRICSET_MULTI and PROMETHEUS_MULTI with expressionType=SIMPLE. The value corresponds to QueryConfigUnified.queries[].name.
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// The severity level. Priority order: CRITICAL > ERROR > WARN / WARNING > INFO. Multiple triggers are sorted by this priority, and the first match triggers the alert. This parameter takes effect when the type is SLS_MULTI_CONDITION or CLOUD_MONITORING_CONDITION with expressionType=SIMPLE.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The statistical method. This parameter applies to CLOUD_MONITORING_CONDITION.
	Statistics *string `json:"statistics,omitempty" xml:"statistics,omitempty"`
	// The threshold value. For CLOUD_MONITORING_CONDITION, this is a string. For UMODEL_METRICSET_MULTI and PROMETHEUS_MULTI, this is a numeric value.
	Threshold interface{} `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// The number of consecutive times the condition must be met to trigger the alert. Each entry has its own setting. This parameter applies to CLOUD_MONITORING_CONDITION with expressionType=SIMPLE.
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s Triggers) String() string {
	return dara.Prettify(s)
}

func (s Triggers) GoString() string {
	return s.String()
}

func (s *Triggers) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *Triggers) GetCondition() *string {
	return s.Condition
}

func (s *Triggers) GetConditions() []*TriggerConditions {
	return s.Conditions
}

func (s *Triggers) GetCountCondition() *string {
	return s.CountCondition
}

func (s *Triggers) GetCountOperator() *string {
	return s.CountOperator
}

func (s *Triggers) GetCountThreshold() *int64 {
	return s.CountThreshold
}

func (s *Triggers) GetDurationSecs() *int32 {
	return s.DurationSecs
}

func (s *Triggers) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *Triggers) GetLogicOperator() *string {
	return s.LogicOperator
}

func (s *Triggers) GetMatchField() *string {
	return s.MatchField
}

func (s *Triggers) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *Triggers) GetMatchValue() *string {
	return s.MatchValue
}

func (s *Triggers) GetMax() *float64 {
	return s.Max
}

func (s *Triggers) GetMetricName() *string {
	return s.MetricName
}

func (s *Triggers) GetMin() *float64 {
	return s.Min
}

func (s *Triggers) GetOperator() *string {
	return s.Operator
}

func (s *Triggers) GetPeriod() *int32 {
	return s.Period
}

func (s *Triggers) GetPreCondition() *string {
	return s.PreCondition
}

func (s *Triggers) GetQueryName() *string {
	return s.QueryName
}

func (s *Triggers) GetSeverity() *string {
	return s.Severity
}

func (s *Triggers) GetStatistics() *string {
	return s.Statistics
}

func (s *Triggers) GetThreshold() interface{} {
	return s.Threshold
}

func (s *Triggers) GetTimes() *int32 {
	return s.Times
}

func (s *Triggers) SetComparisonOperator(v string) *Triggers {
	s.ComparisonOperator = &v
	return s
}

func (s *Triggers) SetCondition(v string) *Triggers {
	s.Condition = &v
	return s
}

func (s *Triggers) SetConditions(v []*TriggerConditions) *Triggers {
	s.Conditions = v
	return s
}

func (s *Triggers) SetCountCondition(v string) *Triggers {
	s.CountCondition = &v
	return s
}

func (s *Triggers) SetCountOperator(v string) *Triggers {
	s.CountOperator = &v
	return s
}

func (s *Triggers) SetCountThreshold(v int64) *Triggers {
	s.CountThreshold = &v
	return s
}

func (s *Triggers) SetDurationSecs(v int32) *Triggers {
	s.DurationSecs = &v
	return s
}

func (s *Triggers) SetExpressionType(v string) *Triggers {
	s.ExpressionType = &v
	return s
}

func (s *Triggers) SetLogicOperator(v string) *Triggers {
	s.LogicOperator = &v
	return s
}

func (s *Triggers) SetMatchField(v string) *Triggers {
	s.MatchField = &v
	return s
}

func (s *Triggers) SetMatchOperator(v string) *Triggers {
	s.MatchOperator = &v
	return s
}

func (s *Triggers) SetMatchValue(v string) *Triggers {
	s.MatchValue = &v
	return s
}

func (s *Triggers) SetMax(v float64) *Triggers {
	s.Max = &v
	return s
}

func (s *Triggers) SetMetricName(v string) *Triggers {
	s.MetricName = &v
	return s
}

func (s *Triggers) SetMin(v float64) *Triggers {
	s.Min = &v
	return s
}

func (s *Triggers) SetOperator(v string) *Triggers {
	s.Operator = &v
	return s
}

func (s *Triggers) SetPeriod(v int32) *Triggers {
	s.Period = &v
	return s
}

func (s *Triggers) SetPreCondition(v string) *Triggers {
	s.PreCondition = &v
	return s
}

func (s *Triggers) SetQueryName(v string) *Triggers {
	s.QueryName = &v
	return s
}

func (s *Triggers) SetSeverity(v string) *Triggers {
	s.Severity = &v
	return s
}

func (s *Triggers) SetStatistics(v string) *Triggers {
	s.Statistics = &v
	return s
}

func (s *Triggers) SetThreshold(v interface{}) *Triggers {
	s.Threshold = v
	return s
}

func (s *Triggers) SetTimes(v int32) *Triggers {
	s.Times = &v
	return s
}

func (s *Triggers) Validate() error {
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
