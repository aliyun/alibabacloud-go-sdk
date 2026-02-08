// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileId(v string) *ModifyAgentProfileShrinkRequest
	GetAgentProfileId() *string
	SetApiPluginJson(v string) *ModifyAgentProfileShrinkRequest
	GetApiPluginJson() *string
	SetDescription(v string) *ModifyAgentProfileShrinkRequest
	GetDescription() *string
	SetFaqCategoryIdsShrink(v string) *ModifyAgentProfileShrinkRequest
	GetFaqCategoryIdsShrink() *string
	SetInstanceId(v string) *ModifyAgentProfileShrinkRequest
	GetInstanceId() *string
	SetInstructionJson(v string) *ModifyAgentProfileShrinkRequest
	GetInstructionJson() *string
	SetLabelsJson(v string) *ModifyAgentProfileShrinkRequest
	GetLabelsJson() *string
	SetModel(v string) *ModifyAgentProfileShrinkRequest
	GetModel() *string
	SetModelConfig(v string) *ModifyAgentProfileShrinkRequest
	GetModelConfig() *string
	SetPrompt(v string) *ModifyAgentProfileShrinkRequest
	GetPrompt() *string
	SetPromptJson(v string) *ModifyAgentProfileShrinkRequest
	GetPromptJson() *string
	SetScenario(v string) *ModifyAgentProfileShrinkRequest
	GetScenario() *string
	SetVariablesJson(v string) *ModifyAgentProfileShrinkRequest
	GetVariablesJson() *string
}

