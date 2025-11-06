// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCircuitBreakerRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCircuitBreakerRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCircuitBreakerRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListCircuitBreakerRulesResponseBody) *ListCircuitBreakerRulesResponse
	GetBody() *ListCircuitBreakerRulesResponseBody
}

type ListCircuitBreakerRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCircuitBreakerRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCircuitBreakerRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCircuitBreakerRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCircuitBreakerRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCircuitBreakerRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCircuitBreakerRulesResponse) GetBody() *ListCircuitBreakerRulesResponseBody {
	return s.Body
}

func (s *ListCircuitBreakerRulesResponse) SetHeaders(v map[string]*string) *ListCircuitBreakerRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCircuitBreakerRulesResponse) SetStatusCode(v int32) *ListCircuitBreakerRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCircuitBreakerRulesResponse) SetBody(v *ListCircuitBreakerRulesResponseBody) *ListCircuitBreakerRulesResponse {
	s.Body = v
	return s
}

func (s *ListCircuitBreakerRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
