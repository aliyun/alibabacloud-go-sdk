// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNatFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNatFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateNatFirewallControlPolicyResponseBody) *CreateNatFirewallControlPolicyResponse
	GetBody() *CreateNatFirewallControlPolicyResponseBody
}

type CreateNatFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNatFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNatFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNatFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNatFirewallControlPolicyResponse) GetBody() *CreateNatFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *CreateNatFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *CreateNatFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateNatFirewallControlPolicyResponse) SetStatusCode(v int32) *CreateNatFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNatFirewallControlPolicyResponse) SetBody(v *CreateNatFirewallControlPolicyResponseBody) *CreateNatFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateNatFirewallControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}
