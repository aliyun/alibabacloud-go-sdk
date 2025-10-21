// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobDiagnosis interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnoseId(v string) *JobDiagnosis
	GetDiagnoseId() *string
	SetDiagnoseTime(v int64) *JobDiagnosis
	GetDiagnoseTime() *int64
	SetNamespace(v string) *JobDiagnosis
	GetNamespace() *string
	SetRiskLevel(v string) *JobDiagnosis
	GetRiskLevel() *string
	SetSymptoms(v *JobDiagnosisSymptoms) *JobDiagnosis
	GetSymptoms() *JobDiagnosisSymptoms
	SetWorkspace(v string) *JobDiagnosis
	GetWorkspace() *string
}

type JobDiagnosis struct {
	DiagnoseId   *string               `json:"diagnoseId,omitempty" xml:"diagnoseId,omitempty"`
	DiagnoseTime *int64                `json:"diagnoseTime,omitempty" xml:"diagnoseTime,omitempty"`
	Namespace    *string               `json:"namespace,omitempty" xml:"namespace,omitempty"`
	RiskLevel    *string               `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	Symptoms     *JobDiagnosisSymptoms `json:"symptoms,omitempty" xml:"symptoms,omitempty"`
	Workspace    *string               `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s JobDiagnosis) String() string {
	return dara.Prettify(s)
}

func (s JobDiagnosis) GoString() string {
	return s.String()
}

func (s *JobDiagnosis) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *JobDiagnosis) GetDiagnoseTime() *int64 {
	return s.DiagnoseTime
}

func (s *JobDiagnosis) GetNamespace() *string {
	return s.Namespace
}

func (s *JobDiagnosis) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *JobDiagnosis) GetSymptoms() *JobDiagnosisSymptoms {
	return s.Symptoms
}

func (s *JobDiagnosis) GetWorkspace() *string {
	return s.Workspace
}

func (s *JobDiagnosis) SetDiagnoseId(v string) *JobDiagnosis {
	s.DiagnoseId = &v
	return s
}

func (s *JobDiagnosis) SetDiagnoseTime(v int64) *JobDiagnosis {
	s.DiagnoseTime = &v
	return s
}

func (s *JobDiagnosis) SetNamespace(v string) *JobDiagnosis {
	s.Namespace = &v
	return s
}

func (s *JobDiagnosis) SetRiskLevel(v string) *JobDiagnosis {
	s.RiskLevel = &v
	return s
}

func (s *JobDiagnosis) SetSymptoms(v *JobDiagnosisSymptoms) *JobDiagnosis {
	s.Symptoms = v
	return s
}

func (s *JobDiagnosis) SetWorkspace(v string) *JobDiagnosis {
	s.Workspace = &v
	return s
}

func (s *JobDiagnosis) Validate() error {
	if s.Symptoms != nil {
		if err := s.Symptoms.Validate(); err != nil {
			return err
		}
	}
	return nil
}
