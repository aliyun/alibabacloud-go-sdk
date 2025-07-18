// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIndexUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceIndexUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceIndexUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceIndexUsageResponseBody) *DescribeDBInstanceIndexUsageResponse
	GetBody() *DescribeDBInstanceIndexUsageResponseBody
}

type DescribeDBInstanceIndexUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceIndexUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceIndexUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceIndexUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceIndexUsageResponse) GetBody() *DescribeDBInstanceIndexUsageResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceIndexUsageResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceIndexUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponse) SetStatusCode(v int32) *DescribeDBInstanceIndexUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponse) SetBody(v *DescribeDBInstanceIndexUsageResponseBody) *DescribeDBInstanceIndexUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponse) Validate() error {
	return dara.Validate(s)
}
