// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorizationRuleResponseBody) *UpdateAuthorizationRuleResponse
	GetBody() *UpdateAuthorizationRuleResponseBody
}

type UpdateAuthorizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorizationRuleResponse) GetBody() *UpdateAuthorizationRuleResponseBody {
	return s.Body
}

func (s *UpdateAuthorizationRuleResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationRuleResponse) SetStatusCode(v int32) *UpdateAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationRuleResponse) SetBody(v *UpdateAuthorizationRuleResponseBody) *UpdateAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
