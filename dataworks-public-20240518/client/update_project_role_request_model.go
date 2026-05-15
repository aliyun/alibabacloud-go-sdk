// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateProjectRoleRequest
	GetClientToken() *string
	SetCode(v string) *UpdateProjectRoleRequest
	GetCode() *string
	SetModulePermissions(v []*UpdateProjectRoleRequestModulePermissions) *UpdateProjectRoleRequest
	GetModulePermissions() []*UpdateProjectRoleRequestModulePermissions
	SetProjectId(v int64) *UpdateProjectRoleRequest
	GetProjectId() *int64
}

type UpdateProjectRoleRequest struct {
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
	ModulePermissions []*UpdateProjectRoleRequestModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateProjectRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateProjectRoleRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateProjectRoleRequest) GetModulePermissions() []*UpdateProjectRoleRequestModulePermissions {
	return s.ModulePermissions
}

func (s *UpdateProjectRoleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateProjectRoleRequest) SetClientToken(v string) *UpdateProjectRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProjectRoleRequest) SetCode(v string) *UpdateProjectRoleRequest {
	s.Code = &v
	return s
}

func (s *UpdateProjectRoleRequest) SetModulePermissions(v []*UpdateProjectRoleRequestModulePermissions) *UpdateProjectRoleRequest {
	s.ModulePermissions = v
	return s
}

func (s *UpdateProjectRoleRequest) SetProjectId(v int64) *UpdateProjectRoleRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectRoleRequest) Validate() error {
	if s.ModulePermissions != nil {
		for _, item := range s.ModulePermissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProjectRoleRequestModulePermissions struct {
	// example:
	//
	// 2
	ModuleId *int64 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// example:
	//
	// Write
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
}

func (s UpdateProjectRoleRequestModulePermissions) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRoleRequestModulePermissions) GoString() string {
	return s.String()
}

func (s *UpdateProjectRoleRequestModulePermissions) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *UpdateProjectRoleRequestModulePermissions) GetPermissionType() *string {
	return s.PermissionType
}

func (s *UpdateProjectRoleRequestModulePermissions) SetModuleId(v int64) *UpdateProjectRoleRequestModulePermissions {
	s.ModuleId = &v
	return s
}

func (s *UpdateProjectRoleRequestModulePermissions) SetPermissionType(v string) *UpdateProjectRoleRequestModulePermissions {
	s.PermissionType = &v
	return s
}

func (s *UpdateProjectRoleRequestModulePermissions) Validate() error {
	return dara.Validate(s)
}
