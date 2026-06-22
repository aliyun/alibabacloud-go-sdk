// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecreaseNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchInterval(v int32) *DecreaseNodesRequest
	GetBatchInterval() *int32
	SetBatchSize(v int32) *DecreaseNodesRequest
	GetBatchSize() *int32
	SetClusterId(v string) *DecreaseNodesRequest
	GetClusterId() *string
	SetDecreaseNodeCount(v int32) *DecreaseNodesRequest
	GetDecreaseNodeCount() *int32
	SetNodeGroupId(v string) *DecreaseNodesRequest
	GetNodeGroupId() *string
	SetNodeIds(v []*string) *DecreaseNodesRequest
	GetNodeIds() []*string
	SetRegionId(v string) *DecreaseNodesRequest
	GetRegionId() *string
}

type DecreaseNodesRequest struct {
	// The cooldown time between batches.
	BatchInterval *int32 `json:"BatchInterval,omitempty" xml:"BatchInterval,omitempty"`
	// The number of nodes to concurrently take offline in a single batch.
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of nodes to remove. The value must be less than the number of active nodes in the node group.
	//
	// example:
	//
	// 3
	DecreaseNodeCount *int32 `json:"DecreaseNodeCount,omitempty" xml:"DecreaseNodeCount,omitempty"`
	// The node group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// An array of node IDs. The array can contain from 1 to 500 elements.
	//
	// example:
	//
	// ["i-bp1cudc25w2bfwl5****"]
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DecreaseNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DecreaseNodesRequest) GoString() string {
	return s.String()
}

func (s *DecreaseNodesRequest) GetBatchInterval() *int32 {
	return s.BatchInterval
}

func (s *DecreaseNodesRequest) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *DecreaseNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DecreaseNodesRequest) GetDecreaseNodeCount() *int32 {
	return s.DecreaseNodeCount
}

func (s *DecreaseNodesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DecreaseNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *DecreaseNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DecreaseNodesRequest) SetBatchInterval(v int32) *DecreaseNodesRequest {
	s.BatchInterval = &v
	return s
}

func (s *DecreaseNodesRequest) SetBatchSize(v int32) *DecreaseNodesRequest {
	s.BatchSize = &v
	return s
}

func (s *DecreaseNodesRequest) SetClusterId(v string) *DecreaseNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DecreaseNodesRequest) SetDecreaseNodeCount(v int32) *DecreaseNodesRequest {
	s.DecreaseNodeCount = &v
	return s
}

func (s *DecreaseNodesRequest) SetNodeGroupId(v string) *DecreaseNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DecreaseNodesRequest) SetNodeIds(v []*string) *DecreaseNodesRequest {
	s.NodeIds = v
	return s
}

func (s *DecreaseNodesRequest) SetRegionId(v string) *DecreaseNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DecreaseNodesRequest) Validate() error {
	return dara.Validate(s)
}
