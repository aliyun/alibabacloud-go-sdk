// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BeeBotAssociateRequest struct {
	// The ID of chatbot instance.
	ChatBotInstanceId *string `json:"ChatBotInstanceId,omitempty" xml:"ChatBotInstanceId,omitempty"`
	CustSpaceId       *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// ISV verification code, which is used to verify whether the sub-account is authorized by ISV.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The list of codes for answers from different perspectives.
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// The number of recommended questions, which ranges from 1 to 10.
	RecommendNum *int32 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// The ID of the session, which identifies the session and context information of the visitor.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input of the visitor.
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s BeeBotAssociateRequest) String() string {
	return tea.Prettify(s)
}

func (s BeeBotAssociateRequest) GoString() string {
	return s.String()
}

func (s *BeeBotAssociateRequest) SetChatBotInstanceId(v string) *BeeBotAssociateRequest {
	s.ChatBotInstanceId = &v
	return s
}

func (s *BeeBotAssociateRequest) SetCustSpaceId(v string) *BeeBotAssociateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *BeeBotAssociateRequest) SetIsvCode(v string) *BeeBotAssociateRequest {
	s.IsvCode = &v
	return s
}

func (s *BeeBotAssociateRequest) SetPerspective(v []*string) *BeeBotAssociateRequest {
	s.Perspective = v
	return s
}

func (s *BeeBotAssociateRequest) SetRecommendNum(v int32) *BeeBotAssociateRequest {
	s.RecommendNum = &v
	return s
}

func (s *BeeBotAssociateRequest) SetSessionId(v string) *BeeBotAssociateRequest {
	s.SessionId = &v
	return s
}

func (s *BeeBotAssociateRequest) SetUtterance(v string) *BeeBotAssociateRequest {
	s.Utterance = &v
	return s
}

type BeeBotAssociateShrinkRequest struct {
	// The ID of chatbot instance.
	ChatBotInstanceId *string `json:"ChatBotInstanceId,omitempty" xml:"ChatBotInstanceId,omitempty"`
	CustSpaceId       *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// ISV verification code, which is used to verify whether the sub-account is authorized by ISV.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The list of codes for answers from different perspectives.
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
	// The number of recommended questions, which ranges from 1 to 10.
	RecommendNum *int32 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// The ID of the session, which identifies the session and context information of the visitor.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input of the visitor.
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s BeeBotAssociateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BeeBotAssociateShrinkRequest) GoString() string {
	return s.String()
}

func (s *BeeBotAssociateShrinkRequest) SetChatBotInstanceId(v string) *BeeBotAssociateShrinkRequest {
	s.ChatBotInstanceId = &v
	return s
}

func (s *BeeBotAssociateShrinkRequest) SetCustSpaceId(v string) *BeeBotAssociateShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *BeeBotAssociateShrinkRequest) SetIsvCode(v string) *BeeBotAssociateShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *BeeBotAssociateShrinkRequest) SetPerspectiveShrink(v string) *BeeBotAssociateShrinkRequest {
	s.PerspectiveShrink = &v
	return s
}

func (s *BeeBotAssociateShrinkRequest) SetRecommendNum(v int32) *BeeBotAssociateShrinkRequest {
	s.RecommendNum = &v
	return s
}

func (s *BeeBotAssociateShrinkRequest) SetSessionId(v string) *BeeBotAssociateShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *BeeBotAssociateShrinkRequest) SetUtterance(v string) *BeeBotAssociateShrinkRequest {
	s.Utterance = &v
	return s
}

type BeeBotAssociateResponseBody struct {
	// If OK is returned, the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *BeeBotAssociateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BeeBotAssociateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeeBotAssociateResponseBody) GoString() string {
	return s.String()
}

func (s *BeeBotAssociateResponseBody) SetCode(v string) *BeeBotAssociateResponseBody {
	s.Code = &v
	return s
}

func (s *BeeBotAssociateResponseBody) SetData(v *BeeBotAssociateResponseBodyData) *BeeBotAssociateResponseBody {
	s.Data = v
	return s
}

func (s *BeeBotAssociateResponseBody) SetMessage(v string) *BeeBotAssociateResponseBody {
	s.Message = &v
	return s
}

func (s *BeeBotAssociateResponseBody) SetRequestId(v string) *BeeBotAssociateResponseBody {
	s.RequestId = &v
	return s
}

type BeeBotAssociateResponseBodyData struct {
	// The list of associated recommendations.
	Associate []*BeeBotAssociateResponseBodyDataAssociate `json:"Associate,omitempty" xml:"Associate,omitempty" type:"Repeated"`
	// The ID of the response message.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The ID of the session.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s BeeBotAssociateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BeeBotAssociateResponseBodyData) GoString() string {
	return s.String()
}

func (s *BeeBotAssociateResponseBodyData) SetAssociate(v []*BeeBotAssociateResponseBodyDataAssociate) *BeeBotAssociateResponseBodyData {
	s.Associate = v
	return s
}

func (s *BeeBotAssociateResponseBodyData) SetMessageId(v string) *BeeBotAssociateResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *BeeBotAssociateResponseBodyData) SetSessionId(v string) *BeeBotAssociateResponseBodyData {
	s.SessionId = &v
	return s
}

type BeeBotAssociateResponseBodyDataAssociate struct {
	// The additional information.
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The title of the associated question.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s BeeBotAssociateResponseBodyDataAssociate) String() string {
	return tea.Prettify(s)
}

func (s BeeBotAssociateResponseBodyDataAssociate) GoString() string {
	return s.String()
}

func (s *BeeBotAssociateResponseBodyDataAssociate) SetMeta(v string) *BeeBotAssociateResponseBodyDataAssociate {
	s.Meta = &v
	return s
}

func (s *BeeBotAssociateResponseBodyDataAssociate) SetTitle(v string) *BeeBotAssociateResponseBodyDataAssociate {
	s.Title = &v
	return s
}

type BeeBotAssociateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BeeBotAssociateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BeeBotAssociateResponse) String() string {
	return tea.Prettify(s)
}

func (s BeeBotAssociateResponse) GoString() string {
	return s.String()
}

func (s *BeeBotAssociateResponse) SetHeaders(v map[string]*string) *BeeBotAssociateResponse {
	s.Headers = v
	return s
}

func (s *BeeBotAssociateResponse) SetStatusCode(v int32) *BeeBotAssociateResponse {
	s.StatusCode = &v
	return s
}

func (s *BeeBotAssociateResponse) SetBody(v *BeeBotAssociateResponseBody) *BeeBotAssociateResponse {
	s.Body = v
	return s
}

type BeeBotChatRequest struct {
	// The ID of chatbot instance.
	ChatBotInstanceId *string `json:"ChatBotInstanceId,omitempty" xml:"ChatBotInstanceId,omitempty"`
	CustSpaceId       *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the intent in the dialog flow. When this parameter is specified, the bot will conduct a Q\&A according to the intent.
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// ISV verification code, which is used to verify whether the sub-account is authorized by ISV.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The ID of the knowledge title in the knowledge base.
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The list of codes for answers from different perspectives.
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// The ID of the visitor, which is used to identify users in the current session.
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The nickname of the visitor in the current session.
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// The ID of the session, which identifies the session and context information of the visitor.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input of the visitor.
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	// The user-defined parameter set in JSON format. You can specify user-defined parameters for conversation engines.
	VendorParam map[string]interface{} `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
}

func (s BeeBotChatRequest) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatRequest) GoString() string {
	return s.String()
}

func (s *BeeBotChatRequest) SetChatBotInstanceId(v string) *BeeBotChatRequest {
	s.ChatBotInstanceId = &v
	return s
}

func (s *BeeBotChatRequest) SetCustSpaceId(v string) *BeeBotChatRequest {
	s.CustSpaceId = &v
	return s
}

func (s *BeeBotChatRequest) SetIntentName(v string) *BeeBotChatRequest {
	s.IntentName = &v
	return s
}

func (s *BeeBotChatRequest) SetIsvCode(v string) *BeeBotChatRequest {
	s.IsvCode = &v
	return s
}

func (s *BeeBotChatRequest) SetKnowledgeId(v string) *BeeBotChatRequest {
	s.KnowledgeId = &v
	return s
}

func (s *BeeBotChatRequest) SetPerspective(v []*string) *BeeBotChatRequest {
	s.Perspective = v
	return s
}

func (s *BeeBotChatRequest) SetSenderId(v string) *BeeBotChatRequest {
	s.SenderId = &v
	return s
}

func (s *BeeBotChatRequest) SetSenderNick(v string) *BeeBotChatRequest {
	s.SenderNick = &v
	return s
}

func (s *BeeBotChatRequest) SetSessionId(v string) *BeeBotChatRequest {
	s.SessionId = &v
	return s
}

func (s *BeeBotChatRequest) SetUtterance(v string) *BeeBotChatRequest {
	s.Utterance = &v
	return s
}

func (s *BeeBotChatRequest) SetVendorParam(v map[string]interface{}) *BeeBotChatRequest {
	s.VendorParam = v
	return s
}

type BeeBotChatShrinkRequest struct {
	// The ID of chatbot instance.
	ChatBotInstanceId *string `json:"ChatBotInstanceId,omitempty" xml:"ChatBotInstanceId,omitempty"`
	CustSpaceId       *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the intent in the dialog flow. When this parameter is specified, the bot will conduct a Q\&A according to the intent.
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// ISV verification code, which is used to verify whether the sub-account is authorized by ISV.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The ID of the knowledge title in the knowledge base.
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The list of codes for answers from different perspectives.
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
	// The ID of the visitor, which is used to identify users in the current session.
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The nickname of the visitor in the current session.
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// The ID of the session, which identifies the session and context information of the visitor.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input of the visitor.
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	// The user-defined parameter set in JSON format. You can specify user-defined parameters for conversation engines.
	VendorParamShrink *string `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
}

