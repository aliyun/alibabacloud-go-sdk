// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAgentChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommunication(v *MedicalAgentChatRequestCommunication) *MedicalAgentChatRequest
	GetCommunication() *MedicalAgentChatRequestCommunication
	SetData(v *MedicalAgentChatRequestData) *MedicalAgentChatRequest
	GetData() *MedicalAgentChatRequestData
	SetParams(v *MedicalAgentChatRequestParams) *MedicalAgentChatRequest
	GetParams() *MedicalAgentChatRequestParams
}

type MedicalAgentChatRequest struct {
	// The correlation ID, session, and application identity information for this request.
	//
	// This parameter is required.
	Communication *MedicalAgentChatRequestCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// The business data that carries the message for this request.
	//
	// This parameter is required.
	Data *MedicalAgentChatRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The optional extensible JSON object. The directAgent parameter can be omitted or set to med_qa. Other parameters are passed based on the parameters supported by the medical agent.
	Params *MedicalAgentChatRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
}

func (s MedicalAgentChatRequest) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatRequest) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatRequest) GetCommunication() *MedicalAgentChatRequestCommunication {
	return s.Communication
}

func (s *MedicalAgentChatRequest) GetData() *MedicalAgentChatRequestData {
	return s.Data
}

func (s *MedicalAgentChatRequest) GetParams() *MedicalAgentChatRequestParams {
	return s.Params
}

func (s *MedicalAgentChatRequest) SetCommunication(v *MedicalAgentChatRequestCommunication) *MedicalAgentChatRequest {
	s.Communication = v
	return s
}

func (s *MedicalAgentChatRequest) SetData(v *MedicalAgentChatRequestData) *MedicalAgentChatRequest {
	s.Data = v
	return s
}

func (s *MedicalAgentChatRequest) SetParams(v *MedicalAgentChatRequestParams) *MedicalAgentChatRequest {
	s.Params = v
	return s
}

func (s *MedicalAgentChatRequest) Validate() error {
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

type MedicalAgentChatRequestCommunication struct {
	// The business application ID. This parameter is required if no fixed application is configured.
	//
	// example:
	//
	// your-app-id
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The optional business scenario tag.
	//
	// example:
	//
	// health_qa
	AppScene *string `json:"appScene,omitempty" xml:"appScene,omitempty"`
	// The session history owner. This value must be the same as the one used in the first call of the session.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLIENT
	HistoryOwner *string `json:"historyOwner,omitempty" xml:"historyOwner,omitempty"`
	// The customer business request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The customer session ID. Reuse this ID to continue the conversation.
	//
	// This parameter is required.
	//
	// example:
	//
	// med-session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The business user and device information provided by the customer. Identity fields are passed based on the configuration of the business gateway credentials.
	UserInfo *MedicalAgentChatRequestCommunicationUserInfo `json:"userInfo,omitempty" xml:"userInfo,omitempty" type:"Struct"`
}

func (s MedicalAgentChatRequestCommunication) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatRequestCommunication) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatRequestCommunication) GetAppId() *string {
	return s.AppId
}

func (s *MedicalAgentChatRequestCommunication) GetAppScene() *string {
	return s.AppScene
}

func (s *MedicalAgentChatRequestCommunication) GetHistoryOwner() *string {
	return s.HistoryOwner
}

func (s *MedicalAgentChatRequestCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatRequestCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *MedicalAgentChatRequestCommunication) GetUserInfo() *MedicalAgentChatRequestCommunicationUserInfo {
	return s.UserInfo
}

func (s *MedicalAgentChatRequestCommunication) SetAppId(v string) *MedicalAgentChatRequestCommunication {
	s.AppId = &v
	return s
}

func (s *MedicalAgentChatRequestCommunication) SetAppScene(v string) *MedicalAgentChatRequestCommunication {
	s.AppScene = &v
	return s
}

func (s *MedicalAgentChatRequestCommunication) SetHistoryOwner(v string) *MedicalAgentChatRequestCommunication {
	s.HistoryOwner = &v
	return s
}

func (s *MedicalAgentChatRequestCommunication) SetRequestId(v string) *MedicalAgentChatRequestCommunication {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatRequestCommunication) SetSessionId(v string) *MedicalAgentChatRequestCommunication {
	s.SessionId = &v
	return s
}

func (s *MedicalAgentChatRequestCommunication) SetUserInfo(v *MedicalAgentChatRequestCommunicationUserInfo) *MedicalAgentChatRequestCommunication {
	s.UserInfo = v
	return s
}

