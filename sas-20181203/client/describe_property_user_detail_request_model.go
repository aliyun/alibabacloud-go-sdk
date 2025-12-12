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
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the fuzzy search by account name is supported. If you want to use fuzzy search, set the parameter to **1**. If you set the parameter to a different value or leave the parameter empty, fuzzy search is not supported.
	//
	// example:
	//
	// 1
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Specifies whether the account has root permissions. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	IsRoot *string `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	// The timestamp of the last logoff from the account. Unit: milliseconds.
	//
	// example:
	//
	// 164922523600
	LastLoginTimeEnd *int64 `json:"LastLoginTimeEnd,omitempty" xml:"LastLoginTimeEnd,omitempty"`
	// The timestamp of the last logon to the account. Unit: milliseconds.
	//
	// example:
	//
	// 164922523600
	LastLoginTimeStart *int64  `json:"LastLoginTimeStart,omitempty" xml:"LastLoginTimeStart,omitempty"`
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UseNextToken *bool   `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The name of the account to which the server belongs.
	//
	// example:
	//
	// bin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server.
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
