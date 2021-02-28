// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivatePerspectiveRequest struct {
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s ActivatePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *ActivatePerspectiveRequest) SetPerspectiveId(v string) *ActivatePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type ActivatePerspectiveResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivatePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivatePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *ActivatePerspectiveResponseBody) SetRequestId(v string) *ActivatePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

type ActivatePerspectiveResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ActivatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivatePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivatePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *ActivatePerspectiveResponse) SetHeaders(v map[string]*string) *ActivatePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *ActivatePerspectiveResponse) SetBody(v *ActivatePerspectiveResponseBody) *ActivatePerspectiveResponse {
	s.Body = v
	return s
}

type AddSynonymRequest struct {
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	Synonym      *string `json:"Synonym,omitempty" xml:"Synonym,omitempty"`
}

func (s AddSynonymRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSynonymRequest) GoString() string {
	return s.String()
}

func (s *AddSynonymRequest) SetCoreWordName(v string) *AddSynonymRequest {
	s.CoreWordName = &v
	return s
}

func (s *AddSynonymRequest) SetSynonym(v string) *AddSynonymRequest {
	s.Synonym = &v
	return s
}

type AddSynonymResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSynonymResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSynonymResponseBody) GoString() string {
	return s.String()
}

func (s *AddSynonymResponseBody) SetRequestId(v string) *AddSynonymResponseBody {
	s.RequestId = &v
	return s
}

type AddSynonymResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSynonymResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSynonymResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSynonymResponse) GoString() string {
	return s.String()
}

func (s *AddSynonymResponse) SetHeaders(v map[string]*string) *AddSynonymResponse {
	s.Headers = v
	return s
}

func (s *AddSynonymResponse) SetBody(v *AddSynonymResponseBody) *AddSynonymResponse {
	s.Body = v
	return s
}

type AppendEntityMemberRequest struct {
	EntityId  *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ApplyType *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	Member    *string `json:"Member,omitempty" xml:"Member,omitempty"`
}

func (s AppendEntityMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendEntityMemberRequest) GoString() string {
	return s.String()
}

func (s *AppendEntityMemberRequest) SetEntityId(v int64) *AppendEntityMemberRequest {
	s.EntityId = &v
	return s
}

func (s *AppendEntityMemberRequest) SetApplyType(v string) *AppendEntityMemberRequest {
	s.ApplyType = &v
	return s
}

func (s *AppendEntityMemberRequest) SetMember(v string) *AppendEntityMemberRequest {
	s.Member = &v
	return s
}

type AppendEntityMemberResponseBody struct {
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppendEntityMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppendEntityMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AppendEntityMemberResponseBody) SetEntityId(v string) *AppendEntityMemberResponseBody {
	s.EntityId = &v
	return s
}

func (s *AppendEntityMemberResponseBody) SetRequestId(v string) *AppendEntityMemberResponseBody {
	s.RequestId = &v
	return s
}

type AppendEntityMemberResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppendEntityMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppendEntityMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendEntityMemberResponse) GoString() string {
	return s.String()
}

func (s *AppendEntityMemberResponse) SetHeaders(v map[string]*string) *AppendEntityMemberResponse {
	s.Headers = v
	return s
}

func (s *AppendEntityMemberResponse) SetBody(v *AppendEntityMemberResponseBody) *AppendEntityMemberResponse {
	s.Body = v
	return s
}

type AssociateRequest struct {
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Utterance     *string   `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	SessionId     *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	BusinessScope *string   `json:"BusinessScope,omitempty" xml:"BusinessScope,omitempty"`
	RecommendNum  *int32    `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	Perspective   []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
}

func (s AssociateRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateRequest) GoString() string {
	return s.String()
}

func (s *AssociateRequest) SetInstanceId(v string) *AssociateRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateRequest) SetUtterance(v string) *AssociateRequest {
	s.Utterance = &v
	return s
}

func (s *AssociateRequest) SetSessionId(v string) *AssociateRequest {
	s.SessionId = &v
	return s
}

func (s *AssociateRequest) SetBusinessScope(v string) *AssociateRequest {
	s.BusinessScope = &v
	return s
}

func (s *AssociateRequest) SetRecommendNum(v int32) *AssociateRequest {
	s.RecommendNum = &v
	return s
}

func (s *AssociateRequest) SetPerspective(v []*string) *AssociateRequest {
	s.Perspective = v
	return s
}

type AssociateResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Associate []*AssociateResponseBodyAssociate `json:"Associate,omitempty" xml:"Associate,omitempty" type:"Repeated"`
	SessionId *string                           `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	MessageId *string                           `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s AssociateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResponseBody) SetRequestId(v string) *AssociateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateResponseBody) SetAssociate(v []*AssociateResponseBodyAssociate) *AssociateResponseBody {
	s.Associate = v
	return s
}

func (s *AssociateResponseBody) SetSessionId(v string) *AssociateResponseBody {
	s.SessionId = &v
	return s
}

func (s *AssociateResponseBody) SetMessageId(v string) *AssociateResponseBody {
	s.MessageId = &v
	return s
}

type AssociateResponseBodyAssociate struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AssociateResponseBodyAssociate) String() string {
	return tea.Prettify(s)
}

func (s AssociateResponseBodyAssociate) GoString() string {
	return s.String()
}

func (s *AssociateResponseBodyAssociate) SetTitle(v string) *AssociateResponseBodyAssociate {
	s.Title = &v
	return s
}

type AssociateResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateResponse) GoString() string {
	return s.String()
}

func (s *AssociateResponse) SetHeaders(v map[string]*string) *AssociateResponse {
	s.Headers = v
	return s
}

func (s *AssociateResponse) SetBody(v *AssociateResponseBody) *AssociateResponse {
	s.Body = v
	return s
}

type ChatRequest struct {
	InstanceId         *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SessionId          *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	KnowledgeId        *string   `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	SenderId           *string   `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	SenderNick         *string   `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	Tag                *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Utterance          *string   `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	Recommend          *bool     `json:"Recommend,omitempty" xml:"Recommend,omitempty"`
	RecommendNum       *int32    `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	IntentName         *string   `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	DefaultPerspective *string   `json:"DefaultPerspective,omitempty" xml:"DefaultPerspective,omitempty"`
	BusinessScope      *string   `json:"BusinessScope,omitempty" xml:"BusinessScope,omitempty"`
	VendorParam        *string   `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
	Emotion            *bool     `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	SandBox            *bool     `json:"SandBox,omitempty" xml:"SandBox,omitempty"`
	Perspective        []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
}

func (s ChatRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) SetInstanceId(v string) *ChatRequest {
	s.InstanceId = &v
	return s
}

func (s *ChatRequest) SetSessionId(v string) *ChatRequest {
	s.SessionId = &v
	return s
}

func (s *ChatRequest) SetKnowledgeId(v string) *ChatRequest {
	s.KnowledgeId = &v
	return s
}

func (s *ChatRequest) SetSenderId(v string) *ChatRequest {
	s.SenderId = &v
	return s
}

func (s *ChatRequest) SetSenderNick(v string) *ChatRequest {
	s.SenderNick = &v
	return s
}

func (s *ChatRequest) SetTag(v string) *ChatRequest {
	s.Tag = &v
	return s
}

func (s *ChatRequest) SetUtterance(v string) *ChatRequest {
	s.Utterance = &v
	return s
}

func (s *ChatRequest) SetRecommend(v bool) *ChatRequest {
	s.Recommend = &v
	return s
}

func (s *ChatRequest) SetRecommendNum(v int32) *ChatRequest {
	s.RecommendNum = &v
	return s
}

func (s *ChatRequest) SetIntentName(v string) *ChatRequest {
	s.IntentName = &v
	return s
}

func (s *ChatRequest) SetDefaultPerspective(v string) *ChatRequest {
	s.DefaultPerspective = &v
	return s
}

func (s *ChatRequest) SetBusinessScope(v string) *ChatRequest {
	s.BusinessScope = &v
	return s
}

func (s *ChatRequest) SetVendorParam(v string) *ChatRequest {
	s.VendorParam = &v
	return s
}

func (s *ChatRequest) SetEmotion(v bool) *ChatRequest {
	s.Emotion = &v
	return s
}

func (s *ChatRequest) SetSandBox(v bool) *ChatRequest {
	s.SandBox = &v
	return s
}

func (s *ChatRequest) SetPerspective(v []*string) *ChatRequest {
	s.Perspective = v
	return s
}

type ChatResponseBody struct {
	Messages  []*ChatResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tag       *string                     `json:"Tag,omitempty" xml:"Tag,omitempty"`
	SessionId *string                     `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	MessageId *string                     `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s ChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) SetMessages(v []*ChatResponseBodyMessages) *ChatResponseBody {
	s.Messages = v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBody) SetTag(v string) *ChatResponseBody {
	s.Tag = &v
	return s
}

func (s *ChatResponseBody) SetSessionId(v string) *ChatResponseBody {
	s.SessionId = &v
	return s
}

func (s *ChatResponseBody) SetMessageId(v string) *ChatResponseBody {
	s.MessageId = &v
	return s
}

