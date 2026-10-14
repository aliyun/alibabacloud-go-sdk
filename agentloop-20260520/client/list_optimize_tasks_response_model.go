// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOptimizeTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOptimizeTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOptimizeTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListOptimizeTasksResponseBody) *ListOptimizeTasksResponse
	GetBody() *ListOptimizeTasksResponseBody
}

type ListOptimizeTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOptimizeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOptimizeTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOptimizeTasksResponse) GoString() string {
	return s.String()
}

func (s *ListOptimizeTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOptimizeTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOptimizeTasksResponse) GetBody() *ListOptimizeTasksResponseBody {
	return s.Body
}

func (s *ListOptimizeTasksResponse) SetHeaders(v map[string]*string) *ListOptimizeTasksResponse {
	s.Headers = v
	return s
}

func (s *ListOptimizeTasksResponse) SetStatusCode(v int32) *ListOptimizeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOptimizeTasksResponse) SetBody(v *ListOptimizeTasksResponseBody) *ListOptimizeTasksResponse {
	s.Body = v
	return s
}

func (s *ListOptimizeTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
