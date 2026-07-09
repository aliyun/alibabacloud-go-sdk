// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitoringCompositeEscalationEntry interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *CloudMonitoringCompositeEscalationEntry
	GetComparisonOperator() *string
	SetMetricName(v string) *CloudMonitoringCompositeEscalationEntry
	GetMetricName() *string
	SetPeriod(v int32) *CloudMonitoringCompositeEscalationEntry
	GetPeriod() *int32
	SetPreCondition(v string) *CloudMonitoringCompositeEscalationEntry
	GetPreCondition() *string
	SetStatistics(v string) *CloudMonitoringCompositeEscalationEntry
	GetStatistics() *string
	SetThreshold(v string) *CloudMonitoringCompositeEscalationEntry
	GetThreshold() *string
}

type CloudMonitoringCompositeEscalationEntry struct {
	ComparisonOperator *string `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	MetricName         *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	Period             *int32  `json:"period,omitempty" xml:"period,omitempty"`
	PreCondition       *string `json:"preCondition,omitempty" xml:"preCondition,omitempty"`
	Statistics         *string `json:"statistics,omitempty" xml:"statistics,omitempty"`
	Threshold          *string `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s CloudMonitoringCompositeEscalationEntry) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitoringCompositeEscalationEntry) GoString() string {
	return s.String()
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

func (s *CloudMonitoringCompositeEscalationEntry) GetStatistics() *string {
	return s.Statistics
}

func (s *CloudMonitoringCompositeEscalationEntry) GetThreshold() *string {
	return s.Threshold
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
