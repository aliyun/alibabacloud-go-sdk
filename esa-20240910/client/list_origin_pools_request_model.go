// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListOriginPoolsRequest
	GetMatchType() *string
	SetName(v string) *ListOriginPoolsRequest
	GetName() *string
	SetOrderBy(v string) *ListOriginPoolsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListOriginPoolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOriginPoolsRequest
	GetPageSize() *int32
	SetSiteId(v int64) *ListOriginPoolsRequest
	GetSiteId() *int64
}

type ListOriginPoolsRequest struct {
	// Type of name match query, supporting the following two types, with exact match as the default.
	//
	// - fuzzy: Fuzzy query.
	//
	// - exact: Exact query.
	//
	// example:
	//
	// exact
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// Name of the origin pool.
	//
	// example:
	//
	// pool1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Sorting, supports ascending and descending order by ID, default is descending by ID, which is positively correlated with creation time.
	//
	// - -id: Sort by ID in descending order.
	//
	// - id: Sort by ID in ascending order.
	//
	// example:
	//
	// id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// Page number, default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, an integer greater than 0, with a maximum of 500. If the value exceeds 500, it will be set to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 151538882642832
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListOriginPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListOriginPoolsRequest) GetName() *string {
	return s.Name
}

func (s *ListOriginPoolsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListOriginPoolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOriginPoolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOriginPoolsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListOriginPoolsRequest) SetMatchType(v string) *ListOriginPoolsRequest {
	s.MatchType = &v
	return s
}

func (s *ListOriginPoolsRequest) SetName(v string) *ListOriginPoolsRequest {
	s.Name = &v
	return s
}

func (s *ListOriginPoolsRequest) SetOrderBy(v string) *ListOriginPoolsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListOriginPoolsRequest) SetPageNumber(v int32) *ListOriginPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOriginPoolsRequest) SetPageSize(v int32) *ListOriginPoolsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOriginPoolsRequest) SetSiteId(v int64) *ListOriginPoolsRequest {
	s.SiteId = &v
	return s
}

func (s *ListOriginPoolsRequest) Validate() error {
	return dara.Validate(s)
}
