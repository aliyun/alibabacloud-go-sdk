// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrievalKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryText(v string) *RetrievalKnowledgeBaseResponseBody
	GetQueryText() *string
	SetRequestId(v string) *RetrievalKnowledgeBaseResponseBody
	GetRequestId() *string
	SetResultCount(v int32) *RetrievalKnowledgeBaseResponseBody
	GetResultCount() *int32
	SetResults(v []*RetrievalKnowledgeBaseResponseBodyResults) *RetrievalKnowledgeBaseResponseBody
	GetResults() []*RetrievalKnowledgeBaseResponseBodyResults
}

type RetrievalKnowledgeBaseResponseBody struct {
	// The query text.
	//
	// example:
	//
	// Financial report
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of results.
	//
	// example:
	//
	// 5
	ResultCount *int32 `json:"ResultCount,omitempty" xml:"ResultCount,omitempty"`
	// The search results.
	Results []*RetrievalKnowledgeBaseResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RetrievalKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrievalKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *RetrievalKnowledgeBaseResponseBody) GetQueryText() *string {
	return s.QueryText
}

func (s *RetrievalKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrievalKnowledgeBaseResponseBody) GetResultCount() *int32 {
	return s.ResultCount
}

func (s *RetrievalKnowledgeBaseResponseBody) GetResults() []*RetrievalKnowledgeBaseResponseBodyResults {
	return s.Results
}

