// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBatchTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBatchTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListBatchTasksResponseBody) *ListBatchTasksResponse
	GetBody() *ListBatchTasksResponseBody
}

type ListBatchTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBatchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBatchTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTasksResponse) GoString() string {
	return s.String()
}

func (s *ListBatchTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBatchTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBatchTasksResponse) GetBody() *ListBatchTasksResponseBody {
	return s.Body
}

func (s *ListBatchTasksResponse) SetHeaders(v map[string]*string) *ListBatchTasksResponse {
	s.Headers = v
	return s
}

func (s *ListBatchTasksResponse) SetStatusCode(v int32) *ListBatchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchTasksResponse) SetBody(v *ListBatchTasksResponseBody) *ListBatchTasksResponse {
	s.Body = v
	return s
}

func (s *ListBatchTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
