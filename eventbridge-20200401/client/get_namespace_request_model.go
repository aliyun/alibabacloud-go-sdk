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
	// 要查询的命名空间名称。需同时指定所属 Catalog。可通过 ListNamespaces 获取已有命名空间列表
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
