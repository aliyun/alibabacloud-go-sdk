// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiagnosisResponseBody) *CreateDiagnosisResponse
	GetBody() *CreateDiagnosisResponseBody
}

type CreateDiagnosisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiagnosisResponse) GetBody() *CreateDiagnosisResponseBody {
	return s.Body
}

func (s *CreateDiagnosisResponse) SetHeaders(v map[string]*string) *CreateDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosisResponse) SetStatusCode(v int32) *CreateDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosisResponse) SetBody(v *CreateDiagnosisResponseBody) *CreateDiagnosisResponse {
	s.Body = v
	return s
}

func (s *CreateDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
