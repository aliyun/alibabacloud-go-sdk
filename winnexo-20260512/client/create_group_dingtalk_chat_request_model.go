// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDingtalkChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupDingtalkChatRequest
	GetChatId() *string
	SetChatName(v string) *CreateGroupDingtalkChatRequest
	GetChatName() *string
	SetDescription(v string) *CreateGroupDingtalkChatRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupDingtalkChatRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupDingtalkChatRequest
	GetGroupId() *string
	SetHistoryStartTime(v string) *CreateGroupDingtalkChatRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreateGroupDingtalkChatRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupDingtalkChatRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupDingtalkChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupDingtalkChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreateGroupDingtalkChatRequestUpdateFrequency) *CreateGroupDingtalkChatRequest
	GetUpdateFrequency() *CreateGroupDingtalkChatRequestUpdateFrequency
}

type CreateGroupDingtalkChatRequest struct {
	// The conversation ID, typically used for JSSDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxxxx
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The chat name.
	//
	// example:
	//
	// Customer Project Chat
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
	// Focus on identifying customer requests and to-do items
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
	// ["Customer","Chat"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The feature update frequency.
	UpdateFrequency *CreateGroupDingtalkChatRequestUpdateFrequency `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty" type:"Struct"`
}

func (s CreateGroupDingtalkChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupDingtalkChatRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupDingtalkChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupDingtalkChatRequest) GetChatName() *string {
	return s.ChatName
}

func (s *CreateGroupDingtalkChatRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupDingtalkChatRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupDingtalkChatRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupDingtalkChatRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreateGroupDingtalkChatRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupDingtalkChatRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupDingtalkChatRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupDingtalkChatRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupDingtalkChatRequest) GetUpdateFrequency() *CreateGroupDingtalkChatRequestUpdateFrequency {
	return s.UpdateFrequency
}

func (s *CreateGroupDingtalkChatRequest) SetChatId(v string) *CreateGroupDingtalkChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetChatName(v string) *CreateGroupDingtalkChatRequest {
	s.ChatName = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetDescription(v string) *CreateGroupDingtalkChatRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetDirectoryId(v string) *CreateGroupDingtalkChatRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetGroupId(v string) *CreateGroupDingtalkChatRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetHistoryStartTime(v string) *CreateGroupDingtalkChatRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetNotes(v string) *CreateGroupDingtalkChatRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetOperatingObjectName(v string) *CreateGroupDingtalkChatRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetSourceTags(v string) *CreateGroupDingtalkChatRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetTenantId(v string) *CreateGroupDingtalkChatRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetUpdateFrequency(v *CreateGroupDingtalkChatRequestUpdateFrequency) *CreateGroupDingtalkChatRequest {
	s.UpdateFrequency = v
	return s
}

func (s *CreateGroupDingtalkChatRequest) Validate() error {
	if s.UpdateFrequency != nil {
		if err := s.UpdateFrequency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGroupDingtalkChatRequestUpdateFrequency struct {
	// The cron expression for timed scheduling.
	//
	// example:
	//
	// 0 2 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// Specifies whether the throttling rule is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The preset mode. You can ignore this parameter.
	//
	// example:
	//
	// hourly
	Preset *string `json:"preset,omitempty" xml:"preset,omitempty"`
}

func (s CreateGroupDingtalkChatRequestUpdateFrequency) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupDingtalkChatRequestUpdateFrequency) GoString() string {
	return s.String()
}

func (s *CreateGroupDingtalkChatRequestUpdateFrequency) GetCron() *string {
	return s.Cron
}

func (s *CreateGroupDingtalkChatRequestUpdateFrequency) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateGroupDingtalkChatRequestUpdateFrequency) GetPreset() *string {
	return s.Preset
}

func (s *CreateGroupDingtalkChatRequestUpdateFrequency) SetCron(v string) *CreateGroupDingtalkChatRequestUpdateFrequency {
	s.Cron = &v
	return s
}

func (s *CreateGroupDingtalkChatRequestUpdateFrequency) SetEnabled(v bool) *CreateGroupDingtalkChatRequestUpdateFrequency {
	s.Enabled = &v
	return s
}

func (s *CreateGroupDingtalkChatRequestUpdateFrequency) SetPreset(v string) *CreateGroupDingtalkChatRequestUpdateFrequency {
	s.Preset = &v
	return s
}

func (s *CreateGroupDingtalkChatRequestUpdateFrequency) Validate() error {
	return dara.Validate(s)
}
