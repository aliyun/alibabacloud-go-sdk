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
	SetUseNextToken(v bool) *DescribePropertyUserDetailRequest
	GetUseNextToken() *bool
	SetUser(v string) *DescribePropertyUserDetailRequest
	GetUser() *string
	SetUuid(v string) *DescribePropertyUserDetailRequest
	GetUuid() *string
}

type DescribePropertyUserDetailRequest struct {
	// Set which page of the returned results to start displaying the query results. The default value is **1**, indicating that the display starts from the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Whether the account name supports fuzzy search. To enable fuzzy search, set this parameter\\"s value to **1**; other values or an empty value indicate that fuzzy search is not supported.
	//
	// example:
	//
	// 1
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Whether the queried account has ROOT privileges. Possible values include:
	//
	// - **0**: No
	//
	// - **1**: Yes
	//
	// example:
	//
	// 0
	IsRoot *string `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	// The end timestamp for the last login retrieval. The unit is milliseconds.
	//
	// example:
	//
	// 1651298836000
	LastLoginTimeEnd *int64 `json:"LastLoginTimeEnd,omitempty" xml:"LastLoginTimeEnd,omitempty"`
	// The start timestamp for the last login retrieval. The unit is milliseconds.
	//
	// example:
	//
	// 164922523600
	LastLoginTimeStart *int64 `json:"LastLoginTimeStart,omitempty" xml:"LastLoginTimeStart,omitempty"`
	// Used to mark the starting position for reading. Leave it empty to start from the beginning.
	//
	// > For the first call, you do not need to fill in this field. The response will include the NextToken for the second call, and each subsequent call will include the NextToken for the next call.
	//
	// example:
	//
	// E17B501887A2D3AA5E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Set the number of account asset fingerprint information items to display per page during pagination. The default value is **10**, indicating that 10 items of account asset fingerprint information are displayed per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or IP address of the server to be queried.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Whether to use the NextToken method to fetch vulnerability list data. If this parameter is used, TotalCount will no longer be returned. Possible values:
	//
	// - **true**: Use the NextToken method.
	//
	// - **false**: Do not use the NextToken method.
	//
	// example:
	//
	// false
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The account name of the server to be queried.
	//
	// example:
	//
	// bin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server to be queried.
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
