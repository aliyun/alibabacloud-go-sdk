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
	// The diagnostic task ID.
	//
	// example:
	//
	// ba30cd99-37a5-4a20-8cd9-ed4b*****
	DiagnoseId *string `json:"diagnoseId,omitempty" xml:"diagnoseId,omitempty"`
	// The time when the deployment is diagnosed.
	//
	// example:
	//
	// 1740389560871
	DiagnoseTime *int64 `json:"diagnoseTime,omitempty" xml:"diagnoseTime,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default-namespace-*****
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The severity level of the risk.
	//
	// Valid values:
	//
	// 	- RISK_LEVEL_HIGH
	//
	// 	- RISK_LEVEL_MID
	//
	// 	- RISK_LEVEL_LOW
	//
	// example:
	//
	// RISK_LEVEL_LOW
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// The diagnostic details.
	Symptoms *JobDiagnosisSymptoms `json:"symptoms,omitempty" xml:"symptoms,omitempty"`
	// The workspace to which the deployment belongs.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
