// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *ChatResponseBody
	GetMessageId() *string
	SetMessages(v []*ChatResponseBodyMessages) *ChatResponseBody
	GetMessages() []*ChatResponseBodyMessages
	SetQuerySegList(v []*string) *ChatResponseBody
	GetQuerySegList() []*string
	SetRequestId(v string) *ChatResponseBody
	GetRequestId() *string
	SetSessionId(v string) *ChatResponseBody
	GetSessionId() *string
}

type ChatResponseBody struct {
	// The unique message ID.
	//
	// example:
	//
	// A2315C4B-A872-5DEE-9DAD-D73B194A4AEC
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// A list of message objects.
	Messages []*ChatResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The words segmented from the query. This field may be empty.
	QuerySegList []*string `json:"QuerySegList,omitempty" xml:"QuerySegList,omitempty" type:"Repeated"`
	// The unique request ID.
	//
	// example:
	//
	// A2315C4B-A872-5DEE-9DAD-D73B194A4AEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unique session ID.
	//
	// example:
	//
	// a6f216a0685c4c8baa0e8beb6d5ec6db
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *ChatResponseBody) GetMessages() []*ChatResponseBodyMessages {
	return s.Messages
}

func (s *ChatResponseBody) GetQuerySegList() []*string {
	return s.QuerySegList
}

func (s *ChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatResponseBody) SetMessageId(v string) *ChatResponseBody {
	s.MessageId = &v
	return s
}

func (s *ChatResponseBody) SetMessages(v []*ChatResponseBodyMessages) *ChatResponseBody {
	s.Messages = v
	return s
}

func (s *ChatResponseBody) SetQuerySegList(v []*string) *ChatResponseBody {
	s.QuerySegList = v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBody) SetSessionId(v string) *ChatResponseBody {
	s.SessionId = &v
	return s
}

func (s *ChatResponseBody) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatResponseBodyMessages struct {
	// Indicates the source of the recommended answer if `AnswerType` is `Recommend`.
	//
	// example:
	//
	// KNOWLEDGE
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// Type of the message.
	//
	// example:
	//
	// Text
	AnswerType *string `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	// Contains the `Knowledge` object if `AnswerType` is `Knowledge`.
	Knowledge *ChatResponseBodyMessagesKnowledge `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
	// Contains a list of `Recommend` objects if `AnswerType` is `Recommend`.
	Recommends []*ChatResponseBodyMessagesRecommends `json:"Recommends,omitempty" xml:"Recommends,omitempty" type:"Repeated"`
	// Contains the `Text` object if `AnswerType` is `Text`.
	Text *ChatResponseBodyMessagesText `json:"Text,omitempty" xml:"Text,omitempty" type:"Struct"`
	// The title of the clarification question for text-based chat scenarios.
	//
	// example:
	//
	// 请问您想咨询的是哪个投保年龄区间呢？
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The clarification content for voice-based scenarios.
	//
	// example:
	//
	// 请问你说的是公积金查询，还是公积金提取
	VoiceTitle *string `json:"VoiceTitle,omitempty" xml:"VoiceTitle,omitempty"`
}

func (s ChatResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessages) GetAnswerSource() *string {
	return s.AnswerSource
}

func (s *ChatResponseBodyMessages) GetAnswerType() *string {
	return s.AnswerType
}

func (s *ChatResponseBodyMessages) GetKnowledge() *ChatResponseBodyMessagesKnowledge {
	return s.Knowledge
}

func (s *ChatResponseBodyMessages) GetRecommends() []*ChatResponseBodyMessagesRecommends {
	return s.Recommends
}

func (s *ChatResponseBodyMessages) GetText() *ChatResponseBodyMessagesText {
	return s.Text
}

func (s *ChatResponseBodyMessages) GetTitle() *string {
	return s.Title
}

func (s *ChatResponseBodyMessages) GetVoiceTitle() *string {
	return s.VoiceTitle
}

func (s *ChatResponseBodyMessages) SetAnswerSource(v string) *ChatResponseBodyMessages {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessages) SetAnswerType(v string) *ChatResponseBodyMessages {
	s.AnswerType = &v
	return s
}

func (s *ChatResponseBodyMessages) SetKnowledge(v *ChatResponseBodyMessagesKnowledge) *ChatResponseBodyMessages {
	s.Knowledge = v
	return s
}

func (s *ChatResponseBodyMessages) SetRecommends(v []*ChatResponseBodyMessagesRecommends) *ChatResponseBodyMessages {
	s.Recommends = v
	return s
}

func (s *ChatResponseBodyMessages) SetText(v *ChatResponseBodyMessagesText) *ChatResponseBodyMessages {
	s.Text = v
	return s
}

func (s *ChatResponseBodyMessages) SetTitle(v string) *ChatResponseBodyMessages {
	s.Title = &v
	return s
}

