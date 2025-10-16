// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallCenListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallCenListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallCenListResponseBody) *DescribeVpcFirewallCenListResponse
	GetBody() *DescribeVpcFirewallCenListResponseBody
}

type DescribeVpcFirewallCenListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallCenListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallCenListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallCenListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallCenListResponse) GetBody() *DescribeVpcFirewallCenListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallCenListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallCenListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallCenListResponse) SetStatusCode(v int32) *DescribeVpcFirewallCenListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponse) SetBody(v *DescribeVpcFirewallCenListResponseBody) *DescribeVpcFirewallCenListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallCenListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
