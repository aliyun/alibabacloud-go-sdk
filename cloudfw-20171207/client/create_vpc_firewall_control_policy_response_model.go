// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcFirewallControlPolicyResponseBody) *CreateVpcFirewallControlPolicyResponse
	GetBody() *CreateVpcFirewallControlPolicyResponseBody
}

type CreateVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcFirewallControlPolicyResponse) GetBody() *CreateVpcFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *CreateVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *CreateVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponse) SetBody(v *CreateVpcFirewallControlPolicyResponseBody) *CreateVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
