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
	// The list of node instance IDs. Separate multiple values with commas (,).
	//
	// example:
	//
	// "i-bp11xjhwkj8k966u****,i-bp1dmhc2bu5igkyq****"
	InstanceIds *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty"`
	// The list of node IP addresses used for filtering. Separate multiple values with commas (,).
	//
	// example:
	//
	// 192.168.0.1
	NodeIps *string `json:"nodeIps,omitempty" xml:"nodeIps,omitempty"`
	// The node labels used for filtering. Specify the labels as key=value pairs separated by commas (,). Multiple labels are evaluated with the AND operator. Only nodes that match all specified labels are returned.
	//
	// example:
	//
	// nodeLabels=app=nginx,env=prod
	NodeLabels *string `json:"nodeLabels,omitempty" xml:"nodeLabels,omitempty"`
	// The list of node names used for filtering. Separate multiple values with commas (,).
	//
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
	// The maximum number of records to return per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The running states of the cluster nodes used for filtering. Valid values:
	//
	// - `all`: does not filter nodes by running state and queries nodes in all states.
	//
	// - `running`: the cluster nodes that are running.
	//
	// - `removing`: the cluster nodes that are being removed.
	//
	// - `initial`: the cluster nodes that are being initialized.
	//
	// - `failed`: the cluster nodes that failed to be created.
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