func (s *RetrievalKnowledgeBaseResponseBody) SetQueryText(v string) *RetrievalKnowledgeBaseResponseBody {
	s.QueryText = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBody) SetRequestId(v string) *RetrievalKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBody) SetResultCount(v int32) *RetrievalKnowledgeBaseResponseBody {
	s.ResultCount = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBody) SetResults(v []*RetrievalKnowledgeBaseResponseBodyResults) *RetrievalKnowledgeBaseResponseBody {
	s.Results = v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RetrievalKnowledgeBaseResponseBodyResults struct {
	// The list of figure or table captions associated with the chunk.
	Captions []*string `json:"Captions,omitempty" xml:"Captions,omitempty" type:"Repeated"`
	// The list of Docling source document structured element references associated with the chunk. You can use these references to precisely locate original elements.
	DocItems []*string `json:"DocItems,omitempty" xml:"DocItems,omitempty" type:"Repeated"`
	// The unique ID of the file.
	//
	// example:
	//
	// 91b97b71-xxxx-xxxx-xxxx-33c6a6341cdc
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// 2024FinancialReport.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The chain of section headings that the chunk belongs to.
	Headings []*string `json:"Headings,omitempty" xml:"Headings,omitempty" type:"Repeated"`
	// The list of image resources referenced by the chunk.
	ImageResources []*RetrievalKnowledgeBaseResponseBodyResultsImageResources `json:"ImageResources,omitempty" xml:"ImageResources,omitempty" type:"Repeated"`
	// The metadata.
	//
	// example:
	//
	// {}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The list of page numbers that the chunk belongs to.
	PageNumbers []*int32 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty" type:"Repeated"`
	// The text content of the chunk.
	//
	// example:
	//
	// Financial report
	ShardContent *string `json:"ShardContent,omitempty" xml:"ShardContent,omitempty"`
	// The index of the chunk.
	//
	// example:
	//
	// 1
	ShardIndex *int32 `json:"ShardIndex,omitempty" xml:"ShardIndex,omitempty"`
	// The similarity score.
	//
	// example:
	//
	// 0.8
	SimilarityScore *float64 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty"`
}

func (s RetrievalKnowledgeBaseResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s RetrievalKnowledgeBaseResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetCaptions() []*string {
	return s.Captions
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetDocItems() []*string {
	return s.DocItems
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetFileId() *string {
	return s.FileId
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetFileName() *string {
	return s.FileName
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetHeadings() []*string {
	return s.Headings
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetImageResources() []*RetrievalKnowledgeBaseResponseBodyResultsImageResources {
	return s.ImageResources
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetMetadata() *string {
	return s.Metadata
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetPageNumbers() []*int32 {
	return s.PageNumbers
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetShardContent() *string {
	return s.ShardContent
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetShardIndex() *int32 {
	return s.ShardIndex
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetSimilarityScore() *float64 {
	return s.SimilarityScore
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetCaptions(v []*string) *RetrievalKnowledgeBaseResponseBodyResults {
	s.Captions = v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetDocItems(v []*string) *RetrievalKnowledgeBaseResponseBodyResults {
	s.DocItems = v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetFileId(v string) *RetrievalKnowledgeBaseResponseBodyResults {
	s.FileId = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetFileName(v string) *RetrievalKnowledgeBaseResponseBodyResults {
	s.FileName = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetHeadings(v []*string) *RetrievalKnowledgeBaseResponseBodyResults {
	s.Headings = v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetImageResources(v []*RetrievalKnowledgeBaseResponseBodyResultsImageResources) *RetrievalKnowledgeBaseResponseBodyResults {
	s.ImageResources = v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetMetadata(v string) *RetrievalKnowledgeBaseResponseBodyResults {
	s.Metadata = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetPageNumbers(v []*int32) *RetrievalKnowledgeBaseResponseBodyResults {
	s.PageNumbers = v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetShardContent(v string) *RetrievalKnowledgeBaseResponseBodyResults {
	s.ShardContent = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetShardIndex(v int32) *RetrievalKnowledgeBaseResponseBodyResults {
	s.ShardIndex = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) SetSimilarityScore(v float64) *RetrievalKnowledgeBaseResponseBodyResults {
	s.SimilarityScore = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) Validate() error {
	if s.ImageResources != nil {
		for _, item := range s.ImageResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RetrievalKnowledgeBaseResponseBodyResultsImageResources struct {
	// The index of the source document that the image belongs to, starting from 0.
	//
	// example:
	//
	// 0
	DocumentIndex *int32 `json:"DocumentIndex,omitempty" xml:"DocumentIndex,omitempty"`
	// The unique ID of the image resource.
	//
	// example:
	//
	// document-0/pictures/1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The element reference of the image in the Docling source document structure.
	//
	// example:
	//
	// #/pictures/1
	ItemRef *string `json:"ItemRef,omitempty" xml:"ItemRef,omitempty"`
	// The media type of the image resource.
	//
	// example:
	//
	// image/png
	MimeType *string `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
	// The OSS URI of the image resource.
	//
	// example:
	//
	// oss://my-bucket/results/my-space/doc-001/artifacts/image-1.png
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s RetrievalKnowledgeBaseResponseBodyResultsImageResources) String() string {
	return dara.Prettify(s)
}

func (s RetrievalKnowledgeBaseResponseBodyResultsImageResources) GoString() string {
	return s.String()
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) GetDocumentIndex() *int32 {
	return s.DocumentIndex
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) GetId() *string {
	return s.Id
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) GetItemRef() *string {
	return s.ItemRef
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) GetMimeType() *string {
	return s.MimeType
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) GetUri() *string {
	return s.Uri
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) SetDocumentIndex(v int32) *RetrievalKnowledgeBaseResponseBodyResultsImageResources {
	s.DocumentIndex = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) SetId(v string) *RetrievalKnowledgeBaseResponseBodyResultsImageResources {
	s.Id = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) SetItemRef(v string) *RetrievalKnowledgeBaseResponseBodyResultsImageResources {
	s.ItemRef = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) SetMimeType(v string) *RetrievalKnowledgeBaseResponseBodyResultsImageResources {
	s.MimeType = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) SetUri(v string) *RetrievalKnowledgeBaseResponseBodyResultsImageResources {
	s.Uri = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponseBodyResultsImageResources) Validate() error {
	return dara.Validate(s)
}
