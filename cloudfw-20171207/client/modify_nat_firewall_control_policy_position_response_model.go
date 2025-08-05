// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatFirewallControlPolicyPositionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNatFirewallControlPolicyPositionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNatFirewallControlPolicyPositionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNatFirewallControlPolicyPositionResponseBody) *ModifyNatFirewallControlPolicyPositionResponse
	GetBody() *ModifyNatFirewallControlPolicyPositionResponseBody
}

type ModifyNatFirewallControlPolicyPositionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatFirewallControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatFirewallControlPolicyPositionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyPositionResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) GetBody() *ModifyNatFirewallControlPolicyPositionResponseBody {
	return s.Body
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) SetHeaders(v map[string]*string) *ModifyNatFirewallControlPolicyPositionResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) SetStatusCode(v int32) *ModifyNatFirewallControlPolicyPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) SetBody(v *ModifyNatFirewallControlPolicyPositionResponseBody) *ModifyNatFirewallControlPolicyPositionResponse {
	s.Body = v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) Validate() error {
	return dara.Validate(s)
}
