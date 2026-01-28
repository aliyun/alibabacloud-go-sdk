// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForAuthorizationRuleResponseBody) *ListApplicationsForAuthorizationRuleResponse
	GetBody() *ListApplicationsForAuthorizationRuleResponseBody
}

type ListApplicationsForAuthorizationRuleResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForAuthorizationRuleResponse) GetBody() *ListApplicationsForAuthorizationRuleResponseBody {
	return s.Body
}

func (s *ListApplicationsForAuthorizationRuleResponse) SetHeaders(v map[string]*string) *ListApplicationsForAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponse) SetStatusCode(v int32) *ListApplicationsForAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponse) SetBody(v *ListApplicationsForAuthorizationRuleResponseBody) *ListApplicationsForAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
