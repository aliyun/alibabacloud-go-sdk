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
	SetName(v string) *CreateTemplateRequest
	GetName() *string
	SetTags(v string) *CreateTemplateRequest
	GetTags() *string
	SetTemplate(v string) *CreateTemplateRequest
	GetTemplate() *string
	SetTemplateType(v string) *CreateTemplateRequest
	GetTemplateType() *string
}

type CreateTemplateRequest struct {
	// The description of the template.
	//
	// example:
	//
	// this is test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the orchestration template.
	//
	// The name must be 1 to 63 characters in length, and can contain digits, letters, and hyphens (-). It cannot start with a hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// service-account-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The label of the template.
	//
	// example:
	//
	// test
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The template content in the YAML format.
	//
	// This parameter is required.
	//
	// example:
	//
	// apiVersion: v1\\nkind: ServiceAccount\\nmetadata:\\n  name: test-sa
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// The template type.
	//
	// 	- If the parameter is set to `kubernetes`, the template is displayed on the Templates page in the console.
	//
	// 	- If this parameter is not specified or the value is set to another value, the template is not displayed on the Orchestration Template page in the console.
	//
	// We recommend that you set the parameter to `kubernetes`.
	//
	// example:
	//
	// kubernetes
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
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

func (s *CreateTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateTemplateRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreateTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetName(v string) *CreateTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateTemplateRequest) SetTags(v string) *CreateTemplateRequest {
	s.Tags = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplate(v string) *CreateTemplateRequest {
	s.Template = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateType(v string) *CreateTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
