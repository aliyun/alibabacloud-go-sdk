// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetFileMeta interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *DatasetFileMeta
	GetContentType() *string
	SetDataSize(v int64) *DatasetFileMeta
	GetDataSize() *int64
	SetDatasetFileMetaId(v string) *DatasetFileMeta
	GetDatasetFileMetaId() *string
	SetDownloadUrl(v string) *DatasetFileMeta
	GetDownloadUrl() *string
	SetFileCreateTime(v string) *DatasetFileMeta
	GetFileCreateTime() *string
	SetFileFingerPrint(v string) *DatasetFileMeta
	GetFileFingerPrint() *string
	SetFileName(v string) *DatasetFileMeta
	GetFileName() *string
	SetFileType(v string) *DatasetFileMeta
	GetFileType() *string
	SetFileUpdateTime(v string) *DatasetFileMeta
	GetFileUpdateTime() *string
	SetMetaAttributes(v string) *DatasetFileMeta
	GetMetaAttributes() *string
	SetScore(v float32) *DatasetFileMeta
	GetScore() *float32
	SetSemanticIndexJobId(v string) *DatasetFileMeta
	GetSemanticIndexJobId() *string
	SetSemanticIndexUpdateTime(v string) *DatasetFileMeta
	GetSemanticIndexUpdateTime() *string
	SetStatus(v string) *DatasetFileMeta
	GetStatus() *string
	SetTags(v string) *DatasetFileMeta
	GetTags() *string
	SetThumbnailUrl(v string) *DatasetFileMeta
	GetThumbnailUrl() *string
	SetUri(v string) *DatasetFileMeta
	GetUri() *string
}

type DatasetFileMeta struct {
	// example:
	//
	// image/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 12
	DataSize          *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	DownloadUrl       *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	FileCreateTime  *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	FileFingerPrint *string `json:"FileFingerPrint,omitempty" xml:"FileFingerPrint,omitempty"`
	// example:
	//
	// car.png
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// image
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	FileUpdateTime *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	// example:
	//
	// {     "ImageHeight": 400,     "ImageWidth": 800 }
	MetaAttributes *string `json:"MetaAttributes,omitempty" xml:"MetaAttributes,omitempty"`
	// example:
	//
	// 0.6
	Score              *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	SemanticIndexJobId *string  `json:"SemanticIndexJobId,omitempty" xml:"SemanticIndexJobId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	SemanticIndexUpdateTime *string `json:"SemanticIndexUpdateTime,omitempty" xml:"SemanticIndexUpdateTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                    *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ThumbnailUrl            *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// example:
	//
	// oss://test-bucket/dataset/car.png
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DatasetFileMeta) String() string {
	return dara.Prettify(s)
}

func (s DatasetFileMeta) GoString() string {
	return s.String()
}

func (s *DatasetFileMeta) GetContentType() *string {
	return s.ContentType
}

func (s *DatasetFileMeta) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DatasetFileMeta) GetDatasetFileMetaId() *string {
	return s.DatasetFileMetaId
}

func (s *DatasetFileMeta) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DatasetFileMeta) GetFileCreateTime() *string {
	return s.FileCreateTime
}

func (s *DatasetFileMeta) GetFileFingerPrint() *string {
	return s.FileFingerPrint
}

func (s *DatasetFileMeta) GetFileName() *string {
	return s.FileName
}

func (s *DatasetFileMeta) GetFileType() *string {
	return s.FileType
}

func (s *DatasetFileMeta) GetFileUpdateTime() *string {
	return s.FileUpdateTime
}

func (s *DatasetFileMeta) GetMetaAttributes() *string {
	return s.MetaAttributes
}

func (s *DatasetFileMeta) GetScore() *float32 {
	return s.Score
}

func (s *DatasetFileMeta) GetSemanticIndexJobId() *string {
	return s.SemanticIndexJobId
}

func (s *DatasetFileMeta) GetSemanticIndexUpdateTime() *string {
	return s.SemanticIndexUpdateTime
}

func (s *DatasetFileMeta) GetStatus() *string {
	return s.Status
}

func (s *DatasetFileMeta) GetTags() *string {
	return s.Tags
}

func (s *DatasetFileMeta) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *DatasetFileMeta) GetUri() *string {
	return s.Uri
}

func (s *DatasetFileMeta) SetContentType(v string) *DatasetFileMeta {
	s.ContentType = &v
	return s
}

func (s *DatasetFileMeta) SetDataSize(v int64) *DatasetFileMeta {
	s.DataSize = &v
	return s
}

func (s *DatasetFileMeta) SetDatasetFileMetaId(v string) *DatasetFileMeta {
	s.DatasetFileMetaId = &v
	return s
}

func (s *DatasetFileMeta) SetDownloadUrl(v string) *DatasetFileMeta {
	s.DownloadUrl = &v
	return s
}

func (s *DatasetFileMeta) SetFileCreateTime(v string) *DatasetFileMeta {
	s.FileCreateTime = &v
	return s
}

func (s *DatasetFileMeta) SetFileFingerPrint(v string) *DatasetFileMeta {
	s.FileFingerPrint = &v
	return s
}

func (s *DatasetFileMeta) SetFileName(v string) *DatasetFileMeta {
	s.FileName = &v
	return s
}

func (s *DatasetFileMeta) SetFileType(v string) *DatasetFileMeta {
	s.FileType = &v
	return s
}

func (s *DatasetFileMeta) SetFileUpdateTime(v string) *DatasetFileMeta {
	s.FileUpdateTime = &v
	return s
}

func (s *DatasetFileMeta) SetMetaAttributes(v string) *DatasetFileMeta {
	s.MetaAttributes = &v
	return s
}

func (s *DatasetFileMeta) SetScore(v float32) *DatasetFileMeta {
	s.Score = &v
	return s
}

func (s *DatasetFileMeta) SetSemanticIndexJobId(v string) *DatasetFileMeta {
	s.SemanticIndexJobId = &v
	return s
}

func (s *DatasetFileMeta) SetSemanticIndexUpdateTime(v string) *DatasetFileMeta {
	s.SemanticIndexUpdateTime = &v
	return s
}

func (s *DatasetFileMeta) SetStatus(v string) *DatasetFileMeta {
	s.Status = &v
	return s
}

func (s *DatasetFileMeta) SetTags(v string) *DatasetFileMeta {
	s.Tags = &v
	return s
}

func (s *DatasetFileMeta) SetThumbnailUrl(v string) *DatasetFileMeta {
	s.ThumbnailUrl = &v
	return s
}

func (s *DatasetFileMeta) SetUri(v string) *DatasetFileMeta {
	s.Uri = &v
	return s
}

func (s *DatasetFileMeta) Validate() error {
	return dara.Validate(s)
}
