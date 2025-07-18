// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForPrivateAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserGroupsForPrivateAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserGroupsForPrivateAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListUserGroupsForPrivateAccessPolicyResponseBody) *ListUserGroupsForPrivateAccessPolicyResponse
	GetBody() *ListUserGroupsForPrivateAccessPolicyResponseBody
}

type ListUserGroupsForPrivateAccessPolicyResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) GetBody() *ListUserGroupsForPrivateAccessPolicyResponseBody {
	return s.Body
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *ListUserGroupsForPrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) SetStatusCode(v int32) *ListUserGroupsForPrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) SetBody(v *ListUserGroupsForPrivateAccessPolicyResponseBody) *ListUserGroupsForPrivateAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
