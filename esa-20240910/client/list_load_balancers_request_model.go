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
	MatchType  *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
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
