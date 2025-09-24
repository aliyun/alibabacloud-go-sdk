// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetFileMetaContentGet interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DatasetFileMetaContentGet
	GetComment() *string
	SetContentType(v string) *DatasetFileMetaContentGet
	GetContentType() *string
	SetDataSize(v int64) *DatasetFileMetaContentGet
	GetDataSize() *int64
	SetDatasetFileMetaId(v string) *DatasetFileMetaContentGet
	GetDatasetFileMetaId() *string
	SetFileCreateTime(v string) *DatasetFileMetaContentGet
	GetFileCreateTime() *string
	SetFileDir(v string) *DatasetFileMetaContentGet
	GetFileDir() *string
	SetFileFingerPrint(v string) *DatasetFileMetaContentGet
	GetFileFingerPrint() *string
	SetFileName(v string) *DatasetFileMetaContentGet
	GetFileName() *string
	SetFileType(v string) *DatasetFileMetaContentGet
	GetFileType() *string
	SetFileUpdateTime(v string) *DatasetFileMetaContentGet
	GetFileUpdateTime() *string
	SetMetaAttributes(v string) *DatasetFileMetaContentGet
	GetMetaAttributes() *string
	SetSemanticIndexJobId(v string) *DatasetFileMetaContentGet
	GetSemanticIndexJobId() *string
	SetSemanticIndexUpdateTime(v string) *DatasetFileMetaContentGet
	GetSemanticIndexUpdateTime() *string
	SetTagUpdateTime(v string) *DatasetFileMetaContentGet
	GetTagUpdateTime() *string
	SetTags(v string) *DatasetFileMetaContentGet
	GetTags() *string
	SetUri(v string) *DatasetFileMetaContentGet
	GetUri() *string
}

type DatasetFileMetaContentGet struct {
	Comment           *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ContentType       *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DataSize          *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	FileCreateTime  *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	FileDir         *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	FileFingerPrint *string `json:"FileFingerPrint,omitempty" xml:"FileFingerPrint,omitempty"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileType        *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	FileUpdateTime     *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	MetaAttributes     *string `json:"MetaAttributes,omitempty" xml:"MetaAttributes,omitempty"`
	SemanticIndexJobId *string `json:"SemanticIndexJobId,omitempty" xml:"SemanticIndexJobId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	SemanticIndexUpdateTime *string `json:"SemanticIndexUpdateTime,omitempty" xml:"SemanticIndexUpdateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	TagUpdateTime *string `json:"TagUpdateTime,omitempty" xml:"TagUpdateTime,omitempty"`
	Tags          *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Uri           *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DatasetFileMetaContentGet) String() string {
	return dara.Prettify(s)
}

func (s DatasetFileMetaContentGet) GoString() string {
	return s.String()
}

func (s *DatasetFileMetaContentGet) GetComment() *string {
	return s.Comment
}

func (s *DatasetFileMetaContentGet) GetContentType() *string {
	return s.ContentType
}

func (s *DatasetFileMetaContentGet) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DatasetFileMetaContentGet) GetDatasetFileMetaId() *string {
	return s.DatasetFileMetaId
}

func (s *DatasetFileMetaContentGet) GetFileCreateTime() *string {
	return s.FileCreateTime
}

func (s *DatasetFileMetaContentGet) GetFileDir() *string {
	return s.FileDir
}

func (s *DatasetFileMetaContentGet) GetFileFingerPrint() *string {
	return s.FileFingerPrint
}

func (s *DatasetFileMetaContentGet) GetFileName() *string {
	return s.FileName
}

func (s *DatasetFileMetaContentGet) GetFileType() *string {
	return s.FileType
}

func (s *DatasetFileMetaContentGet) GetFileUpdateTime() *string {
	return s.FileUpdateTime
}

func (s *DatasetFileMetaContentGet) GetMetaAttributes() *string {
	return s.MetaAttributes
}

func (s *DatasetFileMetaContentGet) GetSemanticIndexJobId() *string {
	return s.SemanticIndexJobId
}

func (s *DatasetFileMetaContentGet) GetSemanticIndexUpdateTime() *string {
	return s.SemanticIndexUpdateTime
}

func (s *DatasetFileMetaContentGet) GetTagUpdateTime() *string {
	return s.TagUpdateTime
}

func (s *DatasetFileMetaContentGet) GetTags() *string {
	return s.Tags
}

func (s *DatasetFileMetaContentGet) GetUri() *string {
	return s.Uri
}

func (s *DatasetFileMetaContentGet) SetComment(v string) *DatasetFileMetaContentGet {
	s.Comment = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetContentType(v string) *DatasetFileMetaContentGet {
	s.ContentType = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetDataSize(v int64) *DatasetFileMetaContentGet {
	s.DataSize = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetDatasetFileMetaId(v string) *DatasetFileMetaContentGet {
	s.DatasetFileMetaId = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetFileCreateTime(v string) *DatasetFileMetaContentGet {
	s.FileCreateTime = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetFileDir(v string) *DatasetFileMetaContentGet {
	s.FileDir = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetFileFingerPrint(v string) *DatasetFileMetaContentGet {
	s.FileFingerPrint = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetFileName(v string) *DatasetFileMetaContentGet {
	s.FileName = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetFileType(v string) *DatasetFileMetaContentGet {
	s.FileType = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetFileUpdateTime(v string) *DatasetFileMetaContentGet {
	s.FileUpdateTime = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetMetaAttributes(v string) *DatasetFileMetaContentGet {
	s.MetaAttributes = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetSemanticIndexJobId(v string) *DatasetFileMetaContentGet {
	s.SemanticIndexJobId = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetSemanticIndexUpdateTime(v string) *DatasetFileMetaContentGet {
	s.SemanticIndexUpdateTime = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetTagUpdateTime(v string) *DatasetFileMetaContentGet {
	s.TagUpdateTime = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetTags(v string) *DatasetFileMetaContentGet {
	s.Tags = &v
	return s
}

func (s *DatasetFileMetaContentGet) SetUri(v string) *DatasetFileMetaContentGet {
	s.Uri = &v
	return s
}

func (s *DatasetFileMetaContentGet) Validate() error {
	return dara.Validate(s)
}
