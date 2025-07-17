// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstEntityId(v string) *ListLineagesRequest
	GetDstEntityId() *string
	SetDstEntityName(v string) *ListLineagesRequest
	GetDstEntityName() *string
	SetNeedAttachRelationship(v bool) *ListLineagesRequest
	GetNeedAttachRelationship() *bool
	SetOrder(v string) *ListLineagesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListLineagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLineagesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListLineagesRequest
	GetSortBy() *string
	SetSrcEntityId(v string) *ListLineagesRequest
	GetSrcEntityId() *string
	SetSrcEntityName(v string) *ListLineagesRequest
	GetSrcEntityName() *string
}

type ListLineagesRequest struct {
	// The destination entity ID. For more information, see the table ID or field ID in the response returned by the ListTables or ListColumns operation. You can also specify a custom entity ID.
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
	// dstName1
	DstEntityName *string `json:"DstEntityName,omitempty" xml:"DstEntityName,omitempty"`
	// example:
	//
	// false
	NeedAttachRelationship *bool `json:"NeedAttachRelationship,omitempty" xml:"NeedAttachRelationship,omitempty"`
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
	// srcName1
	SrcEntityName *string `json:"SrcEntityName,omitempty" xml:"SrcEntityName,omitempty"`
}

func (s ListLineagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLineagesRequest) GoString() string {
	return s.String()
}

func (s *ListLineagesRequest) GetDstEntityId() *string {
	return s.DstEntityId
}

func (s *ListLineagesRequest) GetDstEntityName() *string {
	return s.DstEntityName
}

func (s *ListLineagesRequest) GetNeedAttachRelationship() *bool {
	return s.NeedAttachRelationship
}

func (s *ListLineagesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListLineagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLineagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLineagesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLineagesRequest) GetSrcEntityId() *string {
	return s.SrcEntityId
}

func (s *ListLineagesRequest) GetSrcEntityName() *string {
	return s.SrcEntityName
}

func (s *ListLineagesRequest) SetDstEntityId(v string) *ListLineagesRequest {
	s.DstEntityId = &v
	return s
}

func (s *ListLineagesRequest) SetDstEntityName(v string) *ListLineagesRequest {
	s.DstEntityName = &v
	return s
}

func (s *ListLineagesRequest) SetNeedAttachRelationship(v bool) *ListLineagesRequest {
	s.NeedAttachRelationship = &v
	return s
}

func (s *ListLineagesRequest) SetOrder(v string) *ListLineagesRequest {
	s.Order = &v
	return s
}

func (s *ListLineagesRequest) SetPageNumber(v int32) *ListLineagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLineagesRequest) SetPageSize(v int32) *ListLineagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLineagesRequest) SetSortBy(v string) *ListLineagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListLineagesRequest) SetSrcEntityId(v string) *ListLineagesRequest {
	s.SrcEntityId = &v
	return s
}

func (s *ListLineagesRequest) SetSrcEntityName(v string) *ListLineagesRequest {
	s.SrcEntityName = &v
	return s
}

func (s *ListLineagesRequest) Validate() error {
	return dara.Validate(s)
}
