// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserToAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserToAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddUserToAuthorizationRuleResponseBody) *AddUserToAuthorizationRuleResponse
	GetBody() *AddUserToAuthorizationRuleResponseBody
}

type AddUserToAuthorizationRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserToAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *AddUserToAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserToAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserToAuthorizationRuleResponse) GetBody() *AddUserToAuthorizationRuleResponseBody {
	return s.Body
}

func (s *AddUserToAuthorizationRuleResponse) SetHeaders(v map[string]*string) *AddUserToAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *AddUserToAuthorizationRuleResponse) SetStatusCode(v int32) *AddUserToAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToAuthorizationRuleResponse) SetBody(v *AddUserToAuthorizationRuleResponseBody) *AddUserToAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *AddUserToAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
