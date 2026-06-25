// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkspaceRoleResponseBody) *UpdateWorkspaceRoleResponse
	GetBody() *UpdateWorkspaceRoleResponseBody
}

type UpdateWorkspaceRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceRoleResponse) GetBody() *UpdateWorkspaceRoleResponseBody {
	return s.Body
}

func (s *UpdateWorkspaceRoleResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceRoleResponse) SetStatusCode(v int32) *UpdateWorkspaceRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceRoleResponse) SetBody(v *UpdateWorkspaceRoleResponseBody) *UpdateWorkspaceRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
