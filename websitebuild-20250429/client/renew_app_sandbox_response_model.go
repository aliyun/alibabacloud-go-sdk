// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppSandboxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewAppSandboxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewAppSandboxResponse
	GetStatusCode() *int32
	SetBody(v *RenewAppSandboxResponseBody) *RenewAppSandboxResponse
	GetBody() *RenewAppSandboxResponseBody
}

type RenewAppSandboxResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAppSandboxResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAppSandboxResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewAppSandboxResponse) GoString() string {
	return s.String()
}

func (s *RenewAppSandboxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewAppSandboxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewAppSandboxResponse) GetBody() *RenewAppSandboxResponseBody {
	return s.Body
}

func (s *RenewAppSandboxResponse) SetHeaders(v map[string]*string) *RenewAppSandboxResponse {
	s.Headers = v
	return s
}

func (s *RenewAppSandboxResponse) SetStatusCode(v int32) *RenewAppSandboxResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAppSandboxResponse) SetBody(v *RenewAppSandboxResponseBody) *RenewAppSandboxResponse {
	s.Body = v
	return s
}

func (s *RenewAppSandboxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
