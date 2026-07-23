// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *CreateNamespaceRequest
	GetCatalog() *string
	SetClientToken(v string) *CreateNamespaceRequest
	GetClientToken() *string
	SetComment(v string) *CreateNamespaceRequest
	GetComment() *string
	SetName(v string) *CreateNamespaceRequest
	GetName() *string
}

type CreateNamespaceRequest struct {
	// Data catalog
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Idempotent Token
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Remarks
	//
	// example:
	//
	// 测试命名空间
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Namespace name
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateNamespaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNamespaceRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateNamespaceRequest) SetCatalog(v string) *CreateNamespaceRequest {
	s.Catalog = &v
	return s
}

func (s *CreateNamespaceRequest) SetClientToken(v string) *CreateNamespaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNamespaceRequest) SetComment(v string) *CreateNamespaceRequest {
	s.Comment = &v
	return s
}

func (s *CreateNamespaceRequest) SetName(v string) *CreateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
