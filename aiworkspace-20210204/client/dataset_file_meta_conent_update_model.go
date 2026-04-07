// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetFileMetaConentUpdate interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DatasetFileMetaConentUpdate
	GetComment() *string
	SetContentType(v string) *DatasetFileMetaConentUpdate
	GetContentType() *string
	SetDataSize(v int64) *DatasetFileMetaConentUpdate
	GetDataSize() *int64
	SetDatasetFileMetaId(v string) *DatasetFileMetaConentUpdate
	GetDatasetFileMetaId() *string
	SetFileCreateTime(v string) *DatasetFileMetaConentUpdate
	GetFileCreateTime() *string
	SetFileFingerPrint(v string) *DatasetFileMetaConentUpdate
	GetFileFingerPrint() *string
	SetFileName(v string) *DatasetFileMetaConentUpdate
	GetFileName() *string
	SetFileType(v string) *DatasetFileMetaConentUpdate
	GetFileType() *string
	SetFileUpdateTime(v string) *DatasetFileMetaConentUpdate
	GetFileUpdateTime() *string
	SetMetaAttributes(v string) *DatasetFileMetaConentUpdate
	GetMetaAttributes() *string
	SetSemanticIndexJobId(v string) *DatasetFileMetaConentUpdate
	GetSemanticIndexJobId() *string
	SetSemanticIndexUpdateTime(v string) *DatasetFileMetaConentUpdate
	GetSemanticIndexUpdateTime() *string
	SetTags(v string) *DatasetFileMetaConentUpdate
	GetTags() *string
}

type DatasetFileMetaConentUpdate struct {
	// The file comment.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The MIME type of the file. The value consists of a type and a subtype.
	//
	// Valid values:
	//
	// 	- image/png
	//
	// 	- image/svg+xml
	//
	// 	- image/jpeg
	//
	// 	- image/tiff
	//
	// 	- image/gif
	//
	// 	- image/bmp
	//
	// 	- image/x-icon
	//
	// 	- image/heic
	//
	// 	- image/webp
	//
	// example:
	//
	// image/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The file size. Unit: byte.
	//
	// example:
	//
	// 10000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The metadata ID of the dataset file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 07914c9534586e4e7aa6e9dbca5009082df******fd8a0d857b33296c59bf6
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	// The time when the file is created. The time follows the ISO 8601 standard.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// The fingerprint information of the file.
	//
	// example:
	//
	// 124FB71******7BE45608CDEA4DE54B3
	FileFingerPrint *string `json:"FileFingerPrint,omitempty" xml:"FileFingerPrint,omitempty"`
	// The file name.
	//
	// example:
	//
	// 00001.jpeg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type, which is the same as Multipurpose Internet Mail Extensions (MIME) type.
	//
	// Valid values:
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
	// The time when the file is last modified. The time follows the ISO 8601 standard.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-12T14:36:01Z
	FileUpdateTime *string `json:"FileUpdateTime,omitempty" xml:"FileUpdateTime,omitempty"`
	// The specific metadata of the file, such as the width and height of an image and the bitrate and resolution of a video file. You cannot retrieve the metadata. The value is a JSON string.
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
	// dsjob-klfwt*****l0escvt3
	SemanticIndexJobId *string `json:"SemanticIndexJobId,omitempty" xml:"SemanticIndexJobId,omitempty"`
	// The time when the semantic index is created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	SemanticIndexUpdateTime *string `json:"SemanticIndexUpdateTime,omitempty" xml:"SemanticIndexUpdateTime,omitempty"`
	// The tags to be updated.
	//
	// 	- Update an algorithm tag group (a valid TagJobId must be set):
	//
	// <!---->
	//
	//     {
	//
	//        "ai":["Lane line", "Water horse", "Sunny day"]
	//
	//     }
	//
	// 	- Update a user-defined tag group (add or remove indicates that tags are added or deleted): Tag groups that can be updated:
	//
	//     	- user: a list of user-defined tags that can be added to or deleted from a single piece of metadata.
	//
	//     	- user-delete-ai-tags: a list of tags that you want to delete from an algorithm tag group.
	//
	// <!---->
	//
	//     {
	//
	//         "user":{
	//
	//             "add":["Lane line","Sunny day"],
	//
	//             "remove":["Water horse"]    },
	//
	//         "user-delete-ai-tags":{
	//
	//             "add": ["Ground shadow"],
	//
	//             "remove": []
	//
	//         }
	//
	//     }
	//
	// example:
	//
	// {"ai":["cat"], "user":["black"]}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DatasetFileMetaConentUpdate) String() string {
	return dara.Prettify(s)
}

func (s DatasetFileMetaConentUpdate) GoString() string {
	return s.String()
}

func (s *DatasetFileMetaConentUpdate) GetComment() *string {
	return s.Comment
}

func (s *DatasetFileMetaConentUpdate) GetContentType() *string {
	return s.ContentType
}

func (s *DatasetFileMetaConentUpdate) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DatasetFileMetaConentUpdate) GetDatasetFileMetaId() *string {
	return s.DatasetFileMetaId
}

func (s *DatasetFileMetaConentUpdate) GetFileCreateTime() *string {
	return s.FileCreateTime
}

func (s *DatasetFileMetaConentUpdate) GetFileFingerPrint() *string {
	return s.FileFingerPrint
}

func (s *DatasetFileMetaConentUpdate) GetFileName() *string {
	return s.FileName
}

func (s *DatasetFileMetaConentUpdate) GetFileType() *string {
	return s.FileType
}

func (s *DatasetFileMetaConentUpdate) GetFileUpdateTime() *string {
	return s.FileUpdateTime
}

func (s *DatasetFileMetaConentUpdate) GetMetaAttributes() *string {
	return s.MetaAttributes
}

func (s *DatasetFileMetaConentUpdate) GetSemanticIndexJobId() *string {
	return s.SemanticIndexJobId
}

func (s *DatasetFileMetaConentUpdate) GetSemanticIndexUpdateTime() *string {
	return s.SemanticIndexUpdateTime
}

func (s *DatasetFileMetaConentUpdate) GetTags() *string {
	return s.Tags
}

func (s *DatasetFileMetaConentUpdate) SetComment(v string) *DatasetFileMetaConentUpdate {
	s.Comment = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetContentType(v string) *DatasetFileMetaConentUpdate {
	s.ContentType = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetDataSize(v int64) *DatasetFileMetaConentUpdate {
	s.DataSize = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetDatasetFileMetaId(v string) *DatasetFileMetaConentUpdate {
	s.DatasetFileMetaId = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetFileCreateTime(v string) *DatasetFileMetaConentUpdate {
	s.FileCreateTime = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetFileFingerPrint(v string) *DatasetFileMetaConentUpdate {
	s.FileFingerPrint = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetFileName(v string) *DatasetFileMetaConentUpdate {
	s.FileName = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetFileType(v string) *DatasetFileMetaConentUpdate {
	s.FileType = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetFileUpdateTime(v string) *DatasetFileMetaConentUpdate {
	s.FileUpdateTime = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetMetaAttributes(v string) *DatasetFileMetaConentUpdate {
	s.MetaAttributes = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetSemanticIndexJobId(v string) *DatasetFileMetaConentUpdate {
	s.SemanticIndexJobId = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetSemanticIndexUpdateTime(v string) *DatasetFileMetaConentUpdate {
	s.SemanticIndexUpdateTime = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) SetTags(v string) *DatasetFileMetaConentUpdate {
	s.Tags = &v
	return s
}

func (s *DatasetFileMetaConentUpdate) Validate() error {
	return dara.Validate(s)
}
