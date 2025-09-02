// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectRoleList(v []*ListProjectRolesResponseBodyProjectRoleList) *ListProjectRolesResponseBody
	GetProjectRoleList() []*ListProjectRolesResponseBodyProjectRoleList
	SetRequestId(v string) *ListProjectRolesResponseBody
	GetRequestId() *string
}

type ListProjectRolesResponseBody struct {
	// The roles in the DataWorks workspace.
	ProjectRoleList []*ListProjectRolesResponseBodyProjectRoleList `json:"ProjectRoleList,omitempty" xml:"ProjectRoleList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBody) GetProjectRoleList() []*ListProjectRolesResponseBodyProjectRoleList {
	return s.ProjectRoleList
}

func (s *ListProjectRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectRolesResponseBody) SetProjectRoleList(v []*ListProjectRolesResponseBodyProjectRoleList) *ListProjectRolesResponseBody {
	s.ProjectRoleList = v
	return s
}

func (s *ListProjectRolesResponseBody) SetRequestId(v string) *ListProjectRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectRolesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProjectRolesResponseBodyProjectRoleList struct {
	// The code of the role in the DataWorks workspace.
	//
	// example:
	//
	// role_project_guest
	ProjectRoleCode *string `json:"ProjectRoleCode,omitempty" xml:"ProjectRoleCode,omitempty"`
	// The ID of the role in the DataWorks workspace.
	//
	// example:
	//
	// 1
	ProjectRoleId *int32 `json:"ProjectRoleId,omitempty" xml:"ProjectRoleId,omitempty"`
	// The name of the role in the DataWorks workspace.
	//
	// example:
	//
	// visitor
	ProjectRoleName *string `json:"ProjectRoleName,omitempty" xml:"ProjectRoleName,omitempty"`
	// The type of the role in the DataWorks workspace.
	//
	// example:
	//
	// 0
	ProjectRoleType *string `json:"ProjectRoleType,omitempty" xml:"ProjectRoleType,omitempty"`
}

func (s ListProjectRolesResponseBodyProjectRoleList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesResponseBodyProjectRoleList) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBodyProjectRoleList) GetProjectRoleCode() *string {
	return s.ProjectRoleCode
}

func (s *ListProjectRolesResponseBodyProjectRoleList) GetProjectRoleId() *int32 {
	return s.ProjectRoleId
}

func (s *ListProjectRolesResponseBodyProjectRoleList) GetProjectRoleName() *string {
	return s.ProjectRoleName
}

func (s *ListProjectRolesResponseBodyProjectRoleList) GetProjectRoleType() *string {
	return s.ProjectRoleType
}

func (s *ListProjectRolesResponseBodyProjectRoleList) SetProjectRoleCode(v string) *ListProjectRolesResponseBodyProjectRoleList {
	s.ProjectRoleCode = &v
	return s
}

func (s *ListProjectRolesResponseBodyProjectRoleList) SetProjectRoleId(v int32) *ListProjectRolesResponseBodyProjectRoleList {
	s.ProjectRoleId = &v
	return s
}

func (s *ListProjectRolesResponseBodyProjectRoleList) SetProjectRoleName(v string) *ListProjectRolesResponseBodyProjectRoleList {
	s.ProjectRoleName = &v
	return s
}

func (s *ListProjectRolesResponseBodyProjectRoleList) SetProjectRoleType(v string) *ListProjectRolesResponseBodyProjectRoleList {
	s.ProjectRoleType = &v
	return s
}

func (s *ListProjectRolesResponseBodyProjectRoleList) Validate() error {
	return dara.Validate(s)
}
