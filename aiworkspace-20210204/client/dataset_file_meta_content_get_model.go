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
	SetStatus(v string) *DatasetFileMetaContentGet
	GetStatus() *string
	SetTagUpdateTime(v string) *DatasetFileMetaContentGet
	GetTagUpdateTime() *string
	SetTags(v string) *DatasetFileMetaContentGet
	GetTags() *string
	SetUri(v string) *DatasetFileMetaContentGet
	GetUri() *string
}

type DatasetFileMetaContentGet struct {
	// The file comment.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The MIME type of the file. It contains a Type and a SubType.
	//
	// Valid value:
	//
	// 	- image/png: PNG
	//
	// 	- image/jpeg: JPEG
	//
	// 	- image/tiff: TIFF
	//
	// 	- image/bmp: BMP
	//
	// 	- image/gif: GIF
	//
	// 	- image/x-icon: ICON
	//
	// 	- image/svg + xml: SVG
	//
	// 	- image/heic: HEIC
	//
	// 	- image/webp: WEBP
	//
	// example:
	//
	// text/png
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The file size. Unit: byte.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The metadata ID of the dataset file.
	//
	// example:
	//
	// 07914c9534586e4e7aa6e9dbca5009082df******fd8a0d857b33296c59bf6
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	// The time when the file was created. Format: ISO8601.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// The directory of the file that is stored in OSS, NAS, or Cloud Parallel File Storage (CPFS).
	//
	// example:
	//
	// icp_certificate_card/icp/1577179298694813/1716429710367
	FileDir *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	// The fingerprint value of the file. Used to check the uniqueness of the file. This value changes after the file content is modified. OSS files use ETags, and NAS files use MD5.
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
	// The file type. The same as MIME type.
	//
	// Valid value:
	//
	// 	- image
	//
	// 	- application
	//
	// 	- audio
	//
	// 	- video
	//
	// 	- text
	//
	// example:
	//
	// image
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The time when the file was last modified. Format: ISO8601.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileUpdateTime *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	// The specific metadata of the file. You cannot retrieve the metadata. In JSON String format.
	//
	// example:
	//
	// {
	//
	//     "Image":
	//
	//     {
	//
	//         "Width": 1920,
	//
	//         "Height": 1080,
	//
	//         "Channel": 3
	//
	//     }
	//
	// }
	MetaAttributes *string `json:"MetaAttributes,omitempty" xml:"MetaAttributes,omitempty"`
	// The ID of the semantic index-based job.
	//
	// example:
	//
	// dsjob-klfwtjtov*****scvt3
	SemanticIndexJobId *string `json:"SemanticIndexJobId,omitempty" xml:"SemanticIndexJobId,omitempty"`
	// The time when the semantic index-based job is created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	SemanticIndexUpdateTime *string `json:"SemanticIndexUpdateTime,omitempty" xml:"SemanticIndexUpdateTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the tag is last modified. The time follows the ISO 8601 standard.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	TagUpdateTime *string `json:"TagUpdateTime,omitempty" xml:"TagUpdateTime,omitempty"`
	// The tags for the metadata. The tags are divided into the following groups:
	//
	// 	- Algorithm tag group:
	//
	//     	- ai: a list of tags that are aggregated by all algorithm tagging tasks for a single piece of metadata.
	//
	// 	- User-defined tag groups:
	//
	//     	- user: a list of user-defined tags that are added to a single piece of metadata.
	//
	//     	- user-delete-ai-tags: a list of tags that you want to delete from an algorithm tag group.
	//
	// example:
	//
	// {
	//
	//     "ai":
	//
	//     [
	//
	//         "Felis catus",
	//
	//         "Shorthair"
	//
	//     ],
	//
	//     "user":
	//
	//     [
	//
	//         "cat",
	//
	//         "White"
	//
	//     ]
	//
	// }
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The unique URI of the file. Used to record the unique path of the file. File paths in OSS and NAS are supported.
	//
	// **OSS**
	//
	// oss://${bucket}/${path}
	//
	// **NAS**
	//
	// nas://${fileSystemId}/${path}
	//
	// example:
	//
	// oss://*****-test/dataset/1653421.jpg
	//
	// nas://0e25d***dff/dataset/1653421.jpg
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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

func (s *DatasetFileMetaContentGet) GetStatus() *string {
	return s.Status
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

func (s *DatasetFileMetaContentGet) SetStatus(v string) *DatasetFileMetaContentGet {
	s.Status = &v
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
