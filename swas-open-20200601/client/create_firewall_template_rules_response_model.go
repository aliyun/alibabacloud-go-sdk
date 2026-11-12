// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallTemplateRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFirewallTemplateRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFirewallTemplateRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateFirewallTemplateRulesResponseBody) *CreateFirewallTemplateRulesResponse
	GetBody() *CreateFirewallTemplateRulesResponseBody
}

type CreateFirewallTemplateRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirewallTemplateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirewallTemplateRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFirewallTemplateRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFirewallTemplateRulesResponse) GetBody() *CreateFirewallTemplateRulesResponseBody {
	return s.Body
}

func (s *CreateFirewallTemplateRulesResponse) SetHeaders(v map[string]*string) *CreateFirewallTemplateRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallTemplateRulesResponse) SetStatusCode(v int32) *CreateFirewallTemplateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponse) SetBody(v *CreateFirewallTemplateRulesResponseBody) *CreateFirewallTemplateRulesResponse {
	s.Body = v
	return s
}

func (s *CreateFirewallTemplateRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
