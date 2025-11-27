// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DeleteRCClusterNodesRequest
	GetInstanceIds() []*string
	SetNodes(v []*string) *DeleteRCClusterNodesRequest
	GetNodes() []*string
	SetRegionId(v string) *DeleteRCClusterNodesRequest
	GetRegionId() *string
	SetVpcId(v string) *DeleteRCClusterNodesRequest
	GetVpcId() *string
}

type DeleteRCClusterNodesRequest struct {
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The node information.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// >  This is a reserved parameter.
	//
	// example:
	//
	// None
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteRCClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCClusterNodesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DeleteRCClusterNodesRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *DeleteRCClusterNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCClusterNodesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteRCClusterNodesRequest) SetInstanceIds(v []*string) *DeleteRCClusterNodesRequest {
	s.InstanceIds = v
	return s
}

func (s *DeleteRCClusterNodesRequest) SetNodes(v []*string) *DeleteRCClusterNodesRequest {
	s.Nodes = v
	return s
}

func (s *DeleteRCClusterNodesRequest) SetRegionId(v string) *DeleteRCClusterNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCClusterNodesRequest) SetVpcId(v string) *DeleteRCClusterNodesRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteRCClusterNodesRequest) Validate() error {
	return dara.Validate(s)
}
