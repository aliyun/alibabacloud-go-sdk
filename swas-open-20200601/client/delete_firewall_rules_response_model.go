// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFirewallRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFirewallRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFirewallRulesResponseBody) *DeleteFirewallRulesResponse
	GetBody() *DeleteFirewallRulesResponseBody
}

type DeleteFirewallRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFirewallRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFirewallRulesResponse) GetBody() *DeleteFirewallRulesResponseBody {
	return s.Body
}

func (s *DeleteFirewallRulesResponse) SetHeaders(v map[string]*string) *DeleteFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallRulesResponse) SetStatusCode(v int32) *DeleteFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallRulesResponse) SetBody(v *DeleteFirewallRulesResponseBody) *DeleteFirewallRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteFirewallRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
