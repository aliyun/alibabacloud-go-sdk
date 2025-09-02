// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMemberToRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddProjectMemberToRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddProjectMemberToRoleResponse
	GetStatusCode() *int32
	SetBody(v *AddProjectMemberToRoleResponseBody) *AddProjectMemberToRoleResponse
	GetBody() *AddProjectMemberToRoleResponseBody
}

type AddProjectMemberToRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddProjectMemberToRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddProjectMemberToRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberToRoleResponse) GoString() string {
	return s.String()
}

func (s *AddProjectMemberToRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddProjectMemberToRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddProjectMemberToRoleResponse) GetBody() *AddProjectMemberToRoleResponseBody {
	return s.Body
}

func (s *AddProjectMemberToRoleResponse) SetHeaders(v map[string]*string) *AddProjectMemberToRoleResponse {
	s.Headers = v
	return s
}

func (s *AddProjectMemberToRoleResponse) SetStatusCode(v int32) *AddProjectMemberToRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProjectMemberToRoleResponse) SetBody(v *AddProjectMemberToRoleResponseBody) *AddProjectMemberToRoleResponse {
	s.Body = v
	return s
}

func (s *AddProjectMemberToRoleResponse) Validate() error {
	return dara.Validate(s)
}
