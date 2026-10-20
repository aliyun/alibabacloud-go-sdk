// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDenseSimilarityTopK(v int32) *UpdateIndexRequest
	GetDenseSimilarityTopK() *int32
	SetDescription(v string) *UpdateIndexRequest
	GetDescription() *string
	SetId(v string) *UpdateIndexRequest
	GetId() *string
	SetName(v string) *UpdateIndexRequest
	GetName() *string
	SetPipelineCommercialCu(v int32) *UpdateIndexRequest
	GetPipelineCommercialCu() *int32
	SetPipelineCommercialType(v string) *UpdateIndexRequest
	GetPipelineCommercialType() *string
	SetRerankMinScore(v string) *UpdateIndexRequest
	GetRerankMinScore() *string
	SetSparseSimilarityTopK(v int32) *UpdateIndexRequest
	GetSparseSimilarityTopK() *int32
}

type UpdateIndexRequest struct {
	// The top K value for AISearch. The system generates a vector for the input text and retrieves the K text chunks with the most similar vector representations from the knowledge base. Valid values: 0 to 100.
	//
	// The sum of `DenseSimilarityTopK` and `SparseSimilarityTopK` must be less than or equal to 200.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	DenseSimilarityTopK *int32 `json:"DenseSimilarityTopK,omitempty" xml:"DenseSimilarityTopK,omitempty"`
	// The knowledge base description.
	//
	// example:
	//
	// Enterprise knowledge base
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The knowledge base name.
	//
	// example:
	//
	// Enterprise help documentation library
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of RCUs for an Ultimate Edition knowledge base. This parameter is required only when pipelineCommercialType is set to enterprise.
	//
	// Valid values: 1 to 200.
	//
	// example:
	//
	// 3
	PipelineCommercialCu *int32 `json:"PipelineCommercialCu,omitempty" xml:"PipelineCommercialCu,omitempty"`
	// The specification type of the knowledge base. Valid values:
	//
	// - standard: Standard Edition.
	//
	// - enterprise: Ultimate Edition.
	//
	// example:
	//
	// standard
	PipelineCommercialType *string `json:"PipelineCommercialType,omitempty" xml:"PipelineCommercialType,omitempty"`
	// The minimum reranking score. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.01
	RerankMinScore *string `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The top K value for keyword retrieval. The system retrieves text chunks in the knowledge base that are an exact match for the keywords of the input text. This helps filter out irrelevant text chunks and provides more accurate results.
	//
	// Valid values: 0 to 100.
	//
	// The sum of `DenseSimilarityTopK` and `SparseSimilarityTopK` must be less than or equal to 200.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	SparseSimilarityTopK *int32 `json:"SparseSimilarityTopK,omitempty" xml:"SparseSimilarityTopK,omitempty"`
}

func (s UpdateIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIndexRequest) GoString() string {
	return s.String()
}

func (s *UpdateIndexRequest) GetDenseSimilarityTopK() *int32 {
	return s.DenseSimilarityTopK
}

func (s *UpdateIndexRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateIndexRequest) GetId() *string {
	return s.Id
}

func (s *UpdateIndexRequest) GetName() *string {
	return s.Name
}

func (s *UpdateIndexRequest) GetPipelineCommercialCu() *int32 {
	return s.PipelineCommercialCu
}

func (s *UpdateIndexRequest) GetPipelineCommercialType() *string {
	return s.PipelineCommercialType
}

func (s *UpdateIndexRequest) GetRerankMinScore() *string {
	return s.RerankMinScore
}

func (s *UpdateIndexRequest) GetSparseSimilarityTopK() *int32 {
	return s.SparseSimilarityTopK
}

func (s *UpdateIndexRequest) SetDenseSimilarityTopK(v int32) *UpdateIndexRequest {
	s.DenseSimilarityTopK = &v
	return s
}

func (s *UpdateIndexRequest) SetDescription(v string) *UpdateIndexRequest {
	s.Description = &v
	return s
}

func (s *UpdateIndexRequest) SetId(v string) *UpdateIndexRequest {
	s.Id = &v
	return s
}

func (s *UpdateIndexRequest) SetName(v string) *UpdateIndexRequest {
	s.Name = &v
	return s
}

func (s *UpdateIndexRequest) SetPipelineCommercialCu(v int32) *UpdateIndexRequest {
	s.PipelineCommercialCu = &v
	return s
}

func (s *UpdateIndexRequest) SetPipelineCommercialType(v string) *UpdateIndexRequest {
	s.PipelineCommercialType = &v
	return s
}

func (s *UpdateIndexRequest) SetRerankMinScore(v string) *UpdateIndexRequest {
	s.RerankMinScore = &v
	return s
}

func (s *UpdateIndexRequest) SetSparseSimilarityTopK(v int32) *UpdateIndexRequest {
	s.SparseSimilarityTopK = &v
	return s
}

func (s *UpdateIndexRequest) Validate() error {
	return dara.Validate(s)
}
