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
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAgentProfilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 1721701525327
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// e49ad122-15a1-443a-a102-84a78a93be79
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
	// “”
	Prompt     *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptJson *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	Scenario   *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// example:
	//
	// 43972417-0657-452a-81c2-c59d8245a9b2
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// example:
	//
	// 1721701525327
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
