// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCircuitBreakerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCircuitBreakerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCircuitBreakerRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCircuitBreakerRuleResponseBody) *UpdateCircuitBreakerRuleResponse
	GetBody() *UpdateCircuitBreakerRuleResponseBody
}

type UpdateCircuitBreakerRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCircuitBreakerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCircuitBreakerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCircuitBreakerRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateCircuitBreakerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCircuitBreakerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCircuitBreakerRuleResponse) GetBody() *UpdateCircuitBreakerRuleResponseBody {
	return s.Body
}

func (s *UpdateCircuitBreakerRuleResponse) SetHeaders(v map[string]*string) *UpdateCircuitBreakerRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateCircuitBreakerRuleResponse) SetStatusCode(v int32) *UpdateCircuitBreakerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCircuitBreakerRuleResponse) SetBody(v *UpdateCircuitBreakerRuleResponseBody) *UpdateCircuitBreakerRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateCircuitBreakerRuleResponse) Validate() error {
	return dara.Validate(s)
}
