// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListTablesRequest
	GetCatalog() *string
	SetLimit(v int32) *ListTablesRequest
	GetLimit() *int32
	SetNamespace(v string) *ListTablesRequest
	GetNamespace() *string
	SetNextToken(v string) *ListTablesRequest
	GetNextToken() *string
}

type ListTablesRequest struct {
	// 表所属的数据目录名称。可通过 ListCatalogs 获取
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// 每页返回的最大数据条数。不传时默认 10，最大 100
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 表所属的命名空间名称。可通过 ListNamespaces 获取
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 分页查询的起始Token。首次查询不传或传 "0"；后续翻页使用上一次响应中返回的 NextToken 值
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListTablesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListTablesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTablesRequest) SetCatalog(v string) *ListTablesRequest {
	s.Catalog = &v
	return s
}

func (s *ListTablesRequest) SetLimit(v int32) *ListTablesRequest {
	s.Limit = &v
	return s
}

func (s *ListTablesRequest) SetNamespace(v string) *ListTablesRequest {
	s.Namespace = &v
	return s
}

func (s *ListTablesRequest) SetNextToken(v string) *ListTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}
