// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateProjectRoleRequest
	GetClientToken() *string
	SetModulePermissions(v []*CreateProjectRoleRequestModulePermissions) *CreateProjectRoleRequest
	GetModulePermissions() []*CreateProjectRoleRequestModulePermissions
	SetName(v string) *CreateProjectRoleRequest
	GetName() *string
	SetProjectId(v int64) *CreateProjectRoleRequest
	GetProjectId() *int64
}

type CreateProjectRoleRequest struct {
	ClientToken       *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ModulePermissions []*CreateProjectRoleRequestModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
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

func (s CreateProjectRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProjectRoleRequest) GetModulePermissions() []*CreateProjectRoleRequestModulePermissions {
	return s.ModulePermissions
}

func (s *CreateProjectRoleRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectRoleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateProjectRoleRequest) SetClientToken(v string) *CreateProjectRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectRoleRequest) SetModulePermissions(v []*CreateProjectRoleRequestModulePermissions) *CreateProjectRoleRequest {
	s.ModulePermissions = v
	return s
}

func (s *CreateProjectRoleRequest) SetName(v string) *CreateProjectRoleRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRoleRequest) SetProjectId(v int64) *CreateProjectRoleRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectRoleRequest) Validate() error {
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

type CreateProjectRoleRequestModulePermissions struct {
	// example:
	//
	// 2
	ModuleId *int64 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// example:
	//
	// Write
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
}

func (s CreateProjectRoleRequestModulePermissions) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRoleRequestModulePermissions) GoString() string {
	return s.String()
}

func (s *CreateProjectRoleRequestModulePermissions) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *CreateProjectRoleRequestModulePermissions) GetPermissionType() *string {
	return s.PermissionType
}

func (s *CreateProjectRoleRequestModulePermissions) SetModuleId(v int64) *CreateProjectRoleRequestModulePermissions {
	s.ModuleId = &v
	return s
}

func (s *CreateProjectRoleRequestModulePermissions) SetPermissionType(v string) *CreateProjectRoleRequestModulePermissions {
	s.PermissionType = &v
	return s
}

func (s *CreateProjectRoleRequestModulePermissions) Validate() error {
	return dara.Validate(s)
}
