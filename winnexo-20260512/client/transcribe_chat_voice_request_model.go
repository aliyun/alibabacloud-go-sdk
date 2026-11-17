// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranscribeChatVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *TranscribeChatVoiceRequest
	GetContentType() *string
	SetFileContentBase64(v string) *TranscribeChatVoiceRequest
	GetFileContentBase64() *string
	SetFileName(v string) *TranscribeChatVoiceRequest
	GetFileName() *string
	SetFileUrl(v string) *TranscribeChatVoiceRequest
	GetFileUrl() *string
	SetTenantId(v string) *TranscribeChatVoiceRequest
	GetTenantId() *string
}

type TranscribeChatVoiceRequest struct {
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

func (s *TranscribeChatVoiceRequest) GetTenantId() *string {
	return s.TenantId
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

func (s *TranscribeChatVoiceRequest) SetTenantId(v string) *TranscribeChatVoiceRequest {
	s.TenantId = &v
	return s
}

func (s *TranscribeChatVoiceRequest) Validate() error {
	return dara.Validate(s)
}
