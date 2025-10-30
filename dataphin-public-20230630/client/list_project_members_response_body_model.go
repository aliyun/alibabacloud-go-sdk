// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProjectMembersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListProjectMembersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListProjectMembersResponseBody
	GetMessage() *string
	SetPageResult(v *ListProjectMembersResponseBodyPageResult) *ListProjectMembersResponseBody
	GetPageResult() *ListProjectMembersResponseBodyPageResult
	SetRequestId(v string) *ListProjectMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectMembersResponseBody
	GetSuccess() *bool
}

type ListProjectMembersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListProjectMembersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProjectMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProjectMembersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListProjectMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProjectMembersResponseBody) GetPageResult() *ListProjectMembersResponseBodyPageResult {
	return s.PageResult
}

func (s *ListProjectMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectMembersResponseBody) SetCode(v string) *ListProjectMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetHttpStatusCode(v int32) *ListProjectMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetMessage(v string) *ListProjectMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetPageResult(v *ListProjectMembersResponseBodyPageResult) *ListProjectMembersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListProjectMembersResponseBody) SetRequestId(v string) *ListProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetSuccess(v bool) *ListProjectMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectMembersResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectMembersResponseBodyPageResult struct {
	ProjectMemberList []*ListProjectMembersResponseBodyPageResultProjectMemberList `json:"ProjectMemberList,omitempty" xml:"ProjectMemberList,omitempty" type:"Repeated"`
	// example:
	//
	// 101
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectMembersResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyPageResult) GetProjectMemberList() []*ListProjectMembersResponseBodyPageResultProjectMemberList {
	return s.ProjectMemberList
}

func (s *ListProjectMembersResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectMembersResponseBodyPageResult) SetProjectMemberList(v []*ListProjectMembersResponseBodyPageResultProjectMemberList) *ListProjectMembersResponseBodyPageResult {
	s.ProjectMemberList = v
	return s
}

func (s *ListProjectMembersResponseBodyPageResult) SetTotalCount(v int32) *ListProjectMembersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListProjectMembersResponseBodyPageResult) Validate() error {
	if s.ProjectMemberList != nil {
		for _, item := range s.ProjectMemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectMembersResponseBodyPageResultProjectMemberList struct {
	// example:
	//
	// 1702692675000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1721720955000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 12356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 101111
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// test
	LastModifierName *string  `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	RoleIdList       []*int32 `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 101111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 张三
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListProjectMembersResponseBodyPageResultProjectMemberList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyPageResultProjectMemberList) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) GetId() *int64 {
	return s.Id
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) GetLastModifier() *string {
	return s.LastModifier
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) GetRoleIdList() []*int32 {
	return s.RoleIdList
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) GetUserId() *string {
	return s.UserId
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) GetUserName() *string {
	return s.UserName
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) SetGmtCreate(v string) *ListProjectMembersResponseBodyPageResultProjectMemberList {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) SetGmtModified(v string) *ListProjectMembersResponseBodyPageResultProjectMemberList {
	s.GmtModified = &v
	return s
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) SetId(v int64) *ListProjectMembersResponseBodyPageResultProjectMemberList {
	s.Id = &v
	return s
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) SetLastModifier(v string) *ListProjectMembersResponseBodyPageResultProjectMemberList {
	s.LastModifier = &v
	return s
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) SetLastModifierName(v string) *ListProjectMembersResponseBodyPageResultProjectMemberList {
	s.LastModifierName = &v
	return s
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) SetRoleIdList(v []*int32) *ListProjectMembersResponseBodyPageResultProjectMemberList {
	s.RoleIdList = v
	return s
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) SetUserId(v string) *ListProjectMembersResponseBodyPageResultProjectMemberList {
	s.UserId = &v
	return s
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) SetUserName(v string) *ListProjectMembersResponseBodyPageResultProjectMemberList {
	s.UserName = &v
	return s
}

func (s *ListProjectMembersResponseBodyPageResultProjectMemberList) Validate() error {
	return dara.Validate(s)
}
