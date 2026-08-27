// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalDingtalkChatRequest
	GetChatId() *string
	SetChatName(v string) *CreatePersonalDingtalkChatRequest
	GetChatName() *string
	SetDescription(v string) *CreatePersonalDingtalkChatRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkChatRequest
	GetDirectoryId() *string
	SetHistoryStartTime(v string) *CreatePersonalDingtalkChatRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreatePersonalDingtalkChatRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalDingtalkChatRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreatePersonalDingtalkChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalDingtalkChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreatePersonalDingtalkChatRequestUpdateFrequency) *CreatePersonalDingtalkChatRequest
	GetUpdateFrequency() *CreatePersonalDingtalkChatRequestUpdateFrequency
}

type CreatePersonalDingtalkChatRequest struct {
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
	UpdateFrequency *CreatePersonalDingtalkChatRequestUpdateFrequency `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty" type:"Struct"`
}

func (s CreatePersonalDingtalkChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkChatRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalDingtalkChatRequest) GetChatName() *string {
	return s.ChatName
}

func (s *CreatePersonalDingtalkChatRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDingtalkChatRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkChatRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreatePersonalDingtalkChatRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalDingtalkChatRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDingtalkChatRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalDingtalkChatRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDingtalkChatRequest) GetUpdateFrequency() *CreatePersonalDingtalkChatRequestUpdateFrequency {
	return s.UpdateFrequency
}

func (s *CreatePersonalDingtalkChatRequest) SetChatId(v string) *CreatePersonalDingtalkChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetChatName(v string) *CreatePersonalDingtalkChatRequest {
	s.ChatName = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetDescription(v string) *CreatePersonalDingtalkChatRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetDirectoryId(v string) *CreatePersonalDingtalkChatRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetHistoryStartTime(v string) *CreatePersonalDingtalkChatRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetNotes(v string) *CreatePersonalDingtalkChatRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetOperatingObjectName(v string) *CreatePersonalDingtalkChatRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetSourceTags(v string) *CreatePersonalDingtalkChatRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetTenantId(v string) *CreatePersonalDingtalkChatRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) SetUpdateFrequency(v *CreatePersonalDingtalkChatRequestUpdateFrequency) *CreatePersonalDingtalkChatRequest {
	s.UpdateFrequency = v
	return s
}

func (s *CreatePersonalDingtalkChatRequest) Validate() error {
	if s.UpdateFrequency != nil {
		if err := s.UpdateFrequency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePersonalDingtalkChatRequestUpdateFrequency struct {
	// The cron expression for timed scheduling.
	//
	// example:
	//
	// 0 2 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// Specifies whether to enable or disable the feature.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The preset mode (can be ignored).
	//
	// example:
	//
	// hourly
	Preset *string `json:"preset,omitempty" xml:"preset,omitempty"`
}

func (s CreatePersonalDingtalkChatRequestUpdateFrequency) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkChatRequestUpdateFrequency) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkChatRequestUpdateFrequency) GetCron() *string {
	return s.Cron
}

func (s *CreatePersonalDingtalkChatRequestUpdateFrequency) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreatePersonalDingtalkChatRequestUpdateFrequency) GetPreset() *string {
	return s.Preset
}

func (s *CreatePersonalDingtalkChatRequestUpdateFrequency) SetCron(v string) *CreatePersonalDingtalkChatRequestUpdateFrequency {
	s.Cron = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequestUpdateFrequency) SetEnabled(v bool) *CreatePersonalDingtalkChatRequestUpdateFrequency {
	s.Enabled = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequestUpdateFrequency) SetPreset(v string) *CreatePersonalDingtalkChatRequestUpdateFrequency {
	s.Preset = &v
	return s
}

func (s *CreatePersonalDingtalkChatRequestUpdateFrequency) Validate() error {
	return dara.Validate(s)
}
