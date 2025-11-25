// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallPrecheckDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallPrecheckDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallPrecheckDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallPrecheckDetailResponseBody) *DescribeNatFirewallPrecheckDetailResponse
	GetBody() *DescribeNatFirewallPrecheckDetailResponseBody
}

type DescribeNatFirewallPrecheckDetailResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallPrecheckDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallPrecheckDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPrecheckDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPrecheckDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallPrecheckDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallPrecheckDetailResponse) GetBody() *DescribeNatFirewallPrecheckDetailResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallPrecheckDetailResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallPrecheckDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponse) SetStatusCode(v int32) *DescribeNatFirewallPrecheckDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponse) SetBody(v *DescribeNatFirewallPrecheckDetailResponseBody) *DescribeNatFirewallPrecheckDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
