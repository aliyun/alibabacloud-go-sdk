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
	// The query result.
	PageResult *ListProjectsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 20658801****
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
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyPageResult struct {
	// The current page number.
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
	// The list of DataWorks workspaces.
	ProjectList []*ListProjectsResponseBodyPageResultProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 68
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
	if s.ProjectList != nil {
		for _, item := range s.ProjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyPageResultProjectList struct {
	// Indicates whether the development role is disabled. Valid values:
	//
	// -  **false**: The development role is enabled.
	//
	// -  **true**: The development role is disabled.
	//
	// example:
	//
	// true
	DisableDevelopment *bool `json:"DisableDevelopment,omitempty" xml:"DisableDevelopment,omitempty"`
	// Indicates whether the workspace is the default workspace. Valid values:
	//
	// - **1**: yes.
	//
	// - **0**: no.
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
	// 466230
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// test_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The user ID of the workspace owner.
	//
	// example:
	//
	// 13969939245****
	ProjectOwnerBaseId *string `json:"ProjectOwnerBaseId,omitempty" xml:"ProjectOwnerBaseId,omitempty"`
	// The status of the workspace. Valid values:
	//
	// - AVAILABLE: The status value is 0, which indicates that the workspace is Normal.
	//
	// - DELETED: The status value is 1, which indicates that the workspace is deleted.
	//
	// - INITIALIZING: The status value is 2, which indicates that the workspace is being initialized.
	//
	// - INIT_FAILED: The status value is 3, which indicates that the workspace failed to be initialized.
	//
	// - FORBIDDEN: The status value is 4, which indicates that the workspace is manually disabled.
	//
	// - DELETING: The status value is 5, which indicates that the workspace is being deleted.
	//
	// - DEL_FAILED: The status value is 6, which indicates that the workspace failed to be deleted.
	//
	// - FROZEN: The status value is 7, which indicates that the workspace is frozen due to overdue payment.
	//
	// - UPDATING: The status value is 8, which indicates that the workspace is being updated (a compute engine is being added and initialized for the project).
	//
	// - UPDATE_FAILED: The status value is 9, which indicates that the workspace failed to be updated (a compute engine failed to be added and initialized for the project).
	//
	// example:
	//
	// 0
	ProjectStatus *int32 `json:"ProjectStatus,omitempty" xml:"ProjectStatus,omitempty"`
	// The status code of the workspace. Valid values:
	//
	// - AVAILABLE: The status value is 0, which indicates that the workspace is Normal.
	//
	// - DELETED: The status value is 1, which indicates that the workspace is deleted.
	//
	// - INITIALIZING: The status value is 2, which indicates that the workspace is being initialized.
	//
	// - INIT_FAILED: The status value is 3, which indicates that the workspace failed to be initialized.
	//
	// - FORBIDDEN: The status value is 4, which indicates that the workspace is manually disabled.
	//
	// - DELETING: The status value is 5, which indicates that the workspace is being deleted.
	//
	// - DEL_FAILED: The status value is 6, which indicates that the workspace failed to be deleted.
	//
	// - FROZEN: The status value is 7, which indicates that the workspace is frozen due to overdue payment.
	//
	// - UPDATING: The status value is 8, which indicates that the workspace is being updated (a compute engine is being added and initialized for the project).
	//
	// - UPDATE_FAILED: The status value is 9, which indicates that the workspace failed to be updated (a compute engine failed to be added and initialized for the project).
	//
	// example:
	//
	// AVAILABLE
	ProjectStatusCode *string `json:"ProjectStatusCode,omitempty" xml:"ProjectStatusCode,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzbn7****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The visibility permission of MaxCompute tables. Valid values:
	//
	// - **0**: MaxCompute tables are not visible to users within the tenant.
	//
	// - **1**: MaxCompute tables are visible to users within the tenant.
	//
	// example:
	//
	// 1
	TablePrivacyMode *int32 `json:"TablePrivacyMode,omitempty" xml:"TablePrivacyMode,omitempty"`
	// The list of tags bound to the workspace.
	Tags []*ListProjectsResponseBodyPageResultProjectListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether a proxy account is used to access the MaxCompute engine. Valid values:
	//
	// - **false**: A proxy account is not used.
	//
	// - **true**: A proxy account is used.
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyPageResultProjectListTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
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
