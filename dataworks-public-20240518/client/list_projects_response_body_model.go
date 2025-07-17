// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListProjectsResponseBodyPagingInfo) *ListProjectsResponseBody
	GetPagingInfo() *ListProjectsResponseBodyPagingInfo
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
}

type ListProjectsResponseBody struct {
	// The pagination information.
	PagingInfo *ListProjectsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6D24AD9A-652F-59E2-AC1F-05029300F8A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetPagingInfo() *ListProjectsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) SetPagingInfo(v *ListProjectsResponseBodyPagingInfo) *ListProjectsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspaces.
	Projects []*ListProjectsResponseBodyPagingInfoProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsResponseBodyPagingInfo) GetProjects() []*ListProjectsResponseBodyPagingInfoProjects {
	return s.Projects
}

func (s *ListProjectsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectsResponseBodyPagingInfo) SetPageNumber(v int32) *ListProjectsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetPageSize(v int32) *ListProjectsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetProjects(v []*ListProjectsResponseBodyPagingInfoProjects) *ListProjectsResponseBodyPagingInfo {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetTotalCount(v int32) *ListProjectsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyPagingInfoProjects struct {
	// The ID of the Alibaba Cloud resource group to which the workspace belongs.
	//
	// example:
	//
	// rg-acfmzbn7pti3zfa
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The tags.
	AliyunResourceTags []*ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// The description of the workspace.
	//
	// example:
	//
	// Financial analysis group project data development
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the development environment is enabled. Valid values:
	//
	// 	- true: The development environment is enabled. In this case, the development environment is isolated from the production environment in the workspace.
	//
	// 	- false: The development environment is disabled. In this case, only the production environment is used in the workspace.
	//
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// Indicates whether the Develop role is disabled. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// Sora financial analysis
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account to which the workspace belongs.
	//
	// example:
	//
	// 123532153125
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Indicates whether scheduling of PAI tasks is enabled. Valid values:
	//
	// 	- true: Scheduling of PAI tasks is enabled. In this case, you can create a PAI node in a DataWorks workspace and configure scheduling properties for the node to implement periodic scheduling of PAI tasks.
	//
	// 	- false: Scheduling of PAI tasks is disabled.
	//
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// The status of the workspace. Valid values:
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

func (s ListProjectsResponseBodyPagingInfoProjects) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfoProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetAliyunResourceTags() []*ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags {
	return s.AliyunResourceTags
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetDescription() *string {
	return s.Description
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetDevEnvironmentEnabled() *bool {
	return s.DevEnvironmentEnabled
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetDevRoleDisabled() *bool {
	return s.DevRoleDisabled
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetId() *int64 {
	return s.Id
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetName() *string {
	return s.Name
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetOwner() *string {
	return s.Owner
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetPaiTaskEnabled() *bool {
	return s.PaiTaskEnabled
}

func (s *ListProjectsResponseBodyPagingInfoProjects) GetStatus() *string {
	return s.Status
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetAliyunResourceGroupId(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetAliyunResourceTags(v []*ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) *ListProjectsResponseBodyPagingInfoProjects {
	s.AliyunResourceTags = v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDescription(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDevEnvironmentEnabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDevRoleDisabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDisplayName(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.DisplayName = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetId(v int64) *ListProjectsResponseBodyPagingInfoProjects {
	s.Id = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetName(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetOwner(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetPaiTaskEnabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetStatus(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Status = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags struct {
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

func (s ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) GetKey() *string {
	return s.Key
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) GetValue() *string {
	return s.Value
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) SetKey(v string) *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) SetValue(v string) *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags {
	s.Value = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) Validate() error {
	return dara.Validate(s)
}
