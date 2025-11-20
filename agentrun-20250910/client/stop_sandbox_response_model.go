// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSandboxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopSandboxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopSandboxResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSandboxResult) *StopSandboxResponse
	GetBody() *DeleteSandboxResult
}

type StopSandboxResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSandboxResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSandboxResponse) String() string {
	return dara.Prettify(s)
}

func (s StopSandboxResponse) GoString() string {
	return s.String()
}

func (s *StopSandboxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopSandboxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopSandboxResponse) GetBody() *DeleteSandboxResult {
	return s.Body
}

func (s *StopSandboxResponse) SetHeaders(v map[string]*string) *StopSandboxResponse {
	s.Headers = v
	return s
}

func (s *StopSandboxResponse) SetStatusCode(v int32) *StopSandboxResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSandboxResponse) SetBody(v *DeleteSandboxResult) *StopSandboxResponse {
	s.Body = v
	return s
}

func (s *StopSandboxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
