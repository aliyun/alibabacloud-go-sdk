// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadChatFileByBase64Request interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UploadChatFileByBase64Request
	GetAppToken() *string
	SetContentType(v string) *UploadChatFileByBase64Request
	GetContentType() *string
	SetFileContentBase64(v string) *UploadChatFileByBase64Request
	GetFileContentBase64() *string
	SetFileName(v string) *UploadChatFileByBase64Request
	GetFileName() *string
	SetOpsUserCode(v string) *UploadChatFileByBase64Request
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UploadChatFileByBase64Request
	GetOpsUserId() *string
	SetTenantId(v string) *UploadChatFileByBase64Request
	GetTenantId() *string
}

type UploadChatFileByBase64Request struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The MIME type of the file.
	//
	// example:
	//
	// application/pdf
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The Base64-encoded string of the file content.
	//
	// This parameter is required.
	//
	// example:
	//
	// YWJj
	FileContentBase64 *string `json:"fileContentBase64,omitempty" xml:"fileContentBase64,omitempty"`
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// report.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
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
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UploadChatFileByBase64Request) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileByBase64Request) GoString() string {
	return s.String()
}

func (s *UploadChatFileByBase64Request) GetAppToken() *string {
	return s.AppToken
}

func (s *UploadChatFileByBase64Request) GetContentType() *string {
	return s.ContentType
}

func (s *UploadChatFileByBase64Request) GetFileContentBase64() *string {
	return s.FileContentBase64
}

func (s *UploadChatFileByBase64Request) GetFileName() *string {
	return s.FileName
}

func (s *UploadChatFileByBase64Request) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UploadChatFileByBase64Request) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UploadChatFileByBase64Request) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadChatFileByBase64Request) SetAppToken(v string) *UploadChatFileByBase64Request {
	s.AppToken = &v
	return s
}

func (s *UploadChatFileByBase64Request) SetContentType(v string) *UploadChatFileByBase64Request {
	s.ContentType = &v
	return s
}

func (s *UploadChatFileByBase64Request) SetFileContentBase64(v string) *UploadChatFileByBase64Request {
	s.FileContentBase64 = &v
	return s
}

func (s *UploadChatFileByBase64Request) SetFileName(v string) *UploadChatFileByBase64Request {
	s.FileName = &v
	return s
}

func (s *UploadChatFileByBase64Request) SetOpsUserCode(v string) *UploadChatFileByBase64Request {
	s.OpsUserCode = &v
	return s
}

func (s *UploadChatFileByBase64Request) SetOpsUserId(v string) *UploadChatFileByBase64Request {
	s.OpsUserId = &v
	return s
}

func (s *UploadChatFileByBase64Request) SetTenantId(v string) *UploadChatFileByBase64Request {
	s.TenantId = &v
	return s
}

func (s *UploadChatFileByBase64Request) Validate() error {
	return dara.Validate(s)
}
