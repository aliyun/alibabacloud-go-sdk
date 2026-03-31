// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVisitUasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeVisitUasRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeVisitUasRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeVisitUasRequest
	GetRegionId() *string
	SetResource(v string) *DescribeVisitUasRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeVisitUasRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribeVisitUasRequest
	GetStartTimestamp() *string
}

type DescribeVisitUasRequest struct {
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
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
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

func (s DescribeVisitUasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVisitUasRequest) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeVisitUasRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVisitUasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVisitUasRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeVisitUasRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeVisitUasRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeVisitUasRequest) SetEndTimestamp(v string) *DescribeVisitUasRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeVisitUasRequest) SetInstanceId(v string) *DescribeVisitUasRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVisitUasRequest) SetRegionId(v string) *DescribeVisitUasRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVisitUasRequest) SetResource(v string) *DescribeVisitUasRequest {
	s.Resource = &v
	return s
}

func (s *DescribeVisitUasRequest) SetResourceManagerResourceGroupId(v string) *DescribeVisitUasRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeVisitUasRequest) SetStartTimestamp(v string) *DescribeVisitUasRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeVisitUasRequest) Validate() error {
	return dara.Validate(s)
}
