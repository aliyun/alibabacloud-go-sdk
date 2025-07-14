// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceRoleUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkspaceRoleUsersResponseBody
	GetRequestId() *string
	SetResult(v *ListWorkspaceRoleUsersResponseBodyResult) *ListWorkspaceRoleUsersResponseBody
	GetResult() *ListWorkspaceRoleUsersResponseBodyResult
	SetSuccess(v bool) *ListWorkspaceRoleUsersResponseBody
	GetSuccess() *bool
}

type ListWorkspaceRoleUsersResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the list of users under the specified workspace role.
	Result *ListWorkspaceRoleUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s ListWorkspaceRoleUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspaceRoleUsersResponseBody) GetResult() *ListWorkspaceRoleUsersResponseBodyResult {
	return s.Result
}

func (s *ListWorkspaceRoleUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkspaceRoleUsersResponseBody) SetRequestId(v string) *ListWorkspaceRoleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBody) SetResult(v *ListWorkspaceRoleUsersResponseBodyResult) *ListWorkspaceRoleUsersResponseBody {
	s.Result = v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBody) SetSuccess(v bool) *ListWorkspaceRoleUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWorkspaceRoleUsersResponseBodyResult struct {
	// User list.
	Data []*ListWorkspaceRoleUsersResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of items per page as set in the request.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of items.
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

func (s ListWorkspaceRoleUsersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRoleUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) GetData() []*ListWorkspaceRoleUsersResponseBodyResultData {
	return s.Data
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetData(v []*ListWorkspaceRoleUsersResponseBodyResultData) *ListWorkspaceRoleUsersResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetPageNum(v int32) *ListWorkspaceRoleUsersResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetPageSize(v int32) *ListWorkspaceRoleUsersResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetTotalNum(v int32) *ListWorkspaceRoleUsersResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetTotalPages(v int32) *ListWorkspaceRoleUsersResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListWorkspaceRoleUsersResponseBodyResultData struct {
	// Nickname of the organization member.
	//
	// example:
	//
	// Test user
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// UserID of the organization member in Quick BI.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Workspace ID.
	//
	// example:
	//
	// 7350a155-0e94-4c6c-8620-57bbec38****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// Workspace name.
	//
	// example:
	//
	// Test space
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListWorkspaceRoleUsersResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRoleUsersResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) GetNickName() *string {
	return s.NickName
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) GetUserId() *string {
	return s.UserId
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) SetNickName(v string) *ListWorkspaceRoleUsersResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) SetUserId(v string) *ListWorkspaceRoleUsersResponseBodyResultData {
	s.UserId = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) SetWorkspaceId(v string) *ListWorkspaceRoleUsersResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) SetWorkspaceName(v string) *ListWorkspaceRoleUsersResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
