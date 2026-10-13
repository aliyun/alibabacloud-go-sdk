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
	// The business space key. If this parameter is not specified, the default business space is accessed. Obtain the key value from the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The unique chatbot identifier (chatbot ID). Log on to the Chatbot console and navigate to Chatbot Details > Session Interface to view the chatbot instance information and obtain the instance ID.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The intent name in the dialog flow. If this name is specified, the chatbot directly enters this intent for Q&A.
	//
	// example:
	//
	// CheckWeatherIntent
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// The ID of a knowledge title in the knowledge base. If this ID is specified, the chatbot directly returns the answer for the specified knowledge title.
	//
	// example:
	//
	// 30002406051
	KnowledgeId *string `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The array of perspective codes, which is used to call answers from different perspectives under the same knowledge title. Example: Perspective=["FZJBY3raWr"]. When using the SDK, use the parameters defined in the SDK.
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
	// Specifies whether to use the test or production environment. Default value: false (production environment). Valid values:
	//
	// - true: Test environment. This is for testing only. Using it in production may cause instability and QPS throttling.
	//
	// - false: Production environment.
	//
	// example:
	//
	// true
	SandBox *bool `json:"SandBox,omitempty" xml:"SandBox,omitempty"`
	// The visitor ID, which identifies the user in the current session.
	//
	// example:
	//
	// custumer_123456
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The nickname of the visitor in the current session.
	//
	// example:
	//
	// User123456
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// The session ID, which identifies a visitor\\"s session and maintains context information. For a new visitor, you do not need to pass this field when calling the Chat operation for the first time. The chatbot starts a session and returns the SessionId in the response of the Chat operation. For subsequent conversation turns of the visitor, pass the SessionId of the current session when calling the Chat operation so that the chatbot can continue the conversation based on the SessionId. Maximum length: 64 characters.
	//
	// example:
	//
	// 9c6ebdc6e66f46ecadab3434314f6959
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The input from the chatbot visitor.
	//
	// example:
	//
	// How is the weather today?
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	// A JSON-formatted set of user-defined parameters that can be passed to various dialog engines.
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
