// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizationRulesForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizationRulesForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizationRulesForUserResponseBody) *ListAuthorizationRulesForUserResponse
	GetBody() *ListAuthorizationRulesForUserResponseBody
}

type ListAuthorizationRulesForUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizationRulesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizationRulesForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizationRulesForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizationRulesForUserResponse) GetBody() *ListAuthorizationRulesForUserResponseBody {
	return s.Body
}

func (s *ListAuthorizationRulesForUserResponse) SetHeaders(v map[string]*string) *ListAuthorizationRulesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizationRulesForUserResponse) SetStatusCode(v int32) *ListAuthorizationRulesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponse) SetBody(v *ListAuthorizationRulesForUserResponseBody) *ListAuthorizationRulesForUserResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizationRulesForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
