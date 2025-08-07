// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterSSLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterSSLResponseBody) *DescribeDBClusterSSLResponse
	GetBody() *DescribeDBClusterSSLResponseBody
}

type DescribeDBClusterSSLResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterSSLResponse) GetBody() *DescribeDBClusterSSLResponseBody {
	return s.Body
}

func (s *DescribeDBClusterSSLResponse) SetHeaders(v map[string]*string) *DescribeDBClusterSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterSSLResponse) SetStatusCode(v int32) *DescribeDBClusterSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterSSLResponse) SetBody(v *DescribeDBClusterSSLResponseBody) *DescribeDBClusterSSLResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterSSLResponse) Validate() error {
	return dara.Validate(s)
}
