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
	// Data catalog
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Items per page
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// Namespace
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Pagination token
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
