// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageResult(v *ListProjectsResponseBodyPageResult) *ListProjectsResponseBody
	GetPageResult() *ListProjectsResponseBodyPageResult
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
}

type ListProjectsResponseBody struct {
	// The results that are returned.
	PageResult *ListProjectsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 20658801***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetPageResult() *ListProjectsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) SetPageResult(v *ListProjectsResponseBodyPageResult) *ListProjectsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyPageResult struct {
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
	// The DataWorks workspaces.
	ProjectList []*ListProjectsResponseBodyPageResultProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 123
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPageResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsResponseBodyPageResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsResponseBodyPageResult) GetProjectList() []*ListProjectsResponseBodyPageResultProjectList {
	return s.ProjectList
}

func (s *ListProjectsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectsResponseBodyPageResult) SetPageNumber(v int32) *ListProjectsResponseBodyPageResult {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBodyPageResult) SetPageSize(v int32) *ListProjectsResponseBodyPageResult {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBodyPageResult) SetProjectList(v []*ListProjectsResponseBodyPageResultProjectList) *ListProjectsResponseBodyPageResult {
	s.ProjectList = v
	return s
}

func (s *ListProjectsResponseBodyPageResult) SetTotalCount(v int32) *ListProjectsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsResponseBodyPageResult) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyPageResultProjectList struct {
	// Indicates whether the Development role is disabled. Valid values:
	//
	// 	- **false**: enabled
	//
	// 	- **true**: disabled
	//
	// example:
	//
	// true
	DisableDevelopment *bool `json:"DisableDevelopment,omitempty" xml:"DisableDevelopment,omitempty"`
	// Indicates whether the workspace is a default workspace. Valid values:
	//
	// 	- **1**: The workspace is a default workspace.
	//
	// 	- **0**: The workspace is not a default workspace.
	//
	// example:
	//
	// 1
	IsDefault *int32 `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// test_describe
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1212
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// test
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the user used by the workspace owner.
	//
	// example:
	//
	// 122222
	ProjectOwnerBaseId *string `json:"ProjectOwnerBaseId,omitempty" xml:"ProjectOwnerBaseId,omitempty"`
	// The status of the workspace. Valid values:
	//
	// 	- 0: AVAILABLE, which indicates that the workspace is running as expected.
	//
	// 	- 1: DELETED, which indicates that the workspace is deleted.
	//
	// 	- 2: INITIALIZING, which indicates that the workspace is being initialized.
	//
	// 	- 3: INIT_FAILED, which indicates that the workspace fails to be initialized.
	//
	// 	- 4: FORBIDDEN, which indicates that the workspace is manually disabled.
	//
	// 	- 5: DELETING, which indicates that the workspace is being deleted.
	//
	// 	- 6: DEL_FAILED, which indicates that the workspace fails to be deleted.
	//
	// 	- 7: FROZEN, which indicates that the workspace is frozen due to overdue payments.
	//
	// 	- 8: UPDATING, which indicates that the workspace is being updated. After you associate a compute engine with the workspace, the system initializes the compute engine and updates the workspace.
	//
	// 	- 9: UPDATE_FAILED, which indicates that the workspace fails to be updated.
	//
	// example:
	//
	// 0
	ProjectStatus *int32 `json:"ProjectStatus,omitempty" xml:"ProjectStatus,omitempty"`
	// The status code of the workspace. Valid values:
	//
	// 	- AVAILABLE: 0, which indicates that the workspace is running as expected.
	//
	// 	- DELETED: 1, which indicates that the workspace is deleted.
	//
	// 	- INITIALIZING: 2, which indicates that the workspace is being initialized.
	//
	// 	- INIT_FAILED: 3, which indicates that the workspace fails to be initialized.
	//
	// 	- FORBIDDEN: 4, which indicates that the workspace is manually disabled.
	//
	// 	- DELETING: 5, which indicates that the workspace is being deleted.
	//
	// 	- DEL_FAILED: 6, which indicates that the workspace fails to be deleted.
	//
	// 	- FROZEN: 7, which indicates that the workspace is frozen due to overdue payments.
	//
	// 	- UPDATING: 8, which indicates that the workspace is being updated. After you associate a compute engine with the workspace, the system initializes the compute engine and updates the workspace.
	//
	// 	- UPDATE_FAILED: 9, which indicates that the workspace fails to be updated.
	//
	// example:
	//
	// AVAILABLE
	ProjectStatusCode *string `json:"ProjectStatusCode,omitempty" xml:"ProjectStatusCode,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzbn7pti3zfa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Indicates whether the MaxCompute tables in the workspace are visible to the users within a tenant. Valid values:
	//
	// 	- **0**: invisible
	//
	// 	- **1**: visible
	//
	// example:
	//
	// 1
	TablePrivacyMode *int32 `json:"TablePrivacyMode,omitempty" xml:"TablePrivacyMode,omitempty"`
	// The tags added to the workspace.
	Tags []*ListProjectsResponseBodyPageResultProjectListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether a proxy account is used to access the MaxCompute compute engine associated with the workspace. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	UseProxyOdpsAccount *bool `json:"UseProxyOdpsAccount,omitempty" xml:"UseProxyOdpsAccount,omitempty"`
}

