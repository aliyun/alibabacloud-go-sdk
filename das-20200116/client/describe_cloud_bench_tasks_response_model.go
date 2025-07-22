// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudBenchTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudBenchTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudBenchTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudBenchTasksResponseBody) *DescribeCloudBenchTasksResponse
	GetBody() *DescribeCloudBenchTasksResponseBody
}

type DescribeCloudBenchTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudBenchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudBenchTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudBenchTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudBenchTasksResponse) GetBody() *DescribeCloudBenchTasksResponseBody {
	return s.Body
}

func (s *DescribeCloudBenchTasksResponse) SetHeaders(v map[string]*string) *DescribeCloudBenchTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudBenchTasksResponse) SetStatusCode(v int32) *DescribeCloudBenchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudBenchTasksResponse) SetBody(v *DescribeCloudBenchTasksResponseBody) *DescribeCloudBenchTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudBenchTasksResponse) Validate() error {
	return dara.Validate(s)
}
