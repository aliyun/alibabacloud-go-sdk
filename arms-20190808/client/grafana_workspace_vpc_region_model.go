// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceVpcRegion interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GrafanaWorkspaceVpcRegion
	GetRegionId() *string
	SetRegionName(v string) *GrafanaWorkspaceVpcRegion
	GetRegionName() *string
}

type GrafanaWorkspaceVpcRegion struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 华北2（北京）
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
}

func (s GrafanaWorkspaceVpcRegion) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceVpcRegion) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceVpcRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *GrafanaWorkspaceVpcRegion) GetRegionName() *string {
	return s.RegionName
}

func (s *GrafanaWorkspaceVpcRegion) SetRegionId(v string) *GrafanaWorkspaceVpcRegion {
	s.RegionId = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegion) SetRegionName(v string) *GrafanaWorkspaceVpcRegion {
	s.RegionName = &v
	return s
}

func (s *GrafanaWorkspaceVpcRegion) Validate() error {
	return dara.Validate(s)
}
