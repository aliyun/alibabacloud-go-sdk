// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterResourcePoolPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterResourcePoolPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterResourcePoolPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterResourcePoolPerformanceResponseBody) *DescribeDBClusterResourcePoolPerformanceResponse
	GetBody() *DescribeDBClusterResourcePoolPerformanceResponseBody
}

type DescribeDBClusterResourcePoolPerformanceResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterResourcePoolPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) GetBody() *DescribeDBClusterResourcePoolPerformanceResponseBody {
	return s.Body
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterResourcePoolPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) SetStatusCode(v int32) *DescribeDBClusterResourcePoolPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) SetBody(v *DescribeDBClusterResourcePoolPerformanceResponseBody) *DescribeDBClusterResourcePoolPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
