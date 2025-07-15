// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateTemplateRequest
	GetContent() *string
	SetRegionId(v string) *CreateTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTemplateRequest
	GetResourceGroupId() *string
	SetTags(v map[string]interface{}) *CreateTemplateRequest
	GetTags() map[string]interface{}
	SetTemplateName(v string) *CreateTemplateRequest
	GetTemplateName() *string
	SetVersionName(v string) *CreateTemplateRequest
	GetVersionName() *string
}

type CreateTemplateRequest struct {
	// The content of the template. The content must be in the JSON or YAML format, and its maximum size is 64 KB.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"FormatVersion": "OOS-2019-06-01", "Description": "Describe instances of given status", "Parameters": {"Status": {"Type": "String", "Description": "(Required) The status of the Ecs instance."}}, "Tasks": [{"Properties": {"Parameters": {"Status": "{{ Status }}"}, "API": "DescribeInstances", "Service": "Ecs"}, "Name": "foo", "Action": "ACS::ExecuteApi"}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag keys and tag values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name can be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The name cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the version of the template.
	//
	// example:
	//
	// v2
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *CreateTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTemplateRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateTemplateRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateTemplateRequest) SetContent(v string) *CreateTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateRequest) SetRegionId(v string) *CreateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceGroupId(v string) *CreateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateRequest) SetTags(v map[string]interface{}) *CreateTemplateRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetVersionName(v string) *CreateTemplateRequest {
	s.VersionName = &v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
