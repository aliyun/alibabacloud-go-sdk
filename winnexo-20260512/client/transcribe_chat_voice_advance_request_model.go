// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iTranscribeChatVoiceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *TranscribeChatVoiceAdvanceRequest
	GetAppToken() *string
	SetContentType(v string) *TranscribeChatVoiceAdvanceRequest
	GetContentType() *string
	SetFileContentBase64(v string) *TranscribeChatVoiceAdvanceRequest
	GetFileContentBase64() *string
	SetFileName(v string) *TranscribeChatVoiceAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *TranscribeChatVoiceAdvanceRequest
	GetFileUrlObject() io.Reader
	SetOpsUserCode(v string) *TranscribeChatVoiceAdvanceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *TranscribeChatVoiceAdvanceRequest
	GetOpsUserId() *string
	SetTenantId(v string) *TranscribeChatVoiceAdvanceRequest
	GetTenantId() *string
}

type TranscribeChatVoiceAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
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

func (s TranscribeChatVoiceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TranscribeChatVoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TranscribeChatVoiceAdvanceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *TranscribeChatVoiceAdvanceRequest) GetContentType() *string {
	return s.ContentType
}

func (s *TranscribeChatVoiceAdvanceRequest) GetFileContentBase64() *string {
	return s.FileContentBase64
}

func (s *TranscribeChatVoiceAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *TranscribeChatVoiceAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *TranscribeChatVoiceAdvanceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *TranscribeChatVoiceAdvanceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *TranscribeChatVoiceAdvanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TranscribeChatVoiceAdvanceRequest) SetAppToken(v string) *TranscribeChatVoiceAdvanceRequest {
	s.AppToken = &v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) SetContentType(v string) *TranscribeChatVoiceAdvanceRequest {
	s.ContentType = &v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) SetFileContentBase64(v string) *TranscribeChatVoiceAdvanceRequest {
	s.FileContentBase64 = &v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) SetFileName(v string) *TranscribeChatVoiceAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) SetFileUrlObject(v io.Reader) *TranscribeChatVoiceAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) SetOpsUserCode(v string) *TranscribeChatVoiceAdvanceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) SetOpsUserId(v string) *TranscribeChatVoiceAdvanceRequest {
	s.OpsUserId = &v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) SetTenantId(v string) *TranscribeChatVoiceAdvanceRequest {
	s.TenantId = &v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
