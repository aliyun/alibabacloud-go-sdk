// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreatePersonalAliDingChatShrinkRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *CreatePersonalAliDingChatShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreatePersonalAliDingChatShrinkRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreatePersonalAliDingChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalAliDingChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreatePersonalAliDingChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreatePersonalAliDingChatShrinkRequest struct {
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
	UpdateFrequencyShrink *string `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty"`
}

func (s CreatePersonalAliDingChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *CreatePersonalAliDingChatShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreatePersonalAliDingChatShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *CreatePersonalAliDingChatShrinkRequest) SetAppToken(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.AppToken = &v
	return s
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

func (s *CreatePersonalAliDingChatShrinkRequest) SetOpsUserCode(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreatePersonalAliDingChatShrinkRequest) SetOpsUserId(v string) *CreatePersonalAliDingChatShrinkRequest {
	s.OpsUserId = &v
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
