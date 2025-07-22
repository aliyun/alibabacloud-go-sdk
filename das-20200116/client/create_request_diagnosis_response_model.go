// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRequestDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRequestDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRequestDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *CreateRequestDiagnosisResponseBody) *CreateRequestDiagnosisResponse
	GetBody() *CreateRequestDiagnosisResponseBody
}

type CreateRequestDiagnosisResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRequestDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRequestDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRequestDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CreateRequestDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRequestDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRequestDiagnosisResponse) GetBody() *CreateRequestDiagnosisResponseBody {
	return s.Body
}

func (s *CreateRequestDiagnosisResponse) SetHeaders(v map[string]*string) *CreateRequestDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CreateRequestDiagnosisResponse) SetStatusCode(v int32) *CreateRequestDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRequestDiagnosisResponse) SetBody(v *CreateRequestDiagnosisResponseBody) *CreateRequestDiagnosisResponse {
	s.Body = v
	return s
}

func (s *CreateRequestDiagnosisResponse) Validate() error {
	return dara.Validate(s)
}
