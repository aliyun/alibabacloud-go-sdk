// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentMemoryResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentMemoryResponseBody) *ListDataAgentMemoryResponse
	GetBody() *ListDataAgentMemoryResponseBody
}

type ListDataAgentMemoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMemoryResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentMemoryResponse) GetBody() *ListDataAgentMemoryResponseBody {
	return s.Body
}

func (s *ListDataAgentMemoryResponse) SetHeaders(v map[string]*string) *ListDataAgentMemoryResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentMemoryResponse) SetStatusCode(v int32) *ListDataAgentMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentMemoryResponse) SetBody(v *ListDataAgentMemoryResponseBody) *ListDataAgentMemoryResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
