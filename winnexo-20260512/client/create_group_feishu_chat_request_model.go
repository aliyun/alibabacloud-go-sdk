// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupFeishuChatRequest
	GetChatId() *string
	SetDescription(v string) *CreateGroupFeishuChatRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFeishuChatRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupFeishuChatRequest
	GetGroupId() *string
	SetHistoryStartTime(v string) *CreateGroupFeishuChatRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreateGroupFeishuChatRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupFeishuChatRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupFeishuChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupFeishuChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreateGroupFeishuChatRequestUpdateFrequency) *CreateGroupFeishuChatRequest
	GetUpdateFrequency() *CreateGroupFeishuChatRequestUpdateFrequency
}

type CreateGroupFeishuChatRequest struct {
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
	UpdateFrequency *CreateGroupFeishuChatRequestUpdateFrequency `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty" type:"Struct"`
}

func (s CreateGroupFeishuChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuChatRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupFeishuChatRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFeishuChatRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuChatRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuChatRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreateGroupFeishuChatRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupFeishuChatRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupFeishuChatRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFeishuChatRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFeishuChatRequest) GetUpdateFrequency() *CreateGroupFeishuChatRequestUpdateFrequency {
	return s.UpdateFrequency
}

func (s *CreateGroupFeishuChatRequest) SetChatId(v string) *CreateGroupFeishuChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetDescription(v string) *CreateGroupFeishuChatRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetDirectoryId(v string) *CreateGroupFeishuChatRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetGroupId(v string) *CreateGroupFeishuChatRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetHistoryStartTime(v string) *CreateGroupFeishuChatRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetNotes(v string) *CreateGroupFeishuChatRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetOperatingObjectName(v string) *CreateGroupFeishuChatRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetSourceTags(v string) *CreateGroupFeishuChatRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetTenantId(v string) *CreateGroupFeishuChatRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetUpdateFrequency(v *CreateGroupFeishuChatRequestUpdateFrequency) *CreateGroupFeishuChatRequest {
	s.UpdateFrequency = v
	return s
}

func (s *CreateGroupFeishuChatRequest) Validate() error {
	if s.UpdateFrequency != nil {
		if err := s.UpdateFrequency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGroupFeishuChatRequestUpdateFrequency struct {
	// The cron expression for the timed scheduling task.
	//
	// example:
	//
	// 0 2 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// **Enable/Disable**
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

func (s CreateGroupFeishuChatRequestUpdateFrequency) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuChatRequestUpdateFrequency) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) GetCron() *string {
	return s.Cron
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) GetPreset() *string {
	return s.Preset
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) SetCron(v string) *CreateGroupFeishuChatRequestUpdateFrequency {
	s.Cron = &v
	return s
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) SetEnabled(v bool) *CreateGroupFeishuChatRequestUpdateFrequency {
	s.Enabled = &v
	return s
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) SetPreset(v string) *CreateGroupFeishuChatRequestUpdateFrequency {
	s.Preset = &v
	return s
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) Validate() error {
	return dara.Validate(s)
}
