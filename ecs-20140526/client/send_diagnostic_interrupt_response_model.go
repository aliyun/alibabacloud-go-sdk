// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDiagnosticInterruptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendDiagnosticInterruptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendDiagnosticInterruptResponse
	GetStatusCode() *int32
	SetBody(v *SendDiagnosticInterruptResponseBody) *SendDiagnosticInterruptResponse
	GetBody() *SendDiagnosticInterruptResponseBody
}

type SendDiagnosticInterruptResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendDiagnosticInterruptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendDiagnosticInterruptResponse) String() string {
	return dara.Prettify(s)
}

func (s SendDiagnosticInterruptResponse) GoString() string {
	return s.String()
}

func (s *SendDiagnosticInterruptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendDiagnosticInterruptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendDiagnosticInterruptResponse) GetBody() *SendDiagnosticInterruptResponseBody {
	return s.Body
}

func (s *SendDiagnosticInterruptResponse) SetHeaders(v map[string]*string) *SendDiagnosticInterruptResponse {
	s.Headers = v
	return s
}

func (s *SendDiagnosticInterruptResponse) SetStatusCode(v int32) *SendDiagnosticInterruptResponse {
	s.StatusCode = &v
	return s
}

func (s *SendDiagnosticInterruptResponse) SetBody(v *SendDiagnosticInterruptResponseBody) *SendDiagnosticInterruptResponse {
	s.Body = v
	return s
}

func (s *SendDiagnosticInterruptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
