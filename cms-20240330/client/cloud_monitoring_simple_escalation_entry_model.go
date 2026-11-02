// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitoringSimpleEscalationEntry interface {
	dara.Model
	String() string
	GoString() string
	SetAbsDeviation(v float64) *CloudMonitoringSimpleEscalationEntry
	GetAbsDeviation() *float64
	SetBaselinePeriod(v string) *CloudMonitoringSimpleEscalationEntry
	GetBaselinePeriod() *string
	SetComparisonOperator(v string) *CloudMonitoringSimpleEscalationEntry
	GetComparisonOperator() *string
	SetPreCondition(v string) *CloudMonitoringSimpleEscalationEntry
	GetPreCondition() *string
	SetSensitivity(v string) *CloudMonitoringSimpleEscalationEntry
	GetSensitivity() *string
	SetSeverity(v string) *CloudMonitoringSimpleEscalationEntry
	GetSeverity() *string
	SetStatistics(v string) *CloudMonitoringSimpleEscalationEntry
	GetStatistics() *string
	SetThreshold(v string) *CloudMonitoringSimpleEscalationEntry
	GetThreshold() *string
	SetTimes(v int32) *CloudMonitoringSimpleEscalationEntry
	GetTimes() *int32
}

type CloudMonitoringSimpleEscalationEntry struct {
	// The minimum dynamic baseline deviation or absolute deviation dead zone, which takes effect only when the baseline comparisonOperator is used. If the absolute difference between the current value and the boundary is less than the absDeviation value, no alert is triggered. The unit is the same as the metric. The value must be greater than or equal to 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 0.0
	AbsDeviation *float64 `json:"absDeviation,omitempty" xml:"absDeviation,omitempty"`
	// The baseline period, which takes effect only when the baseline comparisonOperator is used. Valid values: AUTO for automatic identification, DAILY for daily, WEEKLY for weekly, and NONE for no period. For weekly periods, the backend automatically extends the historical training window to at least 14 days. Automatic identification cannot return specific identification results.
	//
	// example:
	//
	// AUTO
	BaselinePeriod *string `json:"baselinePeriod,omitempty" xml:"baselinePeriod,omitempty"`
	// The comparison operator.
	//
	// example:
	//
	// GT
	ComparisonOperator *string `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	// The precondition.
	//
	// example:
	//
	// 连续3次
	PreCondition *string `json:"preCondition,omitempty" xml:"preCondition,omitempty"`
	// The sensitivity of the dynamic baseline, which takes effect only when the baseline comparisonOperator is used. Valid values: HIGH for the narrowest and most sensitive band, MEDIUM for medium sensitivity, and LOW for the widest and least sensitive band.
	//
	// example:
	//
	// MEDIUM
	Sensitivity *string `json:"sensitivity,omitempty" xml:"sensitivity,omitempty"`
	// The severity level.
	//
	// example:
	//
	// CRITICAL
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The statistical method.
	//
	// example:
	//
	// Average
	Statistics *string `json:"statistics,omitempty" xml:"statistics,omitempty"`
	// The threshold.
	//
	// example:
	//
	// 80
	Threshold *string `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// The number of consecutive triggers.
	//
	// example:
	//
	// 3
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s CloudMonitoringSimpleEscalationEntry) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitoringSimpleEscalationEntry) GoString() string {
	return s.String()
}

func (s *CloudMonitoringSimpleEscalationEntry) GetAbsDeviation() *float64 {
	return s.AbsDeviation
}

func (s *CloudMonitoringSimpleEscalationEntry) GetBaselinePeriod() *string {
	return s.BaselinePeriod
}

func (s *CloudMonitoringSimpleEscalationEntry) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CloudMonitoringSimpleEscalationEntry) GetPreCondition() *string {
	return s.PreCondition
}

func (s *CloudMonitoringSimpleEscalationEntry) GetSensitivity() *string {
	return s.Sensitivity
}

func (s *CloudMonitoringSimpleEscalationEntry) GetSeverity() *string {
	return s.Severity
}

func (s *CloudMonitoringSimpleEscalationEntry) GetStatistics() *string {
	return s.Statistics
}

func (s *CloudMonitoringSimpleEscalationEntry) GetThreshold() *string {
	return s.Threshold
}

func (s *CloudMonitoringSimpleEscalationEntry) GetTimes() *int32 {
	return s.Times
}

func (s *CloudMonitoringSimpleEscalationEntry) SetAbsDeviation(v float64) *CloudMonitoringSimpleEscalationEntry {
	s.AbsDeviation = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetBaselinePeriod(v string) *CloudMonitoringSimpleEscalationEntry {
	s.BaselinePeriod = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetComparisonOperator(v string) *CloudMonitoringSimpleEscalationEntry {
	s.ComparisonOperator = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetPreCondition(v string) *CloudMonitoringSimpleEscalationEntry {
	s.PreCondition = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetSensitivity(v string) *CloudMonitoringSimpleEscalationEntry {
	s.Sensitivity = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetSeverity(v string) *CloudMonitoringSimpleEscalationEntry {
	s.Severity = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetStatistics(v string) *CloudMonitoringSimpleEscalationEntry {
	s.Statistics = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetThreshold(v string) *CloudMonitoringSimpleEscalationEntry {
	s.Threshold = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetTimes(v int32) *CloudMonitoringSimpleEscalationEntry {
	s.Times = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) Validate() error {
	return dara.Validate(s)
}
