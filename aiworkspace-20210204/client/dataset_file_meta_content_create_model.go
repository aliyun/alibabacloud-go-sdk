// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetFileMetaContentCreate interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DatasetFileMetaContentCreate
	GetComment() *string
	SetContentType(v string) *DatasetFileMetaContentCreate
	GetContentType() *string
	SetDataSize(v int64) *DatasetFileMetaContentCreate
	GetDataSize() *int64
	SetFileCreateTime(v string) *DatasetFileMetaContentCreate
	GetFileCreateTime() *string
	SetFileFingerPrint(v string) *DatasetFileMetaContentCreate
	GetFileFingerPrint() *string
	SetFileName(v string) *DatasetFileMetaContentCreate
	GetFileName() *string
	SetFileType(v string) *DatasetFileMetaContentCreate
	GetFileType() *string
	SetFileUpdateTime(v string) *DatasetFileMetaContentCreate
	GetFileUpdateTime() *string
	SetMetaAttributes(v string) *DatasetFileMetaContentCreate
	GetMetaAttributes() *string
	SetTags(v string) *DatasetFileMetaContentCreate
	GetTags() *string
	SetUri(v string) *DatasetFileMetaContentCreate
	GetUri() *string
}

type DatasetFileMetaContentCreate struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DataSize    *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// This parameter is required.
	FileFingerPrint *string `json:"FileFingerPrint,omitempty" xml:"FileFingerPrint,omitempty"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	FileUpdateTime *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	MetaAttributes *string `json:"MetaAttributes,omitempty" xml:"MetaAttributes,omitempty"`
	// example:
	//
	// {"user":{"add":["cat"]}}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DatasetFileMetaContentCreate) String() string {
	return dara.Prettify(s)
}

func (s DatasetFileMetaContentCreate) GoString() string {
	return s.String()
}

func (s *DatasetFileMetaContentCreate) GetComment() *string {
	return s.Comment
}

func (s *DatasetFileMetaContentCreate) GetContentType() *string {
	return s.ContentType
}

func (s *DatasetFileMetaContentCreate) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DatasetFileMetaContentCreate) GetFileCreateTime() *string {
	return s.FileCreateTime
}

func (s *DatasetFileMetaContentCreate) GetFileFingerPrint() *string {
	return s.FileFingerPrint
}

func (s *DatasetFileMetaContentCreate) GetFileName() *string {
	return s.FileName
}

func (s *DatasetFileMetaContentCreate) GetFileType() *string {
	return s.FileType
}

func (s *DatasetFileMetaContentCreate) GetFileUpdateTime() *string {
	return s.FileUpdateTime
}

func (s *DatasetFileMetaContentCreate) GetMetaAttributes() *string {
	return s.MetaAttributes
}

func (s *DatasetFileMetaContentCreate) GetTags() *string {
	return s.Tags
}

func (s *DatasetFileMetaContentCreate) GetUri() *string {
	return s.Uri
}

func (s *DatasetFileMetaContentCreate) SetComment(v string) *DatasetFileMetaContentCreate {
	s.Comment = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetContentType(v string) *DatasetFileMetaContentCreate {
	s.ContentType = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetDataSize(v int64) *DatasetFileMetaContentCreate {
	s.DataSize = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetFileCreateTime(v string) *DatasetFileMetaContentCreate {
	s.FileCreateTime = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetFileFingerPrint(v string) *DatasetFileMetaContentCreate {
	s.FileFingerPrint = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetFileName(v string) *DatasetFileMetaContentCreate {
	s.FileName = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetFileType(v string) *DatasetFileMetaContentCreate {
	s.FileType = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetFileUpdateTime(v string) *DatasetFileMetaContentCreate {
	s.FileUpdateTime = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetMetaAttributes(v string) *DatasetFileMetaContentCreate {
	s.MetaAttributes = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetTags(v string) *DatasetFileMetaContentCreate {
	s.Tags = &v
	return s
}

func (s *DatasetFileMetaContentCreate) SetUri(v string) *DatasetFileMetaContentCreate {
	s.Uri = &v
	return s
}

func (s *DatasetFileMetaContentCreate) Validate() error {
	return dara.Validate(s)
}
