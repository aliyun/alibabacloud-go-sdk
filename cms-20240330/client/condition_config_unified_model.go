// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConditionConfigUnified interface {
	dara.Model
	String() string
	GoString() string
	SetAbsDeviation(v float64) *ConditionConfigUnified
	GetAbsDeviation() *float64
	SetAggregate(v string) *ConditionConfigUnified
	GetAggregate() *string
	SetAlertCount(v int32) *ConditionConfigUnified
	GetAlertCount() *int32
	SetBaselinePeriod(v string) *ConditionConfigUnified
	GetBaselinePeriod() *string
	SetCompareList(v []*CompareList) *ConditionConfigUnified
	GetCompareList() []*CompareList
	SetCompositeEscalation(v *CloudMonitoringCompositeEscalation) *ConditionConfigUnified
	GetCompositeEscalation() *CloudMonitoringCompositeEscalation
	SetCountOperator(v string) *ConditionConfigUnified
	GetCountOperator() *string
	SetCountThreshold(v int64) *ConditionConfigUnified
	GetCountThreshold() *int64
	SetDurationSecs(v int32) *ConditionConfigUnified
	GetDurationSecs() *int32
	SetEnableSeveritySuppression(v bool) *ConditionConfigUnified
	GetEnableSeveritySuppression() *bool
	SetEscalationType(v string) *ConditionConfigUnified
	GetEscalationType() *string
	SetExpressEscalation(v *CloudMonitoringExpressEscalation) *ConditionConfigUnified
	GetExpressEscalation() *CloudMonitoringExpressEscalation
	SetLegacyRaw(v string) *ConditionConfigUnified
	GetLegacyRaw() *string
	SetLegacyType(v string) *ConditionConfigUnified
	GetLegacyType() *string
	SetMatchField(v string) *ConditionConfigUnified
	GetMatchField() *string
	SetMatchOperator(v string) *ConditionConfigUnified
	GetMatchOperator() *string
	SetMatchValue(v string) *ConditionConfigUnified
	GetMatchValue() *string
	SetMax(v float64) *ConditionConfigUnified
	GetMax() *float64
	SetMin(v float64) *ConditionConfigUnified
	GetMin() *float64
	SetNoDataAlertLevel(v string) *ConditionConfigUnified
	GetNoDataAlertLevel() *string
	SetNoDataAlertSeverity(v string) *ConditionConfigUnified
	GetNoDataAlertSeverity() *string
	SetNoDataAppendValue(v float64) *ConditionConfigUnified
	GetNoDataAppendValue() *float64
	SetNoDataDurationSecs(v int32) *ConditionConfigUnified
	GetNoDataDurationSecs() *int32
	SetNoDataPolicy(v string) *ConditionConfigUnified
	GetNoDataPolicy() *string
	SetOperator(v string) *ConditionConfigUnified
	GetOperator() *string
	SetPrometheus(v *CloudMonitoringPrometheusEscalation) *ConditionConfigUnified
	GetPrometheus() *CloudMonitoringPrometheusEscalation
	SetRelation(v string) *ConditionConfigUnified
	GetRelation() *string
	SetSensitivity(v string) *ConditionConfigUnified
	GetSensitivity() *string
	SetSeverity(v string) *ConditionConfigUnified
	GetSeverity() *string
	SetSimpleEscalation(v *CloudMonitoringSimpleEscalation) *ConditionConfigUnified
	GetSimpleEscalation() *CloudMonitoringSimpleEscalation
	SetThreshold(v float64) *ConditionConfigUnified
	GetThreshold() *float64
	SetThresholdList(v []*ThresholdList) *ConditionConfigUnified
	GetThresholdList() []*ThresholdList
	SetTriggers(v []*Triggers) *ConditionConfigUnified
	GetTriggers() []*Triggers
	SetType(v string) *ConditionConfigUnified
	GetType() *string
	SetYoyTimeUnit(v string) *ConditionConfigUnified
	GetYoyTimeUnit() *string
	SetYoyTimeValue(v int32) *ConditionConfigUnified
	GetYoyTimeValue() *int32
}

