// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ChatRequest
	GetAgentKey() *string
	SetInstanceId(v string) *ChatRequest
	GetInstanceId() *string
	SetIntentName(v string) *ChatRequest
	GetIntentName() *string
	SetKnowledgeId(v string) *ChatRequest
	GetKnowledgeId() *string
	SetPerspective(v []*string) *ChatRequest
	GetPerspective() []*string
	SetSandBox(v bool) *ChatRequest
	GetSandBox() *bool
	SetSenderId(v string) *ChatRequest
	GetSenderId() *string
	SetSenderNick(v string) *ChatRequest
	GetSenderNick() *string
	SetSessionId(v string) *ChatRequest
	GetSessionId() *string
	SetUtterance(v string) *ChatRequest
	GetUtterance() *string
	SetVendorParam(v string) *ChatRequest
	GetVendorParam() *string
}

type ChatRequest struct {
	// The key for the business space. If omitted, the request is routed to the default business space. You can get this key from the **Business Management*	- page of your main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The unique ID of the chatbot instance. To get this ID, log in to the Alibaba Cloud Chatbot console and go to **Chatbot Details*	- > **Session API**.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of an intent within a dialog flow. If specified, the chatbot directly activates this intent to process the user\\"s request.
	//
	// example:
	//
	// 查天气意图
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// The ID of an entry in the knowledge base. If you specify this ID, the chatbot directly returns the corresponding answer.
	//
	// example:
	//
	// 30002406051
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// An array of perspective codes. Use these codes to retrieve answers from different perspectives for the same knowledge entry. Example: `Perspective=["FZJBY3raWr"]`. When using an SDK, refer to its parameter definitions.
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// Specifies the environment to use. The default value is `false`, which indicates the production environment.
	//
	// - `true`: The test environment. This environment is for testing only. Do not use it in production due to potential instability and QPS limitations.
	//
	// - `false`: The production environment.
	//
	// example:
	//
	// true
	SandBox *bool `json:"SandBox,omitempty" xml:"SandBox,omitempty"`
	// The unique ID of the user in the current session.
	//
	// example:
	//
	// custumer_123456
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The nickname of the user in the current session.
	//
	// example:
	//
	// 用户123456
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// The session ID, used to identify a user session and maintain context. For a new user, omit this parameter in the first call to the `Chat` API. The chatbot automatically starts a session and returns the `SessionId` in the response. To continue the conversation, include this `SessionId` in all subsequent requests. The maximum length is 64 characters.
	//
	// example:
	//
	// 9c6ebdc6e66f46ecadab3434314f6959
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The user\\"s input text.
	//
	// example:
	//
	// 今天天气怎么样？
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	// A JSON-formatted string containing custom parameters to pass to various dialog engines.
	//
	// example:
	//
	// {"phone":123456789}
	VendorParam *string `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
}

func (s ChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ChatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChatRequest) GetIntentName() *string {
	return s.IntentName
}

func (s *ChatRequest) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *ChatRequest) GetPerspective() []*string {
	return s.Perspective
}

func (s *ChatRequest) GetSandBox() *bool {
	return s.SandBox
}

func (s *ChatRequest) GetSenderId() *string {
	return s.SenderId
}

func (s *ChatRequest) GetSenderNick() *string {
	return s.SenderNick
}

func (s *ChatRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *ChatRequest) GetVendorParam() *string {
	return s.VendorParam
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

func (s *ChatRequest) Validate() error {
	return dara.Validate(s)
}
