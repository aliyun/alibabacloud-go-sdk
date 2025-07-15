// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailoverTestJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFailoverTestJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFailoverTestJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFailoverTestJobsResponseBody) *DescribeFailoverTestJobsResponse
	GetBody() *DescribeFailoverTestJobsResponseBody
}

type DescribeFailoverTestJobsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFailoverTestJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFailoverTestJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFailoverTestJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFailoverTestJobsResponse) GetBody() *DescribeFailoverTestJobsResponseBody {
	return s.Body
}

func (s *DescribeFailoverTestJobsResponse) SetHeaders(v map[string]*string) *DescribeFailoverTestJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFailoverTestJobsResponse) SetStatusCode(v int32) *DescribeFailoverTestJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFailoverTestJobsResponse) SetBody(v *DescribeFailoverTestJobsResponseBody) *DescribeFailoverTestJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeFailoverTestJobsResponse) Validate() error {
	return dara.Validate(s)
}
