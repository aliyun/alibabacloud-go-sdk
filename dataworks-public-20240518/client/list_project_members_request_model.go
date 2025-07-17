// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProjectMembersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectMembersRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListProjectMembersRequest
	GetProjectId() *int64
	SetRoleCodes(v []*string) *ListProjectMembersRequest
	GetRoleCodes() []*string
	SetUserIds(v []*string) *ListProjectMembersRequest
	GetUserIds() []*string
}

type ListProjectMembersRequest struct {
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
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 62136
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The codes of the roles in the workspace. You can call the [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) operation to query the codes of all roles in the workspace.
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// The IDs of the accounts used by the members in the workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/product/ms_menu), choose More > Management Center in the left-side navigation pane, select the desired workspace on the Management Center page, and then click Go to Management Center. In the left-side navigation pane of the SettingCenter page, click Tenant Members and Roles. On the Tenant Members and Roles page, view the IDs of the accounts used by the members in the workspace.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ListProjectMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectMembersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectMembersRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectMembersRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *ListProjectMembersRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *ListProjectMembersRequest) SetPageNumber(v int32) *ListProjectMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectMembersRequest) SetPageSize(v int32) *ListProjectMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersRequest) SetProjectId(v int64) *ListProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectMembersRequest) SetRoleCodes(v []*string) *ListProjectMembersRequest {
	s.RoleCodes = v
	return s
}

func (s *ListProjectMembersRequest) SetUserIds(v []*string) *ListProjectMembersRequest {
	s.UserIds = v
	return s
}

func (s *ListProjectMembersRequest) Validate() error {
	return dara.Validate(s)
}
