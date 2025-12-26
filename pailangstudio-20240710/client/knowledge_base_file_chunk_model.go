// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseFileChunk interface {
	dara.Model
	String() string
	GoString() string
	SetChunkAttachment(v []*KnowledgeBaseFileChunkChunkAttachment) *KnowledgeBaseFileChunk
	GetChunkAttachment() []*KnowledgeBaseFileChunkChunkAttachment
	SetChunkContent(v string) *KnowledgeBaseFileChunk
	GetChunkContent() *string
	SetChunkEnd(v int32) *KnowledgeBaseFileChunk
	GetChunkEnd() *int32
	SetChunkId(v string) *KnowledgeBaseFileChunk
	GetChunkId() *string
	SetChunkSequence(v int32) *KnowledgeBaseFileChunk
	GetChunkSequence() *int32
	SetChunkSize(v int32) *KnowledgeBaseFileChunk
	GetChunkSize() *int32
	SetChunkStart(v int32) *KnowledgeBaseFileChunk
	GetChunkStart() *int32
	SetChunkStatus(v string) *KnowledgeBaseFileChunk
	GetChunkStatus() *string
	SetDownloadUrl(v string) *KnowledgeBaseFileChunk
	GetDownloadUrl() *string
	SetMetaData(v *KnowledgeBaseFileChunkMetaData) *KnowledgeBaseFileChunk
	GetMetaData() *KnowledgeBaseFileChunkMetaData
	SetScore(v float32) *KnowledgeBaseFileChunk
	GetScore() *float32
	SetThumbnailUrl(v string) *KnowledgeBaseFileChunk
	GetThumbnailUrl() *string
}

type KnowledgeBaseFileChunk struct {
	ChunkAttachment []*KnowledgeBaseFileChunkChunkAttachment `json:"ChunkAttachment,omitempty" xml:"ChunkAttachment,omitempty" type:"Repeated"`
	ChunkContent    *string                                  `json:"ChunkContent,omitempty" xml:"ChunkContent,omitempty"`
	ChunkEnd        *int32                                   `json:"ChunkEnd,omitempty" xml:"ChunkEnd,omitempty"`
	ChunkId         *string                                  `json:"ChunkId,omitempty" xml:"ChunkId,omitempty"`
	ChunkSequence   *int32                                   `json:"ChunkSequence,omitempty" xml:"ChunkSequence,omitempty"`
	ChunkSize       *int32                                   `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	ChunkStart      *int32                                   `json:"ChunkStart,omitempty" xml:"ChunkStart,omitempty"`
	ChunkStatus     *string                                  `json:"ChunkStatus,omitempty" xml:"ChunkStatus,omitempty"`
	DownloadUrl     *string                                  `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	MetaData        *KnowledgeBaseFileChunkMetaData          `json:"MetaData,omitempty" xml:"MetaData,omitempty" type:"Struct"`
	Score           *float32                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	ThumbnailUrl    *string                                  `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
}

func (s KnowledgeBaseFileChunk) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseFileChunk) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseFileChunk) GetChunkAttachment() []*KnowledgeBaseFileChunkChunkAttachment {
	return s.ChunkAttachment
}

func (s *KnowledgeBaseFileChunk) GetChunkContent() *string {
	return s.ChunkContent
}

func (s *KnowledgeBaseFileChunk) GetChunkEnd() *int32 {
	return s.ChunkEnd
}

func (s *KnowledgeBaseFileChunk) GetChunkId() *string {
	return s.ChunkId
}

func (s *KnowledgeBaseFileChunk) GetChunkSequence() *int32 {
	return s.ChunkSequence
}

func (s *KnowledgeBaseFileChunk) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *KnowledgeBaseFileChunk) GetChunkStart() *int32 {
	return s.ChunkStart
}

func (s *KnowledgeBaseFileChunk) GetChunkStatus() *string {
	return s.ChunkStatus
}

func (s *KnowledgeBaseFileChunk) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *KnowledgeBaseFileChunk) GetMetaData() *KnowledgeBaseFileChunkMetaData {
	return s.MetaData
}

func (s *KnowledgeBaseFileChunk) GetScore() *float32 {
	return s.Score
}

func (s *KnowledgeBaseFileChunk) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *KnowledgeBaseFileChunk) SetChunkAttachment(v []*KnowledgeBaseFileChunkChunkAttachment) *KnowledgeBaseFileChunk {
	s.ChunkAttachment = v
	return s
}

func (s *KnowledgeBaseFileChunk) SetChunkContent(v string) *KnowledgeBaseFileChunk {
	s.ChunkContent = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetChunkEnd(v int32) *KnowledgeBaseFileChunk {
	s.ChunkEnd = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetChunkId(v string) *KnowledgeBaseFileChunk {
	s.ChunkId = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetChunkSequence(v int32) *KnowledgeBaseFileChunk {
	s.ChunkSequence = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetChunkSize(v int32) *KnowledgeBaseFileChunk {
	s.ChunkSize = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetChunkStart(v int32) *KnowledgeBaseFileChunk {
	s.ChunkStart = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetChunkStatus(v string) *KnowledgeBaseFileChunk {
	s.ChunkStatus = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetDownloadUrl(v string) *KnowledgeBaseFileChunk {
	s.DownloadUrl = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetMetaData(v *KnowledgeBaseFileChunkMetaData) *KnowledgeBaseFileChunk {
	s.MetaData = v
	return s
}

func (s *KnowledgeBaseFileChunk) SetScore(v float32) *KnowledgeBaseFileChunk {
	s.Score = &v
	return s
}

func (s *KnowledgeBaseFileChunk) SetThumbnailUrl(v string) *KnowledgeBaseFileChunk {
	s.ThumbnailUrl = &v
	return s
}

func (s *KnowledgeBaseFileChunk) Validate() error {
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

type KnowledgeBaseFileChunkChunkAttachment struct {
	DownloadUrl   *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	PlaceholderId *string `json:"PlaceholderId,omitempty" xml:"PlaceholderId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri           *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s KnowledgeBaseFileChunkChunkAttachment) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseFileChunkChunkAttachment) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseFileChunkChunkAttachment) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *KnowledgeBaseFileChunkChunkAttachment) GetPlaceholderId() *string {
	return s.PlaceholderId
}

