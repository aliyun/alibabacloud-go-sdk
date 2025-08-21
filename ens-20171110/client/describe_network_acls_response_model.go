// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAclsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkAclsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkAclsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkAclsResponseBody) *DescribeNetworkAclsResponse
	GetBody() *DescribeNetworkAclsResponseBody
}

type DescribeNetworkAclsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkAclsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkAclsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkAclsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkAclsResponse) GetBody() *DescribeNetworkAclsResponseBody {
	return s.Body
}

func (s *DescribeNetworkAclsResponse) SetHeaders(v map[string]*string) *DescribeNetworkAclsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkAclsResponse) SetStatusCode(v int32) *DescribeNetworkAclsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkAclsResponse) SetBody(v *DescribeNetworkAclsResponseBody) *DescribeNetworkAclsResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkAclsResponse) Validate() error {
	return dara.Validate(s)
}
