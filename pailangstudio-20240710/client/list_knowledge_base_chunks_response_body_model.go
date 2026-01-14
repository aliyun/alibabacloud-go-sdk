// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseChunks(v []*ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) *ListKnowledgeBaseChunksResponseBody
	GetKnowledgeBaseChunks() []*ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks
	SetMaxResults(v int32) *ListKnowledgeBaseChunksResponseBody
	GetMaxResults() *int32
	SetRequestId(v string) *ListKnowledgeBaseChunksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKnowledgeBaseChunksResponseBody
	GetTotalCount() *int32
}

type ListKnowledgeBaseChunksResponseBody struct {
	// 切片列表
	KnowledgeBaseChunks []*ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks `json:"KnowledgeBaseChunks,omitempty" xml:"KnowledgeBaseChunks,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 963BD7F9-0C02-5594-9550-BCC6DD43E3C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKnowledgeBaseChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseChunksResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseChunksResponseBody) GetKnowledgeBaseChunks() []*ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	return s.KnowledgeBaseChunks
}

func (s *ListKnowledgeBaseChunksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBaseChunksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKnowledgeBaseChunksResponseBody) SetKnowledgeBaseChunks(v []*ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) *ListKnowledgeBaseChunksResponseBody {
	s.KnowledgeBaseChunks = v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBody) SetMaxResults(v int32) *ListKnowledgeBaseChunksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBody) SetRequestId(v string) *ListKnowledgeBaseChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBody) SetTotalCount(v int32) *ListKnowledgeBaseChunksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBody) Validate() error {
	if s.KnowledgeBaseChunks != nil {
		for _, item := range s.KnowledgeBaseChunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks struct {
	// 切片附属信息
	ChunkAttachment []*ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment `json:"ChunkAttachment,omitempty" xml:"ChunkAttachment,omitempty" type:"Repeated"`
	// 切片内容
	//
	// example:
	//
	// content
	ChunkContent *string `json:"ChunkContent,omitempty" xml:"ChunkContent,omitempty"`
	// 切片结束位置
	//
	// example:
	//
	// 30000
	ChunkEnd *int32 `json:"ChunkEnd,omitempty" xml:"ChunkEnd,omitempty"`
	// 切片顺序
	//
	// example:
	//
	// 1
	ChunkSequence *int32 `json:"ChunkSequence,omitempty" xml:"ChunkSequence,omitempty"`
	// 切片大小
	//
	// example:
	//
	// 1873
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// 切片起始位置
	//
	// example:
	//
	// 0
	ChunkStart *int32 `json:"ChunkStart,omitempty" xml:"ChunkStart,omitempty"`
	// 切片状态
	//
	// example:
	//
	// Enable
	ChunkStatus *string `json:"ChunkStatus,omitempty" xml:"ChunkStatus,omitempty"`
	// 切片下载地址
	//
	// example:
	//
	// https://mybucket.oss-cn-shanghai.aliyuncs.com/...
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// 切片ID
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	KnowledgeBaseChunkId *string `json:"KnowledgeBaseChunkId,omitempty" xml:"KnowledgeBaseChunkId,omitempty"`
	// 知识库ID
	//
	// example:
	//
	// d-ix****o9
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// 原始文件信息
	MetaData *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData `json:"MetaData,omitempty" xml:"MetaData,omitempty" type:"Struct"`
	// 切片缩略图
	//
	// example:
	//
	// https://mybucket.oss-cn-shanghai.aliyuncs.com/...
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// 知识库版本
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetChunkAttachment() []*ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment {
	return s.ChunkAttachment
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetChunkContent() *string {
	return s.ChunkContent
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetChunkEnd() *int32 {
	return s.ChunkEnd
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetChunkSequence() *int32 {
	return s.ChunkSequence
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetChunkStart() *int32 {
	return s.ChunkStart
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetChunkStatus() *string {
	return s.ChunkStatus
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetKnowledgeBaseChunkId() *string {
	return s.KnowledgeBaseChunkId
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetMetaData() *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData {
	return s.MetaData
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) GetVersionName() *string {
	return s.VersionName
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetChunkAttachment(v []*ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.ChunkAttachment = v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetChunkContent(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.ChunkContent = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetChunkEnd(v int32) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.ChunkEnd = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetChunkSequence(v int32) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.ChunkSequence = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetChunkSize(v int32) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.ChunkSize = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetChunkStart(v int32) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.ChunkStart = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetChunkStatus(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.ChunkStatus = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetDownloadUrl(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.DownloadUrl = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetKnowledgeBaseChunkId(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.KnowledgeBaseChunkId = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetKnowledgeBaseId(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetMetaData(v *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.MetaData = v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetThumbnailUrl(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.ThumbnailUrl = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) SetVersionName(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks {
	s.VersionName = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunks) Validate() error {
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

type ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment struct {
	// 下载地址
	//
	// example:
	//
	// https://mybucket.oss-cn-shanghai.aliyuncs.com/...
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// 占位符ID
	//
	// example:
	//
	// IMAGE_PLACEHOLDER_0
	PlaceholderId *string `json:"PlaceholderId,omitempty" xml:"PlaceholderId,omitempty"`
	// 类型
	//
	// example:
	//
	// image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 路径
	//
	// example:
	//
	// oss://mybucket/path/file.txt
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) GetPlaceholderId() *string {
	return s.PlaceholderId
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) GetType() *string {
	return s.Type
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) GetUri() *string {
	return s.Uri
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) SetDownloadUrl(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment {
	s.DownloadUrl = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) SetPlaceholderId(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment {
	s.PlaceholderId = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) SetType(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment {
	s.Type = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) SetUri(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment {
	s.Uri = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksChunkAttachment) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData struct {
	// 文件元数据ID
	//
	// example:
	//
	// xd8e****79du
	FileMetaId *string `json:"FileMetaId,omitempty" xml:"FileMetaId,omitempty"`
	// 文件名
	//
	// example:
	//
	// file1.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件地址
	//
	// example:
	//
	// oss://mybucketpath/file1.txt
	FileUri *string `json:"FileUri,omitempty" xml:"FileUri,omitempty"`
}

func (s ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) GetFileMetaId() *string {
	return s.FileMetaId
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) GetFileName() *string {
	return s.FileName
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) GetFileUri() *string {
	return s.FileUri
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) SetFileMetaId(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData {
	s.FileMetaId = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) SetFileName(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData {
	s.FileName = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) SetFileUri(v string) *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData {
	s.FileUri = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponseBodyKnowledgeBaseChunksMetaData) Validate() error {
	return dara.Validate(s)
}
