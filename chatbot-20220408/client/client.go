// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApplyForStreamAccessTokenRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ApplyForStreamAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyForStreamAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *ApplyForStreamAccessTokenRequest) SetAgentKey(v string) *ApplyForStreamAccessTokenRequest {
	s.AgentKey = &v
	return s
}

type ApplyForStreamAccessTokenResponseBody struct {
	// example:
	//
	// 63ba97b4f18a4a04f715c81e8e643938
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// cc9e88c0-4f41-4f1d-a1a9-91a72d2aa27d
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 736994BD-AA35-4742-88C9-E64BE4BAA14B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// cc9e88c0-4f41-4f1d-a1a9-91a72d2aa27d
	StreamSecret *string `json:"StreamSecret,omitempty" xml:"StreamSecret,omitempty"`
}

func (s ApplyForStreamAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyForStreamAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyForStreamAccessTokenResponseBody) SetAccessToken(v string) *ApplyForStreamAccessTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponseBody) SetChannelId(v string) *ApplyForStreamAccessTokenResponseBody {
	s.ChannelId = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponseBody) SetRequestId(v string) *ApplyForStreamAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponseBody) SetStreamSecret(v string) *ApplyForStreamAccessTokenResponseBody {
	s.StreamSecret = &v
	return s
}

type ApplyForStreamAccessTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyForStreamAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyForStreamAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyForStreamAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *ApplyForStreamAccessTokenResponse) SetHeaders(v map[string]*string) *ApplyForStreamAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *ApplyForStreamAccessTokenResponse) SetStatusCode(v int32) *ApplyForStreamAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponse) SetBody(v *ApplyForStreamAccessTokenResponseBody) *ApplyForStreamAccessTokenResponse {
	s.Body = v
	return s
}

type AssociateRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["qyzzVfyFfa"]
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	RecommendNum *int64 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// example:
	//
	// 7c3cec23cc8940bc9db4a318c8f4f0aa
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 公积金提取
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s AssociateRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateRequest) GoString() string {
	return s.String()
}

func (s *AssociateRequest) SetAgentKey(v string) *AssociateRequest {
	s.AgentKey = &v
	return s
}

func (s *AssociateRequest) SetInstanceId(v string) *AssociateRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateRequest) SetPerspective(v []*string) *AssociateRequest {
	s.Perspective = v
	return s
}

func (s *AssociateRequest) SetRecommendNum(v int64) *AssociateRequest {
	s.RecommendNum = &v
	return s
}

func (s *AssociateRequest) SetSessionId(v string) *AssociateRequest {
	s.SessionId = &v
	return s
}

func (s *AssociateRequest) SetUtterance(v string) *AssociateRequest {
	s.Utterance = &v
	return s
}

type AssociateShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["qyzzVfyFfa"]
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
	// example:
	//
	// 8
	RecommendNum *int64 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// example:
	//
	// 7c3cec23cc8940bc9db4a318c8f4f0aa
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 公积金提取
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s AssociateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssociateShrinkRequest) SetAgentKey(v string) *AssociateShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *AssociateShrinkRequest) SetInstanceId(v string) *AssociateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateShrinkRequest) SetPerspectiveShrink(v string) *AssociateShrinkRequest {
	s.PerspectiveShrink = &v
	return s
}

func (s *AssociateShrinkRequest) SetRecommendNum(v int64) *AssociateShrinkRequest {
	s.RecommendNum = &v
	return s
}

func (s *AssociateShrinkRequest) SetSessionId(v string) *AssociateShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *AssociateShrinkRequest) SetUtterance(v string) *AssociateShrinkRequest {
	s.Utterance = &v
	return s
}

type AssociateResponseBody struct {
	Associate []*AssociateResponseBodyAssociate `json:"Associate,omitempty" xml:"Associate,omitempty" type:"Repeated"`
	// example:
	//
	// 5ADF0EBD-7C50-1922-A28B-43215B47CC1A
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 5C20F0D4-9721-178A-8236-3BF990634962
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1531ded6b3df4afca4be63943f708bb7
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AssociateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResponseBody) SetAssociate(v []*AssociateResponseBodyAssociate) *AssociateResponseBody {
	s.Associate = v
	return s
}

func (s *AssociateResponseBody) SetMessageId(v string) *AssociateResponseBody {
	s.MessageId = &v
	return s
}

func (s *AssociateResponseBody) SetRequestId(v string) *AssociateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateResponseBody) SetSessionId(v string) *AssociateResponseBody {
	s.SessionId = &v
	return s
}

type AssociateResponseBodyAssociate struct {
	// example:
	//
	// {}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 公积金提取的政策
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AssociateResponseBodyAssociate) String() string {
	return tea.Prettify(s)
}

func (s AssociateResponseBodyAssociate) GoString() string {
	return s.String()
}

func (s *AssociateResponseBodyAssociate) SetMeta(v string) *AssociateResponseBodyAssociate {
	s.Meta = &v
	return s
}

func (s *AssociateResponseBodyAssociate) SetTitle(v string) *AssociateResponseBodyAssociate {
	s.Title = &v
	return s
}

type AssociateResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AssociateResponse) SetStatusCode(v int32) *AssociateResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateResponse) SetBody(v *AssociateResponseBody) *AssociateResponse {
	s.Body = v
	return s
}

type BeginSessionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s BeginSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s BeginSessionRequest) GoString() string {
	return s.String()
}

func (s *BeginSessionRequest) SetAgentKey(v string) *BeginSessionRequest {
	s.AgentKey = &v
	return s
}

func (s *BeginSessionRequest) SetInstanceId(v string) *BeginSessionRequest {
	s.InstanceId = &v
	return s
}

type BeginSessionResponseBody struct {
	AsrMaxEndSilence *int32 `json:"AsrMaxEndSilence,omitempty" xml:"AsrMaxEndSilence,omitempty"`
	Interruptible    *bool  `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 149C7528-C104-1B50-A4F9-0C5907A8AD9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 静默超时时间
	//
	// example:
	//
	// 5
	SilenceReplyTimeout *int32 `json:"SilenceReplyTimeout,omitempty" xml:"SilenceReplyTimeout,omitempty"`
	// example:
	//
	// 智能对话机器人为您服务，请问有什么可以帮您？
	WelcomeMessage *string `json:"WelcomeMessage,omitempty" xml:"WelcomeMessage,omitempty"`
}

func (s BeginSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeginSessionResponseBody) GoString() string {
	return s.String()
}

func (s *BeginSessionResponseBody) SetAsrMaxEndSilence(v int32) *BeginSessionResponseBody {
	s.AsrMaxEndSilence = &v
	return s
}

func (s *BeginSessionResponseBody) SetInterruptible(v bool) *BeginSessionResponseBody {
	s.Interruptible = &v
	return s
}

func (s *BeginSessionResponseBody) SetRequestId(v string) *BeginSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginSessionResponseBody) SetSilenceReplyTimeout(v int32) *BeginSessionResponseBody {
	s.SilenceReplyTimeout = &v
	return s
}

func (s *BeginSessionResponseBody) SetWelcomeMessage(v string) *BeginSessionResponseBody {
	s.WelcomeMessage = &v
	return s
}

type BeginSessionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BeginSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BeginSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s BeginSessionResponse) GoString() string {
	return s.String()
}

func (s *BeginSessionResponse) SetHeaders(v map[string]*string) *BeginSessionResponse {
	s.Headers = v
	return s
}

func (s *BeginSessionResponse) SetStatusCode(v int32) *BeginSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *BeginSessionResponse) SetBody(v *BeginSessionResponseBody) *BeginSessionResponse {
	s.Body = v
	return s
}

type CancelInstancePublishTaskRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CancelInstancePublishTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelInstancePublishTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelInstancePublishTaskRequest) SetAgentKey(v string) *CancelInstancePublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelInstancePublishTaskRequest) SetId(v int64) *CancelInstancePublishTaskRequest {
	s.Id = &v
	return s
}

func (s *CancelInstancePublishTaskRequest) SetInstanceId(v string) *CancelInstancePublishTaskRequest {
	s.InstanceId = &v
	return s
}

type CancelInstancePublishTaskResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_ABORTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelInstancePublishTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelInstancePublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelInstancePublishTaskResponseBody) SetBizTypeList(v []*string) *CancelInstancePublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetCreateTime(v string) *CancelInstancePublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetError(v string) *CancelInstancePublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetId(v int64) *CancelInstancePublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetModifyTime(v string) *CancelInstancePublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetRequestId(v string) *CancelInstancePublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetResponse(v string) *CancelInstancePublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *CancelInstancePublishTaskResponseBody) SetStatus(v string) *CancelInstancePublishTaskResponseBody {
	s.Status = &v
	return s
}

type CancelInstancePublishTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelInstancePublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelInstancePublishTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelInstancePublishTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelInstancePublishTaskResponse) SetHeaders(v map[string]*string) *CancelInstancePublishTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelInstancePublishTaskResponse) SetStatusCode(v int32) *CancelInstancePublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelInstancePublishTaskResponse) SetBody(v *CancelInstancePublishTaskResponseBody) *CancelInstancePublishTaskResponse {
	s.Body = v
	return s
}

type CancelPublishTaskRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CancelPublishTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelPublishTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelPublishTaskRequest) SetAgentKey(v string) *CancelPublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelPublishTaskRequest) SetId(v int64) *CancelPublishTaskRequest {
	s.Id = &v
	return s
}

type CancelPublishTaskResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_ABORTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelPublishTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelPublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPublishTaskResponseBody) SetBizTypeList(v []*string) *CancelPublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *CancelPublishTaskResponseBody) SetCreateTime(v string) *CancelPublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetError(v string) *CancelPublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetId(v int64) *CancelPublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetModifyTime(v string) *CancelPublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetRequestId(v string) *CancelPublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetResponse(v string) *CancelPublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *CancelPublishTaskResponseBody) SetStatus(v string) *CancelPublishTaskResponseBody {
	s.Status = &v
	return s
}

type CancelPublishTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPublishTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelPublishTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelPublishTaskResponse) SetHeaders(v map[string]*string) *CancelPublishTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelPublishTaskResponse) SetStatusCode(v int32) *CancelPublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPublishTaskResponse) SetBody(v *CancelPublishTaskResponseBody) *CancelPublishTaskResponse {
	s.Body = v
	return s
}

type ChatRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// 30002406051
	KnowledgeId *string   `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SandBox *bool `json:"SandBox,omitempty" xml:"SandBox,omitempty"`
	// example:
	//
	// custumer_123456
	SenderId   *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// example:
	//
	// 9c6ebdc6e66f46ecadab3434314f6959
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	// example:
	//
	// {"phone":123456789}
	VendorParam *string `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
}

func (s ChatRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) SetAgentKey(v string) *ChatRequest {
	s.AgentKey = &v
	return s
}

func (s *ChatRequest) SetInstanceId(v string) *ChatRequest {
	s.InstanceId = &v
	return s
}

func (s *ChatRequest) SetIntentName(v string) *ChatRequest {
	s.IntentName = &v
	return s
}

func (s *ChatRequest) SetKnowledgeId(v string) *ChatRequest {
	s.KnowledgeId = &v
	return s
}

func (s *ChatRequest) SetPerspective(v []*string) *ChatRequest {
	s.Perspective = v
	return s
}

func (s *ChatRequest) SetSandBox(v bool) *ChatRequest {
	s.SandBox = &v
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

func (s *ChatRequest) SetSessionId(v string) *ChatRequest {
	s.SessionId = &v
	return s
}

func (s *ChatRequest) SetUtterance(v string) *ChatRequest {
	s.Utterance = &v
	return s
}

func (s *ChatRequest) SetVendorParam(v string) *ChatRequest {
	s.VendorParam = &v
	return s
}

type ChatShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// 30002406051
	KnowledgeId       *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
	// example:
	//
	// true
	SandBox *bool `json:"SandBox,omitempty" xml:"SandBox,omitempty"`
	// example:
	//
	// custumer_123456
	SenderId   *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// example:
	//
	// 9c6ebdc6e66f46ecadab3434314f6959
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	// example:
	//
	// {"phone":123456789}
	VendorParam *string `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
}

func (s ChatShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatShrinkRequest) SetAgentKey(v string) *ChatShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ChatShrinkRequest) SetInstanceId(v string) *ChatShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ChatShrinkRequest) SetIntentName(v string) *ChatShrinkRequest {
	s.IntentName = &v
	return s
}

func (s *ChatShrinkRequest) SetKnowledgeId(v string) *ChatShrinkRequest {
	s.KnowledgeId = &v
	return s
}

func (s *ChatShrinkRequest) SetPerspectiveShrink(v string) *ChatShrinkRequest {
	s.PerspectiveShrink = &v
	return s
}

func (s *ChatShrinkRequest) SetSandBox(v bool) *ChatShrinkRequest {
	s.SandBox = &v
	return s
}

func (s *ChatShrinkRequest) SetSenderId(v string) *ChatShrinkRequest {
	s.SenderId = &v
	return s
}

func (s *ChatShrinkRequest) SetSenderNick(v string) *ChatShrinkRequest {
	s.SenderNick = &v
	return s
}

func (s *ChatShrinkRequest) SetSessionId(v string) *ChatShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *ChatShrinkRequest) SetUtterance(v string) *ChatShrinkRequest {
	s.Utterance = &v
	return s
}

func (s *ChatShrinkRequest) SetVendorParam(v string) *ChatShrinkRequest {
	s.VendorParam = &v
	return s
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
	return tea.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessages) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesKnowledge) GoString() string {
	return s.String()
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

type ChatResponseBodyMessagesKnowledgeRelatedKnowledges struct {
	// example:
	//
	// 735899
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
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesRecommends) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesText) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s ChatResponseBodyMessagesTextSlots) GoString() string {
	return s.String()
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

type ChatResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ChatResponse) SetStatusCode(v int32) *ChatResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatResponse) SetBody(v *ChatResponseBody) *ChatResponse {
	s.Body = v
	return s
}

type ContinueInstancePublishTaskRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ContinueInstancePublishTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueInstancePublishTaskRequest) GoString() string {
	return s.String()
}

func (s *ContinueInstancePublishTaskRequest) SetAgentKey(v string) *ContinueInstancePublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *ContinueInstancePublishTaskRequest) SetId(v int64) *ContinueInstancePublishTaskRequest {
	s.Id = &v
	return s
}

func (s *ContinueInstancePublishTaskRequest) SetInstanceId(v string) *ContinueInstancePublishTaskRequest {
	s.InstanceId = &v
	return s
}

type ContinueInstancePublishTaskResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// {}
	Errors map[string]interface{} `json:"Errors,omitempty" xml:"Errors,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {         "category_bind_faq": [             "以下类目没有发布到正式环境: 项目交付信息汇总"         ]     }
	Warnings map[string]interface{} `json:"Warnings,omitempty" xml:"Warnings,omitempty"`
}

func (s ContinueInstancePublishTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueInstancePublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueInstancePublishTaskResponseBody) SetBizTypeList(v []*string) *ContinueInstancePublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetCreateTime(v string) *ContinueInstancePublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetError(v string) *ContinueInstancePublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetErrors(v map[string]interface{}) *ContinueInstancePublishTaskResponseBody {
	s.Errors = v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetId(v int64) *ContinueInstancePublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetModifyTime(v string) *ContinueInstancePublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetRequestId(v string) *ContinueInstancePublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetResponse(v string) *ContinueInstancePublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetStatus(v string) *ContinueInstancePublishTaskResponseBody {
	s.Status = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetWarnings(v map[string]interface{}) *ContinueInstancePublishTaskResponseBody {
	s.Warnings = v
	return s
}

type ContinueInstancePublishTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueInstancePublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueInstancePublishTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueInstancePublishTaskResponse) GoString() string {
	return s.String()
}

func (s *ContinueInstancePublishTaskResponse) SetHeaders(v map[string]*string) *ContinueInstancePublishTaskResponse {
	s.Headers = v
	return s
}

func (s *ContinueInstancePublishTaskResponse) SetStatusCode(v int32) *ContinueInstancePublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueInstancePublishTaskResponse) SetBody(v *ContinueInstancePublishTaskResponseBody) *ContinueInstancePublishTaskResponse {
	s.Body = v
	return s
}

type CreateCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	BizCode       *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	KnowledgeType *int32  `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s CreateCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateCategoryRequest) SetAgentKey(v string) *CreateCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateCategoryRequest) SetBizCode(v string) *CreateCategoryRequest {
	s.BizCode = &v
	return s
}

func (s *CreateCategoryRequest) SetKnowledgeType(v int32) *CreateCategoryRequest {
	s.KnowledgeType = &v
	return s
}

func (s *CreateCategoryRequest) SetName(v string) *CreateCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateCategoryRequest) SetParentCategoryId(v int64) *CreateCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

type CreateCategoryResponseBody struct {
	Category *CreateCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// example:
	//
	// A629A28F-F25E-5572-A679-FA46FB0151D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponseBody) SetCategory(v *CreateCategoryResponseBodyCategory) *CreateCategoryResponseBody {
	s.Category = v
	return s
}

func (s *CreateCategoryResponseBody) SetRequestId(v string) *CreateCategoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateCategoryResponseBodyCategory struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 30000049006
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateCategoryResponseBodyCategory) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponseBodyCategory) SetBizCode(v string) *CreateCategoryResponseBodyCategory {
	s.BizCode = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) SetCategoryId(v int64) *CreateCategoryResponseBodyCategory {
	s.CategoryId = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) SetName(v string) *CreateCategoryResponseBodyCategory {
	s.Name = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) SetParentCategoryId(v int64) *CreateCategoryResponseBodyCategory {
	s.ParentCategoryId = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) SetStatus(v int32) *CreateCategoryResponseBodyCategory {
	s.Status = &v
	return s
}

type CreateCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateCategoryResponse) SetStatusCode(v int32) *CreateCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCategoryResponse) SetBody(v *CreateCategoryResponseBody) *CreateCategoryResponse {
	s.Body = v
	return s
}

type CreateConnQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30002654628
	ConnQuestionId *int64 `json:"ConnQuestionId,omitempty" xml:"ConnQuestionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30002174773
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s CreateConnQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConnQuestionRequest) GoString() string {
	return s.String()
}

func (s *CreateConnQuestionRequest) SetAgentKey(v string) *CreateConnQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateConnQuestionRequest) SetConnQuestionId(v int64) *CreateConnQuestionRequest {
	s.ConnQuestionId = &v
	return s
}

func (s *CreateConnQuestionRequest) SetKnowledgeId(v int64) *CreateConnQuestionRequest {
	s.KnowledgeId = &v
	return s
}

type CreateConnQuestionResponseBody struct {
	// example:
	//
	// 1000002123
	OutlineId *int64 `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	// example:
	//
	// C191B48B-9268-4FB1-A3C2-5143B4A91D0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConnQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConnQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnQuestionResponseBody) SetOutlineId(v int64) *CreateConnQuestionResponseBody {
	s.OutlineId = &v
	return s
}

func (s *CreateConnQuestionResponseBody) SetRequestId(v string) *CreateConnQuestionResponseBody {
	s.RequestId = &v
	return s
}

type CreateConnQuestionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConnQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConnQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConnQuestionResponse) GoString() string {
	return s.String()
}

func (s *CreateConnQuestionResponse) SetHeaders(v map[string]*string) *CreateConnQuestionResponse {
	s.Headers = v
	return s
}

func (s *CreateConnQuestionResponse) SetStatusCode(v int32) *CreateConnQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConnQuestionResponse) SetBody(v *CreateConnQuestionResponseBody) *CreateConnQuestionResponse {
	s.Body = v
	return s
}

type CreateDSEntityRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 实体名称
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// synonyms
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDSEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDSEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateDSEntityRequest) SetAgentKey(v string) *CreateDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDSEntityRequest) SetEntityName(v string) *CreateDSEntityRequest {
	s.EntityName = &v
	return s
}

func (s *CreateDSEntityRequest) SetEntityType(v string) *CreateDSEntityRequest {
	s.EntityType = &v
	return s
}

func (s *CreateDSEntityRequest) SetInstanceId(v string) *CreateDSEntityRequest {
	s.InstanceId = &v
	return s
}

type CreateDSEntityResponseBody struct {
	// example:
	//
	// 23436345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// adfad2343f1f2r
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDSEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDSEntityResponseBody) SetEntityId(v int64) *CreateDSEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *CreateDSEntityResponseBody) SetRequestId(v string) *CreateDSEntityResponseBody {
	s.RequestId = &v
	return s
}

type CreateDSEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDSEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDSEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateDSEntityResponse) SetHeaders(v map[string]*string) *CreateDSEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateDSEntityResponse) SetStatusCode(v int32) *CreateDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDSEntityResponse) SetBody(v *CreateDSEntityResponseBody) *CreateDSEntityResponse {
	s.Body = v
	return s
}

type CreateDSEntityValueRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ada
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s CreateDSEntityValueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDSEntityValueRequest) GoString() string {
	return s.String()
}

func (s *CreateDSEntityValueRequest) SetAgentKey(v string) *CreateDSEntityValueRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDSEntityValueRequest) SetContent(v string) *CreateDSEntityValueRequest {
	s.Content = &v
	return s
}

func (s *CreateDSEntityValueRequest) SetEntityId(v int64) *CreateDSEntityValueRequest {
	s.EntityId = &v
	return s
}

func (s *CreateDSEntityValueRequest) SetInstanceId(v string) *CreateDSEntityValueRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDSEntityValueRequest) SetSynonyms(v []*string) *CreateDSEntityValueRequest {
	s.Synonyms = v
	return s
}

type CreateDSEntityValueShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ada
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SynonymsShrink *string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty"`
}

func (s CreateDSEntityValueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDSEntityValueShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDSEntityValueShrinkRequest) SetAgentKey(v string) *CreateDSEntityValueShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) SetContent(v string) *CreateDSEntityValueShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) SetEntityId(v int64) *CreateDSEntityValueShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) SetInstanceId(v string) *CreateDSEntityValueShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) SetSynonymsShrink(v string) *CreateDSEntityValueShrinkRequest {
	s.SynonymsShrink = &v
	return s
}

type CreateDSEntityValueResponseBody struct {
	// example:
	//
	// 2434543453
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// example:
	//
	// g763hg48j3f3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDSEntityValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDSEntityValueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDSEntityValueResponseBody) SetEntityValueId(v int64) *CreateDSEntityValueResponseBody {
	s.EntityValueId = &v
	return s
}

func (s *CreateDSEntityValueResponseBody) SetRequestId(v string) *CreateDSEntityValueResponseBody {
	s.RequestId = &v
	return s
}

type CreateDSEntityValueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDSEntityValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDSEntityValueResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDSEntityValueResponse) GoString() string {
	return s.String()
}

func (s *CreateDSEntityValueResponse) SetHeaders(v map[string]*string) *CreateDSEntityValueResponse {
	s.Headers = v
	return s
}

func (s *CreateDSEntityValueResponse) SetStatusCode(v int32) *CreateDSEntityValueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDSEntityValueResponse) SetBody(v *CreateDSEntityValueResponseBody) *CreateDSEntityValueResponse {
	s.Body = v
	return s
}

type CreateDocRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":""}
	Config      *string                        `json:"Config,omitempty" xml:"Config,omitempty"`
	Content     *string                        `json:"Content,omitempty" xml:"Content,omitempty"`
	DocMetadata []*CreateDocRequestDocMetadata `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty" type:"Repeated"`
	// example:
	//
	// 2032-05-25T16:28:36Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate *string  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIds    []*int64 `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://example.com/example.pdf
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDocRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocRequest) GoString() string {
	return s.String()
}

func (s *CreateDocRequest) SetAgentKey(v string) *CreateDocRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDocRequest) SetCategoryId(v int64) *CreateDocRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateDocRequest) SetConfig(v string) *CreateDocRequest {
	s.Config = &v
	return s
}

func (s *CreateDocRequest) SetContent(v string) *CreateDocRequest {
	s.Content = &v
	return s
}

func (s *CreateDocRequest) SetDocMetadata(v []*CreateDocRequestDocMetadata) *CreateDocRequest {
	s.DocMetadata = v
	return s
}

func (s *CreateDocRequest) SetEndDate(v string) *CreateDocRequest {
	s.EndDate = &v
	return s
}

func (s *CreateDocRequest) SetMeta(v string) *CreateDocRequest {
	s.Meta = &v
	return s
}

func (s *CreateDocRequest) SetStartDate(v string) *CreateDocRequest {
	s.StartDate = &v
	return s
}

func (s *CreateDocRequest) SetTagIds(v []*int64) *CreateDocRequest {
	s.TagIds = v
	return s
}

func (s *CreateDocRequest) SetTitle(v string) *CreateDocRequest {
	s.Title = &v
	return s
}

func (s *CreateDocRequest) SetUrl(v string) *CreateDocRequest {
	s.Url = &v
	return s
}

type CreateDocRequestDocMetadata struct {
	BusinessViewId      *string                                           `json:"BusinessViewId,omitempty" xml:"BusinessViewId,omitempty"`
	BusinessViewName    *string                                           `json:"BusinessViewName,omitempty" xml:"BusinessViewName,omitempty"`
	MetaCellInfoDTOList []*CreateDocRequestDocMetadataMetaCellInfoDTOList `json:"MetaCellInfoDTOList,omitempty" xml:"MetaCellInfoDTOList,omitempty" type:"Repeated"`
}

func (s CreateDocRequestDocMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateDocRequestDocMetadata) GoString() string {
	return s.String()
}

func (s *CreateDocRequestDocMetadata) SetBusinessViewId(v string) *CreateDocRequestDocMetadata {
	s.BusinessViewId = &v
	return s
}

func (s *CreateDocRequestDocMetadata) SetBusinessViewName(v string) *CreateDocRequestDocMetadata {
	s.BusinessViewName = &v
	return s
}

func (s *CreateDocRequestDocMetadata) SetMetaCellInfoDTOList(v []*CreateDocRequestDocMetadataMetaCellInfoDTOList) *CreateDocRequestDocMetadata {
	s.MetaCellInfoDTOList = v
	return s
}

type CreateDocRequestDocMetadataMetaCellInfoDTOList struct {
	FieldCode *string `json:"FieldCode,omitempty" xml:"FieldCode,omitempty"`
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDocRequestDocMetadataMetaCellInfoDTOList) String() string {
	return tea.Prettify(s)
}

func (s CreateDocRequestDocMetadataMetaCellInfoDTOList) GoString() string {
	return s.String()
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) SetFieldCode(v string) *CreateDocRequestDocMetadataMetaCellInfoDTOList {
	s.FieldCode = &v
	return s
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) SetFieldName(v string) *CreateDocRequestDocMetadataMetaCellInfoDTOList {
	s.FieldName = &v
	return s
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) SetValue(v string) *CreateDocRequestDocMetadataMetaCellInfoDTOList {
	s.Value = &v
	return s
}

type CreateDocShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":""}
	Config            *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DocMetadataShrink *string `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty"`
	// example:
	//
	// 2032-05-25T16:28:36Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://example.com/example.pdf
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDocShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDocShrinkRequest) SetAgentKey(v string) *CreateDocShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDocShrinkRequest) SetCategoryId(v int64) *CreateDocShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateDocShrinkRequest) SetConfig(v string) *CreateDocShrinkRequest {
	s.Config = &v
	return s
}

func (s *CreateDocShrinkRequest) SetContent(v string) *CreateDocShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateDocShrinkRequest) SetDocMetadataShrink(v string) *CreateDocShrinkRequest {
	s.DocMetadataShrink = &v
	return s
}

func (s *CreateDocShrinkRequest) SetEndDate(v string) *CreateDocShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *CreateDocShrinkRequest) SetMeta(v string) *CreateDocShrinkRequest {
	s.Meta = &v
	return s
}

func (s *CreateDocShrinkRequest) SetStartDate(v string) *CreateDocShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *CreateDocShrinkRequest) SetTagIdsShrink(v string) *CreateDocShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

func (s *CreateDocShrinkRequest) SetTitle(v string) *CreateDocShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateDocShrinkRequest) SetUrl(v string) *CreateDocShrinkRequest {
	s.Url = &v
	return s
}

type CreateDocResponseBody struct {
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 07B270A4-61D8-57F6-A609-A3C216CFB872
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocResponseBody) SetKnowledgeId(v int64) *CreateDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *CreateDocResponseBody) SetRequestId(v string) *CreateDocResponseBody {
	s.RequestId = &v
	return s
}

type CreateDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDocResponse) GoString() string {
	return s.String()
}

func (s *CreateDocResponse) SetHeaders(v map[string]*string) *CreateDocResponse {
	s.Headers = v
	return s
}

func (s *CreateDocResponse) SetStatusCode(v int32) *CreateDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocResponse) SetBody(v *CreateDocResponseBody) *CreateDocResponse {
	s.Body = v
	return s
}

type CreateFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000053274
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2030-12-31T16:00:00Z
	EndDate         *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	SolutionContent *string `json:"SolutionContent,omitempty" xml:"SolutionContent,omitempty"`
	// example:
	//
	// 0
	SolutionType *int32 `json:"SolutionType,omitempty" xml:"SolutionType,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateFaqRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFaqRequest) GoString() string {
	return s.String()
}

func (s *CreateFaqRequest) SetAgentKey(v string) *CreateFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateFaqRequest) SetCategoryId(v int64) *CreateFaqRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateFaqRequest) SetEndDate(v string) *CreateFaqRequest {
	s.EndDate = &v
	return s
}

func (s *CreateFaqRequest) SetSolutionContent(v string) *CreateFaqRequest {
	s.SolutionContent = &v
	return s
}

func (s *CreateFaqRequest) SetSolutionType(v int32) *CreateFaqRequest {
	s.SolutionType = &v
	return s
}

func (s *CreateFaqRequest) SetStartDate(v string) *CreateFaqRequest {
	s.StartDate = &v
	return s
}

func (s *CreateFaqRequest) SetTitle(v string) *CreateFaqRequest {
	s.Title = &v
	return s
}

type CreateFaqResponseBody struct {
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 28805A7C-D695-548C-A31B-67E52C2C274F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFaqResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFaqResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFaqResponseBody) SetKnowledgeId(v int64) *CreateFaqResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *CreateFaqResponseBody) SetRequestId(v string) *CreateFaqResponseBody {
	s.RequestId = &v
	return s
}

type CreateFaqResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFaqResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFaqResponse) GoString() string {
	return s.String()
}

func (s *CreateFaqResponse) SetHeaders(v map[string]*string) *CreateFaqResponse {
	s.Headers = v
	return s
}

func (s *CreateFaqResponse) SetStatusCode(v int32) *CreateFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFaqResponse) SetBody(v *CreateFaqResponseBody) *CreateFaqResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 用于手机app的小蜜机器人
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// example:
	//
	// 小蜜机器人
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// scenario_im
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAgentKey(v string) *CreateInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateInstanceRequest) SetIntroduction(v string) *CreateInstanceRequest {
	s.Introduction = &v
	return s
}

func (s *CreateInstanceRequest) SetLanguageCode(v string) *CreateInstanceRequest {
	s.LanguageCode = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetRobotType(v string) *CreateInstanceRequest {
	s.RobotType = &v
	return s
}

type CreateInstanceResponseBody struct {
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2E2C40D-AB09-45A1-B5C5-EB9F5C4E4E4A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateInstancePublishTaskRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateInstancePublishTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancePublishTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInstancePublishTaskRequest) SetAgentKey(v string) *CreateInstancePublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateInstancePublishTaskRequest) SetInstanceId(v string) *CreateInstancePublishTaskRequest {
	s.InstanceId = &v
	return s
}

type CreateInstancePublishTaskResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateInstancePublishTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancePublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstancePublishTaskResponseBody) SetBizTypeList(v []*string) *CreateInstancePublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetCreateTime(v string) *CreateInstancePublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetError(v string) *CreateInstancePublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetId(v int64) *CreateInstancePublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetModifyTime(v string) *CreateInstancePublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetRequestId(v string) *CreateInstancePublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetResponse(v string) *CreateInstancePublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *CreateInstancePublishTaskResponseBody) SetStatus(v string) *CreateInstancePublishTaskResponseBody {
	s.Status = &v
	return s
}

type CreateInstancePublishTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstancePublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstancePublishTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancePublishTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateInstancePublishTaskResponse) SetHeaders(v map[string]*string) *CreateInstancePublishTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateInstancePublishTaskResponse) SetStatusCode(v int32) *CreateInstancePublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstancePublishTaskResponse) SetBody(v *CreateInstancePublishTaskResponseBody) *CreateInstancePublishTaskResponse {
	s.Body = v
	return s
}

type CreateIntentRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId       *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentDefinition *CreateIntentRequestIntentDefinition `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty" type:"Struct"`
}

func (s CreateIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentRequest) SetAgentKey(v string) *CreateIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateIntentRequest) SetInstanceId(v string) *CreateIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIntentRequest) SetIntentDefinition(v *CreateIntentRequestIntentDefinition) *CreateIntentRequest {
	s.IntentDefinition = v
	return s
}

type CreateIntentRequestIntentDefinition struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// This parameter is required.
	IntentName *string                                         `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	SlotInfos  []*CreateIntentRequestIntentDefinitionSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s CreateIntentRequestIntentDefinition) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentRequestIntentDefinition) GoString() string {
	return s.String()
}

func (s *CreateIntentRequestIntentDefinition) SetAliasName(v string) *CreateIntentRequestIntentDefinition {
	s.AliasName = &v
	return s
}

func (s *CreateIntentRequestIntentDefinition) SetIntentName(v string) *CreateIntentRequestIntentDefinition {
	s.IntentName = &v
	return s
}

func (s *CreateIntentRequestIntentDefinition) SetSlotInfos(v []*CreateIntentRequestIntentDefinitionSlotInfos) *CreateIntentRequestIntentDefinition {
	s.SlotInfos = v
	return s
}

type CreateIntentRequestIntentDefinitionSlotInfos struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// example:
	//
	// false
	Interactive *bool `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fg452dfg3df23
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIntentRequestIntentDefinitionSlotInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentRequestIntentDefinitionSlotInfos) GoString() string {
	return s.String()
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetArray(v bool) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Array = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetEncrypt(v bool) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Encrypt = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetInteractive(v bool) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Interactive = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetName(v string) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Name = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetSlotId(v string) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.SlotId = &v
	return s
}

func (s *CreateIntentRequestIntentDefinitionSlotInfos) SetValue(v string) *CreateIntentRequestIntentDefinitionSlotInfos {
	s.Value = &v
	return s
}

type CreateIntentShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentDefinitionShrink *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
}

func (s CreateIntentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentShrinkRequest) SetAgentKey(v string) *CreateIntentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateIntentShrinkRequest) SetInstanceId(v string) *CreateIntentShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIntentShrinkRequest) SetIntentDefinitionShrink(v string) *CreateIntentShrinkRequest {
	s.IntentDefinitionShrink = &v
	return s
}

type CreateIntentResponseBody struct {
	// example:
	//
	// 43546474
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// df56gjh5et34g3g3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntentResponseBody) SetIntentId(v int64) *CreateIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *CreateIntentResponseBody) SetRequestId(v string) *CreateIntentResponseBody {
	s.RequestId = &v
	return s
}

type CreateIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateIntentResponse) SetStatusCode(v int32) *CreateIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntentResponse) SetBody(v *CreateIntentResponseBody) *CreateIntentResponse {
	s.Body = v
	return s
}

type CreateLgfRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId    *string                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LgfDefinition *CreateLgfRequestLgfDefinition `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty" type:"Struct"`
}

func (s CreateLgfRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLgfRequest) GoString() string {
	return s.String()
}

func (s *CreateLgfRequest) SetAgentKey(v string) *CreateLgfRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateLgfRequest) SetInstanceId(v string) *CreateLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLgfRequest) SetLgfDefinition(v *CreateLgfRequestLgfDefinition) *CreateLgfRequest {
	s.LgfDefinition = v
	return s
}

type CreateLgfRequestLgfDefinition struct {
	// This parameter is required.
	//
	// example:
	//
	// 4675678567
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	RuleText *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s CreateLgfRequestLgfDefinition) String() string {
	return tea.Prettify(s)
}

func (s CreateLgfRequestLgfDefinition) GoString() string {
	return s.String()
}

func (s *CreateLgfRequestLgfDefinition) SetIntentId(v int64) *CreateLgfRequestLgfDefinition {
	s.IntentId = &v
	return s
}

func (s *CreateLgfRequestLgfDefinition) SetRuleText(v string) *CreateLgfRequestLgfDefinition {
	s.RuleText = &v
	return s
}

type CreateLgfShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LgfDefinitionShrink *string `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty"`
}

func (s CreateLgfShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLgfShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLgfShrinkRequest) SetAgentKey(v string) *CreateLgfShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateLgfShrinkRequest) SetInstanceId(v string) *CreateLgfShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLgfShrinkRequest) SetLgfDefinitionShrink(v string) *CreateLgfShrinkRequest {
	s.LgfDefinitionShrink = &v
	return s
}

type CreateLgfResponseBody struct {
	// LGF ID
	//
	// example:
	//
	// 123453433453
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// af5fg3sdf457j5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLgfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLgfResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLgfResponseBody) SetLgfId(v int64) *CreateLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *CreateLgfResponseBody) SetRequestId(v string) *CreateLgfResponseBody {
	s.RequestId = &v
	return s
}

type CreateLgfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLgfResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLgfResponse) GoString() string {
	return s.String()
}

func (s *CreateLgfResponse) SetHeaders(v map[string]*string) *CreateLgfResponse {
	s.Headers = v
	return s
}

func (s *CreateLgfResponse) SetStatusCode(v int32) *CreateLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLgfResponse) SetBody(v *CreateLgfResponseBody) *CreateLgfResponse {
	s.Body = v
	return s
}

type CreatePerspectiveRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 用于购物APP的移动端视角
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 移动端视角
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveRequest) SetAgentKey(v string) *CreatePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *CreatePerspectiveRequest) SetDescription(v string) *CreatePerspectiveRequest {
	s.Description = &v
	return s
}

func (s *CreatePerspectiveRequest) SetName(v string) *CreatePerspectiveRequest {
	s.Name = &v
	return s
}

type CreatePerspectiveResponseBody struct {
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	// example:
	//
	// F285D735-D580-18A8-B97F-B2E72B00F101
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveResponseBody) SetPerspectiveId(v string) *CreatePerspectiveResponseBody {
	s.PerspectiveId = &v
	return s
}

func (s *CreatePerspectiveResponseBody) SetRequestId(v string) *CreatePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

type CreatePerspectiveResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveResponse) SetHeaders(v map[string]*string) *CreatePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *CreatePerspectiveResponse) SetStatusCode(v int32) *CreatePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePerspectiveResponse) SetBody(v *CreatePerspectiveResponseBody) *CreatePerspectiveResponse {
	s.Body = v
	return s
}

type CreatePublishTaskRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// faq
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// ["8521"]
	DataIdList []*string `json:"DataIdList,omitempty" xml:"DataIdList,omitempty" type:"Repeated"`
}

func (s CreatePublishTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePublishTaskRequest) SetAgentKey(v string) *CreatePublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CreatePublishTaskRequest) SetBizType(v string) *CreatePublishTaskRequest {
	s.BizType = &v
	return s
}

func (s *CreatePublishTaskRequest) SetDataIdList(v []*string) *CreatePublishTaskRequest {
	s.DataIdList = v
	return s
}

type CreatePublishTaskShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// faq
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// ["8521"]
	DataIdListShrink *string `json:"DataIdList,omitempty" xml:"DataIdList,omitempty"`
}

func (s CreatePublishTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePublishTaskShrinkRequest) SetAgentKey(v string) *CreatePublishTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreatePublishTaskShrinkRequest) SetBizType(v string) *CreatePublishTaskShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CreatePublishTaskShrinkRequest) SetDataIdListShrink(v string) *CreatePublishTaskShrinkRequest {
	s.DataIdListShrink = &v
	return s
}

type CreatePublishTaskResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreatePublishTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublishTaskResponseBody) SetBizTypeList(v []*string) *CreatePublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *CreatePublishTaskResponseBody) SetCreateTime(v string) *CreatePublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetError(v string) *CreatePublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetId(v int64) *CreatePublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetModifyTime(v string) *CreatePublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetRequestId(v string) *CreatePublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetResponse(v string) *CreatePublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *CreatePublishTaskResponseBody) SetStatus(v string) *CreatePublishTaskResponseBody {
	s.Status = &v
	return s
}

type CreatePublishTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublishTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePublishTaskResponse) SetHeaders(v map[string]*string) *CreatePublishTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePublishTaskResponse) SetStatusCode(v int32) *CreatePublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublishTaskResponse) SetBody(v *CreatePublishTaskResponseBody) *CreatePublishTaskResponse {
	s.Body = v
	return s
}

type CreateSimQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateSimQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSimQuestionRequest) GoString() string {
	return s.String()
}

func (s *CreateSimQuestionRequest) SetAgentKey(v string) *CreateSimQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateSimQuestionRequest) SetKnowledgeId(v int64) *CreateSimQuestionRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateSimQuestionRequest) SetTitle(v string) *CreateSimQuestionRequest {
	s.Title = &v
	return s
}

type CreateSimQuestionResponseBody struct {
	// example:
	//
	// 16AC1B3C-66E0-438B-BB7C-71B692407B67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000002788
	SimQuestionId *int64 `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
}

func (s CreateSimQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSimQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimQuestionResponseBody) SetRequestId(v string) *CreateSimQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimQuestionResponseBody) SetSimQuestionId(v int64) *CreateSimQuestionResponseBody {
	s.SimQuestionId = &v
	return s
}

type CreateSimQuestionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSimQuestionResponse) GoString() string {
	return s.String()
}

func (s *CreateSimQuestionResponse) SetHeaders(v map[string]*string) *CreateSimQuestionResponse {
	s.Headers = v
	return s
}

func (s *CreateSimQuestionResponse) SetStatusCode(v int32) *CreateSimQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimQuestionResponse) SetBody(v *CreateSimQuestionResponseBody) *CreateSimQuestionResponse {
	s.Body = v
	return s
}

type CreateSolutionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// This parameter is required.
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
}

func (s CreateSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSolutionRequest) GoString() string {
	return s.String()
}

func (s *CreateSolutionRequest) SetAgentKey(v string) *CreateSolutionRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateSolutionRequest) SetContent(v string) *CreateSolutionRequest {
	s.Content = &v
	return s
}

func (s *CreateSolutionRequest) SetContentType(v int32) *CreateSolutionRequest {
	s.ContentType = &v
	return s
}

func (s *CreateSolutionRequest) SetKnowledgeId(v int64) *CreateSolutionRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateSolutionRequest) SetPerspectiveCodes(v []*string) *CreateSolutionRequest {
	s.PerspectiveCodes = v
	return s
}

type CreateSolutionResponseBody struct {
	// example:
	//
	// F55D90C1-31BE-4B2A-AA3F-25EFC36F9419
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100001089003
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s CreateSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSolutionResponseBody) SetRequestId(v string) *CreateSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSolutionResponseBody) SetSolutionId(v int64) *CreateSolutionResponseBody {
	s.SolutionId = &v
	return s
}

type CreateSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSolutionResponse) GoString() string {
	return s.String()
}

func (s *CreateSolutionResponse) SetHeaders(v map[string]*string) *CreateSolutionResponse {
	s.Headers = v
	return s
}

func (s *CreateSolutionResponse) SetStatusCode(v int32) *CreateSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSolutionResponse) SetBody(v *CreateSolutionResponseBody) *CreateSolutionResponse {
	s.Body = v
	return s
}

type CreateUserSayRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId        *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserSayDefinition *CreateUserSayRequestUserSayDefinition `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty" type:"Struct"`
}

func (s CreateUserSayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSayRequest) GoString() string {
	return s.String()
}

func (s *CreateUserSayRequest) SetAgentKey(v string) *CreateUserSayRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateUserSayRequest) SetInstanceId(v string) *CreateUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserSayRequest) SetUserSayDefinition(v *CreateUserSayRequestUserSayDefinition) *CreateUserSayRequest {
	s.UserSayDefinition = v
	return s
}

type CreateUserSayRequestUserSayDefinition struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123232
	IntentId  *int64                                            `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	SlotInfos []*CreateUserSayRequestUserSayDefinitionSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s CreateUserSayRequestUserSayDefinition) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSayRequestUserSayDefinition) GoString() string {
	return s.String()
}

func (s *CreateUserSayRequestUserSayDefinition) SetContent(v string) *CreateUserSayRequestUserSayDefinition {
	s.Content = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinition) SetIntentId(v int64) *CreateUserSayRequestUserSayDefinition {
	s.IntentId = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinition) SetSlotInfos(v []*CreateUserSayRequestUserSayDefinitionSlotInfos) *CreateUserSayRequestUserSayDefinition {
	s.SlotInfos = v
	return s
}

type CreateUserSayRequestUserSayDefinitionSlotInfos struct {
	// example:
	//
	// 6
	EndIndex *int32 `json:"EndIndex,omitempty" xml:"EndIndex,omitempty"`
	// example:
	//
	// fb34adf2fv43f2
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// example:
	//
	// 4
	StartIndex *int32 `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
}

func (s CreateUserSayRequestUserSayDefinitionSlotInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSayRequestUserSayDefinitionSlotInfos) GoString() string {
	return s.String()
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) SetEndIndex(v int32) *CreateUserSayRequestUserSayDefinitionSlotInfos {
	s.EndIndex = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) SetSlotId(v string) *CreateUserSayRequestUserSayDefinitionSlotInfos {
	s.SlotId = &v
	return s
}

func (s *CreateUserSayRequestUserSayDefinitionSlotInfos) SetStartIndex(v int32) *CreateUserSayRequestUserSayDefinitionSlotInfos {
	s.StartIndex = &v
	return s
}

type CreateUserSayShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserSayDefinitionShrink *string `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty"`
}

func (s CreateUserSayShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSayShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserSayShrinkRequest) SetAgentKey(v string) *CreateUserSayShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateUserSayShrinkRequest) SetInstanceId(v string) *CreateUserSayShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserSayShrinkRequest) SetUserSayDefinitionShrink(v string) *CreateUserSayShrinkRequest {
	s.UserSayDefinitionShrink = &v
	return s
}

type CreateUserSayResponseBody struct {
	// example:
	//
	// 8g4n8bnd236fg79
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 46456176856
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s CreateUserSayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserSayResponseBody) SetRequestId(v string) *CreateUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserSayResponseBody) SetUserSayId(v int64) *CreateUserSayResponseBody {
	s.UserSayId = &v
	return s
}

type CreateUserSayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserSayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSayResponse) GoString() string {
	return s.String()
}

func (s *CreateUserSayResponse) SetHeaders(v map[string]*string) *CreateUserSayResponse {
	s.Headers = v
	return s
}

func (s *CreateUserSayResponse) SetStatusCode(v int32) *CreateUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserSayResponse) SetBody(v *CreateUserSayResponseBody) *CreateUserSayResponse {
	s.Body = v
	return s
}

type DeleteCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DeleteCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteCategoryRequest) SetAgentKey(v string) *DeleteCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCategoryRequest) SetCategoryId(v int64) *DeleteCategoryRequest {
	s.CategoryId = &v
	return s
}

type DeleteCategoryResponseBody struct {
	// example:
	//
	// 4B3E0DE3-DC57-5BFE-88D4-ADD8ED024F55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DeleteCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteCategoryResponse) SetStatusCode(v int32) *DeleteCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCategoryResponse) SetBody(v *DeleteCategoryResponseBody) *DeleteCategoryResponse {
	s.Body = v
	return s
}

type DeleteConnQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 877397683
	OutlineId *int64 `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
}

func (s DeleteConnQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnQuestionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnQuestionRequest) SetAgentKey(v string) *DeleteConnQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteConnQuestionRequest) SetOutlineId(v int64) *DeleteConnQuestionRequest {
	s.OutlineId = &v
	return s
}

type DeleteConnQuestionResponseBody struct {
	// example:
	//
	// FC323352-3AD7-59A1-9088-A64470BAFC9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConnQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnQuestionResponseBody) SetRequestId(v string) *DeleteConnQuestionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConnQuestionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConnQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConnQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnQuestionResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnQuestionResponse) SetHeaders(v map[string]*string) *DeleteConnQuestionResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnQuestionResponse) SetStatusCode(v int32) *DeleteConnQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnQuestionResponse) SetBody(v *DeleteConnQuestionResponseBody) *DeleteConnQuestionResponse {
	s.Body = v
	return s
}

type DeleteDSEntityRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDSEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDSEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityRequest) SetAgentKey(v string) *DeleteDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDSEntityRequest) SetEntityId(v int64) *DeleteDSEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteDSEntityRequest) SetInstanceId(v string) *DeleteDSEntityRequest {
	s.InstanceId = &v
	return s
}

type DeleteDSEntityResponseBody struct {
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 4dfghf56235asdf452
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDSEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityResponseBody) SetEntityId(v int64) *DeleteDSEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *DeleteDSEntityResponseBody) SetRequestId(v string) *DeleteDSEntityResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDSEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDSEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDSEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityResponse) SetHeaders(v map[string]*string) *DeleteDSEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteDSEntityResponse) SetStatusCode(v int32) *DeleteDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDSEntityResponse) SetBody(v *DeleteDSEntityResponseBody) *DeleteDSEntityResponse {
	s.Body = v
	return s
}

type DeleteDSEntityValueRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 345346223452
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3453453452
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDSEntityValueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDSEntityValueRequest) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityValueRequest) SetAgentKey(v string) *DeleteDSEntityValueRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDSEntityValueRequest) SetEntityId(v int64) *DeleteDSEntityValueRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteDSEntityValueRequest) SetEntityValueId(v int64) *DeleteDSEntityValueRequest {
	s.EntityValueId = &v
	return s
}

func (s *DeleteDSEntityValueRequest) SetInstanceId(v string) *DeleteDSEntityValueRequest {
	s.InstanceId = &v
	return s
}

type DeleteDSEntityValueResponseBody struct {
	// example:
	//
	// 3453453452
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// example:
	//
	// dfdf2t3rfvb45y
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDSEntityValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDSEntityValueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityValueResponseBody) SetEntityValueId(v int64) *DeleteDSEntityValueResponseBody {
	s.EntityValueId = &v
	return s
}

func (s *DeleteDSEntityValueResponseBody) SetRequestId(v string) *DeleteDSEntityValueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDSEntityValueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDSEntityValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDSEntityValueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDSEntityValueResponse) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityValueResponse) SetHeaders(v map[string]*string) *DeleteDSEntityValueResponse {
	s.Headers = v
	return s
}

func (s *DeleteDSEntityValueResponse) SetStatusCode(v int32) *DeleteDSEntityValueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDSEntityValueResponse) SetBody(v *DeleteDSEntityValueResponseBody) *DeleteDSEntityValueResponse {
	s.Body = v
	return s
}

type DeleteDocRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DeleteDocRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocRequest) SetAgentKey(v string) *DeleteDocRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDocRequest) SetKnowledgeId(v int64) *DeleteDocRequest {
	s.KnowledgeId = &v
	return s
}

type DeleteDocResponseBody struct {
	// example:
	//
	// 30002406051
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// DFB71B34-4188-4EA2-9988-EF3014E75910
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocResponseBody) SetKnowledgeId(v int64) *DeleteDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *DeleteDocResponseBody) SetRequestId(v string) *DeleteDocResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocResponse) SetHeaders(v map[string]*string) *DeleteDocResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocResponse) SetStatusCode(v int32) *DeleteDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocResponse) SetBody(v *DeleteDocResponseBody) *DeleteDocResponse {
	s.Body = v
	return s
}

type DeleteFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DeleteFaqRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaqRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaqRequest) SetAgentKey(v string) *DeleteFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteFaqRequest) SetKnowledgeId(v int64) *DeleteFaqRequest {
	s.KnowledgeId = &v
	return s
}

type DeleteFaqResponseBody struct {
	// example:
	//
	// F79E7305-5314-5069-A701-9591AD051902
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaqResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaqResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaqResponseBody) SetRequestId(v string) *DeleteFaqResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaqResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaqResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaqResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaqResponse) SetHeaders(v map[string]*string) *DeleteFaqResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaqResponse) SetStatusCode(v int32) *DeleteFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaqResponse) SetBody(v *DeleteFaqResponseBody) *DeleteFaqResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetAgentKey(v string) *DeleteInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-09-11T09:26:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 9052
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// xuqiang_test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	Error          *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8521
	Response *int64 `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetBizTypeList(v []*string) *DeleteInstanceResponseBody {
	s.BizTypeList = v
	return s
}

func (s *DeleteInstanceResponseBody) SetCreateTime(v string) *DeleteInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetCreateUserId(v int64) *DeleteInstanceResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetCreateUserName(v string) *DeleteInstanceResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetError(v string) *DeleteInstanceResponseBody {
	s.Error = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetId(v int64) *DeleteInstanceResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetResponse(v int64) *DeleteInstanceResponseBody {
	s.Response = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetStatus(v string) *DeleteInstanceResponseBody {
	s.Status = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteIntentRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s DeleteIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntentRequest) SetAgentKey(v string) *DeleteIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteIntentRequest) SetInstanceId(v string) *DeleteIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIntentRequest) SetIntentId(v int64) *DeleteIntentRequest {
	s.IntentId = &v
	return s
}

type DeleteIntentResponseBody struct {
	// example:
	//
	// 12345
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// 3464dfg3qwr34tf34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntentResponseBody) SetIntentId(v int64) *DeleteIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *DeleteIntentResponseBody) SetRequestId(v string) *DeleteIntentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteIntentResponse) SetStatusCode(v int32) *DeleteIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntentResponse) SetBody(v *DeleteIntentResponseBody) *DeleteIntentResponse {
	s.Body = v
	return s
}

type DeleteLgfRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23242342
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// lgf Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 2342424
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
}

func (s DeleteLgfRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLgfRequest) GoString() string {
	return s.String()
}

func (s *DeleteLgfRequest) SetAgentKey(v string) *DeleteLgfRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteLgfRequest) SetInstanceId(v string) *DeleteLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteLgfRequest) SetIntentId(v int64) *DeleteLgfRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteLgfRequest) SetLgfId(v int64) *DeleteLgfRequest {
	s.LgfId = &v
	return s
}

type DeleteLgfResponseBody struct {
	// LGF ID
	//
	// example:
	//
	// 2342424
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// dgw2342424qw42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLgfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLgfResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLgfResponseBody) SetLgfId(v int64) *DeleteLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *DeleteLgfResponseBody) SetRequestId(v string) *DeleteLgfResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLgfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLgfResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLgfResponse) GoString() string {
	return s.String()
}

func (s *DeleteLgfResponse) SetHeaders(v map[string]*string) *DeleteLgfResponse {
	s.Headers = v
	return s
}

func (s *DeleteLgfResponse) SetStatusCode(v int32) *DeleteLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLgfResponse) SetBody(v *DeleteLgfResponseBody) *DeleteLgfResponse {
	s.Body = v
	return s
}

type DeletePerspectiveRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s DeletePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveRequest) SetAgentKey(v string) *DeletePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *DeletePerspectiveRequest) SetPerspectiveId(v string) *DeletePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type DeletePerspectiveResponseBody struct {
	// example:
	//
	// FC384CE1-8D42-1900-84E1-F33F990F2B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeletePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveResponseBody) SetRequestId(v string) *DeletePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePerspectiveResponseBody) SetResult(v bool) *DeletePerspectiveResponseBody {
	s.Result = &v
	return s
}

type DeletePerspectiveResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePerspectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveResponse) SetHeaders(v map[string]*string) *DeletePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *DeletePerspectiveResponse) SetStatusCode(v int32) *DeletePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePerspectiveResponse) SetBody(v *DeletePerspectiveResponseBody) *DeletePerspectiveResponse {
	s.Body = v
	return s
}

type DeleteSimQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000002788
	SimQuestionId *int64 `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
}

func (s DeleteSimQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimQuestionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSimQuestionRequest) SetAgentKey(v string) *DeleteSimQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteSimQuestionRequest) SetSimQuestionId(v int64) *DeleteSimQuestionRequest {
	s.SimQuestionId = &v
	return s
}

type DeleteSimQuestionResponseBody struct {
	// example:
	//
	// 6419BA93-D111-5225-8998-13E63E6D3940
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSimQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSimQuestionResponseBody) SetRequestId(v string) *DeleteSimQuestionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSimQuestionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSimQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSimQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimQuestionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSimQuestionResponse) SetHeaders(v map[string]*string) *DeleteSimQuestionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSimQuestionResponse) SetStatusCode(v int32) *DeleteSimQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSimQuestionResponse) SetBody(v *DeleteSimQuestionResponseBody) *DeleteSimQuestionResponse {
	s.Body = v
	return s
}

type DeleteSolutionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100001321580
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DeleteSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSolutionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSolutionRequest) SetAgentKey(v string) *DeleteSolutionRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteSolutionRequest) SetSolutionId(v int64) *DeleteSolutionRequest {
	s.SolutionId = &v
	return s
}

type DeleteSolutionResponseBody struct {
	// example:
	//
	// F79E7305-5314-5069-A701-9591AD051902
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSolutionResponseBody) SetRequestId(v string) *DeleteSolutionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSolutionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSolutionResponse) SetHeaders(v map[string]*string) *DeleteSolutionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSolutionResponse) SetStatusCode(v int32) *DeleteSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSolutionResponse) SetBody(v *DeleteSolutionResponseBody) *DeleteSolutionResponse {
	s.Body = v
	return s
}

type DeleteUserSayRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 343df2sdf23
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5564564546
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4562121234
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s DeleteUserSayRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserSayRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserSayRequest) SetAgentKey(v string) *DeleteUserSayRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteUserSayRequest) SetInstanceId(v string) *DeleteUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserSayRequest) SetIntentId(v int64) *DeleteUserSayRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteUserSayRequest) SetUserSayId(v int64) *DeleteUserSayRequest {
	s.UserSayId = &v
	return s
}

type DeleteUserSayResponseBody struct {
	// example:
	//
	// dfgdg324gf34t34g34g3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4562121234
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s DeleteUserSayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserSayResponseBody) SetRequestId(v string) *DeleteUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserSayResponseBody) SetUserSayId(v int64) *DeleteUserSayResponseBody {
	s.UserSayId = &v
	return s
}

type DeleteUserSayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserSayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserSayResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserSayResponse) SetHeaders(v map[string]*string) *DeleteUserSayResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserSayResponse) SetStatusCode(v int32) *DeleteUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserSayResponse) SetBody(v *DeleteUserSayResponseBody) *DeleteUserSayResponse {
	s.Body = v
	return s
}

type DescribeCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DescribeCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCategoryRequest) SetAgentKey(v string) *DescribeCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeCategoryRequest) SetCategoryId(v int64) *DescribeCategoryRequest {
	s.CategoryId = &v
	return s
}

type DescribeCategoryResponseBody struct {
	Category *DescribeCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// example:
	//
	// 2B0304FD-3804-5C06-9A83-77F5523664AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponseBody) SetCategory(v *DescribeCategoryResponseBodyCategory) *DescribeCategoryResponseBody {
	s.Category = v
	return s
}

func (s *DescribeCategoryResponseBody) SetRequestId(v string) *DescribeCategoryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCategoryResponseBodyCategory struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 30000049006
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCategoryResponseBodyCategory) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponseBodyCategory) SetBizCode(v string) *DescribeCategoryResponseBodyCategory {
	s.BizCode = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) SetCategoryId(v int64) *DescribeCategoryResponseBodyCategory {
	s.CategoryId = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) SetName(v string) *DescribeCategoryResponseBodyCategory {
	s.Name = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) SetParentCategoryId(v int64) *DescribeCategoryResponseBodyCategory {
	s.ParentCategoryId = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) SetStatus(v int32) *DescribeCategoryResponseBodyCategory {
	s.Status = &v
	return s
}

type DescribeCategoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeCategoryResponse) SetStatusCode(v int32) *DescribeCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCategoryResponse) SetBody(v *DescribeCategoryResponseBody) *DescribeCategoryResponse {
	s.Body = v
	return s
}

type DescribeDSEntityRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDSEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDSEntityRequest) GoString() string {
	return s.String()
}

func (s *DescribeDSEntityRequest) SetAgentKey(v string) *DescribeDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeDSEntityRequest) SetEntityId(v int64) *DescribeDSEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeDSEntityRequest) SetInstanceId(v string) *DescribeDSEntityRequest {
	s.InstanceId = &v
	return s
}

type DescribeDSEntityResponseBody struct {
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123231
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 123
	EntityId   *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// synonyms
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 123231
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// test
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// ad23234dsf234fga
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SysEntityCode *string `json:"SysEntityCode,omitempty" xml:"SysEntityCode,omitempty"`
}

func (s DescribeDSEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDSEntityResponseBody) SetCreateTime(v string) *DescribeDSEntityResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetCreateUserId(v string) *DescribeDSEntityResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetCreateUserName(v string) *DescribeDSEntityResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetEntityId(v int64) *DescribeDSEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetEntityName(v string) *DescribeDSEntityResponseBody {
	s.EntityName = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetEntityType(v string) *DescribeDSEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetModifyTime(v string) *DescribeDSEntityResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetModifyUserId(v string) *DescribeDSEntityResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetModifyUserName(v string) *DescribeDSEntityResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetRequestId(v string) *DescribeDSEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetSysEntityCode(v string) *DescribeDSEntityResponseBody {
	s.SysEntityCode = &v
	return s
}

type DescribeDSEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDSEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDSEntityResponse) GoString() string {
	return s.String()
}

func (s *DescribeDSEntityResponse) SetHeaders(v map[string]*string) *DescribeDSEntityResponse {
	s.Headers = v
	return s
}

func (s *DescribeDSEntityResponse) SetStatusCode(v int32) *DescribeDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDSEntityResponse) SetBody(v *DescribeDSEntityResponseBody) *DescribeDSEntityResponse {
	s.Body = v
	return s
}

type DescribeDocRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// false
	ShowDetail *bool `json:"ShowDetail,omitempty" xml:"ShowDetail,omitempty"`
}

func (s DescribeDocRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocRequest) SetAgentKey(v string) *DescribeDocRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeDocRequest) SetKnowledgeId(v int64) *DescribeDocRequest {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeDocRequest) SetShowDetail(v bool) *DescribeDocRequest {
	s.ShowDetail = &v
	return s
}

type DescribeDocResponseBody struct {
	// example:
	//
	// bizcode123
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":"docName"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1111111111
	CreateUserId   *int64                                `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string                               `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DocInfo        *DescribeDocResponseBodyDocInfo       `json:"DocInfo,omitempty" xml:"DocInfo,omitempty" type:"Struct"`
	DocMetadata    []*DescribeDocResponseBodyDocMetadata `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty" type:"Repeated"`
	DocName        *string                               `json:"DocName,omitempty" xml:"DocName,omitempty"`
	DocTags        []*DescribeDocResponseBodyDocTags     `json:"DocTags,omitempty" xml:"DocTags,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	EffectStatus *int32 `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// example:
	//
	// 2023-04-27T06:08:54Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2020-11-25T08:56:55Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 2222222222
	ModifyUserId   *int64  `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// true
	ProcessCanRetry *bool   `json:"ProcessCanRetry,omitempty" xml:"ProcessCanRetry,omitempty"`
	ProcessMessage  *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty"`
	// example:
	//
	// 0
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F132693-212A-40A9-8A81-11E7694E478B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1979-12-31T16:00:00Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 1
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBody) SetBizCode(v string) *DescribeDocResponseBody {
	s.BizCode = &v
	return s
}

func (s *DescribeDocResponseBody) SetCategoryId(v int64) *DescribeDocResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeDocResponseBody) SetConfig(v string) *DescribeDocResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDocResponseBody) SetCreateTime(v string) *DescribeDocResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDocResponseBody) SetCreateUserId(v int64) *DescribeDocResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeDocResponseBody) SetCreateUserName(v string) *DescribeDocResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeDocResponseBody) SetDocInfo(v *DescribeDocResponseBodyDocInfo) *DescribeDocResponseBody {
	s.DocInfo = v
	return s
}

func (s *DescribeDocResponseBody) SetDocMetadata(v []*DescribeDocResponseBodyDocMetadata) *DescribeDocResponseBody {
	s.DocMetadata = v
	return s
}

func (s *DescribeDocResponseBody) SetDocName(v string) *DescribeDocResponseBody {
	s.DocName = &v
	return s
}

func (s *DescribeDocResponseBody) SetDocTags(v []*DescribeDocResponseBodyDocTags) *DescribeDocResponseBody {
	s.DocTags = v
	return s
}

func (s *DescribeDocResponseBody) SetEffectStatus(v int32) *DescribeDocResponseBody {
	s.EffectStatus = &v
	return s
}

func (s *DescribeDocResponseBody) SetEndDate(v string) *DescribeDocResponseBody {
	s.EndDate = &v
	return s
}

func (s *DescribeDocResponseBody) SetKnowledgeId(v int64) *DescribeDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeDocResponseBody) SetMeta(v string) *DescribeDocResponseBody {
	s.Meta = &v
	return s
}

func (s *DescribeDocResponseBody) SetModifyTime(v string) *DescribeDocResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDocResponseBody) SetModifyUserId(v int64) *DescribeDocResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeDocResponseBody) SetModifyUserName(v string) *DescribeDocResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeDocResponseBody) SetProcessCanRetry(v bool) *DescribeDocResponseBody {
	s.ProcessCanRetry = &v
	return s
}

func (s *DescribeDocResponseBody) SetProcessMessage(v string) *DescribeDocResponseBody {
	s.ProcessMessage = &v
	return s
}

func (s *DescribeDocResponseBody) SetProcessStatus(v int32) *DescribeDocResponseBody {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeDocResponseBody) SetRequestId(v string) *DescribeDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocResponseBody) SetStartDate(v string) *DescribeDocResponseBody {
	s.StartDate = &v
	return s
}

func (s *DescribeDocResponseBody) SetStatus(v int32) *DescribeDocResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDocResponseBody) SetTitle(v string) *DescribeDocResponseBody {
	s.Title = &v
	return s
}

func (s *DescribeDocResponseBody) SetUrl(v string) *DescribeDocResponseBody {
	s.Url = &v
	return s
}

type DescribeDocResponseBodyDocInfo struct {
	DocParas []*DescribeDocResponseBodyDocInfoDocParas `json:"DocParas,omitempty" xml:"DocParas,omitempty" type:"Repeated"`
}

func (s DescribeDocResponseBodyDocInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponseBodyDocInfo) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocInfo) SetDocParas(v []*DescribeDocResponseBodyDocInfoDocParas) *DescribeDocResponseBodyDocInfo {
	s.DocParas = v
	return s
}

type DescribeDocResponseBodyDocInfoDocParas struct {
	// example:
	//
	// 1
	ParaLevel *int32 `json:"ParaLevel,omitempty" xml:"ParaLevel,omitempty"`
	// example:
	//
	// 1
	ParaNo   *int32  `json:"ParaNo,omitempty" xml:"ParaNo,omitempty"`
	ParaText *string `json:"ParaText,omitempty" xml:"ParaText,omitempty"`
	// example:
	//
	// text
	ParaType *string `json:"ParaType,omitempty" xml:"ParaType,omitempty"`
}

func (s DescribeDocResponseBodyDocInfoDocParas) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponseBodyDocInfoDocParas) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocInfoDocParas) SetParaLevel(v int32) *DescribeDocResponseBodyDocInfoDocParas {
	s.ParaLevel = &v
	return s
}

func (s *DescribeDocResponseBodyDocInfoDocParas) SetParaNo(v int32) *DescribeDocResponseBodyDocInfoDocParas {
	s.ParaNo = &v
	return s
}

func (s *DescribeDocResponseBodyDocInfoDocParas) SetParaText(v string) *DescribeDocResponseBodyDocInfoDocParas {
	s.ParaText = &v
	return s
}

func (s *DescribeDocResponseBodyDocInfoDocParas) SetParaType(v string) *DescribeDocResponseBodyDocInfoDocParas {
	s.ParaType = &v
	return s
}

type DescribeDocResponseBodyDocMetadata struct {
	BusinessViewId      *string                                                  `json:"BusinessViewId,omitempty" xml:"BusinessViewId,omitempty"`
	BusinessViewName    *string                                                  `json:"BusinessViewName,omitempty" xml:"BusinessViewName,omitempty"`
	MetaCellInfoDTOList []*DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList `json:"MetaCellInfoDTOList,omitempty" xml:"MetaCellInfoDTOList,omitempty" type:"Repeated"`
}

func (s DescribeDocResponseBodyDocMetadata) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponseBodyDocMetadata) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocMetadata) SetBusinessViewId(v string) *DescribeDocResponseBodyDocMetadata {
	s.BusinessViewId = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadata) SetBusinessViewName(v string) *DescribeDocResponseBodyDocMetadata {
	s.BusinessViewName = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadata) SetMetaCellInfoDTOList(v []*DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) *DescribeDocResponseBodyDocMetadata {
	s.MetaCellInfoDTOList = v
	return s
}

type DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList struct {
	FieldCode *string `json:"FieldCode,omitempty" xml:"FieldCode,omitempty"`
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) SetFieldCode(v string) *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList {
	s.FieldCode = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) SetFieldName(v string) *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList {
	s.FieldName = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) SetValue(v string) *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList {
	s.Value = &v
	return s
}

type DescribeDocResponseBodyDocTags struct {
	DefaultTag *bool   `json:"DefaultTag,omitempty" xml:"DefaultTag,omitempty"`
	GroupId    *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	TagId      *int64  `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName    *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeDocResponseBodyDocTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponseBodyDocTags) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocTags) SetDefaultTag(v bool) *DescribeDocResponseBodyDocTags {
	s.DefaultTag = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) SetGroupId(v int64) *DescribeDocResponseBodyDocTags {
	s.GroupId = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) SetGroupName(v string) *DescribeDocResponseBodyDocTags {
	s.GroupName = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) SetTagId(v int64) *DescribeDocResponseBodyDocTags {
	s.TagId = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) SetTagName(v string) *DescribeDocResponseBodyDocTags {
	s.TagName = &v
	return s
}

type DescribeDocResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDocResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocResponse) SetHeaders(v map[string]*string) *DescribeDocResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocResponse) SetStatusCode(v int32) *DescribeDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocResponse) SetBody(v *DescribeDocResponseBody) *DescribeDocResponse {
	s.Body = v
	return s
}

type DescribeFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DescribeFaqRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaqRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaqRequest) SetAgentKey(v string) *DescribeFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeFaqRequest) SetKnowledgeId(v int64) *DescribeFaqRequest {
	s.KnowledgeId = &v
	return s
}

type DescribeFaqResponseBody struct {
	// example:
	//
	// 30000055617
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2020-11-30T03:03:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test01
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 20
	EffectStatus *int32 `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// example:
	//
	// 2023-04-27T06:08:54Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 2020-12-02T06:35:50Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test01
	ModifyUserName *string                            `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Outlines       []*DescribeFaqResponseBodyOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 8AD9FA10-7780-5E12-B701-13C928524F32
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SimQuestions []*DescribeFaqResponseBodySimQuestions `json:"SimQuestions,omitempty" xml:"SimQuestions,omitempty" type:"Repeated"`
	Solutions    []*DescribeFaqResponseBodySolutions    `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-27T07:04:39Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 3
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeFaqResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaqResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponseBody) SetCategoryId(v int64) *DescribeFaqResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeFaqResponseBody) SetCreateTime(v string) *DescribeFaqResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaqResponseBody) SetCreateUserName(v string) *DescribeFaqResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeFaqResponseBody) SetEffectStatus(v int32) *DescribeFaqResponseBody {
	s.EffectStatus = &v
	return s
}

func (s *DescribeFaqResponseBody) SetEndDate(v string) *DescribeFaqResponseBody {
	s.EndDate = &v
	return s
}

func (s *DescribeFaqResponseBody) SetKnowledgeId(v int64) *DescribeFaqResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeFaqResponseBody) SetModifyTime(v string) *DescribeFaqResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeFaqResponseBody) SetModifyUserName(v string) *DescribeFaqResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeFaqResponseBody) SetOutlines(v []*DescribeFaqResponseBodyOutlines) *DescribeFaqResponseBody {
	s.Outlines = v
	return s
}

func (s *DescribeFaqResponseBody) SetRequestId(v string) *DescribeFaqResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaqResponseBody) SetSimQuestions(v []*DescribeFaqResponseBodySimQuestions) *DescribeFaqResponseBody {
	s.SimQuestions = v
	return s
}

func (s *DescribeFaqResponseBody) SetSolutions(v []*DescribeFaqResponseBodySolutions) *DescribeFaqResponseBody {
	s.Solutions = v
	return s
}

func (s *DescribeFaqResponseBody) SetStartDate(v string) *DescribeFaqResponseBody {
	s.StartDate = &v
	return s
}

func (s *DescribeFaqResponseBody) SetStatus(v int32) *DescribeFaqResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFaqResponseBody) SetTitle(v string) *DescribeFaqResponseBody {
	s.Title = &v
	return s
}

type DescribeFaqResponseBodyOutlines struct {
	// example:
	//
	// 1000098002
	ConnQuestionId *int64 `json:"ConnQuestionId,omitempty" xml:"ConnQuestionId,omitempty"`
	// example:
	//
	// 2022-05-26T10:24:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-26T18:12:02Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 797
	OutlineId *int64  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeFaqResponseBodyOutlines) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaqResponseBodyOutlines) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponseBodyOutlines) SetConnQuestionId(v int64) *DescribeFaqResponseBodyOutlines {
	s.ConnQuestionId = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) SetCreateTime(v string) *DescribeFaqResponseBodyOutlines {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) SetModifyTime(v string) *DescribeFaqResponseBodyOutlines {
	s.ModifyTime = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) SetOutlineId(v int64) *DescribeFaqResponseBodyOutlines {
	s.OutlineId = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) SetTitle(v string) *DescribeFaqResponseBodyOutlines {
	s.Title = &v
	return s
}

type DescribeFaqResponseBodySimQuestions struct {
	// example:
	//
	// 2022-05-26T10:24:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-29T03:55:07Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 10000000581
	SimQuestionId *int64  `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeFaqResponseBodySimQuestions) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaqResponseBodySimQuestions) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponseBodySimQuestions) SetCreateTime(v string) *DescribeFaqResponseBodySimQuestions {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaqResponseBodySimQuestions) SetModifyTime(v string) *DescribeFaqResponseBodySimQuestions {
	s.ModifyTime = &v
	return s
}

func (s *DescribeFaqResponseBodySimQuestions) SetSimQuestionId(v int64) *DescribeFaqResponseBodySimQuestions {
	s.SimQuestionId = &v
	return s
}

func (s *DescribeFaqResponseBodySimQuestions) SetTitle(v string) *DescribeFaqResponseBodySimQuestions {
	s.Title = &v
	return s
}

type DescribeFaqResponseBodySolutions struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 0
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 2022-05-26T10:24:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-29T07:07:13Z
	ModifyTime       *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
	PlainText        *string   `json:"PlainText,omitempty" xml:"PlainText,omitempty"`
	// example:
	//
	// 10000003071
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DescribeFaqResponseBodySolutions) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaqResponseBodySolutions) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponseBodySolutions) SetContent(v string) *DescribeFaqResponseBodySolutions {
	s.Content = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetContentType(v int32) *DescribeFaqResponseBodySolutions {
	s.ContentType = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetCreateTime(v string) *DescribeFaqResponseBodySolutions {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetModifyTime(v string) *DescribeFaqResponseBodySolutions {
	s.ModifyTime = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetPerspectiveCodes(v []*string) *DescribeFaqResponseBodySolutions {
	s.PerspectiveCodes = v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetPlainText(v string) *DescribeFaqResponseBodySolutions {
	s.PlainText = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetSolutionId(v int64) *DescribeFaqResponseBodySolutions {
	s.SolutionId = &v
	return s
}

type DescribeFaqResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaqResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaqResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponse) SetHeaders(v map[string]*string) *DescribeFaqResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaqResponse) SetStatusCode(v int32) *DescribeFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaqResponse) SetBody(v *DescribeFaqResponseBody) *DescribeFaqResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetAgentKey(v string) *DescribeInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceResponseBody struct {
	// example:
	//
	// /alimefe/meebot/robot/0.0.5/img/xxx-90-97.png
	Avatar     *string                                   `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Categories []*DescribeInstanceResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// PUBLISHED
	EditStatus *string `json:"EditStatus,omitempty" xml:"EditStatus,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 用于C端问答的机器人
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// example:
	//
	// 智能客服-小C
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 907AA5F2-0521-49AB-80AB-1ADEFAB2B901
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// scenario_im
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
	// example:
	//
	// Asia/Chongqing
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetAvatar(v string) *DescribeInstanceResponseBody {
	s.Avatar = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCategories(v []*DescribeInstanceResponseBodyCategories) *DescribeInstanceResponseBody {
	s.Categories = v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreateTime(v string) *DescribeInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEditStatus(v string) *DescribeInstanceResponseBody {
	s.EditStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIntroduction(v string) *DescribeInstanceResponseBody {
	s.Introduction = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetLanguageCode(v string) *DescribeInstanceResponseBody {
	s.LanguageCode = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetName(v string) *DescribeInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRobotType(v string) *DescribeInstanceResponseBody {
	s.RobotType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTimeZone(v string) *DescribeInstanceResponseBody {
	s.TimeZone = &v
	return s
}

type DescribeInstanceResponseBodyCategories struct {
	// example:
	//
	// 30000066832
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 杭州市防疫政策
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s DescribeInstanceResponseBodyCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyCategories) SetCategoryId(v int64) *DescribeInstanceResponseBodyCategories {
	s.CategoryId = &v
	return s
}

func (s *DescribeInstanceResponseBodyCategories) SetName(v string) *DescribeInstanceResponseBodyCategories {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyCategories) SetParentCategoryId(v int64) *DescribeInstanceResponseBodyCategories {
	s.ParentCategoryId = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeIntentRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s DescribeIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntentRequest) SetAgentKey(v string) *DescribeIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeIntentRequest) SetInstanceId(v string) *DescribeIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntentRequest) SetIntentId(v int64) *DescribeIntentRequest {
	s.IntentId = &v
	return s
}

type DescribeIntentResponseBody struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123231
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 84243341
	IntentId   *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 123231
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// test
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// a22afaf2adfasf2gr345fga45ada
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlotInfos []*DescribeIntentResponseBodySlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s DescribeIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBody) SetAliasName(v string) *DescribeIntentResponseBody {
	s.AliasName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetCreateTime(v string) *DescribeIntentResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeIntentResponseBody) SetCreateUserId(v string) *DescribeIntentResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetCreateUserName(v string) *DescribeIntentResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetIntentId(v int64) *DescribeIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetIntentName(v string) *DescribeIntentResponseBody {
	s.IntentName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyTime(v string) *DescribeIntentResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyUserId(v string) *DescribeIntentResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyUserName(v string) *DescribeIntentResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetRequestId(v string) *DescribeIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetSlotInfos(v []*DescribeIntentResponseBodySlotInfos) *DescribeIntentResponseBody {
	s.SlotInfos = v
	return s
}

type DescribeIntentResponseBodySlotInfos struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// example:
	//
	// false
	Interactive *bool   `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// aa4d2a343a3ad4afad
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIntentResponseBodySlotInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponseBodySlotInfos) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBodySlotInfos) SetArray(v bool) *DescribeIntentResponseBodySlotInfos {
	s.Array = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetEncrypt(v bool) *DescribeIntentResponseBodySlotInfos {
	s.Encrypt = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetInteractive(v bool) *DescribeIntentResponseBodySlotInfos {
	s.Interactive = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetName(v string) *DescribeIntentResponseBodySlotInfos {
	s.Name = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetSlotId(v string) *DescribeIntentResponseBodySlotInfos {
	s.SlotId = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetValue(v string) *DescribeIntentResponseBodySlotInfos {
	s.Value = &v
	return s
}

type DescribeIntentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIntentResponse) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponse) SetHeaders(v map[string]*string) *DescribeIntentResponse {
	s.Headers = v
	return s
}

func (s *DescribeIntentResponse) SetStatusCode(v int32) *DescribeIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIntentResponse) SetBody(v *DescribeIntentResponseBody) *DescribeIntentResponse {
	s.Body = v
	return s
}

type DescribePerspectiveRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s DescribePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveRequest) SetAgentKey(v string) *DescribePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribePerspectiveRequest) SetPerspectiveId(v string) *DescribePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type DescribePerspectiveResponseBody struct {
	// example:
	//
	// 2021-07-27T07:05:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2021-07-26T07:05:37Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 移动端视角
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FZJBY3raWr
	PerspectiveCode *string `json:"PerspectiveCode,omitempty" xml:"PerspectiveCode,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	// example:
	//
	// F285D735-D580-18A8-B97F-B2E72B00F101
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SelfDefine *bool `json:"SelfDefine,omitempty" xml:"SelfDefine,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePerspectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveResponseBody) SetCreateTime(v string) *DescribePerspectiveResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetModifyTime(v string) *DescribePerspectiveResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetName(v string) *DescribePerspectiveResponseBody {
	s.Name = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetPerspectiveCode(v string) *DescribePerspectiveResponseBody {
	s.PerspectiveCode = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetPerspectiveId(v string) *DescribePerspectiveResponseBody {
	s.PerspectiveId = &v
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

func (s *DescribePerspectiveResponseBody) SetStatus(v int32) *DescribePerspectiveResponseBody {
	s.Status = &v
	return s
}

type DescribePerspectiveResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribePerspectiveResponse) SetStatusCode(v int32) *DescribePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePerspectiveResponse) SetBody(v *DescribePerspectiveResponseBody) *DescribePerspectiveResponse {
	s.Body = v
	return s
}

type FeedbackRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// good
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// example:
	//
	// 这个回答很棒
	FeedbackContent *string `json:"FeedbackContent,omitempty" xml:"FeedbackContent,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5ca40988-4f99-47ad-ac96-9060d0f81db9
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 7c3cec23cc8940bc9db4a318c8f4f0aa
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackRequest) GoString() string {
	return s.String()
}

func (s *FeedbackRequest) SetAgentKey(v string) *FeedbackRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackRequest) SetFeedback(v string) *FeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *FeedbackRequest) SetFeedbackContent(v string) *FeedbackRequest {
	s.FeedbackContent = &v
	return s
}

func (s *FeedbackRequest) SetInstanceId(v string) *FeedbackRequest {
	s.InstanceId = &v
	return s
}

func (s *FeedbackRequest) SetMessageId(v string) *FeedbackRequest {
	s.MessageId = &v
	return s
}

func (s *FeedbackRequest) SetSessionId(v string) *FeedbackRequest {
	s.SessionId = &v
	return s
}

type FeedbackResponseBody struct {
	// example:
	//
	// good
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// example:
	//
	// 5ca40988-4f99-47ad-ac96-9060d0f81db9
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 4e5eea71-f326-450c-8849-49515473ef64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *FeedbackResponseBody) SetFeedback(v string) *FeedbackResponseBody {
	s.Feedback = &v
	return s
}

func (s *FeedbackResponseBody) SetMessageId(v string) *FeedbackResponseBody {
	s.MessageId = &v
	return s
}

func (s *FeedbackResponseBody) SetRequestId(v string) *FeedbackResponseBody {
	s.RequestId = &v
	return s
}

type FeedbackResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *FeedbackResponse) SetStatusCode(v int32) *FeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackResponse) SetBody(v *FeedbackResponseBody) *FeedbackResponse {
	s.Body = v
	return s
}

type GenerateUserAccessTokenRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// a***bcx@gmail.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 1000
	ExpireTime *int32  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExtraInfo  *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8882022040000000171
	ForeignId *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	// This parameter is required.
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// example:
	//
	// 1381111****
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
}

func (s GenerateUserAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUserAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateUserAccessTokenRequest) SetAgentKey(v string) *GenerateUserAccessTokenRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetEmail(v string) *GenerateUserAccessTokenRequest {
	s.Email = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetExpireTime(v int32) *GenerateUserAccessTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetExtraInfo(v string) *GenerateUserAccessTokenRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetForeignId(v string) *GenerateUserAccessTokenRequest {
	s.ForeignId = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetNick(v string) *GenerateUserAccessTokenRequest {
	s.Nick = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetTelephone(v string) *GenerateUserAccessTokenRequest {
	s.Telephone = &v
	return s
}

type GenerateUserAccessTokenResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// DDEXEDAFWAGASDFWAEFFAWEFAWFWEAFWAFWAEF
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Parameter.Invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E6988CE6-41CF-1103-9BEC-2B20D26C0B51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUserAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUserAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUserAccessTokenResponseBody) SetCode(v string) *GenerateUserAccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) SetData(v string) *GenerateUserAccessTokenResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) SetMessage(v string) *GenerateUserAccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) SetRequestId(v string) *GenerateUserAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) SetSuccess(v bool) *GenerateUserAccessTokenResponseBody {
	s.Success = &v
	return s
}

type GenerateUserAccessTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUserAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUserAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUserAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateUserAccessTokenResponse) SetHeaders(v map[string]*string) *GenerateUserAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateUserAccessTokenResponse) SetStatusCode(v int32) *GenerateUserAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUserAccessTokenResponse) SetBody(v *GenerateUserAccessTokenResponseBody) *GenerateUserAccessTokenResponse {
	s.Body = v
	return s
}

type GetAgentInfoRequest struct {
	// example:
	//
	// beebot_bot_public_cn-ca36x8v3n1x
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAgentInfoRequest) SetInstanceId(v string) *GetAgentInfoRequest {
	s.InstanceId = &v
	return s
}

type GetAgentInfoResponseBody struct {
	Data *GetAgentInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter.Invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FC384CE1-8D42-1900-84E1-F33F990F2B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentInfoResponseBody) SetData(v *GetAgentInfoResponseBodyData) *GetAgentInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentInfoResponseBody) SetMessage(v string) *GetAgentInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentInfoResponseBody) SetRequestId(v string) *GetAgentInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentInfoResponseBody) SetSuccess(v bool) *GetAgentInfoResponseBody {
	s.Success = &v
	return s
}

type GetAgentInfoResponseBodyData struct {
	// example:
	//
	// 4e7400028e6f4a7393ed3acf6a7b8927_p_beebot_public
	AgentKey  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s GetAgentInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentInfoResponseBodyData) SetAgentKey(v string) *GetAgentInfoResponseBodyData {
	s.AgentKey = &v
	return s
}

func (s *GetAgentInfoResponseBodyData) SetAgentName(v string) *GetAgentInfoResponseBodyData {
	s.AgentName = &v
	return s
}

type GetAgentInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAgentInfoResponse) SetHeaders(v map[string]*string) *GetAgentInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAgentInfoResponse) SetStatusCode(v int32) *GetAgentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentInfoResponse) SetBody(v *GetAgentInfoResponseBody) *GetAgentInfoResponse {
	s.Body = v
	return s
}

type GetAsyncResultRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAsyncResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncResultRequest) SetAgentKey(v string) *GetAsyncResultRequest {
	s.AgentKey = &v
	return s
}

func (s *GetAsyncResultRequest) SetTaskId(v string) *GetAsyncResultRequest {
	s.TaskId = &v
	return s
}

type GetAsyncResultResponseBody struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// F79E7305-5314-5069-A701-9591AD051902
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponseBody) SetData(v string) *GetAsyncResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetRequestId(v string) *GetAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetStatus(v string) *GetAsyncResultResponseBody {
	s.Status = &v
	return s
}

type GetAsyncResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponse) SetHeaders(v map[string]*string) *GetAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncResultResponse) SetStatusCode(v int32) *GetAsyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncResultResponse) SetBody(v *GetAsyncResultResponseBody) *GetAsyncResultResponse {
	s.Body = v
	return s
}

type GetBotSessionDataRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20240605
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-7QuUfaqMQe
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20240505
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetBotSessionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBotSessionDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataRequest) SetAgentKey(v string) *GetBotSessionDataRequest {
	s.AgentKey = &v
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

func (s *GetBotSessionDataRequest) SetStartTime(v string) *GetBotSessionDataRequest {
	s.StartTime = &v
	return s
}

type GetBotSessionDataResponseBody struct {
	// example:
	//
	// 116
	CostTime *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas    []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// example:
	//
	// 15CD94CC-CBEB-4189-806C-A132D1F45D51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetBotSessionDataResponseBody) SetDatas(v []map[string]interface{}) *GetBotSessionDataResponseBody {
	s.Datas = v
	return s
}

func (s *GetBotSessionDataResponseBody) SetRequestId(v string) *GetBotSessionDataResponseBody {
	s.RequestId = &v
	return s
}

type GetBotSessionDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBotSessionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetBotSessionDataResponse) SetStatusCode(v int32) *GetBotSessionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBotSessionDataResponse) SetBody(v *GetBotSessionDataResponseBody) *GetBotSessionDataResponse {
	s.Body = v
	return s
}

type GetInstancePublishTaskStateRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstancePublishTaskStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstancePublishTaskStateRequest) GoString() string {
	return s.String()
}

func (s *GetInstancePublishTaskStateRequest) SetAgentKey(v string) *GetInstancePublishTaskStateRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInstancePublishTaskStateRequest) SetId(v int64) *GetInstancePublishTaskStateRequest {
	s.Id = &v
	return s
}

func (s *GetInstancePublishTaskStateRequest) SetInstanceId(v string) *GetInstancePublishTaskStateRequest {
	s.InstanceId = &v
	return s
}

type GetInstancePublishTaskStateResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string                `json:"Error,omitempty" xml:"Error,omitempty"`
	Errors     map[string]interface{} `json:"Errors,omitempty" xml:"Errors,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status   *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Warnings map[string]interface{} `json:"Warnings,omitempty" xml:"Warnings,omitempty"`
}

