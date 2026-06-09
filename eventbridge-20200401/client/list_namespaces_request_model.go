// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListNamespacesRequest
	GetCatalog() *string
	SetLimit(v int32) *ListNamespacesRequest
	GetLimit() *int32
	SetNextToken(v string) *ListNamespacesRequest
	GetNextToken() *string
}

type ListNamespacesRequest struct {
	// 要查询的数据目录名称。可通过 ListCatalogs 接口获取
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
	// 分页查询的起始Token。首次查询不传或传 "0"；后续翻页使用上一次响应中返回的 NextToken 值
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListNamespacesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListNamespacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNamespacesRequest) SetCatalog(v string) *ListNamespacesRequest {
	s.Catalog = &v
	return s
}

func (s *ListNamespacesRequest) SetLimit(v int32) *ListNamespacesRequest {
	s.Limit = &v
	return s
}

func (s *ListNamespacesRequest) SetNextToken(v string) *ListNamespacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
