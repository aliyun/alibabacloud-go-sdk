// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindAssets(v []*string) *ModifyTemplateResourcesRequest
	GetBindAssets() []*string
	SetBindResourceGroups(v []*string) *ModifyTemplateResourcesRequest
	GetBindResourceGroups() []*string
	SetBindResources(v []*string) *ModifyTemplateResourcesRequest
	GetBindResources() []*string
	SetDryRun(v bool) *ModifyTemplateResourcesRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ModifyTemplateResourcesRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyTemplateResourcesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyTemplateResourcesRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *ModifyTemplateResourcesRequest
	GetTemplateId() *int64
	SetUnbindAssets(v []*string) *ModifyTemplateResourcesRequest
	GetUnbindAssets() []*string
	SetUnbindResourceGroups(v []*string) *ModifyTemplateResourcesRequest
	GetUnbindResourceGroups() []*string
	SetUnbindResources(v []*string) *ModifyTemplateResourcesRequest
	GetUnbindResources() []*string
}

type ModifyTemplateResourcesRequest struct {
	// The IDs of the protected assets to associate, in the format of ["XX1","XX2",...].
	BindAssets []*string `json:"BindAssets,omitempty" xml:"BindAssets,omitempty" type:"Repeated"`
	// The protected object groups to associate, in the format of [**"group1","group2",...**].
	BindResourceGroups []*string `json:"BindResourceGroups,omitempty" xml:"BindResourceGroups,omitempty" type:"Repeated"`
	// The protected objects to associate, in the format of [**"XX1","XX2",...**].
	BindResources []*string `json:"BindResources,omitempty" xml:"BindResources,omitempty" type:"Repeated"`
	// Specifies whether to enable the dry run mode. If you do not specify this parameter, a normal request is sent. Valid values:
	//
	// - **true**: A dry run request is sent. The system checks whether the request meets the execution conditions without performing the specified operation. If the dry run fails, the corresponding error code is returned. If the dry run succeeds, the error code Defense.Control.DryRunOperation is returned.
	//
	// - **false**: A normal request is sent. The specified operation is performed after the request passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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
	// The ID of the protection template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2291
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The IDs of the protected assets to disassociate, in the format of ["XX1","XX2",...].
	UnbindAssets []*string `json:"UnbindAssets,omitempty" xml:"UnbindAssets,omitempty" type:"Repeated"`
	// The protected object groups to disassociate, in the format of [**"group1","group2",...**].
	UnbindResourceGroups []*string `json:"UnbindResourceGroups,omitempty" xml:"UnbindResourceGroups,omitempty" type:"Repeated"`
	// The protected objects to disassociate, in the format of [**"XX1","XX2",...**].
	UnbindResources []*string `json:"UnbindResources,omitempty" xml:"UnbindResources,omitempty" type:"Repeated"`
}

func (s ModifyTemplateResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateResourcesRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesRequest) GetBindAssets() []*string {
	return s.BindAssets
}

func (s *ModifyTemplateResourcesRequest) GetBindResourceGroups() []*string {
	return s.BindResourceGroups
}

func (s *ModifyTemplateResourcesRequest) GetBindResources() []*string {
	return s.BindResources
}

func (s *ModifyTemplateResourcesRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *ModifyTemplateResourcesRequest) GetUnbindAssets() []*string {
	return s.UnbindAssets
}

func (s *ModifyTemplateResourcesRequest) GetUnbindResourceGroups() []*string {
	return s.UnbindResourceGroups
}

func (s *ModifyTemplateResourcesRequest) GetUnbindResources() []*string {
	return s.UnbindResources
}

func (s *ModifyTemplateResourcesRequest) SetBindAssets(v []*string) *ModifyTemplateResourcesRequest {
	s.BindAssets = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetBindResourceGroups(v []*string) *ModifyTemplateResourcesRequest {
	s.BindResourceGroups = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetBindResources(v []*string) *ModifyTemplateResourcesRequest {
	s.BindResources = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetDryRun(v bool) *ModifyTemplateResourcesRequest {
	s.DryRun = &v
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

func (s *ModifyTemplateResourcesRequest) SetUnbindAssets(v []*string) *ModifyTemplateResourcesRequest {
	s.UnbindAssets = v
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
