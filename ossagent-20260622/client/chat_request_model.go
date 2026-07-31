// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*ChatRequestMessages) *ChatRequest
	GetMessages() []*ChatRequestMessages
	SetSessionId(v string) *ChatRequest
	GetSessionId() *string
}

type ChatRequest struct {
	Messages  []*ChatRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	SessionId *string                `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) GetMessages() []*ChatRequestMessages {
	return s.Messages
}

func (s *ChatRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatRequest) SetMessages(v []*ChatRequestMessages) *ChatRequest {
	s.Messages = v
	return s
}

func (s *ChatRequest) SetSessionId(v string) *ChatRequest {
	s.SessionId = &v
	return s
}

func (s *ChatRequest) Validate() error {
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

type ChatRequestMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ChatRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestMessages) GoString() string {
	return s.String()
}

func (s *ChatRequestMessages) GetContent() *string {
	return s.Content
}

func (s *ChatRequestMessages) GetRole() *string {
	return s.Role
}

func (s *ChatRequestMessages) SetContent(v string) *ChatRequestMessages {
	s.Content = &v
	return s
}

func (s *ChatRequestMessages) SetRole(v string) *ChatRequestMessages {
	s.Role = &v
	return s
}

func (s *ChatRequestMessages) Validate() error {
	return dara.Validate(s)
}
