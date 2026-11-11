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
	// The match type for name-based queries. Valid values:
	//
	// - fuzzy: fuzzy search.
	//
	// - exact: exact search.
	//
	// Default value: exact.
	//
	// example:
	//
	// exact
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The name of the origin address pool.
	//
	// example:
	//
	// pool1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Supports ascending or descending order by ID. The ID is positively correlated with the creation time. Default value: descending order by ID. Valid values:
	//
	// - id: descending order by ID.
	//
	// - id: ascending order by ID.
	//
	// example:
	//
	// id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value must be a positive integer. Maximum value: 500. If the value exceeds 500, it is reset to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
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
