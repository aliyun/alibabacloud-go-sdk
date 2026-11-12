// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFirewallRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableFirewallRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableFirewallRuleResponse
	GetStatusCode() *int32
	SetBody(v *DisableFirewallRuleResponseBody) *DisableFirewallRuleResponse
	GetBody() *DisableFirewallRuleResponseBody
}

type DisableFirewallRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableFirewallRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableFirewallRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableFirewallRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableFirewallRuleResponse) GetBody() *DisableFirewallRuleResponseBody {
	return s.Body
}

func (s *DisableFirewallRuleResponse) SetHeaders(v map[string]*string) *DisableFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableFirewallRuleResponse) SetStatusCode(v int32) *DisableFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableFirewallRuleResponse) SetBody(v *DisableFirewallRuleResponseBody) *DisableFirewallRuleResponse {
	s.Body = v
	return s
}

func (s *DisableFirewallRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
