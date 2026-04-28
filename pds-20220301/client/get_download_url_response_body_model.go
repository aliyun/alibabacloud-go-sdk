// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCdnUrl(v string) *GetDownloadUrlResponseBody
	GetCdnUrl() *string
	SetContentHash(v string) *GetDownloadUrlResponseBody
	GetContentHash() *string
	SetContentHashName(v string) *GetDownloadUrlResponseBody
	GetContentHashName() *string
	SetCrc64Hash(v string) *GetDownloadUrlResponseBody
	GetCrc64Hash() *string
	SetExpiration(v string) *GetDownloadUrlResponseBody
	GetExpiration() *string
	SetInternalUrl(v string) *GetDownloadUrlResponseBody
	GetInternalUrl() *string
	SetSize(v int64) *GetDownloadUrlResponseBody
	GetSize() *int64
	SetUrl(v string) *GetDownloadUrlResponseBody
	GetUrl() *string
}

type GetDownloadUrlResponseBody struct {
	// The download URL of a file that is downloaded by using Alibaba Cloud CDN.
	//
	// example:
	//
	// https://data-cdn.aliyunpds.com/hz22%2F5d79219b0aa9a7c995a94a96993ba3205cd91c5a%2F5d79219bf3261a5d38744da0834ed489b677a27a?Expires=xxxOSSAccessKeyId=xxx&Signature=xxx&response-content-disposition=attachment%3Bfilename%3DtBiZAoJPC2c8b13450eda4292b7f5f8010618e078.txt
	CdnUrl *string `json:"cdn_url,omitempty" xml:"cdn_url,omitempty"`
	// The hash value of the file content.
	//
	// example:
	//
	// EA4942AA8761213890A5C386F88E6464D2C31CA1
	ContentHash *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	// The name of the algorithm that is used to calculate the hash value of the file content.
	//
	// example:
	//
	// sha1
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	// The hash value calculated by using 64-bit cyclic redundancy check (CRC-64).
	//
	// example:
	//
	// 5498595269368962671
	Crc64Hash *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	// The time when the download URL expires.
	//
	// example:
	//
	// 2022-01-02T15:04:05.999Z07:00
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// The download URL of a file that is downloaded over a virtual private cloud (VPC).
	//
	// example:
	//
	// https://data-vpc.aliyunpds.com/hz22%2F5d79219b0aa9a7c995a94a96993ba3205cd91c5a%2F5d79219bf3261a5d38744da0834ed489b677a27a?Expires=xxxOSSAccessKeyId=xxx&Signature=xxx&response-content-disposition=attachment%3Bfilename%3DtBiZAoJPC2c8b13450eda4292b7f5f8010618e078.txt
	InternalUrl *string `json:"internal_url,omitempty" xml:"internal_url,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The download URL of a file that is downloaded over the Internet.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d79219b0aa9a7c995a94a96993ba3205cd91c5a%2F5d79219bf3261a5d38744da0834ed489b677a27a?Expires=xxxOSSAccessKeyId=xxx&Signature=xxx&response-content-disposition=attachment%3Bfilename%3DtBiZAoJPC2c8b13450eda4292b7f5f8010618e078.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlResponseBody) GetCdnUrl() *string {
	return s.CdnUrl
}

func (s *GetDownloadUrlResponseBody) GetContentHash() *string {
	return s.ContentHash
}

func (s *GetDownloadUrlResponseBody) GetContentHashName() *string {
	return s.ContentHashName
}

func (s *GetDownloadUrlResponseBody) GetCrc64Hash() *string {
	return s.Crc64Hash
}

func (s *GetDownloadUrlResponseBody) GetExpiration() *string {
	return s.Expiration
}

func (s *GetDownloadUrlResponseBody) GetInternalUrl() *string {
	return s.InternalUrl
}

func (s *GetDownloadUrlResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *GetDownloadUrlResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetDownloadUrlResponseBody) SetCdnUrl(v string) *GetDownloadUrlResponseBody {
	s.CdnUrl = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetContentHash(v string) *GetDownloadUrlResponseBody {
	s.ContentHash = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetContentHashName(v string) *GetDownloadUrlResponseBody {
	s.ContentHashName = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetCrc64Hash(v string) *GetDownloadUrlResponseBody {
	s.Crc64Hash = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetExpiration(v string) *GetDownloadUrlResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetInternalUrl(v string) *GetDownloadUrlResponseBody {
	s.InternalUrl = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetSize(v int64) *GetDownloadUrlResponseBody {
	s.Size = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetUrl(v string) *GetDownloadUrlResponseBody {
	s.Url = &v
	return s
}

func (s *GetDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
