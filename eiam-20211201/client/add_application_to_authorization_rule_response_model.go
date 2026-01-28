// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApplicationToAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddApplicationToAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddApplicationToAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddApplicationToAuthorizationRuleResponseBody) *AddApplicationToAuthorizationRuleResponse
	GetBody() *AddApplicationToAuthorizationRuleResponseBody
}

type AddApplicationToAuthorizationRuleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddApplicationToAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddApplicationToAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddApplicationToAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *AddApplicationToAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddApplicationToAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddApplicationToAuthorizationRuleResponse) GetBody() *AddApplicationToAuthorizationRuleResponseBody {
	return s.Body
}

func (s *AddApplicationToAuthorizationRuleResponse) SetHeaders(v map[string]*string) *AddApplicationToAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *AddApplicationToAuthorizationRuleResponse) SetStatusCode(v int32) *AddApplicationToAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleResponse) SetBody(v *AddApplicationToAuthorizationRuleResponseBody) *AddApplicationToAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *AddApplicationToAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
