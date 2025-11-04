// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFilters(v string) *SearchMediaRequest
	GetCustomFilters() *string
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
	// {}
	CustomFilters *string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	// The ID of the entity.
	//
	// example:
	//
	// 2d3bf1e35a1e42b5ab338d701efa7603
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The filter conditions. For more information about the parameter syntax
	//
	// <props="china">, see [Media asset search protocols](https://help.aliyun.com/document_detail/2584256.html).
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pagination identifier. The value can be up to 32 characters in length. The first time you call this operation for each new search, you do not need to specify this parameter. The value of this parameter is returned each time data records that meet the specified filter condition are found. The value is used to record the current position of queried data. Record the returned parameter value and set this parameter according to the following requirements during the next search: If you need to traverse all data that meets the filter criteria, you must set the ScrollToken parameter. If the value of the PageNo parameter exceeds 200, we recommend that you set this parameter to optimize search performance. You can only page backward. You can page a maximum of 1,000 entries in an operation.
	//
	// example:
	//
	// F8C4F642184DBDA5D93907A70AAE****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// The search library.
	//
	// example:
	//
	// test-1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// The sort field and order. Separate multiple parameters with commas (,).
	//
	// example:
	//
	// utcCreate:Desc, utcModified:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s SearchMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaRequest) GetCustomFilters() *string {
	return s.CustomFilters
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

func (s *SearchMediaRequest) SetCustomFilters(v string) *SearchMediaRequest {
	s.CustomFilters = &v
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
