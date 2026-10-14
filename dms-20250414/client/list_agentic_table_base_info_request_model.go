// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticTableBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseUuid(v string) *ListAgenticTableBaseInfoRequest
	GetDatabaseUuid() *string
	SetMaxResults(v int32) *ListAgenticTableBaseInfoRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgenticTableBaseInfoRequest
	GetNextToken() *string
	SetSearchKey(v string) *ListAgenticTableBaseInfoRequest
	GetSearchKey() *string
	SetSortFieldName(v string) *ListAgenticTableBaseInfoRequest
	GetSortFieldName() *string
	SetSortOrder(v string) *ListAgenticTableBaseInfoRequest
	GetSortOrder() *string
	SetTableType(v string) *ListAgenticTableBaseInfoRequest
	GetTableType() *string
}

type ListAgenticTableBaseInfoRequest struct {
	// The UUID of the database to which the table belongs. You can obtain the UUID by calling the ListAgenticDatabase or SearchAgenticDatabase operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
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
	// The search keyword. The system matches the keyword against table names. The results are sorted by exact match, prefix match, and contains match. Within the same match level, the results are sorted by the field specified by SortFieldName.
	//
	// example:
	//
	// order
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The field by which to sort the results. This parameter is case-sensitive. Valid values:
	//
	// - **name**: Sorts the results by table name. This value is recommended.
	//
	// - **numRows**: Sorts the results by the number of rows.
	//
	// - **storeCapacity**: Sorts the results by the storage capacity.
	//
	// If you specify other values, the request fails. If you do not specify this parameter, the results are returned in the default order, and the SortOrder parameter does not take effect.
	//
	// example:
	//
	// name
	SortFieldName *string `json:"SortFieldName,omitempty" xml:"SortFieldName,omitempty"`
	// The sort order. This parameter takes effect only when SortFieldName is specified. This parameter is not case-sensitive. Valid values:
	//
	// - **ASC**: Sorts the results in ascending order. This is the default value and is recommended.
	//
	// - **DESC**: Sorts the results in descending order.
	//
	// If you specify other values, the request fails.
	//
	// example:
	//
	// ASC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The table type. The results are strictly filtered by this type. Valid values:
	//
	// - **NORMAL**: Regular table.
	//
	// - **VIEW**: View.
	//
	// - **FOREIGN**: Foreign table.
	//
	// - **SYNONYM**: Synonym. This value is supported only for Oracle data sources.
	//
	// example:
	//
	// NORMAL
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ListAgenticTableBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticTableBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *ListAgenticTableBaseInfoRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *ListAgenticTableBaseInfoRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgenticTableBaseInfoRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgenticTableBaseInfoRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAgenticTableBaseInfoRequest) GetSortFieldName() *string {
	return s.SortFieldName
}

func (s *ListAgenticTableBaseInfoRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListAgenticTableBaseInfoRequest) GetTableType() *string {
	return s.TableType
}

func (s *ListAgenticTableBaseInfoRequest) SetDatabaseUuid(v string) *ListAgenticTableBaseInfoRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *ListAgenticTableBaseInfoRequest) SetMaxResults(v int32) *ListAgenticTableBaseInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgenticTableBaseInfoRequest) SetNextToken(v string) *ListAgenticTableBaseInfoRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgenticTableBaseInfoRequest) SetSearchKey(v string) *ListAgenticTableBaseInfoRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAgenticTableBaseInfoRequest) SetSortFieldName(v string) *ListAgenticTableBaseInfoRequest {
	s.SortFieldName = &v
	return s
}

func (s *ListAgenticTableBaseInfoRequest) SetSortOrder(v string) *ListAgenticTableBaseInfoRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAgenticTableBaseInfoRequest) SetTableType(v string) *ListAgenticTableBaseInfoRequest {
	s.TableType = &v
	return s
}

func (s *ListAgenticTableBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
