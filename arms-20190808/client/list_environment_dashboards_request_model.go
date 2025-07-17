// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentDashboardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListEnvironmentDashboardsRequest
	GetAddonName() *string
	SetEnvironmentId(v string) *ListEnvironmentDashboardsRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListEnvironmentDashboardsRequest
	GetRegionId() *string
	SetScene(v string) *ListEnvironmentDashboardsRequest
	GetScene() *string
}

type ListEnvironmentDashboardsRequest struct {
	// Name of Addon,One of AddonName and Scene must be filled in.
	//
	// example:
	//
	// trace-java
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scenario of Addon. Either AddonName or Scene is required.
	//
	// example:
	//
	// database
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListEnvironmentDashboardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentDashboardsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDashboardsRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListEnvironmentDashboardsRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentDashboardsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentDashboardsRequest) GetScene() *string {
	return s.Scene
}

func (s *ListEnvironmentDashboardsRequest) SetAddonName(v string) *ListEnvironmentDashboardsRequest {
	s.AddonName = &v
	return s
}

func (s *ListEnvironmentDashboardsRequest) SetEnvironmentId(v string) *ListEnvironmentDashboardsRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentDashboardsRequest) SetRegionId(v string) *ListEnvironmentDashboardsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentDashboardsRequest) SetScene(v string) *ListEnvironmentDashboardsRequest {
	s.Scene = &v
	return s
}

func (s *ListEnvironmentDashboardsRequest) Validate() error {
	return dara.Validate(s)
}
