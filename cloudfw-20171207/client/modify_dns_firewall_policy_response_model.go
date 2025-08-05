// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDnsFirewallPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDnsFirewallPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDnsFirewallPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDnsFirewallPolicyResponseBody) *ModifyDnsFirewallPolicyResponse
	GetBody() *ModifyDnsFirewallPolicyResponseBody
}

type ModifyDnsFirewallPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDnsFirewallPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDnsFirewallPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDnsFirewallPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDnsFirewallPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDnsFirewallPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDnsFirewallPolicyResponse) GetBody() *ModifyDnsFirewallPolicyResponseBody {
	return s.Body
}

func (s *ModifyDnsFirewallPolicyResponse) SetHeaders(v map[string]*string) *ModifyDnsFirewallPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDnsFirewallPolicyResponse) SetStatusCode(v int32) *ModifyDnsFirewallPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDnsFirewallPolicyResponse) SetBody(v *ModifyDnsFirewallPolicyResponseBody) *ModifyDnsFirewallPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyDnsFirewallPolicyResponse) Validate() error {
	return dara.Validate(s)
}
