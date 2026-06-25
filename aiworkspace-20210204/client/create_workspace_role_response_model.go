// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkspaceRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkspaceRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkspaceRoleResponseBody) *CreateWorkspaceRoleResponse
	GetBody() *CreateWorkspaceRoleResponseBody
}

type CreateWorkspaceRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkspaceRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkspaceRoleResponse) GetBody() *CreateWorkspaceRoleResponseBody {
	return s.Body
}

func (s *CreateWorkspaceRoleResponse) SetHeaders(v map[string]*string) *CreateWorkspaceRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceRoleResponse) SetStatusCode(v int32) *CreateWorkspaceRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceRoleResponse) SetBody(v *CreateWorkspaceRoleResponseBody) *CreateWorkspaceRoleResponse {
	s.Body = v
	return s
}

func (s *CreateWorkspaceRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
