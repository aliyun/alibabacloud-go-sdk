// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSandboxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeSandboxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeSandboxResponse
	GetStatusCode() *int32
	SetBody(v *SandboxResult) *ResumeSandboxResponse
	GetBody() *SandboxResult
}

type ResumeSandboxResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SandboxResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeSandboxResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeSandboxResponse) GoString() string {
	return s.String()
}

func (s *ResumeSandboxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeSandboxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeSandboxResponse) GetBody() *SandboxResult {
	return s.Body
}

func (s *ResumeSandboxResponse) SetHeaders(v map[string]*string) *ResumeSandboxResponse {
	s.Headers = v
	return s
}

func (s *ResumeSandboxResponse) SetStatusCode(v int32) *ResumeSandboxResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeSandboxResponse) SetBody(v *SandboxResult) *ResumeSandboxResponse {
	s.Body = v
	return s
}

func (s *ResumeSandboxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
