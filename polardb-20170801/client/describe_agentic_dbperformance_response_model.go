// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBPerformanceResponseBody) *DescribeAgenticDBPerformanceResponse
	GetBody() *DescribeAgenticDBPerformanceResponseBody
}

type DescribeAgenticDBPerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBPerformanceResponse) GetBody() *DescribeAgenticDBPerformanceResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBPerformanceResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBPerformanceResponse) SetStatusCode(v int32) *DescribeAgenticDBPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBPerformanceResponse) SetBody(v *DescribeAgenticDBPerformanceResponseBody) *DescribeAgenticDBPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
