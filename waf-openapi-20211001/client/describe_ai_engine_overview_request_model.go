// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeAiEngineOverviewRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeAiEngineOverviewRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAiEngineOverviewRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineOverviewRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeAiEngineOverviewRequest
	GetStartTime() *int64
}

type DescribeAiEngineOverviewRequest struct {
	// The end time. The time is a UNIX timestamp in UTC. Unit: seconds. Data is filtered based on the time when the event was last detected. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792973400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-dss****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-demo****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The start time. The time is a UNIX timestamp in UTC. Unit: seconds. Data is filtered based on the time when the event was last detected. You can query data within the last 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792973400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAiEngineOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineOverviewRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAiEngineOverviewRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineOverviewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineOverviewRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineOverviewRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAiEngineOverviewRequest) SetEndTime(v int64) *DescribeAiEngineOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAiEngineOverviewRequest) SetInstanceId(v string) *DescribeAiEngineOverviewRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineOverviewRequest) SetRegionId(v string) *DescribeAiEngineOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineOverviewRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineOverviewRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineOverviewRequest) SetStartTime(v int64) *DescribeAiEngineOverviewRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAiEngineOverviewRequest) Validate() error {
	return dara.Validate(s)
}
