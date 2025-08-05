// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody
	GetData() *ListUsersResponseBodyData
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
}

type ListUsersResponseBody struct {
	// The returned data.
	Data *ListUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dd4816687424611405643e3730
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetData() *ListUsersResponseBodyData {
	return s.Data
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 64
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The users.
	Users []*ListUsersResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUsersResponseBodyData) GetUsers() []*ListUsersResponseBodyDataUsers {
	return s.Users
}

func (s *ListUsersResponseBodyData) SetPageNumber(v int32) *ListUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPageSize(v int32) *ListUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalCount(v int32) *ListUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUsers(v []*ListUsersResponseBodyDataUsers) *ListUsersResponseBodyData {
	s.Users = v
	return s
}

func (s *ListUsersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyDataUsers struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 167835629082
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// Bob@
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// The type of the account.
	//
	// example:
	//
	// ALIYUN
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// The display name.
	//
	// example:
	//
	// Bob
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1567253789
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUsersResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataUsers) GetAccountId() *string {
	return s.AccountId
}

func (s *ListUsersResponseBodyDataUsers) GetAccountName() *string {
	return s.AccountName
}

func (s *ListUsersResponseBodyDataUsers) GetAccountType() *string {
	return s.AccountType
}

func (s *ListUsersResponseBodyDataUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyDataUsers) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUsersResponseBodyDataUsers) SetAccountId(v string) *ListUsersResponseBodyDataUsers {
	s.AccountId = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetAccountName(v string) *ListUsersResponseBodyDataUsers {
	s.AccountName = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetAccountType(v string) *ListUsersResponseBodyDataUsers {
	s.AccountType = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetDisplayName(v string) *ListUsersResponseBodyDataUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetTenantId(v string) *ListUsersResponseBodyDataUsers {
	s.TenantId = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
