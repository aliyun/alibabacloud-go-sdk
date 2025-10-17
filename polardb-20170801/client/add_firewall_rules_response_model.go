// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFirewallRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFirewallRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFirewallRulesResponse
	GetStatusCode() *int32
	SetBody(v *AddFirewallRulesResponseBody) *AddFirewallRulesResponse
	GetBody() *AddFirewallRulesResponseBody
}

type AddFirewallRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFirewallRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *AddFirewallRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFirewallRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFirewallRulesResponse) GetBody() *AddFirewallRulesResponseBody {
	return s.Body
}

func (s *AddFirewallRulesResponse) SetHeaders(v map[string]*string) *AddFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *AddFirewallRulesResponse) SetStatusCode(v int32) *AddFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFirewallRulesResponse) SetBody(v *AddFirewallRulesResponseBody) *AddFirewallRulesResponse {
	s.Body = v
	return s
}

func (s *AddFirewallRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
