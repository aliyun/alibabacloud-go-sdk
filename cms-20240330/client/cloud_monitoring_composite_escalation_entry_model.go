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
	// The minimum deviation or absolute deviation dead zone for dynamic baselines. This parameter takes effect only when the comparisonOperator is set to a baseline operator. If the absolute difference between the current value and the boundary is less than the absDeviation value, no alert is triggered. The unit is the same as the metric unit. The value must be greater than or equal to 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 0.0
	AbsDeviation *float64 `json:"absDeviation,omitempty" xml:"absDeviation,omitempty"`
	// The baseline period. This parameter takes effect only when the comparisonOperator is set to a baseline operator. Valid values: AUTO for automatic identification, DAILY for daily, WEEKLY for weekly, and NONE for no period. When set to WEEKLY, the backend automatically extends the historical training window to 14 days or more. Automatic identification cannot return the specific identification result.
	//
	// example:
	//
	// AUTO
	BaselinePeriod *string `json:"baselinePeriod,omitempty" xml:"baselinePeriod,omitempty"`
	// The comparison operator. Valid values include GreaterThanThreshold and LessThanThreshold for static thresholds, and ABOVE_UPPER, BELOW_LOWER, and OUT_OF_BAND for dynamic baseline spikes, drops, and bidirectional changes. For dynamic baseline operators, you must specify the sensitivity parameter. The threshold parameter is not used for evaluation and can be left empty.
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
	// The dynamic baseline sensitivity. This parameter takes effect only when the comparisonOperator is set to a baseline operator. Valid values: HIGH for the narrowest and most sensitive band, MEDIUM for a medium band, and LOW for the widest and least sensitive band.
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
