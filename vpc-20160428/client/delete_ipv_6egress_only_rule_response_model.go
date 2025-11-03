// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6EgressOnlyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpv6EgressOnlyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpv6EgressOnlyRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpv6EgressOnlyRuleResponseBody) *DeleteIpv6EgressOnlyRuleResponse
	GetBody() *DeleteIpv6EgressOnlyRuleResponseBody
}

type DeleteIpv6EgressOnlyRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpv6EgressOnlyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpv6EgressOnlyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6EgressOnlyRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpv6EgressOnlyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpv6EgressOnlyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpv6EgressOnlyRuleResponse) GetBody() *DeleteIpv6EgressOnlyRuleResponseBody {
	return s.Body
}

func (s *DeleteIpv6EgressOnlyRuleResponse) SetHeaders(v map[string]*string) *DeleteIpv6EgressOnlyRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleResponse) SetStatusCode(v int32) *DeleteIpv6EgressOnlyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleResponse) SetBody(v *DeleteIpv6EgressOnlyRuleResponseBody) *DeleteIpv6EgressOnlyRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
