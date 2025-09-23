// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFirewallRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFirewallRulesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFirewallRulesResponseBody) *ModifyFirewallRulesResponse
	GetBody() *ModifyFirewallRulesResponseBody
}

type ModifyFirewallRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFirewallRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFirewallRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFirewallRulesResponse) GetBody() *ModifyFirewallRulesResponseBody {
	return s.Body
}

func (s *ModifyFirewallRulesResponse) SetHeaders(v map[string]*string) *ModifyFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirewallRulesResponse) SetStatusCode(v int32) *ModifyFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirewallRulesResponse) SetBody(v *ModifyFirewallRulesResponseBody) *ModifyFirewallRulesResponse {
	s.Body = v
	return s
}

func (s *ModifyFirewallRulesResponse) Validate() error {
	return dara.Validate(s)
}
