// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListProjectMembersResponseBodyPagingInfo) *ListProjectMembersResponseBody
	GetPagingInfo() *ListProjectMembersResponseBodyPagingInfo
	SetRequestId(v string) *ListProjectMembersResponseBody
	GetRequestId() *string
}

type ListProjectMembersResponseBody struct {
	// The pagination information.
	PagingInfo *ListProjectMembersResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 9FBBBB1F-DD5E-5D8E-8F50-37F77460F056
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBody) GetPagingInfo() *ListProjectMembersResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListProjectMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectMembersResponseBody) SetPagingInfo(v *ListProjectMembersResponseBodyPagingInfo) *ListProjectMembersResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListProjectMembersResponseBody) SetRequestId(v string) *ListProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectMembersResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectMembersResponseBodyPagingInfo struct {
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
	// The members in the workspace.
	ProjectMembers []*ListProjectMembersResponseBodyPagingInfoProjectMembers `json:"ProjectMembers,omitempty" xml:"ProjectMembers,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectMembersResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectMembersResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectMembersResponseBodyPagingInfo) GetProjectMembers() []*ListProjectMembersResponseBodyPagingInfoProjectMembers {
	return s.ProjectMembers
}

func (s *ListProjectMembersResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectMembersResponseBodyPagingInfo) SetPageNumber(v int32) *ListProjectMembersResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfo) SetPageSize(v int32) *ListProjectMembersResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfo) SetProjectMembers(v []*ListProjectMembersResponseBodyPagingInfoProjectMembers) *ListProjectMembersResponseBodyPagingInfo {
	s.ProjectMembers = v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfo) SetTotalCount(v int32) *ListProjectMembersResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfo) Validate() error {
	if s.ProjectMembers != nil {
		for _, item := range s.ProjectMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectMembersResponseBodyPagingInfoProjectMembers struct {
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 62136
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The roles that are assigned to the member.
	Roles []*ListProjectMembersResponseBodyPagingInfoProjectMembersRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The status of the member. Valid values:
	//
	// 	- Normal
	//
	// 	- Forbidden
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the account used by the member.
	//
	// example:
	//
	// 123422344899
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListProjectMembersResponseBodyPagingInfoProjectMembers) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyPagingInfoProjectMembers) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) GetRoles() []*ListProjectMembersResponseBodyPagingInfoProjectMembersRoles {
	return s.Roles
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) GetStatus() *string {
	return s.Status
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) GetUserId() *string {
	return s.UserId
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) GetUserName() *string {
	return s.UserName
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetProjectId(v int64) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.ProjectId = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetRoles(v []*ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.Roles = v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetStatus(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.Status = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetUserId(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.UserId = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetUserName(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.UserName = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectMembersResponseBodyPagingInfoProjectMembersRoles struct {
	// The code of the role.
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
	// The type of the role. Valid values:
	//
	// 	- UserCustom: user-defined role
	//
	// 	- System: system role
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) GetCode() *string {
	return s.Code
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) GetName() *string {
	return s.Name
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) GetType() *string {
	return s.Type
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) SetCode(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles {
	s.Code = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) SetName(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles {
	s.Name = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) SetType(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles {
	s.Type = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) Validate() error {
	return dara.Validate(s)
}
