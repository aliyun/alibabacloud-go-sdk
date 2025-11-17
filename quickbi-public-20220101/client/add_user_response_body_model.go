// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserResponseBody
	GetRequestId() *string
	SetResult(v *AddUserResponseBodyResult) *AddUserResponseBody
	GetResult() *AddUserResponseBodyResult
	SetSuccess(v bool) *AddUserResponseBody
	GetSuccess() *bool
}

type AddUserResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns detailed information about the newly added Aliyun user.
	Result *AddUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserResponseBody) GetResult() *AddUserResponseBodyResult {
	return s.Result
}

func (s *AddUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddUserResponseBody) SetRequestId(v string) *AddUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserResponseBody) SetResult(v *AddUserResponseBodyResult) *AddUserResponseBody {
	s.Result = v
	return s
}

func (s *AddUserResponseBody) SetSuccess(v bool) *AddUserResponseBody {
	s.Success = &v
	return s
}

func (s *AddUserResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddUserResponseBodyResult struct {
	// Aliyun account.
	//
	// example:
	//
	// xxxxxx@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Whether the organization administrator role is assigned. Value range:
	//
	// - true: Yes
	//
	// - false: No
	//
	// <notice>This parameter is deprecated and not recommended for use. It is invalid when RoleIdList is provided.</notice>
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Whether the permission administrator role is assigned. Value range:
	//
	// - true: Yes
	//
	// - false: No
	//
	// <notice>This parameter is deprecated and not recommended for use. It is invalid when RoleIdList is provided.</notice>
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// Aliyun account nickname.
	//
	// example:
	//
	// ddd
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// List of organization role IDs bound to the user.
	RoleIdList []*int64 `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Repeated"`
	// UserID in Quick BI.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// User type of the organization member. Value range:
	//
	// - 1: Developer
	//
	// - 2: Visitor
	//
	// - 3: Analyst
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s AddUserResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddUserResponseBodyResult) GetAccountName() *string {
	return s.AccountName
}

func (s *AddUserResponseBodyResult) GetAdminUser() *bool {
	return s.AdminUser
}

func (s *AddUserResponseBodyResult) GetAuthAdminUser() *bool {
	return s.AuthAdminUser
}

func (s *AddUserResponseBodyResult) GetNickName() *string {
	return s.NickName
}

func (s *AddUserResponseBodyResult) GetRoleIdList() []*int64 {
	return s.RoleIdList
}

func (s *AddUserResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *AddUserResponseBodyResult) GetUserType() *int32 {
	return s.UserType
}

func (s *AddUserResponseBodyResult) SetAccountName(v string) *AddUserResponseBodyResult {
	s.AccountName = &v
	return s
}

func (s *AddUserResponseBodyResult) SetAdminUser(v bool) *AddUserResponseBodyResult {
	s.AdminUser = &v
	return s
}

func (s *AddUserResponseBodyResult) SetAuthAdminUser(v bool) *AddUserResponseBodyResult {
	s.AuthAdminUser = &v
	return s
}

func (s *AddUserResponseBodyResult) SetNickName(v string) *AddUserResponseBodyResult {
	s.NickName = &v
	return s
}

func (s *AddUserResponseBodyResult) SetRoleIdList(v []*int64) *AddUserResponseBodyResult {
	s.RoleIdList = v
	return s
}

func (s *AddUserResponseBodyResult) SetUserId(v string) *AddUserResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *AddUserResponseBodyResult) SetUserType(v int32) *AddUserResponseBodyResult {
	s.UserType = &v
	return s
}

func (s *AddUserResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
