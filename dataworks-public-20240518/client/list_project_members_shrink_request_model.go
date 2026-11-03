// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProjectMembersShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectMembersShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListProjectMembersShrinkRequest
	GetProjectId() *int64
	SetRoleCodesShrink(v string) *ListProjectMembersShrinkRequest
	GetRoleCodesShrink() *string
	SetUserIdsShrink(v string) *ListProjectMembersShrinkRequest
	GetUserIdsShrink() *string
}

type ListProjectMembersShrinkRequest struct {
	// The page number of the requested data.
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
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace used when you invoke this API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 62136
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of workspace role codes. You can call the [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) operation to obtain the role codes.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// The list of DataWorks account IDs. You can log on to the [DataWorks console - Administration](https://dataworks.console.aliyun.com/product/ms_menu), select the workspace that you want to manage, go to the Administration page, and then go to the Tenant Members and Roles page to view the account IDs.
	UserIdsShrink *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s ListProjectMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectMembersShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectMembersShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectMembersShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *ListProjectMembersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *ListProjectMembersShrinkRequest) SetPageNumber(v int32) *ListProjectMembersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetPageSize(v int32) *ListProjectMembersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetProjectId(v int64) *ListProjectMembersShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetRoleCodesShrink(v string) *ListProjectMembersShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetUserIdsShrink(v string) *ListProjectMembersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
