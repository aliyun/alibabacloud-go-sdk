// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceSSLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceSSLResponseBody) *DescribeDBInstanceSSLResponse
	GetBody() *DescribeDBInstanceSSLResponseBody
}

type DescribeDBInstanceSSLResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceSSLResponse) GetBody() *DescribeDBInstanceSSLResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceSSLResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetStatusCode(v int32) *DescribeDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetBody(v *DescribeDBInstanceSSLResponseBody) *DescribeDBInstanceSSLResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceSSLResponse) Validate() error {
	return dara.Validate(s)
}
