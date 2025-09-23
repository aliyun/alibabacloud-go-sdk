// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterProxyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterProxyResponseBody) *DescribeDBClusterProxyResponse
	GetBody() *DescribeDBClusterProxyResponseBody
}

type DescribeDBClusterProxyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterProxyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterProxyResponse) GetBody() *DescribeDBClusterProxyResponseBody {
	return s.Body
}

func (s *DescribeDBClusterProxyResponse) SetHeaders(v map[string]*string) *DescribeDBClusterProxyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterProxyResponse) SetStatusCode(v int32) *DescribeDBClusterProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterProxyResponse) SetBody(v *DescribeDBClusterProxyResponseBody) *DescribeDBClusterProxyResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterProxyResponse) Validate() error {
	return dara.Validate(s)
}
