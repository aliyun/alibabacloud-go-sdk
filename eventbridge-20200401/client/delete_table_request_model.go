// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *DeleteTableRequest
	GetCatalog() *string
	SetClientToken(v string) *DeleteTableRequest
	GetClientToken() *string
	SetName(v string) *DeleteTableRequest
	GetName() *string
	SetNamespace(v string) *DeleteTableRequest
	GetNamespace() *string
}

type DeleteTableRequest struct {
	// Data catalog the table belongs to
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Idempotent token
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Table name
	//
	// This parameter is required.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Namespace the table belongs to
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *DeleteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTableRequest) GetName() *string {
	return s.Name
}

func (s *DeleteTableRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteTableRequest) SetCatalog(v string) *DeleteTableRequest {
	s.Catalog = &v
	return s
}

func (s *DeleteTableRequest) SetClientToken(v string) *DeleteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTableRequest) SetName(v string) *DeleteTableRequest {
	s.Name = &v
	return s
}

func (s *DeleteTableRequest) SetNamespace(v string) *DeleteTableRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteTableRequest) Validate() error {
	return dara.Validate(s)
}
