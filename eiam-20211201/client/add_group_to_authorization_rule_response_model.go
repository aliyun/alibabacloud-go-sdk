// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupToAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGroupToAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGroupToAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddGroupToAuthorizationRuleResponseBody) *AddGroupToAuthorizationRuleResponse
	GetBody() *AddGroupToAuthorizationRuleResponseBody
}

type AddGroupToAuthorizationRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGroupToAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGroupToAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGroupToAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *AddGroupToAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGroupToAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGroupToAuthorizationRuleResponse) GetBody() *AddGroupToAuthorizationRuleResponseBody {
	return s.Body
}

func (s *AddGroupToAuthorizationRuleResponse) SetHeaders(v map[string]*string) *AddGroupToAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *AddGroupToAuthorizationRuleResponse) SetStatusCode(v int32) *AddGroupToAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGroupToAuthorizationRuleResponse) SetBody(v *AddGroupToAuthorizationRuleResponseBody) *AddGroupToAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *AddGroupToAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
