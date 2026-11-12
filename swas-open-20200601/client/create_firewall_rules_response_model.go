// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFirewallRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFirewallRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateFirewallRulesResponseBody) *CreateFirewallRulesResponse
	GetBody() *CreateFirewallRulesResponseBody
}

type CreateFirewallRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirewallRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFirewallRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFirewallRulesResponse) GetBody() *CreateFirewallRulesResponseBody {
	return s.Body
}

func (s *CreateFirewallRulesResponse) SetHeaders(v map[string]*string) *CreateFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallRulesResponse) SetStatusCode(v int32) *CreateFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallRulesResponse) SetBody(v *CreateFirewallRulesResponseBody) *CreateFirewallRulesResponse {
	s.Body = v
	return s
}

func (s *CreateFirewallRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
