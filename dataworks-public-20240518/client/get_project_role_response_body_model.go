// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectRole(v *GetProjectRoleResponseBodyProjectRole) *GetProjectRoleResponseBody
	GetProjectRole() *GetProjectRoleResponseBodyProjectRole
	SetRequestId(v string) *GetProjectRoleResponseBody
	GetRequestId() *string
}

type GetProjectRoleResponseBody struct {
	// The role in the DataWorks workspace.
	ProjectRole *GetProjectRoleResponseBodyProjectRole `json:"ProjectRole,omitempty" xml:"ProjectRole,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 82F28E60-CF48-5EDF-AB25-D806847B97D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectRoleResponseBody) GetProjectRole() *GetProjectRoleResponseBodyProjectRole {
	return s.ProjectRole
}

func (s *GetProjectRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectRoleResponseBody) SetProjectRole(v *GetProjectRoleResponseBodyProjectRole) *GetProjectRoleResponseBody {
	s.ProjectRole = v
	return s
}

func (s *GetProjectRoleResponseBody) SetRequestId(v string) *GetProjectRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectRoleResponseBody) Validate() error {
	if s.ProjectRole != nil {
		if err := s.ProjectRole.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectRoleResponseBodyProjectRole struct {
	// The code of the role in the DataWorks workspace.
	//
	// example:
	//
	// role_project_guest
	Code              *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	ModulePermissions []*GetProjectRoleResponseBodyProjectRoleModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
	// The name of the role in the DataWorks workspace.
	//
	// example:
	//
	// Visitors
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the role in the DataWorks workspace. Valid values:
	//
	// 	- UserCustom: user-defined role
	//
	// 	- System: system role
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProjectRoleResponseBodyProjectRole) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRoleResponseBodyProjectRole) GoString() string {
	return s.String()
}

func (s *GetProjectRoleResponseBodyProjectRole) GetCode() *string {
	return s.Code
}

func (s *GetProjectRoleResponseBodyProjectRole) GetModulePermissions() []*GetProjectRoleResponseBodyProjectRoleModulePermissions {
	return s.ModulePermissions
}

func (s *GetProjectRoleResponseBodyProjectRole) GetName() *string {
	return s.Name
}

func (s *GetProjectRoleResponseBodyProjectRole) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetProjectRoleResponseBodyProjectRole) GetType() *string {
	return s.Type
}

func (s *GetProjectRoleResponseBodyProjectRole) SetCode(v string) *GetProjectRoleResponseBodyProjectRole {
	s.Code = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRole) SetModulePermissions(v []*GetProjectRoleResponseBodyProjectRoleModulePermissions) *GetProjectRoleResponseBodyProjectRole {
	s.ModulePermissions = v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRole) SetName(v string) *GetProjectRoleResponseBodyProjectRole {
	s.Name = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRole) SetProjectId(v int64) *GetProjectRoleResponseBodyProjectRole {
	s.ProjectId = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRole) SetType(v string) *GetProjectRoleResponseBodyProjectRole {
	s.Type = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRole) Validate() error {
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

type GetProjectRoleResponseBodyProjectRoleModulePermissions struct {
	// example:
	//
	// 2
	ModuleId *int64 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// example:
	//
	// HoloStudio
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// Read
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
}

func (s GetProjectRoleResponseBodyProjectRoleModulePermissions) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRoleResponseBodyProjectRoleModulePermissions) GoString() string {
	return s.String()
}

func (s *GetProjectRoleResponseBodyProjectRoleModulePermissions) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *GetProjectRoleResponseBodyProjectRoleModulePermissions) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetProjectRoleResponseBodyProjectRoleModulePermissions) GetPermissionType() *string {
	return s.PermissionType
}

func (s *GetProjectRoleResponseBodyProjectRoleModulePermissions) SetModuleId(v int64) *GetProjectRoleResponseBodyProjectRoleModulePermissions {
	s.ModuleId = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRoleModulePermissions) SetModuleName(v string) *GetProjectRoleResponseBodyProjectRoleModulePermissions {
	s.ModuleName = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRoleModulePermissions) SetPermissionType(v string) *GetProjectRoleResponseBodyProjectRoleModulePermissions {
	s.PermissionType = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRoleModulePermissions) Validate() error {
	return dara.Validate(s)
}
