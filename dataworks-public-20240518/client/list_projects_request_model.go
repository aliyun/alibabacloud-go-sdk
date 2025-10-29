// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *ListProjectsRequest
	GetAliyunResourceGroupId() *string
	SetAliyunResourceTags(v []*ListProjectsRequestAliyunResourceTags) *ListProjectsRequest
	GetAliyunResourceTags() []*ListProjectsRequestAliyunResourceTags
	SetDevEnvironmentEnabled(v bool) *ListProjectsRequest
	GetDevEnvironmentEnabled() *bool
	SetDevRoleDisabled(v bool) *ListProjectsRequest
	GetDevRoleDisabled() *bool
	SetIds(v []*int64) *ListProjectsRequest
	GetIds() []*int64
	SetNames(v []*string) *ListProjectsRequest
	GetNames() []*string
	SetPageNumber(v int32) *ListProjectsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectsRequest
	GetPageSize() *int32
	SetPaiTaskEnabled(v bool) *ListProjectsRequest
	GetPaiTaskEnabled() *bool
	SetStatus(v string) *ListProjectsRequest
	GetStatus() *string
}

type ListProjectsRequest struct {
	// The ID of the Alibaba Cloud resource group to which the workspaces belong. You can log on to the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups) and go to the Resource Group page to query the ID.
	//
	// This parameter is used to query the information about workspaces that belong to a specific resource group.
	//
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The tags.
	AliyunResourceTags []*ListProjectsRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
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
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The names of the DataWorks workspaces.
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
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

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *ListProjectsRequest) GetAliyunResourceTags() []*ListProjectsRequestAliyunResourceTags {
	return s.AliyunResourceTags
}

func (s *ListProjectsRequest) GetDevEnvironmentEnabled() *bool {
	return s.DevEnvironmentEnabled
}

func (s *ListProjectsRequest) GetDevRoleDisabled() *bool {
	return s.DevRoleDisabled
}

func (s *ListProjectsRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *ListProjectsRequest) GetNames() []*string {
	return s.Names
}

func (s *ListProjectsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsRequest) GetPaiTaskEnabled() *bool {
	return s.PaiTaskEnabled
}

func (s *ListProjectsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListProjectsRequest) SetAliyunResourceGroupId(v string) *ListProjectsRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsRequest) SetAliyunResourceTags(v []*ListProjectsRequestAliyunResourceTags) *ListProjectsRequest {
	s.AliyunResourceTags = v
	return s
}

func (s *ListProjectsRequest) SetDevEnvironmentEnabled(v bool) *ListProjectsRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsRequest) SetDevRoleDisabled(v bool) *ListProjectsRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsRequest) SetIds(v []*int64) *ListProjectsRequest {
	s.Ids = v
	return s
}

func (s *ListProjectsRequest) SetNames(v []*string) *ListProjectsRequest {
	s.Names = v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v int32) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetPaiTaskEnabled(v bool) *ListProjectsRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsRequest) SetStatus(v string) *ListProjectsRequest {
	s.Status = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	if s.AliyunResourceTags != nil {
		for _, item := range s.AliyunResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsRequestAliyunResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectsRequestAliyunResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequestAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *ListProjectsRequestAliyunResourceTags) GetKey() *string {
	return s.Key
}

func (s *ListProjectsRequestAliyunResourceTags) GetValue() *string {
	return s.Value
}

func (s *ListProjectsRequestAliyunResourceTags) SetKey(v string) *ListProjectsRequestAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *ListProjectsRequestAliyunResourceTags) SetValue(v string) *ListProjectsRequestAliyunResourceTags {
	s.Value = &v
	return s
}

func (s *ListProjectsRequestAliyunResourceTags) Validate() error {
	return dara.Validate(s)
}