func (s GetInstancePublishTaskStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstancePublishTaskStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancePublishTaskStateResponseBody) SetBizTypeList(v []*string) *GetInstancePublishTaskStateResponseBody {
	s.BizTypeList = v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetCreateTime(v string) *GetInstancePublishTaskStateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetError(v string) *GetInstancePublishTaskStateResponseBody {
	s.Error = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetErrors(v map[string]interface{}) *GetInstancePublishTaskStateResponseBody {
	s.Errors = v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetId(v int64) *GetInstancePublishTaskStateResponseBody {
	s.Id = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetModifyTime(v string) *GetInstancePublishTaskStateResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetRequestId(v string) *GetInstancePublishTaskStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetResponse(v string) *GetInstancePublishTaskStateResponseBody {
	s.Response = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetStatus(v string) *GetInstancePublishTaskStateResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetWarnings(v map[string]interface{}) *GetInstancePublishTaskStateResponseBody {
	s.Warnings = v
	return s
}

type GetInstancePublishTaskStateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstancePublishTaskStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstancePublishTaskStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstancePublishTaskStateResponse) GoString() string {
	return s.String()
}

func (s *GetInstancePublishTaskStateResponse) SetHeaders(v map[string]*string) *GetInstancePublishTaskStateResponse {
	s.Headers = v
	return s
}

func (s *GetInstancePublishTaskStateResponse) SetStatusCode(v int32) *GetInstancePublishTaskStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancePublishTaskStateResponse) SetBody(v *GetInstancePublishTaskStateResponseBody) *GetInstancePublishTaskStateResponse {
	s.Body = v
	return s
}

type GetPublishTaskStateRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetPublishTaskStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublishTaskStateRequest) GoString() string {
	return s.String()
}

func (s *GetPublishTaskStateRequest) SetAgentKey(v string) *GetPublishTaskStateRequest {
	s.AgentKey = &v
	return s
}

func (s *GetPublishTaskStateRequest) SetId(v int64) *GetPublishTaskStateRequest {
	s.Id = &v
	return s
}

type GetPublishTaskStateResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string                `json:"Error,omitempty" xml:"Error,omitempty"`
	Errors     map[string]interface{} `json:"Errors,omitempty" xml:"Errors,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status   *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Warnings map[string]interface{} `json:"Warnings,omitempty" xml:"Warnings,omitempty"`
}

func (s GetPublishTaskStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublishTaskStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublishTaskStateResponseBody) SetBizTypeList(v []*string) *GetPublishTaskStateResponseBody {
	s.BizTypeList = v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetCreateTime(v string) *GetPublishTaskStateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetError(v string) *GetPublishTaskStateResponseBody {
	s.Error = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetErrors(v map[string]interface{}) *GetPublishTaskStateResponseBody {
	s.Errors = v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetId(v int64) *GetPublishTaskStateResponseBody {
	s.Id = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetModifyTime(v string) *GetPublishTaskStateResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetRequestId(v string) *GetPublishTaskStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetResponse(v string) *GetPublishTaskStateResponseBody {
	s.Response = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetStatus(v string) *GetPublishTaskStateResponseBody {
	s.Status = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetWarnings(v map[string]interface{}) *GetPublishTaskStateResponseBody {
	s.Warnings = v
	return s
}

type GetPublishTaskStateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublishTaskStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublishTaskStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublishTaskStateResponse) GoString() string {
	return s.String()
}

func (s *GetPublishTaskStateResponse) SetHeaders(v map[string]*string) *GetPublishTaskStateResponse {
	s.Headers = v
	return s
}

func (s *GetPublishTaskStateResponse) SetStatusCode(v int32) *GetPublishTaskStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublishTaskStateResponse) SetBody(v *GetPublishTaskStateResponseBody) *GetPublishTaskStateResponse {
	s.Body = v
	return s
}

type InitIMConnectRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8XNBzDucJv
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// QUM4SndaY3VPMjhkQldDZUNOR0ZaTmZ5R3NBY0FKWHJ4OGc4dERZbEJzcjNIKzFiS1RyTjhXRUpBYmVpQlpsakprNDRFVkdxcy9HWVk2RXZvalU3bHhxRkJlc1NBUXZwdHFKOTE2UTNwamQ4b1U4N3dEbmhyRjc4R2hOQStvMnMrYkV2dlVpSHNvWC96SEVNZWRqMjBuMXdjNklpamJzaDNWYllnUldDZGhJPQ==
	UserAccessToken *string `json:"UserAccessToken,omitempty" xml:"UserAccessToken,omitempty"`
}

func (s InitIMConnectRequest) String() string {
	return tea.Prettify(s)
}

func (s InitIMConnectRequest) GoString() string {
	return s.String()
}

func (s *InitIMConnectRequest) SetAgentKey(v string) *InitIMConnectRequest {
	s.AgentKey = &v
	return s
}

func (s *InitIMConnectRequest) SetFrom(v string) *InitIMConnectRequest {
	s.From = &v
	return s
}

func (s *InitIMConnectRequest) SetUserAccessToken(v string) *InitIMConnectRequest {
	s.UserAccessToken = &v
	return s
}

type InitIMConnectResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {     "imDomain": "im.alimebot.com",     "appKey": "WDg2VfNv",     "token": "QUM4SndaY3VPMjhkQldDZUNOR0ZaTmZ5R3NBY0FKWHJ4OGc4dERZbEJzcjNIKzFiS1RyTjhXRUpBYmVpQlpsakprNDRFVkdxcy9HWVk2RXZvalU3bHhxRkJlc1NBUXZwdHFKOTE2UTNwamQ4b1U4N3dEbmhyRjc4R2hOQStvMnMrYkV2dlVpSHNvWC96SEVNZWRqMjBuMXdjNklpamJzaDNWYllnUldDZGhJPQ=="   }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Parameter.Invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E6988CE6-41CF-1103-9BEC-2B20D26C0B52
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitIMConnectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitIMConnectResponseBody) GoString() string {
	return s.String()
}

func (s *InitIMConnectResponseBody) SetCode(v string) *InitIMConnectResponseBody {
	s.Code = &v
	return s
}

func (s *InitIMConnectResponseBody) SetData(v string) *InitIMConnectResponseBody {
	s.Data = &v
	return s
}

func (s *InitIMConnectResponseBody) SetMessage(v string) *InitIMConnectResponseBody {
	s.Message = &v
	return s
}

func (s *InitIMConnectResponseBody) SetRequestId(v string) *InitIMConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitIMConnectResponseBody) SetSuccess(v bool) *InitIMConnectResponseBody {
	s.Success = &v
	return s
}

type InitIMConnectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitIMConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitIMConnectResponse) String() string {
	return tea.Prettify(s)
}

func (s InitIMConnectResponse) GoString() string {
	return s.String()
}

func (s *InitIMConnectResponse) SetHeaders(v map[string]*string) *InitIMConnectResponse {
	s.Headers = v
	return s
}

func (s *InitIMConnectResponse) SetStatusCode(v int32) *InitIMConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *InitIMConnectResponse) SetBody(v *InitIMConnectResponseBody) *InitIMConnectResponse {
	s.Body = v
	return s
}

type LinkInstanceCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// [\\"30000065789\\"]
	CategoryIds *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s LinkInstanceCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s LinkInstanceCategoryRequest) GoString() string {
	return s.String()
}

func (s *LinkInstanceCategoryRequest) SetAgentKey(v string) *LinkInstanceCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *LinkInstanceCategoryRequest) SetCategoryIds(v string) *LinkInstanceCategoryRequest {
	s.CategoryIds = &v
	return s
}

func (s *LinkInstanceCategoryRequest) SetInstanceId(v string) *LinkInstanceCategoryRequest {
	s.InstanceId = &v
	return s
}

type LinkInstanceCategoryResponseBody struct {
	// example:
	//
	// D8C96601-E645-1BD7-99F3-04EADAB84E29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LinkInstanceCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LinkInstanceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *LinkInstanceCategoryResponseBody) SetRequestId(v string) *LinkInstanceCategoryResponseBody {
	s.RequestId = &v
	return s
}

type LinkInstanceCategoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LinkInstanceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LinkInstanceCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s LinkInstanceCategoryResponse) GoString() string {
	return s.String()
}

func (s *LinkInstanceCategoryResponse) SetHeaders(v map[string]*string) *LinkInstanceCategoryResponse {
	s.Headers = v
	return s
}

func (s *LinkInstanceCategoryResponse) SetStatusCode(v int32) *LinkInstanceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *LinkInstanceCategoryResponse) SetBody(v *LinkInstanceCategoryResponseBody) *LinkInstanceCategoryResponse {
	s.Body = v
	return s
}

type ListAgentRequest struct {
	AgentName  *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	GoodsCodes *string `json:"GoodsCodes,omitempty" xml:"GoodsCodes,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentRequest) GoString() string {
	return s.String()
}

func (s *ListAgentRequest) SetAgentName(v string) *ListAgentRequest {
	s.AgentName = &v
	return s
}

func (s *ListAgentRequest) SetGoodsCodes(v string) *ListAgentRequest {
	s.GoodsCodes = &v
	return s
}

func (s *ListAgentRequest) SetPageNumber(v int32) *ListAgentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRequest) SetPageSize(v int32) *ListAgentRequest {
	s.PageSize = &v
	return s
}

type ListAgentResponseBody struct {
	Data []*ListAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F55D90C1-31BE-4B2A-AA3F-25EFC36F9419
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentResponseBody) SetData(v []*ListAgentResponseBodyData) *ListAgentResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentResponseBody) SetPageNumber(v int32) *ListAgentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAgentResponseBody) SetPageSize(v int32) *ListAgentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAgentResponseBody) SetRequestId(v string) *ListAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentResponseBody) SetTotalCount(v int32) *ListAgentResponseBody {
	s.TotalCount = &v
	return s
}

type ListAgentResponseBodyData struct {
	// example:
	//
	// 881
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 4e7400028e6f4a7393ed3acf6a7b8927_p_beebot_public
	AgentKey      *string                `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentName     *string                `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	InstanceInfos map[string]interface{} `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty"`
}

func (s ListAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentResponseBodyData) SetAgentId(v int64) *ListAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListAgentResponseBodyData) SetAgentKey(v string) *ListAgentResponseBodyData {
	s.AgentKey = &v
	return s
}

func (s *ListAgentResponseBodyData) SetAgentName(v string) *ListAgentResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *ListAgentResponseBodyData) SetInstanceInfos(v map[string]interface{}) *ListAgentResponseBodyData {
	s.InstanceInfos = v
	return s
}

type ListAgentResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentResponse) GoString() string {
	return s.String()
}

func (s *ListAgentResponse) SetHeaders(v map[string]*string) *ListAgentResponse {
	s.Headers = v
	return s
}

func (s *ListAgentResponse) SetStatusCode(v int32) *ListAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentResponse) SetBody(v *ListAgentResponseBody) *ListAgentResponse {
	s.Body = v
	return s
}

type ListCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	KnowledgeType *int32  `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s ListCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListCategoryRequest) SetAgentKey(v string) *ListCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *ListCategoryRequest) SetKnowledgeType(v int32) *ListCategoryRequest {
	s.KnowledgeType = &v
	return s
}

func (s *ListCategoryRequest) SetParentCategoryId(v int64) *ListCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

type ListCategoryResponseBody struct {
	Categories []*ListCategoryResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// 9C5F8186-2D22-433E-9545-606D344F30B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBody) SetCategories(v []*ListCategoryResponseBodyCategories) *ListCategoryResponseBody {
	s.Categories = v
	return s
}

func (s *ListCategoryResponseBody) SetRequestId(v string) *ListCategoryResponseBody {
	s.RequestId = &v
	return s
}

type ListCategoryResponseBodyCategories struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 231001028593
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCategoryResponseBodyCategories) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBodyCategories) SetBizCode(v string) *ListCategoryResponseBodyCategories {
	s.BizCode = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) SetCategoryId(v int64) *ListCategoryResponseBodyCategories {
	s.CategoryId = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) SetName(v string) *ListCategoryResponseBodyCategories {
	s.Name = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) SetParentCategoryId(v int64) *ListCategoryResponseBodyCategories {
	s.ParentCategoryId = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) SetStatus(v int32) *ListCategoryResponseBodyCategories {
	s.Status = &v
	return s
}

type ListCategoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListCategoryResponse) SetHeaders(v map[string]*string) *ListCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListCategoryResponse) SetStatusCode(v int32) *ListCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCategoryResponse) SetBody(v *ListCategoryResponseBody) *ListCategoryResponse {
	s.Body = v
	return s
}

type ListConnQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s ListConnQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnQuestionRequest) GoString() string {
	return s.String()
}

func (s *ListConnQuestionRequest) SetAgentKey(v string) *ListConnQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *ListConnQuestionRequest) SetKnowledgeId(v int64) *ListConnQuestionRequest {
	s.KnowledgeId = &v
	return s
}

type ListConnQuestionResponseBody struct {
	Outlines []*ListConnQuestionResponseBodyOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 92B81548-42B9-4B34-924B-4E778AEB412B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConnQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnQuestionResponseBody) SetOutlines(v []*ListConnQuestionResponseBodyOutlines) *ListConnQuestionResponseBody {
	s.Outlines = v
	return s
}

func (s *ListConnQuestionResponseBody) SetRequestId(v string) *ListConnQuestionResponseBody {
	s.RequestId = &v
	return s
}

type ListConnQuestionResponseBodyOutlines struct {
	// example:
	//
	// 30001979424
	ConnQuestionId *int64 `json:"ConnQuestionId,omitempty" xml:"ConnQuestionId,omitempty"`
	// example:
	//
	// 2022-02-25T02:47:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-26T10:18:15Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 797
	OutlineId *int64  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListConnQuestionResponseBodyOutlines) String() string {
	return tea.Prettify(s)
}

func (s ListConnQuestionResponseBodyOutlines) GoString() string {
	return s.String()
}

func (s *ListConnQuestionResponseBodyOutlines) SetConnQuestionId(v int64) *ListConnQuestionResponseBodyOutlines {
	s.ConnQuestionId = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) SetCreateTime(v string) *ListConnQuestionResponseBodyOutlines {
	s.CreateTime = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) SetModifyTime(v string) *ListConnQuestionResponseBodyOutlines {
	s.ModifyTime = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) SetOutlineId(v int64) *ListConnQuestionResponseBodyOutlines {
	s.OutlineId = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) SetTitle(v string) *ListConnQuestionResponseBodyOutlines {
	s.Title = &v
	return s
}

type ListConnQuestionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnQuestionResponse) GoString() string {
	return s.String()
}

func (s *ListConnQuestionResponse) SetHeaders(v map[string]*string) *ListConnQuestionResponse {
	s.Headers = v
	return s
}

func (s *ListConnQuestionResponse) SetStatusCode(v int32) *ListConnQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnQuestionResponse) SetBody(v *ListConnQuestionResponseBody) *ListConnQuestionResponse {
	s.Body = v
	return s
}

type ListDSEntityRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// system
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 实体
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDSEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDSEntityRequest) GoString() string {
	return s.String()
}

func (s *ListDSEntityRequest) SetAgentKey(v string) *ListDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDSEntityRequest) SetEntityType(v string) *ListDSEntityRequest {
	s.EntityType = &v
	return s
}

func (s *ListDSEntityRequest) SetInstanceId(v string) *ListDSEntityRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDSEntityRequest) SetKeyword(v string) *ListDSEntityRequest {
	s.Keyword = &v
	return s
}

func (s *ListDSEntityRequest) SetPageNumber(v int32) *ListDSEntityRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDSEntityRequest) SetPageSize(v int32) *ListDSEntityRequest {
	s.PageSize = &v
	return s
}

type ListDSEntityResponseBody struct {
	Entities []*ListDSEntityResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ga4h345defgwet2sdf223
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDSEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *ListDSEntityResponseBody) SetEntities(v []*ListDSEntityResponseBodyEntities) *ListDSEntityResponseBody {
	s.Entities = v
	return s
}

func (s *ListDSEntityResponseBody) SetPageNumber(v int32) *ListDSEntityResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDSEntityResponseBody) SetPageSize(v int32) *ListDSEntityResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDSEntityResponseBody) SetRequestId(v string) *ListDSEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDSEntityResponseBody) SetTotalCount(v int32) *ListDSEntityResponseBody {
	s.TotalCount = &v
	return s
}

type ListDSEntityResponseBodyEntities struct {
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123231
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 234564567445
	EntityId   *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// synonyms
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 123231
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// test
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	SysEntityCode  *string `json:"SysEntityCode,omitempty" xml:"SysEntityCode,omitempty"`
}

func (s ListDSEntityResponseBodyEntities) String() string {
	return tea.Prettify(s)
}

func (s ListDSEntityResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *ListDSEntityResponseBodyEntities) SetCreateTime(v string) *ListDSEntityResponseBodyEntities {
	s.CreateTime = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetCreateUserId(v string) *ListDSEntityResponseBodyEntities {
	s.CreateUserId = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetCreateUserName(v string) *ListDSEntityResponseBodyEntities {
	s.CreateUserName = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetEntityId(v int64) *ListDSEntityResponseBodyEntities {
	s.EntityId = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetEntityName(v string) *ListDSEntityResponseBodyEntities {
	s.EntityName = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetEntityType(v string) *ListDSEntityResponseBodyEntities {
	s.EntityType = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetModifyTime(v string) *ListDSEntityResponseBodyEntities {
	s.ModifyTime = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetModifyUserId(v string) *ListDSEntityResponseBodyEntities {
	s.ModifyUserId = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetModifyUserName(v string) *ListDSEntityResponseBodyEntities {
	s.ModifyUserName = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetSysEntityCode(v string) *ListDSEntityResponseBodyEntities {
	s.SysEntityCode = &v
	return s
}

type ListDSEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDSEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDSEntityResponse) GoString() string {
	return s.String()
}

func (s *ListDSEntityResponse) SetHeaders(v map[string]*string) *ListDSEntityResponse {
	s.Headers = v
	return s
}

func (s *ListDSEntityResponse) SetStatusCode(v int32) *ListDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDSEntityResponse) SetBody(v *ListDSEntityResponseBody) *ListDSEntityResponse {
	s.Body = v
	return s
}

type ListDSEntityValueRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 234
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDSEntityValueRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDSEntityValueRequest) GoString() string {
	return s.String()
}

func (s *ListDSEntityValueRequest) SetAgentKey(v string) *ListDSEntityValueRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDSEntityValueRequest) SetEntityId(v int64) *ListDSEntityValueRequest {
	s.EntityId = &v
	return s
}

func (s *ListDSEntityValueRequest) SetEntityValueId(v int64) *ListDSEntityValueRequest {
	s.EntityValueId = &v
	return s
}

func (s *ListDSEntityValueRequest) SetInstanceId(v string) *ListDSEntityValueRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDSEntityValueRequest) SetKeyword(v string) *ListDSEntityValueRequest {
	s.Keyword = &v
	return s
}

func (s *ListDSEntityValueRequest) SetPageNumber(v int32) *ListDSEntityValueRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDSEntityValueRequest) SetPageSize(v int32) *ListDSEntityValueRequest {
	s.PageSize = &v
	return s
}

type ListDSEntityValueResponseBody struct {
	EntityValues []*ListDSEntityValueResponseBodyEntityValues `json:"EntityValues,omitempty" xml:"EntityValues,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// n3fg34gbfj8adf2gj923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDSEntityValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDSEntityValueResponseBody) GoString() string {
	return s.String()
}

func (s *ListDSEntityValueResponseBody) SetEntityValues(v []*ListDSEntityValueResponseBodyEntityValues) *ListDSEntityValueResponseBody {
	s.EntityValues = v
	return s
}

func (s *ListDSEntityValueResponseBody) SetPageNumber(v int32) *ListDSEntityValueResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDSEntityValueResponseBody) SetPageSize(v int32) *ListDSEntityValueResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDSEntityValueResponseBody) SetRequestId(v string) *ListDSEntityValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDSEntityValueResponseBody) SetTotalCount(v int32) *ListDSEntityValueResponseBody {
	s.TotalCount = &v
	return s
}

type ListDSEntityValueResponseBodyEntityValues struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 34313785463
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 3467858234534534532
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s ListDSEntityValueResponseBodyEntityValues) String() string {
	return tea.Prettify(s)
}

func (s ListDSEntityValueResponseBodyEntityValues) GoString() string {
	return s.String()
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetContent(v string) *ListDSEntityValueResponseBodyEntityValues {
	s.Content = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetCreateTime(v string) *ListDSEntityValueResponseBodyEntityValues {
	s.CreateTime = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetEntityId(v int64) *ListDSEntityValueResponseBodyEntityValues {
	s.EntityId = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetEntityValueId(v int64) *ListDSEntityValueResponseBodyEntityValues {
	s.EntityValueId = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetModifyTime(v string) *ListDSEntityValueResponseBodyEntityValues {
	s.ModifyTime = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetSynonyms(v []*string) *ListDSEntityValueResponseBodyEntityValues {
	s.Synonyms = v
	return s
}

type ListDSEntityValueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDSEntityValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDSEntityValueResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDSEntityValueResponse) GoString() string {
	return s.String()
}

func (s *ListDSEntityValueResponse) SetHeaders(v map[string]*string) *ListDSEntityValueResponse {
	s.Headers = v
	return s
}

func (s *ListDSEntityValueResponse) SetStatusCode(v int32) *ListDSEntityValueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDSEntityValueResponse) SetBody(v *ListDSEntityValueResponseBody) *ListDSEntityValueResponse {
	s.Body = v
	return s
}

type ListInstanceRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// scenario_im
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// scenario_im
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) SetAgentKey(v string) *ListInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInstanceRequest) SetName(v string) *ListInstanceRequest {
	s.Name = &v
	return s
}

func (s *ListInstanceRequest) SetPageNumber(v int64) *ListInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceRequest) SetPageSize(v int64) *ListInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRequest) SetRobotType(v string) *ListInstanceRequest {
	s.RobotType = &v
	return s
}

type ListInstanceResponseBody struct {
	Instances []*ListInstanceResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 92B81548-42B9-4B34-924B-4E778AEB412B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) SetInstances(v []*ListInstanceResponseBodyInstances) *ListInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstanceResponseBody) SetPageNumber(v int64) *ListInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceResponseBody) SetPageSize(v int64) *ListInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetTotalCount(v int64) *ListInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstanceResponseBodyInstances struct {
	// example:
	//
	// /alimefe/meebot/robot/0.0.5/img/xxx-90-97.png
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 用于C端问答的机器人
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// example:
	//
	// 智能客服-小C
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// scenario_im
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s ListInstanceResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyInstances) SetAvatar(v string) *ListInstanceResponseBodyInstances {
	s.Avatar = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetCreateTime(v string) *ListInstanceResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceId(v string) *ListInstanceResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetIntroduction(v string) *ListInstanceResponseBodyInstances {
	s.Introduction = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetLanguageCode(v string) *ListInstanceResponseBodyInstances {
	s.LanguageCode = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetName(v string) *ListInstanceResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetRobotType(v string) *ListInstanceResponseBodyInstances {
	s.RobotType = &v
	return s
}

type ListInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetStatusCode(v int32) *ListInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

type ListIntentRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 查天气
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntentRequest) GoString() string {
	return s.String()
}

func (s *ListIntentRequest) SetAgentKey(v string) *ListIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *ListIntentRequest) SetInstanceId(v string) *ListIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntentRequest) SetIntentName(v string) *ListIntentRequest {
	s.IntentName = &v
	return s
}

func (s *ListIntentRequest) SetPageNumber(v int32) *ListIntentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIntentRequest) SetPageSize(v int32) *ListIntentRequest {
	s.PageSize = &v
	return s
}

type ListIntentResponseBody struct {
	Intents []*ListIntentResponseBodyIntents `json:"Intents,omitempty" xml:"Intents,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 23dsfa34r2s2s2sd12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntentResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntentResponseBody) SetIntents(v []*ListIntentResponseBodyIntents) *ListIntentResponseBody {
	s.Intents = v
	return s
}

func (s *ListIntentResponseBody) SetPageNumber(v int32) *ListIntentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIntentResponseBody) SetPageSize(v int32) *ListIntentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIntentResponseBody) SetRequestId(v string) *ListIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntentResponseBody) SetTotalCount(v int32) *ListIntentResponseBody {
	s.TotalCount = &v
	return s
}

type ListIntentResponseBodyIntents struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123231
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 234234234234
	IntentId   *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 123231
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// test
	ModifyUserName *string                                   `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	SlotInfos      []*ListIntentResponseBodyIntentsSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s ListIntentResponseBodyIntents) String() string {
	return tea.Prettify(s)
}

func (s ListIntentResponseBodyIntents) GoString() string {
	return s.String()
}

func (s *ListIntentResponseBodyIntents) SetAliasName(v string) *ListIntentResponseBodyIntents {
	s.AliasName = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetCreateTime(v string) *ListIntentResponseBodyIntents {
	s.CreateTime = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetCreateUserId(v string) *ListIntentResponseBodyIntents {
	s.CreateUserId = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetCreateUserName(v string) *ListIntentResponseBodyIntents {
	s.CreateUserName = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetIntentId(v int64) *ListIntentResponseBodyIntents {
	s.IntentId = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetIntentName(v string) *ListIntentResponseBodyIntents {
	s.IntentName = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetModifyTime(v string) *ListIntentResponseBodyIntents {
	s.ModifyTime = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetModifyUserId(v string) *ListIntentResponseBodyIntents {
	s.ModifyUserId = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetModifyUserName(v string) *ListIntentResponseBodyIntents {
	s.ModifyUserName = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetSlotInfos(v []*ListIntentResponseBodyIntentsSlotInfos) *ListIntentResponseBodyIntents {
	s.SlotInfos = v
	return s
}

type ListIntentResponseBodyIntentsSlotInfos struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// example:
	//
	// false
	Interactive *bool   `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12134223
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIntentResponseBodyIntentsSlotInfos) String() string {
	return tea.Prettify(s)
}

func (s ListIntentResponseBodyIntentsSlotInfos) GoString() string {
	return s.String()
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetArray(v bool) *ListIntentResponseBodyIntentsSlotInfos {
	s.Array = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetEncrypt(v bool) *ListIntentResponseBodyIntentsSlotInfos {
	s.Encrypt = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetInteractive(v bool) *ListIntentResponseBodyIntentsSlotInfos {
	s.Interactive = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetName(v string) *ListIntentResponseBodyIntentsSlotInfos {
	s.Name = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetSlotId(v string) *ListIntentResponseBodyIntentsSlotInfos {
	s.SlotId = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetValue(v string) *ListIntentResponseBodyIntentsSlotInfos {
	s.Value = &v
	return s
}

type ListIntentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntentResponse) GoString() string {
	return s.String()
}

func (s *ListIntentResponse) SetHeaders(v map[string]*string) *ListIntentResponse {
	s.Headers = v
	return s
}

func (s *ListIntentResponse) SetStatusCode(v int32) *ListIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntentResponse) SetBody(v *ListIntentResponseBody) *ListIntentResponse {
	s.Body = v
	return s
}

type ListLgfRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// .{0,5}北京天气
	LgfText *string `json:"LgfText,omitempty" xml:"LgfText,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLgfRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLgfRequest) GoString() string {
	return s.String()
}

func (s *ListLgfRequest) SetAgentKey(v string) *ListLgfRequest {
	s.AgentKey = &v
	return s
}

func (s *ListLgfRequest) SetInstanceId(v string) *ListLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLgfRequest) SetIntentId(v int64) *ListLgfRequest {
	s.IntentId = &v
	return s
}

func (s *ListLgfRequest) SetLgfText(v string) *ListLgfRequest {
	s.LgfText = &v
	return s
}

func (s *ListLgfRequest) SetPageNumber(v int32) *ListLgfRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLgfRequest) SetPageSize(v int32) *ListLgfRequest {
	s.PageSize = &v
	return s
}

type ListLgfResponseBody struct {
	Lgfs []*ListLgfResponseBodyLgfs `json:"Lgfs,omitempty" xml:"Lgfs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 34fg57h2gh5783
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLgfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLgfResponseBody) GoString() string {
	return s.String()
}

func (s *ListLgfResponseBody) SetLgfs(v []*ListLgfResponseBodyLgfs) *ListLgfResponseBody {
	s.Lgfs = v
	return s
}

func (s *ListLgfResponseBody) SetPageNumber(v int32) *ListLgfResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLgfResponseBody) SetPageSize(v int32) *ListLgfResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLgfResponseBody) SetRequestId(v string) *ListLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLgfResponseBody) SetTotalCount(v int32) *ListLgfResponseBody {
	s.TotalCount = &v
	return s
}

type ListLgfResponseBodyLgfs struct {
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 256756734345
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// LGF ID
	//
	// example:
	//
	// 123
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RuleText   *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s ListLgfResponseBodyLgfs) String() string {
	return tea.Prettify(s)
}

func (s ListLgfResponseBodyLgfs) GoString() string {
	return s.String()
}

func (s *ListLgfResponseBodyLgfs) SetCreateTime(v string) *ListLgfResponseBodyLgfs {
	s.CreateTime = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) SetIntentId(v int64) *ListLgfResponseBodyLgfs {
	s.IntentId = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) SetLgfId(v int64) *ListLgfResponseBodyLgfs {
	s.LgfId = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) SetModifyTime(v string) *ListLgfResponseBodyLgfs {
	s.ModifyTime = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) SetRuleText(v string) *ListLgfResponseBodyLgfs {
	s.RuleText = &v
	return s
}

type ListLgfResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLgfResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLgfResponse) GoString() string {
	return s.String()
}

func (s *ListLgfResponse) SetHeaders(v map[string]*string) *ListLgfResponse {
	s.Headers = v
	return s
}

func (s *ListLgfResponse) SetStatusCode(v int32) *ListLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLgfResponse) SetBody(v *ListLgfResponseBody) *ListLgfResponse {
	s.Body = v
	return s
}

type ListSaasInfoRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// DS,FAQ
	SaasGroupCodes *string `json:"SaasGroupCodes,omitempty" xml:"SaasGroupCodes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// userTest
	SaasName *string `json:"SaasName,omitempty" xml:"SaasName,omitempty"`
}

func (s ListSaasInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSaasInfoRequest) GoString() string {
	return s.String()
}

func (s *ListSaasInfoRequest) SetAgentKey(v string) *ListSaasInfoRequest {
	s.AgentKey = &v
	return s
}

func (s *ListSaasInfoRequest) SetSaasGroupCodes(v string) *ListSaasInfoRequest {
	s.SaasGroupCodes = &v
	return s
}

func (s *ListSaasInfoRequest) SetSaasName(v string) *ListSaasInfoRequest {
	s.SaasName = &v
	return s
}

type ListSaasInfoResponseBody struct {
	Data []*ListSaasInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A629A28F-F25E-5572-A679-FA46FB0151D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 06614fdb-c72f-436e-8003-dfe8a2854a15
	SaasToken *string `json:"SaasToken,omitempty" xml:"SaasToken,omitempty"`
}

func (s ListSaasInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSaasInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListSaasInfoResponseBody) SetData(v []*ListSaasInfoResponseBodyData) *ListSaasInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListSaasInfoResponseBody) SetRequestId(v string) *ListSaasInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSaasInfoResponseBody) SetSaasToken(v string) *ListSaasInfoResponseBody {
	s.SaasToken = &v
	return s
}

type ListSaasInfoResponseBodyData struct {
	// example:
	//
	// GLOBAL_SERVICE
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// GLOBAL SERVICE
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://pre-alime4service.console.aliyun.com/?productCode=p_beebot_public&switchAgent=1204001&saasCode=Robot&saasToken=06614fdb-c72f-436e-8003-dfe8a2854a15&saasName=123#/robot
	ServiceUrl *string `json:"ServiceUrl,omitempty" xml:"ServiceUrl,omitempty"`
	// example:
	//
	// https://alime.console.aliyun.com/?productCode=p_beebot_public&switchAgent=1204001&saasCode=Robot&saasToken=06614fdb-c72f-436e-8003-dfe8a2854a15&saasName=123#/robot
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSaasInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSaasInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSaasInfoResponseBodyData) SetCode(v string) *ListSaasInfoResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) SetEnName(v string) *ListSaasInfoResponseBodyData {
	s.EnName = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) SetName(v string) *ListSaasInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) SetServiceUrl(v string) *ListSaasInfoResponseBodyData {
	s.ServiceUrl = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) SetUrl(v string) *ListSaasInfoResponseBodyData {
	s.Url = &v
	return s
}

type ListSaasInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSaasInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSaasInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSaasInfoResponse) GoString() string {
	return s.String()
}

func (s *ListSaasInfoResponse) SetHeaders(v map[string]*string) *ListSaasInfoResponse {
	s.Headers = v
	return s
}

func (s *ListSaasInfoResponse) SetStatusCode(v int32) *ListSaasInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSaasInfoResponse) SetBody(v *ListSaasInfoResponseBody) *ListSaasInfoResponse {
	s.Body = v
	return s
}

type ListSaasPermissionGroupInfosRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ListSaasPermissionGroupInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSaasPermissionGroupInfosRequest) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosRequest) SetAgentKey(v string) *ListSaasPermissionGroupInfosRequest {
	s.AgentKey = &v
	return s
}

type ListSaasPermissionGroupInfosResponseBody struct {
	Data []*ListSaasPermissionGroupInfosResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 8AD9FA10-7780-5E12-B701-13C928524F32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSaasPermissionGroupInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSaasPermissionGroupInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosResponseBody) SetData(v []*ListSaasPermissionGroupInfosResponseBodyData) *ListSaasPermissionGroupInfosResponseBody {
	s.Data = v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBody) SetRequestId(v string) *ListSaasPermissionGroupInfosResponseBody {
	s.RequestId = &v
	return s
}

type ListSaasPermissionGroupInfosResponseBodyData struct {
	// example:
	//
	// Release Center
	EnName  *string                                                `json:"EnName,omitempty" xml:"EnName,omitempty"`
	Name    *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PgInfos []*ListSaasPermissionGroupInfosResponseBodyDataPgInfos `json:"PgInfos,omitempty" xml:"PgInfos,omitempty" type:"Repeated"`
	// example:
	//
	// FAQ
	SaasCode *string `json:"SaasCode,omitempty" xml:"SaasCode,omitempty"`
}

func (s ListSaasPermissionGroupInfosResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSaasPermissionGroupInfosResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) SetEnName(v string) *ListSaasPermissionGroupInfosResponseBodyData {
	s.EnName = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) SetName(v string) *ListSaasPermissionGroupInfosResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) SetPgInfos(v []*ListSaasPermissionGroupInfosResponseBodyDataPgInfos) *ListSaasPermissionGroupInfosResponseBodyData {
	s.PgInfos = v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) SetSaasCode(v string) *ListSaasPermissionGroupInfosResponseBodyData {
	s.SaasCode = &v
	return s
}

type ListSaasPermissionGroupInfosResponseBodyDataPgInfos struct {
	// example:
	//
	// FAQ
	PgCode *string `json:"PgCode,omitempty" xml:"PgCode,omitempty"`
	// example:
	//
	// FAQ
	PgEnName *string `json:"PgEnName,omitempty" xml:"PgEnName,omitempty"`
	PgName   *string `json:"PgName,omitempty" xml:"PgName,omitempty"`
}

func (s ListSaasPermissionGroupInfosResponseBodyDataPgInfos) String() string {
	return tea.Prettify(s)
}

func (s ListSaasPermissionGroupInfosResponseBodyDataPgInfos) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) SetPgCode(v string) *ListSaasPermissionGroupInfosResponseBodyDataPgInfos {
	s.PgCode = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) SetPgEnName(v string) *ListSaasPermissionGroupInfosResponseBodyDataPgInfos {
	s.PgEnName = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) SetPgName(v string) *ListSaasPermissionGroupInfosResponseBodyDataPgInfos {
	s.PgName = &v
	return s
}

type ListSaasPermissionGroupInfosResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSaasPermissionGroupInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSaasPermissionGroupInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSaasPermissionGroupInfosResponse) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosResponse) SetHeaders(v map[string]*string) *ListSaasPermissionGroupInfosResponse {
	s.Headers = v
	return s
}

func (s *ListSaasPermissionGroupInfosResponse) SetStatusCode(v int32) *ListSaasPermissionGroupInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponse) SetBody(v *ListSaasPermissionGroupInfosResponseBody) *ListSaasPermissionGroupInfosResponse {
	s.Body = v
	return s
}

type ListSimQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30002299537
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s ListSimQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSimQuestionRequest) GoString() string {
	return s.String()
}

func (s *ListSimQuestionRequest) SetAgentKey(v string) *ListSimQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *ListSimQuestionRequest) SetKnowledgeId(v int64) *ListSimQuestionRequest {
	s.KnowledgeId = &v
	return s
}

type ListSimQuestionResponseBody struct {
	// example:
	//
	// 15CD94CC-CBEB-4189-806C-A132D1F45D51
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SimQuestions []*ListSimQuestionResponseBodySimQuestions `json:"SimQuestions,omitempty" xml:"SimQuestions,omitempty" type:"Repeated"`
}

func (s ListSimQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSimQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *ListSimQuestionResponseBody) SetRequestId(v string) *ListSimQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSimQuestionResponseBody) SetSimQuestions(v []*ListSimQuestionResponseBodySimQuestions) *ListSimQuestionResponseBody {
	s.SimQuestions = v
	return s
}

type ListSimQuestionResponseBodySimQuestions struct {
	// example:
	//
	// 2022-05-30T02:08:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-13T03:49:28Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 30001979424
	SimQuestionId *int64  `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListSimQuestionResponseBodySimQuestions) String() string {
	return tea.Prettify(s)
}

func (s ListSimQuestionResponseBodySimQuestions) GoString() string {
	return s.String()
}

func (s *ListSimQuestionResponseBodySimQuestions) SetCreateTime(v string) *ListSimQuestionResponseBodySimQuestions {
	s.CreateTime = &v
	return s
}

func (s *ListSimQuestionResponseBodySimQuestions) SetModifyTime(v string) *ListSimQuestionResponseBodySimQuestions {
	s.ModifyTime = &v
	return s
}

func (s *ListSimQuestionResponseBodySimQuestions) SetSimQuestionId(v int64) *ListSimQuestionResponseBodySimQuestions {
	s.SimQuestionId = &v
	return s
}

func (s *ListSimQuestionResponseBodySimQuestions) SetTitle(v string) *ListSimQuestionResponseBodySimQuestions {
	s.Title = &v
	return s
}

type ListSimQuestionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSimQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSimQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSimQuestionResponse) GoString() string {
	return s.String()
}

func (s *ListSimQuestionResponse) SetHeaders(v map[string]*string) *ListSimQuestionResponse {
	s.Headers = v
	return s
}

func (s *ListSimQuestionResponse) SetStatusCode(v int32) *ListSimQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSimQuestionResponse) SetBody(v *ListSimQuestionResponseBody) *ListSimQuestionResponse {
	s.Body = v
	return s
}

type ListSolutionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s ListSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSolutionRequest) GoString() string {
	return s.String()
}

func (s *ListSolutionRequest) SetAgentKey(v string) *ListSolutionRequest {
	s.AgentKey = &v
	return s
}

func (s *ListSolutionRequest) SetKnowledgeId(v int64) *ListSolutionRequest {
	s.KnowledgeId = &v
	return s
}

type ListSolutionResponseBody struct {
	// example:
	//
	// 5B29DB5E-251D-5A73-84B5-A12DF795F231
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Solutions []*ListSolutionResponseBodySolutions `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
}

func (s ListSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *ListSolutionResponseBody) SetRequestId(v string) *ListSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSolutionResponseBody) SetSolutions(v []*ListSolutionResponseBodySolutions) *ListSolutionResponseBody {
	s.Solutions = v
	return s
}

type ListSolutionResponseBodySolutions struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 2022-03-29T03:55:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-03-29T06:23:53Z
	ModifyTime       *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
	PlainText        *string   `json:"PlainText,omitempty" xml:"PlainText,omitempty"`
	// example:
	//
	// 496
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s ListSolutionResponseBodySolutions) String() string {
	return tea.Prettify(s)
}

func (s ListSolutionResponseBodySolutions) GoString() string {
	return s.String()
}

func (s *ListSolutionResponseBodySolutions) SetContent(v string) *ListSolutionResponseBodySolutions {
	s.Content = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetContentType(v int32) *ListSolutionResponseBodySolutions {
	s.ContentType = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetCreateTime(v string) *ListSolutionResponseBodySolutions {
	s.CreateTime = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetModifyTime(v string) *ListSolutionResponseBodySolutions {
	s.ModifyTime = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetPerspectiveCodes(v []*string) *ListSolutionResponseBodySolutions {
	s.PerspectiveCodes = v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetPlainText(v string) *ListSolutionResponseBodySolutions {
	s.PlainText = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetSolutionId(v int64) *ListSolutionResponseBodySolutions {
	s.SolutionId = &v
	return s
}

type ListSolutionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSolutionResponse) GoString() string {
	return s.String()
}

func (s *ListSolutionResponse) SetHeaders(v map[string]*string) *ListSolutionResponse {
	s.Headers = v
	return s
}

func (s *ListSolutionResponse) SetStatusCode(v int32) *ListSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSolutionResponse) SetBody(v *ListSolutionResponseBody) *ListSolutionResponse {
	s.Body = v
	return s
}

type ListTongyiChatHistorysRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-04-01 08:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 30
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-7QuUfaqMQe
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-04-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListTongyiChatHistorysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTongyiChatHistorysRequest) GoString() string {
	return s.String()
}

func (s *ListTongyiChatHistorysRequest) SetAgentKey(v string) *ListTongyiChatHistorysRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) SetEndTime(v string) *ListTongyiChatHistorysRequest {
	s.EndTime = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) SetLimit(v int32) *ListTongyiChatHistorysRequest {
	s.Limit = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) SetRobotInstanceId(v string) *ListTongyiChatHistorysRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) SetStartTime(v string) *ListTongyiChatHistorysRequest {
	s.StartTime = &v
	return s
}

type ListTongyiChatHistorysResponseBody struct {
	// example:
	//
	// 116
	CostTime *string                  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Datas    []map[string]interface{} `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// example:
	//
	// D0DDFC4C-D66D-4787-9AE4-4D757481EDEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTongyiChatHistorysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTongyiChatHistorysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTongyiChatHistorysResponseBody) SetCostTime(v string) *ListTongyiChatHistorysResponseBody {
	s.CostTime = &v
	return s
}

func (s *ListTongyiChatHistorysResponseBody) SetDatas(v []map[string]interface{}) *ListTongyiChatHistorysResponseBody {
	s.Datas = v
	return s
}

func (s *ListTongyiChatHistorysResponseBody) SetRequestId(v string) *ListTongyiChatHistorysResponseBody {
	s.RequestId = &v
	return s
}

type ListTongyiChatHistorysResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTongyiChatHistorysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTongyiChatHistorysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTongyiChatHistorysResponse) GoString() string {
	return s.String()
}

func (s *ListTongyiChatHistorysResponse) SetHeaders(v map[string]*string) *ListTongyiChatHistorysResponse {
	s.Headers = v
	return s
}

func (s *ListTongyiChatHistorysResponse) SetStatusCode(v int32) *ListTongyiChatHistorysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTongyiChatHistorysResponse) SetBody(v *ListTongyiChatHistorysResponseBody) *ListTongyiChatHistorysResponse {
	s.Body = v
	return s
}

type ListUserSayRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 您做核酸了嘛
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 232
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserSayRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserSayRequest) GoString() string {
	return s.String()
}

func (s *ListUserSayRequest) SetAgentKey(v string) *ListUserSayRequest {
	s.AgentKey = &v
	return s
}

func (s *ListUserSayRequest) SetContent(v string) *ListUserSayRequest {
	s.Content = &v
	return s
}

func (s *ListUserSayRequest) SetInstanceId(v string) *ListUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserSayRequest) SetIntentId(v int64) *ListUserSayRequest {
	s.IntentId = &v
	return s
}

func (s *ListUserSayRequest) SetPageNumber(v int32) *ListUserSayRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserSayRequest) SetPageSize(v int32) *ListUserSayRequest {
	s.PageSize = &v
	return s
}

type ListUserSayResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fs1fg4512v43572v23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserSays   []*ListUserSayResponseBodyUserSays `json:"UserSays,omitempty" xml:"UserSays,omitempty" type:"Repeated"`
}

func (s ListUserSayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserSayResponseBody) SetPageNumber(v int32) *ListUserSayResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserSayResponseBody) SetPageSize(v int32) *ListUserSayResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserSayResponseBody) SetRequestId(v string) *ListUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserSayResponseBody) SetTotalCount(v int32) *ListUserSayResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserSayResponseBody) SetUserSays(v []*ListUserSayResponseBodyUserSays) *ListUserSayResponseBody {
	s.UserSays = v
	return s
}

type ListUserSayResponseBodyUserSays struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 235564564
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string                                     `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SlotInfos  []*ListUserSayResponseBodyUserSaysSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 3453452138
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s ListUserSayResponseBodyUserSays) String() string {
	return tea.Prettify(s)
}

func (s ListUserSayResponseBodyUserSays) GoString() string {
	return s.String()
}

func (s *ListUserSayResponseBodyUserSays) SetContent(v string) *ListUserSayResponseBodyUserSays {
	s.Content = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetCreateTime(v string) *ListUserSayResponseBodyUserSays {
	s.CreateTime = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetIntentId(v int64) *ListUserSayResponseBodyUserSays {
	s.IntentId = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetModifyTime(v string) *ListUserSayResponseBodyUserSays {
	s.ModifyTime = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetSlotInfos(v []*ListUserSayResponseBodyUserSaysSlotInfos) *ListUserSayResponseBodyUserSays {
	s.SlotInfos = v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetUserSayId(v int64) *ListUserSayResponseBodyUserSays {
	s.UserSayId = &v
	return s
}

type ListUserSayResponseBodyUserSaysSlotInfos struct {
	// example:
	//
	// 4
	EndIndex *int32 `json:"EndIndex,omitempty" xml:"EndIndex,omitempty"`
	// example:
	//
	// 3456sdfg3tu
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// example:
	//
	// 2
	StartIndex *int32 `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
}

func (s ListUserSayResponseBodyUserSaysSlotInfos) String() string {
	return tea.Prettify(s)
}

func (s ListUserSayResponseBodyUserSaysSlotInfos) GoString() string {
	return s.String()
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) SetEndIndex(v int32) *ListUserSayResponseBodyUserSaysSlotInfos {
	s.EndIndex = &v
	return s
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) SetSlotId(v string) *ListUserSayResponseBodyUserSaysSlotInfos {
	s.SlotId = &v
	return s
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) SetStartIndex(v int32) *ListUserSayResponseBodyUserSaysSlotInfos {
	s.StartIndex = &v
	return s
}

type ListUserSayResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserSayResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserSayResponse) GoString() string {
	return s.String()
}

func (s *ListUserSayResponse) SetHeaders(v map[string]*string) *ListUserSayResponse {
	s.Headers = v
	return s
}

func (s *ListUserSayResponse) SetStatusCode(v int32) *ListUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserSayResponse) SetBody(v *ListUserSayResponseBody) *ListUserSayResponse {
	s.Body = v
	return s
}

type NluRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 北京的天气怎么样
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s NluRequest) String() string {
	return tea.Prettify(s)
}

func (s NluRequest) GoString() string {
	return s.String()
}

func (s *NluRequest) SetAgentKey(v string) *NluRequest {
	s.AgentKey = &v
	return s
}

func (s *NluRequest) SetInstanceId(v string) *NluRequest {
	s.InstanceId = &v
	return s
}

func (s *NluRequest) SetUtterance(v string) *NluRequest {
	s.Utterance = &v
	return s
}

type NluResponseBody struct {
	// example:
	//
	// 2828708A-2C7A-1BAE-B810-87DB9DA9C661
	MessageId *string                    `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Messages  []*NluResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// A6357C1B-1D79-1382-B259-BD9E80751B42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NluResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBody) GoString() string {
	return s.String()
}

func (s *NluResponseBody) SetMessageId(v string) *NluResponseBody {
	s.MessageId = &v
	return s
}

func (s *NluResponseBody) SetMessages(v []*NluResponseBodyMessages) *NluResponseBody {
	s.Messages = v
	return s
}

func (s *NluResponseBody) SetRequestId(v string) *NluResponseBody {
	s.RequestId = &v
	return s
}

type NluResponseBodyMessages struct {
	DialogHubNluInfo *NluResponseBodyMessagesDialogHubNluInfo `json:"DialogHubNluInfo,omitempty" xml:"DialogHubNluInfo,omitempty" type:"Struct"`
	DsNluInfo        *NluResponseBodyMessagesDsNluInfo        `json:"DsNluInfo,omitempty" xml:"DsNluInfo,omitempty" type:"Struct"`
}

func (s NluResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessages) SetDialogHubNluInfo(v *NluResponseBodyMessagesDialogHubNluInfo) *NluResponseBodyMessages {
	s.DialogHubNluInfo = v
	return s
}

func (s *NluResponseBodyMessages) SetDsNluInfo(v *NluResponseBodyMessagesDsNluInfo) *NluResponseBodyMessages {
	s.DsNluInfo = v
	return s
}

type NluResponseBodyMessagesDialogHubNluInfo struct {
	GlobalDictList          []*NluResponseBodyMessagesDialogHubNluInfoGlobalDictList          `json:"GlobalDictList,omitempty" xml:"GlobalDictList,omitempty" type:"Repeated"`
	GlobalSensitiveWordList []*NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList `json:"GlobalSensitiveWordList,omitempty" xml:"GlobalSensitiveWordList,omitempty" type:"Repeated"`
}

func (s NluResponseBodyMessagesDialogHubNluInfo) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBodyMessagesDialogHubNluInfo) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDialogHubNluInfo) SetGlobalDictList(v []*NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) *NluResponseBodyMessagesDialogHubNluInfo {
	s.GlobalDictList = v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfo) SetGlobalSensitiveWordList(v []*NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) *NluResponseBodyMessagesDialogHubNluInfo {
	s.GlobalSensitiveWordList = v
	return s
}

type NluResponseBodyMessagesDialogHubNluInfoGlobalDictList struct {
	// example:
	//
	// 天气
	StandardWord *string `json:"StandardWord,omitempty" xml:"StandardWord,omitempty"`
	// example:
	//
	// 天气
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) SetStandardWord(v string) *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList {
	s.StandardWord = &v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) SetWord(v string) *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList {
	s.Word = &v
	return s
}

type NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList struct {
	// example:
	//
	// 天气
	StandardWord *string `json:"StandardWord,omitempty" xml:"StandardWord,omitempty"`
	// example:
	//
	// 天气
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) SetStandardWord(v string) *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList {
	s.StandardWord = &v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) SetWord(v string) *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList {
	s.Word = &v
	return s
}

type NluResponseBodyMessagesDsNluInfo struct {
	EntityList []*NluResponseBodyMessagesDsNluInfoEntityList `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Repeated"`
	IntentList []*NluResponseBodyMessagesDsNluInfoIntentList `json:"IntentList,omitempty" xml:"IntentList,omitempty" type:"Repeated"`
}

func (s NluResponseBodyMessagesDsNluInfo) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBodyMessagesDsNluInfo) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDsNluInfo) SetEntityList(v []*NluResponseBodyMessagesDsNluInfoEntityList) *NluResponseBodyMessagesDsNluInfo {
	s.EntityList = v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfo) SetIntentList(v []*NluResponseBodyMessagesDsNluInfoIntentList) *NluResponseBodyMessagesDsNluInfo {
	s.IntentList = v
	return s
}

type NluResponseBodyMessagesDsNluInfoEntityList struct {
	// example:
	//
	// @城市
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 北京
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 首都
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s NluResponseBodyMessagesDsNluInfoEntityList) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBodyMessagesDsNluInfoEntityList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) SetName(v string) *NluResponseBodyMessagesDsNluInfoEntityList {
	s.Name = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) SetOrigin(v string) *NluResponseBodyMessagesDsNluInfoEntityList {
	s.Origin = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) SetType(v string) *NluResponseBodyMessagesDsNluInfoEntityList {
	s.Type = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) SetValue(v string) *NluResponseBodyMessagesDsNluInfoEntityList {
	s.Value = &v
	return s
}

type NluResponseBodyMessagesDsNluInfoIntentList struct {
	// example:
	//
	// 724387
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// classifierType=Fewshot,from=Fewshot,content=[我要查北京的天气, 帮我查北京的天气, 北京天气怎么样, 北京今天下雨吗, 北京今天多少度]
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
	// example:
	//
	// FewShotLearning
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// example:
	//
	// 查天气意图
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.995
	Score    *float64                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	SlotList []*NluResponseBodyMessagesDsNluInfoIntentListSlotList `json:"SlotList,omitempty" xml:"SlotList,omitempty" type:"Repeated"`
}

func (s NluResponseBodyMessagesDsNluInfoIntentList) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBodyMessagesDsNluInfoIntentList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetIntentId(v int64) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.IntentId = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetMatchDetail(v string) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.MatchDetail = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetMatchType(v string) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.MatchType = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetName(v string) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.Name = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetScore(v float64) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.Score = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetSlotList(v []*NluResponseBodyMessagesDsNluInfoIntentListSlotList) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.SlotList = v
	return s
}

type NluResponseBodyMessagesDsNluInfoIntentListSlotList struct {
	// example:
	//
	// @城市
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 北京
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 首都
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s NluResponseBodyMessagesDsNluInfoIntentListSlotList) String() string {
	return tea.Prettify(s)
}

func (s NluResponseBodyMessagesDsNluInfoIntentListSlotList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) SetName(v string) *NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	s.Name = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) SetOrigin(v string) *NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	s.Origin = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) SetType(v string) *NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	s.Type = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) SetValue(v string) *NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	s.Value = &v
	return s
}

type NluResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NluResponseBody   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NluResponse) String() string {
	return tea.Prettify(s)
}

func (s NluResponse) GoString() string {
	return s.String()
}

func (s *NluResponse) SetHeaders(v map[string]*string) *NluResponse {
	s.Headers = v
	return s
}

func (s *NluResponse) SetStatusCode(v int32) *NluResponse {
	s.StatusCode = &v
	return s
}

func (s *NluResponse) SetBody(v *NluResponseBody) *NluResponse {
	s.Body = v
	return s
}

type QueryPerspectivesRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s QueryPerspectivesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesRequest) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesRequest) SetAgentKey(v string) *QueryPerspectivesRequest {
	s.AgentKey = &v
	return s
}

type QueryPerspectivesResponseBody struct {
	Perspectives []*QueryPerspectivesResponseBodyPerspectives `json:"Perspectives,omitempty" xml:"Perspectives,omitempty" type:"Repeated"`
	// example:
	//
	// F285D735-D580-18A8-B97F-B2E72B00F101
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPerspectivesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponseBody) SetPerspectives(v []*QueryPerspectivesResponseBodyPerspectives) *QueryPerspectivesResponseBody {
	s.Perspectives = v
	return s
}

func (s *QueryPerspectivesResponseBody) SetRequestId(v string) *QueryPerspectivesResponseBody {
	s.RequestId = &v
	return s
}

type QueryPerspectivesResponseBodyPerspectives struct {
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-04-29T03:38:54Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 移动端视角
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FZJBY3raWr
	PerspectiveCode *string `json:"PerspectiveCode,omitempty" xml:"PerspectiveCode,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	// example:
	//
	// true
	SelfDefine *bool `json:"SelfDefine,omitempty" xml:"SelfDefine,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryPerspectivesResponseBodyPerspectives) String() string {
	return tea.Prettify(s)
}

func (s QueryPerspectivesResponseBodyPerspectives) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetCreateTime(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.CreateTime = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetModifyTime(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.ModifyTime = &v
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

func (s *QueryPerspectivesResponseBodyPerspectives) SetPerspectiveId(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.PerspectiveId = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetSelfDefine(v bool) *QueryPerspectivesResponseBodyPerspectives {
	s.SelfDefine = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetStatus(v int32) *QueryPerspectivesResponseBodyPerspectives {
	s.Status = &v
	return s
}

type QueryPerspectivesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPerspectivesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryPerspectivesResponse) SetStatusCode(v int32) *QueryPerspectivesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPerspectivesResponse) SetBody(v *QueryPerspectivesResponseBody) *QueryPerspectivesResponse {
	s.Body = v
	return s
}

type RetryDocRequest struct {
	// example:
	//
	// e2a20f74cd9042558002c0f7dc873739_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s RetryDocRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryDocRequest) GoString() string {
	return s.String()
}

func (s *RetryDocRequest) SetAgentKey(v string) *RetryDocRequest {
	s.AgentKey = &v
	return s
}

func (s *RetryDocRequest) SetKnowledgeId(v int64) *RetryDocRequest {
	s.KnowledgeId = &v
	return s
}

type RetryDocResponseBody struct {
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6419BA93-D111-5225-8998-13E63E6D3940
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryDocResponseBody) GoString() string {
	return s.String()
}

func (s *RetryDocResponseBody) SetKnowledgeId(v int64) *RetryDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *RetryDocResponseBody) SetRequestId(v string) *RetryDocResponseBody {
	s.RequestId = &v
	return s
}

type RetryDocResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryDocResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryDocResponse) GoString() string {
	return s.String()
}

func (s *RetryDocResponse) SetHeaders(v map[string]*string) *RetryDocResponse {
	s.Headers = v
	return s
}

func (s *RetryDocResponse) SetStatusCode(v int32) *RetryDocResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryDocResponse) SetBody(v *RetryDocResponseBody) *RetryDocResponse {
	s.Body = v
	return s
}

type SearchDocRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey    *string  `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryIds []*int64 `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	CreateTimeBegin *string `json:"CreateTimeBegin,omitempty" xml:"CreateTimeBegin,omitempty"`
	// example:
	//
	// 2022-05-02T03:09:30Z
	CreateTimeEnd  *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" xml:"EndTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	ModifyTimeBegin *string `json:"ModifyTimeBegin,omitempty" xml:"ModifyTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	ModifyTimeEnd  *string `json:"ModifyTimeEnd,omitempty" xml:"ModifyTimeEnd,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// example:
	//
	// 1
	SearchScope *int32 `json:"SearchScope,omitempty" xml:"SearchScope,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" xml:"StartTimeBegin,omitempty"`
	// example:
	//
	// 2022-04-03T03:09:30Z
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" xml:"StartTimeEnd,omitempty"`
	// example:
	//
	// 1
	Status *int32   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*int64 `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s SearchDocRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDocRequest) GoString() string {
	return s.String()
}

func (s *SearchDocRequest) SetAgentKey(v string) *SearchDocRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchDocRequest) SetCategoryIds(v []*int64) *SearchDocRequest {
	s.CategoryIds = v
	return s
}

func (s *SearchDocRequest) SetCreateTimeBegin(v string) *SearchDocRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchDocRequest) SetCreateTimeEnd(v string) *SearchDocRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchDocRequest) SetCreateUserName(v string) *SearchDocRequest {
	s.CreateUserName = &v
	return s
}

func (s *SearchDocRequest) SetEndTimeBegin(v string) *SearchDocRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *SearchDocRequest) SetEndTimeEnd(v string) *SearchDocRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *SearchDocRequest) SetKeyword(v string) *SearchDocRequest {
	s.Keyword = &v
	return s
}

func (s *SearchDocRequest) SetModifyTimeBegin(v string) *SearchDocRequest {
	s.ModifyTimeBegin = &v
	return s
}

func (s *SearchDocRequest) SetModifyTimeEnd(v string) *SearchDocRequest {
	s.ModifyTimeEnd = &v
	return s
}

func (s *SearchDocRequest) SetModifyUserName(v string) *SearchDocRequest {
	s.ModifyUserName = &v
	return s
}

func (s *SearchDocRequest) SetPageNumber(v int32) *SearchDocRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchDocRequest) SetPageSize(v int32) *SearchDocRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDocRequest) SetProcessStatus(v int32) *SearchDocRequest {
	s.ProcessStatus = &v
	return s
}

func (s *SearchDocRequest) SetSearchScope(v int32) *SearchDocRequest {
	s.SearchScope = &v
	return s
}

func (s *SearchDocRequest) SetStartTimeBegin(v string) *SearchDocRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *SearchDocRequest) SetStartTimeEnd(v string) *SearchDocRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *SearchDocRequest) SetStatus(v int32) *SearchDocRequest {
	s.Status = &v
	return s
}

func (s *SearchDocRequest) SetTagIds(v []*int64) *SearchDocRequest {
	s.TagIds = v
	return s
}

type SearchDocShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey          *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	CreateTimeBegin *string `json:"CreateTimeBegin,omitempty" xml:"CreateTimeBegin,omitempty"`
	// example:
	//
	// 2022-05-02T03:09:30Z
	CreateTimeEnd  *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" xml:"EndTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	ModifyTimeBegin *string `json:"ModifyTimeBegin,omitempty" xml:"ModifyTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	ModifyTimeEnd  *string `json:"ModifyTimeEnd,omitempty" xml:"ModifyTimeEnd,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// example:
	//
	// 1
	SearchScope *int32 `json:"SearchScope,omitempty" xml:"SearchScope,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" xml:"StartTimeBegin,omitempty"`
	// example:
	//
	// 2022-04-03T03:09:30Z
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" xml:"StartTimeEnd,omitempty"`
	// example:
	//
	// 1
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
}

func (s SearchDocShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchDocShrinkRequest) SetAgentKey(v string) *SearchDocShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchDocShrinkRequest) SetCategoryIdsShrink(v string) *SearchDocShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *SearchDocShrinkRequest) SetCreateTimeBegin(v string) *SearchDocShrinkRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchDocShrinkRequest) SetCreateTimeEnd(v string) *SearchDocShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchDocShrinkRequest) SetCreateUserName(v string) *SearchDocShrinkRequest {
	s.CreateUserName = &v
	return s
}

func (s *SearchDocShrinkRequest) SetEndTimeBegin(v string) *SearchDocShrinkRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *SearchDocShrinkRequest) SetEndTimeEnd(v string) *SearchDocShrinkRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *SearchDocShrinkRequest) SetKeyword(v string) *SearchDocShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *SearchDocShrinkRequest) SetModifyTimeBegin(v string) *SearchDocShrinkRequest {
	s.ModifyTimeBegin = &v
	return s
}

func (s *SearchDocShrinkRequest) SetModifyTimeEnd(v string) *SearchDocShrinkRequest {
	s.ModifyTimeEnd = &v
	return s
}

func (s *SearchDocShrinkRequest) SetModifyUserName(v string) *SearchDocShrinkRequest {
	s.ModifyUserName = &v
	return s
}

func (s *SearchDocShrinkRequest) SetPageNumber(v int32) *SearchDocShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchDocShrinkRequest) SetPageSize(v int32) *SearchDocShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDocShrinkRequest) SetProcessStatus(v int32) *SearchDocShrinkRequest {
	s.ProcessStatus = &v
	return s
}

func (s *SearchDocShrinkRequest) SetSearchScope(v int32) *SearchDocShrinkRequest {
	s.SearchScope = &v
	return s
}

func (s *SearchDocShrinkRequest) SetStartTimeBegin(v string) *SearchDocShrinkRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *SearchDocShrinkRequest) SetStartTimeEnd(v string) *SearchDocShrinkRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *SearchDocShrinkRequest) SetStatus(v int32) *SearchDocShrinkRequest {
	s.Status = &v
	return s
}

func (s *SearchDocShrinkRequest) SetTagIdsShrink(v string) *SearchDocShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

type SearchDocResponseBody struct {
	DocHits []*SearchDocResponseBodyDocHits `json:"DocHits,omitempty" xml:"DocHits,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E3E5C779-A630-45AC-B0F2-A4506A4212F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 141
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchDocResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDocResponseBody) SetDocHits(v []*SearchDocResponseBodyDocHits) *SearchDocResponseBody {
	s.DocHits = v
	return s
}

func (s *SearchDocResponseBody) SetPageNumber(v int32) *SearchDocResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchDocResponseBody) SetPageSize(v int32) *SearchDocResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchDocResponseBody) SetRequestId(v string) *SearchDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchDocResponseBody) SetTotalCount(v int32) *SearchDocResponseBody {
	s.TotalCount = &v
	return s
}

type SearchDocResponseBodyDocHits struct {
	// example:
	//
	// cn_dytns
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 30000135654
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":"docName"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2023-06-22T03:53:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 111111111
	CreateUserId   *int64                                 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string                                `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DocName        *string                                `json:"DocName,omitempty" xml:"DocName,omitempty"`
	DocTags        []*SearchDocResponseBodyDocHitsDocTags `json:"DocTags,omitempty" xml:"DocTags,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	EffectStatus *int32 `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// example:
	//
	// 2099-12-31T16:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 30002692007
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2023-06-25T02:27:42Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 222222222
	ModifyUserId   *int64  `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// true
	ProcessCanRetry *bool   `json:"ProcessCanRetry,omitempty" xml:"ProcessCanRetry,omitempty"`
	ProcessMessage  *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty"`
	// example:
	//
	// 0
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// example:
	//
	// 2023-02-28T11:40:18Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 1
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SearchDocResponseBodyDocHits) String() string {
	return tea.Prettify(s)
}

func (s SearchDocResponseBodyDocHits) GoString() string {
	return s.String()
}

func (s *SearchDocResponseBodyDocHits) SetBizCode(v string) *SearchDocResponseBodyDocHits {
	s.BizCode = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetCategoryId(v int64) *SearchDocResponseBodyDocHits {
	s.CategoryId = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetConfig(v string) *SearchDocResponseBodyDocHits {
	s.Config = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetCreateTime(v string) *SearchDocResponseBodyDocHits {
	s.CreateTime = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetCreateUserId(v int64) *SearchDocResponseBodyDocHits {
	s.CreateUserId = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetCreateUserName(v string) *SearchDocResponseBodyDocHits {
	s.CreateUserName = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetDocName(v string) *SearchDocResponseBodyDocHits {
	s.DocName = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetDocTags(v []*SearchDocResponseBodyDocHitsDocTags) *SearchDocResponseBodyDocHits {
	s.DocTags = v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetEffectStatus(v int32) *SearchDocResponseBodyDocHits {
	s.EffectStatus = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetEndDate(v string) *SearchDocResponseBodyDocHits {
	s.EndDate = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetKnowledgeId(v int64) *SearchDocResponseBodyDocHits {
	s.KnowledgeId = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetMeta(v string) *SearchDocResponseBodyDocHits {
	s.Meta = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetModifyTime(v string) *SearchDocResponseBodyDocHits {
	s.ModifyTime = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetModifyUserId(v int64) *SearchDocResponseBodyDocHits {
	s.ModifyUserId = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetModifyUserName(v string) *SearchDocResponseBodyDocHits {
	s.ModifyUserName = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetProcessCanRetry(v bool) *SearchDocResponseBodyDocHits {
	s.ProcessCanRetry = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetProcessMessage(v string) *SearchDocResponseBodyDocHits {
	s.ProcessMessage = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetProcessStatus(v int32) *SearchDocResponseBodyDocHits {
	s.ProcessStatus = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetStartDate(v string) *SearchDocResponseBodyDocHits {
	s.StartDate = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetStatus(v int32) *SearchDocResponseBodyDocHits {
	s.Status = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetUrl(v string) *SearchDocResponseBodyDocHits {
	s.Url = &v
	return s
}

type SearchDocResponseBodyDocHitsDocTags struct {
	DefaultTag *bool   `json:"DefaultTag,omitempty" xml:"DefaultTag,omitempty"`
	GroupId    *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	TagId      *int64  `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName    *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s SearchDocResponseBodyDocHitsDocTags) String() string {
	return tea.Prettify(s)
}

func (s SearchDocResponseBodyDocHitsDocTags) GoString() string {
	return s.String()
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetDefaultTag(v bool) *SearchDocResponseBodyDocHitsDocTags {
	s.DefaultTag = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetGroupId(v int64) *SearchDocResponseBodyDocHitsDocTags {
	s.GroupId = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetGroupName(v string) *SearchDocResponseBodyDocHitsDocTags {
	s.GroupName = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetTagId(v int64) *SearchDocResponseBodyDocHitsDocTags {
	s.TagId = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetTagName(v string) *SearchDocResponseBodyDocHitsDocTags {
	s.TagName = &v
	return s
}

type SearchDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDocResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDocResponse) GoString() string {
	return s.String()
}

func (s *SearchDocResponse) SetHeaders(v map[string]*string) *SearchDocResponse {
	s.Headers = v
	return s
}

func (s *SearchDocResponse) SetStatusCode(v int32) *SearchDocResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDocResponse) SetBody(v *SearchDocResponseBody) *SearchDocResponse {
	s.Body = v
	return s
}

type SearchFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey    *string  `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryIds []*int64 `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	CreateTimeBegin *string `json:"CreateTimeBegin,omitempty" xml:"CreateTimeBegin,omitempty"`
	// example:
	//
	// 2022-05-02T03:09:30Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// test01
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" xml:"EndTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	ModifyTimeBegin *string `json:"ModifyTimeBegin,omitempty" xml:"ModifyTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	ModifyTimeEnd *string `json:"ModifyTimeEnd,omitempty" xml:"ModifyTimeEnd,omitempty"`
	// example:
	//
	// test01
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	SearchScope *int32 `json:"SearchScope,omitempty" xml:"SearchScope,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" xml:"StartTimeBegin,omitempty"`
	// example:
	//
	// 2022-04-03T03:09:30Z
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" xml:"StartTimeEnd,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SearchFaqRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaqRequest) GoString() string {
	return s.String()
}

func (s *SearchFaqRequest) SetAgentKey(v string) *SearchFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchFaqRequest) SetCategoryIds(v []*int64) *SearchFaqRequest {
	s.CategoryIds = v
	return s
}

func (s *SearchFaqRequest) SetCreateTimeBegin(v string) *SearchFaqRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchFaqRequest) SetCreateTimeEnd(v string) *SearchFaqRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchFaqRequest) SetCreateUserName(v string) *SearchFaqRequest {
	s.CreateUserName = &v
	return s
}

func (s *SearchFaqRequest) SetEndTimeBegin(v string) *SearchFaqRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *SearchFaqRequest) SetEndTimeEnd(v string) *SearchFaqRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *SearchFaqRequest) SetKeyword(v string) *SearchFaqRequest {
	s.Keyword = &v
	return s
}

func (s *SearchFaqRequest) SetModifyTimeBegin(v string) *SearchFaqRequest {
	s.ModifyTimeBegin = &v
	return s
}

func (s *SearchFaqRequest) SetModifyTimeEnd(v string) *SearchFaqRequest {
	s.ModifyTimeEnd = &v
	return s
}

func (s *SearchFaqRequest) SetModifyUserName(v string) *SearchFaqRequest {
	s.ModifyUserName = &v
	return s
}

func (s *SearchFaqRequest) SetPageNumber(v int32) *SearchFaqRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFaqRequest) SetPageSize(v int32) *SearchFaqRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFaqRequest) SetSearchScope(v int32) *SearchFaqRequest {
	s.SearchScope = &v
	return s
}

func (s *SearchFaqRequest) SetStartTimeBegin(v string) *SearchFaqRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *SearchFaqRequest) SetStartTimeEnd(v string) *SearchFaqRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *SearchFaqRequest) SetStatus(v int32) *SearchFaqRequest {
	s.Status = &v
	return s
}

type SearchFaqShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey          *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	CreateTimeBegin *string `json:"CreateTimeBegin,omitempty" xml:"CreateTimeBegin,omitempty"`
	// example:
	//
	// 2022-05-02T03:09:30Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// test01
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" xml:"EndTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	ModifyTimeBegin *string `json:"ModifyTimeBegin,omitempty" xml:"ModifyTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	ModifyTimeEnd *string `json:"ModifyTimeEnd,omitempty" xml:"ModifyTimeEnd,omitempty"`
	// example:
	//
	// test01
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	SearchScope *int32 `json:"SearchScope,omitempty" xml:"SearchScope,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" xml:"StartTimeBegin,omitempty"`
	// example:
	//
	// 2022-04-03T03:09:30Z
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" xml:"StartTimeEnd,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SearchFaqShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaqShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchFaqShrinkRequest) SetAgentKey(v string) *SearchFaqShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetCategoryIdsShrink(v string) *SearchFaqShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetCreateTimeBegin(v string) *SearchFaqShrinkRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetCreateTimeEnd(v string) *SearchFaqShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetCreateUserName(v string) *SearchFaqShrinkRequest {
	s.CreateUserName = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetEndTimeBegin(v string) *SearchFaqShrinkRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetEndTimeEnd(v string) *SearchFaqShrinkRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetKeyword(v string) *SearchFaqShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetModifyTimeBegin(v string) *SearchFaqShrinkRequest {
	s.ModifyTimeBegin = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetModifyTimeEnd(v string) *SearchFaqShrinkRequest {
	s.ModifyTimeEnd = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetModifyUserName(v string) *SearchFaqShrinkRequest {
	s.ModifyUserName = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetPageNumber(v int32) *SearchFaqShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetPageSize(v int32) *SearchFaqShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetSearchScope(v int32) *SearchFaqShrinkRequest {
	s.SearchScope = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetStartTimeBegin(v string) *SearchFaqShrinkRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetStartTimeEnd(v string) *SearchFaqShrinkRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetStatus(v int32) *SearchFaqShrinkRequest {
	s.Status = &v
	return s
}

type SearchFaqResponseBody struct {
	FaqHits []*SearchFaqResponseBodyFaqHits `json:"FaqHits,omitempty" xml:"FaqHits,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// E45491D5-7E0A-42C6-9B21-91D1066B1475
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1075
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchFaqResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFaqResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFaqResponseBody) SetFaqHits(v []*SearchFaqResponseBodyFaqHits) *SearchFaqResponseBody {
	s.FaqHits = v
	return s
}

func (s *SearchFaqResponseBody) SetPageNumber(v int32) *SearchFaqResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFaqResponseBody) SetPageSize(v int32) *SearchFaqResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchFaqResponseBody) SetRequestId(v string) *SearchFaqResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFaqResponseBody) SetTotalCount(v int32) *SearchFaqResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFaqResponseBodyFaqHits struct {
	// example:
	//
	// 30000055639
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 18453
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// test01
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 20
	EffectStatus     *int32    `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	HitSimilarTitles []*string `json:"HitSimilarTitles,omitempty" xml:"HitSimilarTitles,omitempty" type:"Repeated"`
	HitSolutions     []*string `json:"HitSolutions,omitempty" xml:"HitSolutions,omitempty" type:"Repeated"`
	// example:
	//
	// 30002145804
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 18453
	ModifyUserId *int64 `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// test01
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 3
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchFaqResponseBodyFaqHits) String() string {
	return tea.Prettify(s)
}

func (s SearchFaqResponseBodyFaqHits) GoString() string {
	return s.String()
}

func (s *SearchFaqResponseBodyFaqHits) SetCategoryId(v int64) *SearchFaqResponseBodyFaqHits {
	s.CategoryId = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetCreateTime(v string) *SearchFaqResponseBodyFaqHits {
	s.CreateTime = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetCreateUserId(v int64) *SearchFaqResponseBodyFaqHits {
	s.CreateUserId = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetCreateUserName(v string) *SearchFaqResponseBodyFaqHits {
	s.CreateUserName = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetEffectStatus(v int32) *SearchFaqResponseBodyFaqHits {
	s.EffectStatus = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetHitSimilarTitles(v []*string) *SearchFaqResponseBodyFaqHits {
	s.HitSimilarTitles = v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetHitSolutions(v []*string) *SearchFaqResponseBodyFaqHits {
	s.HitSolutions = v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetKnowledgeId(v int64) *SearchFaqResponseBodyFaqHits {
	s.KnowledgeId = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetModifyTime(v string) *SearchFaqResponseBodyFaqHits {
	s.ModifyTime = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetModifyUserId(v int64) *SearchFaqResponseBodyFaqHits {
	s.ModifyUserId = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetModifyUserName(v string) *SearchFaqResponseBodyFaqHits {
	s.ModifyUserName = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetStatus(v int32) *SearchFaqResponseBodyFaqHits {
	s.Status = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetTitle(v string) *SearchFaqResponseBodyFaqHits {
	s.Title = &v
	return s
}

type SearchFaqResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFaqResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFaqResponse) GoString() string {
	return s.String()
}

func (s *SearchFaqResponse) SetHeaders(v map[string]*string) *SearchFaqResponse {
	s.Headers = v
	return s
}

func (s *SearchFaqResponse) SetStatusCode(v int32) *SearchFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFaqResponse) SetBody(v *SearchFaqResponseBody) *SearchFaqResponse {
	s.Body = v
	return s
}

type UpdateCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	BizCode  *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 231001028593
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategoryRequest) SetAgentKey(v string) *UpdateCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateCategoryRequest) SetBizCode(v string) *UpdateCategoryRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateCategoryRequest) SetCategoryId(v int64) *UpdateCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateCategoryRequest) SetName(v string) *UpdateCategoryRequest {
	s.Name = &v
	return s
}

type UpdateCategoryResponseBody struct {
	// example:
	//
	// F79E7305-5314-5069-A701-9591AD051902
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type UpdateCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateCategoryResponse) SetStatusCode(v int32) *UpdateCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCategoryResponse) SetBody(v *UpdateCategoryResponseBody) *UpdateCategoryResponse {
	s.Body = v
	return s
}

type UpdateConnQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000000295
	ConnQuestionId *int64 `json:"ConnQuestionId,omitempty" xml:"ConnQuestionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 877397683
	OutlineId *int64 `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
}

func (s UpdateConnQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnQuestionRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnQuestionRequest) SetAgentKey(v string) *UpdateConnQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateConnQuestionRequest) SetConnQuestionId(v int64) *UpdateConnQuestionRequest {
	s.ConnQuestionId = &v
	return s
}

func (s *UpdateConnQuestionRequest) SetOutlineId(v int64) *UpdateConnQuestionRequest {
	s.OutlineId = &v
	return s
}

type UpdateConnQuestionResponseBody struct {
	// example:
	//
	// 004EB5C0-9DEB-53BF-A57A-0407A6D6B3C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConnQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnQuestionResponseBody) SetRequestId(v string) *UpdateConnQuestionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConnQuestionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConnQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConnQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnQuestionResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnQuestionResponse) SetHeaders(v map[string]*string) *UpdateConnQuestionResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnQuestionResponse) SetStatusCode(v int32) *UpdateConnQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnQuestionResponse) SetBody(v *UpdateConnQuestionResponseBody) *UpdateConnQuestionResponse {
	s.Body = v
	return s
}

type UpdateDSEntityRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 实体名称
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// synonyms
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDSEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDSEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityRequest) SetAgentKey(v string) *UpdateDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDSEntityRequest) SetEntityId(v int64) *UpdateDSEntityRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateDSEntityRequest) SetEntityName(v string) *UpdateDSEntityRequest {
	s.EntityName = &v
	return s
}

func (s *UpdateDSEntityRequest) SetEntityType(v string) *UpdateDSEntityRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateDSEntityRequest) SetInstanceId(v string) *UpdateDSEntityRequest {
	s.InstanceId = &v
	return s
}

type UpdateDSEntityResponseBody struct {
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// df23fgh4hyj67hn56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDSEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityResponseBody) SetEntityId(v int64) *UpdateDSEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *UpdateDSEntityResponseBody) SetRequestId(v string) *UpdateDSEntityResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDSEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDSEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDSEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDSEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityResponse) SetHeaders(v map[string]*string) *UpdateDSEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateDSEntityResponse) SetStatusCode(v int32) *UpdateDSEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDSEntityResponse) SetBody(v *UpdateDSEntityResponseBody) *UpdateDSEntityResponse {
	s.Body = v
	return s
}

type UpdateDSEntityValueRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 223423423
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2342377423
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s UpdateDSEntityValueRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDSEntityValueRequest) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityValueRequest) SetAgentKey(v string) *UpdateDSEntityValueRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetContent(v string) *UpdateDSEntityValueRequest {
	s.Content = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetEntityId(v int64) *UpdateDSEntityValueRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetEntityValueId(v int64) *UpdateDSEntityValueRequest {
	s.EntityValueId = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetInstanceId(v string) *UpdateDSEntityValueRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetSynonyms(v []*string) *UpdateDSEntityValueRequest {
	s.Synonyms = v
	return s
}

type UpdateDSEntityValueShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 223423423
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2342377423
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SynonymsShrink *string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty"`
}

func (s UpdateDSEntityValueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDSEntityValueShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityValueShrinkRequest) SetAgentKey(v string) *UpdateDSEntityValueShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetContent(v string) *UpdateDSEntityValueShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetEntityId(v int64) *UpdateDSEntityValueShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetEntityValueId(v int64) *UpdateDSEntityValueShrinkRequest {
	s.EntityValueId = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetInstanceId(v string) *UpdateDSEntityValueShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetSynonymsShrink(v string) *UpdateDSEntityValueShrinkRequest {
	s.SynonymsShrink = &v
	return s
}

type UpdateDSEntityValueResponseBody struct {
	// example:
	//
	// 2342377423
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// example:
	//
	// sDag3g43wesf2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDSEntityValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDSEntityValueResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityValueResponseBody) SetEntityValueId(v int64) *UpdateDSEntityValueResponseBody {
	s.EntityValueId = &v
	return s
}

func (s *UpdateDSEntityValueResponseBody) SetRequestId(v string) *UpdateDSEntityValueResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDSEntityValueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDSEntityValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDSEntityValueResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDSEntityValueResponse) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityValueResponse) SetHeaders(v map[string]*string) *UpdateDSEntityValueResponse {
	s.Headers = v
	return s
}

func (s *UpdateDSEntityValueResponse) SetStatusCode(v int32) *UpdateDSEntityValueResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDSEntityValueResponse) SetBody(v *UpdateDSEntityValueResponseBody) *UpdateDSEntityValueResponse {
	s.Body = v
	return s
}

type UpdateDocRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 231001028593
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":""}
	Config      *string                        `json:"Config,omitempty" xml:"Config,omitempty"`
	Content     *string                        `json:"Content,omitempty" xml:"Content,omitempty"`
	DocMetadata []*UpdateDocRequestDocMetadata `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty" type:"Repeated"`
	DocName     *string                        `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// 2023-03-11T23:59:59Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate *string  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIds    []*int64 `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	Title     *string  `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateDocRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocRequest) SetAgentKey(v string) *UpdateDocRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDocRequest) SetCategoryId(v int64) *UpdateDocRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateDocRequest) SetConfig(v string) *UpdateDocRequest {
	s.Config = &v
	return s
}

func (s *UpdateDocRequest) SetContent(v string) *UpdateDocRequest {
	s.Content = &v
	return s
}

func (s *UpdateDocRequest) SetDocMetadata(v []*UpdateDocRequestDocMetadata) *UpdateDocRequest {
	s.DocMetadata = v
	return s
}

func (s *UpdateDocRequest) SetDocName(v string) *UpdateDocRequest {
	s.DocName = &v
	return s
}

func (s *UpdateDocRequest) SetEndDate(v string) *UpdateDocRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateDocRequest) SetKnowledgeId(v int64) *UpdateDocRequest {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateDocRequest) SetMeta(v string) *UpdateDocRequest {
	s.Meta = &v
	return s
}

func (s *UpdateDocRequest) SetStartDate(v string) *UpdateDocRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateDocRequest) SetTagIds(v []*int64) *UpdateDocRequest {
	s.TagIds = v
	return s
}

func (s *UpdateDocRequest) SetTitle(v string) *UpdateDocRequest {
	s.Title = &v
	return s
}

type UpdateDocRequestDocMetadata struct {
	BusinessViewId      *string                                           `json:"BusinessViewId,omitempty" xml:"BusinessViewId,omitempty"`
	BusinessViewName    *string                                           `json:"BusinessViewName,omitempty" xml:"BusinessViewName,omitempty"`
	MetaCellInfoDTOList []*UpdateDocRequestDocMetadataMetaCellInfoDTOList `json:"MetaCellInfoDTOList,omitempty" xml:"MetaCellInfoDTOList,omitempty" type:"Repeated"`
}

func (s UpdateDocRequestDocMetadata) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocRequestDocMetadata) GoString() string {
	return s.String()
}

func (s *UpdateDocRequestDocMetadata) SetBusinessViewId(v string) *UpdateDocRequestDocMetadata {
	s.BusinessViewId = &v
	return s
}

func (s *UpdateDocRequestDocMetadata) SetBusinessViewName(v string) *UpdateDocRequestDocMetadata {
	s.BusinessViewName = &v
	return s
}

func (s *UpdateDocRequestDocMetadata) SetMetaCellInfoDTOList(v []*UpdateDocRequestDocMetadataMetaCellInfoDTOList) *UpdateDocRequestDocMetadata {
	s.MetaCellInfoDTOList = v
	return s
}

type UpdateDocRequestDocMetadataMetaCellInfoDTOList struct {
	FieldCode *string `json:"FieldCode,omitempty" xml:"FieldCode,omitempty"`
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDocRequestDocMetadataMetaCellInfoDTOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocRequestDocMetadataMetaCellInfoDTOList) GoString() string {
	return s.String()
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) SetFieldCode(v string) *UpdateDocRequestDocMetadataMetaCellInfoDTOList {
	s.FieldCode = &v
	return s
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) SetFieldName(v string) *UpdateDocRequestDocMetadataMetaCellInfoDTOList {
	s.FieldName = &v
	return s
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) SetValue(v string) *UpdateDocRequestDocMetadataMetaCellInfoDTOList {
	s.Value = &v
	return s
}

type UpdateDocShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 231001028593
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":""}
	Config            *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DocMetadataShrink *string `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty"`
	DocName           *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// 2023-03-11T23:59:59Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateDocShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocShrinkRequest) SetAgentKey(v string) *UpdateDocShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetCategoryId(v int64) *UpdateDocShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetConfig(v string) *UpdateDocShrinkRequest {
	s.Config = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetContent(v string) *UpdateDocShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetDocMetadataShrink(v string) *UpdateDocShrinkRequest {
	s.DocMetadataShrink = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetDocName(v string) *UpdateDocShrinkRequest {
	s.DocName = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetEndDate(v string) *UpdateDocShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetKnowledgeId(v int64) *UpdateDocShrinkRequest {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetMeta(v string) *UpdateDocShrinkRequest {
	s.Meta = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetStartDate(v string) *UpdateDocShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetTagIdsShrink(v string) *UpdateDocShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetTitle(v string) *UpdateDocShrinkRequest {
	s.Title = &v
	return s
}

type UpdateDocResponseBody struct {
	// example:
	//
	// 30002406051
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0F9F136A-1BF6-5CC1-9D57-9717761F03B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocResponseBody) SetKnowledgeId(v int64) *UpdateDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateDocResponseBody) SetRequestId(v string) *UpdateDocResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDocResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocResponse) SetHeaders(v map[string]*string) *UpdateDocResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocResponse) SetStatusCode(v int32) *UpdateDocResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocResponse) SetBody(v *UpdateDocResponseBody) *UpdateDocResponse {
	s.Body = v
	return s
}

type UpdateFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2030-12-31T16:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 2022-05-27T05:18:20Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateFaqRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaqRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaqRequest) SetAgentKey(v string) *UpdateFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateFaqRequest) SetCategoryId(v int64) *UpdateFaqRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateFaqRequest) SetEndDate(v string) *UpdateFaqRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateFaqRequest) SetKnowledgeId(v int64) *UpdateFaqRequest {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateFaqRequest) SetStartDate(v string) *UpdateFaqRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateFaqRequest) SetTitle(v string) *UpdateFaqRequest {
	s.Title = &v
	return s
}

type UpdateFaqResponseBody struct {
	// example:
	//
	// 736994BD-AA35-4742-88C9-E64BE4BAA14B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFaqResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaqResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaqResponseBody) SetRequestId(v string) *UpdateFaqResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFaqResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFaqResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaqResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaqResponse) SetHeaders(v map[string]*string) *UpdateFaqResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaqResponse) SetStatusCode(v int32) *UpdateFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaqResponse) SetBody(v *UpdateFaqResponseBody) *UpdateFaqResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 用于C端问答的机器人
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// 智能客服-小C
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetAgentKey(v string) *UpdateInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetIntroduction(v string) *UpdateInstanceRequest {
	s.Introduction = &v
	return s
}

func (s *UpdateInstanceRequest) SetName(v string) *UpdateInstanceRequest {
	s.Name = &v
	return s
}

type UpdateInstanceResponseBody struct {
	// example:
	//
	// C45BFEE4-F657-1332-8B47-2C757B94C972
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

type UpdateIntentRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId       *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentDefinition *UpdateIntentRequestIntentDefinition `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 234234234534
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s UpdateIntentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntentRequest) SetAgentKey(v string) *UpdateIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateIntentRequest) SetInstanceId(v string) *UpdateIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIntentRequest) SetIntentDefinition(v *UpdateIntentRequestIntentDefinition) *UpdateIntentRequest {
	s.IntentDefinition = v
	return s
}

func (s *UpdateIntentRequest) SetIntentId(v int64) *UpdateIntentRequest {
	s.IntentId = &v
	return s
}

type UpdateIntentRequestIntentDefinition struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// This parameter is required.
	IntentName *string                                         `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	SlotInfos  []*UpdateIntentRequestIntentDefinitionSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s UpdateIntentRequestIntentDefinition) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentRequestIntentDefinition) GoString() string {
	return s.String()
}

