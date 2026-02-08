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
	// Status Code
	//
	// example:
	//
	// Ok
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	Data *CreateAgentProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message
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
	// Indicates whether the operation succeeded
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
	// Agent configuration ID
	//
	// example:
	//
	// d31794e2a51f47d2901b4094d88311d7
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// Agent configuration template ID
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
	// 1739333534000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// 这是一个智能体。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID
	//
	// example:
	//
	// 1e16c663-0339-4064-9d57-07f772e78f21
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instruction Configuration
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
	// Model ID
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Model Configuration
	//
	// example:
	//
	// ""
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// Prompt (Professional Mode)
	//
	// example:
	//
	// \\n名称：安妮 \\n身份：游戏推广员\\n背景：通过电话向玩家宣传最新福利活动，确保每一位玩家都能及时掌握信息，享受游戏的乐趣。\\n技能：作为游戏客服，向用户推荐限时游戏福利活动，强调福利亮点并告知参与流程。\\n约束条件：你需要记住你是安妮，来自于热门游戏客服团队，专注于通过电话向玩家宣传最新福利活动。\\n作为主动联系玩家的客服，需要根据玩家是否想了解活动的意图来选择是否介绍。\\n保持通话的专业性，同时语言平易近人，确保玩家轻松理解活动内容。\\n明确活动的起止时间，确保玩家不会混淆。\\n强调活动的重点奖励和参与方式，提升玩家的参与兴趣。
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Agent configuration
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
	// Scenario
	//
	// example:
	//
	// 场景
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// d13ad2d3-3fe6-4352-b38b-bd6559047de8
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Indicates whether it is a system template.
	//
	// example:
	//
	// True
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// Updated At
	//
	// example:
	//
	// 1737077564000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Variable configuration
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