func (s BeeBotChatShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *BeeBotChatShrinkRequest) SetChatBotInstanceId(v string) *BeeBotChatShrinkRequest {
	s.ChatBotInstanceId = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetCustSpaceId(v string) *BeeBotChatShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetIntentName(v string) *BeeBotChatShrinkRequest {
	s.IntentName = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetIsvCode(v string) *BeeBotChatShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetKnowledgeId(v string) *BeeBotChatShrinkRequest {
	s.KnowledgeId = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetPerspectiveShrink(v string) *BeeBotChatShrinkRequest {
	s.PerspectiveShrink = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetSenderId(v string) *BeeBotChatShrinkRequest {
	s.SenderId = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetSenderNick(v string) *BeeBotChatShrinkRequest {
	s.SenderNick = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetSessionId(v string) *BeeBotChatShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetUtterance(v string) *BeeBotChatShrinkRequest {
	s.Utterance = &v
	return s
}

func (s *BeeBotChatShrinkRequest) SetVendorParamShrink(v string) *BeeBotChatShrinkRequest {
	s.VendorParamShrink = &v
	return s
}

type BeeBotChatResponseBody struct {
	// If OK is returned, the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *BeeBotChatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BeeBotChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBody) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBody) SetCode(v string) *BeeBotChatResponseBody {
	s.Code = &v
	return s
}

func (s *BeeBotChatResponseBody) SetData(v *BeeBotChatResponseBodyData) *BeeBotChatResponseBody {
	s.Data = v
	return s
}

func (s *BeeBotChatResponseBody) SetMessage(v string) *BeeBotChatResponseBody {
	s.Message = &v
	return s
}

func (s *BeeBotChatResponseBody) SetRequestId(v string) *BeeBotChatResponseBody {
	s.RequestId = &v
	return s
}

type BeeBotChatResponseBodyData struct {
	// The ID of the response message.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The list of the message.
	Messages []*BeeBotChatResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The ID of the session.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s BeeBotChatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBodyData) SetMessageId(v string) *BeeBotChatResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *BeeBotChatResponseBodyData) SetMessages(v []*BeeBotChatResponseBodyDataMessages) *BeeBotChatResponseBodyData {
	s.Messages = v
	return s
}

func (s *BeeBotChatResponseBodyData) SetSessionId(v string) *BeeBotChatResponseBodyData {
	s.SessionId = &v
	return s
}

type BeeBotChatResponseBodyDataMessages struct {
	// If the AnswerType parameter is set to Recommend, this parameter indicates the source of the recommended answer.
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The type of the message.
	AnswerType *string `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	// When the AnswerType parameter is set to Knowledge, this parameter contains the Knowledge object returned by the bot.
	Knowledge *BeeBotChatResponseBodyDataMessagesKnowledge `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
	// When the AnswerType parameter is set to Recommend, this parameter contains a list of recommends returned by the bot.
	Recommends []*BeeBotChatResponseBodyDataMessagesRecommends `json:"Recommends,omitempty" xml:"Recommends,omitempty" type:"Repeated"`
	// When the AnswerType parameter is set to Text, this parameter contains the Text object returned by the bot.
	Text *BeeBotChatResponseBodyDataMessagesText `json:"Text,omitempty" xml:"Text,omitempty" type:"Struct"`
}

func (s BeeBotChatResponseBodyDataMessages) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBodyDataMessages) SetAnswerSource(v string) *BeeBotChatResponseBodyDataMessages {
	s.AnswerSource = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessages) SetAnswerType(v string) *BeeBotChatResponseBodyDataMessages {
	s.AnswerType = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessages) SetKnowledge(v *BeeBotChatResponseBodyDataMessagesKnowledge) *BeeBotChatResponseBodyDataMessages {
	s.Knowledge = v
	return s
}

func (s *BeeBotChatResponseBodyDataMessages) SetRecommends(v []*BeeBotChatResponseBodyDataMessagesRecommends) *BeeBotChatResponseBodyDataMessages {
	s.Recommends = v
	return s
}

func (s *BeeBotChatResponseBodyDataMessages) SetText(v *BeeBotChatResponseBodyDataMessagesText) *BeeBotChatResponseBodyDataMessages {
	s.Text = v
	return s
}

type BeeBotChatResponseBodyDataMessagesKnowledge struct {
	// The source of the answer.
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The category of the knowledge.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the hit question.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether the answer is in plain text or rich text.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The hit text.
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// The ID of the hit problem in the knowledge base.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of the related knowledge.
	RelatedKnowledges []*BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges `json:"RelatedKnowledges,omitempty" xml:"RelatedKnowledges,omitempty" type:"Repeated"`
	// The summary to the hit question.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The title of the hit question.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s BeeBotChatResponseBodyDataMessagesKnowledge) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBodyDataMessagesKnowledge) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetAnswerSource(v string) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.AnswerSource = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetCategory(v string) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.Category = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetContent(v string) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.Content = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetContentType(v string) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.ContentType = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetHitStatement(v string) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.HitStatement = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetId(v string) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.Id = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetRelatedKnowledges(v []*BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.RelatedKnowledges = v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetSummary(v string) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.Summary = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledge) SetTitle(v string) *BeeBotChatResponseBodyDataMessagesKnowledge {
	s.Title = &v
	return s
}

type BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges struct {
	// The ID of the related knowledge.
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The title of the related knowledge.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges) SetKnowledgeId(v string) *BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges {
	s.KnowledgeId = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges) SetTitle(v string) *BeeBotChatResponseBodyDataMessagesKnowledgeRelatedKnowledges {
	s.Title = &v
	return s
}

type BeeBotChatResponseBodyDataMessagesRecommends struct {
	// The source of the recommended answer.
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The ID of the recommended knowledge.
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The recommended content, which may be the entity in graph-based question answering, the standard knowledge in knowledge-based question answering, or the column value in table-based question answering.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s BeeBotChatResponseBodyDataMessagesRecommends) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBodyDataMessagesRecommends) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBodyDataMessagesRecommends) SetAnswerSource(v string) *BeeBotChatResponseBodyDataMessagesRecommends {
	s.AnswerSource = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesRecommends) SetKnowledgeId(v string) *BeeBotChatResponseBodyDataMessagesRecommends {
	s.KnowledgeId = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesRecommends) SetTitle(v string) *BeeBotChatResponseBodyDataMessagesRecommends {
	s.Title = &v
	return s
}

type BeeBotChatResponseBodyDataMessagesText struct {
	// The source of the answer.
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The content of the text message.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether the answer is in plain text or rich text.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	DialogName *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	// The passthrough parameters are returned.
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	ExternalFlags map[string]interface{} `json:"ExternalFlags,omitempty" xml:"ExternalFlags,omitempty"`
	// The hit text.
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// The metadata.
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The list of slots.
	Slots []*BeeBotChatResponseBodyDataMessagesTextSlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
	// The title of the chitchat.
	UserDefinedChatTitle *string `json:"UserDefinedChatTitle,omitempty" xml:"UserDefinedChatTitle,omitempty"`
}

func (s BeeBotChatResponseBodyDataMessagesText) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBodyDataMessagesText) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetAnswerSource(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.AnswerSource = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetContent(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.Content = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetContentType(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.ContentType = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetDialogName(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.DialogName = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetExt(v map[string]interface{}) *BeeBotChatResponseBodyDataMessagesText {
	s.Ext = v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetExternalFlags(v map[string]interface{}) *BeeBotChatResponseBodyDataMessagesText {
	s.ExternalFlags = v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetHitStatement(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.HitStatement = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetIntentName(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.IntentName = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetMetaData(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.MetaData = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetNodeId(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.NodeId = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetNodeName(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.NodeName = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetSlots(v []*BeeBotChatResponseBodyDataMessagesTextSlots) *BeeBotChatResponseBodyDataMessagesText {
	s.Slots = v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesText) SetUserDefinedChatTitle(v string) *BeeBotChatResponseBodyDataMessagesText {
	s.UserDefinedChatTitle = &v
	return s
}

type BeeBotChatResponseBodyDataMessagesTextSlots struct {
	// Indicates whether the slot is hit.
	Hit *bool `json:"Hit,omitempty" xml:"Hit,omitempty"`
	// The name of the slot.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The original value.
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The specific value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s BeeBotChatResponseBodyDataMessagesTextSlots) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBodyDataMessagesTextSlots) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBodyDataMessagesTextSlots) SetHit(v bool) *BeeBotChatResponseBodyDataMessagesTextSlots {
	s.Hit = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesTextSlots) SetName(v string) *BeeBotChatResponseBodyDataMessagesTextSlots {
	s.Name = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesTextSlots) SetOrigin(v string) *BeeBotChatResponseBodyDataMessagesTextSlots {
	s.Origin = &v
	return s
}

func (s *BeeBotChatResponseBodyDataMessagesTextSlots) SetValue(v string) *BeeBotChatResponseBodyDataMessagesTextSlots {
	s.Value = &v
	return s
}

type BeeBotChatResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BeeBotChatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BeeBotChatResponse) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponse) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponse) SetHeaders(v map[string]*string) *BeeBotChatResponse {
	s.Headers = v
	return s
}

func (s *BeeBotChatResponse) SetStatusCode(v int32) *BeeBotChatResponse {
	s.StatusCode = &v
	return s
}

func (s *BeeBotChatResponse) SetBody(v *BeeBotChatResponseBody) *BeeBotChatResponse {
	s.Body = v
	return s
}

type CreateChatappTemplateRequest struct {
	// The category of the message template. Valid values:
	//
	// *   **ACCOUNT_UPDATE**: account update
	// *   **PAYMENT_UPDATE**: payment update
	// *   **PERSONAL_FINANCE\_UPDATE**: personal finance update
	// *   **SHIPPING_UPDATE**: traffic update
	// *   **RESERVATION_UPDATE**: reservation update
	// *   **ISSUE_RESOLUTION**: issue resolution
	// *   **APPOINTMENT_UPDATE**: appointment update
	// *   **TRANSPORTATION_UPDATE**: logistics information update
	// *   **TICKET_UPDATE**: ticket update
	// *   **ALERT_UPDATE**: alert update
	// *   **AUTO_REPLY**: auto reply
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The components of the message template.
	Components  []*CreateChatappTemplateRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	CustSpaceId *string                                   `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the message template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the message template. Valid values:
	//
	// *   **WHATSAPP**
	// *   VIBER (under development)
	// *   LINE (under development)
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequest) SetCategory(v string) *CreateChatappTemplateRequest {
	s.Category = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetComponents(v []*CreateChatappTemplateRequestComponents) *CreateChatappTemplateRequest {
	s.Components = v
	return s
}

func (s *CreateChatappTemplateRequest) SetCustSpaceId(v string) *CreateChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetCustWabaId(v string) *CreateChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetExample(v map[string]*string) *CreateChatappTemplateRequest {
	s.Example = v
	return s
}

func (s *CreateChatappTemplateRequest) SetIsvCode(v string) *CreateChatappTemplateRequest {
	s.IsvCode = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetLanguage(v string) *CreateChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetName(v string) *CreateChatappTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetTemplateType(v string) *CreateChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

type CreateChatappTemplateRequestComponents struct {
	// This parameter applies only to components of the **BUTTONS** type.
	Buttons []*CreateChatappTemplateRequestComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description of the file.
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The name of the file.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The format of the message.
	//
	// *   **TEXT**: text
	// *   **IMGAGE**: image
	// *   **DOCUMENT**: document
	// *   **VIDEO**: video
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The text of the message to be sent.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the component. Valid values:
	//
	// *   **BODY**
	// *   **HEADER**
	// *   **FOOTER**
	// *   **BUTTONS**
	//
	// > A component of the **BODY** type cannot exceed 1,024 characters in length. A component of the **HEADER** or **FOOTER** type cannot exceed 60 characters in length.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the material.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateChatappTemplateRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponents) SetButtons(v []*CreateChatappTemplateRequestComponentsButtons) *CreateChatappTemplateRequestComponents {
	s.Buttons = v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetCaption(v string) *CreateChatappTemplateRequestComponents {
	s.Caption = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFileName(v string) *CreateChatappTemplateRequestComponents {
	s.FileName = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFormat(v string) *CreateChatappTemplateRequestComponents {
	s.Format = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetText(v string) *CreateChatappTemplateRequestComponents {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetType(v string) *CreateChatappTemplateRequestComponents {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetUrl(v string) *CreateChatappTemplateRequestComponents {
	s.Url = &v
	return s
}

type CreateChatappTemplateRequestComponentsButtons struct {
	// The mobile phone number. This parameter is valid only if the Type parameter is set to **PHONE_NUMBER**.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The display name of the button.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the button. Valid values:
	//
	// *   **PHONE_NUMBER**: a phone call button
	// *   **URL**: a URL button
	// *   **QUICK_REPLY**: a quick reply button
	//
	// > *   A quick reply button cannot coexist with a phone call button or a URL button in a message template.
	// > *   You can add a combination of two URL buttons or a combination of a URL button and a phone call button to a message template.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to be accessed when you click the URL button.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the URL. Valid values:
	//
	// *   **static**: a static URL
	// *   **dynamic**: a dynamic URL
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsButtons) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetPhoneNumber(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetText(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetType(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetUrl(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Url = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetUrlType(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.UrlType = &v
	return s
}

type CreateChatappTemplateShrinkRequest struct {
	// The category of the message template. Valid values:
	//
	// *   **ACCOUNT_UPDATE**: account update
	// *   **PAYMENT_UPDATE**: payment update
	// *   **PERSONAL_FINANCE\_UPDATE**: personal finance update
	// *   **SHIPPING_UPDATE**: traffic update
	// *   **RESERVATION_UPDATE**: reservation update
	// *   **ISSUE_RESOLUTION**: issue resolution
	// *   **APPOINTMENT_UPDATE**: appointment update
	// *   **TRANSPORTATION_UPDATE**: logistics information update
	// *   **TICKET_UPDATE**: ticket update
	// *   **ALERT_UPDATE**: alert update
	// *   **AUTO_REPLY**: auto reply
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The components of the message template.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	CustSpaceId      *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	ExampleShrink *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the message template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the message template. Valid values:
	//
	// *   **WHATSAPP**
	// *   VIBER (under development)
	// *   LINE (under development)
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateChatappTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateShrinkRequest) SetCategory(v string) *CreateChatappTemplateShrinkRequest {
	s.Category = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetComponentsShrink(v string) *CreateChatappTemplateShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetCustSpaceId(v string) *CreateChatappTemplateShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetCustWabaId(v string) *CreateChatappTemplateShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetExampleShrink(v string) *CreateChatappTemplateShrinkRequest {
	s.ExampleShrink = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetIsvCode(v string) *CreateChatappTemplateShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetLanguage(v string) *CreateChatappTemplateShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetName(v string) *CreateChatappTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetTemplateType(v string) *CreateChatappTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

type CreateChatappTemplateResponseBody struct {
	// The HTTP status code returned.
	//
	// *   A code of OK indicates that the call is successful.
	// *   Other codes indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CreateChatappTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponseBody) SetCode(v string) *CreateChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetData(v *CreateChatappTemplateResponseBodyData) *CreateChatappTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetMessage(v string) *CreateChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetRequestId(v string) *CreateChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateChatappTemplateResponseBodyData struct {
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateChatappTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponseBodyData) SetTemplateCode(v string) *CreateChatappTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *CreateChatappTemplateResponseBodyData) SetTemplateName(v string) *CreateChatappTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

type CreateChatappTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateChatappTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponse) SetHeaders(v map[string]*string) *CreateChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateChatappTemplateResponse) SetStatusCode(v int32) *CreateChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatappTemplateResponse) SetBody(v *CreateChatappTemplateResponseBody) *CreateChatappTemplateResponse {
	s.Body = v
	return s
}

type DeleteChatappTemplateRequest struct {
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s DeleteChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateRequest) SetCustSpaceId(v string) *DeleteChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetCustWabaId(v string) *DeleteChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetIsvCode(v string) *DeleteChatappTemplateRequest {
	s.IsvCode = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetTemplateCode(v string) *DeleteChatappTemplateRequest {
	s.TemplateCode = &v
	return s
}

type DeleteChatappTemplateResponseBody struct {
	// The HTTP status code returned.
	//
	// *   A code of OK indicates that the call is successful.
	// *   Other codes indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateResponseBody) SetCode(v string) *DeleteChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) SetMessage(v string) *DeleteChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) SetRequestId(v string) *DeleteChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChatappTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteChatappTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateResponse) SetHeaders(v map[string]*string) *DeleteChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatappTemplateResponse) SetStatusCode(v int32) *DeleteChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatappTemplateResponse) SetBody(v *DeleteChatappTemplateResponseBody) *DeleteChatappTemplateResponse {
	s.Body = v
	return s
}

type GetChatappTemplateDetailRequest struct {
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetChatappTemplateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailRequest) SetCustSpaceId(v string) *GetChatappTemplateDetailRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetCustWabaId(v string) *GetChatappTemplateDetailRequest {
	s.CustWabaId = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetIsvCode(v string) *GetChatappTemplateDetailRequest {
	s.IsvCode = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetLanguage(v string) *GetChatappTemplateDetailRequest {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetTemplateCode(v string) *GetChatappTemplateDetailRequest {
	s.TemplateCode = &v
	return s
}

type GetChatappTemplateDetailResponseBody struct {
	// The HTTP status code returned.
	//
	// *   A code of OK indicates that the call is successful.
	// *   Other codes indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetChatappTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappTemplateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBody) SetCode(v string) *GetChatappTemplateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetData(v *GetChatappTemplateDetailResponseBodyData) *GetChatappTemplateDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetMessage(v string) *GetChatappTemplateDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetRequestId(v string) *GetChatappTemplateDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetChatappTemplateDetailResponseBodyData struct {
	// The review status of the message template. Valid values:
	//
	// *   **pass**: The message template is approved.
	// *   **fail**: The message template is rejected.
	// *   **auditing**: The message template is being reviewed.
	// *   **unaudit**: The review is suspended.
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The category of the message template. Valid values:
	//
	// *   **ACCOUNT_UPDATE**: account update
	// *   **PAYMENT_UPDATE**: payment update
	// *   **PERSONAL_FINANCE\_UPDATE**: personal finance update
	// *   **SHIPPING_UPDATE**: traffic update
	// *   **RESERVATION_UPDATE**: reservation update
	// *   **ISSUE_RESOLUTION**: issue resolution
	// *   **APPOINTMENT_UPDATE**: appointment update
	// *   **TRANSPORTATION_UPDATE**: logistics information update
	// *   **TICKET_UPDATE**: ticket update
	// *   **ALERT_UPDATE**: alert update
	// *   **AUTO_REPLY**: auto reply
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The components of the message template.
	Components []*GetChatappTemplateDetailResponseBodyDataComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The examples of variables.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the message template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyData) SetAuditStatus(v string) *GetChatappTemplateDetailResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetCategory(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetComponents(v []*GetChatappTemplateDetailResponseBodyDataComponents) *GetChatappTemplateDetailResponseBodyData {
	s.Components = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetExample(v map[string]*string) *GetChatappTemplateDetailResponseBodyData {
	s.Example = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetLanguage(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetName(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetTemplateCode(v string) *GetChatappTemplateDetailResponseBodyData {
	s.TemplateCode = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponents struct {
	// This parameter applies only to components of the **BUTTONS** type. This parameter is passed in by converting its original JSON structure into a string.
	Buttons []*GetChatappTemplateDetailResponseBodyDataComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description of the file.
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The name of the file.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The format.
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The text of the message to be sent.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the component. Valid values:
	//
	// *   **BODY**
	// *   **HEADER**
	// *   **FOOTER**
	// *   **BUTTONS**
	//
	// **
	//
	// **Note** A component of the **BODY** type cannot exceed 1,024 characters in length. A component of the **HEADER** or **FOOTER** type cannot exceed 60 characters in length.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the material.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponents) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponents) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetButtons(v []*GetChatappTemplateDetailResponseBodyDataComponentsButtons) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Buttons = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetCaption(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Caption = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFileName(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.FileName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFormat(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Format = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Url = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponentsButtons struct {
	// The mobile phone number. This parameter is valid only if the Type parameter is set to **PHONE_NUMBER**.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The display name of the button.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the button. Valid values:
	//
	// *   **PHONE_NUMBER**: a phone call button
	// *   **URL**: a URL button
	// *   **QUICK_REPLY**: a quick reply button
	//
	// **
	//
	// **Note**
	//
	// *   A quick reply button cannot coexist with a phone call button or a URL button in a message template.
	//
	// *   You can add a combination of two URL buttons or a combination of a URL button and a phone call button to a message template.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to be accessed when you click the URL button.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the URL. Valid values:
	//
	// *   **static**: a static URL
	// *   **dynamic**: a dynamic URL
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtons) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetPhoneNumber(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Url = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetUrlType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.UrlType = &v
	return s
}

type GetChatappTemplateDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetChatappTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChatappTemplateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponse) SetHeaders(v map[string]*string) *GetChatappTemplateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetChatappTemplateDetailResponse) SetStatusCode(v int32) *GetChatappTemplateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponse) SetBody(v *GetChatappTemplateDetailResponseBody) *GetChatappTemplateDetailResponse {
	s.Body = v
	return s
}

type ListChatappTemplateRequest struct {
	// The review status of the message template. Valid values:
	//
	// *   **pass**: The message template is approved.
	// *   **fail**: The message template is rejected.
	// *   **auditing**: The message template is being reviewed.
	// *   **unaudit**: The review is suspended.
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the message template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The paging settings.
	Page *ListChatappTemplateRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateRequest) SetAuditStatus(v string) *ListChatappTemplateRequest {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateRequest) SetCustSpaceId(v string) *ListChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatappTemplateRequest) SetCustWabaId(v string) *ListChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *ListChatappTemplateRequest) SetIsvCode(v string) *ListChatappTemplateRequest {
	s.IsvCode = &v
	return s
}

func (s *ListChatappTemplateRequest) SetLanguage(v string) *ListChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateRequest) SetName(v string) *ListChatappTemplateRequest {
	s.Name = &v
	return s
}

func (s *ListChatappTemplateRequest) SetPage(v *ListChatappTemplateRequestPage) *ListChatappTemplateRequest {
	s.Page = v
	return s
}

type ListChatappTemplateRequestPage struct {
	// The number of the page to return. Default value: 1.
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The number of message templates to return on each page. Default value: 10.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListChatappTemplateRequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateRequestPage) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateRequestPage) SetIndex(v int32) *ListChatappTemplateRequestPage {
	s.Index = &v
	return s
}

func (s *ListChatappTemplateRequestPage) SetSize(v int32) *ListChatappTemplateRequestPage {
	s.Size = &v
	return s
}

type ListChatappTemplateShrinkRequest struct {
	// The review status of the message template. Valid values:
	//
	// *   **pass**: The message template is approved.
	// *   **fail**: The message template is rejected.
	// *   **auditing**: The message template is being reviewed.
	// *   **unaudit**: The review is suspended.
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the message template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The paging settings.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s ListChatappTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateShrinkRequest) SetAuditStatus(v string) *ListChatappTemplateShrinkRequest {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetCustSpaceId(v string) *ListChatappTemplateShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetCustWabaId(v string) *ListChatappTemplateShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetIsvCode(v string) *ListChatappTemplateShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetLanguage(v string) *ListChatappTemplateShrinkRequest {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetName(v string) *ListChatappTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetPageShrink(v string) *ListChatappTemplateShrinkRequest {
	s.PageShrink = &v
	return s
}

type ListChatappTemplateResponseBody struct {
	// The HTTP status code returned.
	//
	// *   A code of OK indicates that the call is successful.
	// *   Other codes indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message templates.
	ListTemplate []*ListChatappTemplateResponseBodyListTemplate `json:"ListTemplate,omitempty" xml:"ListTemplate,omitempty" type:"Repeated"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponseBody) SetCode(v string) *ListChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetListTemplate(v []*ListChatappTemplateResponseBodyListTemplate) *ListChatappTemplateResponseBody {
	s.ListTemplate = v
	return s
}

func (s *ListChatappTemplateResponseBody) SetMessage(v string) *ListChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetRequestId(v string) *ListChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ListChatappTemplateResponseBodyListTemplate struct {
	// The review status of the message template. Valid values:
	//
	// *   **pass**: The message template is approved.
	// *   **fail**: The message template is rejected.
	// *   **auditing**: The message template is being reviewed.
	// *   **unaudit**: The review is suspended.
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The category of the message template. Valid values:
	//
	// *   **ACCOUNT_UPDATE**: account update
	// *   **PAYMENT_UPDATE**: payment update
	// *   **PERSONAL_FINANCE\_UPDATE**: personal finance update
	// *   **SHIPPING_UPDATE**: traffic update
	// *   **RESERVATION_UPDATE**: reservation update
	// *   **ISSUE_RESOLUTION**: issue resolution
	// *   **APPOINTMENT_UPDATE**: appointment update
	// *   **TRANSPORTATION_UPDATE**: logistics information update
	// *   **TICKET_UPDATE**: ticket update
	// *   **ALERT_UPDATE**: alert update
	// *   **AUTO_REPLY**: auto reply
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListChatappTemplateResponseBodyListTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateResponseBodyListTemplate) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetAuditStatus(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetCategory(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.Category = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetLanguage(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetTemplateCode(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.TemplateCode = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetTemplateName(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.TemplateName = &v
	return s
}

type ListChatappTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChatappTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponse) SetHeaders(v map[string]*string) *ListChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListChatappTemplateResponse) SetStatusCode(v int32) *ListChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatappTemplateResponse) SetBody(v *ListChatappTemplateResponseBody) *ListChatappTemplateResponse {
	s.Body = v
	return s
}

type ModifyChatappTemplateRequest struct {
	// The list of components of the message template.
	Components []*ModifyChatappTemplateRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// Isv customer space id
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The ID of the WhatApp Business account of the ISV customer.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// ISV verification code, which is used to verify whether the sub-account is authorized by ISV.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s ModifyChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequest) SetComponents(v []*ModifyChatappTemplateRequestComponents) *ModifyChatappTemplateRequest {
	s.Components = v
	return s
}

func (s *ModifyChatappTemplateRequest) SetCustSpaceId(v string) *ModifyChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetCustWabaId(v string) *ModifyChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetExample(v map[string]*string) *ModifyChatappTemplateRequest {
	s.Example = v
	return s
}

func (s *ModifyChatappTemplateRequest) SetIsvCode(v string) *ModifyChatappTemplateRequest {
	s.IsvCode = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetLanguage(v string) *ModifyChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetTemplateCode(v string) *ModifyChatappTemplateRequest {
	s.TemplateCode = &v
	return s
}

type ModifyChatappTemplateRequestComponents struct {
	// The list of buttons, which applies only to the **BUTTONS** component.
	Buttons []*ModifyChatappTemplateRequestComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description.
	//
	// >  When the Type parameter is set to **HEADER** and the Format parameter is set to **IMAGE/DOCUMENT/VIDEO**, you can specify the description.
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The name of the file.
	//
	// > : When the Type parameter is set to **HEADER** and the Format parameter is set to **DOCUMENT**, you can specify a name of the file.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the media resource.
	//
	// *   **TEXT**: text
	// *   **IMAGE**: image
	// *   **DOCUMENT**: document
	// *   **VIDEO**: video
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The text of the message to be sent.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the component.
	//
	// *   **BODY**
	// *   **HEADER**
	// *   **FOOTER**
	// *   **BUTTONS**
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the material.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ModifyChatappTemplateRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponents) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponents) SetButtons(v []*ModifyChatappTemplateRequestComponentsButtons) *ModifyChatappTemplateRequestComponents {
	s.Buttons = v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetCaption(v string) *ModifyChatappTemplateRequestComponents {
	s.Caption = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetFileName(v string) *ModifyChatappTemplateRequestComponents {
	s.FileName = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetFormat(v string) *ModifyChatappTemplateRequestComponents {
	s.Format = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetText(v string) *ModifyChatappTemplateRequestComponents {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetType(v string) *ModifyChatappTemplateRequestComponents {
	s.Type = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetUrl(v string) *ModifyChatappTemplateRequestComponents {
	s.Url = &v
	return s
}

type ModifyChatappTemplateRequestComponentsButtons struct {
	// The phone number.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The text of the message to be sent.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the button.
	//
	// *   **PHONE_NUMBER**: the phone call button
	// *   **URL**: the URL button
	// *   **QUICK_REPLY**: the quick reply button
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to be visited after clicking the button.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the URL. Valid values:
	//
	// *   **static**: the static URL
	// *   **dynamic**: the dynamic URL
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsButtons) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetPhoneNumber(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetText(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetType(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.Type = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetUrl(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.Url = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetUrlType(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.UrlType = &v
	return s
}

type ModifyChatappTemplateShrinkRequest struct {
	// The list of components of the message template.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// Isv customer space id
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The ID of the WhatApp Business account of the ISV customer.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	ExampleShrink *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// ISV verification code, which is used to verify whether the sub-account is authorized by ISV.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s ModifyChatappTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateShrinkRequest) SetComponentsShrink(v string) *ModifyChatappTemplateShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetCustSpaceId(v string) *ModifyChatappTemplateShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetCustWabaId(v string) *ModifyChatappTemplateShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetExampleShrink(v string) *ModifyChatappTemplateShrinkRequest {
	s.ExampleShrink = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetIsvCode(v string) *ModifyChatappTemplateShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetLanguage(v string) *ModifyChatappTemplateShrinkRequest {
	s.Language = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetTemplateCode(v string) *ModifyChatappTemplateShrinkRequest {
	s.TemplateCode = &v
	return s
}

type ModifyChatappTemplateResponseBody struct {
	// The HTTP status code.
	//
	// *   If OK is returned, the request is successful.
	// *   Other codes indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ModifyChatappTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateResponseBody) SetCode(v string) *ModifyChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyChatappTemplateResponseBody) SetData(v *ModifyChatappTemplateResponseBodyData) *ModifyChatappTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ModifyChatappTemplateResponseBody) SetMessage(v string) *ModifyChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyChatappTemplateResponseBody) SetRequestId(v string) *ModifyChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyChatappTemplateResponseBodyData struct {
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifyChatappTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateResponseBodyData) SetTemplateCode(v string) *ModifyChatappTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *ModifyChatappTemplateResponseBodyData) SetTemplateName(v string) *ModifyChatappTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

type ModifyChatappTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyChatappTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateResponse) SetHeaders(v map[string]*string) *ModifyChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyChatappTemplateResponse) SetStatusCode(v int32) *ModifyChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyChatappTemplateResponse) SetBody(v *ModifyChatappTemplateResponseBody) *ModifyChatappTemplateResponse {
	s.Body = v
	return s
}

type QueryChatappBindWabaRequest struct {
	// The space ID of the user under the ISV account.
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by ISV.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
}

func (s QueryChatappBindWabaRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappBindWabaRequest) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaRequest) SetCustSpaceId(v string) *QueryChatappBindWabaRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryChatappBindWabaRequest) SetIsvCode(v string) *QueryChatappBindWabaRequest {
	s.IsvCode = &v
	return s
}

type QueryChatappBindWabaResponseBody struct {
	// The HTTP status code returned.
	//
	// *   A value of OK indicates that the call is successful.
	// *   Other values indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryChatappBindWabaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryChatappBindWabaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappBindWabaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaResponseBody) SetCode(v string) *QueryChatappBindWabaResponseBody {
	s.Code = &v
	return s
}

func (s *QueryChatappBindWabaResponseBody) SetData(v *QueryChatappBindWabaResponseBodyData) *QueryChatappBindWabaResponseBody {
	s.Data = v
	return s
}

func (s *QueryChatappBindWabaResponseBody) SetMessage(v string) *QueryChatappBindWabaResponseBody {
	s.Message = &v
	return s
}

func (s *QueryChatappBindWabaResponseBody) SetRequestId(v string) *QueryChatappBindWabaResponseBody {
	s.RequestId = &v
	return s
}

type QueryChatappBindWabaResponseBodyData struct {
	// The audit status of the WhatApp Business account.
	AccountReviewStatus *string `json:"AccountReviewStatus,omitempty" xml:"AccountReviewStatus,omitempty"`
	// Currency
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// WabaID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The namespace of the message template.
	MessageTemplateNamespace *string `json:"MessageTemplateNamespace,omitempty" xml:"MessageTemplateNamespace,omitempty"`
	// The name of the WhatApp Business account.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryChatappBindWabaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappBindWabaResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaResponseBodyData) SetAccountReviewStatus(v string) *QueryChatappBindWabaResponseBodyData {
	s.AccountReviewStatus = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetCurrency(v string) *QueryChatappBindWabaResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetId(v string) *QueryChatappBindWabaResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetMessageTemplateNamespace(v string) *QueryChatappBindWabaResponseBodyData {
	s.MessageTemplateNamespace = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetName(v string) *QueryChatappBindWabaResponseBodyData {
	s.Name = &v
	return s
}

type QueryChatappBindWabaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryChatappBindWabaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryChatappBindWabaResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappBindWabaResponse) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaResponse) SetHeaders(v map[string]*string) *QueryChatappBindWabaResponse {
	s.Headers = v
	return s
}

func (s *QueryChatappBindWabaResponse) SetStatusCode(v int32) *QueryChatappBindWabaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChatappBindWabaResponse) SetBody(v *QueryChatappBindWabaResponseBody) *QueryChatappBindWabaResponse {
	s.Body = v
	return s
}

type QueryChatappPhoneNumbersRequest struct {
	// The space ID of the user under the ISV account.
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by ISV.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
}

func (s QueryChatappPhoneNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappPhoneNumbersRequest) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersRequest) SetCustSpaceId(v string) *QueryChatappPhoneNumbersRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryChatappPhoneNumbersRequest) SetIsvCode(v string) *QueryChatappPhoneNumbersRequest {
	s.IsvCode = &v
	return s
}

type QueryChatappPhoneNumbersResponseBody struct {
	// The HTTP status code returned.
	//
	// *   A value of OK indicates that the call is successful.
	// *   Other values indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of phone numbers.
	PhoneNumbers []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBody) SetCode(v string) *QueryChatappPhoneNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetMessage(v string) *QueryChatappPhoneNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetPhoneNumbers(v []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers) *QueryChatappPhoneNumbersResponseBody {
	s.PhoneNumbers = v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetRequestId(v string) *QueryChatappPhoneNumbersResponseBody {
	s.RequestId = &v
	return s
}

type QueryChatappPhoneNumbersResponseBodyPhoneNumbers struct {
	// 号码校验状态。
	CodeVerificationStatus *string `json:"CodeVerificationStatus,omitempty" xml:"CodeVerificationStatus,omitempty"`
	// 号码发送量。
	MessagingLimitTier *string `json:"MessagingLimitTier,omitempty" xml:"MessagingLimitTier,omitempty"`
	// 名称状态。
	NameStatus *string `json:"NameStatus,omitempty" xml:"NameStatus,omitempty"`
	// 新名称审核状态。
	NewNameStatus *string `json:"NewNameStatus,omitempty" xml:"NewNameStatus,omitempty"`
	// The phone number.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The quality rating of the phone number. Valid values: GREEN, YELLOW, RED, and UNKNOWN.
	QualityRating *string `json:"QualityRating,omitempty" xml:"QualityRating,omitempty"`
	// The status of the phone number. Valid values: PENDING, DELETED, MIGRATED, BANNED, RESTRICTED, RATE_LIMITED, FLAGGED, CONNECTED, DISCONNECTED, UNKNOWN, and UNVERIFIED.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status report notification URL.
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// The status report notification queue.
	StatusQueue *string `json:"StatusQueue,omitempty" xml:"StatusQueue,omitempty"`
	// The MO message notification URL.
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
	// The mobile originated (MO) message notification queue.
	UpQueue *string `json:"UpQueue,omitempty" xml:"UpQueue,omitempty"`
	// The name of the company to which the phone number is associated with.
	VerifiedName *string `json:"VerifiedName,omitempty" xml:"VerifiedName,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbers) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetCodeVerificationStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.CodeVerificationStatus = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetMessagingLimitTier(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.MessagingLimitTier = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetNameStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.NameStatus = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetNewNameStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.NewNameStatus = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetPhoneNumber(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.PhoneNumber = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetQualityRating(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.QualityRating = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.Status = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetStatusCallbackUrl(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.StatusCallbackUrl = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetStatusQueue(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.StatusQueue = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetUpCallbackUrl(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.UpCallbackUrl = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetUpQueue(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.UpQueue = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetVerifiedName(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.VerifiedName = &v
	return s
}

type QueryChatappPhoneNumbersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryChatappPhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryChatappPhoneNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponse) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponse) SetHeaders(v map[string]*string) *QueryChatappPhoneNumbersResponse {
	s.Headers = v
	return s
}

func (s *QueryChatappPhoneNumbersResponse) SetStatusCode(v int32) *QueryChatappPhoneNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponse) SetBody(v *QueryChatappPhoneNumbersResponseBody) *QueryChatappPhoneNumbersResponse {
	s.Body = v
	return s
}

type SendChatappMassMessageRequest struct {
	// The type of the message channel. Valid values:
	//
	// *   **whatsapp**
	// *   viber (under development)
	// *   line (under development)
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// Fallback message content.
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// Fallback strategy id. Fallback Strategy can be created on the ChatApp console.
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The mobile phone number of the message sender.
	//
	// <notice>You can specify a mobile phone number that is registered for a WhatsApp account and is approved in the ChatApp console.</notice>
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Target number and parameter list.
	SenderList []*SendChatappMassMessageRequestSenderList `json:"SenderList,omitempty" xml:"SenderList,omitempty" type:"Repeated"`
	// User-define ID to identify a single batch of messages.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s SendChatappMassMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequest) SetChannelType(v string) *SendChatappMassMessageRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetCustSpaceId(v string) *SendChatappMassMessageRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetCustWabaId(v string) *SendChatappMassMessageRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFallBackContent(v string) *SendChatappMassMessageRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFallBackId(v string) *SendChatappMassMessageRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFrom(v string) *SendChatappMassMessageRequest {
	s.From = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetIsvCode(v string) *SendChatappMassMessageRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetLanguage(v string) *SendChatappMassMessageRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetSenderList(v []*SendChatappMassMessageRequestSenderList) *SendChatappMassMessageRequest {
	s.SenderList = v
	return s
}

func (s *SendChatappMassMessageRequest) SetTaskId(v string) *SendChatappMassMessageRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetTemplateCode(v string) *SendChatappMassMessageRequest {
	s.TemplateCode = &v
	return s
}

type SendChatappMassMessageRequestSenderList struct {
	// Payload list.
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// Template parameters.
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// Target number.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendChatappMassMessageRequestSenderList) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderList) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderList) SetPayload(v []*string) *SendChatappMassMessageRequestSenderList {
	s.Payload = v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) SetTemplateParams(v map[string]*string) *SendChatappMassMessageRequestSenderList {
	s.TemplateParams = v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) SetTo(v string) *SendChatappMassMessageRequestSenderList {
	s.To = &v
	return s
}

type SendChatappMassMessageShrinkRequest struct {
	// The type of the message channel. Valid values:
	//
	// *   **whatsapp**
	// *   viber (under development)
	// *   line (under development)
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// Fallback message content.
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// Fallback strategy id. Fallback Strategy can be created on the ChatApp console.
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The mobile phone number of the message sender.
	//
	// <notice>You can specify a mobile phone number that is registered for a WhatsApp account and is approved in the ChatApp console.</notice>
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Target number and parameter list.
	SenderListShrink *string `json:"SenderList,omitempty" xml:"SenderList,omitempty"`
	// User-define ID to identify a single batch of messages.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The code of the message template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s SendChatappMassMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageShrinkRequest) SetChannelType(v string) *SendChatappMassMessageShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetCustSpaceId(v string) *SendChatappMassMessageShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetCustWabaId(v string) *SendChatappMassMessageShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackContent(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackId(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFrom(v string) *SendChatappMassMessageShrinkRequest {
	s.From = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetIsvCode(v string) *SendChatappMassMessageShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetLanguage(v string) *SendChatappMassMessageShrinkRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetSenderListShrink(v string) *SendChatappMassMessageShrinkRequest {
	s.SenderListShrink = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTaskId(v string) *SendChatappMassMessageShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTemplateCode(v string) *SendChatappMassMessageShrinkRequest {
	s.TemplateCode = &v
	return s
}

type SendChatappMassMessageResponseBody struct {
	// The HTTP status code returned.
	//
	// *   A code of OK indicates that the call is successful.
	// *   Other codes indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Batch send message ID.
	GroupMessageId *string `json:"GroupMessageId,omitempty" xml:"GroupMessageId,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendChatappMassMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageResponseBody) SetCode(v string) *SendChatappMassMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendChatappMassMessageResponseBody) SetGroupMessageId(v string) *SendChatappMassMessageResponseBody {
	s.GroupMessageId = &v
	return s
}

func (s *SendChatappMassMessageResponseBody) SetMessage(v string) *SendChatappMassMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendChatappMassMessageResponseBody) SetRequestId(v string) *SendChatappMassMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendChatappMassMessageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendChatappMassMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendChatappMassMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageResponse) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageResponse) SetHeaders(v map[string]*string) *SendChatappMassMessageResponse {
	s.Headers = v
	return s
}

func (s *SendChatappMassMessageResponse) SetStatusCode(v int32) *SendChatappMassMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendChatappMassMessageResponse) SetBody(v *SendChatappMassMessageResponseBody) *SendChatappMassMessageResponse {
	s.Body = v
	return s
}

type SendChatappMessageRequest struct {
	// The type of the message channel. Valid values:
	//
	// *   **whatsapp**
	// *   viber (under development)
	// *   line (under development)
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The content of the message.
	//
	// **
	//
	// **Note** The **Content** parameter is required if you set the **Type** parameter to **message**.
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ContextMessageId *string `json:"ContextMessageId,omitempty" xml:"ContextMessageId,omitempty"`
	CustSpaceId      *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// Fallback message content.
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// Fallback strategy id. Fallback Strategy can be created on the ChatApp console.
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The mobile phone number of the message sender.
	//
	// <notice>You can specify a mobile phone number that is registered for a WhatsApp account and is approved in the ChatApp console.</notice>
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The type of the message. This parameter is required if you set the Type parameter to **message**. Valid values:
	//
	// *   **text**: a text message. The **Text** parameter is required if you set the MessageType parameter to text.
	// *   **image**: an image message. The **Link** parameter is required and the **Caption** parameter is optional if you set the MessageType parameter to image.
	// *   **video**: a video message. The **Link** parameter is required and the **Caption** parameter is optional if you set the MessageType parameter to video.
	// *   **audio**: an audio message. The **Link** parameter is required and the **Caption** parameter is invalid if you set the MessageType parameter to audio.
	// *   **document**: a document message. The **Link** and **FileName** parameters are required and the **Caption** parameter is invalid if you set the MessageType parameter to document.
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The payload of the button.
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	Tag     *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The code of the message template. This parameter is required if you set the Type parameter to **template**.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The variables of the message template.
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The mobile phone number of the message recipient.
	To           *string `json:"To,omitempty" xml:"To,omitempty"`
	TrackingData *string `json:"TrackingData,omitempty" xml:"TrackingData,omitempty"`
	Ttl          *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the message. Valid values:
	//
	// *   **template**: a template message. A template message is sent based on a template that is created in the ChatApp console and is approved. You can send template messages based on your business requirements.
	// *   **message**: a custom message. You can send a custom message to a user only within 24 hours after you receive the last message from the user.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendChatappMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequest) SetChannelType(v string) *SendChatappMessageRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMessageRequest) SetContent(v string) *SendChatappMessageRequest {
	s.Content = &v
	return s
}

func (s *SendChatappMessageRequest) SetContextMessageId(v string) *SendChatappMessageRequest {
	s.ContextMessageId = &v
	return s
}

func (s *SendChatappMessageRequest) SetCustSpaceId(v string) *SendChatappMessageRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMessageRequest) SetCustWabaId(v string) *SendChatappMessageRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackContent(v string) *SendChatappMessageRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackId(v string) *SendChatappMessageRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMessageRequest) SetFrom(v string) *SendChatappMessageRequest {
	s.From = &v
	return s
}

func (s *SendChatappMessageRequest) SetIsvCode(v string) *SendChatappMessageRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMessageRequest) SetLabel(v string) *SendChatappMessageRequest {
	s.Label = &v
	return s
}

func (s *SendChatappMessageRequest) SetLanguage(v string) *SendChatappMessageRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMessageRequest) SetMessageType(v string) *SendChatappMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendChatappMessageRequest) SetPayload(v []*string) *SendChatappMessageRequest {
	s.Payload = v
	return s
}

func (s *SendChatappMessageRequest) SetTag(v string) *SendChatappMessageRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateCode(v string) *SendChatappMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateParams(v map[string]*string) *SendChatappMessageRequest {
	s.TemplateParams = v
	return s
}

func (s *SendChatappMessageRequest) SetTo(v string) *SendChatappMessageRequest {
	s.To = &v
	return s
}

func (s *SendChatappMessageRequest) SetTrackingData(v string) *SendChatappMessageRequest {
	s.TrackingData = &v
	return s
}

func (s *SendChatappMessageRequest) SetTtl(v int32) *SendChatappMessageRequest {
	s.Ttl = &v
	return s
}

func (s *SendChatappMessageRequest) SetType(v string) *SendChatappMessageRequest {
	s.Type = &v
	return s
}

type SendChatappMessageShrinkRequest struct {
	// The type of the message channel. Valid values:
	//
	// *   **whatsapp**
	// *   viber (under development)
	// *   line (under development)
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The content of the message.
	//
	// **
	//
	// **Note** The **Content** parameter is required if you set the **Type** parameter to **message**.
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ContextMessageId *string `json:"ContextMessageId,omitempty" xml:"ContextMessageId,omitempty"`
	CustSpaceId      *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	// The unique identifier of the WhatsApp account that you register.
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// Fallback message content.
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// Fallback strategy id. Fallback Strategy can be created on the ChatApp console.
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The mobile phone number of the message sender.
	//
	// <notice>You can specify a mobile phone number that is registered for a WhatsApp account and is approved in the ChatApp console.</notice>
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Assigned by ISV for RAM user authentication and authorization.
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language that is used in the message template.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The type of the message. This parameter is required if you set the Type parameter to **message**. Valid values:
	//
	// *   **text**: a text message. The **Text** parameter is required if you set the MessageType parameter to text.
	// *   **image**: an image message. The **Link** parameter is required and the **Caption** parameter is optional if you set the MessageType parameter to image.
	// *   **video**: a video message. The **Link** parameter is required and the **Caption** parameter is optional if you set the MessageType parameter to video.
	// *   **audio**: an audio message. The **Link** parameter is required and the **Caption** parameter is invalid if you set the MessageType parameter to audio.
	// *   **document**: a document message. The **Link** and **FileName** parameters are required and the **Caption** parameter is invalid if you set the MessageType parameter to document.
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The payload of the button.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	Tag           *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The code of the message template. This parameter is required if you set the Type parameter to **template**.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The variables of the message template.
	TemplateParamsShrink *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The mobile phone number of the message recipient.
	To           *string `json:"To,omitempty" xml:"To,omitempty"`
	TrackingData *string `json:"TrackingData,omitempty" xml:"TrackingData,omitempty"`
	Ttl          *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the message. Valid values:
	//
	// *   **template**: a template message. A template message is sent based on a template that is created in the ChatApp console and is approved. You can send template messages based on your business requirements.
	// *   **message**: a custom message. You can send a custom message to a user only within 24 hours after you receive the last message from the user.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendChatappMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMessageShrinkRequest) SetChannelType(v string) *SendChatappMessageShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetContent(v string) *SendChatappMessageShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetContextMessageId(v string) *SendChatappMessageShrinkRequest {
	s.ContextMessageId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetCustSpaceId(v string) *SendChatappMessageShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetCustWabaId(v string) *SendChatappMessageShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackContent(v string) *SendChatappMessageShrinkRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackId(v string) *SendChatappMessageShrinkRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFrom(v string) *SendChatappMessageShrinkRequest {
	s.From = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetIsvCode(v string) *SendChatappMessageShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetLabel(v string) *SendChatappMessageShrinkRequest {
	s.Label = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetLanguage(v string) *SendChatappMessageShrinkRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetMessageType(v string) *SendChatappMessageShrinkRequest {
	s.MessageType = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetPayloadShrink(v string) *SendChatappMessageShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTag(v string) *SendChatappMessageShrinkRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateCode(v string) *SendChatappMessageShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateParamsShrink(v string) *SendChatappMessageShrinkRequest {
	s.TemplateParamsShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTo(v string) *SendChatappMessageShrinkRequest {
	s.To = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTrackingData(v string) *SendChatappMessageShrinkRequest {
	s.TrackingData = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTtl(v int32) *SendChatappMessageShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetType(v string) *SendChatappMessageShrinkRequest {
	s.Type = &v
	return s
}

type SendChatappMessageResponseBody struct {
	// The HTTP status code returned.
	//
	// *   A code of OK indicates that the call is successful.
	// *   Other codes indicate that the call fails. For more information, see [Error codes](~~196974~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the message.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendChatappMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendChatappMessageResponseBody) SetCode(v string) *SendChatappMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetMessage(v string) *SendChatappMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetMessageId(v string) *SendChatappMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetRequestId(v string) *SendChatappMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendChatappMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendChatappMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendChatappMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageResponse) GoString() string {
	return s.String()
}

func (s *SendChatappMessageResponse) SetHeaders(v map[string]*string) *SendChatappMessageResponse {
	s.Headers = v
	return s
}

func (s *SendChatappMessageResponse) SetStatusCode(v int32) *SendChatappMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendChatappMessageResponse) SetBody(v *SendChatappMessageResponseBody) *SendChatappMessageResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cams"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param tmpReq BeeBotAssociateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BeeBotAssociateResponse
 */
func (client *Client) BeeBotAssociateWithOptions(tmpReq *BeeBotAssociateRequest, runtime *util.RuntimeOptions) (_result *BeeBotAssociateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BeeBotAssociateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Perspective)) {
		request.PerspectiveShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Perspective, tea.String("Perspective"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatBotInstanceId)) {
		body["ChatBotInstanceId"] = request.ChatBotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveShrink)) {
		body["Perspective"] = request.PerspectiveShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RecommendNum)) {
		body["RecommendNum"] = request.RecommendNum
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		body["Utterance"] = request.Utterance
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BeeBotAssociate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BeeBotAssociateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request BeeBotAssociateRequest
 * @return BeeBotAssociateResponse
 */
func (client *Client) BeeBotAssociate(request *BeeBotAssociateRequest) (_result *BeeBotAssociateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BeeBotAssociateResponse{}
	_body, _err := client.BeeBotAssociateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param tmpReq BeeBotChatRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BeeBotChatResponse
 */
func (client *Client) BeeBotChatWithOptions(tmpReq *BeeBotChatRequest, runtime *util.RuntimeOptions) (_result *BeeBotChatResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BeeBotChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Perspective)) {
		request.PerspectiveShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Perspective, tea.String("Perspective"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VendorParam)) {
		request.VendorParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VendorParam, tea.String("VendorParam"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatBotInstanceId)) {
		body["ChatBotInstanceId"] = request.ChatBotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentName)) {
		body["IntentName"] = request.IntentName
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveShrink)) {
		body["Perspective"] = request.PerspectiveShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		body["SenderId"] = request.SenderId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderNick)) {
		body["SenderNick"] = request.SenderNick
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		body["Utterance"] = request.Utterance
	}

	if !tea.BoolValue(util.IsUnset(request.VendorParamShrink)) {
		body["VendorParam"] = request.VendorParamShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BeeBotChat"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BeeBotChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request BeeBotChatRequest
 * @return BeeBotChatResponse
 */
func (client *Client) BeeBotChat(request *BeeBotChatRequest) (_result *BeeBotChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BeeBotChatResponse{}
	_body, _err := client.BeeBotChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param tmpReq CreateChatappTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateChatappTemplateResponse
 */
func (client *Client) CreateChatappTemplateWithOptions(tmpReq *CreateChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateChatappTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Components)) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, tea.String("Components"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Example)) {
		request.ExampleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Example, tea.String("Example"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentsShrink)) {
		body["Components"] = request.ComponentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.ExampleShrink)) {
		body["Example"] = request.ExampleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChatappTemplate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChatappTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request CreateChatappTemplateRequest
 * @return CreateChatappTemplateResponse
 */
func (client *Client) CreateChatappTemplate(request *CreateChatappTemplateRequest) (_result *CreateChatappTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChatappTemplateResponse{}
	_body, _err := client.CreateChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request DeleteChatappTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteChatappTemplateResponse
 */
func (client *Client) DeleteChatappTemplateWithOptions(request *DeleteChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteChatappTemplate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteChatappTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request DeleteChatappTemplateRequest
 * @return DeleteChatappTemplateResponse
 */
func (client *Client) DeleteChatappTemplate(request *DeleteChatappTemplateRequest) (_result *DeleteChatappTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChatappTemplateResponse{}
	_body, _err := client.DeleteChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request GetChatappTemplateDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetChatappTemplateDetailResponse
 */
func (client *Client) GetChatappTemplateDetailWithOptions(request *GetChatappTemplateDetailRequest, runtime *util.RuntimeOptions) (_result *GetChatappTemplateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatappTemplateDetail"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatappTemplateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request GetChatappTemplateDetailRequest
 * @return GetChatappTemplateDetailResponse
 */
func (client *Client) GetChatappTemplateDetail(request *GetChatappTemplateDetailRequest) (_result *GetChatappTemplateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChatappTemplateDetailResponse{}
	_body, _err := client.GetChatappTemplateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param tmpReq ListChatappTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListChatappTemplateResponse
 */
func (client *Client) ListChatappTemplateWithOptions(tmpReq *ListChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *ListChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListChatappTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Page)) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, tea.String("Page"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		query["Page"] = request.PageShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChatappTemplate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChatappTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 5 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request ListChatappTemplateRequest
 * @return ListChatappTemplateResponse
 */
func (client *Client) ListChatappTemplate(request *ListChatappTemplateRequest) (_result *ListChatappTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChatappTemplateResponse{}
	_body, _err := client.ListChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyChatappTemplateWithOptions(tmpReq *ModifyChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyChatappTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Components)) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, tea.String("Components"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Example)) {
		request.ExampleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Example, tea.String("Example"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentsShrink)) {
		body["Components"] = request.ComponentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.ExampleShrink)) {
		body["Example"] = request.ExampleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyChatappTemplate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyChatappTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyChatappTemplate(request *ModifyChatappTemplateRequest) (_result *ModifyChatappTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyChatappTemplateResponse{}
	_body, _err := client.ModifyChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request QueryChatappBindWabaRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryChatappBindWabaResponse
 */
func (client *Client) QueryChatappBindWabaWithOptions(request *QueryChatappBindWabaRequest, runtime *util.RuntimeOptions) (_result *QueryChatappBindWabaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["IsvCode"] = request.IsvCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryChatappBindWaba"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryChatappBindWabaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request QueryChatappBindWabaRequest
 * @return QueryChatappBindWabaResponse
 */
func (client *Client) QueryChatappBindWaba(request *QueryChatappBindWabaRequest) (_result *QueryChatappBindWabaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryChatappBindWabaResponse{}
	_body, _err := client.QueryChatappBindWabaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request QueryChatappPhoneNumbersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryChatappPhoneNumbersResponse
 */
func (client *Client) QueryChatappPhoneNumbersWithOptions(request *QueryChatappPhoneNumbersRequest, runtime *util.RuntimeOptions) (_result *QueryChatappPhoneNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["IsvCode"] = request.IsvCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryChatappPhoneNumbers"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryChatappPhoneNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request QueryChatappPhoneNumbersRequest
 * @return QueryChatappPhoneNumbersResponse
 */
func (client *Client) QueryChatappPhoneNumbers(request *QueryChatappPhoneNumbersRequest) (_result *QueryChatappPhoneNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryChatappPhoneNumbersResponse{}
	_body, _err := client.QueryChatappPhoneNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 * Can support up to 1,000 numbers per request.
 *
 * @param tmpReq SendChatappMassMessageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendChatappMassMessageResponse
 */
func (client *Client) SendChatappMassMessageWithOptions(tmpReq *SendChatappMassMessageRequest, runtime *util.RuntimeOptions) (_result *SendChatappMassMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendChatappMassMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SenderList)) {
		request.SenderListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SenderList, tea.String("SenderList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackContent)) {
		body["FallBackContent"] = request.FallBackContent
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackId)) {
		body["FallBackId"] = request.FallBackId
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SenderListShrink)) {
		body["SenderList"] = request.SenderListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendChatappMassMessage"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendChatappMassMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 * Can support up to 1,000 numbers per request.
 *
 * @param request SendChatappMassMessageRequest
 * @return SendChatappMassMessageResponse
 */
func (client *Client) SendChatappMassMessage(request *SendChatappMassMessageRequest) (_result *SendChatappMassMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendChatappMassMessageResponse{}
	_body, _err := client.SendChatappMassMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 200 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param tmpReq SendChatappMessageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendChatappMessageResponse
 */
func (client *Client) SendChatappMessageWithOptions(tmpReq *SendChatappMessageRequest, runtime *util.RuntimeOptions) (_result *SendChatappMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendChatappMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TemplateParams)) {
		request.TemplateParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateParams, tea.String("TemplateParams"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.ContextMessageId)) {
		body["ContextMessageId"] = request.ContextMessageId
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackContent)) {
		body["FallBackContent"] = request.FallBackContent
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackId)) {
		body["FallBackId"] = request.FallBackId
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParamsShrink)) {
		body["TemplateParams"] = request.TemplateParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.TrackingData)) {
		body["TrackingData"] = request.TrackingData
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["Ttl"] = request.Ttl
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendChatappMessage"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendChatappMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 200 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request SendChatappMessageRequest
 * @return SendChatappMessageResponse
 */
func (client *Client) SendChatappMessage(request *SendChatappMessageRequest) (_result *SendChatappMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendChatappMessageResponse{}
	_body, _err := client.SendChatappMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
