// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerConditions interface {
	dara.Model
	String() string
	GoString() string
	SetAbsDeviation(v float64) *TriggerConditions
	GetAbsDeviation() *float64
	SetBaselinePeriod(v string) *TriggerConditions
	GetBaselinePeriod() *string
	SetExpressionType(v string) *TriggerConditions
	GetExpressionType() *string
	SetMax(v float64) *TriggerConditions
	GetMax() *float64
	SetMin(v float64) *TriggerConditions
	GetMin() *float64
	SetOperator(v string) *TriggerConditions
	GetOperator() *string
	SetQueryName(v string) *TriggerConditions
	GetQueryName() *string
	SetSensitivity(v string) *TriggerConditions
	GetSensitivity() *string
	SetThreshold(v float64) *TriggerConditions
	GetThreshold() *float64
}

type TriggerConditions struct {
	// The minimum deviation or absolute deviation dead zone for dynamic baselines, which takes effect only for baseline operators. No alert is triggered if the absolute difference between the current value and the boundary is less than the absDeviation value. The unit is the same as the metric unit. The value must be greater than or equal to 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 0.0
	AbsDeviation *float64 `json:"absDeviation,omitempty" xml:"absDeviation,omitempty"`
	// The baseline period, which takes effect only for baseline operators. Valid values: AUTO for automatic identification, DAILY for daily, WEEKLY for weekly, and NONE for no period. If you set this parameter to WEEKLY, the backend automatically extends the historical training window to at least 14 days.
	//
	// example:
	//
	// AUTO
	BaselinePeriod *string `json:"baselinePeriod,omitempty" xml:"baselinePeriod,omitempty"`
	// The expression type. The value is fixed as SIMPLE, which is used for MetricSet multi-threshold triggers.
	//
	// example:
	//
	// SIMPLE
	ExpressionType *string `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	// The upper bound of the range. This parameter is required when the operator is IN_RANGE or OUT_OF_RANGE, and the value must be greater than or equal to the min value.
	//
	// example:
	//
	// 100
	Max *float64 `json:"max,omitempty" xml:"max,omitempty"`
	// The lower bound of the range. This parameter is required when the operator is IN_RANGE or OUT_OF_RANGE.
	//
	// example:
	//
	// 0
	Min *float64 `json:"min,omitempty" xml:"min,omitempty"`
	// The comparison operator. Valid values: GT, GE, LT, LE, EQ, NE, IN_RANGE (requires both min and max), OUT_OF_RANGE (requires both min and max), PRESENT, and NOT_PRESENT.
	//
	// example:
	//
	// GT
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The referenced query name, which corresponds to QueryConfigUnified.queries[].name.
	//
	// example:
	//
	// cpuQuery
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// The dynamic baseline sensitivity, which takes effect only when the operator is ABOVE_UPPER, BELOW_LOWER, or OUT_OF_BAND. Valid values: HIGH for the narrowest and most sensitive band, MEDIUM for medium sensitivity, and LOW for the widest and least sensitive band. You can configure different sensitivities for multiple subconditions.
	//
	// example:
	//
	// MEDIUM
	Sensitivity *string `json:"sensitivity,omitempty" xml:"sensitivity,omitempty"`
	// The comparison threshold. This parameter is used when the operator is GT, GE, LT, LE, EQ, or NE. For IN_RANGE or OUT_OF_RANGE, use the min and max parameters. For PRESENT or NOT_PRESENT, leave this parameter empty.
	//
	// example:
	//
	// 80
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s TriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s TriggerConditions) GoString() string {
	return s.String()
}

func (s *TriggerConditions) GetAbsDeviation() *float64 {
	return s.AbsDeviation
}

func (s *TriggerConditions) GetBaselinePeriod() *string {
	return s.BaselinePeriod
}

func (s *TriggerConditions) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *TriggerConditions) GetMax() *float64 {
	return s.Max
}

func (s *TriggerConditions) GetMin() *float64 {
	return s.Min
}

func (s *TriggerConditions) GetOperator() *string {
	return s.Operator
}

func (s *TriggerConditions) GetQueryName() *string {
	return s.QueryName
}

func (s *TriggerConditions) GetSensitivity() *string {
	return s.Sensitivity
}

func (s *TriggerConditions) GetThreshold() *float64 {
	return s.Threshold
}

func (s *TriggerConditions) SetAbsDeviation(v float64) *TriggerConditions {
	s.AbsDeviation = &v
	return s
}

func (s *TriggerConditions) SetBaselinePeriod(v string) *TriggerConditions {
	s.BaselinePeriod = &v
	return s
}

func (s *TriggerConditions) SetExpressionType(v string) *TriggerConditions {
	s.ExpressionType = &v
	return s
}

func (s *TriggerConditions) SetMax(v float64) *TriggerConditions {
	s.Max = &v
	return s
}

func (s *TriggerConditions) SetMin(v float64) *TriggerConditions {
	s.Min = &v
	return s
}

func (s *TriggerConditions) SetOperator(v string) *TriggerConditions {
	s.Operator = &v
	return s
}

func (s *TriggerConditions) SetQueryName(v string) *TriggerConditions {
	s.QueryName = &v
	return s
}

func (s *TriggerConditions) SetSensitivity(v string) *TriggerConditions {
	s.Sensitivity = &v
	return s
}

func (s *TriggerConditions) SetThreshold(v float64) *TriggerConditions {
	s.Threshold = &v
	return s
}

func (s *TriggerConditions) Validate() error {
	return dara.Validate(s)
}
