// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadChatFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UploadChatFileRequest
	GetAppToken() *string
	SetContentType(v string) *UploadChatFileRequest
	GetContentType() *string
	SetFileName(v string) *UploadChatFileRequest
	GetFileName() *string
	SetFileUrl(v string) *UploadChatFileRequest
	GetFileUrl() *string
	SetOperatingObjectName(v string) *UploadChatFileRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *UploadChatFileRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UploadChatFileRequest
	GetOpsUserId() *string
	SetTenantId(v string) *UploadChatFileRequest
	GetTenantId() *string
}

type UploadChatFileRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The content type of the file. Valid values:
	//
	// - **image**: Image.
	//
	// - **document**: General document.
	//
	// - **alidoc**: Alibaba document.
	//
	// - **text**: Text.
	//
	// - **video**: Video.
	//
	// - **audio**: Audio.
	//
	// - **archive**: Archive.
	//
	// - **app**: Application.
	//
	// - **link**: Shortcut.
	//
	// - **other**: Other.
	//
	// example:
	//
	// application/pdf
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The full path name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// report.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The Yida attachment address.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://winnexo-file-transfer.oss-cn-hangzhou.aliyuncs.com/openapi/2026-08-06/9f8c2a1b
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// The name of the digital employee (operating object name). This parameter is optional.
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
	// The tenant ID that takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UploadChatFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileRequest) GoString() string {
	return s.String()
}

func (s *UploadChatFileRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UploadChatFileRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UploadChatFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadChatFileRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadChatFileRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *UploadChatFileRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UploadChatFileRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UploadChatFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadChatFileRequest) SetAppToken(v string) *UploadChatFileRequest {
	s.AppToken = &v
	return s
}

func (s *UploadChatFileRequest) SetContentType(v string) *UploadChatFileRequest {
	s.ContentType = &v
	return s
}

func (s *UploadChatFileRequest) SetFileName(v string) *UploadChatFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadChatFileRequest) SetFileUrl(v string) *UploadChatFileRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadChatFileRequest) SetOperatingObjectName(v string) *UploadChatFileRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *UploadChatFileRequest) SetOpsUserCode(v string) *UploadChatFileRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UploadChatFileRequest) SetOpsUserId(v string) *UploadChatFileRequest {
	s.OpsUserId = &v
	return s
}

func (s *UploadChatFileRequest) SetTenantId(v string) *UploadChatFileRequest {
	s.TenantId = &v
	return s
}

func (s *UploadChatFileRequest) Validate() error {
	return dara.Validate(s)
}
