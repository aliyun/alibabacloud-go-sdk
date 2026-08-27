// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupAliDingChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupAliDingChatRequest
	GetChatId() *string
	SetChatName(v string) *CreateGroupAliDingChatRequest
	GetChatName() *string
	SetDescription(v string) *CreateGroupAliDingChatRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupAliDingChatRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupAliDingChatRequest
	GetGroupId() *string
	SetHistoryStartTime(v string) *CreateGroupAliDingChatRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreateGroupAliDingChatRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupAliDingChatRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupAliDingChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupAliDingChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreateGroupAliDingChatRequestUpdateFrequency) *CreateGroupAliDingChatRequest
	GetUpdateFrequency() *CreateGroupAliDingChatRequestUpdateFrequency
}

type CreateGroupAliDingChatRequest struct {
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
	// Customer group chat history
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The group ID.
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
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
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
	UpdateFrequency *CreateGroupAliDingChatRequestUpdateFrequency `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty" type:"Struct"`
}

func (s CreateGroupAliDingChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingChatRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupAliDingChatRequest) GetChatName() *string {
	return s.ChatName
}

func (s *CreateGroupAliDingChatRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupAliDingChatRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupAliDingChatRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupAliDingChatRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreateGroupAliDingChatRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupAliDingChatRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupAliDingChatRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupAliDingChatRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupAliDingChatRequest) GetUpdateFrequency() *CreateGroupAliDingChatRequestUpdateFrequency {
	return s.UpdateFrequency
}

func (s *CreateGroupAliDingChatRequest) SetChatId(v string) *CreateGroupAliDingChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetChatName(v string) *CreateGroupAliDingChatRequest {
	s.ChatName = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetDescription(v string) *CreateGroupAliDingChatRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetDirectoryId(v string) *CreateGroupAliDingChatRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetGroupId(v string) *CreateGroupAliDingChatRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetHistoryStartTime(v string) *CreateGroupAliDingChatRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetNotes(v string) *CreateGroupAliDingChatRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetOperatingObjectName(v string) *CreateGroupAliDingChatRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetSourceTags(v string) *CreateGroupAliDingChatRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetTenantId(v string) *CreateGroupAliDingChatRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupAliDingChatRequest) SetUpdateFrequency(v *CreateGroupAliDingChatRequestUpdateFrequency) *CreateGroupAliDingChatRequest {
	s.UpdateFrequency = v
	return s
}

func (s *CreateGroupAliDingChatRequest) Validate() error {
	if s.UpdateFrequency != nil {
		if err := s.UpdateFrequency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGroupAliDingChatRequestUpdateFrequency struct {
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

func (s CreateGroupAliDingChatRequestUpdateFrequency) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingChatRequestUpdateFrequency) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingChatRequestUpdateFrequency) GetCron() *string {
	return s.Cron
}

func (s *CreateGroupAliDingChatRequestUpdateFrequency) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateGroupAliDingChatRequestUpdateFrequency) GetPreset() *string {
	return s.Preset
}

func (s *CreateGroupAliDingChatRequestUpdateFrequency) SetCron(v string) *CreateGroupAliDingChatRequestUpdateFrequency {
	s.Cron = &v
	return s
}

func (s *CreateGroupAliDingChatRequestUpdateFrequency) SetEnabled(v bool) *CreateGroupAliDingChatRequestUpdateFrequency {
	s.Enabled = &v
	return s
}

func (s *CreateGroupAliDingChatRequestUpdateFrequency) SetPreset(v string) *CreateGroupAliDingChatRequestUpdateFrequency {
	s.Preset = &v
	return s
}

func (s *CreateGroupAliDingChatRequestUpdateFrequency) Validate() error {
	return dara.Validate(s)
}
