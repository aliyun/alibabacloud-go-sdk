// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallTemplateRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFirewallTemplateRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFirewallTemplateRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFirewallTemplateRulesResponseBody) *DeleteFirewallTemplateRulesResponse
	GetBody() *DeleteFirewallTemplateRulesResponseBody
}

type DeleteFirewallTemplateRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallTemplateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallTemplateRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallTemplateRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplateRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFirewallTemplateRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFirewallTemplateRulesResponse) GetBody() *DeleteFirewallTemplateRulesResponseBody {
	return s.Body
}

func (s *DeleteFirewallTemplateRulesResponse) SetHeaders(v map[string]*string) *DeleteFirewallTemplateRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallTemplateRulesResponse) SetStatusCode(v int32) *DeleteFirewallTemplateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallTemplateRulesResponse) SetBody(v *DeleteFirewallTemplateRulesResponseBody) *DeleteFirewallTemplateRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteFirewallTemplateRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
