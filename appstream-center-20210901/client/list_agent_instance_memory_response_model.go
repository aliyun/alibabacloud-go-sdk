// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstanceMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentInstanceMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentInstanceMemoryResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentInstanceMemoryResponseBody) *ListAgentInstanceMemoryResponse
	GetBody() *ListAgentInstanceMemoryResponseBody
}

type ListAgentInstanceMemoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentInstanceMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentInstanceMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceMemoryResponse) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentInstanceMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentInstanceMemoryResponse) GetBody() *ListAgentInstanceMemoryResponseBody {
	return s.Body
}

func (s *ListAgentInstanceMemoryResponse) SetHeaders(v map[string]*string) *ListAgentInstanceMemoryResponse {
	s.Headers = v
	return s
}

func (s *ListAgentInstanceMemoryResponse) SetStatusCode(v int32) *ListAgentInstanceMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentInstanceMemoryResponse) SetBody(v *ListAgentInstanceMemoryResponseBody) *ListAgentInstanceMemoryResponse {
	s.Body = v
	return s
}

func (s *ListAgentInstanceMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
