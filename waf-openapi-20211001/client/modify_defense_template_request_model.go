// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyDefenseTemplateRequest
	GetDescription() *string
	SetDetail(v string) *ModifyDefenseTemplateRequest
	GetDetail() *string
	SetDryRun(v bool) *ModifyDefenseTemplateRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ModifyDefenseTemplateRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseTemplateRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseTemplateRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *ModifyDefenseTemplateRequest
	GetTemplateId() *int64
	SetTemplateName(v string) *ModifyDefenseTemplateRequest
	GetTemplateName() *string
}

type ModifyDefenseTemplateRequest struct {
	// The description of the protection template that you want to modify.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the template. For more information, see the Detail parameter in [CreateDefenseTemplate](https://help.aliyun.com/document_detail/461613.html).
	//
	// example:
	//
	// {"trafficFeature":"{\\"global\\":0,\\"excludeStatus\\":1,\\"conditions\\":[{\\"key\\":\\"URL\\",\\"opValue\\":\\"not-contain\\",\\"values\\":\\"test\\"}]}"}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
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
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
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
	// The ID of the protection template that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7392
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection template that you want to modify. The name must be 1 to 255 characters in length and can contain Chinese characters, letters, digits, underscores (_), periods (.), and hyphens (-).
	//
	// > Template names within the same protection scenario (**DefenseScene**) must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifyDefenseTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDefenseTemplateRequest) GetDetail() *string {
	return s.Detail
}

func (s *ModifyDefenseTemplateRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDefenseTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseTemplateRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyDefenseTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyDefenseTemplateRequest) SetDescription(v string) *ModifyDefenseTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetDetail(v string) *ModifyDefenseTemplateRequest {
	s.Detail = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetDryRun(v bool) *ModifyDefenseTemplateRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetInstanceId(v string) *ModifyDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetRegionId(v string) *ModifyDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetTemplateId(v int64) *ModifyDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetTemplateName(v string) *ModifyDefenseTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) Validate() error {
	return dara.Validate(s)
}
