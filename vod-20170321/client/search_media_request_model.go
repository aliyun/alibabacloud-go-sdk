// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v string) *SearchMediaRequest
	GetFields() *string
	SetMatch(v string) *SearchMediaRequest
	GetMatch() *string
	SetPageNo(v int32) *SearchMediaRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchMediaRequest
	GetPageSize() *int32
	SetScrollToken(v string) *SearchMediaRequest
	GetScrollToken() *string
	SetSearchType(v string) *SearchMediaRequest
	GetSearchType() *string
	SetSortBy(v string) *SearchMediaRequest
	GetSortBy() *string
}

type SearchMediaRequest struct {
	// The media asset fields to return in the search results.
	//
	// By default, only basic media asset fields are returned. You can specify additional media asset fields to return. For more information, see [Usage examples](https://help.aliyun.com/document_detail/99179.html).
	//
	// example:
	//
	// Title,CoverURL
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The filter conditions. For syntax rules, see [Search protocol syntax](https://help.aliyun.com/document_detail/86991.html).
	//
	// example:
	//
	// field = value
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The page number. Default value: **1**.
	//
	// > If this parameter exceeds **200**, set the ScrollToken parameter as well.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of records per page. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pagination token. The value is a 32-character string.
	//
	// You do not need to set this parameter for the first search request. When the search request matches data, the server returns this parameter value, which records the current position of the search data. Record the returned value and set this parameter in the next search request based on the following requirements or recommendations:
	//
	// - If SearchType is set to **video*	- or **audio*	- and you need to traverse all data that matches the search conditions, this parameter is required.
	//
	// - If PageNo exceeds **200**, set this parameter to optimize search performance.
	//
	// example:
	//
	// 24e0fba7188fae707e146esa54****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// The type of media asset to search. Valid values:
	//
	// - **video*	- (default): video.
	//
	// - **audio**: audio.
	//
	// - **image**: image.
	//
	// - **attached**: auxiliary media asset.
	//
	// > If this parameter is set to **video*	- or **audio*	- and you need to traverse all data that matches the search conditions, you must set the ScrollToken parameter.
	//
	// example:
	//
	// video
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// The sort field and sort order. Separate multiple values with commas (,). Valid values:
	//
	// - **CreationTime:Desc*	- (default): sorts by creation time in descending order.
	//
	// - **CreationTime:Asc**: sorts by creation time in ascending order.
	//
	// > - For sort field examples, see [Sort fields](https://help.aliyun.com/document_detail/99179.html).
	//
	// > - When retrieving the first 5,000 records of search results, up to three sort fields are supported.
	//
	// > - When retrieving all data that matches the search conditions, only one sort field is supported.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s SearchMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaRequest) GetFields() *string {
	return s.Fields
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

func (s *SearchMediaRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *SearchMediaRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchMediaRequest) SetFields(v string) *SearchMediaRequest {
	s.Fields = &v
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

func (s *SearchMediaRequest) SetSearchType(v string) *SearchMediaRequest {
	s.SearchType = &v
	return s
}

func (s *SearchMediaRequest) SetSortBy(v string) *SearchMediaRequest {
	s.SortBy = &v
	return s
}

func (s *SearchMediaRequest) Validate() error {
	return dara.Validate(s)
}
