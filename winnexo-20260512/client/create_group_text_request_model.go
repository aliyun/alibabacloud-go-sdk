// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateGroupTextRequest
	GetAppToken() *string
	SetDescription(v string) *CreateGroupTextRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupTextRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupTextRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupTextRequest
	GetName() *string
	SetOpsUserCode(v string) *CreateGroupTextRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateGroupTextRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreateGroupTextRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupTextRequest
	GetTenantId() *string
	SetTextContent(v string) *CreateGroupTextRequest
	GetTextContent() *string
}

type CreateGroupTextRequest struct {
	// The platform application token issued by the console. If this parameter is specified, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the AI assistant.
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The mirror name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Project materials
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// example
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The message content for text messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
}

func (s CreateGroupTextRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupTextRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupTextRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateGroupTextRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupTextRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupTextRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupTextRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupTextRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateGroupTextRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateGroupTextRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupTextRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupTextRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *CreateGroupTextRequest) SetAppToken(v string) *CreateGroupTextRequest {
	s.AppToken = &v
	return s
}

func (s *CreateGroupTextRequest) SetDescription(v string) *CreateGroupTextRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupTextRequest) SetDirectoryId(v string) *CreateGroupTextRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupTextRequest) SetGroupId(v string) *CreateGroupTextRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupTextRequest) SetName(v string) *CreateGroupTextRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupTextRequest) SetOpsUserCode(v string) *CreateGroupTextRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateGroupTextRequest) SetOpsUserId(v string) *CreateGroupTextRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateGroupTextRequest) SetSourceTags(v string) *CreateGroupTextRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupTextRequest) SetTenantId(v string) *CreateGroupTextRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupTextRequest) SetTextContent(v string) *CreateGroupTextRequest {
	s.TextContent = &v
	return s
}

func (s *CreateGroupTextRequest) Validate() error {
	return dara.Validate(s)
}
