// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemorySessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemorySessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemorySessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListMemorySessionsResponseBody) *ListMemorySessionsResponse
	GetBody() *ListMemorySessionsResponseBody
}

type ListMemorySessionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemorySessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemorySessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionsResponse) GoString() string {
	return s.String()
}

func (s *ListMemorySessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemorySessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemorySessionsResponse) GetBody() *ListMemorySessionsResponseBody {
	return s.Body
}

func (s *ListMemorySessionsResponse) SetHeaders(v map[string]*string) *ListMemorySessionsResponse {
	s.Headers = v
	return s
}

func (s *ListMemorySessionsResponse) SetStatusCode(v int32) *ListMemorySessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemorySessionsResponse) SetBody(v *ListMemorySessionsResponseBody) *ListMemorySessionsResponse {
	s.Body = v
	return s
}

func (s *ListMemorySessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
