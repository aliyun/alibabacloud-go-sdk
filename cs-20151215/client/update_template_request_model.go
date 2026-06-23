// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTemplateRequest
	GetDescription() *string
	SetName(v string) *UpdateTemplateRequest
	GetName() *string
	SetTags(v string) *UpdateTemplateRequest
	GetTags() *string
	SetTemplate(v string) *UpdateTemplateRequest
	GetTemplate() *string
	SetTemplateType(v string) *UpdateTemplateRequest
	GetTemplateType() *string
}

type UpdateTemplateRequest struct {
	// The description of the deployment template.
	//
	// example:
	//
	// web server cluster
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the deployment template.
	//
	// example:
	//
	// webserver01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The tags of the deployment template.
	//
	// example:
	//
	// web
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The template content in YAML format.
	//
	// example:
	//
	// apiVersion: apps/v1\\\\nkind: Deployment\\\\nmetadata:\\\\n  name: nginx-deployment-basic\\\\n  labels:\\\\n    app: nginx\\\\nspec:\\\\n  replicas: 2\\\\n  selector:\\\\n    matchLabels:\\\\n      app: nginx\\\\n  template:\\\\n    metadata:\\\\n      labels:\\\\n        app: nginx\\\\n    spec:\\\\n      containers:\\\\n      - name: nginx\\\\n        image: busybox:latest\\\\n        ports:\\\\n        - containerPort: 8080
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// The templatetype.
	//
	// - If you set this parameter to `kubernetes`, the template is displayed on the Orchestration Templates page in the console.
	//
	// - If you leave this parameter empty or set it to other values, the template is not displayed on the Orchestration Templates page in the console.
	//
	// Settings this parameter to `kubernetes` is recommended.
	//
	// example:
	//
	// kubernetes
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTemplateRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *UpdateTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *UpdateTemplateRequest) SetDescription(v string) *UpdateTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateRequest) SetName(v string) *UpdateTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateTemplateRequest) SetTags(v string) *UpdateTemplateRequest {
	s.Tags = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplate(v string) *UpdateTemplateRequest {
	s.Template = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateType(v string) *UpdateTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
