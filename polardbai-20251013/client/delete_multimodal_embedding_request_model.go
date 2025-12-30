// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalEmbeddingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteMultimodalEmbeddingRequest
	GetDBClusterId() *string
	SetEmbedding(v string) *DeleteMultimodalEmbeddingRequest
	GetEmbedding() *string
}

type DeleteMultimodalEmbeddingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polar4ai_multimodal_embedding_****
	Embedding *string `json:"Embedding,omitempty" xml:"Embedding,omitempty"`
}

func (s DeleteMultimodalEmbeddingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalEmbeddingRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteMultimodalEmbeddingRequest) GetEmbedding() *string {
	return s.Embedding
}

func (s *DeleteMultimodalEmbeddingRequest) SetDBClusterId(v string) *DeleteMultimodalEmbeddingRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteMultimodalEmbeddingRequest) SetEmbedding(v string) *DeleteMultimodalEmbeddingRequest {
	s.Embedding = &v
	return s
}

func (s *DeleteMultimodalEmbeddingRequest) Validate() error {
	return dara.Validate(s)
}
