// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConsumerAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConsumerAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConsumerAuthorizationRuleResponseBody) *UpdateConsumerAuthorizationRuleResponse
	GetBody() *UpdateConsumerAuthorizationRuleResponseBody
}

type UpdateConsumerAuthorizationRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConsumerAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConsumerAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConsumerAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConsumerAuthorizationRuleResponse) GetBody() *UpdateConsumerAuthorizationRuleResponseBody {
	return s.Body
}

func (s *UpdateConsumerAuthorizationRuleResponse) SetHeaders(v map[string]*string) *UpdateConsumerAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerAuthorizationRuleResponse) SetStatusCode(v int32) *UpdateConsumerAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConsumerAuthorizationRuleResponse) SetBody(v *UpdateConsumerAuthorizationRuleResponseBody) *UpdateConsumerAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateConsumerAuthorizationRuleResponse) Validate() error {
	return dara.Validate(s)
}
