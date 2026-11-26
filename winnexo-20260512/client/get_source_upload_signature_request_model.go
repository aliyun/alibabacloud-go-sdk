// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceUploadSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetSourceUploadSignatureRequest
	GetAppToken() *string
	SetContentType(v string) *GetSourceUploadSignatureRequest
	GetContentType() *string
	SetExpires(v int64) *GetSourceUploadSignatureRequest
	GetExpires() *int64
	SetFilename(v string) *GetSourceUploadSignatureRequest
	GetFilename() *string
	SetGroupId(v string) *GetSourceUploadSignatureRequest
	GetGroupId() *string
	SetOperatingObjectName(v string) *GetSourceUploadSignatureRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *GetSourceUploadSignatureRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetSourceUploadSignatureRequest
	GetOpsUserId() *string
	SetScope(v string) *GetSourceUploadSignatureRequest
	GetScope() *string
	SetTenantId(v string) *GetSourceUploadSignatureRequest
	GetTenantId() *string
}

type GetSourceUploadSignatureRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The content type. Valid values:
	//
	// - Text
	//
	// - Markdown
	//
	// example:
	//
	// string_value
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The expiration time of the signed URL, in seconds. Default value: 3600.
	//
	// example:
	//
	// 3600
	Expires *int64 `json:"expires,omitempty" xml:"expires,omitempty"`
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	Filename *string `json:"filename,omitempty" xml:"filename,omitempty"`
	// The collaboration space ID passed from the frontend. This parameter is required when scope is set to group. The value must be a space accessible to the current user and is used in the OSS path for locating. The value must be 1 to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-). This parameter is not used for other scope values.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
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
	// The permission scope.
	//
	// example:
	//
	// source
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The tenant ID to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSourceUploadSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceUploadSignatureRequest) GoString() string {
	return s.String()
}

func (s *GetSourceUploadSignatureRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetSourceUploadSignatureRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GetSourceUploadSignatureRequest) GetExpires() *int64 {
	return s.Expires
}

func (s *GetSourceUploadSignatureRequest) GetFilename() *string {
	return s.Filename
}

func (s *GetSourceUploadSignatureRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetSourceUploadSignatureRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *GetSourceUploadSignatureRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetSourceUploadSignatureRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetSourceUploadSignatureRequest) GetScope() *string {
	return s.Scope
}

func (s *GetSourceUploadSignatureRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSourceUploadSignatureRequest) SetAppToken(v string) *GetSourceUploadSignatureRequest {
	s.AppToken = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetContentType(v string) *GetSourceUploadSignatureRequest {
	s.ContentType = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetExpires(v int64) *GetSourceUploadSignatureRequest {
	s.Expires = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetFilename(v string) *GetSourceUploadSignatureRequest {
	s.Filename = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetGroupId(v string) *GetSourceUploadSignatureRequest {
	s.GroupId = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetOperatingObjectName(v string) *GetSourceUploadSignatureRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetOpsUserCode(v string) *GetSourceUploadSignatureRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetOpsUserId(v string) *GetSourceUploadSignatureRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetScope(v string) *GetSourceUploadSignatureRequest {
	s.Scope = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetTenantId(v string) *GetSourceUploadSignatureRequest {
	s.TenantId = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) Validate() error {
	return dara.Validate(s)
}
