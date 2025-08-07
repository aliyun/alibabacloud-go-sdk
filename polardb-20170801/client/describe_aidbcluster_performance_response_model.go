// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterPerformanceResponseBody) *DescribeAIDBClusterPerformanceResponse
	GetBody() *DescribeAIDBClusterPerformanceResponseBody
}

type DescribeAIDBClusterPerformanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterPerformanceResponse) GetBody() *DescribeAIDBClusterPerformanceResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterPerformanceResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponse) SetStatusCode(v int32) *DescribeAIDBClusterPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponse) SetBody(v *DescribeAIDBClusterPerformanceResponseBody) *DescribeAIDBClusterPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterPerformanceResponse) Validate() error {
	return dara.Validate(s)
}
