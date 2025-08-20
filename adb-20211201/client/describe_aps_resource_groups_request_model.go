// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeApsResourceGroupsRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeApsResourceGroupsRequest
	GetRegionId() *string
	SetWorkloadId(v string) *DescribeApsResourceGroupsRequest
	GetWorkloadId() *string
}

type DescribeApsResourceGroupsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1t6rym21****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the data synchronization job.
	//
	// example:
	//
	// aps-hz1686v37sx****
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
}

func (s DescribeApsResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsResourceGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsResourceGroupsRequest) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *DescribeApsResourceGroupsRequest) SetDBClusterId(v string) *DescribeApsResourceGroupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsResourceGroupsRequest) SetRegionId(v string) *DescribeApsResourceGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsResourceGroupsRequest) SetWorkloadId(v string) *DescribeApsResourceGroupsRequest {
	s.WorkloadId = &v
	return s
}

func (s *DescribeApsResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
