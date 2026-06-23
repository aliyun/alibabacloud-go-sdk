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
	// The client token.
	//
	// example:
	//
	// 保留字段
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of DataWorks module permissions.
	ModulePermissions []*CreateProjectRoleRequestModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
	// The role name.
	//
	// This parameter is required.
	//
	// example:
	//
	// category_role
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the workspace management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace on which the API operation is performed.
	//
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
	// The DataWorks module ID. Valid values:
	//
	// - 2: HoloStudio
	//
	// - 3: StreamStudio
	//
	// - 4: Deployment management
	//
	// - 6: Data Security Guard
	//
	// - 7: Data Map
	//
	// - 8: Data Service
	//
	// - 9: Data Integration
	//
	// - 10: Data Modeling (DataBlau DDM)
	//
	// - 11: Data Studio
	//
	// - 12: Data Quality
	//
	// - 13: Data Governance
	//
	// - 14: Operation Center
	//
	// - 15: Resource optimization
	//
	// - 16: Migration Assistant
	//
	// - 17: Data Analysis
	//
	// - 18: Approval center
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
	// - Write: Read-only
	//
	// - Read: Edit
	//
	// - NotSet: Not controlled
	//
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
