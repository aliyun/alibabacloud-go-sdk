// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseTemplateStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ModifyDefenseTemplateStatusRequest
	GetDryRun() *bool
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
	// Specifies whether to enable the dry run mode. If you do not specify this parameter, a normal request is sent. Valid values:
	//
	// - **true**: Sends a dry run request. The system checks whether the request meets the execution conditions without performing the specified operation. If the dry run fails, the corresponding error code is returned. If the dry run succeeds, the error code Defense.Control.DryRunOperation is returned.
	//
	// - **false**: Sends a normal request. The specified operation is performed after the request passes the check.
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
	// The region where the WAF instance is deployed. Valid values:
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
	// The ID of the protection rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2249
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The status of the protection template that you want to set. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
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

func (s *ModifyDefenseTemplateStatusRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *ModifyDefenseTemplateStatusRequest) SetDryRun(v bool) *ModifyDefenseTemplateStatusRequest {
	s.DryRun = &v
	return s
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
