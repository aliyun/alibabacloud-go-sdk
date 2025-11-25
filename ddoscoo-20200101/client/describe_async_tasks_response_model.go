// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAsyncTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAsyncTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAsyncTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAsyncTasksResponseBody) *DescribeAsyncTasksResponse
	GetBody() *DescribeAsyncTasksResponseBody
}

type DescribeAsyncTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAsyncTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAsyncTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAsyncTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAsyncTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAsyncTasksResponse) GetBody() *DescribeAsyncTasksResponseBody {
	return s.Body
}

func (s *DescribeAsyncTasksResponse) SetHeaders(v map[string]*string) *DescribeAsyncTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAsyncTasksResponse) SetStatusCode(v int32) *DescribeAsyncTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAsyncTasksResponse) SetBody(v *DescribeAsyncTasksResponseBody) *DescribeAsyncTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeAsyncTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
