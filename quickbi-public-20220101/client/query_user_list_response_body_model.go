// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserListResponseBody
	GetRequestId() *string
	SetResult(v *QueryUserListResponseBodyResult) *QueryUserListResponseBody
	GetResult() *QueryUserListResponseBodyResult
	SetSuccess(v bool) *QueryUserListResponseBody
	GetSuccess() *bool
}

type QueryUserListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The pagination result of the user list is returned. The detailed information list of organization members is stored in the response parameter Data.
	Result *QueryUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s QueryUserListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserListResponseBody) GetResult() *QueryUserListResponseBodyResult {
	return s.Result
}

func (s *QueryUserListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserListResponseBody) SetRequestId(v string) *QueryUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserListResponseBody) SetResult(v *QueryUserListResponseBodyResult) *QueryUserListResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserListResponseBody) SetSuccess(v bool) *QueryUserListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUserListResponseBodyResult struct {
	// Returns the list of requested users.
	Data []*QueryUserListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page set when the interface is requested.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of rows in the table.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryUserListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserListResponseBodyResult) GetData() []*QueryUserListResponseBodyResultData {
	return s.Data
}

func (s *QueryUserListResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryUserListResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryUserListResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *QueryUserListResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryUserListResponseBodyResult) SetData(v []*QueryUserListResponseBodyResultData) *QueryUserListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryUserListResponseBodyResult) SetPageNum(v int32) *QueryUserListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryUserListResponseBodyResult) SetPageSize(v int32) *QueryUserListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryUserListResponseBodyResult) SetTotalNum(v int32) *QueryUserListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryUserListResponseBodyResult) SetTotalPages(v int32) *QueryUserListResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryUserListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type QueryUserListResponseBodyResultData struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1355********
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account that corresponds to the member.
	//
	// example:
	//
	// Test user
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the organization administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Indicate whether the RAM user is a permission administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// User status:
	//
	// - Active - false
	//
	// - Inactive - true
	//
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// Join Date
	//
	// example:
	//
	// 1718691704000
	JoinedDate *int64 `json:"JoinedDate,omitempty" xml:"JoinedDate,omitempty"`
	// Last login time.
	//
	// example:
	//
	// 1718761320681
	LastLoginTime *int64 `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	// The nickname of the organization member.
	//
	// example:
	//
	// Test user
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// List of organization role IDs bound to the user.
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

func (s QueryUserListResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryUserListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryUserListResponseBodyResultData) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryUserListResponseBodyResultData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryUserListResponseBodyResultData) GetAdminUser() *bool {
	return s.AdminUser
}

func (s *QueryUserListResponseBodyResultData) GetAuthAdminUser() *bool {
	return s.AuthAdminUser
}

func (s *QueryUserListResponseBodyResultData) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *QueryUserListResponseBodyResultData) GetJoinedDate() *int64 {
	return s.JoinedDate
}

func (s *QueryUserListResponseBodyResultData) GetLastLoginTime() *int64 {
	return s.LastLoginTime
}

func (s *QueryUserListResponseBodyResultData) GetNickName() *string {
	return s.NickName
}

func (s *QueryUserListResponseBodyResultData) GetRoleIdList() []*int64 {
	return s.RoleIdList
}

func (s *QueryUserListResponseBodyResultData) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserListResponseBodyResultData) GetUserType() *int32 {
	return s.UserType
}

func (s *QueryUserListResponseBodyResultData) SetAccountId(v string) *QueryUserListResponseBodyResultData {
	s.AccountId = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetAccountName(v string) *QueryUserListResponseBodyResultData {
	s.AccountName = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetAdminUser(v bool) *QueryUserListResponseBodyResultData {
	s.AdminUser = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetAuthAdminUser(v bool) *QueryUserListResponseBodyResultData {
	s.AuthAdminUser = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetIsDeleted(v bool) *QueryUserListResponseBodyResultData {
	s.IsDeleted = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetJoinedDate(v int64) *QueryUserListResponseBodyResultData {
	s.JoinedDate = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetLastLoginTime(v int64) *QueryUserListResponseBodyResultData {
	s.LastLoginTime = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetNickName(v string) *QueryUserListResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetRoleIdList(v []*int64) *QueryUserListResponseBodyResultData {
	s.RoleIdList = v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetUserId(v string) *QueryUserListResponseBodyResultData {
	s.UserId = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetUserType(v int32) *QueryUserListResponseBodyResultData {
	s.UserType = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
