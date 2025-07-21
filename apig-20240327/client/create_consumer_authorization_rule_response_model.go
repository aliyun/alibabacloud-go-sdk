// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConsumerAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConsumerAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateConsumerAuthorizationRuleResponseBody) *CreateConsumerAuthorizationRuleResponse
	GetBody() *CreateConsumerAuthorizationRuleResponseBody
}

type CreateConsumerAuthorizationRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConsumerAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConsumerAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConsumerAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConsumerAuthorizationRuleResponse) GetBody() *CreateConsumerAuthorizationRuleResponseBody {
	return s.Body
}

func (s *CreateConsumerAuthorizationRuleResponse) SetHeaders(v map[string]*string) *CreateConsumerAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerAuthorizationRuleResponse) SetStatusCode(v int32) *CreateConsumerAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleResponse) SetBody(v *CreateConsumerAuthorizationRuleResponseBody) *CreateConsumerAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateConsumerAuthorizationRuleResponse) Validate() error {
	return dara.Validate(s)
}
