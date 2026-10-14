// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAgenticCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogType(v string) *SearchAgenticCatalogRequest
	GetCatalogType() *string
	SetCatalogUuids(v string) *SearchAgenticCatalogRequest
	GetCatalogUuids() *string
	SetDbType(v string) *SearchAgenticCatalogRequest
	GetDbType() *string
	SetMaxResults(v int32) *SearchAgenticCatalogRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchAgenticCatalogRequest
	GetNextToken() *string
	SetSearchKey(v string) *SearchAgenticCatalogRequest
	GetSearchKey() *string
}

type SearchAgenticCatalogRequest struct {
	// The catalog type. Filters catalogs by exact type. Valid values:
	//
	// - **JDBC**: relational database catalog.
	//
	// - **DataLake**: data lake catalog.
	//
	// - **Dataset**: knowledge base catalog.
	//
	// example:
	//
	// JDBC
	CatalogType *string `json:"CatalogType,omitempty" xml:"CatalogType,omitempty"`
	// The list of catalog UUIDs. Filters catalogs exactly by UUID. Separate multiple UUIDs with commas (,). You can specify up to 500 UUIDs, and null values are not allowed. If this parameter is specified along with other filter conditions, the conditions are combined by using the AND logical operator.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***,mc-HZ-Bq8Lx3wT***
	CatalogUuids *string `json:"CatalogUuids,omitempty" xml:"CatalogUuids,omitempty"`
	// The database type. Filters catalogs exactly by the database type of the associated data source. The value is the uppercase name of the database type, such as MYSQL.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. This parameter is not required for the first query. To query the next page, set this parameter to the NextToken value returned in the previous response.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search keyword. It matches the catalog name, data source alias, or the prefix of the data source endpoint. A catalog is matched if any of these conditions are met.
	//
	// example:
	//
	// prod_
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s SearchAgenticCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAgenticCatalogRequest) GoString() string {
	return s.String()
}

func (s *SearchAgenticCatalogRequest) GetCatalogType() *string {
	return s.CatalogType
}

func (s *SearchAgenticCatalogRequest) GetCatalogUuids() *string {
	return s.CatalogUuids
}

func (s *SearchAgenticCatalogRequest) GetDbType() *string {
	return s.DbType
}

func (s *SearchAgenticCatalogRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchAgenticCatalogRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchAgenticCatalogRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *SearchAgenticCatalogRequest) SetCatalogType(v string) *SearchAgenticCatalogRequest {
	s.CatalogType = &v
	return s
}

func (s *SearchAgenticCatalogRequest) SetCatalogUuids(v string) *SearchAgenticCatalogRequest {
	s.CatalogUuids = &v
	return s
}

func (s *SearchAgenticCatalogRequest) SetDbType(v string) *SearchAgenticCatalogRequest {
	s.DbType = &v
	return s
}

func (s *SearchAgenticCatalogRequest) SetMaxResults(v int32) *SearchAgenticCatalogRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchAgenticCatalogRequest) SetNextToken(v string) *SearchAgenticCatalogRequest {
	s.NextToken = &v
	return s
}

func (s *SearchAgenticCatalogRequest) SetSearchKey(v string) *SearchAgenticCatalogRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchAgenticCatalogRequest) Validate() error {
	return dara.Validate(s)
}
