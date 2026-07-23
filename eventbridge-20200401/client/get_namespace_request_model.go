// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *GetNamespaceRequest
	GetCatalog() *string
	SetClientToken(v string) *GetNamespaceRequest
	GetClientToken() *string
	SetName(v string) *GetNamespaceRequest
	GetName() *string
}

type GetNamespaceRequest struct {
	// Data catalog
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Idempotency token
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Namespace name
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetNamespaceRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetNamespaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *GetNamespaceRequest) SetCatalog(v string) *GetNamespaceRequest {
	s.Catalog = &v
	return s
}

func (s *GetNamespaceRequest) SetClientToken(v string) *GetNamespaceRequest {
	s.ClientToken = &v
	return s
}

func (s *GetNamespaceRequest) SetName(v string) *GetNamespaceRequest {
	s.Name = &v
	return s
}

func (s *GetNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
