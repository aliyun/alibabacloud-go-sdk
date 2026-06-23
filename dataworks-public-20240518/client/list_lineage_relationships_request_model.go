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
	// The ID of the destination entity. You can get the ID for a table or column from the response of the `ListTables` or `ListColumns` operation, or specify the ID of a custom entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlf-table::catalog_id:database_name::table_name
	DstEntityId *string `json:"DstEntityId,omitempty" xml:"DstEntityId,omitempty"`
	// The name of the destination entity. Supports fuzzy matching.
	//
	// example:
	//
	// dstName
	DstEntityName *string `json:"DstEntityName,omitempty" xml:"DstEntityName,omitempty"`
	// The sort order. The default value is `Asc`. Valid values:
	//
	// - `Asc`: ascending order
	//
	// - `Desc`: descending order
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. The default value is 10. The maximum value is 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field to sort the results by. The default value is `Name`.
	//
	// example:
	//
	// Name
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The ID of the source entity. You can get the ID for a table or column from the response of the `ListTables` or `ListColumns` operation, or specify the ID of a custom entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:::project_name:[schema_name]:table_name
	SrcEntityId *string `json:"SrcEntityId,omitempty" xml:"SrcEntityId,omitempty"`
	// The name of the source entity. Supports fuzzy matching.
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
