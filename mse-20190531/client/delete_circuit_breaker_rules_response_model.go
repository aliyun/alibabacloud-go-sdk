// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCircuitBreakerRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCircuitBreakerRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCircuitBreakerRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCircuitBreakerRulesResponseBody) *DeleteCircuitBreakerRulesResponse
	GetBody() *DeleteCircuitBreakerRulesResponseBody
}

type DeleteCircuitBreakerRulesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCircuitBreakerRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCircuitBreakerRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCircuitBreakerRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCircuitBreakerRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCircuitBreakerRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCircuitBreakerRulesResponse) GetBody() *DeleteCircuitBreakerRulesResponseBody {
	return s.Body
}

func (s *DeleteCircuitBreakerRulesResponse) SetHeaders(v map[string]*string) *DeleteCircuitBreakerRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCircuitBreakerRulesResponse) SetStatusCode(v int32) *DeleteCircuitBreakerRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCircuitBreakerRulesResponse) SetBody(v *DeleteCircuitBreakerRulesResponseBody) *DeleteCircuitBreakerRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteCircuitBreakerRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
