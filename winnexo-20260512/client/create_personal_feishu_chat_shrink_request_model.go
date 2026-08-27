// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalFeishuChatShrinkRequest
	GetChatId() *string
	SetDescription(v string) *CreatePersonalFeishuChatShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalFeishuChatShrinkRequest
	GetDirectoryId() *string
	SetHistoryStartTime(v string) *CreatePersonalFeishuChatShrinkRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreatePersonalFeishuChatShrinkRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalFeishuChatShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreatePersonalFeishuChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalFeishuChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreatePersonalFeishuChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreatePersonalFeishuChatShrinkRequest struct {
	// The group chat session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// oc_abc123
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The description of the source.
	//
	// example:
	//
	// Product R&D group chat records
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// dir_personal_1
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The start time for historical messages. Supports YYYY-MM-DD or YYYY-MM-DD HH:MM:SS. If not specified, all visible history is pulled.
	//
	// example:
	//
	// 2026-08-01 00:00:00
	HistoryStartTime *string `json:"historyStartTime,omitempty" xml:"historyStartTime,omitempty"`
	// The meeting notes content (optional). Used for auxiliary analysis.
	//
	// example:
	//
	// Focus on extracting decisions and action items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// R&D Assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The source tags.
	//
	// example:
	//
	// ["R&D"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID to take effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The update frequency.
	UpdateFrequencyShrink *string `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty"`
}

func (s CreatePersonalFeishuChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalFeishuChatShrinkRequest) GetUpdateFrequencyShrink() *string {
	return s.UpdateFrequencyShrink
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetChatId(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetDescription(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetDirectoryId(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetHistoryStartTime(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetNotes(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetOperatingObjectName(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetSourceTags(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetTenantId(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) SetUpdateFrequencyShrink(v string) *CreatePersonalFeishuChatShrinkRequest {
	s.UpdateFrequencyShrink = &v
	return s
}

func (s *CreatePersonalFeishuChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
