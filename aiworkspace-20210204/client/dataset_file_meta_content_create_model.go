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
	// The file comment.
	//
	// example:
	//
	// The first image file in the dataset.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The MIME type of the file. Contains Type and SubType.
	//
	// This parameter is required.
	//
	// example:
	//
	// image/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The file size in bytes.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The file creation time in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// The file fingerprint value. Used to determine the uniqueness of file content. This value changes when the file content is modified. OSS files use ETag, and NAS files use MD5 values.
	//
	// This parameter is required.
	//
	// example:
	//
	// D41D8CD98F*****E9800998ECF8
	FileFingerPrint *string `json:"FileFingerPrint,omitempty" xml:"FileFingerPrint,omitempty"`
	// The file name.
	//
	// example:
	//
	// 00001.jpeg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type. Same as MIME Type.
	//
	// This parameter is required.
	//
	// example:
	//
	// image
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The last modification time of the file in ISO 8601 format.
	//
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileUpdateTime *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	// The specific metadata of the file, not searchable. In JSON string format.
	//
	// example:
	//
	// {
	//
	//     "ImageHeight": 1080,
	//
	//     "ImageWidth": 1920
	//
	// }
	MetaAttributes *string `json:"MetaAttributes,omitempty" xml:"MetaAttributes,omitempty"`
	// User manual tagging: (add indicates adding tags to the tag group). In JSON string format.
	//
	// The operable tag groups are:
	//
	// - user: The list of tag names manually added by the user for a single metadata entry.
	//
	// ```
	//
	// {
	//
	//     "user":{
	//
	//         "add":["lane_line","sunny"]
	//
	//     }
	//
	// }
	//
	// ```
	//
	// example:
	//
	// {
	//
	//     "user":{
	//
	//         "add":["Lane line","Sunny day"]
	//
	//     }
	//
	// }
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The unique URI of the file. Used to record the unique file path. Supports file paths in OSS and NAS.
	//
	// <details>
	//
	// <summary>OSS</summary>
	//
	// oss://${bucket}/${path}
	//
	// </details>
	//
	// <details>
	//
	// <summary>NAS</summary>
	//
	// nas://${fileSystemId}/${path}
	//
	// </details>
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://*****-test/dataset/1653421.jpg
	//
	// nas://0e25d***dff/dataset/1653421.jpg
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
