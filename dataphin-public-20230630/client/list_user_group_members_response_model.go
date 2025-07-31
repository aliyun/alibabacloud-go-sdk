// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListUserGroupMembersResponseBody) *ListUserGroupMembersResponse
	GetBody() *ListUserGroupMembersResponseBody
}

type ListUserGroupMembersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserGroupMembersResponse) GetBody() *ListUserGroupMembersResponseBody {
	return s.Body
}

func (s *ListUserGroupMembersResponse) SetHeaders(v map[string]*string) *ListUserGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupMembersResponse) SetStatusCode(v int32) *ListUserGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupMembersResponse) SetBody(v *ListUserGroupMembersResponseBody) *ListUserGroupMembersResponse {
	s.Body = v
	return s
}

func (s *ListUserGroupMembersResponse) Validate() error {
	return dara.Validate(s)
}
