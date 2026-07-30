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
	SetMatch(v string) *SearchMediaRequest
	GetMatch() *string
	SetPageNo(v int32) *SearchMediaRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchMediaRequest
	GetPageSize() *int32
	SetScrollToken(v string) *SearchMediaRequest
	GetScrollToken() *string
	SetSortBy(v string) *SearchMediaRequest
	GetSortBy() *string
}

type SearchMediaRequest struct {
	// The category ID. You can obtain the ID by using the following methods:
	//
	// example:
	//
	// 10
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The filter condition. For syntax rules, see [Media asset search protocol](https://www.alibabacloud.com/help/en/ims/developer-reference/media-asset-search-filter-description).
	//
	// example:
	//
	// title = \\"China\\" and utcCreate = [\\"1693367158561\\",\\"1693367158562\\"]
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pagination token. A 32-character string. You do not need to set this parameter for the first search request. When the search request matches data, the server returns this parameter value to record the current position of the search data. Record the returned parameter value and set this parameter in the next search request based on the following requirements or suggestions: This parameter must be set if you want to traverse all data that matches the search conditions. If the PageNo parameter value exceeds 200, set this parameter to optimize search performance. You can only page forward, with a maximum paging distance of 1000 media assets.
	//
	// example:
	//
	// F8C4F642184DBDA5D93907A70AAE****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// The sort fields and sort orders, separated by commas (,). Format: field1:Desc,field2:Asc. The direction can only be Asc or Desc.
	//
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

func (s *SearchMediaRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchMediaRequest) SetCategoryId(v int64) *SearchMediaRequest {
	s.CategoryId = &v
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

func (s *SearchMediaRequest) SetSortBy(v string) *SearchMediaRequest {
	s.SortBy = &v
	return s
}

func (s *SearchMediaRequest) Validate() error {
	return dara.Validate(s)
}
