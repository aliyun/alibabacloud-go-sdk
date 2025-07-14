// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUserRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceUserRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceUserRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkspaceUserRoleResponseBody) *UpdateWorkspaceUserRoleResponse
	GetBody() *UpdateWorkspaceUserRoleResponseBody
}

type UpdateWorkspaceUserRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceUserRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceUserRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUserRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceUserRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceUserRoleResponse) GetBody() *UpdateWorkspaceUserRoleResponseBody {
	return s.Body
}

func (s *UpdateWorkspaceUserRoleResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceUserRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceUserRoleResponse) SetStatusCode(v int32) *UpdateWorkspaceUserRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceUserRoleResponse) SetBody(v *UpdateWorkspaceUserRoleResponseBody) *UpdateWorkspaceUserRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceUserRoleResponse) Validate() error {
	return dara.Validate(s)
}
