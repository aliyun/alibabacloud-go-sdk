// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountDelegatedStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAccountDelegatedStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAccountDelegatedStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAccountDelegatedStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeAccountDelegatedStatusRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
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

func (s DescribeAccountDelegatedStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountDelegatedStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountDelegatedStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAccountDelegatedStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountDelegatedStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAccountDelegatedStatusRequest) SetInstanceId(v string) *DescribeAccountDelegatedStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccountDelegatedStatusRequest) SetRegionId(v string) *DescribeAccountDelegatedStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountDelegatedStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeAccountDelegatedStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAccountDelegatedStatusRequest) Validate() error {
	return dara.Validate(s)
}
