// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateProjectRoleShrinkRequest
	GetClientToken() *string
	SetCode(v string) *UpdateProjectRoleShrinkRequest
	GetCode() *string
	SetModulePermissionsShrink(v string) *UpdateProjectRoleShrinkRequest
	GetModulePermissionsShrink() *string
	SetProjectId(v int64) *UpdateProjectRoleShrinkRequest
	GetProjectId() *int64
}

type UpdateProjectRoleShrinkRequest struct {
	// example:
	//
	// 0000-ABCD-EFG****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// base_role_xx
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	ModulePermissionsShrink *string `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateProjectRoleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRoleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateProjectRoleShrinkRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateProjectRoleShrinkRequest) GetModulePermissionsShrink() *string {
	return s.ModulePermissionsShrink
}

func (s *UpdateProjectRoleShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateProjectRoleShrinkRequest) SetClientToken(v string) *UpdateProjectRoleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProjectRoleShrinkRequest) SetCode(v string) *UpdateProjectRoleShrinkRequest {
	s.Code = &v
	return s
}

func (s *UpdateProjectRoleShrinkRequest) SetModulePermissionsShrink(v string) *UpdateProjectRoleShrinkRequest {
	s.ModulePermissionsShrink = &v
	return s
}

func (s *UpdateProjectRoleShrinkRequest) SetProjectId(v int64) *UpdateProjectRoleShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectRoleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
