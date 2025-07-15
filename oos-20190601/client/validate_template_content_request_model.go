// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTemplateContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ValidateTemplateContentRequest
	GetContent() *string
	SetRegionId(v string) *ValidateTemplateContentRequest
	GetRegionId() *string
	SetTemplateURL(v string) *ValidateTemplateContentRequest
	GetTemplateURL() *string
}

type ValidateTemplateContentRequest struct {
	// The content of the template.
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
	// The URL that is used to store the content of the Operation Orchestration Service (OOS) template in the Alibaba Cloud Object Storage Service (OSS). Only the public-read URL is supported. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// http:/oos-template.cn-hangzhou.oss.aliyun-inc.com/oos-template.json
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
}

func (s ValidateTemplateContentRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateContentRequest) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentRequest) GetContent() *string {
	return s.Content
}

func (s *ValidateTemplateContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ValidateTemplateContentRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *ValidateTemplateContentRequest) SetContent(v string) *ValidateTemplateContentRequest {
	s.Content = &v
	return s
}

func (s *ValidateTemplateContentRequest) SetRegionId(v string) *ValidateTemplateContentRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateTemplateContentRequest) SetTemplateURL(v string) *ValidateTemplateContentRequest {
	s.TemplateURL = &v
	return s
}

func (s *ValidateTemplateContentRequest) Validate() error {
	return dara.Validate(s)
}
