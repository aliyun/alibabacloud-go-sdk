// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallListResponseBody) *DescribeNatFirewallListResponse
	GetBody() *DescribeNatFirewallListResponseBody
}

type DescribeNatFirewallListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallListResponse) GetBody() *DescribeNatFirewallListResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallListResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallListResponse) SetStatusCode(v int32) *DescribeNatFirewallListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallListResponse) SetBody(v *DescribeNatFirewallListResponseBody) *DescribeNatFirewallListResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallListResponse) Validate() error {
	return dara.Validate(s)
}
