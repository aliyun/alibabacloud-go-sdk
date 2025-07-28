// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeResourceLogStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeResourceLogStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeResourceLogStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetResources(v string) *DescribeResourceLogStatusRequest
	GetResources() *string
}

type DescribeResourceLogStatusRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11zcl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The protected object that you want to query. You can specify multiple protected objects. Separate the protected objects with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-wewbb23dfsetetcic1242-0****,alb-wewbb23dfsetetcic1242-1****
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s DescribeResourceLogStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResourceLogStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceLogStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeResourceLogStatusRequest) GetResources() *string {
	return s.Resources
}

func (s *DescribeResourceLogStatusRequest) SetInstanceId(v string) *DescribeResourceLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceLogStatusRequest) SetRegionId(v string) *DescribeResourceLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceLogStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceLogStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResourceLogStatusRequest) SetResources(v string) *DescribeResourceLogStatusRequest {
	s.Resources = &v
	return s
}

func (s *DescribeResourceLogStatusRequest) Validate() error {
	return dara.Validate(s)
}
