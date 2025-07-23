// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterSubnetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ConfigClusterSubnetRequest
	GetClusterId() *string
	SetRegionId(v string) *ConfigClusterSubnetRequest
	GetRegionId() *string
	SetVSwitchIds(v []*string) *ConfigClusterSubnetRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *ConfigClusterSubnetRequest
	GetVpcId() *string
}

type ConfigClusterSubnetRequest struct {
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
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-7xvkh90cw39p0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ConfigClusterSubnetRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterSubnetRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ConfigClusterSubnetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigClusterSubnetRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ConfigClusterSubnetRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ConfigClusterSubnetRequest) SetClusterId(v string) *ConfigClusterSubnetRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterSubnetRequest) SetRegionId(v string) *ConfigClusterSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigClusterSubnetRequest) SetVSwitchIds(v []*string) *ConfigClusterSubnetRequest {
	s.VSwitchIds = v
	return s
}

func (s *ConfigClusterSubnetRequest) SetVpcId(v string) *ConfigClusterSubnetRequest {
	s.VpcId = &v
	return s
}

func (s *ConfigClusterSubnetRequest) Validate() error {
	return dara.Validate(s)
}
