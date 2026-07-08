// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatasetDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchDatasetDocumentsResponseBody
	GetCode() *string
	SetData(v *SearchDatasetDocumentsResponseBodyData) *SearchDatasetDocumentsResponseBody
	GetData() *SearchDatasetDocumentsResponseBodyData
	SetHttpStatusCode(v int32) *SearchDatasetDocumentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SearchDatasetDocumentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchDatasetDocumentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchDatasetDocumentsResponseBody
	GetSuccess() *bool
}

type SearchDatasetDocumentsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *SearchDatasetDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates if the request succeeded (`true`) or failed (`false`).
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchDatasetDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchDatasetDocumentsResponseBody) GetData() *SearchDatasetDocumentsResponseBodyData {
	return s.Data
}

func (s *SearchDatasetDocumentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchDatasetDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchDatasetDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchDatasetDocumentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchDatasetDocumentsResponseBody) SetCode(v string) *SearchDatasetDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetData(v *SearchDatasetDocumentsResponseBodyData) *SearchDatasetDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetHttpStatusCode(v int32) *SearchDatasetDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetMessage(v string) *SearchDatasetDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetRequestId(v string) *SearchDatasetDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) SetSuccess(v bool) *SearchDatasetDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchDatasetDocumentsResponseBodyData struct {
	// The document list.
	Documents []*SearchDatasetDocumentsResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
}

func (s SearchDatasetDocumentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsResponseBodyData) GetDocuments() []*SearchDatasetDocumentsResponseBodyDataDocuments {
	return s.Documents
}

func (s *SearchDatasetDocumentsResponseBodyData) SetDocuments(v []*SearchDatasetDocumentsResponseBodyDataDocuments) *SearchDatasetDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyData) Validate() error {
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchDatasetDocumentsResponseBodyDataDocuments struct {
	// The unique identifier for the category.
	//
	// example:
	//
	// xx
	CategoryUuid *string `json:"CategoryUuid,omitempty" xml:"CategoryUuid,omitempty"`
	// The content of the relevant chunk. This field is returned only in `chunk` mode.
	//
	// example:
	//
	// xx
	Chunk *string `json:"Chunk,omitempty" xml:"Chunk,omitempty"`
	// A list of relevant chunks from the document. This field is returned only in `document` mode.
	ChunkInfos []*SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos `json:"ChunkInfos,omitempty" xml:"ChunkInfos,omitempty" type:"Repeated"`
	// The content of the document.
	//
	// example:
	//
	// xx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The user-defined unique ID for the document.
	//
	// example:
	//
	// xx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// The document type.
	//
	// example:
	//
	// text
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The unique system ID of the document.
	//
	// example:
	//
	// xxx
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// Custom extension field 1.
	//
	// example:
	//
	// xx
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	// Custom extension field 2.
	//
	// example:
	//
	// xx
	Extend2 *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	// Custom extension field 3.
	//
	// example:
	//
	// xx
	Extend3 *string `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	// The publication time, in `yyyy-MM-dd HH:mm:ss` format.
	//
	// example:
	//
	// 2024-12-09 17:09:40
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// The relevance score.
	//
	// example:
	//
	// 0.5
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The unique identifier for the dataset.
	//
	// example:
	//
	// xx
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// xx
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// The dataset type.
	//
	// example:
	//
	// xx
	SearchSourceType *string `json:"SearchSourceType,omitempty" xml:"SearchSourceType,omitempty"`
	// The source of the document.
	//
	// example:
	//
	// 来源
	SourceFrom *string `json:"SourceFrom,omitempty" xml:"SourceFrom,omitempty"`
	// The summary of the document.
	//
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// A list of tags.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The title of the document.
	//
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the document.
	//
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SearchDatasetDocumentsResponseBodyDataDocuments) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetCategoryUuid() *string {
	return s.CategoryUuid
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetChunk() *string {
	return s.Chunk
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetChunkInfos() []*SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos {
	return s.ChunkInfos
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetContent() *string {
	return s.Content
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetDocId() *string {
	return s.DocId
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetDocType() *string {
	return s.DocType
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetDocUuid() *string {
	return s.DocUuid
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetExtend1() *string {
	return s.Extend1
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetExtend2() *string {
	return s.Extend2
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetExtend3() *string {
	return s.Extend3
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetPubTime() *string {
	return s.PubTime
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetScore() *float64 {
	return s.Score
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetSearchSource() *string {
	return s.SearchSource
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetSearchSourceType() *string {
	return s.SearchSourceType
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetSourceFrom() *string {
	return s.SourceFrom
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetSummary() *string {
	return s.Summary
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetTags() []*string {
	return s.Tags
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetTitle() *string {
	return s.Title
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) GetUrl() *string {
	return s.Url
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetCategoryUuid(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.CategoryUuid = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetChunk(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Chunk = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetChunkInfos(v []*SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.ChunkInfos = v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetContent(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Content = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetDocId(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.DocId = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetDocType(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.DocType = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetDocUuid(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.DocUuid = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetExtend1(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Extend1 = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetExtend2(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Extend2 = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetExtend3(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Extend3 = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetPubTime(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.PubTime = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetScore(v float64) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Score = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetSearchSource(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.SearchSource = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetSearchSourceName(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.SearchSourceName = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetSearchSourceType(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.SearchSourceType = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetSourceFrom(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.SourceFrom = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetSummary(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Summary = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetTags(v []*string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Tags = v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetTitle(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Title = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) SetUrl(v string) *SearchDatasetDocumentsResponseBodyDataDocuments {
	s.Url = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocuments) Validate() error {
	if s.ChunkInfos != nil {
		for _, item := range s.ChunkInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos struct {
	// The content of the chunk.
	//
	// example:
	//
	// xx
	Chunk *string `json:"Chunk,omitempty" xml:"Chunk,omitempty"`
	// The relevance score of the chunk.
	//
	// example:
	//
	// 0.77
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos) GetChunk() *string {
	return s.Chunk
}

func (s *SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos) GetScore() *float64 {
	return s.Score
}

func (s *SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos) SetChunk(v string) *SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos {
	s.Chunk = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos) SetScore(v float64) *SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos {
	s.Score = &v
	return s
}

func (s *SearchDatasetDocumentsResponseBodyDataDocumentsChunkInfos) Validate() error {
	return dara.Validate(s)
}
