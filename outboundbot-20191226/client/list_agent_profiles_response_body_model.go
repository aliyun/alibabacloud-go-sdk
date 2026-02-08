// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentProfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentProfilesResponseBody
	GetCode() *string
	SetData(v []*ListAgentProfilesResponseBodyData) *ListAgentProfilesResponseBody
	GetData() []*ListAgentProfilesResponseBodyData
	SetHttpStatusCode(v int32) *ListAgentProfilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAgentProfilesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAgentProfilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentProfilesResponseBody
	GetSuccess() *bool
}

type ListAgentProfilesResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	Data []*ListAgentProfilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the return result.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgentProfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentProfilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentProfilesResponseBody) GetData() []*ListAgentProfilesResponseBodyData {
	return s.Data
}

func (s *ListAgentProfilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentProfilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentProfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentProfilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentProfilesResponseBody) SetCode(v string) *ListAgentProfilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentProfilesResponseBody) SetData(v []*ListAgentProfilesResponseBodyData) *ListAgentProfilesResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentProfilesResponseBody) SetHttpStatusCode(v int32) *ListAgentProfilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentProfilesResponseBody) SetMessage(v string) *ListAgentProfilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentProfilesResponseBody) SetRequestId(v string) *ListAgentProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentProfilesResponseBody) SetSuccess(v bool) *ListAgentProfilesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentProfilesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentProfilesResponseBodyData struct {
	// Agent configuration ID
	//
	// example:
	//
	// 37ca3ca1ac4b4e57adf3da5b5d939d04
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// Agent template ID
	//
	// example:
	//
	// default-survey
	AgentProfileTemplateId *string `json:"AgentProfileTemplateId,omitempty" xml:"AgentProfileTemplateId,omitempty"`
	// agent type
	//
	// example:
	//
	// “”
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 1721701525327
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// 这是一个充满智慧的智能体。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID
	//
	// example:
	//
	// e49ad122-15a1-443a-a102-84a78a93be79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instruction configuration
	//
	// example:
	//
	// [{"type":"TransferToAgent","instructions":[{"code":"Transfer0","skillGroupId":"123","skillGroupName":"123"}],"timeoutEnable":false},{"type":"CollectNumber","instructions":[]}]
	InstructionJson *string `json:"InstructionJson,omitempty" xml:"InstructionJson,omitempty"`
	// Tag Description
	//
	// example:
	//
	// [{\\"name\\":\\"是否送达\\",\\"description\\":\\"购买的家电是否已经送达\\",\\"valueList\\":\\"[\\\\\\"是\\\\\\",\\\\\\"否\\\\\\"]\\"},{\\"name\\":\\"预约上门时间\\",\\"description\\":\\"收集客户期望的上门安装时间\\",\\"valueList\\":\\"[]\\"}]
	LabelsJson *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// Model ID
	//
	// example:
	//
	// model_002
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Model Configuration
	//
	// example:
	//
	// {}
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// Prompt (Professional Mode)
	//
	// example:
	//
	// “”
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Agent configuration
	//
	// example:
	//
	// {"name":"小x","gender":"男","age":18,"role":"游戏推广员","communicationStyle":["亲切"],"goals":"你好","background":"不是很好","openingPrompt":"你好，我是xxx"}
	PromptJson *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	// Scenario
	//
	// example:
	//
	// 测试场景
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// 43972417-0657-452a-81c2-c59d8245a9b2
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Indicates whether it is a system template
	//
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// Updated At
	//
	// example:
	//
	// 1721701525327
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Variable configuration
	//
	// example:
	//
	// [{\\"name\\":\\"name\\",\\"description\\":\\"客户姓名\\"},{\\"name\\":\\"gender\\",\\"description\\":\\"客户性别\\"}]
	VariablesJson *string `json:"VariablesJson,omitempty" xml:"VariablesJson,omitempty"`
}

func (s ListAgentProfilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentProfilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentProfilesResponseBodyData) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *ListAgentProfilesResponseBodyData) GetAgentProfileTemplateId() *string {
	return s.AgentProfileTemplateId
}

func (s *ListAgentProfilesResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *ListAgentProfilesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAgentProfilesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListAgentProfilesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentProfilesResponseBodyData) GetInstructionJson() *string {
	return s.InstructionJson
}

func (s *ListAgentProfilesResponseBodyData) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *ListAgentProfilesResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ListAgentProfilesResponseBodyData) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *ListAgentProfilesResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *ListAgentProfilesResponseBodyData) GetPromptJson() *string {
	return s.PromptJson
}

func (s *ListAgentProfilesResponseBodyData) GetScenario() *string {
	return s.Scenario
}

func (s *ListAgentProfilesResponseBodyData) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListAgentProfilesResponseBodyData) GetSystem() *bool {
	return s.System
}

func (s *ListAgentProfilesResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAgentProfilesResponseBodyData) GetVariablesJson() *string {
	return s.VariablesJson
}

func (s *ListAgentProfilesResponseBodyData) SetAgentProfileId(v string) *ListAgentProfilesResponseBodyData {
	s.AgentProfileId = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetAgentProfileTemplateId(v string) *ListAgentProfilesResponseBodyData {
	s.AgentProfileTemplateId = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetAgentType(v string) *ListAgentProfilesResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetCreateTime(v string) *ListAgentProfilesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetDescription(v string) *ListAgentProfilesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetInstanceId(v string) *ListAgentProfilesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetInstructionJson(v string) *ListAgentProfilesResponseBodyData {
	s.InstructionJson = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetLabelsJson(v string) *ListAgentProfilesResponseBodyData {
	s.LabelsJson = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetModel(v string) *ListAgentProfilesResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetModelConfig(v string) *ListAgentProfilesResponseBodyData {
	s.ModelConfig = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetPrompt(v string) *ListAgentProfilesResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetPromptJson(v string) *ListAgentProfilesResponseBodyData {
	s.PromptJson = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetScenario(v string) *ListAgentProfilesResponseBodyData {
	s.Scenario = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetScriptId(v string) *ListAgentProfilesResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetSystem(v bool) *ListAgentProfilesResponseBodyData {
	s.System = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetUpdateTime(v string) *ListAgentProfilesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) SetVariablesJson(v string) *ListAgentProfilesResponseBodyData {
	s.VariablesJson = &v
	return s
}

func (s *ListAgentProfilesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
