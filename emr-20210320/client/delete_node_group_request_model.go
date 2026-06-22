// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteNodeGroupRequest
	GetClusterId() *string
	SetDescription(v string) *DeleteNodeGroupRequest
	GetDescription() *string
	SetNodeGroupId(v string) *DeleteNodeGroupRequest
	GetNodeGroupId() *string
	SetRegionId(v string) *DeleteNodeGroupRequest
	GetRegionId() *string
}

type DeleteNodeGroupRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The description.
	//
	// example:
	//
	// 机器组不再需要
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The node group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The region ID. You can call [ListRegions](url) to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteNodeGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *DeleteNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DeleteNodeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNodeGroupRequest) SetClusterId(v string) *DeleteNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodeGroupRequest) SetDescription(v string) *DeleteNodeGroupRequest {
	s.Description = &v
	return s
}

func (s *DeleteNodeGroupRequest) SetNodeGroupId(v string) *DeleteNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DeleteNodeGroupRequest) SetRegionId(v string) *DeleteNodeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
