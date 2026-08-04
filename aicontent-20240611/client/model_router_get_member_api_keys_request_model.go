// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ModelRouterGetMemberApiKeysRequest
	GetKeyword() *string
	SetPageIndex(v int32) *ModelRouterGetMemberApiKeysRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterGetMemberApiKeysRequest
	GetPageSize() *int32
}

type ModelRouterGetMemberApiKeysRequest struct {
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

func (s ModelRouterGetMemberApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberApiKeysRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterGetMemberApiKeysRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterGetMemberApiKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterGetMemberApiKeysRequest) SetKeyword(v string) *ModelRouterGetMemberApiKeysRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysRequest) SetPageIndex(v int32) *ModelRouterGetMemberApiKeysRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysRequest) SetPageSize(v int32) *ModelRouterGetMemberApiKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
