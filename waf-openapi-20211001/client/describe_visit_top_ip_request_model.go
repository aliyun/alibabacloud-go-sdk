// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVisitTopIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeVisitTopIpRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeVisitTopIpRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeVisitTopIpRequest
	GetRegionId() *string
	SetResource(v string) *DescribeVisitTopIpRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeVisitTopIpRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribeVisitTopIpRequest
	GetStartTimestamp() *string
}

type DescribeVisitTopIpRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386280
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

func (s DescribeVisitTopIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVisitTopIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeVisitTopIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVisitTopIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVisitTopIpRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeVisitTopIpRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeVisitTopIpRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeVisitTopIpRequest) SetEndTimestamp(v string) *DescribeVisitTopIpRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetInstanceId(v string) *DescribeVisitTopIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetRegionId(v string) *DescribeVisitTopIpRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetResource(v string) *DescribeVisitTopIpRequest {
	s.Resource = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetResourceManagerResourceGroupId(v string) *DescribeVisitTopIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetStartTimestamp(v string) *DescribeVisitTopIpRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeVisitTopIpRequest) Validate() error {
	return dara.Validate(s)
}
