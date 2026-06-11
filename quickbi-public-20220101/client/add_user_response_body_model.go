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
	// The request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the newly added Alibaba Cloud user.
	Result *AddUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true
	//
	// - false
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
	// The Alibaba Cloud account.
	//
	// example:
	//
	// xxxxxx@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the organization administrator role is assigned. Valid values:
	//
	// - true
	//
	// - false
	//
	// 	Notice:
	//
	// This parameter is deprecated. It does not take effect when `RoleIdList` is specified.
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Indicates whether the permission administrator role is assigned. Valid values:
	//
	// - true
	//
	// - false
	//
	// 	Notice:
	//
	// This parameter is deprecated. It does not take effect when `RoleIdList` is specified.
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// The Copilot modules for which the user has a quota.
	CopilotModules []*string `json:"CopilotModules,omitempty" xml:"CopilotModules,omitempty" type:"Repeated"`
	// The nickname of the Alibaba Cloud account.
	//
	// example:
	//
	// 张三
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// A list of organization role IDs assigned to the user.
	RoleIdList []*int64 `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Repeated"`
	// The user ID in Quick BI.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The user type. Valid values:
	//
	// - 1: developer
	//
	// - 2: viewer
	//
	// - 3: analyst
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

func (s *AddUserResponseBodyResult) GetCopilotModules() []*string {
	return s.CopilotModules
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

func (s *AddUserResponseBodyResult) SetCopilotModules(v []*string) *AddUserResponseBodyResult {
	s.CopilotModules = v
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
