// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAbnormalTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeUserAbnormalTypeRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeUserAbnormalTypeRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeUserAbnormalTypeRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserAbnormalTypeRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserAbnormalTypeRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeUserAbnormalTypeRequest
	GetStartTime() *int64
}

type DescribeUserAbnormalTypeRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 993
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1726113600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-g4***201
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-ac***lani
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1723435200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserAbnormalTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAbnormalTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAbnormalTypeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeUserAbnormalTypeRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeUserAbnormalTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserAbnormalTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserAbnormalTypeRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserAbnormalTypeRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeUserAbnormalTypeRequest) SetClusterId(v string) *DescribeUserAbnormalTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeUserAbnormalTypeRequest) SetEndTime(v int64) *DescribeUserAbnormalTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserAbnormalTypeRequest) SetInstanceId(v string) *DescribeUserAbnormalTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserAbnormalTypeRequest) SetRegionId(v string) *DescribeUserAbnormalTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserAbnormalTypeRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserAbnormalTypeRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserAbnormalTypeRequest) SetStartTime(v int64) *DescribeUserAbnormalTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUserAbnormalTypeRequest) Validate() error {
	return dara.Validate(s)
}
