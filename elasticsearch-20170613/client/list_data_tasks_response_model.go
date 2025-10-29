// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListDataTasksResponseBody) *ListDataTasksResponse
	GetBody() *ListDataTasksResponseBody
}

type ListDataTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataTasksResponse) GetBody() *ListDataTasksResponseBody {
	return s.Body
}

func (s *ListDataTasksResponse) SetHeaders(v map[string]*string) *ListDataTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDataTasksResponse) SetStatusCode(v int32) *ListDataTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataTasksResponse) SetBody(v *ListDataTasksResponseBody) *ListDataTasksResponse {
	s.Body = v
	return s
}

func (s *ListDataTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