type ModifyAgentProfileShrinkRequest struct {
	// agent profile ID
	//
	// > You can obtain the agent profile ID of an already created scenario from the ChatbotId response parameter of the DescribeScript operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 37ca3ca1ac4b4e57adf3da5b5d939d04
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// API plugin configuration
	//
	// > If you need to attach multiple APIs, retrieve the UUIDs of the APIs by calling ListApiPlugins, and enter them in the format [{"uuid":"xxx"},{"uuid":"xxxx"},...].
	//
	// example:
	//
	// [{"uuid":"a55d171da96c41a************"},{"uuid":"d41a3e5195d1402********"}]
	ApiPluginJson *string `json:"ApiPluginJson,omitempty" xml:"ApiPluginJson,omitempty"`
	// Description
	//
	// example:
	//
	// 第一版本提交审核
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// FAQ Folder ID
	FaqCategoryIdsShrink *string `json:"FaqCategoryIds,omitempty" xml:"FaqCategoryIds,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// instruction configuration
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
	// 		"name": "收号测试",
	//
	// 		"collectVoice": true,
	//
	// 		"terminator": "#"
	//
	// 	}]
	//
	// }]
	InstructionJson *string `json:"InstructionJson,omitempty" xml:"InstructionJson,omitempty"`
	// Tag Description
	//
	// example:
	//
	// [{
	//
	// 	"name": "是否满意",
	//
	// 	"description": "对介绍的游戏内容是否感兴趣",
	//
	// 	"valueList": "[\\"有兴趣\\",\\"没兴趣\\"]"
	//
	// }]
	LabelsJson *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// model ID
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Model Configuration
	//
	// example:
	//
	// {}
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// prompt (Professional mode)
	//
	// example:
	//
	// \\n名称：安妮 \\n身份：游戏推广员\\n背景：通过电话向玩家宣传最新福利活动，确保每一位玩家都能及时掌握信息，享受游戏的乐趣。\\n技能：作为游戏客服，向用户推荐限时游戏福利活动，强调福利亮点并告知参与流程。\\n约束条件：你需要记住你是安妮，来自于热门游戏客服团队，专注于通过电话向玩家宣传最新福利活动。\\n作为主动联系玩家的客服，需要根据玩家是否想了解活动的意图来选择是否介绍。\\n保持通话的专业性，同时语言平易近人，确保玩家轻松理解活动内容。\\n明确活动的起止时间，确保玩家不会混淆。\\n强调活动的重点奖励和参与方式，提升玩家的参与兴趣。
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// agent configuration
	//
	// example:
	//
	// {
	//
	// 	"name": "安妮",
	//
	// 	"gender": "女",
	//
	// 	"age": 20,
	//
	// 	"role": "游戏推广员",
	//
	// 	"communicationStyle": ["亲切", "口语化", "活泼"],
	//
	// 	"goals": "通过电话向玩家宣传最新福利活动，确保每一位玩家都能及时掌握信息，享受游戏的乐趣。",
	//
	// 	"background": "福利介绍\\n从${开始时间}开始，直至${结束时间}结束，为期一周的时间里，我们精心准备了\\"限时寻宝\\"福利活动。只要您在这段时间内每日登录游戏，即可领取登录礼包，内含珍稀材料、能量药剂和限定外观奖励\\n\\n福利亮点\\n本次活动的亮点在于，累积登录达到5天的玩家，将额外获得一份\\"神秘宝箱\\"，开启宝箱获得稀有宠物，将为您的战斗增添强大助力！千万别错过这次增强实力的绝佳机会。\\n\\n参与方式\\n登录游戏后，在活动面板中找到\\"限时寻宝\\"，点击即可领取对应的奖励。记得每天上线，奖励不等人哦！",
	//
	// 	"skills": "作为游戏客服，向用户推荐限时游戏福利活动，强调福利亮点并告知参与流程。",
	//
	// 	"workflow": "1.亲切问候与自我介绍\\n- 开场白：“尊敬的冒险者，您好！我是安妮，来自游戏客服团队。今天有幸与您连线，是想分享一项即将启动的独家福利活动，绝对会让您的探险之旅更加精彩纷呈！”\\n- 确认玩家想了解活动后，再进行福利活动的介绍\\n2.福利活动介绍\\n3.强调福利亮点\\n4.说明参与方式与提醒\\n5.鼓励分享与反馈\\n- 互动号召：“如果您觉得这个活动不错，不妨和您的公会伙伴或者游戏朋友分享这个好消息。同时，我们也非常欢迎您在活动结束后，通过游戏内置的反馈系统告诉我们您的体验感受和建议。”\\n技能6: 礼貌结束通话\\n- 结束语：“好了，尊敬的冒险者，以上就是这次活动的主要内容。希望您能在活动中满载而归，祝您在旅程中所向披靡！如果还有任何疑问，欢迎随时联系我们客服团队。感谢您的接听，期待再次为您服务，再见！”",
	//
	// 	"constraint": "你需要记住你是安妮，来自于热门游戏客服团队，专注于通过电话向玩家宣传最新福利活动。\\n作为主动联系玩家的客服，需要根据玩家是否想了解活动的意图来选择是否介绍。\\n保持通话的专业性，同时语言平易近人，确保玩家轻松理解活动内容。\\n明确活动的起止时间，确保玩家不会混淆。\\n强调活动的重点奖励和参与方式，提升玩家的参与兴趣。\\n鼓励玩家间的互动和反馈，以促进游戏社区的活跃度。\\n结束通话时保持礼貌，给玩家留下良好印象。\\n用户询问的问题与目标或流程无关时， 请礼貌拒绝，并引导到目标问题上。\\n对于不知道的问题，请不要胡乱回复，需要礼貌回复“不清楚”，并引导到目标问题上。",
	//
	// 	"openingPrompt": "尊敬的冒险者，您好！我是安妮，来自游戏客服团队。今天有幸与您连线，是想分享一项即将启动的独家福利活动，绝对会让您的探险之旅更加精彩纷呈！",
	//
	// 	"output": "",
	//
	// 	"aiHangupOutput": "",
	//
	// 	"aiSilenceTimeoutOutput": ""
	//
	// }
	PromptJson *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	// scenario
	//
	// example:
	//
	// 测试场景
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// Variable Configuration
	//
	// example:
	//
	// [{
	//
	// 	"name": "开始时间",
	//
	// 	"description": "开始时间"
	//
	// }, {
	//
	// 	"name": "结束时间",
	//
	// 	"description": "结束时间"
	//
	// }]
	VariablesJson *string `json:"VariablesJson,omitempty" xml:"VariablesJson,omitempty"`
}

func (s ModifyAgentProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgentProfileShrinkRequest) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *ModifyAgentProfileShrinkRequest) GetApiPluginJson() *string {
	return s.ApiPluginJson
}

func (s *ModifyAgentProfileShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAgentProfileShrinkRequest) GetFaqCategoryIdsShrink() *string {
	return s.FaqCategoryIdsShrink
}

func (s *ModifyAgentProfileShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAgentProfileShrinkRequest) GetInstructionJson() *string {
	return s.InstructionJson
}

func (s *ModifyAgentProfileShrinkRequest) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *ModifyAgentProfileShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ModifyAgentProfileShrinkRequest) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *ModifyAgentProfileShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *ModifyAgentProfileShrinkRequest) GetPromptJson() *string {
	return s.PromptJson
}

func (s *ModifyAgentProfileShrinkRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ModifyAgentProfileShrinkRequest) GetVariablesJson() *string {
	return s.VariablesJson
}

func (s *ModifyAgentProfileShrinkRequest) SetAgentProfileId(v string) *ModifyAgentProfileShrinkRequest {
	s.AgentProfileId = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetApiPluginJson(v string) *ModifyAgentProfileShrinkRequest {
	s.ApiPluginJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetDescription(v string) *ModifyAgentProfileShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetFaqCategoryIdsShrink(v string) *ModifyAgentProfileShrinkRequest {
	s.FaqCategoryIdsShrink = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetInstanceId(v string) *ModifyAgentProfileShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetInstructionJson(v string) *ModifyAgentProfileShrinkRequest {
	s.InstructionJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetLabelsJson(v string) *ModifyAgentProfileShrinkRequest {
	s.LabelsJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetModel(v string) *ModifyAgentProfileShrinkRequest {
	s.Model = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetModelConfig(v string) *ModifyAgentProfileShrinkRequest {
	s.ModelConfig = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetPrompt(v string) *ModifyAgentProfileShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetPromptJson(v string) *ModifyAgentProfileShrinkRequest {
	s.PromptJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetScenario(v string) *ModifyAgentProfileShrinkRequest {
	s.Scenario = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) SetVariablesJson(v string) *ModifyAgentProfileShrinkRequest {
	s.VariablesJson = &v
	return s
}

func (s *ModifyAgentProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
