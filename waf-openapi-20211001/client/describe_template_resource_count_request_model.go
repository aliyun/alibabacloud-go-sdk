// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateResourceCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTemplateResourceCountRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeTemplateResourceCountRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeTemplateResourceCountRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateIds(v string) *DescribeTemplateResourceCountRequest
	GetTemplateIds() *string
}

type DescribeTemplateResourceCountRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-lbj****gx08
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
	// rg-aekzwwkpn****5i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The IDs of the protection templates that you want to query. Separate multiple template IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345,12346
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
}

func (s DescribeTemplateResourceCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResourceCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourceCountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTemplateResourceCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTemplateResourceCountRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeTemplateResourceCountRequest) GetTemplateIds() *string {
	return s.TemplateIds
}

func (s *DescribeTemplateResourceCountRequest) SetInstanceId(v string) *DescribeTemplateResourceCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTemplateResourceCountRequest) SetRegionId(v string) *DescribeTemplateResourceCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTemplateResourceCountRequest) SetResourceManagerResourceGroupId(v string) *DescribeTemplateResourceCountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeTemplateResourceCountRequest) SetTemplateIds(v string) *DescribeTemplateResourceCountRequest {
	s.TemplateIds = &v
	return s
}

func (s *DescribeTemplateResourceCountRequest) Validate() error {
	return dara.Validate(s)
}
