// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterTasksResponseBody) *DescribeClusterTasksResponse
	GetBody() *DescribeClusterTasksResponseBody
}

type DescribeClusterTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterTasksResponse) GetBody() *DescribeClusterTasksResponseBody {
	return s.Body
}

func (s *DescribeClusterTasksResponse) SetHeaders(v map[string]*string) *DescribeClusterTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterTasksResponse) SetStatusCode(v int32) *DescribeClusterTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterTasksResponse) SetBody(v *DescribeClusterTasksResponseBody) *DescribeClusterTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
