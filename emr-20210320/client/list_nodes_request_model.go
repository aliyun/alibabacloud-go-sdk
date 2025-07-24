// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodesRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListNodesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodesRequest
	GetNextToken() *string
	SetNodeGroupIds(v []*string) *ListNodesRequest
	GetNodeGroupIds() []*string
	SetNodeIds(v []*string) *ListNodesRequest
	GetNodeIds() []*string
	SetNodeNames(v []*string) *ListNodesRequest
	GetNodeNames() []*string
	SetNodeStates(v []*string) *ListNodesRequest
	GetNodeStates() []*string
	SetPrivateIps(v []*string) *ListNodesRequest
	GetPrivateIps() []*string
	SetPublicIps(v []*string) *ListNodesRequest
	GetPublicIps() []*string
	SetRegionId(v string) *ListNodesRequest
	GetRegionId() *string
	SetTags(v []*Tag) *ListNodesRequest
	GetTags() []*Tag
}

type ListNodesRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of maximum number of records to obtain at a time. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Marks the current position where reading starts. If you set this value to null, you can start from the beginning.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of node groups.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NodeGroupIds []*string `json:"NodeGroupIds,omitempty" xml:"NodeGroupIds,omitempty" type:"Repeated"`
	// An array that consists of information about the ID of the node.
	//
	// example:
	//
	// c-b933c5aac8fe****
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The names of the nodes.
	//
	// example:
	//
	// 20
	NodeNames []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	// The status of the node.
	//
	// example:
	//
	// ["CREATED"]
	NodeStates []*string `json:"NodeStates,omitempty" xml:"NodeStates,omitempty" type:"Repeated"`
	// The private IP address.
	//
	// example:
	//
	// ["172.12.0.91"]
	PrivateIps []*string `json:"PrivateIps,omitempty" xml:"PrivateIps,omitempty" type:"Repeated"`
	// The public IP address.
	//
	// example:
	//
	// ["120.13.14.38"]
	PublicIps []*string `json:"PublicIps,omitempty" xml:"PublicIps,omitempty" type:"Repeated"`
	// The ID of the region in which you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodesRequest) GetNodeGroupIds() []*string {
	return s.NodeGroupIds
}

func (s *ListNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *ListNodesRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *ListNodesRequest) GetNodeStates() []*string {
	return s.NodeStates
}

func (s *ListNodesRequest) GetPrivateIps() []*string {
	return s.PrivateIps
}

func (s *ListNodesRequest) GetPublicIps() []*string {
	return s.PublicIps
}

func (s *ListNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNodesRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *ListNodesRequest) SetClusterId(v string) *ListNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesRequest) SetMaxResults(v int32) *ListNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodesRequest) SetNextToken(v string) *ListNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodesRequest) SetNodeGroupIds(v []*string) *ListNodesRequest {
	s.NodeGroupIds = v
	return s
}

func (s *ListNodesRequest) SetNodeIds(v []*string) *ListNodesRequest {
	s.NodeIds = v
	return s
}

func (s *ListNodesRequest) SetNodeNames(v []*string) *ListNodesRequest {
	s.NodeNames = v
	return s
}

func (s *ListNodesRequest) SetNodeStates(v []*string) *ListNodesRequest {
	s.NodeStates = v
	return s
}

func (s *ListNodesRequest) SetPrivateIps(v []*string) *ListNodesRequest {
	s.PrivateIps = v
	return s
}

func (s *ListNodesRequest) SetPublicIps(v []*string) *ListNodesRequest {
	s.PublicIps = v
	return s
}

func (s *ListNodesRequest) SetRegionId(v string) *ListNodesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNodesRequest) SetTags(v []*Tag) *ListNodesRequest {
	s.Tags = v
	return s
}

func (s *ListNodesRequest) Validate() error {
	return dara.Validate(s)
}
