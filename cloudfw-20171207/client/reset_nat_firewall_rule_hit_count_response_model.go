// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetNatFirewallRuleHitCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetNatFirewallRuleHitCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetNatFirewallRuleHitCountResponse
	GetStatusCode() *int32
	SetBody(v *ResetNatFirewallRuleHitCountResponseBody) *ResetNatFirewallRuleHitCountResponse
	GetBody() *ResetNatFirewallRuleHitCountResponseBody
}

type ResetNatFirewallRuleHitCountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetNatFirewallRuleHitCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetNatFirewallRuleHitCountResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetNatFirewallRuleHitCountResponse) GoString() string {
	return s.String()
}

func (s *ResetNatFirewallRuleHitCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetNatFirewallRuleHitCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetNatFirewallRuleHitCountResponse) GetBody() *ResetNatFirewallRuleHitCountResponseBody {
	return s.Body
}

func (s *ResetNatFirewallRuleHitCountResponse) SetHeaders(v map[string]*string) *ResetNatFirewallRuleHitCountResponse {
	s.Headers = v
	return s
}

func (s *ResetNatFirewallRuleHitCountResponse) SetStatusCode(v int32) *ResetNatFirewallRuleHitCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetNatFirewallRuleHitCountResponse) SetBody(v *ResetNatFirewallRuleHitCountResponseBody) *ResetNatFirewallRuleHitCountResponse {
	s.Body = v
	return s
}

func (s *ResetNatFirewallRuleHitCountResponse) Validate() error {
	return dara.Validate(s)
}
