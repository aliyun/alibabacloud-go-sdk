// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalFeishuChatRequest
	GetChatId() *string
	SetDescription(v string) *CreatePersonalFeishuChatRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalFeishuChatRequest
	GetDirectoryId() *string
	SetHistoryStartTime(v string) *CreatePersonalFeishuChatRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreatePersonalFeishuChatRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalFeishuChatRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreatePersonalFeishuChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalFeishuChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreatePersonalFeishuChatRequestUpdateFrequency) *CreatePersonalFeishuChatRequest
	GetUpdateFrequency() *CreatePersonalFeishuChatRequestUpdateFrequency
}

type CreatePersonalFeishuChatRequest struct {
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
	UpdateFrequency *CreatePersonalFeishuChatRequestUpdateFrequency `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty" type:"Struct"`
}

func (s CreatePersonalFeishuChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuChatRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalFeishuChatRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalFeishuChatRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuChatRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreatePersonalFeishuChatRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalFeishuChatRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalFeishuChatRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalFeishuChatRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalFeishuChatRequest) GetUpdateFrequency() *CreatePersonalFeishuChatRequestUpdateFrequency {
	return s.UpdateFrequency
}

func (s *CreatePersonalFeishuChatRequest) SetChatId(v string) *CreatePersonalFeishuChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalFeishuChatRequest) SetDescription(v string) *CreatePersonalFeishuChatRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalFeishuChatRequest) SetDirectoryId(v string) *CreatePersonalFeishuChatRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuChatRequest) SetHistoryStartTime(v string) *CreatePersonalFeishuChatRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreatePersonalFeishuChatRequest) SetNotes(v string) *CreatePersonalFeishuChatRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalFeishuChatRequest) SetOperatingObjectName(v string) *CreatePersonalFeishuChatRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalFeishuChatRequest) SetSourceTags(v string) *CreatePersonalFeishuChatRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalFeishuChatRequest) SetTenantId(v string) *CreatePersonalFeishuChatRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalFeishuChatRequest) SetUpdateFrequency(v *CreatePersonalFeishuChatRequestUpdateFrequency) *CreatePersonalFeishuChatRequest {
	s.UpdateFrequency = v
	return s
}

func (s *CreatePersonalFeishuChatRequest) Validate() error {
	if s.UpdateFrequency != nil {
		if err := s.UpdateFrequency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePersonalFeishuChatRequestUpdateFrequency struct {
	// The cron expression for the timed scheduling node.
	//
	// example:
	//
	// 0 	- 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// Specifies whether to enable the scheduled synchronization.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The synchronization preset: hourly or daily_2am.
	//
	// example:
	//
	// hourly
	Preset *string `json:"preset,omitempty" xml:"preset,omitempty"`
}

func (s CreatePersonalFeishuChatRequestUpdateFrequency) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuChatRequestUpdateFrequency) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuChatRequestUpdateFrequency) GetCron() *string {
	return s.Cron
}

func (s *CreatePersonalFeishuChatRequestUpdateFrequency) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreatePersonalFeishuChatRequestUpdateFrequency) GetPreset() *string {
	return s.Preset
}

func (s *CreatePersonalFeishuChatRequestUpdateFrequency) SetCron(v string) *CreatePersonalFeishuChatRequestUpdateFrequency {
	s.Cron = &v
	return s
}

func (s *CreatePersonalFeishuChatRequestUpdateFrequency) SetEnabled(v bool) *CreatePersonalFeishuChatRequestUpdateFrequency {
	s.Enabled = &v
	return s
}

func (s *CreatePersonalFeishuChatRequestUpdateFrequency) SetPreset(v string) *CreatePersonalFeishuChatRequestUpdateFrequency {
	s.Preset = &v
	return s
}

func (s *CreatePersonalFeishuChatRequestUpdateFrequency) Validate() error {
	return dara.Validate(s)
}
