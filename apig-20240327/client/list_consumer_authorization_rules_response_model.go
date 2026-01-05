// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerAuthorizationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerAuthorizationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerAuthorizationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerAuthorizationRulesResponseBody) *ListConsumerAuthorizationRulesResponse
	GetBody() *ListConsumerAuthorizationRulesResponseBody
}

type ListConsumerAuthorizationRulesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerAuthorizationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerAuthorizationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerAuthorizationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerAuthorizationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerAuthorizationRulesResponse) GetBody() *ListConsumerAuthorizationRulesResponseBody {
	return s.Body
}

func (s *ListConsumerAuthorizationRulesResponse) SetHeaders(v map[string]*string) *ListConsumerAuthorizationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerAuthorizationRulesResponse) SetStatusCode(v int32) *ListConsumerAuthorizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponse) SetBody(v *ListConsumerAuthorizationRulesResponseBody) *ListConsumerAuthorizationRulesResponse {
	s.Body = v
	return s
}

func (s *ListConsumerAuthorizationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
