// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversationDTO interface {
	dara.Model
	String() string
	GoString() string
	SetChatData(v string) *ConversationDTO
	GetChatData() *string
	SetDeleteTag(v int32) *ConversationDTO
	GetDeleteTag() *int32
	SetGmtCreate(v string) *ConversationDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ConversationDTO
	GetGmtModified() *string
	SetId(v int64) *ConversationDTO
	GetId() *int64
	SetMessageCount(v int32) *ConversationDTO
	GetMessageCount() *int32
	SetModelIds(v string) *ConversationDTO
	GetModelIds() *string
	SetTitle(v string) *ConversationDTO
	GetTitle() *string
}

type ConversationDTO struct {
	// example:
	//
	// {}
	ChatData *string `json:"chatData,omitempty" xml:"chatData,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 10
	MessageCount *int32 `json:"messageCount,omitempty" xml:"messageCount,omitempty"`
	// example:
	//
	// [1,2,3]
	ModelIds *string `json:"modelIds,omitempty" xml:"modelIds,omitempty"`
	// example:
	//
	// 我的对话
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ConversationDTO) String() string {
	return dara.Prettify(s)
}

func (s ConversationDTO) GoString() string {
	return s.String()
}

func (s *ConversationDTO) GetChatData() *string {
	return s.ChatData
}

func (s *ConversationDTO) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ConversationDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ConversationDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ConversationDTO) GetId() *int64 {
	return s.Id
}

func (s *ConversationDTO) GetMessageCount() *int32 {
	return s.MessageCount
}

func (s *ConversationDTO) GetModelIds() *string {
	return s.ModelIds
}

func (s *ConversationDTO) GetTitle() *string {
	return s.Title
}

func (s *ConversationDTO) SetChatData(v string) *ConversationDTO {
	s.ChatData = &v
	return s
}

func (s *ConversationDTO) SetDeleteTag(v int32) *ConversationDTO {
	s.DeleteTag = &v
	return s
}

func (s *ConversationDTO) SetGmtCreate(v string) *ConversationDTO {
	s.GmtCreate = &v
	return s
}

func (s *ConversationDTO) SetGmtModified(v string) *ConversationDTO {
	s.GmtModified = &v
	return s
}

func (s *ConversationDTO) SetId(v int64) *ConversationDTO {
	s.Id = &v
	return s
}

func (s *ConversationDTO) SetMessageCount(v int32) *ConversationDTO {
	s.MessageCount = &v
	return s
}

func (s *ConversationDTO) SetModelIds(v string) *ConversationDTO {
	s.ModelIds = &v
	return s
}

func (s *ConversationDTO) SetTitle(v string) *ConversationDTO {
	s.Title = &v
	return s
}

func (s *ConversationDTO) Validate() error {
	return dara.Validate(s)
}
