// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeViewJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeViewJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeViewJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeViewJobsResponseBody) *DescribeViewJobsResponse
	GetBody() *DescribeViewJobsResponseBody
}

type DescribeViewJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeViewJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeViewJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeViewJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeViewJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeViewJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeViewJobsResponse) GetBody() *DescribeViewJobsResponseBody {
	return s.Body
}

func (s *DescribeViewJobsResponse) SetHeaders(v map[string]*string) *DescribeViewJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeViewJobsResponse) SetStatusCode(v int32) *DescribeViewJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeViewJobsResponse) SetBody(v *DescribeViewJobsResponseBody) *DescribeViewJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeViewJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
