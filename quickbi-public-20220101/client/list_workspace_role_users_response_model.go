// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceRoleUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkspaceRoleUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkspaceRoleUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkspaceRoleUsersResponseBody) *ListWorkspaceRoleUsersResponse
	GetBody() *ListWorkspaceRoleUsersResponseBody
}

type ListWorkspaceRoleUsersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceRoleUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkspaceRoleUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkspaceRoleUsersResponse) GetBody() *ListWorkspaceRoleUsersResponseBody {
	return s.Body
}

func (s *ListWorkspaceRoleUsersResponse) SetHeaders(v map[string]*string) *ListWorkspaceRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceRoleUsersResponse) SetStatusCode(v int32) *ListWorkspaceRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponse) SetBody(v *ListWorkspaceRoleUsersResponseBody) *ListWorkspaceRoleUsersResponse {
	s.Body = v
	return s
}

func (s *ListWorkspaceRoleUsersResponse) Validate() error {
	return dara.Validate(s)
}
