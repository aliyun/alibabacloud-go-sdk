// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShoppingAssistantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ShoppingAssistantRequest
	GetConfig() *string
	SetContents(v *ShoppingAssistantRequestContents) *ShoppingAssistantRequest
	GetContents() *ShoppingAssistantRequestContents
	SetConversationId(v string) *ShoppingAssistantRequest
	GetConversationId() *string
	SetEnvironment(v string) *ShoppingAssistantRequest
	GetEnvironment() *string
	SetInputMessage(v *ShoppingAssistantRequestInputMessage) *ShoppingAssistantRequest
	GetInputMessage() *ShoppingAssistantRequestInputMessage
	SetInstanceId(v string) *ShoppingAssistantRequest
	GetInstanceId() *string
	SetLanguage(v string) *ShoppingAssistantRequest
	GetLanguage() *string
	SetSceneId(v string) *ShoppingAssistantRequest
	GetSceneId() *string
	SetServiceId(v string) *ShoppingAssistantRequest
	GetServiceId() *string
	SetSessionId(v string) *ShoppingAssistantRequest
	GetSessionId() *string
	SetUid(v string) *ShoppingAssistantRequest
	GetUid() *string
}

type ShoppingAssistantRequest struct {
	// The additional configuration.
	//
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The contents.
	Contents *ShoppingAssistantRequestContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Struct"`
	// The conversation ID. This parameter is not yet effective.
	//
	// example:
	//
	// e47cfae9-c0cc-42e1-91e2-e67cdb0e7b96
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// **The environment.**
	//
	// example:
	//
	// Prod: productionPre: pre-release.
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The input message.
	InputMessage *ShoppingAssistantRequestInputMessage `json:"InputMessage,omitempty" xml:"InputMessage,omitempty" type:"Struct"`
	// **The instance ID.**
	//
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language.
	//
	// example:
	//
	// zh/en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// **The scene ID.**
	//
	// example:
	//
	// ai_shopping
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// **The service ID.**
	//
	// example:
	//
	// ServiceId.
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// e47cfae9-c0cc-42e1-91e2-e67cdb0e7b96
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// user id。
	//
	// example:
	//
	// 0001
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ShoppingAssistantRequest) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantRequest) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantRequest) GetConfig() *string {
	return s.Config
}

func (s *ShoppingAssistantRequest) GetContents() *ShoppingAssistantRequestContents {
	return s.Contents
}

func (s *ShoppingAssistantRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ShoppingAssistantRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ShoppingAssistantRequest) GetInputMessage() *ShoppingAssistantRequestInputMessage {
	return s.InputMessage
}

func (s *ShoppingAssistantRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ShoppingAssistantRequest) GetLanguage() *string {
	return s.Language
}

func (s *ShoppingAssistantRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ShoppingAssistantRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ShoppingAssistantRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ShoppingAssistantRequest) GetUid() *string {
	return s.Uid
}

func (s *ShoppingAssistantRequest) SetConfig(v string) *ShoppingAssistantRequest {
	s.Config = &v
	return s
}

func (s *ShoppingAssistantRequest) SetContents(v *ShoppingAssistantRequestContents) *ShoppingAssistantRequest {
	s.Contents = v
	return s
}

func (s *ShoppingAssistantRequest) SetConversationId(v string) *ShoppingAssistantRequest {
	s.ConversationId = &v
	return s
}

func (s *ShoppingAssistantRequest) SetEnvironment(v string) *ShoppingAssistantRequest {
	s.Environment = &v
	return s
}

func (s *ShoppingAssistantRequest) SetInputMessage(v *ShoppingAssistantRequestInputMessage) *ShoppingAssistantRequest {
	s.InputMessage = v
	return s
}

func (s *ShoppingAssistantRequest) SetInstanceId(v string) *ShoppingAssistantRequest {
	s.InstanceId = &v
	return s
}

func (s *ShoppingAssistantRequest) SetLanguage(v string) *ShoppingAssistantRequest {
	s.Language = &v
	return s
}

func (s *ShoppingAssistantRequest) SetSceneId(v string) *ShoppingAssistantRequest {
	s.SceneId = &v
	return s
}

func (s *ShoppingAssistantRequest) SetServiceId(v string) *ShoppingAssistantRequest {
	s.ServiceId = &v
	return s
}

func (s *ShoppingAssistantRequest) SetSessionId(v string) *ShoppingAssistantRequest {
	s.SessionId = &v
	return s
}

func (s *ShoppingAssistantRequest) SetUid(v string) *ShoppingAssistantRequest {
	s.Uid = &v
	return s
}

func (s *ShoppingAssistantRequest) Validate() error {
	if s.Contents != nil {
		if err := s.Contents.Validate(); err != nil {
			return err
		}
	}
	if s.InputMessage != nil {
		if err := s.InputMessage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ShoppingAssistantRequestContents struct {
	// The message content.
	//
	// example:
	//
	// Recommend some light-colored long-sleeve shirts suitable for spring, budget under 300
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The message type.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ShoppingAssistantRequestContents) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantRequestContents) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantRequestContents) GetText() *string {
	return s.Text
}

func (s *ShoppingAssistantRequestContents) GetType() *string {
	return s.Type
}

func (s *ShoppingAssistantRequestContents) SetText(v string) *ShoppingAssistantRequestContents {
	s.Text = &v
	return s
}

func (s *ShoppingAssistantRequestContents) SetType(v string) *ShoppingAssistantRequestContents {
	s.Type = &v
	return s
}

func (s *ShoppingAssistantRequestContents) Validate() error {
	return dara.Validate(s)
}

type ShoppingAssistantRequestInputMessage struct {
	// The message content.
	Content []*ShoppingAssistantRequestInputMessageContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
}

func (s ShoppingAssistantRequestInputMessage) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantRequestInputMessage) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantRequestInputMessage) GetContent() []*ShoppingAssistantRequestInputMessageContent {
	return s.Content
}

func (s *ShoppingAssistantRequestInputMessage) SetContent(v []*ShoppingAssistantRequestInputMessageContent) *ShoppingAssistantRequestInputMessage {
	s.Content = v
	return s
}

func (s *ShoppingAssistantRequestInputMessage) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ShoppingAssistantRequestInputMessageContent struct {
	// The message content.
	//
	// example:
	//
	// Recommend some light-colored long-sleeve shirts suitable for spring, budget under 300
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The message type.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ShoppingAssistantRequestInputMessageContent) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantRequestInputMessageContent) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantRequestInputMessageContent) GetText() *string {
	return s.Text
}

func (s *ShoppingAssistantRequestInputMessageContent) GetType() *string {
	return s.Type
}

func (s *ShoppingAssistantRequestInputMessageContent) SetText(v string) *ShoppingAssistantRequestInputMessageContent {
	s.Text = &v
	return s
}

func (s *ShoppingAssistantRequestInputMessageContent) SetType(v string) *ShoppingAssistantRequestInputMessageContent {
	s.Type = &v
	return s
}

func (s *ShoppingAssistantRequestInputMessageContent) Validate() error {
	return dara.Validate(s)
}
