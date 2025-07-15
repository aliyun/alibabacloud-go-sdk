// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatasetDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *AddDatasetDocumentRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *AddDatasetDocumentRequest
	GetDatasetName() *string
	SetDocument(v *AddDatasetDocumentRequestDocument) *AddDatasetDocumentRequest
	GetDocument() *AddDatasetDocumentRequestDocument
	SetWorkspaceId(v string) *AddDatasetDocumentRequest
	GetWorkspaceId() *string
}

type AddDatasetDocumentRequest struct {
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	Document *AddDatasetDocumentRequestDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddDatasetDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentRequest) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *AddDatasetDocumentRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *AddDatasetDocumentRequest) GetDocument() *AddDatasetDocumentRequestDocument {
	return s.Document
}

func (s *AddDatasetDocumentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddDatasetDocumentRequest) SetDatasetId(v int64) *AddDatasetDocumentRequest {
	s.DatasetId = &v
	return s
}

func (s *AddDatasetDocumentRequest) SetDatasetName(v string) *AddDatasetDocumentRequest {
	s.DatasetName = &v
	return s
}

func (s *AddDatasetDocumentRequest) SetDocument(v *AddDatasetDocumentRequestDocument) *AddDatasetDocumentRequest {
	s.Document = v
	return s
}

func (s *AddDatasetDocumentRequest) SetWorkspaceId(v string) *AddDatasetDocumentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddDatasetDocumentRequest) Validate() error {
	return dara.Validate(s)
}

type AddDatasetDocumentRequestDocument struct {
	// example:
	//
	// xxx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// false
	DisableHandleMultimodalMedia *bool `json:"DisableHandleMultimodalMedia,omitempty" xml:"DisableHandleMultimodalMedia,omitempty"`
	// example:
	//
	// 业务文档唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 文档类型
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// 内部文档唯一ID
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 扩展字段1
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	// example:
	//
	// 扩展字段2
	Extend2 *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	// example:
	//
	// 扩展字段3
	Extend3 *string `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	// example:
	//
	// 模型名称 todo 商业化 仅个别账号可传入
	MultimodalIndexName *string                                              `json:"MultimodalIndexName,omitempty" xml:"MultimodalIndexName,omitempty"`
	MultimodalMedias    []*AddDatasetDocumentRequestDocumentMultimodalMedias `json:"MultimodalMedias,omitempty" xml:"MultimodalMedias,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-12-09 13:35:40
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 来源
	SourceFrom *string `json:"SourceFrom,omitempty" xml:"SourceFrom,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// xxxxx@xxxxx.com
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s AddDatasetDocumentRequestDocument) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentRequestDocument) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentRequestDocument) GetContent() *string {
	return s.Content
}

func (s *AddDatasetDocumentRequestDocument) GetDisableHandleMultimodalMedia() *bool {
	return s.DisableHandleMultimodalMedia
}

func (s *AddDatasetDocumentRequestDocument) GetDocId() *string {
	return s.DocId
}

func (s *AddDatasetDocumentRequestDocument) GetDocType() *string {
	return s.DocType
}

func (s *AddDatasetDocumentRequestDocument) GetDocUuid() *string {
	return s.DocUuid
}

func (s *AddDatasetDocumentRequestDocument) GetExtend1() *string {
	return s.Extend1
}

func (s *AddDatasetDocumentRequestDocument) GetExtend2() *string {
	return s.Extend2
}

func (s *AddDatasetDocumentRequestDocument) GetExtend3() *string {
	return s.Extend3
}

func (s *AddDatasetDocumentRequestDocument) GetMultimodalIndexName() *string {
	return s.MultimodalIndexName
}

func (s *AddDatasetDocumentRequestDocument) GetMultimodalMedias() []*AddDatasetDocumentRequestDocumentMultimodalMedias {
	return s.MultimodalMedias
}

func (s *AddDatasetDocumentRequestDocument) GetPubTime() *string {
	return s.PubTime
}

func (s *AddDatasetDocumentRequestDocument) GetSourceFrom() *string {
	return s.SourceFrom
}

func (s *AddDatasetDocumentRequestDocument) GetSummary() *string {
	return s.Summary
}

func (s *AddDatasetDocumentRequestDocument) GetTitle() *string {
	return s.Title
}

func (s *AddDatasetDocumentRequestDocument) GetUrl() *string {
	return s.Url
}

func (s *AddDatasetDocumentRequestDocument) SetContent(v string) *AddDatasetDocumentRequestDocument {
	s.Content = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetDisableHandleMultimodalMedia(v bool) *AddDatasetDocumentRequestDocument {
	s.DisableHandleMultimodalMedia = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetDocId(v string) *AddDatasetDocumentRequestDocument {
	s.DocId = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetDocType(v string) *AddDatasetDocumentRequestDocument {
	s.DocType = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetDocUuid(v string) *AddDatasetDocumentRequestDocument {
	s.DocUuid = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetExtend1(v string) *AddDatasetDocumentRequestDocument {
	s.Extend1 = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetExtend2(v string) *AddDatasetDocumentRequestDocument {
	s.Extend2 = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetExtend3(v string) *AddDatasetDocumentRequestDocument {
	s.Extend3 = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetMultimodalIndexName(v string) *AddDatasetDocumentRequestDocument {
	s.MultimodalIndexName = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetMultimodalMedias(v []*AddDatasetDocumentRequestDocumentMultimodalMedias) *AddDatasetDocumentRequestDocument {
	s.MultimodalMedias = v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetPubTime(v string) *AddDatasetDocumentRequestDocument {
	s.PubTime = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetSourceFrom(v string) *AddDatasetDocumentRequestDocument {
	s.SourceFrom = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetSummary(v string) *AddDatasetDocumentRequestDocument {
	s.Summary = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetTitle(v string) *AddDatasetDocumentRequestDocument {
	s.Title = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) SetUrl(v string) *AddDatasetDocumentRequestDocument {
	s.Url = &v
	return s
}

func (s *AddDatasetDocumentRequestDocument) Validate() error {
	return dara.Validate(s)
}

type AddDatasetDocumentRequestDocumentMultimodalMedias struct {
	// example:
	//
	// 图片或视频文件地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 多模态数据唯一标识
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 多模态数据类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s AddDatasetDocumentRequestDocumentMultimodalMedias) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentRequestDocumentMultimodalMedias) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentRequestDocumentMultimodalMedias) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AddDatasetDocumentRequestDocumentMultimodalMedias) GetMediaId() *string {
	return s.MediaId
}

func (s *AddDatasetDocumentRequestDocumentMultimodalMedias) GetMediaType() *string {
	return s.MediaType
}

func (s *AddDatasetDocumentRequestDocumentMultimodalMedias) SetFileUrl(v string) *AddDatasetDocumentRequestDocumentMultimodalMedias {
	s.FileUrl = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMultimodalMedias) SetMediaId(v string) *AddDatasetDocumentRequestDocumentMultimodalMedias {
	s.MediaId = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMultimodalMedias) SetMediaType(v string) *AddDatasetDocumentRequestDocumentMultimodalMedias {
	s.MediaType = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMultimodalMedias) Validate() error {
	return dara.Validate(s)
}
