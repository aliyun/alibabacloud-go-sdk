// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreatePersonalTextRequest
	GetAppToken() *string
	SetDescription(v string) *CreatePersonalTextRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalTextRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalTextRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalTextRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *CreatePersonalTextRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreatePersonalTextRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreatePersonalTextRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalTextRequest
	GetTenantId() *string
	SetTextContent(v string) *CreatePersonalTextRequest
	GetTextContent() *string
}

type CreatePersonalTextRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The pipeline description.
	//
	// example:
	//
	// PublicApplication
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_mysql_10_34_4_255_6306_password
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The list of resource tag JSON strings.
	//
	// example:
	//
	// ["Key","Text"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 3668
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The message content for text messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
}

func (s CreatePersonalTextRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTextRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalTextRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreatePersonalTextRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalTextRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalTextRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalTextRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalTextRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreatePersonalTextRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreatePersonalTextRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalTextRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalTextRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *CreatePersonalTextRequest) SetAppToken(v string) *CreatePersonalTextRequest {
	s.AppToken = &v
	return s
}

func (s *CreatePersonalTextRequest) SetDescription(v string) *CreatePersonalTextRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalTextRequest) SetDirectoryId(v string) *CreatePersonalTextRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalTextRequest) SetName(v string) *CreatePersonalTextRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalTextRequest) SetOperatingObjectName(v string) *CreatePersonalTextRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalTextRequest) SetOpsUserCode(v string) *CreatePersonalTextRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreatePersonalTextRequest) SetOpsUserId(v string) *CreatePersonalTextRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreatePersonalTextRequest) SetSourceTags(v string) *CreatePersonalTextRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalTextRequest) SetTenantId(v string) *CreatePersonalTextRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalTextRequest) SetTextContent(v string) *CreatePersonalTextRequest {
	s.TextContent = &v
	return s
}

func (s *CreatePersonalTextRequest) Validate() error {
	return dara.Validate(s)
}
