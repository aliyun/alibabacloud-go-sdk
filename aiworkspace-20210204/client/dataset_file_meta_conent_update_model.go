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
	Comment     *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DataSize    *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// This parameter is required.
	DatasetFileMetaId *string `json:"DatasetFileMetaId,omitempty" xml:"DatasetFileMetaId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:01.000Z
	FileCreateTime  *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
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
	// example:
	//
	// ● 执行算法打标更新（必须设置有效的TagJobId）： {    "ai":["车道线", "水马", "晴天"] } ● 用户手动打标：(add/remove表示对标签组内的标签进行增加/删除操作) 用户可操作的标签组为：   ○ user: 对单个元数据，用户自行添加的标签名列表。   ○ user-delete-ai-tags: 对单个元数据，算法标签组中用户需要删除的标签名列表。 {     "user":{         "add":["车道线","晴天"],         "remove":["水马"]     },     "user-delete-ai-tags":{         "add": ["地面阴影"],         "remove": []     } }
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
