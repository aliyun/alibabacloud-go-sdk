// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFirewallRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFirewallRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFirewallRuleResponseBody) *DeleteFirewallRuleResponse
	GetBody() *DeleteFirewallRuleResponseBody
}

type DeleteFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFirewallRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFirewallRuleResponse) GetBody() *DeleteFirewallRuleResponseBody {
	return s.Body
}

func (s *DeleteFirewallRuleResponse) SetHeaders(v map[string]*string) *DeleteFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallRuleResponse) SetStatusCode(v int32) *DeleteFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallRuleResponse) SetBody(v *DeleteFirewallRuleResponseBody) *DeleteFirewallRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteFirewallRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
