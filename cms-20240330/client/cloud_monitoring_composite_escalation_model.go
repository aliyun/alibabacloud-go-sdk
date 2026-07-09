// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitoringCompositeEscalation interface {
	dara.Model
	String() string
	GoString() string
	SetEscalations(v []*CloudMonitoringCompositeEscalationEntry) *CloudMonitoringCompositeEscalation
	GetEscalations() []*CloudMonitoringCompositeEscalationEntry
	SetRelation(v string) *CloudMonitoringCompositeEscalation
	GetRelation() *string
	SetSeverity(v string) *CloudMonitoringCompositeEscalation
	GetSeverity() *string
	SetTimes(v int32) *CloudMonitoringCompositeEscalation
	GetTimes() *int32
}

type CloudMonitoringCompositeEscalation struct {
	// A single entry in the escalation policy. See the `CloudMonitoringCompositeEscalationEntry` object for details.
	Escalations []*CloudMonitoringCompositeEscalationEntry `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// Specifies the logical relationship for evaluating the conditions of the composite alert rule. Valid values: `and` and `or`.
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// Specifies the severity level of the alert. For example: `Critical`, `Warning`, and `Info`.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// Specifies the number of times the alert conditions must be met to trigger this escalation policy.
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s CloudMonitoringCompositeEscalation) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitoringCompositeEscalation) GoString() string {
	return s.String()
}

func (s *CloudMonitoringCompositeEscalation) GetEscalations() []*CloudMonitoringCompositeEscalationEntry {
	return s.Escalations
}

func (s *CloudMonitoringCompositeEscalation) GetRelation() *string {
	return s.Relation
}

func (s *CloudMonitoringCompositeEscalation) GetSeverity() *string {
	return s.Severity
}

func (s *CloudMonitoringCompositeEscalation) GetTimes() *int32 {
	return s.Times
}

func (s *CloudMonitoringCompositeEscalation) SetEscalations(v []*CloudMonitoringCompositeEscalationEntry) *CloudMonitoringCompositeEscalation {
	s.Escalations = v
	return s
}

func (s *CloudMonitoringCompositeEscalation) SetRelation(v string) *CloudMonitoringCompositeEscalation {
	s.Relation = &v
	return s
}

func (s *CloudMonitoringCompositeEscalation) SetSeverity(v string) *CloudMonitoringCompositeEscalation {
	s.Severity = &v
	return s
}

func (s *CloudMonitoringCompositeEscalation) SetTimes(v int32) *CloudMonitoringCompositeEscalation {
	s.Times = &v
	return s
}

func (s *CloudMonitoringCompositeEscalation) Validate() error {
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
