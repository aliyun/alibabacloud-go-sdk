// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv6EgressOnlyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpv6EgressOnlyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpv6EgressOnlyRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpv6EgressOnlyRuleResponseBody) *CreateIpv6EgressOnlyRuleResponse
	GetBody() *CreateIpv6EgressOnlyRuleResponseBody
}

type CreateIpv6EgressOnlyRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpv6EgressOnlyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpv6EgressOnlyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv6EgressOnlyRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateIpv6EgressOnlyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpv6EgressOnlyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpv6EgressOnlyRuleResponse) GetBody() *CreateIpv6EgressOnlyRuleResponseBody {
	return s.Body
}

func (s *CreateIpv6EgressOnlyRuleResponse) SetHeaders(v map[string]*string) *CreateIpv6EgressOnlyRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateIpv6EgressOnlyRuleResponse) SetStatusCode(v int32) *CreateIpv6EgressOnlyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleResponse) SetBody(v *CreateIpv6EgressOnlyRuleResponseBody) *CreateIpv6EgressOnlyRuleResponse {
	s.Body = v
	return s
}

func (s *CreateIpv6EgressOnlyRuleResponse) Validate() error {
	return dara.Validate(s)
}
