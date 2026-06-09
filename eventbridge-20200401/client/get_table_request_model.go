// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *GetTableRequest
	GetCatalog() *string
	SetClientToken(v string) *GetTableRequest
	GetClientToken() *string
	SetName(v string) *GetTableRequest
	GetName() *string
	SetNamespace(v string) *GetTableRequest
	GetNamespace() *string
}

type GetTableRequest struct {
	// 表所属的数据目录名称。可通过 ListCatalogs 获取
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
	// 要查询的事件表名称。需同时指定所属 Catalog 和 Namespace。可通过 ListTables 获取已有表列表
	//
	// This parameter is required.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 表所属的命名空间名称。可通过 ListNamespaces 获取
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableRequest) GoString() string {
	return s.String()
}

func (s *GetTableRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetTableRequest) GetName() *string {
	return s.Name
}

func (s *GetTableRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetTableRequest) SetCatalog(v string) *GetTableRequest {
	s.Catalog = &v
	return s
}

func (s *GetTableRequest) SetClientToken(v string) *GetTableRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTableRequest) SetName(v string) *GetTableRequest {
	s.Name = &v
	return s
}

func (s *GetTableRequest) SetNamespace(v string) *GetTableRequest {
	s.Namespace = &v
	return s
}

func (s *GetTableRequest) Validate() error {
	return dara.Validate(s)
}
