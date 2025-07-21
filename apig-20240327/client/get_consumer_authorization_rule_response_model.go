// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerAuthorizationRuleResponseBody) *GetConsumerAuthorizationRuleResponse
	GetBody() *GetConsumerAuthorizationRuleResponseBody
}

type GetConsumerAuthorizationRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerAuthorizationRuleResponse) GetBody() *GetConsumerAuthorizationRuleResponseBody {
	return s.Body
}

func (s *GetConsumerAuthorizationRuleResponse) SetHeaders(v map[string]*string) *GetConsumerAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerAuthorizationRuleResponse) SetStatusCode(v int32) *GetConsumerAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponse) SetBody(v *GetConsumerAuthorizationRuleResponseBody) *GetConsumerAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *GetConsumerAuthorizationRuleResponse) Validate() error {
	return dara.Validate(s)
}
