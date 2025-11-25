// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAssetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallAssetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallAssetListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallAssetListResponseBody) *DescribeVpcFirewallAssetListResponse
	GetBody() *DescribeVpcFirewallAssetListResponseBody
}

type DescribeVpcFirewallAssetListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallAssetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallAssetListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAssetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAssetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallAssetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallAssetListResponse) GetBody() *DescribeVpcFirewallAssetListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallAssetListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallAssetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallAssetListResponse) SetStatusCode(v int32) *DescribeVpcFirewallAssetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponse) SetBody(v *DescribeVpcFirewallAssetListResponseBody) *DescribeVpcFirewallAssetListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallAssetListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
