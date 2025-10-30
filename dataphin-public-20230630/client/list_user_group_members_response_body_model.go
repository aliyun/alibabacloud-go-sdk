// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserGroupMembersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListUserGroupMembersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUserGroupMembersResponseBody
	GetMessage() *string
	SetPageResult(v *ListUserGroupMembersResponseBodyPageResult) *ListUserGroupMembersResponseBody
	GetPageResult() *ListUserGroupMembersResponseBodyPageResult
	SetRequestId(v string) *ListUserGroupMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserGroupMembersResponseBody
	GetSuccess() *bool
}

type ListUserGroupMembersResponseBody struct {
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
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListUserGroupMembersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserGroupMembersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUserGroupMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserGroupMembersResponseBody) GetPageResult() *ListUserGroupMembersResponseBodyPageResult {
	return s.PageResult
}

func (s *ListUserGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserGroupMembersResponseBody) SetCode(v string) *ListUserGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetHttpStatusCode(v int32) *ListUserGroupMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetMessage(v string) *ListUserGroupMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetPageResult(v *ListUserGroupMembersResponseBodyPageResult) *ListUserGroupMembersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetRequestId(v string) *ListUserGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetSuccess(v bool) *ListUserGroupMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserGroupMembersResponseBodyPageResult struct {
	MemberList []*ListUserGroupMembersResponseBodyPageResultMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// example:
	//
	// 217
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserGroupMembersResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBodyPageResult) GetMemberList() []*ListUserGroupMembersResponseBodyPageResultMemberList {
	return s.MemberList
}

func (s *ListUserGroupMembersResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserGroupMembersResponseBodyPageResult) SetMemberList(v []*ListUserGroupMembersResponseBodyPageResultMemberList) *ListUserGroupMembersResponseBodyPageResult {
	s.MemberList = v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResult) SetTotalCount(v int32) *ListUserGroupMembersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResult) Validate() error {
	if s.MemberList != nil {
		for _, item := range s.MemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserGroupMembersResponseBodyPageResultMemberList struct {
	Creator *ListUserGroupMembersResponseBodyPageResultMemberListCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// zhangsan
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2324211
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 231111
	UserGroupId *string                                                       `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserInfo    *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	// example:
	//
	// SECURITY_ADMIN
	UserRole *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
}

func (s ListUserGroupMembersResponseBodyPageResultMemberList) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersResponseBodyPageResultMemberList) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) GetCreator() *ListUserGroupMembersResponseBodyPageResultMemberListCreator {
	return s.Creator
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) GetId() *string {
	return s.Id
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) GetUserInfo() *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo {
	return s.UserInfo
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) GetUserRole() *string {
	return s.UserRole
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetCreator(v *ListUserGroupMembersResponseBodyPageResultMemberListCreator) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.Creator = v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetGmtCreate(v int64) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.GmtCreate = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetId(v string) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.Id = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetUserGroupId(v string) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetUserInfo(v *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.UserInfo = v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetUserRole(v string) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.UserRole = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserGroupMembersResponseBodyPageResultMemberListCreator struct {
	// example:
	//
	// 12121111
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 12121111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserGroupMembersResponseBodyPageResultMemberListCreator) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersResponseBodyPageResultMemberListCreator) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) GetAccountName() *string {
	return s.AccountName
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) GetId() *string {
	return s.Id
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) SetAccountName(v string) *ListUserGroupMembersResponseBodyPageResultMemberListCreator {
	s.AccountName = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) SetDisplayName(v string) *ListUserGroupMembersResponseBodyPageResultMemberListCreator {
	s.DisplayName = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) SetId(v string) *ListUserGroupMembersResponseBodyPageResultMemberListCreator {
	s.Id = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) Validate() error {
	return dara.Validate(s)
}

type ListUserGroupMembersResponseBodyPageResultMemberListUserInfo struct {
	// example:
	//
	// atest
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 13232
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) GetAccountName() *string {
	return s.AccountName
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) GetId() *string {
	return s.Id
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) SetAccountName(v string) *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo {
	s.AccountName = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) SetDisplayName(v string) *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo {
	s.DisplayName = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) SetId(v string) *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo {
	s.Id = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) Validate() error {
	return dara.Validate(s)
}
