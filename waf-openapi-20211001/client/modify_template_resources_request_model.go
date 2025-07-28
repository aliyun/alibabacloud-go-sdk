// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindResourceGroups(v []*string) *ModifyTemplateResourcesRequest
	GetBindResourceGroups() []*string
	SetBindResources(v []*string) *ModifyTemplateResourcesRequest
	GetBindResources() []*string
	SetInstanceId(v string) *ModifyTemplateResourcesRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyTemplateResourcesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyTemplateResourcesRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *ModifyTemplateResourcesRequest
	GetTemplateId() *int64
	SetUnbindResourceGroups(v []*string) *ModifyTemplateResourcesRequest
	GetUnbindResourceGroups() []*string
	SetUnbindResources(v []*string) *ModifyTemplateResourcesRequest
	GetUnbindResources() []*string
}

type ModifyTemplateResourcesRequest struct {
	// The protected object groups that you want to associate with the template. Specify the value in the [**"group1","group2",...**] format.
	BindResourceGroups []*string `json:"BindResourceGroups,omitempty" xml:"BindResourceGroups,omitempty" type:"Repeated"`
	// The protected objects that you want to associate with the template. Specify the value in the [**"XX1","XX2",...**] format.
	BindResources []*string `json:"BindResources,omitempty" xml:"BindResources,omitempty" type:"Repeated"`
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2291
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The protected object groups that you want to disassociate from the template. Specify the value in the [**"group1","group2",...**] format.
	UnbindResourceGroups []*string `json:"UnbindResourceGroups,omitempty" xml:"UnbindResourceGroups,omitempty" type:"Repeated"`
	// The protected objects that you want to disassociate from the template. Specify the value in the [**"XX1","XX2",...**] format.
	UnbindResources []*string `json:"UnbindResources,omitempty" xml:"UnbindResources,omitempty" type:"Repeated"`
}

func (s ModifyTemplateResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateResourcesRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesRequest) GetBindResourceGroups() []*string {
	return s.BindResourceGroups
}

func (s *ModifyTemplateResourcesRequest) GetBindResources() []*string {
	return s.BindResources
}

func (s *ModifyTemplateResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTemplateResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTemplateResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyTemplateResourcesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyTemplateResourcesRequest) GetUnbindResourceGroups() []*string {
	return s.UnbindResourceGroups
}

func (s *ModifyTemplateResourcesRequest) GetUnbindResources() []*string {
	return s.UnbindResources
}

func (s *ModifyTemplateResourcesRequest) SetBindResourceGroups(v []*string) *ModifyTemplateResourcesRequest {
	s.BindResourceGroups = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetBindResources(v []*string) *ModifyTemplateResourcesRequest {
	s.BindResources = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetInstanceId(v string) *ModifyTemplateResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetRegionId(v string) *ModifyTemplateResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetResourceManagerResourceGroupId(v string) *ModifyTemplateResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetTemplateId(v int64) *ModifyTemplateResourcesRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetUnbindResourceGroups(v []*string) *ModifyTemplateResourcesRequest {
	s.UnbindResourceGroups = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetUnbindResources(v []*string) *ModifyTemplateResourcesRequest {
	s.UnbindResources = v
	return s
}

func (s *ModifyTemplateResourcesRequest) Validate() error {
	return dara.Validate(s)
}
