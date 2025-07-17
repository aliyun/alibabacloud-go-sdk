// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *ListProjectsShrinkRequest
	GetAliyunResourceGroupId() *string
	SetAliyunResourceTagsShrink(v string) *ListProjectsShrinkRequest
	GetAliyunResourceTagsShrink() *string
	SetDevEnvironmentEnabled(v bool) *ListProjectsShrinkRequest
	GetDevEnvironmentEnabled() *bool
	SetDevRoleDisabled(v bool) *ListProjectsShrinkRequest
	GetDevRoleDisabled() *bool
	SetIdsShrink(v string) *ListProjectsShrinkRequest
	GetIdsShrink() *string
	SetNamesShrink(v string) *ListProjectsShrinkRequest
	GetNamesShrink() *string
	SetPageNumber(v int32) *ListProjectsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectsShrinkRequest
	GetPageSize() *int32
	SetPaiTaskEnabled(v bool) *ListProjectsShrinkRequest
	GetPaiTaskEnabled() *bool
	SetStatus(v string) *ListProjectsShrinkRequest
	GetStatus() *string
}

type ListProjectsShrinkRequest struct {
	// The ID of the Alibaba Cloud resource group to which the workspaces belong. You can log on to the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups) and go to the Resource Group page to query the ID.
	//
	// This parameter is used to query the information about workspaces that belong to a specific resource group.
	//
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The tags.
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	// Specifies whether the development environment is enabled. Valid values:
	//
	// 	- true: The development environment is enabled. In this case, the development environment is isolated from the production environment in a workspace.
	//
	// 	- false: The development environment is disabled. In this case, only the production environment is used in a workspace.
	//
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// Specifies whether the Develop role is disabled. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// The IDs of the DataWorks workspaces.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The names of the DataWorks workspaces.
	NamesShrink *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether scheduling of Platform for AI (PAI) tasks is enabled. Valid values:
	//
	// 	- true: Scheduling of PAI tasks is enabled. In this case, you can create a PAI node in a DataWorks workspace and configure scheduling properties for the node to implement periodic scheduling of PAI tasks.
	//
	// 	- false: Scheduling of PAI tasks is disabled.
	//
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// The status of the workspaces. Valid values:
	//
	// 	- Available
	//
	// 	- Initializing
	//
	// 	- InitFailed
	//
	// 	- Forbidden
	//
	// 	- Deleting
	//
	// 	- DeleteFailed
	//
	// 	- Frozen
	//
	// 	- Updating
	//
	// 	- UpdateFailed
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *ListProjectsShrinkRequest) GetAliyunResourceTagsShrink() *string {
	return s.AliyunResourceTagsShrink
}

func (s *ListProjectsShrinkRequest) GetDevEnvironmentEnabled() *bool {
	return s.DevEnvironmentEnabled
}

func (s *ListProjectsShrinkRequest) GetDevRoleDisabled() *bool {
	return s.DevRoleDisabled
}

func (s *ListProjectsShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *ListProjectsShrinkRequest) GetNamesShrink() *string {
	return s.NamesShrink
}

func (s *ListProjectsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsShrinkRequest) GetPaiTaskEnabled() *bool {
	return s.PaiTaskEnabled
}

func (s *ListProjectsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListProjectsShrinkRequest) SetAliyunResourceGroupId(v string) *ListProjectsShrinkRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetAliyunResourceTagsShrink(v string) *ListProjectsShrinkRequest {
	s.AliyunResourceTagsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetDevEnvironmentEnabled(v bool) *ListProjectsShrinkRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetDevRoleDisabled(v bool) *ListProjectsShrinkRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetIdsShrink(v string) *ListProjectsShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetNamesShrink(v string) *ListProjectsShrinkRequest {
	s.NamesShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageNumber(v int32) *ListProjectsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageSize(v int32) *ListProjectsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPaiTaskEnabled(v bool) *ListProjectsShrinkRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetStatus(v string) *ListProjectsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListProjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
