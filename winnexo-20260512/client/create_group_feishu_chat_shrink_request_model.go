// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupFeishuChatShrinkRequest
	GetChatId() *string
	SetDescription(v string) *CreateGroupFeishuChatShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFeishuChatShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupFeishuChatShrinkRequest
	GetGroupId() *string
	SetHistoryStartTime(v string) *CreateGroupFeishuChatShrinkRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreateGroupFeishuChatShrinkRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupFeishuChatShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupFeishuChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupFeishuChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreateGroupFeishuChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreateGroupFeishuChatShrinkRequest struct {
	// The DingTalk group chat session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxxxx
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The pipeline description.
	//
	// example:
	//
	// string_value
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
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
	// The start time for historical messages. The value must be in the YYYY-MM-DD or YYYY-MM-DD HH:MM:SS format. If this parameter is not specified, all visible historical messages are retrieved.
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
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The source tags.
	//
	// example:
	//
	// ["Key","File"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. You can pass it explicitly by using --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The feature update frequency.
	UpdateFrequencyShrink *string `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty"`
}

func (s CreateGroupFeishuChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuChatShrinkRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupFeishuChatShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFeishuChatShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuChatShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuChatShrinkRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreateGroupFeishuChatShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupFeishuChatShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupFeishuChatShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFeishuChatShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFeishuChatShrinkRequest) GetUpdateFrequencyShrink() *string {
	return s.UpdateFrequencyShrink
}

func (s *CreateGroupFeishuChatShrinkRequest) SetChatId(v string) *CreateGroupFeishuChatShrinkRequest {
	s.ChatId = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetDescription(v string) *CreateGroupFeishuChatShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetDirectoryId(v string) *CreateGroupFeishuChatShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetGroupId(v string) *CreateGroupFeishuChatShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetHistoryStartTime(v string) *CreateGroupFeishuChatShrinkRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetNotes(v string) *CreateGroupFeishuChatShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetOperatingObjectName(v string) *CreateGroupFeishuChatShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetSourceTags(v string) *CreateGroupFeishuChatShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetTenantId(v string) *CreateGroupFeishuChatShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetUpdateFrequencyShrink(v string) *CreateGroupFeishuChatShrinkRequest {
	s.UpdateFrequencyShrink = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
