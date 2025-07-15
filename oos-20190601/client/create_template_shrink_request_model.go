// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateTemplateShrinkRequest
	GetContent() *string
	SetRegionId(v string) *CreateTemplateShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTemplateShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *CreateTemplateShrinkRequest
	GetTagsShrink() *string
	SetTemplateName(v string) *CreateTemplateShrinkRequest
	GetTemplateName() *string
	SetVersionName(v string) *CreateTemplateShrinkRequest
	GetVersionName() *string
}

type CreateTemplateShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s CreateTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateTemplateShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTemplateShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTemplateShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateTemplateShrinkRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateTemplateShrinkRequest) SetContent(v string) *CreateTemplateShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetRegionId(v string) *CreateTemplateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetResourceGroupId(v string) *CreateTemplateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetTagsShrink(v string) *CreateTemplateShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetTemplateName(v string) *CreateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetVersionName(v string) *CreateTemplateShrinkRequest {
	s.VersionName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
