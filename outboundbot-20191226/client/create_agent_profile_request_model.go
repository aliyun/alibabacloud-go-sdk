// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileTemplateId(v string) *CreateAgentProfileRequest
	GetAgentProfileTemplateId() *string
	SetAppIp(v string) *CreateAgentProfileRequest
	GetAppIp() *string
	SetDescription(v string) *CreateAgentProfileRequest
	GetDescription() *string
	SetFaqCategoryIds(v string) *CreateAgentProfileRequest
	GetFaqCategoryIds() *string
	SetInstanceId(v string) *CreateAgentProfileRequest
	GetInstanceId() *string
	SetInstructionJson(v string) *CreateAgentProfileRequest
	GetInstructionJson() *string
	SetLabelsJson(v string) *CreateAgentProfileRequest
	GetLabelsJson() *string
	SetModel(v string) *CreateAgentProfileRequest
	GetModel() *string
	SetModelConfig(v string) *CreateAgentProfileRequest
	GetModelConfig() *string
	SetPrompt(v string) *CreateAgentProfileRequest
	GetPrompt() *string
	SetPromptJson(v string) *CreateAgentProfileRequest
	GetPromptJson() *string
	SetScenario(v string) *CreateAgentProfileRequest
	GetScenario() *string
	SetScriptId(v string) *CreateAgentProfileRequest
	GetScriptId() *string
	SetVariablesJson(v string) *CreateAgentProfileRequest
	GetVariablesJson() *string
}

type CreateAgentProfileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default-survey
	AgentProfileTemplateId *string `json:"AgentProfileTemplateId,omitempty" xml:"AgentProfileTemplateId,omitempty"`
	// example:
	//
	// 127.0.0.1
	AppIp       *string `json:"AppIp,omitempty" xml:"AppIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// [30000474726]
	FaqCategoryIds *string `json:"FaqCategoryIds,omitempty" xml:"FaqCategoryIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// [{\\"type\\":\\"TransferToAgent\\",\\"instructions\\":[{\\"code\\":\\"Transfer0\\",\\"skillGroupId\\":\\"123\\",\\"skillGroupName\\":\\"123\\"}]
	InstructionJson *string `json:"InstructionJson,omitempty" xml:"InstructionJson,omitempty"`
	LabelsJson      *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// ""
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	Prompt      *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptJson  *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	Scenario    *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aa279896-64a6-4182-864c-4f2b04ec8d17
	ScriptId      *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	VariablesJson *string `json:"VariablesJson,omitempty" xml:"VariablesJson,omitempty"`
}

func (s CreateAgentProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentProfileRequest) GetAgentProfileTemplateId() *string {
	return s.AgentProfileTemplateId
}

func (s *CreateAgentProfileRequest) GetAppIp() *string {
	return s.AppIp
}

func (s *CreateAgentProfileRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentProfileRequest) GetFaqCategoryIds() *string {
	return s.FaqCategoryIds
}

func (s *CreateAgentProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAgentProfileRequest) GetInstructionJson() *string {
	return s.InstructionJson
}

func (s *CreateAgentProfileRequest) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *CreateAgentProfileRequest) GetModel() *string {
	return s.Model
}

func (s *CreateAgentProfileRequest) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *CreateAgentProfileRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateAgentProfileRequest) GetPromptJson() *string {
	return s.PromptJson
}

func (s *CreateAgentProfileRequest) GetScenario() *string {
	return s.Scenario
}

func (s *CreateAgentProfileRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateAgentProfileRequest) GetVariablesJson() *string {
	return s.VariablesJson
}

func (s *CreateAgentProfileRequest) SetAgentProfileTemplateId(v string) *CreateAgentProfileRequest {
	s.AgentProfileTemplateId = &v
	return s
}

func (s *CreateAgentProfileRequest) SetAppIp(v string) *CreateAgentProfileRequest {
	s.AppIp = &v
	return s
}

func (s *CreateAgentProfileRequest) SetDescription(v string) *CreateAgentProfileRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentProfileRequest) SetFaqCategoryIds(v string) *CreateAgentProfileRequest {
	s.FaqCategoryIds = &v
	return s
}

func (s *CreateAgentProfileRequest) SetInstanceId(v string) *CreateAgentProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAgentProfileRequest) SetInstructionJson(v string) *CreateAgentProfileRequest {
	s.InstructionJson = &v
	return s
}

func (s *CreateAgentProfileRequest) SetLabelsJson(v string) *CreateAgentProfileRequest {
	s.LabelsJson = &v
	return s
}

func (s *CreateAgentProfileRequest) SetModel(v string) *CreateAgentProfileRequest {
	s.Model = &v
	return s
}

func (s *CreateAgentProfileRequest) SetModelConfig(v string) *CreateAgentProfileRequest {
	s.ModelConfig = &v
	return s
}

func (s *CreateAgentProfileRequest) SetPrompt(v string) *CreateAgentProfileRequest {
	s.Prompt = &v
	return s
}

func (s *CreateAgentProfileRequest) SetPromptJson(v string) *CreateAgentProfileRequest {
	s.PromptJson = &v
	return s
}

func (s *CreateAgentProfileRequest) SetScenario(v string) *CreateAgentProfileRequest {
	s.Scenario = &v
	return s
}

func (s *CreateAgentProfileRequest) SetScriptId(v string) *CreateAgentProfileRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateAgentProfileRequest) SetVariablesJson(v string) *CreateAgentProfileRequest {
	s.VariablesJson = &v
	return s
}

func (s *CreateAgentProfileRequest) Validate() error {
	return dara.Validate(s)
}
