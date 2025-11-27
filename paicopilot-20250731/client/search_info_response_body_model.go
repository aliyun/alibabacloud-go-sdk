// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseSearchResults(v []*SearchInfoResponseBodyKnowledgeBaseSearchResults) *SearchInfoResponseBody
	GetKnowledgeBaseSearchResults() []*SearchInfoResponseBodyKnowledgeBaseSearchResults
	SetRequestId(v string) *SearchInfoResponseBody
	GetRequestId() *string
	SetWebSearchResults(v []*SearchInfoResponseBodyWebSearchResults) *SearchInfoResponseBody
	GetWebSearchResults() []*SearchInfoResponseBodyWebSearchResults
}

type SearchInfoResponseBody struct {
	KnowledgeBaseSearchResults []*SearchInfoResponseBodyKnowledgeBaseSearchResults `json:"KnowledgeBaseSearchResults,omitempty" xml:"KnowledgeBaseSearchResults,omitempty" type:"Repeated"`
	// example:
	//
	// 44553E9A-******-37ADC33FE2
	RequestId        *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebSearchResults []*SearchInfoResponseBodyWebSearchResults `json:"WebSearchResults,omitempty" xml:"WebSearchResults,omitempty" type:"Repeated"`
}

func (s SearchInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SearchInfoResponseBody) GetKnowledgeBaseSearchResults() []*SearchInfoResponseBodyKnowledgeBaseSearchResults {
	return s.KnowledgeBaseSearchResults
}

func (s *SearchInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchInfoResponseBody) GetWebSearchResults() []*SearchInfoResponseBodyWebSearchResults {
	return s.WebSearchResults
}

func (s *SearchInfoResponseBody) SetKnowledgeBaseSearchResults(v []*SearchInfoResponseBodyKnowledgeBaseSearchResults) *SearchInfoResponseBody {
	s.KnowledgeBaseSearchResults = v
	return s
}

func (s *SearchInfoResponseBody) SetRequestId(v string) *SearchInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchInfoResponseBody) SetWebSearchResults(v []*SearchInfoResponseBodyWebSearchResults) *SearchInfoResponseBody {
	s.WebSearchResults = v
	return s
}

func (s *SearchInfoResponseBody) Validate() error {
	if s.KnowledgeBaseSearchResults != nil {
		for _, item := range s.KnowledgeBaseSearchResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebSearchResults != nil {
		for _, item := range s.WebSearchResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchInfoResponseBodyKnowledgeBaseSearchResults struct {
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// DSW is *****	- in pai.
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	// example:
	//
	// 0.6215165367
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchInfoResponseBodyKnowledgeBaseSearchResults) String() string {
	return dara.Prettify(s)
}

func (s SearchInfoResponseBodyKnowledgeBaseSearchResults) GoString() string {
	return s.String()
}

func (s *SearchInfoResponseBodyKnowledgeBaseSearchResults) GetIndex() *string {
	return s.Index
}

func (s *SearchInfoResponseBodyKnowledgeBaseSearchResults) GetResultContent() *string {
	return s.ResultContent
}

func (s *SearchInfoResponseBodyKnowledgeBaseSearchResults) GetScore() *float64 {
	return s.Score
}

func (s *SearchInfoResponseBodyKnowledgeBaseSearchResults) SetIndex(v string) *SearchInfoResponseBodyKnowledgeBaseSearchResults {
	s.Index = &v
	return s
}

func (s *SearchInfoResponseBodyKnowledgeBaseSearchResults) SetResultContent(v string) *SearchInfoResponseBodyKnowledgeBaseSearchResults {
	s.ResultContent = &v
	return s
}

func (s *SearchInfoResponseBodyKnowledgeBaseSearchResults) SetScore(v float64) *SearchInfoResponseBodyKnowledgeBaseSearchResults {
	s.Score = &v
	return s
}

func (s *SearchInfoResponseBodyKnowledgeBaseSearchResults) Validate() error {
	return dara.Validate(s)
}

type SearchInfoResponseBodyWebSearchResults struct {
	// example:
	//
	// 2
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// Qwen3-Coder*****is good
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	// example:
	//
	// 0.5215170567
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// https://help.aliyun.com/zh/pai/use-cases/dsw-use-qwen3-coder-to-assist-in-code-development
	SourceLink *string `json:"SourceLink,omitempty" xml:"SourceLink,omitempty"`
}

func (s SearchInfoResponseBodyWebSearchResults) String() string {
	return dara.Prettify(s)
}

func (s SearchInfoResponseBodyWebSearchResults) GoString() string {
	return s.String()
}

func (s *SearchInfoResponseBodyWebSearchResults) GetIndex() *string {
	return s.Index
}

func (s *SearchInfoResponseBodyWebSearchResults) GetResultContent() *string {
	return s.ResultContent
}

func (s *SearchInfoResponseBodyWebSearchResults) GetScore() *float64 {
	return s.Score
}

func (s *SearchInfoResponseBodyWebSearchResults) GetSourceLink() *string {
	return s.SourceLink
}

func (s *SearchInfoResponseBodyWebSearchResults) SetIndex(v string) *SearchInfoResponseBodyWebSearchResults {
	s.Index = &v
	return s
}

func (s *SearchInfoResponseBodyWebSearchResults) SetResultContent(v string) *SearchInfoResponseBodyWebSearchResults {
	s.ResultContent = &v
	return s
}

func (s *SearchInfoResponseBodyWebSearchResults) SetScore(v float64) *SearchInfoResponseBodyWebSearchResults {
	s.Score = &v
	return s
}

func (s *SearchInfoResponseBodyWebSearchResults) SetSourceLink(v string) *SearchInfoResponseBodyWebSearchResults {
	s.SourceLink = &v
	return s
}

func (s *SearchInfoResponseBodyWebSearchResults) Validate() error {
	return dara.Validate(s)
}
