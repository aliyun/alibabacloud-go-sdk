// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSharePerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataSharePerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataSharePerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataSharePerformanceResponseBody) *DescribeDataSharePerformanceResponse
	GetBody() *DescribeDataSharePerformanceResponseBody
}

type DescribeDataSharePerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSharePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSharePerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSharePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataSharePerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataSharePerformanceResponse) GetBody() *DescribeDataSharePerformanceResponseBody {
	return s.Body
}

func (s *DescribeDataSharePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDataSharePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSharePerformanceResponse) SetStatusCode(v int32) *DescribeDataSharePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSharePerformanceResponse) SetBody(v *DescribeDataSharePerformanceResponseBody) *DescribeDataSharePerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDataSharePerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
