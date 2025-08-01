// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayCircuitBreakerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayCircuitBreakerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayCircuitBreakerRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayCircuitBreakerRuleResponseBody) *ListGatewayCircuitBreakerRuleResponse
	GetBody() *ListGatewayCircuitBreakerRuleResponseBody
}

type ListGatewayCircuitBreakerRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayCircuitBreakerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayCircuitBreakerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayCircuitBreakerRuleResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayCircuitBreakerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayCircuitBreakerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayCircuitBreakerRuleResponse) GetBody() *ListGatewayCircuitBreakerRuleResponseBody {
	return s.Body
}

func (s *ListGatewayCircuitBreakerRuleResponse) SetHeaders(v map[string]*string) *ListGatewayCircuitBreakerRuleResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponse) SetStatusCode(v int32) *ListGatewayCircuitBreakerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponse) SetBody(v *ListGatewayCircuitBreakerRuleResponseBody) *ListGatewayCircuitBreakerRuleResponse {
	s.Body = v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponse) Validate() error {
	return dara.Validate(s)
}
