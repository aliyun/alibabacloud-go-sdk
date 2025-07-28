// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAssetTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecAssetTrendRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeApisecAssetTrendRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeApisecAssetTrendRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeApisecAssetTrendRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecAssetTrendRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeApisecAssetTrendRequest
	GetStartTime() *int64
}

type DescribeApisecAssetTrendRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// example:
	//
	// 590
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1683183599
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-ww**b06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek**7uq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1668496310
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApisecAssetTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAssetTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecAssetTrendRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecAssetTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeApisecAssetTrendRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecAssetTrendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecAssetTrendRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecAssetTrendRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeApisecAssetTrendRequest) SetClusterId(v string) *DescribeApisecAssetTrendRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecAssetTrendRequest) SetEndTime(v int64) *DescribeApisecAssetTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApisecAssetTrendRequest) SetInstanceId(v string) *DescribeApisecAssetTrendRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecAssetTrendRequest) SetRegionId(v string) *DescribeApisecAssetTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecAssetTrendRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecAssetTrendRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecAssetTrendRequest) SetStartTime(v int64) *DescribeApisecAssetTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApisecAssetTrendRequest) Validate() error {
	return dara.Validate(s)
}
