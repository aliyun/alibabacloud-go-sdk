// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForRegistrationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserGroupsForRegistrationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserGroupsForRegistrationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListUserGroupsForRegistrationPolicyResponseBody) *ListUserGroupsForRegistrationPolicyResponse
	GetBody() *ListUserGroupsForRegistrationPolicyResponseBody
}

type ListUserGroupsForRegistrationPolicyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsForRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsForRegistrationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserGroupsForRegistrationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserGroupsForRegistrationPolicyResponse) GetBody() *ListUserGroupsForRegistrationPolicyResponseBody {
	return s.Body
}

func (s *ListUserGroupsForRegistrationPolicyResponse) SetHeaders(v map[string]*string) *ListUserGroupsForRegistrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponse) SetStatusCode(v int32) *ListUserGroupsForRegistrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponse) SetBody(v *ListUserGroupsForRegistrationPolicyResponseBody) *ListUserGroupsForRegistrationPolicyResponse {
	s.Body = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
