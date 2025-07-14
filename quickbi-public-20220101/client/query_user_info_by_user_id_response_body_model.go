// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserInfoByUserIdResponseBody
	GetRequestId() *string
	SetResult(v *QueryUserInfoByUserIdResponseBodyResult) *QueryUserInfoByUserIdResponseBody
	GetResult() *QueryUserInfoByUserIdResponseBodyResult
	SetSuccess(v bool) *QueryUserInfoByUserIdResponseBody
	GetSuccess() *bool
}

type QueryUserInfoByUserIdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned organization user information.
	Result *QueryUserInfoByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s QueryUserInfoByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserInfoByUserIdResponseBody) GetResult() *QueryUserInfoByUserIdResponseBodyResult {
	return s.Result
}

func (s *QueryUserInfoByUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserInfoByUserIdResponseBody) SetRequestId(v string) *QueryUserInfoByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBody) SetResult(v *QueryUserInfoByUserIdResponseBodyResult) *QueryUserInfoByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserInfoByUserIdResponseBody) SetSuccess(v bool) *QueryUserInfoByUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUserInfoByUserIdResponseBodyResult struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 135****5848
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account that corresponds to the member.
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
	Phone      *string  `json:"Phone,omitempty" xml:"Phone,omitempty"`
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

func (s QueryUserInfoByUserIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetAdminUser() *bool {
	return s.AdminUser
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetAuthAdminUser() *bool {
	return s.AuthAdminUser
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetNickName() *string {
	return s.NickName
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetPhone() *string {
	return s.Phone
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetRoleIdList() []*int64 {
	return s.RoleIdList
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserInfoByUserIdResponseBodyResult) GetUserType() *int32 {
	return s.UserType
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetAccountId(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.AccountId = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetAccountName(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.AccountName = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetAdminUser(v bool) *QueryUserInfoByUserIdResponseBodyResult {
	s.AdminUser = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetAuthAdminUser(v bool) *QueryUserInfoByUserIdResponseBodyResult {
	s.AuthAdminUser = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetEmail(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.Email = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetNickName(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.NickName = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetPhone(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.Phone = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetRoleIdList(v []*int64) *QueryUserInfoByUserIdResponseBodyResult {
	s.RoleIdList = v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetUserId(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetUserType(v int32) *QueryUserInfoByUserIdResponseBodyResult {
	s.UserType = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
