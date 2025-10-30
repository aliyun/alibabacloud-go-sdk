// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancePerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancePerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancePerformanceResponseBody) *DescribeDBInstancePerformanceResponse
	GetBody() *DescribeDBInstancePerformanceResponseBody
}

type DescribeDBInstancePerformanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancePerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancePerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancePerformanceResponse) GetBody() *DescribeDBInstancePerformanceResponseBody {
	return s.Body
}

func (s *DescribeDBInstancePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBInstancePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetStatusCode(v int32) *DescribeDBInstancePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetBody(v *DescribeDBInstancePerformanceResponseBody) *DescribeDBInstancePerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
