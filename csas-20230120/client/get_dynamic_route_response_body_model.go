// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDynamicRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRoute(v *GetDynamicRouteResponseBodyDynamicRoute) *GetDynamicRouteResponseBody
	GetDynamicRoute() *GetDynamicRouteResponseBodyDynamicRoute
	SetRequestId(v string) *GetDynamicRouteResponseBody
	GetRequestId() *string
}

type GetDynamicRouteResponseBody struct {
	DynamicRoute *GetDynamicRouteResponseBodyDynamicRoute `json:"DynamicRoute,omitempty" xml:"DynamicRoute,omitempty" type:"Struct"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDynamicRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetDynamicRouteResponseBody) GetDynamicRoute() *GetDynamicRouteResponseBodyDynamicRoute {
	return s.DynamicRoute
}

func (s *GetDynamicRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDynamicRouteResponseBody) SetDynamicRoute(v *GetDynamicRouteResponseBodyDynamicRoute) *GetDynamicRouteResponseBody {
	s.DynamicRoute = v
	return s
}

func (s *GetDynamicRouteResponseBody) SetRequestId(v string) *GetDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDynamicRouteResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDynamicRouteResponseBodyDynamicRoute struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2023-02-09 10:31:47
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dr-16ff07c8207d****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// example:
	//
	// connector
	DynamicRouteType *string `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// example:
	//
	// 1
	Priority  *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s GetDynamicRouteResponseBodyDynamicRoute) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicRouteResponseBodyDynamicRoute) GoString() string {
	return s.String()
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetDescription() *string {
	return s.Description
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetDynamicRouteId() *string {
	return s.DynamicRouteId
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetDynamicRouteType() *string {
	return s.DynamicRouteType
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetName() *string {
	return s.Name
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetNextHop() *string {
	return s.NextHop
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetPriority() *int32 {
	return s.Priority
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetStatus() *string {
	return s.Status
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) GetTagIds() []*string {
	return s.TagIds
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetApplicationIds(v []*string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.ApplicationIds = v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetApplicationType(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.ApplicationType = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetCreateTime(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.CreateTime = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetDescription(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.Description = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetDynamicRouteId(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.DynamicRouteId = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetDynamicRouteType(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.DynamicRouteType = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetName(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.Name = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetNextHop(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.NextHop = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetPriority(v int32) *GetDynamicRouteResponseBodyDynamicRoute {
	s.Priority = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetRegionIds(v []*string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.RegionIds = v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetStatus(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.Status = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetTagIds(v []*string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.TagIds = v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) Validate() error {
	return dara.Validate(s)
}
