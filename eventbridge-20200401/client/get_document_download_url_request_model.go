// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *GetDocumentDownloadUrlRequest
	GetCatalog() *string
	SetDocumentId(v string) *GetDocumentDownloadUrlRequest
	GetDocumentId() *string
	SetKnowledgeBaseName(v string) *GetDocumentDownloadUrlRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *GetDocumentDownloadUrlRequest
	GetNamespace() *string
	SetNetworkType(v string) *GetDocumentDownloadUrlRequest
	GetNetworkType() *string
}

type GetDocumentDownloadUrlRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The document ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc-bp1xxxxxxxxxxxx
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// The name of the knowledge base. The name is unique within the namespace and is determined at creation time. It cannot be modified.
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
	// The network type of the pre-signed URL. Valid values:
	//
	// - PUBLIC (default): accessible over the Internet.
	//
	// - INTERNAL: accessible only within the same region on the Alibaba Cloud internal network, with no Internet data transfer fees.
	//
	// example:
	//
	// PUBLIC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s GetDocumentDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetDocumentDownloadUrlRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetDocumentDownloadUrlRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *GetDocumentDownloadUrlRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetDocumentDownloadUrlRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetDocumentDownloadUrlRequest) SetCatalog(v string) *GetDocumentDownloadUrlRequest {
	s.Catalog = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) SetDocumentId(v string) *GetDocumentDownloadUrlRequest {
	s.DocumentId = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) SetKnowledgeBaseName(v string) *GetDocumentDownloadUrlRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) SetNamespace(v string) *GetDocumentDownloadUrlRequest {
	s.Namespace = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) SetNetworkType(v string) *GetDocumentDownloadUrlRequest {
	s.NetworkType = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
