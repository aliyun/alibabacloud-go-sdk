// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDingtalkChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupDingtalkChatShrinkRequest
	GetChatId() *string
	SetChatName(v string) *CreateGroupDingtalkChatShrinkRequest
	GetChatName() *string
	SetDescription(v string) *CreateGroupDingtalkChatShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupDingtalkChatShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupDingtalkChatShrinkRequest
	GetGroupId() *string
	SetHistoryStartTime(v string) *CreateGroupDingtalkChatShrinkRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreateGroupDingtalkChatShrinkRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupDingtalkChatShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupDingtalkChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupDingtalkChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreateGroupDingtalkChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreateGroupDingtalkChatShrinkRequest struct {
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
	// Customer chat history
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The project group ID.
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
	// The meeting notes content (optional). This participates in auxiliary analysis.
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
	// The source tags.
	//
	// example:
	//
	// ["Customer","GroupChat"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The feature update frequency.
	UpdateFrequencyShrink *string `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty"`
}

func (s CreateGroupDingtalkChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupDingtalkChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetChatName() *string {
	return s.ChatName
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupDingtalkChatShrinkRequest) GetUpdateFrequencyShrink() *string {
	return s.UpdateFrequencyShrink
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetChatId(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.ChatId = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetChatName(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.ChatName = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetDescription(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetDirectoryId(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetGroupId(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetHistoryStartTime(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetNotes(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetOperatingObjectName(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetSourceTags(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetTenantId(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) SetUpdateFrequencyShrink(v string) *CreateGroupDingtalkChatShrinkRequest {
	s.UpdateFrequencyShrink = &v
	return s
}

func (s *CreateGroupDingtalkChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
