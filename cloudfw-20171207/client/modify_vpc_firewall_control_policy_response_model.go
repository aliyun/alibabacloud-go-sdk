// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallControlPolicyResponseBody) *ModifyVpcFirewallControlPolicyResponse
	GetBody() *ModifyVpcFirewallControlPolicyResponseBody
}

type ModifyVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallControlPolicyResponse) GetBody() *ModifyVpcFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *ModifyVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetBody(v *ModifyVpcFirewallControlPolicyResponseBody) *ModifyVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
