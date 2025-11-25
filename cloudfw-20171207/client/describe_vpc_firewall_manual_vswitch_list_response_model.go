// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallManualVSwitchListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallManualVSwitchListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallManualVSwitchListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallManualVSwitchListResponseBody) *DescribeVpcFirewallManualVSwitchListResponse
	GetBody() *DescribeVpcFirewallManualVSwitchListResponseBody
}

type DescribeVpcFirewallManualVSwitchListResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallManualVSwitchListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallManualVSwitchListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallManualVSwitchListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallManualVSwitchListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallManualVSwitchListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallManualVSwitchListResponse) GetBody() *DescribeVpcFirewallManualVSwitchListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallManualVSwitchListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallManualVSwitchListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponse) SetStatusCode(v int32) *DescribeVpcFirewallManualVSwitchListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponse) SetBody(v *DescribeVpcFirewallManualVSwitchListResponseBody) *DescribeVpcFirewallManualVSwitchListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
