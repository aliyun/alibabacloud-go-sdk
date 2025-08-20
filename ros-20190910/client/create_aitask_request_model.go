// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAITaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *CreateAITaskRequest
	GetPrompt() *string
	SetTaskType(v string) *CreateAITaskRequest
	GetTaskType() *string
	SetTemplate(v string) *CreateAITaskRequest
	GetTemplate() *string
	SetTemplateType(v string) *CreateAITaskRequest
	GetTemplateType() *string
}

type CreateAITaskRequest struct {
	// The input description for the AI task.
	//
	// - When the task type is Generate Template, this parameter specifies the functionality of the template to be generated.
	//
	// - When the task type is FixTemplate, this parameter can describe how the template should be repaired.
	//
	// example:
	//
	// 创建一台ECS，部署nignx服务
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The type of AI task. Values:
	//
	// - GenerateTemplate: AI template generation
	//
	// - FixTemplate: AI template repair
	//
	// example:
	//
	// GenerateTemplate
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// When the task type is AI template repair, specify the original template that needs to be fixed or modified.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion": "2015-09-01"}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the template to be generated or repaired. Default is ROS.
	//
	// example:
	//
	// ROS
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateAITaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAITaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAITaskRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateAITaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateAITaskRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreateAITaskRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateAITaskRequest) SetPrompt(v string) *CreateAITaskRequest {
	s.Prompt = &v
	return s
}

func (s *CreateAITaskRequest) SetTaskType(v string) *CreateAITaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateAITaskRequest) SetTemplate(v string) *CreateAITaskRequest {
	s.Template = &v
	return s
}

func (s *CreateAITaskRequest) SetTemplateType(v string) *CreateAITaskRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateAITaskRequest) Validate() error {
	return dara.Validate(s)
}
