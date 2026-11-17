// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseBuildTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentLocale(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetContentLocale() *string
	SetDescription(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetDescription() *string
	SetInputShrink(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetInputShrink() *string
	SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetKnowledgeBaseId() *string
	SetLlmModel(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetLlmModel() *string
	SetName(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetName() *string
	SetProjectId(v int64) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetProjectId() *int64
	SetRuntimeResourceShrink(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetRuntimeResourceShrink() *string
	SetTaskId(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetTaskId() *string
	SetTriggerShrink(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest
	GetTriggerShrink() *string
}

type UpdateKnowledgeBaseBuildTaskShrinkRequest struct {
	// The content language.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The new description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input configuration that completely replaces the original input.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The associated knowledge base.
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
	// AUTO
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The new task name.
	//
	// example:
	//
	// Sales Project Knowledge Building
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workspace used for future runs.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Updates the runtime resources based on the submitted fields. If ResourceGroupId or Cu is not provided, the original value is retained.
	RuntimeResourceShrink *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The trigger configuration that completely replaces the original trigger.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s UpdateKnowledgeBaseBuildTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetLlmModel() *string {
	return s.LlmModel
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetRuntimeResourceShrink() *string {
	return s.RuntimeResourceShrink
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetContentLocale(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.ContentLocale = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetDescription(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetInputShrink(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetLlmModel(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.LlmModel = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetName(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetProjectId(v int64) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetRuntimeResourceShrink(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.RuntimeResourceShrink = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetTaskId(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) SetTriggerShrink(v string) *UpdateKnowledgeBaseBuildTaskShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
