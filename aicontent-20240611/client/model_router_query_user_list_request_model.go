// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryUserListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ModelRouterQueryUserListRequest
	GetKeyword() *string
	SetPageIndex(v int32) *ModelRouterQueryUserListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryUserListRequest
	GetPageSize() *int32
}

type ModelRouterQueryUserListRequest struct {
	// The search keyword.
	//
	// example:
	//
	// John
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ModelRouterQueryUserListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryUserListRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryUserListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterQueryUserListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryUserListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryUserListRequest) SetKeyword(v string) *ModelRouterQueryUserListRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterQueryUserListRequest) SetPageIndex(v int32) *ModelRouterQueryUserListRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryUserListRequest) SetPageSize(v int32) *ModelRouterQueryUserListRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryUserListRequest) Validate() error {
	return dara.Validate(s)
}
