// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiagnosticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiagnosticResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiagnosticResponseBody) *CreateDiagnosticResponse
	GetBody() *CreateDiagnosticResponseBody
}

type CreateDiagnosticResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosticResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiagnosticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiagnosticResponse) GetBody() *CreateDiagnosticResponseBody {
	return s.Body
}

func (s *CreateDiagnosticResponse) SetHeaders(v map[string]*string) *CreateDiagnosticResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticResponse) SetStatusCode(v int32) *CreateDiagnosticResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosticResponse) SetBody(v *CreateDiagnosticResponseBody) *CreateDiagnosticResponse {
	s.Body = v
	return s
}

func (s *CreateDiagnosticResponse) Validate() error {
	return dara.Validate(s)
}
