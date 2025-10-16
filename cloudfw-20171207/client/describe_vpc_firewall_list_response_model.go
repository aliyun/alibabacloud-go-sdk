// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallListResponseBody) *DescribeVpcFirewallListResponse
	GetBody() *DescribeVpcFirewallListResponseBody
}

type DescribeVpcFirewallListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallListResponse) GetBody() *DescribeVpcFirewallListResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallListResponse) SetStatusCode(v int32) *DescribeVpcFirewallListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallListResponse) SetBody(v *DescribeVpcFirewallListResponseBody) *DescribeVpcFirewallListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
