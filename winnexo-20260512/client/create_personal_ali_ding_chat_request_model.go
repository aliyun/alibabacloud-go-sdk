// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreatePersonalAliDingChatRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *CreatePersonalAliDingChatRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreatePersonalAliDingChatRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreatePersonalAliDingChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalAliDingChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreatePersonalAliDingChatRequestUpdateFrequency) *CreatePersonalAliDingChatRequest
	GetUpdateFrequency() *CreatePersonalAliDingChatRequestUpdateFrequency
}

type CreatePersonalAliDingChatRequest struct {
	// The platform application token issued by the console. If provided, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the DingTalk group chat session.
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
	// The content of meeting notes (optional), which is used for auxiliary analysis.
	//
	// example:
	//
	// Focus on identifying customer requirements and to-do items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the digital employee (the name of the operating object, optional).
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
	// The resource tags (optional). The value is a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// ["Customer","Group chat"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter and is explicitly passed by using the --tenant-id option in winnexo-cli.
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

func (s *CreatePersonalAliDingChatRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *CreatePersonalAliDingChatRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreatePersonalAliDingChatRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *CreatePersonalAliDingChatRequest) SetAppToken(v string) *CreatePersonalAliDingChatRequest {
	s.AppToken = &v
	return s
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

func (s *CreatePersonalAliDingChatRequest) SetOpsUserCode(v string) *CreatePersonalAliDingChatRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreatePersonalAliDingChatRequest) SetOpsUserId(v string) *CreatePersonalAliDingChatRequest {
	s.OpsUserId = &v
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
	// The cron expression for the timed scheduling node.
	//
	// example:
	//
	// 0 2 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// **Enable or disable**
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The preset mode, which can be ignored.
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
