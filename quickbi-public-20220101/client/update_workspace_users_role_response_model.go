// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUsersRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceUsersRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceUsersRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkspaceUsersRoleResponseBody) *UpdateWorkspaceUsersRoleResponse
	GetBody() *UpdateWorkspaceUsersRoleResponseBody
}

type UpdateWorkspaceUsersRoleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceUsersRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceUsersRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUsersRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUsersRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceUsersRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceUsersRoleResponse) GetBody() *UpdateWorkspaceUsersRoleResponseBody {
	return s.Body
}

func (s *UpdateWorkspaceUsersRoleResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceUsersRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponse) SetStatusCode(v int32) *UpdateWorkspaceUsersRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponse) SetBody(v *UpdateWorkspaceUsersRoleResponseBody) *UpdateWorkspaceUsersRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponse) Validate() error {
	return dara.Validate(s)
}
