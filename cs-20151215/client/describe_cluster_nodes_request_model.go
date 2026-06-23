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
	SetNodeIps(v string) *DescribeClusterNodesRequest
	GetNodeIps() *string
	SetNodeLabels(v string) *DescribeClusterNodesRequest
	GetNodeLabels() *string
	SetNodeNames(v string) *DescribeClusterNodesRequest
	GetNodeNames() *string
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
	// The instance IDs of nodes. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// "i-bp11xjhwkj8k966u****,i-bp1dmhc2bu5igkyq****"
	InstanceIds *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty"`
	// example:
	//
	// 192.168.0.1
	NodeIps *string `json:"nodeIps,omitempty" xml:"nodeIps,omitempty"`
	// example:
	//
	// nodeLabels=app=nginx,env=prod
	NodeLabels *string `json:"nodeLabels,omitempty" xml:"nodeLabels,omitempty"`
	// example:
	//
	// cn-hangzhou.192.168.0.1
	NodeNames *string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// npe25633140a7d4fbea56cd0479c******
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The page number of the current query.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The maximum number of records that can be displayed on each page. Valid values: [1, 100].
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The status of cluster nodes. Used to filter by node running status. Valid values:
	//
	// - `all`: does not filter by running status. All nodes are returned.
	//
	// - `running`: running nodes.
	//
	// - `removing`: nodes that are being removed.
	//
	// - `initial`: nodes that are being initialized.
	//
	// - `failed`: nodes that failed to be created.
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

func (s *DescribeClusterNodesRequest) GetNodeIps() *string {
	return s.NodeIps
}

func (s *DescribeClusterNodesRequest) GetNodeLabels() *string {
	return s.NodeLabels
}

func (s *DescribeClusterNodesRequest) GetNodeNames() *string {
	return s.NodeNames
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

func (s *DescribeClusterNodesRequest) SetNodeIps(v string) *DescribeClusterNodesRequest {
	s.NodeIps = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetNodeLabels(v string) *DescribeClusterNodesRequest {
	s.NodeLabels = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetNodeNames(v string) *DescribeClusterNodesRequest {
	s.NodeNames = &v
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
