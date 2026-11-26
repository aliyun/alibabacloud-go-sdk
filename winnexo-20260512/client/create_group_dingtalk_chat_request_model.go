// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDingtalkChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateGroupDingtalkChatRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *CreateGroupDingtalkChatRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateGroupDingtalkChatRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreateGroupDingtalkChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupDingtalkChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreateGroupDingtalkChatRequestUpdateFrequency) *CreateGroupDingtalkChatRequest
	GetUpdateFrequency() *CreateGroupDingtalkChatRequestUpdateFrequency
}

type CreateGroupDingtalkChatRequest struct {
	// The platform application token issued by the console. If specified, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The chat ID, which is generally used for JSSDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxxxx
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The name of the group chat.
	//
	// example:
	//
	// Customer project group
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
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The start time for collecting chat records.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-01
	HistoryStartTime *string `json:"historyStartTime,omitempty" xml:"historyStartTime,omitempty"`
	// The content of the meeting notes. This parameter is optional and is used for auxiliary analysis.
	//
	// example:
	//
	// Focus on identifying customer demands and to-do items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the digital employee, which is the name of the operating object. This parameter is optional.
	//
	// example:
	//
	// my-agent
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The source tags.
	//
	// example:
	//
	// ["Customer","Group chat"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. For winnexo-cli, explicitly pass this parameter by using --tenant-id.
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

func (s *CreateGroupDingtalkChatRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *CreateGroupDingtalkChatRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateGroupDingtalkChatRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *CreateGroupDingtalkChatRequest) SetAppToken(v string) *CreateGroupDingtalkChatRequest {
	s.AppToken = &v
	return s
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

func (s *CreateGroupDingtalkChatRequest) SetOpsUserCode(v string) *CreateGroupDingtalkChatRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateGroupDingtalkChatRequest) SetOpsUserId(v string) *CreateGroupDingtalkChatRequest {
	s.OpsUserId = &v
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
	// The cron expression for the timed scheduling node.
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
