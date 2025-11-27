// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCClusterNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *DeleteRCClusterNodesShrinkRequest
	GetInstanceIdsShrink() *string
	SetNodesShrink(v string) *DeleteRCClusterNodesShrinkRequest
	GetNodesShrink() *string
	SetRegionId(v string) *DeleteRCClusterNodesShrinkRequest
	GetRegionId() *string
	SetVpcId(v string) *DeleteRCClusterNodesShrinkRequest
	GetVpcId() *string
}

type DeleteRCClusterNodesShrinkRequest struct {
	// The instance IDs.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The node information.
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
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

func (s DeleteRCClusterNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCClusterNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCClusterNodesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DeleteRCClusterNodesShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *DeleteRCClusterNodesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCClusterNodesShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteRCClusterNodesShrinkRequest) SetInstanceIdsShrink(v string) *DeleteRCClusterNodesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DeleteRCClusterNodesShrinkRequest) SetNodesShrink(v string) *DeleteRCClusterNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *DeleteRCClusterNodesShrinkRequest) SetRegionId(v string) *DeleteRCClusterNodesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCClusterNodesShrinkRequest) SetVpcId(v string) *DeleteRCClusterNodesShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteRCClusterNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
