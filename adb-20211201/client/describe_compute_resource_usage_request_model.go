// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeComputeResourceUsageRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeComputeResourceUsageRequest
	GetEndTime() *string
	SetResourceGroupName(v string) *DescribeComputeResourceUsageRequest
	GetResourceGroupName() *string
	SetStartTime(v string) *DescribeComputeResourceUsageRequest
	GetStartTime() *string
}

type DescribeComputeResourceUsageRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-05T03:45:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-04T03:45:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeComputeResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeComputeResourceUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeComputeResourceUsageRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeComputeResourceUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeComputeResourceUsageRequest) SetDBClusterId(v string) *DescribeComputeResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeComputeResourceUsageRequest) SetEndTime(v string) *DescribeComputeResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeComputeResourceUsageRequest) SetResourceGroupName(v string) *DescribeComputeResourceUsageRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeComputeResourceUsageRequest) SetStartTime(v string) *DescribeComputeResourceUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeComputeResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
