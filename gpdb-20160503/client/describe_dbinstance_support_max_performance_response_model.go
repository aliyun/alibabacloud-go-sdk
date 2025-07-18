// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSupportMaxPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceSupportMaxPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceSupportMaxPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceSupportMaxPerformanceResponseBody) *DescribeDBInstanceSupportMaxPerformanceResponse
	GetBody() *DescribeDBInstanceSupportMaxPerformanceResponseBody
}

type DescribeDBInstanceSupportMaxPerformanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceSupportMaxPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceSupportMaxPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSupportMaxPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponse) GetBody() *DescribeDBInstanceSupportMaxPerformanceResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSupportMaxPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponse) SetStatusCode(v int32) *DescribeDBInstanceSupportMaxPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponse) SetBody(v *DescribeDBInstanceSupportMaxPerformanceResponseBody) *DescribeDBInstanceSupportMaxPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponse) Validate() error {
	return dara.Validate(s)
}
