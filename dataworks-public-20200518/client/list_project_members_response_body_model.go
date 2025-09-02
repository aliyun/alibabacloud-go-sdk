// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListProjectMembersResponseBodyData) *ListProjectMembersResponseBody
	GetData() *ListProjectMembersResponseBodyData
	SetRequestId(v string) *ListProjectMembersResponseBody
	GetRequestId() *string
}

type ListProjectMembersResponseBody struct {
	// The returned results.
	Data *ListProjectMembersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBody) GetData() *ListProjectMembersResponseBodyData {
	return s.Data
}

func (s *ListProjectMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectMembersResponseBody) SetData(v *ListProjectMembersResponseBodyData) *ListProjectMembersResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectMembersResponseBody) SetRequestId(v string) *ListProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectMembersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProjectMembersResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about members in the DataWorks workspace.
	ProjectMemberList []*ListProjectMembersResponseBodyDataProjectMemberList `json:"ProjectMemberList,omitempty" xml:"ProjectMemberList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectMembersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectMembersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectMembersResponseBodyData) GetProjectMemberList() []*ListProjectMembersResponseBodyDataProjectMemberList {
	return s.ProjectMemberList
}

func (s *ListProjectMembersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectMembersResponseBodyData) SetPageNumber(v int32) *ListProjectMembersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListProjectMembersResponseBodyData) SetPageSize(v int32) *ListProjectMembersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersResponseBodyData) SetProjectMemberList(v []*ListProjectMembersResponseBodyDataProjectMemberList) *ListProjectMembersResponseBodyData {
	s.ProjectMemberList = v
	return s
}

func (s *ListProjectMembersResponseBodyData) SetTotalCount(v int32) *ListProjectMembersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListProjectMembersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListProjectMembersResponseBodyDataProjectMemberList struct {
	// The nickname of the member.
	//
	// example:
	//
	// zhangsan
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// The member ID.
	//
	// example:
	//
	// 121
	ProjectMemberId *string `json:"ProjectMemberId,omitempty" xml:"ProjectMemberId,omitempty"`
	// The name of the member.
	//
	// example:
	//
	// zhangsan
	ProjectMemberName *string `json:"ProjectMemberName,omitempty" xml:"ProjectMemberName,omitempty"`
	// The type of the member. Valid values:
	//
	// 	- 1: USER_ALIYUN, which indicates that the member is an Alibaba Cloud account.
	//
	// 	- 5: USER_UBACCOUNT, which indicates that the member is a RAM user.
	//
	// 	- 6: USER_STS_ROLE, which indicates that the member is a RAM role.
	//
	// example:
	//
	// 1
	ProjectMemberType *string `json:"ProjectMemberType,omitempty" xml:"ProjectMemberType,omitempty"`
	// The roles that are assigned to the member.
	ProjectRoleList []*ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList `json:"ProjectRoleList,omitempty" xml:"ProjectRoleList,omitempty" type:"Repeated"`
	// The status of the member. Valid values:
	//
	// 	- 0: NORMAL, which indicates that the member is in a normal state.
	//
	// 	- 1: FORBIDDEN, which indicates that the member is disabled.
	//
	// 	- 2: DELETED, which indicates that the member is deleted.
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListProjectMembersResponseBodyDataProjectMemberList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyDataProjectMemberList) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) GetNick() *string {
	return s.Nick
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) GetProjectMemberId() *string {
	return s.ProjectMemberId
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) GetProjectMemberName() *string {
	return s.ProjectMemberName
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) GetProjectMemberType() *string {
	return s.ProjectMemberType
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) GetProjectRoleList() []*ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList {
	return s.ProjectRoleList
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) GetStatus() *string {
	return s.Status
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) SetNick(v string) *ListProjectMembersResponseBodyDataProjectMemberList {
	s.Nick = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) SetProjectMemberId(v string) *ListProjectMembersResponseBodyDataProjectMemberList {
	s.ProjectMemberId = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) SetProjectMemberName(v string) *ListProjectMembersResponseBodyDataProjectMemberList {
	s.ProjectMemberName = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) SetProjectMemberType(v string) *ListProjectMembersResponseBodyDataProjectMemberList {
	s.ProjectMemberType = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) SetProjectRoleList(v []*ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) *ListProjectMembersResponseBodyDataProjectMemberList {
	s.ProjectRoleList = v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) SetStatus(v string) *ListProjectMembersResponseBodyDataProjectMemberList {
	s.Status = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberList) Validate() error {
	return dara.Validate(s)
}

type ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList struct {
	// The code of the role. DataWorks provides built-in roles and allows you to create custom roles based on your business requirements. For more information about roles, see [Overview of users, roles, and permissions](https://help.aliyun.com/document_detail/295463.html).
	//
	// example:
	//
	// role_project_guest
	ProjectRoleCode *string `json:"ProjectRoleCode,omitempty" xml:"ProjectRoleCode,omitempty"`
	// The role ID.
	//
	// example:
	//
	// 1
	ProjectRoleId *int32 `json:"ProjectRoleId,omitempty" xml:"ProjectRoleId,omitempty"`
	// The name of the role. DataWorks provides built-in roles and allows you to create custom roles based on your business requirements. For more information about roles, see [Overview of users, roles, and permissions](https://help.aliyun.com/document_detail/295463.html).
	//
	// example:
	//
	// test
	ProjectRoleName *string `json:"ProjectRoleName,omitempty" xml:"ProjectRoleName,omitempty"`
	// The type of the role. Valid values:
	//
	// 	- 0: SYSTEM, which indicates that the role is a built-in role.
	//
	// 	- 2: USER_CUSTOM, which indicates that the role is a custom role.
	//
	// example:
	//
	// 0
	ProjectRoleType *string `json:"ProjectRoleType,omitempty" xml:"ProjectRoleType,omitempty"`
}

func (s ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) GetProjectRoleCode() *string {
	return s.ProjectRoleCode
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) GetProjectRoleId() *int32 {
	return s.ProjectRoleId
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) GetProjectRoleName() *string {
	return s.ProjectRoleName
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) GetProjectRoleType() *string {
	return s.ProjectRoleType
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) SetProjectRoleCode(v string) *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList {
	s.ProjectRoleCode = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) SetProjectRoleId(v int32) *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList {
	s.ProjectRoleId = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) SetProjectRoleName(v string) *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList {
	s.ProjectRoleName = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) SetProjectRoleType(v string) *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList {
	s.ProjectRoleType = &v
	return s
}

func (s *ListProjectMembersResponseBodyDataProjectMemberListProjectRoleList) Validate() error {
	return dara.Validate(s)
}
