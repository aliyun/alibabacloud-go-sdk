// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseMeta interface {
	dara.Model
	String() string
	GoString() string
	SetChunkAttachment(v []*KnowledgeBaseMetaChunkAttachment) *KnowledgeBaseMeta
	GetChunkAttachment() []*KnowledgeBaseMetaChunkAttachment
	SetChunkContent(v string) *KnowledgeBaseMeta
	GetChunkContent() *string
	SetChunkEnd(v int32) *KnowledgeBaseMeta
	GetChunkEnd() *int32
	SetChunkId(v string) *KnowledgeBaseMeta
	GetChunkId() *string
	SetChunkSequence(v int32) *KnowledgeBaseMeta
	GetChunkSequence() *int32
	SetChunkSize(v int32) *KnowledgeBaseMeta
	GetChunkSize() *int32
	SetChunkStart(v int32) *KnowledgeBaseMeta
	GetChunkStart() *int32
	SetChunkStatus(v string) *KnowledgeBaseMeta
	GetChunkStatus() *string
	SetDownloadUrl(v string) *KnowledgeBaseMeta
	GetDownloadUrl() *string
	SetFileContentType(v string) *KnowledgeBaseMeta
	GetFileContentType() *string
	SetFileCreateTime(v string) *KnowledgeBaseMeta
	GetFileCreateTime() *string
	SetFileMetaData(v string) *KnowledgeBaseMeta
	GetFileMetaData() *string
	SetFileMetaId(v string) *KnowledgeBaseMeta
	GetFileMetaId() *string
	SetFileName(v string) *KnowledgeBaseMeta
	GetFileName() *string
	SetFileSize(v int32) *KnowledgeBaseMeta
	GetFileSize() *int32
	SetFileStatus(v string) *KnowledgeBaseMeta
	GetFileStatus() *string
	SetFileType(v string) *KnowledgeBaseMeta
	GetFileType() *string
	SetFileUpdateTime(v string) *KnowledgeBaseMeta
	GetFileUpdateTime() *string
	SetFileUri(v string) *KnowledgeBaseMeta
	GetFileUri() *string
	SetMetaData(v *KnowledgeBaseMetaMetaData) *KnowledgeBaseMeta
	GetMetaData() *KnowledgeBaseMetaMetaData
	SetThumbnailUrl(v string) *KnowledgeBaseMeta
	GetThumbnailUrl() *string
}

