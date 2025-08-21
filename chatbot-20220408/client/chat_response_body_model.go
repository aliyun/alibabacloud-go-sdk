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
	// example:
	//
	// A2315C4B-A872-5DEE-9DAD-D73B194A4AEC
	MessageId    *string                     `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Messages     []*ChatResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	QuerySegList []*string                   `json:"QuerySegList,omitempty" xml:"QuerySegList,omitempty" type:"Repeated"`
	// example:
	//
	// A2315C4B-A872-5DEE-9DAD-D73B194A4AEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type ChatResponseBodyMessages struct {
	// example:
	//
	// KNOWLEDGE
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// example:
	//
	// Text
	AnswerType *string                               `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Knowledge  *ChatResponseBodyMessagesKnowledge    `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
	Recommends []*ChatResponseBodyMessagesRecommends `json:"Recommends,omitempty" xml:"Recommends,omitempty" type:"Repeated"`
	Text       *ChatResponseBodyMessagesText         `json:"Text,omitempty" xml:"Text,omitempty" type:"Struct"`
	Title      *string                               `json:"Title,omitempty" xml:"Title,omitempty"`
	VoiceTitle *string                               `json:"VoiceTitle,omitempty" xml:"VoiceTitle,omitempty"`
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
	return dara.Validate(s)
}

type ChatResponseBodyMessagesKnowledge struct {
	// example:
	//
	// KnowledgeBase
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// PLAIN_TEXT
	ContentType  *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// example:
	//
	// 735898
	Id                *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	RelatedKnowledges []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges `json:"RelatedKnowledges,omitempty" xml:"RelatedKnowledges,omitempty" type:"Repeated"`
	// example:
	//
	// 0.998
	Score   *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	Summary *string  `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title   *string  `json:"Title,omitempty" xml:"Title,omitempty"`
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
	return dara.Validate(s)
}

type ChatResponseBodyMessagesKnowledgeRelatedKnowledges struct {
	// example:
	//
	// 735899
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// example:
	//
	// KNOWLEDGE
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// example:
	//
	// 4548
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 0.46
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	Title *string  `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// example:
	//
	// BotFramework
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	ArticleTitle *string `json:"ArticleTitle,omitempty" xml:"ArticleTitle,omitempty"`
	// example:
	//
	// {
	//
	// 	"sysToAgent": "{\\"skillGroup\\":\\"12\\"}"
	//
	// }
	Commands map[string]interface{} `json:"Commands,omitempty" xml:"Commands,omitempty"`
	Content  *string                `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// PLAIN_TEXT
	ContentType   *string                `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DialogName    *string                `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	Ext           map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	ExternalFlags map[string]interface{} `json:"ExternalFlags,omitempty" xml:"ExternalFlags,omitempty"`
	HitStatement  *string                `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	IntentName    *string                `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	MetaData      *string                `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// example:
	//
	// 1410-c7a72a78.__city
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// SSML
	ResponseType *string `json:"ResponseType,omitempty" xml:"ResponseType,omitempty"`
	// example:
	//
	// 100.0
	Score                *float64                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Slots                []*ChatResponseBodyMessagesTextSlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
	UserDefinedChatTitle *string                              `json:"UserDefinedChatTitle,omitempty" xml:"UserDefinedChatTitle,omitempty"`
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
	return dara.Validate(s)
}

type ChatResponseBodyMessagesTextSlots struct {
	// example:
	//
	// false
	Hit    *bool   `json:"Hit,omitempty" xml:"Hit,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
