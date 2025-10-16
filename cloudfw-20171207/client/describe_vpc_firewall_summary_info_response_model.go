// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallSummaryInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallSummaryInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallSummaryInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallSummaryInfoResponseBody) *DescribeVpcFirewallSummaryInfoResponse
	GetBody() *DescribeVpcFirewallSummaryInfoResponseBody
}

type DescribeVpcFirewallSummaryInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallSummaryInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallSummaryInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallSummaryInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallSummaryInfoResponse) GetBody() *DescribeVpcFirewallSummaryInfoResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallSummaryInfoResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponse) SetStatusCode(v int32) *DescribeVpcFirewallSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponse) SetBody(v *DescribeVpcFirewallSummaryInfoResponseBody) *DescribeVpcFirewallSummaryInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
