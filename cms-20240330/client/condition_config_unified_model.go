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
	SetNoDataPolicy(v string) *ConditionConfigUnified
	GetNoDataPolicy() *string
	SetOperator(v string) *ConditionConfigUnified
	GetOperator() *string
	SetPrometheus(v *CloudMonitoringPrometheusEscalation) *ConditionConfigUnified
	GetPrometheus() *CloudMonitoringPrometheusEscalation
	SetRelation(v string) *ConditionConfigUnified
	GetRelation() *string
	SetSeverity(v string) *ConditionConfigUnified
	GetSeverity() *string
	SetSimpleEscalation(v *CloudMonitoringSimpleEscalation) *ConditionConfigUnified
	GetSimpleEscalation() *CloudMonitoringSimpleEscalation
	SetThreshold(v float64) *ConditionConfigUnified
	GetThreshold() *float64
	SetThresholdList(v []*ApmThresholdConfig) *ConditionConfigUnified
	GetThresholdList() []*ApmThresholdConfig
	SetTriggers(v []*MetricSetMultiTrigger) *ConditionConfigUnified
	GetTriggers() []*MetricSetMultiTrigger
	SetType(v string) *ConditionConfigUnified
	GetType() *string
	SetYoyTimeUnit(v string) *ConditionConfigUnified
	GetYoyTimeUnit() *string
	SetYoyTimeValue(v int32) *ConditionConfigUnified
	GetYoyTimeValue() *int32
}

type ConditionConfigUnified struct {
	Aggregate                 *string                              `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	CompareList               []*ApmCompositeCompareConfig         `json:"compareList,omitempty" xml:"compareList,omitempty" type:"Repeated"`
	CompositeEscalation       *CloudMonitoringCompositeEscalation  `json:"compositeEscalation,omitempty" xml:"compositeEscalation,omitempty"`
	CountOperator             *string                              `json:"countOperator,omitempty" xml:"countOperator,omitempty"`
	CountThreshold            *int64                               `json:"countThreshold,omitempty" xml:"countThreshold,omitempty"`
	DurationSecs              *int32                               `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	EnableSeveritySuppression *bool                                `json:"enableSeveritySuppression,omitempty" xml:"enableSeveritySuppression,omitempty"`
	EscalationType            *string                              `json:"escalationType,omitempty" xml:"escalationType,omitempty"`
	ExpressEscalation         *CloudMonitoringExpressEscalation    `json:"expressEscalation,omitempty" xml:"expressEscalation,omitempty"`
	LegacyRaw                 *string                              `json:"legacyRaw,omitempty" xml:"legacyRaw,omitempty"`
	LegacyType                *string                              `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
	MatchField                *string                              `json:"matchField,omitempty" xml:"matchField,omitempty"`
	MatchOperator             *string                              `json:"matchOperator,omitempty" xml:"matchOperator,omitempty"`
	MatchValue                *string                              `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
	Max                       *float64                             `json:"max,omitempty" xml:"max,omitempty"`
	Min                       *float64                             `json:"min,omitempty" xml:"min,omitempty"`
	NoDataPolicy              *string                              `json:"noDataPolicy,omitempty" xml:"noDataPolicy,omitempty"`
	Operator                  *string                              `json:"operator,omitempty" xml:"operator,omitempty"`
	Prometheus                *CloudMonitoringPrometheusEscalation `json:"prometheus,omitempty" xml:"prometheus,omitempty"`
	Relation                  *string                              `json:"relation,omitempty" xml:"relation,omitempty"`
	Severity                  *string                              `json:"severity,omitempty" xml:"severity,omitempty"`
	SimpleEscalation          *CloudMonitoringSimpleEscalation     `json:"simpleEscalation,omitempty" xml:"simpleEscalation,omitempty"`
	Threshold                 *float64                             `json:"threshold,omitempty" xml:"threshold,omitempty"`
	ThresholdList             []*ApmThresholdConfig                `json:"thresholdList,omitempty" xml:"thresholdList,omitempty" type:"Repeated"`
	Triggers                  []*MetricSetMultiTrigger             `json:"triggers,omitempty" xml:"triggers,omitempty" type:"Repeated"`
	// This parameter is required.
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
	YoyTimeUnit  *string `json:"yoyTimeUnit,omitempty" xml:"yoyTimeUnit,omitempty"`
	YoyTimeValue *int32  `json:"yoyTimeValue,omitempty" xml:"yoyTimeValue,omitempty"`
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

func (s *ConditionConfigUnified) GetSeverity() *string {
	return s.Severity
}

func (s *ConditionConfigUnified) GetSimpleEscalation() *CloudMonitoringSimpleEscalation {
	return s.SimpleEscalation
}

func (s *ConditionConfigUnified) GetThreshold() *float64 {
	return s.Threshold
}

func (s *ConditionConfigUnified) GetThresholdList() []*ApmThresholdConfig {
	return s.ThresholdList
}

func (s *ConditionConfigUnified) GetTriggers() []*MetricSetMultiTrigger {
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

func (s *ConditionConfigUnified) SetAggregate(v string) *ConditionConfigUnified {
	s.Aggregate = &v
	return s
}

func (s *ConditionConfigUnified) SetCompareList(v []*ApmCompositeCompareConfig) *ConditionConfigUnified {
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

func (s *ConditionConfigUnified) SetThresholdList(v []*ApmThresholdConfig) *ConditionConfigUnified {
	s.ThresholdList = v
	return s
}

func (s *ConditionConfigUnified) SetTriggers(v []*MetricSetMultiTrigger) *ConditionConfigUnified {
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
