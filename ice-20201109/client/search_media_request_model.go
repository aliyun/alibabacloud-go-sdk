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
	// The custom filter, specified as a JSON string. Supports the integer field intField1 and the string fields strField1 and strField2. Use only one match type per field. Conditions on different fields are combined with a logical AND.
	//
	// - Exact match: `{"intField1":12,"strField1":"abc"}`
	//
	// - Multi-value match: `{"intField1":[12,13],"strField1":["abc","cd"]}`
	//
	// - Range match: `{"intField1":{"gte":12,"lte":13}}`
	//
	// example:
	//
	// {"intField1":{"gte":12,"lte":13},"strField2":["cd","de"],"strField1":"abc"}
	CustomFilters *string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	// The entity ID.
	//
	// example:
	//
	// 2d3bf1e35a1e42b5ab338d701efa****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The filter condition for the search. <props="china">For syntax rules, see the [Media Search Protocol](https://help.aliyun.com/document_detail/2584256.html).
	//
	// example:
	//
	// Title = \\"China\\" and utcCreate = [\\"1693367158561\\",\\"1693367158562\\"]
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The page number to return. The default value is 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of results per page. The default value is 10, and the maximum value is 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scroll token for deep pagination. It is a 32-character string. This parameter is not required for the first search request. If a search is successful, the response includes a `ScrollToken` to mark the current position. Use this token in subsequent requests to retrieve the next page of results. This parameter is required to iterate through all matching results. For optimal performance, use this parameter when the `PageNo` value exceeds 200. You can scroll only forward, up to a maximum of 1,000 media assets.
	//
	// example:
	//
	// F8C4F642184DBDA5D93907A70AAE****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// The name of the search library.
	//
	// example:
	//
	// test-1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// The sort field and sort order. Separate multiple sort criteria with a comma (,).
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
