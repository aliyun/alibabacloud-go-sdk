// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *UpdateKnowledgeBaseShrinkRequest
	GetCatalog() *string
	SetChunkConfigurationShrink(v string) *UpdateKnowledgeBaseShrinkRequest
	GetChunkConfigurationShrink() *string
	SetDescription(v string) *UpdateKnowledgeBaseShrinkRequest
	GetDescription() *string
	SetKnowledgeBaseName(v string) *UpdateKnowledgeBaseShrinkRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *UpdateKnowledgeBaseShrinkRequest
	GetNamespace() *string
	SetSearchConfigurationShrink(v string) *UpdateKnowledgeBaseShrinkRequest
	GetSearchConfigurationShrink() *string
}

type UpdateKnowledgeBaseShrinkRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Optional. Updates the default chunking strategy of the knowledge base. The update takes effect only for documents uploaded after the update. Existing documents are not re-chunked. If this parameter is not specified, the configuration remains unchanged.
	ChunkConfigurationShrink *string `json:"ChunkConfiguration,omitempty" xml:"ChunkConfiguration,omitempty"`
	// The description of the knowledge base to update. If this parameter is not specified, the description remains unchanged.
	//
	// example:
	//
	// Product documentation knowledge base
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the knowledge base. The name must be unique within the namespace. The name is specified during creation and cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListNamespaces to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Optional. Updates the default search configuration at the knowledge base level. The update takes effect immediately for subsequent search requests. If this parameter is not specified, the configuration remains unchanged.
	SearchConfigurationShrink *string `json:"SearchConfiguration,omitempty" xml:"SearchConfiguration,omitempty"`
}

func (s UpdateKnowledgeBaseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseShrinkRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *UpdateKnowledgeBaseShrinkRequest) GetChunkConfigurationShrink() *string {
	return s.ChunkConfigurationShrink
}

func (s *UpdateKnowledgeBaseShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseShrinkRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *UpdateKnowledgeBaseShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateKnowledgeBaseShrinkRequest) GetSearchConfigurationShrink() *string {
	return s.SearchConfigurationShrink
}

func (s *UpdateKnowledgeBaseShrinkRequest) SetCatalog(v string) *UpdateKnowledgeBaseShrinkRequest {
	s.Catalog = &v
	return s
}

func (s *UpdateKnowledgeBaseShrinkRequest) SetChunkConfigurationShrink(v string) *UpdateKnowledgeBaseShrinkRequest {
	s.ChunkConfigurationShrink = &v
	return s
}

func (s *UpdateKnowledgeBaseShrinkRequest) SetDescription(v string) *UpdateKnowledgeBaseShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseShrinkRequest) SetKnowledgeBaseName(v string) *UpdateKnowledgeBaseShrinkRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *UpdateKnowledgeBaseShrinkRequest) SetNamespace(v string) *UpdateKnowledgeBaseShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateKnowledgeBaseShrinkRequest) SetSearchConfigurationShrink(v string) *UpdateKnowledgeBaseShrinkRequest {
	s.SearchConfigurationShrink = &v
	return s
}

func (s *UpdateKnowledgeBaseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
