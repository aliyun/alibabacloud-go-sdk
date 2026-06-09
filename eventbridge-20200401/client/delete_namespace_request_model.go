// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *DeleteNamespaceRequest
	GetCatalog() *string
	SetClientToken(v string) *DeleteNamespaceRequest
	GetClientToken() *string
	SetName(v string) *DeleteNamespaceRequest
	GetName() *string
}

type DeleteNamespaceRequest struct {
	// 命名空间所属的数据目录名称。可通过 ListCatalogs 接口获取已有目录列表
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// 用于保证请求幂等性的Token。建议使用 UUID
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 要删除的命名空间名称。删除后不可恢复，命名空间下的所有表会一并删除。需同时指定所属 Catalog。可通过 ListNamespaces 获取
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *DeleteNamespaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *DeleteNamespaceRequest) SetCatalog(v string) *DeleteNamespaceRequest {
	s.Catalog = &v
	return s
}

func (s *DeleteNamespaceRequest) SetClientToken(v string) *DeleteNamespaceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteNamespaceRequest) SetName(v string) *DeleteNamespaceRequest {
	s.Name = &v
	return s
}

func (s *DeleteNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
