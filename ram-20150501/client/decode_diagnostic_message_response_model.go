// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecodeDiagnosticMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DecodeDiagnosticMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DecodeDiagnosticMessageResponse
	GetStatusCode() *int32
	SetBody(v *DecodeDiagnosticMessageResponseBody) *DecodeDiagnosticMessageResponse
	GetBody() *DecodeDiagnosticMessageResponseBody
}

type DecodeDiagnosticMessageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecodeDiagnosticMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecodeDiagnosticMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s DecodeDiagnosticMessageResponse) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DecodeDiagnosticMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DecodeDiagnosticMessageResponse) GetBody() *DecodeDiagnosticMessageResponseBody {
	return s.Body
}

func (s *DecodeDiagnosticMessageResponse) SetHeaders(v map[string]*string) *DecodeDiagnosticMessageResponse {
	s.Headers = v
	return s
}

func (s *DecodeDiagnosticMessageResponse) SetStatusCode(v int32) *DecodeDiagnosticMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DecodeDiagnosticMessageResponse) SetBody(v *DecodeDiagnosticMessageResponseBody) *DecodeDiagnosticMessageResponse {
	s.Body = v
	return s
}

func (s *DecodeDiagnosticMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
