// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSandboxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSandboxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSandboxResponse
	GetStatusCode() *int32
	SetBody(v *SandboxResult) *CreateSandboxResponse
	GetBody() *SandboxResult
}

type CreateSandboxResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SandboxResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSandboxResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSandboxResponse) GoString() string {
	return s.String()
}

func (s *CreateSandboxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSandboxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSandboxResponse) GetBody() *SandboxResult {
	return s.Body
}

func (s *CreateSandboxResponse) SetHeaders(v map[string]*string) *CreateSandboxResponse {
	s.Headers = v
	return s
}

func (s *CreateSandboxResponse) SetStatusCode(v int32) *CreateSandboxResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSandboxResponse) SetBody(v *SandboxResult) *CreateSandboxResponse {
	s.Body = v
	return s
}

func (s *CreateSandboxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
