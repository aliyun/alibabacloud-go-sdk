// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConstraints(v map[string]interface{}) *UpdateJobTemplateRequest
	GetConstraints() map[string]interface{}
	SetContent(v string) *UpdateJobTemplateRequest
	GetContent() *string
	SetDescription(v string) *UpdateJobTemplateRequest
	GetDescription() *string
	SetMetadata(v map[string]interface{}) *UpdateJobTemplateRequest
	GetMetadata() map[string]interface{}
	SetSetAsDefault(v bool) *UpdateJobTemplateRequest
	GetSetAsDefault() *bool
	SetTemplateName(v string) *UpdateJobTemplateRequest
	GetTemplateName() *string
	SetVersion(v int32) *UpdateJobTemplateRequest
	GetVersion() *int32
}

type UpdateJobTemplateRequest struct {
	// The field constraint rules. The key is a JSONPath expression and the value is the constraint type: locked (cannot be overridden), overridable (can be overridden), or required (mandatory). Must be provided together with Content. You cannot update this field independently.
	//
	// example:
	//
	// {\\"JobSpecs[0].Image\\":\\"locked\\",\\"UserCommand\\":\\"locked\\",\\"JobType\\":\\"locked\\"}
	Constraints map[string]interface{} `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The configuration content of the task template. Supports all parameter fields of the CreateJob operation, passed in JSON format. Providing this field creates a new version.
	//
	// example:
	//
	// {\\"WorkspaceId\\":\\"15****05\\",\\"JobType\\":\\"PyTorchJob\\",\\"UserCommand\\":\\"echo hello\\",\\"JobSpecs\\":[{\\"Type\\":\\"Worker\\",\\"PodCount\\":1,\\"Image\\":\\"dsw-registry-vpc.cn-hangzhou.cr.aliyuncs.com/pai/pytorch:2.8.0-gpu-py313-cu129-ubuntu22.04-3995b779-1764361782\\",\\"EcsSpec\\":\\"ecs.gn7i-c8g1.2xlarge\\"}],\\"ResourceType\\":\\"ECS\\",\\"_ResourcePaymentType\\":\\"PostPaid\\",\\"CredentialConfig\\":{\\"EnableCredentialInject\\":false},\\"Accessibility\\":\\"PRIVATE\\",\\"Settings\\":{\\"JobReservedMinutes\\":0,\\"Tags\\":{}}}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the task template.
	//
	// example:
	//
	// Template description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user-defined key-value pair metadata.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Specifies whether to set the new version as the default version when a new version is created.
	//
	// example:
	//
	// true
	SetAsDefault *bool `json:"SetAsDefault,omitempty" xml:"SetAsDefault,omitempty"`
	// The name of the task template.
	//
	// example:
	//
	// job-template-example-1778047****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Invalid field.
	//
	// example:
	//
	// Invalid field
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequest) GetConstraints() map[string]interface{} {
	return s.Constraints
}

func (s *UpdateJobTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateJobTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateJobTemplateRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *UpdateJobTemplateRequest) GetSetAsDefault() *bool {
	return s.SetAsDefault
}

func (s *UpdateJobTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateJobTemplateRequest) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateJobTemplateRequest) SetConstraints(v map[string]interface{}) *UpdateJobTemplateRequest {
	s.Constraints = v
	return s
}

func (s *UpdateJobTemplateRequest) SetContent(v string) *UpdateJobTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdateJobTemplateRequest) SetDescription(v string) *UpdateJobTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateJobTemplateRequest) SetMetadata(v map[string]interface{}) *UpdateJobTemplateRequest {
	s.Metadata = v
	return s
}

func (s *UpdateJobTemplateRequest) SetSetAsDefault(v bool) *UpdateJobTemplateRequest {
	s.SetAsDefault = &v
	return s
}

func (s *UpdateJobTemplateRequest) SetTemplateName(v string) *UpdateJobTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateJobTemplateRequest) SetVersion(v int32) *UpdateJobTemplateRequest {
	s.Version = &v
	return s
}

func (s *UpdateJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}
