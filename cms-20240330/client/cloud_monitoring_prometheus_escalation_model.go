// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitoringPrometheusEscalation interface {
	dara.Model
	String() string
	GoString() string
	SetPromQl(v string) *CloudMonitoringPrometheusEscalation
	GetPromQl() *string
	SetSeverity(v string) *CloudMonitoringPrometheusEscalation
	GetSeverity() *string
	SetTimes(v int32) *CloudMonitoringPrometheusEscalation
	GetTimes() *int32
}

type CloudMonitoringPrometheusEscalation struct {
	// The PromQL expression that defines the alert condition. This parameter is required.
	PromQl *string `json:"promQl,omitempty" xml:"promQl,omitempty"`
	// The severity of the alert that triggers the escalation. This parameter is required.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The number of consecutive times the condition must be met to trigger an escalation. This parameter is required.
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s CloudMonitoringPrometheusEscalation) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitoringPrometheusEscalation) GoString() string {
	return s.String()
}

func (s *CloudMonitoringPrometheusEscalation) GetPromQl() *string {
	return s.PromQl
}

func (s *CloudMonitoringPrometheusEscalation) GetSeverity() *string {
	return s.Severity
}

func (s *CloudMonitoringPrometheusEscalation) GetTimes() *int32 {
	return s.Times
}

func (s *CloudMonitoringPrometheusEscalation) SetPromQl(v string) *CloudMonitoringPrometheusEscalation {
	s.PromQl = &v
	return s
}

func (s *CloudMonitoringPrometheusEscalation) SetSeverity(v string) *CloudMonitoringPrometheusEscalation {
	s.Severity = &v
	return s
}

func (s *CloudMonitoringPrometheusEscalation) SetTimes(v int32) *CloudMonitoringPrometheusEscalation {
	s.Times = &v
	return s
}

func (s *CloudMonitoringPrometheusEscalation) Validate() error {
	return dara.Validate(s)
}
