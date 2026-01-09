// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallTrafficAssetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallTrafficAssetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallTrafficAssetListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallTrafficAssetListResponseBody) *DescribeVpcFirewallTrafficAssetListResponse
	GetBody() *DescribeVpcFirewallTrafficAssetListResponseBody
}

type DescribeVpcFirewallTrafficAssetListResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallTrafficAssetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallTrafficAssetListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallTrafficAssetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallTrafficAssetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallTrafficAssetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallTrafficAssetListResponse) GetBody() *DescribeVpcFirewallTrafficAssetListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallTrafficAssetListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallTrafficAssetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponse) SetStatusCode(v int32) *DescribeVpcFirewallTrafficAssetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponse) SetBody(v *DescribeVpcFirewallTrafficAssetListResponseBody) *DescribeVpcFirewallTrafficAssetListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
