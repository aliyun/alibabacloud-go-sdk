// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePauseProtectionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePauseProtectionStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribePauseProtectionStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribePauseProtectionStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribePauseProtectionStatusRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
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

func (s DescribePauseProtectionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePauseProtectionStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePauseProtectionStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePauseProtectionStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePauseProtectionStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribePauseProtectionStatusRequest) SetInstanceId(v string) *DescribePauseProtectionStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePauseProtectionStatusRequest) SetRegionId(v string) *DescribePauseProtectionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePauseProtectionStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribePauseProtectionStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribePauseProtectionStatusRequest) Validate() error {
	return dara.Validate(s)
}