func (s *UpdateIntentRequestIntentDefinition) SetAliasName(v string) *UpdateIntentRequestIntentDefinition {
	s.AliasName = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinition) SetIntentName(v string) *UpdateIntentRequestIntentDefinition {
	s.IntentName = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinition) SetSlotInfos(v []*UpdateIntentRequestIntentDefinitionSlotInfos) *UpdateIntentRequestIntentDefinition {
	s.SlotInfos = v
	return s
}

type UpdateIntentRequestIntentDefinitionSlotInfos struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// example:
	//
	// false
	Interactive *bool `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dgadf23dfg2f
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIntentRequestIntentDefinitionSlotInfos) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentRequestIntentDefinitionSlotInfos) GoString() string {
	return s.String()
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetArray(v bool) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Array = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetEncrypt(v bool) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Encrypt = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetInteractive(v bool) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Interactive = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetName(v string) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Name = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetSlotId(v string) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.SlotId = &v
	return s
}

func (s *UpdateIntentRequestIntentDefinitionSlotInfos) SetValue(v string) *UpdateIntentRequestIntentDefinitionSlotInfos {
	s.Value = &v
	return s
}

type UpdateIntentShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentDefinitionShrink *string `json:"IntentDefinition,omitempty" xml:"IntentDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 234234234534
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s UpdateIntentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntentShrinkRequest) SetAgentKey(v string) *UpdateIntentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateIntentShrinkRequest) SetInstanceId(v string) *UpdateIntentShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIntentShrinkRequest) SetIntentDefinitionShrink(v string) *UpdateIntentShrinkRequest {
	s.IntentDefinitionShrink = &v
	return s
}

func (s *UpdateIntentShrinkRequest) SetIntentId(v int64) *UpdateIntentShrinkRequest {
	s.IntentId = &v
	return s
}

type UpdateIntentResponseBody struct {
	// example:
	//
	// 234234234534
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// dfaf23dfas234234234534
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIntentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntentResponseBody) SetIntentId(v int64) *UpdateIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *UpdateIntentResponseBody) SetRequestId(v string) *UpdateIntentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateIntentResponse) SetStatusCode(v int32) *UpdateIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntentResponse) SetBody(v *UpdateIntentResponseBody) *UpdateIntentResponse {
	s.Body = v
	return s
}

type UpdateLgfRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	LgfDefinition *UpdateLgfRequestLgfDefinition `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty" type:"Struct"`
	// LGF ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 12121
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
}

func (s UpdateLgfRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLgfRequest) GoString() string {
	return s.String()
}

func (s *UpdateLgfRequest) SetAgentKey(v string) *UpdateLgfRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateLgfRequest) SetInstanceId(v string) *UpdateLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLgfRequest) SetLgfDefinition(v *UpdateLgfRequestLgfDefinition) *UpdateLgfRequest {
	s.LgfDefinition = v
	return s
}

func (s *UpdateLgfRequest) SetLgfId(v int64) *UpdateLgfRequest {
	s.LgfId = &v
	return s
}

type UpdateLgfRequestLgfDefinition struct {
	// This parameter is required.
	//
	// example:
	//
	// 23234523522
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	RuleText *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s UpdateLgfRequestLgfDefinition) String() string {
	return tea.Prettify(s)
}

func (s UpdateLgfRequestLgfDefinition) GoString() string {
	return s.String()
}

func (s *UpdateLgfRequestLgfDefinition) SetIntentId(v int64) *UpdateLgfRequestLgfDefinition {
	s.IntentId = &v
	return s
}

func (s *UpdateLgfRequestLgfDefinition) SetRuleText(v string) *UpdateLgfRequestLgfDefinition {
	s.RuleText = &v
	return s
}

type UpdateLgfShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	LgfDefinitionShrink *string `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty"`
	// LGF ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 12121
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
}

func (s UpdateLgfShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLgfShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLgfShrinkRequest) SetAgentKey(v string) *UpdateLgfShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateLgfShrinkRequest) SetInstanceId(v string) *UpdateLgfShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLgfShrinkRequest) SetLgfDefinitionShrink(v string) *UpdateLgfShrinkRequest {
	s.LgfDefinitionShrink = &v
	return s
}

func (s *UpdateLgfShrinkRequest) SetLgfId(v int64) *UpdateLgfShrinkRequest {
	s.LgfId = &v
	return s
}

type UpdateLgfResponseBody struct {
	// example:
	//
	// 2342556223532
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// 289dfa131adf23wqe2r
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLgfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLgfResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLgfResponseBody) SetLgfId(v int64) *UpdateLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *UpdateLgfResponseBody) SetRequestId(v string) *UpdateLgfResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLgfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLgfResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLgfResponse) GoString() string {
	return s.String()
}

func (s *UpdateLgfResponse) SetHeaders(v map[string]*string) *UpdateLgfResponse {
	s.Headers = v
	return s
}

func (s *UpdateLgfResponse) SetStatusCode(v int32) *UpdateLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLgfResponse) SetBody(v *UpdateLgfResponseBody) *UpdateLgfResponse {
	s.Body = v
	return s
}

type UpdatePerspectiveRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 客户端视角
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s UpdatePerspectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveRequest) SetAgentKey(v string) *UpdatePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdatePerspectiveRequest) SetName(v string) *UpdatePerspectiveRequest {
	s.Name = &v
	return s
}

func (s *UpdatePerspectiveRequest) SetPerspectiveId(v string) *UpdatePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

type UpdatePerspectiveResponseBody struct {
	// example:
	//
	// FC384CE1-8D42-1900-84E1-F33F990F2B5E
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdatePerspectiveResponse) SetStatusCode(v int32) *UpdatePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePerspectiveResponse) SetBody(v *UpdatePerspectiveResponseBody) *UpdatePerspectiveResponse {
	s.Body = v
	return s
}

type UpdateSimQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000002788
	SimQuestionId *int64 `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateSimQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSimQuestionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSimQuestionRequest) SetAgentKey(v string) *UpdateSimQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateSimQuestionRequest) SetSimQuestionId(v int64) *UpdateSimQuestionRequest {
	s.SimQuestionId = &v
	return s
}

func (s *UpdateSimQuestionRequest) SetTitle(v string) *UpdateSimQuestionRequest {
	s.Title = &v
	return s
}

type UpdateSimQuestionResponseBody struct {
	// example:
	//
	// DFB71B34-4188-4EA2-9988-EF3014E75910
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSimQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSimQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSimQuestionResponseBody) SetRequestId(v string) *UpdateSimQuestionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSimQuestionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSimQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSimQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSimQuestionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSimQuestionResponse) SetHeaders(v map[string]*string) *UpdateSimQuestionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSimQuestionResponse) SetStatusCode(v int32) *UpdateSimQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSimQuestionResponse) SetBody(v *UpdateSimQuestionResponseBody) *UpdateSimQuestionResponse {
	s.Body = v
	return s
}

type UpdateSolutionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 100001333260
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s UpdateSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSolutionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSolutionRequest) SetAgentKey(v string) *UpdateSolutionRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateSolutionRequest) SetContent(v string) *UpdateSolutionRequest {
	s.Content = &v
	return s
}

func (s *UpdateSolutionRequest) SetContentType(v int32) *UpdateSolutionRequest {
	s.ContentType = &v
	return s
}

func (s *UpdateSolutionRequest) SetPerspectiveCodes(v []*string) *UpdateSolutionRequest {
	s.PerspectiveCodes = v
	return s
}

func (s *UpdateSolutionRequest) SetSolutionId(v int64) *UpdateSolutionRequest {
	s.SolutionId = &v
	return s
}

type UpdateSolutionResponseBody struct {
	// example:
	//
	// 8B8F098D-A338-54DD-B19C-24BBBCBD8498
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSolutionResponseBody) SetRequestId(v string) *UpdateSolutionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSolutionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSolutionResponse) SetHeaders(v map[string]*string) *UpdateSolutionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSolutionResponse) SetStatusCode(v int32) *UpdateSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSolutionResponse) SetBody(v *UpdateSolutionResponseBody) *UpdateSolutionResponse {
	s.Body = v
	return s
}

type UpdateUserSayRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId        *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserSayDefinition *UpdateUserSayRequestUserSayDefinition `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 34512323
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s UpdateUserSayRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserSayRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserSayRequest) SetAgentKey(v string) *UpdateUserSayRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateUserSayRequest) SetInstanceId(v string) *UpdateUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserSayRequest) SetUserSayDefinition(v *UpdateUserSayRequestUserSayDefinition) *UpdateUserSayRequest {
	s.UserSayDefinition = v
	return s
}

func (s *UpdateUserSayRequest) SetUserSayId(v int64) *UpdateUserSayRequest {
	s.UserSayId = &v
	return s
}

type UpdateUserSayRequestUserSayDefinition struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	IntentId  *int64                                            `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	SlotInfos []*UpdateUserSayRequestUserSayDefinitionSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s UpdateUserSayRequestUserSayDefinition) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserSayRequestUserSayDefinition) GoString() string {
	return s.String()
}

func (s *UpdateUserSayRequestUserSayDefinition) SetContent(v string) *UpdateUserSayRequestUserSayDefinition {
	s.Content = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinition) SetIntentId(v int64) *UpdateUserSayRequestUserSayDefinition {
	s.IntentId = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinition) SetSlotInfos(v []*UpdateUserSayRequestUserSayDefinitionSlotInfos) *UpdateUserSayRequestUserSayDefinition {
	s.SlotInfos = v
	return s
}

type UpdateUserSayRequestUserSayDefinitionSlotInfos struct {
	// example:
	//
	// 3
	EndIndex *int32 `json:"EndIndex,omitempty" xml:"EndIndex,omitempty"`
	// example:
	//
	// 346ffg3q23dv
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// example:
	//
	// 1
	StartIndex *int32 `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
}

func (s UpdateUserSayRequestUserSayDefinitionSlotInfos) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserSayRequestUserSayDefinitionSlotInfos) GoString() string {
	return s.String()
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) SetEndIndex(v int32) *UpdateUserSayRequestUserSayDefinitionSlotInfos {
	s.EndIndex = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) SetSlotId(v string) *UpdateUserSayRequestUserSayDefinitionSlotInfos {
	s.SlotId = &v
	return s
}

func (s *UpdateUserSayRequestUserSayDefinitionSlotInfos) SetStartIndex(v int32) *UpdateUserSayRequestUserSayDefinitionSlotInfos {
	s.StartIndex = &v
	return s
}

type UpdateUserSayShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserSayDefinitionShrink *string `json:"UserSayDefinition,omitempty" xml:"UserSayDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 34512323
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s UpdateUserSayShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserSayShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserSayShrinkRequest) SetAgentKey(v string) *UpdateUserSayShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateUserSayShrinkRequest) SetInstanceId(v string) *UpdateUserSayShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserSayShrinkRequest) SetUserSayDefinitionShrink(v string) *UpdateUserSayShrinkRequest {
	s.UserSayDefinitionShrink = &v
	return s
}

func (s *UpdateUserSayShrinkRequest) SetUserSayId(v int64) *UpdateUserSayShrinkRequest {
	s.UserSayId = &v
	return s
}

type UpdateUserSayResponseBody struct {
	// example:
	//
	// 2356fg3wf34634vdt23wef2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 34512323
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s UpdateUserSayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserSayResponseBody) SetRequestId(v string) *UpdateUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserSayResponseBody) SetUserSayId(v int64) *UpdateUserSayResponseBody {
	s.UserSayId = &v
	return s
}

type UpdateUserSayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserSayResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserSayResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserSayResponse) SetHeaders(v map[string]*string) *UpdateUserSayResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserSayResponse) SetStatusCode(v int32) *UpdateUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserSayResponse) SetBody(v *UpdateUserSayResponseBody) *UpdateUserSayResponse {
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

// Summary:
//
// 申请流式网关AccessToken
//
// @param request - ApplyForStreamAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyForStreamAccessTokenResponse
func (client *Client) ApplyForStreamAccessTokenWithOptions(request *ApplyForStreamAccessTokenRequest, runtime *util.RuntimeOptions) (_result *ApplyForStreamAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyForStreamAccessToken"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyForStreamAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请流式网关AccessToken
//
// @param request - ApplyForStreamAccessTokenRequest
//
// @return ApplyForStreamAccessTokenResponse
func (client *Client) ApplyForStreamAccessToken(request *ApplyForStreamAccessTokenRequest) (_result *ApplyForStreamAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyForStreamAccessTokenResponse{}
	_body, _err := client.ApplyForStreamAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话-联想API
//
// @param tmpReq - AssociateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateResponse
func (client *Client) AssociateWithOptions(tmpReq *AssociateRequest, runtime *util.RuntimeOptions) (_result *AssociateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AssociateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Perspective)) {
		request.PerspectiveShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Perspective, tea.String("Perspective"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveShrink)) {
		query["Perspective"] = request.PerspectiveShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RecommendNum)) {
		query["RecommendNum"] = request.RecommendNum
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		query["Utterance"] = request.Utterance
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Associate"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话-联想API
//
// @param request - AssociateRequest
//
// @return AssociateResponse
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

// Summary:
//
// 获取欢迎语
//
// @param request - BeginSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeginSessionResponse
func (client *Client) BeginSessionWithOptions(request *BeginSessionRequest, runtime *util.RuntimeOptions) (_result *BeginSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BeginSession"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BeginSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取欢迎语
//
// @param request - BeginSessionRequest
//
// @return BeginSessionResponse
func (client *Client) BeginSession(request *BeginSessionRequest) (_result *BeginSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BeginSessionResponse{}
	_body, _err := client.BeginSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消机器人发布
//
// @param request - CancelInstancePublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelInstancePublishTaskResponse
func (client *Client) CancelInstancePublishTaskWithOptions(request *CancelInstancePublishTaskRequest, runtime *util.RuntimeOptions) (_result *CancelInstancePublishTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelInstancePublishTask"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelInstancePublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消机器人发布
//
// @param request - CancelInstancePublishTaskRequest
//
// @return CancelInstancePublishTaskResponse
func (client *Client) CancelInstancePublishTask(request *CancelInstancePublishTaskRequest) (_result *CancelInstancePublishTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelInstancePublishTaskResponse{}
	_body, _err := client.CancelInstancePublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消发布任务
//
// @param request - CancelPublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPublishTaskResponse
func (client *Client) CancelPublishTaskWithOptions(request *CancelPublishTaskRequest, runtime *util.RuntimeOptions) (_result *CancelPublishTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelPublishTask"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelPublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消发布任务
//
// @param request - CancelPublishTaskRequest
//
// @return CancelPublishTaskResponse
func (client *Client) CancelPublishTask(request *CancelPublishTaskRequest) (_result *CancelPublishTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelPublishTaskResponse{}
	_body, _err := client.CancelPublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话API
//
// @param tmpReq - ChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatResponse
func (client *Client) ChatWithOptions(tmpReq *ChatRequest, runtime *util.RuntimeOptions) (_result *ChatResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Perspective)) {
		request.PerspectiveShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Perspective, tea.String("Perspective"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentName)) {
		query["IntentName"] = request.IntentName
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveShrink)) {
		query["Perspective"] = request.PerspectiveShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SandBox)) {
		query["SandBox"] = request.SandBox
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		query["SenderId"] = request.SenderId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderNick)) {
		query["SenderNick"] = request.SenderNick
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		query["Utterance"] = request.Utterance
	}

	if !tea.BoolValue(util.IsUnset(request.VendorParam)) {
		query["VendorParam"] = request.VendorParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Chat"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话API
//
// @param request - ChatRequest
//
// @return ChatResponse
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

// Summary:
//
// 继续机器人发布
//
// @param request - ContinueInstancePublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueInstancePublishTaskResponse
func (client *Client) ContinueInstancePublishTaskWithOptions(request *ContinueInstancePublishTaskRequest, runtime *util.RuntimeOptions) (_result *ContinueInstancePublishTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinueInstancePublishTask"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinueInstancePublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 继续机器人发布
//
// @param request - ContinueInstancePublishTaskRequest
//
// @return ContinueInstancePublishTaskResponse
func (client *Client) ContinueInstancePublishTask(request *ContinueInstancePublishTaskRequest) (_result *ContinueInstancePublishTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinueInstancePublishTaskResponse{}
	_body, _err := client.ContinueInstancePublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增类目
//
// @param request - CreateCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCategoryResponse
func (client *Client) CreateCategoryWithOptions(request *CreateCategoryRequest, runtime *util.RuntimeOptions) (_result *CreateCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeType)) {
		body["KnowledgeType"] = request.KnowledgeType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategoryId)) {
		body["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCategory"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增类目
//
// @param request - CreateCategoryRequest
//
// @return CreateCategoryResponse
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

// Summary:
//
// 新建FAQ关联问
//
// @param request - CreateConnQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConnQuestionResponse
func (client *Client) CreateConnQuestionWithOptions(request *CreateConnQuestionRequest, runtime *util.RuntimeOptions) (_result *CreateConnQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnQuestionId)) {
		body["ConnQuestionId"] = request.ConnQuestionId
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConnQuestion"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConnQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建FAQ关联问
//
// @param request - CreateConnQuestionRequest
//
// @return CreateConnQuestionResponse
func (client *Client) CreateConnQuestion(request *CreateConnQuestionRequest) (_result *CreateConnQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConnQuestionResponse{}
	_body, _err := client.CreateConnQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-创建
//
// @param request - CreateDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDSEntityResponse
func (client *Client) CreateDSEntityWithOptions(request *CreateDSEntityRequest, runtime *util.RuntimeOptions) (_result *CreateDSEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		query["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDSEntity"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-创建
//
// @param request - CreateDSEntityRequest
//
// @return CreateDSEntityResponse
func (client *Client) CreateDSEntity(request *CreateDSEntityRequest) (_result *CreateDSEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDSEntityResponse{}
	_body, _err := client.CreateDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体成员-创建
//
// @param tmpReq - CreateDSEntityValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDSEntityValueResponse
func (client *Client) CreateDSEntityValueWithOptions(tmpReq *CreateDSEntityValueRequest, runtime *util.RuntimeOptions) (_result *CreateDSEntityValueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDSEntityValueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Synonyms)) {
		request.SynonymsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Synonyms, tea.String("Synonyms"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SynonymsShrink)) {
		body["Synonyms"] = request.SynonymsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDSEntityValue"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDSEntityValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体成员-创建
//
// @param request - CreateDSEntityValueRequest
//
// @return CreateDSEntityValueResponse
func (client *Client) CreateDSEntityValue(request *CreateDSEntityValueRequest) (_result *CreateDSEntityValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDSEntityValueResponse{}
	_body, _err := client.CreateDSEntityValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文档
//
// @param tmpReq - CreateDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocResponse
func (client *Client) CreateDocWithOptions(tmpReq *CreateDocRequest, runtime *util.RuntimeOptions) (_result *CreateDocResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocMetadata)) {
		request.DocMetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocMetadata, tea.String("DocMetadata"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DocMetadataShrink)) {
		query["DocMetadata"] = request.DocMetadataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Meta)) {
		query["Meta"] = request.Meta
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		query["TagIds"] = request.TagIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDoc"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文档
//
// @param request - CreateDocRequest
//
// @return CreateDocResponse
func (client *Client) CreateDoc(request *CreateDocRequest) (_result *CreateDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDocResponse{}
	_body, _err := client.CreateDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建FAQ
//
// @param request - CreateFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFaqResponse
func (client *Client) CreateFaqWithOptions(request *CreateFaqRequest, runtime *util.RuntimeOptions) (_result *CreateFaqResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionContent)) {
		body["SolutionContent"] = request.SolutionContent
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionType)) {
		body["SolutionType"] = request.SolutionType
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFaq"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建FAQ
//
// @param request - CreateFaqRequest
//
// @return CreateFaqResponse
func (client *Client) CreateFaq(request *CreateFaqRequest) (_result *CreateFaqResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFaqResponse{}
	_body, _err := client.CreateFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-创建
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		query["Introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		query["LanguageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RobotType)) {
		query["RobotType"] = request.RobotType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-创建
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建机器人发布任务
//
// @param request - CreateInstancePublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstancePublishTaskResponse
func (client *Client) CreateInstancePublishTaskWithOptions(request *CreateInstancePublishTaskRequest, runtime *util.RuntimeOptions) (_result *CreateInstancePublishTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstancePublishTask"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstancePublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建机器人发布任务
//
// @param request - CreateInstancePublishTaskRequest
//
// @return CreateInstancePublishTaskResponse
func (client *Client) CreateInstancePublishTask(request *CreateInstancePublishTaskRequest) (_result *CreateInstancePublishTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstancePublishTaskResponse{}
	_body, _err := client.CreateInstancePublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-创建
//
// @param tmpReq - CreateIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntentResponse
func (client *Client) CreateIntentWithOptions(tmpReq *CreateIntentRequest, runtime *util.RuntimeOptions) (_result *CreateIntentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateIntentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IntentDefinition)) {
		request.IntentDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentDefinition, tea.String("IntentDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentDefinitionShrink)) {
		query["IntentDefinition"] = request.IntentDefinitionShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIntent"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-创建
//
// @param request - CreateIntentRequest
//
// @return CreateIntentResponse
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

// Summary:
//
// 意图-LGF-创建
//
// @param tmpReq - CreateLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLgfResponse
func (client *Client) CreateLgfWithOptions(tmpReq *CreateLgfRequest, runtime *util.RuntimeOptions) (_result *CreateLgfResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateLgfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LgfDefinition)) {
		request.LgfDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LgfDefinition, tea.String("LgfDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LgfDefinitionShrink)) {
		query["LgfDefinition"] = request.LgfDefinitionShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLgf"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLgfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-LGF-创建
//
// @param request - CreateLgfRequest
//
// @return CreateLgfResponse
func (client *Client) CreateLgf(request *CreateLgfRequest) (_result *CreateLgfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLgfResponse{}
	_body, _err := client.CreateLgfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-创建
//
// @param request - CreatePerspectiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePerspectiveResponse
func (client *Client) CreatePerspectiveWithOptions(request *CreatePerspectiveRequest, runtime *util.RuntimeOptions) (_result *CreatePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePerspective"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-创建
//
// @param request - CreatePerspectiveRequest
//
// @return CreatePerspectiveResponse
func (client *Client) CreatePerspective(request *CreatePerspectiveRequest) (_result *CreatePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePerspectiveResponse{}
	_body, _err := client.CreatePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建发布任务
//
// @param tmpReq - CreatePublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePublishTaskResponse
func (client *Client) CreatePublishTaskWithOptions(tmpReq *CreatePublishTaskRequest, runtime *util.RuntimeOptions) (_result *CreatePublishTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePublishTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DataIdList)) {
		request.DataIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataIdList, tea.String("DataIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.DataIdListShrink)) {
		query["DataIdList"] = request.DataIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePublishTask"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建发布任务
//
// @param request - CreatePublishTaskRequest
//
// @return CreatePublishTaskResponse
func (client *Client) CreatePublishTask(request *CreatePublishTaskRequest) (_result *CreatePublishTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePublishTaskResponse{}
	_body, _err := client.CreatePublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建FAQ相似问
//
// @param request - CreateSimQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimQuestionResponse
func (client *Client) CreateSimQuestionWithOptions(request *CreateSimQuestionRequest, runtime *util.RuntimeOptions) (_result *CreateSimQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSimQuestion"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSimQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建FAQ相似问
//
// @param request - CreateSimQuestionRequest
//
// @return CreateSimQuestionResponse
func (client *Client) CreateSimQuestion(request *CreateSimQuestionRequest) (_result *CreateSimQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSimQuestionResponse{}
	_body, _err := client.CreateSimQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建FAQ答案
//
// @param request - CreateSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSolutionResponse
func (client *Client) CreateSolutionWithOptions(request *CreateSolutionRequest, runtime *util.RuntimeOptions) (_result *CreateSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		query["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveCodes)) {
		query["PerspectiveCodes"] = request.PerspectiveCodes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSolution"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建FAQ答案
//
// @param request - CreateSolutionRequest
//
// @return CreateSolutionResponse
func (client *Client) CreateSolution(request *CreateSolutionRequest) (_result *CreateSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSolutionResponse{}
	_body, _err := client.CreateSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-话术-创建
//
// @param tmpReq - CreateUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserSayResponse
func (client *Client) CreateUserSayWithOptions(tmpReq *CreateUserSayRequest, runtime *util.RuntimeOptions) (_result *CreateUserSayResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateUserSayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserSayDefinition)) {
		request.UserSayDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserSayDefinition, tea.String("UserSayDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSayDefinitionShrink)) {
		query["UserSayDefinition"] = request.UserSayDefinitionShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserSay"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserSayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-话术-创建
//
// @param request - CreateUserSayRequest
//
// @return CreateUserSayResponse
func (client *Client) CreateUserSay(request *CreateUserSayRequest) (_result *CreateUserSayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserSayResponse{}
	_body, _err := client.CreateUserSayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除类目
//
// @param request - DeleteCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategoryWithOptions(request *DeleteCategoryRequest, runtime *util.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCategory"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除类目
//
// @param request - DeleteCategoryRequest
//
// @return DeleteCategoryResponse
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

// Summary:
//
// 删除FAQ关联问
//
// @param request - DeleteConnQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnQuestionResponse
func (client *Client) DeleteConnQuestionWithOptions(request *DeleteConnQuestionRequest, runtime *util.RuntimeOptions) (_result *DeleteConnQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutlineId)) {
		body["OutlineId"] = request.OutlineId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConnQuestion"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConnQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除FAQ关联问
//
// @param request - DeleteConnQuestionRequest
//
// @return DeleteConnQuestionResponse
func (client *Client) DeleteConnQuestion(request *DeleteConnQuestionRequest) (_result *DeleteConnQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConnQuestionResponse{}
	_body, _err := client.DeleteConnQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-删除
//
// @param request - DeleteDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDSEntityResponse
func (client *Client) DeleteDSEntityWithOptions(request *DeleteDSEntityRequest, runtime *util.RuntimeOptions) (_result *DeleteDSEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDSEntity"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-删除
//
// @param request - DeleteDSEntityRequest
//
// @return DeleteDSEntityResponse
func (client *Client) DeleteDSEntity(request *DeleteDSEntityRequest) (_result *DeleteDSEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDSEntityResponse{}
	_body, _err := client.DeleteDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体成员-删除
//
// @param request - DeleteDSEntityValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDSEntityValueResponse
func (client *Client) DeleteDSEntityValueWithOptions(request *DeleteDSEntityValueRequest, runtime *util.RuntimeOptions) (_result *DeleteDSEntityValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityValueId)) {
		query["EntityValueId"] = request.EntityValueId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDSEntityValue"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDSEntityValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体成员-删除
//
// @param request - DeleteDSEntityValueRequest
//
// @return DeleteDSEntityValueResponse
func (client *Client) DeleteDSEntityValue(request *DeleteDSEntityValueRequest) (_result *DeleteDSEntityValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDSEntityValueResponse{}
	_body, _err := client.DeleteDSEntityValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档删除
//
// @param request - DeleteDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocResponse
func (client *Client) DeleteDocWithOptions(request *DeleteDocRequest, runtime *util.RuntimeOptions) (_result *DeleteDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDoc"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档删除
//
// @param request - DeleteDocRequest
//
// @return DeleteDocResponse
func (client *Client) DeleteDoc(request *DeleteDocRequest) (_result *DeleteDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDocResponse{}
	_body, _err := client.DeleteDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除FAQ，如果是已发布的知识，删除之后，变成已删除未发布，需要发布才能真正删除
//
// @param request - DeleteFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaqResponse
func (client *Client) DeleteFaqWithOptions(request *DeleteFaqRequest, runtime *util.RuntimeOptions) (_result *DeleteFaqResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaq"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除FAQ，如果是已发布的知识，删除之后，变成已删除未发布，需要发布才能真正删除
//
// @param request - DeleteFaqRequest
//
// @return DeleteFaqResponse
func (client *Client) DeleteFaq(request *DeleteFaqRequest) (_result *DeleteFaqResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaqResponse{}
	_body, _err := client.DeleteFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-删除
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-删除
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-删除
//
// @param request - DeleteIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIntentResponse
func (client *Client) DeleteIntentWithOptions(request *DeleteIntentRequest, runtime *util.RuntimeOptions) (_result *DeleteIntentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIntent"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-删除
//
// @param request - DeleteIntentRequest
//
// @return DeleteIntentResponse
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

// Summary:
//
// 意图-LGF-删除
//
// @param request - DeleteLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLgfResponse
func (client *Client) DeleteLgfWithOptions(request *DeleteLgfRequest, runtime *util.RuntimeOptions) (_result *DeleteLgfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	if !tea.BoolValue(util.IsUnset(request.LgfId)) {
		query["LgfId"] = request.LgfId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLgf"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLgfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-LGF-删除
//
// @param request - DeleteLgfRequest
//
// @return DeleteLgfResponse
func (client *Client) DeleteLgf(request *DeleteLgfRequest) (_result *DeleteLgfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLgfResponse{}
	_body, _err := client.DeleteLgfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-删除
//
// @param request - DeletePerspectiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePerspectiveResponse
func (client *Client) DeletePerspectiveWithOptions(request *DeletePerspectiveRequest, runtime *util.RuntimeOptions) (_result *DeletePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveId)) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePerspective"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-删除
//
// @param request - DeletePerspectiveRequest
//
// @return DeletePerspectiveResponse
func (client *Client) DeletePerspective(request *DeletePerspectiveRequest) (_result *DeletePerspectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePerspectiveResponse{}
	_body, _err := client.DeletePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除FAQ相似问
//
// @param request - DeleteSimQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSimQuestionResponse
func (client *Client) DeleteSimQuestionWithOptions(request *DeleteSimQuestionRequest, runtime *util.RuntimeOptions) (_result *DeleteSimQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SimQuestionId)) {
		body["SimQuestionId"] = request.SimQuestionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSimQuestion"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSimQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除FAQ相似问
//
// @param request - DeleteSimQuestionRequest
//
// @return DeleteSimQuestionResponse
func (client *Client) DeleteSimQuestion(request *DeleteSimQuestionRequest) (_result *DeleteSimQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSimQuestionResponse{}
	_body, _err := client.DeleteSimQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除FAQ答案
//
// @param request - DeleteSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSolutionResponse
func (client *Client) DeleteSolutionWithOptions(request *DeleteSolutionRequest, runtime *util.RuntimeOptions) (_result *DeleteSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		body["SolutionId"] = request.SolutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSolution"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除FAQ答案
//
// @param request - DeleteSolutionRequest
//
// @return DeleteSolutionResponse
func (client *Client) DeleteSolution(request *DeleteSolutionRequest) (_result *DeleteSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSolutionResponse{}
	_body, _err := client.DeleteSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-用户话术-删除
//
// @param request - DeleteUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserSayResponse
func (client *Client) DeleteUserSayWithOptions(request *DeleteUserSayRequest, runtime *util.RuntimeOptions) (_result *DeleteUserSayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSayId)) {
		query["UserSayId"] = request.UserSayId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserSay"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserSayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-用户话术-删除
//
// @param request - DeleteUserSayRequest
//
// @return DeleteUserSayResponse
func (client *Client) DeleteUserSay(request *DeleteUserSayRequest) (_result *DeleteUserSayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserSayResponse{}
	_body, _err := client.DeleteUserSayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看单个类目信息
//
// @param request - DescribeCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCategoryResponse
func (client *Client) DescribeCategoryWithOptions(request *DescribeCategoryRequest, runtime *util.RuntimeOptions) (_result *DescribeCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCategory"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看单个类目信息
//
// @param request - DescribeCategoryRequest
//
// @return DescribeCategoryResponse
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

// Summary:
//
// 实体-详情
//
// @param request - DescribeDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDSEntityResponse
func (client *Client) DescribeDSEntityWithOptions(request *DescribeDSEntityRequest, runtime *util.RuntimeOptions) (_result *DescribeDSEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDSEntity"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-详情
//
// @param request - DescribeDSEntityRequest
//
// @return DescribeDSEntityResponse
func (client *Client) DescribeDSEntity(request *DescribeDSEntityRequest) (_result *DescribeDSEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDSEntityResponse{}
	_body, _err := client.DescribeDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档详情
//
// @param request - DescribeDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocResponse
func (client *Client) DescribeDocWithOptions(request *DescribeDocRequest, runtime *util.RuntimeOptions) (_result *DescribeDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDetail)) {
		query["ShowDetail"] = request.ShowDetail
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDoc"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档详情
//
// @param request - DescribeDocRequest
//
// @return DescribeDocResponse
func (client *Client) DescribeDoc(request *DescribeDocRequest) (_result *DescribeDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDocResponse{}
	_body, _err := client.DescribeDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 知识详情
//
// @param request - DescribeFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaqResponse
func (client *Client) DescribeFaqWithOptions(request *DescribeFaqRequest, runtime *util.RuntimeOptions) (_result *DescribeFaqResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaq"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 知识详情
//
// @param request - DescribeFaqRequest
//
// @return DescribeFaqResponse
func (client *Client) DescribeFaq(request *DescribeFaqRequest) (_result *DescribeFaqResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaqResponse{}
	_body, _err := client.DescribeFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-详情
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-详情
//
// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-详情
//
// @param request - DescribeIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIntentResponse
func (client *Client) DescribeIntentWithOptions(request *DescribeIntentRequest, runtime *util.RuntimeOptions) (_result *DescribeIntentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		body["IntentId"] = request.IntentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIntent"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-详情
//
// @param request - DescribeIntentRequest
//
// @return DescribeIntentResponse
func (client *Client) DescribeIntent(request *DescribeIntentRequest) (_result *DescribeIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIntentResponse{}
	_body, _err := client.DescribeIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-详情
//
// @param request - DescribePerspectiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePerspectiveResponse
func (client *Client) DescribePerspectiveWithOptions(request *DescribePerspectiveRequest, runtime *util.RuntimeOptions) (_result *DescribePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveId)) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePerspective"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-详情
//
// @param request - DescribePerspectiveRequest
//
// @return DescribePerspectiveResponse
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

// Summary:
//
// 问答点赞、点踩API
//
// @param request - FeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FeedbackResponse
func (client *Client) FeedbackWithOptions(request *FeedbackRequest, runtime *util.RuntimeOptions) (_result *FeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		query["Feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.FeedbackContent)) {
		query["FeedbackContent"] = request.FeedbackContent
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Feedback"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 问答点赞、点踩API
//
// @param request - FeedbackRequest
//
// @return FeedbackResponse
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

// Summary:
//
// 生成用户免登Token
//
// @param request - GenerateUserAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUserAccessTokenResponse
func (client *Client) GenerateUserAccessTokenWithOptions(request *GenerateUserAccessTokenRequest, runtime *util.RuntimeOptions) (_result *GenerateUserAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		query["ForeignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.Nick)) {
		query["Nick"] = request.Nick
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		query["Telephone"] = request.Telephone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUserAccessToken"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUserAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成用户免登Token
//
// @param request - GenerateUserAccessTokenRequest
//
// @return GenerateUserAccessTokenResponse
func (client *Client) GenerateUserAccessToken(request *GenerateUserAccessTokenRequest) (_result *GenerateUserAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUserAccessTokenResponse{}
	_body, _err := client.GenerateUserAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取业务空间信息
//
// @param request - GetAgentInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentInfoResponse
func (client *Client) GetAgentInfoWithOptions(request *GetAgentInfoRequest, runtime *util.RuntimeOptions) (_result *GetAgentInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgentInfo"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务空间信息
//
// @param request - GetAgentInfoRequest
//
// @return GetAgentInfoResponse
func (client *Client) GetAgentInfo(request *GetAgentInfoRequest) (_result *GetAgentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentInfoResponse{}
	_body, _err := client.GetAgentInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步函数执行结果接口
//
// @param request - GetAsyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncResultResponse
func (client *Client) GetAsyncResultWithOptions(request *GetAsyncResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncResult"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步函数执行结果接口
//
// @param request - GetAsyncResultRequest
//
// @return GetAsyncResultResponse
func (client *Client) GetAsyncResult(request *GetAsyncResultRequest) (_result *GetAsyncResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncResultResponse{}
	_body, _err := client.GetAsyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询机器人接待人次和对话轮次
//
// @param request - GetBotSessionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBotSessionDataResponse
func (client *Client) GetBotSessionDataWithOptions(request *GetBotSessionDataRequest, runtime *util.RuntimeOptions) (_result *GetBotSessionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBotSessionData"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBotSessionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机器人接待人次和对话轮次
//
// @param request - GetBotSessionDataRequest
//
// @return GetBotSessionDataResponse
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

// Summary:
//
// 查询机器人发布进度
//
// @param request - GetInstancePublishTaskStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstancePublishTaskStateResponse
func (client *Client) GetInstancePublishTaskStateWithOptions(request *GetInstancePublishTaskStateRequest, runtime *util.RuntimeOptions) (_result *GetInstancePublishTaskStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstancePublishTaskState"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstancePublishTaskStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机器人发布进度
//
// @param request - GetInstancePublishTaskStateRequest
//
// @return GetInstancePublishTaskStateResponse
func (client *Client) GetInstancePublishTaskState(request *GetInstancePublishTaskStateRequest) (_result *GetInstancePublishTaskStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstancePublishTaskStateResponse{}
	_body, _err := client.GetInstancePublishTaskStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询发布进度
//
// @param request - GetPublishTaskStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublishTaskStateResponse
func (client *Client) GetPublishTaskStateWithOptions(request *GetPublishTaskStateRequest, runtime *util.RuntimeOptions) (_result *GetPublishTaskStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublishTaskState"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublishTaskStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询发布进度
//
// @param request - GetPublishTaskStateRequest
//
// @return GetPublishTaskStateResponse
func (client *Client) GetPublishTaskState(request *GetPublishTaskStateRequest) (_result *GetPublishTaskStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublishTaskStateResponse{}
	_body, _err := client.GetPublishTaskStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 初始化im连接信息
//
// @param request - InitIMConnectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitIMConnectResponse
func (client *Client) InitIMConnectWithOptions(request *InitIMConnectRequest, runtime *util.RuntimeOptions) (_result *InitIMConnectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessToken)) {
		query["UserAccessToken"] = request.UserAccessToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InitIMConnect"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitIMConnectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 初始化im连接信息
//
// @param request - InitIMConnectRequest
//
// @return InitIMConnectResponse
func (client *Client) InitIMConnect(request *InitIMConnectRequest) (_result *InitIMConnectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitIMConnectResponse{}
	_body, _err := client.InitIMConnectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-绑定类目
//
// @param request - LinkInstanceCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LinkInstanceCategoryResponse
func (client *Client) LinkInstanceCategoryWithOptions(request *LinkInstanceCategoryRequest, runtime *util.RuntimeOptions) (_result *LinkInstanceCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryIds)) {
		body["CategoryIds"] = request.CategoryIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LinkInstanceCategory"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LinkInstanceCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-绑定类目
//
// @param request - LinkInstanceCategoryRequest
//
// @return LinkInstanceCategoryResponse
func (client *Client) LinkInstanceCategory(request *LinkInstanceCategoryRequest) (_result *LinkInstanceCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LinkInstanceCategoryResponse{}
	_body, _err := client.LinkInstanceCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取业务空间列表
//
// @param request - ListAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentResponse
func (client *Client) ListAgentWithOptions(request *ListAgentRequest, runtime *util.RuntimeOptions) (_result *ListAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentName)) {
		query["AgentName"] = request.AgentName
	}

	if !tea.BoolValue(util.IsUnset(request.GoodsCodes)) {
		query["GoodsCodes"] = request.GoodsCodes
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAgent"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务空间列表
//
// @param request - ListAgentRequest
//
// @return ListAgentResponse
func (client *Client) ListAgent(request *ListAgentRequest) (_result *ListAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAgentResponse{}
	_body, _err := client.ListAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 类目列表
//
// @param request - ListCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoryResponse
func (client *Client) ListCategoryWithOptions(request *ListCategoryRequest, runtime *util.RuntimeOptions) (_result *ListCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeType)) {
		body["KnowledgeType"] = request.KnowledgeType
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategoryId)) {
		body["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCategory"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 类目列表
//
// @param request - ListCategoryRequest
//
// @return ListCategoryResponse
func (client *Client) ListCategory(request *ListCategoryRequest) (_result *ListCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCategoryResponse{}
	_body, _err := client.ListCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询FAQ关联问列表
//
// @param request - ListConnQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnQuestionResponse
func (client *Client) ListConnQuestionWithOptions(request *ListConnQuestionRequest, runtime *util.RuntimeOptions) (_result *ListConnQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnQuestion"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询FAQ关联问列表
//
// @param request - ListConnQuestionRequest
//
// @return ListConnQuestionResponse
func (client *Client) ListConnQuestion(request *ListConnQuestionRequest) (_result *ListConnQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnQuestionResponse{}
	_body, _err := client.ListConnQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-列表
//
// @param request - ListDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDSEntityResponse
func (client *Client) ListDSEntityWithOptions(request *ListDSEntityRequest, runtime *util.RuntimeOptions) (_result *ListDSEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDSEntity"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-列表
//
// @param request - ListDSEntityRequest
//
// @return ListDSEntityResponse
func (client *Client) ListDSEntity(request *ListDSEntityRequest) (_result *ListDSEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDSEntityResponse{}
	_body, _err := client.ListDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体成员-列表
//
// @param request - ListDSEntityValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDSEntityValueResponse
func (client *Client) ListDSEntityValueWithOptions(request *ListDSEntityValueRequest, runtime *util.RuntimeOptions) (_result *ListDSEntityValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityValueId)) {
		body["EntityValueId"] = request.EntityValueId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDSEntityValue"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDSEntityValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体成员-列表
//
// @param request - ListDSEntityValueRequest
//
// @return ListDSEntityValueResponse
func (client *Client) ListDSEntityValue(request *ListDSEntityValueRequest) (_result *ListDSEntityValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDSEntityValueResponse{}
	_body, _err := client.ListDSEntityValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-修改
//
// @param request - ListInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, runtime *util.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RobotType)) {
		query["RobotType"] = request.RobotType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstance"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-修改
//
// @param request - ListInstanceRequest
//
// @return ListInstanceResponse
func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-列表
//
// @param request - ListIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntentResponse
func (client *Client) ListIntentWithOptions(request *ListIntentRequest, runtime *util.RuntimeOptions) (_result *ListIntentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentName)) {
		query["IntentName"] = request.IntentName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIntent"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-列表
//
// @param request - ListIntentRequest
//
// @return ListIntentResponse
func (client *Client) ListIntent(request *ListIntentRequest) (_result *ListIntentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntentResponse{}
	_body, _err := client.ListIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-LGF-列表
//
// @param request - ListLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLgfResponse
func (client *Client) ListLgfWithOptions(request *ListLgfRequest, runtime *util.RuntimeOptions) (_result *ListLgfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	if !tea.BoolValue(util.IsUnset(request.LgfText)) {
		query["LgfText"] = request.LgfText
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLgf"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLgfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-LGF-列表
//
// @param request - ListLgfRequest
//
// @return ListLgfResponse
func (client *Client) ListLgf(request *ListLgfRequest) (_result *ListLgfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLgfResponse{}
	_body, _err := client.ListLgfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取业务空间下可集成的SaaS信息列表
//
// @param request - ListSaasInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSaasInfoResponse
func (client *Client) ListSaasInfoWithOptions(request *ListSaasInfoRequest, runtime *util.RuntimeOptions) (_result *ListSaasInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.SaasGroupCodes)) {
		query["SaasGroupCodes"] = request.SaasGroupCodes
	}

	if !tea.BoolValue(util.IsUnset(request.SaasName)) {
		query["SaasName"] = request.SaasName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSaasInfo"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSaasInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务空间下可集成的SaaS信息列表
//
// @param request - ListSaasInfoRequest
//
// @return ListSaasInfoResponse
func (client *Client) ListSaasInfo(request *ListSaasInfoRequest) (_result *ListSaasInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSaasInfoResponse{}
	_body, _err := client.ListSaasInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取业务空间下可集成的权限组信息
//
// @param request - ListSaasPermissionGroupInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSaasPermissionGroupInfosResponse
func (client *Client) ListSaasPermissionGroupInfosWithOptions(request *ListSaasPermissionGroupInfosRequest, runtime *util.RuntimeOptions) (_result *ListSaasPermissionGroupInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSaasPermissionGroupInfos"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSaasPermissionGroupInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务空间下可集成的权限组信息
//
// @param request - ListSaasPermissionGroupInfosRequest
//
// @return ListSaasPermissionGroupInfosResponse
func (client *Client) ListSaasPermissionGroupInfos(request *ListSaasPermissionGroupInfosRequest) (_result *ListSaasPermissionGroupInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSaasPermissionGroupInfosResponse{}
	_body, _err := client.ListSaasPermissionGroupInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// FAQ相似问列表
//
// @param request - ListSimQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSimQuestionResponse
func (client *Client) ListSimQuestionWithOptions(request *ListSimQuestionRequest, runtime *util.RuntimeOptions) (_result *ListSimQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSimQuestion"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSimQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// FAQ相似问列表
//
// @param request - ListSimQuestionRequest
//
// @return ListSimQuestionResponse
func (client *Client) ListSimQuestion(request *ListSimQuestionRequest) (_result *ListSimQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSimQuestionResponse{}
	_body, _err := client.ListSimQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// FAQ答案列表
//
// @param request - ListSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSolutionResponse
func (client *Client) ListSolutionWithOptions(request *ListSolutionRequest, runtime *util.RuntimeOptions) (_result *ListSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSolution"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// FAQ答案列表
//
// @param request - ListSolutionRequest
//
// @return ListSolutionResponse
func (client *Client) ListSolution(request *ListSolutionRequest) (_result *ListSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSolutionResponse{}
	_body, _err := client.ListSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tongyi对话明细查询接口
//
// @param request - ListTongyiChatHistorysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTongyiChatHistorysResponse
func (client *Client) ListTongyiChatHistorysWithOptions(request *ListTongyiChatHistorysRequest, runtime *util.RuntimeOptions) (_result *ListTongyiChatHistorysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.RobotInstanceId)) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTongyiChatHistorys"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTongyiChatHistorysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tongyi对话明细查询接口
//
// @param request - ListTongyiChatHistorysRequest
//
// @return ListTongyiChatHistorysResponse
func (client *Client) ListTongyiChatHistorys(request *ListTongyiChatHistorysRequest) (_result *ListTongyiChatHistorysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTongyiChatHistorysResponse{}
	_body, _err := client.ListTongyiChatHistorysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 话术-列表
//
// @param request - ListUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserSayResponse
func (client *Client) ListUserSayWithOptions(request *ListUserSayRequest, runtime *util.RuntimeOptions) (_result *ListUserSayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserSay"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserSayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 话术-列表
//
// @param request - ListUserSayRequest
//
// @return ListUserSayResponse
func (client *Client) ListUserSay(request *ListUserSayRequest) (_result *ListUserSayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserSayResponse{}
	_body, _err := client.ListUserSayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 统一NLU接口
//
// @param request - NluRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NluResponse
func (client *Client) NluWithOptions(request *NluRequest, runtime *util.RuntimeOptions) (_result *NluResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		query["Utterance"] = request.Utterance
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Nlu"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NluResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 统一NLU接口
//
// @param request - NluRequest
//
// @return NluResponse
func (client *Client) Nlu(request *NluRequest) (_result *NluResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NluResponse{}
	_body, _err := client.NluWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-列表
//
// @param request - QueryPerspectivesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPerspectivesResponse
func (client *Client) QueryPerspectivesWithOptions(request *QueryPerspectivesRequest, runtime *util.RuntimeOptions) (_result *QueryPerspectivesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPerspectives"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPerspectivesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-列表
//
// @param request - QueryPerspectivesRequest
//
// @return QueryPerspectivesResponse
func (client *Client) QueryPerspectives(request *QueryPerspectivesRequest) (_result *QueryPerspectivesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPerspectivesResponse{}
	_body, _err := client.QueryPerspectivesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档重试
//
// @param request - RetryDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryDocResponse
func (client *Client) RetryDocWithOptions(request *RetryDocRequest, runtime *util.RuntimeOptions) (_result *RetryDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryDoc"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档重试
//
// @param request - RetryDocRequest
//
// @return RetryDocResponse
func (client *Client) RetryDoc(request *RetryDocRequest) (_result *RetryDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryDocResponse{}
	_body, _err := client.RetryDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档搜索
//
// @param tmpReq - SearchDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDocResponse
func (client *Client) SearchDocWithOptions(tmpReq *SearchDocRequest, runtime *util.RuntimeOptions) (_result *SearchDocResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CategoryIds)) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, tea.String("CategoryIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryIdsShrink)) {
		query["CategoryIds"] = request.CategoryIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeBegin)) {
		query["CreateTimeBegin"] = request.CreateTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserName)) {
		query["CreateUserName"] = request.CreateUserName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeBegin)) {
		query["EndTimeBegin"] = request.EndTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeEnd)) {
		query["EndTimeEnd"] = request.EndTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyTimeBegin)) {
		query["ModifyTimeBegin"] = request.ModifyTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyTimeEnd)) {
		query["ModifyTimeEnd"] = request.ModifyTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyUserName)) {
		query["ModifyUserName"] = request.ModifyUserName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessStatus)) {
		query["ProcessStatus"] = request.ProcessStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SearchScope)) {
		query["SearchScope"] = request.SearchScope
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeBegin)) {
		query["StartTimeBegin"] = request.StartTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeEnd)) {
		query["StartTimeEnd"] = request.StartTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		query["TagIds"] = request.TagIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchDoc"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档搜索
//
// @param request - SearchDocRequest
//
// @return SearchDocResponse
func (client *Client) SearchDoc(request *SearchDocRequest) (_result *SearchDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchDocResponse{}
	_body, _err := client.SearchDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 知识搜索
//
// @param tmpReq - SearchFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFaqResponse
func (client *Client) SearchFaqWithOptions(tmpReq *SearchFaqRequest, runtime *util.RuntimeOptions) (_result *SearchFaqResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchFaqShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CategoryIds)) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, tea.String("CategoryIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryIdsShrink)) {
		body["CategoryIds"] = request.CategoryIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeBegin)) {
		body["CreateTimeBegin"] = request.CreateTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserName)) {
		body["CreateUserName"] = request.CreateUserName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeBegin)) {
		body["EndTimeBegin"] = request.EndTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeEnd)) {
		body["EndTimeEnd"] = request.EndTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyTimeBegin)) {
		body["ModifyTimeBegin"] = request.ModifyTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyTimeEnd)) {
		body["ModifyTimeEnd"] = request.ModifyTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyUserName)) {
		body["ModifyUserName"] = request.ModifyUserName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchScope)) {
		body["SearchScope"] = request.SearchScope
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeBegin)) {
		body["StartTimeBegin"] = request.StartTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeEnd)) {
		body["StartTimeEnd"] = request.StartTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchFaq"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 知识搜索
//
// @param request - SearchFaqRequest
//
// @return SearchFaqResponse
func (client *Client) SearchFaq(request *SearchFaqRequest) (_result *SearchFaqResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchFaqResponse{}
	_body, _err := client.SearchFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑类目
//
// @param request - UpdateCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCategoryResponse
func (client *Client) UpdateCategoryWithOptions(request *UpdateCategoryRequest, runtime *util.RuntimeOptions) (_result *UpdateCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCategory"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑类目
//
// @param request - UpdateCategoryRequest
//
// @return UpdateCategoryResponse
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

// Summary:
//
// 更新FAQ关联问
//
// @param request - UpdateConnQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnQuestionResponse
func (client *Client) UpdateConnQuestionWithOptions(request *UpdateConnQuestionRequest, runtime *util.RuntimeOptions) (_result *UpdateConnQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnQuestionId)) {
		body["ConnQuestionId"] = request.ConnQuestionId
	}

	if !tea.BoolValue(util.IsUnset(request.OutlineId)) {
		body["OutlineId"] = request.OutlineId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConnQuestion"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConnQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新FAQ关联问
//
// @param request - UpdateConnQuestionRequest
//
// @return UpdateConnQuestionResponse
func (client *Client) UpdateConnQuestion(request *UpdateConnQuestionRequest) (_result *UpdateConnQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConnQuestionResponse{}
	_body, _err := client.UpdateConnQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-更新
//
// @param request - UpdateDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDSEntityResponse
func (client *Client) UpdateDSEntityWithOptions(request *UpdateDSEntityRequest, runtime *util.RuntimeOptions) (_result *UpdateDSEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		query["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDSEntity"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-更新
//
// @param request - UpdateDSEntityRequest
//
// @return UpdateDSEntityResponse
func (client *Client) UpdateDSEntity(request *UpdateDSEntityRequest) (_result *UpdateDSEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDSEntityResponse{}
	_body, _err := client.UpdateDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体成员-更新
//
// @param tmpReq - UpdateDSEntityValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDSEntityValueResponse
func (client *Client) UpdateDSEntityValueWithOptions(tmpReq *UpdateDSEntityValueRequest, runtime *util.RuntimeOptions) (_result *UpdateDSEntityValueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDSEntityValueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Synonyms)) {
		request.SynonymsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Synonyms, tea.String("Synonyms"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityValueId)) {
		query["EntityValueId"] = request.EntityValueId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SynonymsShrink)) {
		body["Synonyms"] = request.SynonymsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDSEntityValue"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDSEntityValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体成员-更新
//
// @param request - UpdateDSEntityValueRequest
//
// @return UpdateDSEntityValueResponse
func (client *Client) UpdateDSEntityValue(request *UpdateDSEntityValueRequest) (_result *UpdateDSEntityValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDSEntityValueResponse{}
	_body, _err := client.UpdateDSEntityValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档变更
//
// @param tmpReq - UpdateDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocResponse
func (client *Client) UpdateDocWithOptions(tmpReq *UpdateDocRequest, runtime *util.RuntimeOptions) (_result *UpdateDocResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocMetadata)) {
		request.DocMetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocMetadata, tea.String("DocMetadata"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DocMetadataShrink)) {
		query["DocMetadata"] = request.DocMetadataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DocName)) {
		query["DocName"] = request.DocName
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !tea.BoolValue(util.IsUnset(request.Meta)) {
		query["Meta"] = request.Meta
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		query["TagIds"] = request.TagIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDoc"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档变更
//
// @param request - UpdateDocRequest
//
// @return UpdateDocResponse
func (client *Client) UpdateDoc(request *UpdateDocRequest) (_result *UpdateDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDocResponse{}
	_body, _err := client.UpdateDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新FAQ
//
// @param request - UpdateFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFaqResponse
func (client *Client) UpdateFaqWithOptions(request *UpdateFaqRequest, runtime *util.RuntimeOptions) (_result *UpdateFaqResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.KnowledgeId)) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFaq"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新FAQ
//
// @param request - UpdateFaqRequest
//
// @return UpdateFaqResponse
func (client *Client) UpdateFaq(request *UpdateFaqRequest) (_result *UpdateFaqResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaqResponse{}
	_body, _err := client.UpdateFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-修改
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		query["Introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-修改
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-更新
//
// @param tmpReq - UpdateIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntentResponse
func (client *Client) UpdateIntentWithOptions(tmpReq *UpdateIntentRequest, runtime *util.RuntimeOptions) (_result *UpdateIntentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateIntentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IntentDefinition)) {
		request.IntentDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentDefinition, tea.String("IntentDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentDefinitionShrink)) {
		query["IntentDefinition"] = request.IntentDefinitionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IntentId)) {
		query["IntentId"] = request.IntentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIntent"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-更新
//
// @param request - UpdateIntentRequest
//
// @return UpdateIntentResponse
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

// Summary:
//
// 意图-LGF-更新
//
// @param tmpReq - UpdateLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLgfResponse
func (client *Client) UpdateLgfWithOptions(tmpReq *UpdateLgfRequest, runtime *util.RuntimeOptions) (_result *UpdateLgfResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateLgfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LgfDefinition)) {
		request.LgfDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LgfDefinition, tea.String("LgfDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LgfDefinitionShrink)) {
		query["LgfDefinition"] = request.LgfDefinitionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LgfId)) {
		query["LgfId"] = request.LgfId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLgf"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLgfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-LGF-更新
//
// @param request - UpdateLgfRequest
//
// @return UpdateLgfResponse
func (client *Client) UpdateLgf(request *UpdateLgfRequest) (_result *UpdateLgfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLgfResponse{}
	_body, _err := client.UpdateLgfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-修改
//
// @param request - UpdatePerspectiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePerspectiveResponse
func (client *Client) UpdatePerspectiveWithOptions(request *UpdatePerspectiveRequest, runtime *util.RuntimeOptions) (_result *UpdatePerspectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveId)) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePerspective"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-修改
//
// @param request - UpdatePerspectiveRequest
//
// @return UpdatePerspectiveResponse
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

// Summary:
//
// 更新FAQ相似问
//
// @param request - UpdateSimQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSimQuestionResponse
func (client *Client) UpdateSimQuestionWithOptions(request *UpdateSimQuestionRequest, runtime *util.RuntimeOptions) (_result *UpdateSimQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SimQuestionId)) {
		body["SimQuestionId"] = request.SimQuestionId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSimQuestion"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSimQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新FAQ相似问
//
// @param request - UpdateSimQuestionRequest
//
// @return UpdateSimQuestionResponse
func (client *Client) UpdateSimQuestion(request *UpdateSimQuestionRequest) (_result *UpdateSimQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSimQuestionResponse{}
	_body, _err := client.UpdateSimQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新FAQ答案
//
// @param request - UpdateSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSolutionResponse
func (client *Client) UpdateSolutionWithOptions(request *UpdateSolutionRequest, runtime *util.RuntimeOptions) (_result *UpdateSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveCodes)) {
		body["PerspectiveCodes"] = request.PerspectiveCodes
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		body["SolutionId"] = request.SolutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSolution"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新FAQ答案
//
// @param request - UpdateSolutionRequest
//
// @return UpdateSolutionResponse
func (client *Client) UpdateSolution(request *UpdateSolutionRequest) (_result *UpdateSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSolutionResponse{}
	_body, _err := client.UpdateSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-话术-更新
//
// @param tmpReq - UpdateUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserSayResponse
func (client *Client) UpdateUserSayWithOptions(tmpReq *UpdateUserSayRequest, runtime *util.RuntimeOptions) (_result *UpdateUserSayResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateUserSayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserSayDefinition)) {
		request.UserSayDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserSayDefinition, tea.String("UserSayDefinition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSayDefinitionShrink)) {
		query["UserSayDefinition"] = request.UserSayDefinitionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserSayId)) {
		query["UserSayId"] = request.UserSayId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserSay"),
		Version:     tea.String("2022-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserSayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-话术-更新
//
// @param request - UpdateUserSayRequest
//
// @return UpdateUserSayResponse
func (client *Client) UpdateUserSay(request *UpdateUserSayRequest) (_result *UpdateUserSayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserSayResponse{}
	_body, _err := client.UpdateUserSayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
