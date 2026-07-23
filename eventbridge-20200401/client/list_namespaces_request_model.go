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
	// Pagination Token
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
