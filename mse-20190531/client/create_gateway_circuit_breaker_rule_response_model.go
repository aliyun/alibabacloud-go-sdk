// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayCircuitBreakerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayCircuitBreakerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayCircuitBreakerRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayCircuitBreakerRuleResponseBody) *CreateGatewayCircuitBreakerRuleResponse
	GetBody() *CreateGatewayCircuitBreakerRuleResponseBody
}

type CreateGatewayCircuitBreakerRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayCircuitBreakerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayCircuitBreakerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayCircuitBreakerRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayCircuitBreakerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayCircuitBreakerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayCircuitBreakerRuleResponse) GetBody() *CreateGatewayCircuitBreakerRuleResponseBody {
	return s.Body
}

func (s *CreateGatewayCircuitBreakerRuleResponse) SetHeaders(v map[string]*string) *CreateGatewayCircuitBreakerRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleResponse) SetStatusCode(v int32) *CreateGatewayCircuitBreakerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleResponse) SetBody(v *CreateGatewayCircuitBreakerRuleResponseBody) *CreateGatewayCircuitBreakerRuleResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
