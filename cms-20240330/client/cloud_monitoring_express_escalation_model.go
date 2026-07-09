// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitoringExpressEscalation interface {
	dara.Model
	String() string
	GoString() string
	SetRawExpression(v string) *CloudMonitoringExpressEscalation
	GetRawExpression() *string
	SetSeverity(v string) *CloudMonitoringExpressEscalation
	GetSeverity() *string
	SetTimes(v int32) *CloudMonitoringExpressEscalation
	GetTimes() *int32
}

type CloudMonitoringExpressEscalation struct {
	// The expression that defines the alert condition.
	RawExpression *string `json:"rawExpression,omitempty" xml:"rawExpression,omitempty"`
	// The alert severity that triggers the escalation.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The number of alert occurrences required to trigger the escalation.
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s CloudMonitoringExpressEscalation) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitoringExpressEscalation) GoString() string {
	return s.String()
}

func (s *CloudMonitoringExpressEscalation) GetRawExpression() *string {
	return s.RawExpression
}

func (s *CloudMonitoringExpressEscalation) GetSeverity() *string {
	return s.Severity
}

func (s *CloudMonitoringExpressEscalation) GetTimes() *int32 {
	return s.Times
}

func (s *CloudMonitoringExpressEscalation) SetRawExpression(v string) *CloudMonitoringExpressEscalation {
	s.RawExpression = &v
	return s
}

func (s *CloudMonitoringExpressEscalation) SetSeverity(v string) *CloudMonitoringExpressEscalation {
	s.Severity = &v
	return s
}

func (s *CloudMonitoringExpressEscalation) SetTimes(v int32) *CloudMonitoringExpressEscalation {
	s.Times = &v
	return s
}

func (s *CloudMonitoringExpressEscalation) Validate() error {
	return dara.Validate(s)
}
