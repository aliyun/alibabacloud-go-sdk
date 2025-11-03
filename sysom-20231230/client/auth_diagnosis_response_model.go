// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *AuthDiagnosisResponseBody) *AuthDiagnosisResponse
	GetBody() *AuthDiagnosisResponseBody
}

type AuthDiagnosisResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthDiagnosisResponse) GetBody() *AuthDiagnosisResponseBody {
	return s.Body
}

func (s *AuthDiagnosisResponse) SetHeaders(v map[string]*string) *AuthDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *AuthDiagnosisResponse) SetStatusCode(v int32) *AuthDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthDiagnosisResponse) SetBody(v *AuthDiagnosisResponseBody) *AuthDiagnosisResponse {
	s.Body = v
	return s
}

func (s *AuthDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
