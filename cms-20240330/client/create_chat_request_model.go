// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateChatRequest
	GetAction() *string
	SetDigitalEmployeeName(v string) *CreateChatRequest
	GetDigitalEmployeeName() *string
	SetMessages(v []*CreateChatRequestMessages) *CreateChatRequest
	GetMessages() []*CreateChatRequestMessages
	SetThreadId(v string) *CreateChatRequest
	GetThreadId() *string
	SetVariables(v map[string]interface{}) *CreateChatRequest
	GetVariables() map[string]interface{}
}

type CreateChatRequest struct {
	// example:
	//
	// create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// test
	DigitalEmployeeName *string                      `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	Messages            []*CreateChatRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// thread_id01
	ThreadId  *string                `json:"threadId,omitempty" xml:"threadId,omitempty"`
	Variables map[string]interface{} `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s CreateChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatRequest) GoString() string {
	return s.String()
}

func (s *CreateChatRequest) GetAction() *string {
	return s.Action
}

func (s *CreateChatRequest) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *CreateChatRequest) GetMessages() []*CreateChatRequestMessages {
	return s.Messages
}

func (s *CreateChatRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *CreateChatRequest) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *CreateChatRequest) SetAction(v string) *CreateChatRequest {
	s.Action = &v
	return s
}

func (s *CreateChatRequest) SetDigitalEmployeeName(v string) *CreateChatRequest {
	s.DigitalEmployeeName = &v
	return s
}

func (s *CreateChatRequest) SetMessages(v []*CreateChatRequestMessages) *CreateChatRequest {
	s.Messages = v
	return s
}

func (s *CreateChatRequest) SetThreadId(v string) *CreateChatRequest {
	s.ThreadId = &v
	return s
}

func (s *CreateChatRequest) SetVariables(v map[string]interface{}) *CreateChatRequest {
	s.Variables = v
	return s
}

func (s *CreateChatRequest) Validate() error {
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

type CreateChatRequestMessages struct {
	Contents []*CreateChatRequestMessagesContents `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	// example:
	//
	// message_id02
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// example:
	//
	// system
	Role  *string                  `json:"role,omitempty" xml:"role,omitempty"`
	Tools []map[string]interface{} `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s CreateChatRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateChatRequestMessages) GoString() string {
	return s.String()
}

func (s *CreateChatRequestMessages) GetContents() []*CreateChatRequestMessagesContents {
	return s.Contents
}

func (s *CreateChatRequestMessages) GetMessageId() *string {
	return s.MessageId
}

func (s *CreateChatRequestMessages) GetRole() *string {
	return s.Role
}

func (s *CreateChatRequestMessages) GetTools() []map[string]interface{} {
	return s.Tools
}

func (s *CreateChatRequestMessages) SetContents(v []*CreateChatRequestMessagesContents) *CreateChatRequestMessages {
	s.Contents = v
	return s
}

func (s *CreateChatRequestMessages) SetMessageId(v string) *CreateChatRequestMessages {
	s.MessageId = &v
	return s
}

func (s *CreateChatRequestMessages) SetRole(v string) *CreateChatRequestMessages {
	s.Role = &v
	return s
}

func (s *CreateChatRequestMessages) SetTools(v []map[string]interface{}) *CreateChatRequestMessages {
	s.Tools = v
	return s
}

func (s *CreateChatRequestMessages) Validate() error {
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateChatRequestMessagesContents struct {
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateChatRequestMessagesContents) String() string {
	return dara.Prettify(s)
}

func (s CreateChatRequestMessagesContents) GoString() string {
	return s.String()
}

func (s *CreateChatRequestMessagesContents) GetType() *string {
	return s.Type
}

func (s *CreateChatRequestMessagesContents) GetValue() *string {
	return s.Value
}

func (s *CreateChatRequestMessagesContents) SetType(v string) *CreateChatRequestMessagesContents {
	s.Type = &v
	return s
}

func (s *CreateChatRequestMessagesContents) SetValue(v string) *CreateChatRequestMessagesContents {
	s.Value = &v
	return s
}

func (s *CreateChatRequestMessagesContents) Validate() error {
	return dara.Validate(s)
}
