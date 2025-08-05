// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAclGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallAclGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallAclGroupListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallAclGroupListResponseBody) *DescribeVpcFirewallAclGroupListResponse
	GetBody() *DescribeVpcFirewallAclGroupListResponseBody
}

type DescribeVpcFirewallAclGroupListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallAclGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallAclGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallAclGroupListResponse) GetBody() *DescribeVpcFirewallAclGroupListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallAclGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetStatusCode(v int32) *DescribeVpcFirewallAclGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetBody(v *DescribeVpcFirewallAclGroupListResponseBody) *DescribeVpcFirewallAclGroupListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponse) Validate() error {
	return dara.Validate(s)
}
