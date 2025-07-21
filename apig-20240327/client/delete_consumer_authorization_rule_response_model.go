// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConsumerAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConsumerAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConsumerAuthorizationRuleResponseBody) *DeleteConsumerAuthorizationRuleResponse
	GetBody() *DeleteConsumerAuthorizationRuleResponseBody
}

type DeleteConsumerAuthorizationRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConsumerAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConsumerAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConsumerAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConsumerAuthorizationRuleResponse) GetBody() *DeleteConsumerAuthorizationRuleResponseBody {
	return s.Body
}

func (s *DeleteConsumerAuthorizationRuleResponse) SetHeaders(v map[string]*string) *DeleteConsumerAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerAuthorizationRuleResponse) SetStatusCode(v int32) *DeleteConsumerAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConsumerAuthorizationRuleResponse) SetBody(v *DeleteConsumerAuthorizationRuleResponseBody) *DeleteConsumerAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteConsumerAuthorizationRuleResponse) Validate() error {
	return dara.Validate(s)
}
