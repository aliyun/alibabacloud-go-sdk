// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListProjectRolesResponseBodyPagingInfo) *ListProjectRolesResponseBody
	GetPagingInfo() *ListProjectRolesResponseBodyPagingInfo
	SetRequestId(v string) *ListProjectRolesResponseBody
	GetRequestId() *string
}

type ListProjectRolesResponseBody struct {
	// The pagination information.
	PagingInfo *ListProjectRolesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 61649187-0BCF-5E75-8D4B-64FDBEBBB447
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBody) GetPagingInfo() *ListProjectRolesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListProjectRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectRolesResponseBody) SetPagingInfo(v *ListProjectRolesResponseBodyPagingInfo) *ListProjectRolesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListProjectRolesResponseBody) SetRequestId(v string) *ListProjectRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectRolesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProjectRolesResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The roles in the DataWorks workspace.
	ProjectRoles []*ListProjectRolesResponseBodyPagingInfoProjectRoles `json:"ProjectRoles,omitempty" xml:"ProjectRoles,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 42
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectRolesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBodyPagingInfo) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListProjectRolesResponseBodyPagingInfo) GetPageSize() *string {
	return s.PageSize
}

func (s *ListProjectRolesResponseBodyPagingInfo) GetProjectRoles() []*ListProjectRolesResponseBodyPagingInfoProjectRoles {
	return s.ProjectRoles
}

func (s *ListProjectRolesResponseBodyPagingInfo) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListProjectRolesResponseBodyPagingInfo) SetPageNumber(v string) *ListProjectRolesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfo) SetPageSize(v string) *ListProjectRolesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfo) SetProjectRoles(v []*ListProjectRolesResponseBodyPagingInfoProjectRoles) *ListProjectRolesResponseBodyPagingInfo {
	s.ProjectRoles = v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfo) SetTotalCount(v string) *ListProjectRolesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListProjectRolesResponseBodyPagingInfoProjectRoles struct {
	// The code of the role in the DataWorks workspace.
	//
	// example:
	//
	// role_project_guest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// Visitors
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 21229
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the role in the DataWorks workspace.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectRolesResponseBodyPagingInfoProjectRoles) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesResponseBodyPagingInfoProjectRoles) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) GetCode() *string {
	return s.Code
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) GetName() *string {
	return s.Name
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) GetType() *string {
	return s.Type
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) SetCode(v string) *ListProjectRolesResponseBodyPagingInfoProjectRoles {
	s.Code = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) SetName(v string) *ListProjectRolesResponseBodyPagingInfoProjectRoles {
	s.Name = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) SetProjectId(v int64) *ListProjectRolesResponseBodyPagingInfoProjectRoles {
	s.ProjectId = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) SetType(v string) *ListProjectRolesResponseBodyPagingInfoProjectRoles {
	s.Type = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) Validate() error {
	return dara.Validate(s)
}
