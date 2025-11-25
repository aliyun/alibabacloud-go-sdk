// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAssetRegionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallAssetRegionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallAssetRegionListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallAssetRegionListResponseBody) *DescribeVpcFirewallAssetRegionListResponse
	GetBody() *DescribeVpcFirewallAssetRegionListResponseBody
}

type DescribeVpcFirewallAssetRegionListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallAssetRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallAssetRegionListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAssetRegionListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAssetRegionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallAssetRegionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallAssetRegionListResponse) GetBody() *DescribeVpcFirewallAssetRegionListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallAssetRegionListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallAssetRegionListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallAssetRegionListResponse) SetStatusCode(v int32) *DescribeVpcFirewallAssetRegionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallAssetRegionListResponse) SetBody(v *DescribeVpcFirewallAssetRegionListResponseBody) *DescribeVpcFirewallAssetRegionListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallAssetRegionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
