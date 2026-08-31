// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMiguSourceDownloadDTO interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *MiguSourceDownloadDTO
	GetDownloadUrl() *string
	SetExpiresAt(v string) *MiguSourceDownloadDTO
	GetExpiresAt() *string
	SetMethod(v string) *MiguSourceDownloadDTO
	GetMethod() *string
	SetSourceId(v string) *MiguSourceDownloadDTO
	GetSourceId() *string
}

type MiguSourceDownloadDTO struct {
	// The OSS pre-signed download URL.
	//
	// example:
	//
	// https://bucket.oss-cn-beijing.aliyuncs.com/pipeline/source/xxx.mp4?Expires=1700000000&Signature=xxx
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// The expiration time of the download URL, in RFC 3339 format.
	//
	// example:
	//
	// 2026-08-28T12:00:00Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The download request method. The value is fixed to GET.
	//
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The unique identifier of the source file.
	//
	// example:
	//
	// 3f2a1b9c8d7e4f60a1b2c3d4e5f6a7b8
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s MiguSourceDownloadDTO) String() string {
	return dara.Prettify(s)
}

func (s MiguSourceDownloadDTO) GoString() string {
	return s.String()
}

func (s *MiguSourceDownloadDTO) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *MiguSourceDownloadDTO) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *MiguSourceDownloadDTO) GetMethod() *string {
	return s.Method
}

func (s *MiguSourceDownloadDTO) GetSourceId() *string {
	return s.SourceId
}

func (s *MiguSourceDownloadDTO) SetDownloadUrl(v string) *MiguSourceDownloadDTO {
	s.DownloadUrl = &v
	return s
}

func (s *MiguSourceDownloadDTO) SetExpiresAt(v string) *MiguSourceDownloadDTO {
	s.ExpiresAt = &v
	return s
}

func (s *MiguSourceDownloadDTO) SetMethod(v string) *MiguSourceDownloadDTO {
	s.Method = &v
	return s
}

func (s *MiguSourceDownloadDTO) SetSourceId(v string) *MiguSourceDownloadDTO {
	s.SourceId = &v
	return s
}

func (s *MiguSourceDownloadDTO) Validate() error {
	return dara.Validate(s)
}
