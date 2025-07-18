// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataShareInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDataShareInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDataShareInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataShareInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDataShareInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDataShareInstancesRequest
	GetResourceGroupId() *string
	SetSearchValue(v string) *DescribeDataShareInstancesRequest
	GetSearchValue() *string
}

type DescribeDataShareInstancesRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The keyword used to filter instances, which can be an instance ID or instance description.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs and instance descriptions.
	//
	// example:
	//
	// gp-bp***************
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s DescribeDataShareInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataShareInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDataShareInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataShareInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataShareInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataShareInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDataShareInstancesRequest) GetSearchValue() *string {
	return s.SearchValue
}

func (s *DescribeDataShareInstancesRequest) SetOwnerId(v int64) *DescribeDataShareInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetPageNumber(v int32) *DescribeDataShareInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetPageSize(v int32) *DescribeDataShareInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetRegionId(v string) *DescribeDataShareInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetResourceGroupId(v string) *DescribeDataShareInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetSearchValue(v string) *DescribeDataShareInstancesRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) Validate() error {
	return dara.Validate(s)
}
