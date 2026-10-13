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
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ModifyAgentProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request was successful.
	//
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
	// The agent profile ID.
	//
	// example:
	//
	// 3d7d253cfb77476da0cf3681bcf7b4e8
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// Agent Template ID
	//
	// example:
	//
	// default-survey
	AgentProfileTemplateId *string `json:"AgentProfileTemplateId,omitempty" xml:"AgentProfileTemplateId,omitempty"`
	// Agent Type
	//
	// example:
	//
	// Human
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The API plugin configuration.
	//
	// example:
	//
	// []
	ApiPluginJson *string `json:"ApiPluginJson,omitempty" xml:"ApiPluginJson,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1721356124220
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a large model scenario
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 788066f2-f160-458e-a3bb-83e1c9d5606d
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
	// The label description.
	//
	// example:
	//
	// [{
	//
	// 	"name": "Satisfied",
	//
	// 	"description": "Whether interested in the introduced game content",
	//
	// 	"valueList": "[\\"Interested\\",\\"Not interested\\"]"
	//
	// }]
	LabelsJson *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// The model ID.
	//
	// example:
	//
	// model_002
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The model configuration.
	//
	// example:
	//
	// ""
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// The NLU configuration. Parameter settings:
	//
	// - chatbotInstanceId: the Chatbot large model robot instance ID.
	//
	// - faqCategoryIds: the attached FAQ category IDs.
	//
	// - llmAgentId: the Chatbot large model workspace ID.
	//
	// - llmAgentKey: the Chatbot large model workspace key.
	//
	// - llmAgentInstanceId: the Chatbot large model workspace instance ID.
	//
	// example:
	//
	// {
	//
	// 	"chatbotInstanceId": "chatbot-cn-j7eiVJZRmb",
	//
	// 	"faqCategoryIds": [30000474726],
	//
	// 	"llmAgentId": "1246206",
	//
	// 	"llmAgentInstanceId": "outbound_05efb75a-95df-438e-9b9b-8f2c857d5498",
	//
	// 	"llmAgentKey": "d682716514814815ae77757c0bcbda01_p_outbound_public"
	//
	// }
	NluConfigJson *string `json:"NluConfigJson,omitempty" xml:"NluConfigJson,omitempty"`
	// The prompt in professional mode.
	//
	// example:
	//
	// \\nName: Annie \\nIdentity: Game Promoter\\nBackground: Promote the latest promotional events to players over the phone, ensuring that every player stays informed and enjoys the game.\\nSkills: As a game customer service representative, recommend limited-time game promotional events to users, highlight the event benefits, and inform them of the participation process.\\nConstraints: You need to remember that you are Annie, from the customer service team of a popular game, focusing on promoting the latest promotional events to players over the phone.\\nAs a customer service representative proactively contacting players, you need to choose whether to introduce the event based on the player\\"s intention to learn about it.\\nMaintain the professionalism of the call while keeping the language approachable to ensure players can easily understand the event content.\\nClarify the start and end times of the event to ensure players do not get confused.\\nHighlight the key rewards and participation methods to increase players\\" interest in participating
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
	// 	"goals": "Promote the latest promotional events to players over the phone, ensuring that every player stays informed and enjoys the game.",
	//
	// 	"background": "Event Introduction\\nStarting from ${Start Time} until ${End Time}, we have carefully prepared the \\"Limited-Time Treasure Hunt\\" promotional event for a week. As long as you log in to the game daily during this period, you can claim a login gift pack containing rare materials, energy potions, and exclusive cosmetic rewards.\\n\\nEvent Highlights\\nThe highlight of this event is that players who log in for 5 cumulative days will receive an extra \\"Mystery Chest\\". Opening the chest to obtain a rare pet will add a powerful boost to your battles! Do not miss this excellent opportunity to enhance your strength.\\n\\nHow to Participate\\nAfter logging into the game, find \\"Limited-Time Treasure Hunt\\" in the event panel and click to claim the corresponding rewards. Remember to log in every day, the rewards wait for no one!",
	//
	// 	"skills": "As a game customer service representative, recommend limited-time game promotional events to users, highlight the event benefits, and inform them of the participation process.",
	//
	// 	"workflow": "1. Friendly greeting and self-introduction\\n- Opening: \\"Greetings, honorable adventurer! I am Annie from the game customer service team. I am honored to connect with you today to share an exclusive promotional event that is about to launch, which will definitely make your adventure more exciting!\\"\\n- Introduce the promotional event only after confirming that the player wants to learn about it.\\n2. Promotional event introduction\\n3. Highlight event benefits\\n4. Explain how to participate and remind\\n5. Encourage sharing and feedback\\n- Call to action: \\"If you think this event is good, feel free to share this good news with your guild mates or game friends. At the same time, we also welcome you to share your experience and suggestions through the in-game feedback system after the event ends.\\"\\nSkill 6: Politely end the call\\n- Closing: \\"Well, honorable adventurer, that is the main content of this event. I hope you will return with full rewards from the event, and I wish you invincible in your journey! If you have any questions, feel free to contact our customer service team. Thank you for taking my call, I look forward to serving you again, goodbye!\\"",
	//
	// 	"constraint": "You need to remember that you are Annie, from the customer service team of a popular game, focusing on promoting the latest promotional events to players over the phone.\\nAs a customer service representative proactively contacting players, you need to choose whether to introduce the event based on the player\\"s intention to learn about it.\\nMaintain the professionalism of the call while keeping the language approachable to ensure players can easily understand the event content.\\nClarify the start and end times of the event to ensure players do not get confused.\\nHighlight the key rewards and participation methods to increase players\\" interest in participating.\\nEncourage interaction and feedback among players to promote the activity of the game community.\\nMaintain politeness when ending the call to leave a good impression on the player.\\nWhen the user asks questions unrelated to the goal or process, politely refuse and guide them back to the target topic.\\nFor questions you do not know the answer to, do not reply randomly. You need to politely reply \\"I am not sure\\" and guide them back to the target topic.",
	//
	// 	"openingPrompt": "Greetings, honorable adventurer! I am Annie from the game customer service team. I am honored to connect with you today to share an exclusive promotional event that is about to launch, which will definitely make your adventure more exciting!",
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
	// default
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// 3eacaec0-64ba-4008-9392-1d419b0d2673
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Indicates whether the template is a system template.
	//
	// example:
	//
	// true
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// The update time. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1715416630.0
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
