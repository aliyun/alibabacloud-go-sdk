// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDefenseResourceRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDefenseResourceRequest
	GetRegionId() *string
	SetResource(v string) *DescribeDefenseResourceRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDefenseResourceRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-4xl****i60i
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: The Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object that you want to query. Only exact queries are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseResourceRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourceRequest) SetInstanceId(v string) *DescribeDefenseResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceRequest) SetRegionId(v string) *DescribeDefenseResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceRequest) SetResource(v string) *DescribeDefenseResourceRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourceRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceRequest) Validate() error {
	return dara.Validate(s)
}
