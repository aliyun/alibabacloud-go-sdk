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
	// An optional description of the workspace.
	//
	// example:
	//
	// Financial analysis group project data development
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the development environment. Valid values:
	//
	// - `true`: Enables the development environment for the workspace and isolates it from the production environment.
	//
	// - `false`: Uses only the production environment.
	//
	// **Important**: You cannot disable the development environment after you enable it.
	//
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// Specifies whether to disable the development role, which grants permissions for workflow and code editing. Valid values:
	//
	// - `false`: Enables the development role. This is the default value.
	//
	// - `true`: Disables the development role.
	//
	// **Important**: After you enable the development role (by setting this parameter to `false`), you cannot disable it.
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
	// The ID of the DataWorks workspace. To find the workspace ID, log in to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the Workspace Management page.
	//
	// This parameter specifies the DataWorks workspace to use for the API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to enable task scheduling for Machine Learning Platform for AI (PAI). Valid values:
	//
	// - `true`: You can create PAI nodes in the DataWorks workspace and run them on a schedule.
	//
	// - `false`: Disables task scheduling for PAI.
	//
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// Specifies whether to enable or disable the workspace. Valid values:
	//
	// - `Available`: Enables the workspace.
	//
	// - `Forbidden`: Disables the workspace.
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
