// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayCircuitBreakerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayCircuitBreakerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayCircuitBreakerRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayCircuitBreakerRuleResponseBody) *UpdateGatewayCircuitBreakerRuleResponse
	GetBody() *UpdateGatewayCircuitBreakerRuleResponseBody
}

type UpdateGatewayCircuitBreakerRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayCircuitBreakerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayCircuitBreakerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayCircuitBreakerRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayCircuitBreakerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayCircuitBreakerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayCircuitBreakerRuleResponse) GetBody() *UpdateGatewayCircuitBreakerRuleResponseBody {
	return s.Body
}

func (s *UpdateGatewayCircuitBreakerRuleResponse) SetHeaders(v map[string]*string) *UpdateGatewayCircuitBreakerRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponse) SetStatusCode(v int32) *UpdateGatewayCircuitBreakerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponse) SetBody(v *UpdateGatewayCircuitBreakerRuleResponseBody) *UpdateGatewayCircuitBreakerRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
