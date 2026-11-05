// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUserDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertyUserDetailRequest
	GetCurrentPage() *int32
	SetExtend(v string) *DescribePropertyUserDetailRequest
	GetExtend() *string
	SetIsRoot(v string) *DescribePropertyUserDetailRequest
	GetIsRoot() *string
	SetLastLoginTimeEnd(v int64) *DescribePropertyUserDetailRequest
	GetLastLoginTimeEnd() *int64
	SetLastLoginTimeStart(v int64) *DescribePropertyUserDetailRequest
	GetLastLoginTimeStart() *int64
	SetNextToken(v string) *DescribePropertyUserDetailRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribePropertyUserDetailRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribePropertyUserDetailRequest
	GetRemark() *string
	SetResourceDirectoryAccountId(v int64) *DescribePropertyUserDetailRequest
	GetResourceDirectoryAccountId() *int64
	SetUseNextToken(v bool) *DescribePropertyUserDetailRequest
	GetUseNextToken() *bool
	SetUser(v string) *DescribePropertyUserDetailRequest
	GetUser() *string
	SetUuid(v string) *DescribePropertyUserDetailRequest
	GetUuid() *string
}

type DescribePropertyUserDetailRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether fuzzy match is supported for the account name. To enable fuzzy match, set this parameter to **1**. Other values or an empty value indicate that fuzzy match is not supported.
	//
	// example:
	//
	// 1
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Specifies whether the account to query has root permissions. Valid values:
	//
	// - **0**: No.
	//
	// - **1**: Yes.
	//
	// example:
	//
	// 0
	IsRoot *string `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	// The end timestamp for retrieving the last logon time of the account. Unit: milliseconds.
	//
	// example:
	//
	// 1651298836000
	LastLoginTimeEnd *int64 `json:"LastLoginTimeEnd,omitempty" xml:"LastLoginTimeEnd,omitempty"`
	// The start timestamp for retrieving the last logon time of the account. Unit: milliseconds.
	//
	// example:
	//
	// 164922523600
	LastLoginTimeStart *int64 `json:"LastLoginTimeStart,omitempty" xml:"LastLoginTimeStart,omitempty"`
	// The token that specifies the position from which to start reading. If this parameter is left empty, the reading starts from the beginning.
	//
	// > You do not need to specify this parameter for the first call. The response will contain the NextToken for the second call, and each subsequent response will contain the NextToken for the next call.
	//
	// example:
	//
	// E17B501887A2D3AA5E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The Settings for the number of account Asset fingerprint entries to return on each page during paging for a paged query. Default value: **10**, which indicates that 10 account Asset fingerprint entries are returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or IP address of the server to query.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the management account for the member accounts in the resource folder.
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1234567890123456
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether to use the NextToken method to retrieve vulnerability list data. If this parameter is used, TotalCount is no longer returned. Valid values:
	//
	// - **true**: Use the NextToken method.
	//
	// - **false**: Do not use the NextToken method.
	//
	// example:
	//
	// false
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The name of the account on the server to query.
	//
	// example:
	//
	// bin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server to query.
	//
	// example:
	//
	// 50d213b4-3a35-427a-b8a5-04b0c7e1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyUserDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyUserDetailRequest) GetExtend() *string {
	return s.Extend
}

func (s *DescribePropertyUserDetailRequest) GetIsRoot() *string {
	return s.IsRoot
}

func (s *DescribePropertyUserDetailRequest) GetLastLoginTimeEnd() *int64 {
	return s.LastLoginTimeEnd
}

func (s *DescribePropertyUserDetailRequest) GetLastLoginTimeStart() *int64 {
	return s.LastLoginTimeStart
}

func (s *DescribePropertyUserDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePropertyUserDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyUserDetailRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribePropertyUserDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribePropertyUserDetailRequest) GetUseNextToken() *bool {
	return s.UseNextToken
}

func (s *DescribePropertyUserDetailRequest) GetUser() *string {
	return s.User
}

func (s *DescribePropertyUserDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyUserDetailRequest) SetCurrentPage(v int32) *DescribePropertyUserDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetExtend(v string) *DescribePropertyUserDetailRequest {
	s.Extend = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetIsRoot(v string) *DescribePropertyUserDetailRequest {
	s.IsRoot = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetLastLoginTimeEnd(v int64) *DescribePropertyUserDetailRequest {
	s.LastLoginTimeEnd = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetLastLoginTimeStart(v int64) *DescribePropertyUserDetailRequest {
	s.LastLoginTimeStart = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetNextToken(v string) *DescribePropertyUserDetailRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetPageSize(v int32) *DescribePropertyUserDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetRemark(v string) *DescribePropertyUserDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribePropertyUserDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetUseNextToken(v bool) *DescribePropertyUserDetailRequest {
	s.UseNextToken = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetUser(v string) *DescribePropertyUserDetailRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetUuid(v string) *DescribePropertyUserDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) Validate() error {
	return dara.Validate(s)
}
