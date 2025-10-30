// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextEmbeddingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *TextEmbeddingResponseBody
	GetMessage() *string
	SetRequestId(v string) *TextEmbeddingResponseBody
	GetRequestId() *string
	SetResults(v *TextEmbeddingResponseBodyResults) *TextEmbeddingResponseBody
	GetResults() *TextEmbeddingResponseBodyResults
	SetStatus(v string) *TextEmbeddingResponseBody
	GetStatus() *string
	SetTextTokens(v int32) *TextEmbeddingResponseBody
	GetTextTokens() *int32
}

type TextEmbeddingResponseBody struct {
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *TextEmbeddingResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1000
	TextTokens *int32 `json:"TextTokens,omitempty" xml:"TextTokens,omitempty"`
}

func (s TextEmbeddingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TextEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *TextEmbeddingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TextEmbeddingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TextEmbeddingResponseBody) GetResults() *TextEmbeddingResponseBodyResults {
	return s.Results
}

func (s *TextEmbeddingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *TextEmbeddingResponseBody) GetTextTokens() *int32 {
	return s.TextTokens
}

func (s *TextEmbeddingResponseBody) SetMessage(v string) *TextEmbeddingResponseBody {
	s.Message = &v
	return s
}

func (s *TextEmbeddingResponseBody) SetRequestId(v string) *TextEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *TextEmbeddingResponseBody) SetResults(v *TextEmbeddingResponseBodyResults) *TextEmbeddingResponseBody {
	s.Results = v
	return s
}

func (s *TextEmbeddingResponseBody) SetStatus(v string) *TextEmbeddingResponseBody {
	s.Status = &v
	return s
}

func (s *TextEmbeddingResponseBody) SetTextTokens(v int32) *TextEmbeddingResponseBody {
	s.TextTokens = &v
	return s
}

func (s *TextEmbeddingResponseBody) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextEmbeddingResponseBodyResults struct {
	Results []*TextEmbeddingResponseBodyResultsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s TextEmbeddingResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s TextEmbeddingResponseBodyResults) GoString() string {
	return s.String()
}

func (s *TextEmbeddingResponseBodyResults) GetResults() []*TextEmbeddingResponseBodyResultsResults {
	return s.Results
}

func (s *TextEmbeddingResponseBodyResults) SetResults(v []*TextEmbeddingResponseBodyResultsResults) *TextEmbeddingResponseBodyResults {
	s.Results = v
	return s
}

func (s *TextEmbeddingResponseBodyResults) Validate() error {
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

type TextEmbeddingResponseBodyResultsResults struct {
	Embedding *TextEmbeddingResponseBodyResultsResultsEmbedding `json:"Embedding,omitempty" xml:"Embedding,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s TextEmbeddingResponseBodyResultsResults) String() string {
	return dara.Prettify(s)
}

func (s TextEmbeddingResponseBodyResultsResults) GoString() string {
	return s.String()
}

func (s *TextEmbeddingResponseBodyResultsResults) GetEmbedding() *TextEmbeddingResponseBodyResultsResultsEmbedding {
	return s.Embedding
}

func (s *TextEmbeddingResponseBodyResultsResults) GetIndex() *int32 {
	return s.Index
}

func (s *TextEmbeddingResponseBodyResultsResults) SetEmbedding(v *TextEmbeddingResponseBodyResultsResultsEmbedding) *TextEmbeddingResponseBodyResultsResults {
	s.Embedding = v
	return s
}

func (s *TextEmbeddingResponseBodyResultsResults) SetIndex(v int32) *TextEmbeddingResponseBodyResultsResults {
	s.Index = &v
	return s
}

func (s *TextEmbeddingResponseBodyResultsResults) Validate() error {
	if s.Embedding != nil {
		if err := s.Embedding.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextEmbeddingResponseBodyResultsResultsEmbedding struct {
	Embedding []*float64 `json:"Embedding,omitempty" xml:"Embedding,omitempty" type:"Repeated"`
}

func (s TextEmbeddingResponseBodyResultsResultsEmbedding) String() string {
	return dara.Prettify(s)
}

func (s TextEmbeddingResponseBodyResultsResultsEmbedding) GoString() string {
	return s.String()
}

func (s *TextEmbeddingResponseBodyResultsResultsEmbedding) GetEmbedding() []*float64 {
	return s.Embedding
}

func (s *TextEmbeddingResponseBodyResultsResultsEmbedding) SetEmbedding(v []*float64) *TextEmbeddingResponseBodyResultsResultsEmbedding {
	s.Embedding = v
	return s
}

func (s *TextEmbeddingResponseBodyResultsResultsEmbedding) Validate() error {
	return dara.Validate(s)
}
