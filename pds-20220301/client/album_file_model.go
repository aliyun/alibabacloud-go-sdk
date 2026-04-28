// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlbumFile interface {
	dara.Model
	String() string
	GoString() string
	SetAlbumId(v string) *AlbumFile
	GetAlbumId() *string
	SetCategory(v string) *AlbumFile
	GetCategory() *string
	SetContentHash(v string) *AlbumFile
	GetContentHash() *string
	SetContentHashName(v string) *AlbumFile
	GetContentHashName() *string
	SetContentType(v string) *AlbumFile
	GetContentType() *string
	SetCrc64Hash(v string) *AlbumFile
	GetCrc64Hash() *string
	SetCreatedAt(v string) *AlbumFile
	GetCreatedAt() *string
	SetDescription(v string) *AlbumFile
	GetDescription() *string
	SetDomainId(v string) *AlbumFile
	GetDomainId() *string
	SetDownloadUrl(v string) *AlbumFile
	GetDownloadUrl() *string
	SetDriveId(v string) *AlbumFile
	GetDriveId() *string
	SetExFieldsInfo(v map[string]interface{}) *AlbumFile
	GetExFieldsInfo() map[string]interface{}
	SetFileExtension(v string) *AlbumFile
	GetFileExtension() *string
	SetFileId(v string) *AlbumFile
	GetFileId() *string
	SetHidden(v bool) *AlbumFile
	GetHidden() *bool
	SetImageMediaMetadata(v *ImageMediaMetadata) *AlbumFile
	GetImageMediaMetadata() *ImageMediaMetadata
	SetInvestigationInfo(v *InvestigationInfo) *AlbumFile
	GetInvestigationInfo() *InvestigationInfo
	SetJoinedAt(v int64) *AlbumFile
	GetJoinedAt() *int64
	SetLabels(v []*string) *AlbumFile
	GetLabels() []*string
	SetLocalCreatedAt(v string) *AlbumFile
	GetLocalCreatedAt() *string
	SetLocalModifiedAt(v string) *AlbumFile
	GetLocalModifiedAt() *string
	SetMimeType(v string) *AlbumFile
	GetMimeType() *string
	SetName(v string) *AlbumFile
	GetName() *string
	SetObjectUri(v string) *AlbumFile
	GetObjectUri() *string
	SetParentFileId(v string) *AlbumFile
	GetParentFileId() *string
	SetRevisionId(v string) *AlbumFile
	GetRevisionId() *string
	SetSize(v int64) *AlbumFile
	GetSize() *int64
	SetStarred(v bool) *AlbumFile
	GetStarred() *bool
	SetStatus(v string) *AlbumFile
	GetStatus() *string
	SetThumbnail(v string) *AlbumFile
	GetThumbnail() *string
	SetThumbnailUrls(v map[string]*string) *AlbumFile
	GetThumbnailUrls() map[string]*string
	SetTranshedAt(v string) *AlbumFile
	GetTranshedAt() *string
	SetType(v string) *AlbumFile
	GetType() *string
	SetUpdatedAt(v string) *AlbumFile
	GetUpdatedAt() *string
	SetUploadId(v string) *AlbumFile
	GetUploadId() *string
	SetUserMeta(v string) *AlbumFile
	GetUserMeta() *string
}

type AlbumFile struct {
	AlbumId            *string                `json:"album_id,omitempty" xml:"album_id,omitempty"`
	Category           *string                `json:"category,omitempty" xml:"category,omitempty"`
	ContentHash        *string                `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName    *string                `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentType        *string                `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Crc64Hash          *string                `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	CreatedAt          *string                `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description        *string                `json:"description,omitempty" xml:"description,omitempty"`
	DomainId           *string                `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DownloadUrl        *string                `json:"download_url,omitempty" xml:"download_url,omitempty"`
	DriveId            *string                `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	ExFieldsInfo       map[string]interface{} `json:"ex_fields_info,omitempty" xml:"ex_fields_info,omitempty"`
	FileExtension      *string                `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	FileId             *string                `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Hidden             *bool                  `json:"hidden,omitempty" xml:"hidden,omitempty"`
	ImageMediaMetadata *ImageMediaMetadata    `json:"image_media_metadata,omitempty" xml:"image_media_metadata,omitempty"`
	InvestigationInfo  *InvestigationInfo     `json:"investigation_info,omitempty" xml:"investigation_info,omitempty"`
	JoinedAt           *int64                 `json:"joined_at,omitempty" xml:"joined_at,omitempty"`
	Labels             []*string              `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	LocalCreatedAt     *string                `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	LocalModifiedAt    *string                `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	MimeType           *string                `json:"mime_type,omitempty" xml:"mime_type,omitempty"`
	Name               *string                `json:"name,omitempty" xml:"name,omitempty"`
	ObjectUri          *string                `json:"object_uri,omitempty" xml:"object_uri,omitempty"`
	ParentFileId       *string                `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	RevisionId         *string                `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	Size               *int64                 `json:"size,omitempty" xml:"size,omitempty"`
	Starred            *bool                  `json:"starred,omitempty" xml:"starred,omitempty"`
	Status             *string                `json:"status,omitempty" xml:"status,omitempty"`
	Thumbnail          *string                `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	ThumbnailUrls      map[string]*string     `json:"thumbnail_urls,omitempty" xml:"thumbnail_urls,omitempty"`
	TranshedAt         *string                `json:"transhed_at,omitempty" xml:"transhed_at,omitempty"`
	Type               *string                `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt          *string                `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UploadId           *string                `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	UserMeta           *string                `json:"user_meta,omitempty" xml:"user_meta,omitempty"`
}

