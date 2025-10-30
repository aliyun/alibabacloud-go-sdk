// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *RerankResponseBody
	GetMessage() *string
	SetRequestId(v string) *RerankResponseBody
	GetRequestId() *string
	SetResults(v *RerankResponseBodyResults) *RerankResponseBody
	GetResults() *RerankResponseBodyResults
	SetStatus(v string) *RerankResponseBody
	GetStatus() *string
	SetTokens(v int32) *RerankResponseBody
	GetTokens() *int32
}

type RerankResponseBody struct {
	// Detailed information returned by the interface.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Rerank results.
	Results *RerankResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// API execution status, value description:
	//
	// - **success**: Execution succeeded.
	//
	// - **fail**: Execution failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Number of consumed tokens.
	//
	// example:
	//
	// 100
	Tokens *int32 `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
}

func (s RerankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RerankResponseBody) GoString() string {
	return s.String()
}

func (s *RerankResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RerankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RerankResponseBody) GetResults() *RerankResponseBodyResults {
	return s.Results
}

func (s *RerankResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RerankResponseBody) GetTokens() *int32 {
	return s.Tokens
}

func (s *RerankResponseBody) SetMessage(v string) *RerankResponseBody {
	s.Message = &v
	return s
}

func (s *RerankResponseBody) SetRequestId(v string) *RerankResponseBody {
	s.RequestId = &v
	return s
}

func (s *RerankResponseBody) SetResults(v *RerankResponseBodyResults) *RerankResponseBody {
	s.Results = v
	return s
}

func (s *RerankResponseBody) SetStatus(v string) *RerankResponseBody {
	s.Status = &v
	return s
}

func (s *RerankResponseBody) SetTokens(v int32) *RerankResponseBody {
	s.Tokens = &v
	return s
}

func (s *RerankResponseBody) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RerankResponseBodyResults struct {
	Results []*RerankResponseBodyResultsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RerankResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s RerankResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RerankResponseBodyResults) GetResults() []*RerankResponseBodyResultsResults {
	return s.Results
}

func (s *RerankResponseBodyResults) SetResults(v []*RerankResponseBodyResultsResults) *RerankResponseBodyResults {
	s.Results = v
	return s
}

func (s *RerankResponseBodyResults) Validate() error {
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

type RerankResponseBodyResultsResults struct {
	// Re-ordered document information.
	//
	// example:
	//
	// ADBPG is the OLAP database of Alibaba Cloud.
	Document *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// Index of this document in the request parameter Documents, starting from 0.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Rerank similarity score.
	//
	// example:
	//
	// 2.31412
	RelevanceScore *float32 `json:"RelevanceScore,omitempty" xml:"RelevanceScore,omitempty"`
}

func (s RerankResponseBodyResultsResults) String() string {
	return dara.Prettify(s)
}

func (s RerankResponseBodyResultsResults) GoString() string {
	return s.String()
}

func (s *RerankResponseBodyResultsResults) GetDocument() *string {
	return s.Document
}

func (s *RerankResponseBodyResultsResults) GetIndex() *int32 {
	return s.Index
}

func (s *RerankResponseBodyResultsResults) GetRelevanceScore() *float32 {
	return s.RelevanceScore
}

func (s *RerankResponseBodyResultsResults) SetDocument(v string) *RerankResponseBodyResultsResults {
	s.Document = &v
	return s
}

func (s *RerankResponseBodyResultsResults) SetIndex(v int32) *RerankResponseBodyResultsResults {
	s.Index = &v
	return s
}

func (s *RerankResponseBodyResultsResults) SetRelevanceScore(v float32) *RerankResponseBodyResultsResults {
	s.RelevanceScore = &v
	return s
}

func (s *RerankResponseBodyResultsResults) Validate() error {
	return dara.Validate(s)
}
