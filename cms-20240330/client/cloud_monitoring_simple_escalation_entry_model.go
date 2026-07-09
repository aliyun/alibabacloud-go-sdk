// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitoringSimpleEscalationEntry interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *CloudMonitoringSimpleEscalationEntry
	GetComparisonOperator() *string
	SetPreCondition(v string) *CloudMonitoringSimpleEscalationEntry
	GetPreCondition() *string
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
	ComparisonOperator *string `json:"comparisonOperator,omitempty" xml:"comparisonOperator,omitempty"`
	PreCondition       *string `json:"preCondition,omitempty" xml:"preCondition,omitempty"`
	Severity           *string `json:"severity,omitempty" xml:"severity,omitempty"`
	Statistics         *string `json:"statistics,omitempty" xml:"statistics,omitempty"`
	Threshold          *string `json:"threshold,omitempty" xml:"threshold,omitempty"`
	Times              *int32  `json:"times,omitempty" xml:"times,omitempty"`
}

func (s CloudMonitoringSimpleEscalationEntry) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitoringSimpleEscalationEntry) GoString() string {
	return s.String()
}

func (s *CloudMonitoringSimpleEscalationEntry) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CloudMonitoringSimpleEscalationEntry) GetPreCondition() *string {
	return s.PreCondition
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

func (s *CloudMonitoringSimpleEscalationEntry) SetComparisonOperator(v string) *CloudMonitoringSimpleEscalationEntry {
	s.ComparisonOperator = &v
	return s
}

func (s *CloudMonitoringSimpleEscalationEntry) SetPreCondition(v string) *CloudMonitoringSimpleEscalationEntry {
	s.PreCondition = &v
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
