// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleStreamInfo interface {
	dara.Model
	String() string
	GoString() string
	SetContentHash(v string) *SimpleStreamInfo
	GetContentHash() *string
	SetContentHashName(v string) *SimpleStreamInfo
	GetContentHashName() *string
	SetCrc64Hash(v string) *SimpleStreamInfo
	GetCrc64Hash() *string
	SetDownloadUrl(v string) *SimpleStreamInfo
	GetDownloadUrl() *string
	SetSize(v int64) *SimpleStreamInfo
	GetSize() *int64
	SetThumbnail(v string) *SimpleStreamInfo
	GetThumbnail() *string
	SetUrl(v string) *SimpleStreamInfo
	GetUrl() *string
}

type SimpleStreamInfo struct {
	ContentHash     *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	Crc64Hash       *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	DownloadUrl     *string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	Size            *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Thumbnail       *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SimpleStreamInfo) String() string {
	return dara.Prettify(s)
}

func (s SimpleStreamInfo) GoString() string {
	return s.String()
}

func (s *SimpleStreamInfo) GetContentHash() *string {
	return s.ContentHash
}

func (s *SimpleStreamInfo) GetContentHashName() *string {
	return s.ContentHashName
}

func (s *SimpleStreamInfo) GetCrc64Hash() *string {
	return s.Crc64Hash
}

func (s *SimpleStreamInfo) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *SimpleStreamInfo) GetSize() *int64 {
	return s.Size
}

func (s *SimpleStreamInfo) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *SimpleStreamInfo) GetUrl() *string {
	return s.Url
}

func (s *SimpleStreamInfo) SetContentHash(v string) *SimpleStreamInfo {
	s.ContentHash = &v
	return s
}

func (s *SimpleStreamInfo) SetContentHashName(v string) *SimpleStreamInfo {
	s.ContentHashName = &v
	return s
}

func (s *SimpleStreamInfo) SetCrc64Hash(v string) *SimpleStreamInfo {
	s.Crc64Hash = &v
	return s
}

func (s *SimpleStreamInfo) SetDownloadUrl(v string) *SimpleStreamInfo {
	s.DownloadUrl = &v
	return s
}

func (s *SimpleStreamInfo) SetSize(v int64) *SimpleStreamInfo {
	s.Size = &v
	return s
}

func (s *SimpleStreamInfo) SetThumbnail(v string) *SimpleStreamInfo {
	s.Thumbnail = &v
	return s
}

func (s *SimpleStreamInfo) SetUrl(v string) *SimpleStreamInfo {
	s.Url = &v
	return s
}

func (s *SimpleStreamInfo) Validate() error {
	return dara.Validate(s)
}
