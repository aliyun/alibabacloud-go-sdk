// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGroupFromAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveGroupFromAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveGroupFromAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *RemoveGroupFromAuthorizationRuleResponseBody) *RemoveGroupFromAuthorizationRuleResponse
	GetBody() *RemoveGroupFromAuthorizationRuleResponseBody
}

type RemoveGroupFromAuthorizationRuleResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveGroupFromAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveGroupFromAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveGroupFromAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupFromAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveGroupFromAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveGroupFromAuthorizationRuleResponse) GetBody() *RemoveGroupFromAuthorizationRuleResponseBody {
	return s.Body
}

func (s *RemoveGroupFromAuthorizationRuleResponse) SetHeaders(v map[string]*string) *RemoveGroupFromAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupFromAuthorizationRuleResponse) SetStatusCode(v int32) *RemoveGroupFromAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveGroupFromAuthorizationRuleResponse) SetBody(v *RemoveGroupFromAuthorizationRuleResponseBody) *RemoveGroupFromAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *RemoveGroupFromAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
