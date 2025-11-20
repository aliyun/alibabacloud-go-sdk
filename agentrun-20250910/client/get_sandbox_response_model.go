// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSandboxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSandboxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSandboxResponse
	GetStatusCode() *int32
	SetBody(v *SandboxResult) *GetSandboxResponse
	GetBody() *SandboxResult
}

type GetSandboxResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SandboxResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSandboxResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSandboxResponse) GoString() string {
	return s.String()
}

func (s *GetSandboxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSandboxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSandboxResponse) GetBody() *SandboxResult {
	return s.Body
}

func (s *GetSandboxResponse) SetHeaders(v map[string]*string) *GetSandboxResponse {
	s.Headers = v
	return s
}

func (s *GetSandboxResponse) SetStatusCode(v int32) *GetSandboxResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSandboxResponse) SetBody(v *SandboxResult) *GetSandboxResponse {
	s.Body = v
	return s
}

func (s *GetSandboxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
