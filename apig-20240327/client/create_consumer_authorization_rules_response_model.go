// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAuthorizationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConsumerAuthorizationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConsumerAuthorizationRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateConsumerAuthorizationRulesResponseBody) *CreateConsumerAuthorizationRulesResponse
	GetBody() *CreateConsumerAuthorizationRulesResponseBody
}

type CreateConsumerAuthorizationRulesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConsumerAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConsumerAuthorizationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConsumerAuthorizationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConsumerAuthorizationRulesResponse) GetBody() *CreateConsumerAuthorizationRulesResponseBody {
	return s.Body
}

func (s *CreateConsumerAuthorizationRulesResponse) SetHeaders(v map[string]*string) *CreateConsumerAuthorizationRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerAuthorizationRulesResponse) SetStatusCode(v int32) *CreateConsumerAuthorizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesResponse) SetBody(v *CreateConsumerAuthorizationRulesResponseBody) *CreateConsumerAuthorizationRulesResponse {
	s.Body = v
	return s
}

func (s *CreateConsumerAuthorizationRulesResponse) Validate() error {
	return dara.Validate(s)
}
