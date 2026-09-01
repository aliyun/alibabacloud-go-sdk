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
	// example:
	//
	// 财报
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	ResultCount *int32                                       `json:"ResultCount,omitempty" xml:"ResultCount,omitempty"`
	Results     []*RetrievalKnowledgeBaseResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// example:
	//
	// 91b97b71-xxxx-xxxx-xxxx-33c6a6341cdc
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 2024财报.pdf
	FileName *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Headings []*string `json:"Headings,omitempty" xml:"Headings,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata    *string  `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	PageNumbers []*int32 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty" type:"Repeated"`
	// example:
	//
	// 财报
	ShardContent *string `json:"ShardContent,omitempty" xml:"ShardContent,omitempty"`
	// example:
	//
	// 1
	ShardIndex *int32 `json:"ShardIndex,omitempty" xml:"ShardIndex,omitempty"`
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

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetFileId() *string {
	return s.FileId
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetFileName() *string {
	return s.FileName
}

func (s *RetrievalKnowledgeBaseResponseBodyResults) GetHeadings() []*string {
	return s.Headings
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
	return dara.Validate(s)
}
