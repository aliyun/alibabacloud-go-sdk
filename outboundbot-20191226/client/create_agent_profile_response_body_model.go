// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentProfileResponseBody
	GetCode() *string
	SetData(v *CreateAgentProfileResponseBodyData) *CreateAgentProfileResponseBody
	GetData() *CreateAgentProfileResponseBodyData
	SetHttpStatusCode(v int32) *CreateAgentProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAgentProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentProfileResponseBody
	GetSuccess() *bool
}

type CreateAgentProfileResponseBody struct {
	// example:
	//
	// Ok
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateAgentProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateAgentProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentProfileResponseBody) GetData() *CreateAgentProfileResponseBodyData {
	return s.Data
}

func (s *CreateAgentProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAgentProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentProfileResponseBody) SetCode(v string) *CreateAgentProfileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentProfileResponseBody) SetData(v *CreateAgentProfileResponseBodyData) *CreateAgentProfileResponseBody {
	s.Data = v
	return s
}

func (s *CreateAgentProfileResponseBody) SetHttpStatusCode(v int32) *CreateAgentProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAgentProfileResponseBody) SetMessage(v string) *CreateAgentProfileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentProfileResponseBody) SetRequestId(v string) *CreateAgentProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentProfileResponseBody) SetSuccess(v bool) *CreateAgentProfileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentProfileResponseBodyData struct {
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
	// “”
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// example:
	//
	// 1739333534000
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1e16c663-0339-4064-9d57-07f772e78f21
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// [{\\"type\\":\\"TransferToAgent\\",\\"instructions\\":[{\\"code\\":\\"Transfer0\\",\\"skillGroupId\\":\\"123\\",\\"skillGroupName\\":\\"123\\"}]
	InstructionJson *string `json:"InstructionJson,omitempty" xml:"InstructionJson,omitempty"`
	LabelsJson      *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// ""
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// example:
	//
	// prompt
	Prompt     *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptJson *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	Scenario   *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// example:
	//
	// d13ad2d3-3fe6-4352-b38b-bd6559047de8
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// True
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// example:
	//
	// 1737077564000
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VariablesJson *string `json:"VariablesJson,omitempty" xml:"VariablesJson,omitempty"`
}

func (s CreateAgentProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgentProfileResponseBodyData) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *CreateAgentProfileResponseBodyData) GetAgentProfileTemplateId() *string {
	return s.AgentProfileTemplateId
}

func (s *CreateAgentProfileResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateAgentProfileResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAgentProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentProfileResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAgentProfileResponseBodyData) GetInstructionJson() *string {
	return s.InstructionJson
}

func (s *CreateAgentProfileResponseBodyData) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *CreateAgentProfileResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *CreateAgentProfileResponseBodyData) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *CreateAgentProfileResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateAgentProfileResponseBodyData) GetPromptJson() *string {
	return s.PromptJson
}

func (s *CreateAgentProfileResponseBodyData) GetScenario() *string {
	return s.Scenario
}

func (s *CreateAgentProfileResponseBodyData) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateAgentProfileResponseBodyData) GetSystem() *bool {
	return s.System
}

func (s *CreateAgentProfileResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateAgentProfileResponseBodyData) GetVariablesJson() *string {
	return s.VariablesJson
}

func (s *CreateAgentProfileResponseBodyData) SetAgentProfileId(v string) *CreateAgentProfileResponseBodyData {
	s.AgentProfileId = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetAgentProfileTemplateId(v string) *CreateAgentProfileResponseBodyData {
	s.AgentProfileTemplateId = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetAgentType(v string) *CreateAgentProfileResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetCreateTime(v string) *CreateAgentProfileResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetDescription(v string) *CreateAgentProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetInstanceId(v string) *CreateAgentProfileResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetInstructionJson(v string) *CreateAgentProfileResponseBodyData {
	s.InstructionJson = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetLabelsJson(v string) *CreateAgentProfileResponseBodyData {
	s.LabelsJson = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetModel(v string) *CreateAgentProfileResponseBodyData {
	s.Model = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetModelConfig(v string) *CreateAgentProfileResponseBodyData {
	s.ModelConfig = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetPrompt(v string) *CreateAgentProfileResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetPromptJson(v string) *CreateAgentProfileResponseBodyData {
	s.PromptJson = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetScenario(v string) *CreateAgentProfileResponseBodyData {
	s.Scenario = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetScriptId(v string) *CreateAgentProfileResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetSystem(v bool) *CreateAgentProfileResponseBodyData {
	s.System = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetUpdateTime(v string) *CreateAgentProfileResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) SetVariablesJson(v string) *CreateAgentProfileResponseBodyData {
	s.VariablesJson = &v
	return s
}

func (s *CreateAgentProfileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
