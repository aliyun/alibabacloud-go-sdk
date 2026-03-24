// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecStatisticsRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeApisecStatisticsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeApisecStatisticsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeApisecStatisticsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecStatisticsRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeApisecStatisticsRequest
	GetStartTime() *int64
	SetType(v string) *DescribeApisecStatisticsRequest
	GetType() *string
	SetUserStatusList(v []*string) *DescribeApisecStatisticsRequest
	GetUserStatusList() []*string
}

type DescribeApisecStatisticsRequest struct {
	// The ID of the Hybrid Cloud WAF cluster.
	//
	// > This parameter is available only in hybrid cloud scenarios. Call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query Hybrid Cloud WAF clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1726113600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-uax***b09
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2***uwbs5q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1668496310
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of detection statistics. Valid values:
	//
	// - **risk**: statistics on security risks.
	//
	// - **event**: statistics on attacks.
	//
	// example:
	//
	// risk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The list of statuses used to filter the detection statistics.
	UserStatusList []*string `json:"UserStatusList,omitempty" xml:"UserStatusList,omitempty" type:"Repeated"`
}

func (s DescribeApisecStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecStatisticsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecStatisticsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeApisecStatisticsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecStatisticsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecStatisticsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeApisecStatisticsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApisecStatisticsRequest) GetUserStatusList() []*string {
	return s.UserStatusList
}

func (s *DescribeApisecStatisticsRequest) SetClusterId(v string) *DescribeApisecStatisticsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetEndTime(v int64) *DescribeApisecStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetInstanceId(v string) *DescribeApisecStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetRegionId(v string) *DescribeApisecStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecStatisticsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetStartTime(v int64) *DescribeApisecStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetType(v string) *DescribeApisecStatisticsRequest {
	s.Type = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetUserStatusList(v []*string) *DescribeApisecStatisticsRequest {
	s.UserStatusList = v
	return s
}

func (s *DescribeApisecStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
