// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallAclGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallAclGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallAclGroupListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallAclGroupListResponseBody) *DescribeNatFirewallAclGroupListResponse
	GetBody() *DescribeNatFirewallAclGroupListResponseBody
}

type DescribeNatFirewallAclGroupListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallAclGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallAclGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallAclGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallAclGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallAclGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallAclGroupListResponse) GetBody() *DescribeNatFirewallAclGroupListResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallAclGroupListResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallAclGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponse) SetStatusCode(v int32) *DescribeNatFirewallAclGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponse) SetBody(v *DescribeNatFirewallAclGroupListResponseBody) *DescribeNatFirewallAclGroupListResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
