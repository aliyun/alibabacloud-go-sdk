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
	// The list of multi-condition configurations.
	Escalations []*CloudMonitoringCompositeEscalationEntry `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// The logical relationship between conditions (AND/OR).
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// The severity level.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The number of consecutive times the conditions are met before the alert is triggered.
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
