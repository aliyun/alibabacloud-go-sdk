// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerformanceViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePerformanceViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePerformanceViewsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePerformanceViewsResponseBody) *DescribePerformanceViewsResponse
	GetBody() *DescribePerformanceViewsResponseBody
}

type DescribePerformanceViewsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePerformanceViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePerformanceViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewsResponse) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePerformanceViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePerformanceViewsResponse) GetBody() *DescribePerformanceViewsResponseBody {
	return s.Body
}

func (s *DescribePerformanceViewsResponse) SetHeaders(v map[string]*string) *DescribePerformanceViewsResponse {
	s.Headers = v
	return s
}

func (s *DescribePerformanceViewsResponse) SetStatusCode(v int32) *DescribePerformanceViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePerformanceViewsResponse) SetBody(v *DescribePerformanceViewsResponseBody) *DescribePerformanceViewsResponse {
	s.Body = v
	return s
}

func (s *DescribePerformanceViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