type KnowledgeBaseMeta struct {
	ChunkAttachment []*KnowledgeBaseMetaChunkAttachment `json:"ChunkAttachment,omitempty" xml:"ChunkAttachment,omitempty" type:"Repeated"`
	ChunkContent    *string                             `json:"ChunkContent,omitempty" xml:"ChunkContent,omitempty"`
	ChunkEnd        *int32                              `json:"ChunkEnd,omitempty" xml:"ChunkEnd,omitempty"`
	ChunkId         *string                             `json:"ChunkId,omitempty" xml:"ChunkId,omitempty"`
	ChunkSequence   *int32                              `json:"ChunkSequence,omitempty" xml:"ChunkSequence,omitempty"`
	ChunkSize       *int32                              `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	ChunkStart      *int32                              `json:"ChunkStart,omitempty" xml:"ChunkStart,omitempty"`
	ChunkStatus     *string                             `json:"ChunkStatus,omitempty" xml:"ChunkStatus,omitempty"`
	DownloadUrl     *string                             `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FileContentType *string                             `json:"FileContentType,omitempty" xml:"FileContentType,omitempty"`
	FileCreateTime  *string                             `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	FileMetaData    *string                             `json:"FileMetaData,omitempty" xml:"FileMetaData,omitempty"`
	FileMetaId      *string                             `json:"FileMetaId,omitempty" xml:"FileMetaId,omitempty"`
	FileName        *string                             `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize        *int32                              `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileStatus      *string                             `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	FileType        *string                             `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FileUpdateTime  *string                             `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	FileUri         *string                             `json:"FileUri,omitempty" xml:"FileUri,omitempty"`
	MetaData        *KnowledgeBaseMetaMetaData          `json:"MetaData,omitempty" xml:"MetaData,omitempty" type:"Struct"`
	ThumbnailUrl    *string                             `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
}

func (s KnowledgeBaseMeta) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseMeta) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseMeta) GetChunkAttachment() []*KnowledgeBaseMetaChunkAttachment {
	return s.ChunkAttachment
}

func (s *KnowledgeBaseMeta) GetChunkContent() *string {
	return s.ChunkContent
}

func (s *KnowledgeBaseMeta) GetChunkEnd() *int32 {
	return s.ChunkEnd
}

func (s *KnowledgeBaseMeta) GetChunkId() *string {
	return s.ChunkId
}

func (s *KnowledgeBaseMeta) GetChunkSequence() *int32 {
	return s.ChunkSequence
}

func (s *KnowledgeBaseMeta) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *KnowledgeBaseMeta) GetChunkStart() *int32 {
	return s.ChunkStart
}

func (s *KnowledgeBaseMeta) GetChunkStatus() *string {
	return s.ChunkStatus
}

func (s *KnowledgeBaseMeta) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *KnowledgeBaseMeta) GetFileContentType() *string {
	return s.FileContentType
}

func (s *KnowledgeBaseMeta) GetFileCreateTime() *string {
	return s.FileCreateTime
}

func (s *KnowledgeBaseMeta) GetFileMetaData() *string {
	return s.FileMetaData
}

func (s *KnowledgeBaseMeta) GetFileMetaId() *string {
	return s.FileMetaId
}

func (s *KnowledgeBaseMeta) GetFileName() *string {
	return s.FileName
}

func (s *KnowledgeBaseMeta) GetFileSize() *int32 {
	return s.FileSize
}

func (s *KnowledgeBaseMeta) GetFileStatus() *string {
	return s.FileStatus
}

func (s *KnowledgeBaseMeta) GetFileType() *string {
	return s.FileType
}

func (s *KnowledgeBaseMeta) GetFileUpdateTime() *string {
	return s.FileUpdateTime
}

func (s *KnowledgeBaseMeta) GetFileUri() *string {
	return s.FileUri
}

func (s *KnowledgeBaseMeta) GetMetaData() *KnowledgeBaseMetaMetaData {
	return s.MetaData
}

func (s *KnowledgeBaseMeta) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *KnowledgeBaseMeta) SetChunkAttachment(v []*KnowledgeBaseMetaChunkAttachment) *KnowledgeBaseMeta {
	s.ChunkAttachment = v
	return s
}

func (s *KnowledgeBaseMeta) SetChunkContent(v string) *KnowledgeBaseMeta {
	s.ChunkContent = &v
	return s
}

func (s *KnowledgeBaseMeta) SetChunkEnd(v int32) *KnowledgeBaseMeta {
	s.ChunkEnd = &v
	return s
}

func (s *KnowledgeBaseMeta) SetChunkId(v string) *KnowledgeBaseMeta {
	s.ChunkId = &v
	return s
}

func (s *KnowledgeBaseMeta) SetChunkSequence(v int32) *KnowledgeBaseMeta {
	s.ChunkSequence = &v
	return s
}

func (s *KnowledgeBaseMeta) SetChunkSize(v int32) *KnowledgeBaseMeta {
	s.ChunkSize = &v
	return s
}

func (s *KnowledgeBaseMeta) SetChunkStart(v int32) *KnowledgeBaseMeta {
	s.ChunkStart = &v
	return s
}

func (s *KnowledgeBaseMeta) SetChunkStatus(v string) *KnowledgeBaseMeta {
	s.ChunkStatus = &v
	return s
}

func (s *KnowledgeBaseMeta) SetDownloadUrl(v string) *KnowledgeBaseMeta {
	s.DownloadUrl = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileContentType(v string) *KnowledgeBaseMeta {
	s.FileContentType = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileCreateTime(v string) *KnowledgeBaseMeta {
	s.FileCreateTime = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileMetaData(v string) *KnowledgeBaseMeta {
	s.FileMetaData = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileMetaId(v string) *KnowledgeBaseMeta {
	s.FileMetaId = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileName(v string) *KnowledgeBaseMeta {
	s.FileName = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileSize(v int32) *KnowledgeBaseMeta {
	s.FileSize = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileStatus(v string) *KnowledgeBaseMeta {
	s.FileStatus = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileType(v string) *KnowledgeBaseMeta {
	s.FileType = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileUpdateTime(v string) *KnowledgeBaseMeta {
	s.FileUpdateTime = &v
	return s
}

func (s *KnowledgeBaseMeta) SetFileUri(v string) *KnowledgeBaseMeta {
	s.FileUri = &v
	return s
}

func (s *KnowledgeBaseMeta) SetMetaData(v *KnowledgeBaseMetaMetaData) *KnowledgeBaseMeta {
	s.MetaData = v
	return s
}

func (s *KnowledgeBaseMeta) SetThumbnailUrl(v string) *KnowledgeBaseMeta {
	s.ThumbnailUrl = &v
	return s
}

func (s *KnowledgeBaseMeta) Validate() error {
	if s.ChunkAttachment != nil {
		for _, item := range s.ChunkAttachment {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KnowledgeBaseMetaChunkAttachment struct {
	DownloadUrl   *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	PlaceholderId *string `json:"PlaceholderId,omitempty" xml:"PlaceholderId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri           *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s KnowledgeBaseMetaChunkAttachment) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseMetaChunkAttachment) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseMetaChunkAttachment) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *KnowledgeBaseMetaChunkAttachment) GetPlaceholderId() *string {
	return s.PlaceholderId
}

