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
	SetRecentDays(v int32) *ListLineagesRequest
	GetRecentDays() *int32
	SetSortBy(v string) *ListLineagesRequest
	GetSortBy() *string
	SetSrcEntityId(v string) *ListLineagesRequest
	GetSrcEntityId() *string
	SetSrcEntityName(v string) *ListLineagesRequest
	GetSrcEntityName() *string
}

type ListLineagesRequest struct {
	// The destination entity ID. You can use the table or column ID returned by the ListTables or ListColumns operation, or a custom entity ID.
	//
	// example:
	//
	// dlf-table::catalog_id:database_name::table_name
	DstEntityId *string `json:"DstEntityId,omitempty" xml:"DstEntityId,omitempty"`
	// The destination entity name. Fuzzy match is supported.
	//
	// example:
	//
	// dstName1
	DstEntityName *string `json:"DstEntityName,omitempty" xml:"DstEntityName,omitempty"`
	// Specifies whether to return lineage relationship information. Default value: false.
	//
	// example:
	//
	// false
	NeedAttachRelationship *bool `json:"NeedAttachRelationship,omitempty" xml:"NeedAttachRelationship,omitempty"`
	// The sort direction. Default value: Asc. Valid values:
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
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RecentDays *int32 `json:"RecentDays,omitempty" xml:"RecentDays,omitempty"`
	// The sort field. Default value: Name, which sorts by lineage entity name.
	//
	// example:
	//
	// Name
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The source entity ID. You can use the table or column ID returned by the ListTables or ListColumns operation, or a custom entity ID.
	//
	// example:
	//
	// maxcompute-table:::project_name:[schema_name]:table_name
	SrcEntityId *string `json:"SrcEntityId,omitempty" xml:"SrcEntityId,omitempty"`
	// The source entity name. Fuzzy match is supported.
	//
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

func (s *ListLineagesRequest) GetRecentDays() *int32 {
	return s.RecentDays
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

func (s *ListLineagesRequest) SetRecentDays(v int32) *ListLineagesRequest {
	s.RecentDays = &v
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
