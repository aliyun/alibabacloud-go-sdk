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
	// The ID of the Data Lakehouse Edition cluster.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to view the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1t6rym21****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// > Call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to view the region ID of a cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the data synchronization task.
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
