// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudProtectableCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHybridCloudProtectableCountRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeHybridCloudProtectableCountRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudProtectableCountRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudProtectableCountRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-4xl3fdi4u01**fe23
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 阿里云资源组ID。
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudProtectableCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudProtectableCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudProtectableCountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudProtectableCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudProtectableCountRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudProtectableCountRequest) SetInstanceId(v string) *DescribeHybridCloudProtectableCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudProtectableCountRequest) SetRegionId(v string) *DescribeHybridCloudProtectableCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudProtectableCountRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudProtectableCountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudProtectableCountRequest) Validate() error {
	return dara.Validate(s)
}
