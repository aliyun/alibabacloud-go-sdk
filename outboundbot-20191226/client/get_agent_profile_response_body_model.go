// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentProfileResponseBody
	GetCode() *string
	SetData(v *GetAgentProfileResponseBodyData) *GetAgentProfileResponseBody
	GetData() *GetAgentProfileResponseBodyData
	SetHttpStatusCode(v int32) *GetAgentProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAgentProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentProfileResponseBody
	GetSuccess() *bool
}

type GetAgentProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAgentProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 7A573837-3AD3-54CF-930A-07A3287042C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentProfileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentProfileResponseBody) GetData() *GetAgentProfileResponseBodyData {
	return s.Data
}

func (s *GetAgentProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAgentProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentProfileResponseBody) SetCode(v string) *GetAgentProfileResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentProfileResponseBody) SetData(v *GetAgentProfileResponseBodyData) *GetAgentProfileResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentProfileResponseBody) SetHttpStatusCode(v int32) *GetAgentProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentProfileResponseBody) SetMessage(v string) *GetAgentProfileResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentProfileResponseBody) SetRequestId(v string) *GetAgentProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentProfileResponseBody) SetSuccess(v bool) *GetAgentProfileResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentProfileResponseBodyData struct {
	// example:
	//
	// 37ca3ca1ac4b4e57adf3da5b5d939d04
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// example:
	//
	// default-survey
	AgentProfileTemplateId *string `json:"AgentProfileTemplateId,omitempty" xml:"AgentProfileTemplateId,omitempty"`
	// agent type
	//
	// example:
	//
	// ""
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// example:
	//
	// []
	ApiPluginJson *string `json:"ApiPluginJson,omitempty" xml:"ApiPluginJson,omitempty"`
	// example:
	//
	// 1741338619000
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 7f04f92c-ccfc-4f8f-a816-6902023be5c6
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
	// {}
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// example:
	//
	// {"chatbotInstanceId":"chatbot-cn-j7eiVJZRmb","faqCategoryIds":[30000474726],"llmAgentId":"1246206","llmAgentInstanceId":"outbound_05efb75a-95df-438e-9b9b-8f2c857d5498","llmAgentKey":"d682716514814815ae77757c0bcbda01_p_outbound_public"}
	NluConfigJson *string `json:"NluConfigJson,omitempty" xml:"NluConfigJson,omitempty"`
	// example:
	//
	// ""
	Prompt     *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptJson *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	Scenario   *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// example:
	//
	// d13ad2d3-3fe6-4352-b38b-bd6559047de8
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// example:
	//
	// 1741338619000
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VariablesJson *string `json:"VariablesJson,omitempty" xml:"VariablesJson,omitempty"`
}

func (s GetAgentProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentProfileResponseBodyData) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *GetAgentProfileResponseBodyData) GetAgentProfileTemplateId() *string {
	return s.AgentProfileTemplateId
}

func (s *GetAgentProfileResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *GetAgentProfileResponseBodyData) GetApiPluginJson() *string {
	return s.ApiPluginJson
}

func (s *GetAgentProfileResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAgentProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgentProfileResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentProfileResponseBodyData) GetInstructionJson() *string {
	return s.InstructionJson
}

func (s *GetAgentProfileResponseBodyData) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *GetAgentProfileResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *GetAgentProfileResponseBodyData) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *GetAgentProfileResponseBodyData) GetNluConfigJson() *string {
	return s.NluConfigJson
}

func (s *GetAgentProfileResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *GetAgentProfileResponseBodyData) GetPromptJson() *string {
	return s.PromptJson
}

func (s *GetAgentProfileResponseBodyData) GetScenario() *string {
	return s.Scenario
}

func (s *GetAgentProfileResponseBodyData) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetAgentProfileResponseBodyData) GetSystem() *bool {
	return s.System
}

func (s *GetAgentProfileResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAgentProfileResponseBodyData) GetVariablesJson() *string {
	return s.VariablesJson
}

func (s *GetAgentProfileResponseBodyData) SetAgentProfileId(v string) *GetAgentProfileResponseBodyData {
	s.AgentProfileId = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetAgentProfileTemplateId(v string) *GetAgentProfileResponseBodyData {
	s.AgentProfileTemplateId = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetAgentType(v string) *GetAgentProfileResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetApiPluginJson(v string) *GetAgentProfileResponseBodyData {
	s.ApiPluginJson = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetCreateTime(v string) *GetAgentProfileResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetDescription(v string) *GetAgentProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetInstanceId(v string) *GetAgentProfileResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetInstructionJson(v string) *GetAgentProfileResponseBodyData {
	s.InstructionJson = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetLabelsJson(v string) *GetAgentProfileResponseBodyData {
	s.LabelsJson = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetModel(v string) *GetAgentProfileResponseBodyData {
	s.Model = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetModelConfig(v string) *GetAgentProfileResponseBodyData {
	s.ModelConfig = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetNluConfigJson(v string) *GetAgentProfileResponseBodyData {
	s.NluConfigJson = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetPrompt(v string) *GetAgentProfileResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetPromptJson(v string) *GetAgentProfileResponseBodyData {
	s.PromptJson = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetScenario(v string) *GetAgentProfileResponseBodyData {
	s.Scenario = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetScriptId(v string) *GetAgentProfileResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetSystem(v bool) *GetAgentProfileResponseBodyData {
	s.System = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetUpdateTime(v string) *GetAgentProfileResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) SetVariablesJson(v string) *GetAgentProfileResponseBodyData {
	s.VariablesJson = &v
	return s
}

func (s *GetAgentProfileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