func (s AlbumFile) String() string {
	return dara.Prettify(s)
}

func (s AlbumFile) GoString() string {
	return s.String()
}

func (s *AlbumFile) GetAlbumId() *string {
	return s.AlbumId
}

func (s *AlbumFile) GetCategory() *string {
	return s.Category
}

func (s *AlbumFile) GetContentHash() *string {
	return s.ContentHash
}

func (s *AlbumFile) GetContentHashName() *string {
	return s.ContentHashName
}

func (s *AlbumFile) GetContentType() *string {
	return s.ContentType
}

func (s *AlbumFile) GetCrc64Hash() *string {
	return s.Crc64Hash
}

func (s *AlbumFile) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *AlbumFile) GetDescription() *string {
	return s.Description
}

func (s *AlbumFile) GetDomainId() *string {
	return s.DomainId
}

func (s *AlbumFile) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *AlbumFile) GetDriveId() *string {
	return s.DriveId
}

func (s *AlbumFile) GetExFieldsInfo() map[string]interface{} {
	return s.ExFieldsInfo
}

func (s *AlbumFile) GetFileExtension() *string {
	return s.FileExtension
}

func (s *AlbumFile) GetFileId() *string {
	return s.FileId
}

func (s *AlbumFile) GetHidden() *bool {
	return s.Hidden
}

func (s *AlbumFile) GetImageMediaMetadata() *ImageMediaMetadata {
	return s.ImageMediaMetadata
}

func (s *AlbumFile) GetInvestigationInfo() *InvestigationInfo {
	return s.InvestigationInfo
}

func (s *AlbumFile) GetJoinedAt() *int64 {
	return s.JoinedAt
}

func (s *AlbumFile) GetLabels() []*string {
	return s.Labels
}

func (s *AlbumFile) GetLocalCreatedAt() *string {
	return s.LocalCreatedAt
}

func (s *AlbumFile) GetLocalModifiedAt() *string {
	return s.LocalModifiedAt
}

func (s *AlbumFile) GetMimeType() *string {
	return s.MimeType
}

func (s *AlbumFile) GetName() *string {
	return s.Name
}

func (s *AlbumFile) GetObjectUri() *string {
	return s.ObjectUri
}

func (s *AlbumFile) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *AlbumFile) GetRevisionId() *string {
	return s.RevisionId
}

func (s *AlbumFile) GetSize() *int64 {
	return s.Size
}

func (s *AlbumFile) GetStarred() *bool {
	return s.Starred
}

func (s *AlbumFile) GetStatus() *string {
	return s.Status
}

func (s *AlbumFile) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *AlbumFile) GetThumbnailUrls() map[string]*string {
	return s.ThumbnailUrls
}

func (s *AlbumFile) GetTranshedAt() *string {
	return s.TranshedAt
}

func (s *AlbumFile) GetType() *string {
	return s.Type
}

func (s *AlbumFile) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *AlbumFile) GetUploadId() *string {
	return s.UploadId
}

func (s *AlbumFile) GetUserMeta() *string {
	return s.UserMeta
}

func (s *AlbumFile) SetAlbumId(v string) *AlbumFile {
	s.AlbumId = &v
	return s
}

func (s *AlbumFile) SetCategory(v string) *AlbumFile {
	s.Category = &v
	return s
}

