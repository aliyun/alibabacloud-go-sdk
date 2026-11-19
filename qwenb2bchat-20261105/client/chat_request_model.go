// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommunication(v *ChatRequestCommunication) *ChatRequest
	GetCommunication() *ChatRequestCommunication
	SetData(v *ChatRequestData) *ChatRequest
	GetData() *ChatRequestData
	SetParams(v *ChatRequestParams) *ChatRequest
	GetParams() *ChatRequestParams
}

type ChatRequest struct {
	// The correlation ID, session, and application identity information for the current request.
	//
	// This parameter is required.
	Communication *ChatRequestCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// The business data that carries the messages for the current request.
	//
	// This parameter is required.
	Data *ChatRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// An optional extensible JSON object. responseStreaming is an optional boolean value. A value of false returns a complete JSON response, and a value of true returns a Server-Sent Events (SSE) response.
	Params *ChatRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
}

func (s ChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) GetCommunication() *ChatRequestCommunication {
	return s.Communication
}

func (s *ChatRequest) GetData() *ChatRequestData {
	return s.Data
}

func (s *ChatRequest) GetParams() *ChatRequestParams {
	return s.Params
}

func (s *ChatRequest) SetCommunication(v *ChatRequestCommunication) *ChatRequest {
	s.Communication = v
	return s
}

func (s *ChatRequest) SetData(v *ChatRequestData) *ChatRequest {
	s.Data = v
	return s
}

func (s *ChatRequest) SetParams(v *ChatRequestParams) *ChatRequest {
	s.Params = v
	return s
}

func (s *ChatRequest) Validate() error {
	if s.Communication != nil {
		if err := s.Communication.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatRequestCommunication struct {
	// The application ID provided based on the gateway credential identity configuration.
	//
	// example:
	//
	// example-app
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The optional business scene marker.
	//
	// example:
	//
	// travel
	AppScene *string `json:"appScene,omitempty" xml:"appScene,omitempty"`
	// The owner of the session history, which must be consistent with the first call of the session.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLIENT
	HistoryOwner *string `json:"historyOwner,omitempty" xml:"historyOwner,omitempty"`
	// The business request ID of the customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session ID of the customer, which is reused for continued chats.
	//
	// This parameter is required.
	//
	// example:
	//
	// session-example-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The business user and device information provided by the customer. Identity fields are passed based on the configuration of the business gateway credentials.
	UserInfo *ChatRequestCommunicationUserInfo `json:"userInfo,omitempty" xml:"userInfo,omitempty" type:"Struct"`
}

func (s ChatRequestCommunication) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestCommunication) GoString() string {
	return s.String()
}

func (s *ChatRequestCommunication) GetAppId() *string {
	return s.AppId
}

func (s *ChatRequestCommunication) GetAppScene() *string {
	return s.AppScene
}

func (s *ChatRequestCommunication) GetHistoryOwner() *string {
	return s.HistoryOwner
}

func (s *ChatRequestCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatRequestCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatRequestCommunication) GetUserInfo() *ChatRequestCommunicationUserInfo {
	return s.UserInfo
}

func (s *ChatRequestCommunication) SetAppId(v string) *ChatRequestCommunication {
	s.AppId = &v
	return s
}

func (s *ChatRequestCommunication) SetAppScene(v string) *ChatRequestCommunication {
	s.AppScene = &v
	return s
}

func (s *ChatRequestCommunication) SetHistoryOwner(v string) *ChatRequestCommunication {
	s.HistoryOwner = &v
	return s
}

func (s *ChatRequestCommunication) SetRequestId(v string) *ChatRequestCommunication {
	s.RequestId = &v
	return s
}

func (s *ChatRequestCommunication) SetSessionId(v string) *ChatRequestCommunication {
	s.SessionId = &v
	return s
}

func (s *ChatRequestCommunication) SetUserInfo(v *ChatRequestCommunicationUserInfo) *ChatRequestCommunication {
	s.UserInfo = v
	return s
}

func (s *ChatRequestCommunication) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatRequestCommunicationUserInfo struct {
	// The terminal device identifier provided by the customer.
	//
	// example:
	//
	// device-example-001
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The business user identifier within the customer application, provided based on the user binding or delegation configuration of the business gateway credentials.
	//
	// example:
	//
	// example-user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The IP address of the business user provided by the customer.
	//
	// example:
	//
	// 192.0.2.10
	UserIp *string `json:"userIp,omitempty" xml:"userIp,omitempty"`
	// The username of the business user provided by the customer.
	//
	// example:
	//
	// 示例用户
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ChatRequestCommunicationUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestCommunicationUserInfo) GoString() string {
	return s.String()
}

func (s *ChatRequestCommunicationUserInfo) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ChatRequestCommunicationUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *ChatRequestCommunicationUserInfo) GetUserIp() *string {
	return s.UserIp
}

func (s *ChatRequestCommunicationUserInfo) GetUserName() *string {
	return s.UserName
}

func (s *ChatRequestCommunicationUserInfo) SetDeviceId(v string) *ChatRequestCommunicationUserInfo {
	s.DeviceId = &v
	return s
}

