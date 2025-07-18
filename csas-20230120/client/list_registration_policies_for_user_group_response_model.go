// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistrationPoliciesForUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRegistrationPoliciesForUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRegistrationPoliciesForUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListRegistrationPoliciesForUserGroupResponseBody) *ListRegistrationPoliciesForUserGroupResponse
	GetBody() *ListRegistrationPoliciesForUserGroupResponseBody
}

type ListRegistrationPoliciesForUserGroupResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegistrationPoliciesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegistrationPoliciesForUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRegistrationPoliciesForUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRegistrationPoliciesForUserGroupResponse) GetBody() *ListRegistrationPoliciesForUserGroupResponseBody {
	return s.Body
}

func (s *ListRegistrationPoliciesForUserGroupResponse) SetHeaders(v map[string]*string) *ListRegistrationPoliciesForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponse) SetStatusCode(v int32) *ListRegistrationPoliciesForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponse) SetBody(v *ListRegistrationPoliciesForUserGroupResponseBody) *ListRegistrationPoliciesForUserGroupResponse {
	s.Body = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
