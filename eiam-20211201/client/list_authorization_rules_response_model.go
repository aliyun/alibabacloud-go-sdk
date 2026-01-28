// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizationRulesResponseBody) *ListAuthorizationRulesResponse
	GetBody() *ListAuthorizationRulesResponseBody
}

type ListAuthorizationRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizationRulesResponse) GetBody() *ListAuthorizationRulesResponseBody {
	return s.Body
}

func (s *ListAuthorizationRulesResponse) SetHeaders(v map[string]*string) *ListAuthorizationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizationRulesResponse) SetStatusCode(v int32) *ListAuthorizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizationRulesResponse) SetBody(v *ListAuthorizationRulesResponseBody) *ListAuthorizationRulesResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
