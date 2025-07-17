// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateProjectRequest
	GetDescription() *string
	SetDevEnvironmentEnabled(v bool) *UpdateProjectRequest
	GetDevEnvironmentEnabled() *bool
	SetDevRoleDisabled(v bool) *UpdateProjectRequest
	GetDevRoleDisabled() *bool
	SetDisplayName(v string) *UpdateProjectRequest
	GetDisplayName() *string
	SetId(v int64) *UpdateProjectRequest
	GetId() *int64
	SetPaiTaskEnabled(v bool) *UpdateProjectRequest
	GetPaiTaskEnabled() *bool
	SetStatus(v string) *UpdateProjectRequest
	GetStatus() *string
}

type UpdateProjectRequest struct {
	// The description of the workspace.
	//
	// example:
	//
	// Financial analysis group project data development
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the development environment. Valid values:
	//
	// 	- true: enables the development environment. In this case, the development environment is isolated from the production environment in the workspace.
	//
	// 	- false: disables the development environment. In this case, only the production environment is used in the workspace.
	//
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// Specifies whether to disable the Develop role. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// Note: If you disable the Develop role, you cannot assume the Develop role to develop nodes in workflows and edit node code. The Develop role cannot be enabled again after it is disabled.
	//
	// example:
	//
	// true
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// Sora financial analysis Space
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the workspace management page to obtain the ID.
	//
	// This parameter is used to determine the DataWorks workspaces used for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to enable scheduling of Platform for AI (PAI) tasks. Valid values:
	//
	// 	- true: enables scheduling of PAI tasks. In this case, you can create a PAI node in a DataWorks workspace and configure scheduling properties for the node to implement periodic scheduling of PAI tasks.
	//
	// 	- false: disables scheduling of PAI tasks.
	//
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// Specifies whether to disable or enable the workspace. Valid values:
	//
	// 	- Available: enables the workspace.
	//
	// 	- Forbidden: disables the workspace.
	//
	// example:
	//
	// Forbidden
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProjectRequest) GetDevEnvironmentEnabled() *bool {
	return s.DevEnvironmentEnabled
}

func (s *UpdateProjectRequest) GetDevRoleDisabled() *bool {
	return s.DevRoleDisabled
}

func (s *UpdateProjectRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateProjectRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateProjectRequest) GetPaiTaskEnabled() *bool {
	return s.PaiTaskEnabled
}

func (s *UpdateProjectRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetDevEnvironmentEnabled(v bool) *UpdateProjectRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *UpdateProjectRequest) SetDevRoleDisabled(v bool) *UpdateProjectRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *UpdateProjectRequest) SetDisplayName(v string) *UpdateProjectRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateProjectRequest) SetId(v int64) *UpdateProjectRequest {
	s.Id = &v
	return s
}

func (s *UpdateProjectRequest) SetPaiTaskEnabled(v bool) *UpdateProjectRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *UpdateProjectRequest) SetStatus(v string) *UpdateProjectRequest {
	s.Status = &v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}
