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
	SetDurationSecs(v int32) *ConditionConfigUnified
	GetDurationSecs() *int32
	SetOperator(v string) *ConditionConfigUnified
	GetOperator() *string
	SetRelation(v string) *ConditionConfigUnified
	GetRelation() *string
	SetSeverity(v string) *ConditionConfigUnified
	GetSeverity() *string
	SetThreshold(v float64) *ConditionConfigUnified
	GetThreshold() *float64
	SetThresholdList(v []*ApmThresholdConfig) *ConditionConfigUnified
	GetThresholdList() []*ApmThresholdConfig
	SetType(v string) *ConditionConfigUnified
	GetType() *string
}

type ConditionConfigUnified struct {
	// The aggregation method for metric data points over the evaluation period. Valid values include `AVG`, `SUM`, and `MAX`.
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// A list of composite comparison configurations for APM alerts. Each item is an `ApmCompositeCompareConfig` object.
	CompareList []*ApmCompositeCompareConfig `json:"compareList,omitempty" xml:"compareList,omitempty" type:"Repeated"`
	// The number of seconds a condition must be true before triggering an alert.
	DurationSecs *int32 `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	// The comparison operator used to evaluate the metric against the threshold.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The logical relationship between multiple conditions. Valid values are `AND` and `OR`.
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// The alert severity. Valid values are `CRITICAL`, `WARNING`, and `INFO`.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The value against which the metric is evaluated to trigger an alert.
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// A list of threshold configurations for Application Performance Monitoring (APM) alerts. Each item is an `ApmThresholdConfig` object.
	ThresholdList []*ApmThresholdConfig `json:"thresholdList,omitempty" xml:"thresholdList,omitempty" type:"Repeated"`
	// The type of the alert condition.
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *ConditionConfigUnified) GetDurationSecs() *int32 {
	return s.DurationSecs
}

func (s *ConditionConfigUnified) GetOperator() *string {
	return s.Operator
}

func (s *ConditionConfigUnified) GetRelation() *string {
	return s.Relation
}

func (s *ConditionConfigUnified) GetSeverity() *string {
	return s.Severity
}

func (s *ConditionConfigUnified) GetThreshold() *float64 {
	return s.Threshold
}

func (s *ConditionConfigUnified) GetThresholdList() []*ApmThresholdConfig {
	return s.ThresholdList
}

func (s *ConditionConfigUnified) GetType() *string {
	return s.Type
}

func (s *ConditionConfigUnified) SetAggregate(v string) *ConditionConfigUnified {
	s.Aggregate = &v
	return s
}

func (s *ConditionConfigUnified) SetCompareList(v []*ApmCompositeCompareConfig) *ConditionConfigUnified {
	s.CompareList = v
	return s
}

func (s *ConditionConfigUnified) SetDurationSecs(v int32) *ConditionConfigUnified {
	s.DurationSecs = &v
	return s
}

func (s *ConditionConfigUnified) SetOperator(v string) *ConditionConfigUnified {
	s.Operator = &v
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

func (s *ConditionConfigUnified) SetThreshold(v float64) *ConditionConfigUnified {
	s.Threshold = &v
	return s
}

func (s *ConditionConfigUnified) SetThresholdList(v []*ApmThresholdConfig) *ConditionConfigUnified {
	s.ThresholdList = v
	return s
}

func (s *ConditionConfigUnified) SetType(v string) *ConditionConfigUnified {
	s.Type = &v
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
	if s.ThresholdList != nil {
		for _, item := range s.ThresholdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
