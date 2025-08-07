// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodePerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBNodePerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBNodePerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBNodePerformanceResponseBody) *DescribeDBNodePerformanceResponse
	GetBody() *DescribeDBNodePerformanceResponseBody
}

type DescribeDBNodePerformanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBNodePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBNodePerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBNodePerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBNodePerformanceResponse) GetBody() *DescribeDBNodePerformanceResponseBody {
	return s.Body
}

func (s *DescribeDBNodePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBNodePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetStatusCode(v int32) *DescribeDBNodePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetBody(v *DescribeDBNodePerformanceResponseBody) *DescribeDBNodePerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDBNodePerformanceResponse) Validate() error {
	return dara.Validate(s)
}