func (s *ChatResponseBodyMessages) SetVoiceTitle(v string) *ChatResponseBodyMessages {
	s.VoiceTitle = &v
	return s
}

func (s *ChatResponseBodyMessages) Validate() error {
	if s.Knowledge != nil {
		if err := s.Knowledge.Validate(); err != nil {
			return err
		}
	}
	if s.Recommends != nil {
		for _, item := range s.Recommends {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Text != nil {
		if err := s.Text.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatResponseBodyMessagesKnowledge struct {
	// The source of the answer.
	//
	// `KnowledgeBase`: The answer is from the knowledge base.
	//
	// example:
	//
	// KnowledgeBase
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The category of the knowledge entry.
	//
	// example:
	//
	// 公积金
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the matched knowledge entry.
	//
	// example:
	//
	// 公积金提取，请在首页搜索公积金提取，提交办事的表单
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Content format of the answer.
	//
	// example:
	//
	// PLAIN_TEXT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The hit statement matching the query.
	//
	// example:
	//
	// 公积金
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// ID of the matched knowledge entry in the knowledge base.
	//
	// example:
	//
	// 735898
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// A list of related knowledge objects.
	RelatedKnowledges []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges `json:"RelatedKnowledges,omitempty" xml:"RelatedKnowledges,omitempty" type:"Repeated"`
	// The confidence score.
	//
	// example:
	//
	// 0.998
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// Summary of the matched knowledge entry.
	//
	// example:
	//
	// 公积金提取
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The title of the matched knowledge entry.
	//
	// example:
	//
	// 公积金提取
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ChatResponseBodyMessagesKnowledge) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyMessagesKnowledge) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesKnowledge) GetAnswerSource() *string {
	return s.AnswerSource
}

func (s *ChatResponseBodyMessagesKnowledge) GetCategory() *string {
	return s.Category
}

func (s *ChatResponseBodyMessagesKnowledge) GetContent() *string {
	return s.Content
}

func (s *ChatResponseBodyMessagesKnowledge) GetContentType() *string {
	return s.ContentType
}

func (s *ChatResponseBodyMessagesKnowledge) GetHitStatement() *string {
	return s.HitStatement
}

func (s *ChatResponseBodyMessagesKnowledge) GetId() *string {
	return s.Id
}

func (s *ChatResponseBodyMessagesKnowledge) GetRelatedKnowledges() []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges {
	return s.RelatedKnowledges
}

func (s *ChatResponseBodyMessagesKnowledge) GetScore() *float64 {
	return s.Score
}

func (s *ChatResponseBodyMessagesKnowledge) GetSummary() *string {
	return s.Summary
}

func (s *ChatResponseBodyMessagesKnowledge) GetTitle() *string {
	return s.Title
}

func (s *ChatResponseBodyMessagesKnowledge) SetAnswerSource(v string) *ChatResponseBodyMessagesKnowledge {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetCategory(v string) *ChatResponseBodyMessagesKnowledge {
	s.Category = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetContent(v string) *ChatResponseBodyMessagesKnowledge {
	s.Content = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetContentType(v string) *ChatResponseBodyMessagesKnowledge {
	s.ContentType = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetHitStatement(v string) *ChatResponseBodyMessagesKnowledge {
	s.HitStatement = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetId(v string) *ChatResponseBodyMessagesKnowledge {
	s.Id = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetRelatedKnowledges(v []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges) *ChatResponseBodyMessagesKnowledge {
	s.RelatedKnowledges = v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetScore(v float64) *ChatResponseBodyMessagesKnowledge {
	s.Score = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetSummary(v string) *ChatResponseBodyMessagesKnowledge {
	s.Summary = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetTitle(v string) *ChatResponseBodyMessagesKnowledge {
	s.Title = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) Validate() error {
	if s.RelatedKnowledges != nil {
		for _, item := range s.RelatedKnowledges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatResponseBodyMessagesKnowledgeRelatedKnowledges struct {
	// The ID of the related knowledge entry.
	//
	// example:
	//
	// 735899
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The title of the related knowledge entry.
	//
	// example:
	//
	// 公积金查询
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ChatResponseBodyMessagesKnowledgeRelatedKnowledges) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyMessagesKnowledgeRelatedKnowledges) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) GetTitle() *string {
	return s.Title
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) SetKnowledgeId(v string) *ChatResponseBodyMessagesKnowledgeRelatedKnowledges {
	s.KnowledgeId = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) SetTitle(v string) *ChatResponseBodyMessagesKnowledgeRelatedKnowledges {
	s.Title = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyMessagesRecommends struct {
	// Source of the clarification.
	//
	// example:
	//
	// KNOWLEDGE
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The knowledge ID for the clarification.
	//
	// example:
	//
	// 4548
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The score of the recommended content. Returned only if `AnswerSource` is `KNOWLEDGE`.
	//
	// example:
	//
	// 0.46
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// Clarification content. This can be an entity from knowledge graph QA, a knowledge title from FAQ-based QA, or a column value from table QA.
	//
	// example:
	//
	// 测试纯文本
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ChatResponseBodyMessagesRecommends) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyMessagesRecommends) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesRecommends) GetAnswerSource() *string {
	return s.AnswerSource
}

func (s *ChatResponseBodyMessagesRecommends) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *ChatResponseBodyMessagesRecommends) GetScore() *float64 {
	return s.Score
}

func (s *ChatResponseBodyMessagesRecommends) GetTitle() *string {
	return s.Title
}

func (s *ChatResponseBodyMessagesRecommends) SetAnswerSource(v string) *ChatResponseBodyMessagesRecommends {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetKnowledgeId(v string) *ChatResponseBodyMessagesRecommends {
	s.KnowledgeId = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetScore(v float64) *ChatResponseBodyMessagesRecommends {
	s.Score = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetTitle(v string) *ChatResponseBodyMessagesRecommends {
	s.Title = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyMessagesText struct {
	// The source of the answer.
	//
	// example:
	//
	// BotFramework
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// Title of the matched article. Returned only if `AnswerSource` is `MACHINE_READ`.
	//
	// example:
	//
	// 备案十一
	ArticleTitle *string `json:"ArticleTitle,omitempty" xml:"ArticleTitle,omitempty"`
	// Command parameters, such as the skill group for transferring to a human agent.
	//
	// example:
	//
	// {
	//
	// 	"sysToAgent": "{\\"skillGroup\\":\\"12\\"}"
	//
	// }
	Commands map[string]interface{} `json:"Commands,omitempty" xml:"Commands,omitempty"`
	// The content of the text message.
	//
	// example:
	//
	// 请问您要查哪里的天气？
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content format of the answer.
	//
	// example:
	//
	// PLAIN_TEXT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Name of the dialog. Returned only if `AnswerSource` is `BotFramework`.
	//
	// example:
	//
	// 示例_查天气
	DialogName *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	// Contains passthrough parameters.
	//
	// example:
	//
	// {
	//
	// 	"MATCHED_INTENT_SOURCE": "",
	//
	// 	"INTENT_ID": 724414,
	//
	// 	"IntentName": "查天气意图",
	//
	// 	"INTENT_DETAIL": "[我想|我要]查天气",
	//
	// 	"LGF_EXPRESSION": "[我想|我要]查天气",
	//
	// 	"IS_SESSION_FINISHED": false,
	//
	// 	"DsScore": 100.0,
	//
	// 	"DIALOG_ID": "299034",
	//
	// 	"FINISH_LABEL": false,
	//
	// 	"MODULE_START": false,
	//
	// 	"INTENT_NAME": "查天气意图",
	//
	// 	"INTENT_SOURCE": "Lgf",
	//
	// 	"DIALOG_NAME": "示例_查天气"
	//
	// }
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// Passthrough parameters. Returned only if `AnswerSource` is `BotFramework`.
	ExternalFlags map[string]interface{} `json:"ExternalFlags,omitempty" xml:"ExternalFlags,omitempty"`
	// The hit statement.
	//
	// example:
	//
	// 查天气
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// The name of the intent. This field is returned when `AnswerSource` is `BotFramework`.
	//
	// example:
	//
	// 查天气意图
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// Metadata.
	//
	// example:
	//
	// [[{\\"columnName\\":\\"姓名\\",\\"stringValue\\":\\"王珊珊\\"}]]
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// The node ID. Returned only if `AnswerSource` is `BotFramework`.
	//
	// example:
	//
	// 1410-c7a72a78.__city
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node. This field is returned when `AnswerSource` is `BotFramework`.
	//
	// example:
	//
	// 示例_查天气.查天气填槽.__city
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// A value of `SSML` indicates that an interactive slot-filling process has started in the dialog factory. This field is returned only if `AnswerSource` is `BotFramework`.
	//
	// example:
	//
	// SSML
	ResponseType *string `json:"ResponseType,omitempty" xml:"ResponseType,omitempty"`
	// The confidence score.
	//
	// example:
	//
	// 100.0
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// A list of slot objects. Returned only if `AnswerSource` is `BotFramework`.
	Slots []*ChatResponseBodyMessagesTextSlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
	// The title of a custom chit-chat topic.
	//
	// example:
	//
	// 问候
	UserDefinedChatTitle *string `json:"UserDefinedChatTitle,omitempty" xml:"UserDefinedChatTitle,omitempty"`
}

func (s ChatResponseBodyMessagesText) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyMessagesText) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesText) GetAnswerSource() *string {
	return s.AnswerSource
}

func (s *ChatResponseBodyMessagesText) GetArticleTitle() *string {
	return s.ArticleTitle
}

func (s *ChatResponseBodyMessagesText) GetCommands() map[string]interface{} {
	return s.Commands
}

func (s *ChatResponseBodyMessagesText) GetContent() *string {
	return s.Content
}

func (s *ChatResponseBodyMessagesText) GetContentType() *string {
	return s.ContentType
}

func (s *ChatResponseBodyMessagesText) GetDialogName() *string {
	return s.DialogName
}

func (s *ChatResponseBodyMessagesText) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *ChatResponseBodyMessagesText) GetExternalFlags() map[string]interface{} {
	return s.ExternalFlags
}

func (s *ChatResponseBodyMessagesText) GetHitStatement() *string {
	return s.HitStatement
}

func (s *ChatResponseBodyMessagesText) GetIntentName() *string {
	return s.IntentName
}

func (s *ChatResponseBodyMessagesText) GetMetaData() *string {
	return s.MetaData
}

func (s *ChatResponseBodyMessagesText) GetNodeId() *string {
	return s.NodeId
}

func (s *ChatResponseBodyMessagesText) GetNodeName() *string {
	return s.NodeName
}

func (s *ChatResponseBodyMessagesText) GetResponseType() *string {
	return s.ResponseType
}

func (s *ChatResponseBodyMessagesText) GetScore() *float64 {
	return s.Score
}

func (s *ChatResponseBodyMessagesText) GetSlots() []*ChatResponseBodyMessagesTextSlots {
	return s.Slots
}

func (s *ChatResponseBodyMessagesText) GetUserDefinedChatTitle() *string {
	return s.UserDefinedChatTitle
}

func (s *ChatResponseBodyMessagesText) SetAnswerSource(v string) *ChatResponseBodyMessagesText {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetArticleTitle(v string) *ChatResponseBodyMessagesText {
	s.ArticleTitle = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetCommands(v map[string]interface{}) *ChatResponseBodyMessagesText {
	s.Commands = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetContent(v string) *ChatResponseBodyMessagesText {
	s.Content = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetContentType(v string) *ChatResponseBodyMessagesText {
	s.ContentType = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetDialogName(v string) *ChatResponseBodyMessagesText {
	s.DialogName = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetExt(v map[string]interface{}) *ChatResponseBodyMessagesText {
	s.Ext = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetExternalFlags(v map[string]interface{}) *ChatResponseBodyMessagesText {
	s.ExternalFlags = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetHitStatement(v string) *ChatResponseBodyMessagesText {
	s.HitStatement = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetIntentName(v string) *ChatResponseBodyMessagesText {
	s.IntentName = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetMetaData(v string) *ChatResponseBodyMessagesText {
	s.MetaData = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetNodeId(v string) *ChatResponseBodyMessagesText {
	s.NodeId = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetNodeName(v string) *ChatResponseBodyMessagesText {
	s.NodeName = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetResponseType(v string) *ChatResponseBodyMessagesText {
	s.ResponseType = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetScore(v float64) *ChatResponseBodyMessagesText {
	s.Score = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetSlots(v []*ChatResponseBodyMessagesTextSlots) *ChatResponseBodyMessagesText {
	s.Slots = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetUserDefinedChatTitle(v string) *ChatResponseBodyMessagesText {
	s.UserDefinedChatTitle = &v
	return s
}

func (s *ChatResponseBodyMessagesText) Validate() error {
	if s.Slots != nil {
		for _, item := range s.Slots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatResponseBodyMessagesTextSlots struct {
	// Indicates whether the slot was filled.
	//
	// example:
	//
	// false
	Hit *bool `json:"Hit,omitempty" xml:"Hit,omitempty"`
	// The name of the slot.
	//
	// example:
	//
	// 查天气意图.city
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The original value from the user\\"s input.
	//
	// example:
	//
	// 北京
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// Extracted value of the slot.
	//
	// example:
	//
	// 北京
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ChatResponseBodyMessagesTextSlots) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyMessagesTextSlots) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesTextSlots) GetHit() *bool {
	return s.Hit
}

func (s *ChatResponseBodyMessagesTextSlots) GetName() *string {
	return s.Name
}

func (s *ChatResponseBodyMessagesTextSlots) GetOrigin() *string {
	return s.Origin
}

func (s *ChatResponseBodyMessagesTextSlots) GetValue() *string {
	return s.Value
}

func (s *ChatResponseBodyMessagesTextSlots) SetHit(v bool) *ChatResponseBodyMessagesTextSlots {
	s.Hit = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetName(v string) *ChatResponseBodyMessagesTextSlots {
	s.Name = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetOrigin(v string) *ChatResponseBodyMessagesTextSlots {
	s.Origin = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetValue(v string) *ChatResponseBodyMessagesTextSlots {
	s.Value = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) Validate() error {
	return dara.Validate(s)
}
