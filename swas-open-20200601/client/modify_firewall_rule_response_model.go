// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFirewallRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFirewallRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFirewallRuleResponseBody) *ModifyFirewallRuleResponse
	GetBody() *ModifyFirewallRuleResponseBody
}

type ModifyFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFirewallRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFirewallRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFirewallRuleResponse) GetBody() *ModifyFirewallRuleResponseBody {
	return s.Body
}

func (s *ModifyFirewallRuleResponse) SetHeaders(v map[string]*string) *ModifyFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirewallRuleResponse) SetStatusCode(v int32) *ModifyFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirewallRuleResponse) SetBody(v *ModifyFirewallRuleResponseBody) *ModifyFirewallRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyFirewallRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
