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
	SetContentType(v string) *TranscribeChatVoiceAdvanceRequest
	GetContentType() *string
	SetFileName(v string) *TranscribeChatVoiceAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *TranscribeChatVoiceAdvanceRequest
	GetFileUrlObject() io.Reader
	SetTenantId(v string) *TranscribeChatVoiceAdvanceRequest
	GetTenantId() *string
}

type TranscribeChatVoiceAdvanceRequest struct {
	// The content type of the file. Valid values:
	//
	// - **image**: Image.
	//
	// - **document**: General document.
	//
	// - **alidoc**: Alibaba Cloud document.
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
	// The full path name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// meeting.mp3
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The publicly accessible URL of the attachment.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://winnexo-file-transfer.oss-cn-hangzhou.aliyuncs.com/openapi/2026-09-08/9f8c2a1b
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// The ID of the tenant for which the operation takes effect.
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

func (s *TranscribeChatVoiceAdvanceRequest) GetContentType() *string {
	return s.ContentType
}

func (s *TranscribeChatVoiceAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *TranscribeChatVoiceAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *TranscribeChatVoiceAdvanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TranscribeChatVoiceAdvanceRequest) SetContentType(v string) *TranscribeChatVoiceAdvanceRequest {
	s.ContentType = &v
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

func (s *TranscribeChatVoiceAdvanceRequest) SetTenantId(v string) *TranscribeChatVoiceAdvanceRequest {
	s.TenantId = &v
	return s
}

func (s *TranscribeChatVoiceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
