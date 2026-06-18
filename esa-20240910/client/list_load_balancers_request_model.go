// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListLoadBalancersRequest
	GetMatchType() *string
	SetName(v string) *ListLoadBalancersRequest
	GetName() *string
	SetOrderBy(v string) *ListLoadBalancersRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListLoadBalancersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLoadBalancersRequest
	GetPageSize() *int32
	SetSiteId(v int64) *ListLoadBalancersRequest
	GetSiteId() *int64
}

type ListLoadBalancersRequest struct {
	// The matching strategy to use when querying by name. Valid values:
	//
	// - `fuzzy`: Performs a fuzzy match.
	//
	// - `exact`: Performs an exact match.
	//
	// example:
	//
	// fuzzy
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The name of the load balancer.
	//
	// example:
	//
	// lb.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort field. Only sorting by ID is supported. IDs are assigned chronologically. Specify `id` for ascending order or `-id` for descending order. If this parameter is omitted, the results are sorted by ID in descending order.
	//
	// example:
	//
	// id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number for pagination.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The site ID. You can obtain this ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21655860979****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListLoadBalancersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListLoadBalancersRequest) GetName() *string {
	return s.Name
}

func (s *ListLoadBalancersRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListLoadBalancersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLoadBalancersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLoadBalancersRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListLoadBalancersRequest) SetMatchType(v string) *ListLoadBalancersRequest {
	s.MatchType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetName(v string) *ListLoadBalancersRequest {
	s.Name = &v
	return s
}

func (s *ListLoadBalancersRequest) SetOrderBy(v string) *ListLoadBalancersRequest {
	s.OrderBy = &v
	return s
}

func (s *ListLoadBalancersRequest) SetPageNumber(v int32) *ListLoadBalancersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLoadBalancersRequest) SetPageSize(v int32) *ListLoadBalancersRequest {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalancersRequest) SetSiteId(v int64) *ListLoadBalancersRequest {
	s.SiteId = &v
	return s
}

func (s *ListLoadBalancersRequest) Validate() error {
	return dara.Validate(s)
}
