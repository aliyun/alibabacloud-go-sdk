// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse
	GetBody() *DescribeDBClusterPerformanceResponseBody
}

type DescribeDBClusterPerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterPerformanceResponse) GetBody() *DescribeDBClusterPerformanceResponseBody {
	return s.Body
}

func (s *DescribeDBClusterPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetStatusCode(v int32) *DescribeDBClusterPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
