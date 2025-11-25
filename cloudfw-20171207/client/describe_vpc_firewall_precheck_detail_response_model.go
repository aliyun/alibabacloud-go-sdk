// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallPrecheckDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallPrecheckDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallPrecheckDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallPrecheckDetailResponseBody) *DescribeVpcFirewallPrecheckDetailResponse
	GetBody() *DescribeVpcFirewallPrecheckDetailResponseBody
}

type DescribeVpcFirewallPrecheckDetailResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallPrecheckDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallPrecheckDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPrecheckDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPrecheckDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallPrecheckDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallPrecheckDetailResponse) GetBody() *DescribeVpcFirewallPrecheckDetailResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallPrecheckDetailResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallPrecheckDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponse) SetStatusCode(v int32) *DescribeVpcFirewallPrecheckDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponse) SetBody(v *DescribeVpcFirewallPrecheckDetailResponseBody) *DescribeVpcFirewallPrecheckDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
