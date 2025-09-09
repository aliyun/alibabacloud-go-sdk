// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDSPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRDSPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRDSPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRDSPerformanceResponseBody) *DescribeRDSPerformanceResponse
	GetBody() *DescribeRDSPerformanceResponseBody
}

type DescribeRDSPerformanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRDSPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRDSPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDSPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRDSPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRDSPerformanceResponse) GetBody() *DescribeRDSPerformanceResponseBody {
	return s.Body
}

func (s *DescribeRDSPerformanceResponse) SetHeaders(v map[string]*string) *DescribeRDSPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRDSPerformanceResponse) SetStatusCode(v int32) *DescribeRDSPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRDSPerformanceResponse) SetBody(v *DescribeRDSPerformanceResponseBody) *DescribeRDSPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeRDSPerformanceResponse) Validate() error {
	return dara.Validate(s)
}
