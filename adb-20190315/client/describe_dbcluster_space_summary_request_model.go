// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSpaceSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterSpaceSummaryRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeDBClusterSpaceSummaryRequest
	GetRegionId() *string
}

type DescribeDBClusterSpaceSummaryRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1u8c0mgfg58****
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

func (s DescribeDBClusterSpaceSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterSpaceSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetDBClusterId(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetRegionId(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) Validate() error {
	return dara.Validate(s)
}