type ChatResponseBodyMessages struct {
	Type       *string                               `json:"Type,omitempty" xml:"Type,omitempty"`
	Knowledge  *ChatResponseBodyMessagesKnowledge    `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
	Text       *ChatResponseBodyMessagesText         `json:"Text,omitempty" xml:"Text,omitempty" type:"Struct"`
	Recommends []*ChatResponseBodyMessagesRecommends `json:"Recommends,omitempty" xml:"Recommends,omitempty" type:"Repeated"`
}

func (s ChatResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessages) SetType(v string) *ChatResponseBodyMessages {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyMessages) SetKnowledge(v *ChatResponseBodyMessagesKnowledge) *ChatResponseBodyMessages {
	s.Knowledge = v
	return s
}

func (s *ChatResponseBodyMessages) SetText(v *ChatResponseBodyMessagesText) *ChatResponseBodyMessages {
	s.Text = v
	return s
}

func (s *ChatResponseBodyMessages) SetRecommends(v []*ChatResponseBodyMessagesRecommends) *ChatResponseBodyMessages {
	s.Recommends = v
	return s
}

type ChatResponseBodyMessagesKnowledge struct {
	HitStatement      *string                                               `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	Summary           *string                                               `json:"Summary,omitempty" xml:"Summary,omitempty"`
	RelatedKnowledges []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges `json:"RelatedKnowledges,omitempty" xml:"RelatedKnowledges,omitempty" type:"Repeated"`
	Category          *string                                               `json:"Category,omitempty" xml:"Category,omitempty"`
	Title             *string                                               `json:"Title,omitempty" xml:"Title,omitempty"`
	Content           *string                                               `json:"Content,omitempty" xml:"Content,omitempty"`
	AnswerSource      *string                                               `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	Id                *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ChatResponseBodyMessagesKnowledge) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesKnowledge) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesKnowledge) SetHitStatement(v string) *ChatResponseBodyMessagesKnowledge {
	s.HitStatement = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetSummary(v string) *ChatResponseBodyMessagesKnowledge {
	s.Summary = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetRelatedKnowledges(v []*ChatResponseBodyMessagesKnowledgeRelatedKnowledges) *ChatResponseBodyMessagesKnowledge {
	s.RelatedKnowledges = v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetCategory(v string) *ChatResponseBodyMessagesKnowledge {
	s.Category = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetTitle(v string) *ChatResponseBodyMessagesKnowledge {
	s.Title = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetContent(v string) *ChatResponseBodyMessagesKnowledge {
	s.Content = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetAnswerSource(v string) *ChatResponseBodyMessagesKnowledge {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledge) SetId(v string) *ChatResponseBodyMessagesKnowledge {
	s.Id = &v
	return s
}

type ChatResponseBodyMessagesKnowledgeRelatedKnowledges struct {
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ChatResponseBodyMessagesKnowledgeRelatedKnowledges) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesKnowledgeRelatedKnowledges) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) SetKnowledgeId(v string) *ChatResponseBodyMessagesKnowledgeRelatedKnowledges {
	s.KnowledgeId = &v
	return s
}

func (s *ChatResponseBodyMessagesKnowledgeRelatedKnowledges) SetTitle(v string) *ChatResponseBodyMessagesKnowledgeRelatedKnowledges {
	s.Title = &v
	return s
}

type ChatResponseBodyMessagesText struct {
	HitStatement         *string                              `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	DialogName           *string                              `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	ArticleTitle         *string                              `json:"ArticleTitle,omitempty" xml:"ArticleTitle,omitempty"`
	AnswerSource         *string                              `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	Slots                []*ChatResponseBodyMessagesTextSlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
	IntentName           *string                              `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	MetaData             *string                              `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	NodeName             *string                              `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	ExternalFlags        map[string]interface{}               `json:"ExternalFlags,omitempty" xml:"ExternalFlags,omitempty"`
	Ext                  map[string]interface{}               `json:"Ext,omitempty" xml:"Ext,omitempty"`
	UserDefinedChatTitle *string                              `json:"UserDefinedChatTitle,omitempty" xml:"UserDefinedChatTitle,omitempty"`
	Content              *string                              `json:"Content,omitempty" xml:"Content,omitempty"`
	NodeId               *string                              `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ChatResponseBodyMessagesText) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesText) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesText) SetHitStatement(v string) *ChatResponseBodyMessagesText {
	s.HitStatement = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetDialogName(v string) *ChatResponseBodyMessagesText {
	s.DialogName = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetArticleTitle(v string) *ChatResponseBodyMessagesText {
	s.ArticleTitle = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetAnswerSource(v string) *ChatResponseBodyMessagesText {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetSlots(v []*ChatResponseBodyMessagesTextSlots) *ChatResponseBodyMessagesText {
	s.Slots = v
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

func (s *ChatResponseBodyMessagesText) SetNodeName(v string) *ChatResponseBodyMessagesText {
	s.NodeName = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetExternalFlags(v map[string]interface{}) *ChatResponseBodyMessagesText {
	s.ExternalFlags = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetExt(v map[string]interface{}) *ChatResponseBodyMessagesText {
	s.Ext = v
	return s
}

func (s *ChatResponseBodyMessagesText) SetUserDefinedChatTitle(v string) *ChatResponseBodyMessagesText {
	s.UserDefinedChatTitle = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetContent(v string) *ChatResponseBodyMessagesText {
	s.Content = &v
	return s
}

func (s *ChatResponseBodyMessagesText) SetNodeId(v string) *ChatResponseBodyMessagesText {
	s.NodeId = &v
	return s
}

type ChatResponseBodyMessagesTextSlots struct {
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	IsHit  *bool   `json:"IsHit,omitempty" xml:"IsHit,omitempty"`
}

func (s ChatResponseBodyMessagesTextSlots) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesTextSlots) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesTextSlots) SetValue(v string) *ChatResponseBodyMessagesTextSlots {
	s.Value = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetOrigin(v string) *ChatResponseBodyMessagesTextSlots {
	s.Origin = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetName(v string) *ChatResponseBodyMessagesTextSlots {
	s.Name = &v
	return s
}

func (s *ChatResponseBodyMessagesTextSlots) SetIsHit(v bool) *ChatResponseBodyMessagesTextSlots {
	s.IsHit = &v
	return s
}

type ChatResponseBodyMessagesRecommends struct {
	Summary      *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	KnowledgeId  *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ChatResponseBodyMessagesRecommends) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesRecommends) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyMessagesRecommends) SetSummary(v string) *ChatResponseBodyMessagesRecommends {
	s.Summary = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetKnowledgeId(v string) *ChatResponseBodyMessagesRecommends {
	s.KnowledgeId = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetCategory(v string) *ChatResponseBodyMessagesRecommends {
	s.Category = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetTitle(v string) *ChatResponseBodyMessagesRecommends {
	s.Title = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetAnswerSource(v string) *ChatResponseBodyMessagesRecommends {
	s.AnswerSource = &v
	return s
}

func (s *ChatResponseBodyMessagesRecommends) SetContent(v string) *ChatResponseBodyMessagesRecommends {
	s.Content = &v
	return s
}

type ChatResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChatResponseBody  `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChatResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatResponse) GoString() string {
	return s.String()
}

func (s *ChatResponse) SetHeaders(v map[string]*string) *ChatResponse {
	s.Headers = v
	return s
}

func (s *ChatResponse) SetBody(v *ChatResponseBody) *ChatResponse {
	s.Body = v
	return s
}

type CreateBotRequest struct {
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	TimeZone     *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Avatar       *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	RobotType    *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s CreateBotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBotRequest) GoString() string {
	return s.String()
}

func (s *CreateBotRequest) SetLanguageCode(v string) *CreateBotRequest {
	s.LanguageCode = &v
	return s
}

func (s *CreateBotRequest) SetTimeZone(v string) *CreateBotRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateBotRequest) SetName(v string) *CreateBotRequest {
	s.Name = &v
	return s
}

func (s *CreateBotRequest) SetAvatar(v string) *CreateBotRequest {
	s.Avatar = &v
	return s
}

func (s *CreateBotRequest) SetIntroduction(v string) *CreateBotRequest {
	s.Introduction = &v
	return s
}

func (s *CreateBotRequest) SetRobotType(v string) *CreateBotRequest {
	s.RobotType = &v
	return s
}

type CreateBotResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateBotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBotResponseBody) SetRequestId(v string) *CreateBotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBotResponseBody) SetInstanceId(v string) *CreateBotResponseBody {
	s.InstanceId = &v
	return s
}

type CreateBotResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBotResponse) GoString() string {
	return s.String()
}

func (s *CreateBotResponse) SetHeaders(v map[string]*string) *CreateBotResponse {
	s.Headers = v
	return s
}

func (s *CreateBotResponse) SetBody(v *CreateBotResponseBody) *CreateBotResponse {
	s.Body = v
	return s
}

type CreateCategoryRequest struct {
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateCategoryRequest) SetParentCategoryId(v int64) *CreateCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *CreateCategoryRequest) SetName(v string) *CreateCategoryRequest {
	s.Name = &v
	return s
}

type CreateCategoryResponseBody struct {
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponseBody) SetCategoryId(v int64) *CreateCategoryResponseBody {
	s.CategoryId = &v
	return s
}

func (s *CreateCategoryResponseBody) SetRequestId(v string) *CreateCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCategoryResponseBody) SetSuccess(v bool) *CreateCategoryResponseBody {
	s.Success = &v
	return s
}

type CreateCategoryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponse) SetHeaders(v map[string]*string) *CreateCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateCategoryResponse) SetBody(v *CreateCategoryResponseBody) *CreateCategoryResponse {
	s.Body = v
	return s
}

type CreateCoreWordRequest struct {
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
}

func (s CreateCoreWordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCoreWordRequest) GoString() string {
	return s.String()
}

func (s *CreateCoreWordRequest) SetCoreWordName(v string) *CreateCoreWordRequest {
	s.CoreWordName = &v
	return s
}

type CreateCoreWordResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CoreWordCode *string `json:"CoreWordCode,omitempty" xml:"CoreWordCode,omitempty"`
}

func (s CreateCoreWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCoreWordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCoreWordResponseBody) SetRequestId(v string) *CreateCoreWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCoreWordResponseBody) SetCoreWordCode(v string) *CreateCoreWordResponseBody {
	s.CoreWordCode = &v
	return s
}

type CreateCoreWordResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCoreWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCoreWordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCoreWordResponse) GoString() string {
	return s.String()
}

func (s *CreateCoreWordResponse) SetHeaders(v map[string]*string) *CreateCoreWordResponse {
	s.Headers = v
	return s
}

func (s *CreateCoreWordResponse) SetBody(v *CreateCoreWordResponseBody) *CreateCoreWordResponse {
	s.Body = v
	return s
}

type CreateDialogRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DialogName  *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogRequest) GoString() string {
	return s.String()
}

func (s *CreateDialogRequest) SetInstanceId(v string) *CreateDialogRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDialogRequest) SetDialogName(v string) *CreateDialogRequest {
	s.DialogName = &v
	return s
}

func (s *CreateDialogRequest) SetDescription(v string) *CreateDialogRequest {
	s.Description = &v
	return s
}

type CreateDialogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DialogId  *string `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s CreateDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDialogResponseBody) SetRequestId(v string) *CreateDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDialogResponseBody) SetDialogId(v string) *CreateDialogResponseBody {
	s.DialogId = &v
	return s
}

type CreateDialogResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDialogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogResponse) GoString() string {
	return s.String()
}

func (s *CreateDialogResponse) SetHeaders(v map[string]*string) *CreateDialogResponse {
	s.Headers = v
	return s
}

func (s *CreateDialogResponse) SetBody(v *CreateDialogResponseBody) *CreateDialogResponse {
	s.Body = v
	return s
}

type CreateEntityRequest struct {
	DialogId   *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Regex      *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
	Members    *string `json:"Members,omitempty" xml:"Members,omitempty"`
}

func (s CreateEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityRequest) SetDialogId(v int64) *CreateEntityRequest {
	s.DialogId = &v
	return s
}

func (s *CreateEntityRequest) SetEntityName(v string) *CreateEntityRequest {
	s.EntityName = &v
	return s
}

func (s *CreateEntityRequest) SetEntityType(v string) *CreateEntityRequest {
	s.EntityType = &v
	return s
}

func (s *CreateEntityRequest) SetRegex(v string) *CreateEntityRequest {
	s.Regex = &v
	return s
}

func (s *CreateEntityRequest) SetMembers(v string) *CreateEntityRequest {
	s.Members = &v
	return s
}

type CreateEntityResponseBody struct {
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEntityResponseBody) SetEntityId(v string) *CreateEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *CreateEntityResponseBody) SetRequestId(v string) *CreateEntityResponseBody {
	s.RequestId = &v
	return s
}

type CreateEntityResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateEntityResponse) SetHeaders(v map[string]*string) *CreateEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateEntityResponse) SetBody(v *CreateEntityResponseBody) *CreateEntityResponse {
	s.Body = v
	return s
}

type CreateIntentRequest struct {
	IntentDefinition *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
	DialogId         *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s CreateIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentRequest) SetIntentDefinition(v string) *CreateIntentRequest {
	s.IntentDefinition = &v
	return s
}

func (s *CreateIntentRequest) SetDialogId(v int64) *CreateIntentRequest {
	s.DialogId = &v
	return s
}

type CreateIntentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IntentId  *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s CreateIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntentResponseBody) SetRequestId(v string) *CreateIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntentResponseBody) SetIntentId(v string) *CreateIntentResponseBody {
	s.IntentId = &v
	return s
}

type CreateIntentResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIntentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentResponse) GoString() string {
	return s.String()
}

func (s *CreateIntentResponse) SetHeaders(v map[string]*string) *CreateIntentResponse {
	s.Headers = v
	return s
}

func (s *CreateIntentResponse) SetBody(v *CreateIntentResponseBody) *CreateIntentResponse {
	s.Body = v
	return s
}

type CreateKnowledgeRequest struct {
	Knowledge *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
}

func (s CreateKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeRequest) SetKnowledge(v string) *CreateKnowledgeRequest {
	s.Knowledge = &v
	return s
}

type CreateKnowledgeResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s CreateKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeResponseBody) SetRequestId(v string) *CreateKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeResponseBody) SetKnowledgeId(v int64) *CreateKnowledgeResponseBody {
	s.KnowledgeId = &v
	return s
}

type CreateKnowledgeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeResponse) SetHeaders(v map[string]*string) *CreateKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeResponse) SetBody(v *CreateKnowledgeResponseBody) *CreateKnowledgeResponse {
	s.Body = v
	return s
}

type DeleteBotRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteBotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBotRequest) GoString() string {
	return s.String()
}

func (s *DeleteBotRequest) SetInstanceId(v string) *DeleteBotRequest {
	s.InstanceId = &v
	return s
}

type DeleteBotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBotResponseBody) SetRequestId(v string) *DeleteBotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBotResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBotResponse) GoString() string {
	return s.String()
}

func (s *DeleteBotResponse) SetHeaders(v map[string]*string) *DeleteBotResponse {
	s.Headers = v
	return s
}

func (s *DeleteBotResponse) SetBody(v *DeleteBotResponseBody) *DeleteBotResponse {
	s.Body = v
	return s
}

type DeleteCategoryRequest struct {
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DeleteCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteCategoryRequest) SetCategoryId(v int64) *DeleteCategoryRequest {
	s.CategoryId = &v
	return s
}

type DeleteCategoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponseBody) SetRequestId(v string) *DeleteCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetSuccess(v bool) *DeleteCategoryResponseBody {
	s.Success = &v
	return s
}

type DeleteCategoryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponse) SetHeaders(v map[string]*string) *DeleteCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteCategoryResponse) SetBody(v *DeleteCategoryResponseBody) *DeleteCategoryResponse {
	s.Body = v
	return s
}

type DeleteCoreWordRequest struct {
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
}

func (s DeleteCoreWordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCoreWordRequest) GoString() string {
	return s.String()
}

func (s *DeleteCoreWordRequest) SetCoreWordName(v string) *DeleteCoreWordRequest {
	s.CoreWordName = &v
	return s
}

type DeleteCoreWordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCoreWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCoreWordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCoreWordResponseBody) SetRequestId(v string) *DeleteCoreWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCoreWordResponseBody) SetSuccess(v bool) *DeleteCoreWordResponseBody {
	s.Success = &v
	return s
}

type DeleteCoreWordResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCoreWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCoreWordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCoreWordResponse) GoString() string {
	return s.String()
}

func (s *DeleteCoreWordResponse) SetHeaders(v map[string]*string) *DeleteCoreWordResponse {
	s.Headers = v
	return s
}

func (s *DeleteCoreWordResponse) SetBody(v *DeleteCoreWordResponseBody) *DeleteCoreWordResponse {
	s.Body = v
	return s
}

type DeleteDialogRequest struct {
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s DeleteDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDialogRequest) GoString() string {
	return s.String()
}

func (s *DeleteDialogRequest) SetDialogId(v int64) *DeleteDialogRequest {
	s.DialogId = &v
	return s
}

type DeleteDialogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDialogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDialogResponseBody) SetRequestId(v string) *DeleteDialogResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDialogResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDialogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDialogResponse) GoString() string {
	return s.String()
}

func (s *DeleteDialogResponse) SetHeaders(v map[string]*string) *DeleteDialogResponse {
	s.Headers = v
	return s
}

func (s *DeleteDialogResponse) SetBody(v *DeleteDialogResponseBody) *DeleteDialogResponse {
	s.Body = v
	return s
}

type DeleteEntityRequest struct {
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s DeleteEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityRequest) SetEntityId(v int64) *DeleteEntityRequest {
	s.EntityId = &v
	return s
}

type DeleteEntityResponseBody struct {
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityResponseBody) SetEntityId(v string) *DeleteEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *DeleteEntityResponseBody) SetRequestId(v string) *DeleteEntityResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEntityResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityResponse) SetHeaders(v map[string]*string) *DeleteEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityResponse) SetBody(v *DeleteEntityResponseBody) *DeleteEntityResponse {
	s.Body = v
	return s
}

type DeleteIntentRequest struct {
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s DeleteIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntentRequest) SetIntentId(v int64) *DeleteIntentRequest {
	s.IntentId = &v
	return s
}

type DeleteIntentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IntentId  *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s DeleteIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntentResponseBody) SetRequestId(v string) *DeleteIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIntentResponseBody) SetIntentId(v string) *DeleteIntentResponseBody {
	s.IntentId = &v
	return s
}

type DeleteIntentResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIntentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntentResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntentResponse) SetHeaders(v map[string]*string) *DeleteIntentResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntentResponse) SetBody(v *DeleteIntentResponseBody) *DeleteIntentResponse {
	s.Body = v
	return s
}

type DeleteKnowledgeRequest struct {
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DeleteKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeRequest) SetKnowledgeId(v int64) *DeleteKnowledgeRequest {
	s.KnowledgeId = &v
	return s
}

type DeleteKnowledgeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeResponseBody) SetRequestId(v string) *DeleteKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKnowledgeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeResponse) SetBody(v *DeleteKnowledgeResponseBody) *DeleteKnowledgeResponse {
	s.Body = v
	return s
}

type DescribeBotRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeBotRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBotRequest) GoString() string {
	return s.String()
}

func (s *DescribeBotRequest) SetInstanceId(v string) *DescribeBotRequest {
	s.InstanceId = &v
	return s
}

type DescribeBotResponseBody struct {
	LanguageCode *string                              `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	TimeZone     *string                              `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Introduction *string                              `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	InstanceId   *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Categories   []*DescribeBotResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	CreateTime   *string                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Avatar       *string                              `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Logo         *string                              `json:"Logo,omitempty" xml:"Logo,omitempty"`
	Name         *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBotResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBotResponseBody) SetLanguageCode(v string) *DescribeBotResponseBody {
	s.LanguageCode = &v
	return s
}

func (s *DescribeBotResponseBody) SetTimeZone(v string) *DescribeBotResponseBody {
	s.TimeZone = &v
	return s
}

func (s *DescribeBotResponseBody) SetRequestId(v string) *DescribeBotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBotResponseBody) SetIntroduction(v string) *DescribeBotResponseBody {
	s.Introduction = &v
	return s
}

func (s *DescribeBotResponseBody) SetInstanceId(v string) *DescribeBotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeBotResponseBody) SetCategories(v []*DescribeBotResponseBodyCategories) *DescribeBotResponseBody {
	s.Categories = v
	return s
}

func (s *DescribeBotResponseBody) SetCreateTime(v string) *DescribeBotResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeBotResponseBody) SetAvatar(v string) *DescribeBotResponseBody {
	s.Avatar = &v
	return s
}

func (s *DescribeBotResponseBody) SetLogo(v string) *DescribeBotResponseBody {
	s.Logo = &v
	return s
}

func (s *DescribeBotResponseBody) SetName(v string) *DescribeBotResponseBody {
	s.Name = &v
	return s
}

type DescribeBotResponseBodyCategories struct {
	CategoryId       *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s DescribeBotResponseBodyCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeBotResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *DescribeBotResponseBodyCategories) SetCategoryId(v int64) *DescribeBotResponseBodyCategories {
	s.CategoryId = &v
	return s
}

func (s *DescribeBotResponseBodyCategories) SetName(v string) *DescribeBotResponseBodyCategories {
	s.Name = &v
	return s
}

func (s *DescribeBotResponseBodyCategories) SetParentCategoryId(v int64) *DescribeBotResponseBodyCategories {
	s.ParentCategoryId = &v
	return s
}

type DescribeBotResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBotResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBotResponse) GoString() string {
	return s.String()
}

func (s *DescribeBotResponse) SetHeaders(v map[string]*string) *DescribeBotResponse {
	s.Headers = v
	return s
}

func (s *DescribeBotResponse) SetBody(v *DescribeBotResponseBody) *DescribeBotResponse {
	s.Body = v
	return s
}

type DescribeCategoryRequest struct {
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DescribeCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCategoryRequest) SetCategoryId(v int64) *DescribeCategoryRequest {
	s.CategoryId = &v
	return s
}

type DescribeCategoryResponseBody struct {
	CategoryId       *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponseBody) SetCategoryId(v int64) *DescribeCategoryResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeCategoryResponseBody) SetRequestId(v string) *DescribeCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCategoryResponseBody) SetParentCategoryId(v int64) *DescribeCategoryResponseBody {
	s.ParentCategoryId = &v
	return s
}

func (s *DescribeCategoryResponseBody) SetName(v string) *DescribeCategoryResponseBody {
	s.Name = &v
	return s
}

type DescribeCategoryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponse) SetHeaders(v map[string]*string) *DescribeCategoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCategoryResponse) SetBody(v *DescribeCategoryResponseBody) *DescribeCategoryResponse {
	s.Body = v
	return s
}

type DescribeCoreWordRequest struct {
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
}

func (s DescribeCoreWordRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCoreWordRequest) GoString() string {
	return s.String()
}

func (s *DescribeCoreWordRequest) SetCoreWordName(v string) *DescribeCoreWordRequest {
	s.CoreWordName = &v
	return s
}

type DescribeCoreWordResponseBody struct {
	CoreWordName *string   `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	Synonyms     []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
	ModifyTime   *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime   *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CoreWordCode *string   `json:"CoreWordCode,omitempty" xml:"CoreWordCode,omitempty"`
}

func (s DescribeCoreWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCoreWordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCoreWordResponseBody) SetCoreWordName(v string) *DescribeCoreWordResponseBody {
	s.CoreWordName = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetSynonyms(v []*string) *DescribeCoreWordResponseBody {
	s.Synonyms = v
	return s
}

func (s *DescribeCoreWordResponseBody) SetModifyTime(v string) *DescribeCoreWordResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetRequestId(v string) *DescribeCoreWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetCreateTime(v string) *DescribeCoreWordResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeCoreWordResponseBody) SetCoreWordCode(v string) *DescribeCoreWordResponseBody {
	s.CoreWordCode = &v
	return s
}

type DescribeCoreWordResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCoreWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCoreWordResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCoreWordResponse) GoString() string {
	return s.String()
}

func (s *DescribeCoreWordResponse) SetHeaders(v map[string]*string) *DescribeCoreWordResponse {
	s.Headers = v
	return s
}

func (s *DescribeCoreWordResponse) SetBody(v *DescribeCoreWordResponseBody) *DescribeCoreWordResponse {
	s.Body = v
	return s
}

type DescribeDialogRequest struct {
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s DescribeDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDialogRequest) SetDialogId(v int64) *DescribeDialogRequest {
	s.DialogId = &v
	return s
}

type DescribeDialogResponseBody struct {
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId   *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	DialogId       *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	IsOnline       *bool   `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	DialogName     *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	ModifyUserId   *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	IsSampleDialog *bool   `json:"IsSampleDialog,omitempty" xml:"IsSampleDialog,omitempty"`
}

func (s DescribeDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDialogResponseBody) SetStatus(v int32) *DescribeDialogResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDialogResponseBody) SetModifyTime(v string) *DescribeDialogResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDialogResponseBody) SetDescription(v string) *DescribeDialogResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDialogResponseBody) SetRequestId(v string) *DescribeDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDialogResponseBody) SetCreateTime(v string) *DescribeDialogResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDialogResponseBody) SetCreateUserId(v string) *DescribeDialogResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeDialogResponseBody) SetDialogId(v int64) *DescribeDialogResponseBody {
	s.DialogId = &v
	return s
}

func (s *DescribeDialogResponseBody) SetCreateUserName(v string) *DescribeDialogResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeDialogResponseBody) SetIsOnline(v bool) *DescribeDialogResponseBody {
	s.IsOnline = &v
	return s
}

func (s *DescribeDialogResponseBody) SetDialogName(v string) *DescribeDialogResponseBody {
	s.DialogName = &v
	return s
}

func (s *DescribeDialogResponseBody) SetModifyUserId(v string) *DescribeDialogResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeDialogResponseBody) SetModifyUserName(v string) *DescribeDialogResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeDialogResponseBody) SetIsSampleDialog(v bool) *DescribeDialogResponseBody {
	s.IsSampleDialog = &v
	return s
}

type DescribeDialogResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDialogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDialogResponse) SetHeaders(v map[string]*string) *DescribeDialogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDialogResponse) SetBody(v *DescribeDialogResponseBody) *DescribeDialogResponse {
	s.Body = v
	return s
}

