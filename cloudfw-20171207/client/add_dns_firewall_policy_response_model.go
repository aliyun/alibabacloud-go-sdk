// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsFirewallPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDnsFirewallPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDnsFirewallPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AddDnsFirewallPolicyResponseBody) *AddDnsFirewallPolicyResponse
	GetBody() *AddDnsFirewallPolicyResponseBody
}

type AddDnsFirewallPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDnsFirewallPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDnsFirewallPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDnsFirewallPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddDnsFirewallPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDnsFirewallPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDnsFirewallPolicyResponse) GetBody() *AddDnsFirewallPolicyResponseBody {
	return s.Body
}

func (s *AddDnsFirewallPolicyResponse) SetHeaders(v map[string]*string) *AddDnsFirewallPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddDnsFirewallPolicyResponse) SetStatusCode(v int32) *AddDnsFirewallPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDnsFirewallPolicyResponse) SetBody(v *AddDnsFirewallPolicyResponseBody) *AddDnsFirewallPolicyResponse {
	s.Body = v
	return s
}

func (s *AddDnsFirewallPolicyResponse) Validate() error {
	return dara.Validate(s)
}
