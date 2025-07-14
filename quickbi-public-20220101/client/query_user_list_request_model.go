// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryUserListRequest
	GetKeyword() *string
	SetPageNum(v int32) *QueryUserListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryUserListRequest
	GetPageSize() *int32
}

type QueryUserListRequest struct {
	// The keyword of the username or nickname of the organization member.
	//
	// example:
	//
	// Test user
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Current page number for organization member list:
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page in a paged query.
	//
	// 	- Default value: 10.
	//
	// 	- Maximum value: 1,000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryUserListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserListRequest) GoString() string {
	return s.String()
}

func (s *QueryUserListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryUserListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryUserListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryUserListRequest) SetKeyword(v string) *QueryUserListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryUserListRequest) SetPageNum(v int32) *QueryUserListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryUserListRequest) SetPageSize(v int32) *QueryUserListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryUserListRequest) Validate() error {
	return dara.Validate(s)
}
