// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationPerformanceResponseBody) *DescribeApplicationPerformanceResponse
	GetBody() *DescribeApplicationPerformanceResponseBody
}

type DescribeApplicationPerformanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationPerformanceResponse) GetBody() *DescribeApplicationPerformanceResponseBody {
	return s.Body
}

func (s *DescribeApplicationPerformanceResponse) SetHeaders(v map[string]*string) *DescribeApplicationPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationPerformanceResponse) SetStatusCode(v int32) *DescribeApplicationPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationPerformanceResponse) SetBody(v *DescribeApplicationPerformanceResponseBody) *DescribeApplicationPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
