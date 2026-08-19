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
	// A reserved field.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique identifier of the custom role.
	//
	// This parameter is required.
	//
	// example:
	//
	// base_role_xx
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of DataWorks module permissions.
	//
	// This parameter is required.
	ModulePermissions []*UpdateProjectRoleRequestModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the Storage Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace for this API invocation.
	//
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
	// The DataWorks module ID. Valid values:
	//
	// - 2: HoloStudio
	//
	// - 3: StreamStudio
	//
	// - 4: Deploy Management
	//
	// - 6: Data Protection Umbrella
	//
	// - 7: Data Map
	//
	// - 8: DataService Studio
	//
	// - 9: Data Integration
	//
	// - 10: Data Modeling (DataBlau DDM)
	//
	// - 11: DataStudio
	//
	// - 12: Data Quality
	//
	// - 13: Data Governance Center
	//
	// - 14: Operation Center
	//
	// - 15: Resource Optimization
	//
	// - 16: Migration Assistant
	//
	// - 17: Data Analytics
	//
	// - 18: Approval Center
	//
	// - 19: Security Center
	//
	// - 20: Intelligent Data Modeling
	//
	// example:
	//
	// 2
	ModuleId *int64 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// The permission type. Valid values:
	//
	// - Write: Edit.
	//
	// - Read: Read-only.
	//
	// - NotSet: Not controlled.
	//
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
