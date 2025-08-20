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
}

type DescribeClusterNetInfoRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9dqvn0o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine of the cluster. Valid values:
	//
	// 	- **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// 	- **Clickhouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
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

func (s *DescribeClusterNetInfoRequest) SetDBClusterId(v string) *DescribeClusterNetInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterNetInfoRequest) SetEngine(v string) *DescribeClusterNetInfoRequest {
	s.Engine = &v
	return s
}

func (s *DescribeClusterNetInfoRequest) Validate() error {
	return dara.Validate(s)
}
