// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDefaultIPSConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallDefaultIPSConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallDefaultIPSConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallDefaultIPSConfigResponseBody) *DescribeVpcFirewallDefaultIPSConfigResponse
	GetBody() *DescribeVpcFirewallDefaultIPSConfigResponseBody
}

type DescribeVpcFirewallDefaultIPSConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallDefaultIPSConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) GetBody() *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetStatusCode(v int32) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetBody(v *DescribeVpcFirewallDefaultIPSConfigResponseBody) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) Validate() error {
	return dara.Validate(s)
}
