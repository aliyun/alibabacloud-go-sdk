// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowTopUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeFlowTopUrlRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeFlowTopUrlRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeFlowTopUrlRequest
	GetRegionId() *string
	SetResource(v string) *DescribeFlowTopUrlRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeFlowTopUrlRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribeFlowTopUrlRequest
	GetStartTimestamp() *string
}

type DescribeFlowTopUrlRequest struct {
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
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
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

func (s DescribeFlowTopUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeFlowTopUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFlowTopUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowTopUrlRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeFlowTopUrlRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeFlowTopUrlRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeFlowTopUrlRequest) SetEndTimestamp(v string) *DescribeFlowTopUrlRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetInstanceId(v string) *DescribeFlowTopUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetRegionId(v string) *DescribeFlowTopUrlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetResource(v string) *DescribeFlowTopUrlRequest {
	s.Resource = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetResourceManagerResourceGroupId(v string) *DescribeFlowTopUrlRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetStartTimestamp(v string) *DescribeFlowTopUrlRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) Validate() error {
	return dara.Validate(s)
}
