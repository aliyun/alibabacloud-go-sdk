// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceUsageResponseBody) *DescribeDBInstanceUsageResponse
	GetBody() *DescribeDBInstanceUsageResponseBody
}

type DescribeDBInstanceUsageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceUsageResponse) GetBody() *DescribeDBInstanceUsageResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceUsageResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceUsageResponse) SetStatusCode(v int32) *DescribeDBInstanceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceUsageResponse) SetBody(v *DescribeDBInstanceUsageResponseBody) *DescribeDBInstanceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
