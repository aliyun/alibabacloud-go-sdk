// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiModalEmbeddingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetMultiModalEmbeddingResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetMultiModalEmbeddingResponseBody
	GetRequestId() *string
	SetResult(v *GetMultiModalEmbeddingResponseBodyResult) *GetMultiModalEmbeddingResponseBody
	GetResult() *GetMultiModalEmbeddingResponseBodyResult
	SetUsage(v *GetMultiModalEmbeddingResponseBodyUsage) *GetMultiModalEmbeddingResponseBody
	GetUsage() *GetMultiModalEmbeddingResponseBodyUsage
}

type GetMultiModalEmbeddingResponseBody struct {
	Latency   *int32                                    `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                   `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetMultiModalEmbeddingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetMultiModalEmbeddingResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetMultiModalEmbeddingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiModalEmbeddingResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetMultiModalEmbeddingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiModalEmbeddingResponseBody) GetResult() *GetMultiModalEmbeddingResponseBodyResult {
	return s.Result
}

func (s *GetMultiModalEmbeddingResponseBody) GetUsage() *GetMultiModalEmbeddingResponseBodyUsage {
	return s.Usage
}

func (s *GetMultiModalEmbeddingResponseBody) SetLatency(v int32) *GetMultiModalEmbeddingResponseBody {
	s.Latency = &v
	return s
}

func (s *GetMultiModalEmbeddingResponseBody) SetRequestId(v string) *GetMultiModalEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiModalEmbeddingResponseBody) SetResult(v *GetMultiModalEmbeddingResponseBodyResult) *GetMultiModalEmbeddingResponseBody {
	s.Result = v
	return s
}

func (s *GetMultiModalEmbeddingResponseBody) SetUsage(v *GetMultiModalEmbeddingResponseBodyUsage) *GetMultiModalEmbeddingResponseBody {
	s.Usage = v
	return s
}

func (s *GetMultiModalEmbeddingResponseBody) Validate() error {
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

type GetMultiModalEmbeddingResponseBodyResult struct {
	Embeddings []*GetMultiModalEmbeddingResponseBodyResultEmbeddings `json:"embeddings,omitempty" xml:"embeddings,omitempty" type:"Repeated"`
}

func (s GetMultiModalEmbeddingResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalEmbeddingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMultiModalEmbeddingResponseBodyResult) GetEmbeddings() []*GetMultiModalEmbeddingResponseBodyResultEmbeddings {
	return s.Embeddings
}

func (s *GetMultiModalEmbeddingResponseBodyResult) SetEmbeddings(v []*GetMultiModalEmbeddingResponseBodyResultEmbeddings) *GetMultiModalEmbeddingResponseBodyResult {
	s.Embeddings = v
	return s
}

func (s *GetMultiModalEmbeddingResponseBodyResult) Validate() error {
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

type GetMultiModalEmbeddingResponseBodyResultEmbeddings struct {
	Embedding []*float64 `json:"embedding,omitempty" xml:"embedding,omitempty" type:"Repeated"`
	Index     *int64     `json:"index,omitempty" xml:"index,omitempty"`
}

func (s GetMultiModalEmbeddingResponseBodyResultEmbeddings) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalEmbeddingResponseBodyResultEmbeddings) GoString() string {
	return s.String()
}

func (s *GetMultiModalEmbeddingResponseBodyResultEmbeddings) GetEmbedding() []*float64 {
	return s.Embedding
}

func (s *GetMultiModalEmbeddingResponseBodyResultEmbeddings) GetIndex() *int64 {
	return s.Index
}

func (s *GetMultiModalEmbeddingResponseBodyResultEmbeddings) SetEmbedding(v []*float64) *GetMultiModalEmbeddingResponseBodyResultEmbeddings {
	s.Embedding = v
	return s
}

func (s *GetMultiModalEmbeddingResponseBodyResultEmbeddings) SetIndex(v int64) *GetMultiModalEmbeddingResponseBodyResultEmbeddings {
	s.Index = &v
	return s
}

func (s *GetMultiModalEmbeddingResponseBodyResultEmbeddings) Validate() error {
	return dara.Validate(s)
}

type GetMultiModalEmbeddingResponseBodyUsage struct {
	Image      *int64 `json:"image,omitempty" xml:"image,omitempty"`
	ImageToken *int64 `json:"image_token,omitempty" xml:"image_token,omitempty"`
	TextToken  *int64 `json:"text_token,omitempty" xml:"text_token,omitempty"`
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetMultiModalEmbeddingResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalEmbeddingResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) GetImage() *int64 {
	return s.Image
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) GetImageToken() *int64 {
	return s.ImageToken
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) GetTextToken() *int64 {
	return s.TextToken
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) GetTokenCount() *int64 {
	return s.TokenCount
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) SetImage(v int64) *GetMultiModalEmbeddingResponseBodyUsage {
	s.Image = &v
	return s
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) SetImageToken(v int64) *GetMultiModalEmbeddingResponseBodyUsage {
	s.ImageToken = &v
	return s
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) SetTextToken(v int64) *GetMultiModalEmbeddingResponseBodyUsage {
	s.TextToken = &v
	return s
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) SetTokenCount(v int64) *GetMultiModalEmbeddingResponseBodyUsage {
	s.TokenCount = &v
	return s
}

func (s *GetMultiModalEmbeddingResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
