// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaDocumentDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetLumaDocumentDownloadUrlRequest
	GetAgentName() *string
	SetCatalog(v string) *GetLumaDocumentDownloadUrlRequest
	GetCatalog() *string
	SetDocumentId(v string) *GetLumaDocumentDownloadUrlRequest
	GetDocumentId() *string
	SetKnowledgeBaseName(v string) *GetLumaDocumentDownloadUrlRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *GetLumaDocumentDownloadUrlRequest
	GetNamespace() *string
	SetNetworkType(v string) *GetLumaDocumentDownloadUrlRequest
	GetNetworkType() *string
}

type GetLumaDocumentDownloadUrlRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the Agent. You can call ListLumaCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The unique identifier of the document. You can call ListLumaDocuments to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc-a1b2c3d4
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// The name of the knowledge base bound to the Agent. You can call ListLumaKnowledgeBases to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The name of the namespace bound to the Agent. You can call ListLumaNamespaces to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Valid values:
	//
	// - public: public network.
	//
	// - vpc: internal network.
	//
	// Default value: public.
	//
	// example:
	//
	// public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s GetLumaDocumentDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLumaDocumentDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetLumaDocumentDownloadUrlRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetLumaDocumentDownloadUrlRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetLumaDocumentDownloadUrlRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetLumaDocumentDownloadUrlRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *GetLumaDocumentDownloadUrlRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLumaDocumentDownloadUrlRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetLumaDocumentDownloadUrlRequest) SetAgentName(v string) *GetLumaDocumentDownloadUrlRequest {
	s.AgentName = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlRequest) SetCatalog(v string) *GetLumaDocumentDownloadUrlRequest {
	s.Catalog = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlRequest) SetDocumentId(v string) *GetLumaDocumentDownloadUrlRequest {
	s.DocumentId = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlRequest) SetKnowledgeBaseName(v string) *GetLumaDocumentDownloadUrlRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlRequest) SetNamespace(v string) *GetLumaDocumentDownloadUrlRequest {
	s.Namespace = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlRequest) SetNetworkType(v string) *GetLumaDocumentDownloadUrlRequest {
	s.NetworkType = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
