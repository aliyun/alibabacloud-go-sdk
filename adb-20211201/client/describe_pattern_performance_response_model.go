// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePatternPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePatternPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePatternPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribePatternPerformanceResponseBody) *DescribePatternPerformanceResponse
	GetBody() *DescribePatternPerformanceResponseBody
}

type DescribePatternPerformanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePatternPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePatternPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePatternPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePatternPerformanceResponse) GetBody() *DescribePatternPerformanceResponseBody {
	return s.Body
}

func (s *DescribePatternPerformanceResponse) SetHeaders(v map[string]*string) *DescribePatternPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribePatternPerformanceResponse) SetStatusCode(v int32) *DescribePatternPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePatternPerformanceResponse) SetBody(v *DescribePatternPerformanceResponseBody) *DescribePatternPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribePatternPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
