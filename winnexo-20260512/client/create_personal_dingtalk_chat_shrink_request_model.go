// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetChatId() *string
	SetChatName(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetChatName() *string
	SetDescription(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetDirectoryId() *string
	SetHistoryStartTime(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreatePersonalDingtalkChatShrinkRequest struct {
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
	// The directory ID.
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

func (s CreatePersonalDingtalkChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetChatName() *string {
	return s.ChatName
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetUpdateFrequencyShrink() *string {
	return s.UpdateFrequencyShrink
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetChatId(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetChatName(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.ChatName = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetDescription(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetDirectoryId(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetHistoryStartTime(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetNotes(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetOperatingObjectName(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetSourceTags(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetTenantId(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetUpdateFrequencyShrink(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.UpdateFrequencyShrink = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
