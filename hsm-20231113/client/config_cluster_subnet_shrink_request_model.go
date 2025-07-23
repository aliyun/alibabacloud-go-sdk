// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterSubnetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ConfigClusterSubnetShrinkRequest
	GetClusterId() *string
	SetRegionId(v string) *ConfigClusterSubnetShrinkRequest
	GetRegionId() *string
	SetVSwitchIdsShrink(v string) *ConfigClusterSubnetShrinkRequest
	GetVSwitchIdsShrink() *string
	SetVpcId(v string) *ConfigClusterSubnetShrinkRequest
	GetVpcId() *string
}

type ConfigClusterSubnetShrinkRequest struct {
	// The ID of the cluster. You can call the ListCluster operation to obtain cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsg****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// A list of vSwitches that are associated with the cluster. Note: You must include all vSwitches that you want to associate with the cluster.
	//
	// This parameter is required.
	VSwitchIdsShrink *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-7xvkh90cw39p0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ConfigClusterSubnetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterSubnetShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ConfigClusterSubnetShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigClusterSubnetShrinkRequest) GetVSwitchIdsShrink() *string {
	return s.VSwitchIdsShrink
}

func (s *ConfigClusterSubnetShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ConfigClusterSubnetShrinkRequest) SetClusterId(v string) *ConfigClusterSubnetShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetRegionId(v string) *ConfigClusterSubnetShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetVSwitchIdsShrink(v string) *ConfigClusterSubnetShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetVpcId(v string) *ConfigClusterSubnetShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
