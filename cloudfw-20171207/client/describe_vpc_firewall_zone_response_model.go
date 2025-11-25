// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallZoneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallZoneResponseBody) *DescribeVpcFirewallZoneResponse
	GetBody() *DescribeVpcFirewallZoneResponseBody
}

type DescribeVpcFirewallZoneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallZoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallZoneResponse) GetBody() *DescribeVpcFirewallZoneResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallZoneResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallZoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallZoneResponse) SetStatusCode(v int32) *DescribeVpcFirewallZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallZoneResponse) SetBody(v *DescribeVpcFirewallZoneResponseBody) *DescribeVpcFirewallZoneResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
