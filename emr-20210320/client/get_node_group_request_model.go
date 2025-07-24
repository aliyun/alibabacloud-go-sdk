// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetNodeGroupRequest
	GetClusterId() *string
	SetNodeGroupId(v string) *GetNodeGroupRequest
	GetNodeGroupId() *string
	SetRegionId(v string) *GetNodeGroupRequest
	GetRegionId() *string
}

type GetNodeGroupRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The ID of the region in which you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *GetNodeGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *GetNodeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNodeGroupRequest) SetClusterId(v string) *GetNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *GetNodeGroupRequest) SetNodeGroupId(v string) *GetNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *GetNodeGroupRequest) SetRegionId(v string) *GetNodeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
