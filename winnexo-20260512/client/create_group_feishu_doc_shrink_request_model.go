// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateGroupFeishuDocShrinkRequest
	GetAppToken() *string
	SetDescription(v string) *CreateGroupFeishuDocShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFeishuDocShrinkRequest
	GetDirectoryId() *string
	SetDocUrl(v string) *CreateGroupFeishuDocShrinkRequest
	GetDocUrl() *string
	SetGroupId(v string) *CreateGroupFeishuDocShrinkRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupFeishuDocShrinkRequest
	GetName() *string
	SetNotes(v string) *CreateGroupFeishuDocShrinkRequest
	GetNotes() *string
	SetObjectBindingsShrink(v string) *CreateGroupFeishuDocShrinkRequest
	GetObjectBindingsShrink() *string
	SetOperatingObjectName(v string) *CreateGroupFeishuDocShrinkRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *CreateGroupFeishuDocShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateGroupFeishuDocShrinkRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreateGroupFeishuDocShrinkRequest
	GetSourceTags() *string
	SetSyncConfigShrink(v string) *CreateGroupFeishuDocShrinkRequest
	GetSyncConfigShrink() *string
	SetTenantId(v string) *CreateGroupFeishuDocShrinkRequest
	GetTenantId() *string
}

type CreateGroupFeishuDocShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is performed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the AI assistant.
	//
	// example:
	//
	// Group collaboration document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// dir_group_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The document URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.feishu.cn/docx/doxcnExample
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_delivery
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The mirror name.
	//
	// example:
	//
	// Project plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content. This parameter is optional and is used for auxiliary analysis.
	//
	// example:
	//
	// Extract decisions and to-dos
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The object bindings.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
	// The name of the operation object.
	//
	// example:
	//
	// R&D assistant
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
	// The resource tags. This parameter is optional and accepts a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// ["R&D"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The synchronization settings.
	SyncConfigShrink *string `json:"syncConfig,omitempty" xml:"syncConfig,omitempty"`
	// The tenant ID. This is a common parameter. For winnexo-cli, this parameter is explicitly passed by using the --tenant-id option.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupFeishuDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuDocShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateGroupFeishuDocShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFeishuDocShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuDocShrinkRequest) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreateGroupFeishuDocShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuDocShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupFeishuDocShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupFeishuDocShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *CreateGroupFeishuDocShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupFeishuDocShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateGroupFeishuDocShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateGroupFeishuDocShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFeishuDocShrinkRequest) GetSyncConfigShrink() *string {
	return s.SyncConfigShrink
}

func (s *CreateGroupFeishuDocShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFeishuDocShrinkRequest) SetAppToken(v string) *CreateGroupFeishuDocShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetDescription(v string) *CreateGroupFeishuDocShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetDirectoryId(v string) *CreateGroupFeishuDocShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetDocUrl(v string) *CreateGroupFeishuDocShrinkRequest {
	s.DocUrl = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetGroupId(v string) *CreateGroupFeishuDocShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetName(v string) *CreateGroupFeishuDocShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetNotes(v string) *CreateGroupFeishuDocShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetObjectBindingsShrink(v string) *CreateGroupFeishuDocShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetOperatingObjectName(v string) *CreateGroupFeishuDocShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetOpsUserCode(v string) *CreateGroupFeishuDocShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetOpsUserId(v string) *CreateGroupFeishuDocShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetSourceTags(v string) *CreateGroupFeishuDocShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetSyncConfigShrink(v string) *CreateGroupFeishuDocShrinkRequest {
	s.SyncConfigShrink = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetTenantId(v string) *CreateGroupFeishuDocShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}