func (s *ChatRequestCommunicationUserInfo) SetUserId(v string) *ChatRequestCommunicationUserInfo {
	s.UserId = &v
	return s
}

func (s *ChatRequestCommunicationUserInfo) SetUserIp(v string) *ChatRequestCommunicationUserInfo {
	s.UserIp = &v
	return s
}

func (s *ChatRequestCommunicationUserInfo) SetUserName(v string) *ChatRequestCommunicationUserInfo {
	s.UserName = &v
	return s
}

func (s *ChatRequestCommunicationUserInfo) Validate() error {
	return dara.Validate(s)
}

type ChatRequestData struct {
	// The list of messages arranged in conversational order. When history is managed by the client, it can include the session history provided by the client.
	//
	// This parameter is required.
	Messages []*ChatRequestDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s ChatRequestData) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestData) GoString() string {
	return s.String()
}

func (s *ChatRequestData) GetMessages() []*ChatRequestDataMessages {
	return s.Messages
}

func (s *ChatRequestData) SetMessages(v []*ChatRequestDataMessages) *ChatRequestData {
	s.Messages = v
	return s
}

func (s *ChatRequestData) Validate() error {
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

type ChatRequestDataMessages struct {
	// The text or an array of supported multimodal content blocks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 请介绍杭州西湖。
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The sender of the message. user indicates the business user, and assistant indicates the assistant.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ChatRequestDataMessages) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestDataMessages) GoString() string {
	return s.String()
}

func (s *ChatRequestDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *ChatRequestDataMessages) GetRole() *string {
	return s.Role
}

func (s *ChatRequestDataMessages) SetContent(v interface{}) *ChatRequestDataMessages {
	s.Content = v
	return s
}

func (s *ChatRequestDataMessages) SetRole(v string) *ChatRequestDataMessages {
	s.Role = &v
	return s
}

func (s *ChatRequestDataMessages) Validate() error {
	return dara.Validate(s)
}

type ChatRequestParams struct {
	// The deep thinking switch.
	//
	// example:
	//
	// true
	Deepthink *string `json:"deepthink,omitempty" xml:"deepthink,omitempty"`
	// The agent selector for the current turn.
	//
	// example:
	//
	// study_qa
	DirectAgent *string `json:"directAgent,omitempty" xml:"directAgent,omitempty"`
	// The mixed text and image output switch.
	//
	// example:
	//
	// true
	EnableTextImageMixed *bool `json:"enableTextImageMixed,omitempty" xml:"enableTextImageMixed,omitempty"`
	// The maximum number of output tokens for the model.
	//
	// example:
	//
	// 512
	MaxTokens *int32 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The HTTP audio output switch.
	OutputAudio *bool `json:"outputAudio,omitempty" xml:"outputAudio,omitempty"`
	// Specifies whether to use Server-Sent Events (SSE) for conversation responses.
	ResponseStreaming *bool `json:"responseStreaming,omitempty" xml:"responseStreaming,omitempty"`
	// The model sampling temperature.
	//
	// example:
	//
	// 0.7
	Temperature *float32 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// The nucleus sampling probability for the model.
	//
	// example:
	//
	// 0.8
	TopP *float32 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s ChatRequestParams) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestParams) GoString() string {
	return s.String()
}

func (s *ChatRequestParams) GetDeepthink() *string {
	return s.Deepthink
}

func (s *ChatRequestParams) GetDirectAgent() *string {
	return s.DirectAgent
}

func (s *ChatRequestParams) GetEnableTextImageMixed() *bool {
	return s.EnableTextImageMixed
}

func (s *ChatRequestParams) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ChatRequestParams) GetOutputAudio() *bool {
	return s.OutputAudio
}

func (s *ChatRequestParams) GetResponseStreaming() *bool {
	return s.ResponseStreaming
}

func (s *ChatRequestParams) GetTemperature() *float32 {
	return s.Temperature
}

func (s *ChatRequestParams) GetTopP() *float32 {
	return s.TopP
}

func (s *ChatRequestParams) SetDeepthink(v string) *ChatRequestParams {
	s.Deepthink = &v
	return s
}

func (s *ChatRequestParams) SetDirectAgent(v string) *ChatRequestParams {
	s.DirectAgent = &v
	return s
}

func (s *ChatRequestParams) SetEnableTextImageMixed(v bool) *ChatRequestParams {
	s.EnableTextImageMixed = &v
	return s
}

func (s *ChatRequestParams) SetMaxTokens(v int32) *ChatRequestParams {
	s.MaxTokens = &v
	return s
}

func (s *ChatRequestParams) SetOutputAudio(v bool) *ChatRequestParams {
	s.OutputAudio = &v
	return s
}

func (s *ChatRequestParams) SetResponseStreaming(v bool) *ChatRequestParams {
	s.ResponseStreaming = &v
	return s
}

func (s *ChatRequestParams) SetTemperature(v float32) *ChatRequestParams {
	s.Temperature = &v
	return s
}

func (s *ChatRequestParams) SetTopP(v float32) *ChatRequestParams {
	s.TopP = &v
	return s
}

func (s *ChatRequestParams) Validate() error {
	return dara.Validate(s)
}
