// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddChatappPhoneNumberRequest struct {
	// The country code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	Cc *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 939283893939
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number without a country code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13800000000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The ID of the pre-registered phone number.
	//
	// example:
	//
	// 10202020
	PreValidateId *string `json:"PreValidateId,omitempty" xml:"PreValidateId,omitempty"`
	// The display name of the phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba
	VerifiedName *string `json:"VerifiedName,omitempty" xml:"VerifiedName,omitempty"`
}

func (s AddChatappPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddChatappPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *AddChatappPhoneNumberRequest) SetCc(v string) *AddChatappPhoneNumberRequest {
	s.Cc = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetCustSpaceId(v string) *AddChatappPhoneNumberRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetPhoneNumber(v string) *AddChatappPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetPreValidateId(v string) *AddChatappPhoneNumberRequest {
	s.PreValidateId = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetVerifiedName(v string) *AddChatappPhoneNumberRequest {
	s.VerifiedName = &v
	return s
}

type AddChatappPhoneNumberResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddChatappPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddChatappPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *AddChatappPhoneNumberResponseBody) SetAccessDeniedDetail(v string) *AddChatappPhoneNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddChatappPhoneNumberResponseBody) SetCode(v string) *AddChatappPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *AddChatappPhoneNumberResponseBody) SetMessage(v string) *AddChatappPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *AddChatappPhoneNumberResponseBody) SetRequestId(v string) *AddChatappPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

type AddChatappPhoneNumberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddChatappPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddChatappPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddChatappPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *AddChatappPhoneNumberResponse) SetHeaders(v map[string]*string) *AddChatappPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *AddChatappPhoneNumberResponse) SetStatusCode(v int32) *AddChatappPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddChatappPhoneNumberResponse) SetBody(v *AddChatappPhoneNumberResponseBody) *AddChatappPhoneNumberResponse {
	s.Body = v
	return s
}

type BeeBotAssociateRequest struct {
	// The ID of a bot instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ChatBotInstanceId *string `json:"ChatBotInstanceId,omitempty" xml:"ChatBotInstanceId,omitempty"`
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by ISV.
	//
	// example:
	//
	// ksiekdki39ksks93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The list of codes for answers from different perspectives.
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// The number of recommended questions. The value ranges from 1 to 10.
	//
	// example:
	//
	// 3
	RecommendNum *int32 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// The ID of the session, which is used to identify the session and store context information in the session.
	//
	// example:
	//
	// 2334324234
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input of the visitor.
	//
	// example:
	//
	// 你好
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
	// The ID of a bot instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ChatBotInstanceId *string `json:"ChatBotInstanceId,omitempty" xml:"ChatBotInstanceId,omitempty"`
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by ISV.
	//
	// example:
	//
	// ksiekdki39ksks93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The list of codes for answers from different perspectives.
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
	// The number of recommended questions. The value ranges from 1 to 10.
	//
	// example:
	//
	// 3
	RecommendNum *int32 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// The ID of the session, which is used to identify the session and store context information in the session.
	//
	// example:
	//
	// 2334324234
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input of the visitor.
	//
	// example:
	//
	// 你好
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *BeeBotAssociateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BeeBotAssociateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeeBotAssociateResponseBody) GoString() string {
	return s.String()
}

func (s *BeeBotAssociateResponseBody) SetAccessDeniedDetail(v string) *BeeBotAssociateResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	//
	// example:
	//
	// 1eb47d7a1706429081e90c83c62c2f00
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The ID of the session.
	//
	// example:
	//
	// 93f11165a2a24289a6f869760e8cb3f3
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
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BeeBotAssociateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bot instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ChatBotInstanceId *string `json:"ChatBotInstanceId,omitempty" xml:"ChatBotInstanceId,omitempty"`
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the intent in the dialog flow. When this parameter is specified, the bot conducts a Q\\&A based on the intent.
	//
	// example:
	//
	// intent
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// ksiekdki39ksks93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The ID of the knowledge title in the knowledge base.
	//
	// example:
	//
	// 1
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The list of codes for answers from different perspectives.
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// The ID of the visitor, which is used to identify users in the current session.
	//
	// example:
	//
	// 861500000000
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The nickname of the visitor in the current session.
	//
	// example:
	//
	// nick
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// The ID of the session, which is used to identify the session and store context information of the session.
	//
	// example:
	//
	// en
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input of the visitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// 659216218162179
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
	// The ID of the bot instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ChatBotInstanceId *string `json:"ChatBotInstanceId,omitempty" xml:"ChatBotInstanceId,omitempty"`
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the intent in the dialog flow. When this parameter is specified, the bot conducts a Q\\&A based on the intent.
	//
	// example:
	//
	// intent
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// ksiekdki39ksks93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The ID of the knowledge title in the knowledge base.
	//
	// example:
	//
	// 1
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The list of codes for answers from different perspectives.
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
	// The ID of the visitor, which is used to identify users in the current session.
	//
	// example:
	//
	// 861500000000
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The nickname of the visitor in the current session.
	//
	// example:
	//
	// nick
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// The ID of the session, which is used to identify the session and store context information of the session.
	//
	// example:
	//
	// en
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input of the visitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// 659216218162179
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *BeeBotChatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// none
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BeeBotChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeeBotChatResponseBody) GoString() string {
	return s.String()
}

func (s *BeeBotChatResponseBody) SetAccessDeniedDetail(v string) *BeeBotChatResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	//
	// example:
	//
	// ab6be8af-cee4-40c3-9919-2ac7461d7d98
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The list of messages.
	Messages []*BeeBotChatResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The ID of the session.
	//
	// example:
	//
	// 1234
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
	// When the AnswerType parameter is set to Recommend, this parameter indicates the source of the recommended answer.
	//
	// example:
	//
	// KNOWLEDGE
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The type of the answer.
	//
	// example:
	//
	// Text
	AnswerType *string `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	// When the AnswerType parameter is set to Knowledge, this parameter contains the Knowledge object returned by the bot.
	Knowledge *BeeBotChatResponseBodyDataMessagesKnowledge `json:"Knowledge,omitempty" xml:"Knowledge,omitempty" type:"Struct"`
	// The list of recommended knowledge. When the AnswerType parameter is set to Recommend, this parameter is returned.
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
	//
	// example:
	//
	// KnowledgeBase
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The category of the knowledge.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the hit question.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether the answer is in plain text or rich text.
	//
	// example:
	//
	// PLAIN_TEXT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The hit text.
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// The ID of the hit question in the knowledge base.
	//
	// example:
	//
	// 735898
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
	//
	// example:
	//
	// 735899
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
	//
	// example:
	//
	// KNOWLEDGE
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The ID of the recommended knowledge.
	//
	// example:
	//
	// 4548
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The title of the recommended knowledge. Valid values: the entity in graph-based question answering, the knowledge title in knowledge-based question answering, or the column value in table-based question answering.
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
	//
	// example:
	//
	// BotFramework
	AnswerSource *string `json:"AnswerSource,omitempty" xml:"AnswerSource,omitempty"`
	// The content of the text message.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether the answer is in plain text or rich text.
	//
	// example:
	//
	// PLAIN_TEXT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The name of the dialog. When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	DialogName *string `json:"DialogName,omitempty" xml:"DialogName,omitempty"`
	// The passthrough parameters are returned.
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	ExternalFlags map[string]interface{} `json:"ExternalFlags,omitempty" xml:"ExternalFlags,omitempty"`
	// The hit text.
	HitStatement *string `json:"HitStatement,omitempty" xml:"HitStatement,omitempty"`
	// The name of the intent. When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// The metadata.
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// The ID of the node. When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
	//
	// example:
	//
	// 1410-c7a72a78.__city
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node. When the AnswerSource parameter is set to BotFramework, the value of this parameter is returned.
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
	//
	// example:
	//
	// false
	Hit *bool `json:"Hit,omitempty" xml:"Hit,omitempty"`
	// The name.
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BeeBotChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ChatappBindWabaRequest struct {
	// The ID of the WhatsApp Business account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2293938222343****
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s ChatappBindWabaRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatappBindWabaRequest) GoString() string {
	return s.String()
}

func (s *ChatappBindWabaRequest) SetWabaId(v string) *ChatappBindWabaRequest {
	s.WabaId = &v
	return s
}

type ChatappBindWabaResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ChatappBindWabaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappBindWabaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatappBindWabaResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappBindWabaResponseBody) SetAccessDeniedDetail(v string) *ChatappBindWabaResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappBindWabaResponseBody) SetCode(v string) *ChatappBindWabaResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappBindWabaResponseBody) SetData(v *ChatappBindWabaResponseBodyData) *ChatappBindWabaResponseBody {
	s.Data = v
	return s
}

func (s *ChatappBindWabaResponseBody) SetMessage(v string) *ChatappBindWabaResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappBindWabaResponseBody) SetRequestId(v string) *ChatappBindWabaResponseBody {
	s.RequestId = &v
	return s
}

type ChatappBindWabaResponseBodyData struct {
	// The space ID of the user.
	//
	// example:
	//
	// C02029392939939
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ID of the WhatsApp Business account.
	//
	// example:
	//
	// 2939828282
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s ChatappBindWabaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChatappBindWabaResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChatappBindWabaResponseBodyData) SetCustSpaceId(v string) *ChatappBindWabaResponseBodyData {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappBindWabaResponseBodyData) SetWabaId(v string) *ChatappBindWabaResponseBodyData {
	s.WabaId = &v
	return s
}

type ChatappBindWabaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappBindWabaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappBindWabaResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatappBindWabaResponse) GoString() string {
	return s.String()
}

func (s *ChatappBindWabaResponse) SetHeaders(v map[string]*string) *ChatappBindWabaResponse {
	s.Headers = v
	return s
}

func (s *ChatappBindWabaResponse) SetStatusCode(v int32) *ChatappBindWabaResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappBindWabaResponse) SetBody(v *ChatappBindWabaResponseBody) *ChatappBindWabaResponse {
	s.Body = v
	return s
}

type ChatappEmbedSignUpRequest struct {
	// The InputToken returned after the embedded signup flow is complete.
	//
	// This parameter is required.
	//
	// example:
	//
	// wlelkelwidilwloe-ewlwols0lwsllsld
	InputToken *string `json:"InputToken,omitempty" xml:"InputToken,omitempty"`
}

func (s ChatappEmbedSignUpRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatappEmbedSignUpRequest) GoString() string {
	return s.String()
}

func (s *ChatappEmbedSignUpRequest) SetInputToken(v string) *ChatappEmbedSignUpRequest {
	s.InputToken = &v
	return s
}

type ChatappEmbedSignUpResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of the WhatsApp Business accounts.
	Wabas []*ChatappEmbedSignUpResponseBodyWabas `json:"Wabas,omitempty" xml:"Wabas,omitempty" type:"Repeated"`
}

func (s ChatappEmbedSignUpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatappEmbedSignUpResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappEmbedSignUpResponseBody) SetAccessDeniedDetail(v string) *ChatappEmbedSignUpResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) SetCode(v string) *ChatappEmbedSignUpResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) SetMessage(v string) *ChatappEmbedSignUpResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) SetRequestId(v string) *ChatappEmbedSignUpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBody) SetWabas(v []*ChatappEmbedSignUpResponseBodyWabas) *ChatappEmbedSignUpResponseBody {
	s.Wabas = v
	return s
}

type ChatappEmbedSignUpResponseBodyWabas struct {
	// The review state of the WABA.
	//
	// example:
	//
	// VERIFIED
	AccountReviewStatus *string `json:"AccountReviewStatus,omitempty" xml:"AccountReviewStatus,omitempty"`
	// The currency.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The ID of the WABA.
	//
	// example:
	//
	// 2939933992*****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The namespace of the message template.
	//
	// example:
	//
	// alals-lsslls-slslsos-slsl
	MessageTemplateNamespace *string `json:"MessageTemplateNamespace,omitempty" xml:"MessageTemplateNamespace,omitempty"`
	// The name of the WABA.
	//
	// example:
	//
	// Alibaba
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChatappEmbedSignUpResponseBodyWabas) String() string {
	return tea.Prettify(s)
}

func (s ChatappEmbedSignUpResponseBodyWabas) GoString() string {
	return s.String()
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetAccountReviewStatus(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.AccountReviewStatus = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetCurrency(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.Currency = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetId(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.Id = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetMessageTemplateNamespace(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.MessageTemplateNamespace = &v
	return s
}

func (s *ChatappEmbedSignUpResponseBodyWabas) SetName(v string) *ChatappEmbedSignUpResponseBodyWabas {
	s.Name = &v
	return s
}

type ChatappEmbedSignUpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappEmbedSignUpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappEmbedSignUpResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatappEmbedSignUpResponse) GoString() string {
	return s.String()
}

func (s *ChatappEmbedSignUpResponse) SetHeaders(v map[string]*string) *ChatappEmbedSignUpResponse {
	s.Headers = v
	return s
}

func (s *ChatappEmbedSignUpResponse) SetStatusCode(v int32) *ChatappEmbedSignUpResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappEmbedSignUpResponse) SetBody(v *ChatappEmbedSignUpResponseBody) *ChatappEmbedSignUpResponse {
	s.Body = v
	return s
}

type ChatappMigrationRegisterRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86138001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s ChatappMigrationRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatappMigrationRegisterRequest) GoString() string {
	return s.String()
}

func (s *ChatappMigrationRegisterRequest) SetCustSpaceId(v string) *ChatappMigrationRegisterRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappMigrationRegisterRequest) SetPhoneNumber(v string) *ChatappMigrationRegisterRequest {
	s.PhoneNumber = &v
	return s
}

type ChatappMigrationRegisterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappMigrationRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatappMigrationRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappMigrationRegisterResponseBody) SetAccessDeniedDetail(v string) *ChatappMigrationRegisterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappMigrationRegisterResponseBody) SetCode(v string) *ChatappMigrationRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappMigrationRegisterResponseBody) SetMessage(v string) *ChatappMigrationRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappMigrationRegisterResponseBody) SetRequestId(v string) *ChatappMigrationRegisterResponseBody {
	s.RequestId = &v
	return s
}

type ChatappMigrationRegisterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappMigrationRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappMigrationRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatappMigrationRegisterResponse) GoString() string {
	return s.String()
}

func (s *ChatappMigrationRegisterResponse) SetHeaders(v map[string]*string) *ChatappMigrationRegisterResponse {
	s.Headers = v
	return s
}

func (s *ChatappMigrationRegisterResponse) SetStatusCode(v int32) *ChatappMigrationRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappMigrationRegisterResponse) SetBody(v *ChatappMigrationRegisterResponseBody) *ChatappMigrationRegisterResponse {
	s.Body = v
	return s
}

type ChatappMigrationVerifiedRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861380001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 828798
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s ChatappMigrationVerifiedRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatappMigrationVerifiedRequest) GoString() string {
	return s.String()
}

func (s *ChatappMigrationVerifiedRequest) SetCustSpaceId(v string) *ChatappMigrationVerifiedRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappMigrationVerifiedRequest) SetPhoneNumber(v string) *ChatappMigrationVerifiedRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappMigrationVerifiedRequest) SetVerifyCode(v string) *ChatappMigrationVerifiedRequest {
	s.VerifyCode = &v
	return s
}

type ChatappMigrationVerifiedResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ChatappMigrationVerifiedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappMigrationVerifiedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatappMigrationVerifiedResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappMigrationVerifiedResponseBody) SetAccessDeniedDetail(v string) *ChatappMigrationVerifiedResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) SetCode(v string) *ChatappMigrationVerifiedResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) SetData(v *ChatappMigrationVerifiedResponseBodyData) *ChatappMigrationVerifiedResponseBody {
	s.Data = v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) SetMessage(v string) *ChatappMigrationVerifiedResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) SetRequestId(v string) *ChatappMigrationVerifiedResponseBody {
	s.RequestId = &v
	return s
}

type ChatappMigrationVerifiedResponseBodyData struct {
	// The ID of the phone number.
	//
	// example:
	//
	// 82828893332
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s ChatappMigrationVerifiedResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChatappMigrationVerifiedResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChatappMigrationVerifiedResponseBodyData) SetId(v string) *ChatappMigrationVerifiedResponseBodyData {
	s.Id = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBodyData) SetPhoneNumber(v string) *ChatappMigrationVerifiedResponseBodyData {
	s.PhoneNumber = &v
	return s
}

type ChatappMigrationVerifiedResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappMigrationVerifiedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappMigrationVerifiedResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatappMigrationVerifiedResponse) GoString() string {
	return s.String()
}

func (s *ChatappMigrationVerifiedResponse) SetHeaders(v map[string]*string) *ChatappMigrationVerifiedResponse {
	s.Headers = v
	return s
}

func (s *ChatappMigrationVerifiedResponse) SetStatusCode(v int32) *ChatappMigrationVerifiedResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappMigrationVerifiedResponse) SetBody(v *ChatappMigrationVerifiedResponseBody) *ChatappMigrationVerifiedResponse {
	s.Body = v
	return s
}

type ChatappPhoneNumberDeregisterRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 939283893939
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800000000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s ChatappPhoneNumberDeregisterRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatappPhoneNumberDeregisterRequest) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberDeregisterRequest) SetCustSpaceId(v string) *ChatappPhoneNumberDeregisterRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterRequest) SetPhoneNumber(v string) *ChatappPhoneNumberDeregisterRequest {
	s.PhoneNumber = &v
	return s
}

type ChatappPhoneNumberDeregisterResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappPhoneNumberDeregisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatappPhoneNumberDeregisterResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberDeregisterResponseBody) SetAccessDeniedDetail(v string) *ChatappPhoneNumberDeregisterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponseBody) SetCode(v string) *ChatappPhoneNumberDeregisterResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponseBody) SetMessage(v string) *ChatappPhoneNumberDeregisterResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponseBody) SetRequestId(v string) *ChatappPhoneNumberDeregisterResponseBody {
	s.RequestId = &v
	return s
}

type ChatappPhoneNumberDeregisterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappPhoneNumberDeregisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappPhoneNumberDeregisterResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatappPhoneNumberDeregisterResponse) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberDeregisterResponse) SetHeaders(v map[string]*string) *ChatappPhoneNumberDeregisterResponse {
	s.Headers = v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponse) SetStatusCode(v int32) *ChatappPhoneNumberDeregisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponse) SetBody(v *ChatappPhoneNumberDeregisterResponseBody) *ChatappPhoneNumberDeregisterResponse {
	s.Body = v
	return s
}

type ChatappPhoneNumberRegisterRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 939283893939***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s ChatappPhoneNumberRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatappPhoneNumberRegisterRequest) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberRegisterRequest) SetCustSpaceId(v string) *ChatappPhoneNumberRegisterRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappPhoneNumberRegisterRequest) SetPhoneNumber(v string) *ChatappPhoneNumberRegisterRequest {
	s.PhoneNumber = &v
	return s
}

type ChatappPhoneNumberRegisterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// None.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappPhoneNumberRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatappPhoneNumberRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetAccessDeniedDetail(v string) *ChatappPhoneNumberRegisterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetCode(v string) *ChatappPhoneNumberRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetMessage(v string) *ChatappPhoneNumberRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponseBody) SetRequestId(v string) *ChatappPhoneNumberRegisterResponseBody {
	s.RequestId = &v
	return s
}

type ChatappPhoneNumberRegisterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappPhoneNumberRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappPhoneNumberRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatappPhoneNumberRegisterResponse) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberRegisterResponse) SetHeaders(v map[string]*string) *ChatappPhoneNumberRegisterResponse {
	s.Headers = v
	return s
}

func (s *ChatappPhoneNumberRegisterResponse) SetStatusCode(v int32) *ChatappPhoneNumberRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponse) SetBody(v *ChatappPhoneNumberRegisterResponseBody) *ChatappPhoneNumberRegisterResponse {
	s.Body = v
	return s
}

type ChatappSyncPhoneNumberRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
}

func (s ChatappSyncPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatappSyncPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberRequest) SetCustSpaceId(v string) *ChatappSyncPhoneNumberRequest {
	s.CustSpaceId = &v
	return s
}

type ChatappSyncPhoneNumberResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// None.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Details of the phone numbers.
	PhoneNumbers []*ChatappSyncPhoneNumberResponseBodyPhoneNumbers `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappSyncPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBody) SetAccessDeniedDetail(v string) *ChatappSyncPhoneNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetCode(v string) *ChatappSyncPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetMessage(v string) *ChatappSyncPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetPhoneNumbers(v []*ChatappSyncPhoneNumberResponseBodyPhoneNumbers) *ChatappSyncPhoneNumberResponseBody {
	s.PhoneNumbers = v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetRequestId(v string) *ChatappSyncPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

type ChatappSyncPhoneNumberResponseBodyPhoneNumbers struct {
	// The verification status.
	//
	// example:
	//
	// VERIFIED
	CodeVerificationStatus *string `json:"CodeVerificationStatus,omitempty" xml:"CodeVerificationStatus,omitempty"`
	// The number of phone numbers to which messages can be sent in a day.
	//
	// example:
	//
	// TIER_10
	MessagingLimitTier *string `json:"MessagingLimitTier,omitempty" xml:"MessagingLimitTier,omitempty"`
	// The review status of the business display name.
	//
	// example:
	//
	// Approval
	NameStatus *string `json:"NameStatus,omitempty" xml:"NameStatus,omitempty"`
	// The review status of the new business display name.
	//
	// example:
	//
	// Approval
	NewNameStatus *string `json:"NewNameStatus,omitempty" xml:"NewNameStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The quality rating of the phone number. Valid values: GREEN, YELLOW, and RED.
	//
	// example:
	//
	// GREEN
	QualityRating *string `json:"QualityRating,omitempty" xml:"QualityRating,omitempty"`
	// The status of the phone number.
	//
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The callback URL to which status reports are sent by using HTTP callbacks.
	//
	// example:
	//
	// https://www.alibaba.com/status
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// The status report queue.
	//
	// example:
	//
	// alicom-09399200-queue
	StatusQueue *string `json:"StatusQueue,omitempty" xml:"StatusQueue,omitempty"`
	// The callback URL to which MO messages are sent by using HTTP callbacks.
	//
	// example:
	//
	// https://www.alibaba.com/inbound
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
	// The mobile originated (MO) message queue.
	//
	// example:
	//
	// alicom-09399200-queue
	UpQueue *string `json:"UpQueue,omitempty" xml:"UpQueue,omitempty"`
	// The display name of the business to which the phone number belongs.
	//
	// example:
	//
	// Alibaba
	VerifiedName *string `json:"VerifiedName,omitempty" xml:"VerifiedName,omitempty"`
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbers) String() string {
	return tea.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetCodeVerificationStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.CodeVerificationStatus = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetMessagingLimitTier(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.MessagingLimitTier = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetNameStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.NameStatus = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetNewNameStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.NewNameStatus = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetPhoneNumber(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetQualityRating(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.QualityRating = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.Status = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetStatusCallbackUrl(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.StatusCallbackUrl = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetStatusQueue(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.StatusQueue = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetUpCallbackUrl(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.UpCallbackUrl = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetUpQueue(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.UpQueue = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetVerifiedName(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.VerifiedName = &v
	return s
}

type ChatappSyncPhoneNumberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappSyncPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappSyncPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponse) SetHeaders(v map[string]*string) *ChatappSyncPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ChatappSyncPhoneNumberResponse) SetStatusCode(v int32) *ChatappSyncPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponse) SetBody(v *ChatappSyncPhoneNumberResponseBody) *ChatappSyncPhoneNumberResponse {
	s.Body = v
	return s
}

type ChatappVerifyAndRegisterRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 29389299388383
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 223222
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s ChatappVerifyAndRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatappVerifyAndRegisterRequest) GoString() string {
	return s.String()
}

func (s *ChatappVerifyAndRegisterRequest) SetCustSpaceId(v string) *ChatappVerifyAndRegisterRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappVerifyAndRegisterRequest) SetPhoneNumber(v string) *ChatappVerifyAndRegisterRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappVerifyAndRegisterRequest) SetVerifyCode(v string) *ChatappVerifyAndRegisterRequest {
	s.VerifyCode = &v
	return s
}

type ChatappVerifyAndRegisterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappVerifyAndRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatappVerifyAndRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappVerifyAndRegisterResponseBody) SetAccessDeniedDetail(v string) *ChatappVerifyAndRegisterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponseBody) SetCode(v string) *ChatappVerifyAndRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponseBody) SetMessage(v string) *ChatappVerifyAndRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponseBody) SetRequestId(v string) *ChatappVerifyAndRegisterResponseBody {
	s.RequestId = &v
	return s
}

type ChatappVerifyAndRegisterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappVerifyAndRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappVerifyAndRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatappVerifyAndRegisterResponse) GoString() string {
	return s.String()
}

func (s *ChatappVerifyAndRegisterResponse) SetHeaders(v map[string]*string) *ChatappVerifyAndRegisterResponse {
	s.Headers = v
	return s
}

func (s *ChatappVerifyAndRegisterResponse) SetStatusCode(v int32) *ChatappVerifyAndRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappVerifyAndRegisterResponse) SetBody(v *ChatappVerifyAndRegisterResponseBody) *ChatappVerifyAndRegisterResponse {
	s.Body = v
	return s
}

type CreateChatappMigrationInitiateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 86
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13900001234
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
}

func (s CreateChatappMigrationInitiateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappMigrationInitiateRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappMigrationInitiateRequest) SetCountryCode(v string) *CreateChatappMigrationInitiateRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateChatappMigrationInitiateRequest) SetCustSpaceId(v string) *CreateChatappMigrationInitiateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateChatappMigrationInitiateRequest) SetMobileNumber(v string) *CreateChatappMigrationInitiateRequest {
	s.MobileNumber = &v
	return s
}

type CreateChatappMigrationInitiateResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateChatappMigrationInitiateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChatappMigrationInitiateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappMigrationInitiateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatappMigrationInitiateResponseBody) SetAccessDeniedDetail(v string) *CreateChatappMigrationInitiateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) SetCode(v string) *CreateChatappMigrationInitiateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) SetData(v *CreateChatappMigrationInitiateResponseBodyData) *CreateChatappMigrationInitiateResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) SetMessage(v string) *CreateChatappMigrationInitiateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) SetRequestId(v string) *CreateChatappMigrationInitiateResponseBody {
	s.RequestId = &v
	return s
}

type CreateChatappMigrationInitiateResponseBodyData struct {
	// example:
	//
	// 82828893332
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 8613900001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// MIGRATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateChatappMigrationInitiateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappMigrationInitiateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatappMigrationInitiateResponseBodyData) SetId(v string) *CreateChatappMigrationInitiateResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBodyData) SetPhoneNumber(v string) *CreateChatappMigrationInitiateResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBodyData) SetStatus(v string) *CreateChatappMigrationInitiateResponseBodyData {
	s.Status = &v
	return s
}

type CreateChatappMigrationInitiateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatappMigrationInitiateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatappMigrationInitiateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappMigrationInitiateResponse) GoString() string {
	return s.String()
}

func (s *CreateChatappMigrationInitiateResponse) SetHeaders(v map[string]*string) *CreateChatappMigrationInitiateResponse {
	s.Headers = v
	return s
}

func (s *CreateChatappMigrationInitiateResponse) SetStatusCode(v int32) *CreateChatappMigrationInitiateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponse) SetBody(v *CreateChatappMigrationInitiateResponseBody) *CreateChatappMigrationInitiateResponse {
	s.Body = v
	return s
}

type CreateChatappTemplateRequest struct {
	// Specifies whether to allow Facebook to automatically change the directory of the template. If you set this parameter to true, the review success rate of the template is improved. This parameter is valid only when TemplateType is set to WHATSAPP.
	//
	// example:
	//
	// true
	AllowCategoryChange *bool `json:"AllowCategoryChange,omitempty" xml:"AllowCategoryChange,omitempty"`
	// The category of the template if TemplateType is set to WHATSAPP. Valid values:
	//
	// 	- **UTILITY**: the transaction template
	//
	// 	- **MARKETING**: the marketing template
	//
	// 	- **AUTHENTICATION**: the authentication template
	//
	// The category of the template if TemplateType is set to VIBER. Valid values:
	//
	// 	- **text**: the template that contains only text
	//
	// 	- **image**: the template that contains only images
	//
	// 	- **text_image_button**: the template that contains text, images, and buttons
	//
	// 	- **text_button**: the template that contains text and buttons
	//
	// 	- **document**: the template that contains only documents
	//
	// 	- **video**: the template that contains only videos
	//
	// 	- **text_video**: the template that contains text and videos
	//
	// 	- **text_video_button**: the template that contains text, videos, and buttons
	//
	// 	- **text_image**: the template that contains text and images
	//
	// This parameter is required.
	//
	// example:
	//
	// The code of the message template.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The components of the message template.
	//
	// >  If Category is set to AUTHENTICATION, the Type sub-parameter of the Components parameter cannot be set to HEADER. If the Type sub-parameter is set to BODY or FOOTER, the Text sub-parameter of the Components parameter must be empty.
	//
	// This parameter is required.
	Components []*CreateChatappTemplateRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// > CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Validity period of authentication template message sending in WhatsApp
	//
	// > This attribute requires providing waba in advance to Alibaba operators to open the whitelist, otherwise it will result in template submission failure
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The name of the message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE: the Line message template. This type of message template will be released later.
	//
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequest) SetAllowCategoryChange(v bool) *CreateChatappTemplateRequest {
	s.AllowCategoryChange = &v
	return s
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

func (s *CreateChatappTemplateRequest) SetMessageSendTtlSeconds(v int32) *CreateChatappTemplateRequest {
	s.MessageSendTtlSeconds = &v
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
	// The note indicating that customers cannot share verification codes with others. The note is displayed in the message body. This parameter is valid only if Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to BODY for a WhatsApp message template.
	//
	// example:
	//
	// true
	AddSecretRecommendation *bool `json:"AddSecretRecommendation,omitempty" xml:"AddSecretRecommendation,omitempty"`
	// The buttons. Specify this parameter only if you set the Type sub-parameter of the Components parameter to **BUTTONS**.
	//
	// >  The following section describes the limits on the number of buttons in a WhatsApp message.
	//
	// 	- A marketing or utility WhatsApp message template supports up to 10 buttons.
	//
	// 	- A WhatsApp message template can contain only one phone call button.
	//
	// 	- A WhatsApp message template can contain up to two URL buttons.
	//
	// 	- In a WhatsApp message template, the quick reply button cannot be used together with the phone call button or the URL button.
	Buttons []*CreateChatappTemplateRequestComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description of the document.
	//
	// example:
	//
	// This is a video
	Caption *string                                        `json:"Caption,omitempty" xml:"Caption,omitempty"`
	Cards   []*CreateChatappTemplateRequestComponentsCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// The validity period of the verification code in the WhatsApp authentication template. Unit: minutes. This parameter is valid only when Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to FOOTER in a WhatsApp message template. The validity period of the verification code is displayed in the footer.
	//
	// example:
	//
	// 5
	CodeExpirationMinutes *int32 `json:"CodeExpirationMinutes,omitempty" xml:"CodeExpirationMinutes,omitempty"`
	// The length of the video in the Viber message template. Unit: seconds. Valid values: 0 to 600.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the document.
	//
	// example:
	//
	// video name
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the document attached in the Viber message template.
	//
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The type of the media resources that are included in the message. Valid values:
	//
	// 	- **TEXT**
	//
	// 	- **IMAGE**
	//
	// 	- **DOCUMENT**
	//
	// 	- **VIDEO**
	//
	// example:
	//
	// TEXT
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// true
	HasExpiration *bool `json:"HasExpiration,omitempty" xml:"HasExpiration,omitempty"`
	// The text of the message that you want to send.
	//
	// >  If Category is set to AUTHENTICATION, the Text sub-parameter of the Components parameter is empty.
	//
	// example:
	//
	// hello whatsapp
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The thumbnail URL of the video in the Viber message template.
	//
	// example:
	//
	// https://cdn.multiplymall.mobiapp.cloud/yunmall/B-LM-LMALL202207130001/20220730/d712a057-a6af-4513-bbe6-7ee57ea60983.png?x-oss-process=image/resize,w_100
	ThumbUrl *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	// The type of the component. Valid values:
	//
	// 	- **BODY**
	//
	// 	- **HEADER**
	//
	// 	- **FOOTER**
	//
	// 	- **BUTTONS**
	//
	// >
	//
	// 	- In WhatsApp message templates, a **BODY*	- component cannot exceed 1,024 characters in length, and a **HEADER*	- or **FOOTER*	- component cannot exceed 60 characters in length.
	//
	// 	- **FOOTER*	- components are not supported in Viber message templates.
	//
	// 	- In Viber message templates, media resources such as images, videos, and documents are placed in the **HEADER*	- component. If a Viber message contains both text and an image, the image is placed below the text in the message received on a device.
	//
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the media resource.
	//
	// >  We recommend that you use 800 × 800 images in Viber message templates.
	//
	// example:
	//
	// https://image.developer.aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateChatappTemplateRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponents) SetAddSecretRecommendation(v bool) *CreateChatappTemplateRequestComponents {
	s.AddSecretRecommendation = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetButtons(v []*CreateChatappTemplateRequestComponentsButtons) *CreateChatappTemplateRequestComponents {
	s.Buttons = v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetCaption(v string) *CreateChatappTemplateRequestComponents {
	s.Caption = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetCards(v []*CreateChatappTemplateRequestComponentsCards) *CreateChatappTemplateRequestComponents {
	s.Cards = v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetCodeExpirationMinutes(v int32) *CreateChatappTemplateRequestComponents {
	s.CodeExpirationMinutes = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetDuration(v int32) *CreateChatappTemplateRequestComponents {
	s.Duration = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFileName(v string) *CreateChatappTemplateRequestComponents {
	s.FileName = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFileType(v string) *CreateChatappTemplateRequestComponents {
	s.FileType = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFormat(v string) *CreateChatappTemplateRequestComponents {
	s.Format = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetHasExpiration(v bool) *CreateChatappTemplateRequestComponents {
	s.HasExpiration = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetText(v string) *CreateChatappTemplateRequestComponents {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetThumbUrl(v string) *CreateChatappTemplateRequestComponents {
	s.ThumbUrl = &v
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
	// The text of the one-tap autofill button. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// Autofill
	AutofillText *string `json:"AutofillText,omitempty" xml:"AutofillText,omitempty"`
	// example:
	//
	// 120293
	CouponCode *string `json:"CouponCode,omitempty" xml:"CouponCode,omitempty"`
	// example:
	//
	// NAVIGATE
	FlowAction *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// example:
	//
	// 28383872***
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The unsubscribe button. This parameter is valid if Category is set to MARKETING and the Type sub-parameter of the Buttons parameter is set to QUICK_REPLY for a WhatsApp message template. After you configure message sending in the ChatApp Message Service console, marketing messages are not sent to customers if they click this button.
	//
	// example:
	//
	// false
	IsOptOut *bool `json:"IsOptOut,omitempty" xml:"IsOptOut,omitempty"`
	// example:
	//
	// SIGN_UP
	NavigateScreen *string `json:"NavigateScreen,omitempty" xml:"NavigateScreen,omitempty"`
	// Deprecated
	//
	// The app package name that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// com.demo
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The phone number. This parameter is valid only if the Type sub-parameter of the Buttons parameter is set to **PHONE_NUMBER**.
	//
	// example:
	//
	// +861368897****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Deprecated
	//
	// The app signing key hash that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// wi299382
	SignatureHash *string                                                       `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
	SupportedApps []*CreateChatappTemplateRequestComponentsButtonsSupportedApps `json:"SupportedApps,omitempty" xml:"SupportedApps,omitempty" type:"Repeated"`
	// The display name of the button.
	//
	// example:
	//
	// Call Me
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the button. Valid values:
	//
	// 	- **PHONE_NUMBER**: phone call button
	//
	// 	- **URL**: URL button
	//
	// 	- **QUICK_REPLY**: quick reply button
	//
	// 	- **COPY_CODE**: copy code button if Category is set to AUTHENTICATION
	//
	// 	- **ONE_TAP**: one-tap autofill button if Category is set to AUTHENTICATION
	//
	// >
	//
	// 	- In a WhatsApp message template, a quick reply button cannot be used together with a phone call button or a URL button.
	//
	// 	- You can add a combination of two URL buttons or a combination of a URL button and a phone call button to a WhatsApp message template.
	//
	// 	- If Category is set to AUTHENTICATION for a WhatsApp message template, you can add only one button to the WhatsApp message template and you must set the Type sub-parameter of the Buttons parameter to COPY_CODE or ONE_TAP. If the Type sub-parameter of the Buttons parameter is set to COPY_CODE, the Text sub-parameter of the Buttons parameter is required. If the Type sub-parameter of the Buttons parameter is set to ONE_TAP, the Text, SignatureHash, PackageName, and AutofillText sub-parameters of the Buttons parameter are required. The value of Text is displayed if the desired app is not installed on the device. The value of Text indicates that you must manually copy the verification code.
	//
	// 	- You can add only one button to a Viber message template, and you must set the Type sub-parameter of the Buttons parameter to URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to which you are redirected when you click the URL button.
	//
	// example:
	//
	// https://example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the URL. Valid values:
	//
	// 	- **static**
	//
	// 	- **dynamic**
	//
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsButtons) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetAutofillText(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.AutofillText = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetCouponCode(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.CouponCode = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetFlowAction(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.FlowAction = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetFlowId(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.FlowId = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetIsOptOut(v bool) *CreateChatappTemplateRequestComponentsButtons {
	s.IsOptOut = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetNavigateScreen(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.NavigateScreen = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetPackageName(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.PackageName = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetPhoneNumber(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetSignatureHash(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.SignatureHash = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetSupportedApps(v []*CreateChatappTemplateRequestComponentsButtonsSupportedApps) *CreateChatappTemplateRequestComponentsButtons {
	s.SupportedApps = v
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

type CreateChatappTemplateRequestComponentsButtonsSupportedApps struct {
	// example:
	//
	// com.aliyun.pack.***
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// kdikie9398k
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsButtonsSupportedApps) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsButtonsSupportedApps) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsButtonsSupportedApps) SetPackageName(v string) *CreateChatappTemplateRequestComponentsButtonsSupportedApps {
	s.PackageName = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtonsSupportedApps) SetSignatureHash(v string) *CreateChatappTemplateRequestComponentsButtonsSupportedApps {
	s.SignatureHash = &v
	return s
}

type CreateChatappTemplateRequestComponentsCards struct {
	// This parameter is required.
	CardComponents []*CreateChatappTemplateRequestComponentsCardsCardComponents `json:"CardComponents,omitempty" xml:"CardComponents,omitempty" type:"Repeated"`
}

func (s CreateChatappTemplateRequestComponentsCards) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsCards) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsCards) SetCardComponents(v []*CreateChatappTemplateRequestComponentsCardsCardComponents) *CreateChatappTemplateRequestComponentsCards {
	s.CardComponents = v
	return s
}

type CreateChatappTemplateRequestComponentsCardsCardComponents struct {
	Buttons []*CreateChatappTemplateRequestComponentsCardsCardComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// example:
	//
	// IMAGE
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// Who is the very powerful team
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://alibaba.com/img.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsCardsCardComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsCardsCardComponents) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetButtons(v []*CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Buttons = v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetFormat(v string) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Format = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetText(v string) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetType(v string) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponents) SetUrl(v string) *CreateChatappTemplateRequestComponentsCardsCardComponents {
	s.Url = &v
	return s
}

type CreateChatappTemplateRequestComponentsCardsCardComponentsButtons struct {
	// example:
	//
	// +8613800
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// Call me
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://alibaba.com/xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetPhoneNumber(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetText(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetType(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetUrl(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Url = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons) SetUrlType(v string) *CreateChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.UrlType = &v
	return s
}

type CreateChatappTemplateShrinkRequest struct {
	// Specifies whether to allow Facebook to automatically change the directory of the template. If you set this parameter to true, the review success rate of the template is improved. This parameter is valid only when TemplateType is set to WHATSAPP.
	//
	// example:
	//
	// true
	AllowCategoryChange *bool `json:"AllowCategoryChange,omitempty" xml:"AllowCategoryChange,omitempty"`
	// The category of the template if TemplateType is set to WHATSAPP. Valid values:
	//
	// 	- **UTILITY**: the transaction template
	//
	// 	- **MARKETING**: the marketing template
	//
	// 	- **AUTHENTICATION**: the authentication template
	//
	// The category of the template if TemplateType is set to VIBER. Valid values:
	//
	// 	- **text**: the template that contains only text
	//
	// 	- **image**: the template that contains only images
	//
	// 	- **text_image_button**: the template that contains text, images, and buttons
	//
	// 	- **text_button**: the template that contains text and buttons
	//
	// 	- **document**: the template that contains only documents
	//
	// 	- **video**: the template that contains only videos
	//
	// 	- **text_video**: the template that contains text and videos
	//
	// 	- **text_video_button**: the template that contains text, videos, and buttons
	//
	// 	- **text_image**: the template that contains text and images
	//
	// This parameter is required.
	//
	// example:
	//
	// The code of the message template.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The components of the message template.
	//
	// >  If Category is set to AUTHENTICATION, the Type sub-parameter of the Components parameter cannot be set to HEADER. If the Type sub-parameter is set to BODY or FOOTER, the Text sub-parameter of the Components parameter must be empty.
	//
	// This parameter is required.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// > CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	ExampleShrink *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Validity period of authentication template message sending in WhatsApp
	//
	// > This attribute requires providing waba in advance to Alibaba operators to open the whitelist, otherwise it will result in template submission failure
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The name of the message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE: the Line message template. This type of message template will be released later.
	//
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateChatappTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateShrinkRequest) SetAllowCategoryChange(v bool) *CreateChatappTemplateShrinkRequest {
	s.AllowCategoryChange = &v
	return s
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

func (s *CreateChatappTemplateShrinkRequest) SetMessageSendTtlSeconds(v int32) *CreateChatappTemplateShrinkRequest {
	s.MessageSendTtlSeconds = &v
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {"templateCode": "****4b5c79c9432497a075bdfca36bf5"，"templateName": "hello_whatsapp"}
	Data *CreateChatappTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *CreateChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	//
	// example:
	//
	// SMS_232907****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateFlowRequest struct {
	// The categories of the Flow.
	//
	// This parameter is required.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 93994848
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the Flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) SetCategories(v []*string) *CreateFlowRequest {
	s.Categories = v
	return s
}

func (s *CreateFlowRequest) SetCustSpaceId(v string) *CreateFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateFlowRequest) SetFlowName(v string) *CreateFlowRequest {
	s.FlowName = &v
	return s
}

type CreateFlowShrinkRequest struct {
	// The categories of the Flow.
	//
	// This parameter is required.
	CategoriesShrink *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 93994848
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the Flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s CreateFlowShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowShrinkRequest) SetCategoriesShrink(v string) *CreateFlowShrinkRequest {
	s.CategoriesShrink = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetCustSpaceId(v string) *CreateFlowShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetFlowName(v string) *CreateFlowShrinkRequest {
	s.FlowName = &v
	return s
}

type CreateFlowResponseBody struct {
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) SetCode(v string) *CreateFlowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlowResponseBody) SetData(v *CreateFlowResponseBodyData) *CreateFlowResponseBody {
	s.Data = v
	return s
}

func (s *CreateFlowResponseBody) SetMessage(v string) *CreateFlowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowResponseBodyData struct {
	// The categories of the Flow.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The Flow ID.
	//
	// example:
	//
	// 333993838***
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the Flow.
	//
	// example:
	//
	// test1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s CreateFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBodyData) SetCategories(v []*string) *CreateFlowResponseBodyData {
	s.Categories = v
	return s
}

func (s *CreateFlowResponseBodyData) SetFlowId(v string) *CreateFlowResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *CreateFlowResponseBodyData) SetFlowName(v string) *CreateFlowResponseBodyData {
	s.FlowName = &v
	return s
}

type CreateFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowResponse) SetHeaders(v map[string]*string) *CreateFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowResponse) SetStatusCode(v int32) *CreateFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowResponse) SetBody(v *CreateFlowResponseBody) *CreateFlowResponse {
	s.Body = v
	return s
}

type CreatePhoneMessageQrdlRequest struct {
	// example:
	//
	// 838833
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PNG
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hello
	PrefilledMessage *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
}

func (s CreatePhoneMessageQrdlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *CreatePhoneMessageQrdlRequest) SetCustSpaceId(v string) *CreatePhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetGenerateQrImage(v string) *CreatePhoneMessageQrdlRequest {
	s.GenerateQrImage = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetPhoneNumber(v string) *CreatePhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetPrefilledMessage(v string) *CreatePhoneMessageQrdlRequest {
	s.PrefilledMessage = &v
	return s
}

type CreatePhoneMessageQrdlResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePhoneMessageQrdlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// none
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePhoneMessageQrdlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePhoneMessageQrdlResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhoneMessageQrdlResponseBody) SetCode(v string) *CreatePhoneMessageQrdlResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBody) SetData(v *CreatePhoneMessageQrdlResponseBodyData) *CreatePhoneMessageQrdlResponseBody {
	s.Data = v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBody) SetMessage(v string) *CreatePhoneMessageQrdlResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBody) SetRequestId(v string) *CreatePhoneMessageQrdlResponseBody {
	s.RequestId = &v
	return s
}

type CreatePhoneMessageQrdlResponseBodyData struct {
	// example:
	//
	// https://wa.qrdl/
	DeepLinkUrl *string `json:"DeepLinkUrl,omitempty" xml:"DeepLinkUrl,omitempty"`
	// example:
	//
	// PNG
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	// example:
	//
	// 8613800
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// Hello
	PrefilledMessage *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
	// example:
	//
	// http://img.png
	QrImageUrl *string `json:"QrImageUrl,omitempty" xml:"QrImageUrl,omitempty"`
	// example:
	//
	// D9II3***
	QrdlCode *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
}

func (s CreatePhoneMessageQrdlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePhoneMessageQrdlResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetDeepLinkUrl(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.DeepLinkUrl = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetGenerateQrImage(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.GenerateQrImage = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetPhoneNumber(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetPrefilledMessage(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.PrefilledMessage = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetQrImageUrl(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.QrImageUrl = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetQrdlCode(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.QrdlCode = &v
	return s
}

type CreatePhoneMessageQrdlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePhoneMessageQrdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePhoneMessageQrdlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePhoneMessageQrdlResponse) GoString() string {
	return s.String()
}

func (s *CreatePhoneMessageQrdlResponse) SetHeaders(v map[string]*string) *CreatePhoneMessageQrdlResponse {
	s.Headers = v
	return s
}

func (s *CreatePhoneMessageQrdlResponse) SetStatusCode(v int32) *CreatePhoneMessageQrdlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponse) SetBody(v *CreatePhoneMessageQrdlResponseBody) *CreatePhoneMessageQrdlResponse {
	s.Body = v
	return s
}

type DeleteChatappTemplateRequest struct {
	// The space ID of the user under the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// >  CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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

func (s *DeleteChatappTemplateRequest) SetLanguage(v string) *DeleteChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetTemplateCode(v string) *DeleteChatappTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetTemplateName(v string) *DeleteChatappTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetTemplateType(v string) *DeleteChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

type DeleteChatappTemplateResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *DeleteChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteFlowRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 393983883
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The Flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DeleteFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) SetCustSpaceId(v string) *DeleteFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteFlowRequest) SetFlowId(v string) *DeleteFlowRequest {
	s.FlowId = &v
	return s
}

type DeleteFlowResponseBody struct {
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponseBody) SetCode(v string) *DeleteFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlowResponseBody) SetMessage(v string) *DeleteFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlowResponseBody) SetRequestId(v string) *DeleteFlowResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponse) SetHeaders(v map[string]*string) *DeleteFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowResponse) SetStatusCode(v int32) *DeleteFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowResponse) SetBody(v *DeleteFlowResponseBody) *DeleteFlowResponse {
	s.Body = v
	return s
}

