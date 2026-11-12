// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitoringCompositeEscalationEntry interface {
	dara.Model
	String() string
	GoString() string
	SetAbsDeviation(v float64) *CloudMonitoringCompositeEscalationEntry
	GetAbsDeviation() *float64
	SetBaselinePeriod(v string) *CloudMonitoringCompositeEscalationEntry
	GetBaselinePeriod() *string
	SetComparisonOperator(v string) *CloudMonitoringCompositeEscalationEntry
	GetComparisonOperator() *string
	SetMetricName(v string) *CloudMonitoringCompositeEscalationEntry
	GetMetricName() *string
	SetPeriod(v int32) *CloudMonitoringCompositeEscalationEntry
	GetPeriod() *int32
	SetPreCondition(v string) *CloudMonitoringCompositeEscalationEntry
	GetPreCondition() *string
	SetSensitivity(v string) *CloudMonitoringCompositeEscalationEntry
	GetSensitivity() *string
	SetStatistics(v string) *CloudMonitoringCompositeEscalationEntry
	GetStatistics() *string
	SetThreshold(v string) *CloudMonitoringCompositeEscalationEntry
	GetThreshold() *string
}

type CloudMonitoringCompositeEscalationEntry struct {
	// The minimum deviation or absolute deviation dead zone for the dynamic baseline. This parameter takes effect only when the comparisonOperator is set to a baseline operator. If the absolute difference between the current value and the boundary is less than the absDeviation value, no alert is triggered. The unit is the same as the metric unit. The value must be greater than or equal to 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 0.0
	AbsDeviation *float64 `json:"absDeviation,omitempty" xml:"absDeviation,omitempty"`
	// The baseline period. This parameter takes effect only when the comparisonOperator is set to a baseline operator. Valid values: AUTO (automatic identification), DAILY (daily), WEEKLY (weekly), and NONE (no period). If you set this parameter to WEEKLY, the backend automatically extends the historical training window to at least 14 days. Specific identification results are not returned for automatic identification.
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
	// The metric name.
	//
	// example:
	//
	// cpuUsage
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The collection period in seconds.
	//
	// example:
	//
	// 60
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The precondition.
	//
	// example:
	//
	// 连续3次
	PreCondition *string `json:"preCondition,omitempty" xml:"preCondition,omitempty"`
	// The sensitivity of the dynamic baseline. This parameter takes effect only when the comparisonOperator is set to a baseline operator. Valid values: HIGH (narrowest band and most sensitive), MEDIUM (medium), and LOW (widest band and least sensitive).
	//
	// example:
	//
	// MEDIUM
	Sensitivity *string `json:"sensitivity,omitempty" xml:"sensitivity,omitempty"`
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
}

func (s CloudMonitoringCompositeEscalationEntry) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitoringCompositeEscalationEntry) GoString() string {
	return s.String()
}

func (s *CloudMonitoringCompositeEscalationEntry) GetAbsDeviation() *float64 {
	return s.AbsDeviation
}

func (s *CloudMonitoringCompositeEscalationEntry) GetBaselinePeriod() *string {
	return s.BaselinePeriod
}

func (s *CloudMonitoringCompositeEscalationEntry) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CloudMonitoringCompositeEscalationEntry) GetMetricName() *string {
	return s.MetricName
}

func (s *CloudMonitoringCompositeEscalationEntry) GetPeriod() *int32 {
	return s.Period
}

func (s *CloudMonitoringCompositeEscalationEntry) GetPreCondition() *string {
	return s.PreCondition
}

func (s *CloudMonitoringCompositeEscalationEntry) GetSensitivity() *string {
	return s.Sensitivity
}

func (s *CloudMonitoringCompositeEscalationEntry) GetStatistics() *string {
	return s.Statistics
}

func (s *CloudMonitoringCompositeEscalationEntry) GetThreshold() *string {
	return s.Threshold
}

func (s *CloudMonitoringCompositeEscalationEntry) SetAbsDeviation(v float64) *CloudMonitoringCompositeEscalationEntry {
	s.AbsDeviation = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) SetBaselinePeriod(v string) *CloudMonitoringCompositeEscalationEntry {
	s.BaselinePeriod = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) SetComparisonOperator(v string) *CloudMonitoringCompositeEscalationEntry {
	s.ComparisonOperator = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) SetMetricName(v string) *CloudMonitoringCompositeEscalationEntry {
	s.MetricName = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) SetPeriod(v int32) *CloudMonitoringCompositeEscalationEntry {
	s.Period = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) SetPreCondition(v string) *CloudMonitoringCompositeEscalationEntry {
	s.PreCondition = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) SetSensitivity(v string) *CloudMonitoringCompositeEscalationEntry {
	s.Sensitivity = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) SetStatistics(v string) *CloudMonitoringCompositeEscalationEntry {
	s.Statistics = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) SetThreshold(v string) *CloudMonitoringCompositeEscalationEntry {
	s.Threshold = &v
	return s
}

func (s *CloudMonitoringCompositeEscalationEntry) Validate() error {
	return dara.Validate(s)
}
