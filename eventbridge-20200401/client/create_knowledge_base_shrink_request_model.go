// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *CreateKnowledgeBaseShrinkRequest
	GetCatalog() *string
	SetChunkConfigurationShrink(v string) *CreateKnowledgeBaseShrinkRequest
	GetChunkConfigurationShrink() *string
	SetDescription(v string) *CreateKnowledgeBaseShrinkRequest
	GetDescription() *string
	SetEmbeddingDimension(v int32) *CreateKnowledgeBaseShrinkRequest
	GetEmbeddingDimension() *int32
	SetEmbeddingModel(v string) *CreateKnowledgeBaseShrinkRequest
	GetEmbeddingModel() *string
	SetKnowledgeBaseName(v string) *CreateKnowledgeBaseShrinkRequest
	GetKnowledgeBaseName() *string
	SetMetadataSchemaShrink(v string) *CreateKnowledgeBaseShrinkRequest
	GetMetadataSchemaShrink() *string
	SetNamespace(v string) *CreateKnowledgeBaseShrinkRequest
	GetNamespace() *string
	SetSearchConfigurationShrink(v string) *CreateKnowledgeBaseShrinkRequest
	GetSearchConfigurationShrink() *string
}

type CreateKnowledgeBaseShrinkRequest struct {
	// The EventHouse catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies the knowledge base. This parameter cannot be modified after the knowledge base is created. System catalogs cannot be bound.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Optional. The default chunking strategy for the knowledge base. This strategy applies only to documents uploaded after the configuration is set. If this parameter is not specified, the system default chunking strategy is used.
	ChunkConfigurationShrink *string `json:"ChunkConfiguration,omitempty" xml:"ChunkConfiguration,omitempty"`
	// The description of the knowledge base.
	//
	// example:
	//
	// Product documentation knowledge base
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Optional. The vector dimensions. The value is validated against the embedding model: text-embedding-v3 supports 64, 128, 256, 512, 768, and 1024. text-embedding-v4 supports 64, 128, 256, 512, 768, 1024, 1536, and 2048. qwen3.7-text-embedding supports 256, 512, 768, 1024, 1536, 2048, and 2560. qwen3.7-text-embedding-flash supports 256, 512, 768, and 1024. Default value: 1024 (the default dimension of the model). This parameter cannot be modified after the knowledge base is created. Even if the dimensions are the same, you must rebuild the knowledge base when switching models.
	//
	// example:
	//
	// 1024
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// Optional. The embedding model used for vectorization. This parameter cannot be modified after the knowledge base is created. Valid values: text-embedding-v3, text-embedding-v4, qwen3.7-text-embedding, and qwen3.7-text-embedding-flash. Only Bailian Tongyi models are supported. Third-party models are not supported. Default value: text-embedding-v4.
	//
	// example:
	//
	// text-embedding-v4
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// The name of the knowledge base. The name must be unique within the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// product-docs
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// Declares the metadata fields of the knowledge base. When you upload documents, the metadata can contain only the fields declared here. This parameter cannot be modified after the knowledge base is created.
	//
	// example:
	//
	// [{"Name":"department","Type":"STRING"}]
	MetadataSchemaShrink *string `json:"MetadataSchema,omitempty" xml:"MetadataSchema,omitempty"`
	// The EventHouse namespace to which the knowledge base belongs. The namespace must belong to the specified catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies the knowledge base. This parameter cannot be modified after the knowledge base is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Optional. The default search configuration at the knowledge base level. This configuration takes effect when the corresponding parameters are not specified in search requests. You can modify this configuration after creation by calling UpdateKnowledgeBase.
	SearchConfigurationShrink *string `json:"SearchConfiguration,omitempty" xml:"SearchConfiguration,omitempty"`
}

func (s CreateKnowledgeBaseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseShrinkRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateKnowledgeBaseShrinkRequest) GetChunkConfigurationShrink() *string {
	return s.ChunkConfigurationShrink
}

func (s *CreateKnowledgeBaseShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseShrinkRequest) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *CreateKnowledgeBaseShrinkRequest) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateKnowledgeBaseShrinkRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *CreateKnowledgeBaseShrinkRequest) GetMetadataSchemaShrink() *string {
	return s.MetadataSchemaShrink
}

func (s *CreateKnowledgeBaseShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateKnowledgeBaseShrinkRequest) GetSearchConfigurationShrink() *string {
	return s.SearchConfigurationShrink
}

func (s *CreateKnowledgeBaseShrinkRequest) SetCatalog(v string) *CreateKnowledgeBaseShrinkRequest {
	s.Catalog = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) SetChunkConfigurationShrink(v string) *CreateKnowledgeBaseShrinkRequest {
	s.ChunkConfigurationShrink = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) SetDescription(v string) *CreateKnowledgeBaseShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) SetEmbeddingDimension(v int32) *CreateKnowledgeBaseShrinkRequest {
	s.EmbeddingDimension = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) SetEmbeddingModel(v string) *CreateKnowledgeBaseShrinkRequest {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) SetKnowledgeBaseName(v string) *CreateKnowledgeBaseShrinkRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) SetMetadataSchemaShrink(v string) *CreateKnowledgeBaseShrinkRequest {
	s.MetadataSchemaShrink = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) SetNamespace(v string) *CreateKnowledgeBaseShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) SetSearchConfigurationShrink(v string) *CreateKnowledgeBaseShrinkRequest {
	s.SearchConfigurationShrink = &v
	return s
}

func (s *CreateKnowledgeBaseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
