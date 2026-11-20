// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalModelInferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommunication(v *MedicalModelInferenceRequestCommunication) *MedicalModelInferenceRequest
	GetCommunication() *MedicalModelInferenceRequestCommunication
	SetData(v *MedicalModelInferenceRequestData) *MedicalModelInferenceRequest
	GetData() *MedicalModelInferenceRequestData
	SetParams(v *MedicalModelInferenceRequestParams) *MedicalModelInferenceRequest
	GetParams() *MedicalModelInferenceRequestParams
}

type MedicalModelInferenceRequest struct {
	// The association ID, session, and application identity information for the current request.
	//
	// This parameter is required.
	Communication *MedicalModelInferenceRequestCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// The business data that carries the messages for the current request.
	//
	// This parameter is required.
	Data *MedicalModelInferenceRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The extensible JSON object. The modelName parameter is required and its value must be an activated model selector. The directAgent parameter cannot be passed.
	//
	// This parameter is required.
	Params *MedicalModelInferenceRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
}

func (s MedicalModelInferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceRequest) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceRequest) GetCommunication() *MedicalModelInferenceRequestCommunication {
	return s.Communication
}

func (s *MedicalModelInferenceRequest) GetData() *MedicalModelInferenceRequestData {
	return s.Data
}

func (s *MedicalModelInferenceRequest) GetParams() *MedicalModelInferenceRequestParams {
	return s.Params
}

func (s *MedicalModelInferenceRequest) SetCommunication(v *MedicalModelInferenceRequestCommunication) *MedicalModelInferenceRequest {
	s.Communication = v
	return s
}

func (s *MedicalModelInferenceRequest) SetData(v *MedicalModelInferenceRequestData) *MedicalModelInferenceRequest {
	s.Data = v
	return s
}

func (s *MedicalModelInferenceRequest) SetParams(v *MedicalModelInferenceRequestParams) *MedicalModelInferenceRequest {
	s.Params = v
	return s
}

func (s *MedicalModelInferenceRequest) Validate() error {
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

type MedicalModelInferenceRequestCommunication struct {
	// The business application ID. This parameter is required when no fixed application is configured.
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
	// The host of the session history. This value must be the same as the value specified in the first call of the session.
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
	// The customer session ID, which is reused for continued conversations.
	//
	// This parameter is required.
	//
	// example:
	//
	// med-session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The business user and device information provided by the customer. Identity fields are passed based on the configuration of the business gateway credentials.
	UserInfo *MedicalModelInferenceRequestCommunicationUserInfo `json:"userInfo,omitempty" xml:"userInfo,omitempty" type:"Struct"`
}

func (s MedicalModelInferenceRequestCommunication) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceRequestCommunication) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceRequestCommunication) GetAppId() *string {
	return s.AppId
}

func (s *MedicalModelInferenceRequestCommunication) GetAppScene() *string {
	return s.AppScene
}

func (s *MedicalModelInferenceRequestCommunication) GetHistoryOwner() *string {
	return s.HistoryOwner
}

func (s *MedicalModelInferenceRequestCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceRequestCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *MedicalModelInferenceRequestCommunication) GetUserInfo() *MedicalModelInferenceRequestCommunicationUserInfo {
	return s.UserInfo
}

func (s *MedicalModelInferenceRequestCommunication) SetAppId(v string) *MedicalModelInferenceRequestCommunication {
	s.AppId = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunication) SetAppScene(v string) *MedicalModelInferenceRequestCommunication {
	s.AppScene = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunication) SetHistoryOwner(v string) *MedicalModelInferenceRequestCommunication {
	s.HistoryOwner = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunication) SetRequestId(v string) *MedicalModelInferenceRequestCommunication {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunication) SetSessionId(v string) *MedicalModelInferenceRequestCommunication {
	s.SessionId = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunication) SetUserInfo(v *MedicalModelInferenceRequestCommunicationUserInfo) *MedicalModelInferenceRequestCommunication {
	s.UserInfo = v
	return s
}

