// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MedicalKnowInput) *MedicalKnowledgeRequest
	GetBody() *MedicalKnowInput
}

type MedicalKnowledgeRequest struct {
	Body *MedicalKnowInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MedicalKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s MedicalKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *MedicalKnowledgeRequest) GetBody() *MedicalKnowInput {
	return s.Body
}

func (s *MedicalKnowledgeRequest) SetBody(v *MedicalKnowInput) *MedicalKnowledgeRequest {
	s.Body = v
	return s
}

func (s *MedicalKnowledgeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
