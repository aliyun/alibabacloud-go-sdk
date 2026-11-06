// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataProcessJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataProcessJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataProcessJobsResponseBody) *DescribeDataProcessJobsResponse
	GetBody() *DescribeDataProcessJobsResponseBody
}

type DescribeDataProcessJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataProcessJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataProcessJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataProcessJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataProcessJobsResponse) GetBody() *DescribeDataProcessJobsResponseBody {
	return s.Body
}

func (s *DescribeDataProcessJobsResponse) SetHeaders(v map[string]*string) *DescribeDataProcessJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataProcessJobsResponse) SetStatusCode(v int32) *DescribeDataProcessJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataProcessJobsResponse) SetBody(v *DescribeDataProcessJobsResponseBody) *DescribeDataProcessJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataProcessJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
