// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MedicalKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MedicalKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *MedicalKnowResult) *MedicalKnowledgeResponse
	GetBody() *MedicalKnowResult
}

type MedicalKnowledgeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MedicalKnowResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MedicalKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s MedicalKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *MedicalKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MedicalKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MedicalKnowledgeResponse) GetBody() *MedicalKnowResult {
	return s.Body
}

func (s *MedicalKnowledgeResponse) SetHeaders(v map[string]*string) *MedicalKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *MedicalKnowledgeResponse) SetStatusCode(v int32) *MedicalKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *MedicalKnowledgeResponse) SetBody(v *MedicalKnowResult) *MedicalKnowledgeResponse {
	s.Body = v
	return s
}

func (s *MedicalKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
