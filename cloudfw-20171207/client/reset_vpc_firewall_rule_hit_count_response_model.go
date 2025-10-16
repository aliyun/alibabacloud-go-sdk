// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetVpcFirewallRuleHitCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetVpcFirewallRuleHitCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetVpcFirewallRuleHitCountResponse
	GetStatusCode() *int32
	SetBody(v *ResetVpcFirewallRuleHitCountResponseBody) *ResetVpcFirewallRuleHitCountResponse
	GetBody() *ResetVpcFirewallRuleHitCountResponseBody
}

type ResetVpcFirewallRuleHitCountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetVpcFirewallRuleHitCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountResponse) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetVpcFirewallRuleHitCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetVpcFirewallRuleHitCountResponse) GetBody() *ResetVpcFirewallRuleHitCountResponseBody {
	return s.Body
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetHeaders(v map[string]*string) *ResetVpcFirewallRuleHitCountResponse {
	s.Headers = v
	return s
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetStatusCode(v int32) *ResetVpcFirewallRuleHitCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetBody(v *ResetVpcFirewallRuleHitCountResponseBody) *ResetVpcFirewallRuleHitCountResponse {
	s.Body = v
	return s
}

func (s *ResetVpcFirewallRuleHitCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
