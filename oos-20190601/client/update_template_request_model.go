// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateTemplateRequest
	GetContent() *string
	SetRegionId(v string) *UpdateTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateTemplateRequest
	GetResourceGroupId() *string
	SetTags(v map[string]interface{}) *UpdateTemplateRequest
	GetTags() map[string]interface{}
	SetTemplateName(v string) *UpdateTemplateRequest
	GetTemplateName() *string
	SetVersionName(v string) *UpdateTemplateRequest
	GetVersionName() *string
}

type UpdateTemplateRequest struct {
	// The content of the template. The content must be in the JSON or YAML format, and its maximum size is 64 KB.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "FormatVersion": "OOS-2019-06-01", "Description": { "en": "Bulky starts the ECS instances", "name-en": "Bulky Start Instances", }, "Parameters": { "regionId": { "Type": "String", "Label": { "en": "RegionId", }, "AssociationProperty": "RegionId", "Default": "{{ ACS::RegionId }}" }, "targets": { "Type": "Json", "Label": { "en": "TargetInstance", }, "AssociationProperty": "Targets", "AssociationPropertyMetadata": { "ResourceType": "ALIYUN::ECS::Instance", "RegionId": "regionId" } }, "rateControl": { "Label": { "en": "RateControl", }, "Type": "Json", "AssociationProperty": "RateControl", "Default": { "Mode": "Concurrency", "MaxErrors": 0, "Concurrency": 10 } }, "OOSAssumeRole": { "Label": { "en": "OOSAssumeRole", }, "Type": "String", "Default": "OOSServiceRole" } }, "RamRole": "{{ OOSAssumeRole }}", "Tasks": [ { "Name": "getInstance", "Description": { "en": "Views the ECS instances", }, "Action": "ACS::SelectTargets", "Properties": { "ResourceType": "ALIYUN::ECS::Instance", "RegionId": "{{ regionId }}", "Filters": [ "{{ targets }}" ] }, "Outputs": { "instanceIds": { "Type": "List", "ValueSelector": "Instances.Instance[].InstanceId" } } }, { "Name": "startInstance", "Action": "ACS::ECS::StartInstance", "Description": { "en": "Starts the ECS instances", }, "Properties": { "regionId": "{{ regionId }}", "instanceId": "{{ ACS::TaskLoopItem }}" }, "Loop": { "RateControl": "{{ rateControl }}", "Items": "{{ getInstance.instanceIds }}" } } ], "Outputs": { "instanceIds": { "Type": "List", "Value": "{{ getInstance.instanceIds }}" } } }
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
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name can be up to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The name cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the template version.
	//
	// example:
	//
	// v2
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTemplateRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateTemplateRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *UpdateTemplateRequest) SetContent(v string) *UpdateTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdateTemplateRequest) SetRegionId(v string) *UpdateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateRequest) SetResourceGroupId(v string) *UpdateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateRequest) SetTags(v map[string]interface{}) *UpdateTemplateRequest {
	s.Tags = v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetVersionName(v string) *UpdateTemplateRequest {
	s.VersionName = &v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