func (s *MedicalModelInferenceRequestCommunication) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MedicalModelInferenceRequestCommunicationUserInfo struct {
	// The terminal device identifier provided by the customer.
	//
	// example:
	//
	// demo-device-001
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The business user identifier in the customer application. This value is provided based on the user binding or delegation configuration of the business gateway credentials.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The business user IP address provided by the customer.
	//
	// example:
	//
	// 192.0.2.10
	UserIp *string `json:"userIp,omitempty" xml:"userIp,omitempty"`
	// The business user name provided by the customer.
	//
	// example:
	//
	// Sample user
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s MedicalModelInferenceRequestCommunicationUserInfo) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceRequestCommunicationUserInfo) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) GetUserIp() *string {
	return s.UserIp
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) GetUserName() *string {
	return s.UserName
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) SetDeviceId(v string) *MedicalModelInferenceRequestCommunicationUserInfo {
	s.DeviceId = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) SetUserId(v string) *MedicalModelInferenceRequestCommunicationUserInfo {
	s.UserId = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) SetUserIp(v string) *MedicalModelInferenceRequestCommunicationUserInfo {
	s.UserIp = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) SetUserName(v string) *MedicalModelInferenceRequestCommunicationUserInfo {
	s.UserName = &v
	return s
}

func (s *MedicalModelInferenceRequestCommunicationUserInfo) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceRequestData struct {
	// The list of messages arranged in conversation order. When hosted by the client, this list can include the session history provided by the client.
	//
	// This parameter is required.
	Messages []*MedicalModelInferenceRequestDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s MedicalModelInferenceRequestData) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceRequestData) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceRequestData) GetMessages() []*MedicalModelInferenceRequestDataMessages {
	return s.Messages
}

func (s *MedicalModelInferenceRequestData) SetMessages(v []*MedicalModelInferenceRequestDataMessages) *MedicalModelInferenceRequestData {
	s.Messages = v
	return s
}

func (s *MedicalModelInferenceRequestData) Validate() error {
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

type MedicalModelInferenceRequestDataMessages struct {
	// The string or array of multimodal content blocks.
	//
	// This parameter is required.
	//
	// example:
	//
	// What does hemoglobin mean in the physical examination report?
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The sender of the message. A value of user indicates the business user, and a value of assistant indicates the assistant.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s MedicalModelInferenceRequestDataMessages) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceRequestDataMessages) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceRequestDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *MedicalModelInferenceRequestDataMessages) GetRole() *string {
	return s.Role
}

func (s *MedicalModelInferenceRequestDataMessages) SetContent(v interface{}) *MedicalModelInferenceRequestDataMessages {
	s.Content = v
	return s
}

func (s *MedicalModelInferenceRequestDataMessages) SetRole(v string) *MedicalModelInferenceRequestDataMessages {
	s.Role = &v
	return s
}

func (s *MedicalModelInferenceRequestDataMessages) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceRequestParams struct {
	// Required. The model selector for this medical API is med_qwen4-vl-125.
	//
	// This parameter is required.
	//
	// example:
	//
	// med_qwen4-vl-125
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The response mode. A value of false indicates that a complete JSON object is returned. A value of true indicates that Server-Sent Events (SSE) is returned.
	//
	// example:
	//
	// false
	ResponseStreaming *bool `json:"responseStreaming,omitempty" xml:"responseStreaming,omitempty"`
}

func (s MedicalModelInferenceRequestParams) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceRequestParams) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceRequestParams) GetModelName() *string {
	return s.ModelName
}

func (s *MedicalModelInferenceRequestParams) GetResponseStreaming() *bool {
	return s.ResponseStreaming
}

func (s *MedicalModelInferenceRequestParams) SetModelName(v string) *MedicalModelInferenceRequestParams {
	s.ModelName = &v
	return s
}

func (s *MedicalModelInferenceRequestParams) SetResponseStreaming(v bool) *MedicalModelInferenceRequestParams {
	s.ResponseStreaming = &v
	return s
}

func (s *MedicalModelInferenceRequestParams) Validate() error {
	return dara.Validate(s)
}
