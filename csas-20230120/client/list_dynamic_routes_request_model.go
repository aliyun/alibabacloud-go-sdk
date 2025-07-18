// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListDynamicRoutesRequest
	GetApplicationId() *string
	SetCurrentPage(v int32) *ListDynamicRoutesRequest
	GetCurrentPage() *int32
	SetDynamicRouteIds(v []*string) *ListDynamicRoutesRequest
	GetDynamicRouteIds() []*string
	SetName(v string) *ListDynamicRoutesRequest
	GetName() *string
	SetNextHop(v string) *ListDynamicRoutesRequest
	GetNextHop() *string
	SetPageSize(v int32) *ListDynamicRoutesRequest
	GetPageSize() *int32
	SetRegionIds(v []*string) *ListDynamicRoutesRequest
	GetRegionIds() []*string
	SetStatus(v string) *ListDynamicRoutesRequest
	GetStatus() *string
	SetTagId(v string) *ListDynamicRoutesRequest
	GetTagId() *string
}

type ListDynamicRoutesRequest struct {
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage     *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DynamicRouteIds []*string `json:"DynamicRouteIds,omitempty" xml:"DynamicRouteIds,omitempty" type:"Repeated"`
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListDynamicRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListDynamicRoutesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListDynamicRoutesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDynamicRoutesRequest) GetDynamicRouteIds() []*string {
	return s.DynamicRouteIds
}

func (s *ListDynamicRoutesRequest) GetName() *string {
	return s.Name
}

func (s *ListDynamicRoutesRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *ListDynamicRoutesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDynamicRoutesRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *ListDynamicRoutesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDynamicRoutesRequest) GetTagId() *string {
	return s.TagId
}

func (s *ListDynamicRoutesRequest) SetApplicationId(v string) *ListDynamicRoutesRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetCurrentPage(v int32) *ListDynamicRoutesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetDynamicRouteIds(v []*string) *ListDynamicRoutesRequest {
	s.DynamicRouteIds = v
	return s
}

func (s *ListDynamicRoutesRequest) SetName(v string) *ListDynamicRoutesRequest {
	s.Name = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetNextHop(v string) *ListDynamicRoutesRequest {
	s.NextHop = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetPageSize(v int32) *ListDynamicRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetRegionIds(v []*string) *ListDynamicRoutesRequest {
	s.RegionIds = v
	return s
}

func (s *ListDynamicRoutesRequest) SetStatus(v string) *ListDynamicRoutesRequest {
	s.Status = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetTagId(v string) *ListDynamicRoutesRequest {
	s.TagId = &v
	return s
}

func (s *ListDynamicRoutesRequest) Validate() error {
	return dara.Validate(s)
}