type DeletePhoneMessageQrdlRequest struct {
	// example:
	//
	// 883873773
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 29338838
	QrdlCode *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
}

func (s DeletePhoneMessageQrdlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *DeletePhoneMessageQrdlRequest) SetCustSpaceId(v string) *DeletePhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetPhoneNumber(v string) *DeletePhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetQrdlCode(v string) *DeletePhoneMessageQrdlRequest {
	s.QrdlCode = &v
	return s
}

type DeletePhoneMessageQrdlResponseBody struct {
	// example:
	//
	// OK
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePhoneMessageQrdlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePhoneMessageQrdlResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePhoneMessageQrdlResponseBody) SetCode(v string) *DeletePhoneMessageQrdlResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePhoneMessageQrdlResponseBody) SetMessage(v string) *DeletePhoneMessageQrdlResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePhoneMessageQrdlResponseBody) SetRequestId(v string) *DeletePhoneMessageQrdlResponseBody {
	s.RequestId = &v
	return s
}

type DeletePhoneMessageQrdlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePhoneMessageQrdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePhoneMessageQrdlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePhoneMessageQrdlResponse) GoString() string {
	return s.String()
}

func (s *DeletePhoneMessageQrdlResponse) SetHeaders(v map[string]*string) *DeletePhoneMessageQrdlResponse {
	s.Headers = v
	return s
}

func (s *DeletePhoneMessageQrdlResponse) SetStatusCode(v int32) *DeletePhoneMessageQrdlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePhoneMessageQrdlResponse) SetBody(v *DeletePhoneMessageQrdlResponseBody) *DeletePhoneMessageQrdlResponse {
	s.Body = v
	return s
}

type DeprecateFlowRequest struct {
	// example:
	//
	// 38877483
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Flow ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DeprecateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeprecateFlowRequest) GoString() string {
	return s.String()
}

func (s *DeprecateFlowRequest) SetCustSpaceId(v string) *DeprecateFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeprecateFlowRequest) SetFlowId(v string) *DeprecateFlowRequest {
	s.FlowId = &v
	return s
}

type DeprecateFlowResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeprecateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeprecateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeprecateFlowResponseBody) SetCode(v string) *DeprecateFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DeprecateFlowResponseBody) SetMessage(v string) *DeprecateFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DeprecateFlowResponseBody) SetRequestId(v string) *DeprecateFlowResponseBody {
	s.RequestId = &v
	return s
}

type DeprecateFlowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeprecateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeprecateFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeprecateFlowResponse) GoString() string {
	return s.String()
}

func (s *DeprecateFlowResponse) SetHeaders(v map[string]*string) *DeprecateFlowResponse {
	s.Headers = v
	return s
}

func (s *DeprecateFlowResponse) SetStatusCode(v int32) *DeprecateFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeprecateFlowResponse) SetBody(v *DeprecateFlowResponseBody) *DeprecateFlowResponse {
	s.Body = v
	return s
}

type EnableWhatsappROIMetricRequest struct {
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The verification code used to verify whether the RAM user is authorized by the independent software vendor (ISV) account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
}

func (s EnableWhatsappROIMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableWhatsappROIMetricRequest) GoString() string {
	return s.String()
}

func (s *EnableWhatsappROIMetricRequest) SetCustSpaceId(v string) *EnableWhatsappROIMetricRequest {
	s.CustSpaceId = &v
	return s
}

func (s *EnableWhatsappROIMetricRequest) SetIsvCode(v string) *EnableWhatsappROIMetricRequest {
	s.IsvCode = &v
	return s
}

type EnableWhatsappROIMetricResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// NONE
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The value OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableWhatsappROIMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableWhatsappROIMetricResponseBody) GoString() string {
	return s.String()
}

func (s *EnableWhatsappROIMetricResponseBody) SetAccessDeniedDetail(v string) *EnableWhatsappROIMetricResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *EnableWhatsappROIMetricResponseBody) SetCode(v string) *EnableWhatsappROIMetricResponseBody {
	s.Code = &v
	return s
}

func (s *EnableWhatsappROIMetricResponseBody) SetMessage(v string) *EnableWhatsappROIMetricResponseBody {
	s.Message = &v
	return s
}

func (s *EnableWhatsappROIMetricResponseBody) SetRequestId(v string) *EnableWhatsappROIMetricResponseBody {
	s.RequestId = &v
	return s
}

type EnableWhatsappROIMetricResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableWhatsappROIMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableWhatsappROIMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableWhatsappROIMetricResponse) GoString() string {
	return s.String()
}

func (s *EnableWhatsappROIMetricResponse) SetHeaders(v map[string]*string) *EnableWhatsappROIMetricResponse {
	s.Headers = v
	return s
}

func (s *EnableWhatsappROIMetricResponse) SetStatusCode(v int32) *EnableWhatsappROIMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableWhatsappROIMetricResponse) SetBody(v *EnableWhatsappROIMetricResponseBody) *EnableWhatsappROIMetricResponse {
	s.Body = v
	return s
}

