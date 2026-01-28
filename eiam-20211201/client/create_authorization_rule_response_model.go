// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateAuthorizationRuleResponseBody) *CreateAuthorizationRuleResponse
	GetBody() *CreateAuthorizationRuleResponseBody
}

type CreateAuthorizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAuthorizationRuleResponse) GetBody() *CreateAuthorizationRuleResponseBody {
	return s.Body
}

func (s *CreateAuthorizationRuleResponse) SetHeaders(v map[string]*string) *CreateAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthorizationRuleResponse) SetStatusCode(v int32) *CreateAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthorizationRuleResponse) SetBody(v *CreateAuthorizationRuleResponseBody) *CreateAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