type DescribePerspectiveRequest struct {
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s DescribePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveRequest) SetPerspectiveId(v string) *DescribePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type DescribePerspectiveResponseBody struct {
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	PerspectiveCode *string `json:"PerspectiveCode,omitempty" xml:"PerspectiveCode,omitempty"`
	ModifyTime      *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SelfDefine      *bool   `json:"SelfDefine,omitempty" xml:"SelfDefine,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyUserName  *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	PerspectiveId   *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	CreateUserName  *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveResponseBody) SetStatus(v int32) *DescribePerspectiveResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetPerspectiveCode(v string) *DescribePerspectiveResponseBody {
	s.PerspectiveCode = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetModifyTime(v string) *DescribePerspectiveResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetRequestId(v string) *DescribePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetSelfDefine(v bool) *DescribePerspectiveResponseBody {
	s.SelfDefine = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetCreateTime(v string) *DescribePerspectiveResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetModifyUserName(v string) *DescribePerspectiveResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetPerspectiveId(v string) *DescribePerspectiveResponseBody {
	s.PerspectiveId = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetCreateUserName(v string) *DescribePerspectiveResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetName(v string) *DescribePerspectiveResponseBody {
	s.Name = &v
	return s
}

type DescribePerspectiveResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveResponse) SetHeaders(v map[string]*string) *DescribePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *DescribePerspectiveResponse) SetBody(v *DescribePerspectiveResponseBody) *DescribePerspectiveResponse {
	s.Body = v
	return s
}

type DisableDialogFlowRequest struct {
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s DisableDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *DisableDialogFlowRequest) SetDialogId(v int64) *DisableDialogFlowRequest {
	s.DialogId = &v
	return s
}

type DisableDialogFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDialogFlowResponseBody) SetRequestId(v string) *DisableDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

type DisableDialogFlowResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *DisableDialogFlowResponse) SetHeaders(v map[string]*string) *DisableDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *DisableDialogFlowResponse) SetBody(v *DisableDialogFlowResponseBody) *DisableDialogFlowResponse {
	s.Body = v
	return s
}

type DisableKnowledgeRequest struct {
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DisableKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DisableKnowledgeRequest) SetKnowledgeId(v int64) *DisableKnowledgeRequest {
	s.KnowledgeId = &v
	return s
}

type DisableKnowledgeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *DisableKnowledgeResponseBody) SetRequestId(v string) *DisableKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

type DisableKnowledgeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *DisableKnowledgeResponse) SetHeaders(v map[string]*string) *DisableKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *DisableKnowledgeResponse) SetBody(v *DisableKnowledgeResponseBody) *DisableKnowledgeResponse {
	s.Body = v
	return s
}

type FeedbackRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	MessageId  *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Feedback   *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
}

func (s FeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackRequest) GoString() string {
	return s.String()
}

func (s *FeedbackRequest) SetInstanceId(v string) *FeedbackRequest {
	s.InstanceId = &v
	return s
}

func (s *FeedbackRequest) SetSessionId(v string) *FeedbackRequest {
	s.SessionId = &v
	return s
}

func (s *FeedbackRequest) SetMessageId(v string) *FeedbackRequest {
	s.MessageId = &v
	return s
}

func (s *FeedbackRequest) SetFeedback(v string) *FeedbackRequest {
	s.Feedback = &v
	return s
}

type FeedbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Feedback  *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s FeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *FeedbackResponseBody) SetRequestId(v string) *FeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *FeedbackResponseBody) SetFeedback(v string) *FeedbackResponseBody {
	s.Feedback = &v
	return s
}

func (s *FeedbackResponseBody) SetMessageId(v string) *FeedbackResponseBody {
	s.MessageId = &v
	return s
}

type FeedbackResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackResponse) GoString() string {
	return s.String()
}

func (s *FeedbackResponse) SetHeaders(v map[string]*string) *FeedbackResponse {
	s.Headers = v
	return s
}

func (s *FeedbackResponse) SetBody(v *FeedbackResponseBody) *FeedbackResponse {
	s.Body = v
	return s
}

type GetBotChatDataRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
}

func (s GetBotChatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotChatDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotChatDataRequest) SetCubeId(v string) *GetBotChatDataRequest {
	s.CubeId = &v
	return s
}

func (s *GetBotChatDataRequest) SetMeasures(v string) *GetBotChatDataRequest {
	s.Measures = &v
	return s
}

func (s *GetBotChatDataRequest) SetStartTime(v string) *GetBotChatDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetBotChatDataRequest) SetEndTime(v string) *GetBotChatDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotChatDataRequest) SetRobotInstanceId(v string) *GetBotChatDataRequest {
	s.RobotInstanceId = &v
	return s
}

type GetBotChatDataResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s GetBotChatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotChatDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotChatDataResponseBody) SetCostTime(v string) *GetBotChatDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotChatDataResponseBody) SetRequestId(v string) *GetBotChatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBotChatDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotChatDataResponseBody {
	s.Datas = v
	return s
}

type GetBotChatDataResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBotChatDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBotChatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotChatDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotChatDataResponse) SetHeaders(v map[string]*string) *GetBotChatDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotChatDataResponse) SetBody(v *GetBotChatDataResponseBody) *GetBotChatDataResponse {
	s.Body = v
	return s
}

type GetBotDsStatDataRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
}

func (s GetBotDsStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotDsStatDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotDsStatDataRequest) SetCubeId(v string) *GetBotDsStatDataRequest {
	s.CubeId = &v
	return s
}

func (s *GetBotDsStatDataRequest) SetMeasures(v string) *GetBotDsStatDataRequest {
	s.Measures = &v
	return s
}

func (s *GetBotDsStatDataRequest) SetStartTime(v string) *GetBotDsStatDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetBotDsStatDataRequest) SetEndTime(v string) *GetBotDsStatDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotDsStatDataRequest) SetRobotInstanceId(v string) *GetBotDsStatDataRequest {
	s.RobotInstanceId = &v
	return s
}

type GetBotDsStatDataResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s GetBotDsStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotDsStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotDsStatDataResponseBody) SetCostTime(v string) *GetBotDsStatDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotDsStatDataResponseBody) SetRequestId(v string) *GetBotDsStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBotDsStatDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotDsStatDataResponseBody {
	s.Datas = v
	return s
}

type GetBotDsStatDataResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBotDsStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBotDsStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotDsStatDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotDsStatDataResponse) SetHeaders(v map[string]*string) *GetBotDsStatDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotDsStatDataResponse) SetBody(v *GetBotDsStatDataResponseBody) *GetBotDsStatDataResponse {
	s.Body = v
	return s
}

type GetBotKnowledgeStatDataRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
}

func (s GetBotKnowledgeStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotKnowledgeStatDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotKnowledgeStatDataRequest) SetCubeId(v string) *GetBotKnowledgeStatDataRequest {
	s.CubeId = &v
	return s
}

func (s *GetBotKnowledgeStatDataRequest) SetMeasures(v string) *GetBotKnowledgeStatDataRequest {
	s.Measures = &v
	return s
}

func (s *GetBotKnowledgeStatDataRequest) SetStartTime(v string) *GetBotKnowledgeStatDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetBotKnowledgeStatDataRequest) SetEndTime(v string) *GetBotKnowledgeStatDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotKnowledgeStatDataRequest) SetRobotInstanceId(v string) *GetBotKnowledgeStatDataRequest {
	s.RobotInstanceId = &v
	return s
}

type GetBotKnowledgeStatDataResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s GetBotKnowledgeStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotKnowledgeStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotKnowledgeStatDataResponseBody) SetCostTime(v string) *GetBotKnowledgeStatDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotKnowledgeStatDataResponseBody) SetRequestId(v string) *GetBotKnowledgeStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBotKnowledgeStatDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotKnowledgeStatDataResponseBody {
	s.Datas = v
	return s
}

type GetBotKnowledgeStatDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBotKnowledgeStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBotKnowledgeStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotKnowledgeStatDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotKnowledgeStatDataResponse) SetHeaders(v map[string]*string) *GetBotKnowledgeStatDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotKnowledgeStatDataResponse) SetBody(v *GetBotKnowledgeStatDataResponseBody) *GetBotKnowledgeStatDataResponse {
	s.Body = v
	return s
}

type GetBotSessionDataRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
}

func (s GetBotSessionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotSessionDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataRequest) SetCubeId(v string) *GetBotSessionDataRequest {
	s.CubeId = &v
	return s
}

func (s *GetBotSessionDataRequest) SetMeasures(v string) *GetBotSessionDataRequest {
	s.Measures = &v
	return s
}

func (s *GetBotSessionDataRequest) SetStartTime(v string) *GetBotSessionDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetBotSessionDataRequest) SetEndTime(v string) *GetBotSessionDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotSessionDataRequest) SetRobotInstanceId(v string) *GetBotSessionDataRequest {
	s.RobotInstanceId = &v
	return s
}

type GetBotSessionDataResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s GetBotSessionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBotSessionDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataResponseBody) SetCostTime(v string) *GetBotSessionDataResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetBotSessionDataResponseBody) SetRequestId(v string) *GetBotSessionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBotSessionDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotSessionDataResponseBody {
	s.Datas = v
	return s
}

type GetBotSessionDataResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBotSessionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBotSessionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBotSessionDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataResponse) SetHeaders(v map[string]*string) *GetBotSessionDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotSessionDataResponse) SetBody(v *GetBotSessionDataResponseBody) *GetBotSessionDataResponse {
	s.Body = v
	return s
}

type GetConversationListRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SenderId   *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetConversationListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationListRequest) GoString() string {
	return s.String()
}

func (s *GetConversationListRequest) SetInstanceId(v string) *GetConversationListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConversationListRequest) SetSessionId(v string) *GetConversationListRequest {
	s.SessionId = &v
	return s
}

func (s *GetConversationListRequest) SetSenderId(v string) *GetConversationListRequest {
	s.SenderId = &v
	return s
}

func (s *GetConversationListRequest) SetStartDate(v string) *GetConversationListRequest {
	s.StartDate = &v
	return s
}

func (s *GetConversationListRequest) SetEndDate(v string) *GetConversationListRequest {
	s.EndDate = &v
	return s
}

func (s *GetConversationListRequest) SetPageNumber(v string) *GetConversationListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetConversationListRequest) SetPageSize(v string) *GetConversationListRequest {
	s.PageSize = &v
	return s
}

type GetConversationListResponseBody struct {
	Messages    []map[string]interface{} `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	RequestId   *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int64                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int64                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCounts *int64                   `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s GetConversationListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationListResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationListResponseBody) SetMessages(v []map[string]interface{}) *GetConversationListResponseBody {
	s.Messages = v
	return s
}

func (s *GetConversationListResponseBody) SetRequestId(v string) *GetConversationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConversationListResponseBody) SetPageSize(v int64) *GetConversationListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetConversationListResponseBody) SetPageNumber(v int64) *GetConversationListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetConversationListResponseBody) SetTotalCounts(v int64) *GetConversationListResponseBody {
	s.TotalCounts = &v
	return s
}

type GetConversationListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConversationListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConversationListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationListResponse) GoString() string {
	return s.String()
}

func (s *GetConversationListResponse) SetHeaders(v map[string]*string) *GetConversationListResponse {
	s.Headers = v
	return s
}

func (s *GetConversationListResponse) SetBody(v *GetConversationListResponseBody) *GetConversationListResponse {
	s.Body = v
	return s
}

type ListBotChatHistorysRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	Dimensions      *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	IsDetail        *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	Orders          *string `json:"Orders,omitempty" xml:"Orders,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListBotChatHistorysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotChatHistorysRequest) GoString() string {
	return s.String()
}

func (s *ListBotChatHistorysRequest) SetCubeId(v string) *ListBotChatHistorysRequest {
	s.CubeId = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetStartTime(v string) *ListBotChatHistorysRequest {
	s.StartTime = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetEndTime(v string) *ListBotChatHistorysRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetRobotInstanceId(v string) *ListBotChatHistorysRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetDimensions(v string) *ListBotChatHistorysRequest {
	s.Dimensions = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetIsDetail(v bool) *ListBotChatHistorysRequest {
	s.IsDetail = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetOrders(v string) *ListBotChatHistorysRequest {
	s.Orders = &v
	return s
}

func (s *ListBotChatHistorysRequest) SetLimit(v int32) *ListBotChatHistorysRequest {
	s.Limit = &v
	return s
}

type ListBotChatHistorysResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s ListBotChatHistorysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotChatHistorysResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotChatHistorysResponseBody) SetCostTime(v string) *ListBotChatHistorysResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotChatHistorysResponseBody) SetRequestId(v string) *ListBotChatHistorysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotChatHistorysResponseBody) SetDatas(v []map[string]interface{}) *ListBotChatHistorysResponseBody {
	s.Datas = v
	return s
}

type ListBotChatHistorysResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBotChatHistorysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBotChatHistorysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotChatHistorysResponse) GoString() string {
	return s.String()
}

func (s *ListBotChatHistorysResponse) SetHeaders(v map[string]*string) *ListBotChatHistorysResponse {
	s.Headers = v
	return s
}

func (s *ListBotChatHistorysResponse) SetBody(v *ListBotChatHistorysResponseBody) *ListBotChatHistorysResponse {
	s.Body = v
	return s
}

type ListBotColdDsDatasRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	Dimensions      *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Filters         *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListBotColdDsDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdDsDatasRequest) GoString() string {
	return s.String()
}

func (s *ListBotColdDsDatasRequest) SetCubeId(v string) *ListBotColdDsDatasRequest {
	s.CubeId = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetStartTime(v string) *ListBotColdDsDatasRequest {
	s.StartTime = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetEndTime(v string) *ListBotColdDsDatasRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetRobotInstanceId(v string) *ListBotColdDsDatasRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetDimensions(v string) *ListBotColdDsDatasRequest {
	s.Dimensions = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetFilters(v string) *ListBotColdDsDatasRequest {
	s.Filters = &v
	return s
}

func (s *ListBotColdDsDatasRequest) SetLimit(v int32) *ListBotColdDsDatasRequest {
	s.Limit = &v
	return s
}

type ListBotColdDsDatasResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s ListBotColdDsDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdDsDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotColdDsDatasResponseBody) SetCostTime(v string) *ListBotColdDsDatasResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotColdDsDatasResponseBody) SetRequestId(v string) *ListBotColdDsDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotColdDsDatasResponseBody) SetDatas(v []map[string]interface{}) *ListBotColdDsDatasResponseBody {
	s.Datas = v
	return s
}

type ListBotColdDsDatasResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBotColdDsDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBotColdDsDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdDsDatasResponse) GoString() string {
	return s.String()
}

func (s *ListBotColdDsDatasResponse) SetHeaders(v map[string]*string) *ListBotColdDsDatasResponse {
	s.Headers = v
	return s
}

func (s *ListBotColdDsDatasResponse) SetBody(v *ListBotColdDsDatasResponseBody) *ListBotColdDsDatasResponse {
	s.Body = v
	return s
}

type ListBotColdKnowledgesRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	Dimensions      *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Filters         *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListBotColdKnowledgesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdKnowledgesRequest) GoString() string {
	return s.String()
}

func (s *ListBotColdKnowledgesRequest) SetCubeId(v string) *ListBotColdKnowledgesRequest {
	s.CubeId = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetStartTime(v string) *ListBotColdKnowledgesRequest {
	s.StartTime = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetEndTime(v string) *ListBotColdKnowledgesRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetRobotInstanceId(v string) *ListBotColdKnowledgesRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetDimensions(v string) *ListBotColdKnowledgesRequest {
	s.Dimensions = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetFilters(v string) *ListBotColdKnowledgesRequest {
	s.Filters = &v
	return s
}

func (s *ListBotColdKnowledgesRequest) SetLimit(v int32) *ListBotColdKnowledgesRequest {
	s.Limit = &v
	return s
}

type ListBotColdKnowledgesResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s ListBotColdKnowledgesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdKnowledgesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotColdKnowledgesResponseBody) SetCostTime(v string) *ListBotColdKnowledgesResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotColdKnowledgesResponseBody) SetRequestId(v string) *ListBotColdKnowledgesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotColdKnowledgesResponseBody) SetDatas(v []map[string]interface{}) *ListBotColdKnowledgesResponseBody {
	s.Datas = v
	return s
}

type ListBotColdKnowledgesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBotColdKnowledgesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBotColdKnowledgesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotColdKnowledgesResponse) GoString() string {
	return s.String()
}

func (s *ListBotColdKnowledgesResponse) SetHeaders(v map[string]*string) *ListBotColdKnowledgesResponse {
	s.Headers = v
	return s
}

