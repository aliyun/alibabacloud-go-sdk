// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallIPSWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallIPSWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallIPSWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallIPSWhitelistResponseBody) *DescribeVpcFirewallIPSWhitelistResponse
	GetBody() *DescribeVpcFirewallIPSWhitelistResponseBody
}

type DescribeVpcFirewallIPSWhitelistResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallIPSWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallIPSWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallIPSWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) GetBody() *DescribeVpcFirewallIPSWhitelistResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallIPSWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) SetStatusCode(v int32) *DescribeVpcFirewallIPSWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) SetBody(v *DescribeVpcFirewallIPSWhitelistResponseBody) *DescribeVpcFirewallIPSWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