type GetChatappPhoneNumberMetricRequest struct {
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1693407714687
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The metric granularity.
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The verification code used to verify whether the RAM user is authorized by the independent software vendor (ISV) account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The business phone number.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1693107714687
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetChatappPhoneNumberMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatappPhoneNumberMetricRequest) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberMetricRequest) SetCustSpaceId(v string) *GetChatappPhoneNumberMetricRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetEnd(v int64) *GetChatappPhoneNumberMetricRequest {
	s.End = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetGranularity(v string) *GetChatappPhoneNumberMetricRequest {
	s.Granularity = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetIsvCode(v string) *GetChatappPhoneNumberMetricRequest {
	s.IsvCode = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetPhoneNumber(v string) *GetChatappPhoneNumberMetricRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetStart(v int64) *GetChatappPhoneNumberMetricRequest {
	s.Start = &v
	return s
}

type GetChatappPhoneNumberMetricResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// NONE
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The value OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*GetChatappPhoneNumberMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1612C226-E271-4CFE-9F18-4066D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappPhoneNumberMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatappPhoneNumberMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetAccessDeniedDetail(v string) *GetChatappPhoneNumberMetricResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetCode(v string) *GetChatappPhoneNumberMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetData(v []*GetChatappPhoneNumberMetricResponseBodyData) *GetChatappPhoneNumberMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetMessage(v string) *GetChatappPhoneNumberMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetRequestId(v string) *GetChatappPhoneNumberMetricResponseBody {
	s.RequestId = &v
	return s
}

type GetChatappPhoneNumberMetricResponseBodyData struct {
	// The number of delivered messages.
	//
	// example:
	//
	// 5
	DeliveredCount *int32 `json:"DeliveredCount,omitempty" xml:"DeliveredCount,omitempty"`
	// The end of the time range that you queried.
	//
	// example:
	//
	// 1667196043904
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The metric granularity.
	//
	// >  Valid values:
	//
	// 	- DAILY
	//
	// 	- HALF_HOUR
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The business phone number.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The number of sent messages.
	//
	// example:
	//
	// 10
	SentCount *int32 `json:"SentCount,omitempty" xml:"SentCount,omitempty"`
	// The beginning of the time range that you queried.
	//
	// example:
	//
	// 1669619491000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetChatappPhoneNumberMetricResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChatappPhoneNumberMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetDeliveredCount(v int32) *GetChatappPhoneNumberMetricResponseBodyData {
	s.DeliveredCount = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetEnd(v int64) *GetChatappPhoneNumberMetricResponseBodyData {
	s.End = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetGranularity(v string) *GetChatappPhoneNumberMetricResponseBodyData {
	s.Granularity = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetPhoneNumber(v string) *GetChatappPhoneNumberMetricResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetSentCount(v int32) *GetChatappPhoneNumberMetricResponseBodyData {
	s.SentCount = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetStart(v int64) *GetChatappPhoneNumberMetricResponseBodyData {
	s.Start = &v
	return s
}

type GetChatappPhoneNumberMetricResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappPhoneNumberMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappPhoneNumberMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatappPhoneNumberMetricResponse) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberMetricResponse) SetHeaders(v map[string]*string) *GetChatappPhoneNumberMetricResponse {
	s.Headers = v
	return s
}

func (s *GetChatappPhoneNumberMetricResponse) SetStatusCode(v int32) *GetChatappPhoneNumberMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponse) SetBody(v *GetChatappPhoneNumberMetricResponseBody) *GetChatappPhoneNumberMetricResponse {
	s.Body = v
	return s
}

type GetChatappTemplateDetailRequest struct {
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// >  CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The independent software vendor (ISV) verification code. This parameter is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en_US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// ****4b5c79c9432497a075bdfca36bf5
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message template. Valid values:
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE (developing)
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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

func (s *GetChatappTemplateDetailRequest) SetTemplateName(v string) *GetChatappTemplateDetailRequest {
	s.TemplateName = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetTemplateType(v string) *GetChatappTemplateDetailRequest {
	s.TemplateType = &v
	return s
}

type GetChatappTemplateDetailResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code.
	//
	// 	- Example: OK. This value indicates that the request is successful.
	//
	// 	- Other codes indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// { 		"category": "ACCOUNT_UPDATE", 		"name": "account_notice", 		"language": "en_US", 		"templateCode": "744c4b5c79c9432497a075bdfca3****", 		"auditStatus": "APPROVED", 		"components": "[{\\"type\\":\\"BODY\\",\\"text\\":\\"body_text$(textVariable)\\"},{\\"type\\":\\"HEADER\\",\\"formate\\":\\"IMAGE\\",\\"url\\":\\"$(linkVariable)\\"},{\\"type\\":\\"FOOTER\\",\\"text\\":\\"footer-text\\"},{\\"type\\":\\"BUTTONS\\",\\"buttons\\":[{\\"type\\":\\"PHONE_NUMBER\\",\\"text\\":\\"phone-button-text\\",\\"phone_number\\":\\"+861388888****\\"},{\\"type\\":\\"URL\\",\\"text\\":\\"url-button-text\\",\\"url\\":\\"https://www.website.com/\\"}]}]", 		"example": "{\\"textVariable\\": \\"text\\", \\"linkVariable\\": \\"link\\"}" 	}
	Data *GetChatappTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappTemplateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBody) SetAccessDeniedDetail(v string) *GetChatappTemplateDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// 	- **pass**: The message template is approved.
	//
	// 	- **fail**: The message template is rejected.
	//
	// 	- **auditing**: The message template is being reviewed.
	//
	// 	- **unaudit**: The review is suspended.
	//
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The category of the template when the returned value of TemplateType is WHATSAPP. Valid values:
	//
	// 	- **UTILITY**: a transactional template
	//
	// 	- **MARKETING**: a marketing template
	//
	// 	- **AUTHENTICATION**: an identity authentication template
	//
	// The category of the template when the returned value of the TemplateType parameter is VIBER. Valid values:
	//
	// 	- **text**: a template that contains only text
	//
	// 	- **image**: a template that contains only images
	//
	// 	- **text_image_button**: a template that contains text, images, and buttons
	//
	// 	- **text_button**: a template that contains text and buttons
	//
	// 	- **document**: a template that contains only files
	//
	// 	- **video**: a template that contains only videos
	//
	// 	- **text_video**: a template that contains text and videos
	//
	// 	- **text_video_button**: a template that contains text, videos, and buttons
	//
	// 	- **text_image**: a template that contains text and images
	//
	// > If Category is set to text_video_button, users cannot open a web page by clicking the button. Users can open only the video in the message. In this case, you do not need to specify the Url parameter for the URL button in the template.
	//
	// example:
	//
	// TRANSACTIONAL
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The components of the message template.
	Components []*GetChatappTemplateDetailResponseBodyDataComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The examples of variables.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en_US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The validity period of the WhatsApp authentication message.
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The quality of the template.
	//
	// example:
	//
	// GREEN
	QualityScore *string `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	// The reason why the template was rejected.
	//
	// example:
	//
	// None
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The type of the message template. Valid values:
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE (developing)
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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

func (s *GetChatappTemplateDetailResponseBodyData) SetMessageSendTtlSeconds(v int32) *GetChatappTemplateDetailResponseBodyData {
	s.MessageSendTtlSeconds = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetName(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetQualityScore(v string) *GetChatappTemplateDetailResponseBodyData {
	s.QualityScore = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetReason(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetTemplateCode(v string) *GetChatappTemplateDetailResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetTemplateType(v string) *GetChatappTemplateDetailResponseBodyData {
	s.TemplateType = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponents struct {
	// The note indicating that customers cannot share verification codes with others. The note is displayed in the message body. This parameter is valid if Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to BODY in a WhatsApp message template.
	//
	// example:
	//
	// false
	AddSecretRecommendation *bool `json:"AddSecretRecommendation,omitempty" xml:"AddSecretRecommendation,omitempty"`
	// The buttons. Specify this parameter only if you set the Type sub-parameter of the Components parameter to **BUTTONS**. Before you specify this parameter, the format of the value must be changed from JSON to String.
	Buttons []*GetChatappTemplateDetailResponseBodyDataComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description of the document.
	//
	// example:
	//
	// The new file has been uploaded.
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The carousel cards.
	Cards []*GetChatappTemplateDetailResponseBodyDataComponentsCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// The validity period of the verification code in the WhatsApp authentication template. Unit: minutes. This parameter is valid only when Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to FOOTER in a WhatsApp message template. The validity period of the verification code is displayed in the footer.
	//
	// example:
	//
	// 5
	CodeExpirationMinutes *int32 `json:"CodeExpirationMinutes,omitempty" xml:"CodeExpirationMinutes,omitempty"`
	// The length of the video in the Viber message template. Valid values: 0 to 600. Unit: seconds.
	//
	// example:
	//
	// 50
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the document.
	//
	// example:
	//
	// Express file
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the document attached in the Viber message template.
	//
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The format.
	//
	// example:
	//
	// TEXT
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The latitude of the location.
	//
	// example:
	//
	// 28.001
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// The address of the location.
	//
	// example:
	//
	// Hangzhou
	LocationAddress *string `json:"LocationAddress,omitempty" xml:"LocationAddress,omitempty"`
	// The name of the location.
	//
	// example:
	//
	// Hangzhou
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// The longitude of the location.
	//
	// example:
	//
	// 120.002
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// The variable when the coupon code expires in the limited-time offer template.
	//
	// example:
	//
	// $(offerExpirationTimeMs)
	OfferExpirationTimeMs *string `json:"OfferExpirationTimeMs,omitempty" xml:"OfferExpirationTimeMs,omitempty"`
	// The text of the message that you want to send.
	//
	// example:
	//
	// Hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The thumbnail URL of the video in the Viber message template.
	//
	// example:
	//
	// https://img.png
	ThumbUrl *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	// The type of the component. Valid values:
	//
	// 	- **BODY**
	//
	// 	- **HEADER**
	//
	// 	- **FOOTER**
	//
	// 	- **BUTTONS**
	//
	// >
	//
	// 	- The following limits apply to components in WhatsApp message templates: A **BODY*	- component cannot exceed 1,024 characters in length. A **HEADER*	- or **FOOTER*	- component cannot exceed 60 characters in length.
	//
	// 	- **FOOTER*	- components are not supported in Viber message templates.
	//
	// 	- In a Viber message template, media resources such as images, videos, and documents are placed in the **HEADER*	- component. If a Viber message contains both text and an image, the image is placed below the text in the message received on a device.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the media resource.
	//
	// example:
	//
	// https://image.developer.aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Indicates whether the coupon code will expire in the limited-time offer template.
	//
	// example:
	//
	// true
	HasExpiration *bool `json:"hasExpiration,omitempty" xml:"hasExpiration,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponents) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponents) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetAddSecretRecommendation(v bool) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.AddSecretRecommendation = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetButtons(v []*GetChatappTemplateDetailResponseBodyDataComponentsButtons) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Buttons = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetCaption(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Caption = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetCards(v []*GetChatappTemplateDetailResponseBodyDataComponentsCards) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Cards = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetCodeExpirationMinutes(v int32) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.CodeExpirationMinutes = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetDuration(v int32) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Duration = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFileName(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.FileName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFileType(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.FileType = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFormat(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Format = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetLatitude(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Latitude = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetLocationAddress(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.LocationAddress = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetLocationName(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.LocationName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetLongitude(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Longitude = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetOfferExpirationTimeMs(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.OfferExpirationTimeMs = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetThumbUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.ThumbUrl = &v
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

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetHasExpiration(v bool) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.HasExpiration = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponentsButtons struct {
	// The text of the one-tap autofill button. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP in a WhatsApp message template.
	//
	// example:
	//
	// Autofill
	AutofillText *string `json:"AutofillText,omitempty" xml:"AutofillText,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 202039ksjs
	CouponCode *string `json:"CouponCode,omitempty" xml:"CouponCode,omitempty"`
	// The extended fields.
	ExtendAttrs *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs `json:"ExtendAttrs,omitempty" xml:"ExtendAttrs,omitempty" type:"Struct"`
	// flow 数据事件类型
	//
	// 取值范围：NAVIGATE/DATA_EXCHANGE
	//
	// example:
	//
	// NAVIGATE
	FlowAction *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// Flow ID
	//
	// example:
	//
	// 3838292983
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The unsubscribe button. This parameter is valid if Category is set to MARKETING and the Type sub-parameter of the Buttons parameter is set to QUICK_REPLY in a WhatsApp message template. After you configure message sending in the ChatApp Message Service console, marketing messages will not be sent to customers if they click this button.
	//
	// example:
	//
	// false
	IsOptOut *bool `json:"IsOptOut,omitempty" xml:"IsOptOut,omitempty"`
	// 跳转屏
	//
	// 在FlowAction=NAVIGATE时必填
	//
	// example:
	//
	// DETAILS
	NavigateScreen *string `json:"NavigateScreen,omitempty" xml:"NavigateScreen,omitempty"`
	// The app package name that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP in a WhatsApp message template.
	//
	// example:
	//
	// com.aliyun
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The phone number. This parameter is valid only if the Type sub-parameter of the Buttons parameter is set to **PHONE_NUMBER**.
	//
	// example:
	//
	// 861398745****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The app signing key hash that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP in a WhatsApp message template.
	//
	// example:
	//
	// 2993839
	SignatureHash *string                                                                   `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
	SupportedApps []*GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps `json:"SupportedApps,omitempty" xml:"SupportedApps,omitempty" type:"Repeated"`
	// The display name of the button.
	//
	// example:
	//
	// Call
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the button. Valid values:
	//
	// 	- **PHONE_NUMBER**: phone call button
	//
	// 	- **URL**: URL button
	//
	// 	- **QUICK_REPLY**: quick reply button
	//
	// >
	//
	// 	- In a message template, a quick reply button cannot be used together with a phone call button or a URL button.
	//
	// 	- You can add a combination of two URL buttons or a combination of a URL button and a phone call button to a message template.
	//
	// 	- You can add only one button to a Viber message template, and you must set the Type sub-parameter of the Buttons parameter to URL.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to which you are redirected when you click the URL button.
	//
	// example:
	//
	// https://example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the URL. Valid values:
	//
	// 	- **static**
	//
	// 	- **dynamic**
	//
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtons) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetAutofillText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.AutofillText = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetCouponCode(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.CouponCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetExtendAttrs(v *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.ExtendAttrs = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetFlowAction(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.FlowAction = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetFlowId(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.FlowId = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetIsOptOut(v bool) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.IsOptOut = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetNavigateScreen(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.NavigateScreen = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetPackageName(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.PackageName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetPhoneNumber(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetSignatureHash(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.SignatureHash = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetSupportedApps(v []*GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.SupportedApps = v
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

type GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs struct {
	// The event type.
	//
	// example:
	//
	// nextCard
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The intent code.
	//
	// example:
	//
	// test
	IntentCode *string `json:"IntentCode,omitempty" xml:"IntentCode,omitempty"`
	// The language of the next template.
	//
	// example:
	//
	// en
	NextLanguageCode *string `json:"NextLanguageCode,omitempty" xml:"NextLanguageCode,omitempty"`
	// The code of the next template.
	//
	// example:
	//
	// 20939920093993
	NextTemplateCode *string `json:"NextTemplateCode,omitempty" xml:"NextTemplateCode,omitempty"`
	// The name of the next template.
	//
	// example:
	//
	// abc
	NextTemplateName *string `json:"NextTemplateName,omitempty" xml:"NextTemplateName,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetAction(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.Action = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetIntentCode(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.IntentCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetNextLanguageCode(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.NextLanguageCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetNextTemplateCode(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.NextTemplateCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs) SetNextTemplateName(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsExtendAttrs {
	s.NextTemplateName = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps struct {
	PackageName   *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) SetPackageName(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps {
	s.PackageName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps) SetSignatureHash(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtonsSupportedApps {
	s.SignatureHash = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponentsCards struct {
	// The card components.
	CardComponents []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents `json:"CardComponents,omitempty" xml:"CardComponents,omitempty" type:"Repeated"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCards) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCards) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCards) SetCardComponents(v []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) *GetChatappTemplateDetailResponseBodyDataComponentsCards {
	s.CardComponents = v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents struct {
	// The card buttons.
	Buttons []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The type of the header in the carousel template. The header can only be an image or a video. The headers of all cards must be the same.
	//
	// example:
	//
	// HEADER
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The card text.
	//
	// example:
	//
	// Body
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the component.
	//
	// example:
	//
	// HEADER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL address.
	//
	// example:
	//
	// https://aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetButtons(v []*GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Buttons = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetFormat(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Format = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponents {
	s.Url = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons struct {
	// The phone number.
	//
	// example:
	//
	// +86138000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The button content.
	//
	// example:
	//
	// Button text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the button in the carousel template. Valid values: URL, PHONE_NUMBER, and QUICK_REQLY.
	//
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to be accessed when you click the URL button.
	//
	// example:
	//
	// https://aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the URL. Valid values: static and dynamic.
	//
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetPhoneNumber(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.Url = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons) SetUrlType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsCardsCardComponentsButtons {
	s.UrlType = &v
	return s
}

type GetChatappTemplateDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetChatappTemplateMetricRequest struct {
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1693407714687
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The metric granularity.
	//
	// >  Valid values:
	//
	// 	- DAILY
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The verification code used to verify whether the RAM user is authorized by the independent software vendor (ISV) account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The template language.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1693107714687
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The template code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca36bf5
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template type. Valid values: WHATSAPP and VIBER. If you do not specify this parameter, WHATSAPP is used by default.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetChatappTemplateMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateMetricRequest) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricRequest) SetCustSpaceId(v string) *GetChatappTemplateMetricRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetEnd(v int64) *GetChatappTemplateMetricRequest {
	s.End = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetGranularity(v string) *GetChatappTemplateMetricRequest {
	s.Granularity = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetIsvCode(v string) *GetChatappTemplateMetricRequest {
	s.IsvCode = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetLanguage(v string) *GetChatappTemplateMetricRequest {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetStart(v int64) *GetChatappTemplateMetricRequest {
	s.Start = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetTemplateCode(v string) *GetChatappTemplateMetricRequest {
	s.TemplateCode = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetTemplateType(v string) *GetChatappTemplateMetricRequest {
	s.TemplateType = &v
	return s
}

type GetChatappTemplateMetricResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The value OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*GetChatappTemplateMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappTemplateMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricResponseBody) SetAccessDeniedDetail(v string) *GetChatappTemplateMetricResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) SetCode(v string) *GetChatappTemplateMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) SetData(v []*GetChatappTemplateMetricResponseBodyData) *GetChatappTemplateMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) SetMessage(v string) *GetChatappTemplateMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) SetRequestId(v string) *GetChatappTemplateMetricResponseBody {
	s.RequestId = &v
	return s
}

type GetChatappTemplateMetricResponseBodyData struct {
	// The statistics on button clicks.
	Cliented []*GetChatappTemplateMetricResponseBodyDataCliented `json:"Cliented,omitempty" xml:"Cliented,omitempty" type:"Repeated"`
	// The number of delivered messages.
	//
	// example:
	//
	// 6
	DeliveredCount *int32 `json:"DeliveredCount,omitempty" xml:"DeliveredCount,omitempty"`
	// The end of the time range you queried.
	//
	// example:
	//
	// 1668138331485
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The template language.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of read messages.
	//
	// example:
	//
	// 3
	ReadCount *int32 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// The number of sent messages.
	//
	// example:
	//
	// 10
	SentCount *int32 `json:"SentCount,omitempty" xml:"SentCount,omitempty"`
	// The beginning of the time range you queried.
	//
	// example:
	//
	// 1673919240001
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The template code.
	//
	// example:
	//
	// 83837774838*****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetChatappTemplateMetricResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricResponseBodyData) SetCliented(v []*GetChatappTemplateMetricResponseBodyDataCliented) *GetChatappTemplateMetricResponseBodyData {
	s.Cliented = v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetDeliveredCount(v int32) *GetChatappTemplateMetricResponseBodyData {
	s.DeliveredCount = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetEnd(v int64) *GetChatappTemplateMetricResponseBodyData {
	s.End = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetLanguage(v string) *GetChatappTemplateMetricResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetReadCount(v int32) *GetChatappTemplateMetricResponseBodyData {
	s.ReadCount = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetSentCount(v int32) *GetChatappTemplateMetricResponseBodyData {
	s.SentCount = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetStart(v int64) *GetChatappTemplateMetricResponseBodyData {
	s.Start = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetTemplateCode(v string) *GetChatappTemplateMetricResponseBodyData {
	s.TemplateCode = &v
	return s
}

type GetChatappTemplateMetricResponseBodyDataCliented struct {
	// The text on the button.
	//
	// example:
	//
	// Open url
	ButtonContent *string `json:"ButtonContent,omitempty" xml:"ButtonContent,omitempty"`
	// The number of clicks.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The button type.
	//
	// >  Valid values:
	//
	// 	- url_button
	//
	// 	- quick_relpy_button
	//
	// 	- phone_number_button
	//
	// example:
	//
	// quick_reply_button
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetChatappTemplateMetricResponseBodyDataCliented) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateMetricResponseBodyDataCliented) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) SetButtonContent(v string) *GetChatappTemplateMetricResponseBodyDataCliented {
	s.ButtonContent = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) SetCount(v int32) *GetChatappTemplateMetricResponseBodyDataCliented {
	s.Count = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) SetType(v string) *GetChatappTemplateMetricResponseBodyDataCliented {
	s.Type = &v
	return s
}

type GetChatappTemplateMetricResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappTemplateMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappTemplateMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateMetricResponse) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricResponse) SetHeaders(v map[string]*string) *GetChatappTemplateMetricResponse {
	s.Headers = v
	return s
}

func (s *GetChatappTemplateMetricResponse) SetStatusCode(v int32) *GetChatappTemplateMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappTemplateMetricResponse) SetBody(v *GetChatappTemplateMetricResponseBody) *GetChatappTemplateMetricResponse {
	s.Body = v
	return s
}

type GetChatappUploadAuthorizationRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
}

func (s GetChatappUploadAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatappUploadAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *GetChatappUploadAuthorizationRequest) SetCustSpaceId(v string) *GetChatappUploadAuthorizationRequest {
	s.CustSpaceId = &v
	return s
}

type GetChatappUploadAuthorizationResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetChatappUploadAuthorizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappUploadAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatappUploadAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappUploadAuthorizationResponseBody) SetAccessDeniedDetail(v string) *GetChatappUploadAuthorizationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) SetCode(v string) *GetChatappUploadAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) SetData(v *GetChatappUploadAuthorizationResponseBodyData) *GetChatappUploadAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) SetMessage(v string) *GetChatappUploadAuthorizationResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) SetRequestId(v string) *GetChatappUploadAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

type GetChatappUploadAuthorizationResponseBodyData struct {
	// The AccessKey ID that is used to authorize a user to upload a file to Object Storage Service (OSS).
	//
	// example:
	//
	// 2skeuurfj****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret that is used to authorize a user to upload a file to OSS.
	//
	// example:
	//
	// skdkdukeuuuu****
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The name of the bucket to which a file is uploaded in OSS.
	//
	// example:
	//
	// oss
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The directory to which a file is uploaded in OSS.
	//
	// example:
	//
	// 1000102939
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The address of the OSS server to which a file is uploaded.
	//
	// example:
	//
	// https://oss.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 3600
	Expire *int32 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The security token.
	//
	// example:
	//
	// dkdieiii**
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetChatappUploadAuthorizationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChatappUploadAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetAccessKeyId(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetAccessKeySecret(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetBucketName(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetDir(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.Dir = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetEndPoint(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetExpire(v int32) *GetChatappUploadAuthorizationResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetSecurityToken(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.SecurityToken = &v
	return s
}

type GetChatappUploadAuthorizationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappUploadAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappUploadAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatappUploadAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *GetChatappUploadAuthorizationResponse) SetHeaders(v map[string]*string) *GetChatappUploadAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *GetChatappUploadAuthorizationResponse) SetStatusCode(v int32) *GetChatappUploadAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponse) SetBody(v *GetChatappUploadAuthorizationResponseBody) *GetChatappUploadAuthorizationResponse {
	s.Body = v
	return s
}

type GetChatappVerifyCodeRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 229393838*****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The language.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh_CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The method to obtain the verification code. Valid values: SMS and VOICE.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetChatappVerifyCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatappVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *GetChatappVerifyCodeRequest) SetCustSpaceId(v string) *GetChatappVerifyCodeRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetLocale(v string) *GetChatappVerifyCodeRequest {
	s.Locale = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetMethod(v string) *GetChatappVerifyCodeRequest {
	s.Method = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetPhoneNumber(v string) *GetChatappVerifyCodeRequest {
	s.PhoneNumber = &v
	return s
}

type GetChatappVerifyCodeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// None.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1612C226-E271-4CFE-9F18-4066D550F91B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappVerifyCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatappVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *GetChatappVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappVerifyCodeResponseBody) SetCode(v string) *GetChatappVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappVerifyCodeResponseBody) SetMessage(v string) *GetChatappVerifyCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappVerifyCodeResponseBody) SetRequestId(v string) *GetChatappVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

type GetChatappVerifyCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappVerifyCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatappVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *GetChatappVerifyCodeResponse) SetHeaders(v map[string]*string) *GetChatappVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *GetChatappVerifyCodeResponse) SetStatusCode(v int32) *GetChatappVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappVerifyCodeResponse) SetBody(v *GetChatappVerifyCodeResponseBody) *GetChatappVerifyCodeResponse {
	s.Body = v
	return s
}

type GetCommerceSettingRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1380000****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetCommerceSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCommerceSettingRequest) GoString() string {
	return s.String()
}

func (s *GetCommerceSettingRequest) SetCustSpaceId(v string) *GetCommerceSettingRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetCommerceSettingRequest) SetPhoneNumber(v string) *GetCommerceSettingRequest {
	s.PhoneNumber = &v
	return s
}

type GetCommerceSettingResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetCommerceSettingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCommerceSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCommerceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommerceSettingResponseBody) SetCode(v string) *GetCommerceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *GetCommerceSettingResponseBody) SetData(v *GetCommerceSettingResponseBodyData) *GetCommerceSettingResponseBody {
	s.Data = v
	return s
}

func (s *GetCommerceSettingResponseBody) SetMessage(v string) *GetCommerceSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetCommerceSettingResponseBody) SetRequestId(v string) *GetCommerceSettingResponseBody {
	s.RequestId = &v
	return s
}

type GetCommerceSettingResponseBodyData struct {
	// Indicates whether the shopping cart button is displayed.
	//
	// example:
	//
	// false
	CartEnable *bool `json:"CartEnable,omitempty" xml:"CartEnable,omitempty"`
	// Indicates whether the catalog button is displayed.
	//
	// example:
	//
	// false
	CatalogVisible *bool `json:"CatalogVisible,omitempty" xml:"CatalogVisible,omitempty"`
}

func (s GetCommerceSettingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCommerceSettingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommerceSettingResponseBodyData) SetCartEnable(v bool) *GetCommerceSettingResponseBodyData {
	s.CartEnable = &v
	return s
}

func (s *GetCommerceSettingResponseBodyData) SetCatalogVisible(v bool) *GetCommerceSettingResponseBodyData {
	s.CatalogVisible = &v
	return s
}

type GetCommerceSettingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommerceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommerceSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCommerceSettingResponse) GoString() string {
	return s.String()
}

func (s *GetCommerceSettingResponse) SetHeaders(v map[string]*string) *GetCommerceSettingResponse {
	s.Headers = v
	return s
}

func (s *GetCommerceSettingResponse) SetStatusCode(v int32) *GetCommerceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommerceSettingResponse) SetBody(v *GetCommerceSettingResponseBody) *GetCommerceSettingResponse {
	s.Body = v
	return s
}

type GetFlowRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 99384883
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The Flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowRequest) GoString() string {
	return s.String()
}

func (s *GetFlowRequest) SetCustSpaceId(v string) *GetFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetFlowRequest) SetFlowId(v string) *GetFlowRequest {
	s.FlowId = &v
	return s
}

type GetFlowResponseBody struct {
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBody) SetCode(v string) *GetFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetFlowResponseBody) SetData(v *GetFlowResponseBodyData) *GetFlowResponseBody {
	s.Data = v
	return s
}

func (s *GetFlowResponseBody) SetMessage(v string) *GetFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetFlowResponseBody) SetRequestId(v string) *GetFlowResponseBody {
	s.RequestId = &v
	return s
}

type GetFlowResponseBodyData struct {
	// The categories of the Flow.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The API version.
	//
	// example:
	//
	// 3.0
	DataApiVersion *string `json:"DataApiVersion,omitempty" xml:"DataApiVersion,omitempty"`
	// The Flow ID.
	//
	// example:
	//
	// flow_id_arms
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the Flow.
	//
	// example:
	//
	// dnjn
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The JSON version.
	//
	// example:
	//
	// 2.1
	JSONVersion *string `json:"JSONVersion,omitempty" xml:"JSONVersion,omitempty"`
	// The temporary preview URL.
	//
	// example:
	//
	// https://pre-url
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// The time when the preview URL expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1700617436633
	PreviewUrlExpires *int64 `json:"PreviewUrlExpires,omitempty" xml:"PreviewUrlExpires,omitempty"`
	// The state of the Flow.
	//
	// Valid values:
	//
	// 	- PUBLISHED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DRAFT
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DEPRECATED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBodyData) SetCategories(v []*string) *GetFlowResponseBodyData {
	s.Categories = v
	return s
}

func (s *GetFlowResponseBodyData) SetDataApiVersion(v string) *GetFlowResponseBodyData {
	s.DataApiVersion = &v
	return s
}

func (s *GetFlowResponseBodyData) SetFlowId(v string) *GetFlowResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *GetFlowResponseBodyData) SetFlowName(v string) *GetFlowResponseBodyData {
	s.FlowName = &v
	return s
}

func (s *GetFlowResponseBodyData) SetJSONVersion(v string) *GetFlowResponseBodyData {
	s.JSONVersion = &v
	return s
}

func (s *GetFlowResponseBodyData) SetPreviewUrl(v string) *GetFlowResponseBodyData {
	s.PreviewUrl = &v
	return s
}

func (s *GetFlowResponseBodyData) SetPreviewUrlExpires(v int64) *GetFlowResponseBodyData {
	s.PreviewUrlExpires = &v
	return s
}

func (s *GetFlowResponseBodyData) SetStatus(v string) *GetFlowResponseBodyData {
	s.Status = &v
	return s
}

type GetFlowResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowResponse) GoString() string {
	return s.String()
}

func (s *GetFlowResponse) SetHeaders(v map[string]*string) *GetFlowResponse {
	s.Headers = v
	return s
}

func (s *GetFlowResponse) SetStatusCode(v int32) *GetFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowResponse) SetBody(v *GetFlowResponseBody) *GetFlowResponse {
	s.Body = v
	return s
}

type GetFlowJSONAssestRequest struct {
	// example:
	//
	// 83883873
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Flow ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetFlowJSONAssestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowJSONAssestRequest) GoString() string {
	return s.String()
}

func (s *GetFlowJSONAssestRequest) SetCustSpaceId(v string) *GetFlowJSONAssestRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetFlowJSONAssestRequest) SetFlowId(v string) *GetFlowJSONAssestRequest {
	s.FlowId = &v
	return s
}

type GetFlowJSONAssestResponseBody struct {
	// example:
	//
	// OK
	Code    *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetFlowJSONAssestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFlowJSONAssestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowJSONAssestResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowJSONAssestResponseBody) SetCode(v string) *GetFlowJSONAssestResponseBody {
	s.Code = &v
	return s
}

func (s *GetFlowJSONAssestResponseBody) SetData(v *GetFlowJSONAssestResponseBodyData) *GetFlowJSONAssestResponseBody {
	s.Data = v
	return s
}

func (s *GetFlowJSONAssestResponseBody) SetMessage(v string) *GetFlowJSONAssestResponseBody {
	s.Message = &v
	return s
}

func (s *GetFlowJSONAssestResponseBody) SetRequestId(v string) *GetFlowJSONAssestResponseBody {
	s.RequestId = &v
	return s
}

type GetFlowJSONAssestResponseBodyData struct {
	// example:
	//
	// https://url.com/json.json
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// flow ID。
	//
	// example:
	//
	// flow_id_arms
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetFlowJSONAssestResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFlowJSONAssestResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowJSONAssestResponseBodyData) SetFilePath(v string) *GetFlowJSONAssestResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetFlowJSONAssestResponseBodyData) SetFlowId(v string) *GetFlowJSONAssestResponseBodyData {
	s.FlowId = &v
	return s
}

type GetFlowJSONAssestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowJSONAssestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowJSONAssestResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowJSONAssestResponse) GoString() string {
	return s.String()
}

func (s *GetFlowJSONAssestResponse) SetHeaders(v map[string]*string) *GetFlowJSONAssestResponse {
	s.Headers = v
	return s
}

func (s *GetFlowJSONAssestResponse) SetStatusCode(v int32) *GetFlowJSONAssestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowJSONAssestResponse) SetBody(v *GetFlowJSONAssestResponseBody) *GetFlowJSONAssestResponse {
	s.Body = v
	return s
}

type GetFlowPreviewUrlRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 939399383
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The Flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetFlowPreviewUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowPreviewUrlRequest) GoString() string {
	return s.String()
}

func (s *GetFlowPreviewUrlRequest) SetCustSpaceId(v string) *GetFlowPreviewUrlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetFlowPreviewUrlRequest) SetFlowId(v string) *GetFlowPreviewUrlRequest {
	s.FlowId = &v
	return s
}

type GetFlowPreviewUrlResponseBody struct {
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetFlowPreviewUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFlowPreviewUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowPreviewUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowPreviewUrlResponseBody) SetCode(v string) *GetFlowPreviewUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBody) SetData(v *GetFlowPreviewUrlResponseBodyData) *GetFlowPreviewUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetFlowPreviewUrlResponseBody) SetMessage(v string) *GetFlowPreviewUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBody) SetRequestId(v string) *GetFlowPreviewUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetFlowPreviewUrlResponseBodyData struct {
	// The Flow ID.
	//
	// example:
	//
	// 6dd31e1b7cc940fc99e293d9952b5b79
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The temporary preview URL.
	//
	// example:
	//
	// https://url
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// The time when the preview URL expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1700617436633
	PreviewUrlExpires *int64 `json:"PreviewUrlExpires,omitempty" xml:"PreviewUrlExpires,omitempty"`
}

func (s GetFlowPreviewUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFlowPreviewUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowPreviewUrlResponseBodyData) SetFlowId(v string) *GetFlowPreviewUrlResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBodyData) SetPreviewUrl(v string) *GetFlowPreviewUrlResponseBodyData {
	s.PreviewUrl = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBodyData) SetPreviewUrlExpires(v int64) *GetFlowPreviewUrlResponseBodyData {
	s.PreviewUrlExpires = &v
	return s
}

type GetFlowPreviewUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowPreviewUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowPreviewUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowPreviewUrlResponse) GoString() string {
	return s.String()
}

func (s *GetFlowPreviewUrlResponse) SetHeaders(v map[string]*string) *GetFlowPreviewUrlResponse {
	s.Headers = v
	return s
}

func (s *GetFlowPreviewUrlResponse) SetStatusCode(v int32) *GetFlowPreviewUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowPreviewUrlResponse) SetBody(v *GetFlowPreviewUrlResponseBody) *GetFlowPreviewUrlResponse {
	s.Body = v
	return s
}

type GetMigrationVerifyCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh_CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sms
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetMigrationVerifyCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationVerifyCodeRequest) SetCustSpaceId(v string) *GetMigrationVerifyCodeRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetLocale(v string) *GetMigrationVerifyCodeRequest {
	s.Locale = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetMethod(v string) *GetMigrationVerifyCodeRequest {
	s.Method = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetPhoneNumber(v string) *GetMigrationVerifyCodeRequest {
	s.PhoneNumber = &v
	return s
}

type GetMigrationVerifyCodeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetMigrationVerifyCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMigrationVerifyCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *GetMigrationVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) SetCode(v string) *GetMigrationVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) SetData(v *GetMigrationVerifyCodeResponseBodyData) *GetMigrationVerifyCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) SetMessage(v string) *GetMigrationVerifyCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) SetRequestId(v string) *GetMigrationVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

type GetMigrationVerifyCodeResponseBodyData struct {
	// example:
	//
	// 82828893332
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetMigrationVerifyCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationVerifyCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMigrationVerifyCodeResponseBodyData) SetId(v string) *GetMigrationVerifyCodeResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBodyData) SetPhoneNumber(v string) *GetMigrationVerifyCodeResponseBodyData {
	s.PhoneNumber = &v
	return s
}

type GetMigrationVerifyCodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMigrationVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMigrationVerifyCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationVerifyCodeResponse) SetHeaders(v map[string]*string) *GetMigrationVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationVerifyCodeResponse) SetStatusCode(v int32) *GetMigrationVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrationVerifyCodeResponse) SetBody(v *GetMigrationVerifyCodeResponseBody) *GetMigrationVerifyCodeResponse {
	s.Body = v
	return s
}

type GetPermissionByCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 393847477
	CustSpaceId *string   `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	Permissions []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
}

func (s GetPermissionByCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionByCodeRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionByCodeRequest) SetCode(v string) *GetPermissionByCodeRequest {
	s.Code = &v
	return s
}

func (s *GetPermissionByCodeRequest) SetCustSpaceId(v string) *GetPermissionByCodeRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPermissionByCodeRequest) SetPermissions(v []*string) *GetPermissionByCodeRequest {
	s.Permissions = v
	return s
}

type GetPermissionByCodeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 393847477
	CustSpaceId       *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	PermissionsShrink *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
}

func (s GetPermissionByCodeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionByCodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionByCodeShrinkRequest) SetCode(v string) *GetPermissionByCodeShrinkRequest {
	s.Code = &v
	return s
}

func (s *GetPermissionByCodeShrinkRequest) SetCustSpaceId(v string) *GetPermissionByCodeShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPermissionByCodeShrinkRequest) SetPermissionsShrink(v string) *GetPermissionByCodeShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

type GetPermissionByCodeResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPermissionByCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionByCodeResponseBody) SetCode(v string) *GetPermissionByCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetPermissionByCodeResponseBody) SetMessage(v string) *GetPermissionByCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetPermissionByCodeResponseBody) SetRequestId(v string) *GetPermissionByCodeResponseBody {
	s.RequestId = &v
	return s
}

type GetPermissionByCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermissionByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermissionByCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionByCodeResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionByCodeResponse) SetHeaders(v map[string]*string) *GetPermissionByCodeResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionByCodeResponse) SetStatusCode(v int32) *GetPermissionByCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionByCodeResponse) SetBody(v *GetPermissionByCodeResponseBody) *GetPermissionByCodeResponse {
	s.Body = v
	return s
}

type GetPhoneEncryptionPublicKeyRequest struct {
	// example:
	//
	// 393838848
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetPhoneEncryptionPublicKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneEncryptionPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneEncryptionPublicKeyRequest) SetCustSpaceId(v string) *GetPhoneEncryptionPublicKeyRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyRequest) SetPhoneNumber(v string) *GetPhoneEncryptionPublicKeyRequest {
	s.PhoneNumber = &v
	return s
}

type GetPhoneEncryptionPublicKeyResponseBody struct {
	// example:
	//
	// OK
	Code    *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetPhoneEncryptionPublicKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneEncryptionPublicKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneEncryptionPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetCode(v string) *GetPhoneEncryptionPublicKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetData(v *GetPhoneEncryptionPublicKeyResponseBodyData) *GetPhoneEncryptionPublicKeyResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetMessage(v string) *GetPhoneEncryptionPublicKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetRequestId(v string) *GetPhoneEncryptionPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

type GetPhoneEncryptionPublicKeyResponseBodyData struct {
	// example:
	//
	// -----BEGIN PUBLIC KEY-----
	//
	// AAA
	//
	// BBB
	//
	// CCC
	//
	// DDD
	//
	// EEE
	//
	// FFF
	//
	// GGG
	//
	// -----END PUBLIC KEY-----
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" xml:"EncryptionPublicKey,omitempty"`
	// example:
	//
	// VALID
	EncryptionPublicKeyStatus *string `json:"EncryptionPublicKeyStatus,omitempty" xml:"EncryptionPublicKeyStatus,omitempty"`
	// example:
	//
	// 86138000**
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetPhoneEncryptionPublicKeyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneEncryptionPublicKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) SetEncryptionPublicKey(v string) *GetPhoneEncryptionPublicKeyResponseBodyData {
	s.EncryptionPublicKey = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) SetEncryptionPublicKeyStatus(v string) *GetPhoneEncryptionPublicKeyResponseBodyData {
	s.EncryptionPublicKeyStatus = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) SetPhoneNumber(v string) *GetPhoneEncryptionPublicKeyResponseBodyData {
	s.PhoneNumber = &v
	return s
}

type GetPhoneEncryptionPublicKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneEncryptionPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneEncryptionPublicKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneEncryptionPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneEncryptionPublicKeyResponse) SetHeaders(v map[string]*string) *GetPhoneEncryptionPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponse) SetStatusCode(v int32) *GetPhoneEncryptionPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponse) SetBody(v *GetPhoneEncryptionPublicKeyResponseBody) *GetPhoneEncryptionPublicKeyResponse {
	s.Body = v
	return s
}

type GetPhoneNumberVerificationStatusRequest struct {
	// The space ID of the user under the ISV account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 229393838****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613900001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetPhoneNumberVerificationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberVerificationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberVerificationStatusRequest) SetCustSpaceId(v string) *GetPhoneNumberVerificationStatusRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusRequest) SetPhoneNumber(v string) *GetPhoneNumberVerificationStatusRequest {
	s.PhoneNumber = &v
	return s
}

type GetPhoneNumberVerificationStatusResponseBody struct {
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetPhoneNumberVerificationStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneNumberVerificationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberVerificationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetCode(v string) *GetPhoneNumberVerificationStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetData(v *GetPhoneNumberVerificationStatusResponseBodyData) *GetPhoneNumberVerificationStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetMessage(v string) *GetPhoneNumberVerificationStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetRequestId(v string) *GetPhoneNumberVerificationStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetPhoneNumberVerificationStatusResponseBodyData struct {
	// The verification status.
	//
	// example:
	//
	// VERIFIED
	CodeVerificationStatus *string `json:"CodeVerificationStatus,omitempty" xml:"CodeVerificationStatus,omitempty"`
	// The ID of the number.
	//
	// example:
	//
	// 2224342624
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613900001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetPhoneNumberVerificationStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberVerificationStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) SetCodeVerificationStatus(v string) *GetPhoneNumberVerificationStatusResponseBodyData {
	s.CodeVerificationStatus = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) SetId(v string) *GetPhoneNumberVerificationStatusResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) SetPhoneNumber(v string) *GetPhoneNumberVerificationStatusResponseBodyData {
	s.PhoneNumber = &v
	return s
}

type GetPhoneNumberVerificationStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneNumberVerificationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneNumberVerificationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberVerificationStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberVerificationStatusResponse) SetHeaders(v map[string]*string) *GetPhoneNumberVerificationStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponse) SetStatusCode(v int32) *GetPhoneNumberVerificationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponse) SetBody(v *GetPhoneNumberVerificationStatusResponseBody) *GetPhoneNumberVerificationStatusResponse {
	s.Body = v
	return s
}

type GetPreValidatePhoneIdRequest struct {
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The verification code provided when you purchased the pre-registered phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 208393
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s GetPreValidatePhoneIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPreValidatePhoneIdRequest) GoString() string {
	return s.String()
}

func (s *GetPreValidatePhoneIdRequest) SetPhoneNumber(v string) *GetPreValidatePhoneIdRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetPreValidatePhoneIdRequest) SetVerifyCode(v string) *GetPreValidatePhoneIdRequest {
	s.VerifyCode = &v
	return s
}

type GetPreValidatePhoneIdResponseBody struct {
	// The response code. The value OK indicates the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetPreValidatePhoneIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPreValidatePhoneIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPreValidatePhoneIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPreValidatePhoneIdResponseBody) SetCode(v string) *GetPreValidatePhoneIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetPreValidatePhoneIdResponseBody) SetData(v *GetPreValidatePhoneIdResponseBodyData) *GetPreValidatePhoneIdResponseBody {
	s.Data = v
	return s
}

func (s *GetPreValidatePhoneIdResponseBody) SetMessage(v string) *GetPreValidatePhoneIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetPreValidatePhoneIdResponseBody) SetRequestId(v string) *GetPreValidatePhoneIdResponseBody {
	s.RequestId = &v
	return s
}

type GetPreValidatePhoneIdResponseBodyData struct {
	// The phone number.
	//
	// example:
	//
	// 929833
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The ID of the phone number.
	//
	// example:
	//
	// 8613800000000
	PhoneNumberId *string `json:"PhoneNumberId,omitempty" xml:"PhoneNumberId,omitempty"`
}

func (s GetPreValidatePhoneIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPreValidatePhoneIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPreValidatePhoneIdResponseBodyData) SetPhoneNumber(v string) *GetPreValidatePhoneIdResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetPreValidatePhoneIdResponseBodyData) SetPhoneNumberId(v string) *GetPreValidatePhoneIdResponseBodyData {
	s.PhoneNumberId = &v
	return s
}

type GetPreValidatePhoneIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPreValidatePhoneIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPreValidatePhoneIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPreValidatePhoneIdResponse) GoString() string {
	return s.String()
}

func (s *GetPreValidatePhoneIdResponse) SetHeaders(v map[string]*string) *GetPreValidatePhoneIdResponse {
	s.Headers = v
	return s
}

func (s *GetPreValidatePhoneIdResponse) SetStatusCode(v int32) *GetPreValidatePhoneIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPreValidatePhoneIdResponse) SetBody(v *GetPreValidatePhoneIdResponseBody) *GetPreValidatePhoneIdResponse {
	s.Body = v
	return s
}

type GetWhatsappConnectionCatalogRequest struct {
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// C2020939922929292
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The WABA ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 292939399393
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s GetWhatsappConnectionCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWhatsappConnectionCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetWhatsappConnectionCatalogRequest) SetCustSpaceId(v string) *GetWhatsappConnectionCatalogRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) SetOwnerId(v int64) *GetWhatsappConnectionCatalogRequest {
	s.OwnerId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) SetResourceOwnerAccount(v string) *GetWhatsappConnectionCatalogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) SetResourceOwnerId(v int64) *GetWhatsappConnectionCatalogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) SetWabaId(v string) *GetWhatsappConnectionCatalogRequest {
	s.WabaId = &v
	return s
}

type GetWhatsappConnectionCatalogResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {"id":"200292992"}
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWhatsappConnectionCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWhatsappConnectionCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetAccessDeniedDetail(v string) *GetWhatsappConnectionCatalogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetCode(v string) *GetWhatsappConnectionCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetMessage(v string) *GetWhatsappConnectionCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetModel(v map[string]interface{}) *GetWhatsappConnectionCatalogResponseBody {
	s.Model = v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetRequestId(v string) *GetWhatsappConnectionCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponseBody) SetSuccess(v bool) *GetWhatsappConnectionCatalogResponseBody {
	s.Success = &v
	return s
}

type GetWhatsappConnectionCatalogResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWhatsappConnectionCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWhatsappConnectionCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWhatsappConnectionCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetWhatsappConnectionCatalogResponse) SetHeaders(v map[string]*string) *GetWhatsappConnectionCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetWhatsappConnectionCatalogResponse) SetStatusCode(v int32) *GetWhatsappConnectionCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponse) SetBody(v *GetWhatsappConnectionCatalogResponseBody) *GetWhatsappConnectionCatalogResponse {
	s.Body = v
	return s
}

type IsvGetAppIdRequest struct {
	// 权限
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The type of the application. Set the value to WHATSAPP.
	//
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s IsvGetAppIdRequest) String() string {
	return tea.Prettify(s)
}

func (s IsvGetAppIdRequest) GoString() string {
	return s.String()
}

func (s *IsvGetAppIdRequest) SetPermissions(v string) *IsvGetAppIdRequest {
	s.Permissions = &v
	return s
}

func (s *IsvGetAppIdRequest) SetType(v string) *IsvGetAppIdRequest {
	s.Type = &v
	return s
}

type IsvGetAppIdResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 23hr3v
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IsvGetAppIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsvGetAppIdResponseBody) GoString() string {
	return s.String()
}

func (s *IsvGetAppIdResponseBody) SetAccessDeniedDetail(v string) *IsvGetAppIdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetAppId(v string) *IsvGetAppIdResponseBody {
	s.AppId = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetCode(v string) *IsvGetAppIdResponseBody {
	s.Code = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetConfigId(v string) *IsvGetAppIdResponseBody {
	s.ConfigId = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetMessage(v string) *IsvGetAppIdResponseBody {
	s.Message = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetRequestId(v string) *IsvGetAppIdResponseBody {
	s.RequestId = &v
	return s
}

type IsvGetAppIdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsvGetAppIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsvGetAppIdResponse) String() string {
	return tea.Prettify(s)
}

func (s IsvGetAppIdResponse) GoString() string {
	return s.String()
}

func (s *IsvGetAppIdResponse) SetHeaders(v map[string]*string) *IsvGetAppIdResponse {
	s.Headers = v
	return s
}

func (s *IsvGetAppIdResponse) SetStatusCode(v int32) *IsvGetAppIdResponse {
	s.StatusCode = &v
	return s
}

func (s *IsvGetAppIdResponse) SetBody(v *IsvGetAppIdResponseBody) *IsvGetAppIdResponse {
	s.Body = v
	return s
}

type ListChatappTemplateRequest struct {
	// The review status of the message template. Valid values:
	//
	// 	- **pass**: The message template is approved.
	//
	// 	- **fail**: The message template is rejected.
	//
	// 	- **auditing**: The message template is being reviewed.
	//
	// 	- **unaudit**: The review is suspended.
	//
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// 838888822*****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The space ID of the user under the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// >  CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination settings.
	//
	// example:
	//
	// "page": "{\\"index\\": 1,\\"size\\": 20}
	Page *ListChatappTemplateRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE: the Line message template. This type of message template will be released later.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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

func (s *ListChatappTemplateRequest) SetCode(v string) *ListChatappTemplateRequest {
	s.Code = &v
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

func (s *ListChatappTemplateRequest) SetTemplateType(v string) *ListChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

type ListChatappTemplateRequestPage struct {
	// The page number. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
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
	// 	- **pass**: The message template is approved.
	//
	// 	- **fail**: The message template is rejected.
	//
	// 	- **auditing**: The message template is being reviewed.
	//
	// 	- **unaudit**: The review is suspended.
	//
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// 838888822*****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The space ID of the user under the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// >  CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination settings.
	//
	// example:
	//
	// "page": "{\\"index\\": 1,\\"size\\": 20}
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE: the Line message template. This type of message template will be released later.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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

func (s *ListChatappTemplateShrinkRequest) SetCode(v string) *ListChatappTemplateShrinkRequest {
	s.Code = &v
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

func (s *ListChatappTemplateShrinkRequest) SetTemplateType(v string) *ListChatappTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

type ListChatappTemplateResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message templates.
	ListTemplate []*ListChatappTemplateResponseBodyListTemplate `json:"ListTemplate,omitempty" xml:"ListTemplate,omitempty" type:"Repeated"`
	// The error message returned.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *ListChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
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

func (s *ListChatappTemplateResponseBody) SetTotal(v int32) *ListChatappTemplateResponseBody {
	s.Total = &v
	return s
}

type ListChatappTemplateResponseBodyListTemplate struct {
	// The review state of the message template. Valid values:
	//
	// 	- **pass**: The message template is approved.
	//
	// 	- **fail**: The message template is rejected.
	//
	// 	- **auditing**: The message template is being reviewed.
	//
	// 	- **unaudit**: The review is suspended.
	//
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The category of the WhatsApp template. Valid values:
	//
	// 	- **UTILITY**: utility template
	//
	// 	- **MARKETING**: marketing template
	//
	// 	- **AUTHENTICATION**: authentication template
	//
	// The category of the Viber template. Valid values:
	//
	// 	- **text**: template that contains only text
	//
	// 	- **image**: template that contains only an image
	//
	// 	- **text_image_button**: template that contains text, an image, and a button
	//
	// 	- **text_button**: template that contains text and a button
	//
	// 	- **document**: template that contains only a document
	//
	// 	- **video**: template that contains only a video
	//
	// 	- **text_video**: template that contains text and a video
	//
	// 	- **text_video_button**: template that contains text, a video, and a button
	//
	// 	- **text_image**: template that contains text and an image
	//
	// example:
	//
	// TRANSACTIONAL
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The reason for the review failure.
	//
	// example:
	//
	// None
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the template. Valid values: WHATSAPP and VIBER.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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

func (s *ListChatappTemplateResponseBodyListTemplate) SetReason(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.Reason = &v
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

func (s *ListChatappTemplateResponseBodyListTemplate) SetTemplateType(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.TemplateType = &v
	return s
}

type ListChatappTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListFlowRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 99948484
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the Flow.
	//
	// example:
	//
	// flow_001
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The returned pages.
	Page *ListFlowRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRequest) GoString() string {
	return s.String()
}

func (s *ListFlowRequest) SetCustSpaceId(v string) *ListFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListFlowRequest) SetFlowName(v string) *ListFlowRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowRequest) SetPage(v *ListFlowRequestPage) *ListFlowRequest {
	s.Page = v
	return s
}

type ListFlowRequestPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListFlowRequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRequestPage) GoString() string {
	return s.String()
}

func (s *ListFlowRequestPage) SetIndex(v int32) *ListFlowRequestPage {
	s.Index = &v
	return s
}

func (s *ListFlowRequestPage) SetSize(v int32) *ListFlowRequestPage {
	s.Size = &v
	return s
}

type ListFlowShrinkRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 99948484
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the Flow.
	//
	// example:
	//
	// flow_001
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The returned pages.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s ListFlowShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFlowShrinkRequest) SetCustSpaceId(v string) *ListFlowShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListFlowShrinkRequest) SetFlowName(v string) *ListFlowShrinkRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowShrinkRequest) SetPageShrink(v string) *ListFlowShrinkRequest {
	s.PageShrink = &v
	return s
}

type ListFlowResponseBody struct {
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1612C226-E271-4CFE-9F18-4066D550F91B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBody) SetCode(v string) *ListFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowResponseBody) SetData(v []*ListFlowResponseBodyData) *ListFlowResponseBody {
	s.Data = v
	return s
}

func (s *ListFlowResponseBody) SetMessage(v string) *ListFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowResponseBody) SetRequestId(v string) *ListFlowResponseBody {
	s.RequestId = &v
	return s
}

type ListFlowResponseBodyData struct {
	// The categories of the Flows.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The Flow ID.
	//
	// example:
	//
	// 3939393***
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the Flow.
	//
	// example:
	//
	// flow-02020
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s ListFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBodyData) SetCategories(v []*string) *ListFlowResponseBodyData {
	s.Categories = v
	return s
}

func (s *ListFlowResponseBodyData) SetFlowId(v string) *ListFlowResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *ListFlowResponseBodyData) SetFlowName(v string) *ListFlowResponseBodyData {
	s.FlowName = &v
	return s
}

type ListFlowResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponse) GoString() string {
	return s.String()
}

func (s *ListFlowResponse) SetHeaders(v map[string]*string) *ListFlowResponse {
	s.Headers = v
	return s
}

func (s *ListFlowResponse) SetStatusCode(v int32) *ListFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowResponse) SetBody(v *ListFlowResponseBody) *ListFlowResponse {
	s.Body = v
	return s
}

type ListPhoneMessageQrdlRequest struct {
	// example:
	//
	// 9383883
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s ListPhoneMessageQrdlRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *ListPhoneMessageQrdlRequest) SetCustSpaceId(v string) *ListPhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListPhoneMessageQrdlRequest) SetPhoneNumber(v string) *ListPhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

type ListPhoneMessageQrdlResponseBody struct {
	// example:
	//
	// OK
	Code    *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*ListPhoneMessageQrdlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPhoneMessageQrdlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneMessageQrdlResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhoneMessageQrdlResponseBody) SetCode(v string) *ListPhoneMessageQrdlResponseBody {
	s.Code = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBody) SetData(v []*ListPhoneMessageQrdlResponseBodyData) *ListPhoneMessageQrdlResponseBody {
	s.Data = v
	return s
}

func (s *ListPhoneMessageQrdlResponseBody) SetMessage(v string) *ListPhoneMessageQrdlResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBody) SetRequestId(v string) *ListPhoneMessageQrdlResponseBody {
	s.RequestId = &v
	return s
}

type ListPhoneMessageQrdlResponseBodyData struct {
	// example:
	//
	// https://wa.msg/
	DeepLinkUrl *string `json:"DeepLinkUrl,omitempty" xml:"DeepLinkUrl,omitempty"`
	// example:
	//
	// PNG
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	// example:
	//
	// 8613800
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// Hello
	PrefilledMessage *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
	// example:
	//
	// https://img.png
	QrImageUrl *string `json:"QrImageUrl,omitempty" xml:"QrImageUrl,omitempty"`
	// example:
	//
	// IUIED999
	QrdlCode *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
}

func (s ListPhoneMessageQrdlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneMessageQrdlResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetDeepLinkUrl(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.DeepLinkUrl = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetGenerateQrImage(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.GenerateQrImage = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetPhoneNumber(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetPrefilledMessage(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.PrefilledMessage = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetQrImageUrl(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.QrImageUrl = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetQrdlCode(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.QrdlCode = &v
	return s
}

type ListPhoneMessageQrdlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPhoneMessageQrdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPhoneMessageQrdlResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneMessageQrdlResponse) GoString() string {
	return s.String()
}

func (s *ListPhoneMessageQrdlResponse) SetHeaders(v map[string]*string) *ListPhoneMessageQrdlResponse {
	s.Headers = v
	return s
}

func (s *ListPhoneMessageQrdlResponse) SetStatusCode(v int32) *ListPhoneMessageQrdlResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPhoneMessageQrdlResponse) SetBody(v *ListPhoneMessageQrdlResponseBody) *ListPhoneMessageQrdlResponse {
	s.Body = v
	return s
}

type ListProductRequest struct {
	// The cursor that points to the end of the page of the returned data.
	//
	// example:
	//
	// kdkii48jfjjei3
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// The cursor that points to the beginning of the page of the returned data.
	//
	// example:
	//
	// wiidkd939kek93
	Before *string `json:"Before,omitempty" xml:"Before,omitempty"`
	// The catalog ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 29398389292
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// C29398388383
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The fields. Separate multiple fields with commas (,).
	//
	//  see [product fields](https://help.aliyun.com/document_detail/2579419.html)
	//
	// example:
	//
	// id,name
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The number of products to be queried. Valid values: 1 to 1000.
	//
	// example:
	//
	// 73
	Limit                *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the WhatsApp Business account (WABA).
	//
	// This parameter is required.
	//
	// example:
	//
	// 38487474747
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s ListProductRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductRequest) GoString() string {
	return s.String()
}

func (s *ListProductRequest) SetAfter(v string) *ListProductRequest {
	s.After = &v
	return s
}

func (s *ListProductRequest) SetBefore(v string) *ListProductRequest {
	s.Before = &v
	return s
}

func (s *ListProductRequest) SetCatalogId(v string) *ListProductRequest {
	s.CatalogId = &v
	return s
}

func (s *ListProductRequest) SetCustSpaceId(v string) *ListProductRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListProductRequest) SetFields(v string) *ListProductRequest {
	s.Fields = &v
	return s
}

func (s *ListProductRequest) SetLimit(v int64) *ListProductRequest {
	s.Limit = &v
	return s
}

func (s *ListProductRequest) SetOwnerId(v int64) *ListProductRequest {
	s.OwnerId = &v
	return s
}

func (s *ListProductRequest) SetResourceOwnerAccount(v string) *ListProductRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListProductRequest) SetResourceOwnerId(v int64) *ListProductRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListProductRequest) SetWabaId(v string) *ListProductRequest {
	s.WabaId = &v
	return s
}

type ListProductResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model *ListProductResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductResponseBody) SetAccessDeniedDetail(v string) *ListProductResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListProductResponseBody) SetCode(v string) *ListProductResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductResponseBody) SetMessage(v string) *ListProductResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductResponseBody) SetModel(v *ListProductResponseBodyModel) *ListProductResponseBody {
	s.Model = v
	return s
}

func (s *ListProductResponseBody) SetRequestId(v string) *ListProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductResponseBody) SetSuccess(v bool) *ListProductResponseBody {
	s.Success = &v
	return s
}

type ListProductResponseBodyModel struct {
	// The returned data.
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination details.
	Paging *ListProductResponseBodyModelPaging `json:"Paging,omitempty" xml:"Paging,omitempty" type:"Struct"`
}

func (s ListProductResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ListProductResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListProductResponseBodyModel) SetData(v []map[string]interface{}) *ListProductResponseBodyModel {
	s.Data = v
	return s
}

func (s *ListProductResponseBodyModel) SetPaging(v *ListProductResponseBodyModelPaging) *ListProductResponseBodyModel {
	s.Paging = v
	return s
}

type ListProductResponseBodyModelPaging struct {
	// The cursors for pagination.
	Cursors *ListProductResponseBodyModelPagingCursors `json:"Cursors,omitempty" xml:"Cursors,omitempty" type:"Struct"`
}

func (s ListProductResponseBodyModelPaging) String() string {
	return tea.Prettify(s)
}

func (s ListProductResponseBodyModelPaging) GoString() string {
	return s.String()
}

func (s *ListProductResponseBodyModelPaging) SetCursors(v *ListProductResponseBodyModelPagingCursors) *ListProductResponseBodyModelPaging {
	s.Cursors = v
	return s
}

type ListProductResponseBodyModelPagingCursors struct {
	// The cursor that points to the end of the page of the returned data.
	//
	// example:
	//
	// sjsuueu83838
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// The cursor that points to the beginning of the page of the returned data.
	//
	// example:
	//
	// sjjsjdjjdjd83883
	Before *string `json:"Before,omitempty" xml:"Before,omitempty"`
}

func (s ListProductResponseBodyModelPagingCursors) String() string {
	return tea.Prettify(s)
}

func (s ListProductResponseBodyModelPagingCursors) GoString() string {
	return s.String()
}

func (s *ListProductResponseBodyModelPagingCursors) SetAfter(v string) *ListProductResponseBodyModelPagingCursors {
	s.After = &v
	return s
}

func (s *ListProductResponseBodyModelPagingCursors) SetBefore(v string) *ListProductResponseBodyModelPagingCursors {
	s.Before = &v
	return s
}

type ListProductResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductResponse) GoString() string {
	return s.String()
}

func (s *ListProductResponse) SetHeaders(v map[string]*string) *ListProductResponse {
	s.Headers = v
	return s
}

func (s *ListProductResponse) SetStatusCode(v int32) *ListProductResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductResponse) SetBody(v *ListProductResponseBody) *ListProductResponse {
	s.Body = v
	return s
}

type ListProductCatalogRequest struct {
	// The cursor that points to the end of the page of the returned data.
	//
	// example:
	//
	// kdkii48jfjjei3
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// The cursor that points to the beginning of the page of the returned data.
	//
	// example:
	//
	// wiidkd939kek93
	Before *string `json:"Before,omitempty" xml:"Before,omitempty"`
	// The Business Manager ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 28
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The fields. Separate multiple fields with commas (,).
	//
	// see  [catalog fields](https://help.aliyun.com/document_detail/2579419.html)
	//
	// example:
	//
	// id,name
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The number of catalogs to be queried. Valid values: 1 to 1000.
	//
	// example:
	//
	// 73
	Limit                *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListProductCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductCatalogRequest) GoString() string {
	return s.String()
}

func (s *ListProductCatalogRequest) SetAfter(v string) *ListProductCatalogRequest {
	s.After = &v
	return s
}

func (s *ListProductCatalogRequest) SetBefore(v string) *ListProductCatalogRequest {
	s.Before = &v
	return s
}

func (s *ListProductCatalogRequest) SetBusinessId(v int64) *ListProductCatalogRequest {
	s.BusinessId = &v
	return s
}

func (s *ListProductCatalogRequest) SetCustSpaceId(v string) *ListProductCatalogRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListProductCatalogRequest) SetFields(v string) *ListProductCatalogRequest {
	s.Fields = &v
	return s
}

func (s *ListProductCatalogRequest) SetLimit(v int64) *ListProductCatalogRequest {
	s.Limit = &v
	return s
}

func (s *ListProductCatalogRequest) SetOwnerId(v int64) *ListProductCatalogRequest {
	s.OwnerId = &v
	return s
}

func (s *ListProductCatalogRequest) SetResourceOwnerAccount(v string) *ListProductCatalogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListProductCatalogRequest) SetResourceOwnerId(v int64) *ListProductCatalogRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListProductCatalogResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model *ListProductCatalogResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponseBody) SetAccessDeniedDetail(v string) *ListProductCatalogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListProductCatalogResponseBody) SetCode(v string) *ListProductCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductCatalogResponseBody) SetMessage(v string) *ListProductCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductCatalogResponseBody) SetModel(v *ListProductCatalogResponseBodyModel) *ListProductCatalogResponseBody {
	s.Model = v
	return s
}

