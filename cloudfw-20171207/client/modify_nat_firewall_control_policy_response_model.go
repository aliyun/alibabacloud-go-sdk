// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNatFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNatFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNatFirewallControlPolicyResponseBody) *ModifyNatFirewallControlPolicyResponse
	GetBody() *ModifyNatFirewallControlPolicyResponseBody
}

type ModifyNatFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNatFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNatFirewallControlPolicyResponse) GetBody() *ModifyNatFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *ModifyNatFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *ModifyNatFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatFirewallControlPolicyResponse) SetStatusCode(v int32) *ModifyNatFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyResponse) SetBody(v *ModifyNatFirewallControlPolicyResponseBody) *ModifyNatFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyNatFirewallControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
