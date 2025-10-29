// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConsumerAuthorizationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConsumerAuthorizationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConsumerAuthorizationRulesResponse
	GetStatusCode() *int32
	SetBody(v *QueryConsumerAuthorizationRulesResponseBody) *QueryConsumerAuthorizationRulesResponse
	GetBody() *QueryConsumerAuthorizationRulesResponseBody
}

type QueryConsumerAuthorizationRulesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConsumerAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConsumerAuthorizationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConsumerAuthorizationRulesResponse) GoString() string {
	return s.String()
}

func (s *QueryConsumerAuthorizationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConsumerAuthorizationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConsumerAuthorizationRulesResponse) GetBody() *QueryConsumerAuthorizationRulesResponseBody {
	return s.Body
}

func (s *QueryConsumerAuthorizationRulesResponse) SetHeaders(v map[string]*string) *QueryConsumerAuthorizationRulesResponse {
	s.Headers = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponse) SetStatusCode(v int32) *QueryConsumerAuthorizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponse) SetBody(v *QueryConsumerAuthorizationRulesResponseBody) *QueryConsumerAuthorizationRulesResponse {
	s.Body = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
