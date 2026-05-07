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
	// 字段约束规则。Key 为 JSONPath 表达式，Value 为约束类型：locked（锁定不可覆盖）、overridable（可覆盖）、required（必填）。需与 Content 同时提供，不允许单独更新。
	//
	// example:
	//
	// {\\"JobSpecs[0].Image\\":\\"locked\\",\\"UserCommand\\":\\"locked\\",\\"JobType\\":\\"locked\\"}
	Constraints map[string]interface{} `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// 任务模板的配置内容，支持 CreateJob 接口的所有参数字段，以 JSON 格式传入。提供时会创建新版本。
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
	// 当创建了新版本时，是否将新版本设为默认版本。
	//
	// example:
	//
	// true
	SetAsDefault *bool `json:"SetAsDefault,omitempty" xml:"SetAsDefault,omitempty"`
	// example:
	//
	// job-template-example-1778047****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Version      *int32  `json:"version,omitempty" xml:"version,omitempty"`
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
