// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDynamicRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *UpdateDynamicRouteRequest
	GetApplicationIds() []*string
	SetApplicationType(v string) *UpdateDynamicRouteRequest
	GetApplicationType() *string
	SetDescription(v string) *UpdateDynamicRouteRequest
	GetDescription() *string
	SetDynamicRouteId(v string) *UpdateDynamicRouteRequest
	GetDynamicRouteId() *string
	SetDynamicRouteType(v string) *UpdateDynamicRouteRequest
	GetDynamicRouteType() *string
	SetModifyType(v string) *UpdateDynamicRouteRequest
	GetModifyType() *string
	SetName(v string) *UpdateDynamicRouteRequest
	GetName() *string
	SetNextHop(v string) *UpdateDynamicRouteRequest
	GetNextHop() *string
	SetPriority(v int32) *UpdateDynamicRouteRequest
	GetPriority() *int32
	SetRegionIds(v []*string) *UpdateDynamicRouteRequest
	GetRegionIds() []*string
	SetStatus(v string) *UpdateDynamicRouteRequest
	GetStatus() *string
	SetTagIds(v []*string) *UpdateDynamicRouteRequest
	GetTagIds() []*string
}

type UpdateDynamicRouteRequest struct {
	// A collection of private network access application IDs for the dynamic route. You can specify up to 200 IDs. This parameter is required when **ApplicationType*	- is set to **Application**. Specify either this parameter or **TagIds**, but not both. Do not specify this parameter when **ApplicationType*	- is set to **All**.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The application type of the dynamic route. Valid values:
	//
	// - **All**: All applications.
	//
	// - **Application**: Application.
	//
	// - **Tag**: Tag.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The description of the dynamic route. The description must be 1 to 128 characters long and can contain letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// example:
	//
	// 这是一条动态路由
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the dynamic route.
	//
	// This parameter is required.
	//
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// The type of the dynamic route. Valid value: **connector**: leased line.
	//
	// example:
	//
	// connector
	DynamicRouteType *string `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	// The modification type of the dynamic route. Valid values:
	//
	// - **Cover*	- (default): Use the values of **RegionIds**, **ApplicationIds**, and **TagIds*	- to overwrite the existing regions, private network access application IDs, and private network access tag IDs supported by the SASE POP cluster access points.
	//
	// - **Append**: Add the values specified in **RegionIds**, **ApplicationIds**, and **TagIds*	- to the existing regions, private network access application IDs, and private network access tag IDs supported by the SASE POP cluster access points.
	//
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The name of the dynamic route. The name must be 1 to 128 characters long and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next hop instance ID of the dynamic route. To get valid values, see:
	//
	// - [ListConnectors](~~ListConnectors~~): Query connectors in bulk.
	//
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The route priority. A value of 1 indicates the highest priority. Valid values: 1 to 99.
	//
	// example:
	//
	// 99
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// A list of regions supported by the SASE POP cluster access points.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The status of the dynamic route. Valid values:
	//
	// - **Enabled**: Enable the route.
	//
	// - **Disabled**: Disable the route.
	//
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A collection of private network access tag IDs for the dynamic route. You can specify up to 40 IDs. This parameter is required when **ApplicationType*	- is set to **Tag**. Specify either this parameter or **ApplicationIds**, but not both. Do not specify this parameter when **ApplicationType*	- is set to **All**.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s UpdateDynamicRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateDynamicRouteRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *UpdateDynamicRouteRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *UpdateDynamicRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDynamicRouteRequest) GetDynamicRouteId() *string {
	return s.DynamicRouteId
}

func (s *UpdateDynamicRouteRequest) GetDynamicRouteType() *string {
	return s.DynamicRouteType
}

func (s *UpdateDynamicRouteRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *UpdateDynamicRouteRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDynamicRouteRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *UpdateDynamicRouteRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateDynamicRouteRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *UpdateDynamicRouteRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateDynamicRouteRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *UpdateDynamicRouteRequest) SetApplicationIds(v []*string) *UpdateDynamicRouteRequest {
	s.ApplicationIds = v
	return s
}

func (s *UpdateDynamicRouteRequest) SetApplicationType(v string) *UpdateDynamicRouteRequest {
	s.ApplicationType = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetDescription(v string) *UpdateDynamicRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetDynamicRouteId(v string) *UpdateDynamicRouteRequest {
	s.DynamicRouteId = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetDynamicRouteType(v string) *UpdateDynamicRouteRequest {
	s.DynamicRouteType = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetModifyType(v string) *UpdateDynamicRouteRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetName(v string) *UpdateDynamicRouteRequest {
	s.Name = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetNextHop(v string) *UpdateDynamicRouteRequest {
	s.NextHop = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetPriority(v int32) *UpdateDynamicRouteRequest {
	s.Priority = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetRegionIds(v []*string) *UpdateDynamicRouteRequest {
	s.RegionIds = v
	return s
}

func (s *UpdateDynamicRouteRequest) SetStatus(v string) *UpdateDynamicRouteRequest {
	s.Status = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetTagIds(v []*string) *UpdateDynamicRouteRequest {
	s.TagIds = v
	return s
}

func (s *UpdateDynamicRouteRequest) Validate() error {
	return dara.Validate(s)
}