func (s *ListBotColdKnowledgesResponse) SetBody(v *ListBotColdKnowledgesResponseBody) *ListBotColdKnowledgesResponse {
	s.Body = v
	return s
}

type ListBotDsDetailsRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	Dimensions      *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Orders          *string `json:"Orders,omitempty" xml:"Orders,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListBotDsDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotDsDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListBotDsDetailsRequest) SetCubeId(v string) *ListBotDsDetailsRequest {
	s.CubeId = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetMeasures(v string) *ListBotDsDetailsRequest {
	s.Measures = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetStartTime(v string) *ListBotDsDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetEndTime(v string) *ListBotDsDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetRobotInstanceId(v string) *ListBotDsDetailsRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetDimensions(v string) *ListBotDsDetailsRequest {
	s.Dimensions = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetOrders(v string) *ListBotDsDetailsRequest {
	s.Orders = &v
	return s
}

func (s *ListBotDsDetailsRequest) SetLimit(v int32) *ListBotDsDetailsRequest {
	s.Limit = &v
	return s
}

type ListBotDsDetailsResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s ListBotDsDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotDsDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotDsDetailsResponseBody) SetCostTime(v string) *ListBotDsDetailsResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotDsDetailsResponseBody) SetRequestId(v string) *ListBotDsDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotDsDetailsResponseBody) SetDatas(v []map[string]interface{}) *ListBotDsDetailsResponseBody {
	s.Datas = v
	return s
}

type ListBotDsDetailsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBotDsDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBotDsDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotDsDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListBotDsDetailsResponse) SetHeaders(v map[string]*string) *ListBotDsDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListBotDsDetailsResponse) SetBody(v *ListBotDsDetailsResponseBody) *ListBotDsDetailsResponse {
	s.Body = v
	return s
}

type ListBotHotDsDatasRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	Orders          *string `json:"Orders,omitempty" xml:"Orders,omitempty"`
	Dimensions      *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Filters         *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListBotHotDsDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotDsDatasRequest) GoString() string {
	return s.String()
}

func (s *ListBotHotDsDatasRequest) SetCubeId(v string) *ListBotHotDsDatasRequest {
	s.CubeId = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetMeasures(v string) *ListBotHotDsDatasRequest {
	s.Measures = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetStartTime(v string) *ListBotHotDsDatasRequest {
	s.StartTime = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetEndTime(v string) *ListBotHotDsDatasRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetRobotInstanceId(v string) *ListBotHotDsDatasRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetOrders(v string) *ListBotHotDsDatasRequest {
	s.Orders = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetDimensions(v string) *ListBotHotDsDatasRequest {
	s.Dimensions = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetFilters(v string) *ListBotHotDsDatasRequest {
	s.Filters = &v
	return s
}

func (s *ListBotHotDsDatasRequest) SetLimit(v int32) *ListBotHotDsDatasRequest {
	s.Limit = &v
	return s
}

type ListBotHotDsDatasResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s ListBotHotDsDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotDsDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotHotDsDatasResponseBody) SetCostTime(v string) *ListBotHotDsDatasResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotHotDsDatasResponseBody) SetRequestId(v string) *ListBotHotDsDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotHotDsDatasResponseBody) SetDatas(v []map[string]interface{}) *ListBotHotDsDatasResponseBody {
	s.Datas = v
	return s
}

type ListBotHotDsDatasResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBotHotDsDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBotHotDsDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotDsDatasResponse) GoString() string {
	return s.String()
}

func (s *ListBotHotDsDatasResponse) SetHeaders(v map[string]*string) *ListBotHotDsDatasResponse {
	s.Headers = v
	return s
}

func (s *ListBotHotDsDatasResponse) SetBody(v *ListBotHotDsDatasResponseBody) *ListBotHotDsDatasResponse {
	s.Body = v
	return s
}

type ListBotHotKnowledgesRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	Dimensions      *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Orders          *string `json:"Orders,omitempty" xml:"Orders,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	Filters         *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListBotHotKnowledgesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotKnowledgesRequest) GoString() string {
	return s.String()
}

func (s *ListBotHotKnowledgesRequest) SetCubeId(v string) *ListBotHotKnowledgesRequest {
	s.CubeId = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetStartTime(v string) *ListBotHotKnowledgesRequest {
	s.StartTime = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetEndTime(v string) *ListBotHotKnowledgesRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetRobotInstanceId(v string) *ListBotHotKnowledgesRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetDimensions(v string) *ListBotHotKnowledgesRequest {
	s.Dimensions = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetOrders(v string) *ListBotHotKnowledgesRequest {
	s.Orders = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetMeasures(v string) *ListBotHotKnowledgesRequest {
	s.Measures = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetFilters(v string) *ListBotHotKnowledgesRequest {
	s.Filters = &v
	return s
}

func (s *ListBotHotKnowledgesRequest) SetLimit(v int32) *ListBotHotKnowledgesRequest {
	s.Limit = &v
	return s
}

type ListBotHotKnowledgesResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s ListBotHotKnowledgesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotKnowledgesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotHotKnowledgesResponseBody) SetCostTime(v string) *ListBotHotKnowledgesResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotHotKnowledgesResponseBody) SetRequestId(v string) *ListBotHotKnowledgesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotHotKnowledgesResponseBody) SetDatas(v []map[string]interface{}) *ListBotHotKnowledgesResponseBody {
	s.Datas = v
	return s
}

type ListBotHotKnowledgesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBotHotKnowledgesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBotHotKnowledgesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotHotKnowledgesResponse) GoString() string {
	return s.String()
}

func (s *ListBotHotKnowledgesResponse) SetHeaders(v map[string]*string) *ListBotHotKnowledgesResponse {
	s.Headers = v
	return s
}

func (s *ListBotHotKnowledgesResponse) SetBody(v *ListBotHotKnowledgesResponseBody) *ListBotHotKnowledgesResponse {
	s.Body = v
	return s
}

type ListBotKnowledgeDetailsRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	Limit           *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Dimensions      *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Orders          *string `json:"Orders,omitempty" xml:"Orders,omitempty"`
}

func (s ListBotKnowledgeDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotKnowledgeDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListBotKnowledgeDetailsRequest) SetCubeId(v string) *ListBotKnowledgeDetailsRequest {
	s.CubeId = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetMeasures(v string) *ListBotKnowledgeDetailsRequest {
	s.Measures = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetStartTime(v string) *ListBotKnowledgeDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetEndTime(v string) *ListBotKnowledgeDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetRobotInstanceId(v string) *ListBotKnowledgeDetailsRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetLimit(v string) *ListBotKnowledgeDetailsRequest {
	s.Limit = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetDimensions(v string) *ListBotKnowledgeDetailsRequest {
	s.Dimensions = &v
	return s
}

func (s *ListBotKnowledgeDetailsRequest) SetOrders(v string) *ListBotKnowledgeDetailsRequest {
	s.Orders = &v
	return s
}

type ListBotKnowledgeDetailsResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s ListBotKnowledgeDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotKnowledgeDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotKnowledgeDetailsResponseBody) SetCostTime(v string) *ListBotKnowledgeDetailsResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotKnowledgeDetailsResponseBody) SetRequestId(v string) *ListBotKnowledgeDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotKnowledgeDetailsResponseBody) SetDatas(v []map[string]interface{}) *ListBotKnowledgeDetailsResponseBody {
	s.Datas = v
	return s
}

type ListBotKnowledgeDetailsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBotKnowledgeDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBotKnowledgeDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotKnowledgeDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListBotKnowledgeDetailsResponse) SetHeaders(v map[string]*string) *ListBotKnowledgeDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListBotKnowledgeDetailsResponse) SetBody(v *ListBotKnowledgeDetailsResponseBody) *ListBotKnowledgeDetailsResponse {
	s.Body = v
	return s
}

type ListBotReceptionDetailDatasRequest struct {
	CubeId          *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	Measures        *string `json:"Measures,omitempty" xml:"Measures,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	Dimensions      *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
}

func (s ListBotReceptionDetailDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBotReceptionDetailDatasRequest) GoString() string {
	return s.String()
}

func (s *ListBotReceptionDetailDatasRequest) SetCubeId(v string) *ListBotReceptionDetailDatasRequest {
	s.CubeId = &v
	return s
}

func (s *ListBotReceptionDetailDatasRequest) SetMeasures(v string) *ListBotReceptionDetailDatasRequest {
	s.Measures = &v
	return s
}

func (s *ListBotReceptionDetailDatasRequest) SetStartTime(v string) *ListBotReceptionDetailDatasRequest {
	s.StartTime = &v
	return s
}

func (s *ListBotReceptionDetailDatasRequest) SetEndTime(v string) *ListBotReceptionDetailDatasRequest {
	s.EndTime = &v
	return s
}

func (s *ListBotReceptionDetailDatasRequest) SetRobotInstanceId(v string) *ListBotReceptionDetailDatasRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListBotReceptionDetailDatasRequest) SetDimensions(v string) *ListBotReceptionDetailDatasRequest {
	s.Dimensions = &v
	return s
}

type ListBotReceptionDetailDatasResponseBody struct {
	CostTime  *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datas     []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
}

func (s ListBotReceptionDetailDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBotReceptionDetailDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotReceptionDetailDatasResponseBody) SetCostTime(v string) *ListBotReceptionDetailDatasResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListBotReceptionDetailDatasResponseBody) SetRequestId(v string) *ListBotReceptionDetailDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotReceptionDetailDatasResponseBody) SetDatas(v []map[string]interface{}) *ListBotReceptionDetailDatasResponseBody {
	s.Datas = v
	return s
}

type ListBotReceptionDetailDatasResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBotReceptionDetailDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBotReceptionDetailDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBotReceptionDetailDatasResponse) GoString() string {
	return s.String()
}

func (s *ListBotReceptionDetailDatasResponse) SetHeaders(v map[string]*string) *ListBotReceptionDetailDatasResponse {
	s.Headers = v
	return s
}

func (s *ListBotReceptionDetailDatasResponse) SetBody(v *ListBotReceptionDetailDatasResponseBody) *ListBotReceptionDetailDatasResponse {
	s.Body = v
	return s
}

type ListConversationLogsRequest struct {
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ListConversationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConversationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListConversationLogsRequest) SetSessionId(v string) *ListConversationLogsRequest {
	s.SessionId = &v
	return s
}

type ListConversationLogsResponseBody struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChatLogs  []map[string]interface{} `json:"ChatLogs,omitempty" xml:"ChatLogs,omitempty" type:"Repeated"`
	Rounds    *int64                   `json:"Rounds,omitempty" xml:"Rounds,omitempty"`
}

func (s ListConversationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConversationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConversationLogsResponseBody) SetRequestId(v string) *ListConversationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConversationLogsResponseBody) SetChatLogs(v []map[string]interface{}) *ListConversationLogsResponseBody {
	s.ChatLogs = v
	return s
}

func (s *ListConversationLogsResponseBody) SetRounds(v int64) *ListConversationLogsResponseBody {
	s.Rounds = &v
	return s
}

type ListConversationLogsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConversationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConversationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConversationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListConversationLogsResponse) SetHeaders(v map[string]*string) *ListConversationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListConversationLogsResponse) SetBody(v *ListConversationLogsResponseBody) *ListConversationLogsResponse {
	s.Body = v
	return s
}

type MoveKnowledgeCategoryRequest struct {
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	CategoryId  *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s MoveKnowledgeCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveKnowledgeCategoryRequest) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeCategoryRequest) SetKnowledgeId(v int64) *MoveKnowledgeCategoryRequest {
	s.KnowledgeId = &v
	return s
}

func (s *MoveKnowledgeCategoryRequest) SetCategoryId(v int64) *MoveKnowledgeCategoryRequest {
	s.CategoryId = &v
	return s
}

type MoveKnowledgeCategoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveKnowledgeCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveKnowledgeCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeCategoryResponseBody) SetRequestId(v string) *MoveKnowledgeCategoryResponseBody {
	s.RequestId = &v
	return s
}

type MoveKnowledgeCategoryResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveKnowledgeCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveKnowledgeCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveKnowledgeCategoryResponse) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeCategoryResponse) SetHeaders(v map[string]*string) *MoveKnowledgeCategoryResponse {
	s.Headers = v
	return s
}

func (s *MoveKnowledgeCategoryResponse) SetBody(v *MoveKnowledgeCategoryResponseBody) *MoveKnowledgeCategoryResponse {
	s.Body = v
	return s
}

type PublishDialogFlowRequest struct {
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s PublishDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *PublishDialogFlowRequest) SetDialogId(v int64) *PublishDialogFlowRequest {
	s.DialogId = &v
	return s
}

type PublishDialogFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *PublishDialogFlowResponseBody) SetRequestId(v string) *PublishDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

type PublishDialogFlowResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *PublishDialogFlowResponse) SetHeaders(v map[string]*string) *PublishDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *PublishDialogFlowResponse) SetBody(v *PublishDialogFlowResponseBody) *PublishDialogFlowResponse {
	s.Body = v
	return s
}

type PublishKnowledgeRequest struct {
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s PublishKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *PublishKnowledgeRequest) SetKnowledgeId(v int64) *PublishKnowledgeRequest {
	s.KnowledgeId = &v
	return s
}

type PublishKnowledgeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *PublishKnowledgeResponseBody) SetRequestId(v string) *PublishKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

type PublishKnowledgeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *PublishKnowledgeResponse) SetHeaders(v map[string]*string) *PublishKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *PublishKnowledgeResponse) SetBody(v *PublishKnowledgeResponseBody) *PublishKnowledgeResponse {
	s.Body = v
	return s
}

type QueryBotsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryBotsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBotsRequest) GoString() string {
	return s.String()
}

func (s *QueryBotsRequest) SetPageNumber(v int32) *QueryBotsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryBotsRequest) SetPageSize(v int32) *QueryBotsRequest {
	s.PageSize = &v
	return s
}

type QueryBotsResponseBody struct {
	TotalCount *int32                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Bots       []*QueryBotsResponseBodyBots `json:"Bots,omitempty" xml:"Bots,omitempty" type:"Repeated"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s QueryBotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBotsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBotsResponseBody) SetTotalCount(v int32) *QueryBotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryBotsResponseBody) SetBots(v []*QueryBotsResponseBodyBots) *QueryBotsResponseBody {
	s.Bots = v
	return s
}

