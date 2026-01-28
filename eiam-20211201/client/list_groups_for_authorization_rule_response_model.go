// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupsForAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupsForAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupsForAuthorizationRuleResponseBody) *ListGroupsForAuthorizationRuleResponse
	GetBody() *ListGroupsForAuthorizationRuleResponseBody
}

type ListGroupsForAuthorizationRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupsForAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupsForAuthorizationRuleResponse) GetBody() *ListGroupsForAuthorizationRuleResponseBody {
	return s.Body
}

func (s *ListGroupsForAuthorizationRuleResponse) SetHeaders(v map[string]*string) *ListGroupsForAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponse) SetStatusCode(v int32) *ListGroupsForAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponse) SetBody(v *ListGroupsForAuthorizationRuleResponseBody) *ListGroupsForAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
