// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineageRelationshipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstEntityId(v string) *ListLineageRelationshipsRequest
	GetDstEntityId() *string
	SetDstEntityName(v string) *ListLineageRelationshipsRequest
	GetDstEntityName() *string
	SetOrder(v string) *ListLineageRelationshipsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListLineageRelationshipsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLineageRelationshipsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListLineageRelationshipsRequest
	GetSortBy() *string
	SetSrcEntityId(v string) *ListLineageRelationshipsRequest
	GetSrcEntityId() *string
	SetSrcEntityName(v string) *ListLineageRelationshipsRequest
	GetSrcEntityName() *string
}

type ListLineageRelationshipsRequest struct {
	// The destination entity ID. You can use the table or field ID returned by the ListTables or ListColumns operation, or use a custom entity ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlf-table::catalog_id:database_name::table_name
	DstEntityId *string `json:"DstEntityId,omitempty" xml:"DstEntityId,omitempty"`
	// The destination entity name. Fuzzy match is supported.
	//
	// example:
	//
	// dstName
	DstEntityName *string `json:"DstEntityName,omitempty" xml:"DstEntityName,omitempty"`
	// The sort order. Default value: Asc. Valid values:
	//
	// - Asc: ascending order.
	//
	// - Desc: descending order.
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort field. Default value: Name.
	//
	// example:
	//
	// Name
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The source entity ID. You can use the table or field ID returned by the ListTables or ListColumns operation, or use a custom entity ID.
	//
	// To obtain the table or field entity ID, first call ListCrawlers to obtain the MetaEntityId of the metadata crawler. For types that contain data catalog levels, such as DLF and StarRocks, call ListCatalogs to obtain the catalog ID. Then call ListDatabases to obtain the database ID. If necessary, call ListSchemas to obtain the schema ID. Finally, call ListTables or ListColumns to obtain the table or field ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:::project_name:[schema_name]:table_name
	SrcEntityId *string `json:"SrcEntityId,omitempty" xml:"SrcEntityId,omitempty"`
	// The source entity name. Fuzzy match is supported.
	//
	// example:
	//
	// srcName
	SrcEntityName *string `json:"SrcEntityName,omitempty" xml:"SrcEntityName,omitempty"`
}

func (s ListLineageRelationshipsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLineageRelationshipsRequest) GoString() string {
	return s.String()
}

func (s *ListLineageRelationshipsRequest) GetDstEntityId() *string {
	return s.DstEntityId
}

func (s *ListLineageRelationshipsRequest) GetDstEntityName() *string {
	return s.DstEntityName
}

func (s *ListLineageRelationshipsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListLineageRelationshipsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLineageRelationshipsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLineageRelationshipsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLineageRelationshipsRequest) GetSrcEntityId() *string {
	return s.SrcEntityId
}

func (s *ListLineageRelationshipsRequest) GetSrcEntityName() *string {
	return s.SrcEntityName
}

func (s *ListLineageRelationshipsRequest) SetDstEntityId(v string) *ListLineageRelationshipsRequest {
	s.DstEntityId = &v
	return s
}

func (s *ListLineageRelationshipsRequest) SetDstEntityName(v string) *ListLineageRelationshipsRequest {
	s.DstEntityName = &v
	return s
}

func (s *ListLineageRelationshipsRequest) SetOrder(v string) *ListLineageRelationshipsRequest {
	s.Order = &v
	return s
}

func (s *ListLineageRelationshipsRequest) SetPageNumber(v int32) *ListLineageRelationshipsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLineageRelationshipsRequest) SetPageSize(v int32) *ListLineageRelationshipsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLineageRelationshipsRequest) SetSortBy(v string) *ListLineageRelationshipsRequest {
	s.SortBy = &v
	return s
}

func (s *ListLineageRelationshipsRequest) SetSrcEntityId(v string) *ListLineageRelationshipsRequest {
	s.SrcEntityId = &v
	return s
}

func (s *ListLineageRelationshipsRequest) SetSrcEntityName(v string) *ListLineageRelationshipsRequest {
	s.SrcEntityName = &v
	return s
}

func (s *ListLineageRelationshipsRequest) Validate() error {
	return dara.Validate(s)
}
