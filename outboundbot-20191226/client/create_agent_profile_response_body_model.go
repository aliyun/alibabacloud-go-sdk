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
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateAgentProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
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
	// The agent profile ID.
	//
	// example:
	//
	// d31794e2a51f47d2901b4094d88311d7
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// The agent profile template ID.
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
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1739333534000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// This is an agent
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1e16c663-0339-4064-9d57-07f772e78f21
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instruction configuration.
	//
	// example:
	//
	// [{
	//
	// 	"type": "TransferToAgent",
	//
	// 	"instructions": [{
	//
	// 		"code": "Transfer0",
	//
	// 		"skillGroupId": "testcode",
	//
	// 		"skillGroupName": "skillgroup"
	//
	// 	}],
	//
	// 	"timeoutEnable": false
	//
	// }, {
	//
	// 	"type": "CollectNumber",
	//
	// 	"instructions": [{
	//
	// 		"code": "DTMF0",
	//
	// 		"name": "Number collection test",
	//
	// 		"collectVoice": true,
	//
	// 		"terminator": "#"
	//
	// 	}]
	//
	// }]
	InstructionJson *string `json:"InstructionJson,omitempty" xml:"InstructionJson,omitempty"`
	// The label descriptions.
	//
	// example:
	//
	// [{
	//
	// 	"name": "Satisfaction",
	//
	// 	"description": "Whether you are interested in the introduced game content",
	//
	// 	"valueList": "[\\"Interested\\",\\"Not interested\\"]"
	//
	// }]
	LabelsJson *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// The model ID.
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The model configuration.
	//
	// example:
	//
	// ""
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// The prompt in professional mode.
	//
	// example:
	//
	// \\nName: Annie \\nIdentity: Game Promoter\\nBackground: Promote the latest welfare events to players over the phone, ensuring that every player can keep up with the information in a timely manner and enjoy the game.\\nSkills: As a game customer service representative, recommend limited-time game welfare events to users, emphasize the event highlights, and inform them of the participation process.\\nConstraints: You need to remember that you are Annie, from the customer service team of a popular game, focusing on promoting the latest welfare events to players over the phone.\\nAs a customer service representative who proactively contacts players, you need to choose whether to introduce the event based on the player\\"s intention to learn about it.\\nMaintain the professionalism of the call, while keeping the language approachable to ensure that players can easily understand the event content.\\nClarify the start and end times of the event to ensure that players do not get confused.\\nEmphasize the key rewards and participation methods of the event to increase players\\" interest in participating
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The agent configuration.
	//
	// example:
	//
	// {
	//
	// 	"name": "Annie",
	//
	// 	"gender": "Female",
	//
	// 	"age": 20,
	//
	// 	"role": "Game Promoter",
	//
	// 	"communicationStyle": ["Friendly", "Conversational", "Lively"],
	//
	// 	"goals": "Promote the latest welfare events to players over the phone, ensuring that every player can keep up with the information in a timely manner and enjoy the game.",
	//
	// 	"background": "Event Introduction\\nFrom ${Start Time} to ${End Time}, we have carefully prepared the \\"Limited-Time Treasure Hunt\\" welfare event for a week. As long as you log in to the game every day during this period, you can claim a login gift pack containing rare materials, energy potions, and limited-edition appearance rewards.\\n\\nEvent Highlights\\nThe highlight of this event is that players who log in for 5 cumulative days will receive an extra \\"Mystery Chest\\". Opening the chest to get a rare pet will add a powerful boost to your battles! Don\\"t miss this great opportunity to enhance your strength.\\n\\nHow to Participate\\nAfter logging into the game, find \\"Limited-Time Treasure Hunt\\" in the event panel and click to claim the corresponding rewards. Remember to log in every day, the rewards won\\"t wait!",
	//
	// 	"skills": "As a game customer service representative, recommend limited-time game welfare events to users, emphasize the event highlights, and inform them of the participation process.",
	//
	// 	"workflow": "1. Friendly greeting and self-introduction\\n- Opening: \\"Hello, respected adventurer! I am Annie from the game customer service team. I am honored to connect with you today to share an exclusive welfare event that is about to launch, which will definitely make your adventure more exciting!\\"\\n- Introduce the welfare event only after confirming that the player wants to learn about it.\\n2. Welfare event introduction\\n3. Emphasize event highlights\\n4. Explain how to participate and remind\\n5. Encourage sharing and feedback\\n- Call to action: \\"If you think this event is good, feel free to share the good news with your guild partners or game friends. At the same time, we also highly welcome you to tell us your experience and suggestions through the built-in feedback system of the game after the event ends.\\"\\nSkill 6: Politely end the call\\n- Closing: \\"Well, respected adventurer, the above is the main content of this event. I hope you can return with a full load from the event, and I wish you invincible in your journey! If you have any questions, feel free to contact our customer service team. Thank you for answering, look forward to serving you again, goodbye!\\"",
	//
	// 	"constraint": "You need to remember that you are Annie, from the customer service team of a popular game, focusing on promoting the latest welfare events to players over the phone.\\nAs a customer service representative who proactively contacts players, you need to choose whether to introduce the event based on the player\\"s intention to learn about it.\\nMaintain the professionalism of the call, while keeping the language approachable to ensure that players can easily understand the event content.\\nClarify the start and end times of the event to ensure that players do not get confused.\\nEmphasize the key rewards and participation methods of the event to increase players\\" interest in participating.\\nEncourage interaction and feedback among players to promote the activity of the game community.\\nMaintain politeness when ending the call to leave a good impression on the player.\\nWhen the user asks questions that are irrelevant to the goal or process, politely refuse and guide them back to the target question.\\nFor questions you do not know, do not reply randomly. You need to politely reply \\"I am not sure\\" and guide them back to the target question.",
	//
	// 	"openingPrompt": "Hello, respected adventurer! I am Annie from the game customer service team. I am honored to connect with you today to share an exclusive welfare event that is about to launch, which will definitely make your adventure more exciting!",
	//
	// 	"output": "",
	//
	// 	"aiHangupOutput": "",
	//
	// 	"aiSilenceTimeoutOutput": ""
	//
	// }
	PromptJson *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	// The scenario.
	//
	// example:
	//
	// Scenario
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// d13ad2d3-3fe6-4352-b38b-bd6559047de8
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Indicates whether the template is a system template.
	//
	// example:
	//
	// True
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// The update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1737077564000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The variable configuration.
	//
	// example:
	//
	// [{
	//
	// 	"name": "Start Time",
	//
	// 	"description": "Start time"
	//
	// }, {
	//
	// 	"name": "End Time",
	//
	// 	"description": "End time"
	//
	// }]
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
