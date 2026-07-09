// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitoringSimpleEscalation interface {
	dara.Model
	String() string
	GoString() string
	SetEscalations(v []*CloudMonitoringSimpleEscalationEntry) *CloudMonitoringSimpleEscalation
	GetEscalations() []*CloudMonitoringSimpleEscalationEntry
	SetMetricName(v string) *CloudMonitoringSimpleEscalation
	GetMetricName() *string
	SetPeriod(v int32) *CloudMonitoringSimpleEscalation
	GetPeriod() *int32
}

type CloudMonitoringSimpleEscalation struct {
	// An object that defines a single escalation rule.
	Escalations []*CloudMonitoringSimpleEscalationEntry `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// The name of the metric.
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The evaluation period for the metric, in seconds.
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
}

func (s CloudMonitoringSimpleEscalation) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitoringSimpleEscalation) GoString() string {
	return s.String()
}

func (s *CloudMonitoringSimpleEscalation) GetEscalations() []*CloudMonitoringSimpleEscalationEntry {
	return s.Escalations
}

func (s *CloudMonitoringSimpleEscalation) GetMetricName() *string {
	return s.MetricName
}

func (s *CloudMonitoringSimpleEscalation) GetPeriod() *int32 {
	return s.Period
}

func (s *CloudMonitoringSimpleEscalation) SetEscalations(v []*CloudMonitoringSimpleEscalationEntry) *CloudMonitoringSimpleEscalation {
	s.Escalations = v
	return s
}

func (s *CloudMonitoringSimpleEscalation) SetMetricName(v string) *CloudMonitoringSimpleEscalation {
	s.MetricName = &v
	return s
}

func (s *CloudMonitoringSimpleEscalation) SetPeriod(v int32) *CloudMonitoringSimpleEscalation {
	s.Period = &v
	return s
}

func (s *CloudMonitoringSimpleEscalation) Validate() error {
	if s.Escalations != nil {
		for _, item := range s.Escalations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
