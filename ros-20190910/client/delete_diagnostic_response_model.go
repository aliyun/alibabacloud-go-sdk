// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDiagnosticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDiagnosticResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDiagnosticResponseBody) *DeleteDiagnosticResponse
	GetBody() *DeleteDiagnosticResponseBody
}

type DeleteDiagnosticResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiagnosticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiagnosticResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDiagnosticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDiagnosticResponse) GetBody() *DeleteDiagnosticResponseBody {
	return s.Body
}

func (s *DeleteDiagnosticResponse) SetHeaders(v map[string]*string) *DeleteDiagnosticResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiagnosticResponse) SetStatusCode(v int32) *DeleteDiagnosticResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiagnosticResponse) SetBody(v *DeleteDiagnosticResponseBody) *DeleteDiagnosticResponse {
	s.Body = v
	return s
}

func (s *DeleteDiagnosticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
