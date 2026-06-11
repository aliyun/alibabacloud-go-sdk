// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RetrieveKnowledgeBaseResponseBodyData) *RetrieveKnowledgeBaseResponseBody
	GetData() *RetrieveKnowledgeBaseResponseBodyData
	SetErrorCode(v string) *RetrieveKnowledgeBaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RetrieveKnowledgeBaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RetrieveKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetrieveKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type RetrieveKnowledgeBaseResponseBody struct {
	// The response payload.
	Data *RetrieveKnowledgeBaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code for a failed request.
	//
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message for a failed request.
	//
	// example:
	//
	// Resource not found kb-***
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique request ID. Use this ID to troubleshoot errors.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetrieveKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseResponseBody) GetData() *RetrieveKnowledgeBaseResponseBodyData {
	return s.Data
}

func (s *RetrieveKnowledgeBaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RetrieveKnowledgeBaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RetrieveKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrieveKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetrieveKnowledgeBaseResponseBody) SetData(v *RetrieveKnowledgeBaseResponseBodyData) *RetrieveKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBody) SetErrorCode(v string) *RetrieveKnowledgeBaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBody) SetErrorMessage(v string) *RetrieveKnowledgeBaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBody) SetRequestId(v string) *RetrieveKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBody) SetSuccess(v bool) *RetrieveKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetrieveKnowledgeBaseResponseBodyData struct {
	// Matched items from the windowing process.
	Matches []*RetrieveKnowledgeBaseResponseBodyDataMatches `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Repeated"`
	// The retrieved results.
	Results []*RetrieveKnowledgeBaseResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RetrieveKnowledgeBaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseResponseBodyData) GetMatches() []*RetrieveKnowledgeBaseResponseBodyDataMatches {
	return s.Matches
}

func (s *RetrieveKnowledgeBaseResponseBodyData) GetResults() []*RetrieveKnowledgeBaseResponseBodyDataResults {
	return s.Results
}

func (s *RetrieveKnowledgeBaseResponseBodyData) SetMatches(v []*RetrieveKnowledgeBaseResponseBodyDataMatches) *RetrieveKnowledgeBaseResponseBodyData {
	s.Matches = v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyData) SetResults(v []*RetrieveKnowledgeBaseResponseBodyDataResults) *RetrieveKnowledgeBaseResponseBodyData {
	s.Results = v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyData) Validate() error {
	if s.Matches != nil {
		for _, item := range s.Matches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type RetrieveKnowledgeBaseResponseBodyDataMatches struct {
	// The text content.
	//
	// example:
	//
	// GraphRAG (Graph Retrieval-Augmented Generation) is an advanced AI framework that combines Knowledge Graphs with Large Language Models (LLMs) to improve the accuracy, context, and reasoning capabilities of information retrieval and generation.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The name of the source document.
	//
	// example:
	//
	// 1.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The unique ID of the vector.
	//
	// example:
	//
	// 3b35c95-57f3-442f-9220-xxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metadata generated by the document loader.
	//
	// example:
	//
	// {"page_pos": 1}
	LoaderMetadata *string `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	// A key-value map of the metadata.
	//
	// example:
	//
	// {"title":"test"}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
}

func (s RetrieveKnowledgeBaseResponseBodyDataMatches) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseResponseBodyDataMatches) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) GetContent() *string {
	return s.Content
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) GetFileName() *string {
	return s.FileName
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) GetId() *string {
	return s.Id
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) GetLoaderMetadata() *string {
	return s.LoaderMetadata
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) SetContent(v string) *RetrieveKnowledgeBaseResponseBodyDataMatches {
	s.Content = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) SetFileName(v string) *RetrieveKnowledgeBaseResponseBodyDataMatches {
	s.FileName = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) SetId(v string) *RetrieveKnowledgeBaseResponseBodyDataMatches {
	s.Id = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) SetLoaderMetadata(v string) *RetrieveKnowledgeBaseResponseBodyDataMatches {
	s.LoaderMetadata = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) SetMetadata(v map[string]interface{}) *RetrieveKnowledgeBaseResponseBodyDataMatches {
	s.Metadata = v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataMatches) Validate() error {
	return dara.Validate(s)
}

type RetrieveKnowledgeBaseResponseBodyDataResults struct {
	// The content of the retrieved text chunk.
	//
	// example:
	//
	// GraphRAG (Graph Retrieval-Augmented Generation) is an advanced AI framework that combines Knowledge Graphs with Large Language Models (LLMs) to improve the accuracy, context, and reasoning capabilities of information retrieval and generation.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The name of the source document.
	//
	// example:
	//
	// 1.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The unique ID of the vector.
	//
	// example:
	//
	// 3b35c95-57f3-442f-9220-xxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metadata generated by the document loader.
	//
	// example:
	//
	// {"page_pos": 1}
	LoaderMetadata *string `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	// A key-value map of the metadata.
	//
	// example:
	//
	// {"title":"test"}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The rerank score.
	//
	// example:
	//
	// 6.2345
	RerankScore *float64 `json:"RerankScore,omitempty" xml:"RerankScore,omitempty"`
	// The method used for retrieval. Valid values: `1` (vector retrieval), `2` (full-text search), and `3` (hybrid search).
	//
	// example:
	//
	// 1
	RetrievalSource *int32 `json:"RetrievalSource,omitempty" xml:"RetrievalSource,omitempty"`
	// The similarity score of the result. This score depends on the distance metric (`l2`, `ip`, or `cosine`) specified when the index was created.
	//
	// example:
	//
	// 0.81
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The vector data as an array of floating-point numbers.
	Vector []*float64 `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Repeated"`
}

func (s RetrieveKnowledgeBaseResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetContent() *string {
	return s.Content
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetFileName() *string {
	return s.FileName
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetId() *string {
	return s.Id
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetLoaderMetadata() *string {
	return s.LoaderMetadata
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetRerankScore() *float64 {
	return s.RerankScore
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetRetrievalSource() *int32 {
	return s.RetrievalSource
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetScore() *float64 {
	return s.Score
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) GetVector() []*float64 {
	return s.Vector
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetContent(v string) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.Content = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetFileName(v string) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.FileName = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetId(v string) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.Id = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetLoaderMetadata(v string) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.LoaderMetadata = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetMetadata(v map[string]interface{}) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.Metadata = v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetRerankScore(v float64) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.RerankScore = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetRetrievalSource(v int32) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.RetrievalSource = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetScore(v float64) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.Score = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) SetVector(v []*float64) *RetrieveKnowledgeBaseResponseBodyDataResults {
	s.Vector = v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
