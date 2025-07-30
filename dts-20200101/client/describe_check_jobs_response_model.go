// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckJobsResponseBody) *DescribeCheckJobsResponse
	GetBody() *DescribeCheckJobsResponseBody
}

type DescribeCheckJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckJobsResponse) GetBody() *DescribeCheckJobsResponseBody {
	return s.Body
}

func (s *DescribeCheckJobsResponse) SetHeaders(v map[string]*string) *DescribeCheckJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckJobsResponse) SetStatusCode(v int32) *DescribeCheckJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckJobsResponse) SetBody(v *DescribeCheckJobsResponseBody) *DescribeCheckJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckJobsResponse) Validate() error {
	return dara.Validate(s)
}