func (s *KnowledgeBaseMetaChunkAttachment) GetType() *string {
	return s.Type
}

func (s *KnowledgeBaseMetaChunkAttachment) GetUri() *string {
	return s.Uri
}

func (s *KnowledgeBaseMetaChunkAttachment) SetDownloadUrl(v string) *KnowledgeBaseMetaChunkAttachment {
	s.DownloadUrl = &v
	return s
}

func (s *KnowledgeBaseMetaChunkAttachment) SetPlaceholderId(v string) *KnowledgeBaseMetaChunkAttachment {
	s.PlaceholderId = &v
	return s
}

func (s *KnowledgeBaseMetaChunkAttachment) SetType(v string) *KnowledgeBaseMetaChunkAttachment {
	s.Type = &v
	return s
}

func (s *KnowledgeBaseMetaChunkAttachment) SetUri(v string) *KnowledgeBaseMetaChunkAttachment {
	s.Uri = &v
	return s
}

func (s *KnowledgeBaseMetaChunkAttachment) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseMetaMetaData struct {
	FileMetaId *string `json:"FileMetaId,omitempty" xml:"FileMetaId,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUri    *string `json:"FileUri,omitempty" xml:"FileUri,omitempty"`
}

func (s KnowledgeBaseMetaMetaData) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseMetaMetaData) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseMetaMetaData) GetFileMetaId() *string {
	return s.FileMetaId
}

func (s *KnowledgeBaseMetaMetaData) GetFileName() *string {
	return s.FileName
}

func (s *KnowledgeBaseMetaMetaData) GetFileUri() *string {
	return s.FileUri
}

func (s *KnowledgeBaseMetaMetaData) SetFileMetaId(v string) *KnowledgeBaseMetaMetaData {
	s.FileMetaId = &v
	return s
}

func (s *KnowledgeBaseMetaMetaData) SetFileName(v string) *KnowledgeBaseMetaMetaData {
	s.FileName = &v
	return s
}

func (s *KnowledgeBaseMetaMetaData) SetFileUri(v string) *KnowledgeBaseMetaMetaData {
	s.FileUri = &v
	return s
}

func (s *KnowledgeBaseMetaMetaData) Validate() error {
	return dara.Validate(s)
}
