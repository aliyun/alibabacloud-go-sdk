// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDefenseTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *CopyDefenseTemplateRequest
	GetDryRun() *bool
	SetInstanceId(v string) *CopyDefenseTemplateRequest
	GetInstanceId() *string
	SetRegionId(v string) *CopyDefenseTemplateRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CopyDefenseTemplateRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *CopyDefenseTemplateRequest
	GetTemplateId() *int64
}

type CopyDefenseTemplateRequest struct {
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
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-lbj****x10g
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
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the mitigation template to copy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CopyDefenseTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *CopyDefenseTemplateRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CopyDefenseTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CopyDefenseTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyDefenseTemplateRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CopyDefenseTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CopyDefenseTemplateRequest) SetDryRun(v bool) *CopyDefenseTemplateRequest {
	s.DryRun = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetInstanceId(v string) *CopyDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetRegionId(v string) *CopyDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *CopyDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetTemplateId(v int64) *CopyDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) Validate() error {
	return dara.Validate(s)
}
