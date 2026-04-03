// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSandboxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseSandboxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseSandboxResponse
	GetStatusCode() *int32
	SetBody(v *SandboxResult) *PauseSandboxResponse
	GetBody() *SandboxResult
}

type PauseSandboxResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SandboxResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseSandboxResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseSandboxResponse) GoString() string {
	return s.String()
}

func (s *PauseSandboxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseSandboxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseSandboxResponse) GetBody() *SandboxResult {
	return s.Body
}

func (s *PauseSandboxResponse) SetHeaders(v map[string]*string) *PauseSandboxResponse {
	s.Headers = v
	return s
}

func (s *PauseSandboxResponse) SetStatusCode(v int32) *PauseSandboxResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseSandboxResponse) SetBody(v *SandboxResult) *PauseSandboxResponse {
	s.Body = v
	return s
}

func (s *PauseSandboxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
