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
	// The page number. Pages start from page 1.
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
	// The ID of the DataWorks workspace. To obtain this ID, log in to the [DataWorks Console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page.
	//
	// This parameter specifies the DataWorks workspace for the API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 62136
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// A list of role codes for the workspace to filter the results. You can call the [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) operation to obtain the role codes.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// A list of DataWorks user IDs to filter the results. You can find these IDs on the [Tenant Members and Roles](https://dataworks.console.aliyun.com/product/ms_menu) page in the Management Center.
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
