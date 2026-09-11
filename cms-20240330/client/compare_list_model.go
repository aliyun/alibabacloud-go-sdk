// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareList interface {
	dara.Model
	String() string
	GoString() string
	SetAbsDeviation(v float64) *CompareList
	GetAbsDeviation() *float64
	SetAggregate(v string) *CompareList
	GetAggregate() *string
	SetBaselinePeriod(v string) *CompareList
	GetBaselinePeriod() *string
	SetOperator(v string) *CompareList
	GetOperator() *string
	SetSensitivity(v string) *CompareList
	GetSensitivity() *string
	SetThreshold(v float32) *CompareList
	GetThreshold() *float32
	SetYoyTimeUnit(v string) *CompareList
	GetYoyTimeUnit() *string
	SetYoyTimeValue(v int32) *CompareList
	GetYoyTimeValue() *int32
}

type CompareList struct {
	// The dynamic baseline minimum deviation or absolute deviation dead zone. This parameter takes effect only when a baseline operator is used. If |current value − boundary| < absDeviation, no alert is triggered. The unit is the same as the metric unit. The value must be greater than or equal to 0. A value of 0 indicates no restriction.
	//
	// example:
	//
	// 0.0
	AbsDeviation *float64 `json:"absDeviation,omitempty" xml:"absDeviation,omitempty"`
	// The aggregation function.
	//
	// This parameter is required.
	//
	// example:
	//
	// AVG
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// The baseline period. This parameter takes effect only when a baseline operator is used. Valid values:
	//
	// - AUTO: Automatically identifies the period. The specific identification result cannot be displayed.
	//
	// - DAILY: Daily period.
	//
	// - WEEKLY: Weekly period. The backend automatically expands the historical training window to at least 14 days.
	//
	// - NONE: No period.
	//
	// example:
	//
	// AUTO
	BaselinePeriod *string `json:"baselinePeriod,omitempty" xml:"baselinePeriod,omitempty"`
	// The comparison operator. Valid values:
	//
	// - GTE: greater than or equal to.
	//
	// - LTE: less than or equal to.
	//
	// - YOY_UP: year-over-year increase. You must also specify yoyTimeUnit and yoyTimeValue.
	//
	// - YOY_DOWN: year-over-year decrease. You must also specify yoyTimeUnit and yoyTimeValue.
	//
	// - ABOVE_UPPER: dynamic baseline spike. You must specify sensitivity. When using a baseline operator, threshold is not used for evaluation. Set it to 0 as a placeholder.
	//
	// - BELOW_LOWER: dynamic baseline drop. You must specify sensitivity. When using a baseline operator, threshold is not used for evaluation. Set it to 0 as a placeholder.
	//
	// - OUT_OF_BAND: dynamic baseline bidirectional. You must specify sensitivity. When using a baseline operator, threshold is not used for evaluation. Set it to 0 as a placeholder.
	//
	// This parameter is required.
	//
	// example:
	//
	// OUT_OF_BAND
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The dynamic baseline sensitivity. This parameter takes effect only when a baseline operator is used. Valid values:
	//
	// - HIGH: The narrowest band and the most sensitive.
	//
	// - MEDIUM: Medium sensitivity.
	//
	// - LOW: The widest band and the least sensitive.
	//
	// example:
	//
	// MEDIUM
	Sensitivity *string `json:"sensitivity,omitempty" xml:"sensitivity,omitempty"`
	// The threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// The year-over-year time unit. This parameter takes effect only when operator is set to YOY_UP or YOY_DOWN.
	//
	// example:
	//
	// day
	YoyTimeUnit *string `json:"yoyTimeUnit,omitempty" xml:"yoyTimeUnit,omitempty"`
	// The year-over-year time value. This parameter takes effect only when operator is set to YOY_UP or YOY_DOWN.
	//
	// example:
	//
	// 1
	YoyTimeValue *int32 `json:"yoyTimeValue,omitempty" xml:"yoyTimeValue,omitempty"`
}

func (s CompareList) String() string {
	return dara.Prettify(s)
}

func (s CompareList) GoString() string {
	return s.String()
}

func (s *CompareList) GetAbsDeviation() *float64 {
	return s.AbsDeviation
}

func (s *CompareList) GetAggregate() *string {
	return s.Aggregate
}

func (s *CompareList) GetBaselinePeriod() *string {
	return s.BaselinePeriod
}

func (s *CompareList) GetOperator() *string {
	return s.Operator
}

func (s *CompareList) GetSensitivity() *string {
	return s.Sensitivity
}

func (s *CompareList) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CompareList) GetYoyTimeUnit() *string {
	return s.YoyTimeUnit
}

func (s *CompareList) GetYoyTimeValue() *int32 {
	return s.YoyTimeValue
}

func (s *CompareList) SetAbsDeviation(v float64) *CompareList {
	s.AbsDeviation = &v
	return s
}

func (s *CompareList) SetAggregate(v string) *CompareList {
	s.Aggregate = &v
	return s
}

func (s *CompareList) SetBaselinePeriod(v string) *CompareList {
	s.BaselinePeriod = &v
	return s
}

func (s *CompareList) SetOperator(v string) *CompareList {
	s.Operator = &v
	return s
}

func (s *CompareList) SetSensitivity(v string) *CompareList {
	s.Sensitivity = &v
	return s
}

func (s *CompareList) SetThreshold(v float32) *CompareList {
	s.Threshold = &v
	return s
}

func (s *CompareList) SetYoyTimeUnit(v string) *CompareList {
	s.YoyTimeUnit = &v
	return s
}

func (s *CompareList) SetYoyTimeValue(v int32) *CompareList {
	s.YoyTimeValue = &v
	return s
}

func (s *CompareList) Validate() error {
	return dara.Validate(s)
}
