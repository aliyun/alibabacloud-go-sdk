// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAnswerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MedicalAnswerInput) *MedicalAnswerRequest
	GetBody() *MedicalAnswerInput
}

type MedicalAnswerRequest struct {
	Body *MedicalAnswerInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MedicalAnswerRequest) String() string {
	return dara.Prettify(s)
}

func (s MedicalAnswerRequest) GoString() string {
	return s.String()
}

func (s *MedicalAnswerRequest) GetBody() *MedicalAnswerInput {
	return s.Body
}

func (s *MedicalAnswerRequest) SetBody(v *MedicalAnswerInput) *MedicalAnswerRequest {
	s.Body = v
	return s
}

func (s *MedicalAnswerRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
