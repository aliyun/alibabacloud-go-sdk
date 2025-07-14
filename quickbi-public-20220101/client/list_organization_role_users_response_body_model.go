// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationRoleUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListOrganizationRoleUsersResponseBody
	GetRequestId() *string
	SetResult(v *ListOrganizationRoleUsersResponseBodyResult) *ListOrganizationRoleUsersResponseBody
	GetResult() *ListOrganizationRoleUsersResponseBodyResult
	SetSuccess(v bool) *ListOrganizationRoleUsersResponseBody
	GetSuccess() *bool
}

type ListOrganizationRoleUsersResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// BCE45E6D-****-4F94-86BB-****2B1615FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the list of users under the organization role.
	Result *ListOrganizationRoleUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOrganizationRoleUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrganizationRoleUsersResponseBody) GetResult() *ListOrganizationRoleUsersResponseBodyResult {
	return s.Result
}

func (s *ListOrganizationRoleUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOrganizationRoleUsersResponseBody) SetRequestId(v string) *ListOrganizationRoleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBody) SetResult(v *ListOrganizationRoleUsersResponseBodyResult) *ListOrganizationRoleUsersResponseBody {
	s.Result = v
	return s
}

func (s *ListOrganizationRoleUsersResponseBody) SetSuccess(v bool) *ListOrganizationRoleUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOrganizationRoleUsersResponseBodyResult struct {
	// User list.
	Data []*ListOrganizationRoleUsersResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of items per page as set in the request.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 10
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListOrganizationRoleUsersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRoleUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersResponseBodyResult) GetData() []*ListOrganizationRoleUsersResponseBodyResultData {
	return s.Data
}

func (s *ListOrganizationRoleUsersResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListOrganizationRoleUsersResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrganizationRoleUsersResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListOrganizationRoleUsersResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetData(v []*ListOrganizationRoleUsersResponseBodyResultData) *ListOrganizationRoleUsersResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetPageNum(v int32) *ListOrganizationRoleUsersResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetPageSize(v int32) *ListOrganizationRoleUsersResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetTotalNum(v int32) *ListOrganizationRoleUsersResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetTotalPages(v int32) *ListOrganizationRoleUsersResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListOrganizationRoleUsersResponseBodyResultData struct {
	// Nickname of the organization member.
	//
	// example:
	//
	// Test User
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// UserID of the organization member in Quick BI.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListOrganizationRoleUsersResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRoleUsersResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersResponseBodyResultData) GetNickName() *string {
	return s.NickName
}

func (s *ListOrganizationRoleUsersResponseBodyResultData) GetUserId() *string {
	return s.UserId
}

func (s *ListOrganizationRoleUsersResponseBodyResultData) SetNickName(v string) *ListOrganizationRoleUsersResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResultData) SetUserId(v string) *ListOrganizationRoleUsersResponseBodyResultData {
	s.UserId = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
