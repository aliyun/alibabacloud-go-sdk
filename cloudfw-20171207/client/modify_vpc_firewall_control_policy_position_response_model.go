// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallControlPolicyPositionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallControlPolicyPositionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallControlPolicyPositionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallControlPolicyPositionResponseBody) *ModifyVpcFirewallControlPolicyPositionResponse
	GetBody() *ModifyVpcFirewallControlPolicyPositionResponseBody
}

type ModifyVpcFirewallControlPolicyPositionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) GetBody() *ModifyVpcFirewallControlPolicyPositionResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetStatusCode(v int32) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetBody(v *ModifyVpcFirewallControlPolicyPositionResponseBody) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) Validate() error {
	return dara.Validate(s)
}
