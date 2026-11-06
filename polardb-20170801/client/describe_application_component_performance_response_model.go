// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationComponentPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationComponentPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationComponentPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationComponentPerformanceResponseBody) *DescribeApplicationComponentPerformanceResponse
	GetBody() *DescribeApplicationComponentPerformanceResponseBody
}

type DescribeApplicationComponentPerformanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationComponentPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationComponentPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationComponentPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationComponentPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationComponentPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationComponentPerformanceResponse) GetBody() *DescribeApplicationComponentPerformanceResponseBody {
	return s.Body
}

func (s *DescribeApplicationComponentPerformanceResponse) SetHeaders(v map[string]*string) *DescribeApplicationComponentPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponse) SetStatusCode(v int32) *DescribeApplicationComponentPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponse) SetBody(v *DescribeApplicationComponentPerformanceResponseBody) *DescribeApplicationComponentPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationComponentPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
