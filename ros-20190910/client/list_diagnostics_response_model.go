// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiagnosticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiagnosticsResponse
	GetStatusCode() *int32
	SetBody(v *ListDiagnosticsResponseBody) *ListDiagnosticsResponse
	GetBody() *ListDiagnosticsResponseBody
}

type ListDiagnosticsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnosticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnosticsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiagnosticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiagnosticsResponse) GetBody() *ListDiagnosticsResponseBody {
	return s.Body
}

func (s *ListDiagnosticsResponse) SetHeaders(v map[string]*string) *ListDiagnosticsResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosticsResponse) SetStatusCode(v int32) *ListDiagnosticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnosticsResponse) SetBody(v *ListDiagnosticsResponseBody) *ListDiagnosticsResponse {
	s.Body = v
	return s
}

func (s *ListDiagnosticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
