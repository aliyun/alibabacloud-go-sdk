// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemoryResponse
	GetStatusCode() *int32
	SetBody(v *ListMemoryResponseBody) *ListMemoryResponse
	GetBody() *ListMemoryResponseBody
}

type ListMemoryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryResponse) GoString() string {
	return s.String()
}

func (s *ListMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemoryResponse) GetBody() *ListMemoryResponseBody {
	return s.Body
}

func (s *ListMemoryResponse) SetHeaders(v map[string]*string) *ListMemoryResponse {
	s.Headers = v
	return s
}

func (s *ListMemoryResponse) SetStatusCode(v int32) *ListMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoryResponse) SetBody(v *ListMemoryResponseBody) *ListMemoryResponse {
	s.Body = v
	return s
}

func (s *ListMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
