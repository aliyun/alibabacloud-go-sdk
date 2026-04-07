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
	// The MIME Type of the file.
	//
	// example:
	//
	// image/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 120000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The metadata ID of the dataset file.
	//
	// example:
	//
	// 07914c9534586e4e7aa6e9dbca5009082df******fd8a0d857b33296c59bf6
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	// The download URL of the file.
	//
	// example:
	//
	// https://test-bucket.oss-cn-shanghai.aliyuncs.com/dataset/cat.png?Expires=1712801702&OSSAccessKeyId=****************&Signature=x4oSjZcXOn7FHMCT1D****NuNjs%3D
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The time when the file was created. A UTC timestamp in the ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// The file fingerprint, used to determine the uniqueness of the file content. This value changes after the file content is modified. OSS files use ETags, and NAS files use MD5 values.
	//
	// example:
	//
	// D41D8CD98F*****E9800998ECF8
	FileFingerPrint *string `json:"FileFingerPrint,omitempty" xml:"FileFingerPrint,omitempty"`
	// The file name.
	//
	// example:
	//
	// cat.png
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type.
	//
	// example:
	//
	// image
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The last modified time of the file. A UTC timestamp in the ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileUpdateTime *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	// The specific metadata of the file. For example, the width and height of an image file, or the bitrate and resolution of a video file. This parameter cannot be used for retrieval. In JSON String format.
	//
	// example:
	//
	// {     "ImageHeight": 400,     "ImageWidth": 800 }
	MetaAttributes *string `json:"MetaAttributes,omitempty" xml:"MetaAttributes,omitempty"`
	// The similarity score.
	//
	// example:
	//
	// 0.6
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The ID of the last semantic indexing job.
	//
	// example:
	//
	// dsjob-klfwtjto****scvt3
	SemanticIndexJobId *string `json:"SemanticIndexJobId,omitempty" xml:"SemanticIndexJobId,omitempty"`
	// The last update time of the semantic index. A UTC timestamp in the ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	SemanticIndexUpdateTime *string `json:"SemanticIndexUpdateTime,omitempty" xml:"SemanticIndexUpdateTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the metadata, in JSON string format. Including:
	//
	// 	- Algorithm tag group:
	//
	//     	- ai: tags from all algorithm tagging tasks on this metadata.
	//
	// 	- User-defined tag group:
	//
	//     	- user: tags added by the user to this metadata.
	//
	//     	- user-delete-ai-tags: tags from the algorithm tag group that the user needs to delete.
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
	// The URL of the thumbnail.
	//
	// example:
	//
	// https://test-bucket.oss-cn-shanghai.aliyuncs.com/dataset/cat.png?Expires=171280****&OSSAccessKeyId=LTAI************&Signature=****jZcXOn7FHMCT1DLE22NuNjs%3D
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// The URI of the file. Used to record the unique path of the file. File paths in Object Storage Service (OSS) and File Storage NAS (NAS) are supported.
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
	// oss://test-bucket/dataset/cat.png
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
