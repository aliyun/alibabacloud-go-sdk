// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *CreateDynamicRouteRequest
	GetApplicationIds() []*string
	SetApplicationType(v string) *CreateDynamicRouteRequest
	GetApplicationType() *string
	SetDescription(v string) *CreateDynamicRouteRequest
	GetDescription() *string
	SetDynamicRouteType(v string) *CreateDynamicRouteRequest
	GetDynamicRouteType() *string
	SetName(v string) *CreateDynamicRouteRequest
	GetName() *string
	SetNextHop(v string) *CreateDynamicRouteRequest
	GetNextHop() *string
	SetPriority(v int32) *CreateDynamicRouteRequest
	GetPriority() *int32
	SetRegionIds(v []*string) *CreateDynamicRouteRequest
	GetRegionIds() []*string
	SetStatus(v string) *CreateDynamicRouteRequest
	GetStatus() *string
	SetTagIds(v []*string) *CreateDynamicRouteRequest
	GetTagIds() []*string
}

type CreateDynamicRouteRequest struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// connector
	DynamicRouteType *string `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 99
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Disabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s CreateDynamicRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateDynamicRouteRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *CreateDynamicRouteRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateDynamicRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDynamicRouteRequest) GetDynamicRouteType() *string {
	return s.DynamicRouteType
}

func (s *CreateDynamicRouteRequest) GetName() *string {
	return s.Name
}

func (s *CreateDynamicRouteRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateDynamicRouteRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateDynamicRouteRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *CreateDynamicRouteRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateDynamicRouteRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *CreateDynamicRouteRequest) SetApplicationIds(v []*string) *CreateDynamicRouteRequest {
	s.ApplicationIds = v
	return s
}

func (s *CreateDynamicRouteRequest) SetApplicationType(v string) *CreateDynamicRouteRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetDescription(v string) *CreateDynamicRouteRequest {
	s.Description = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetDynamicRouteType(v string) *CreateDynamicRouteRequest {
	s.DynamicRouteType = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetName(v string) *CreateDynamicRouteRequest {
	s.Name = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetNextHop(v string) *CreateDynamicRouteRequest {
	s.NextHop = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetPriority(v int32) *CreateDynamicRouteRequest {
	s.Priority = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetRegionIds(v []*string) *CreateDynamicRouteRequest {
	s.RegionIds = v
	return s
}

func (s *CreateDynamicRouteRequest) SetStatus(v string) *CreateDynamicRouteRequest {
	s.Status = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetTagIds(v []*string) *CreateDynamicRouteRequest {
	s.TagIds = v
	return s
}

func (s *CreateDynamicRouteRequest) Validate() error {
	return dara.Validate(s)
}
