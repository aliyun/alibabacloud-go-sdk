// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallCenDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallCenDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallCenDetailResponseBody) *DescribeVpcFirewallCenDetailResponse
	GetBody() *DescribeVpcFirewallCenDetailResponseBody
}

type DescribeVpcFirewallCenDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallCenDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallCenDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallCenDetailResponse) GetBody() *DescribeVpcFirewallCenDetailResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallCenDetailResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallCenDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponse) SetStatusCode(v int32) *DescribeVpcFirewallCenDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponse) SetBody(v *DescribeVpcFirewallCenDetailResponseBody) *DescribeVpcFirewallCenDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponse) Validate() error {
	return dara.Validate(s)
}
