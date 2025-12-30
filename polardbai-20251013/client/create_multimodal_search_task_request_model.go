// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalSearchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalSearchTaskRequest
	GetDBClusterId() *string
	SetDatasetIds(v []*string) *CreateMultimodalSearchTaskRequest
	GetDatasetIds() []*string
	SetEmbeddingModel(v string) *CreateMultimodalSearchTaskRequest
	GetEmbeddingModel() *string
	SetQuery(v string) *CreateMultimodalSearchTaskRequest
	GetQuery() *string
	SetSearchModel(v string) *CreateMultimodalSearchTaskRequest
	GetSearchModel() *string
	SetTopK(v int32) *CreateMultimodalSearchTaskRequest
	GetTopK() *int32
}

type CreateMultimodalSearchTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// ["ds-1", "ds-2"]
	DatasetIds []*string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty" type:"Repeated"`
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

func (s CreateMultimodalSearchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalSearchTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalSearchTaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalSearchTaskRequest) GetDatasetIds() []*string {
	return s.DatasetIds
}

func (s *CreateMultimodalSearchTaskRequest) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateMultimodalSearchTaskRequest) GetQuery() *string {
	return s.Query
}

func (s *CreateMultimodalSearchTaskRequest) GetSearchModel() *string {
	return s.SearchModel
}

func (s *CreateMultimodalSearchTaskRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *CreateMultimodalSearchTaskRequest) SetDBClusterId(v string) *CreateMultimodalSearchTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalSearchTaskRequest) SetDatasetIds(v []*string) *CreateMultimodalSearchTaskRequest {
	s.DatasetIds = v
	return s
}

func (s *CreateMultimodalSearchTaskRequest) SetEmbeddingModel(v string) *CreateMultimodalSearchTaskRequest {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateMultimodalSearchTaskRequest) SetQuery(v string) *CreateMultimodalSearchTaskRequest {
	s.Query = &v
	return s
}

func (s *CreateMultimodalSearchTaskRequest) SetSearchModel(v string) *CreateMultimodalSearchTaskRequest {
	s.SearchModel = &v
	return s
}

func (s *CreateMultimodalSearchTaskRequest) SetTopK(v int32) *CreateMultimodalSearchTaskRequest {
	s.TopK = &v
	return s
}

func (s *CreateMultimodalSearchTaskRequest) Validate() error {
	return dara.Validate(s)
}
