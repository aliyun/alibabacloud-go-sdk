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
	// The ID of the response message in the current session.
	//
	// example:
	//
	// A2315C4B-A872-5DEE-9DAD-D73B194A4AEC
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The list of messages.
	Messages []*ChatResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The tokenization results of the query, which may be empty.
	QuerySegList []*string `json:"QuerySegList,omitempty" xml:"QuerySegList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A2315C4B-A872-5DEE-9DAD-D73B194A4AEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the current session.
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
	// The source of the recommended answer when AnswerType is Recommend.
	//
	// [_single.resp.200.props.Messages.items.AnswerSource.enum. INTENT_FAQ_CLARIFY]mixed clarification from FAQ and dialog flow
	//
	// example:
	//
	// KNOWLEDGE
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The type of this message.
	//
	// example:
	//
	// Text
	AnswerType *string `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	// When AnswerType is Knowledge, this field contains the Knowledge object returned by the chatbot.
	Knowledge *ChatResponseBodyMessagesKnowledge `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
	// The list of recommendations returned by the bot when AnswerType is Recommend.
	Recommends []*ChatResponseBodyMessagesRecommends `json:"Recommends,omitempty" xml:"Recommends,omitempty" type:"Repeated"`
	// When AnswerType is Text, this field contains the Text object returned by the chatbot.
	Text *ChatResponseBodyMessagesText `json:"Text,omitempty" xml:"Text,omitempty" type:"Struct"`
	// The counter-question title in online scenarios.
	//
	// example:
	//
	// Which insurance age range are you inquiring about?
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The clarification content in voice scenarios.
	//
	// example:
	//
	// Are you asking about provident fund inquiry or provident fund withdrawal?
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
	// Distinguishes the answer type.
	//
	// KnowledgeBase: knowledge base.
	//
	// example:
	//
	// KnowledgeBase
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The knowledge category.
	//
	// example:
	//
	// 公积金
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the hit question.
	//
	// example:
	//
	// 公积金提取，请在首页搜索公积金提取，提交办事的表单
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The indicator of whether the answer is plain text or rich text.
	//
	// example:
	//
	// PLAIN_TEXT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The hit statement.
	//
	// example:
	//
	// 公积金
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// The ID of the hit question in the knowledge base.
	//
	// example:
	//
	// 735898
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of related knowledge entries.
	RelatedKnowledges []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges `json:"RelatedKnowledges,omitempty" xml:"RelatedKnowledges,omitempty" type:"Repeated"`
	// The score.
	//
	// example:
	//
	// 0.998
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The summary of the hit question.
	//
	// example:
	//
	// 公积金提取
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The title of the hit question.
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
	// The identifier of the clarification source.
	//
	// example:
	//
	// KNOWLEDGE
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The knowledge ID for clarification.
	//
	// example:
	//
	// 4548
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The score of the recommended content. This field has a value when AnswerSource is KNOWLEDGE.
	//
	// example:
	//
	// 0.46
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The clarification content, which can be an entity from knowledge graph Q&A, a knowledge title from knowledge-based Q&A, or a column value from table-based Q&A.
	//
	// example:
	//
	// Test plain text
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
	// Distinguishes the answer type.
	//
	// example:
	//
	// BotFramework
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// When AnswerSource is MACHINE_READ, this field returns the title of the hit article.
	//
	// example:
	//
	// ICP Filing 11
	//
	// [responses.200.schema.properties.Messages.items.properties.AnswerSource.enumValueTitles. INTENT_FAQ_CLARIFY]FAQ and dialog flow mixed clarification
	ArticleTitle *string `json:"ArticleTitle,omitempty" xml:"ArticleTitle,omitempty"`
	// The instruction parameters, such as the skill group for a transfer-to-agent instruction.
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
	// Which location do you want to check the weather for?
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The indicator of whether the answer is plain text or rich text.
	//
	// example:
	//
	// PLAIN_TEXT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// When AnswerSource is BotFramework, this field returns the dialog flow unit name.
	//
	// example:
	//
	// Sample_CheckWeather
	DialogName *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	// The pass-through parameters returned in this field.
	//
	// example:
	//
	// {
	//
	// 	"MATCHED_INTENT_SOURCE": "",
	//
	// 	"INTENT_ID": 724414,
	//
	// 	"IntentName": "CheckWeatherIntent",
	//
	// 	"INTENT_DETAIL": "[I want|I need to] check the weather",
	//
	// 	"LGF_EXPRESSION": "[I want|I need to] check the weather",
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
	// 	"INTENT_NAME": "CheckWeatherIntent",
	//
	// 	"INTENT_SOURCE": "Lgf",
	//
	// 	"DIALOG_NAME": "Sample_CheckWeather"
	//
	// }
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The pass-through parameters returned when AnswerSource is BotFramework.
	ExternalFlags map[string]interface{} `json:"ExternalFlags,omitempty" xml:"ExternalFlags,omitempty"`
	// The hit statement.
	//
	// example:
	//
	// CheckWeather
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// When AnswerSource is BotFramework, this field returns the intent name.
	//
	// example:
	//
	// CheckWeatherIntent
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// The metadata.
	//
	// example:
	//
	// [[{"columnName":"Name","stringValue":"Wang Shanshan"}]]
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// The node ID returned when AnswerSource is BotFramework.
	//
	// example:
	//
	// 1410-c7a72a78.__city
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// When AnswerSource is BotFramework, this field returns the node name.
	//
	// example:
	//
	// Sample_CheckWeather.CheckWeatherSlotFilling.__city
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// When AnswerSource is BotFramework and ResponseType is SSML, this field indicates that a dialog flow interactive slot filling is hit. Otherwise, this field does not exist.
	//
	// example:
	//
	// SSML
	ResponseType *string `json:"ResponseType,omitempty" xml:"ResponseType,omitempty"`
	// The score.
	//
	// example:
	//
	// 100.0
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// When AnswerSource is BotFramework, this field returns the list of named entities.
	Slots []*ChatResponseBodyMessagesTextSlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
	// The title of the custom chitchat topic.
	//
	// example:
	//
	// Greeting
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
	// Indicates whether the slot is hit.
	//
	// example:
	//
	// false
	Hit *bool `json:"Hit,omitempty" xml:"Hit,omitempty"`
	// The name.
	//
	// example:
	//
	// CheckWeatherIntent.city
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The original value.
	//
	// example:
	//
	// Beijing
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The specific value.
	//
	// example:
	//
	// Beijing
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
