// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeFile interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *KnowledgeFile
	GetCreatorId() *string
	SetDriveId(v string) *KnowledgeFile
	GetDriveId() *string
	SetDriveName(v string) *KnowledgeFile
	GetDriveName() *string
	SetFileCategory(v string) *KnowledgeFile
	GetFileCategory() *string
	SetFileCreatedAt(v int64) *KnowledgeFile
	GetFileCreatedAt() *int64
	SetFileCreatorId(v string) *KnowledgeFile
	GetFileCreatorId() *string
	SetFileId(v string) *KnowledgeFile
	GetFileId() *string
	SetFileImageTime(v int64) *KnowledgeFile
	GetFileImageTime() *int64
	SetFileLastModifierId(v string) *KnowledgeFile
	GetFileLastModifierId() *string
	SetFileLastModifierType(v string) *KnowledgeFile
	GetFileLastModifierType() *string
	SetFileName(v string) *KnowledgeFile
	GetFileName() *string
	SetFileNamePath(v string) *KnowledgeFile
	GetFileNamePath() *string
	SetFileSize(v int64) *KnowledgeFile
	GetFileSize() *int64
	SetFileUpdatedAt(v int64) *KnowledgeFile
	GetFileUpdatedAt() *int64
	SetJoinedAt(v int64) *KnowledgeFile
	GetJoinedAt() *int64
	SetKnowledgeBaseId(v string) *KnowledgeFile
	GetKnowledgeBaseId() *string
	SetKnowledgeCategoryId(v string) *KnowledgeFile
	GetKnowledgeCategoryId() *string
	SetRevisionId(v string) *KnowledgeFile
	GetRevisionId() *string
}

type KnowledgeFile struct {
	CreatorId            *string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	DriveId              *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	DriveName            *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	FileCategory         *string `json:"file_category,omitempty" xml:"file_category,omitempty"`
	FileCreatedAt        *int64  `json:"file_created_at,omitempty" xml:"file_created_at,omitempty"`
	FileCreatorId        *string `json:"file_creator_id,omitempty" xml:"file_creator_id,omitempty"`
	FileId               *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	FileImageTime        *int64  `json:"file_image_time,omitempty" xml:"file_image_time,omitempty"`
	FileLastModifierId   *string `json:"file_last_modifier_id,omitempty" xml:"file_last_modifier_id,omitempty"`
	FileLastModifierType *string `json:"file_last_modifier_type,omitempty" xml:"file_last_modifier_type,omitempty"`
	FileName             *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	FileNamePath         *string `json:"file_name_path,omitempty" xml:"file_name_path,omitempty"`
	FileSize             *int64  `json:"file_size,omitempty" xml:"file_size,omitempty"`
	FileUpdatedAt        *int64  `json:"file_updated_at,omitempty" xml:"file_updated_at,omitempty"`
	JoinedAt             *int64  `json:"joined_at,omitempty" xml:"joined_at,omitempty"`
	KnowledgeBaseId      *string `json:"knowledge_base_id,omitempty" xml:"knowledge_base_id,omitempty"`
	KnowledgeCategoryId  *string `json:"knowledge_category_id,omitempty" xml:"knowledge_category_id,omitempty"`
	RevisionId           *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s KnowledgeFile) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeFile) GoString() string {
	return s.String()
}

func (s *KnowledgeFile) GetCreatorId() *string {
	return s.CreatorId
}

func (s *KnowledgeFile) GetDriveId() *string {
	return s.DriveId
}

func (s *KnowledgeFile) GetDriveName() *string {
	return s.DriveName
}

func (s *KnowledgeFile) GetFileCategory() *string {
	return s.FileCategory
}

func (s *KnowledgeFile) GetFileCreatedAt() *int64 {
	return s.FileCreatedAt
}

func (s *KnowledgeFile) GetFileCreatorId() *string {
	return s.FileCreatorId
}

func (s *KnowledgeFile) GetFileId() *string {
	return s.FileId
}

func (s *KnowledgeFile) GetFileImageTime() *int64 {
	return s.FileImageTime
}

func (s *KnowledgeFile) GetFileLastModifierId() *string {
	return s.FileLastModifierId
}

func (s *KnowledgeFile) GetFileLastModifierType() *string {
	return s.FileLastModifierType
}

func (s *KnowledgeFile) GetFileName() *string {
	return s.FileName
}

func (s *KnowledgeFile) GetFileNamePath() *string {
	return s.FileNamePath
}

func (s *KnowledgeFile) GetFileSize() *int64 {
	return s.FileSize
}

func (s *KnowledgeFile) GetFileUpdatedAt() *int64 {
	return s.FileUpdatedAt
}

func (s *KnowledgeFile) GetJoinedAt() *int64 {
	return s.JoinedAt
}

func (s *KnowledgeFile) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *KnowledgeFile) GetKnowledgeCategoryId() *string {
	return s.KnowledgeCategoryId
}

func (s *KnowledgeFile) GetRevisionId() *string {
	return s.RevisionId
}

func (s *KnowledgeFile) SetCreatorId(v string) *KnowledgeFile {
	s.CreatorId = &v
	return s
}

func (s *KnowledgeFile) SetDriveId(v string) *KnowledgeFile {
	s.DriveId = &v
	return s
}

func (s *KnowledgeFile) SetDriveName(v string) *KnowledgeFile {
	s.DriveName = &v
	return s
}

func (s *KnowledgeFile) SetFileCategory(v string) *KnowledgeFile {
	s.FileCategory = &v
	return s
}

func (s *KnowledgeFile) SetFileCreatedAt(v int64) *KnowledgeFile {
	s.FileCreatedAt = &v
	return s
}

func (s *KnowledgeFile) SetFileCreatorId(v string) *KnowledgeFile {
	s.FileCreatorId = &v
	return s
}

func (s *KnowledgeFile) SetFileId(v string) *KnowledgeFile {
	s.FileId = &v
	return s
}

func (s *KnowledgeFile) SetFileImageTime(v int64) *KnowledgeFile {
	s.FileImageTime = &v
	return s
}

func (s *KnowledgeFile) SetFileLastModifierId(v string) *KnowledgeFile {
	s.FileLastModifierId = &v
	return s
}

func (s *KnowledgeFile) SetFileLastModifierType(v string) *KnowledgeFile {
	s.FileLastModifierType = &v
	return s
}

func (s *KnowledgeFile) SetFileName(v string) *KnowledgeFile {
	s.FileName = &v
	return s
}

func (s *KnowledgeFile) SetFileNamePath(v string) *KnowledgeFile {
	s.FileNamePath = &v
	return s
}

func (s *KnowledgeFile) SetFileSize(v int64) *KnowledgeFile {
	s.FileSize = &v
	return s
}

func (s *KnowledgeFile) SetFileUpdatedAt(v int64) *KnowledgeFile {
	s.FileUpdatedAt = &v
	return s
}

func (s *KnowledgeFile) SetJoinedAt(v int64) *KnowledgeFile {
	s.JoinedAt = &v
	return s
}

func (s *KnowledgeFile) SetKnowledgeBaseId(v string) *KnowledgeFile {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeFile) SetKnowledgeCategoryId(v string) *KnowledgeFile {
	s.KnowledgeCategoryId = &v
	return s
}

func (s *KnowledgeFile) SetRevisionId(v string) *KnowledgeFile {
	s.RevisionId = &v
	return s
}

func (s *KnowledgeFile) Validate() error {
	return dara.Validate(s)
}