func (s *QueryBotsResponseBody) SetRequestId(v string) *QueryBotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBotsResponseBody) SetPageSize(v int32) *QueryBotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBotsResponseBody) SetPageNumber(v int32) *QueryBotsResponseBody {
	s.PageNumber = &v
	return s
}

type QueryBotsResponseBodyBots struct {
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	Avatar       *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	TimeZone     *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryBotsResponseBodyBots) String() string {
	return tea.Prettify(s)
}

func (s QueryBotsResponseBodyBots) GoString() string {
	return s.String()
}

func (s *QueryBotsResponseBodyBots) SetIntroduction(v string) *QueryBotsResponseBodyBots {
	s.Introduction = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetAvatar(v string) *QueryBotsResponseBodyBots {
	s.Avatar = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetTimeZone(v string) *QueryBotsResponseBodyBots {
	s.TimeZone = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetCreateTime(v string) *QueryBotsResponseBodyBots {
	s.CreateTime = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetLanguageCode(v string) *QueryBotsResponseBodyBots {
	s.LanguageCode = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetInstanceId(v string) *QueryBotsResponseBodyBots {
	s.InstanceId = &v
	return s
}

func (s *QueryBotsResponseBodyBots) SetName(v string) *QueryBotsResponseBodyBots {
	s.Name = &v
	return s
}

type QueryBotsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBotsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBotsResponse) GoString() string {
	return s.String()
}

func (s *QueryBotsResponse) SetHeaders(v map[string]*string) *QueryBotsResponse {
	s.Headers = v
	return s
}

func (s *QueryBotsResponse) SetBody(v *QueryBotsResponseBody) *QueryBotsResponse {
	s.Body = v
	return s
}

type QueryDialogsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DialogName *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryDialogsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDialogsRequest) GoString() string {
	return s.String()
}

func (s *QueryDialogsRequest) SetInstanceId(v string) *QueryDialogsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryDialogsRequest) SetDialogName(v string) *QueryDialogsRequest {
	s.DialogName = &v
	return s
}

func (s *QueryDialogsRequest) SetPageNumber(v int32) *QueryDialogsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryDialogsRequest) SetPageSize(v int32) *QueryDialogsRequest {
	s.PageSize = &v
	return s
}

type QueryDialogsResponseBody struct {
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Dialogs    []*QueryDialogsResponseBodyDialogs `json:"Dialogs,omitempty" xml:"Dialogs,omitempty" type:"Repeated"`
}

func (s QueryDialogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDialogsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDialogsResponseBody) SetTotalCount(v int32) *QueryDialogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryDialogsResponseBody) SetRequestId(v string) *QueryDialogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDialogsResponseBody) SetPageSize(v int32) *QueryDialogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDialogsResponseBody) SetPageNumber(v int32) *QueryDialogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryDialogsResponseBody) SetDialogs(v []*QueryDialogsResponseBodyDialogs) *QueryDialogsResponseBody {
	s.Dialogs = v
	return s
}

type QueryDialogsResponseBodyDialogs struct {
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	DialogName     *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	ModifyUserId   *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	IsOnline       *bool   `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId   *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DialogId       *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	IsSampleDialog *bool   `json:"IsSampleDialog,omitempty" xml:"IsSampleDialog,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s QueryDialogsResponseBodyDialogs) String() string {
	return tea.Prettify(s)
}

func (s QueryDialogsResponseBodyDialogs) GoString() string {
	return s.String()
}

func (s *QueryDialogsResponseBodyDialogs) SetStatus(v int32) *QueryDialogsResponseBodyDialogs {
	s.Status = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetDialogName(v string) *QueryDialogsResponseBodyDialogs {
	s.DialogName = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetModifyUserId(v string) *QueryDialogsResponseBodyDialogs {
	s.ModifyUserId = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetIsOnline(v bool) *QueryDialogsResponseBodyDialogs {
	s.IsOnline = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetCreateUserName(v string) *QueryDialogsResponseBodyDialogs {
	s.CreateUserName = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetCreateTime(v string) *QueryDialogsResponseBodyDialogs {
	s.CreateTime = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetCreateUserId(v string) *QueryDialogsResponseBodyDialogs {
	s.CreateUserId = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetModifyUserName(v string) *QueryDialogsResponseBodyDialogs {
	s.ModifyUserName = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetDescription(v string) *QueryDialogsResponseBodyDialogs {
	s.Description = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetDialogId(v int64) *QueryDialogsResponseBodyDialogs {
	s.DialogId = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetIsSampleDialog(v bool) *QueryDialogsResponseBodyDialogs {
	s.IsSampleDialog = &v
	return s
}

func (s *QueryDialogsResponseBodyDialogs) SetModifyTime(v string) *QueryDialogsResponseBodyDialogs {
	s.ModifyTime = &v
	return s
}

type QueryDialogsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDialogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDialogsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDialogsResponse) GoString() string {
	return s.String()
}

func (s *QueryDialogsResponse) SetHeaders(v map[string]*string) *QueryDialogsResponse {
	s.Headers = v
	return s
}

func (s *QueryDialogsResponse) SetBody(v *QueryDialogsResponseBody) *QueryDialogsResponse {
	s.Body = v
	return s
}

type QueryPerspectivesResponseBody struct {
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Perspectives []*QueryPerspectivesResponseBodyPerspectives `json:"Perspectives,omitempty" xml:"Perspectives,omitempty" type:"Repeated"`
}

func (s QueryPerspectivesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponseBody) SetRequestId(v string) *QueryPerspectivesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPerspectivesResponseBody) SetPerspectives(v []*QueryPerspectivesResponseBodyPerspectives) *QueryPerspectivesResponseBody {
	s.Perspectives = v
	return s
}

type QueryPerspectivesResponseBodyPerspectives struct {
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ModifyUserName  *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	CreateUserName  *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PerspectiveId   *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	SelfDefine      *string `json:"SelfDefine,omitempty" xml:"SelfDefine,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PerspectiveCode *string `json:"PerspectiveCode,omitempty" xml:"PerspectiveCode,omitempty"`
	ModifyTime      *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s QueryPerspectivesResponseBodyPerspectives) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesResponseBodyPerspectives) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetStatus(v int32) *QueryPerspectivesResponseBodyPerspectives {
	s.Status = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetModifyUserName(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.ModifyUserName = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetCreateUserName(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.CreateUserName = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetCreateTime(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.CreateTime = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetPerspectiveId(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.PerspectiveId = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetSelfDefine(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.SelfDefine = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetName(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.Name = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetPerspectiveCode(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.PerspectiveCode = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetModifyTime(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.ModifyTime = &v
	return s
}

type QueryPerspectivesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPerspectivesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPerspectivesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesResponse) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponse) SetHeaders(v map[string]*string) *QueryPerspectivesResponse {
	s.Headers = v
	return s
}

func (s *QueryPerspectivesResponse) SetBody(v *QueryPerspectivesResponseBody) *QueryPerspectivesResponse {
	s.Body = v
	return s
}

type QuerySystemEntitiesRequest struct {
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s QuerySystemEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEntitiesRequest) GoString() string {
	return s.String()
}

func (s *QuerySystemEntitiesRequest) SetEntityName(v string) *QuerySystemEntitiesRequest {
	s.EntityName = &v
	return s
}

type QuerySystemEntitiesResponseBody struct {
	SystemEntities []*QuerySystemEntitiesResponseBodySystemEntities `json:"SystemEntities,omitempty" xml:"SystemEntities,omitempty" type:"Repeated"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySystemEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySystemEntitiesResponseBody) SetSystemEntities(v []*QuerySystemEntitiesResponseBodySystemEntities) *QuerySystemEntitiesResponseBody {
	s.SystemEntities = v
	return s
}

func (s *QuerySystemEntitiesResponseBody) SetRequestId(v string) *QuerySystemEntitiesResponseBody {
	s.RequestId = &v
	return s
}

type QuerySystemEntitiesResponseBodySystemEntities struct {
	DefaultQuestion *string `json:"DefaultQuestion,omitempty" xml:"DefaultQuestion,omitempty"`
	EntityCode      *string `json:"EntityCode,omitempty" xml:"EntityCode,omitempty"`
	EntityName      *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s QuerySystemEntitiesResponseBodySystemEntities) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEntitiesResponseBodySystemEntities) GoString() string {
	return s.String()
}

func (s *QuerySystemEntitiesResponseBodySystemEntities) SetDefaultQuestion(v string) *QuerySystemEntitiesResponseBodySystemEntities {
	s.DefaultQuestion = &v
	return s
}

func (s *QuerySystemEntitiesResponseBodySystemEntities) SetEntityCode(v string) *QuerySystemEntitiesResponseBodySystemEntities {
	s.EntityCode = &v
	return s
}

func (s *QuerySystemEntitiesResponseBodySystemEntities) SetEntityName(v string) *QuerySystemEntitiesResponseBodySystemEntities {
	s.EntityName = &v
	return s
}

type QuerySystemEntitiesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySystemEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySystemEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEntitiesResponse) GoString() string {
	return s.String()
}

func (s *QuerySystemEntitiesResponse) SetHeaders(v map[string]*string) *QuerySystemEntitiesResponse {
	s.Headers = v
	return s
}

func (s *QuerySystemEntitiesResponse) SetBody(v *QuerySystemEntitiesResponseBody) *QuerySystemEntitiesResponse {
	s.Body = v
	return s
}

type RemoveEntityMemberRequest struct {
	EntityId   *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RemoveType *string `json:"RemoveType,omitempty" xml:"RemoveType,omitempty"`
	Member     *string `json:"Member,omitempty" xml:"Member,omitempty"`
}

func (s RemoveEntityMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntityMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntityMemberRequest) SetEntityId(v int64) *RemoveEntityMemberRequest {
	s.EntityId = &v
	return s
}

func (s *RemoveEntityMemberRequest) SetRemoveType(v string) *RemoveEntityMemberRequest {
	s.RemoveType = &v
	return s
}

func (s *RemoveEntityMemberRequest) SetMember(v string) *RemoveEntityMemberRequest {
	s.Member = &v
	return s
}

type RemoveEntityMemberResponseBody struct {
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveEntityMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntityMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntityMemberResponseBody) SetEntityId(v string) *RemoveEntityMemberResponseBody {
	s.EntityId = &v
	return s
}

func (s *RemoveEntityMemberResponseBody) SetRequestId(v string) *RemoveEntityMemberResponseBody {
	s.RequestId = &v
	return s
}

type RemoveEntityMemberResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveEntityMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveEntityMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntityMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveEntityMemberResponse) SetHeaders(v map[string]*string) *RemoveEntityMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveEntityMemberResponse) SetBody(v *RemoveEntityMemberResponseBody) *RemoveEntityMemberResponse {
	s.Body = v
	return s
}

type RemoveSynonymRequest struct {
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	Synonym      *string `json:"Synonym,omitempty" xml:"Synonym,omitempty"`
}

func (s RemoveSynonymRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSynonymRequest) GoString() string {
	return s.String()
}

func (s *RemoveSynonymRequest) SetCoreWordName(v string) *RemoveSynonymRequest {
	s.CoreWordName = &v
	return s
}

func (s *RemoveSynonymRequest) SetSynonym(v string) *RemoveSynonymRequest {
	s.Synonym = &v
	return s
}

type RemoveSynonymResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSynonymResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSynonymResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSynonymResponseBody) SetRequestId(v string) *RemoveSynonymResponseBody {
	s.RequestId = &v
	return s
}

type RemoveSynonymResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveSynonymResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSynonymResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSynonymResponse) GoString() string {
	return s.String()
}

func (s *RemoveSynonymResponse) SetHeaders(v map[string]*string) *RemoveSynonymResponse {
	s.Headers = v
	return s
}

func (s *RemoveSynonymResponse) SetBody(v *RemoveSynonymResponseBody) *RemoveSynonymResponse {
	s.Body = v
	return s
}

type TestDialogFlowRequest struct {
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
}

func (s TestDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s TestDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *TestDialogFlowRequest) SetDialogId(v int64) *TestDialogFlowRequest {
	s.DialogId = &v
	return s
}

type TestDialogFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TestDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *TestDialogFlowResponseBody) SetRequestId(v string) *TestDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

type TestDialogFlowResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TestDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TestDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s TestDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *TestDialogFlowResponse) SetHeaders(v map[string]*string) *TestDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *TestDialogFlowResponse) SetBody(v *TestDialogFlowResponseBody) *TestDialogFlowResponse {
	s.Body = v
	return s
}

type UpdateCategoryRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s UpdateCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategoryRequest) SetName(v string) *UpdateCategoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateCategoryRequest) SetCategoryId(v int64) *UpdateCategoryRequest {
	s.CategoryId = &v
	return s
}

type UpdateCategoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCategoryResponseBody) SetRequestId(v string) *UpdateCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCategoryResponseBody) SetSuccess(v bool) *UpdateCategoryResponseBody {
	s.Success = &v
	return s
}

type UpdateCategoryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateCategoryResponse) SetHeaders(v map[string]*string) *UpdateCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateCategoryResponse) SetBody(v *UpdateCategoryResponseBody) *UpdateCategoryResponse {
	s.Body = v
	return s
}

type UpdateCoreWordRequest struct {
	CoreWordName *string `json:"CoreWordName,omitempty" xml:"CoreWordName,omitempty"`
	CoreWordCode *string `json:"CoreWordCode,omitempty" xml:"CoreWordCode,omitempty"`
}

func (s UpdateCoreWordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoreWordRequest) GoString() string {
	return s.String()
}

func (s *UpdateCoreWordRequest) SetCoreWordName(v string) *UpdateCoreWordRequest {
	s.CoreWordName = &v
	return s
}

func (s *UpdateCoreWordRequest) SetCoreWordCode(v string) *UpdateCoreWordRequest {
	s.CoreWordCode = &v
	return s
}

type UpdateCoreWordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCoreWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoreWordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCoreWordResponseBody) SetRequestId(v string) *UpdateCoreWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCoreWordResponseBody) SetSuccess(v bool) *UpdateCoreWordResponseBody {
	s.Success = &v
	return s
}

type UpdateCoreWordResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCoreWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCoreWordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoreWordResponse) GoString() string {
	return s.String()
}

func (s *UpdateCoreWordResponse) SetHeaders(v map[string]*string) *UpdateCoreWordResponse {
	s.Headers = v
	return s
}

func (s *UpdateCoreWordResponse) SetBody(v *UpdateCoreWordResponseBody) *UpdateCoreWordResponse {
	s.Body = v
	return s
}

type UpdateDialogRequest struct {
	DialogId    *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	DialogName  *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogRequest) GoString() string {
	return s.String()
}

func (s *UpdateDialogRequest) SetDialogId(v int64) *UpdateDialogRequest {
	s.DialogId = &v
	return s
}

func (s *UpdateDialogRequest) SetDialogName(v string) *UpdateDialogRequest {
	s.DialogName = &v
	return s
}

func (s *UpdateDialogRequest) SetDescription(v string) *UpdateDialogRequest {
	s.Description = &v
	return s
}

type UpdateDialogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDialogResponseBody) SetRequestId(v string) *UpdateDialogResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDialogResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDialogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogResponse) GoString() string {
	return s.String()
}

func (s *UpdateDialogResponse) SetHeaders(v map[string]*string) *UpdateDialogResponse {
	s.Headers = v
	return s
}

func (s *UpdateDialogResponse) SetBody(v *UpdateDialogResponseBody) *UpdateDialogResponse {
	s.Body = v
	return s
}

type UpdateDialogFlowRequest struct {
	DialogId         *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	ModuleDefinition *string `json:"ModuleDefinition,omitempty" xml:"ModuleDefinition,omitempty"`
}

func (s UpdateDialogFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowRequest) SetDialogId(v int64) *UpdateDialogFlowRequest {
	s.DialogId = &v
	return s
}

func (s *UpdateDialogFlowRequest) SetModuleDefinition(v string) *UpdateDialogFlowRequest {
	s.ModuleDefinition = &v
	return s
}

type UpdateDialogFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDialogFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowResponseBody) SetRequestId(v string) *UpdateDialogFlowResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDialogFlowResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDialogFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDialogFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDialogFlowResponse) GoString() string {
	return s.String()
}

func (s *UpdateDialogFlowResponse) SetHeaders(v map[string]*string) *UpdateDialogFlowResponse {
	s.Headers = v
	return s
}

func (s *UpdateDialogFlowResponse) SetBody(v *UpdateDialogFlowResponseBody) *UpdateDialogFlowResponse {
	s.Body = v
	return s
}

type UpdateEntityRequest struct {
	EntityId   *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Regex      *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
	Members    *string `json:"Members,omitempty" xml:"Members,omitempty"`
}

func (s UpdateEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityRequest) SetEntityId(v int64) *UpdateEntityRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateEntityRequest) SetEntityName(v string) *UpdateEntityRequest {
	s.EntityName = &v
	return s
}

func (s *UpdateEntityRequest) SetEntityType(v string) *UpdateEntityRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityRequest) SetRegex(v string) *UpdateEntityRequest {
	s.Regex = &v
	return s
}

func (s *UpdateEntityRequest) SetMembers(v string) *UpdateEntityRequest {
	s.Members = &v
	return s
}

type UpdateEntityResponseBody struct {
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEntityResponseBody) SetEntityId(v string) *UpdateEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *UpdateEntityResponseBody) SetRequestId(v string) *UpdateEntityResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEntityResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateEntityResponse) SetHeaders(v map[string]*string) *UpdateEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateEntityResponse) SetBody(v *UpdateEntityResponseBody) *UpdateEntityResponse {
	s.Body = v
	return s
}

type UpdateIntentRequest struct {
	IntentDefinition *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
	IntentId         *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s UpdateIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntentRequest) SetIntentDefinition(v string) *UpdateIntentRequest {
	s.IntentDefinition = &v
	return s
}

func (s *UpdateIntentRequest) SetIntentId(v int64) *UpdateIntentRequest {
	s.IntentId = &v
	return s
}

type UpdateIntentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IntentId  *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s UpdateIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntentResponseBody) SetRequestId(v string) *UpdateIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIntentResponseBody) SetIntentId(v string) *UpdateIntentResponseBody {
	s.IntentId = &v
	return s
}

type UpdateIntentResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIntentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntentResponse) SetHeaders(v map[string]*string) *UpdateIntentResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntentResponse) SetBody(v *UpdateIntentResponseBody) *UpdateIntentResponse {
	s.Body = v
	return s
}

type UpdateKnowledgeRequest struct {
	Knowledge *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
}

func (s UpdateKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeRequest) SetKnowledge(v string) *UpdateKnowledgeRequest {
	s.Knowledge = &v
	return s
}

type UpdateKnowledgeResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KnowledgeId *int64  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s UpdateKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeResponseBody) SetRequestId(v string) *UpdateKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeResponseBody) SetKnowledgeId(v int64) *UpdateKnowledgeResponseBody {
	s.KnowledgeId = &v
	return s
}

type UpdateKnowledgeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeResponse) SetBody(v *UpdateKnowledgeResponseBody) *UpdateKnowledgeResponse {
	s.Body = v
	return s
}

type UpdatePerspectiveRequest struct {
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdatePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveRequest) SetPerspectiveId(v string) *UpdatePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

func (s *UpdatePerspectiveRequest) SetName(v string) *UpdatePerspectiveRequest {
	s.Name = &v
	return s
}

type UpdatePerspectiveResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveResponseBody) SetRequestId(v string) *UpdatePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePerspectiveResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveResponse) SetHeaders(v map[string]*string) *UpdatePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *UpdatePerspectiveResponse) SetBody(v *UpdatePerspectiveResponseBody) *UpdatePerspectiveResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("chatbot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActivatePerspectiveWithOptions(request *ActivatePerspectiveRequest, runtime *util.RuntimeOptions) (_result *ActivatePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ActivatePerspectiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ActivatePerspective"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivatePerspective(request *ActivatePerspectiveRequest) (_result *ActivatePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivatePerspectiveResponse{}
	_body, _err := client.ActivatePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSynonymWithOptions(request *AddSynonymRequest, runtime *util.RuntimeOptions) (_result *AddSynonymResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSynonymResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSynonym"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSynonym(request *AddSynonymRequest) (_result *AddSynonymResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSynonymResponse{}
	_body, _err := client.AddSynonymWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppendEntityMemberWithOptions(request *AppendEntityMemberRequest, runtime *util.RuntimeOptions) (_result *AppendEntityMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AppendEntityMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AppendEntityMember"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppendEntityMember(request *AppendEntityMemberRequest) (_result *AppendEntityMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppendEntityMemberResponse{}
	_body, _err := client.AppendEntityMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateWithOptions(request *AssociateRequest, runtime *util.RuntimeOptions) (_result *AssociateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Associate"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Associate(request *AssociateRequest) (_result *AssociateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateResponse{}
	_body, _err := client.AssociateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChatWithOptions(request *ChatRequest, runtime *util.RuntimeOptions) (_result *ChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChatResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Chat"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Chat(request *ChatRequest) (_result *ChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatResponse{}
	_body, _err := client.ChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBotWithOptions(request *CreateBotRequest, runtime *util.RuntimeOptions) (_result *CreateBotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBot"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBot(request *CreateBotRequest) (_result *CreateBotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBotResponse{}
	_body, _err := client.CreateBotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCategoryWithOptions(request *CreateCategoryRequest, runtime *util.RuntimeOptions) (_result *CreateCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCategory"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCategory(request *CreateCategoryRequest) (_result *CreateCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCategoryResponse{}
	_body, _err := client.CreateCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCoreWordWithOptions(request *CreateCoreWordRequest, runtime *util.RuntimeOptions) (_result *CreateCoreWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCoreWordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCoreWord"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCoreWord(request *CreateCoreWordRequest) (_result *CreateCoreWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCoreWordResponse{}
	_body, _err := client.CreateCoreWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDialogWithOptions(request *CreateDialogRequest, runtime *util.RuntimeOptions) (_result *CreateDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDialogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDialog"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDialog(request *CreateDialogRequest) (_result *CreateDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDialogResponse{}
	_body, _err := client.CreateDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEntityWithOptions(request *CreateEntityRequest, runtime *util.RuntimeOptions) (_result *CreateEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEntity"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEntity(request *CreateEntityRequest) (_result *CreateEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEntityResponse{}
	_body, _err := client.CreateEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIntentWithOptions(request *CreateIntentRequest, runtime *util.RuntimeOptions) (_result *CreateIntentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIntentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIntent"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIntent(request *CreateIntentRequest) (_result *CreateIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIntentResponse{}
	_body, _err := client.CreateIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKnowledgeWithOptions(request *CreateKnowledgeRequest, runtime *util.RuntimeOptions) (_result *CreateKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateKnowledgeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateKnowledge"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKnowledge(request *CreateKnowledgeRequest) (_result *CreateKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKnowledgeResponse{}
	_body, _err := client.CreateKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBotWithOptions(request *DeleteBotRequest, runtime *util.RuntimeOptions) (_result *DeleteBotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBot"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBot(request *DeleteBotRequest) (_result *DeleteBotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBotResponse{}
	_body, _err := client.DeleteBotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCategoryWithOptions(request *DeleteCategoryRequest, runtime *util.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCategory"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCategory(request *DeleteCategoryRequest) (_result *DeleteCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DeleteCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCoreWordWithOptions(request *DeleteCoreWordRequest, runtime *util.RuntimeOptions) (_result *DeleteCoreWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCoreWordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCoreWord"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCoreWord(request *DeleteCoreWordRequest) (_result *DeleteCoreWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCoreWordResponse{}
	_body, _err := client.DeleteCoreWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDialogWithOptions(request *DeleteDialogRequest, runtime *util.RuntimeOptions) (_result *DeleteDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDialogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDialog"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDialog(request *DeleteDialogRequest) (_result *DeleteDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDialogResponse{}
	_body, _err := client.DeleteDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEntityWithOptions(request *DeleteEntityRequest, runtime *util.RuntimeOptions) (_result *DeleteEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEntity"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEntity(request *DeleteEntityRequest) (_result *DeleteEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEntityResponse{}
	_body, _err := client.DeleteEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIntentWithOptions(request *DeleteIntentRequest, runtime *util.RuntimeOptions) (_result *DeleteIntentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteIntentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteIntent"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIntent(request *DeleteIntentRequest) (_result *DeleteIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIntentResponse{}
	_body, _err := client.DeleteIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKnowledgeWithOptions(request *DeleteKnowledgeRequest, runtime *util.RuntimeOptions) (_result *DeleteKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteKnowledge"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKnowledge(request *DeleteKnowledgeRequest) (_result *DeleteKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.DeleteKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBotWithOptions(request *DescribeBotRequest, runtime *util.RuntimeOptions) (_result *DescribeBotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBot"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBot(request *DescribeBotRequest) (_result *DescribeBotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBotResponse{}
	_body, _err := client.DescribeBotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCategoryWithOptions(request *DescribeCategoryRequest, runtime *util.RuntimeOptions) (_result *DescribeCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCategory"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCategory(request *DescribeCategoryRequest) (_result *DescribeCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCategoryResponse{}
	_body, _err := client.DescribeCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCoreWordWithOptions(request *DescribeCoreWordRequest, runtime *util.RuntimeOptions) (_result *DescribeCoreWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCoreWordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCoreWord"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCoreWord(request *DescribeCoreWordRequest) (_result *DescribeCoreWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCoreWordResponse{}
	_body, _err := client.DescribeCoreWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDialogWithOptions(request *DescribeDialogRequest, runtime *util.RuntimeOptions) (_result *DescribeDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDialogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDialog"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDialog(request *DescribeDialogRequest) (_result *DescribeDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDialogResponse{}
	_body, _err := client.DescribeDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePerspectiveWithOptions(request *DescribePerspectiveRequest, runtime *util.RuntimeOptions) (_result *DescribePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePerspectiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePerspective"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePerspective(request *DescribePerspectiveRequest) (_result *DescribePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePerspectiveResponse{}
	_body, _err := client.DescribePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableDialogFlowWithOptions(request *DisableDialogFlowRequest, runtime *util.RuntimeOptions) (_result *DisableDialogFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableDialogFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableDialogFlow"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableDialogFlow(request *DisableDialogFlowRequest) (_result *DisableDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDialogFlowResponse{}
	_body, _err := client.DisableDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableKnowledgeWithOptions(request *DisableKnowledgeRequest, runtime *util.RuntimeOptions) (_result *DisableKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableKnowledgeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableKnowledge"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableKnowledge(request *DisableKnowledgeRequest) (_result *DisableKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableKnowledgeResponse{}
	_body, _err := client.DisableKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FeedbackWithOptions(request *FeedbackRequest, runtime *util.RuntimeOptions) (_result *FeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FeedbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Feedback"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Feedback(request *FeedbackRequest) (_result *FeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FeedbackResponse{}
	_body, _err := client.FeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBotChatDataWithOptions(request *GetBotChatDataRequest, runtime *util.RuntimeOptions) (_result *GetBotChatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetBotChatDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBotChatData"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotChatData(request *GetBotChatDataRequest) (_result *GetBotChatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBotChatDataResponse{}
	_body, _err := client.GetBotChatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBotDsStatDataWithOptions(request *GetBotDsStatDataRequest, runtime *util.RuntimeOptions) (_result *GetBotDsStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetBotDsStatDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBotDsStatData"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotDsStatData(request *GetBotDsStatDataRequest) (_result *GetBotDsStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBotDsStatDataResponse{}
	_body, _err := client.GetBotDsStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBotKnowledgeStatDataWithOptions(request *GetBotKnowledgeStatDataRequest, runtime *util.RuntimeOptions) (_result *GetBotKnowledgeStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetBotKnowledgeStatDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBotKnowledgeStatData"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotKnowledgeStatData(request *GetBotKnowledgeStatDataRequest) (_result *GetBotKnowledgeStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBotKnowledgeStatDataResponse{}
	_body, _err := client.GetBotKnowledgeStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBotSessionDataWithOptions(request *GetBotSessionDataRequest, runtime *util.RuntimeOptions) (_result *GetBotSessionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetBotSessionDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBotSessionData"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBotSessionData(request *GetBotSessionDataRequest) (_result *GetBotSessionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBotSessionDataResponse{}
	_body, _err := client.GetBotSessionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConversationListWithOptions(request *GetConversationListRequest, runtime *util.RuntimeOptions) (_result *GetConversationListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConversationListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConversationList"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConversationList(request *GetConversationListRequest) (_result *GetConversationListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConversationListResponse{}
	_body, _err := client.GetConversationListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotChatHistorysWithOptions(request *ListBotChatHistorysRequest, runtime *util.RuntimeOptions) (_result *ListBotChatHistorysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBotChatHistorysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBotChatHistorys"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotChatHistorys(request *ListBotChatHistorysRequest) (_result *ListBotChatHistorysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotChatHistorysResponse{}
	_body, _err := client.ListBotChatHistorysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotColdDsDatasWithOptions(request *ListBotColdDsDatasRequest, runtime *util.RuntimeOptions) (_result *ListBotColdDsDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBotColdDsDatasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBotColdDsDatas"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotColdDsDatas(request *ListBotColdDsDatasRequest) (_result *ListBotColdDsDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotColdDsDatasResponse{}
	_body, _err := client.ListBotColdDsDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotColdKnowledgesWithOptions(request *ListBotColdKnowledgesRequest, runtime *util.RuntimeOptions) (_result *ListBotColdKnowledgesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBotColdKnowledgesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBotColdKnowledges"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotColdKnowledges(request *ListBotColdKnowledgesRequest) (_result *ListBotColdKnowledgesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotColdKnowledgesResponse{}
	_body, _err := client.ListBotColdKnowledgesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotDsDetailsWithOptions(request *ListBotDsDetailsRequest, runtime *util.RuntimeOptions) (_result *ListBotDsDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBotDsDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBotDsDetails"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotDsDetails(request *ListBotDsDetailsRequest) (_result *ListBotDsDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotDsDetailsResponse{}
	_body, _err := client.ListBotDsDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotHotDsDatasWithOptions(request *ListBotHotDsDatasRequest, runtime *util.RuntimeOptions) (_result *ListBotHotDsDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBotHotDsDatasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBotHotDsDatas"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotHotDsDatas(request *ListBotHotDsDatasRequest) (_result *ListBotHotDsDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotHotDsDatasResponse{}
	_body, _err := client.ListBotHotDsDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotHotKnowledgesWithOptions(request *ListBotHotKnowledgesRequest, runtime *util.RuntimeOptions) (_result *ListBotHotKnowledgesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBotHotKnowledgesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBotHotKnowledges"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotHotKnowledges(request *ListBotHotKnowledgesRequest) (_result *ListBotHotKnowledgesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotHotKnowledgesResponse{}
	_body, _err := client.ListBotHotKnowledgesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotKnowledgeDetailsWithOptions(request *ListBotKnowledgeDetailsRequest, runtime *util.RuntimeOptions) (_result *ListBotKnowledgeDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBotKnowledgeDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBotKnowledgeDetails"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotKnowledgeDetails(request *ListBotKnowledgeDetailsRequest) (_result *ListBotKnowledgeDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotKnowledgeDetailsResponse{}
	_body, _err := client.ListBotKnowledgeDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBotReceptionDetailDatasWithOptions(request *ListBotReceptionDetailDatasRequest, runtime *util.RuntimeOptions) (_result *ListBotReceptionDetailDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBotReceptionDetailDatasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBotReceptionDetailDatas"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBotReceptionDetailDatas(request *ListBotReceptionDetailDatasRequest) (_result *ListBotReceptionDetailDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBotReceptionDetailDatasResponse{}
	_body, _err := client.ListBotReceptionDetailDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConversationLogsWithOptions(request *ListConversationLogsRequest, runtime *util.RuntimeOptions) (_result *ListConversationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConversationLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConversationLogs"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConversationLogs(request *ListConversationLogsRequest) (_result *ListConversationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConversationLogsResponse{}
	_body, _err := client.ListConversationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveKnowledgeCategoryWithOptions(request *MoveKnowledgeCategoryRequest, runtime *util.RuntimeOptions) (_result *MoveKnowledgeCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveKnowledgeCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveKnowledgeCategory"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveKnowledgeCategory(request *MoveKnowledgeCategoryRequest) (_result *MoveKnowledgeCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveKnowledgeCategoryResponse{}
	_body, _err := client.MoveKnowledgeCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishDialogFlowWithOptions(request *PublishDialogFlowRequest, runtime *util.RuntimeOptions) (_result *PublishDialogFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishDialogFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishDialogFlow"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishDialogFlow(request *PublishDialogFlowRequest) (_result *PublishDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishDialogFlowResponse{}
	_body, _err := client.PublishDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishKnowledgeWithOptions(request *PublishKnowledgeRequest, runtime *util.RuntimeOptions) (_result *PublishKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishKnowledgeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishKnowledge"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishKnowledge(request *PublishKnowledgeRequest) (_result *PublishKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishKnowledgeResponse{}
	_body, _err := client.PublishKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBotsWithOptions(request *QueryBotsRequest, runtime *util.RuntimeOptions) (_result *QueryBotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryBotsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBots"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBots(request *QueryBotsRequest) (_result *QueryBotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBotsResponse{}
	_body, _err := client.QueryBotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDialogsWithOptions(request *QueryDialogsRequest, runtime *util.RuntimeOptions) (_result *QueryDialogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDialogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDialogs"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDialogs(request *QueryDialogsRequest) (_result *QueryDialogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDialogsResponse{}
	_body, _err := client.QueryDialogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPerspectivesWithOptions(runtime *util.RuntimeOptions) (_result *QueryPerspectivesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &QueryPerspectivesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPerspectives"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPerspectives() (_result *QueryPerspectivesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPerspectivesResponse{}
	_body, _err := client.QueryPerspectivesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySystemEntitiesWithOptions(request *QuerySystemEntitiesRequest, runtime *util.RuntimeOptions) (_result *QuerySystemEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySystemEntitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySystemEntities"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySystemEntities(request *QuerySystemEntitiesRequest) (_result *QuerySystemEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySystemEntitiesResponse{}
	_body, _err := client.QuerySystemEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveEntityMemberWithOptions(request *RemoveEntityMemberRequest, runtime *util.RuntimeOptions) (_result *RemoveEntityMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveEntityMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveEntityMember"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveEntityMember(request *RemoveEntityMemberRequest) (_result *RemoveEntityMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveEntityMemberResponse{}
	_body, _err := client.RemoveEntityMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSynonymWithOptions(request *RemoveSynonymRequest, runtime *util.RuntimeOptions) (_result *RemoveSynonymResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveSynonymResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveSynonym"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSynonym(request *RemoveSynonymRequest) (_result *RemoveSynonymResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSynonymResponse{}
	_body, _err := client.RemoveSynonymWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TestDialogFlowWithOptions(request *TestDialogFlowRequest, runtime *util.RuntimeOptions) (_result *TestDialogFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TestDialogFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TestDialogFlow"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TestDialogFlow(request *TestDialogFlowRequest) (_result *TestDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestDialogFlowResponse{}
	_body, _err := client.TestDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCategoryWithOptions(request *UpdateCategoryRequest, runtime *util.RuntimeOptions) (_result *UpdateCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCategory"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCategory(request *UpdateCategoryRequest) (_result *UpdateCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.UpdateCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCoreWordWithOptions(request *UpdateCoreWordRequest, runtime *util.RuntimeOptions) (_result *UpdateCoreWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCoreWordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCoreWord"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCoreWord(request *UpdateCoreWordRequest) (_result *UpdateCoreWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCoreWordResponse{}
	_body, _err := client.UpdateCoreWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDialogWithOptions(request *UpdateDialogRequest, runtime *util.RuntimeOptions) (_result *UpdateDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDialogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDialog"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDialog(request *UpdateDialogRequest) (_result *UpdateDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDialogResponse{}
	_body, _err := client.UpdateDialogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDialogFlowWithOptions(request *UpdateDialogFlowRequest, runtime *util.RuntimeOptions) (_result *UpdateDialogFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDialogFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDialogFlow"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDialogFlow(request *UpdateDialogFlowRequest) (_result *UpdateDialogFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDialogFlowResponse{}
	_body, _err := client.UpdateDialogFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEntityWithOptions(request *UpdateEntityRequest, runtime *util.RuntimeOptions) (_result *UpdateEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEntity"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEntity(request *UpdateEntityRequest) (_result *UpdateEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEntityResponse{}
	_body, _err := client.UpdateEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIntentWithOptions(request *UpdateIntentRequest, runtime *util.RuntimeOptions) (_result *UpdateIntentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateIntentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateIntent"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIntent(request *UpdateIntentRequest) (_result *UpdateIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIntentResponse{}
	_body, _err := client.UpdateIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKnowledgeWithOptions(request *UpdateKnowledgeRequest, runtime *util.RuntimeOptions) (_result *UpdateKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateKnowledgeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateKnowledge"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKnowledge(request *UpdateKnowledgeRequest) (_result *UpdateKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateKnowledgeResponse{}
	_body, _err := client.UpdateKnowledgeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePerspectiveWithOptions(request *UpdatePerspectiveRequest, runtime *util.RuntimeOptions) (_result *UpdatePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePerspectiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePerspective"), tea.String("2017-10-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePerspective(request *UpdatePerspectiveRequest) (_result *UpdatePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePerspectiveResponse{}
	_body, _err := client.UpdatePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
