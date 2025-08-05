// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsFirewallPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsFirewallPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsFirewallPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsFirewallPolicyResponseBody) *DescribeDnsFirewallPolicyResponse
	GetBody() *DescribeDnsFirewallPolicyResponseBody
}

type DescribeDnsFirewallPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsFirewallPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsFirewallPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsFirewallPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsFirewallPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsFirewallPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsFirewallPolicyResponse) GetBody() *DescribeDnsFirewallPolicyResponseBody {
	return s.Body
}

func (s *DescribeDnsFirewallPolicyResponse) SetHeaders(v map[string]*string) *DescribeDnsFirewallPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsFirewallPolicyResponse) SetStatusCode(v int32) *DescribeDnsFirewallPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponse) SetBody(v *DescribeDnsFirewallPolicyResponseBody) *DescribeDnsFirewallPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsFirewallPolicyResponse) Validate() error {
	return dara.Validate(s)
}
