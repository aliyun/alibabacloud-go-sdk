// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkStatus(v string) *ListKnowledgeBaseChunksRequest
	GetChunkStatus() *string
	SetMetaData(v *ListKnowledgeBaseChunksRequestMetaData) *ListKnowledgeBaseChunksRequest
	GetMetaData() *ListKnowledgeBaseChunksRequestMetaData
	SetPageNumber(v int32) *ListKnowledgeBaseChunksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKnowledgeBaseChunksRequest
	GetPageSize() *int32
	SetVersionName(v string) *ListKnowledgeBaseChunksRequest
	GetVersionName() *string
}

type ListKnowledgeBaseChunksRequest struct {
	// example:
	//
	// Enable
	ChunkStatus *string `json:"ChunkStatus,omitempty" xml:"ChunkStatus,omitempty"`
	// example:
	//
	// 1
	MetaData *ListKnowledgeBaseChunksRequestMetaData `json:"MetaData,omitempty" xml:"MetaData,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListKnowledgeBaseChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseChunksRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseChunksRequest) GetChunkStatus() *string {
	return s.ChunkStatus
}

func (s *ListKnowledgeBaseChunksRequest) GetMetaData() *ListKnowledgeBaseChunksRequestMetaData {
	return s.MetaData
}

func (s *ListKnowledgeBaseChunksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKnowledgeBaseChunksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKnowledgeBaseChunksRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *ListKnowledgeBaseChunksRequest) SetChunkStatus(v string) *ListKnowledgeBaseChunksRequest {
	s.ChunkStatus = &v
	return s
}

func (s *ListKnowledgeBaseChunksRequest) SetMetaData(v *ListKnowledgeBaseChunksRequestMetaData) *ListKnowledgeBaseChunksRequest {
	s.MetaData = v
	return s
}

func (s *ListKnowledgeBaseChunksRequest) SetPageNumber(v int32) *ListKnowledgeBaseChunksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKnowledgeBaseChunksRequest) SetPageSize(v int32) *ListKnowledgeBaseChunksRequest {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeBaseChunksRequest) SetVersionName(v string) *ListKnowledgeBaseChunksRequest {
	s.VersionName = &v
	return s
}

func (s *ListKnowledgeBaseChunksRequest) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeBaseChunksRequestMetaData struct {
	// 文件元数据ID
	//
	// example:
	//
	// xd8e****79du
	FileMetaId *string `json:"FileMetaId,omitempty" xml:"FileMetaId,omitempty"`
	// 文件地址
	//
	// example:
	//
	// oss://mybucketpath/file1.txt
	FileUri *string `json:"FileUri,omitempty" xml:"FileUri,omitempty"`
}

func (s ListKnowledgeBaseChunksRequestMetaData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseChunksRequestMetaData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseChunksRequestMetaData) GetFileMetaId() *string {
	return s.FileMetaId
}

func (s *ListKnowledgeBaseChunksRequestMetaData) GetFileUri() *string {
	return s.FileUri
}

func (s *ListKnowledgeBaseChunksRequestMetaData) SetFileMetaId(v string) *ListKnowledgeBaseChunksRequestMetaData {
	s.FileMetaId = &v
	return s
}

func (s *ListKnowledgeBaseChunksRequestMetaData) SetFileUri(v string) *ListKnowledgeBaseChunksRequestMetaData {
	s.FileUri = &v
	return s
}

func (s *ListKnowledgeBaseChunksRequestMetaData) Validate() error {
	return dara.Validate(s)
}