func (s *AlbumFile) SetContentHash(v string) *AlbumFile {
	s.ContentHash = &v
	return s
}

func (s *AlbumFile) SetContentHashName(v string) *AlbumFile {
	s.ContentHashName = &v
	return s
}

func (s *AlbumFile) SetContentType(v string) *AlbumFile {
	s.ContentType = &v
	return s
}

func (s *AlbumFile) SetCrc64Hash(v string) *AlbumFile {
	s.Crc64Hash = &v
	return s
}

func (s *AlbumFile) SetCreatedAt(v string) *AlbumFile {
	s.CreatedAt = &v
	return s
}

func (s *AlbumFile) SetDescription(v string) *AlbumFile {
	s.Description = &v
	return s
}

func (s *AlbumFile) SetDomainId(v string) *AlbumFile {
	s.DomainId = &v
	return s
}

func (s *AlbumFile) SetDownloadUrl(v string) *AlbumFile {
	s.DownloadUrl = &v
	return s
}

func (s *AlbumFile) SetDriveId(v string) *AlbumFile {
	s.DriveId = &v
	return s
}

func (s *AlbumFile) SetExFieldsInfo(v map[string]interface{}) *AlbumFile {
	s.ExFieldsInfo = v
	return s
}

func (s *AlbumFile) SetFileExtension(v string) *AlbumFile {
	s.FileExtension = &v
	return s
}

func (s *AlbumFile) SetFileId(v string) *AlbumFile {
	s.FileId = &v
	return s
}

func (s *AlbumFile) SetHidden(v bool) *AlbumFile {
	s.Hidden = &v
	return s
}

func (s *AlbumFile) SetImageMediaMetadata(v *ImageMediaMetadata) *AlbumFile {
	s.ImageMediaMetadata = v
	return s
}

func (s *AlbumFile) SetInvestigationInfo(v *InvestigationInfo) *AlbumFile {
	s.InvestigationInfo = v
	return s
}

func (s *AlbumFile) SetJoinedAt(v int64) *AlbumFile {
	s.JoinedAt = &v
	return s
}

func (s *AlbumFile) SetLabels(v []*string) *AlbumFile {
	s.Labels = v
	return s
}

func (s *AlbumFile) SetLocalCreatedAt(v string) *AlbumFile {
	s.LocalCreatedAt = &v
	return s
}

func (s *AlbumFile) SetLocalModifiedAt(v string) *AlbumFile {
	s.LocalModifiedAt = &v
	return s
}

func (s *AlbumFile) SetMimeType(v string) *AlbumFile {
	s.MimeType = &v
	return s
}

func (s *AlbumFile) SetName(v string) *AlbumFile {
	s.Name = &v
	return s
}

func (s *AlbumFile) SetObjectUri(v string) *AlbumFile {
	s.ObjectUri = &v
	return s
}

func (s *AlbumFile) SetParentFileId(v string) *AlbumFile {
	s.ParentFileId = &v
	return s
}

func (s *AlbumFile) SetRevisionId(v string) *AlbumFile {
	s.RevisionId = &v
	return s
}

func (s *AlbumFile) SetSize(v int64) *AlbumFile {
	s.Size = &v
	return s
}

func (s *AlbumFile) SetStarred(v bool) *AlbumFile {
	s.Starred = &v
	return s
}

func (s *AlbumFile) SetStatus(v string) *AlbumFile {
	s.Status = &v
	return s
}

func (s *AlbumFile) SetThumbnail(v string) *AlbumFile {
	s.Thumbnail = &v
	return s
}

func (s *AlbumFile) SetThumbnailUrls(v map[string]*string) *AlbumFile {
	s.ThumbnailUrls = v
	return s
}

func (s *AlbumFile) SetTranshedAt(v string) *AlbumFile {
	s.TranshedAt = &v
	return s
}

func (s *AlbumFile) SetType(v string) *AlbumFile {
	s.Type = &v
	return s
}

func (s *AlbumFile) SetUpdatedAt(v string) *AlbumFile {
	s.UpdatedAt = &v
	return s
}

func (s *AlbumFile) SetUploadId(v string) *AlbumFile {
	s.UploadId = &v
	return s
}

func (s *AlbumFile) SetUserMeta(v string) *AlbumFile {
	s.UserMeta = &v
	return s
}

func (s *AlbumFile) Validate() error {
	if s.ImageMediaMetadata != nil {
		if err := s.ImageMediaMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.InvestigationInfo != nil {
		if err := s.InvestigationInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
