// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoByAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserInfoByAccountResponseBody
	GetRequestId() *string
	SetResult(v *QueryUserInfoByAccountResponseBodyResult) *QueryUserInfoByAccountResponseBody
	GetResult() *QueryUserInfoByAccountResponseBodyResult
	SetSuccess(v bool) *QueryUserInfoByAccountResponseBody
	GetSuccess() *bool
}

type QueryUserInfoByAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned organization user information.
	Result *QueryUserInfoByAccountResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserInfoByAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoByAccountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserInfoByAccountResponseBody) GetResult() *QueryUserInfoByAccountResponseBodyResult {
	return s.Result
}

func (s *QueryUserInfoByAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserInfoByAccountResponseBody) SetRequestId(v string) *QueryUserInfoByAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBody) SetResult(v *QueryUserInfoByAccountResponseBodyResult) *QueryUserInfoByAccountResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserInfoByAccountResponseBody) SetSuccess(v bool) *QueryUserInfoByAccountResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUserInfoByAccountResponseBodyResult struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 135****5848
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account that corresponds to the member. (If you use a RAM user, the domain name information that follows @ is removed. For example, if you use a <test@test.com>, test is returned.)
	//
	// example:
	//
	// 1386587****@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Whether you are an administrator of the organization. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Whether you are a permission administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// 1386587****@163.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The nickname of the account.
	//
	// example:
	//
	// Test user
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The phone number of the alert contact.
	//
	// example:
	//
	// 1386587****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// 用户绑定的组织角色ID列表。
	RoleIdList []*int64 `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Repeated"`
	// The UserID in the Quick BI.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The role type of the organization member. Valid values:
	//
	// 	- 1 : developer
	//
	// 	- 2 : visitors
	//
	// 	- 3 : Analyst
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s QueryUserInfoByAccountResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoByAccountResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetAdminUser() *bool {
	return s.AdminUser
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetAuthAdminUser() *bool {
	return s.AuthAdminUser
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetNickName() *string {
	return s.NickName
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetPhone() *string {
	return s.Phone
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetRoleIdList() []*int64 {
	return s.RoleIdList
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserInfoByAccountResponseBodyResult) GetUserType() *int32 {
	return s.UserType
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetAccountId(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.AccountId = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetAccountName(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.AccountName = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetAdminUser(v bool) *QueryUserInfoByAccountResponseBodyResult {
	s.AdminUser = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetAuthAdminUser(v bool) *QueryUserInfoByAccountResponseBodyResult {
	s.AuthAdminUser = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetEmail(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.Email = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetNickName(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.NickName = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetPhone(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.Phone = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetRoleIdList(v []*int64) *QueryUserInfoByAccountResponseBodyResult {
	s.RoleIdList = v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetUserId(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetUserType(v int32) *QueryUserInfoByAccountResponseBodyResult {
	s.UserType = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
