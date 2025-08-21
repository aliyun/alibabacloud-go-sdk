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
