// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFirewallRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFirewallRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFirewallRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListFirewallRulesResponseBody) *ListFirewallRulesResponse
	GetBody() *ListFirewallRulesResponseBody
}

type ListFirewallRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFirewallRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFirewallRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFirewallRulesResponse) GetBody() *ListFirewallRulesResponseBody {
	return s.Body
}

func (s *ListFirewallRulesResponse) SetHeaders(v map[string]*string) *ListFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *ListFirewallRulesResponse) SetStatusCode(v int32) *ListFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFirewallRulesResponse) SetBody(v *ListFirewallRulesResponseBody) *ListFirewallRulesResponse {
	s.Body = v
	return s
}

func (s *ListFirewallRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
