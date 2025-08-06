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
	// The media asset fields to return in the query results.
	//
	// By default, only the basic media asset fields are returned. You can specify additional media asset fields that need to be returned in the request. For more information, see the "API examples" section of the [Search for media asset information](https://help.aliyun.com/document_detail/99179.html) topic.
	//
	// example:
	//
	// Title,CoverURL
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The filter condition. For more information about the syntax, see [Protocol for media asset search](https://help.aliyun.com/document_detail/86991.html).
	//
	// example:
	//
	// field = value
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// > If the value of this parameter exceeds **200**, we recommend that you set the ScrollToken parameter as well.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pagination identifier. The password must be 32 characters in length The first time you call this operation for each new search, you do not need to specify this parameter. The value of this parameter is returned each time data records that meet the specified filter condition are found. The value is used to record the current position of queried data. Record the returned parameter value and set this parameter according to the following requirements during the next search:
	//
	// 	- If SearchType is set to **video*	- or **audio*	- and you need to traverse all data that meets the filter criteria, you must set the ScrollToken parameter.
	//
	// 	- If the value of the PageNo parameter exceeds **200**, we recommend that you set this parameter to optimize search performance.
	//
	// example:
	//
	// 24e0fba7188fae707e146esa54****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// The type of the media asset that you want to query. Default value: video. Valid values:
	//
	// 	- **video**
	//
	// 	- **audio**
	//
	// 	- **image**
	//
	// 	- **attached**
	//
	// > If this parameter is set to **video*	- or **audio*	- and you want to traverse all data that meets the filter criteria, you must set the ScrollToken parameter.
	//
	// example:
	//
	// video
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// The sort field and order. Separate multiple values with commas (,). Default value: CreationTime:Desc. Valid values:
	//
	// 	- **CreationTime:Desc**: The results are sorted in reverse chronological order based on the creation time.
	//
	// 	- **CreationTime:Asc**: The results are sorted in chronological order based on the creation time.
	//
	// > 	- For more information about the sort field, see "Sort field" in the [Search for media asset information](https://help.aliyun.com/document_detail/99179.html) topic.
	//
	// > 	- To obtain the first 5,000 data records that meet the specified filter criteria, you can specify a maximum of three sort fields.
	//
	// > 	- To obtain all the data records that meet the specified filter criteria, you can specify only one sort field.
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
