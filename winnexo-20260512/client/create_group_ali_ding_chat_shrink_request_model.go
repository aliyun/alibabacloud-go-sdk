// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupAliDingChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupAliDingChatShrinkRequest
	GetChatId() *string
	SetChatName(v string) *CreateGroupAliDingChatShrinkRequest
	GetChatName() *string
	SetDescription(v string) *CreateGroupAliDingChatShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupAliDingChatShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupAliDingChatShrinkRequest
	GetGroupId() *string
	SetHistoryStartTime(v string) *CreateGroupAliDingChatShrinkRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreateGroupAliDingChatShrinkRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupAliDingChatShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupAliDingChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupAliDingChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreateGroupAliDingChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreateGroupAliDingChatShrinkRequest struct {
	// The session ID, typically used for JSSDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxxxx
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The group chat name.
	//
	// example:
	//
	// CustomerProjectGroup
	ChatName *string `json:"chatName,omitempty" xml:"chatName,omitempty"`
	// The description of the AI assistant.
	//
	// example:
	//
	// Customer group chat history
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The start time for collecting chat history.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-01
	HistoryStartTime *string `json:"historyStartTime,omitempty" xml:"historyStartTime,omitempty"`
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
	//
	// example:
	//
	// Focus on identifying customer demands and to-do items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// my-agent
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The resource tags (optional, a JSON string list, such as ["tagA","tagB"]).
	//
	// example:
	//
	// ["Customer","GroupChat"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The feature update frequency.
	UpdateFrequencyShrink *string `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty"`
}

func (s CreateGroupAliDingChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingChatShrinkRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupAliDingChatShrinkRequest) GetChatName() *string {
	return s.ChatName
}

func (s *CreateGroupAliDingChatShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupAliDingChatShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupAliDingChatShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupAliDingChatShrinkRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreateGroupAliDingChatShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupAliDingChatShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupAliDingChatShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupAliDingChatShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupAliDingChatShrinkRequest) GetUpdateFrequencyShrink() *string {
	return s.UpdateFrequencyShrink
}

func (s *CreateGroupAliDingChatShrinkRequest) SetChatId(v string) *CreateGroupAliDingChatShrinkRequest {
	s.ChatId = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetChatName(v string) *CreateGroupAliDingChatShrinkRequest {
	s.ChatName = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetDescription(v string) *CreateGroupAliDingChatShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetDirectoryId(v string) *CreateGroupAliDingChatShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetGroupId(v string) *CreateGroupAliDingChatShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetHistoryStartTime(v string) *CreateGroupAliDingChatShrinkRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetNotes(v string) *CreateGroupAliDingChatShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetOperatingObjectName(v string) *CreateGroupAliDingChatShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetSourceTags(v string) *CreateGroupAliDingChatShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetTenantId(v string) *CreateGroupAliDingChatShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) SetUpdateFrequencyShrink(v string) *CreateGroupAliDingChatShrinkRequest {
	s.UpdateFrequencyShrink = &v
	return s
}

func (s *CreateGroupAliDingChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
