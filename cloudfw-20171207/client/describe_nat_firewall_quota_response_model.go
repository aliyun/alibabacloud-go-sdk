// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallQuotaResponseBody) *DescribeNatFirewallQuotaResponse
	GetBody() *DescribeNatFirewallQuotaResponseBody
}

type DescribeNatFirewallQuotaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallQuotaResponse) GetBody() *DescribeNatFirewallQuotaResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallQuotaResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallQuotaResponse) SetStatusCode(v int32) *DescribeNatFirewallQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallQuotaResponse) SetBody(v *DescribeNatFirewallQuotaResponseBody) *DescribeNatFirewallQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
