// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCircuitBreakerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCircuitBreakerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCircuitBreakerRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCircuitBreakerRuleResponseBody) *CreateCircuitBreakerRuleResponse
	GetBody() *CreateCircuitBreakerRuleResponseBody
}

type CreateCircuitBreakerRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCircuitBreakerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCircuitBreakerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCircuitBreakerRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateCircuitBreakerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCircuitBreakerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCircuitBreakerRuleResponse) GetBody() *CreateCircuitBreakerRuleResponseBody {
	return s.Body
}

func (s *CreateCircuitBreakerRuleResponse) SetHeaders(v map[string]*string) *CreateCircuitBreakerRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateCircuitBreakerRuleResponse) SetStatusCode(v int32) *CreateCircuitBreakerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCircuitBreakerRuleResponse) SetBody(v *CreateCircuitBreakerRuleResponseBody) *CreateCircuitBreakerRuleResponse {
	s.Body = v
	return s
}

func (s *CreateCircuitBreakerRuleResponse) Validate() error {
	return dara.Validate(s)
}
