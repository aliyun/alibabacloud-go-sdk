// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalAliDingChatRequest
	GetChatId() *string
	SetChatName(v string) *CreatePersonalAliDingChatRequest
	GetChatName() *string
	SetDescription(v string) *CreatePersonalAliDingChatRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalAliDingChatRequest
	GetDirectoryId() *string
	SetHistoryStartTime(v string) *CreatePersonalAliDingChatRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreatePersonalAliDingChatRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalAliDingChatRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreatePersonalAliDingChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalAliDingChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreatePersonalAliDingChatRequestUpdateFrequency) *CreatePersonalAliDingChatRequest
	GetUpdateFrequency() *CreatePersonalAliDingChatRequestUpdateFrequency
}

type CreatePersonalAliDingChatRequest struct {
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
	UpdateFrequency *CreatePersonalAliDingChatRequestUpdateFrequency `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty" type:"Struct"`
}

func (s CreatePersonalAliDingChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingChatRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalAliDingChatRequest) GetChatName() *string {
	return s.ChatName
}

func (s *CreatePersonalAliDingChatRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalAliDingChatRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAliDingChatRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreatePersonalAliDingChatRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalAliDingChatRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAliDingChatRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalAliDingChatRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalAliDingChatRequest) GetUpdateFrequency() *CreatePersonalAliDingChatRequestUpdateFrequency {
	return s.UpdateFrequency
}

func (s *CreatePersonalAliDingChatRequest) SetChatId(v string) *CreatePersonalAliDingChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetChatName(v string) *CreatePersonalAliDingChatRequest {
	s.ChatName = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetDescription(v string) *CreatePersonalAliDingChatRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetDirectoryId(v string) *CreatePersonalAliDingChatRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetHistoryStartTime(v string) *CreatePersonalAliDingChatRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetNotes(v string) *CreatePersonalAliDingChatRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetOperatingObjectName(v string) *CreatePersonalAliDingChatRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetSourceTags(v string) *CreatePersonalAliDingChatRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetTenantId(v string) *CreatePersonalAliDingChatRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetUpdateFrequency(v *CreatePersonalAliDingChatRequestUpdateFrequency) *CreatePersonalAliDingChatRequest {
	s.UpdateFrequency = v
	return s
}

func (s *CreatePersonalAliDingChatRequest) Validate() error {
	if s.UpdateFrequency != nil {
		if err := s.UpdateFrequency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePersonalAliDingChatRequestUpdateFrequency struct {
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

func (s CreatePersonalAliDingChatRequestUpdateFrequency) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingChatRequestUpdateFrequency) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingChatRequestUpdateFrequency) GetCron() *string {
	return s.Cron
}

func (s *CreatePersonalAliDingChatRequestUpdateFrequency) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreatePersonalAliDingChatRequestUpdateFrequency) GetPreset() *string {
	return s.Preset
}

func (s *CreatePersonalAliDingChatRequestUpdateFrequency) SetCron(v string) *CreatePersonalAliDingChatRequestUpdateFrequency {
	s.Cron = &v
	return s
}

func (s *CreatePersonalAliDingChatRequestUpdateFrequency) SetEnabled(v bool) *CreatePersonalAliDingChatRequestUpdateFrequency {
	s.Enabled = &v
	return s
}

func (s *CreatePersonalAliDingChatRequestUpdateFrequency) SetPreset(v string) *CreatePersonalAliDingChatRequestUpdateFrequency {
	s.Preset = &v
	return s
}

func (s *CreatePersonalAliDingChatRequestUpdateFrequency) Validate() error {
	return dara.Validate(s)
}
