// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBClusterPerformanceRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBClusterPerformanceRequest
	GetEndTime() *string
	SetKey(v string) *DescribeDBClusterPerformanceRequest
	GetKey() *string
	SetNodeType(v string) *DescribeDBClusterPerformanceRequest
	GetNodeType() *string
	SetNodes(v string) *DescribeDBClusterPerformanceRequest
	GetNodes() *string
	SetResourceGroupName(v string) *DescribeDBClusterPerformanceRequest
	GetResourceGroupName() *string
	SetStartTime(v string) *DescribeDBClusterPerformanceRequest
	GetStartTime() *string
}

type DescribeDBClusterPerformanceRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query details about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDTHH:mmZ` format.
	//
	// > The end time must be later than the start time. The maximum time range that can be specified is seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-03T15:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metric that you want to query. Separate multiple values with commas (,). For more information, see [Performance parameters](https://help.aliyun.com/document_detail/86943.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// adbpg_conn_count
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The node type. Valid values:
	//
	// 	- **master**: coordinator node.
	//
	// 	- **segment**: compute node.
	//
	// > If you do not specify this parameter, the performance metrics of all nodes are returned.
	//
	// example:
	//
	// master
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The nodes for which you want to query performance metrics. Separate multiple values with commas (,). Example: `master-10******1,master-10******2`. You can call the [DescribeDBClusterNode](https://help.aliyun.com/document_detail/390136.html) operation to query the names of nodes.
	//
	// You can also filter the nodes based on their metric values. Valid values:
	//
	// 	- **top10**: the 10 nodes that have the highest metric values.
	//
	// 	- **top20**: the 20 nodes that have the highest metric values.
	//
	// 	- **bottom10**: the 10 nodes that have the lowest metric values.
	//
	// 	- **bottom20**: the 20 nodes that have the lowest metric values.
	//
	// example:
	//
	// top10
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// testgroup
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDTHH:mmZ` format.
	//
	// > You can query monitoring information only within the last 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-03T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBClusterPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterPerformanceRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBClusterPerformanceRequest) GetNodes() *string {
	return s.Nodes
}

func (s *DescribeDBClusterPerformanceRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeDBClusterPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterPerformanceRequest) SetDBInstanceId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetNodeType(v string) *DescribeDBClusterPerformanceRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetNodes(v string) *DescribeDBClusterPerformanceRequest {
	s.Nodes = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourceGroupName(v string) *DescribeDBClusterPerformanceRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
