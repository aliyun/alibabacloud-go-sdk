// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateExecutionPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRamRole(v string) *GenerateExecutionPolicyRequest
	GetRamRole() *string
	SetRegionId(v string) *GenerateExecutionPolicyRequest
	GetRegionId() *string
	SetTemplateContent(v string) *GenerateExecutionPolicyRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *GenerateExecutionPolicyRequest
	GetTemplateName() *string
	SetTemplateVersion(v string) *GenerateExecutionPolicyRequest
	GetTemplateVersion() *string
}

type GenerateExecutionPolicyRequest struct {
	// The RAM role.
	//
	// example:
	//
	// AliyunServiceRoleForOOSBandwidthScheduler
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The content of the template in the JSON or YAML format. This parameter is the same as the Content parameter that you can specify when you call the CreateTemplate operation. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// {   "Description": "Example template, describe instances in some status",   "FormatVersion": "OOS-2019-06-01",   "Parameters": {},   "Tasks": [     {       "Name": "describeInstances",       "Action": "ACS::ExecuteAPI",       "Description": "desc-en",       "Properties": {         "Service": "ECS",         "API": "DescribeInstances",         "Parameters": {           "Status": "Running"         }       }     }   ] }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// vmeixme
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. The default value is the latest version of the template.
	//
	// example:
	//
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GenerateExecutionPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateExecutionPolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *GenerateExecutionPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateExecutionPolicyRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *GenerateExecutionPolicyRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GenerateExecutionPolicyRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GenerateExecutionPolicyRequest) SetRamRole(v string) *GenerateExecutionPolicyRequest {
	s.RamRole = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetRegionId(v string) *GenerateExecutionPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetTemplateContent(v string) *GenerateExecutionPolicyRequest {
	s.TemplateContent = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetTemplateName(v string) *GenerateExecutionPolicyRequest {
	s.TemplateName = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetTemplateVersion(v string) *GenerateExecutionPolicyRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) Validate() error {
	return dara.Validate(s)
}
