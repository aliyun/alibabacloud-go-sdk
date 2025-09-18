// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNodeGroupRequest
	GetClusterId() *string
	SetNodeGroup(v *NodeGroupConfig) *CreateNodeGroupRequest
	GetNodeGroup() *NodeGroupConfig
	SetRegionId(v string) *CreateNodeGroupRequest
	GetRegionId() *string
}

type CreateNodeGroupRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-E525E04F3914****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the node group.
	NodeGroup *NodeGroupConfig `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNodeGroupRequest) GetNodeGroup() *NodeGroupConfig {
	return s.NodeGroup
}

func (s *CreateNodeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNodeGroupRequest) SetClusterId(v string) *CreateNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeGroupRequest) SetNodeGroup(v *NodeGroupConfig) *CreateNodeGroupRequest {
	s.NodeGroup = v
	return s
}

func (s *CreateNodeGroupRequest) SetRegionId(v string) *CreateNodeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
