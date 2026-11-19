// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemoryTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemoryTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListMemoryTasksResponseBody) *ListMemoryTasksResponse
	GetBody() *ListMemoryTasksResponseBody
}

type ListMemoryTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoryTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoryTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMemoryTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemoryTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemoryTasksResponse) GetBody() *ListMemoryTasksResponseBody {
	return s.Body
}

func (s *ListMemoryTasksResponse) SetHeaders(v map[string]*string) *ListMemoryTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMemoryTasksResponse) SetStatusCode(v int32) *ListMemoryTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoryTasksResponse) SetBody(v *ListMemoryTasksResponseBody) *ListMemoryTasksResponse {
	s.Body = v
	return s
}

func (s *ListMemoryTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
