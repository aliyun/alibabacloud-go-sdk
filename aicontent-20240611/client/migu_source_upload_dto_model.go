// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMiguSourceUploadDTO interface {
	dara.Model
	String() string
	GoString() string
	SetExpiresAt(v string) *MiguSourceUploadDTO
	GetExpiresAt() *string
	SetFileType(v string) *MiguSourceUploadDTO
	GetFileType() *string
	SetSourceId(v string) *MiguSourceUploadDTO
	GetSourceId() *string
	SetUploadUrl(v string) *MiguSourceUploadDTO
	GetUploadUrl() *string
}

type MiguSourceUploadDTO struct {
	// The expiration time of the upload URL in RFC 3339 format.
	//
	// example:
	//
	// 2026-08-28T12:00:00Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The type of the source file (uppercase). Valid values: VIDEO, IMAGE, AUDIO, and TEXT.
	//
	// example:
	//
	// VIDEO
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// The unique identifier of the source file, used for subsequent generation tasks and downloads.
	//
	// example:
	//
	// 3f2a1b9c8d7e4f60a1b2c3d4e5f6a7b8
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The OSS pre-signed upload URL. Use the PUT method to upload the file.
	//
	// example:
	//
	// https://bucket.oss-cn-beijing.aliyuncs.com/pipeline/source/xxx.mp4?Expires=1700000000&Signature=xxx
	UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty"`
}

func (s MiguSourceUploadDTO) String() string {
	return dara.Prettify(s)
}

func (s MiguSourceUploadDTO) GoString() string {
	return s.String()
}

func (s *MiguSourceUploadDTO) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *MiguSourceUploadDTO) GetFileType() *string {
	return s.FileType
}

func (s *MiguSourceUploadDTO) GetSourceId() *string {
	return s.SourceId
}

func (s *MiguSourceUploadDTO) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *MiguSourceUploadDTO) SetExpiresAt(v string) *MiguSourceUploadDTO {
	s.ExpiresAt = &v
	return s
}

func (s *MiguSourceUploadDTO) SetFileType(v string) *MiguSourceUploadDTO {
	s.FileType = &v
	return s
}

func (s *MiguSourceUploadDTO) SetSourceId(v string) *MiguSourceUploadDTO {
	s.SourceId = &v
	return s
}

func (s *MiguSourceUploadDTO) SetUploadUrl(v string) *MiguSourceUploadDTO {
	s.UploadUrl = &v
	return s
}

func (s *MiguSourceUploadDTO) Validate() error {
	return dara.Validate(s)
}
