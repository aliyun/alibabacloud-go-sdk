// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *SearchMediaRequest
	GetCategoryId() *int64
	SetEntityId(v string) *SearchMediaRequest
	GetEntityId() *string
	SetMatch(v string) *SearchMediaRequest
	GetMatch() *string
	SetPageNo(v int32) *SearchMediaRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchMediaRequest
	GetPageSize() *int32
	SetScrollToken(v string) *SearchMediaRequest
	GetScrollToken() *string
	SetSearchLibName(v string) *SearchMediaRequest
	GetSearchLibName() *string
	SetSortBy(v string) *SearchMediaRequest
	GetSortBy() *string
}

type SearchMediaRequest struct {
	// example:
	//
	// 10
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// BaseMedia
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// title = \\"中国\\" and utcCreate = [\\"1693367158561\\",\\"1693367158562\\"]
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F8C4F642184DBDA5D93907A70AAE****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// example:
	//
	// test-1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// example:
	//
	// utcCreate:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s SearchMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *SearchMediaRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchMediaRequest) GetMatch() *string {
	return s.Match
}

func (s *SearchMediaRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchMediaRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchMediaRequest) GetScrollToken() *string {
	return s.ScrollToken
}

func (s *SearchMediaRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *SearchMediaRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchMediaRequest) SetCategoryId(v int64) *SearchMediaRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchMediaRequest) SetEntityId(v string) *SearchMediaRequest {
	s.EntityId = &v
	return s
}

func (s *SearchMediaRequest) SetMatch(v string) *SearchMediaRequest {
	s.Match = &v
	return s
}

func (s *SearchMediaRequest) SetPageNo(v int32) *SearchMediaRequest {
	s.PageNo = &v
	return s
}

func (s *SearchMediaRequest) SetPageSize(v int32) *SearchMediaRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMediaRequest) SetScrollToken(v string) *SearchMediaRequest {
	s.ScrollToken = &v
	return s
}

func (s *SearchMediaRequest) SetSearchLibName(v string) *SearchMediaRequest {
	s.SearchLibName = &v
	return s
}

func (s *SearchMediaRequest) SetSortBy(v string) *SearchMediaRequest {
	s.SortBy = &v
	return s
}

func (s *SearchMediaRequest) Validate() error {
	return dara.Validate(s)
}
