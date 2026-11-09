// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizConfig(v string) *SearchMediaRequest
	GetBizConfig() *string
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
	// The business configuration. The value is a JSON string that can contain ProductionId (project ID).
	//
	// example:
	//
	// {"ProductionId":"pd_abc1234567"}
	BizConfig *string `json:"BizConfig,omitempty" xml:"BizConfig,omitempty"`
	// The category ID. You can obtain the category ID by using one of the following methods:
	//
	// - When you create a category by calling the CreateAssetCategory operation, the category ID is the value of the CategoryId response parameter.
	//
	// - When you query categories by calling the ListAssetCategories operation, the category ID is the value of the CategoryId parameter in the corresponding returned entry.
	//
	// example:
	//
	// 10
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The filter conditions. For the syntax rules, see [Media asset search protocol](https://www.alibabacloud.com/help/en/ims/developer-reference/media-asset-search-filter-description).
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
	// The number of entries to return on each page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pagination token, which is a 32-character string. You do not need to specify this parameter when you send a search request for the first time. If the search request returns data, the server returns this parameter to record the position of the current search data. Record the returned value and specify this parameter in subsequent searches based on the following requirements or recommendations: This parameter is required if you need to traverse all data that matches the search conditions. If the value of the PageNo parameter exceeds 200, specify this parameter to optimize search performance. You can only page forward, and the maximum paging distance is 1,000 media assets.
	//
	// example:
	//
	// F8C4F642184DBDA5D93907A70AAE****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// The sort fields and sort orders. Separate multiple values with commas (,). Format: field1:Desc,field2:Asc. The sort direction can only be Asc or Desc.
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

func (s *SearchMediaRequest) GetBizConfig() *string {
	return s.BizConfig
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

func (s *SearchMediaRequest) SetBizConfig(v string) *SearchMediaRequest {
	s.BizConfig = &v
	return s
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
