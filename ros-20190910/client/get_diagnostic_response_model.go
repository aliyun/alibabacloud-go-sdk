// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiagnosticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDiagnosticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDiagnosticResponse
	GetStatusCode() *int32
	SetBody(v *GetDiagnosticResponseBody) *GetDiagnosticResponse
	GetBody() *GetDiagnosticResponseBody
}

type GetDiagnosticResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiagnosticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiagnosticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosticResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDiagnosticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDiagnosticResponse) GetBody() *GetDiagnosticResponseBody {
	return s.Body
}

func (s *GetDiagnosticResponse) SetHeaders(v map[string]*string) *GetDiagnosticResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosticResponse) SetStatusCode(v int32) *GetDiagnosticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiagnosticResponse) SetBody(v *GetDiagnosticResponseBody) *GetDiagnosticResponse {
	s.Body = v
	return s
}

func (s *GetDiagnosticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
