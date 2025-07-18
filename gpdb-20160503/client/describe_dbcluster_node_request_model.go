// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBClusterNodeRequest
	GetDBInstanceId() *string
	SetNodeType(v string) *DescribeDBClusterNodeRequest
	GetNodeType() *string
}

type DescribeDBClusterNodeRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query details about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node type. Valid values:
	//
	// 	- **master**: coordinator node.
	//
	// 	- **segment**: compute node.
	//
	// > If you do not specify this parameter, the information about all nodes is returned.
	//
	// example:
	//
	// master
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeDBClusterNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBClusterNodeRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBClusterNodeRequest) SetDBInstanceId(v string) *DescribeDBClusterNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterNodeRequest) SetNodeType(v string) *DescribeDBClusterNodeRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeDBClusterNodeRequest) Validate() error {
	return dara.Validate(s)
}
