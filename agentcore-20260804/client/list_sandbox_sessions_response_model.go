// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSandboxSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSandboxSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListSandboxSessionsResponseBody) *ListSandboxSessionsResponse
	GetBody() *ListSandboxSessionsResponseBody
}

type ListSandboxSessionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSandboxSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSandboxSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListSandboxSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSandboxSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSandboxSessionsResponse) GetBody() *ListSandboxSessionsResponseBody {
	return s.Body
}

func (s *ListSandboxSessionsResponse) SetHeaders(v map[string]*string) *ListSandboxSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListSandboxSessionsResponse) SetStatusCode(v int32) *ListSandboxSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSandboxSessionsResponse) SetBody(v *ListSandboxSessionsResponseBody) *ListSandboxSessionsResponse {
	s.Body = v
	return s
}

func (s *ListSandboxSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
