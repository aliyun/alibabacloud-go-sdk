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
	// 字段约束规则。Key 为 JSONPath 表达式，Value 为约束类型：locked（锁定不可覆盖）、overridable（可覆盖）、required（必填）。
	//
	// example:
	//
	// {\\"JobSpecs[0].Image\\":\\"locked\\",\\"UserCommand\\":\\"locked\\",\\"JobType\\":\\"locked\\"}
	Constraints map[string]interface{} `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// 任务模板的配置内容，包含作业配置参数，以 JSON 格式传入。
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"WorkspaceId\\":\\"15****05\\",\\"JobType\\":\\"PyTorchJob\\",\\"UserCommand\\":\\"echo hello\\",\\"JobSpecs\\":[{\\"Type\\":\\"Worker\\",\\"PodCount\\":1,\\"Image\\":\\"dsw-registry-vpc.cn-hangzhou.cr.aliyuncs.com/pai/pytorch:2.8.0-gpu-py313-cu129-ubuntu22.04-3995b779-1764361782\\",\\"EcsSpec\\":\\"ecs.gn7i-c8g1.2xlarge\\"}],\\"ResourceType\\":\\"ECS\\",\\"_ResourcePaymentType\\":\\"PostPaid\\",\\"CredentialConfig\\":{\\"EnableCredentialInject\\":false},\\"Accessibility\\":\\"PRIVATE\\",\\"Settings\\":{\\"JobReservedMinutes\\":0,\\"Tags\\":{}}}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// Template description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 用户自定义的键值对元数据，用于存储模板的附加信息。
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-template-example-1778047****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 工作空间 ID。如何获取工作空间 ID，请参见 ListWorkspaces。
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
