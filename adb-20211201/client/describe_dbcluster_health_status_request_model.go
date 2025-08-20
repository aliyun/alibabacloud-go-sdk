// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterHealthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterHealthStatusRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeDBClusterHealthStatusRequest
	GetRegionId() *string
}

type DescribeDBClusterHealthStatusRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6o6m8p6x7v****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterHealthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterHealthStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterHealthStatusRequest) SetDBClusterId(v string) *DescribeDBClusterHealthStatusRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterHealthStatusRequest) SetRegionId(v string) *DescribeDBClusterHealthStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterHealthStatusRequest) Validate() error {
	return dara.Validate(s)
}
