// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewFile interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ViewFile
	GetCategory() *string
	SetContentHash(v string) *ViewFile
	GetContentHash() *string
	SetContentHashName(v string) *ViewFile
	GetContentHashName() *string
	SetContentType(v string) *ViewFile
	GetContentType() *string
	SetCrc64Hash(v string) *ViewFile
	GetCrc64Hash() *string
	SetCreatedAt(v string) *ViewFile
	GetCreatedAt() *string
	SetDescription(v string) *ViewFile
	GetDescription() *string
	SetDomainId(v string) *ViewFile
	GetDomainId() *string
	SetDownloadUrl(v string) *ViewFile
	GetDownloadUrl() *string
	SetDriveId(v string) *ViewFile
	GetDriveId() *string
	SetFields(v map[string]interface{}) *ViewFile
	GetFields() map[string]interface{}
	SetFileExtension(v string) *ViewFile
	GetFileExtension() *string
	SetFileId(v string) *ViewFile
	GetFileId() *string
	SetFileRevisionId(v string) *ViewFile
	GetFileRevisionId() *string
	SetHidden(v bool) *ViewFile
	GetHidden() *bool
	SetInvestigationInfo(v *ViewFileInvestigationInfo) *ViewFile
	GetInvestigationInfo() *ViewFileInvestigationInfo
	SetJoinedAt(v int64) *ViewFile
	GetJoinedAt() *int64
	SetLabels(v []*string) *ViewFile
	GetLabels() []*string
	SetLocalCreatedAt(v string) *ViewFile
	GetLocalCreatedAt() *string
	SetLocalModifiedAt(v string) *ViewFile
	GetLocalModifiedAt() *string
	SetName(v string) *ViewFile
	GetName() *string
	SetParentFileId(v string) *ViewFile
	GetParentFileId() *string
	SetRevisionId(v string) *ViewFile
	GetRevisionId() *string
	SetSize(v int64) *ViewFile
	GetSize() *int64
	SetStarred(v bool) *ViewFile
	GetStarred() *bool
	SetStatus(v string) *ViewFile
	GetStatus() *string
	SetThumbnail(v string) *ViewFile
	GetThumbnail() *string
	SetThumbnailUrls(v map[string]*string) *ViewFile
	GetThumbnailUrls() map[string]*string
	SetTrashedAt(v string) *ViewFile
	GetTrashedAt() *string
	SetType(v string) *ViewFile
	GetType() *string
	SetUpdatedAt(v string) *ViewFile
	GetUpdatedAt() *string
	SetUploadId(v string) *ViewFile
	GetUploadId() *string
	SetViewId(v string) *ViewFile
	GetViewId() *string
}

