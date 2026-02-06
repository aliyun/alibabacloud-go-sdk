// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAgentProfileResponseBody
	GetCode() *string
	SetData(v *ModifyAgentProfileResponseBodyData) *ModifyAgentProfileResponseBody
	GetData() *ModifyAgentProfileResponseBodyData
	SetHttpStatusCode(v int32) *ModifyAgentProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyAgentProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAgentProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAgentProfileResponseBody
	GetSuccess() *bool
}

type ModifyAgentProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ModifyAgentProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAgentProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentProfileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAgentProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAgentProfileResponseBody) GetData() *ModifyAgentProfileResponseBodyData {
	return s.Data
}

func (s *ModifyAgentProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyAgentProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAgentProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAgentProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAgentProfileResponseBody) SetCode(v string) *ModifyAgentProfileResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAgentProfileResponseBody) SetData(v *ModifyAgentProfileResponseBodyData) *ModifyAgentProfileResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAgentProfileResponseBody) SetHttpStatusCode(v int32) *ModifyAgentProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAgentProfileResponseBody) SetMessage(v string) *ModifyAgentProfileResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAgentProfileResponseBody) SetRequestId(v string) *ModifyAgentProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAgentProfileResponseBody) SetSuccess(v bool) *ModifyAgentProfileResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAgentProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAgentProfileResponseBodyData struct {
	// example:
	//
	// 3d7d253cfb77476da0cf3681bcf7b4e8
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// agent template id
	//
	// example:
	//
	// default-survey
	AgentProfileTemplateId *string `json:"AgentProfileTemplateId,omitempty" xml:"AgentProfileTemplateId,omitempty"`
	// agent type
	//
	// example:
	//
	// Human
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// example:
	//
	// []
	ApiPluginJson *string `json:"ApiPluginJson,omitempty" xml:"ApiPluginJson,omitempty"`
	// example:
	//
	// 1721356124220
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 788066f2-f160-458e-a3bb-83e1c9d5606d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// [{"type":"TransferToAgent","instructions":[{"code":"Transfer0","skillGroupId":"123","skillGroupName":"123"}],"timeoutEnable":false},{"type":"CollectNumber","instructions":[]}]
	InstructionJson *string `json:"InstructionJson,omitempty" xml:"InstructionJson,omitempty"`
	LabelsJson      *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// example:
	//
	// model_002
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// ""
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// example:
	//
	// {"chatbotInstanceId":"chatbot-cn-j7eiVJZRmb","faqCategoryIds":[30000474726],"llmAgentId":"1246206","llmAgentInstanceId":"outbound_05efb75a-95df-438e-9b9b-8f2c857d5498","llmAgentKey":"d682716514814815ae77757c0bcbda01_p_outbound_public"}
	NluConfigJson *string `json:"NluConfigJson,omitempty" xml:"NluConfigJson,omitempty"`
	Prompt        *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptJson    *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	// example:
	//
	// default
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// example:
	//
	// 3eacaec0-64ba-4008-9392-1d419b0d2673
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// true
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// example:
	//
	// 1715416630.0
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VariablesJson *string `json:"VariablesJson,omitempty" xml:"VariablesJson,omitempty"`
}

func (s ModifyAgentProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAgentProfileResponseBodyData) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *ModifyAgentProfileResponseBodyData) GetAgentProfileTemplateId() *string {
	return s.AgentProfileTemplateId
}

func (s *ModifyAgentProfileResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *ModifyAgentProfileResponseBodyData) GetApiPluginJson() *string {
	return s.ApiPluginJson
}

func (s *ModifyAgentProfileResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ModifyAgentProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ModifyAgentProfileResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAgentProfileResponseBodyData) GetInstructionJson() *string {
	return s.InstructionJson
}

func (s *ModifyAgentProfileResponseBodyData) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *ModifyAgentProfileResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ModifyAgentProfileResponseBodyData) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *ModifyAgentProfileResponseBodyData) GetNluConfigJson() *string {
	return s.NluConfigJson
}

func (s *ModifyAgentProfileResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *ModifyAgentProfileResponseBodyData) GetPromptJson() *string {
	return s.PromptJson
}

func (s *ModifyAgentProfileResponseBodyData) GetScenario() *string {
	return s.Scenario
}

func (s *ModifyAgentProfileResponseBodyData) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyAgentProfileResponseBodyData) GetSystem() *bool {
	return s.System
}

func (s *ModifyAgentProfileResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ModifyAgentProfileResponseBodyData) GetVariablesJson() *string {
	return s.VariablesJson
}

func (s *ModifyAgentProfileResponseBodyData) SetAgentProfileId(v string) *ModifyAgentProfileResponseBodyData {
	s.AgentProfileId = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetAgentProfileTemplateId(v string) *ModifyAgentProfileResponseBodyData {
	s.AgentProfileTemplateId = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetAgentType(v string) *ModifyAgentProfileResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetApiPluginJson(v string) *ModifyAgentProfileResponseBodyData {
	s.ApiPluginJson = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetCreateTime(v string) *ModifyAgentProfileResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetDescription(v string) *ModifyAgentProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetInstanceId(v string) *ModifyAgentProfileResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetInstructionJson(v string) *ModifyAgentProfileResponseBodyData {
	s.InstructionJson = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetLabelsJson(v string) *ModifyAgentProfileResponseBodyData {
	s.LabelsJson = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetModel(v string) *ModifyAgentProfileResponseBodyData {
	s.Model = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetModelConfig(v string) *ModifyAgentProfileResponseBodyData {
	s.ModelConfig = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetNluConfigJson(v string) *ModifyAgentProfileResponseBodyData {
	s.NluConfigJson = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetPrompt(v string) *ModifyAgentProfileResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetPromptJson(v string) *ModifyAgentProfileResponseBodyData {
	s.PromptJson = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetScenario(v string) *ModifyAgentProfileResponseBodyData {
	s.Scenario = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetScriptId(v string) *ModifyAgentProfileResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetSystem(v bool) *ModifyAgentProfileResponseBodyData {
	s.System = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetUpdateTime(v string) *ModifyAgentProfileResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) SetVariablesJson(v string) *ModifyAgentProfileResponseBodyData {
	s.VariablesJson = &v
	return s
}

func (s *ModifyAgentProfileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
