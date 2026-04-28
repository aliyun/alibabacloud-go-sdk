// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevision interface {
	dara.Model
	String() string
	GoString() string
	SetContentHash(v string) *Revision
	GetContentHash() *string
	SetContentHashName(v string) *Revision
	GetContentHashName() *string
	SetCrc64Hash(v string) *Revision
	GetCrc64Hash() *string
	SetCreatedAt(v string) *Revision
	GetCreatedAt() *string
	SetCreatorId(v string) *Revision
	GetCreatorId() *string
	SetCreatorName(v string) *Revision
	GetCreatorName() *string
	SetDomainId(v string) *Revision
	GetDomainId() *string
	SetDownloadUrl(v string) *Revision
	GetDownloadUrl() *string
	SetDriveId(v string) *Revision
	GetDriveId() *string
	SetFileExtension(v string) *Revision
	GetFileExtension() *string
	SetFileId(v string) *Revision
	GetFileId() *string
	SetIsLatestVersion(v bool) *Revision
	GetIsLatestVersion() *bool
	SetKeepForever(v bool) *Revision
	GetKeepForever() *bool
	SetRevisionDescription(v string) *Revision
	GetRevisionDescription() *string
	SetRevisionId(v string) *Revision
	GetRevisionId() *string
	SetRevisionName(v string) *Revision
	GetRevisionName() *string
	SetRevisionVersion(v int64) *Revision
	GetRevisionVersion() *int64
	SetSize(v int64) *Revision
	GetSize() *int64
	SetThumbnail(v string) *Revision
	GetThumbnail() *string
	SetUpdatedAt(v string) *Revision
	GetUpdatedAt() *string
	SetUrl(v string) *Revision
	GetUrl() *string
}

type Revision struct {
	// The hash value of the content.
	//
	// example:
	//
	// EA4942AA8761213890A5C386F88E6464D2C31CA3
	ContentHash *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	// The name of the hash algorithm.
	//
	// example:
	//
	// sha1
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	// The CRC64 value of the version.
	//
	// example:
	//
	// 3574582125365864471
	Crc64Hash *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	// The time when the version was created.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The ID of the user who created the version.
	CreatorId *string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// The name of the user who created the version.
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The download URL. The ListRevision operation does not return this value. The GetRevision, UpdateRevision, and RestoreRevision operations return this value.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	DownloadUrl *string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file extension.
	//
	// example:
	//
	// mov
	FileExtension *string `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Indicates whether it is the latest version.
	//
	// example:
	//
	// false
	IsLatestVersion *bool `json:"is_latest_version,omitempty" xml:"is_latest_version,omitempty"`
	// Indicates whether the version is permanently retained.
	//
	// example:
	//
	// false
	KeepForever *bool `json:"keep_forever,omitempty" xml:"keep_forever,omitempty"`
	// The description of the version.
	//
	// example:
	//
	// aaa
	RevisionDescription *string `json:"revision_description,omitempty" xml:"revision_description,omitempty"`
	// The version ID.
	//
	// example:
	//
	// 40CB7794C929
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	// The version name.
	//
	// example:
	//
	// 1.mov
	RevisionName *string `json:"revision_name,omitempty" xml:"revision_name,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	RevisionVersion *int64 `json:"revision_version,omitempty" xml:"revision_version,omitempty"`
	// The version size.
	//
	// example:
	//
	// 1024
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	Thumbnail *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	// The time when the version was modified.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The preview URL.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s Revision) String() string {
	return dara.Prettify(s)
}

func (s Revision) GoString() string {
	return s.String()
}

func (s *Revision) GetContentHash() *string {
	return s.ContentHash
}

func (s *Revision) GetContentHashName() *string {
	return s.ContentHashName
}

func (s *Revision) GetCrc64Hash() *string {
	return s.Crc64Hash
}

func (s *Revision) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Revision) GetCreatorId() *string {
	return s.CreatorId
}

func (s *Revision) GetCreatorName() *string {
	return s.CreatorName
}

func (s *Revision) GetDomainId() *string {
	return s.DomainId
}

func (s *Revision) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *Revision) GetDriveId() *string {
	return s.DriveId
}

func (s *Revision) GetFileExtension() *string {
	return s.FileExtension
}

func (s *Revision) GetFileId() *string {
	return s.FileId
}

func (s *Revision) GetIsLatestVersion() *bool {
	return s.IsLatestVersion
}

func (s *Revision) GetKeepForever() *bool {
	return s.KeepForever
}

func (s *Revision) GetRevisionDescription() *string {
	return s.RevisionDescription
}

func (s *Revision) GetRevisionId() *string {
	return s.RevisionId
}

func (s *Revision) GetRevisionName() *string {
	return s.RevisionName
}

func (s *Revision) GetRevisionVersion() *int64 {
	return s.RevisionVersion
}

func (s *Revision) GetSize() *int64 {
	return s.Size
}

func (s *Revision) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *Revision) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Revision) GetUrl() *string {
	return s.Url
}

func (s *Revision) SetContentHash(v string) *Revision {
	s.ContentHash = &v
	return s
}

func (s *Revision) SetContentHashName(v string) *Revision {
	s.ContentHashName = &v
	return s
}

func (s *Revision) SetCrc64Hash(v string) *Revision {
	s.Crc64Hash = &v
	return s
}

func (s *Revision) SetCreatedAt(v string) *Revision {
	s.CreatedAt = &v
	return s
}

func (s *Revision) SetCreatorId(v string) *Revision {
	s.CreatorId = &v
	return s
}

func (s *Revision) SetCreatorName(v string) *Revision {
	s.CreatorName = &v
	return s
}

func (s *Revision) SetDomainId(v string) *Revision {
	s.DomainId = &v
	return s
}

func (s *Revision) SetDownloadUrl(v string) *Revision {
	s.DownloadUrl = &v
	return s
}

func (s *Revision) SetDriveId(v string) *Revision {
	s.DriveId = &v
	return s
}

func (s *Revision) SetFileExtension(v string) *Revision {
	s.FileExtension = &v
	return s
}

func (s *Revision) SetFileId(v string) *Revision {
	s.FileId = &v
	return s
}

func (s *Revision) SetIsLatestVersion(v bool) *Revision {
	s.IsLatestVersion = &v
	return s
}

func (s *Revision) SetKeepForever(v bool) *Revision {
	s.KeepForever = &v
	return s
}

func (s *Revision) SetRevisionDescription(v string) *Revision {
	s.RevisionDescription = &v
	return s
}

func (s *Revision) SetRevisionId(v string) *Revision {
	s.RevisionId = &v
	return s
}

func (s *Revision) SetRevisionName(v string) *Revision {
	s.RevisionName = &v
	return s
}

func (s *Revision) SetRevisionVersion(v int64) *Revision {
	s.RevisionVersion = &v
	return s
}

func (s *Revision) SetSize(v int64) *Revision {
	s.Size = &v
	return s
}

func (s *Revision) SetThumbnail(v string) *Revision {
	s.Thumbnail = &v
	return s
}

func (s *Revision) SetUpdatedAt(v string) *Revision {
	s.UpdatedAt = &v
	return s
}

func (s *Revision) SetUrl(v string) *Revision {
	s.Url = &v
	return s
}

func (s *Revision) Validate() error {
	return dara.Validate(s)
}
