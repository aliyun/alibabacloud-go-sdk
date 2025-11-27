// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceMetricsResponseBody) *DescribeDBInstanceMetricsResponse
	GetBody() *DescribeDBInstanceMetricsResponseBody
}

type DescribeDBInstanceMetricsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceMetricsResponse) GetBody() *DescribeDBInstanceMetricsResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceMetricsResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceMetricsResponse) SetStatusCode(v int32) *DescribeDBInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponse) SetBody(v *DescribeDBInstanceMetricsResponseBody) *DescribeDBInstanceMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
