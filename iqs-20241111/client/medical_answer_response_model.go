// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAnswerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MedicalAnswerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MedicalAnswerResponse
	GetStatusCode() *int32
	SetBody(v *MedicalAnswerResult) *MedicalAnswerResponse
	GetBody() *MedicalAnswerResult
}

type MedicalAnswerResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MedicalAnswerResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MedicalAnswerResponse) String() string {
	return dara.Prettify(s)
}

func (s MedicalAnswerResponse) GoString() string {
	return s.String()
}

func (s *MedicalAnswerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MedicalAnswerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MedicalAnswerResponse) GetBody() *MedicalAnswerResult {
	return s.Body
}

func (s *MedicalAnswerResponse) SetHeaders(v map[string]*string) *MedicalAnswerResponse {
	s.Headers = v
	return s
}

func (s *MedicalAnswerResponse) SetStatusCode(v int32) *MedicalAnswerResponse {
	s.StatusCode = &v
	return s
}

func (s *MedicalAnswerResponse) SetBody(v *MedicalAnswerResult) *MedicalAnswerResponse {
	s.Body = v
	return s
}

func (s *MedicalAnswerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
