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
	// The unique identifier of the dataset.
	//
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The document.
	//
	// This parameter is required.
	Document *AddDatasetDocumentRequestDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The unique identifier of the Model Studio workspace. For more information, see [Obtain a workspaceId](https://help.aliyun.com/document_detail/2782167.html).
	//
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
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDatasetDocumentRequestDocument struct {
	// The unique identifier of the category.
	//
	// example:
	//
	// xx
	CategoryUuid *string `json:"CategoryUuid,omitempty" xml:"CategoryUuid,omitempty"`
	// The content.
	//
	// example:
	//
	// 正文
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to disable the indexing of multimodal data, such as images and videos, in the current record. The default value is true.
	//
	// example:
	//
	// false
	DisableHandleMultimodalMedia *bool `json:"DisableHandleMultimodalMedia,omitempty" xml:"DisableHandleMultimodalMedia,omitempty"`
	// The unique business ID of the document.
	//
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// The type of the document.
	//
	// - plainText: plain text. The content parameter is required.
	//
	// - richText: rich text in HTML format. The content parameter is required.
	//
	// - text: a text file. The url parameter is required.
	//
	// - pdf: a PDF file. The url parameter is required.
	//
	// - word: a Word document. The url parameter is required.
	//
	// - image: an image. The url parameter is required. Most common image formats are supported, such as GIF, PNG, JPG, and JPEG.
	//
	// - video: a video. The url parameter is required. Most common video formats are supported, such as MP4, AVI, WMV, and MOV.
	//
	// example:
	//
	// image
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The unique system ID of the document. The system automatically generates this ID. You do not need to specify this parameter.
	//
	// example:
	//
	// xxxx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// Extension field 1.
	//
	// example:
	//
	// xxx
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	// Extension field 2.
	//
	// example:
	//
	// xxxx
	Extend2 *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	// Extension field 3.
	//
	// example:
	//
	// xxx
	Extend3 *string `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	// The metadata.
	Metadata *AddDatasetDocumentRequestDocumentMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// Deprecated. This parameter is not available.
	//
	// example:
	//
	// xxxx
	MultimodalIndexName *string `json:"MultimodalIndexName,omitempty" xml:"MultimodalIndexName,omitempty"`
	// A list of multimodal data in the document.
	//
	// - If a document, such as a rich text document, contains multimodal data like images or videos, you can pass the data using this parameter. This allows the data to be retrieved in search results.
	//
	// - If the document itself is multimodal data, leave this field empty and specify the data using the docType and url parameters.
	MultimodalMedias []*AddDatasetDocumentRequestDocumentMultimodalMedias `json:"MultimodalMedias,omitempty" xml:"MultimodalMedias,omitempty" type:"Repeated"`
	// The publishing time.
	//
	// example:
	//
	// 2024-12-09 13:35:40
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// The source.
	//
	// example:
	//
	// xxx媒体
	SourceFrom *string `json:"SourceFrom,omitempty" xml:"SourceFrom,omitempty"`
	// The summary of the article.
	//
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The tag name.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The title of the document.
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the article. The URL must be accessible over the public network.
	//
	// example:
	//
	// http://xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s AddDatasetDocumentRequestDocument) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentRequestDocument) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentRequestDocument) GetCategoryUuid() *string {
	return s.CategoryUuid
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

func (s *AddDatasetDocumentRequestDocument) GetMetadata() *AddDatasetDocumentRequestDocumentMetadata {
	return s.Metadata
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

func (s *AddDatasetDocumentRequestDocument) GetTags() []*string {
	return s.Tags
}

func (s *AddDatasetDocumentRequestDocument) GetTitle() *string {
	return s.Title
}

func (s *AddDatasetDocumentRequestDocument) GetUrl() *string {
	return s.Url
}

func (s *AddDatasetDocumentRequestDocument) SetCategoryUuid(v string) *AddDatasetDocumentRequestDocument {
	s.CategoryUuid = &v
	return s
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

func (s *AddDatasetDocumentRequestDocument) SetMetadata(v *AddDatasetDocumentRequestDocumentMetadata) *AddDatasetDocumentRequestDocument {
	s.Metadata = v
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

func (s *AddDatasetDocumentRequestDocument) SetTags(v []*string) *AddDatasetDocumentRequestDocument {
	s.Tags = v
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
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.MultimodalMedias != nil {
		for _, item := range s.MultimodalMedias {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddDatasetDocumentRequestDocumentMetadata struct {
	// The speech or caption information.
	AsrSentences []*AddDatasetDocumentRequestDocumentMetadataAsrSentences `json:"AsrSentences,omitempty" xml:"AsrSentences,omitempty" type:"Repeated"`
	// The metadata in a key-value structure.
	KeyValues []*AddDatasetDocumentRequestDocumentMetadataKeyValues `json:"KeyValues,omitempty" xml:"KeyValues,omitempty" type:"Repeated"`
	// The description of the metadata. This field is deprecated.
	//
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The video shot information.
	VideoShots []*AddDatasetDocumentRequestDocumentMetadataVideoShots `json:"VideoShots,omitempty" xml:"VideoShots,omitempty" type:"Repeated"`
}

func (s AddDatasetDocumentRequestDocumentMetadata) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentRequestDocumentMetadata) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentRequestDocumentMetadata) GetAsrSentences() []*AddDatasetDocumentRequestDocumentMetadataAsrSentences {
	return s.AsrSentences
}

func (s *AddDatasetDocumentRequestDocumentMetadata) GetKeyValues() []*AddDatasetDocumentRequestDocumentMetadataKeyValues {
	return s.KeyValues
}

func (s *AddDatasetDocumentRequestDocumentMetadata) GetText() *string {
	return s.Text
}

func (s *AddDatasetDocumentRequestDocumentMetadata) GetVideoShots() []*AddDatasetDocumentRequestDocumentMetadataVideoShots {
	return s.VideoShots
}

func (s *AddDatasetDocumentRequestDocumentMetadata) SetAsrSentences(v []*AddDatasetDocumentRequestDocumentMetadataAsrSentences) *AddDatasetDocumentRequestDocumentMetadata {
	s.AsrSentences = v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadata) SetKeyValues(v []*AddDatasetDocumentRequestDocumentMetadataKeyValues) *AddDatasetDocumentRequestDocumentMetadata {
	s.KeyValues = v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadata) SetText(v string) *AddDatasetDocumentRequestDocumentMetadata {
	s.Text = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadata) SetVideoShots(v []*AddDatasetDocumentRequestDocumentMetadataVideoShots) *AddDatasetDocumentRequestDocumentMetadata {
	s.VideoShots = v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadata) Validate() error {
	if s.AsrSentences != nil {
		for _, item := range s.AsrSentences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.KeyValues != nil {
		for _, item := range s.KeyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoShots != nil {
		for _, item := range s.VideoShots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddDatasetDocumentRequestDocumentMetadataAsrSentences struct {
	// The end time in milliseconds.
	//
	// example:
	//
	// 2000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time in milliseconds.
	//
	// example:
	//
	// 1000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The speech or caption information.
	//
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AddDatasetDocumentRequestDocumentMetadataAsrSentences) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentRequestDocumentMetadataAsrSentences) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentRequestDocumentMetadataAsrSentences) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AddDatasetDocumentRequestDocumentMetadataAsrSentences) GetStartTime() *int64 {
	return s.StartTime
}

func (s *AddDatasetDocumentRequestDocumentMetadataAsrSentences) GetText() *string {
	return s.Text
}

func (s *AddDatasetDocumentRequestDocumentMetadataAsrSentences) SetEndTime(v int64) *AddDatasetDocumentRequestDocumentMetadataAsrSentences {
	s.EndTime = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadataAsrSentences) SetStartTime(v int64) *AddDatasetDocumentRequestDocumentMetadataAsrSentences {
	s.StartTime = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadataAsrSentences) SetText(v string) *AddDatasetDocumentRequestDocumentMetadataAsrSentences {
	s.Text = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadataAsrSentences) Validate() error {
	return dara.Validate(s)
}

type AddDatasetDocumentRequestDocumentMetadataKeyValues struct {
	// The name.
	//
	// example:
	//
	// xx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// xx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddDatasetDocumentRequestDocumentMetadataKeyValues) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentRequestDocumentMetadataKeyValues) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentRequestDocumentMetadataKeyValues) GetKey() *string {
	return s.Key
}

func (s *AddDatasetDocumentRequestDocumentMetadataKeyValues) GetValue() *string {
	return s.Value
}

func (s *AddDatasetDocumentRequestDocumentMetadataKeyValues) SetKey(v string) *AddDatasetDocumentRequestDocumentMetadataKeyValues {
	s.Key = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadataKeyValues) SetValue(v string) *AddDatasetDocumentRequestDocumentMetadataKeyValues {
	s.Value = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadataKeyValues) Validate() error {
	return dara.Validate(s)
}

type AddDatasetDocumentRequestDocumentMetadataVideoShots struct {
	// The end time in milliseconds.
	//
	// example:
	//
	// 2000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time in milliseconds.
	//
	// example:
	//
	// 1000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The text information from the video shot analysis.
	//
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AddDatasetDocumentRequestDocumentMetadataVideoShots) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentRequestDocumentMetadataVideoShots) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentRequestDocumentMetadataVideoShots) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AddDatasetDocumentRequestDocumentMetadataVideoShots) GetStartTime() *int64 {
	return s.StartTime
}

func (s *AddDatasetDocumentRequestDocumentMetadataVideoShots) GetText() *string {
	return s.Text
}

func (s *AddDatasetDocumentRequestDocumentMetadataVideoShots) SetEndTime(v int64) *AddDatasetDocumentRequestDocumentMetadataVideoShots {
	s.EndTime = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadataVideoShots) SetStartTime(v int64) *AddDatasetDocumentRequestDocumentMetadataVideoShots {
	s.StartTime = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadataVideoShots) SetText(v string) *AddDatasetDocumentRequestDocumentMetadataVideoShots {
	s.Text = &v
	return s
}

func (s *AddDatasetDocumentRequestDocumentMetadataVideoShots) Validate() error {
	return dara.Validate(s)
}

type AddDatasetDocumentRequestDocumentMultimodalMedias struct {
	// The URL of the file. The URL must be accessible over the public network.
	//
	// example:
	//
	// http://xxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The unique identifier of the multimodal data. The system automatically generates this ID. You do not need to specify this parameter.
	//
	// example:
	//
	// xxxx
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the multimodal data.
	//
	// - image: an image
	//
	// - video: a video
	//
	// example:
	//
	// image
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
