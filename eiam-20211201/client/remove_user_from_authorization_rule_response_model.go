// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserFromAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserFromAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserFromAuthorizationRuleResponseBody) *RemoveUserFromAuthorizationRuleResponse
	GetBody() *RemoveUserFromAuthorizationRuleResponseBody
}

type RemoveUserFromAuthorizationRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserFromAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserFromAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserFromAuthorizationRuleResponse) GetBody() *RemoveUserFromAuthorizationRuleResponseBody {
	return s.Body
}

func (s *RemoveUserFromAuthorizationRuleResponse) SetHeaders(v map[string]*string) *RemoveUserFromAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromAuthorizationRuleResponse) SetStatusCode(v int32) *RemoveUserFromAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromAuthorizationRuleResponse) SetBody(v *RemoveUserFromAuthorizationRuleResponseBody) *RemoveUserFromAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *RemoveUserFromAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
