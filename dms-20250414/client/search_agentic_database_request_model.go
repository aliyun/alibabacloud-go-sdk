// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAgenticDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *SearchAgenticDatabaseRequest
	GetCatalogUuid() *string
	SetDbType(v string) *SearchAgenticDatabaseRequest
	GetDbType() *string
	SetEnvType(v string) *SearchAgenticDatabaseRequest
	GetEnvType() *string
	SetMaxResults(v int32) *SearchAgenticDatabaseRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchAgenticDatabaseRequest
	GetNextToken() *string
	SetSearchKey(v string) *SearchAgenticDatabaseRequest
	GetSearchKey() *string
}

type SearchAgenticDatabaseRequest struct {
	// The UUID of the catalog that limits the search scope. If you do not specify this parameter, the search is performed across all catalogs.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The database type. This parameter filters the results exactly by the database type of the data source associated with the catalog. The value is the uppercase name of the database type, such as MYSQL.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The environment type. This parameter filters the results exactly by the environment of the associated data source. Valid values: product, pre, dev, test, sit, uat, pet, and stag.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. You do not need to specify this parameter for the first query. When you query the next page, set this parameter to the NextToken value returned from the previous query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search keyword. It matches the name, fully qualified name, and search name of the database, as well as the name or data source alias of the associated catalog.
	//
	// example:
	//
	// order
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s SearchAgenticDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAgenticDatabaseRequest) GoString() string {
	return s.String()
}

func (s *SearchAgenticDatabaseRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *SearchAgenticDatabaseRequest) GetDbType() *string {
	return s.DbType
}

func (s *SearchAgenticDatabaseRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *SearchAgenticDatabaseRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchAgenticDatabaseRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchAgenticDatabaseRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *SearchAgenticDatabaseRequest) SetCatalogUuid(v string) *SearchAgenticDatabaseRequest {
	s.CatalogUuid = &v
	return s
}

func (s *SearchAgenticDatabaseRequest) SetDbType(v string) *SearchAgenticDatabaseRequest {
	s.DbType = &v
	return s
}

func (s *SearchAgenticDatabaseRequest) SetEnvType(v string) *SearchAgenticDatabaseRequest {
	s.EnvType = &v
	return s
}

func (s *SearchAgenticDatabaseRequest) SetMaxResults(v int32) *SearchAgenticDatabaseRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchAgenticDatabaseRequest) SetNextToken(v string) *SearchAgenticDatabaseRequest {
	s.NextToken = &v
	return s
}

func (s *SearchAgenticDatabaseRequest) SetSearchKey(v string) *SearchAgenticDatabaseRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchAgenticDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
