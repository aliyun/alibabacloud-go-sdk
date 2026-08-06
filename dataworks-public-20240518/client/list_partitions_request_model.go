// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListPartitionsRequest
	GetName() *string
	SetOrder(v string) *ListPartitionsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListPartitionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPartitionsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListPartitionsRequest
	GetSortBy() *string
	SetTableId(v string) *ListPartitionsRequest
	GetTableId() *string
}

type ListPartitionsRequest struct {
	// The partition name. Fuzzy match is supported. This parameter is valid only for MaxCompute tables.
	//
	// example:
	//
	// ds=20250101
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort direction. Default value: Asc. Valid values:
	//
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
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which to sort the results. Default value: CreateTime. Valid values:
	//
	//
	// - CreateTime: the creation time. Only MaxCompute tables are supported.
	//
	// - ModifyTime: the modification time. Only MaxCompute tables are supported.
	//
	// - Name: the name. This is the sort method used for HMS tables.
	//
	// - RecordCount: the number of records. Only MaxCompute tables are supported.
	//
	// - DataSize: the storage size. Only MaxCompute tables are supported.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The ID of the table. You can obtain this value from the response of the [ListTables](https://help.aliyun.com/document_detail/2880092.html) operation. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:::project_name:[schema_name]:table_name
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s ListPartitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsRequest) GoString() string {
	return s.String()
}

func (s *ListPartitionsRequest) GetName() *string {
	return s.Name
}

func (s *ListPartitionsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListPartitionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPartitionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPartitionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListPartitionsRequest) GetTableId() *string {
	return s.TableId
}

func (s *ListPartitionsRequest) SetName(v string) *ListPartitionsRequest {
	s.Name = &v
	return s
}

func (s *ListPartitionsRequest) SetOrder(v string) *ListPartitionsRequest {
	s.Order = &v
	return s
}

func (s *ListPartitionsRequest) SetPageNumber(v int32) *ListPartitionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPartitionsRequest) SetPageSize(v int32) *ListPartitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPartitionsRequest) SetSortBy(v string) *ListPartitionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListPartitionsRequest) SetTableId(v string) *ListPartitionsRequest {
	s.TableId = &v
	return s
}

func (s *ListPartitionsRequest) Validate() error {
	return dara.Validate(s)
}
