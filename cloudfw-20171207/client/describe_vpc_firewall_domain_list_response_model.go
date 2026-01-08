// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDomainListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallDomainListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallDomainListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallDomainListResponseBody) *DescribeVpcFirewallDomainListResponse
	GetBody() *DescribeVpcFirewallDomainListResponseBody
}

type DescribeVpcFirewallDomainListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallDomainListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallDomainListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallDomainListResponse) GetBody() *DescribeVpcFirewallDomainListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallDomainListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDomainListResponse) SetStatusCode(v int32) *DescribeVpcFirewallDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponse) SetBody(v *DescribeVpcFirewallDomainListResponseBody) *DescribeVpcFirewallDomainListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallDomainListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
