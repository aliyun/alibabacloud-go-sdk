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
	// The destination entity ID. For more information, see the table ID or field ID in the response returned by the ListTables or ListColumns operation. You can also specify a custom entity ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:123456XXX::test_project::test_tbl
	//
	// dlf-table:123456XXX:test_catalog:test_db::test_tbl
	//
	// hms-table:c-abc123xxx::test_db::test_tbl
	//
	// holo-table:h-abc123xxx::test_db:test_schema:test_tbl
	//
	// custom-api:api123
	//
	// custom-table:table456
	DstEntityId *string `json:"DstEntityId,omitempty" xml:"DstEntityId,omitempty"`
	// example:
	//
	// dstName
	DstEntityName *string `json:"DstEntityName,omitempty" xml:"DstEntityName,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Name
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The source entity ID. For more information, see the table ID or field ID in the response returned by the ListTables or ListColumns operation. You can also specify a custom entity ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:123456XXX::test_project::test_tbl
	//
	// dlf-table:123456XXX:test_catalog:test_db::test_tbl
	//
	// hms-table:c-abc123xxx::test_db::test_tbl
	//
	// holo-table:h-abc123xxx::test_db:test_schema:test_tbl
	//
	// custom-api:api123
	//
	// custom-table:table456
	SrcEntityId *string `json:"SrcEntityId,omitempty" xml:"SrcEntityId,omitempty"`
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
