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
	// The comment on the file.
	//
	// example:
	//
	// The first image file in the dataset.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The MIME type of the file. It includes a type and a subtype.
	//
	// This parameter is required.
	//
	// example:
	//
	// image/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The size of the file in bytes.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The time when the file was created. The time is in ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// The fingerprint of the file. This value ensures the uniqueness of the file content and changes if the content is modified. For OSS files, this is the ETag. For NAS files, this is the MD5 value.
	//
	// This parameter is required.
	//
	// example:
	//
	// D41D8CD98F*****E9800998ECF8
	FileFingerPrint *string `json:"FileFingerPrint,omitempty" xml:"FileFingerPrint,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// 00001.jpeg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the file. This is the same as the Multipurpose Internet Mail Extensions (MIME) type.
	//
	// This parameter is required.
	//
	// example:
	//
	// image
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The time when the file was last modified. The time is in ISO 8601 format.
	//
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileUpdateTime *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	// The specific metadata of the file. This metadata cannot be used for retrieval. The value must be a JSON string.
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
	// The tags that are manually added by users. The \\`add\\` operation is used to add tags to a tag group. The value must be a JSON string.
	//
	// The following tag group is available:
	//
	// - user: A list of tag names added to a single piece of metadata.
	//
	// ```
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
	// The unique URI of the file. This URI records the unique path of the file. The path can be an OSS or NAS path.
	//
	// <details>
	//
	// <summary>
	//
	// OSS
	//
	// </summary>
	//
	// oss\\://${bucket}/${path}
	//
	// </details>
	//
	// <details>
	//
	// <summary>
	//
	// NAS
	//
	// </summary>
	//
	// nas\\://${fileSystemId}/${path}
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
