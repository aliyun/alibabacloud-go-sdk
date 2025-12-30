// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalSearchTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalSearchTaskShrinkRequest
	GetDBClusterId() *string
	SetDatasetIdsShrink(v string) *CreateMultimodalSearchTaskShrinkRequest
	GetDatasetIdsShrink() *string
	SetEmbeddingModel(v string) *CreateMultimodalSearchTaskShrinkRequest
	GetEmbeddingModel() *string
	SetQuery(v string) *CreateMultimodalSearchTaskShrinkRequest
	GetQuery() *string
	SetSearchModel(v string) *CreateMultimodalSearchTaskShrinkRequest
	GetSearchModel() *string
	SetTopK(v int32) *CreateMultimodalSearchTaskShrinkRequest
	GetTopK() *int32
}

type CreateMultimodalSearchTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// ["ds-1", "ds-2"]
	DatasetIdsShrink *string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty"`
	// example:
	//
	// Multimodal-Embedding
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// example:
	//
	// 指示牌
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// Multimodal-Search
	SearchModel *string `json:"SearchModel,omitempty" xml:"SearchModel,omitempty"`
	// example:
	//
	// 20
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s CreateMultimodalSearchTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalSearchTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalSearchTaskShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalSearchTaskShrinkRequest) GetDatasetIdsShrink() *string {
	return s.DatasetIdsShrink
}

func (s *CreateMultimodalSearchTaskShrinkRequest) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateMultimodalSearchTaskShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *CreateMultimodalSearchTaskShrinkRequest) GetSearchModel() *string {
	return s.SearchModel
}

func (s *CreateMultimodalSearchTaskShrinkRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *CreateMultimodalSearchTaskShrinkRequest) SetDBClusterId(v string) *CreateMultimodalSearchTaskShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalSearchTaskShrinkRequest) SetDatasetIdsShrink(v string) *CreateMultimodalSearchTaskShrinkRequest {
	s.DatasetIdsShrink = &v
	return s
}

func (s *CreateMultimodalSearchTaskShrinkRequest) SetEmbeddingModel(v string) *CreateMultimodalSearchTaskShrinkRequest {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateMultimodalSearchTaskShrinkRequest) SetQuery(v string) *CreateMultimodalSearchTaskShrinkRequest {
	s.Query = &v
	return s
}

func (s *CreateMultimodalSearchTaskShrinkRequest) SetSearchModel(v string) *CreateMultimodalSearchTaskShrinkRequest {
	s.SearchModel = &v
	return s
}

func (s *CreateMultimodalSearchTaskShrinkRequest) SetTopK(v int32) *CreateMultimodalSearchTaskShrinkRequest {
	s.TopK = &v
	return s
}

func (s *CreateMultimodalSearchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
