// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostGroupElasticStrategyParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostGroupName(v string) *DescribeHostGroupElasticStrategyParametersRequest
	GetDedicatedHostGroupName() *string
	SetRegionId(v string) *DescribeHostGroupElasticStrategyParametersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeHostGroupElasticStrategyParametersRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *DescribeHostGroupElasticStrategyParametersRequest
	GetResourceOwnerId() *int64
}

type DescribeHostGroupElasticStrategyParametersRequest struct {
	// The name of the dedicated cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// dhg-d0dwi82293b2w9t5
	DedicatedHostGroupName *string `json:"DedicatedHostGroupName,omitempty" xml:"DedicatedHostGroupName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHostGroupElasticStrategyParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostGroupElasticStrategyParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) GetDedicatedHostGroupName() *string {
	return s.DedicatedHostGroupName
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) SetDedicatedHostGroupName(v string) *DescribeHostGroupElasticStrategyParametersRequest {
	s.DedicatedHostGroupName = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) SetRegionId(v string) *DescribeHostGroupElasticStrategyParametersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) SetResourceGroupId(v string) *DescribeHostGroupElasticStrategyParametersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) SetResourceOwnerId(v int64) *DescribeHostGroupElasticStrategyParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersRequest) Validate() error {
	return dara.Validate(s)
}
