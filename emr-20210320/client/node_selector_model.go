// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeSelector interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *NodeSelector
	GetNodeGroupId() *string
	SetNodeGroupIds(v []*string) *NodeSelector
	GetNodeGroupIds() []*string
	SetNodeGroupName(v string) *NodeSelector
	GetNodeGroupName() *string
	SetNodeGroupNames(v []*string) *NodeSelector
	GetNodeGroupNames() []*string
	SetNodeGroupTypes(v []*string) *NodeSelector
	GetNodeGroupTypes() []*string
	SetNodeNames(v []*string) *NodeSelector
	GetNodeNames() []*string
	SetNodeSelectType(v string) *NodeSelector
	GetNodeSelectType() *string
}

type NodeSelector struct {
	// Deprecated
	//
	// > This parameter is deprecated. Use `NodeGroupIds` instead.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The IDs of the node groups to select.
	NodeGroupIds []*string `json:"NodeGroupIds,omitempty" xml:"NodeGroupIds,omitempty" type:"Repeated"`
	// Deprecated
	//
	// > This parameter is deprecated. Use `NodeGroupNames` instead.
	//
	// example:
	//
	// master-1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The names of the node groups to select.
	NodeGroupNames []*string `json:"NodeGroupNames,omitempty" xml:"NodeGroupNames,omitempty" type:"Repeated"`
	// The types of node groups to select. This parameter applies only when `NodeSelectType` is set to `NODE_GROUP` and `NodeGroupIds` is not specified. The array can contain up to 10 elements.
	//
	// example:
	//
	// ["CORE","TASK"]
	NodeGroupTypes []*string `json:"NodeGroupTypes,omitempty" xml:"NodeGroupTypes,omitempty" type:"Repeated"`
	// The names of the nodes to select. This parameter applies only when `NodeSelectType` is set to `NODE`.
	//
	// example:
	//
	// ["core1-1"]
	NodeNames []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	// The node selection type. Valid values:
	//
	// - `CLUSTER`: Selects all nodes in the cluster.
	//
	// - `NODE_GROUP`: Selects nodes by their node group.
	//
	// - `NODE`: Selects specific nodes by name.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLUSTER
	NodeSelectType *string `json:"NodeSelectType,omitempty" xml:"NodeSelectType,omitempty"`
}

func (s NodeSelector) String() string {
	return dara.Prettify(s)
}

func (s NodeSelector) GoString() string {
	return s.String()
}

func (s *NodeSelector) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *NodeSelector) GetNodeGroupIds() []*string {
	return s.NodeGroupIds
}

func (s *NodeSelector) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *NodeSelector) GetNodeGroupNames() []*string {
	return s.NodeGroupNames
}

func (s *NodeSelector) GetNodeGroupTypes() []*string {
	return s.NodeGroupTypes
}

func (s *NodeSelector) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *NodeSelector) GetNodeSelectType() *string {
	return s.NodeSelectType
}

func (s *NodeSelector) SetNodeGroupId(v string) *NodeSelector {
	s.NodeGroupId = &v
	return s
}

func (s *NodeSelector) SetNodeGroupIds(v []*string) *NodeSelector {
	s.NodeGroupIds = v
	return s
}

func (s *NodeSelector) SetNodeGroupName(v string) *NodeSelector {
	s.NodeGroupName = &v
	return s
}

func (s *NodeSelector) SetNodeGroupNames(v []*string) *NodeSelector {
	s.NodeGroupNames = v
	return s
}

func (s *NodeSelector) SetNodeGroupTypes(v []*string) *NodeSelector {
	s.NodeGroupTypes = v
	return s
}

func (s *NodeSelector) SetNodeNames(v []*string) *NodeSelector {
	s.NodeNames = v
	return s
}

func (s *NodeSelector) SetNodeSelectType(v string) *NodeSelector {
	s.NodeSelectType = &v
	return s
}

func (s *NodeSelector) Validate() error {
	return dara.Validate(s)
}