func (s *MedicalAgentChatRequestCommunication) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MedicalAgentChatRequestCommunicationUserInfo struct {
	// The terminal device identifier provided by the customer.
	//
	// example:
	//
	// demo-device-001
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The business user identifier in the customer application. This is provided based on the user binding or delegation configuration of the business gateway credentials.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The IP address of the business user provided by the customer.
	//
	// example:
	//
	// 192.0.2.10
	UserIp *string `json:"userIp,omitempty" xml:"userIp,omitempty"`
	// The name of the business user provided by the customer.
	//
	// example:
	//
	// Sample user
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s MedicalAgentChatRequestCommunicationUserInfo) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatRequestCommunicationUserInfo) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) GetUserIp() *string {
	return s.UserIp
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) GetUserName() *string {
	return s.UserName
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) SetDeviceId(v string) *MedicalAgentChatRequestCommunicationUserInfo {
	s.DeviceId = &v
	return s
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) SetUserId(v string) *MedicalAgentChatRequestCommunicationUserInfo {
	s.UserId = &v
	return s
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) SetUserIp(v string) *MedicalAgentChatRequestCommunicationUserInfo {
	s.UserIp = &v
	return s
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) SetUserName(v string) *MedicalAgentChatRequestCommunicationUserInfo {
	s.UserName = &v
	return s
}

func (s *MedicalAgentChatRequestCommunicationUserInfo) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatRequestData struct {
	// The list of messages arranged in conversation order. When the history owner is CLIENT, this list can include the session history provided by the client.
	//
	// This parameter is required.
	Messages []*MedicalAgentChatRequestDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s MedicalAgentChatRequestData) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatRequestData) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatRequestData) GetMessages() []*MedicalAgentChatRequestDataMessages {
	return s.Messages
}

func (s *MedicalAgentChatRequestData) SetMessages(v []*MedicalAgentChatRequestDataMessages) *MedicalAgentChatRequestData {
	s.Messages = v
	return s
}

func (s *MedicalAgentChatRequestData) Validate() error {
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

type MedicalAgentChatRequestDataMessages struct {
	// The string or array of multimodal content blocks.
	//
	// This parameter is required.
	//
	// example:
	//
	// What does the hemoglobin level in the physical examination report mean?
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The sender of the message. A value of user indicates the business user, and assistant indicates the assistant.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s MedicalAgentChatRequestDataMessages) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatRequestDataMessages) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatRequestDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *MedicalAgentChatRequestDataMessages) GetRole() *string {
	return s.Role
}

func (s *MedicalAgentChatRequestDataMessages) SetContent(v interface{}) *MedicalAgentChatRequestDataMessages {
	s.Content = v
	return s
}

func (s *MedicalAgentChatRequestDataMessages) SetRole(v string) *MedicalAgentChatRequestDataMessages {
	s.Role = &v
	return s
}

func (s *MedicalAgentChatRequestDataMessages) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatRequestParams struct {
	// Specifies whether to enable deep thinking for the medical agent.
	//
	// example:
	//
	// false
	Deepthink *string `json:"deepthink,omitempty" xml:"deepthink,omitempty"`
	// This parameter can be omitted. If explicitly provided, the value must be med_qa.
	//
	// example:
	//
	// med_qa
	DirectAgent *string `json:"directAgent,omitempty" xml:"directAgent,omitempty"`
	// The response mode. A value of false returns a complete JSON response, and true returns Server-Sent Events (SSE).
	//
	// example:
	//
	// false
	ResponseStreaming *bool `json:"responseStreaming,omitempty" xml:"responseStreaming,omitempty"`
}

func (s MedicalAgentChatRequestParams) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatRequestParams) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatRequestParams) GetDeepthink() *string {
	return s.Deepthink
}

func (s *MedicalAgentChatRequestParams) GetDirectAgent() *string {
	return s.DirectAgent
}

func (s *MedicalAgentChatRequestParams) GetResponseStreaming() *bool {
	return s.ResponseStreaming
}

func (s *MedicalAgentChatRequestParams) SetDeepthink(v string) *MedicalAgentChatRequestParams {
	s.Deepthink = &v
	return s
}

func (s *MedicalAgentChatRequestParams) SetDirectAgent(v string) *MedicalAgentChatRequestParams {
	s.DirectAgent = &v
	return s
}

func (s *MedicalAgentChatRequestParams) SetResponseStreaming(v bool) *MedicalAgentChatRequestParams {
	s.ResponseStreaming = &v
	return s
}

func (s *MedicalAgentChatRequestParams) Validate() error {
	return dara.Validate(s)
}
