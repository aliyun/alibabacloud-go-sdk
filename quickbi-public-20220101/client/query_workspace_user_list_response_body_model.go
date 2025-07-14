// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkspaceUserListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryWorkspaceUserListResponseBody
	GetRequestId() *string
	SetResult(v *QueryWorkspaceUserListResponseBodyResult) *QueryWorkspaceUserListResponseBody
	GetResult() *QueryWorkspaceUserListResponseBodyResult
	SetSuccess(v bool) *QueryWorkspaceUserListResponseBody
	GetSuccess() *bool
}

type QueryWorkspaceUserListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the paginated result of the member list, with detailed information about the members stored in the Data parameter of the response.
	Result *QueryWorkspaceUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s QueryWorkspaceUserListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWorkspaceUserListResponseBody) GetResult() *QueryWorkspaceUserListResponseBodyResult {
	return s.Result
}

func (s *QueryWorkspaceUserListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWorkspaceUserListResponseBody) SetRequestId(v string) *QueryWorkspaceUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBody) SetResult(v *QueryWorkspaceUserListResponseBodyResult) *QueryWorkspaceUserListResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorkspaceUserListResponseBody) SetSuccess(v bool) *QueryWorkspaceUserListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryWorkspaceUserListResponseBodyResult struct {
	// Information about the workspace members.
	Data []*QueryWorkspaceUserListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of rows per page as set in the request.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of rows.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryWorkspaceUserListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceUserListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponseBodyResult) GetData() []*QueryWorkspaceUserListResponseBodyResultData {
	return s.Data
}

func (s *QueryWorkspaceUserListResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryWorkspaceUserListResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryWorkspaceUserListResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *QueryWorkspaceUserListResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetData(v []*QueryWorkspaceUserListResponseBodyResultData) *QueryWorkspaceUserListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetPageNum(v int32) *QueryWorkspaceUserListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetPageSize(v int32) *QueryWorkspaceUserListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetTotalNum(v int32) *QueryWorkspaceUserListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetTotalPages(v int32) *QueryWorkspaceUserListResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type QueryWorkspaceUserListResponseBodyResultData struct {
	// Alibaba Cloud account ID.
	//
	// example:
	//
	// 16020915****8429
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Alibaba Cloud account name.
	//
	// example:
	//
	// pop****@aliyunid.test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Nickname.
	//
	// example:
	//
	// test
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// Preset role information for the workspace member.
	Role *QueryWorkspaceUserListResponseBodyResultDataRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
	// Quick BI user ID.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryWorkspaceUserListResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceUserListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponseBodyResultData) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryWorkspaceUserListResponseBodyResultData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryWorkspaceUserListResponseBodyResultData) GetNickName() *string {
	return s.NickName
}

func (s *QueryWorkspaceUserListResponseBodyResultData) GetRole() *QueryWorkspaceUserListResponseBodyResultDataRole {
	return s.Role
}

func (s *QueryWorkspaceUserListResponseBodyResultData) GetUserId() *string {
	return s.UserId
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetAccountId(v string) *QueryWorkspaceUserListResponseBodyResultData {
	s.AccountId = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetAccountName(v string) *QueryWorkspaceUserListResponseBodyResultData {
	s.AccountName = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetNickName(v string) *QueryWorkspaceUserListResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetRole(v *QueryWorkspaceUserListResponseBodyResultDataRole) *QueryWorkspaceUserListResponseBodyResultData {
	s.Role = v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetUserId(v string) *QueryWorkspaceUserListResponseBodyResultData {
	s.UserId = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

type QueryWorkspaceUserListResponseBodyResultDataRole struct {
	// Code corresponding to the preset role.
	//
	// example:
	//
	// role_workspace_admin
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// Preset role ID. Possible values:
	//
	// - 25: Workspace Administrator
	//
	// - 26: Workspace Developer
	//
	// - 27: Workspace Analyst
	//
	// - 30: Workspace Viewer
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// Name of the preset role.
	//
	// example:
	//
	// test
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s QueryWorkspaceUserListResponseBodyResultDataRole) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceUserListResponseBodyResultDataRole) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) GetRoleCode() *string {
	return s.RoleCode
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) GetRoleId() *int64 {
	return s.RoleId
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) GetRoleName() *string {
	return s.RoleName
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) SetRoleCode(v string) *QueryWorkspaceUserListResponseBodyResultDataRole {
	s.RoleCode = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) SetRoleId(v int64) *QueryWorkspaceUserListResponseBodyResultDataRole {
	s.RoleId = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) SetRoleName(v string) *QueryWorkspaceUserListResponseBodyResultDataRole {
	s.RoleName = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) Validate() error {
	return dara.Validate(s)
}
