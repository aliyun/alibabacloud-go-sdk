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
	// The comment on the file.
	//
	// example:
	//
	// The first image file in the dataset.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The MIME type of the file. It includes a type and a subtype.
	//
	// example:
	//
	// image/png
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The file size in bytes.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The ID of the dataset file metadata.
	//
	// example:
	//
	// 07914c9534586e4e7aa6e9dbca5009082df******fd8a0d857b33296c59bf6
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	// The time when the file was created. The time is in the ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// The path of the folder where the OSS, NAS, or CPFS file is located.
	//
	// example:
	//
	// icp_certificate_card/icp/1577179298694813/1716429710367
	FileDir *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	// The fingerprint of the file. This value ensures the uniqueness of the file content. The value changes if the file content is modified. For OSS files, the ETag is used. For NAS files, the MD5 hash is used.
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
	// The file type. This is the same as the Multipurpose Internet Mail Extensions (MIME) type.
	//
	// example:
	//
	// image
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The time when the file was last modified. The time is in the ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileUpdateTime *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	// The specific metadata of the file. This metadata cannot be used for retrieval. The format is a JSON string.
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
	// The ID of the job that builds the semantic index.
	//
	// example:
	//
	// dsjob-klfwtjtov*****scvt3
	SemanticIndexJobId *string `json:"SemanticIndexJobId,omitempty" xml:"SemanticIndexJobId,omitempty"`
	// The time when the semantic index was built.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	SemanticIndexUpdateTime *string `json:"SemanticIndexUpdateTime,omitempty" xml:"SemanticIndexUpdateTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the tag was last modified. The time is in the ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	TagUpdateTime *string `json:"TagUpdateTime,omitempty" xml:"TagUpdateTime,omitempty"`
	// A collection of tags for the metadata. It includes the following groups:
	//
	// - Algorithm tag group:
	//
	//   - ai: A list of tag names aggregated from all algorithmic tagging tasks for a single metadata record.
	//
	// - User-defined tag group:
	//
	//   - user: A list of tag names that a user adds to a single metadata record.
	//
	//   - user-delete-ai-tags: A list of tag names from the algorithm tag group that the user wants to delete from a single metadata record.
	//
	// example:
	//
	// {
	//
	//     "ai":
	//
	//     [
	//
	//         "Lane line",
	//
	//         "Water horse",
	//
	//         "Sunny day"
	//
	//     ],
	//
	//     "user":
	//
	//     [
	//
	//         "Everett",
	//
	//         "Intelligent driving Dataset 1",
	//
	//         "Cloudy day"
	//
	//     ],
	//
	//     "user-delete-ai-tags":
	//
	//     [
	//
	//         "Sunny day"
	//
	//     ]
	//
	// }
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The unique URI of the file. This URI records the unique path of the file. Paths for files in OSS and NAS are supported.
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
