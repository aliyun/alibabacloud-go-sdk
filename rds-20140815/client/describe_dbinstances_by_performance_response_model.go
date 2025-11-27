// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesByPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancesByPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancesByPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancesByPerformanceResponseBody) *DescribeDBInstancesByPerformanceResponse
	GetBody() *DescribeDBInstancesByPerformanceResponseBody
}

type DescribeDBInstancesByPerformanceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesByPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancesByPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancesByPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancesByPerformanceResponse) GetBody() *DescribeDBInstancesByPerformanceResponseBody {
	return s.Body
}

func (s *DescribeDBInstancesByPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesByPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponse) SetStatusCode(v int32) *DescribeDBInstancesByPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponse) SetBody(v *DescribeDBInstancesByPerformanceResponseBody) *DescribeDBInstancesByPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
