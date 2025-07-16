// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRefreshTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRefreshTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRefreshTasksResponseBody) *DescribeRefreshTasksResponse
	GetBody() *DescribeRefreshTasksResponseBody
}

type DescribeRefreshTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRefreshTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRefreshTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRefreshTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRefreshTasksResponse) GetBody() *DescribeRefreshTasksResponseBody {
	return s.Body
}

func (s *DescribeRefreshTasksResponse) SetHeaders(v map[string]*string) *DescribeRefreshTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeRefreshTasksResponse) SetStatusCode(v int32) *DescribeRefreshTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRefreshTasksResponse) SetBody(v *DescribeRefreshTasksResponseBody) *DescribeRefreshTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeRefreshTasksResponse) Validate() error {
	return dara.Validate(s)
}
