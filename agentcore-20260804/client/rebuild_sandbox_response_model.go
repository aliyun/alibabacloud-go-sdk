// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildSandboxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebuildSandboxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebuildSandboxResponse
	GetStatusCode() *int32
	SetBody(v *RebuildSandboxResponseBody) *RebuildSandboxResponse
	GetBody() *RebuildSandboxResponseBody
}

type RebuildSandboxResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebuildSandboxResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebuildSandboxResponse) String() string {
	return dara.Prettify(s)
}

func (s RebuildSandboxResponse) GoString() string {
	return s.String()
}

func (s *RebuildSandboxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebuildSandboxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebuildSandboxResponse) GetBody() *RebuildSandboxResponseBody {
	return s.Body
}

func (s *RebuildSandboxResponse) SetHeaders(v map[string]*string) *RebuildSandboxResponse {
	s.Headers = v
	return s
}

func (s *RebuildSandboxResponse) SetStatusCode(v int32) *RebuildSandboxResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildSandboxResponse) SetBody(v *RebuildSandboxResponseBody) *RebuildSandboxResponse {
	s.Body = v
	return s
}

func (s *RebuildSandboxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
