// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowTopResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeFlowTopResourceRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeFlowTopResourceRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeFlowTopResourceRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeFlowTopResourceRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribeFlowTopResourceRequest
	GetStartTimestamp() *string
}

type DescribeFlowTopResourceRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386340
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFlowTopResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowTopResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeFlowTopResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFlowTopResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowTopResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeFlowTopResourceRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeFlowTopResourceRequest) SetEndTimestamp(v string) *DescribeFlowTopResourceRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetInstanceId(v string) *DescribeFlowTopResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetRegionId(v string) *DescribeFlowTopResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetResourceManagerResourceGroupId(v string) *DescribeFlowTopResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetStartTimestamp(v string) *DescribeFlowTopResourceRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) Validate() error {
	return dara.Validate(s)
}
