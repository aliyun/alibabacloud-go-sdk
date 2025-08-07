// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBProxyPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBProxyPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBProxyPerformanceResponseBody) *DescribeDBProxyPerformanceResponse
	GetBody() *DescribeDBProxyPerformanceResponseBody
}

type DescribeDBProxyPerformanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBProxyPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBProxyPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBProxyPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBProxyPerformanceResponse) GetBody() *DescribeDBProxyPerformanceResponseBody {
	return s.Body
}

func (s *DescribeDBProxyPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBProxyPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBProxyPerformanceResponse) SetStatusCode(v int32) *DescribeDBProxyPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponse) SetBody(v *DescribeDBProxyPerformanceResponseBody) *DescribeDBProxyPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDBProxyPerformanceResponse) Validate() error {
	return dara.Validate(s)
}
