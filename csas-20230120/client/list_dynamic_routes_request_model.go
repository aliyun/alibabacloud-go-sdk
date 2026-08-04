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
	// The ID of the private access application for the dynamic route. You cannot filter by both the private access application ID and the private access tag ID. You can obtain the ID from the following sources:
	//
	// - [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): Queries multiple private access applications.
	//
	// - [CreatePrivateAccessApplication](~~CreatePrivateAccessApplication~~): Creates a private access application.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The number of the page to return for a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The IDs of the dynamic routes. You can specify up to 100 dynamic route IDs.
	DynamicRouteIds []*string `json:"DynamicRouteIds,omitempty" xml:"DynamicRouteIds,omitempty" type:"Repeated"`
	// The name of the dynamic route. The name must be 1 to 128 characters in length and can contain Chinese characters, letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the next hop instance for the dynamic route. You can obtain the ID from the following source:
	//
	// - [ListConnectors](~~ListConnectors~~): Queries multiple connectors.
	//
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The number of entries to return on each page for a paged query. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of regions where the SASE POP cluster endpoint is supported.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The status of the dynamic route. Valid values:
	//
	// - **Enabled**: The dynamic route is enabled.
	//
	// - **Disabled**: The dynamic route is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the private access tag for the dynamic route. You cannot filter by both the private access tag ID and the private access application ID. You can obtain the ID from the following sources:
	//
	// - [ListPrivateAccessTags](~~ListPrivateAccessTags~~): Queries multiple private access tags.
	//
	// - [CreatePrivateAccessTag](~~CreatePrivateAccessTag~~): Creates a private access tag.
	//
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
