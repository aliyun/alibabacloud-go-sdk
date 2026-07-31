// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeClusterNetInfoRequest
	GetDBClusterId() *string
	SetEngine(v string) *DescribeClusterNetInfoRequest
	GetEngine() *string
	SetResourceGroupName(v string) *DescribeClusterNetInfoRequest
	GetResourceGroupName() *string
}

type DescribeClusterNetInfoRequest struct {
	// <props="china">The ID of an Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the details of clusters in a specific region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9dqvn0o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine. Valid values:
	//
	// - **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// - **ClickHouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine            *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DescribeClusterNetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeClusterNetInfoRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeClusterNetInfoRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeClusterNetInfoRequest) SetDBClusterId(v string) *DescribeClusterNetInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterNetInfoRequest) SetEngine(v string) *DescribeClusterNetInfoRequest {
	s.Engine = &v
	return s
}

func (s *DescribeClusterNetInfoRequest) SetResourceGroupName(v string) *DescribeClusterNetInfoRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeClusterNetInfoRequest) Validate() error {
	return dara.Validate(s)
}
