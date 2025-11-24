// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSandboxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSandboxResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSandboxResult) *DeleteSandboxResponse
	GetBody() *DeleteSandboxResult
}

type DeleteSandboxResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSandboxResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSandboxResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxResponse) GoString() string {
	return s.String()
}

func (s *DeleteSandboxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSandboxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSandboxResponse) GetBody() *DeleteSandboxResult {
	return s.Body
}

func (s *DeleteSandboxResponse) SetHeaders(v map[string]*string) *DeleteSandboxResponse {
	s.Headers = v
	return s
}

func (s *DeleteSandboxResponse) SetStatusCode(v int32) *DeleteSandboxResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSandboxResponse) SetBody(v *DeleteSandboxResult) *DeleteSandboxResponse {
	s.Body = v
	return s
}

func (s *DeleteSandboxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
