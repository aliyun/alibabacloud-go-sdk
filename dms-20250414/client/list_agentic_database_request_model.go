// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *ListAgenticDatabaseRequest
	GetCatalogUuid() *string
	SetMaxResults(v int32) *ListAgenticDatabaseRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgenticDatabaseRequest
	GetNextToken() *string
	SetSearchKey(v string) *ListAgenticDatabaseRequest
	GetSearchKey() *string
	SetSortFieldName(v string) *ListAgenticDatabaseRequest
	GetSortFieldName() *string
	SetSortOrder(v string) *ListAgenticDatabaseRequest
	GetSortOrder() *string
}

type ListAgenticDatabaseRequest struct {
	// The UUID of the catalog to which the database belongs. You can obtain the UUID by calling ListAgenticCatalog or SearchAgenticCatalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. You do not need to specify this parameter for the first query. When you query the next page, set this parameter to the NextToken value returned in the previous response.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search keyword used to match the database name or the prefix of the fully qualified name.
	//
	// example:
	//
	// order
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The field used for sorting. This parameter is case-sensitive. Valid values:
	//
	// - **name**: Sorts by database name.
	//
	// - **qualifiedName**: Sorts by the fully qualified name of the database.
	//
	// If you specify other values, the request fails. If you do not specify this parameter, the results are returned in the default order, and the SortOrder parameter does not take effect.
	//
	// example:
	//
	// name
	SortFieldName *string `json:"SortFieldName,omitempty" xml:"SortFieldName,omitempty"`
	// The sort order. This parameter takes effect only when SortFieldName is specified. This parameter is case-insensitive. Valid values:
	//
	// - **ASC*	- (default): Ascending order.
	//
	// - **DESC**: Descending order.
	//
	// If you specify other values, the request fails.
	//
	// example:
	//
	// ASC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListAgenticDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ListAgenticDatabaseRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ListAgenticDatabaseRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgenticDatabaseRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgenticDatabaseRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAgenticDatabaseRequest) GetSortFieldName() *string {
	return s.SortFieldName
}

func (s *ListAgenticDatabaseRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListAgenticDatabaseRequest) SetCatalogUuid(v string) *ListAgenticDatabaseRequest {
	s.CatalogUuid = &v
	return s
}

func (s *ListAgenticDatabaseRequest) SetMaxResults(v int32) *ListAgenticDatabaseRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgenticDatabaseRequest) SetNextToken(v string) *ListAgenticDatabaseRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgenticDatabaseRequest) SetSearchKey(v string) *ListAgenticDatabaseRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAgenticDatabaseRequest) SetSortFieldName(v string) *ListAgenticDatabaseRequest {
	s.SortFieldName = &v
	return s
}

func (s *ListAgenticDatabaseRequest) SetSortOrder(v string) *ListAgenticDatabaseRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAgenticDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
