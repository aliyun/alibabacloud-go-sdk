// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *DescribeUserTrafficRequest
	GetEndTimestamp() *int64
	SetInstanceId(v string) *DescribeUserTrafficRequest
	GetInstanceId() *string
	SetInterval(v int64) *DescribeUserTrafficRequest
	GetInterval() *int64
	SetRegionId(v string) *DescribeUserTrafficRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserTrafficRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v int64) *DescribeUserTrafficRequest
	GetStartTimestamp() *int64
	SetType(v string) *DescribeUserTrafficRequest
	GetType() *string
}

type DescribeUserTrafficRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 1665386280
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time interval. Unit: seconds.
	//
	// example:
	//
	// 3600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
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
	// rg-ae*******i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The type of real-time user traffic. Valid values:
	//
	// - bot: the number of bot management requests.
	//
	// - risk: the number of times risk identification is triggered.
	//
	// - custom_acl_captcha: the number of times the slider action of custom rules is triggered.
	//
	// - qps: the peak QPS.
	//
	// - apisec: the number of API security requests.
	//
	// - alb: the number of requests connected through ALB.
	//
	// - mse: the number of requests connected through MSE.
	//
	// - fc: the number of requests connected through Function Compute.
	//
	// - sae: the number of requests connected through Serverless App Engine.
	//
	// - apig: the number of requests connected through Cloud Native API Gateway.
	//
	// - nlb: the number of requests connected through NLB.
	//
	// example:
	//
	// qps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUserTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserTrafficRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeUserTrafficRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserTrafficRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeUserTrafficRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserTrafficRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserTrafficRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeUserTrafficRequest) GetType() *string {
	return s.Type
}

func (s *DescribeUserTrafficRequest) SetEndTimestamp(v int64) *DescribeUserTrafficRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetInstanceId(v string) *DescribeUserTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetInterval(v int64) *DescribeUserTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetRegionId(v string) *DescribeUserTrafficRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserTrafficRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetStartTimestamp(v int64) *DescribeUserTrafficRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetType(v string) *DescribeUserTrafficRequest {
	s.Type = &v
	return s
}

func (s *DescribeUserTrafficRequest) Validate() error {
	return dara.Validate(s)
}
