// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConstraints(v map[string]interface{}) *CreateJobTemplateRequest
	GetConstraints() map[string]interface{}
	SetContent(v string) *CreateJobTemplateRequest
	GetContent() *string
	SetDescription(v string) *CreateJobTemplateRequest
	GetDescription() *string
	SetMetadata(v map[string]interface{}) *CreateJobTemplateRequest
	GetMetadata() map[string]interface{}
	SetTemplateName(v string) *CreateJobTemplateRequest
	GetTemplateName() *string
	SetWorkspaceId(v string) *CreateJobTemplateRequest
	GetWorkspaceId() *string
}

type CreateJobTemplateRequest struct {
	// The field constraint rules. The key is a JSONPath expression, and the value is a constraint type: `locked` (cannot be overridden), `overridable` (can be overridden), or `required` (must be specified).
	//
	// example:
	//
	// {\\"JobSpecs[0].Image\\":\\"locked\\",\\"UserCommand\\":\\"locked\\",\\"JobType\\":\\"locked\\"}
	Constraints map[string]interface{} `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The configuration of the job template, which must be a JSON string containing the job configuration parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"WorkspaceId\\":\\"15****05\\",\\"JobType\\":\\"PyTorchJob\\",\\"UserCommand\\":\\"echo hello\\",\\"JobSpecs\\":[{\\"Type\\":\\"Worker\\",\\"PodCount\\":1,\\"Image\\":\\"dsw-registry-vpc.cn-hangzhou.cr.aliyuncs.com/pai/pytorch:2.8.0-gpu-py313-cu129-ubuntu22.04-3995b779-1764361782\\",\\"EcsSpec\\":\\"ecs.gn7i-c8g1.2xlarge\\"}],\\"ResourceType\\":\\"ECS\\",\\"_ResourcePaymentType\\":\\"PostPaid\\",\\"CredentialConfig\\":{\\"EnableCredentialInject\\":false},\\"Accessibility\\":\\"PRIVATE\\",\\"Settings\\":{\\"JobReservedMinutes\\":0,\\"Tags\\":{}}}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the job template.
	//
	// example:
	//
	// Template description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// User-defined key-value metadata.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the job template.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-template-example-1778047****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The ID of the workspace that contains the job template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15****05
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequest) GetConstraints() map[string]interface{} {
	return s.Constraints
}

func (s *CreateJobTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *CreateJobTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateJobTemplateRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *CreateJobTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateJobTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateJobTemplateRequest) SetConstraints(v map[string]interface{}) *CreateJobTemplateRequest {
	s.Constraints = v
	return s
}

func (s *CreateJobTemplateRequest) SetContent(v string) *CreateJobTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateJobTemplateRequest) SetDescription(v string) *CreateJobTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateJobTemplateRequest) SetMetadata(v map[string]interface{}) *CreateJobTemplateRequest {
	s.Metadata = v
	return s
}

func (s *CreateJobTemplateRequest) SetTemplateName(v string) *CreateJobTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateJobTemplateRequest) SetWorkspaceId(v string) *CreateJobTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}
