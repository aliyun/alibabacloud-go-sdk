// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateProjectRoleShrinkRequest
	GetClientToken() *string
	SetModulePermissionsShrink(v string) *CreateProjectRoleShrinkRequest
	GetModulePermissionsShrink() *string
	SetName(v string) *CreateProjectRoleShrinkRequest
	GetName() *string
	SetProjectId(v int64) *CreateProjectRoleShrinkRequest
	GetProjectId() *int64
}

type CreateProjectRoleShrinkRequest struct {
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ModulePermissionsShrink *string `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// category_role
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateProjectRoleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRoleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProjectRoleShrinkRequest) GetModulePermissionsShrink() *string {
	return s.ModulePermissionsShrink
}

func (s *CreateProjectRoleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectRoleShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateProjectRoleShrinkRequest) SetClientToken(v string) *CreateProjectRoleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectRoleShrinkRequest) SetModulePermissionsShrink(v string) *CreateProjectRoleShrinkRequest {
	s.ModulePermissionsShrink = &v
	return s
}

func (s *CreateProjectRoleShrinkRequest) SetName(v string) *CreateProjectRoleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRoleShrinkRequest) SetProjectId(v int64) *CreateProjectRoleShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectRoleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
