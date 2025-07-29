// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DescribeClusterNodesRequest
	GetInstanceIds() *string
	SetNodepoolId(v string) *DescribeClusterNodesRequest
	GetNodepoolId() *string
	SetPageNumber(v string) *DescribeClusterNodesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeClusterNodesRequest
	GetPageSize() *string
	SetState(v string) *DescribeClusterNodesRequest
	GetState() *string
}

type DescribeClusterNodesRequest struct {
	// The IDs of the nodes that you want to query. Separate multiple node IDs with commas (,).
	//
	// example:
	//
	// "i-bp11xjhwkj8k966u****,i-bp1dmhc2bu5igkyq****"
	InstanceIds *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// np****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The node state that you want to use to filter nodes. Valid values:
	//
	// 	- `all`: query nodes in the following four states.
	//
	// 	- `running`: query nodes in the running state.
	//
	// 	- `removing`: query nodes that are being removed.
	//
	// 	- `initial`: query nodes that are being initialized.
	//
	// 	- `failed`: query nodes that fail to be created.
	//
	// Default value: `all`.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s DescribeClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeClusterNodesRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterNodesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeClusterNodesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeClusterNodesRequest) GetState() *string {
	return s.State
}

func (s *DescribeClusterNodesRequest) SetInstanceIds(v string) *DescribeClusterNodesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetNodepoolId(v string) *DescribeClusterNodesRequest {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetPageNumber(v string) *DescribeClusterNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetPageSize(v string) *DescribeClusterNodesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetState(v string) *DescribeClusterNodesRequest {
	s.State = &v
	return s
}

func (s *DescribeClusterNodesRequest) Validate() error {
	return dara.Validate(s)
}
