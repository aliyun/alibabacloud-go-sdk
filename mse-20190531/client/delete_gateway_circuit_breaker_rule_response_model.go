// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayCircuitBreakerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayCircuitBreakerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayCircuitBreakerRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayCircuitBreakerRuleResponseBody) *DeleteGatewayCircuitBreakerRuleResponse
	GetBody() *DeleteGatewayCircuitBreakerRuleResponseBody
}

type DeleteGatewayCircuitBreakerRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayCircuitBreakerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayCircuitBreakerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayCircuitBreakerRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayCircuitBreakerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayCircuitBreakerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayCircuitBreakerRuleResponse) GetBody() *DeleteGatewayCircuitBreakerRuleResponseBody {
	return s.Body
}

func (s *DeleteGatewayCircuitBreakerRuleResponse) SetHeaders(v map[string]*string) *DeleteGatewayCircuitBreakerRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleResponse) SetStatusCode(v int32) *DeleteGatewayCircuitBreakerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleResponse) SetBody(v *DeleteGatewayCircuitBreakerRuleResponseBody) *DeleteGatewayCircuitBreakerRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
