// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ChatShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *ChatShrinkRequest
	GetInstanceId() *string
	SetIntentName(v string) *ChatShrinkRequest
	GetIntentName() *string
	SetKnowledgeId(v string) *ChatShrinkRequest
	GetKnowledgeId() *string
	SetPerspectiveShrink(v string) *ChatShrinkRequest
	GetPerspectiveShrink() *string
	SetSandBox(v bool) *ChatShrinkRequest
	GetSandBox() *bool
	SetSenderId(v string) *ChatShrinkRequest
	GetSenderId() *string
	SetSenderNick(v string) *ChatShrinkRequest
	GetSenderNick() *string
	SetSessionId(v string) *ChatShrinkRequest
	GetSessionId() *string
	SetUtterance(v string) *ChatShrinkRequest
	GetUtterance() *string
	SetVendorParam(v string) *ChatShrinkRequest
	GetVendorParam() *string
}

type ChatShrinkRequest struct {
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
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
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

func (s ChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ChatShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChatShrinkRequest) GetIntentName() *string {
	return s.IntentName
}

func (s *ChatShrinkRequest) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *ChatShrinkRequest) GetPerspectiveShrink() *string {
	return s.PerspectiveShrink
}

func (s *ChatShrinkRequest) GetSandBox() *bool {
	return s.SandBox
}

func (s *ChatShrinkRequest) GetSenderId() *string {
	return s.SenderId
}

func (s *ChatShrinkRequest) GetSenderNick() *string {
	return s.SenderNick
}

func (s *ChatShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatShrinkRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *ChatShrinkRequest) GetVendorParam() *string {
	return s.VendorParam
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

func (s *ChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
