// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNatFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNatFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNatFirewallControlPolicyResponseBody) *DeleteNatFirewallControlPolicyResponse
	GetBody() *DeleteNatFirewallControlPolicyResponseBody
}

type DeleteNatFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNatFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNatFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNatFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNatFirewallControlPolicyResponse) GetBody() *DeleteNatFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *DeleteNatFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteNatFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteNatFirewallControlPolicyResponse) SetStatusCode(v int32) *DeleteNatFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyResponse) SetBody(v *DeleteNatFirewallControlPolicyResponseBody) *DeleteNatFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteNatFirewallControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