type ViewFile struct {
	Category          *string                    `json:"category,omitempty" xml:"category,omitempty"`
	ContentHash       *string                    `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName   *string                    `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentType       *string                    `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Crc64Hash         *string                    `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	CreatedAt         *string                    `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description       *string                    `json:"description,omitempty" xml:"description,omitempty"`
	DomainId          *string                    `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DownloadUrl       *string                    `json:"download_url,omitempty" xml:"download_url,omitempty"`
	DriveId           *string                    `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Fields            map[string]interface{}     `json:"fields,omitempty" xml:"fields,omitempty"`
	FileExtension     *string                    `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	FileId            *string                    `json:"file_id,omitempty" xml:"file_id,omitempty"`
	FileRevisionId    *string                    `json:"file_revision_id,omitempty" xml:"file_revision_id,omitempty"`
	Hidden            *bool                      `json:"hidden,omitempty" xml:"hidden,omitempty"`
	InvestigationInfo *ViewFileInvestigationInfo `json:"investigation_info,omitempty" xml:"investigation_info,omitempty" type:"Struct"`
	JoinedAt          *int64                     `json:"joined_at,omitempty" xml:"joined_at,omitempty"`
	Labels            []*string                  `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	LocalCreatedAt    *string                    `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	LocalModifiedAt   *string                    `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	Name              *string                    `json:"name,omitempty" xml:"name,omitempty"`
	ParentFileId      *string                    `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	RevisionId        *string                    `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	Size              *int64                     `json:"size,omitempty" xml:"size,omitempty"`
	Starred           *bool                      `json:"starred,omitempty" xml:"starred,omitempty"`
	Status            *string                    `json:"status,omitempty" xml:"status,omitempty"`
	Thumbnail         *string                    `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	ThumbnailUrls     map[string]*string         `json:"thumbnail_urls,omitempty" xml:"thumbnail_urls,omitempty"`
	TrashedAt         *string                    `json:"trashed_at,omitempty" xml:"trashed_at,omitempty"`
	Type              *string                    `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt         *string                    `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UploadId          *string                    `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	ViewId            *string                    `json:"view_id,omitempty" xml:"view_id,omitempty"`
}

func (s ViewFile) String() string {
	return dara.Prettify(s)
}

func (s ViewFile) GoString() string {
	return s.String()
}

func (s *ViewFile) GetCategory() *string {
	return s.Category
}

func (s *ViewFile) GetContentHash() *string {
	return s.ContentHash
}

func (s *ViewFile) GetContentHashName() *string {
	return s.ContentHashName
}

func (s *ViewFile) GetContentType() *string {
	return s.ContentType
}

func (s *ViewFile) GetCrc64Hash() *string {
	return s.Crc64Hash
}

func (s *ViewFile) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ViewFile) GetDescription() *string {
	return s.Description
}

func (s *ViewFile) GetDomainId() *string {
	return s.DomainId
}

func (s *ViewFile) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ViewFile) GetDriveId() *string {
	return s.DriveId
}

func (s *ViewFile) GetFields() map[string]interface{} {
	return s.Fields
}

func (s *ViewFile) GetFileExtension() *string {
	return s.FileExtension
}

func (s *ViewFile) GetFileId() *string {
	return s.FileId
}

func (s *ViewFile) GetFileRevisionId() *string {
	return s.FileRevisionId
}

func (s *ViewFile) GetHidden() *bool {
	return s.Hidden
}

func (s *ViewFile) GetInvestigationInfo() *ViewFileInvestigationInfo {
	return s.InvestigationInfo
}

func (s *ViewFile) GetJoinedAt() *int64 {
	return s.JoinedAt
}

func (s *ViewFile) GetLabels() []*string {
	return s.Labels
}

func (s *ViewFile) GetLocalCreatedAt() *string {
	return s.LocalCreatedAt
}

func (s *ViewFile) GetLocalModifiedAt() *string {
	return s.LocalModifiedAt
}

func (s *ViewFile) GetName() *string {
	return s.Name
}

func (s *ViewFile) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *ViewFile) GetRevisionId() *string {
	return s.RevisionId
}

func (s *ViewFile) GetSize() *int64 {
	return s.Size
}

func (s *ViewFile) GetStarred() *bool {
	return s.Starred
}

func (s *ViewFile) GetStatus() *string {
	return s.Status
}

func (s *ViewFile) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *ViewFile) GetThumbnailUrls() map[string]*string {
	return s.ThumbnailUrls
}

func (s *ViewFile) GetTrashedAt() *string {
	return s.TrashedAt
}

func (s *ViewFile) GetType() *string {
	return s.Type
}

func (s *ViewFile) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ViewFile) GetUploadId() *string {
	return s.UploadId
}

func (s *ViewFile) GetViewId() *string {
	return s.ViewId
}

func (s *ViewFile) SetCategory(v string) *ViewFile {
	s.Category = &v
	return s
}

func (s *ViewFile) SetContentHash(v string) *ViewFile {
	s.ContentHash = &v
	return s
}

func (s *ViewFile) SetContentHashName(v string) *ViewFile {
	s.ContentHashName = &v
	return s
}

func (s *ViewFile) SetContentType(v string) *ViewFile {
	s.ContentType = &v
	return s
}

func (s *ViewFile) SetCrc64Hash(v string) *ViewFile {
	s.Crc64Hash = &v
	return s
}

func (s *ViewFile) SetCreatedAt(v string) *ViewFile {
	s.CreatedAt = &v
	return s
}

func (s *ViewFile) SetDescription(v string) *ViewFile {
	s.Description = &v
	return s
}

func (s *ViewFile) SetDomainId(v string) *ViewFile {
	s.DomainId = &v
	return s
}

func (s *ViewFile) SetDownloadUrl(v string) *ViewFile {
	s.DownloadUrl = &v
	return s
}

func (s *ViewFile) SetDriveId(v string) *ViewFile {
	s.DriveId = &v
	return s
}

func (s *ViewFile) SetFields(v map[string]interface{}) *ViewFile {
	s.Fields = v
	return s
}

func (s *ViewFile) SetFileExtension(v string) *ViewFile {
	s.FileExtension = &v
	return s
}

func (s *ViewFile) SetFileId(v string) *ViewFile {
	s.FileId = &v
	return s
}

func (s *ViewFile) SetFileRevisionId(v string) *ViewFile {
	s.FileRevisionId = &v
	return s
}

func (s *ViewFile) SetHidden(v bool) *ViewFile {
	s.Hidden = &v
	return s
}

func (s *ViewFile) SetInvestigationInfo(v *ViewFileInvestigationInfo) *ViewFile {
	s.InvestigationInfo = v
	return s
}

func (s *ViewFile) SetJoinedAt(v int64) *ViewFile {
	s.JoinedAt = &v
	return s
}

func (s *ViewFile) SetLabels(v []*string) *ViewFile {
	s.Labels = v
	return s
}

func (s *ViewFile) SetLocalCreatedAt(v string) *ViewFile {
	s.LocalCreatedAt = &v
	return s
}

func (s *ViewFile) SetLocalModifiedAt(v string) *ViewFile {
	s.LocalModifiedAt = &v
	return s
}

func (s *ViewFile) SetName(v string) *ViewFile {
	s.Name = &v
	return s
}

func (s *ViewFile) SetParentFileId(v string) *ViewFile {
	s.ParentFileId = &v
	return s
}

func (s *ViewFile) SetRevisionId(v string) *ViewFile {
	s.RevisionId = &v
	return s
}

func (s *ViewFile) SetSize(v int64) *ViewFile {
	s.Size = &v
	return s
}

func (s *ViewFile) SetStarred(v bool) *ViewFile {
	s.Starred = &v
	return s
}

func (s *ViewFile) SetStatus(v string) *ViewFile {
	s.Status = &v
	return s
}

func (s *ViewFile) SetThumbnail(v string) *ViewFile {
	s.Thumbnail = &v
	return s
}

func (s *ViewFile) SetThumbnailUrls(v map[string]*string) *ViewFile {
	s.ThumbnailUrls = v
	return s
}

func (s *ViewFile) SetTrashedAt(v string) *ViewFile {
	s.TrashedAt = &v
	return s
}

func (s *ViewFile) SetType(v string) *ViewFile {
	s.Type = &v
	return s
}

func (s *ViewFile) SetUpdatedAt(v string) *ViewFile {
	s.UpdatedAt = &v
	return s
}

func (s *ViewFile) SetUploadId(v string) *ViewFile {
	s.UploadId = &v
	return s
}

func (s *ViewFile) SetViewId(v string) *ViewFile {
	s.ViewId = &v
	return s
}

func (s *ViewFile) Validate() error {
	if s.InvestigationInfo != nil {
		if err := s.InvestigationInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ViewFileInvestigationInfo struct {
	Status     *int64  `json:"status,omitempty" xml:"status,omitempty"`
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}

func (s ViewFileInvestigationInfo) String() string {
	return dara.Prettify(s)
}

func (s ViewFileInvestigationInfo) GoString() string {
	return s.String()
}

func (s *ViewFileInvestigationInfo) GetStatus() *int64 {
	return s.Status
}

func (s *ViewFileInvestigationInfo) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ViewFileInvestigationInfo) SetStatus(v int64) *ViewFileInvestigationInfo {
	s.Status = &v
	return s
}

func (s *ViewFileInvestigationInfo) SetSuggestion(v string) *ViewFileInvestigationInfo {
	s.Suggestion = &v
	return s
}

func (s *ViewFileInvestigationInfo) Validate() error {
	return dara.Validate(s)
}
