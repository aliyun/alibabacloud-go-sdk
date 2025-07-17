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
	// example:
	//
	// ds=20250101
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The ID of the table to which the partitions belong. You can call the ListTables operation to query the ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-column:11075xxxx::test_project:test_schema:test_table
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
