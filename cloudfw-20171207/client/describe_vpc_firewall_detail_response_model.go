// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallDetailResponseBody) *DescribeVpcFirewallDetailResponse
	GetBody() *DescribeVpcFirewallDetailResponseBody
}

type DescribeVpcFirewallDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallDetailResponse) GetBody() *DescribeVpcFirewallDetailResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallDetailResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDetailResponse) SetStatusCode(v int32) *DescribeVpcFirewallDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponse) SetBody(v *DescribeVpcFirewallDetailResponseBody) *DescribeVpcFirewallDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
