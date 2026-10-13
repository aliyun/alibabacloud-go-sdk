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
	// The agent profile template ID (unique value).
	//
	// This parameter is required.
	//
	// example:
	//
	// default-survey
	AgentProfileTemplateId *string `json:"AgentProfileTemplateId,omitempty" xml:"AgentProfileTemplateId,omitempty"`
	// The app_ip (system field). This parameter is optional.
	//
	// example:
	//
	// 127.0.0.1
	AppIp *string `json:"AppIp,omitempty" xml:"AppIp,omitempty"`
	// The description.
	//
	// example:
	//
	// Describe the features of the current agent
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The FAQ category IDs.
	//
	// example:
	//
	// [30000474726]
	FaqCategoryIds *string `json:"FaqCategoryIds,omitempty" xml:"FaqCategoryIds,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
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
	// The model ID (unique value).
	//
	// This parameter is required.
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// **[Deprecated]*	- The model configuration.
	//
	// example:
	//
	// ""
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// The agent configuration in professional mode.
	//
	// example:
	//
	// \\nName: Annie \\nIdentity: Game Promoter\\nBackground: Promote the latest welfare events to players over the phone, ensuring that every player can keep up with the information in a timely manner and enjoy the game.\\nSkills: As a game customer service representative, recommend limited-time game welfare events to users, emphasize the event highlights, and inform them of the participation process.\\nConstraints: You need to remember that you are Annie, from the customer service team of a popular game, focusing on promoting the latest welfare events to players over the phone.\\nAs a customer service representative who proactively contacts players, you need to choose whether to introduce the event based on the player\\"s intention to learn about it.\\nMaintain the professionalism of the call, while keeping the language approachable to ensure that players can easily understand the event content.\\nClarify the start and end times of the event to ensure that players do not get confused.\\nEmphasize the key rewards and participation methods of the event to increase players\\" interest in participating
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The agent configuration in normal mode.
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
	// Test scenario
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aa279896-64a6-4182-864c-4f2b04ec8d17
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
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
