// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFirewallRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFirewallRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateFirewallRuleResponseBody) *CreateFirewallRuleResponse
	GetBody() *CreateFirewallRuleResponseBody
}

type CreateFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirewallRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFirewallRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFirewallRuleResponse) GetBody() *CreateFirewallRuleResponseBody {
	return s.Body
}

func (s *CreateFirewallRuleResponse) SetHeaders(v map[string]*string) *CreateFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallRuleResponse) SetStatusCode(v int32) *CreateFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallRuleResponse) SetBody(v *CreateFirewallRuleResponseBody) *CreateFirewallRuleResponse {
	s.Body = v
	return s
}

func (s *CreateFirewallRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
