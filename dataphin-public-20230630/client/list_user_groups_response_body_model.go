// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserGroupsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListUserGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUserGroupsResponseBody
	GetMessage() *string
	SetPageResult(v *ListUserGroupsResponseBodyPageResult) *ListUserGroupsResponseBody
	GetPageResult() *ListUserGroupsResponseBodyPageResult
	SetRequestId(v string) *ListUserGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserGroupsResponseBody
	GetSuccess() *bool
}

type ListUserGroupsResponseBody struct {
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
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListUserGroupsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUserGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserGroupsResponseBody) GetPageResult() *ListUserGroupsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListUserGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserGroupsResponseBody) SetCode(v string) *ListUserGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetHttpStatusCode(v int32) *ListUserGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetMessage(v string) *ListUserGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetPageResult(v *ListUserGroupsResponseBodyPageResult) *ListUserGroupsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetSuccess(v bool) *ListUserGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserGroupsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserGroupsResponseBodyPageResult struct {
	// example:
	//
	// 49
	TotalCount    *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserGroupList []*ListUserGroupsResponseBodyPageResultUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
}

func (s ListUserGroupsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserGroupsResponseBodyPageResult) GetUserGroupList() []*ListUserGroupsResponseBodyPageResultUserGroupList {
	return s.UserGroupList
}

func (s *ListUserGroupsResponseBodyPageResult) SetTotalCount(v int32) *ListUserGroupsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResult) SetUserGroupList(v []*ListUserGroupsResponseBodyPageResultUserGroupList) *ListUserGroupsResponseBodyPageResult {
	s.UserGroupList = v
	return s
}

func (s *ListUserGroupsResponseBodyPageResult) Validate() error {
	if s.UserGroupList != nil {
		for _, item := range s.UserGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserGroupsResponseBodyPageResultUserGroupList struct {
	// example:
	//
	// true
	Active      *bool                                                         `json:"Active,omitempty" xml:"Active,omitempty"`
	AdminList   []*ListUserGroupsResponseBodyPageResultUserGroupListAdminList `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	Description *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 31313232
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// SECURITY_ADMIN
	MyRole *string `json:"MyRole,omitempty" xml:"MyRole,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUserGroupsResponseBodyPageResultUserGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBodyPageResultUserGroupList) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) GetActive() *bool {
	return s.Active
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) GetAdminList() []*ListUserGroupsResponseBodyPageResultUserGroupListAdminList {
	return s.AdminList
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) GetDescription() *string {
	return s.Description
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) GetId() *string {
	return s.Id
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) GetMyRole() *string {
	return s.MyRole
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) GetName() *string {
	return s.Name
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetActive(v bool) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.Active = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetAdminList(v []*ListUserGroupsResponseBodyPageResultUserGroupListAdminList) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.AdminList = v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetDescription(v string) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.Description = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetId(v string) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.Id = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetMyRole(v string) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.MyRole = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetName(v string) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.Name = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) Validate() error {
	if s.AdminList != nil {
		for _, item := range s.AdminList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserGroupsResponseBodyPageResultUserGroupListAdminList struct {
	// example:
	//
	// zhangsan
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 32313131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserGroupsResponseBodyPageResultUserGroupListAdminList) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBodyPageResultUserGroupListAdminList) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) GetAccountName() *string {
	return s.AccountName
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) GetId() *string {
	return s.Id
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) SetAccountName(v string) *ListUserGroupsResponseBodyPageResultUserGroupListAdminList {
	s.AccountName = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) SetDisplayName(v string) *ListUserGroupsResponseBodyPageResultUserGroupListAdminList {
	s.DisplayName = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) SetId(v string) *ListUserGroupsResponseBodyPageResultUserGroupListAdminList {
	s.Id = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) Validate() error {
	return dara.Validate(s)
}
