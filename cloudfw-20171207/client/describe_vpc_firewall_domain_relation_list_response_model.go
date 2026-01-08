// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDomainRelationListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallDomainRelationListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallDomainRelationListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallDomainRelationListResponseBody) *DescribeVpcFirewallDomainRelationListResponse
	GetBody() *DescribeVpcFirewallDomainRelationListResponseBody
}

type DescribeVpcFirewallDomainRelationListResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallDomainRelationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallDomainRelationListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainRelationListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainRelationListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallDomainRelationListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallDomainRelationListResponse) GetBody() *DescribeVpcFirewallDomainRelationListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallDomainRelationListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDomainRelationListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponse) SetStatusCode(v int32) *DescribeVpcFirewallDomainRelationListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponse) SetBody(v *DescribeVpcFirewallDomainRelationListResponseBody) *DescribeVpcFirewallDomainRelationListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
