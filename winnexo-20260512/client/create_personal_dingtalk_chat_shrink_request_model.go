// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreatePersonalDingtalkChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreatePersonalDingtalkChatShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The DingTalk group chat session ID.
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
	// Customer Project Group
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
	// The start time for chat record collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-01
	HistoryStartTime *string `json:"historyStartTime,omitempty" xml:"historyStartTime,omitempty"`
	// The content of the meeting notes (optional), which is used for auxiliary analysis.
	//
	// example:
	//
	// Focus on identifying customer demands and to-do items
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
	// The ID of the tenant.
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

func (s *CreatePersonalDingtalkChatShrinkRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *CreatePersonalDingtalkChatShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreatePersonalDingtalkChatShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *CreatePersonalDingtalkChatShrinkRequest) SetAppToken(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.AppToken = &v
	return s
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

func (s *CreatePersonalDingtalkChatShrinkRequest) SetOpsUserCode(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreatePersonalDingtalkChatShrinkRequest) SetOpsUserId(v string) *CreatePersonalDingtalkChatShrinkRequest {
	s.OpsUserId = &v
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