func (s ListProjectsResponseBodyPageResultProjectList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyPageResultProjectList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetDisableDevelopment() *bool {
	return s.DisableDevelopment
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetIsDefault() *int32 {
	return s.IsDefault
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetProjectDescription() *string {
	return s.ProjectDescription
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetProjectOwnerBaseId() *string {
	return s.ProjectOwnerBaseId
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetProjectStatus() *int32 {
	return s.ProjectStatus
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetProjectStatusCode() *string {
	return s.ProjectStatusCode
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetTablePrivacyMode() *int32 {
	return s.TablePrivacyMode
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetTags() []*ListProjectsResponseBodyPageResultProjectListTags {
	return s.Tags
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetUseProxyOdpsAccount() *bool {
	return s.UseProxyOdpsAccount
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetDisableDevelopment(v bool) *ListProjectsResponseBodyPageResultProjectList {
	s.DisableDevelopment = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetIsDefault(v int32) *ListProjectsResponseBodyPageResultProjectList {
	s.IsDefault = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetProjectDescription(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.ProjectDescription = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetProjectId(v int64) *ListProjectsResponseBodyPageResultProjectList {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetProjectIdentifier(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.ProjectIdentifier = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetProjectName(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.ProjectName = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetProjectOwnerBaseId(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.ProjectOwnerBaseId = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetProjectStatus(v int32) *ListProjectsResponseBodyPageResultProjectList {
	s.ProjectStatus = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetProjectStatusCode(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.ProjectStatusCode = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetResourceManagerResourceGroupId(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetTablePrivacyMode(v int32) *ListProjectsResponseBodyPageResultProjectList {
	s.TablePrivacyMode = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetTags(v []*ListProjectsResponseBodyPageResultProjectListTags) *ListProjectsResponseBodyPageResultProjectList {
	s.Tags = v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetUseProxyOdpsAccount(v bool) *ListProjectsResponseBodyPageResultProjectList {
	s.UseProxyOdpsAccount = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) Validate() error {
	return dara.Validate(s)
}

type ListProjectsResponseBodyPageResultProjectListTags struct {
	// The key of tag N added to the workspace.
	//
	// example:
	//
	// Env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the workspace.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectsResponseBodyPageResultProjectListTags) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyPageResultProjectListTags) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPageResultProjectListTags) GetKey() *string {
	return s.Key
}

func (s *ListProjectsResponseBodyPageResultProjectListTags) GetValue() *string {
	return s.Value
}

func (s *ListProjectsResponseBodyPageResultProjectListTags) SetKey(v string) *ListProjectsResponseBodyPageResultProjectListTags {
	s.Key = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectListTags) SetValue(v string) *ListProjectsResponseBodyPageResultProjectListTags {
	s.Value = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectListTags) Validate() error {
	return dara.Validate(s)
}
