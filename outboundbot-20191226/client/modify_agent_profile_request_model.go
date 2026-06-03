// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileId(v string) *ModifyAgentProfileRequest
	GetAgentProfileId() *string
	SetApiPluginJson(v string) *ModifyAgentProfileRequest
	GetApiPluginJson() *string
	SetDescription(v string) *ModifyAgentProfileRequest
	GetDescription() *string
	SetFaqCategoryIds(v []*int64) *ModifyAgentProfileRequest
	GetFaqCategoryIds() []*int64
	SetInstanceId(v string) *ModifyAgentProfileRequest
	GetInstanceId() *string
	SetInstructionJson(v string) *ModifyAgentProfileRequest
	GetInstructionJson() *string
	SetLabelsJson(v string) *ModifyAgentProfileRequest
	GetLabelsJson() *string
	SetModel(v string) *ModifyAgentProfileRequest
	GetModel() *string
	SetModelConfig(v string) *ModifyAgentProfileRequest
	GetModelConfig() *string
	SetPrompt(v string) *ModifyAgentProfileRequest
	GetPrompt() *string
	SetPromptJson(v string) *ModifyAgentProfileRequest
	GetPromptJson() *string
	SetScenario(v string) *ModifyAgentProfileRequest
	GetScenario() *string
	SetVariablesJson(v string) *ModifyAgentProfileRequest
	GetVariablesJson() *string
}

type ModifyAgentProfileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 37ca3ca1ac4b4e57adf3da5b5d939d04
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// example:
	//
	// []
	ApiPluginJson  *string  `json:"ApiPluginJson,omitempty" xml:"ApiPluginJson,omitempty"`
	Description    *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	FaqCategoryIds []*int64 `json:"FaqCategoryIds,omitempty" xml:"FaqCategoryIds,omitempty" type:"Repeated"`
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

func (s ModifyAgentProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentProfileRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgentProfileRequest) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *ModifyAgentProfileRequest) GetApiPluginJson() *string {
	return s.ApiPluginJson
}

func (s *ModifyAgentProfileRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAgentProfileRequest) GetFaqCategoryIds() []*int64 {
	return s.FaqCategoryIds
}

func (s *ModifyAgentProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAgentProfileRequest) GetInstructionJson() *string {
	return s.InstructionJson
}

func (s *ModifyAgentProfileRequest) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *ModifyAgentProfileRequest) GetModel() *string {
	return s.Model
}

func (s *ModifyAgentProfileRequest) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *ModifyAgentProfileRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *ModifyAgentProfileRequest) GetPromptJson() *string {
	return s.PromptJson
}

func (s *ModifyAgentProfileRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ModifyAgentProfileRequest) GetVariablesJson() *string {
	return s.VariablesJson
}

func (s *ModifyAgentProfileRequest) SetAgentProfileId(v string) *ModifyAgentProfileRequest {
	s.AgentProfileId = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetApiPluginJson(v string) *ModifyAgentProfileRequest {
	s.ApiPluginJson = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetDescription(v string) *ModifyAgentProfileRequest {
	s.Description = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetFaqCategoryIds(v []*int64) *ModifyAgentProfileRequest {
	s.FaqCategoryIds = v
	return s
}

func (s *ModifyAgentProfileRequest) SetInstanceId(v string) *ModifyAgentProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetInstructionJson(v string) *ModifyAgentProfileRequest {
	s.InstructionJson = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetLabelsJson(v string) *ModifyAgentProfileRequest {
	s.LabelsJson = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetModel(v string) *ModifyAgentProfileRequest {
	s.Model = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetModelConfig(v string) *ModifyAgentProfileRequest {
	s.ModelConfig = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetPrompt(v string) *ModifyAgentProfileRequest {
	s.Prompt = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetPromptJson(v string) *ModifyAgentProfileRequest {
	s.PromptJson = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetScenario(v string) *ModifyAgentProfileRequest {
	s.Scenario = &v
	return s
}

func (s *ModifyAgentProfileRequest) SetVariablesJson(v string) *ModifyAgentProfileRequest {
	s.VariablesJson = &v
	return s
}

func (s *ModifyAgentProfileRequest) Validate() error {
	return dara.Validate(s)
}
