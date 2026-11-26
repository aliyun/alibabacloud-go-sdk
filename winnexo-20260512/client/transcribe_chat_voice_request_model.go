// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranscribeChatVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *TranscribeChatVoiceRequest
	GetAppToken() *string
	SetContentType(v string) *TranscribeChatVoiceRequest
	GetContentType() *string
	SetFileContentBase64(v string) *TranscribeChatVoiceRequest
	GetFileContentBase64() *string
	SetFileName(v string) *TranscribeChatVoiceRequest
	GetFileName() *string
	SetFileUrl(v string) *TranscribeChatVoiceRequest
	GetFileUrl() *string
	SetOpsUserCode(v string) *TranscribeChatVoiceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *TranscribeChatVoiceRequest
	GetOpsUserId() *string
	SetTenantId(v string) *TranscribeChatVoiceRequest
	GetTenantId() *string
}

type TranscribeChatVoiceRequest struct {
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
	// - **alidoc**: AliDoc.
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
	// audio/mpeg
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The standard Base64-encoded audio content, excluding the data URI prefix. This parameter and FileUrl are mutually exclusive.
	//
	// example:
	//
	// UklGRg==
	FileContentBase64 *string `json:"fileContentBase64,omitempty" xml:"fileContentBase64,omitempty"`
	// The full path and name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// meeting.mp3
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The publicly accessible URL of the attachment.
	//
	// example:
	//
	// http://winnexo-file-transfer.oss-cn-hangzhou.aliyuncs.com/openapi/2026-09-08/9f8c2a1b
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
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
	// The ID of the tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s TranscribeChatVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s TranscribeChatVoiceRequest) GoString() string {
	return s.String()
}

func (s *TranscribeChatVoiceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *TranscribeChatVoiceRequest) GetContentType() *string {
	return s.ContentType
}

func (s *TranscribeChatVoiceRequest) GetFileContentBase64() *string {
	return s.FileContentBase64
}

func (s *TranscribeChatVoiceRequest) GetFileName() *string {
	return s.FileName
}

func (s *TranscribeChatVoiceRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *TranscribeChatVoiceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *TranscribeChatVoiceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *TranscribeChatVoiceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TranscribeChatVoiceRequest) SetAppToken(v string) *TranscribeChatVoiceRequest {
	s.AppToken = &v
	return s
}

func (s *TranscribeChatVoiceRequest) SetContentType(v string) *TranscribeChatVoiceRequest {
	s.ContentType = &v
	return s
}

func (s *TranscribeChatVoiceRequest) SetFileContentBase64(v string) *TranscribeChatVoiceRequest {
	s.FileContentBase64 = &v
	return s
}

func (s *TranscribeChatVoiceRequest) SetFileName(v string) *TranscribeChatVoiceRequest {
	s.FileName = &v
	return s
}

func (s *TranscribeChatVoiceRequest) SetFileUrl(v string) *TranscribeChatVoiceRequest {
	s.FileUrl = &v
	return s
}

func (s *TranscribeChatVoiceRequest) SetOpsUserCode(v string) *TranscribeChatVoiceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *TranscribeChatVoiceRequest) SetOpsUserId(v string) *TranscribeChatVoiceRequest {
	s.OpsUserId = &v
	return s
}

func (s *TranscribeChatVoiceRequest) SetTenantId(v string) *TranscribeChatVoiceRequest {
	s.TenantId = &v
	return s
}

func (s *TranscribeChatVoiceRequest) Validate() error {
	return dara.Validate(s)
}