func (s *KnowledgeBaseFileChunkChunkAttachment) GetType() *string {
	return s.Type
}

func (s *KnowledgeBaseFileChunkChunkAttachment) GetUri() *string {
	return s.Uri
}

func (s *KnowledgeBaseFileChunkChunkAttachment) SetDownloadUrl(v string) *KnowledgeBaseFileChunkChunkAttachment {
	s.DownloadUrl = &v
	return s
}

func (s *KnowledgeBaseFileChunkChunkAttachment) SetPlaceholderId(v string) *KnowledgeBaseFileChunkChunkAttachment {
	s.PlaceholderId = &v
	return s
}

func (s *KnowledgeBaseFileChunkChunkAttachment) SetType(v string) *KnowledgeBaseFileChunkChunkAttachment {
	s.Type = &v
	return s
}

func (s *KnowledgeBaseFileChunkChunkAttachment) SetUri(v string) *KnowledgeBaseFileChunkChunkAttachment {
	s.Uri = &v
	return s
}

func (s *KnowledgeBaseFileChunkChunkAttachment) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseFileChunkMetaData struct {
	FileMetaId *string `json:"FileMetaId,omitempty" xml:"FileMetaId,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUri    *string `json:"FileUri,omitempty" xml:"FileUri,omitempty"`
}

func (s KnowledgeBaseFileChunkMetaData) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseFileChunkMetaData) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseFileChunkMetaData) GetFileMetaId() *string {
	return s.FileMetaId
}

func (s *KnowledgeBaseFileChunkMetaData) GetFileName() *string {
	return s.FileName
}

func (s *KnowledgeBaseFileChunkMetaData) GetFileUri() *string {
	return s.FileUri
}

func (s *KnowledgeBaseFileChunkMetaData) SetFileMetaId(v string) *KnowledgeBaseFileChunkMetaData {
	s.FileMetaId = &v
	return s
}

func (s *KnowledgeBaseFileChunkMetaData) SetFileName(v string) *KnowledgeBaseFileChunkMetaData {
	s.FileName = &v
	return s
}

func (s *KnowledgeBaseFileChunkMetaData) SetFileUri(v string) *KnowledgeBaseFileChunkMetaData {
	s.FileUri = &v
	return s
}

func (s *KnowledgeBaseFileChunkMetaData) Validate() error {
	return dara.Validate(s)
}
