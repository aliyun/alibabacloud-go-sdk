// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalAliDingChatShrinkRequest
	GetChatId() *string
	SetChatName(v string) *CreatePersonalAliDingChatShrinkRequest
	GetChatName() *string
	SetDescription(v string) *CreatePersonalAliDingChatShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalAliDingChatShrinkRequest
	GetDirectoryId() *string
	SetHistoryStartTime(v string) *CreatePersonalAliDingChatShrinkRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreatePersonalAliDingChatShrinkRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalAliDingChatShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreatePersonalAliDingChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalAliDingChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreatePersonalAliDingChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreatePersonalAliDingChatShrinkRequest struct {
	// The DingTalk group chat session ID.
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
	// The pipeline description.
	//
	// example:
	//
	// Customer group chat history
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
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
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// my-agent
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The resource tags (optional, a JSON string list such as ["tagA","tagB"]).
	//
	// example:
	//
	// ["Customer","GroupChat"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. The winnexo-cli passes this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The feature update frequency.
	UpdateFrequencyShrink *string `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty"`
}

func (s CreatePersonalAliDingChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetChatName() *string {
	return s.ChatName
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetUpdateFrequencyShrink() *string {
	return s.UpdateFrequencyShrink
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetChatId(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetChatName(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.ChatName = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetDescription(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetDirectoryId(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetHistoryStartTime(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetNotes(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetOperatingObjectName(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetSourceTags(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetTenantId(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetUpdateFrequencyShrink(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.UpdateFrequencyShrink = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
