// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDDoSStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDDoSStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDDoSStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDDoSStatusRequest struct {
	// The ID of the WAF instance that you want to query.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
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
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDDoSStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDDoSStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDDoSStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDDoSStatusRequest) SetInstanceId(v string) *DescribeDDoSStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDDoSStatusRequest) SetRegionId(v string) *DescribeDDoSStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDDoSStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeDDoSStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDDoSStatusRequest) Validate() error {
	return dara.Validate(s)
}
