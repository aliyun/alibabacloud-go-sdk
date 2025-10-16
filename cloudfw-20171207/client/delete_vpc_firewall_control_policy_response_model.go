// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcFirewallControlPolicyResponseBody) *DeleteVpcFirewallControlPolicyResponse
	GetBody() *DeleteVpcFirewallControlPolicyResponseBody
}

type DeleteVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcFirewallControlPolicyResponse) GetBody() *DeleteVpcFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *DeleteVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetBody(v *DeleteVpcFirewallControlPolicyResponseBody) *DeleteVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcFirewallControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
