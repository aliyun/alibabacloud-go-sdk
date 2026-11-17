// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseBuildTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetClientToken() *string
	SetContentLocale(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetContentLocale() *string
	SetDescription(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetDescription() *string
	SetInputShrink(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetInputShrink() *string
	SetKnowledgeBaseId(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetKnowledgeBaseId() *string
	SetLlmModel(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetLlmModel() *string
	SetName(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetName() *string
	SetProjectId(v int64) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetProjectId() *int64
	SetRuntimeResourceShrink(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetRuntimeResourceShrink() *string
	SetTriggerShrink(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetTriggerShrink() *string
	SetType(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest
	GetType() *string
}

type CreateKnowledgeBaseBuildTaskShrinkRequest struct {
	// The idempotency token for the creation request.
	//
	// example:
	//
	// task-create-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the content.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The description of the task. The description can be up to 1,024 characters in length. If this parameter is omitted, no description is set.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies the input scope by selecting projects, entities, sources, or files based on the type.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The ID of the knowledge base to which the task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The specific model ID or AUTO.
	//
	// example:
	//
	// model-001
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// Project space knowledge building
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the workspace where the task runs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder compute units (CUs).
	//
	// This parameter is required.
	RuntimeResourceShrink *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The manual or periodic trigger configuration.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// The available build types.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_SPACE_BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetLlmModel() *string {
	return s.LlmModel
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetRuntimeResourceShrink() *string {
	return s.RuntimeResourceShrink
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetClientToken(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetContentLocale(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.ContentLocale = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetDescription(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetInputShrink(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetKnowledgeBaseId(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetLlmModel(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.LlmModel = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetName(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetProjectId(v int64) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetRuntimeResourceShrink(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.RuntimeResourceShrink = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetTriggerShrink(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) SetType(v string) *CreateKnowledgeBaseBuildTaskShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
