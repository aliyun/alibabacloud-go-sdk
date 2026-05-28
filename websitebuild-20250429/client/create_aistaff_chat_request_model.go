// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIStaffChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateAIStaffChatRequest
	GetBizId() *string
	SetChatId(v string) *CreateAIStaffChatRequest
	GetChatId() *string
	SetConversationId(v string) *CreateAIStaffChatRequest
	GetConversationId() *string
	SetMessages(v []*CreateAIStaffChatRequestMessages) *CreateAIStaffChatRequest
	GetMessages() []*CreateAIStaffChatRequestMessages
	SetMetaData(v map[string]*string) *CreateAIStaffChatRequest
	GetMetaData() map[string]*string
}

type CreateAIStaffChatRequest struct {
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 3b465fe1-6f06-4899-af9f-d43d9338df25
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// example:
	//
	// 5b7105a2-2999-430b-ba23-ba09149d5434
	ConversationId *string                             `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	Messages       []*CreateAIStaffChatRequestMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	MetaData       map[string]*string                  `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
}

func (s CreateAIStaffChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffChatRequest) GoString() string {
	return s.String()
}

func (s *CreateAIStaffChatRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateAIStaffChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateAIStaffChatRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateAIStaffChatRequest) GetMessages() []*CreateAIStaffChatRequestMessages {
	return s.Messages
}

func (s *CreateAIStaffChatRequest) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *CreateAIStaffChatRequest) SetBizId(v string) *CreateAIStaffChatRequest {
	s.BizId = &v
	return s
}

func (s *CreateAIStaffChatRequest) SetChatId(v string) *CreateAIStaffChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreateAIStaffChatRequest) SetConversationId(v string) *CreateAIStaffChatRequest {
	s.ConversationId = &v
	return s
}

func (s *CreateAIStaffChatRequest) SetMessages(v []*CreateAIStaffChatRequestMessages) *CreateAIStaffChatRequest {
	s.Messages = v
	return s
}

func (s *CreateAIStaffChatRequest) SetMetaData(v map[string]*string) *CreateAIStaffChatRequest {
	s.MetaData = v
	return s
}

func (s *CreateAIStaffChatRequest) Validate() error {
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

type CreateAIStaffChatRequestMessages struct {
	// example:
	//
	// {\\"CodeRevision\\": \\"1750040991876284109\\"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// application/octet-stream
	ContentType *string            `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	MetaData    map[string]*string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// example:
	//
	// polarx_dn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// risk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAIStaffChatRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffChatRequestMessages) GoString() string {
	return s.String()
}

func (s *CreateAIStaffChatRequestMessages) GetContent() *string {
	return s.Content
}

func (s *CreateAIStaffChatRequestMessages) GetContentType() *string {
	return s.ContentType
}

func (s *CreateAIStaffChatRequestMessages) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *CreateAIStaffChatRequestMessages) GetRole() *string {
	return s.Role
}

func (s *CreateAIStaffChatRequestMessages) GetType() *string {
	return s.Type
}

func (s *CreateAIStaffChatRequestMessages) SetContent(v string) *CreateAIStaffChatRequestMessages {
	s.Content = &v
	return s
}

func (s *CreateAIStaffChatRequestMessages) SetContentType(v string) *CreateAIStaffChatRequestMessages {
	s.ContentType = &v
	return s
}

func (s *CreateAIStaffChatRequestMessages) SetMetaData(v map[string]*string) *CreateAIStaffChatRequestMessages {
	s.MetaData = v
	return s
}

func (s *CreateAIStaffChatRequestMessages) SetRole(v string) *CreateAIStaffChatRequestMessages {
	s.Role = &v
	return s
}

func (s *CreateAIStaffChatRequestMessages) SetType(v string) *CreateAIStaffChatRequestMessages {
	s.Type = &v
	return s
}

func (s *CreateAIStaffChatRequestMessages) Validate() error {
	return dara.Validate(s)
}
