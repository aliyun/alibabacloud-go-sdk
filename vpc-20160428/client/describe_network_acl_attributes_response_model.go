// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAclAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkAclAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkAclAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkAclAttributesResponseBody) *DescribeNetworkAclAttributesResponse
	GetBody() *DescribeNetworkAclAttributesResponseBody
}

type DescribeNetworkAclAttributesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkAclAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkAclAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkAclAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkAclAttributesResponse) GetBody() *DescribeNetworkAclAttributesResponseBody {
	return s.Body
}

func (s *DescribeNetworkAclAttributesResponse) SetHeaders(v map[string]*string) *DescribeNetworkAclAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkAclAttributesResponse) SetStatusCode(v int32) *DescribeNetworkAclAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponse) SetBody(v *DescribeNetworkAclAttributesResponseBody) *DescribeNetworkAclAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkAclAttributesResponse) Validate() error {
	return dara.Validate(s)
}
