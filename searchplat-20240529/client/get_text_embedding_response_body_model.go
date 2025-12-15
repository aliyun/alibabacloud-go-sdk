// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextEmbeddingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetTextEmbeddingResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetTextEmbeddingResponseBody
	GetRequestId() *string
	SetResult(v *GetTextEmbeddingResponseBodyResult) *GetTextEmbeddingResponseBody
	GetResult() *GetTextEmbeddingResponseBodyResult
	SetUsage(v *GetTextEmbeddingResponseBodyUsage) *GetTextEmbeddingResponseBody
	GetUsage() *GetTextEmbeddingResponseBodyUsage
}

type GetTextEmbeddingResponseBody struct {
	Latency   *int32                              `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                             `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetTextEmbeddingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetTextEmbeddingResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTextEmbeddingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTextEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetTextEmbeddingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTextEmbeddingResponseBody) GetResult() *GetTextEmbeddingResponseBodyResult {
	return s.Result
}

func (s *GetTextEmbeddingResponseBody) GetUsage() *GetTextEmbeddingResponseBodyUsage {
	return s.Usage
}

func (s *GetTextEmbeddingResponseBody) SetLatency(v int32) *GetTextEmbeddingResponseBody {
	s.Latency = &v
	return s
}

func (s *GetTextEmbeddingResponseBody) SetRequestId(v string) *GetTextEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextEmbeddingResponseBody) SetResult(v *GetTextEmbeddingResponseBodyResult) *GetTextEmbeddingResponseBody {
	s.Result = v
	return s
}

func (s *GetTextEmbeddingResponseBody) SetUsage(v *GetTextEmbeddingResponseBodyUsage) *GetTextEmbeddingResponseBody {
	s.Usage = v
	return s
}

func (s *GetTextEmbeddingResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTextEmbeddingResponseBodyResult struct {
	Embeddings []*GetTextEmbeddingResponseBodyResultEmbeddings `json:"embeddings,omitempty" xml:"embeddings,omitempty" type:"Repeated"`
}

func (s GetTextEmbeddingResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetTextEmbeddingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponseBodyResult) GetEmbeddings() []*GetTextEmbeddingResponseBodyResultEmbeddings {
	return s.Embeddings
}

func (s *GetTextEmbeddingResponseBodyResult) SetEmbeddings(v []*GetTextEmbeddingResponseBodyResultEmbeddings) *GetTextEmbeddingResponseBodyResult {
	s.Embeddings = v
	return s
}

func (s *GetTextEmbeddingResponseBodyResult) Validate() error {
	if s.Embeddings != nil {
		for _, item := range s.Embeddings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTextEmbeddingResponseBodyResultEmbeddings struct {
	Embedding []*float64 `json:"embedding,omitempty" xml:"embedding,omitempty" type:"Repeated"`
	Index     *int32     `json:"index,omitempty" xml:"index,omitempty"`
}

func (s GetTextEmbeddingResponseBodyResultEmbeddings) String() string {
	return dara.Prettify(s)
}

func (s GetTextEmbeddingResponseBodyResultEmbeddings) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponseBodyResultEmbeddings) GetEmbedding() []*float64 {
	return s.Embedding
}

func (s *GetTextEmbeddingResponseBodyResultEmbeddings) GetIndex() *int32 {
	return s.Index
}

func (s *GetTextEmbeddingResponseBodyResultEmbeddings) SetEmbedding(v []*float64) *GetTextEmbeddingResponseBodyResultEmbeddings {
	s.Embedding = v
	return s
}

func (s *GetTextEmbeddingResponseBodyResultEmbeddings) SetIndex(v int32) *GetTextEmbeddingResponseBodyResultEmbeddings {
	s.Index = &v
	return s
}

func (s *GetTextEmbeddingResponseBodyResultEmbeddings) Validate() error {
	return dara.Validate(s)
}

type GetTextEmbeddingResponseBodyUsage struct {
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetTextEmbeddingResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetTextEmbeddingResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponseBodyUsage) GetTokenCount() *int64 {
	return s.TokenCount
}

func (s *GetTextEmbeddingResponseBodyUsage) SetTokenCount(v int64) *GetTextEmbeddingResponseBodyUsage {
	s.TokenCount = &v
	return s
}

func (s *GetTextEmbeddingResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
