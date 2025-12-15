// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextSparseEmbeddingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetTextSparseEmbeddingResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetTextSparseEmbeddingResponseBody
	GetRequestId() *string
	SetResult(v *GetTextSparseEmbeddingResponseBodyResult) *GetTextSparseEmbeddingResponseBody
	GetResult() *GetTextSparseEmbeddingResponseBodyResult
	SetUsage(v *GetTextSparseEmbeddingResponseBodyUsage) *GetTextSparseEmbeddingResponseBody
	GetUsage() *GetTextSparseEmbeddingResponseBodyUsage
}

type GetTextSparseEmbeddingResponseBody struct {
	Latency   *int32                                    `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                   `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetTextSparseEmbeddingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetTextSparseEmbeddingResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTextSparseEmbeddingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetTextSparseEmbeddingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTextSparseEmbeddingResponseBody) GetResult() *GetTextSparseEmbeddingResponseBodyResult {
	return s.Result
}

func (s *GetTextSparseEmbeddingResponseBody) GetUsage() *GetTextSparseEmbeddingResponseBodyUsage {
	return s.Usage
}

func (s *GetTextSparseEmbeddingResponseBody) SetLatency(v int32) *GetTextSparseEmbeddingResponseBody {
	s.Latency = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBody) SetRequestId(v string) *GetTextSparseEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBody) SetResult(v *GetTextSparseEmbeddingResponseBodyResult) *GetTextSparseEmbeddingResponseBody {
	s.Result = v
	return s
}

func (s *GetTextSparseEmbeddingResponseBody) SetUsage(v *GetTextSparseEmbeddingResponseBodyUsage) *GetTextSparseEmbeddingResponseBody {
	s.Usage = v
	return s
}

func (s *GetTextSparseEmbeddingResponseBody) Validate() error {
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

type GetTextSparseEmbeddingResponseBodyResult struct {
	SparseEmbeddings []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings `json:"sparse_embeddings,omitempty" xml:"sparse_embeddings,omitempty" type:"Repeated"`
}

func (s GetTextSparseEmbeddingResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBodyResult) GetSparseEmbeddings() []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings {
	return s.SparseEmbeddings
}

func (s *GetTextSparseEmbeddingResponseBodyResult) SetSparseEmbeddings(v []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) *GetTextSparseEmbeddingResponseBodyResult {
	s.SparseEmbeddings = v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResult) Validate() error {
	if s.SparseEmbeddings != nil {
		for _, item := range s.SparseEmbeddings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings struct {
	Embedding []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding `json:"embedding,omitempty" xml:"embedding,omitempty" type:"Repeated"`
	Index     *int32                                                               `json:"index,omitempty" xml:"index,omitempty"`
}

func (s GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) String() string {
	return dara.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) GetEmbedding() []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding {
	return s.Embedding
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) GetIndex() *int32 {
	return s.Index
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) SetEmbedding(v []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings {
	s.Embedding = v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) SetIndex(v int32) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings {
	s.Index = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) Validate() error {
	if s.Embedding != nil {
		for _, item := range s.Embedding {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding struct {
	Token   *string  `json:"token,omitempty" xml:"token,omitempty"`
	TokenId *int32   `json:"token_id,omitempty" xml:"token_id,omitempty"`
	Weight  *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) String() string {
	return dara.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) GetToken() *string {
	return s.Token
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) GetTokenId() *int32 {
	return s.TokenId
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) GetWeight() *float32 {
	return s.Weight
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) SetToken(v string) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding {
	s.Token = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) SetTokenId(v int32) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding {
	s.TokenId = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) SetWeight(v float32) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding {
	s.Weight = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) Validate() error {
	return dara.Validate(s)
}

type GetTextSparseEmbeddingResponseBodyUsage struct {
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetTextSparseEmbeddingResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBodyUsage) GetTokenCount() *int64 {
	return s.TokenCount
}

func (s *GetTextSparseEmbeddingResponseBodyUsage) SetTokenCount(v int64) *GetTextSparseEmbeddingResponseBodyUsage {
	s.TokenCount = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
