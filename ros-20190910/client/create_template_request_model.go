// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTemplateRequest
	GetDescription() *string
	SetResourceGroupId(v string) *CreateTemplateRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateTemplateRequestTags) *CreateTemplateRequest
	GetTags() []*CreateTemplateRequestTags
	SetTemplateBody(v string) *CreateTemplateRequest
	GetTemplateBody() *string
	SetTemplateName(v string) *CreateTemplateRequest
	GetTemplateName() *string
	SetTemplateURL(v string) *CreateTemplateRequest
	GetTemplateURL() *string
	SetValidationOptions(v []*string) *CreateTemplateRequest
	GetValidationOptions() []*string
}

type CreateTemplateRequest struct {
	// The description of the template. The description can be up to 256 characters in length.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the resource group.\\
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.html).
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the template.
	Tags         []*CreateTemplateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateBody *string                      `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The name of the template.\\
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The template body must be 1 to 1,024 bytes in length. If you do not specify the region of the OSS bucket, the value of RegionId is used.
	//
	// > You must specify TemplateBody or TemplateURL.
	//
	// example:
	//
	// oss://ros/template/demo
	TemplateURL       *string   `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	ValidationOptions []*string `json:"ValidationOptions,omitempty" xml:"ValidationOptions,omitempty" type:"Repeated"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTemplateRequest) GetTags() []*CreateTemplateRequestTags {
	return s.Tags
}

func (s *CreateTemplateRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *CreateTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateTemplateRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *CreateTemplateRequest) GetValidationOptions() []*string {
	return s.ValidationOptions
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceGroupId(v string) *CreateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateRequest) SetTags(v []*CreateTemplateRequestTags) *CreateTemplateRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateRequest) SetTemplateBody(v string) *CreateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateURL(v string) *CreateTemplateRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateTemplateRequest) SetValidationOptions(v []*string) *CreateTemplateRequest {
	s.ValidationOptions = v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateRequestTags struct {
	// The tag key of the template.
	//
	// > Tags is optional. If you need to specify Tags, you must also specify Key.
	//
	// This parameter is required.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the template.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTemplateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateTemplateRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateTemplateRequestTags) SetKey(v string) *CreateTemplateRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTemplateRequestTags) SetValue(v string) *CreateTemplateRequestTags {
	s.Value = &v
	return s
}

func (s *CreateTemplateRequestTags) Validate() error {
	return dara.Validate(s)
}
