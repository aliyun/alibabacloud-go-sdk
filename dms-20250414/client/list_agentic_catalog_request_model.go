// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogType(v string) *ListAgenticCatalogRequest
	GetCatalogType() *string
	SetMaxResults(v int32) *ListAgenticCatalogRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgenticCatalogRequest
	GetNextToken() *string
	SetSearchKey(v string) *ListAgenticCatalogRequest
	GetSearchKey() *string
}

type ListAgenticCatalogRequest struct {
	// The catalog type. This parameter is used to filter catalogs by exact type. Valid values:
	//
	// - **JDBC**: relational database catalog
	//
	// - **DataLake**: data lake catalog
	//
	// - **Dataset**: knowledge base catalog
	//
	// example:
	//
	// JDBC
	CatalogType *string `json:"CatalogType,omitempty" xml:"CatalogType,omitempty"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. You do not need to specify this parameter for the first query. To query the next page, set this parameter to the NextToken value returned from the previous query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search keyword for the catalog name. Matches catalogs by name prefix.
	//
	// example:
	//
	// prod_
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s ListAgenticCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticCatalogRequest) GoString() string {
	return s.String()
}

func (s *ListAgenticCatalogRequest) GetCatalogType() *string {
	return s.CatalogType
}

func (s *ListAgenticCatalogRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgenticCatalogRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgenticCatalogRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAgenticCatalogRequest) SetCatalogType(v string) *ListAgenticCatalogRequest {
	s.CatalogType = &v
	return s
}

func (s *ListAgenticCatalogRequest) SetMaxResults(v int32) *ListAgenticCatalogRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgenticCatalogRequest) SetNextToken(v string) *ListAgenticCatalogRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgenticCatalogRequest) SetSearchKey(v string) *ListAgenticCatalogRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAgenticCatalogRequest) Validate() error {
	return dara.Validate(s)
}
