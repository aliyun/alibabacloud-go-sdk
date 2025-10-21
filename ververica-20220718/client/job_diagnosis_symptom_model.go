// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobDiagnosisSymptom interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *JobDiagnosisSymptom
	GetDescription() *string
	SetName(v string) *JobDiagnosisSymptom
	GetName() *string
	SetRecommendation(v string) *JobDiagnosisSymptom
	GetRecommendation() *string
}

type JobDiagnosisSymptom struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Recommendation *string `json:"recommendation,omitempty" xml:"recommendation,omitempty"`
}

func (s JobDiagnosisSymptom) String() string {
	return dara.Prettify(s)
}

func (s JobDiagnosisSymptom) GoString() string {
	return s.String()
}

func (s *JobDiagnosisSymptom) GetDescription() *string {
	return s.Description
}

func (s *JobDiagnosisSymptom) GetName() *string {
	return s.Name
}

func (s *JobDiagnosisSymptom) GetRecommendation() *string {
	return s.Recommendation
}

func (s *JobDiagnosisSymptom) SetDescription(v string) *JobDiagnosisSymptom {
	s.Description = &v
	return s
}

func (s *JobDiagnosisSymptom) SetName(v string) *JobDiagnosisSymptom {
	s.Name = &v
	return s
}

func (s *JobDiagnosisSymptom) SetRecommendation(v string) *JobDiagnosisSymptom {
	s.Recommendation = &v
	return s
}

func (s *JobDiagnosisSymptom) Validate() error {
	return dara.Validate(s)
}