func (s *ListProductCatalogResponseBody) SetRequestId(v string) *ListProductCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductCatalogResponseBody) SetSuccess(v bool) *ListProductCatalogResponseBody {
	s.Success = &v
	return s
}

type ListProductCatalogResponseBodyModel struct {
	// The returned data.
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination details.
	Paging *ListProductCatalogResponseBodyModelPaging `json:"Paging,omitempty" xml:"Paging,omitempty" type:"Struct"`
}

func (s ListProductCatalogResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ListProductCatalogResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponseBodyModel) SetData(v []map[string]interface{}) *ListProductCatalogResponseBodyModel {
	s.Data = v
	return s
}

func (s *ListProductCatalogResponseBodyModel) SetPaging(v *ListProductCatalogResponseBodyModelPaging) *ListProductCatalogResponseBodyModel {
	s.Paging = v
	return s
}

type ListProductCatalogResponseBodyModelPaging struct {
	// The cursors for pagination.
	Cursors *ListProductCatalogResponseBodyModelPagingCursors `json:"Cursors,omitempty" xml:"Cursors,omitempty" type:"Struct"`
}

func (s ListProductCatalogResponseBodyModelPaging) String() string {
	return tea.Prettify(s)
}

func (s ListProductCatalogResponseBodyModelPaging) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponseBodyModelPaging) SetCursors(v *ListProductCatalogResponseBodyModelPagingCursors) *ListProductCatalogResponseBodyModelPaging {
	s.Cursors = v
	return s
}

type ListProductCatalogResponseBodyModelPagingCursors struct {
	// The cursor that points to the end of the page of the returned data.
	//
	// example:
	//
	// sjsuueu83838
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// The cursor that points to the beginning of the page of the returned data.
	//
	// example:
	//
	// sjjsjdjjdjd83883
	Before *string `json:"Before,omitempty" xml:"Before,omitempty"`
}

func (s ListProductCatalogResponseBodyModelPagingCursors) String() string {
	return tea.Prettify(s)
}

func (s ListProductCatalogResponseBodyModelPagingCursors) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponseBodyModelPagingCursors) SetAfter(v string) *ListProductCatalogResponseBodyModelPagingCursors {
	s.After = &v
	return s
}

func (s *ListProductCatalogResponseBodyModelPagingCursors) SetBefore(v string) *ListProductCatalogResponseBodyModelPagingCursors {
	s.Before = &v
	return s
}

type ListProductCatalogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductCatalogResponse) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponse) SetHeaders(v map[string]*string) *ListProductCatalogResponse {
	s.Headers = v
	return s
}

func (s *ListProductCatalogResponse) SetStatusCode(v int32) *ListProductCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductCatalogResponse) SetBody(v *ListProductCatalogResponseBody) *ListProductCatalogResponse {
	s.Body = v
	return s
}