type ConditionConfigUnified struct {
	// The minimum deviation or absolute deviation dead zone for dynamic baselines. This parameter takes effect only for baseline operators in UMODEL_METRICSET_CONDITION and APM_SIMPLE_CONDITION. If the absolute difference between the current value and the boundary is less than the absDeviation value, no alert is triggered. The unit is the same as the metric unit. The value must be greater than or equal to 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 0.0
	AbsDeviation *float64 `json:"absDeviation,omitempty" xml:"absDeviation,omitempty"`
	// The aggregation function used for APM_SIMPLE_CONDITION. For UModel conditions, this semantic has been migrated to QueryConfigUnified, and this field no longer takes effect.
	//
	// example:
	//
	// AVG
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// The consecutive trigger count threshold for SLS_MULTI_CONDITION. An alert is triggered only after the condition is met for the specified number of consecutive times. Default value: 1.
	//
	// example:
	//
	// 1
	AlertCount *int32 `json:"alertCount,omitempty" xml:"alertCount,omitempty"`
	// The baseline period. This parameter takes effect only for baseline operators. Valid values: AUTO (automatic identification), DAILY (daily), WEEKLY (weekly), and NONE (no period). For weekly baselines, the backend automatically extends the historical training window to 14 days or more. Automatic identification cannot echo specific identification conclusions.
	//
	// example:
	//
	// AUTO
	BaselinePeriod *string `json:"baselinePeriod,omitempty" xml:"baselinePeriod,omitempty"`
	// The list of multiple comparisons for APM_COMPOSITE_CONDITION.
	CompareList []*CompareList `json:"compareList,omitempty" xml:"compareList,omitempty" type:"Repeated"`
	// The multi-metric composite trigger configuration for CLOUD_MONITORING_CONDITION when escalationType is set to COMPOSITE. The relation, severity, times, and escalations parameters are required.
	CompositeEscalation *CloudMonitoringCompositeEscalation `json:"compositeEscalation,omitempty" xml:"compositeEscalation,omitempty"`
	// The count comparison operator for UMODEL_LOGSET_CONDITION.
	//
	// example:
	//
	// GTE
	CountOperator *string `json:"countOperator,omitempty" xml:"countOperator,omitempty"`
	// The count threshold for UMODEL_LOGSET_CONDITION.
	//
	// example:
	//
	// 100
	CountThreshold *int64 `json:"countThreshold,omitempty" xml:"countThreshold,omitempty"`
	// The duration in seconds. This parameter is directly used for PROMETHEUS_SIMPLE, UMODEL_METRICSET_CONDITION, and UMODEL_LOGSET_CONDITION. For UMODEL_METRICSET_MULTI_CONDITION, it serves as the global default value, and each trigger can override it with its own durationSecs value.
	//
	// example:
	//
	// 60
	DurationSecs *int32 `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	// Specifies whether to enable severity suppression for UMODEL_METRICSET_MULTI_CONDITION and PROMETHEUS_MULTI_CONDITION. Default value: true. When enabled, only the highest severity trigger is reported for the same entity.
	//
	// example:
	//
	// true
	EnableSeveritySuppression *bool `json:"enableSeveritySuppression,omitempty" xml:"enableSeveritySuppression,omitempty"`
	// The expression type for CLOUD_MONITORING_CONDITION. Valid values: SIMPLE, COMPOSITE, EXPRESS, and PROMETHEUS. The write path only supports SIMPLE and COMPOSITE. Specify the corresponding escalation sub-object based on the type.
	//
	// example:
	//
	// SIMPLE
	EscalationType *string `json:"escalationType,omitempty" xml:"escalationType,omitempty"`
	// The expression trigger configuration for CLOUD_MONITORING_CONDITION when escalationType is set to EXPRESS. This parameter is output only in the read path.
	ExpressEscalation *CloudMonitoringExpressEscalation `json:"expressEscalation,omitempty" xml:"expressEscalation,omitempty"`
	// The original V1 condition JSON string returned as a fallback when parsing fails in the read path and the type is UNKNOWN_CONDITION. The frontend displays this field as read-only when it is not empty.
	//
	// example:
	//
	// 示例值
	LegacyRaw *string `json:"legacyRaw,omitempty" xml:"legacyRaw,omitempty"`
	// The legacy type returned when the type is UNKNOWN_CONDITION. This indicates that the rule cannot be edited by using the new API. Submit a ticket to contact the CloudMonitor product team.
	//
	// example:
	//
	// default
	LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
	// The log field name used for UMODEL_LOGSET_CONDITION when matchOperator is set to CONTAINS, EQUALS, or REGEX.
	//
	// example:
	//
	// 示例值
	MatchField *string `json:"matchField,omitempty" xml:"matchField,omitempty"`
	// The log match operator for UMODEL_LOGSET_CONDITION.
	//
	// example:
	//
	// PRESENT
	MatchOperator *string `json:"matchOperator,omitempty" xml:"matchOperator,omitempty"`
	// The log match value used for UMODEL_LOGSET_CONDITION when matchOperator is set to CONTAINS, EQUALS, or REGEX.
	//
	// example:
	//
	// 示例值
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
	// The upper bound of the range used for UMODEL_METRICSET_CONDITION when operator is set to IN_RANGE or OUT_OF_RANGE.
	//
	// example:
	//
	// 1.0
	Max *float64 `json:"max,omitempty" xml:"max,omitempty"`
	// The lower bound of the range used for UMODEL_METRICSET_CONDITION when operator is set to IN_RANGE or OUT_OF_RANGE.
	//
	// example:
	//
	// 1.0
	Min *float64 `json:"min,omitempty" xml:"min,omitempty"`
	// The no-data alert level for SLS_MULTI_CONDITION. For APM and Prometheus conditions, this parameter has been migrated to noDataPolicy and noDataAlertSeverity.
	//
	// example:
	//
	// INFO
	NoDataAlertLevel *string `json:"noDataAlertLevel,omitempty" xml:"noDataAlertLevel,omitempty"`
	// The no-data alert severity. This parameter takes effect for PROMETHEUS_SIMPLE_CONDITION, PROMETHEUS_MULTI_CONDITION, UMODEL_METRICSET_CONDITION, and UMODEL_METRICSET_MULTI_CONDITION when noDataPolicy is set to NO_DATA_TO_ALERT. For UMODEL_METRICSET conditions, this field is required when noDataPolicy is set to NO_DATA_TO_ALERT, and must not be specified for other policies. SLS_MULTI_CONDITION still uses noDataAlertLevel.
	//
	// example:
	//
	// INFO
	NoDataAlertSeverity *string `json:"noDataAlertSeverity,omitempty" xml:"noDataAlertSeverity,omitempty"`
	// The compensated value when no data is available for APM_SIMPLE_CONDITION and APM_COMPOSITE_CONDITION. This parameter is nullable.
	//
	// example:
	//
	// 1.0
	NoDataAppendValue *float64 `json:"noDataAppendValue,omitempty" xml:"noDataAppendValue,omitempty"`
	// The tolerance duration in seconds before no data is determined. This parameter takes effect for UMODEL_METRICSET_CONDITION and UMODEL_METRICSET_MULTI_CONDITION when noDataPolicy is not empty. The noDataPolicy is applied only after no data is continuously reported for longer than this duration. This prevents empty query results in a single cycle caused by UModel data write jitter. The value must be a positive integer. If not specified, the backend uses the fallback value of max(collection interval × 2, 300).
	//
	// example:
	//
	// 600
	NoDataDurationSecs *int32 `json:"noDataDurationSecs,omitempty" xml:"noDataDurationSecs,omitempty"`
	// The no-data processing policy for CLOUD_MONITORING_CONDITION, PROMETHEUS_MULTI_CONDITION, PROMETHEUS_SIMPLE_CONDITION, APM_SIMPLE_CONDITION, APM_COMPOSITE_CONDITION, UMODEL_METRICSET_CONDITION, and UMODEL_METRICSET_MULTI_CONDITION. Valid values: NO_DATA_TO_OK (treats no data as normal and recovers existing alerts), NO_DATA_TO_ALERT (triggers an alert for no data), KEEP_LAST_STATE (maintains the last state before no data), and APPEND_VALUE (APM only). If not specified, no-data determination is disabled, and the behavior remains the same as before this feature was released. UMODEL_METRICSET_CONDITION and UMODEL_METRICSET_MULTI_CONDITION do not support APPEND_VALUE because their data sources do not support backfilling.
	//
	// example:
	//
	// 示例值
	NoDataPolicy *string `json:"noDataPolicy,omitempty" xml:"noDataPolicy,omitempty"`
	// The comparison operator. For UMODEL_METRICSET_CONDITION, valid values are GT (greater than), GE (greater than or equal to), LT (less than), LE (less than or equal to), EQ (equal to), NE (not equal to), IN_RANGE (within the range, requires min and max), OUT_OF_RANGE (outside the range, requires min and max), PRESENT (field exists), NOT_PRESENT (field does not exist), ABOVE_UPPER (dynamic baseline spike), BELOW_LOWER (dynamic baseline drop), and OUT_OF_BAND (bidirectional dynamic baseline, requires sensitivity, do not specify threshold, min, or max). This field is not used for UMODEL_LOGSET_CONDITION. For APM_SIMPLE_CONDITION, valid values are GT, GTE, LT, LTE, EQ, NE, YOY_UP, YOY_DOWN (YOY operators require yoyTimeUnit and yoyTimeValue), ABOVE_UPPER, BELOW_LOWER, and OUT_OF_BAND (dynamic baseline, requires sensitivity).
	//
	// example:
	//
	// GT
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The PromQL trigger configuration for CLOUD_MONITORING_CONDITION when escalationType is set to PROMETHEUS. This parameter is output only in the read path.
	Prometheus *CloudMonitoringPrometheusEscalation `json:"prometheus,omitempty" xml:"prometheus,omitempty"`
	// The logical relation between conditions for APM_COMPOSITE_CONDITION.
	//
	// example:
	//
	// AND
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// The dynamic baseline sensitivity for UMODEL_METRICSET_CONDITION and APM_SIMPLE_CONDITION. This parameter takes effect only when operator is set to ABOVE_UPPER, BELOW_LOWER, or OUT_OF_BAND. Valid values: HIGH (narrowest band and most sensitive), MEDIUM, and LOW (widest band and least sensitive).
	//
	// example:
	//
	// MEDIUM
	Sensitivity *string `json:"sensitivity,omitempty" xml:"sensitivity,omitempty"`
	// The severity level for UMODEL_METRICSET_CONDITION, UMODEL_LOGSET_CONDITION, PROMETHEUS_SIMPLE, and APM_COMPOSITE.
	//
	// example:
	//
	// INFO
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The single-metric multi-level trigger configuration for CLOUD_MONITORING_CONDITION when escalationType is set to SIMPLE. The metricName, period, and escalations parameters are required.
	SimpleEscalation *CloudMonitoringSimpleEscalation `json:"simpleEscalation,omitempty" xml:"simpleEscalation,omitempty"`
	// The threshold used for UMODEL_METRICSET_CONDITION when a non-range operator is specified.
	//
	// example:
	//
	// 30
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// The list of multiple thresholds for APM_SIMPLE_CONDITION.
	ThresholdList []*ThresholdList `json:"thresholdList,omitempty" xml:"thresholdList,omitempty" type:"Repeated"`
	// The list of triggers. This field is polymorphic based on the type. CLOUD_MONITORING_CONDITION does not use this field; use simpleEscalation.escalations or compositeEscalation.escalations instead. For SLS_MULTI_CONDITION, each case includes matchField, matchOperator, matchValue, countOperator, countThreshold, and severity, and at least one case is required. For UMODEL_METRICSET_MULTI_CONDITION, each trigger includes severity, durationSecs, and an expression (SIMPLE or COMPOSITE). For PROMETHEUS_MULTI_CONDITION, each trigger includes severity, durationSecs, and an expression (SIMPLE or COMPOSITE). Triggers are sorted by severity priority, and the first matched trigger is executed.
	Triggers []*Triggers `json:"triggers,omitempty" xml:"triggers,omitempty" type:"Repeated"`
	// The detection condition type. Valid values and required fields: PROMETHEUS_SIMPLE_CONDITION (requires operator, threshold, durationSecs, and severity), UMODEL_METRICSET_CONDITION (requires operator, durationSecs, and severity; requires threshold for non-range operators; requires min and max when operator is IN_RANGE or OUT_OF_RANGE; noDataPolicy and noDataDurationSecs are optional, but noDataAlertSeverity is required when noDataPolicy is NO_DATA_TO_ALERT), UMODEL_LOGSET_CONDITION (requires matchOperator, durationSecs, and severity; requires matchField and matchValue when matchOperator is CONTAINS, EQUALS, or REGEX; countOperator and countThreshold are optional), UMODEL_METRICSET_MULTI_CONDITION (requires triggers; durationSecs as global default and enableSeveritySuppression are optional; noDataPolicy and noDataDurationSecs are optional, but noDataAlertSeverity is required when noDataPolicy is NO_DATA_TO_ALERT), APM_SIMPLE_CONDITION (requires operator and aggregate; requires thresholdList or threshold; requires yoyTimeUnit and yoyTimeValue when operator is YOY_UP or YOY_DOWN; noDataPolicy and noDataAppendValue are optional), APM_COMPOSITE_CONDITION (requires compareList, relation, and severity; noDataPolicy and noDataAppendValue are optional), CLOUD_MONITORING_CONDITION (requires escalationType; requires simpleEscalation when escalationType is SIMPLE; requires compositeEscalation when escalationType is COMPOSITE; noDataPolicy is optional), and UNKNOWN_CONDITION (read-only fallback, do not use in the write path). Non-enum values such as SLS_CONDITION or CMS_BASIC_CONDITION are prohibited. The backend returns an Invalidtype 400 error for invalid values.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROMETHEUS_SIMPLE_CONDITION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The year-over-year time unit for APM_SIMPLE_CONDITION. This parameter takes effect only when operator is set to YOY_UP or YOY_DOWN.
	//
	// example:
	//
	// minute
	YoyTimeUnit *string `json:"yoyTimeUnit,omitempty" xml:"yoyTimeUnit,omitempty"`
	// The year-over-year time value for APM_SIMPLE_CONDITION. This parameter takes effect only when operator is set to YOY_UP or YOY_DOWN.
	//
	// example:
	//
	// 1
	YoyTimeValue *int32 `json:"yoyTimeValue,omitempty" xml:"yoyTimeValue,omitempty"`
}

func (s ConditionConfigUnified) String() string {
	return dara.Prettify(s)
}

func (s ConditionConfigUnified) GoString() string {
	return s.String()
}

func (s *ConditionConfigUnified) GetAbsDeviation() *float64 {
	return s.AbsDeviation
}

func (s *ConditionConfigUnified) GetAggregate() *string {
	return s.Aggregate
}

func (s *ConditionConfigUnified) GetAlertCount() *int32 {
	return s.AlertCount
}

func (s *ConditionConfigUnified) GetBaselinePeriod() *string {
	return s.BaselinePeriod
}

func (s *ConditionConfigUnified) GetCompareList() []*CompareList {
	return s.CompareList
}

func (s *ConditionConfigUnified) GetCompositeEscalation() *CloudMonitoringCompositeEscalation {
	return s.CompositeEscalation
}

func (s *ConditionConfigUnified) GetCountOperator() *string {
	return s.CountOperator
}

func (s *ConditionConfigUnified) GetCountThreshold() *int64 {
	return s.CountThreshold
}

func (s *ConditionConfigUnified) GetDurationSecs() *int32 {
	return s.DurationSecs
}

func (s *ConditionConfigUnified) GetEnableSeveritySuppression() *bool {
	return s.EnableSeveritySuppression
}

func (s *ConditionConfigUnified) GetEscalationType() *string {
	return s.EscalationType
}

func (s *ConditionConfigUnified) GetExpressEscalation() *CloudMonitoringExpressEscalation {
	return s.ExpressEscalation
}

func (s *ConditionConfigUnified) GetLegacyRaw() *string {
	return s.LegacyRaw
}

func (s *ConditionConfigUnified) GetLegacyType() *string {
	return s.LegacyType
}

func (s *ConditionConfigUnified) GetMatchField() *string {
	return s.MatchField
}

func (s *ConditionConfigUnified) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *ConditionConfigUnified) GetMatchValue() *string {
	return s.MatchValue
}

func (s *ConditionConfigUnified) GetMax() *float64 {
	return s.Max
}

func (s *ConditionConfigUnified) GetMin() *float64 {
	return s.Min
}

func (s *ConditionConfigUnified) GetNoDataAlertLevel() *string {
	return s.NoDataAlertLevel
}

func (s *ConditionConfigUnified) GetNoDataAlertSeverity() *string {
	return s.NoDataAlertSeverity
}

func (s *ConditionConfigUnified) GetNoDataAppendValue() *float64 {
	return s.NoDataAppendValue
}

func (s *ConditionConfigUnified) GetNoDataDurationSecs() *int32 {
	return s.NoDataDurationSecs
}

func (s *ConditionConfigUnified) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *ConditionConfigUnified) GetOperator() *string {
	return s.Operator
}

func (s *ConditionConfigUnified) GetPrometheus() *CloudMonitoringPrometheusEscalation {
	return s.Prometheus
}

func (s *ConditionConfigUnified) GetRelation() *string {
	return s.Relation
}

func (s *ConditionConfigUnified) GetSensitivity() *string {
	return s.Sensitivity
}

func (s *ConditionConfigUnified) GetSeverity() *string {
	return s.Severity
}

func (s *ConditionConfigUnified) GetSimpleEscalation() *CloudMonitoringSimpleEscalation {
	return s.SimpleEscalation
}

func (s *ConditionConfigUnified) GetThreshold() *float64 {
	return s.Threshold
}

func (s *ConditionConfigUnified) GetThresholdList() []*ThresholdList {
	return s.ThresholdList
}

func (s *ConditionConfigUnified) GetTriggers() []*Triggers {
	return s.Triggers
}

func (s *ConditionConfigUnified) GetType() *string {
	return s.Type
}

func (s *ConditionConfigUnified) GetYoyTimeUnit() *string {
	return s.YoyTimeUnit
}

func (s *ConditionConfigUnified) GetYoyTimeValue() *int32 {
	return s.YoyTimeValue
}

func (s *ConditionConfigUnified) SetAbsDeviation(v float64) *ConditionConfigUnified {
	s.AbsDeviation = &v
	return s
}

func (s *ConditionConfigUnified) SetAggregate(v string) *ConditionConfigUnified {
	s.Aggregate = &v
	return s
}

func (s *ConditionConfigUnified) SetAlertCount(v int32) *ConditionConfigUnified {
	s.AlertCount = &v
	return s
}

func (s *ConditionConfigUnified) SetBaselinePeriod(v string) *ConditionConfigUnified {
	s.BaselinePeriod = &v
	return s
}

func (s *ConditionConfigUnified) SetCompareList(v []*CompareList) *ConditionConfigUnified {
	s.CompareList = v
	return s
}

func (s *ConditionConfigUnified) SetCompositeEscalation(v *CloudMonitoringCompositeEscalation) *ConditionConfigUnified {
	s.CompositeEscalation = v
	return s
}

func (s *ConditionConfigUnified) SetCountOperator(v string) *ConditionConfigUnified {
	s.CountOperator = &v
	return s
}

func (s *ConditionConfigUnified) SetCountThreshold(v int64) *ConditionConfigUnified {
	s.CountThreshold = &v
	return s
}

func (s *ConditionConfigUnified) SetDurationSecs(v int32) *ConditionConfigUnified {
	s.DurationSecs = &v
	return s
}

func (s *ConditionConfigUnified) SetEnableSeveritySuppression(v bool) *ConditionConfigUnified {
	s.EnableSeveritySuppression = &v
	return s
}

func (s *ConditionConfigUnified) SetEscalationType(v string) *ConditionConfigUnified {
	s.EscalationType = &v
	return s
}

func (s *ConditionConfigUnified) SetExpressEscalation(v *CloudMonitoringExpressEscalation) *ConditionConfigUnified {
	s.ExpressEscalation = v
	return s
}

func (s *ConditionConfigUnified) SetLegacyRaw(v string) *ConditionConfigUnified {
	s.LegacyRaw = &v
	return s
}

func (s *ConditionConfigUnified) SetLegacyType(v string) *ConditionConfigUnified {
	s.LegacyType = &v
	return s
}

func (s *ConditionConfigUnified) SetMatchField(v string) *ConditionConfigUnified {
	s.MatchField = &v
	return s
}

func (s *ConditionConfigUnified) SetMatchOperator(v string) *ConditionConfigUnified {
	s.MatchOperator = &v
	return s
}

func (s *ConditionConfigUnified) SetMatchValue(v string) *ConditionConfigUnified {
	s.MatchValue = &v
	return s
}

func (s *ConditionConfigUnified) SetMax(v float64) *ConditionConfigUnified {
	s.Max = &v
	return s
}

func (s *ConditionConfigUnified) SetMin(v float64) *ConditionConfigUnified {
	s.Min = &v
	return s
}

func (s *ConditionConfigUnified) SetNoDataAlertLevel(v string) *ConditionConfigUnified {
	s.NoDataAlertLevel = &v
	return s
}

func (s *ConditionConfigUnified) SetNoDataAlertSeverity(v string) *ConditionConfigUnified {
	s.NoDataAlertSeverity = &v
	return s
}

func (s *ConditionConfigUnified) SetNoDataAppendValue(v float64) *ConditionConfigUnified {
	s.NoDataAppendValue = &v
	return s
}

func (s *ConditionConfigUnified) SetNoDataDurationSecs(v int32) *ConditionConfigUnified {
	s.NoDataDurationSecs = &v
	return s
}

func (s *ConditionConfigUnified) SetNoDataPolicy(v string) *ConditionConfigUnified {
	s.NoDataPolicy = &v
	return s
}

func (s *ConditionConfigUnified) SetOperator(v string) *ConditionConfigUnified {
	s.Operator = &v
	return s
}

func (s *ConditionConfigUnified) SetPrometheus(v *CloudMonitoringPrometheusEscalation) *ConditionConfigUnified {
	s.Prometheus = v
	return s
}

func (s *ConditionConfigUnified) SetRelation(v string) *ConditionConfigUnified {
	s.Relation = &v
	return s
}

func (s *ConditionConfigUnified) SetSensitivity(v string) *ConditionConfigUnified {
	s.Sensitivity = &v
	return s
}

func (s *ConditionConfigUnified) SetSeverity(v string) *ConditionConfigUnified {
	s.Severity = &v
	return s
}

func (s *ConditionConfigUnified) SetSimpleEscalation(v *CloudMonitoringSimpleEscalation) *ConditionConfigUnified {
	s.SimpleEscalation = v
	return s
}

func (s *ConditionConfigUnified) SetThreshold(v float64) *ConditionConfigUnified {
	s.Threshold = &v
	return s
}

func (s *ConditionConfigUnified) SetThresholdList(v []*ThresholdList) *ConditionConfigUnified {
	s.ThresholdList = v
	return s
}

func (s *ConditionConfigUnified) SetTriggers(v []*Triggers) *ConditionConfigUnified {
	s.Triggers = v
	return s
}

func (s *ConditionConfigUnified) SetType(v string) *ConditionConfigUnified {
	s.Type = &v
	return s
}

func (s *ConditionConfigUnified) SetYoyTimeUnit(v string) *ConditionConfigUnified {
	s.YoyTimeUnit = &v
	return s
}

func (s *ConditionConfigUnified) SetYoyTimeValue(v int32) *ConditionConfigUnified {
	s.YoyTimeValue = &v
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
	if s.Prometheus != nil {
		if err := s.Prometheus.Validate(); err != nil {
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
