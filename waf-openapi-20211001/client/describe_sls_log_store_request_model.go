// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSlsLogStoreRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSlsLogStoreRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSlsLogStoreRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSlsLogStoreRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-zpr3d****0o
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
	// rg-aek2wf3mn****vq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSlsLogStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogStoreRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSlsLogStoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlsLogStoreRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSlsLogStoreRequest) SetInstanceId(v string) *DescribeSlsLogStoreRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlsLogStoreRequest) SetRegionId(v string) *DescribeSlsLogStoreRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlsLogStoreRequest) SetResourceManagerResourceGroupId(v string) *DescribeSlsLogStoreRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSlsLogStoreRequest) Validate() error {
	return dara.Validate(s)
}