type ModifyChatappTemplateRequest struct {
	// The category of the Viber message template. Valid values:
	//
	// 	- **text**: the template that contains only text
	//
	// 	- **image**: the template that contains only images
	//
	// 	- **text_image_button**: the template that contains text, images, and buttons
	//
	// 	- **text_button**: the template that contains text and buttons
	//
	// 	- **document**: the template that contains only documents
	//
	// 	- **video**: the template that contains only videos
	//
	// 	- **text_video**: the template that contains text and videos
	//
	// 	- **text_video_button**: the template that contains text, videos, and buttons
	//
	// 	- **text_image**: the template that contains text and images
	//
	// > This parameter applies only to Viber message templates.
	//
	// example:
	//
	// text
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The components of the message template.
	//
	// >  If Category is set to AUTHENTICATION, the Type sub-parameter of the Components parameter cannot be set to HEADER. If the Type sub-parameter is set to BODY or FOOTER, the Text sub-parameter of the Components parameter is empty and text in the body or footer is automatically generated.
	//
	// This parameter is required.
	Components []*ModifyChatappTemplateRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// > CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 659216218162179
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// ksiekdki39ksks93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Validity period of authentication template message sending in WhatsApp
	//
	// >This attribute requires providing waba in advance to Alibaba operators to open the whitelist, otherwise it will result in template submission failure
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The message template code.
	//
	// example:
	//
	// 8472929283883
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE: the Line message template. This type of message template will be released later.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ModifyChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequest) SetCategory(v string) *ModifyChatappTemplateRequest {
	s.Category = &v
	return s
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

func (s *ModifyChatappTemplateRequest) SetMessageSendTtlSeconds(v int32) *ModifyChatappTemplateRequest {
	s.MessageSendTtlSeconds = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetTemplateCode(v string) *ModifyChatappTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetTemplateName(v string) *ModifyChatappTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyChatappTemplateRequest) SetTemplateType(v string) *ModifyChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

type ModifyChatappTemplateRequestComponents struct {
	// The note indicating that customers cannot share verification codes with others. The note is displayed in the message body. This parameter is valid only if Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to BODY for a WhatsApp message template.
	//
	// example:
	//
	// false
	AddSecretRecommendation *bool `json:"AddSecretRecommendation,omitempty" xml:"AddSecretRecommendation,omitempty"`
	// The buttons. Specify this parameter only if you set the Type sub-parameter of the Components parameter to **BUTTONS**.
	//
	// >  The following section describes the limits on the number of buttons in a WhatsApp message.
	//
	// 	- A marketing or utility WhatsApp message template supports up to 10 buttons.
	//
	// 	- A WhatsApp message template can contain only one phone call button.
	//
	// 	- A WhatsApp message template can contain up to two URL buttons.
	//
	// 	- In a WhatsApp message template, the quick reply button cannot be used together with the phone call button or the URL button.
	Buttons []*ModifyChatappTemplateRequestComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// The description.
	//
	// >  You can specify this parameter if the Type sub-parameter of the Components parameter is set to **HEADER*	- and the Format sub-parameter of the Components parameter is set to **IMAGE, DOCUMENT, or VIDEO**.
	//
	// example:
	//
	// This is a video
	Caption *string                                        `json:"Caption,omitempty" xml:"Caption,omitempty"`
	Cards   []*ModifyChatappTemplateRequestComponentsCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// The validity period of the verification code in the WhatsApp authentication template. Unit: minutes. This parameter is valid only when Category is set to AUTHENTICATION and the Type sub-parameter of the Components parameter is set to FOOTER in a WhatsApp message template. The validity period of the verification code is displayed in the footer.
	//
	// example:
	//
	// 5
	CodeExpirationMinutes *int32 `json:"CodeExpirationMinutes,omitempty" xml:"CodeExpirationMinutes,omitempty"`
	// The length of the video in the Viber message template. Unit: seconds. Valid values: 0 to 600.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the document.
	//
	// >  You can specify this parameter if the Type sub-parameter of the Components parameter is set to **HEADER*	- and the Format sub-parameter of the Components parameter is set to **DOCUMENT**.
	//
	// example:
	//
	// video name
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the document attached in the Viber message template.
	//
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The type of the media resources that are included in the message. Valid values:
	//
	// 	- **TEXT**
	//
	// 	- **IMAGE**
	//
	// 	- **DOCUMENT**
	//
	// 	- **VIDEO**
	//
	// example:
	//
	// TEXT
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// true
	HasExpiration *bool `json:"HasExpiration,omitempty" xml:"HasExpiration,omitempty"`
	// The text of the message that you want to send.
	//
	// >  If Category is set to AUTHENTICATION, the Text sub-parameter of the Components parameter is empty.
	//
	// example:
	//
	// hello chatapp
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The thumbnail URL of the video in the Viber message template.
	//
	// example:
	//
	// https://cdn.multiplymall.mobiapp.cloud/cloudcode/yc-165407506207478-165511576113195/20220905/ec5b9737-1507-4208-bb27-8da3958da961.jpg?x-oss-process=image/resize,w_100
	ThumbUrl *string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`
	// The type of the component. Valid values:
	//
	// 	- **BODY**
	//
	// 	- **HEADER**
	//
	// 	- **FOOTER**
	//
	// 	- **BUTTONS**
	//
	// >
	//
	// 	- The following limits apply to components in WhatsApp message templates: A **BODY*	- component cannot exceed 1,024 characters in length. A **HEADER*	- or **FOOTER*	- component cannot exceed 60 characters in length.
	//
	// 	- **FOOTER*	- components are not supported in Viber message templates.
	//
	// 	- In a Viber message template, media resources such as images, videos, and documents are placed in the **HEADER*	- component. If a Viber message contains text and images, the images are placed under the text in the message received on a device.
	//
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the media resource.
	//
	// example:
	//
	// https://img.tukuppt.com/png_preview/00/10/24/1GygxVK3F4.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ModifyChatappTemplateRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponents) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponents) SetAddSecretRecommendation(v bool) *ModifyChatappTemplateRequestComponents {
	s.AddSecretRecommendation = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetButtons(v []*ModifyChatappTemplateRequestComponentsButtons) *ModifyChatappTemplateRequestComponents {
	s.Buttons = v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetCaption(v string) *ModifyChatappTemplateRequestComponents {
	s.Caption = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetCards(v []*ModifyChatappTemplateRequestComponentsCards) *ModifyChatappTemplateRequestComponents {
	s.Cards = v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetCodeExpirationMinutes(v int32) *ModifyChatappTemplateRequestComponents {
	s.CodeExpirationMinutes = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetDuration(v int32) *ModifyChatappTemplateRequestComponents {
	s.Duration = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetFileName(v string) *ModifyChatappTemplateRequestComponents {
	s.FileName = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetFileType(v string) *ModifyChatappTemplateRequestComponents {
	s.FileType = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetFormat(v string) *ModifyChatappTemplateRequestComponents {
	s.Format = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetHasExpiration(v bool) *ModifyChatappTemplateRequestComponents {
	s.HasExpiration = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetText(v string) *ModifyChatappTemplateRequestComponents {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponents) SetThumbUrl(v string) *ModifyChatappTemplateRequestComponents {
	s.ThumbUrl = &v
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
	// The text of the one-tap autofill button. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// Autofill
	AutofillText *string `json:"AutofillText,omitempty" xml:"AutofillText,omitempty"`
	// example:
	//
	// 120293
	CouponCode *string `json:"CouponCode,omitempty" xml:"CouponCode,omitempty"`
	// example:
	//
	// NAVIGATE
	FlowAction *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// example:
	//
	// 28383872***
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The unsubscribe button. This parameter is valid if Category is set to MARKETING and the Type sub-parameter of the Buttons parameter is set to QUICK_REPLY for a WhatsApp message template. After you configure message sending in the ChatApp Message Service console, marketing messages are not sent to customers if they click this button.
	//
	// example:
	//
	// false
	IsOptOut *bool `json:"IsOptOut,omitempty" xml:"IsOptOut,omitempty"`
	// example:
	//
	// SIGN_UP
	NavigateScreen *string `json:"NavigateScreen,omitempty" xml:"NavigateScreen,omitempty"`
	// Deprecated
	//
	// The app package name that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// com.demo
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The phone number.
	//
	// example:
	//
	// +8613888887889
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Deprecated
	//
	// The app signing key hash that WhatsApp uses to load your app. This parameter is required if Category is set to AUTHENTICATION and the Type sub-parameter of the Buttons parameter is set to ONE_TAP for a WhatsApp message template.
	//
	// example:
	//
	// 29dkeke
	SignatureHash *string                                                       `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
	SupportedApps []*ModifyChatappTemplateRequestComponentsButtonsSupportedApps `json:"SupportedApps,omitempty" xml:"SupportedApps,omitempty" type:"Repeated"`
	// The text of the button.
	//
	// example:
	//
	// phone-button-text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the button. Valid values:
	//
	// 	- **PHONE_NUMBER**: phone call button
	//
	// 	- **URL**: URL button
	//
	// 	- **QUICK_REPLY**: quick reply button
	//
	// 	- **COPY_CODE**: copy code button if Category is set to AUTHENTICATION
	//
	// 	- **ONE_TAP**: one-tap autofill button if Category is set to AUTHENTICATION
	//
	// >
	//
	// 	- In a WhatsApp message template, a quick reply button cannot be used together with a phone call button or a URL button.
	//
	// 	- You can add a combination of two URL buttons or a combination of a URL button and a phone call button to a WhatsApp message template.
	//
	// 	- If Category is set to AUTHENTICATION for a WhatsApp message template, you can add only one button to the WhatsApp message template and you must set the Type sub-parameter of the Buttons parameter to COPY_CODE or ONE_TAP. If the Type sub-parameter of the Buttons parameter is set to COPY_CODE, the Text sub-parameter of the Buttons parameter is required. If the Type sub-parameter of the Buttons parameter is set to ONE_TAP, the Text, SignatureHash, PackageName, and AutofillText sub-parameters of the Buttons parameter are required. The value of Text is displayed if the desired app is not installed on the device. The value of Text indicates that you must manually copy the verification code.
	//
	// 	- You can add only one button to a Viber message template, and you must set the Type sub-parameter of the Buttons parameter to URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL to which you are redirected when you click the URL button.
	//
	// example:
	//
	// https://www.website.com/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the URL. Valid values:
	//
	// 	- **static**
	//
	// 	- **dynamic**
	//
	// example:
	//
	// dynamic
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsButtons) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetAutofillText(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.AutofillText = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetCouponCode(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.CouponCode = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetFlowAction(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.FlowAction = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetFlowId(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.FlowId = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetIsOptOut(v bool) *ModifyChatappTemplateRequestComponentsButtons {
	s.IsOptOut = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetNavigateScreen(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.NavigateScreen = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetPackageName(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.PackageName = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetPhoneNumber(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetSignatureHash(v string) *ModifyChatappTemplateRequestComponentsButtons {
	s.SignatureHash = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtons) SetSupportedApps(v []*ModifyChatappTemplateRequestComponentsButtonsSupportedApps) *ModifyChatappTemplateRequestComponentsButtons {
	s.SupportedApps = v
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

type ModifyChatappTemplateRequestComponentsButtonsSupportedApps struct {
	// example:
	//
	// com.aliyun.pack.***
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// kdikie9398k
	SignatureHash *string `json:"SignatureHash,omitempty" xml:"SignatureHash,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsButtonsSupportedApps) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsButtonsSupportedApps) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsButtonsSupportedApps) SetPackageName(v string) *ModifyChatappTemplateRequestComponentsButtonsSupportedApps {
	s.PackageName = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsButtonsSupportedApps) SetSignatureHash(v string) *ModifyChatappTemplateRequestComponentsButtonsSupportedApps {
	s.SignatureHash = &v
	return s
}

type ModifyChatappTemplateRequestComponentsCards struct {
	// This parameter is required.
	CardComponents []*ModifyChatappTemplateRequestComponentsCardsCardComponents `json:"CardComponents,omitempty" xml:"CardComponents,omitempty" type:"Repeated"`
}

func (s ModifyChatappTemplateRequestComponentsCards) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsCards) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsCards) SetCardComponents(v []*ModifyChatappTemplateRequestComponentsCardsCardComponents) *ModifyChatappTemplateRequestComponentsCards {
	s.CardComponents = v
	return s
}

type ModifyChatappTemplateRequestComponentsCardsCardComponents struct {
	Buttons []*ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// example:
	//
	// IMAGE
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// Who is the very powerful team
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BODY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://alibaba.com/img.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsCardsCardComponents) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsCardsCardComponents) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetButtons(v []*ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Buttons = v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetFormat(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Format = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetText(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetType(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Type = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponents) SetUrl(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponents {
	s.Url = &v
	return s
}

type ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons struct {
	// example:
	//
	// +8613800
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// Call me
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PHONE_NUMBER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://alibaba.com/xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// static
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetPhoneNumber(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetText(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Text = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetType(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Type = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetUrl(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.Url = &v
	return s
}

func (s *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons) SetUrlType(v string) *ModifyChatappTemplateRequestComponentsCardsCardComponentsButtons {
	s.UrlType = &v
	return s
}

type ModifyChatappTemplateShrinkRequest struct {
	// The category of the Viber message template. Valid values:
	//
	// 	- **text**: the template that contains only text
	//
	// 	- **image**: the template that contains only images
	//
	// 	- **text_image_button**: the template that contains text, images, and buttons
	//
	// 	- **text_button**: the template that contains text and buttons
	//
	// 	- **document**: the template that contains only documents
	//
	// 	- **video**: the template that contains only videos
	//
	// 	- **text_video**: the template that contains text and videos
	//
	// 	- **text_video_button**: the template that contains text, videos, and buttons
	//
	// 	- **text_image**: the template that contains text and images
	//
	// > This parameter applies only to Viber message templates.
	//
	// example:
	//
	// text
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The components of the message template.
	//
	// >  If Category is set to AUTHENTICATION, the Type sub-parameter of the Components parameter cannot be set to HEADER. If the Type sub-parameter is set to BODY or FOOTER, the Text sub-parameter of the Components parameter is empty and text in the body or footer is automatically generated.
	//
	// This parameter is required.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The WhatsApp Business account (WABA) ID of the user within the independent software vendor (ISV) account.
	//
	// > CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 659216218162179
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The examples of variables that are used when you create the message template.
	ExampleShrink *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// ksiekdki39ksks93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Validity period of authentication template message sending in WhatsApp
	//
	// >This attribute requires providing waba in advance to Alibaba operators to open the whitelist, otherwise it will result in template submission failure
	//
	// example:
	//
	// 120
	MessageSendTtlSeconds *int32 `json:"MessageSendTtlSeconds,omitempty" xml:"MessageSendTtlSeconds,omitempty"`
	// The message template code.
	//
	// example:
	//
	// 8472929283883
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// 	- LINE: the Line message template. This type of message template will be released later.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ModifyChatappTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateShrinkRequest) SetCategory(v string) *ModifyChatappTemplateShrinkRequest {
	s.Category = &v
	return s
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

func (s *ModifyChatappTemplateShrinkRequest) SetMessageSendTtlSeconds(v int32) *ModifyChatappTemplateShrinkRequest {
	s.MessageSendTtlSeconds = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetTemplateCode(v string) *ModifyChatappTemplateShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetTemplateName(v string) *ModifyChatappTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyChatappTemplateShrinkRequest) SetTemplateType(v string) *ModifyChatappTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

type ModifyChatappTemplateResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ModifyChatappTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// NONE
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *ModifyChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	//
	// example:
	//
	// 8472929283883
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifyFlowRequest struct {
	// The categories of the Flow.
	//
	// This parameter is required.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 9493884
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The Flow ID.
	//
	// example:
	//
	// 2938838
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the Flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s ModifyFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowRequest) SetCategories(v []*string) *ModifyFlowRequest {
	s.Categories = v
	return s
}

func (s *ModifyFlowRequest) SetCustSpaceId(v string) *ModifyFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyFlowRequest) SetFlowId(v string) *ModifyFlowRequest {
	s.FlowId = &v
	return s
}

func (s *ModifyFlowRequest) SetFlowName(v string) *ModifyFlowRequest {
	s.FlowName = &v
	return s
}

type ModifyFlowShrinkRequest struct {
	// The categories of the Flow.
	//
	// This parameter is required.
	CategoriesShrink *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 9493884
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The Flow ID.
	//
	// example:
	//
	// 2938838
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the Flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s ModifyFlowShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowShrinkRequest) SetCategoriesShrink(v string) *ModifyFlowShrinkRequest {
	s.CategoriesShrink = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetCustSpaceId(v string) *ModifyFlowShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetFlowId(v string) *ModifyFlowShrinkRequest {
	s.FlowId = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetFlowName(v string) *ModifyFlowShrinkRequest {
	s.FlowName = &v
	return s
}

type ModifyFlowResponseBody struct {
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ModifyFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1612C226-E271-4CFE-9F18-4066D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponseBody) SetCode(v string) *ModifyFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFlowResponseBody) SetData(v *ModifyFlowResponseBodyData) *ModifyFlowResponseBody {
	s.Data = v
	return s
}

func (s *ModifyFlowResponseBody) SetMessage(v string) *ModifyFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyFlowResponseBody) SetRequestId(v string) *ModifyFlowResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowResponseBodyData struct {
	// The categories of the Flow.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The Flow ID.
	//
	// example:
	//
	// 3939399****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the Flow.
	//
	// example:
	//
	// flow-00203
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s ModifyFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponseBodyData) SetCategories(v []*string) *ModifyFlowResponseBodyData {
	s.Categories = v
	return s
}

func (s *ModifyFlowResponseBodyData) SetFlowId(v string) *ModifyFlowResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *ModifyFlowResponseBodyData) SetFlowName(v string) *ModifyFlowResponseBodyData {
	s.FlowName = &v
	return s
}

type ModifyFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponse) SetHeaders(v map[string]*string) *ModifyFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowResponse) SetStatusCode(v int32) *ModifyFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFlowResponse) SetBody(v *ModifyFlowResponseBody) *ModifyFlowResponse {
	s.Body = v
	return s
}

type ModifyPhoneBusinessProfileRequest struct {
	// example:
	//
	// business profile
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// example:
	//
	// The phone number.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Modifies the business information of the account to which a specified phone number is bound.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The address.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description.
	//
	// example:
	//
	// aa@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The industry.
	//
	// > Enum: {OTHER, AUTO, BEAUTY, APPAREL, EDU, ENTERTAIN, EVENT_PLAN, FINANCE, GROCERY, GOVT, HOTEL, HEALTH, NONPROFIT, PROF_SERVICES, RETAIL, TRAVEL, RESTAURANT}
	//
	// example:
	//
	// http://a.img
	ProfilePictureUrl *string `json:"ProfilePictureUrl,omitempty" xml:"ProfilePictureUrl,omitempty"`
	// Sichuan
	//
	// example:
	//
	// OTHER
	Vertical *string `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
	// The email address.
	Websites []*string `json:"Websites,omitempty" xml:"Websites,omitempty" type:"Repeated"`
}

func (s ModifyPhoneBusinessProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPhoneBusinessProfileRequest) GoString() string {
	return s.String()
}

func (s *ModifyPhoneBusinessProfileRequest) SetAbout(v string) *ModifyPhoneBusinessProfileRequest {
	s.About = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetAddress(v string) *ModifyPhoneBusinessProfileRequest {
	s.Address = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetCustSpaceId(v string) *ModifyPhoneBusinessProfileRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetDescription(v string) *ModifyPhoneBusinessProfileRequest {
	s.Description = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetEmail(v string) *ModifyPhoneBusinessProfileRequest {
	s.Email = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetPhoneNumber(v string) *ModifyPhoneBusinessProfileRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetProfilePictureUrl(v string) *ModifyPhoneBusinessProfileRequest {
	s.ProfilePictureUrl = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetVertical(v string) *ModifyPhoneBusinessProfileRequest {
	s.Vertical = &v
	return s
}

func (s *ModifyPhoneBusinessProfileRequest) SetWebsites(v []*string) *ModifyPhoneBusinessProfileRequest {
	s.Websites = v
	return s
}

type ModifyPhoneBusinessProfileShrinkRequest struct {
	// example:
	//
	// business profile
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// example:
	//
	// The phone number.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Modifies the business information of the account to which a specified phone number is bound.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The address.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description.
	//
	// example:
	//
	// aa@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The industry.
	//
	// > Enum: {OTHER, AUTO, BEAUTY, APPAREL, EDU, ENTERTAIN, EVENT_PLAN, FINANCE, GROCERY, GOVT, HOTEL, HEALTH, NONPROFIT, PROF_SERVICES, RETAIL, TRAVEL, RESTAURANT}
	//
	// example:
	//
	// http://a.img
	ProfilePictureUrl *string `json:"ProfilePictureUrl,omitempty" xml:"ProfilePictureUrl,omitempty"`
	// Sichuan
	//
	// example:
	//
	// OTHER
	Vertical *string `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
	// The email address.
	WebsitesShrink *string `json:"Websites,omitempty" xml:"Websites,omitempty"`
}

func (s ModifyPhoneBusinessProfileShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPhoneBusinessProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetAbout(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.About = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetAddress(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.Address = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetCustSpaceId(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetDescription(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetEmail(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.Email = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetPhoneNumber(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetProfilePictureUrl(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.ProfilePictureUrl = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetVertical(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.Vertical = &v
	return s
}

func (s *ModifyPhoneBusinessProfileShrinkRequest) SetWebsitesShrink(v string) *ModifyPhoneBusinessProfileShrinkRequest {
	s.WebsitesShrink = &v
	return s
}

type ModifyPhoneBusinessProfileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The URL of the website.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The websites.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPhoneBusinessProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPhoneBusinessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetAccessDeniedDetail(v string) *ModifyPhoneBusinessProfileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetCode(v string) *ModifyPhoneBusinessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetMessage(v string) *ModifyPhoneBusinessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponseBody) SetRequestId(v string) *ModifyPhoneBusinessProfileResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPhoneBusinessProfileResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPhoneBusinessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPhoneBusinessProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPhoneBusinessProfileResponse) GoString() string {
	return s.String()
}

func (s *ModifyPhoneBusinessProfileResponse) SetHeaders(v map[string]*string) *ModifyPhoneBusinessProfileResponse {
	s.Headers = v
	return s
}

func (s *ModifyPhoneBusinessProfileResponse) SetStatusCode(v int32) *ModifyPhoneBusinessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponse) SetBody(v *ModifyPhoneBusinessProfileResponseBody) *ModifyPhoneBusinessProfileResponse {
	s.Body = v
	return s
}

type PublishFlowRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The Flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s PublishFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishFlowRequest) GoString() string {
	return s.String()
}

func (s *PublishFlowRequest) SetCustSpaceId(v string) *PublishFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *PublishFlowRequest) SetFlowId(v string) *PublishFlowRequest {
	s.FlowId = &v
	return s
}

type PublishFlowResponseBody struct {
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishFlowResponseBody) GoString() string {
	return s.String()
}

func (s *PublishFlowResponseBody) SetCode(v string) *PublishFlowResponseBody {
	s.Code = &v
	return s
}

func (s *PublishFlowResponseBody) SetMessage(v string) *PublishFlowResponseBody {
	s.Message = &v
	return s
}

func (s *PublishFlowResponseBody) SetRequestId(v string) *PublishFlowResponseBody {
	s.RequestId = &v
	return s
}

type PublishFlowResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishFlowResponse) GoString() string {
	return s.String()
}

func (s *PublishFlowResponse) SetHeaders(v map[string]*string) *PublishFlowResponse {
	s.Headers = v
	return s
}

func (s *PublishFlowResponse) SetStatusCode(v int32) *PublishFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishFlowResponse) SetBody(v *PublishFlowResponseBody) *PublishFlowResponse {
	s.Body = v
	return s
}

type QueryChatappBindWabaRequest struct {
	// The space ID of the user under the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// aksik93kdkkxmwol93939
	IsvCode              *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *QueryChatappBindWabaRequest) SetOwnerId(v int64) *QueryChatappBindWabaRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryChatappBindWabaRequest) SetResourceOwnerAccount(v string) *QueryChatappBindWabaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryChatappBindWabaRequest) SetResourceOwnerId(v int64) *QueryChatappBindWabaRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryChatappBindWabaResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryChatappBindWabaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryChatappBindWabaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappBindWabaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaResponseBody) SetAccessDeniedDetail(v string) *QueryChatappBindWabaResponseBody {
	s.AccessDeniedDetail = &v
	return s
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

func (s *QueryChatappBindWabaResponseBody) SetSuccess(v bool) *QueryChatappBindWabaResponseBody {
	s.Success = &v
	return s
}

type QueryChatappBindWabaResponseBodyData struct {
	// The review state of the WhatsApp Business account (WABA).
	//
	// >  Valid values:
	//
	// 	- PENDING: The WABA is to be reviewed.
	//
	// 	- APPROVED: The WABA was approved.
	//
	// 	- REJECTED: The WABA was rejected.
	//
	// 	- DISABLED: The WABA was forbidden.
	//
	// example:
	//
	// APPROVED
	AccountReviewStatus              *string                `json:"AccountReviewStatus,omitempty" xml:"AccountReviewStatus,omitempty"`
	AuthInternationalRateEligibility map[string]interface{} `json:"AuthInternationalRateEligibility,omitempty" xml:"AuthInternationalRateEligibility,omitempty"`
	// The currency.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The ID of the WhatsApp Business account.
	//
	// example:
	//
	// 20393988393993***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The namespace of the message template.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	MessageTemplateNamespace *string `json:"MessageTemplateNamespace,omitempty" xml:"MessageTemplateNamespace,omitempty"`
	// The name of the WhatsApp Business account.
	//
	// example:
	//
	// Alibaba
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 示例值示例值
	PrimaryBusinessLocation *string `json:"PrimaryBusinessLocation,omitempty" xml:"PrimaryBusinessLocation,omitempty"`
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

func (s *QueryChatappBindWabaResponseBodyData) SetAuthInternationalRateEligibility(v map[string]interface{}) *QueryChatappBindWabaResponseBodyData {
	s.AuthInternationalRateEligibility = v
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

func (s *QueryChatappBindWabaResponseBodyData) SetPrimaryBusinessLocation(v string) *QueryChatappBindWabaResponseBodyData {
	s.PrimaryBusinessLocation = &v
	return s
}

type QueryChatappBindWabaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChatappBindWabaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// aksik93kdkkxmwol93939
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// example:
	//
	// VERIFIED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *QueryChatappPhoneNumbersRequest) SetStatus(v string) *QueryChatappPhoneNumbersRequest {
	s.Status = &v
	return s
}

type QueryChatappPhoneNumbersResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of phone numbers.
	PhoneNumbers []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBody) SetAccessDeniedDetail(v string) *QueryChatappPhoneNumbersResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The verification status of the phone number.
	//
	// example:
	//
	// VERIFIED
	CodeVerificationStatus *string `json:"CodeVerificationStatus,omitempty" xml:"CodeVerificationStatus,omitempty"`
	// The maximum number of messages that can be sent to users by using the phone number.
	//
	// example:
	//
	// TIER_10
	MessagingLimitTier *string `json:"MessagingLimitTier,omitempty" xml:"MessagingLimitTier,omitempty"`
	// The status of the business name.
	//
	// example:
	//
	// Approval
	NameStatus *string `json:"NameStatus,omitempty" xml:"NameStatus,omitempty"`
	// The review status of the new business name.
	//
	// example:
	//
	// Approval
	NewNameStatus *string `json:"NewNameStatus,omitempty" xml:"NewNameStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613800000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The quality rating of the phone number. Valid values:
	//
	// 	- **GREEN**
	//
	// 	- **YELLOW**
	//
	// 	- **RED**
	//
	// 	- **UNKNOWN**
	//
	// example:
	//
	// GREEN
	QualityRating *string `json:"QualityRating,omitempty" xml:"QualityRating,omitempty"`
	// The status of the phone number.
	//
	// 	- PENDING
	//
	// 	- DELETED
	//
	// 	- MIGRATED
	//
	// 	- BANNED
	//
	// 	- RESTRICTED
	//
	// 	- RATE_LIMITED
	//
	// 	- FLAGGED
	//
	// 	- CONNECTED
	//
	// 	- DISCONNECTED
	//
	// 	- UNKNOWN
	//
	// 	- UNVERIFIED
	//
	// example:
	//
	// CONNECTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The callback URL to which status reports are sent by using HTTP callbacks.
	//
	// example:
	//
	// https://ali.com/status
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// The status report notification queue.
	//
	// example:
	//
	// Alicom-Queue-****-ChatAppStatus
	StatusQueue *string `json:"StatusQueue,omitempty" xml:"StatusQueue,omitempty"`
	// The callback URL to which MO messages are sent by using HTTP callbacks.
	//
	// example:
	//
	// https://ali.com/inbound
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
	// The mobile originated (MO) message notification queue.
	//
	// example:
	//
	// Alicom-Queue-****-ChatAppInbound
	UpQueue *string `json:"UpQueue,omitempty" xml:"UpQueue,omitempty"`
	// The name of the company with which the phone number is associated.
	//
	// example:
	//
	// Alibaba
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChatappPhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryPhoneBusinessProfileRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2934839388494***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s QueryPhoneBusinessProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneBusinessProfileRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneBusinessProfileRequest) SetCustSpaceId(v string) *QueryPhoneBusinessProfileRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryPhoneBusinessProfileRequest) SetPhoneNumber(v string) *QueryPhoneBusinessProfileRequest {
	s.PhoneNumber = &v
	return s
}

type QueryPhoneBusinessProfileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryPhoneBusinessProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPhoneBusinessProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneBusinessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneBusinessProfileResponseBody) SetAccessDeniedDetail(v string) *QueryPhoneBusinessProfileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetCode(v string) *QueryPhoneBusinessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetData(v *QueryPhoneBusinessProfileResponseBodyData) *QueryPhoneBusinessProfileResponseBody {
	s.Data = v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetMessage(v string) *QueryPhoneBusinessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetRequestId(v string) *QueryPhoneBusinessProfileResponseBody {
	s.RequestId = &v
	return s
}

type QueryPhoneBusinessProfileResponseBodyData struct {
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// The address.
	//
	// example:
	//
	// Changsha
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The email address.
	//
	// example:
	//
	// aa@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The URL of the profile picture.
	//
	// example:
	//
	// https://....img
	ProfilePictureUrl *string `json:"ProfilePictureUrl,omitempty" xml:"ProfilePictureUrl,omitempty"`
	// The industry.
	//
	// example:
	//
	// Retail
	Vertical *string `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
	// The websites.
	Websites []*string `json:"Websites,omitempty" xml:"Websites,omitempty" type:"Repeated"`
}

func (s QueryPhoneBusinessProfileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneBusinessProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetAbout(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.About = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetAddress(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Address = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetDescription(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetEmail(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Email = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetProfilePictureUrl(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.ProfilePictureUrl = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetVertical(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Vertical = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetWebsites(v []*string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Websites = v
	return s
}

type QueryPhoneBusinessProfileResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPhoneBusinessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPhoneBusinessProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneBusinessProfileResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneBusinessProfileResponse) SetHeaders(v map[string]*string) *QueryPhoneBusinessProfileResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneBusinessProfileResponse) SetStatusCode(v int32) *QueryPhoneBusinessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponse) SetBody(v *QueryPhoneBusinessProfileResponseBody) *QueryPhoneBusinessProfileResponse {
	s.Body = v
	return s
}

type QueryWabaBusinessInfoRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ID of the WABA.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293848822333
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s QueryWabaBusinessInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWabaBusinessInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryWabaBusinessInfoRequest) SetCustSpaceId(v string) *QueryWabaBusinessInfoRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryWabaBusinessInfoRequest) SetWabaId(v string) *QueryWabaBusinessInfoRequest {
	s.WabaId = &v
	return s
}

type QueryWabaBusinessInfoResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business information about the WABA.
	Data *QueryWabaBusinessInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryWabaBusinessInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWabaBusinessInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWabaBusinessInfoResponseBody) SetAccessDeniedDetail(v string) *QueryWabaBusinessInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetCode(v string) *QueryWabaBusinessInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetData(v *QueryWabaBusinessInfoResponseBodyData) *QueryWabaBusinessInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetMessage(v string) *QueryWabaBusinessInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetRequestId(v string) *QueryWabaBusinessInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryWabaBusinessInfoResponseBodyData struct {
	// The ID of the business platform.
	//
	// example:
	//
	// 192882828733
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The name of the business platform.
	//
	// example:
	//
	// Alibaba
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The verification status.
	//
	// example:
	//
	// verified
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
	// The industry.
	//
	// example:
	//
	// Retail
	Vertical *string `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
}

func (s QueryWabaBusinessInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryWabaBusinessInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryWabaBusinessInfoResponseBodyData) SetBusinessId(v string) *QueryWabaBusinessInfoResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBodyData) SetBusinessName(v string) *QueryWabaBusinessInfoResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBodyData) SetVerificationStatus(v string) *QueryWabaBusinessInfoResponseBodyData {
	s.VerificationStatus = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBodyData) SetVertical(v string) *QueryWabaBusinessInfoResponseBodyData {
	s.Vertical = &v
	return s
}

type QueryWabaBusinessInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWabaBusinessInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWabaBusinessInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWabaBusinessInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryWabaBusinessInfoResponse) SetHeaders(v map[string]*string) *QueryWabaBusinessInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryWabaBusinessInfoResponse) SetStatusCode(v int32) *QueryWabaBusinessInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWabaBusinessInfoResponse) SetBody(v *QueryWabaBusinessInfoResponseBody) *QueryWabaBusinessInfoResponse {
	s.Body = v
	return s
}

type SendChatappMassMessageRequest struct {
	// The channel type. Valid values: whatsapp, viber, and line.
	//
	// This parameter is required.
	//
	// example:
	//
	// whatsapp
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The space ID of the user.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The ID of the WhatsApp Business account under the ISV account.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The fallback content.
	//
	// example:
	//
	// Fallback message
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The ID of the fallback strategy.
	//
	// example:
	//
	// S00001
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The fallback rule.
	//
	// >  Valid values:
	//
	// 	- undelivered: A fallback is triggered if the WhatsApp message is not delivered to clients. When the WhatsApp message is being sent, the template parameters are verified. If the parameters fail to pass the verification, the message fails to be sent. Whether the template and phone number are prohibited is not verified. By default, this value is used when FallBackRule is left empty.
	//
	// 	- sentfailed: A fallback is triggered even if the template parameters including variables fail to pass the verification. If the channelType, type, messageType, to, and from parameters fail to pass the verification, a fallback is not triggered.
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// The phone number of the message sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861387777****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The message type when the ChannelType parameter is set to viber. Valid values: promotion and transaction.
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language. For more information about language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The phone numbers to which the message is sent.
	//
	// This parameter is required.
	SenderList []*SendChatappMassMessageRequestSenderList `json:"SenderList,omitempty" xml:"SenderList,omitempty" type:"Repeated"`
	// The tag information when the ChannelType parameter is set to viber.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 100000001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The encoding of the message template.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca36bf5
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The timeout period for sending messages when the ChannelType parameter is set to viber. Valid values: 30 to 1209600. Unit: seconds.
	//
	// example:
	//
	// 50
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
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

func (s *SendChatappMassMessageRequest) SetFallBackDuration(v int32) *SendChatappMassMessageRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFallBackId(v string) *SendChatappMassMessageRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetFallBackRule(v string) *SendChatappMassMessageRequest {
	s.FallBackRule = &v
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

func (s *SendChatappMassMessageRequest) SetLabel(v string) *SendChatappMassMessageRequest {
	s.Label = &v
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

func (s *SendChatappMassMessageRequest) SetTag(v string) *SendChatappMassMessageRequest {
	s.Tag = &v
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

func (s *SendChatappMassMessageRequest) SetTemplateName(v string) *SendChatappMassMessageRequest {
	s.TemplateName = &v
	return s
}

func (s *SendChatappMassMessageRequest) SetTtl(v int64) *SendChatappMassMessageRequest {
	s.Ttl = &v
	return s
}

type SendChatappMassMessageRequestSenderList struct {
	FlowAction *SendChatappMassMessageRequestSenderListFlowAction `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
	// The payload.
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// The information about the product.
	ProductAction *SendChatappMassMessageRequestSenderListProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
	// The parameters of the template.
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The phone number to which the message is sent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861388988****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendChatappMassMessageRequestSenderList) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderList) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderList) SetFlowAction(v *SendChatappMassMessageRequestSenderListFlowAction) *SendChatappMassMessageRequestSenderList {
	s.FlowAction = v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) SetPayload(v []*string) *SendChatappMassMessageRequestSenderList {
	s.Payload = v
	return s
}

func (s *SendChatappMassMessageRequestSenderList) SetProductAction(v *SendChatappMassMessageRequestSenderListProductAction) *SendChatappMassMessageRequestSenderList {
	s.ProductAction = v
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

type SendChatappMassMessageRequestSenderListFlowAction struct {
	FlowActionData map[string]*string `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// example:
	//
	// 1122***
	FlowToken *string `json:"FlowToken,omitempty" xml:"FlowToken,omitempty"`
}

func (s SendChatappMassMessageRequestSenderListFlowAction) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderListFlowAction) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderListFlowAction) SetFlowActionData(v map[string]*string) *SendChatappMassMessageRequestSenderListFlowAction {
	s.FlowActionData = v
	return s
}

func (s *SendChatappMassMessageRequestSenderListFlowAction) SetFlowToken(v string) *SendChatappMassMessageRequestSenderListFlowAction {
	s.FlowToken = &v
	return s
}

type SendChatappMassMessageRequestSenderListProductAction struct {
	// The products. Up to 30 products and 10 categories can be added.
	Sections []*SendChatappMassMessageRequestSenderListProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// The retailer ID of the product.
	//
	// example:
	//
	// skkks999393
	ThumbnailProductRetailerId *string `json:"ThumbnailProductRetailerId,omitempty" xml:"ThumbnailProductRetailerId,omitempty"`
}

func (s SendChatappMassMessageRequestSenderListProductAction) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderListProductAction) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderListProductAction) SetSections(v []*SendChatappMassMessageRequestSenderListProductActionSections) *SendChatappMassMessageRequestSenderListProductAction {
	s.Sections = v
	return s
}

func (s *SendChatappMassMessageRequestSenderListProductAction) SetThumbnailProductRetailerId(v string) *SendChatappMassMessageRequestSenderListProductAction {
	s.ThumbnailProductRetailerId = &v
	return s
}

type SendChatappMassMessageRequestSenderListProductActionSections struct {
	// The products.
	ProductItems []*SendChatappMassMessageRequestSenderListProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// The name of the category.
	//
	// example:
	//
	// abcd
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SendChatappMassMessageRequestSenderListProductActionSections) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderListProductActionSections) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderListProductActionSections) SetProductItems(v []*SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) *SendChatappMassMessageRequestSenderListProductActionSections {
	s.ProductItems = v
	return s
}

func (s *SendChatappMassMessageRequestSenderListProductActionSections) SetTitle(v string) *SendChatappMassMessageRequestSenderListProductActionSections {
	s.Title = &v
	return s
}

type SendChatappMassMessageRequestSenderListProductActionSectionsProductItems struct {
	// The retailer ID of the product.
	//
	// example:
	//
	// ksi399d8
	ProductRetailerId *string `json:"ProductRetailerId,omitempty" xml:"ProductRetailerId,omitempty"`
}

func (s SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageRequestSenderListProductActionSectionsProductItems) SetProductRetailerId(v string) *SendChatappMassMessageRequestSenderListProductActionSectionsProductItems {
	s.ProductRetailerId = &v
	return s
}

type SendChatappMassMessageShrinkRequest struct {
	// The channel type. Valid values: whatsapp, viber, and line.
	//
	// This parameter is required.
	//
	// example:
	//
	// whatsapp
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The space ID of the user.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The ID of the WhatsApp Business account under the ISV account.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The fallback content.
	//
	// example:
	//
	// Fallback message
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The ID of the fallback strategy.
	//
	// example:
	//
	// S00001
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// The fallback rule.
	//
	// >  Valid values:
	//
	// 	- undelivered: A fallback is triggered if the WhatsApp message is not delivered to clients. When the WhatsApp message is being sent, the template parameters are verified. If the parameters fail to pass the verification, the message fails to be sent. Whether the template and phone number are prohibited is not verified. By default, this value is used when FallBackRule is left empty.
	//
	// 	- sentfailed: A fallback is triggered even if the template parameters including variables fail to pass the verification. If the channelType, type, messageType, to, and from parameters fail to pass the verification, a fallback is not triggered.
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// The phone number of the message sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// 861387777****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The message type when the ChannelType parameter is set to viber. Valid values: promotion and transaction.
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language. For more information about language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The phone numbers to which the message is sent.
	//
	// This parameter is required.
	SenderListShrink *string `json:"SenderList,omitempty" xml:"SenderList,omitempty"`
	// The tag information when the ChannelType parameter is set to viber.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 100000001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The encoding of the message template.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca36bf5
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The timeout period for sending messages when the ChannelType parameter is set to viber. Valid values: 30 to 1209600. Unit: seconds.
	//
	// example:
	//
	// 50
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
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

func (s *SendChatappMassMessageShrinkRequest) SetFallBackDuration(v int32) *SendChatappMassMessageShrinkRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackId(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackRule(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackRule = &v
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

func (s *SendChatappMassMessageShrinkRequest) SetLabel(v string) *SendChatappMassMessageShrinkRequest {
	s.Label = &v
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

func (s *SendChatappMassMessageShrinkRequest) SetTag(v string) *SendChatappMassMessageShrinkRequest {
	s.Tag = &v
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

func (s *SendChatappMassMessageShrinkRequest) SetTemplateName(v string) *SendChatappMassMessageShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTtl(v int64) *SendChatappMassMessageShrinkRequest {
	s.Ttl = &v
	return s
}

type SendChatappMassMessageResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the group of messages.
	//
	// example:
	//
	// 890000010002****
	GroupMessageId *string `json:"GroupMessageId,omitempty" xml:"GroupMessageId,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendChatappMassMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMassMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageResponseBody) SetAccessDeniedDetail(v string) *SendChatappMassMessageResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendChatappMassMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// 	- **whatsapp**
	//
	// 	- **viber**
	//
	// 	- line. The feature that ChatAPP sends messages by using Line is under development.
	//
	// This parameter is required.
	//
	// example:
	//
	// whatsapp
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The content of the message.
	//
	// **Usage notes when you set the ChannelType parameter to whatsapp:**
	//
	// 	- When you set the **MessageType*	- parameter to **text**, the **text*	- parameter is required and the **caption*	- parameter cannot be specified.
	//
	// 	- When you set the **MessageType*	- parameter to **image**, the **link*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **video**, the **link*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **audio**, the **link*	- parameter is required and the **caption*	- parameter is invalid.
	//
	// 	- When you set the **MessageType*	- parameter to **document**, the **link*	- and **fileName*	- parameters are required and the **caption*	- parameter is invalid.
	//
	// 	- When you set the **MessageType*	- parameter to **interactive**, the **type*	- and **action*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **contacts**, the **name*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **location**, the **longitude*	- and **latitude*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **sticker**, the **link*	- parameter is required, and the **caption*	- and **fileName*	- parameters are invalid.
	//
	// 	- When you set the **MessageType*	- parameter to **reaction**, the **messageId*	- and **emoji*	- parameters are required.
	//
	// **Usage notes when you set the ChannelType parameter to viber:**
	//
	// 	- When you set the **MessageType*	- parameter to **text**, the **text*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **image**, the **link*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **video**, the **link**, **thumbnail**, **fileSize**, and **duration*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **document**, the **link**, **fileName**, and **fileType*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **text_button**, the **text**, **caption**, and **action*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **text_image_button**, the **text**, **link**, **caption**, and **action*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **text_video**, the **text**, **link**, **thumbnail**, **fileSize**, and **duration*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **text_video_button**, the **text**, **link**, **thumbnail**, **fileSize**, **duration**, and **caption*	- parameters are required. The **action*	- parameter is invalid.
	//
	// example:
	//
	// {\\"text\\": \\"hello whatsapp\\", \\"link\\": \\"\\", \\"caption\\": \\"\\", \\"fileName\\": \\"\\" }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the message to reply to.
	//
	// example:
	//
	// 61851ccb2f1365b16aee****
	ContextMessageId *string `json:"ContextMessageId,omitempty" xml:"ContextMessageId,omitempty"`
	// The space ID of the user.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The ID of the WhatsApp account that you register.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The content of the fallback message.
	//
	// example:
	//
	// This is a fallback message.
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// Specifies the period of time after which the fallback Short Message Service (SMS) message is sent if the message receipt that indicates the message is delivered to customers is not received. If this parameter is left empty, the fallback SMS message is sent only when the **message fails to be sent*	- or **the message receipt that indicates the message is not delivered to customers*	- is received. Valid values: 60 to 43200. Unit: seconds.
	//
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The ID of the fallback strategy. You can create a fallback strategy and view the information in the console.
	//
	// example:
	//
	// S_000001
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// 回落规则。
	//
	// > 取值范围
	//
	// > - undelivered  消息不能发送到端时回落（在发送状态时模板、参数需要校验通过，模板被封、号码被封等不做校验）。参数值为空时默认使用此规则
	//
	// > - sentFailed  消息在校验模板、模板变量等参数时，校验不通过也会回落。只会强校验channelType, type, messageType, to, from(是否存在) 几个参数。
	//
	// <props="china">
	//
	// > 中国站此字段无效
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// Flow发送数据
	FlowAction *SendChatappMessageRequestFlowAction `json:"FlowAction,omitempty" xml:"FlowAction,omitempty" type:"Struct"`
	// The phone number of the message sender.
	//
	// > You can specify a mobile phone number that is registered for a WhatsApp account and is approved in the ChatAPP console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1360000****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The type of the Viber message. This parameter is required if you set the ChannelType parameter to viber. Valid values: promotion and transaction.
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language that is used in the message template. This parameter is required only if you set the Type parameter to **template**. For more information about language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The specific type of the message. This parameter is required only if you set the Type parameter to **message**.
	//
	// **Valid values of MessageType when you set the ChannelType parameter to whatsapp:**
	//
	// 	- **text**: a text message.
	//
	// 	- **image**: an image message.
	//
	// 	- **video**: a video message.
	//
	// 	- **audio**: an audio message.
	//
	// 	- **document**: a document message.
	//
	// 	- **interactive**: an interactive message.
	//
	// 	- **contacts**: a contact message.
	//
	// 	- **location**: a location message.
	//
	// 	- **sticker**: a sticker message.
	//
	// 	- **reaction**: a reaction message.
	//
	// **Valid values of MessageType when you set the ChannelType parameter to viber:**
	//
	// 	- **text**: a text message.
	//
	// 	- **image**: an image message.
	//
	// 	- **video**: a video message.
	//
	// 	- **document**: a document message.
	//
	// 	- **text_button**: a message that contains the text and button media objects.
	//
	// 	- **text_image_button**: a message that contains multiple media objects, including the text, image, and button.
	//
	// 	- **text_video**: a message that contains the text and video media objects.
	//
	// 	- **text_video_button**: a message that contains multiple media objects, including text, video, and button.
	//
	// 	- **text_image**: a message that contains the text and image media objects.
	//
	// > For more information, see [Parameters of a message template](https://help.aliyun.com/document_detail/454530.html).
	//
	// example:
	//
	// text
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The payload of the button.
	//
	// example:
	//
	// payloadtext1,payloadtext2,payloadtext3
	Payload []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// The information about the products included in the WhatsApp catalog message or multi-product message (MPM).
	ProductAction *SendChatappMessageRequestProductAction `json:"ProductAction,omitempty" xml:"ProductAction,omitempty" type:"Struct"`
	// The tag information of the Viber message.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 100000001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The code of the message template. This parameter is required only if you set the Type parameter to **template**.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The variables of the message template.
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The phone number that receives the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The tracking ID of the Viber message.
	//
	// example:
	//
	// tracking_id:123456
	TrackingData *string `json:"TrackingData,omitempty" xml:"TrackingData,omitempty"`
	// The timeout period for sending the Viber message. Valid values: 30 to 1209600. Unit: seconds.
	//
	// example:
	//
	// 50
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **template**: a template message. A template message is sent based on a template that is created in the ChatAPP console and is approved. You can send template messages at any time based on your business requirements.
	//
	// 	- **message**: a custom message. You can send a custom message to a user only within 24 hours after you receive the last message from the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// template
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

func (s *SendChatappMessageRequest) SetFallBackDuration(v int32) *SendChatappMessageRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackId(v string) *SendChatappMessageRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackRule(v string) *SendChatappMessageRequest {
	s.FallBackRule = &v
	return s
}

func (s *SendChatappMessageRequest) SetFlowAction(v *SendChatappMessageRequestFlowAction) *SendChatappMessageRequest {
	s.FlowAction = v
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

func (s *SendChatappMessageRequest) SetProductAction(v *SendChatappMessageRequestProductAction) *SendChatappMessageRequest {
	s.ProductAction = v
	return s
}

func (s *SendChatappMessageRequest) SetTag(v string) *SendChatappMessageRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMessageRequest) SetTaskId(v string) *SendChatappMessageRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateCode(v string) *SendChatappMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateName(v string) *SendChatappMessageRequest {
	s.TemplateName = &v
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

type SendChatappMessageRequestFlowAction struct {
	// flow默认参数
	FlowActionData map[string]*string `json:"FlowActionData,omitempty" xml:"FlowActionData,omitempty"`
	// flow token信息
	//
	// example:
	//
	// 1122***
	FlowToken *string `json:"FlowToken,omitempty" xml:"FlowToken,omitempty"`
}

func (s SendChatappMessageRequestFlowAction) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageRequestFlowAction) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequestFlowAction) SetFlowActionData(v map[string]*string) *SendChatappMessageRequestFlowAction {
	s.FlowActionData = v
	return s
}

func (s *SendChatappMessageRequestFlowAction) SetFlowToken(v string) *SendChatappMessageRequestFlowAction {
	s.FlowToken = &v
	return s
}

type SendChatappMessageRequestProductAction struct {
	// The products. Up to 30 products can be added. The products can be divided into up to 10 categories.
	Sections []*SendChatappMessageRequestProductActionSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// The retailer ID of the product.
	//
	// example:
	//
	// S238SK
	ThumbnailProductRetailerId *string `json:"ThumbnailProductRetailerId,omitempty" xml:"ThumbnailProductRetailerId,omitempty"`
}

func (s SendChatappMessageRequestProductAction) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageRequestProductAction) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequestProductAction) SetSections(v []*SendChatappMessageRequestProductActionSections) *SendChatappMessageRequestProductAction {
	s.Sections = v
	return s
}

func (s *SendChatappMessageRequestProductAction) SetThumbnailProductRetailerId(v string) *SendChatappMessageRequestProductAction {
	s.ThumbnailProductRetailerId = &v
	return s
}

type SendChatappMessageRequestProductActionSections struct {
	// The products.
	ProductItems []*SendChatappMessageRequestProductActionSectionsProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Repeated"`
	// The name of the category.
	//
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SendChatappMessageRequestProductActionSections) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageRequestProductActionSections) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequestProductActionSections) SetProductItems(v []*SendChatappMessageRequestProductActionSectionsProductItems) *SendChatappMessageRequestProductActionSections {
	s.ProductItems = v
	return s
}

func (s *SendChatappMessageRequestProductActionSections) SetTitle(v string) *SendChatappMessageRequestProductActionSections {
	s.Title = &v
	return s
}

type SendChatappMessageRequestProductActionSectionsProductItems struct {
	// The retailer ID of the product.
	//
	// example:
	//
	// 9I39E9E
	ProductRetailerId *string `json:"ProductRetailerId,omitempty" xml:"ProductRetailerId,omitempty"`
}

func (s SendChatappMessageRequestProductActionSectionsProductItems) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageRequestProductActionSectionsProductItems) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequestProductActionSectionsProductItems) SetProductRetailerId(v string) *SendChatappMessageRequestProductActionSectionsProductItems {
	s.ProductRetailerId = &v
	return s
}

type SendChatappMessageShrinkRequest struct {
	// The type of the message channel. Valid values:
	//
	// 	- **whatsapp**
	//
	// 	- **viber**
	//
	// 	- line. The feature that ChatAPP sends messages by using Line is under development.
	//
	// This parameter is required.
	//
	// example:
	//
	// whatsapp
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The content of the message.
	//
	// **Usage notes when you set the ChannelType parameter to whatsapp:**
	//
	// 	- When you set the **MessageType*	- parameter to **text**, the **text*	- parameter is required and the **caption*	- parameter cannot be specified.
	//
	// 	- When you set the **MessageType*	- parameter to **image**, the **link*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **video**, the **link*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **audio**, the **link*	- parameter is required and the **caption*	- parameter is invalid.
	//
	// 	- When you set the **MessageType*	- parameter to **document**, the **link*	- and **fileName*	- parameters are required and the **caption*	- parameter is invalid.
	//
	// 	- When you set the **MessageType*	- parameter to **interactive**, the **type*	- and **action*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **contacts**, the **name*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **location**, the **longitude*	- and **latitude*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **sticker**, the **link*	- parameter is required, and the **caption*	- and **fileName*	- parameters are invalid.
	//
	// 	- When you set the **MessageType*	- parameter to **reaction**, the **messageId*	- and **emoji*	- parameters are required.
	//
	// **Usage notes when you set the ChannelType parameter to viber:**
	//
	// 	- When you set the **MessageType*	- parameter to **text**, the **text*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **image**, the **link*	- parameter is required.
	//
	// 	- When you set the **MessageType*	- parameter to **video**, the **link**, **thumbnail**, **fileSize**, and **duration*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **document**, the **link**, **fileName**, and **fileType*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **text_button**, the **text**, **caption**, and **action*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **text_image_button**, the **text**, **link**, **caption**, and **action*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **text_video**, the **text**, **link**, **thumbnail**, **fileSize**, and **duration*	- parameters are required.
	//
	// 	- When you set the **MessageType*	- parameter to **text_video_button**, the **text**, **link**, **thumbnail**, **fileSize**, **duration**, and **caption*	- parameters are required. The **action*	- parameter is invalid.
	//
	// example:
	//
	// {\\"text\\": \\"hello whatsapp\\", \\"link\\": \\"\\", \\"caption\\": \\"\\", \\"fileName\\": \\"\\" }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the message to reply to.
	//
	// example:
	//
	// 61851ccb2f1365b16aee****
	ContextMessageId *string `json:"ContextMessageId,omitempty" xml:"ContextMessageId,omitempty"`
	// The space ID of the user.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
	// The ID of the WhatsApp account that you register.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The content of the fallback message.
	//
	// example:
	//
	// This is a fallback message.
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// Specifies the period of time after which the fallback Short Message Service (SMS) message is sent if the message receipt that indicates the message is delivered to customers is not received. If this parameter is left empty, the fallback SMS message is sent only when the **message fails to be sent*	- or **the message receipt that indicates the message is not delivered to customers*	- is received. Valid values: 60 to 43200. Unit: seconds.
	//
	// example:
	//
	// 120
	FallBackDuration *int32 `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// The ID of the fallback strategy. You can create a fallback strategy and view the information in the console.
	//
	// example:
	//
	// S_000001
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// 回落规则。
	//
	// > 取值范围
	//
	// > - undelivered  消息不能发送到端时回落（在发送状态时模板、参数需要校验通过，模板被封、号码被封等不做校验）。参数值为空时默认使用此规则
	//
	// > - sentFailed  消息在校验模板、模板变量等参数时，校验不通过也会回落。只会强校验channelType, type, messageType, to, from(是否存在) 几个参数。
	//
	// <props="china">
	//
	// > 中国站此字段无效
	//
	// example:
	//
	// undelivered
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// Flow发送数据
	FlowActionShrink *string `json:"FlowAction,omitempty" xml:"FlowAction,omitempty"`
	// The phone number of the message sender.
	//
	// > You can specify a mobile phone number that is registered for a WhatsApp account and is approved in the ChatAPP console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1360000****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The type of the Viber message. This parameter is required if you set the ChannelType parameter to viber. Valid values: promotion and transaction.
	//
	// example:
	//
	// promotion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The language that is used in the message template. This parameter is required only if you set the Type parameter to **template**. For more information about language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The specific type of the message. This parameter is required only if you set the Type parameter to **message**.
	//
	// **Valid values of MessageType when you set the ChannelType parameter to whatsapp:**
	//
	// 	- **text**: a text message.
	//
	// 	- **image**: an image message.
	//
	// 	- **video**: a video message.
	//
	// 	- **audio**: an audio message.
	//
	// 	- **document**: a document message.
	//
	// 	- **interactive**: an interactive message.
	//
	// 	- **contacts**: a contact message.
	//
	// 	- **location**: a location message.
	//
	// 	- **sticker**: a sticker message.
	//
	// 	- **reaction**: a reaction message.
	//
	// **Valid values of MessageType when you set the ChannelType parameter to viber:**
	//
	// 	- **text**: a text message.
	//
	// 	- **image**: an image message.
	//
	// 	- **video**: a video message.
	//
	// 	- **document**: a document message.
	//
	// 	- **text_button**: a message that contains the text and button media objects.
	//
	// 	- **text_image_button**: a message that contains multiple media objects, including the text, image, and button.
	//
	// 	- **text_video**: a message that contains the text and video media objects.
	//
	// 	- **text_video_button**: a message that contains multiple media objects, including text, video, and button.
	//
	// 	- **text_image**: a message that contains the text and image media objects.
	//
	// > For more information, see [Parameters of a message template](https://help.aliyun.com/document_detail/454530.html).
	//
	// example:
	//
	// text
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The payload of the button.
	//
	// example:
	//
	// payloadtext1,payloadtext2,payloadtext3
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The information about the products included in the WhatsApp catalog message or multi-product message (MPM).
	ProductActionShrink *string `json:"ProductAction,omitempty" xml:"ProductAction,omitempty"`
	// The tag information of the Viber message.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 100000001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The code of the message template. This parameter is required only if you set the Type parameter to **template**.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The variables of the message template.
	TemplateParamsShrink *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The phone number that receives the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The tracking ID of the Viber message.
	//
	// example:
	//
	// tracking_id:123456
	TrackingData *string `json:"TrackingData,omitempty" xml:"TrackingData,omitempty"`
	// The timeout period for sending the Viber message. Valid values: 30 to 1209600. Unit: seconds.
	//
	// example:
	//
	// 50
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **template**: a template message. A template message is sent based on a template that is created in the ChatAPP console and is approved. You can send template messages at any time based on your business requirements.
	//
	// 	- **message**: a custom message. You can send a custom message to a user only within 24 hours after you receive the last message from the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// template
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

func (s *SendChatappMessageShrinkRequest) SetFallBackDuration(v int32) *SendChatappMessageShrinkRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackId(v string) *SendChatappMessageShrinkRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackRule(v string) *SendChatappMessageShrinkRequest {
	s.FallBackRule = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFlowActionShrink(v string) *SendChatappMessageShrinkRequest {
	s.FlowActionShrink = &v
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

func (s *SendChatappMessageShrinkRequest) SetProductActionShrink(v string) *SendChatappMessageShrinkRequest {
	s.ProductActionShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTag(v string) *SendChatappMessageShrinkRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTaskId(v string) *SendChatappMessageShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateCode(v string) *SendChatappMessageShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateName(v string) *SendChatappMessageShrinkRequest {
	s.TemplateName = &v
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
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the message that was sent.
	//
	// example:
	//
	// 61851ccb2f1365b16aee****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendChatappMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SubmitIsvCustomerTermsRequest struct {
	// The use scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// Marketing products
	BusinessDesc *string `json:"BusinessDesc,omitempty" xml:"BusinessDesc,omitempty"`
	// The email address of your business.
	//
	// This parameter is required.
	//
	// example:
	//
	// partner@aliyun.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// The country code.
	//
	// This parameter is required.
	//
	// example:
	//
	// CN
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The display name of your business.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	CustName *string `json:"CustName,omitempty" xml:"CustName,omitempty"`
	// The space ID of the user under the ISV account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV or Client Agreement.
	//
	// > Before you upload files to Object Storage Service (OSS) servers, you must call the GetChatappUploadAuthorization operation to obtain the authentication information required to upload files. You can use the SDK provided by OSS to upload files. When you upload a file, you must set the key parameter value. To set the value for the key parameter, concatenate the value of the Dir parameter and the file name by using a forward slash (/). You can obtain the value of the Dir parameter by calling the GetChatappUploadAuthorization operation.
	//
	// > The value of this parameter is the name of the uploaded file.
	//
	// This parameter is required.
	//
	// example:
	//
	// isvTerms.pdf
	IsvTerms *string `json:"IsvTerms,omitempty" xml:"IsvTerms,omitempty"`
	// The address of your business.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	OfficeAddress *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
}

func (s SubmitIsvCustomerTermsRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitIsvCustomerTermsRequest) GoString() string {
	return s.String()
}

func (s *SubmitIsvCustomerTermsRequest) SetBusinessDesc(v string) *SubmitIsvCustomerTermsRequest {
	s.BusinessDesc = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetContactMail(v string) *SubmitIsvCustomerTermsRequest {
	s.ContactMail = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetCountryId(v string) *SubmitIsvCustomerTermsRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetCustName(v string) *SubmitIsvCustomerTermsRequest {
	s.CustName = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetCustSpaceId(v string) *SubmitIsvCustomerTermsRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetIsvTerms(v string) *SubmitIsvCustomerTermsRequest {
	s.IsvTerms = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetOfficeAddress(v string) *SubmitIsvCustomerTermsRequest {
	s.OfficeAddress = &v
	return s
}

type SubmitIsvCustomerTermsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitIsvCustomerTermsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitIsvCustomerTermsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIsvCustomerTermsResponseBody) SetAccessDeniedDetail(v string) *SubmitIsvCustomerTermsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponseBody) SetCode(v string) *SubmitIsvCustomerTermsResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponseBody) SetMessage(v string) *SubmitIsvCustomerTermsResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponseBody) SetRequestId(v string) *SubmitIsvCustomerTermsResponseBody {
	s.RequestId = &v
	return s
}

type SubmitIsvCustomerTermsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIsvCustomerTermsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIsvCustomerTermsResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitIsvCustomerTermsResponse) GoString() string {
	return s.String()
}

func (s *SubmitIsvCustomerTermsResponse) SetHeaders(v map[string]*string) *SubmitIsvCustomerTermsResponse {
	s.Headers = v
	return s
}

func (s *SubmitIsvCustomerTermsResponse) SetStatusCode(v int32) *SubmitIsvCustomerTermsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponse) SetBody(v *SubmitIsvCustomerTermsResponseBody) *SubmitIsvCustomerTermsResponse {
	s.Body = v
	return s
}

type UpdateAccountWebhookRequest struct {
	// The space ID of the user under the ISV account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Specifies whether to use HTTP callbacks to receive message receipts. Valid values:
	//
	// 	- Y: indicates that HTTP callbacks are used to receive receipts.
	//
	// 	- N: indicates that HTTP callbacks are not used to receive receipts.
	//
	// example:
	//
	// Y
	HttpFlag *string `json:"HttpFlag,omitempty" xml:"HttpFlag,omitempty"`
	// Specifies whether to use Message Service (MNS) queues to receive receipts. Valid values:
	//
	// 	- Y: indicates that MNS queues are used to receive receipts.
	//
	// 	- N: indicates that MNS queues are not used to receive receipts.
	//
	// example:
	//
	// N
	QueueFlag *string `json:"QueueFlag,omitempty" xml:"QueueFlag,omitempty"`
	// The callback URL to which status reports are sent by using HTTP callbacks.
	//
	// example:
	//
	// http://www.aliyun.com
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
}

func (s UpdateAccountWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountWebhookRequest) SetCustSpaceId(v string) *UpdateAccountWebhookRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetHttpFlag(v string) *UpdateAccountWebhookRequest {
	s.HttpFlag = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetQueueFlag(v string) *UpdateAccountWebhookRequest {
	s.QueueFlag = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetStatusCallbackUrl(v string) *UpdateAccountWebhookRequest {
	s.StatusCallbackUrl = &v
	return s
}

type UpdateAccountWebhookResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccountWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountWebhookResponseBody) SetAccessDeniedDetail(v string) *UpdateAccountWebhookResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAccountWebhookResponseBody) SetCode(v string) *UpdateAccountWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAccountWebhookResponseBody) SetMessage(v string) *UpdateAccountWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAccountWebhookResponseBody) SetRequestId(v string) *UpdateAccountWebhookResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAccountWebhookResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountWebhookResponse) SetHeaders(v map[string]*string) *UpdateAccountWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountWebhookResponse) SetStatusCode(v int32) *UpdateAccountWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountWebhookResponse) SetBody(v *UpdateAccountWebhookResponseBody) *UpdateAccountWebhookResponse {
	s.Body = v
	return s
}

type UpdateCommerceSettingRequest struct {
	// Specifies whether to display the shopping cart button.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	CartEnable *bool `json:"CartEnable,omitempty" xml:"CartEnable,omitempty"`
	// Specifies whether to display the catalog button.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CatalogVisible *bool `json:"CatalogVisible,omitempty" xml:"CatalogVisible,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1380000****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s UpdateCommerceSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommerceSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommerceSettingRequest) SetCartEnable(v bool) *UpdateCommerceSettingRequest {
	s.CartEnable = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetCatalogVisible(v bool) *UpdateCommerceSettingRequest {
	s.CatalogVisible = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetCustSpaceId(v string) *UpdateCommerceSettingRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetPhoneNumber(v string) *UpdateCommerceSettingRequest {
	s.PhoneNumber = &v
	return s
}

type UpdateCommerceSettingResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCommerceSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommerceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCommerceSettingResponseBody) SetCode(v string) *UpdateCommerceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCommerceSettingResponseBody) SetMessage(v string) *UpdateCommerceSettingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCommerceSettingResponseBody) SetRequestId(v string) *UpdateCommerceSettingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCommerceSettingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCommerceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCommerceSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommerceSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateCommerceSettingResponse) SetHeaders(v map[string]*string) *UpdateCommerceSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateCommerceSettingResponse) SetStatusCode(v int32) *UpdateCommerceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCommerceSettingResponse) SetBody(v *UpdateCommerceSettingResponseBody) *UpdateCommerceSettingResponse {
	s.Body = v
	return s
}

type UpdateFlowJSONAssetRequest struct {
	// example:
	//
	// 9399393
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun/json.json
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// Flow ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// flow_001
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s UpdateFlowJSONAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowJSONAssetRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowJSONAssetRequest) SetCustSpaceId(v string) *UpdateFlowJSONAssetRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateFlowJSONAssetRequest) SetFilePath(v string) *UpdateFlowJSONAssetRequest {
	s.FilePath = &v
	return s
}

func (s *UpdateFlowJSONAssetRequest) SetFlowId(v string) *UpdateFlowJSONAssetRequest {
	s.FlowId = &v
	return s
}

type UpdateFlowJSONAssetResponseBody struct {
	// example:
	//
	// OK
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *UpdateFlowJSONAssetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFlowJSONAssetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowJSONAssetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowJSONAssetResponseBody) SetCode(v string) *UpdateFlowJSONAssetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFlowJSONAssetResponseBody) SetData(v *UpdateFlowJSONAssetResponseBodyData) *UpdateFlowJSONAssetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFlowJSONAssetResponseBody) SetMessage(v string) *UpdateFlowJSONAssetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFlowJSONAssetResponseBody) SetRequestId(v string) *UpdateFlowJSONAssetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFlowJSONAssetResponseBodyData struct {
	// Flow ID。
	//
	// example:
	//
	// 84848847****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s UpdateFlowJSONAssetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowJSONAssetResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateFlowJSONAssetResponseBodyData) SetFlowId(v string) *UpdateFlowJSONAssetResponseBodyData {
	s.FlowId = &v
	return s
}

type UpdateFlowJSONAssetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlowJSONAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowJSONAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowJSONAssetResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowJSONAssetResponse) SetHeaders(v map[string]*string) *UpdateFlowJSONAssetResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowJSONAssetResponse) SetStatusCode(v int32) *UpdateFlowJSONAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowJSONAssetResponse) SetBody(v *UpdateFlowJSONAssetResponseBody) *UpdateFlowJSONAssetResponse {
	s.Body = v
	return s
}

type UpdatePhoneEncryptionPublicKeyRequest struct {
	// example:
	//
	// 399382882
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY-----
	//
	// AAA
	//
	// BBB
	//
	// CCC
	//
	// DDD
	//
	// EEE
	//
	// FFF
	//
	// GGG
	//
	// -----END PUBLIC KEY-----
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" xml:"EncryptionPublicKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86138000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s UpdatePhoneEncryptionPublicKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneEncryptionPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetCustSpaceId(v string) *UpdatePhoneEncryptionPublicKeyRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetEncryptionPublicKey(v string) *UpdatePhoneEncryptionPublicKeyRequest {
	s.EncryptionPublicKey = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetPhoneNumber(v string) *UpdatePhoneEncryptionPublicKeyRequest {
	s.PhoneNumber = &v
	return s
}

type UpdatePhoneEncryptionPublicKeyResponseBody struct {
	// example:
	//
	// OK
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhoneEncryptionPublicKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneEncryptionPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) SetCode(v string) *UpdatePhoneEncryptionPublicKeyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) SetMessage(v string) *UpdatePhoneEncryptionPublicKeyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) SetRequestId(v string) *UpdatePhoneEncryptionPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePhoneEncryptionPublicKeyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePhoneEncryptionPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePhoneEncryptionPublicKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneEncryptionPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) SetHeaders(v map[string]*string) *UpdatePhoneEncryptionPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) SetStatusCode(v int32) *UpdatePhoneEncryptionPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) SetBody(v *UpdatePhoneEncryptionPublicKeyResponseBody) *UpdatePhoneEncryptionPublicKeyResponse {
	s.Body = v
	return s
}

type UpdatePhoneMessageQrdlRequest struct {
	// example:
	//
	// 9383884
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PNG
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hello
	PrefilledMessage *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 29338838
	QrdlCode *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
}

func (s UpdatePhoneMessageQrdlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhoneMessageQrdlRequest) SetCustSpaceId(v string) *UpdatePhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetGenerateQrImage(v string) *UpdatePhoneMessageQrdlRequest {
	s.GenerateQrImage = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetPhoneNumber(v string) *UpdatePhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetPrefilledMessage(v string) *UpdatePhoneMessageQrdlRequest {
	s.PrefilledMessage = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetQrdlCode(v string) *UpdatePhoneMessageQrdlRequest {
	s.QrdlCode = &v
	return s
}

type UpdatePhoneMessageQrdlResponseBody struct {
	// example:
	//
	// OK
	Code    *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *UpdatePhoneMessageQrdlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// 1612C226-E271-4CFE-9F18-4066D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhoneMessageQrdlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneMessageQrdlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetCode(v string) *UpdatePhoneMessageQrdlResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetData(v *UpdatePhoneMessageQrdlResponseBodyData) *UpdatePhoneMessageQrdlResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetMessage(v string) *UpdatePhoneMessageQrdlResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetRequestId(v string) *UpdatePhoneMessageQrdlResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePhoneMessageQrdlResponseBodyData struct {
	// example:
	//
	// https://wa.msg/
	DeepLinkUrl *string `json:"DeepLinkUrl,omitempty" xml:"DeepLinkUrl,omitempty"`
	// example:
	//
	// PNG
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	// example:
	//
	// 8613800
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// Hello
	PrefilledMessage *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
	// example:
	//
	// https://img.png
	QrImageUrl *string `json:"QrImageUrl,omitempty" xml:"QrImageUrl,omitempty"`
	// example:
	//
	// DEDEE998
	QrdlCode *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
}

func (s UpdatePhoneMessageQrdlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneMessageQrdlResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetDeepLinkUrl(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.DeepLinkUrl = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetGenerateQrImage(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.GenerateQrImage = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetPhoneNumber(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetPrefilledMessage(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.PrefilledMessage = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetQrImageUrl(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.QrImageUrl = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetQrdlCode(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.QrdlCode = &v
	return s
}

type UpdatePhoneMessageQrdlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePhoneMessageQrdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePhoneMessageQrdlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneMessageQrdlResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhoneMessageQrdlResponse) SetHeaders(v map[string]*string) *UpdatePhoneMessageQrdlResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhoneMessageQrdlResponse) SetStatusCode(v int32) *UpdatePhoneMessageQrdlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponse) SetBody(v *UpdatePhoneMessageQrdlResponseBody) *UpdatePhoneMessageQrdlResponse {
	s.Body = v
	return s
}

type UpdatePhoneWebhookRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// Y
	HttpFlag *string `json:"HttpFlag,omitempty" xml:"HttpFlag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// N
	QueueFlag *string `json:"QueueFlag,omitempty" xml:"QueueFlag,omitempty"`
	// example:
	//
	// http://www.aliyun.com
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// example:
	//
	// http://aliyun.com
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
}

func (s UpdatePhoneWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhoneWebhookRequest) SetCustSpaceId(v string) *UpdatePhoneWebhookRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetHttpFlag(v string) *UpdatePhoneWebhookRequest {
	s.HttpFlag = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetPhoneNumber(v string) *UpdatePhoneWebhookRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetQueueFlag(v string) *UpdatePhoneWebhookRequest {
	s.QueueFlag = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetStatusCallbackUrl(v string) *UpdatePhoneWebhookRequest {
	s.StatusCallbackUrl = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetUpCallbackUrl(v string) *UpdatePhoneWebhookRequest {
	s.UpCallbackUrl = &v
	return s
}

type UpdatePhoneWebhookResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhoneWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhoneWebhookResponseBody) SetAccessDeniedDetail(v string) *UpdatePhoneWebhookResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdatePhoneWebhookResponseBody) SetCode(v string) *UpdatePhoneWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePhoneWebhookResponseBody) SetMessage(v string) *UpdatePhoneWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePhoneWebhookResponseBody) SetRequestId(v string) *UpdatePhoneWebhookResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePhoneWebhookResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePhoneWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePhoneWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhoneWebhookResponse) SetHeaders(v map[string]*string) *UpdatePhoneWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhoneWebhookResponse) SetStatusCode(v int32) *UpdatePhoneWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePhoneWebhookResponse) SetBody(v *UpdatePhoneWebhookResponseBody) *UpdatePhoneWebhookResponse {
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

// Summary:
//
// Adds a phone number for a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddChatappPhoneNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddChatappPhoneNumberResponse
func (client *Client) AddChatappPhoneNumberWithOptions(request *AddChatappPhoneNumberRequest, runtime *util.RuntimeOptions) (_result *AddChatappPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cc)) {
		body["Cc"] = request.Cc
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PreValidateId)) {
		body["PreValidateId"] = request.PreValidateId
	}

	if !tea.BoolValue(util.IsUnset(request.VerifiedName)) {
		body["VerifiedName"] = request.VerifiedName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddChatappPhoneNumber"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddChatappPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a phone number for a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddChatappPhoneNumberRequest
//
// @return AddChatappPhoneNumberResponse
func (client *Client) AddChatappPhoneNumber(request *AddChatappPhoneNumberRequest) (_result *AddChatappPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddChatappPhoneNumberResponse{}
	_body, _err := client.AddChatappPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates FAQs in the knowledge base.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - BeeBotAssociateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeeBotAssociateResponse
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

// Summary:
//
// Associates FAQs in the knowledge base.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BeeBotAssociateRequest
//
// @return BeeBotAssociateResponse
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

// Summary:
//
// Conducts sessions with the bot based on its unique identifier (ID).
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - BeeBotChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeeBotChatResponse
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

// Summary:
//
// Conducts sessions with the bot based on its unique identifier (ID).
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - BeeBotChatRequest
//
// @return BeeBotChatResponse
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

// Summary:
//
// Binds the WhatsApp Business account with ChatApp.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappBindWabaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappBindWabaResponse
func (client *Client) ChatappBindWabaWithOptions(request *ChatappBindWabaRequest, runtime *util.RuntimeOptions) (_result *ChatappBindWabaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WabaId)) {
		body["WabaId"] = request.WabaId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatappBindWaba"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatappBindWabaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds the WhatsApp Business account with ChatApp.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappBindWabaRequest
//
// @return ChatappBindWabaResponse
func (client *Client) ChatappBindWaba(request *ChatappBindWabaRequest) (_result *ChatappBindWabaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatappBindWabaResponse{}
	_body, _err := client.ChatappBindWabaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries WhatsApp Business account (WABA) information after embedded signup. You do not need to call this API operation if you use Version 2 of WhatsApp embedded signup.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappEmbedSignUpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappEmbedSignUpResponse
func (client *Client) ChatappEmbedSignUpWithOptions(request *ChatappEmbedSignUpRequest, runtime *util.RuntimeOptions) (_result *ChatappEmbedSignUpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputToken)) {
		body["InputToken"] = request.InputToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatappEmbedSignUp"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatappEmbedSignUpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries WhatsApp Business account (WABA) information after embedded signup. You do not need to call this API operation if you use Version 2 of WhatsApp embedded signup.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappEmbedSignUpRequest
//
// @return ChatappEmbedSignUpResponse
func (client *Client) ChatappEmbedSignUp(request *ChatappEmbedSignUpRequest) (_result *ChatappEmbedSignUpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatappEmbedSignUpResponse{}
	_body, _err := client.ChatappEmbedSignUpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers a phone number for migration.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappMigrationRegisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappMigrationRegisterResponse
func (client *Client) ChatappMigrationRegisterWithOptions(request *ChatappMigrationRegisterRequest, runtime *util.RuntimeOptions) (_result *ChatappMigrationRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatappMigrationRegister"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatappMigrationRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a phone number for migration.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappMigrationRegisterRequest
//
// @return ChatappMigrationRegisterResponse
func (client *Client) ChatappMigrationRegister(request *ChatappMigrationRegisterRequest) (_result *ChatappMigrationRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatappMigrationRegisterResponse{}
	_body, _err := client.ChatappMigrationRegisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies a specified phone number for migration.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappMigrationVerifiedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappMigrationVerifiedResponse
func (client *Client) ChatappMigrationVerifiedWithOptions(request *ChatappMigrationVerifiedRequest, runtime *util.RuntimeOptions) (_result *ChatappMigrationVerifiedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatappMigrationVerified"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatappMigrationVerifiedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies a specified phone number for migration.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappMigrationVerifiedRequest
//
// @return ChatappMigrationVerifiedResponse
func (client *Client) ChatappMigrationVerified(request *ChatappMigrationVerifiedRequest) (_result *ChatappMigrationVerifiedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatappMigrationVerifiedResponse{}
	_body, _err := client.ChatappMigrationVerifiedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deregisters a phone number from a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappPhoneNumberDeregisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappPhoneNumberDeregisterResponse
func (client *Client) ChatappPhoneNumberDeregisterWithOptions(request *ChatappPhoneNumberDeregisterRequest, runtime *util.RuntimeOptions) (_result *ChatappPhoneNumberDeregisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatappPhoneNumberDeregister"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatappPhoneNumberDeregisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deregisters a phone number from a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappPhoneNumberDeregisterRequest
//
// @return ChatappPhoneNumberDeregisterResponse
func (client *Client) ChatappPhoneNumberDeregister(request *ChatappPhoneNumberDeregisterRequest) (_result *ChatappPhoneNumberDeregisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatappPhoneNumberDeregisterResponse{}
	_body, _err := client.ChatappPhoneNumberDeregisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappPhoneNumberRegisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappPhoneNumberRegisterResponse
func (client *Client) ChatappPhoneNumberRegisterWithOptions(request *ChatappPhoneNumberRegisterRequest, runtime *util.RuntimeOptions) (_result *ChatappPhoneNumberRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatappPhoneNumberRegister"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatappPhoneNumberRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappPhoneNumberRegisterRequest
//
// @return ChatappPhoneNumberRegisterResponse
func (client *Client) ChatappPhoneNumberRegister(request *ChatappPhoneNumberRegisterRequest) (_result *ChatappPhoneNumberRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatappPhoneNumberRegisterResponse{}
	_body, _err := client.ChatappPhoneNumberRegisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes phone numbers.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappSyncPhoneNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappSyncPhoneNumberResponse
func (client *Client) ChatappSyncPhoneNumberWithOptions(request *ChatappSyncPhoneNumberRequest, runtime *util.RuntimeOptions) (_result *ChatappSyncPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatappSyncPhoneNumber"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatappSyncPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes phone numbers.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappSyncPhoneNumberRequest
//
// @return ChatappSyncPhoneNumberResponse
func (client *Client) ChatappSyncPhoneNumber(request *ChatappSyncPhoneNumberRequest) (_result *ChatappSyncPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatappSyncPhoneNumberResponse{}
	_body, _err := client.ChatappSyncPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a phone number with a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappVerifyAndRegisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappVerifyAndRegisterResponse
func (client *Client) ChatappVerifyAndRegisterWithOptions(request *ChatappVerifyAndRegisterRequest, runtime *util.RuntimeOptions) (_result *ChatappVerifyAndRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		body["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatappVerifyAndRegister"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatappVerifyAndRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a phone number with a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappVerifyAndRegisterRequest
//
// @return ChatappVerifyAndRegisterResponse
func (client *Client) ChatappVerifyAndRegister(request *ChatappVerifyAndRegisterRequest) (_result *ChatappVerifyAndRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatappVerifyAndRegisterResponse{}
	_body, _err := client.ChatappVerifyAndRegisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the number.
//
// Description:
//
// The status of the phone number.
//
// @param request - CreateChatappMigrationInitiateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatappMigrationInitiateResponse
func (client *Client) CreateChatappMigrationInitiateWithOptions(request *CreateChatappMigrationInitiateRequest, runtime *util.RuntimeOptions) (_result *CreateChatappMigrationInitiateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		query["CountryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		query["MobileNumber"] = request.MobileNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChatappMigrationInitiate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChatappMigrationInitiateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the number.
//
// Description:
//
// The status of the phone number.
//
// @param request - CreateChatappMigrationInitiateRequest
//
// @return CreateChatappMigrationInitiateResponse
func (client *Client) CreateChatappMigrationInitiate(request *CreateChatappMigrationInitiateRequest) (_result *CreateChatappMigrationInitiateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChatappMigrationInitiateResponse{}
	_body, _err := client.CreateChatappMigrationInitiateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The HTTP status code.
//
// \\\\	- Example: OK. This parameter indicates that the request is successful.
//
// \\\\	- Other values indicate that the request fails. For more information, see \\[Error codes]\\\\(https://www.alibabacloud.com/help/zh/cams/latest/api-error-codes).
//
// Description:
//
// The error message.
//
// @param tmpReq - CreateChatappTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatappTemplateResponse
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowCategoryChange)) {
		body["AllowCategoryChange"] = request.AllowCategoryChange
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

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

	if !tea.BoolValue(util.IsUnset(request.MessageSendTtlSeconds)) {
		body["MessageSendTtlSeconds"] = request.MessageSendTtlSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
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

// Summary:
//
// The HTTP status code.
//
// \\\\	- Example: OK. This parameter indicates that the request is successful.
//
// \\\\	- Other values indicate that the request fails. For more information, see \\[Error codes]\\\\(https://www.alibabacloud.com/help/zh/cams/latest/api-error-codes).
//
// Description:
//
// The error message.
//
// @param request - CreateChatappTemplateRequest
//
// @return CreateChatappTemplateResponse
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

// Summary:
//
// Creates a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowResponse
func (client *Client) CreateFlowWithOptions(tmpReq *CreateFlowRequest, runtime *util.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Categories)) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, tea.String("Categories"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoriesShrink)) {
		body["Categories"] = request.CategoriesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlow"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateFlowRequest
//
// @return CreateFlowResponse
func (client *Client) CreateFlow(request *CreateFlowRequest) (_result *CreateFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowResponse{}
	_body, _err := client.CreateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建消息发送二维码
//
// @param request - CreatePhoneMessageQrdlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePhoneMessageQrdlResponse
func (client *Client) CreatePhoneMessageQrdlWithOptions(request *CreatePhoneMessageQrdlRequest, runtime *util.RuntimeOptions) (_result *CreatePhoneMessageQrdlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.GenerateQrImage)) {
		body["GenerateQrImage"] = request.GenerateQrImage
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PrefilledMessage)) {
		body["PrefilledMessage"] = request.PrefilledMessage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePhoneMessageQrdl"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePhoneMessageQrdlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建消息发送二维码
//
// @param request - CreatePhoneMessageQrdlRequest
//
// @return CreatePhoneMessageQrdlResponse
func (client *Client) CreatePhoneMessageQrdl(request *CreatePhoneMessageQrdlRequest) (_result *CreatePhoneMessageQrdlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePhoneMessageQrdlResponse{}
	_body, _err := client.CreatePhoneMessageQrdlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to five times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteChatappTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChatappTemplateResponse
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

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
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

// Summary:
//
// Deletes a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to five times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteChatappTemplateRequest
//
// @return DeleteChatappTemplateResponse
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

// Summary:
//
// Deletes a Flow. Only Flows in the DRAFT state can be deleted.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlowWithOptions(request *DeleteFlowRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		body["FlowId"] = request.FlowId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlow"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Flow. Only Flows in the DRAFT state can be deleted.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteFlowRequest
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlow(request *DeleteFlowRequest) (_result *DeleteFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DeleteFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除消息发送二维码
//
// @param request - DeletePhoneMessageQrdlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePhoneMessageQrdlResponse
func (client *Client) DeletePhoneMessageQrdlWithOptions(request *DeletePhoneMessageQrdlRequest, runtime *util.RuntimeOptions) (_result *DeletePhoneMessageQrdlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.QrdlCode)) {
		body["QrdlCode"] = request.QrdlCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePhoneMessageQrdl"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePhoneMessageQrdlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消息发送二维码
//
// @param request - DeletePhoneMessageQrdlRequest
//
// @return DeletePhoneMessageQrdlResponse
func (client *Client) DeletePhoneMessageQrdl(request *DeletePhoneMessageQrdlRequest) (_result *DeletePhoneMessageQrdlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePhoneMessageQrdlResponse{}
	_body, _err := client.DeletePhoneMessageQrdlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 弃用Flow
//
// @param request - DeprecateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprecateFlowResponse
func (client *Client) DeprecateFlowWithOptions(request *DeprecateFlowRequest, runtime *util.RuntimeOptions) (_result *DeprecateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		body["FlowId"] = request.FlowId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeprecateFlow"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeprecateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 弃用Flow
//
// @param request - DeprecateFlowRequest
//
// @return DeprecateFlowResponse
func (client *Client) DeprecateFlow(request *DeprecateFlowRequest) (_result *DeprecateFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeprecateFlowResponse{}
	_body, _err := client.DeprecateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the statistics on the metrics that are related to WhatsApp.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - EnableWhatsappROIMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableWhatsappROIMetricResponse
func (client *Client) EnableWhatsappROIMetricWithOptions(request *EnableWhatsappROIMetricRequest, runtime *util.RuntimeOptions) (_result *EnableWhatsappROIMetricResponse, _err error) {
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
		Action:      tea.String("EnableWhatsappROIMetric"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableWhatsappROIMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the statistics on the metrics that are related to WhatsApp.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - EnableWhatsappROIMetricRequest
//
// @return EnableWhatsappROIMetricResponse
func (client *Client) EnableWhatsappROIMetric(request *EnableWhatsappROIMetricRequest) (_result *EnableWhatsappROIMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableWhatsappROIMetricResponse{}
	_body, _err := client.EnableWhatsappROIMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of messages that are sent by using a phone number by a specific metric.
//
// Description:
//
// You can call this operation up to 50 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappPhoneNumberMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappPhoneNumberMetricResponse
func (client *Client) GetChatappPhoneNumberMetricWithOptions(request *GetChatappPhoneNumberMetricRequest, runtime *util.RuntimeOptions) (_result *GetChatappPhoneNumberMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatappPhoneNumberMetric"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatappPhoneNumberMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of messages that are sent by using a phone number by a specific metric.
//
// Description:
//
// You can call this operation up to 50 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappPhoneNumberMetricRequest
//
// @return GetChatappPhoneNumberMetricResponse
func (client *Client) GetChatappPhoneNumberMetric(request *GetChatappPhoneNumberMetricRequest) (_result *GetChatappPhoneNumberMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChatappPhoneNumberMetricResponse{}
	_body, _err := client.GetChatappPhoneNumberMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this API operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappTemplateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappTemplateDetailResponse
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

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
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

// Summary:
//
// Queries the information of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this API operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappTemplateDetailRequest
//
// @return GetChatappTemplateDetailResponse
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

// Summary:
//
// Queries the metrics about a marketing template.
//
// Description:
//
// You can call this operation up to 50 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappTemplateMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappTemplateMetricResponse
func (client *Client) GetChatappTemplateMetricWithOptions(request *GetChatappTemplateMetricRequest, runtime *util.RuntimeOptions) (_result *GetChatappTemplateMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["IsvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatappTemplateMetric"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatappTemplateMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metrics about a marketing template.
//
// Description:
//
// You can call this operation up to 50 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappTemplateMetricRequest
//
// @return GetChatappTemplateMetricResponse
func (client *Client) GetChatappTemplateMetric(request *GetChatappTemplateMetricRequest) (_result *GetChatappTemplateMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChatappTemplateMetricResponse{}
	_body, _err := client.GetChatappTemplateMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the authentication information that is used to upload a file.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappUploadAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappUploadAuthorizationResponse
func (client *Client) GetChatappUploadAuthorizationWithOptions(request *GetChatappUploadAuthorizationRequest, runtime *util.RuntimeOptions) (_result *GetChatappUploadAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatappUploadAuthorization"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatappUploadAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the authentication information that is used to upload a file.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappUploadAuthorizationRequest
//
// @return GetChatappUploadAuthorizationResponse
func (client *Client) GetChatappUploadAuthorization(request *GetChatappUploadAuthorizationRequest) (_result *GetChatappUploadAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChatappUploadAuthorizationResponse{}
	_body, _err := client.GetChatappUploadAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a verification code.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappVerifyCodeResponse
func (client *Client) GetChatappVerifyCodeWithOptions(request *GetChatappVerifyCodeRequest, runtime *util.RuntimeOptions) (_result *GetChatappVerifyCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		body["Locale"] = request.Locale
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		body["Method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatappVerifyCode"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatappVerifyCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a verification code.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappVerifyCodeRequest
//
// @return GetChatappVerifyCodeResponse
func (client *Client) GetChatappVerifyCode(request *GetChatappVerifyCodeRequest) (_result *GetChatappVerifyCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChatappVerifyCodeResponse{}
	_body, _err := client.GetChatappVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the business settings of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCommerceSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommerceSettingResponse
func (client *Client) GetCommerceSettingWithOptions(request *GetCommerceSettingRequest, runtime *util.RuntimeOptions) (_result *GetCommerceSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCommerceSetting"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCommerceSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the business settings of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCommerceSettingRequest
//
// @return GetCommerceSettingResponse
func (client *Client) GetCommerceSetting(request *GetCommerceSettingRequest) (_result *GetCommerceSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCommerceSettingResponse{}
	_body, _err := client.GetCommerceSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowResponse
func (client *Client) GetFlowWithOptions(request *GetFlowRequest, runtime *util.RuntimeOptions) (_result *GetFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		body["FlowId"] = request.FlowId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlow"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetFlowRequest
//
// @return GetFlowResponse
func (client *Client) GetFlow(request *GetFlowRequest) (_result *GetFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFlowResponse{}
	_body, _err := client.GetFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取flow的JSON文件
//
// @param request - GetFlowJSONAssestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowJSONAssestResponse
func (client *Client) GetFlowJSONAssestWithOptions(request *GetFlowJSONAssestRequest, runtime *util.RuntimeOptions) (_result *GetFlowJSONAssestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		body["FlowId"] = request.FlowId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlowJSONAssest"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowJSONAssestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取flow的JSON文件
//
// @param request - GetFlowJSONAssestRequest
//
// @return GetFlowJSONAssestResponse
func (client *Client) GetFlowJSONAssest(request *GetFlowJSONAssestRequest) (_result *GetFlowJSONAssestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFlowJSONAssestResponse{}
	_body, _err := client.GetFlowJSONAssestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the preview URL of a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetFlowPreviewUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowPreviewUrlResponse
func (client *Client) GetFlowPreviewUrlWithOptions(request *GetFlowPreviewUrlRequest, runtime *util.RuntimeOptions) (_result *GetFlowPreviewUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		body["FlowId"] = request.FlowId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlowPreviewUrl"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowPreviewUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the preview URL of a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetFlowPreviewUrlRequest
//
// @return GetFlowPreviewUrlResponse
func (client *Client) GetFlowPreviewUrl(request *GetFlowPreviewUrlRequest) (_result *GetFlowPreviewUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFlowPreviewUrlResponse{}
	_body, _err := client.GetFlowPreviewUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The phone number.
//
// Description:
//
// The ID of the phone number.
//
// @param request - GetMigrationVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMigrationVerifyCodeResponse
func (client *Client) GetMigrationVerifyCodeWithOptions(request *GetMigrationVerifyCodeRequest, runtime *util.RuntimeOptions) (_result *GetMigrationVerifyCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		query["Method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMigrationVerifyCode"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMigrationVerifyCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The phone number.
//
// Description:
//
// The ID of the phone number.
//
// @param request - GetMigrationVerifyCodeRequest
//
// @return GetMigrationVerifyCodeResponse
func (client *Client) GetMigrationVerifyCode(request *GetMigrationVerifyCodeRequest) (_result *GetMigrationVerifyCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMigrationVerifyCodeResponse{}
	_body, _err := client.GetMigrationVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据Code获取权限
//
// @param tmpReq - GetPermissionByCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionByCodeResponse
func (client *Client) GetPermissionByCodeWithOptions(tmpReq *GetPermissionByCodeRequest, runtime *util.RuntimeOptions) (_result *GetPermissionByCodeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetPermissionByCodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Permissions)) {
		request.PermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permissions, tea.String("Permissions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionsShrink)) {
		body["Permissions"] = request.PermissionsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPermissionByCode"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPermissionByCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据Code获取权限
//
// @param request - GetPermissionByCodeRequest
//
// @return GetPermissionByCodeResponse
func (client *Client) GetPermissionByCode(request *GetPermissionByCodeRequest) (_result *GetPermissionByCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPermissionByCodeResponse{}
	_body, _err := client.GetPermissionByCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取号码的公钥
//
// @param request - GetPhoneEncryptionPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhoneEncryptionPublicKeyResponse
func (client *Client) GetPhoneEncryptionPublicKeyWithOptions(request *GetPhoneEncryptionPublicKeyRequest, runtime *util.RuntimeOptions) (_result *GetPhoneEncryptionPublicKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhoneEncryptionPublicKey"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhoneEncryptionPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取号码的公钥
//
// @param request - GetPhoneEncryptionPublicKeyRequest
//
// @return GetPhoneEncryptionPublicKeyResponse
func (client *Client) GetPhoneEncryptionPublicKey(request *GetPhoneEncryptionPublicKeyRequest) (_result *GetPhoneEncryptionPublicKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhoneEncryptionPublicKeyResponse{}
	_body, _err := client.GetPhoneEncryptionPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the verification status of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPhoneNumberVerificationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhoneNumberVerificationStatusResponse
func (client *Client) GetPhoneNumberVerificationStatusWithOptions(request *GetPhoneNumberVerificationStatusRequest, runtime *util.RuntimeOptions) (_result *GetPhoneNumberVerificationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhoneNumberVerificationStatus"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhoneNumberVerificationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the verification status of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPhoneNumberVerificationStatusRequest
//
// @return GetPhoneNumberVerificationStatusResponse
func (client *Client) GetPhoneNumberVerificationStatus(request *GetPhoneNumberVerificationStatusRequest) (_result *GetPhoneNumberVerificationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhoneNumberVerificationStatusResponse{}
	_body, _err := client.GetPhoneNumberVerificationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the ID of a pre-registered phone number used for embedded signup without the need to re-obtain a verification code.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPreValidatePhoneIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPreValidatePhoneIdResponse
func (client *Client) GetPreValidatePhoneIdWithOptions(request *GetPreValidatePhoneIdRequest, runtime *util.RuntimeOptions) (_result *GetPreValidatePhoneIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		body["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPreValidatePhoneId"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPreValidatePhoneIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the ID of a pre-registered phone number used for embedded signup without the need to re-obtain a verification code.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPreValidatePhoneIdRequest
//
// @return GetPreValidatePhoneIdResponse
func (client *Client) GetPreValidatePhoneId(request *GetPreValidatePhoneIdRequest) (_result *GetPreValidatePhoneIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPreValidatePhoneIdResponse{}
	_body, _err := client.GetPreValidatePhoneIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the product catalogs that are associated with a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetWhatsappConnectionCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWhatsappConnectionCatalogResponse
func (client *Client) GetWhatsappConnectionCatalogWithOptions(request *GetWhatsappConnectionCatalogRequest, runtime *util.RuntimeOptions) (_result *GetWhatsappConnectionCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.WabaId)) {
		query["WabaId"] = request.WabaId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWhatsappConnectionCatalog"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWhatsappConnectionCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the product catalogs that are associated with a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetWhatsappConnectionCatalogRequest
//
// @return GetWhatsappConnectionCatalogResponse
func (client *Client) GetWhatsappConnectionCatalog(request *GetWhatsappConnectionCatalogRequest) (_result *GetWhatsappConnectionCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWhatsappConnectionCatalogResponse{}
	_body, _err := client.GetWhatsappConnectionCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the application ID under the ISV account.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - IsvGetAppIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsvGetAppIdResponse
func (client *Client) IsvGetAppIdWithOptions(request *IsvGetAppIdRequest, runtime *util.RuntimeOptions) (_result *IsvGetAppIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		body["Permissions"] = request.Permissions
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IsvGetAppId"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IsvGetAppIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the application ID under the ISV account.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - IsvGetAppIdRequest
//
// @return IsvGetAppIdResponse
func (client *Client) IsvGetAppId(request *IsvGetAppIdRequest) (_result *IsvGetAppIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IsvGetAppIdResponse{}
	_body, _err := client.IsvGetAppIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries message templates.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - ListChatappTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatappTemplateResponse
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

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
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

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
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

// Summary:
//
// Queries message templates.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListChatappTemplateRequest
//
// @return ListChatappTemplateResponse
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

// Summary:
//
// Queries a list of Flows.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - ListFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowResponse
func (client *Client) ListFlowWithOptions(tmpReq *ListFlowRequest, runtime *util.RuntimeOptions) (_result *ListFlowResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Page)) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, tea.String("Page"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		body["Page"] = request.PageShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlow"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Flows.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListFlowRequest
//
// @return ListFlowResponse
func (client *Client) ListFlow(request *ListFlowRequest) (_result *ListFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowResponse{}
	_body, _err := client.ListFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询消息发送二维码
//
// @param request - ListPhoneMessageQrdlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPhoneMessageQrdlResponse
func (client *Client) ListPhoneMessageQrdlWithOptions(request *ListPhoneMessageQrdlRequest, runtime *util.RuntimeOptions) (_result *ListPhoneMessageQrdlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPhoneMessageQrdl"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPhoneMessageQrdlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消息发送二维码
//
// @param request - ListPhoneMessageQrdlRequest
//
// @return ListPhoneMessageQrdlResponse
func (client *Client) ListPhoneMessageQrdl(request *ListPhoneMessageQrdlRequest) (_result *ListPhoneMessageQrdlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPhoneMessageQrdlResponse{}
	_body, _err := client.ListPhoneMessageQrdlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries products in a product catalog.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductResponse
func (client *Client) ListProductWithOptions(request *ListProductRequest, runtime *util.RuntimeOptions) (_result *ListProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.After)) {
		query["After"] = request.After
	}

	if !tea.BoolValue(util.IsUnset(request.Before)) {
		query["Before"] = request.Before
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		query["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.WabaId)) {
		query["WabaId"] = request.WabaId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProduct"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries products in a product catalog.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListProductRequest
//
// @return ListProductResponse
func (client *Client) ListProduct(request *ListProductRequest) (_result *ListProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductResponse{}
	_body, _err := client.ListProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the product catalogs on the Business Manager platform of Meta.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListProductCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductCatalogResponse
func (client *Client) ListProductCatalogWithOptions(request *ListProductCatalogRequest, runtime *util.RuntimeOptions) (_result *ListProductCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.After)) {
		query["After"] = request.After
	}

	if !tea.BoolValue(util.IsUnset(request.Before)) {
		query["Before"] = request.Before
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessId)) {
		query["BusinessId"] = request.BusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		query["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductCatalog"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the product catalogs on the Business Manager platform of Meta.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListProductCatalogRequest
//
// @return ListProductCatalogResponse
func (client *Client) ListProductCatalog(request *ListProductCatalogRequest) (_result *ListProductCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductCatalogResponse{}
	_body, _err := client.ListProductCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The code of the message template.
//
// Description:
//
// The name of the message template.
//
// @param tmpReq - ModifyChatappTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyChatappTemplateResponse
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
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

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

	if !tea.BoolValue(util.IsUnset(request.MessageSendTtlSeconds)) {
		body["MessageSendTtlSeconds"] = request.MessageSendTtlSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
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

// Summary:
//
// The code of the message template.
//
// Description:
//
// The name of the message template.
//
// @param request - ModifyChatappTemplateRequest
//
// @return ModifyChatappTemplateResponse
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

// Summary:
//
// Modifies the basic information about a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - ModifyFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFlowResponse
func (client *Client) ModifyFlowWithOptions(tmpReq *ModifyFlowRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Categories)) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, tea.String("Categories"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoriesShrink)) {
		body["Categories"] = request.CategoriesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		body["FlowId"] = request.FlowId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlow"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyFlowRequest
//
// @return ModifyFlowResponse
func (client *Client) ModifyFlow(request *ModifyFlowRequest) (_result *ModifyFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowResponse{}
	_body, _err := client.ModifyFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// ModifyPhoneBusinessProfile
//
// @param tmpReq - ModifyPhoneBusinessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPhoneBusinessProfileResponse
func (client *Client) ModifyPhoneBusinessProfileWithOptions(tmpReq *ModifyPhoneBusinessProfileRequest, runtime *util.RuntimeOptions) (_result *ModifyPhoneBusinessProfileResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyPhoneBusinessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Websites)) {
		request.WebsitesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Websites, tea.String("Websites"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.About)) {
		query["About"] = request.About
	}

	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ProfilePictureUrl)) {
		query["ProfilePictureUrl"] = request.ProfilePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Vertical)) {
		query["Vertical"] = request.Vertical
	}

	if !tea.BoolValue(util.IsUnset(request.WebsitesShrink)) {
		query["Websites"] = request.WebsitesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPhoneBusinessProfile"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPhoneBusinessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// ModifyPhoneBusinessProfile
//
// @param request - ModifyPhoneBusinessProfileRequest
//
// @return ModifyPhoneBusinessProfileResponse
func (client *Client) ModifyPhoneBusinessProfile(request *ModifyPhoneBusinessProfileRequest) (_result *ModifyPhoneBusinessProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPhoneBusinessProfileResponse{}
	_body, _err := client.ModifyPhoneBusinessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PublishFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFlowResponse
func (client *Client) PublishFlowWithOptions(request *PublishFlowRequest, runtime *util.RuntimeOptions) (_result *PublishFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		body["FlowId"] = request.FlowId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishFlow"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PublishFlowRequest
//
// @return PublishFlowResponse
func (client *Client) PublishFlow(request *PublishFlowRequest) (_result *PublishFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishFlowResponse{}
	_body, _err := client.PublishFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the WhatsApp Business account you associate with ChatApp.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryChatappBindWabaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChatappBindWabaResponse
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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

// Summary:
//
// Query the WhatsApp Business account you associate with ChatApp.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryChatappBindWabaRequest
//
// @return QueryChatappBindWabaResponse
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

// Summary:
//
// Queries phone numbers that receive messages and statuses of these numbers under a specified user.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryChatappPhoneNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChatappPhoneNumbersResponse
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
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

// Summary:
//
// Queries phone numbers that receive messages and statuses of these numbers under a specified user.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryChatappPhoneNumbersRequest
//
// @return QueryChatappPhoneNumbersResponse
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

// Summary:
//
// Queries the business information of the account to which a specified phone number is bound.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryPhoneBusinessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneBusinessProfileResponse
func (client *Client) QueryPhoneBusinessProfileWithOptions(request *QueryPhoneBusinessProfileRequest, runtime *util.RuntimeOptions) (_result *QueryPhoneBusinessProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPhoneBusinessProfile"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPhoneBusinessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the business information of the account to which a specified phone number is bound.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryPhoneBusinessProfileRequest
//
// @return QueryPhoneBusinessProfileResponse
func (client *Client) QueryPhoneBusinessProfile(request *QueryPhoneBusinessProfileRequest) (_result *QueryPhoneBusinessProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPhoneBusinessProfileResponse{}
	_body, _err := client.QueryPhoneBusinessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the business information about the WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryWabaBusinessInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWabaBusinessInfoResponse
func (client *Client) QueryWabaBusinessInfoWithOptions(request *QueryWabaBusinessInfoRequest, runtime *util.RuntimeOptions) (_result *QueryWabaBusinessInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WabaId)) {
		query["WabaId"] = request.WabaId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWabaBusinessInfo"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWabaBusinessInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the business information about the WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryWabaBusinessInfoRequest
//
// @return QueryWabaBusinessInfoResponse
func (client *Client) QueryWabaBusinessInfo(request *QueryWabaBusinessInfoRequest) (_result *QueryWabaBusinessInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWabaBusinessInfoResponse{}
	_body, _err := client.QueryWabaBusinessInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a message to multiple phone numbers by using ChatAPP at a time.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// You can send messages to up to 1,000 phone numbers in a single request.
//
// @param tmpReq - SendChatappMassMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendChatappMassMessageResponse
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

	if !tea.BoolValue(util.IsUnset(request.FallBackDuration)) {
		body["FallBackDuration"] = request.FallBackDuration
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackId)) {
		body["FallBackId"] = request.FallBackId
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackRule)) {
		body["FallBackRule"] = request.FallBackRule
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

	if !tea.BoolValue(util.IsUnset(request.SenderListShrink)) {
		body["SenderList"] = request.SenderListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["Ttl"] = request.Ttl
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

// Summary:
//
// Sends a message to multiple phone numbers by using ChatAPP at a time.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// You can send messages to up to 1,000 phone numbers in a single request.
//
// @param request - SendChatappMassMessageRequest
//
// @return SendChatappMassMessageResponse
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

// Summary:
//
// Sends messages by using ChatAPP.
//
// Description:
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - SendChatappMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendChatappMessageResponse
func (client *Client) SendChatappMessageWithOptions(tmpReq *SendChatappMessageRequest, runtime *util.RuntimeOptions) (_result *SendChatappMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendChatappMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FlowAction)) {
		request.FlowActionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FlowAction, tea.String("FlowAction"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ProductAction)) {
		request.ProductActionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProductAction, tea.String("ProductAction"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.FallBackDuration)) {
		body["FallBackDuration"] = request.FallBackDuration
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackId)) {
		body["FallBackId"] = request.FallBackId
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackRule)) {
		body["FallBackRule"] = request.FallBackRule
	}

	if !tea.BoolValue(util.IsUnset(request.FlowActionShrink)) {
		body["FlowAction"] = request.FlowActionShrink
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

	if !tea.BoolValue(util.IsUnset(request.ProductActionShrink)) {
		body["ProductAction"] = request.ProductActionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["TemplateName"] = request.TemplateName
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

// Summary:
//
// Sends messages by using ChatAPP.
//
// Description:
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendChatappMessageRequest
//
// @return SendChatappMessageResponse
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

// Summary:
//
// Submits the agreement information for independent software vendor (ISV) customers.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SubmitIsvCustomerTermsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIsvCustomerTermsResponse
func (client *Client) SubmitIsvCustomerTermsWithOptions(request *SubmitIsvCustomerTermsRequest, runtime *util.RuntimeOptions) (_result *SubmitIsvCustomerTermsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessDesc)) {
		query["BusinessDesc"] = request.BusinessDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ContactMail)) {
		query["ContactMail"] = request.ContactMail
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.CustName)) {
		query["CustName"] = request.CustName
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvTerms)) {
		query["IsvTerms"] = request.IsvTerms
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeAddress)) {
		query["OfficeAddress"] = request.OfficeAddress
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitIsvCustomerTerms"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitIsvCustomerTermsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits the agreement information for independent software vendor (ISV) customers.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SubmitIsvCustomerTermsRequest
//
// @return SubmitIsvCustomerTermsResponse
func (client *Client) SubmitIsvCustomerTerms(request *SubmitIsvCustomerTermsRequest) (_result *SubmitIsvCustomerTermsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitIsvCustomerTermsResponse{}
	_body, _err := client.SubmitIsvCustomerTermsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the callback URL of an account.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateAccountWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountWebhookResponse
func (client *Client) UpdateAccountWebhookWithOptions(request *UpdateAccountWebhookRequest, runtime *util.RuntimeOptions) (_result *UpdateAccountWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.HttpFlag)) {
		query["HttpFlag"] = request.HttpFlag
	}

	if !tea.BoolValue(util.IsUnset(request.QueueFlag)) {
		query["QueueFlag"] = request.QueueFlag
	}

	if !tea.BoolValue(util.IsUnset(request.StatusCallbackUrl)) {
		query["StatusCallbackUrl"] = request.StatusCallbackUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccountWebhook"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccountWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the callback URL of an account.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateAccountWebhookRequest
//
// @return UpdateAccountWebhookResponse
func (client *Client) UpdateAccountWebhook(request *UpdateAccountWebhookRequest) (_result *UpdateAccountWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccountWebhookResponse{}
	_body, _err := client.UpdateAccountWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the business settings of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateCommerceSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCommerceSettingResponse
func (client *Client) UpdateCommerceSettingWithOptions(request *UpdateCommerceSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateCommerceSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CartEnable)) {
		query["CartEnable"] = request.CartEnable
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogVisible)) {
		query["CatalogVisible"] = request.CatalogVisible
	}

	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCommerceSetting"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCommerceSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the business settings of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateCommerceSettingRequest
//
// @return UpdateCommerceSettingResponse
func (client *Client) UpdateCommerceSetting(request *UpdateCommerceSettingRequest) (_result *UpdateCommerceSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCommerceSettingResponse{}
	_body, _err := client.UpdateCommerceSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新flow的表单内容
//
// @param request - UpdateFlowJSONAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowJSONAssetResponse
func (client *Client) UpdateFlowJSONAssetWithOptions(request *UpdateFlowJSONAssetRequest, runtime *util.RuntimeOptions) (_result *UpdateFlowJSONAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		body["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		body["FlowId"] = request.FlowId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFlowJSONAsset"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFlowJSONAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新flow的表单内容
//
// @param request - UpdateFlowJSONAssetRequest
//
// @return UpdateFlowJSONAssetResponse
func (client *Client) UpdateFlowJSONAsset(request *UpdateFlowJSONAssetRequest) (_result *UpdateFlowJSONAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlowJSONAssetResponse{}
	_body, _err := client.UpdateFlowJSONAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新号码的公钥
//
// @param request - UpdatePhoneEncryptionPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePhoneEncryptionPublicKeyResponse
func (client *Client) UpdatePhoneEncryptionPublicKeyWithOptions(request *UpdatePhoneEncryptionPublicKeyRequest, runtime *util.RuntimeOptions) (_result *UpdatePhoneEncryptionPublicKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionPublicKey)) {
		body["EncryptionPublicKey"] = request.EncryptionPublicKey
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePhoneEncryptionPublicKey"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePhoneEncryptionPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新号码的公钥
//
// @param request - UpdatePhoneEncryptionPublicKeyRequest
//
// @return UpdatePhoneEncryptionPublicKeyResponse
func (client *Client) UpdatePhoneEncryptionPublicKey(request *UpdatePhoneEncryptionPublicKeyRequest) (_result *UpdatePhoneEncryptionPublicKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePhoneEncryptionPublicKeyResponse{}
	_body, _err := client.UpdatePhoneEncryptionPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改消息发送二维码
//
// @param request - UpdatePhoneMessageQrdlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePhoneMessageQrdlResponse
func (client *Client) UpdatePhoneMessageQrdlWithOptions(request *UpdatePhoneMessageQrdlRequest, runtime *util.RuntimeOptions) (_result *UpdatePhoneMessageQrdlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.GenerateQrImage)) {
		body["GenerateQrImage"] = request.GenerateQrImage
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PrefilledMessage)) {
		body["PrefilledMessage"] = request.PrefilledMessage
	}

	if !tea.BoolValue(util.IsUnset(request.QrdlCode)) {
		body["QrdlCode"] = request.QrdlCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePhoneMessageQrdl"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePhoneMessageQrdlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改消息发送二维码
//
// @param request - UpdatePhoneMessageQrdlRequest
//
// @return UpdatePhoneMessageQrdlResponse
func (client *Client) UpdatePhoneMessageQrdl(request *UpdatePhoneMessageQrdlRequest) (_result *UpdatePhoneMessageQrdlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePhoneMessageQrdlResponse{}
	_body, _err := client.UpdatePhoneMessageQrdlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The HTTP status code returned.
//
// 	- A value of OK indicates that the call is successful.
//
// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
//
// Description:
//
// The error message returned.
//
// @param request - UpdatePhoneWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePhoneWebhookResponse
func (client *Client) UpdatePhoneWebhookWithOptions(request *UpdatePhoneWebhookRequest, runtime *util.RuntimeOptions) (_result *UpdatePhoneWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustSpaceId)) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.HttpFlag)) {
		query["HttpFlag"] = request.HttpFlag
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.QueueFlag)) {
		query["QueueFlag"] = request.QueueFlag
	}

	if !tea.BoolValue(util.IsUnset(request.StatusCallbackUrl)) {
		query["StatusCallbackUrl"] = request.StatusCallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.UpCallbackUrl)) {
		query["UpCallbackUrl"] = request.UpCallbackUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePhoneWebhook"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePhoneWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The HTTP status code returned.
//
// 	- A value of OK indicates that the call is successful.
//
// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
//
// Description:
//
// The error message returned.
//
// @param request - UpdatePhoneWebhookRequest
//
// @return UpdatePhoneWebhookResponse
func (client *Client) UpdatePhoneWebhook(request *UpdatePhoneWebhookRequest) (_result *UpdatePhoneWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePhoneWebhookResponse{}
	_body, _err := client.UpdatePhoneWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
