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
	// The agent profile ID.
	//
	// > You can obtain the agent profile ID of a created scenario from the ChatbotId response parameter of DescribeScript.
	//
	// This parameter is required.
	//
	// example:
	//
	// 37ca3ca1ac4b4e57adf3da5b5d939d04
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// The API plugin configuration.
	//
	// > To bind multiple APIs, obtain the UUIDs of multiple APIs through ListApiPlugins, and fill them in the format of [{"uuid":"xxx"},{"uuid":"xxxx"},....].
	//
	// example:
	//
	// [{"uuid":"a55d171da96c41a************"},{"uuid":"d41a3e5195d1402********"}]
	ApiPluginJson *string `json:"ApiPluginJson,omitempty" xml:"ApiPluginJson,omitempty"`
	// The description.
	//
	// example:
	//
	// First version submitted for review
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The FAQ category IDs.
	FaqCategoryIds []*int64 `json:"FaqCategoryIds,omitempty" xml:"FaqCategoryIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
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
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The model configuration.
	//
	// example:
	//
	// {}
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
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
	// Test scenario
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
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
