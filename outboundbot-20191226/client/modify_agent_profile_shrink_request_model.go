// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileId(v string) *ModifyAgentProfileShrinkRequest
	GetAgentProfileId() *string
	SetApiPluginJson(v string) *ModifyAgentProfileShrinkRequest
	GetApiPluginJson() *string
	SetDescription(v string) *ModifyAgentProfileShrinkRequest
	GetDescription() *string
	SetFaqCategoryIdsShrink(v string) *ModifyAgentProfileShrinkRequest
	GetFaqCategoryIdsShrink() *string
	SetInstanceId(v string) *ModifyAgentProfileShrinkRequest
	GetInstanceId() *string
	SetInstructionJson(v string) *ModifyAgentProfileShrinkRequest
	GetInstructionJson() *string
	SetLabelsJson(v string) *ModifyAgentProfileShrinkRequest
	GetLabelsJson() *string
	SetModel(v string) *ModifyAgentProfileShrinkRequest
	GetModel() *string
	SetModelConfig(v string) *ModifyAgentProfileShrinkRequest
	GetModelConfig() *string
	SetPrompt(v string) *ModifyAgentProfileShrinkRequest
	GetPrompt() *string
	SetPromptJson(v string) *ModifyAgentProfileShrinkRequest
	GetPromptJson() *string
	SetScenario(v string) *ModifyAgentProfileShrinkRequest
	GetScenario() *string
	SetVariablesJson(v string) *ModifyAgentProfileShrinkRequest
	GetVariablesJson() *string
}

type ModifyAgentProfileShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 37ca3ca1ac4b4e57adf3da5b5d939d04
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// example:
	//
	// []
	ApiPluginJson        *string `json:"ApiPluginJson,omitempty" xml:"ApiPluginJson,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FaqCategoryIdsShrink *string `json:"FaqCategoryIds,omitempty" xml:"FaqCategoryIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// [{"type":"TransferToAgent","instructions":[{"code":"Transfer0","skillGroupId":"123","skillGroupName":"123"}],"timeoutEnable":false},{"type":"CollectNumber","instructions":[]}]
	InstructionJson *string `json:"InstructionJson,omitempty" xml:"InstructionJson,omitempty"`
	LabelsJson      *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// {}
	ModelConfig   *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	Prompt        *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptJson    *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	Scenario      *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	VariablesJson *string `json:"VariablesJson,omitempty" xml:"VariablesJson,omitempty"`
}

func (s ModifyAgentProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgentProfileShrinkRequest) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *ModifyAgentProfileShrinkRequest) GetApiPluginJson() *string {
	return s.ApiPluginJson
}

func (s *ModifyAgentProfileShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAgentProfileShrinkRequest) GetFaqCategoryIdsShrink() *string {
	return s.FaqCategoryIdsShrink
}

func (s *ModifyAgentProfileShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAgentProfileShrinkRequest) GetInstructionJson() *string {
	return s.InstructionJson
}

func (s *ModifyAgentProfileShrinkRequest) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *ModifyAgentProfileShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ModifyAgentProfileShrinkRequest) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *ModifyAgentProfileShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *ModifyAgentProfileShrinkRequest) GetPromptJson() *string {
	return s.PromptJson
}

func (s *ModifyAgentProfileShrinkRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ModifyAgentProfileShrinkRequest) GetVariablesJson() *string {
	return s.VariablesJson
}

func (s *ModifyAgentProfileShrinkRequest) SetAgentProfileId(v string) *ModifyAgentProfileShrinkRequest {
	s.AgentProfileId = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetApiPluginJson(v string) *ModifyAgentProfileShrinkRequest {
	s.ApiPluginJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetDescription(v string) *ModifyAgentProfileShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetFaqCategoryIdsShrink(v string) *ModifyAgentProfileShrinkRequest {
	s.FaqCategoryIdsShrink = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetInstanceId(v string) *ModifyAgentProfileShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetInstructionJson(v string) *ModifyAgentProfileShrinkRequest {
	s.InstructionJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetLabelsJson(v string) *ModifyAgentProfileShrinkRequest {
	s.LabelsJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetModel(v string) *ModifyAgentProfileShrinkRequest {
	s.Model = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetModelConfig(v string) *ModifyAgentProfileShrinkRequest {
	s.ModelConfig = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetPrompt(v string) *ModifyAgentProfileShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetPromptJson(v string) *ModifyAgentProfileShrinkRequest {
	s.PromptJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetScenario(v string) *ModifyAgentProfileShrinkRequest {
	s.Scenario = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetVariablesJson(v string) *ModifyAgentProfileShrinkRequest {
	s.VariablesJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
