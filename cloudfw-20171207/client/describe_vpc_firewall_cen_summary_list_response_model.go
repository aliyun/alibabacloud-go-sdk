// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenSummaryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallCenSummaryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallCenSummaryListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallCenSummaryListResponseBody) *DescribeVpcFirewallCenSummaryListResponse
	GetBody() *DescribeVpcFirewallCenSummaryListResponseBody
}

type DescribeVpcFirewallCenSummaryListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallCenSummaryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallCenSummaryListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenSummaryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenSummaryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallCenSummaryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallCenSummaryListResponse) GetBody() *DescribeVpcFirewallCenSummaryListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallCenSummaryListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallCenSummaryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponse) SetStatusCode(v int32) *DescribeVpcFirewallCenSummaryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponse) SetBody(v *DescribeVpcFirewallCenSummaryListResponseBody) *DescribeVpcFirewallCenSummaryListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
