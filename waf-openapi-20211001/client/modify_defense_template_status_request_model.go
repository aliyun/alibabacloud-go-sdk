// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseTemplateStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyDefenseTemplateStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseTemplateStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseTemplateStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *ModifyDefenseTemplateStatusRequest
	GetTemplateId() *int64
	SetTemplateStatus(v int32) *ModifyDefenseTemplateStatusRequest
	GetTemplateStatus() *int32
}

type ModifyDefenseTemplateStatusRequest struct {
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
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
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
	// The ID of the protection rule template whose status you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2249
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The new status of the protection rule template. Valid values:
	//
	// 	- **0:*	- disabled.
	//
	// 	- **1:*	- enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
}

func (s ModifyDefenseTemplateStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseTemplateStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseTemplateStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseTemplateStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseTemplateStatusRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyDefenseTemplateStatusRequest) GetTemplateStatus() *int32 {
	return s.TemplateStatus
}

func (s *ModifyDefenseTemplateStatusRequest) SetInstanceId(v string) *ModifyDefenseTemplateStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetRegionId(v string) *ModifyDefenseTemplateStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseTemplateStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetTemplateId(v int64) *ModifyDefenseTemplateStatusRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetTemplateStatus(v int32) *ModifyDefenseTemplateStatusRequest {
	s.TemplateStatus = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) Validate() error {
	return dara.Validate(s)
}
